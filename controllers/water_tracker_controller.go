package controllers

import (
	"github.com/RafatMeraz/h20/models"
	"github.com/RafatMeraz/h20/repositories"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strconv"
)

type WaterTrackerController struct {
	waterTrackRepository repositories.WaterTrackerRepository
}

func (controller WaterTrackerController) GetUserWaterTrack(c echo.Context) error {
	var userID uint
	userParam := c.Param("id")
	go log.Println(userParam)
	if userParam == "" {
		selfID := c.Get("user").(uint)
		if selfID == 0 {
			return echo.ErrUnauthorized
		}
		userID = selfID
	} else {
		paramID, err := strconv.Atoi(userParam)
		if err != nil {
			return err
		}
		userID = uint(paramID)
	}
	waterTrackList, err := controller.waterTrackRepository.GetWaterConsumes(userID)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, waterTrackList)
}

func (controller WaterTrackerController) AddNewWaterTrack(c echo.Context) error {
	v := validator.New()
	var waterTrackRequest models.WaterTrackRequest
	userId := c.Get("user").(uint)
	waterTrackRequest.UserId = userId
	if err := c.Bind(&waterTrackRequest); err != nil {
		return err
	}
	if err := v.Struct(waterTrackRequest); err != nil {
		return err
	}
	_, err := controller.waterTrackRepository.AddNewWaterConsume(waterTrackRequest)
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusCreated)
}
