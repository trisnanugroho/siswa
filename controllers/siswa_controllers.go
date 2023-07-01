package controllers

import (
	"siswa/configs"
	"siswa/models"

	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func TambahSiswaController(c echo.Context) error {
	var students models.Students
	c.Bind(&students)

	result := configs.DB.Create(&students)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Message: "Error",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "Success",
		Data:    students,
	})
}

func DetailSiswaController(c echo.Context) error {
	var id, _ = strconv.Atoi(c.Param("id"))

	var data = models.Students{Id: id, Name: "Trisna", Class: 5, Sex: "Men", City: "Cirebon"}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "Success",
		Data:    data,
	})
}

func SiswaController(c echo.Context) error {
	var data []models.Students

	result := configs.DB.Find(&data)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Message: "Error",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "Success",
		Data:    data,
	})
}

func HapusSiswaController(c echo.Context) error {
	var id, _ = strconv.Atoi(c.Param("id"))

	var students models.Students
	c.Bind(&students)

	result := configs.DB.Delete(&students, id)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Message: "Error",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "Success",
		Data:    students,
	})
}
