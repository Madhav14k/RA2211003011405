package models

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Post struct {
	ID      int    `json:"id"`
	UserID  string `json:"userid"`
	Content string `json:"content"`
}

type Comment struct {
	ID      int    `json:"id"`
	PostID  int    `json:"postid"`
	Content string `json:"content"`
}

type UserPostCount struct {
	User      User `json:"user"`
	PostCount int  `json:"postCount"`
}

type PostCommentCount struct {
	Post         Post `json:"post"`
	CommentCount int  `json:"commentCount"`
}

type RegisterRequest struct {
	CompanyName string `json:"companyName"`
	OwnerName   string `json:"ownerName"`
	RollNo      string `json:"rollNo"`
	OwnerEmail  string `json:"ownerEmail"`
	AccessCode  string `json:"accessCode"`
}

type RegisterResponse struct {
	CompanyName  string `json:"companyName"`
	ClientID     string `json:"clientID"`
	ClientSecret string `json:"clientSecret"`
	OwnerName    string `json:"ownerName"`
	OwnerEmail   string `json:"ownerEmail"`
	RollNo       string `json:"rollNo"`
}

type AuthRequest struct {
	CompanyName  string `json:"companyName"`
	ClientID     string `json:"clientID"`
	ClientSecret string `json:"clientSecret"`
	OwnerName    string `json:"ownerName"`
	OwnerEmail   string `json:"ownerEmail"`
	RollNo       string `json:"rollNo"`
}

type AuthResponse struct {
	TokenType   string `json:"token_type"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}
