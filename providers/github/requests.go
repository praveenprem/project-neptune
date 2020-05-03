package github

import (
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/praveenprem/logging"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

/**
 * Package name: github
 * Project name: ProjectNeptune
 * Created by: Praveen Premaratne
 * Created on: 29/03/2020 18:17
 */

type (
	//Claims defines the override structure for the jwt.StandardClaims
	Claims struct {
		jwt.StandardClaims
		Iss int `json:"iss"`
	}
)

//HttpCall makes Http request for the given request and maps the response to given response pinter.
//
//request defines the pointer for  http.Request
//response defines the generic interface pointer to map the response
//statusCode defines the http status code to for successful response
func (c *Configuration) HttpCall(request *http.Request, response interface{}, statusCode int) error {
	var client = &http.Client{}
	request.Header.Add("Accept", c.MediaType)
	if c.Token != "" {
		//logging.Debug("access token found. adding token header")
		//logging.Debug(fmt.Sprintf("Bearer %s", c.Token))
		request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.Token))
	}
	logging.Info(fmt.Sprintf("%s: %s", request.Method, request.URL.String()))
	if resp, err := client.Do(request); err != nil {
		return err
	} else {
		defer resp.Body.Close()

		if resp.StatusCode != statusCode {
			body, _ := ioutil.ReadAll(resp.Body)
			logging.Warning(string(body))
			return err
		}
		logging.Info("decoding response")
		if err := json.NewDecoder(resp.Body).Decode(response); err != nil {
			return err
		}

	}
	logging.Info("decode completed")
	return nil
}

//CreateToken creates a signed JWT token that is used to request the access token
func (c *Claims) CreateToken() (string, error) {
	logging.Info("creating JWT token")
	var now = time.Now().Unix()
	var token = jwt.New(jwt.SigningMethodRS256)

	c.IssuedAt = now
	c.ExpiresAt = now + 600 // 10 minutes expire time
	c.Iss = appId

	token.Claims = c

	if key, err := c.Sign(); err != nil {
		return "", err
	} else {
		return token.SignedString(key)
	}
}

//Sign signs the given token with specified private key
func (c *Claims) Sign() (*rsa.PrivateKey, error) {
	logging.Info("signing JWT token")
	return jwt.ParseRSAPrivateKeyFromPEM(loadPrivateKey())
}

//NewHttpRequest creates a http request for the given parameters
// error out on failure
func NewHttpRequest(method string, url string, body io.Reader) *http.Request {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		logging.Error(err.Error())
	}
	return req
}
