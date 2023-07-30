package repositories

import "github.com/RafatMeraz/h20/models"

type WaterTrackerRepository interface {
	AddNewWaterConsume(request models.WaterTrackRequest) (uint, error)
	GetWaterConsumes(userId uint) ([]models.WaterTrackDTO, error)
	DeleteWaterConsume(userId uint, trackId uint) error
}
