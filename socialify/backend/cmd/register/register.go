package register

import (
	"encoding/json"
	"fmt"
	"socialify/backend/models"
	"socialify/backend/utils"
)

func RegisterWithServer() {
	registerReq := models.RegisterRequest{
		CompanyName: "socialify",
		OwnerName:   "userowner",
		OwnerEmail:  "user@example.com",
		RollNo:      "123456",
		AccessCode:  "FKDLJG", // Replace with your actual access code
	}

	resp, err := utils.RegisterWithTestServer(registerReq)
	if err != nil {
		fmt.Printf("Registration error: %s\n", err)
		return
	}

	respJson, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Printf("Registration successful:\n%s\n", string(respJson))

	authReq := models.AuthRequest{
		CompanyName:  resp.CompanyName,
		ClientID:     resp.ClientID,
		ClientSecret: resp.ClientSecret,
		OwnerName:    resp.OwnerName,
		OwnerEmail:   resp.OwnerEmail,
		RollNo:       resp.RollNo,
	}

	authResp, err := utils.GetAuthToken(authReq)
	if err != nil {
		fmt.Printf("Auth error: %s\n", err)
		return
	}

	authRespJson, _ := json.MarshalIndent(authResp, "", "  ")
	fmt.Printf("Auth successful:\n%s\n", string(authRespJson))

	fmt.Println("\nUse these environment variables:")
	fmt.Printf("export CLIENT_ID=\"%s\"\n", resp.ClientID)
	fmt.Printf("export CLIENT_SECRET=\"%s\"\n", resp.ClientSecret)
	fmt.Printf("export COMPANY_NAME=\"%s\"\n", resp.CompanyName)
	fmt.Printf("export OWNER_NAME=\"%s\"\n", resp.OwnerName)
	fmt.Printf("export OWNER_EMAIL=\"%s\"\n", resp.OwnerEmail)
	fmt.Printf("export ROLL_NO=\"%s\"\n", resp.RollNo)
}
