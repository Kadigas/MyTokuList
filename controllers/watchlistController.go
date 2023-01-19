package controllers

import (
	"mytokulist/database"
	"mytokulist/models"
	"mytokulist/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllWatchlist(c *gin.Context) {
	var (
		result gin.H
	)

	user, _ := c.Get("user")

	watchlist, err := repository.GetAllWatchlist(database.DbConnection, user.(models.Users).ID)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": watchlist,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertWatchlist(c *gin.Context) {
	var wl models.Watchlist

	err := c.ShouldBindJSON(&wl)
	if err != nil {
		panic(err)
	}

	user, _ := c.Get("user")

	err = repository.InsertWatchlist(database.DbConnection, wl, user.(models.Users).ID)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert Watchlist",
	})
}

func DeleteWatchlist(c *gin.Context) {
	var wl models.Watchlist

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		panic(err)
	}

	wl.ID = int64(id)

	user, _ := c.Get("user")

	err = repository.DeleteWatchlist(database.DbConnection, wl, user.(models.Users).ID)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Delete Watchlist",
	})
}
