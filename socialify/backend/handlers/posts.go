package handlers

import (
	"encoding/json"
	"net/http"
	"socialify/backend/models"
	"socialify/backend/utils"
	"sort"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUserPosts(c *gin.Context) {
	userID := c.Param("userId")
	url := "/users/" + userID + "/posts"

	body, err := utils.FetchFromTestServer(url, clientID, clientSecret, companyName, ownerName, ownerEmail, rollNo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var resp struct {
		Posts []models.Post `json:"posts"`
	}
	if err := json.Unmarshal(body, &resp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func GetPostComments(c *gin.Context) {
	postID := c.Param("postId")
	url := "/posts/" + postID + "/comments"

	body, err := utils.FetchFromTestServer(url, clientID, clientSecret, companyName, ownerName, ownerEmail, rollNo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var resp struct {
		Comments []models.Comment `json:"comments"`
	}
	if err := json.Unmarshal(body, &resp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func GetLatestPosts(c *gin.Context) {
	allPosts := make([]models.Post, 0)
	userIDMap := make(map[string]string)

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

	for id, name := range usersResp.Users {
		userIDMap[id] = name
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

		allPosts = append(allPosts, postsResp.Posts...)
	}

	sort.Slice(allPosts, func(i, j int) bool {
		return allPosts[i].ID > allPosts[j].ID
	})

	if len(allPosts) > 5 {
		allPosts = allPosts[:5]
	}

	result := make([]gin.H, 0)
	for _, post := range allPosts {
		result = append(result, gin.H{
			"post": post,
			"user": gin.H{
				"id":   post.UserID,
				"name": userIDMap[post.UserID],
			},
		})
	}

	c.JSON(http.StatusOK, gin.H{"latestPosts": result})
}

func GetPopularPosts(c *gin.Context) {
	postCommentCounts := make([]models.PostCommentCount, 0)
	userIDMap := make(map[string]string)

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

	for id, name := range usersResp.Users {
		userIDMap[id] = name
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

		for _, post := range postsResp.Posts {
			commentsURL := "/posts/" + strconv.Itoa(post.ID) + "/comments"
			commentsBody, err := utils.FetchFromTestServer(commentsURL, clientID, clientSecret, companyName, ownerName, ownerEmail, rollNo)
			if err != nil {
				continue
			}

			var commentsResp struct {
				Comments []models.Comment `json:"comments"`
			}
			if err := json.Unmarshal(commentsBody, &commentsResp); err != nil {
				continue
			}

			postCommentCounts = append(postCommentCounts, models.PostCommentCount{
				Post:         post,
				CommentCount: len(commentsResp.Comments),
			})
		}
	}

	sort.Slice(postCommentCounts, func(i, j int) bool {
		return postCommentCounts[i].CommentCount > postCommentCounts[j].CommentCount
	})

	maxCommentCount := 0
	if len(postCommentCounts) > 0 {
		maxCommentCount = postCommentCounts[0].CommentCount
	}

	popularPosts := make([]gin.H, 0)
	for _, pc := range postCommentCounts {
		if pc.CommentCount == maxCommentCount {
			popularPosts = append(popularPosts, gin.H{
				"post": pc.Post,
				"user": gin.H{
					"id":   pc.Post.UserID,
					"name": userIDMap[pc.Post.UserID],
				},
				"commentCount": pc.CommentCount,
			})
		} else {
			break
		}
	}

	c.JSON(http.StatusOK, gin.H{"popularPosts": popularPosts})
}
