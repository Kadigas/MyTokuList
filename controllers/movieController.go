package controllers

import (
	"mytokulist/database"
	"mytokulist/repository"
	"mytokulist/structs"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllMovie(c *gin.Context) {
	var (
		result gin.H
	)

	categories, err := repository.GetAllMovie(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": categories,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertMovie(c *gin.Context) {
	var tipe structs.Movie

	err := c.ShouldBindJSON(&tipe)
	if err != nil {
		panic(err)
	}

	err = repository.InsertMovie(database.DbConnection, tipe)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert Movie",
	})
}

func UpdateMovie(c *gin.Context) {
	var tipe structs.Movie

	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&tipe)
	if err != nil {
		panic(err)
	}

	tipe.ID = int64(id)

	err = repository.UpdateMovie(database.DbConnection, tipe)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Update Movie",
	})
}

func DeleteMovie(c *gin.Context) {
	var tipe structs.Movie

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		panic(err)
	}

	tipe.ID = int64(id)

	err = repository.DeleteMovie(database.DbConnection, tipe)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Delete Movie",
	})
}
