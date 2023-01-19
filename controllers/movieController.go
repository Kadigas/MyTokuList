package controllers

import (
	"mytokulist/database"
	"mytokulist/models"
	"mytokulist/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllMovie(c *gin.Context) {
	var (
		result gin.H
	)

	movies, err := repository.GetAllMovie(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": movies,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertMovie(c *gin.Context) {
	var movie models.Movie

	err := c.ShouldBindJSON(&movie)
	if err != nil {
		panic(err)
	}

	err = repository.InsertMovie(database.DbConnection, movie)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert Movie",
	})
}

func UpdateMovie(c *gin.Context) {
	var movie models.Movie

	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&movie)
	if err != nil {
		panic(err)
	}

	movie.ID = int64(id)

	err = repository.UpdateMovie(database.DbConnection, movie)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Update Movie",
	})
}

func DeleteMovie(c *gin.Context) {
	var movie models.Movie

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		panic(err)
	}

	movie.ID = int64(id)

	err = repository.DeleteMovie(database.DbConnection, movie)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Delete Movie",
	})
}

func AttachMovie(c *gin.Context) {
	var movie models.Movie

	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&movie)
	if err != nil {
		panic(err)
	}

	movie.ID = int64(id)

	err = repository.AttachMovie(database.DbConnection, movie)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": "Category or Type is not valid.",
		})
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Attach Movie with Category and Type",
	})
}
