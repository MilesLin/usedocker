package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Image struct {
	ImageNameTag string `form:"imageNameTag" json:"imageNameTag" binding:"required"`
}

func RemoveImageApi(c *gin.Context) {

	var err error
	var msg string
	var json Image
	if err = c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	msg, err = RemoveImage(json.ImageNameTag)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"msg": msg, "err": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"msg": msg, "err": ""})
	}
}

func PullImageApi(c *gin.Context) {

	var err error
	var msg string
	var json Image
	if err = c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	msg, err = PullImage(json.ImageNameTag)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"msg": msg, "err": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"msg": msg, "err": ""})
	}
}
