package controllers

import (
	"mytokulist/database"
	"mytokulist/models"
	"mytokulist/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllType(c *gin.Context) {
	var (
		result gin.H
	)

	categories, err := repository.GetAllType(database.DbConnection)

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

func InsertType(c *gin.Context) {
	var tipe models.Type

	err := c.ShouldBindJSON(&tipe)
	if err != nil {
		panic(err)
	}

	err = repository.InsertType(database.DbConnection, tipe)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert Type",
	})
}

func UpdateType(c *gin.Context) {
	var tipe models.Type

	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&tipe)
	if err != nil {
		panic(err)
	}

	tipe.ID = int64(id)

	err = repository.UpdateType(database.DbConnection, tipe)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Update Type",
	})
}

func DeleteType(c *gin.Context) {
	var tipe models.Type

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		panic(err)
	}

	tipe.ID = int64(id)

	err = repository.DeleteType(database.DbConnection, tipe)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Delete Type",
	})
}
