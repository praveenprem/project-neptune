package github

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/praveenprem/logging"
	"github.com/praveenprem/nexus-auth/codes"
	"io/ioutil"
	"os"
)

/**
 * Package name: github
 * Project name: ProjectNeptune
 * Created by: Praveen Premaratne
 * Created on: 29/03/2020 18:14
 */

type (
	//Configuration defines the structure for the configuration needed to complete a query to Github API.
	Configuration struct {
		InstallationId int    `json:"installation_id"`
		AdminRole      string `json:"admin_role"`
		ApiUrl         string `json:"api_url"`
		DefaultRole    string `json:"default_role"`
		Org            string `json:"org"`
		TeamName       string `json:"team_name"`
		MediaType      string `json:"media_type"`
		Token          string `json:"-"`
	}
)

func (c *Configuration) configCast(rawConf interface{}) {
	logging.Info("loading GitHub configuration")
	jsonData, mErr := json.Marshal(rawConf)
	if mErr != nil {
		logging.Error(mErr.Error())
	}
	uErr := json.Unmarshal(jsonData, c)
	if uErr != nil {
		logging.Error(uErr.Error())
	}
	logging.Info("configuration successfully loaded")
}

func (c *Configuration) validate() error {
	logging.Info("validating provider config")
	if c.InstallationId == 0 ||
		c.AdminRole == "" ||
		c.DefaultRole == "" ||
		c.ApiUrl == "" ||
		c.Org == "" ||
		c.MediaType == "" {
		return errors.New(codes.CODE4)
	}

	if c.TeamName == "" {
		logging.Warning(codes.CODE9)
	}
	logging.Info("provider config validated")
	return nil
}

func (c *Configuration) getRole(username string) string {
	switch username {
	case c.AdminRole:
		return c.AdminRole
	default:
		return c.DefaultRole
	}
}

func loadPrivateKey() []byte {
	var PrivateKey = os.Getenv("githubPem")
	if PrivateKey == "" {
		panic(codes.CODE8)
	}
	if pKey, err := ioutil.ReadFile(PrivateKey); err != nil {
		panic(fmt.Sprintf("build failed: %s", err.Error()))
	} else {
		return pKey
	}
	return nil
}
