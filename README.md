# NexusAuth

<img width="100" src="resources/Icon.png" alt="app icon">
NexusAuth is an SSH authentication plugin written for Unix based systems which allow to source public keys from an external source, extending from rudimentary static file-based authentication.

## Table of Contents
- [Getting Started](#getting-started)
- [Pre-requisites](#pre-requisites)
- [Installation](#installation)
    - [Configuration](#configuration)
    - [Sample](#sample)
    - [Integration with system SSH](#integration-with-system-ssh)
    - [Test installation](#test-installation)
- [Alerting](#alerting)
- [Logging](#logging)
- [Logging Example](#logging-example)
- [Development](#development)
- [License](#license)
- [Authors](#authors)
- [Acknowledgments](#acknowledgments)

## Getting Started
This plugin will allow you to configure one of the following services as an authentication mechanism:
- [GitHub](https://github.com/apps/sshauth)
- AWS - _In development_
- MySQL - _In development_

>This is **not** a replacement for the `authrorized_keys` file. It highly encouraged to still populate, few permanent public keys for failsafe purposes.

Following instruction will guide you through how to install and configure this plugin on a Unix/Linux server.

## Pre-requisites
- Root access to the server in question
- Administrative access to one of third-party services listed above
- Uninterrupted connection to the server in questions, especially when making changes to the SSH daemon.

## Installation
- Download the latest release version of the [Nexus-auth](https://github.com/praveenprem/nexus-auth/releases/latest) for your system
- Unzip the archive to current directory. `unzip nexus-auth-*-amd64.zip`
- Give execution permission. `chmod +x nexus-auth`
- Copy the binary file to system path. `install $PWD/nexus-auth /usr/local/bin/nexus-auth`
- Initialise the application configuration. `nexus-auth -init`

### Configuration
Configuration can be located in the `/etc/nexusauth/config.json`. Following is the structure of the configuration file:
- **host** - Unique identifier of server
- **admin_user** - System user with administrator privilege
- **user** - System user _without_ administrator privilege
- **notification** - Notification service configuration. [Supported notification services](resources/Notification.md)
- **provider** - Configuration of authentication provider
    - **Name** - Name of the authentication provider. [Supported providers](resources/Providers.md)
    - **Configuration** - Configuration of the named authentication provider.

#### Sample
```json
{
    "host": "",
    "provider": {
        "Configuration": null,
        "Name": ""
    },
    "admin_user": "",
    "user": "",
    "notification": null
}
```

### Integration with system SSH
Update the /etc/ssh/sshd_config to reflect the following changes:
- `AuthorizedKeysCommand /usr/local/bin/nexus-auth -u %u -k %f`
- `AuthorizedKeysCommandUser root`. This differs on which user owns the **nexus-auth** binary execution file.

Apply the changes made to the SSH daemon using system specific command. I.E. `service ssh restart` for Ubuntu.
It is recommended to test the installation before applying these changes.

### Test installation
For testing the installation and configuration, run the following command:

```
nexus-auth -u <USERNAME> -k "<PUBLIC KEY Fingerprint>"
```

- `<USERNAME>` is the admin or default user of the server, defined under `admin_user` or `user`
- `<PUBLIC KEY>` is a public key from any user who's public key can be retrieved from the third-party service.
    In order to mimic the SSH daemon place the public key with in `" "` and exclude the trailing comment of the key.
    
```bash
sshauth -u ubuntu SHA256:OZvuPKD7k9uS15jeV3HilpDXutQRPrGct2UWhQDRLQA
```

## Alerting

Alert will provide the following information:
   - **User** - User name of the third-party service public key matched
   - **Provider** - Name of the third-party service has been used to authenticate the user
   - **Host** - Host name defined in the configuration file
   - **System User** - System username used for authentication

#### Provider specific notification can be found in the [Notification Service configuration](resources/Notification.md) documentation.

## Logging
This plugin will log for informative and debugging purposes, such as bad configuration.
These logs can be found in `/var/log/nexusauth.log` when the plugin runs as `root`.

Logging is achieve using a third-party library [logging](https://github.com/praveenprem/logging).
### Logging Example
```
2020/04/12 14:20:45 info: =============== starting authentication ===============
2020/04/12 14:20:45 info: input verification successful
2020/04/12 14:20:45 info: loading config from system
2020/04/12 14:20:45 info: loading configuration
2020/04/12 14:20:45 info:  nexusauth found
2020/04/12 14:20:45 info: configuration loaded
2020/04/12 14:20:45 info: validating provider configuration
2020/04/12 14:20:45 info: GitHub provider detected
2020/04/12 14:20:45 info: loading GitHub configuration
2020/04/12 14:20:45 info: configuration successfully loaded
2020/04/12 14:20:45 info: validating provider config
2020/04/12 14:20:45 info: provider config validated
2020/04/12 14:20:45 info: starting application authentication
2020/04/12 14:20:45 info: creating JWT token
2020/04/12 14:20:45 info: signing JWT token
2020/04/12 14:20:45 info: application authentication token created
2020/04/12 14:20:45 info: POST: https://api.github.com/app/installations/000000/access_tokens
2020/04/12 14:20:45 info: decoding response
2020/04/12 14:20:45 info: decode completed
2020/04/12 14:20:45 info: application successfully authenticated
2020/04/12 14:20:45 info: fetching members of team devops
2020/04/12 14:20:45 info: GET: https://api.github.com/orgs/:org/teams/:team/members?role=all
2020/04/12 14:20:46 info: decoding response
2020/04/12 14:20:46 info: decode completed
2020/04/12 14:20:46 info: fetching keys of praveenprem
2020/04/12 14:20:46 info: GET: https://api.github.com/users/praveenprem/keys
2020/04/12 14:20:46 info: decoding response
2020/04/12 14:20:46 info: decode completed
2020/04/12 14:20:47 info: json marshalling provider result
2020/04/12 14:20:47 info: decoding provider results
2020/04/12 14:20:47 info: decoding completed
2020/04/12 14:20:47 info: sending notification to Slack
2020/04/12 14:20:47 info: notification sent
2020/04/12 14:20:47 info: =============== authentication process completed ===============
```


## Development
TODO:

## License
```
MIT License

Copyright (c) 2020 Praveen Premaratne

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```

## Authors
   | <div><a href="https://github.com/praveenprem"><img width="200" src="https://avatars3.githubusercontent.com/u/23165760"/><p></p><p>Praveen Premaratne</p></a></div> |
   | :-------: |
    
## Acknowledgments
[SSHAUTH](https://github.com/apps/sshauth) icon is composed of:
- Icons made by [Freepik](https://www.flaticon.com/authors/freepik) from [www.flaticon.com](https://www.flaticon.com/)
- Icons made by [Those Icons](https://www.flaticon.com/authors/those-icons) from [www.flaticon.com](https://www.flaticon.com/)
- Icons made by [DinosoftLabs](https://www.flaticon.com/authors/dinosoftlabs) from [www.flaticon.com](https://www.flaticon.com/)
