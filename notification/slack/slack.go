package slack

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/praveenprem/logging"
	"github.com/razorcorp/nexus-auth/codes"
	"io/ioutil"
	"net/http"
)

/**
 * Package name: slack
 * Project name: ProjectNeptune
 * Created by: Praveen Premaratne
 * Created on: 10/04/2020 22:05
 */

type (
	Slack struct {
		User       string
		Host       string
		SystemUser string
		Provider   string
		unknown    bool
	}

	Blocks struct {
		Blocks []Block `json:"blocks"`
	}

	Block struct {
		Type     string  `json:"type,omitempty"`
		Text     *Text   `json:"text,omitempty"`
		Elements *[]Text `json:"elements,omitempty"`
		*Fields
	}

	Fields struct {
		Fields    *[]Text    `json:"fields,omitempty"`
		Accessory *Accessory `json:"accessory,omitempty"`
	}

	Text struct {
		Type string `json:"type,omitempty"`
		Text string `json:"text,omitempty"`
	}

	Accessory struct {
		Type     string `json:"type,omitempty"`
		ImageUrl string `json:"image_url,omitempty"`
		AltText  string `json:"alt_text,omitempty"`
	}
)

// SlackSend defines the implementation of Slack notification
func (s *Slack) SlackSend(url string) error {
	logging.Info("sending notification to Slack")
	s.validate()
	if url == "" {
		return errors.New(codes.CODE7)
	}
	var (
		payload Blocks
		app     = Block{
			Type: "context",
			Elements: &[]Text{{
				Type: "mrkdwn",
				Text: "Plugin: <https://github.com/apps/sshauth|SSHAUTH>",
			}},
		}
		conn = Block{
			Type: "section",
			Text: &Text{
				Type: "mrkdwn",
				Text: "*New SSH connection detected*",
			},
		}
		unknownUser = Block{
			Type: "section",
			Text: &Text{
				Type: "mrkdwn",
				Text: fmt.Sprintf(
					"_*`Login attempt detected with system user: %s that was NOT authenticated via SSHAUTH plugin.`*_",
					s.SystemUser),
			},
		}
		userData = Block{
			Type: "section",
			Fields: &Fields{
				Fields: &[]Text{
					{
						Type: "mrkdwn",
						Text: fmt.Sprintf("*User:* \n _%s_", s.User),
					},
					{
						Type: "mrkdwn",
						Text: fmt.Sprintf("*Host:* \n _%s_", s.Host),
					},
					{
						Type: "mrkdwn",
						Text: fmt.Sprintf("*Provider:* \n _%s_", s.Provider),
					},
					{
						Type: "mrkdwn",
						Text: fmt.Sprintf("*System User:* \n _%s_", s.SystemUser),
					},
				},
				Accessory: &Accessory{
					Type:     "image",
					ImageUrl: "https://raw.githubusercontent.com/praveenprem/nexus-auth/develop/resources/Icon.png",
					AltText:  "<https://github.com/apps/sshauth|SSHAUTH>",
				},
			},
		}
		divider = Block{
			Type: "divider",
		}
	)
	payload.Blocks = append(payload.Blocks, divider)
	payload.Blocks = append(payload.Blocks, conn)
	if s.unknown {
		payload.Blocks = append(payload.Blocks, unknownUser)
	}
	payload.Blocks = append(payload.Blocks, userData)
	payload.Blocks = append(payload.Blocks, app)
	payload.Blocks = append(payload.Blocks, divider)

	return payload.post(url)
}

func (b *Blocks) post(url string) error {
	var client = &http.Client{}

	body := bytes.NewBuffer([]byte{})

	enc := json.NewEncoder(body)
	enc.SetEscapeHTML(false)
	if err := enc.Encode(b); err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return err
	}

	req.Header.Add("Content-type", "application/json")

	if resp, err := client.Do(req); err != nil {
		return err
	} else {
		defer resp.Body.Close()

		if resp.StatusCode != 200 {
			body, _ := ioutil.ReadAll(resp.Body)
			return errors.New(string(body))
		}
	}

	logging.Info("notification sent")

	return nil
}

func (s *Slack) validate() {
	if s.User == "" {
		s.User = "`Unknown`"
		s.unknown = true
	}
	if s.Provider == "" {
		s.unknown = true
		s.Provider = "`Unknown`"
	}
}
