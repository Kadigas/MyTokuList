package controllers

import (
	"mytokulist/database"
	"mytokulist/repository"
	"mytokulist/structs"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllStatus(c *gin.Context) {
	var (
		result gin.H
	)

	categories, err := repository.GetAllStatus(database.DbConnection)

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

func InsertStatus(c *gin.Context) {
	var status structs.Status

	err := c.ShouldBindJSON(&status)
	if err != nil {
		panic(err)
	}

	err = repository.InsertStatus(database.DbConnection, status)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert Status",
	})
}

func UpdateStatus(c *gin.Context) {
	var status structs.Status

	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&status)
	if err != nil {
		panic(err)
	}

	status.ID = int64(id)

	err = repository.UpdateStatus(database.DbConnection, status)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Update Status",
	})
}

func DeleteStatus(c *gin.Context) {
	var status structs.Status

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		panic(err)
	}

	status.ID = int64(id)

	err = repository.DeleteStatus(database.DbConnection, status)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Delete Status",
	})
}
