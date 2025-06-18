package spotify

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type SpotifyTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

func (o *outbound) GetTokenDetails() (string, string, error) {
	if o.AccessToken == "" || time.Now().After(o.ExpiredAt) {
		// call spotify api token here
		err := o.generateToken()
		if err != nil {
			return "", "", err
		}
	}
	return o.AccessToken, o.TokenType, nil
}

func (o *outbound) generateToken() error {
	formData := url.Values{}
	formData.Set("grant_type", "client_credentials")
	formData.Set("client_id", o.cfg.SpotifyConfig.ClientID)
	formData.Set("client_secret", o.cfg.SpotifyConfig.ClientSecret)

	encodedURL := formData.Encode()
	request, err := http.NewRequest(http.MethodPost, `https://accounts.spotify.com/api/token`, strings.NewReader(encodedURL))
	if err != nil {
		log.Println("error create request for spotify")
		return err
	}

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := o.client.Do(request)
	if err != nil {
		log.Println("error execute request for spotify")
		return err
	}
	defer resp.Body.Close()

	var response SpotifyTokenResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		log.Println("error unmarshal response from spotify")
		return err
	}

	o.AccessToken = response.AccessToken
	o.TokenType = response.TokenType
	o.ExpiredAt = time.Now().Add(time.Duration(response.ExpiresIn) * time.Second)

	return nil
}