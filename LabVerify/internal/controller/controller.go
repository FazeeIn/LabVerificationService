package controller

import (
	"encoding/json"
	"net/http"

	"github.com/FazeeIn/LabVerificationService/LabVerify/internal/docker"
	"github.com/FazeeIn/LabVerificationService/LabVerify/internal/model"
	"github.com/gin-gonic/gin"
)

func (a app) RunCheck(c *gin.Context) {
	userID := c.Params.ByName("userID")
	decoder := json.NewDecoder(c.Request.Body)
	var testRequest model.TestRequest

	err := decoder.Decode(&testRequest)
	if err != nil {
		http.Error(c.Writer, "Invalid request payload", http.StatusBadRequest)
		return
	}

	report, err := docker.NewContainer(testRequest, model.Python{})
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	a.reports[userID] = report
}

func (a app) DownloadReport(c *gin.Context) {
	userID := c.Params.ByName("userID")
	a.RunCheck(c)
	responseData := a.reports[userID]
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.Write(responseData)
}
