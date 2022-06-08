package controllers

import (
	"strconv"

	"github.com/FredsZDev/WebApi_GO_B2/database"
	"github.com/FredsZDev/WebApi_GO_B2/models"
	"github.com/gin-gonic/gin"
)

func ShowGame(g *gin.Context) {
	id := g.Param("id")

	newid, err := strconv.Atoi(id)
	if err != nil {
		g.JSON(400, gin.H{
			"error": "O ID deve ser um inteiro",
		})
		return
	}

	db := database.GetDatabase()

	var game models.Game
	err = db.First(&game, newid).Error
	if err != nil {
		g.JSON(400, gin.H{
			"error": "Não foi possível encontrar o jogo" + err.Error(),
		})
		return
	}

	g.JSON(200, game)
}

func CreateGame(g *gin.Context) {
	db := database.GetDatabase()

	var game models.Game

	err := g.ShouldBindJSON(&game)
	if err != nil {
		g.JSON(400, gin.H{
			"error": "Nao foi possivel bindar o json" + err.Error(),
		})

		return

	}

	err = db.Create(&game).Error

	if err != nil {
		g.JSON(400, gin.H{
			"error": "sds" + err.Error(),
		})

	}

	g.JSON(200, game)

}

func ShowGames(g *gin.Context) {
	db := database.GetDatabase()

	var games []models.Game
	err := db.Find(&games).Error
	if err != nil {
		g.JSON(400, gin.H{
			"error": "cannot list games:" + err.Error(),
		})
		return
	}

	g.JSON(200, games)
}

func UpdateGames(g *gin.Context) {

	db := database.GetDatabase()

	var game models.Game

	err := g.ShouldBindJSON(&game)
	if err != nil {
		g.JSON(400, gin.H{
			"error": "Nao foi possivel fazer o update" + err.Error(),
		})

		return

	}

	err = db.Save(&game).Error

	if err != nil {
		g.JSON(400, gin.H{
			"error": "sds" + err.Error(),
		})

	}

	g.JSON(200, game)

}

func DeleteGames(g *gin.Context) {

	id := g.Param("id")

	newid, err := strconv.Atoi(id)
	if err != nil {
		g.JSON(400, gin.H{
			"error": "O ID deve ser um inteiro",
		})
		return
	}

	db := database.GetDatabase()
	err = db.Delete(&models.Game{}, newid).Error

	if err != nil {
		g.JSON(400, gin.H{
			"erro": "" + err.Error(),
		})

		return
	}

	g.Status(204)

}
