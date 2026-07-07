package logic

import (
	"chess-room-backend/dao/mysql"
	"chess-room-backend/dao/redis"
	"chess-room-backend/model"
	"chess-room-backend/pkg/errno"
	"strconv"

	"github.com/jinzhu/gorm"
)

func GetCronJobListAdmin() ([]model.CronJob, error) {
	cronJobs, err := mysql.GetCronJobListAdmin()
	if err != nil {
		return nil, errno.New(errno.InternalError)
	}
	return cronJobs, nil
}

func GetCronJobListAdminFiltered(name string, status, page, pageSize int) ([]model.CronJob, int64, error) {
	cronJobs, total, err := mysql.GetCronJobListAdminFiltered(name, status, page, pageSize)
	if err != nil {
		return nil, 0, errno.New(errno.InternalError)
	}
	return cronJobs, total, nil
}

func GetCronJobByID(id string) (*model.CronJob, error) {
	cronJobID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, errno.New(errno.BadRequest)
	}
	cronJob, err := mysql.GetCronJobByID(cronJobID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errno.New(errno.CronJobNotFound)
		}
		return nil, errno.New(errno.InternalError)
	}
	return cronJob, nil
}

func CreateCronJob(cronJob *model.CronJob) error {
	if err := mysql.CreateCronJob(cronJob); err != nil {
		return errno.New(errno.InternalError)
	}
	redis.Del("cron_job:list")
	return nil
}

func UpdateCronJob(id string, cronJob *model.CronJob) error {
	cronJobID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return errno.New(errno.BadRequest)
	}
	existing, err := mysql.GetCronJobByID(cronJobID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errno.New(errno.CronJobNotFound)
		}
		return errno.New(errno.InternalError)
	}
	existing.Name = cronJob.Name
	existing.CronExpression = cronJob.CronExpression
	existing.Handler = cronJob.Handler
	existing.Params = cronJob.Params
	existing.Status = cronJob.Status
	existing.Description = cronJob.Description
	if err := mysql.UpdateCronJob(existing); err != nil {
		return errno.New(errno.InternalError)
	}
	redis.Del("cron_job:list")
	return nil
}

func DeleteCronJob(id string) error {
	cronJobID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return errno.New(errno.BadRequest)
	}
	if err := mysql.DeleteCronJob(cronJobID); err != nil {
		return errno.New(errno.InternalError)
	}
	redis.Del("cron_job:list")
	return nil
}

func BatchDeleteCronJob(ids []string) error {
	var cronJobIDs []int64
	for _, id := range ids {
		cronJobID, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			return errno.New(errno.BadRequest)
		}
		cronJobIDs = append(cronJobIDs, cronJobID)
	}
	if err := mysql.BatchDeleteCronJob(cronJobIDs); err != nil {
		return errno.New(errno.InternalError)
	}
	redis.Del("cron_job:list")
	return nil
}

func BatchUpdateCronJob(reqs []struct {
	ID     int64 `json:"id"`
	Status int   `json:"status"`
}) error {
	for _, req := range reqs {
		cronJob, err := mysql.GetCronJobByID(req.ID)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return errno.New(errno.CronJobNotFound)
			}
			return errno.New(errno.InternalError)
		}

		cronJob.Status = req.Status

		if err := mysql.UpdateCronJob(cronJob); err != nil {
			return errno.New(errno.InternalError)
		}
	}

	redis.Del("cron_job:list")
	return nil
}
