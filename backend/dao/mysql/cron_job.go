package mysql

import (
	"chess-room-backend/model"
)

func GetCronJobListAdmin() ([]model.CronJob, error) {
	var cronJobs []model.CronJob
	err := DB.Order("created_at DESC").Find(&cronJobs).Error
	return cronJobs, err
}

func GetCronJobListAdminFiltered(name string, status, page, pageSize int) ([]model.CronJob, int64, error) {
	var cronJobs []model.CronJob
	var total int64
	db := DB.Model(&model.CronJob{}).Order("created_at DESC")
	if name != "" {
		db = db.Where("name LIKE ?", "%"+name+"%")
	}
	if status >= 0 {
		db = db.Where("status = ?", status)
	}
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if page > 0 && pageSize > 0 {
		db = DB.Order("created_at DESC")
		if name != "" {
			db = db.Where("name LIKE ?", "%"+name+"%")
		}
		if status >= 0 {
			db = db.Where("status = ?", status)
		}
		db = db.Offset((page - 1) * pageSize).Limit(pageSize)
	}
	err := db.Find(&cronJobs).Error
	return cronJobs, total, err
}

func GetCronJobByID(id int64) (*model.CronJob, error) {
	var cronJob model.CronJob
	err := DB.Where("id = ?", id).First(&cronJob).Error
	return &cronJob, err
}

func CreateCronJob(cronJob *model.CronJob) error {
	return DB.Create(cronJob).Error
}

func UpdateCronJob(cronJob *model.CronJob) error {
	return DB.Save(cronJob).Error
}

func DeleteCronJob(id int64) error {
	return DB.Delete(&model.CronJob{}, id).Error
}

func BatchDeleteCronJob(ids []int64) error {
	return DB.Where("id IN (?)", ids).Delete(&model.CronJob{}).Error
}