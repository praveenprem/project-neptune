package configuration

import (
	"encoding/json"
	"fmt"
	"github.com/praveenprem/logging"
	"github.com/praveenprem/nexus-auth/notification"
	"github.com/praveenprem/nexus-auth/providers"
	"os"
)

/**
 * Package name: configuration
 * Project name: ProjectNeptune
 * Created by: Praveen Premaratne
 * Created on: 22/03/2020 18:10
 */

type (
	// System defines the structure of the base application configuration
	// without any specific authentication backend configuration
	System struct {
		Path         string                     `json:"-"`            // Path defines the root directory that holds the configuration
		Provider     providers.Provider         `json:"provider"`     // Provider defines the authentication source
		SudoUser     string                     `json:"admin_user"`   // SudoUser define the system user with privileges
		User         string                     `json:"user"`         // KeyChain defiles the system user without privileges
		Notification *notification.Notification `json:"notification"` // Notification defines the configuration for notification service
		Host         string                     `json:"host"`         // Host defines unique name of the host for notification purposes
	}
	loader interface {
		ConfigInit() error
		ReadConfig() error
		isExist() error
		setup() error
		write(filename string) error
		read(filename string) error
	}
)

// ConfigInit creates a JSON based configuration schema on filesystem.
// It returns an error on failure of nil on success when creating configuration.
func (s *System) ConfigInit() error {
	if err := isExist(s.Path); err != nil {
		if err := s.setup(); err != nil {
			return err
		}
	}

	if err := isExist(fmt.Sprintf("%s/%s", s.Path, "config.json")); err != nil {
		return s.write("config.json")
	} else {
		return err
	}
}

// ReadConfig loads the configuration from the local filesystem.
// It returns a System object and an error on failure to load or nil on success.
func (s *System) ReadConfig() error {
	logging.Info("loading configuration")
	if err := isExist(s.Path); err != nil {
		return err
	} else {
		return s.read("config.json")
	}
}

// isExist validates if the configuration file exist.
// It returns an error if the doesn't exist or nil if it does.
func isExist(path string) error {
	file, err := os.Stat(path)
	if err != nil {
		return err
	}
	logging.Info(fmt.Sprintf(" %s found", file.Name()))
	return nil
}

func (s *System) setup() error {
	logging.Info(fmt.Sprintf("creating directory %s", s.Path))
	return os.Mkdir(s.Path, 0655)
}

func (s *System) write(filename string) error {
	logging.Info("creating configuration template")
	if outputFile, err := os.OpenFile(fmt.Sprintf("%s/%s",
		s.Path, filename), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0655); err != nil {
		return err
	} else {
		defer outputFile.Close()
		if rawData, err := json.MarshalIndent(s, "", "\t"); err != nil {
			return err
		} else {
			if _, err := outputFile.Write(rawData); err != nil {
				return err
			} else {
				logging.Info("configuration template created")
				return nil
			}
		}

	}
}

func (s *System) read(filename string) error {
	if configFile, err := os.OpenFile(fmt.Sprintf("%s/%s",
		s.Path, filename), os.O_RDONLY, 0655); err != nil {
		return err
	} else {
		if err := json.NewDecoder(configFile).Decode(&s); err != nil {
			return err
		} else {
			logging.Info("configuration loaded")
			return nil
		}
	}
}
