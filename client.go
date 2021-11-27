package dropshipping

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func NewClient(email, password string, isDebug bool) (*Client, error) {

	if email == "" || password == "" {
		return nil, errors.New("email or password has null")
	}

	client := &Client{
		Email:    email,
		Http:     &http.Client{},
		Password: password,
	}

	if isDebug {
		client.SetLogger(log.Writer())
	}

	return client, nil
}

// GetAccessToken
// Endpoint /api2.0/v1/authentication/getAccessToken
func (c *Client) GetAccessToken(ctx context.Context) (*AccessTokenResponse, error) {

	url := fmt.Sprintf("%s%s", ApiBaseUrl, UriGetAccessToken)
	payload := strings.NewReader(fmt.Sprintf(`{"email": "%s","password": "%s"}`, c.Email, c.Password))
	request, err := http.NewRequestWithContext(ctx, "POST", url, payload)
	if err != nil {
		return nil, err
	}
	request.Header.Add("content-type", "application/json")

	response := &AccessTokenResponse{}
	err = c.Send(request, response)
	if err != nil {
		return nil, err
	}
	if response.Data.AccessToken != "" {
		c.AccessToken = response
		c.AccessTokenExpiryAt = response.Data.AccessTokenExpiryDate
	}

	return response, nil
}

func (c *Client) SendWithAuth(req *http.Request, v interface{}) error {

	c.Lock()

	if c.AccessToken != nil {
		// TODO 计算 token 失效时间
		req.Header.Add("CJ-Access-Token", c.AccessToken.Data.AccessToken)
	}
	defer c.Unlock()

	return c.Send(req, v)
}

func (c *Client) Send(req *http.Request, v interface{}) error {

	var err error
	var rep *http.Response

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Language", "en_US")
	if req.Header.Get("Content-type") == "" {
		req.Header.Set("Content-type", "application/json")
	}

	rep, err = c.Http.Do(req)
	c.log(req, rep)
	defer rep.Body.Close()

	if err != nil {
		return err
	}

	if v == nil {
		return nil
	}

	//fmt.Println("vvvv", rep.Body)
	err = json.NewDecoder(rep.Body).Decode(v)

	return err
}
