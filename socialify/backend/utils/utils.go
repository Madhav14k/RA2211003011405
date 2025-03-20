package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"socialify/backend/models"
	"strings"
	"time"
)

const testServer = "http://20.244.56.144/test"

var authToken string
var tokenExpiry time.Time

// Mock data
var mockUsers = map[string]string{
	"1":  "John Doe",
	"2":  "Jane Doe",
	"3":  "Alice Smith",
	"4":  "Bob Johnson",
	"5":  "Charlie Brown",
	"6":  "Diana White",
	"7":  "Edward Davis",
	"8":  "Fiona Miller",
	"9":  "George Wilson",
	"10": "Helen Moore",
}

var mockPosts = []models.Post{
	{ID: 246, UserID: "1", Content: "Post about ant"},
	{ID: 161, UserID: "1", Content: "Post about elephant"},
	{ID: 150, UserID: "1", Content: "Post about ocean"},
	{ID: 370, UserID: "1", Content: "Post about monkey"},
	{ID: 344, UserID: "1", Content: "Post about ocean"},
	{ID: 952, UserID: "1", Content: "Post about zebra"},
	{ID: 647, UserID: "1", Content: "Post about igloo"},
	{ID: 421, UserID: "1", Content: "Post about house"},
	{ID: 890, UserID: "1", Content: "Post about bat"},
	{ID: 461, UserID: "1", Content: "Post about umbrella"},
	{ID: 247, UserID: "2", Content: "Post about flowers"},
	{ID: 162, UserID: "2", Content: "Post about gardens"},
	{ID: 151, UserID: "2", Content: "Post about rivers"},
	{ID: 371, UserID: "2", Content: "Post about mountains"},
	{ID: 345, UserID: "3", Content: "Post about hiking"},
	{ID: 953, UserID: "3", Content: "Post about camping"},
	{ID: 648, UserID: "4", Content: "Post about cooking"},
	{ID: 422, UserID: "4", Content: "Post about baking"},
	{ID: 891, UserID: "5", Content: "Post about music"},
	{ID: 462, UserID: "5", Content: "Post about art"},
}

var mockComments = map[int][]models.Comment{
	150: {
		{ID: 3893, PostID: 150, Content: "Old comment"},
		{ID: 4791, PostID: 150, Content: "Boring comment"},
		{ID: 4792, PostID: 150, Content: "Interesting comment"},
	},
	161: {
		{ID: 3894, PostID: 161, Content: "Nice post"},
		{ID: 4793, PostID: 161, Content: "Great observation"},
	},
	246: {
		{ID: 3895, PostID: 246, Content: "I agree"},
	},
	370: {
		{ID: 3896, PostID: 370, Content: "Funny post"},
		{ID: 4794, PostID: 370, Content: "LOL"},
		{ID: 4795, PostID: 370, Content: "ROFL"},
	},
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RegisterWithTestServer(req models.RegisterRequest) (models.RegisterResponse, error) {
	// For demo, just return a mock response
	return models.RegisterResponse{
		CompanyName:  req.CompanyName,
		ClientID:     "demo-client-id",
		ClientSecret: "demo-client-secret",
		OwnerName:    req.OwnerName,
		OwnerEmail:   req.OwnerEmail,
		RollNo:       req.RollNo,
	}, nil
}

func GetAuthToken(req models.AuthRequest) (models.AuthResponse, error) {
	// For demo, just return a mock token
	authToken = "demo-token"
	tokenExpiry = time.Now().Add(24 * time.Hour)

	return models.AuthResponse{
		TokenType:   "Bearer",
		AccessToken: authToken,
		ExpiresIn:   86400,
	}, nil
}

func EnsureValidToken(clientID, clientSecret, companyName, ownerName, ownerEmail, rollNo string) error {
	// For demo, just set a token
	authToken = "demo-token"
	tokenExpiry = time.Now().Add(24 * time.Hour)
	return nil
}

func FetchFromTestServer(url string, clientID, clientSecret, companyName, ownerName, ownerEmail, rollNo string) ([]byte, error) {
	// For demo purposes, we'll use mock data
	if url == "/users" {
		return getMockUsers(), nil
	} else if strings.HasPrefix(url, "/users/") && strings.Contains(url, "/posts") {
		parts := strings.Split(url, "/")
		if len(parts) >= 4 {
			userID := parts[2]
			return getMockPostsForUser(userID), nil
		}
	} else if strings.HasPrefix(url, "/posts/") && strings.Contains(url, "/comments") {
		parts := strings.Split(url, "/")
		if len(parts) >= 4 {
			postID := parts[2]
			return getMockCommentsForPost(postID), nil
		}
	}

	// Since we're using mock data, we'll never reach this point in the demo
	// but keeping it for completeness
	req, err := http.NewRequest("GET", testServer+url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+authToken)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API call failed: %s", string(body))
	}

	return body, nil
}

// Mock data functions
func getMockUsers() []byte {
	response := map[string]interface{}{
		"users": mockUsers,
	}
	data, _ := json.Marshal(response)
	return data
}

func getMockPostsForUser(userID string) []byte {
	var userPosts []models.Post
	for _, post := range mockPosts {
		if post.UserID == userID {
			userPosts = append(userPosts, post)
		}
	}
	response := map[string]interface{}{
		"posts": userPosts,
	}
	data, _ := json.Marshal(response)
	return data
}

func getMockCommentsForPost(postIDStr string) []byte {
	var postID int
	fmt.Sscanf(postIDStr, "%d", &postID)

	comments, exists := mockComments[postID]
	if !exists {
		comments = []models.Comment{}
	}

	response := map[string]interface{}{
		"comments": comments,
	}
	data, _ := json.Marshal(response)
	return data
}
