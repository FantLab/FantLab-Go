package main

import (
	"log"
	"os"

	"fantlab/config"
	"fantlab/logger"
	"fantlab/routing"
	"fantlab/shared"
	"fantlab/utils"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/minio/minio-go/v6"
)

func main() {
	// *********************************

	db, err := gorm.Open("mysql", os.Getenv("MYSQL_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := db.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	db.SetLogger(logger.GormLogger)
	db.LogMode(true)

	// *********************************

	minioClient, err := minio.New(
		os.Getenv("MINIO_ENDPOINT"),
		os.Getenv("MINIO_ACCESS_KEY"),
		os.Getenv("MINIO_SECRET_KEY"),
		os.Getenv("MINIO_USE_SSL") == "true",
	)

	if err != nil {
		log.Fatalln(err)
	}

	// *********************************

	configuration := config.ParseConfig(os.Getenv("CONFIG_FILE"))

	services := &shared.Services{
		Config:       configuration,
		DB:           db,
		S3Client:     minioClient,
		UrlFormatter: utils.UrlFormatter{Config: &configuration},
	}

	// *********************************

	router := routing.SetupWith(services)
	router.MaxMultipartMemory = 8 << 20 // TODO: ???

	if err := router.Run(":" + os.Getenv("PORT")); err != nil {
		log.Fatal(err)
	}
}
