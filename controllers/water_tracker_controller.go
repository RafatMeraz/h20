package controllers

import (
	"github.com/RafatMeraz/h20/repositories"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type WaterTrackerController struct {
	waterTrackRepository repositories.WaterTrackerRepository
}

func (controller WaterTrackerController) GetUserWaterTrack(c echo.Context) error {
	var userID uint
	userParam := c.Param("id")
	if userParam == "" {
		selfID, ok := c.Get("user-id").(uint)
		if !ok {
			return echo.ErrUnauthorized
		}
		userID = selfID
	} else {
		id, err := strconv.Atoi(userParam)
		if err != nil {
			return err
		}
		userID = uint(id)
	}
	waterTrackList, err := controller.waterTrackRepository.GetWaterConsumes(userID)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, waterTrackList)
}
