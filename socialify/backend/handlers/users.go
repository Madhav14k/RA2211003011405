package handlers

import (
	"encoding/json"
	"net/http"
	"socialify/backend/models"
	"socialify/backend/utils"
	"sort"

	"github.com/gin-gonic/gin"
)

type UsersResponse struct {
	Users map[string]string `json:"users"`
}

func GetUsers(c *gin.Context) {
	body, err := utils.FetchFromTestServer("/users", clientID, clientSecret, companyName, ownerName, ownerEmail, rollNo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var resp UsersResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func GetTopUsers(c *gin.Context) {
	body, err := utils.FetchFromTestServer("/users", clientID, clientSecret, companyName, ownerName, ownerEmail, rollNo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var usersResp UsersResponse
	if err := json.Unmarshal(body, &usersResp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	userPostCounts := make([]models.UserPostCount, 0)

	for id, name := range usersResp.Users {
		postsURL := "/users/" + id + "/posts"
		postsBody, err := utils.FetchFromTestServer(postsURL, clientID, clientSecret, companyName, ownerName, ownerEmail, rollNo)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		var postsResp struct {
			Posts []models.Post `json:"posts"`
		}
		if err := json.Unmarshal(postsBody, &postsResp); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		userPostCounts = append(userPostCounts, models.UserPostCount{
			User: models.User{
				ID:   id,
				Name: name,
			},
			PostCount: len(postsResp.Posts),
		})
	}

	sort.Slice(userPostCounts, func(i, j int) bool {
		return userPostCounts[i].PostCount > userPostCounts[j].PostCount
	})

	if len(userPostCounts) > 5 {
		userPostCounts = userPostCounts[:5]
	}

	c.JSON(http.StatusOK, gin.H{"topUsers": userPostCounts})
}
