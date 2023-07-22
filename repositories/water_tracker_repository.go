package repositories

import (
	"errors"
	"github.com/RafatMeraz/h20/database"
	"github.com/RafatMeraz/h20/models"
	"gorm.io/gorm"
)

type WaterTrackerRepository struct{}

func (WaterTrackerRepository) AddNewWaterConsume(request models.WaterTrackRequest) (uint, error) {
	waterTrack := models.WaterTrack{UserId: request.UserId, Amount: request.Amount, ConsumeTime: request.ConsumeTime}
	result := database.Database.Instance().Create(&waterTrack)
	if result.Error != nil {
		return 0, result.Error
	}
	if result.RowsAffected == 0 {
		return 0, errors.New("unable to add record")
	}
	return waterTrack.ID, nil
}

func (WaterTrackerRepository) GetWaterConsumes(userId uint) ([]models.WaterTrackDTO, error) {
	var waterTrackList []models.WaterTrackDTO
	result := database.Database.Instance().Table("water_tracks").Where("user_id = ?", userId).Find(&waterTrackList)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return waterTrackList, nil
		}
		return nil, result.Error
	}
	return waterTrackList, nil
}
