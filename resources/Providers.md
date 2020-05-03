# Provider configuration
Below are a list of authentication service providers, currently supported by NexusAuth

> Keys `Name` and `Configuration` are the same keys also listed in the README.md.

- [Github](#github)
    - [Pre-Requisites](#pre-requisites)
    - [Configuration](#configuration)
        - [Sample](#sample)

## Github
### Pre-requisites
- Install the [SSHAUTH App](https://github.com/apps/sshauth)
- Note down the installation *ID* once installed. This can be found in the URL of the *settings/installations/{ID}*
    > https://github.com/organizations/XXX/settings/installations/123456

### Configuration 
- **installation_id** - ID of the SSHAUTH app installation above 
- **admin_role** - Github role to associate with `admin_user`
- **api_url** - Base URL of the GitHib API
- **default_role** - Github role to associate with `user`
- **org** - Unique organisation ID. This must be same as the name appear on the URL at the organisation page
- **team_name** - Unique name of a team to filter users from. Omitting this field will make *all* users of the organisation a candidate.
- **media_type** - `Accept` header value of the GitHub application API. More information at [Authenticating with GitHub Apps](https://developer.github.com/apps/building-github-apps/authenticating-with-github-apps/) and [Media Types](https://developer.github.com/v3/media/).
 
#### Sample
```json
{
  "Name": "github",
  "Configuration": {
      "installation_id": 0,
      "admin_role": "",
      "api_url": "https://api.github.com",
      "default_role": "",
      "org": "",
      "team_name": "",
      "media_type": "application/vnd.github.machine-man-preview+json"
  }
}
```
