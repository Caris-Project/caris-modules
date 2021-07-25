package apple

import (
	"context"
	"errors"

	"github.com/Timothylock/go-signin-with-apple/apple"
)

// Config ...
type Config struct {
	Secret   string
	TeamID   string
	ClientID string
	KeyID    string
}

// UserInfo ...
type UserInfo struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

// GetUserInfo ...
func GetUserInfo(ctx context.Context, cfg Config, code string) (*UserInfo, error) {
	secret, err := apple.GenerateClientSecret(cfg.Secret, cfg.TeamID, cfg.ClientID, cfg.KeyID)
	if err != nil {
		return nil, err
	}
	client := apple.New()
	vReq := apple.AppValidationTokenRequest{
		ClientID:     cfg.ClientID,
		ClientSecret: secret,
		Code:         code,
	}

	var res apple.ValidationResponse

	// Do the verification
	err = client.VerifyAppToken(ctx, vReq, &res)
	if err != nil {
		return nil, err
	}
	if res.Error != "" {
		return nil, errors.New(res.Error)
	}
	userID, err := apple.GetUniqueID(res.IDToken)
	if err != nil {
		return nil, err
	}
	claim, _ := apple.GetClaims(res.IDToken)
	email, _ := claim.Get("email")

	return &UserInfo{
		ID:    userID,
		Email: email.(string),
	}, nil
}
