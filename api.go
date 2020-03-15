package main

import (
	"fmt"
	"github.com/docker/go-connections/nat"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Image struct {
	ImageNameTag string `form:"imageNameTag" json:"imageNameTag" binding:"required"`
}

type Container struct {
	ContainerName string `form:"containerName" json:"containerName" binding:"required"`
}

type ContainerConfig struct {
	ImageNameTag  string `form:"imageNameTag" json:"imageNameTag" binding:"required"`
	ContainerName string `form:"containerName" json:"containerName" binding:"required"`
	ExportPort    int    `form:"exportPort" json:"exportPort" binding:"required"`
	HostPort      int    `form:"hostPort" json:"hostPort" binding:"required"`
	HostIP        int    `form:"hostIP" json:"hostIP" binding:"required"`
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

func StopContainerApi(c *gin.Context) {

	var err error
	var msg string
	var json Container
	if err = c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	msg, err = StopContainer(json.ContainerName)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"msg": msg, "err": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"msg": msg, "err": ""})
	}
}

func RemoveContainerApi(c *gin.Context) {

	var err error
	var msg string
	var json Container
	if err = c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	msg, err = RemoveContainer(json.ContainerName)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"msg": msg, "err": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"msg": msg, "err": ""})
	}
}

func RunContainerApi(c *gin.Context) {

	var err error
	var msg string
	var json ContainerConfig
	if err = c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	exportSet := nat.Port(fmt.Sprintf("%d/tcp", json.HostPort))
	hostPort := nat.Port(fmt.Sprintf("%d/tcp", json.HostPort))
	portSet := nat.PortSet{exportSet: struct{}{}}
	portBindings := nat.PortMap{
		hostPort: []nat.PortBinding{
			{
				HostIP:   strconv.Itoa(json.HostIP),
				HostPort: strconv.Itoa(json.ExportPort),
			},
		},
	}

	msg, err = RunContainer(json.ImageNameTag,
		json.ContainerName,
		portSet,
		portBindings)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"msg": msg, "err": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"msg": msg, "err": ""})
	}
}
