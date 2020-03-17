package main

import (
	"fmt"
	"github.com/docker/go-connections/nat"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Image struct {
	ImageNameTag string `form:"imageNameTag" json:"imageNameTag" binding:"required"`
}

type Container struct {
	ContainerName string `form:"containerName" json:"containerName" binding:"required"`
}

type ContainerConfig struct {
	ImageNameTag  string   `form:"imageNameTag" json:"imageNameTag" binding:"required"`
	ContainerName string   `form:"containerName" json:"containerName" binding:"required"`
	ExportPort    string   `form:"exportPort" json:"exportPort" binding:"required"`
	HostPort      string   `form:"hostPort" json:"hostPort" binding:"required"`
	HostIP        string   `form:"hostIP" json:"hostIP" binding:"required"`
	RestartPolicy string   `form:"restartPolicy" json:"restartPolicy"`
	Env           []string `form:"env" json:"env"`
}

// ShowAccount godoc
// @Summary Show a account
// @Description get string by ID
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Param id path int true "Account ID"
// @Success 200 {object} main.Image
// @Header 200 {string} Token "qwerty"
// @Router /accounts/{id} [get]
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
	exportSet := nat.Port(fmt.Sprintf("%s/tcp", json.ExportPort))
	portSet := nat.PortSet{exportSet: struct{}{}}

	portBindings := nat.PortMap{
		exportSet: []nat.PortBinding{
			{
				HostIP:   json.HostIP,
				HostPort: json.HostPort,
			},
		},
	}

	msg, err = RunContainer(json.ImageNameTag,
		json.ContainerName,
		portSet,
		portBindings,
		json.Env,
		json.RestartPolicy)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"msg": msg, "err": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"msg": msg, "err": ""})
	}

}

func UpdateRunningContainerApi(c *gin.Context) {
	var err error
	var msg string
	var tempMsg string
	var json ContainerConfig
	if err = c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if tempMsg, err = StopContainer(json.ContainerName); err != nil {
		msg += tempMsg
		c.JSON(http.StatusOK, gin.H{"msg": msg, "err": err.Error()})
		return
	}

	if tempMsg, err = RemoveContainer(json.ContainerName); err != nil {
		msg += tempMsg
		c.JSON(http.StatusOK, gin.H{"msg": msg, "err": err.Error()})
		return
	}

	if tempMsg, err = RemoveImage(json.ImageNameTag); err != nil {
		msg += tempMsg
		c.JSON(http.StatusOK, gin.H{"msg": msg, "err": err.Error()})
		return
	}

	if tempMsg, err = PullImage(json.ImageNameTag); err != nil {
		msg += tempMsg
		c.JSON(http.StatusOK, gin.H{"msg": msg, "err": err.Error()})
		return
	}

	exportSet := nat.Port(fmt.Sprintf("%s/tcp", json.ExportPort))
	portSet := nat.PortSet{exportSet: struct{}{}}

	portBindings := nat.PortMap{
		exportSet: []nat.PortBinding{
			{
				HostIP:   json.HostIP,
				HostPort: json.HostPort,
			},
		},
	}

	tempMsg, err = RunContainer(json.ImageNameTag,
		json.ContainerName,
		portSet,
		portBindings,
		json.Env,
		json.RestartPolicy)

	msg += tempMsg

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"msg": msg, "err": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"msg": msg, "err": ""})
	}
}
