package facebook

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// FailedResponse ...
type FailedResponse struct {
	Error struct {
		Message      string `json:"message"`
		Type         string `json:"type"`
		Code         int    `json:"code"`
		ErrorSubCode int    `json:"error_subcode"`
		FbTraceID    string `json:"fbtrace_id"`
	} `json:"error"`
}

type (
	// UserInfo ...
	UserInfo struct {
		ID      string          `json:"id"`
		Name    string          `json:"name"`
		Email   string          `json:"email"`
		Picture UserInfoPicture `json:"picture"`
	}

	// UserInfoPicture ...
	UserInfoPicture struct {
		Data UserInfoPictureData `json:"data"`
	}

	// UserInfoPictureData ...
	UserInfoPictureData struct {
		Height int    `json:"height"`
		URL    string `json:"url"`
		Width  int    `json:"width"`
	}
)

// Replace facebook endpoint, due to old  version
var endpointMe = "https://graph.facebook.com/me?fields=name,picture,email&access_token=%s"

// var endpointAvatar = "https://graph.facebook.com/%s/picture?height=200&access_token=%s"
var client http.Client

// GetUserInfo return user data from token
func GetUserInfo(token string) (result UserInfo, err error) {
	// Call api for user info
	meURL := fmt.Sprintf(endpointMe, token)
	userInfo, err := callMe(meURL)

	// Return if not found or error
	if userInfo.ID == "" || err != nil {
		return
	}

	// TODO: call for larger picture
	// Call api for picture
	// Save and upload to s3

	// Return
	return userInfo, nil
}

// call api /me, for user data
func callMe(url string) (result UserInfo, err error) {
	resp, err := client.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	// Parse for error first
	var failedResponse FailedResponse
	err = json.Unmarshal(body, &failedResponse)
	// Return if has error code
	if failedResponse.Error.Code != 0 {
		err = errors.New(convertMeErrorCode(failedResponse.Error.Code))
		return
	}

	// Convert to response data
	err = json.Unmarshal(body, &result)
	return
}

// convertMeErrorCode ...
func convertMeErrorCode(code int) string {
	if code == 190 {
		return "invalid_access_token"
	} else if code == 200 {
		return "permission_error"
	}

	fmt.Println("*** Error (other) code when get user info:", code)

	return "cannot_get_user"
}
