package handlers

import (
	"net/http"
	"os"
	"socialify/backend/models"
	"socialify/backend/utils"

	"github.com/gin-gonic/gin"
)

var (
	clientID     string
	clientSecret string
	companyName  string
	ownerName    string
	ownerEmail   string
	rollNo       string
)

func init() {
	clientID = os.Getenv("CLIENT_ID")
	clientSecret = os.Getenv("CLIENT_SECRET")
	companyName = os.Getenv("COMPANY_NAME")
	ownerName = os.Getenv("OWNER_NAME")
	ownerEmail = os.Getenv("OWNER_EMAIL")
	rollNo = os.Getenv("ROLL_NO")
}

func Register(c *gin.Context) {
	var req models.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := utils.RegisterWithTestServer(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func Auth(c *gin.Context) {
	var req models.AuthRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := utils.GetAuthToken(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	clientID = req.ClientID
	clientSecret = req.ClientSecret
	companyName = req.CompanyName
	ownerName = req.OwnerName
	ownerEmail = req.OwnerEmail
	rollNo = req.RollNo

	c.JSON(http.StatusOK, resp)
}
