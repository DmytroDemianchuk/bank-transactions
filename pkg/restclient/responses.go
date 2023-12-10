package restclient

import (
	"encoding/json"
	"errors"
)

type ResponseToken struct {
	Token      string `json:"access_token" validate:"required"`
	Expires_in int    `json:"expires_in" validate:"required"`
	Token_type string `json:"token_type" validate:"required"`
}

func respGetToken(body []byte) (string, error) {
	var rt ResponseToken
	if err := json.Unmarshal(body, &rt); err != nil || rt.Token == "" || rt.Token_type != "bearer" {

		return "", errors.New("api decode error")
	}
	return rt.Token, nil
}
