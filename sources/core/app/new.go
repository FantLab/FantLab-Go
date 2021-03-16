package app

import (
	"context"
	"database/sql"
	"errors"
	"fantlab/base/clients/memcacheclient"
	"fantlab/base/clients/redisclient"
	"fantlab/base/clients/smtpclient"
	"fantlab/base/stdutils"
	"fantlab/base/ttlcache"
	"fantlab/core/config"
	"fantlab/core/db"
	"fantlab/core/logs"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/FantLab/go-kit/codeflow"
	"github.com/FantLab/go-kit/crypto/signed"
	"github.com/FantLab/go-kit/database/sqlapi"
	"github.com/FantLab/go-kit/database/sqldb"
	"github.com/gomodule/redigo/redis"
	"github.com/minio/minio-go/v6"
	"go.elastic.co/apm/module/apmredigo"
	"go.elastic.co/apm/module/apmsql"
	_ "go.elastic.co/apm/module/apmsql/mysql"
	"go.uber.org/zap"
)

func MakeServices() (*Services, error, []func() error) {
	var disposeBag []func() error

	var mysqlDB sqlapi.DB
	var redisClient redisclient.Client
	var memcacheClient memcacheclient.Client
	var smtpClient smtpclient.Client
	var cryptoCoder *signed.Coder
	var minioClient *minio.Client
	var minioBucket string
	var appConfig *config.AppConfig

	err := codeflow.Try(
		func() error { // мускуль
			db, err := apmsql.Open("mysql", os.Getenv("MYSQL_URL"))
			if err == nil {
				err = db.Ping()
			}
			if err != nil {
				return fmt.Errorf("MySQL setup error: %v", err)
			}

			disposeBag = append(disposeBag, db.Close)

			mysqlDB = sqlapi.Log(sqldb.New(db), func(ctx context.Context, entry sqlapi.LogEntry) {
				logger := logs.WithAPM(ctx)
				logger.Info(
					entry.Query(),
					zap.Duration("duration", entry.Duration),
					zap.Int64("rows", entry.Rows),
				)
				if entry.Err != nil && entry.Err != sql.ErrNoRows {
					logger.Error(entry.Err.Error())
				}
			})

			return nil
		},
		func() error { // редис (опционально)
			serverAddr := os.Getenv("RDS_ADDRESS")
			if len(serverAddr) == 0 {
				return nil
			}

			client, close := redisclient.NewPoolClient(serverAddr, 8, func(pool *redis.Pool, ctx context.Context) (redis.Conn, error) {
				return apmredigo.Wrap(pool.Get()).WithContext(ctx), nil
			})
			err := client.Perform(context.Background(), func(conn redisclient.Conn) error {
				_, err := conn.Do("PING")
				return err
			})
			if err != nil {
				return fmt.Errorf("Redis setup error: %v", err)
			}

			disposeBag = append(disposeBag, close)

			redisClient = redisclient.Log(client, func(ctx context.Context, err error) {
				logs.WithAPM(ctx).Error(err.Error())
			})

			return nil
		},
		func() error { // мемкэш (опционально)
			serverAddr := os.Getenv("MC_ADDRESS")
			if len(serverAddr) == 0 {
				return nil
			}

			client := memcacheclient.New(serverAddr)
			err := client.Ping()
			if err != nil {
				return fmt.Errorf("Memcache setup error: %v", err)
			}

			memcacheClient = memcacheclient.APM(memcacheclient.Log(client, func(ctx context.Context, entry *memcacheclient.LogEntry) {
				logger := logs.WithAPM(ctx)
				logger.Info(fmt.Sprintf("%s %s", entry.Operation, entry.Key))
				if entry.Err != nil && !memcacheclient.IsNotFoundError(entry.Err) {
					logger.Error(entry.Err.Error())
				}
			}), "db.memcache")

			return nil
		},
		func() error { // SMTP (опционально)
			smtpAddr := os.Getenv("SMTP_ADDRESS")
			if len(smtpAddr) == 0 {
				return nil
			}

			client, close, err := smtpclient.New(smtpAddr)
			if err == nil {
				err = client.Ping()
			}
			if err != nil {
				return fmt.Errorf("SMTP setup error: %v", err)
			}

			disposeBag = append(disposeBag, close)

			smtpClient = smtpclient.APM(smtpclient.Log(client, func(ctx context.Context, entry *smtpclient.LogEntry) {
				logger := logs.WithAPM(ctx)
				logger.Info(fmt.Sprintf("send email to: %s; subject: %s", entry.To, entry.Subject))
				if entry.Err != nil {
					logger.Error(entry.Err.Error())
				}
			}), "smtp")

			return nil
		},
		func() error { // криптокодер для jwt-like токенов
			coder, err := signed.NewFileCoder64(os.Getenv("JWT_PUBLIC_KEY_FILE"), os.Getenv("JWT_PRIVATE_KEY_FILE"))
			if err != nil {
				coder, err = signed.NewCoder64(os.Getenv("JWT_PUB_KEY"), os.Getenv("JWT_PRIV_KEY"))
				if err != nil {
					return fmt.Errorf("JWT setup error: %v", err)
				}
			}
			cryptoCoder = coder
			return nil
		},
		func() error { // minio
			server := os.Getenv("MINIO_SERVER")
			if len(server) == 0 {
				return errors.New("Min.io setup error: server not found")
			}

			client, err := makeMinioWithSecrets(server, os.Getenv("MINIO_ACCESS_KEY_FILE"), os.Getenv("MINIO_SECRET_KEY_FILE"))
			if err != nil {
				client, err = minio.New(server, os.Getenv("MINIO_ACCESS_KEY"), os.Getenv("MINIO_SECRET_KEY"), false)
				if err != nil {
					return fmt.Errorf("Min.io setup error: %v", err)
				}
			}

			minioBucket = os.Getenv("MINIO_BUCKET")
			bucketExists, err := client.BucketExists(minioBucket)
			if err != nil {
				return fmt.Errorf("Min.io setup error: %v", err)
			}
			if !bucketExists {
				return fmt.Errorf("Min.io setup error: bucket %s not found", minioBucket)
			}

			minioClient = client

			return nil
		},
		func() error { // конфигурация бизнес-логики
			// Все параметры заданы в config/main.cfg и config/misc.cfg Perl-бэка
			appConfig = &config.AppConfig{
				SiteUrl:                                  "https://fantlab.ru",
				SiteName:                                 "fantlab.ru",
				SiteEmail:                                "support@fantlab.ru",
				BaseImageUrl:                             stdutils.Elvis(os.Getenv("BASE_IMAGE_URL"), "https://data.fantlab.ru/images"),
				BaseMinioFileUrl:                         stdutils.Elvis(os.Getenv("BASE_MINIO_FILE_URL"), "https://fantlab.ru/files"),
				BaseForumMessageAttachUrl:                stdutils.Elvis(os.Getenv("BASE_FORUM_MESSAGE_ATTACH_URL"), "https://fantlab.ru/messages"),
				BaseForumMessageDraftAttachUrl:           stdutils.Elvis(os.Getenv("BASE_FORUM_MESSAGE_DRAFT_ATTACH_URL"), "https://fantlab.ru/files/preview"),
				MinResponseLength:                        50,
				MaxUserResponseCountPerWork:              3,
				MinUserOwnResponsesRatingForMinusAbility: 300,
				ForumTopicsInPage:                        20,
				ForumMessagesInPage:                      20,
				MaxMessageLength:                         20000,
				MaxForumMessageEditTimeout:               3600, /* секунд */
				// Первоапрельские форумы, в отличие от Perl-бэка, недоступны для любых действий (поскольку доступ к ним
				// реализован хардкодом в Auth.pm)
				DefaultAccessToForums:   []uint64{1, 2, 3, 5, 6, 7, 8, 10, 12, 13, 14, 15, 16, 17, 22},
				ForumsWithEnabledRating: []uint64{1, 3, 4, 6, 7, 9, 11, 12, 19},
				// В Perl-бэке список задан хардкодом в Forum.pm:
				// https://github.com/parserpro/fantlab/blob/19b295d70fd741e4f9a905e892ecf1358a87e35c/pm/Forum.pm#L1400-L1403
				ForumsWithDisabledMinuses: []uint64{9, 12},
				ReadOnlyForumUsers:        map[uint64][]uint64{9: {18938, 29932, 31525, 42100, 59417, 91217, 122312, 154599, 163940, 166509}},
				BlogsInPage:               50,
				BlogTopicsInPage:          20,
				BlogArticleCommentsInPage: 10,
				CensorshipText:            "[Сообщение изъято модератором]",
				PreModerationText:         "[Сообщение на премодерации]",
				ImageReplacementLinkText:  "[картинка]",
				BotUserId:                 2, // Р. Букашка
				MaxAttachCountPerMessage:  10,
				BookcaseItemInPage:        50,
				Smiles: config.MakeSmiles([]string{"smile", "wink", "glasses", "biggrin", "gigi", "lol", "haha", "frown",
					"confused", "insane", "weep", "abuse", "mad", "dont", "eek", "blush", "super", "pray", "box", "beer",
					"shuffle", "rev", "tired", "hihiks", "superkiss", "hb", "kiss2", "wink2", "love", "gun", "nunu2", "gy",
					"help", "hmm", "kap", "kult", "rom", "sad", "shock", "sla", "wht", "yar", "yaya", "glum", "alc", "friends",
					"appl", "bigeyes2", "bye", "cool", "drink", "facepalm", "kar", "lady", "mafia", "rolleyes", "spy", "roof",
					"silly", "lam", "inv", "leb", "val", "snuffy", "popcorn", "eyes", "glad2", "gorgeous", "kiss", "ninja",
					"nunu", "redface", "sleep", "sleep2", "surprise", "tongue"}),
				FlContestInProgress:               true,
				FlContestAuthorId:                 11677,
				CorrelationUserMarkCountThreshold: 14,
			}
			return nil
		},
	)

	if err != nil {
		fmt.Println(err) // for debug
		return nil, err, disposeBag
	}

	return &Services{
		appConfig:    appConfig,
		cryptoCoder:  cryptoCoder,
		db:           db.NewDB(mysqlDB),
		redis:        redisClient,
		memcache:     memcacheClient,
		smtp:         smtpClient,
		localStorage: ttlcache.NewWithDefaultExpireFunc(),
		minioClient:  minioClient,
		minioBucket:  minioBucket,
	}, nil, disposeBag
}

func makeMinioWithSecrets(server, accessKeyFile, secretKeyFile string) (*minio.Client, error) {
	accessKey, err := ioutil.ReadFile(accessKeyFile)
	if err != nil {
		return nil, err
	}
	secretKey, err := ioutil.ReadFile(secretKeyFile)
	if err != nil {
		return nil, err
	}
	return minio.New(server, string(accessKey), string(secretKey), false)
}
