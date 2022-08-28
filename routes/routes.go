package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Champion struct {
	Name  string `json:"name"`
	Quote string `json:"quote"`
}

var Router = gin.Default()
var championReq Champion
var champions = map[string]string{
	"Yasuo": "This blade never gets any lighter.",
	"Zed":   "Do not fear the shrouded path.",
}

func InitRoutes() {
	publicRoutes := Router.Group("v1/")
	publicRoutes.GET("health", HealthCheck)
	publicRoutes.GET("champion", GetChampion)
	publicRoutes.POST("champion", AddChampion)
	publicRoutes.PUT("champion", UpdateChampion)
	publicRoutes.DELETE("champion", DeleteChampion)
}

func HealthCheck(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "API is up and working fine",
	})
}

func GetChampion(context *gin.Context) {
	champion := context.Query("championName")
	championQuote, ok := champions[champion]
	if !ok {
		context.JSON(http.StatusNotFound, gin.H{
			champion: "Champion not exsits",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		champion: championQuote,
	})
}

func AddChampion(context *gin.Context) {
	context.BindJSON(&championReq)
	if len(championReq.Name) == 0 || len(championReq.Quote) == 0 {
		context.JSON(http.StatusBadRequest, gin.H{
			championReq.Name: championReq.Quote,
		})
		return
	}
	if _, err := champions[championReq.Name]; err {
		context.JSON(http.StatusConflict, gin.H{
			"message": "Champion already exsists",
		})
		return
	}

	champions[championReq.Name] = championReq.Quote
	context.JSON(http.StatusCreated, gin.H{
		"message": "Champion added",
	})
}

func UpdateChampion(context *gin.Context) {
	context.BindJSON(&championReq)
	if len(championReq.Name) == 0 || len(championReq.Quote) == 0 {
		context.JSON(http.StatusBadRequest, gin.H{
			championReq.Name: championReq.Quote,
		})
		return
	}
	_, ok := champions[championReq.Name]
	if !ok {
		context.JSON(http.StatusNotFound, gin.H{
			championReq.Name: "Champion not exsits",
		})
		return
	}

	champions[championReq.Name] = championReq.Quote
	context.JSON(http.StatusCreated, gin.H{
		"message": "Champion updated",
	})
}

func DeleteChampion(context *gin.Context) {
	champion := context.Query("championName")
	delete(champions, champion)
	context.JSON(http.StatusOK, gin.H{
		"message": "Champion deleted successfully",
	})
}
