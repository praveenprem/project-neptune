/*
Package configuration implements access to configuration stored on filesystem

Using configuration, it is possible to generate JSON schema for the new installation or
load existing configuration from the filesystem

Example

	package main

	import (
		"flag"
		"fmt"
		"github.com/praveenprem/ssh-auth/configuration"
		"github.com/praveenprem/ssh-auth/providers"
		"os"
	)

	func main() {
		var app App
		var config = configuration.System{
			Path:     configPath,
			Provider: providers.Provider{},
			SudoUser: "",
			User:     "",
		}
		if app.ConfigInit {
			if err := config.ConfigInit(); err != nil {
				panic(err.Error())
			}
		}
	}
*/
package configuration
