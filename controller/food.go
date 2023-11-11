package controller

import (
	"net/http"
	"nutrition-api/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddFood(context *gin.Context) {
	var input model.Food
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	savedFood, err := input.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": savedFood})
}

func FindFoodByName(context *gin.Context) {
	params := context.Request.URL.Query()
	name := params.Get("name")
	limit, err := strconv.Atoi(params["limit"][0])

	if name == "" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "name is required"})
		return
	}

	food, err := model.FindFoodByName(name, limit)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"food": food})
}

func GetAllFood(context *gin.Context) {
	food, err := model.AllFood()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": food})
}
