// Main package for the Socialify backend server
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"socialify/backend/models"
	"socialify/backend/utils"
	"strconv"
	"time"
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

	// Default values for demo purposes
	if clientID == "" {
		clientID = "demo_client_id"
	}
	if clientSecret == "" {
		clientSecret = "demo_client_secret"
	}
	if companyName == "" {
		companyName = "socialify"
	}
	if ownerName == "" {
		ownerName = "userowner"
	}
	if ownerEmail == "" {
		ownerEmail = "user@example.com"
	}
	if rollNo == "" {
		rollNo = "123456"
	}
}

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req models.RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := utils.RegisterWithTestServer(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func authHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req models.AuthRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := utils.GetAuthToken(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	clientID = req.ClientID
	clientSecret = req.ClientSecret
	companyName = req.CompanyName
	ownerName = req.OwnerName
	ownerEmail = req.OwnerEmail
	rollNo = req.RollNo

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func topUsersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := utils.FetchFromTestServer("/users", clientID, clientSecret, companyName, ownerName, ownerEmail, rollNo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var usersResp struct {
		Users map[string]string `json:"users"`
	}

	if err := json.Unmarshal(body, &usersResp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userPostCounts := make([]models.UserPostCount, 0)

	for id, name := range usersResp.Users {
		postsURL := "/users/" + id + "/posts"
		postsBody, err := utils.FetchFromTestServer(postsURL, clientID, clientSecret, companyName, ownerName, ownerEmail, rollNo)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var postsResp struct {
			Posts []models.Post `json:"posts"`
		}

		if err := json.Unmarshal(postsBody, &postsResp); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
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

	// Sort by post count (descending)
	for i := 0; i < len(userPostCounts); i++ {
		for j := i + 1; j < len(userPostCounts); j++ {
			if userPostCounts[i].PostCount < userPostCounts[j].PostCount {
				userPostCounts[i], userPostCounts[j] = userPostCounts[j], userPostCounts[i]
			}
		}
	}

	if len(userPostCounts) > 5 {
		userPostCounts = userPostCounts[:5]
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"topUsers": userPostCounts,
	})
}

func latestPostsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	allPosts := make([]models.Post, 0)
	userIDMap := make(map[string]string)

	body, err := utils.FetchFromTestServer("/users", clientID, clientSecret, companyName, ownerName, ownerEmail, rollNo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var usersResp struct {
		Users map[string]string `json:"users"`
	}

	if err := json.Unmarshal(body, &usersResp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for id, name := range usersResp.Users {
		userIDMap[id] = name
		postsURL := "/users/" + id + "/posts"
		postsBody, err := utils.FetchFromTestServer(postsURL, clientID, clientSecret, companyName, ownerName, ownerEmail, rollNo)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var postsResp struct {
			Posts []models.Post `json:"posts"`
		}

		if err := json.Unmarshal(postsBody, &postsResp); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		allPosts = append(allPosts, postsResp.Posts...)
	}

	// Sort by post ID (descending) for latest
	for i := 0; i < len(allPosts); i++ {
		for j := i + 1; j < len(allPosts); j++ {
			if allPosts[i].ID < allPosts[j].ID {
				allPosts[i], allPosts[j] = allPosts[j], allPosts[i]
			}
		}
	}

	if len(allPosts) > 5 {
		allPosts = allPosts[:5]
	}

	result := make([]map[string]interface{}, 0)
	for _, post := range allPosts {
		result = append(result, map[string]interface{}{
			"post": post,
			"user": map[string]interface{}{
				"id":   post.UserID,
				"name": userIDMap[post.UserID],
			},
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"latestPosts": result,
	})
}

func popularPostsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	type PostCommentCount struct {
		Post         models.Post
		CommentCount int
	}

	postCommentCounts := make([]PostCommentCount, 0)
	userIDMap := make(map[string]string)

	body, err := utils.FetchFromTestServer("/users", clientID, clientSecret, companyName, ownerName, ownerEmail, rollNo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var usersResp struct {
		Users map[string]string `json:"users"`
	}

	if err := json.Unmarshal(body, &usersResp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for id, name := range usersResp.Users {
		userIDMap[id] = name
		postsURL := "/users/" + id + "/posts"
		postsBody, err := utils.FetchFromTestServer(postsURL, clientID, clientSecret, companyName, ownerName, ownerEmail, rollNo)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var postsResp struct {
			Posts []models.Post `json:"posts"`
		}

		if err := json.Unmarshal(postsBody, &postsResp); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
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

			postCommentCounts = append(postCommentCounts, PostCommentCount{
				Post:         post,
				CommentCount: len(commentsResp.Comments),
			})
		}
	}

	// Sort by comment count (descending)
	for i := 0; i < len(postCommentCounts); i++ {
		for j := i + 1; j < len(postCommentCounts); j++ {
			if postCommentCounts[i].CommentCount < postCommentCounts[j].CommentCount {
				postCommentCounts[i], postCommentCounts[j] = postCommentCounts[j], postCommentCounts[i]
			}
		}
	}

	maxCommentCount := 0
	if len(postCommentCounts) > 0 {
		maxCommentCount = postCommentCounts[0].CommentCount
	}

	popularPosts := make([]map[string]interface{}, 0)
	for _, pc := range postCommentCounts {
		if pc.CommentCount == maxCommentCount {
			popularPosts = append(popularPosts, map[string]interface{}{
				"post": pc.Post,
				"user": map[string]interface{}{
					"id":   pc.Post.UserID,
					"name": userIDMap[pc.Post.UserID],
				},
				"commentCount": pc.CommentCount,
			})
		} else {
			break
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"popularPosts": popularPosts,
	})
}

// SetupRoutes configures all HTTP routes for the server
func SetupRoutes() {
	http.Handle("/api/auth/register", enableCORS(http.HandlerFunc(registerHandler)))
	http.Handle("/api/auth/token", enableCORS(http.HandlerFunc(authHandler)))
	http.Handle("/api/users/top", enableCORS(http.HandlerFunc(topUsersHandler)))
	http.Handle("/api/posts/latest", enableCORS(http.HandlerFunc(latestPostsHandler)))
	http.Handle("/api/posts/popular", enableCORS(http.HandlerFunc(popularPostsHandler)))
}

func main() {
	SetupRoutes()

	server := &http.Server{
		Addr:         ":8081",
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
	}

	fmt.Println("Server is running on port 8081...")
	log.Fatal(server.ListenAndServe())
}
