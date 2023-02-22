package providers

import (
	"encoding/json"
	"errors"
	"github.com/praveenprem/logging"
	"github.com/razorcorp/nexus-auth/codes"
	"github.com/razorcorp/nexus-auth/providers/github"
)

/**
 * Package name: providers
 * Project name: ProjectNeptune
 * Created by: Praveen Premaratne
 * Created on: 25/03/2020 20:09
 */

type (
	//Provider defines the structure for provider specific configuration.
	Provider struct {
		Name          string
		Configuration interface{}
	}
	KeyChain struct {
		Username string   `json:"username"`
		Keys     []string `json:"keys"`
	}
	provider interface {
		Authenticate(username string) (users []KeyChain, err error)
		valid()
	}
)

// Authenticate defines the logic used to determine the authentication provider
// and translation of the data between main application and provider
func (p *Provider) Authenticate(username string) (users []KeyChain, err error) {
	var keyChain []KeyChain

	p.valid()
	switch p.Name {
	case "github":
		logging.Info("GitHub provider detected")
		if k, err := github.Call(username, p.Configuration); err != nil {
			logging.Info(err.Error())
		} else {
			logging.Info("decoding provider results")
			err := json.Unmarshal([]byte(k), &keyChain)
			if err != nil {
				logging.Error(err.Error())
			}
			logging.Info("decoding completed")
		}
	default:
		logging.Warning("Invalid provider given")
		return keyChain, errors.New(codes.CODE3)
	}
	return keyChain, nil
}

// valid define logic to validate if the provider configuration
func (p *Provider) valid() {
	logging.Info("validating provider configuration")
	if p.Name != "" || p.Configuration != nil {
		return
	}
	panic("provider name or configuration not found")
}
