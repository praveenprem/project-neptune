/*
Package slack defines implementation of Slack notification

SlackSend example
	var s slack.Slack
	s.User = m.User
	s.Host = m.Host
	s.Provider = m.Provider
	s.SystemUser = m.SystemUser
	return s.SlackSend(n.Url)
*/
package slack
