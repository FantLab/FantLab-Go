package db

import (
	"context"
	"fantlab/core/db/queries"
	"github.com/FantLab/go-kit/codeflow"
	"github.com/FantLab/go-kit/database/sqlapi"
	"math"
)

type UserMark struct {
	Mark       uint8   `db:"mark"`
	MarkWeight float64 `db:"mark_weight"`
	Sex        uint8   `db:"sex"`
}

func (db *DB) UpsertMark(ctx context.Context, userId, workId uint64, workAuthorIds []uint64, isFlContestWork bool,
	mark uint8, correlationUserMarkCountThreshold uint64) (WorkStats, error) {
	var marks []UserMark
	var userMarkCount uint64
	var stats WorkStats

	err := db.engine.InTransaction(func(rw sqlapi.ReaderWriter) error {
		return codeflow.Try(
			func() error { // Сохраняем или удаляем оценку
				if mark == 0 {
					return rw.Write(ctx, sqlapi.NewQuery(queries.MarkDeleteUserWorkMark).WithArgs(userId, workId)).Error
				} else {
					return rw.Write(ctx, sqlapi.NewQuery(queries.MarkUpsertUserWorkMark).WithArgs(userId, workId, mark, mark)).Error
				}
			},
			func() error { // Получаем список оценок произведения
				if isFlContestWork {
					return nil
				} else {
					return rw.Read(ctx, sqlapi.NewQuery(queries.MarkGetWorkMarks).WithArgs(workId), &marks)
				}
			},
			func() error { // Считаем и сохраняем статистику (или удаляем запись, если это фант. лабораторная работа)
				if isFlContestWork {
					return rw.Write(ctx, sqlapi.NewQuery(queries.WorkStatsDeleteWorkStats).WithArgs(workId)).Error
				} else {
					stats = calcWorkStats(marks)
					return rw.Write(ctx, sqlapi.NewQuery(queries.WorkStatsUpsertWorkStats).WithArgs(workId, stats.AverageMark,
						stats.AverageMarkByWeight, stats.MarkCount, stats.Rating, stats.AverageMarkGenderDelta,
						stats.AverageMarkMale, stats.AverageMarkFemale, stats.MarkCountMale, stats.MarkCountFemale,
						stats.AverageMarkDelta, stats.AverageMark, stats.AverageMarkByWeight, stats.MarkCount, stats.Rating,
						stats.AverageMarkGenderDelta, stats.AverageMarkMale, stats.AverageMarkFemale, stats.MarkCountMale,
						stats.MarkCountFemale, stats.AverageMarkDelta)).Error
				}
			},
			func() error { // Обновляем количество оценок пользователя
				return rw.Write(ctx, sqlapi.NewQuery(queries.UserUpdateMarkCount).WithArgs(userId)).Error
			},
			func() error { // Получаем обновленное количество оценок
				return rw.Read(ctx, sqlapi.NewQuery(queries.UserGetMarkCount).WithArgs(userId), &userMarkCount)
			},
			func() error { // Делаем запись о необходимости пересчета корреляций пользователя
				if userMarkCount > correlationUserMarkCountThreshold {
					return rw.Write(ctx, sqlapi.NewQuery(queries.CorrelationsInsertUserWorkUpdate).WithArgs(userId, workId)).Error
				} else {
					return nil
				}
			},
			func() error { // Выставляем флаги для Cron-а о необходимости пересчета статистики авторов
				if len(workAuthorIds) > 0 {
					return rw.Write(ctx, sqlapi.NewQuery(queries.AutorMarkAutorsNeedRecalcStats).WithArgs(workAuthorIds).FlatArgs()).Error
				} else {
					return nil
				}
			},
		)
	})

	if err != nil {
		return WorkStats{}, err
	}

	return stats, nil
}

type WorkStats struct {
	AverageMark            float64
	AverageMarkByWeight    float64
	MarkCount              uint64
	Rating                 float64
	AverageMarkGenderDelta float64
	AverageMarkMale        float64
	AverageMarkFemale      float64
	MarkCountMale          uint64
	MarkCountFemale        uint64
	AverageMarkDelta       float64
}

func calcWorkStats(usersMarks []UserMark) WorkStats {
	// Число оценок у произведения, когда начинается переход от средней оценки к средневзвешенной.
	// Уже при этом количестве начинают работать доверительные веса
	var minMarkCount = uint64(31)
	// Число оценок у произведения, когда заканчивается переход от средней оценки к средневзвешенной.
	// Это последнее значение, где ещё идёт переход. При этом количестве плюс один будет чистая средневзвешенная.
	var maxMarkCount = uint64(200)

	averageMark := float64(0)
	averageMarkByWeight := float64(0)
	averageMarkMale := float64(0)
	markCountMale := uint64(0)
	averageMarkFemale := float64(0)
	markCountFemale := uint64(0)
	rating := float64(0)
	averageMarkDelta := float64(0)
	averageMarkGenderDelta := float64(0)

	markCount := uint64(len(usersMarks))

	sumOfMarkWeights := float64(0)
	for _, userMark := range usersMarks {
		sumOfMarkWeights += userMark.MarkWeight
	}

	if markCount > 0 && (markCount < minMarkCount || sumOfMarkWeights > 0) {
		var marks []uint8

		sumOfMarks := uint64(0)
		sumOfMarksMale := uint64(0)
		sumOfMarksFemale := uint64(0)
		for _, userMark := range usersMarks {
			sumOfMarks += uint64(userMark.Mark)
			averageMarkByWeight += float64(userMark.Mark) * userMark.MarkWeight
			if userMark.Sex == 1 {
				sumOfMarksMale += uint64(userMark.Mark)
				markCountMale++
			} else {
				sumOfMarksFemale += uint64(userMark.Mark)
				markCountFemale++
			}
			marks = append(marks, userMark.Mark)
		}

		averageMark = float64(sumOfMarks) / float64(markCount)
		if markCountMale > 0 {
			averageMarkMale = float64(sumOfMarksMale) / float64(markCountMale)
		} else {
			averageMarkMale = 0
		}
		if markCountFemale > 0 {
			averageMarkFemale = float64(sumOfMarksFemale) / float64(markCountFemale)
		} else {
			averageMarkFemale = 0
		}
		if markCount < minMarkCount {
			averageMarkByWeight = averageMark
		} else if sumOfMarkWeights > 0 {
			averageMarkByWeight = averageMarkByWeight / sumOfMarkWeights
			if markCount <= maxMarkCount {
				averageMarkByWeight = (averageMark*float64(maxMarkCount-markCount+1) + averageMarkByWeight*float64(markCount-minMarkCount+1)) /
					float64(maxMarkCount-minMarkCount+2)
			}
		}
		var boundedMarkCount = markCount
		if boundedMarkCount > 1000 {
			boundedMarkCount = 1000
		}
		rating = averageMarkByWeight * (1 - 1/float64(boundedMarkCount/10+1))

		sumOfMarkDelta := float64(0)
		for _, mark := range marks {
			sumOfMarkDelta += math.Abs(float64(mark) - averageMark)
		}
		averageMarkDelta = sumOfMarkDelta / float64(markCount)

		averageMarkGenderDelta = averageMarkMale - averageMarkFemale
	}

	return WorkStats{
		AverageMark:            averageMark,
		AverageMarkByWeight:    averageMarkByWeight,
		MarkCount:              markCount,
		Rating:                 rating,
		AverageMarkGenderDelta: averageMarkGenderDelta,
		AverageMarkMale:        averageMarkMale,
		AverageMarkFemale:      averageMarkFemale,
		MarkCountMale:          markCountMale,
		MarkCountFemale:        markCountFemale,
		AverageMarkDelta:       averageMarkDelta,
	}
}
