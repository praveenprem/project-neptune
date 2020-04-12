# NexusAuth

<img width="100" src="resources/Icon.png" alt="app icon">
NexusAuth is an SSH authentication plugin written for Unix based systems which allow to source public keys from an external source, extending from rudimentary static file-based authentication.

This plugin will allow you to configure one of the following services as an authentication mechanism:
- [GitHub](https://github.com/apps/sshauth)
- AWS - _In development_
- MySQL - _In development_

>This is **not** a replacement for the `authrorized_keys` file. It highly encouraged to still populate, few permanent public keys for failsafe purposes.

## Table of Contents

## Getting Started
Following instruction will guide you through how to install and configure this plugin on a Unix/Linux server.

### Prerequisites
- Root access to the server in question
- Administrative access to one of third-party services listed above
- Uninterrupted connection to the server in questions, especially when making changes to the SSH daemon.

## Installation
TODO:

## Alerting
Current version support alerting with [Slack Incoming Webhooks](https://api.slack.com/incoming-webhooks). 

Alert will provide the following information:
   - **User** - User name of the third-party service that public key matched
   - **Provider** - Name of the third-party service that has been used to authenticate the user
   - **Host** - Host name defined in the configuration file
   - **System User** - System username used for authentication

### Sample notifications
All login attempts will have one of the following results.
#### Success
![slack-notification-success]
#### Failure
![slack-notification-fail]

## Logging
This plugin will log for informative and debugging purposes, such as bad configuration.
These logs can be found in `/var/log/nexusauth.log` when the plugin is run as `root`.

Logging is done using a third-party library [logging](https://github.com/praveenprem/logging).
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

[slack-notification-success]: resources/Slack-notification-success.png "Slack notification success"
[slack-notification-fail]: resources/Slack-notification-fail.png "Slack notification failure"
