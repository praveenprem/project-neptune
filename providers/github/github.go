package github

import (
	"encoding/json"
	"fmt"
	"github.com/praveenprem/logging"
	"net/http"
)

/**
 * Package name: github
 * Project name: ProjectNeptune
 * Created by: Praveen Premaratne
 * Created on: 25/03/2020 20:33
 */

type (
	//Key defines the structure for the response of Github when requesting user's keys
	Key struct {
		Id  int    `json:"id"`
		Key string `json:"key"`
	}

	//Team defines the structure for the response of Github when requesting teams in an organisation.
	Team struct {
		Name            string `json:"name"`
		Id              int    `json:"id"`
		NodeId          string `json:"node_id"`
		Slug            string `json:"slug"`
		Description     string `json:"description"`
		Privacy         string `json:"privacy"`
		Url             string `json:"url"`
		MembersUrl      string `json:"members_url"`
		RepositoriesUrl string `json:"repositories_url"`
		Permission      string `json:"permission"`
	}

	//User defines the structure for the response of Github when requesting user.
	User struct {
		Login  string `json:"login"`
		Id     int    `json:"id"`
		NodeId string `json:"node_id"`
		//AvatarUrl         string `json:"avatar_url"`
		//GravatarId        string `json:"gravatar_id"`
		Url               string `json:"url"`
		HtmlUrl           string `json:"html_url"`
		FollowersUrl      string `json:"followers_url"`
		FollowingUrl      string `json:"following_url"`
		GistsUrl          string `json:"gists_url"`
		StarredUrl        string `json:"starred_url"`
		SubscriptionsUrl  string `json:"subscriptions_url"`
		OrganizationsUrl  string `json:"organizations_url"`
		ReposUrl          string `json:"repos_url"`
		EventsUrl         string `json:"events_url"`
		ReceivedEventsUrl string `json:"received_events_url"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
	}

	//AccessToken defines the structure of Github response to access token request.
	AccessToken struct {
		Token       string      `json:"token"`
		Expire      string      `json:"expires_at"`
		Permissions interface{} `json:"permissions"`
	}

	KeyChain struct {
		Username string   `json:"username"`
		Keys     []string `json:"keys"`
	}
)

const (
	appId           = 59005
	appAccTokenPath = "%s/app/installations/%d/access_tokens"
	membersPath     = "%s/orgs/%s/members?role=%s"
	teamPath        = "%s/orgs/%s/teams/%s"
	teamMembersPath = "%s/orgs/%s/teams/%s/members?role=%s"
	keyPath         = "%s/keys"
)

//Call is a package entrypoint for underlying application.
//username defines the system user used for access.
//configuration defines the generic interface to pass the provider configuration from underlying application.
func Call(username string, configuration interface{}) (users string, err error) {
	var keyChain []KeyChain
	var config Configuration
	var members []User

	config.configCast(configuration)
	if err := config.validate(); err != nil {
		logging.Error(err.Error())
	}

	var token AccessToken
	token.AccessToken(config)
	if token.Token != "" {
		config.Token = token.Token
	}

	if config.TeamName != "" {
		members = *token.TeamsMembers(config, username)
	} else {
		members = *token.OrgMembers(config, username)
	}

	for _, user := range members {
		var userKeyChain KeyChain
		userKeyChain.Username = user.Login
		userKeys := *user.UserKeys(config)
		for _, key := range userKeys {
			userKeyChain.Keys = append(userKeyChain.Keys, key.Key)
		}
		keyChain = append(keyChain, userKeyChain)
	}

	return jsonKeyChain(keyChain), nil
}

//AccessToken requests application access token from Github for authentication
//of the other calls
func (t *AccessToken) AccessToken(c Configuration) {
	logging.Info("starting application authentication")
	var claims Claims

	appAuthToken, err := claims.CreateToken()
	if err != nil {
		logging.Error(err.Error())
	}
	logging.Info("application authentication token created")

	req := NewHttpRequest(http.MethodPost, fmt.Sprintf(appAccTokenPath, c.ApiUrl, c.InstallationId), nil)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", appAuthToken))

	if err := c.HttpCall(req, &t, 201); err != nil {
		logging.Error(err.Error())
	}
	logging.Info("application successfully authenticated")
}

//TeamsMembers gets a list of members of the given team according to the role
func (t *AccessToken) TeamsMembers(c Configuration, username string) *[]User {
	logging.Info(fmt.Sprintf("fetching members of team %s", c.TeamName))
	var users []User

	req := NewHttpRequest(http.MethodGet, fmt.Sprintf(teamMembersPath, c.ApiUrl, c.Org, c.TeamName, c.getRole(username)), nil)
	if err := c.HttpCall(req, &users, 200); err != nil {
		logging.Warning(err.Error())
	}

	return &users
}

//OrgMembers gets a list of members of the organisation according to the role
func (t AccessToken) OrgMembers(c Configuration, username string) *[]User {
	logging.Info(fmt.Sprintf("fetching members organisation %s", c.Org))
	var users []User
	req := NewHttpRequest(http.MethodGet, fmt.Sprintf(membersPath, c.ApiUrl, c.Org, c.getRole(username)), nil)

	if err := c.HttpCall(req, &users, 200); err != nil {
		logging.Warning(err.Error())
	}
	logging.Info(fmt.Sprintf("%d organisation members found", len(users)))
	return &users
}

//UserKeys gets keys on the user accounts given to the method
func (u *User) UserKeys(c Configuration) *[]Key {
	logging.Info(fmt.Sprintf("fetching keys of %s", u.Login))
	var keys []Key
	req := NewHttpRequest(http.MethodGet, fmt.Sprintf(keyPath, u.Url), nil)
	if err := c.HttpCall(req, &keys, 200); err != nil {
		logging.Warning(err.Error())
	}
	return &keys
}

func jsonKeyChain(k []KeyChain) string {
	logging.Info("json marshalling provider result")
	rawData, err := json.Marshal(k)
	if err != nil {
		logging.Error(err.Error())
	}
	return string(rawData)
}

//GetTeam gets list of teams in the given organisation
func (t *Team) GetTeam(c Configuration) error {
	req := NewHttpRequest(http.MethodGet, fmt.Sprintf(teamPath, c.ApiUrl, c.Org, c.TeamName), nil)
	if err := c.HttpCall(req, &t, 200); err != nil {
		return err
	}
	return nil

}
