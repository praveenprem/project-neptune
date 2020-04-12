package notification

import (
	"errors"
	"github.com/praveenprem/project-neptune/codes"
	"github.com/praveenprem/project-neptune/notification/slack"
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
