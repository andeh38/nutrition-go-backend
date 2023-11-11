package controller

import (
	"net/http"
	"nutrition-api/helper"
	"nutrition-api/model"

	"github.com/gin-gonic/gin"
)

func AddDay(context *gin.Context) {
	var input model.Day
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := helper.CurrentUser(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input.UserID = user.ID

	savedDay, err := input.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"day": savedDay})
}

func GetDay(context *gin.Context) {
	var input model.Day
	var JSONdate map[string]interface{}
	if err := context.ShouldBindJSON(&JSONdate); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := helper.CurrentUser(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	date, err := helper.UnmarshalJSON(JSONdate["date"].(string))

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input.Date = &date
	input.UserID = user.ID

	day, err := input.FindOrCreate()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"day": day})
}

func Update(context *gin.Context) {
	var input model.Day
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	day, err := input.Update()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"day": day})
}

func AllDays(context *gin.Context) {
	days, err := model.AllDays()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"days": days})
}

func DeleteAll(context *gin.Context) {
	err := model.DeleteAll()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}
