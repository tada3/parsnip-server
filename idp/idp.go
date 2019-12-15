package idp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"time"
)

const (
	idpURL  = "https://account.line-beta.biz/"
	rootURL = "https://console.brain.line-beta.biz/"
)

type Client struct {
	httpClient *http.Client
	baseURL    *url.URL
	token      string
}

type UserResponse struct {
	BusinessID     string `json:"businessId"`
	Language       string `json:"language"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	MID            string `json:"mid"`
	LinePictureURI string `json:"linePictureUri"`
	LineRegionCode string `json:"lineRegionCode"`
}

type SessionResponse struct {
	Status           string       `json:"status"`
	User             UserResponse `json:"user"`
	ErrorDescription string       `json:"errorDescription"`
}

type LogoutURIResponse struct {
	Status           string `json:"status"`
	LogoutURI        string `json:"logoutUri"`
	ErrorDescription string `json:"errorDescription"`
}

func New(baseURL, token string, timeout int) (*Client, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}
	hc := &http.Client{Timeout: time.Duration(timeout) * time.Millisecond}
	return &Client{
		httpClient: hc,
		baseURL:    u,
		token:      token,
	}, nil
}

func (c *Client) GetSession(sessionID string) (*SessionResponse, error) {
	fmt.Println("Call IDP /v1/getSession")

	u := *c.baseURL
	u.Path = path.Join(c.baseURL.Path, "/v1/getSession")

	q := u.Query()
	q.Set("sessionId", sessionID)
	u.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+c.token)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("/v1/getSessin failed: %+v", err)
	}

	sr := new(SessionResponse)
	if err = decodeBody(resp, sr); err != nil {
		return nil, err
	}

	return sr, nil
}

//  Logout returns redirect uri from IDP
func (c *Client) GenerateLogoutURI(sessionID, redirectURI string) (*LogoutURIResponse, error) {

	u := *c.baseURL
	u.Path = path.Join(c.baseURL.Path, "/v1/generateLogoutUri")

	q := u.Query()
	q.Set("sessionId", sessionID)
	q.Set("redirectUri", redirectURI)
	u.RawQuery = q.Encode()

	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+c.token)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("/v1/generateLogoutUri failed: %+v", err)
	}

	logoutURIResponse := &LogoutURIResponse{}
	err = decodeBody(resp, logoutURIResponse)
	if err != nil {
		return nil, err
	}

	//key := i.cacheKey(sessionID)
	//err = i.redisClient.Del(key)
	//if err != nil {
	//	return nil, fmt.Errorf("Unable to delete Redis's key: %+v with error: %+v and", key, err)
	//}

	return logoutURIResponse, nil
}

func GenerateLoginURI() string {

	u, err := url.Parse(idpURL)
	if err != nil {
		panic(err)
	}

	u.Path = path.Join(u.Path, "login")

	q := u.Query()
	q.Set("redirectUri", rootURL)
	u.RawQuery = q.Encode()
	return u.String()
}

func decodeBody(resp *http.Response, out interface{}) error {
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Invalid HTTP Status: %s", resp.Status)
	}
	decoder := json.NewDecoder(resp.Body)
	err := decoder.Decode(out)
	if err != nil {
		return fmt.Errorf("Failed to parse response body: %v", err)
	}
	return nil
}
