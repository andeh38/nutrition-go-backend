package controller

import (
	"net/http"
	"nutrition-api/helper"
	"nutrition-api/model"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Analytics struct {
	Date         time.Time `json:"date"`
	Calories     float64   `json:"calories"`
	Protein      float64   `json:"protein"`
	Fat          float64   `json:"fat"`
	Carbohydrate float64   `json:"carbohydrate"`
}

func accumulate(day model.Day) (float64, float64, float64, float64) {
	var calories, protein, fat, carbohydrate float64
	if day.Breakfast != nil {
		for _, food := range *day.Breakfast {
			calories += food.Calories
			protein += food.Protein
			fat += food.Fat
			carbohydrate += food.Carbohydrate
		}
	}
	if day.Brunch != nil {
		for _, food := range *day.Brunch {
			calories += food.Calories
			protein += food.Protein
			fat += food.Fat
			carbohydrate += food.Carbohydrate
		}
	}
	if day.Lunch != nil {
		for _, food := range *day.Lunch {
			calories += food.Calories
			protein += food.Protein
			fat += food.Fat
			carbohydrate += food.Carbohydrate
		}
	}
	if day.Dinner != nil {
		for _, food := range *day.Dinner {
			calories += food.Calories
			protein += food.Protein
			fat += food.Fat
			carbohydrate += food.Carbohydrate
		}
	}
	return calories, protein, fat, carbohydrate
}

func GetAnalytics(context *gin.Context) {
	params := context.Request.URL.Query()
	limit, err := strconv.Atoi(params["limit"][0])
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	maxDate := time.Now().UTC()
	maxDate.Format(time.RFC3339)
	y, m, d := maxDate.Date()
	minDate := time.Date(y, m, d-limit+1, 0, 0, 0, 0, time.UTC)
	maxDate = time.Date(y, m, d, 0, 0, 0, 0, time.UTC)

	user, err := helper.CurrentUser(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	days, err := model.SliceDays(&minDate, &maxDate, user.ID)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var data []Analytics
	currentDate := minDate

	for i := 0; i < len(days); {
		day := days[i]
		if day.Date.UTC() == currentDate {
			calories, protein, fat, carbohydrate := accumulate(day)
			res := Analytics{
				Date:         day.Date.UTC(),
				Calories:     calories,
				Protein:      protein,
				Fat:          fat,
				Carbohydrate: carbohydrate,
			}
			data = append(data, res)
			i++
		} else {
			res := Analytics{
				Date:         currentDate,
				Calories:     0,
				Protein:      0,
				Fat:          0,
				Carbohydrate: 0,
			}
			data = append(data, res)
		}
		y, m, d := currentDate.Date()
		currentDate = time.Date(y, m, d+1, 0, 0, 0, 0, time.UTC)
	}
	context.JSON(http.StatusOK, gin.H{"analytics": data})
}
