package main

import (
	"fmt"
	"net/http"
	"server/restAPI/models"

	"github.com/gin-gonic/gin"
)

func main() {
	var topic string = "myTopic 1"
	fmt.Print(topic)
	server := gin.Default()

	// GET, POST, PUT, PATCH, DELETE
	server.GET("/jobs/:topicname", getJobs)     //GET => Consume from kafka
	server.POST("/jobs/:topicname", createJobs) //POST => imagine it is send to database :)
	server.Run(":8080")                         //localhost:8080
}

func getJobs(context *gin.Context) {
	topicname := context.Param("topicname")
	jobs := models.ConsumeTopic(topicname)
	context.JSON(http.StatusOK, jobs)
	models.SendEmail(topicname, jobs, "hosamag54@gmail.com")

}

func createJobs(context *gin.Context) {
	topicname := context.Param("topicname")
	jobs := models.ConsumeTopic(topicname)
	context.ShouldBindJSON(&jobs)

	jobs.Save() // for now we do not have database so save it localy
	context.JSON(http.StatusCreated, gin.H{
		"message": "Job created",
		"job":     jobs,
	})
}
