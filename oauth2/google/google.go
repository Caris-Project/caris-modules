package google

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Config ...
type Config struct {
	ClientID string
}

var config Config

// Init ...
func Init(clientID string) {
	config.ClientID = clientID
}

type (
	// Result ...
	Result struct {
		ID      string `json:"sub"`
		AUD     string `json:"aud"`
		Email   string `json:"email"`
		Name    string `json:"name"`
		Picture string `json:"picture"`
	}

	// UserInfo ...
	UserInfo struct {
		ID      string `json:"id"`
		Email   string `json:"email"`
		Name    string `json:"name"`
		Picture string `json:"picture"` // Default size 96x96, but can change "=s96" in value to "=s200" to increase width, height to 200
	}
)

var endpointToken = "https://oauth2.googleapis.com/tokeninfo?id_token=%s"

// GetUserInfo ...
func GetUserInfo(token string) (user UserInfo, err error) {
	result, err := getDataFromToken(token)
	if err != nil {
		err = errors.New("cannot_get_user")
		return
	}

	// Compare AUD
	if result.AUD != config.ClientID {
		err = errors.New("not_valid_client_id")
		return
	}

	user = UserInfo{
		ID:      result.ID,
		Email:   result.Email,
		Name:    result.Name,
		Picture: result.Picture,
	}
	return
}

// Get user data from google token
func getDataFromToken(token string) (result Result, err error) {
	url := fmt.Sprintf(endpointToken, token)
	response, err := http.Get(url)
	if err != nil {
		return
	}
	defer response.Body.Close()

	// Read body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	// Parse to result
	err = json.Unmarshal(body, &result)
	if err != nil {
		return
	}

	return
}
