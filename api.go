package main

import (
	"fmt"
	"github.com/docker/go-connections/nat"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Image struct {
	ImageNameTag string `form:"imageNameTag" json:"imageNameTag" binding:"required" example:"mileslin/dockerlab:latest"`
}

type Container struct {
	ContainerName string `form:"containerName" json:"containerName" binding:"required" example:"dockerapp"`
}

type ContainerConfig struct {
	ImageNameTag  string   `form:"imageNameTag" json:"imageNameTag" binding:"required" example:"mileslin/dockerlab:latest"`
	ContainerName string   `form:"containerName" json:"containerName" binding:"required" example:"dockerapp"`
	ExportPort    string   `form:"exportPort" json:"exportPort" binding:"required" example:"80"`
	HostPort      string   `form:"hostPort" json:"hostPort" binding:"required" example:"8080"`
	HostIP        string   `form:"hostIP" json:"hostIP" binding:"required" example:"0.0.0.0"`
	RestartPolicy string   `form:"restartPolicy" json:"restartPolicy" example:"always"`
	Env           []string `form:"env" json:"env" example:"abc=123,xyz=999"`
}

// @Summary Remove an image
// @Description Remove an image by image name
// @Accept  json
// @Produce  json
// @Param body body Image true "the body content"
// @Success 200 {body} string "the sample of body is {\"msg\": \"message\", \"err\":\"message\"}"
// @Router /rmi [post]
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

// @Summary Pull an image
// @Description Pull an image by image name
// @Accept  json
// @Produce  json
// @Param body body Image true "the body content"
// @Success 200 {body} string "the sample of body is {\"msg\": \"message\", \"err\":\"message\"}"
// @Router /pull [post]
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

// @Summary Stop a container
// @Description Stop a container by container name
// @Accept  json
// @Produce  json
// @Param body body Container true "the body content"
// @Success 200 {body} string "the sample of body is {\"msg\": \"message\", \"err\":\"message\"}"
// @Router /stop [post]
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

// @Summary Remove a container
// @Description Remove a container by container name
// @Accept  json
// @Produce  json
// @Param body body Container true "the body content"
// @Success 200 {body} string "the sample of body is {\"msg\": \"message\", \"err\":\"message\"}"
// @Router /rm [post]
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

// @Summary Run a container
// @Description It do create and start to run the container a container by container name.
// @Accept  json
// @Produce  json
// @Param body body ContainerConfig true "the body content"
// @Success 200 {body} string "the sample of body is {\"msg\": \"message\", \"err\":\"message\"}"
// @Router /run [post]
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

// @Summary Update a running container
// @Description It do 1. stop container, 2. remove container 3. remove image 4. pull image 5. run container.  If one step failed, then it stop immediately.
// @Accept  json
// @Produce  json
// @Param body body ContainerConfig true "the body content"
// @Success 200 {body} string "the sample of body is {\"msg\": \"message\", \"err\":\"message\"}"
// @Router /updaterunningcontainer [post]
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
