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
	RestartPolicy string   `form:"restartPolicy" json:"restartPolicy" example:"always"` // It supports `no`, `always`, `on-failure`, `unless-stopped`
	Env           []string `form:"env" json:"env" example:"abc=123,xyz=999"`
	Mount         []Mount  `form:"mount" json:"mount"`
}

type ContainerConfigWithAuth struct {
	ContainerConfig
	WithAuth bool `form:"withAuth" json:"withAuth" example:"true`
}

type Mount struct {
	Type   string `form:"type" json:"type" example:"volume"` // It supports `bind`, `volume`, `tmpfs`, `npipe`
	Source string `form:"source" json:"source" example:"myvolume"`
	Target string `form:"target" json:"target" example:"/app/appdata"`
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

// @Summary Pull an image with authentication
// @Description Pull an image with authentication by image name. Your should set -cracct and -crpwd flag for username and password when running the console.
// @Accept  json
// @Produce  json
// @Param body body Image true "the body content"
// @Success 200 {body} string "the sample of body is {\"msg\": \"message\", \"err\":\"message\"}"
// @Router /pull [post]
func PullImageWithAuthApi(c *gin.Context) {

	var err error
	var msg string
	var json Image
	if err = c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	msg, err = PullImageWithAuth(json.ImageNameTag)
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
		json.RestartPolicy,
		json.Mount)

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
// @Param body body ContainerConfigWithAuth true "the body content"
// @Success 200 {body} string "the sample of body is {\"msg\": \"message\", \"err\":\"message\"}"
// @Router /updaterunningcontainer [post]
func UpdateRunningContainerApi(c *gin.Context) {
	var err error
	var msg string
	var errMsg string
	var tempMsg string
	var json ContainerConfigWithAuth
	if err = c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tempMsg, err = StopContainer(json.ContainerName)
	msg += tempMsg
	if err != nil {
		errMsg += err.Error()
	}

	tempMsg, err = RemoveContainer(json.ContainerName)
	msg += tempMsg
	if err != nil {
		errMsg += err.Error()
	}

	tempMsg, err = RemoveImage(json.ImageNameTag)
	msg += tempMsg
	if err != nil {
		errMsg += err.Error()
	}
	if json.WithAuth {
		tempMsg, err = PullImageWithAuth(json.ImageNameTag)
		msg += tempMsg
		if err != nil {
			errMsg += err.Error()
		}
	} else {
		tempMsg, err = PullImage(json.ImageNameTag)
		msg += tempMsg
		if err != nil {
			errMsg += err.Error()
		}
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
		json.RestartPolicy,
		json.Mount)

	msg += tempMsg
	if err != nil {
		errMsg += err.Error()
	}
	c.JSON(http.StatusOK, gin.H{"msg": msg, "err": errMsg})

}
