package notification

import (
	"errors"
	"fmt"
	"github.com/praveenprem/logging"
	"github.com/razorcorp/nexus-auth/codes"
	"github.com/razorcorp/nexus-auth/notification/slack"
	"io/ioutil"
	"os"
)

/**
 * Package name: notification
 * Project name: ProjectNeptune
 * Created by: Praveen Premaratne
 * Created on: 10/04/2020 20:10
 */

type (
	Notification struct {
		Service string `json:"service"`
		Url     string `json:"url"`
	}

	Message struct {
		User       string
		Provider   string
		Host       string
		SystemUser string
	}

	Notifier interface {
		Notify() error
	}
)

// Notify defines the interface for different notification providers
func (n *Notification) Notify(m Message) error {
	if isLastUser(m.User) {
		logging.Warning("stopping duplicate alert")
		return nil
	}

	switch n.Service {
	case "slack":
		var s slack.Slack
		s.User = m.User
		s.Host = m.Host
		s.Provider = m.Provider
		s.SystemUser = m.SystemUser
		return s.SlackSend(n.Url)
	default:
		return errors.New(codes.CODE5)
	}
}

func isLastUser(username string) bool {
	logging.Info("checking last logged user")
	var sameUser = false
	if parityFile, err := os.OpenFile("/tmp/login.last", os.O_CREATE|os.O_RDONLY, 0655); err != nil {
		logging.Error(err.Error())
	} else {
		lastUser, readErr := ioutil.ReadAll(parityFile)
		if readErr != nil {
			logging.Error(readErr.Error())
		}

		logging.Info(fmt.Sprintf("last logged user: %s", lastUser))

		if string(lastUser) == username {
			sameUser = true
		}
	}

	if parityFile, err := os.OpenFile("/tmp/login.last", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0655); err != nil {
		logging.Error(err.Error())
	} else {
		if err := parityFile.Truncate(0); err != nil {
			logging.Warning(err.Error())
		}
		if _, err := parityFile.Seek(0, 0); err != nil {
			logging.Warning(err.Error())
		}

		if !sameUser {
			if _, err := parityFile.WriteString(username); err != nil {
				logging.Error(err.Error())
			}
		}
	}

	return sameUser
}
