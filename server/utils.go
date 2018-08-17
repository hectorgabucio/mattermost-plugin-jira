package main

import (
	"net/url"
	"regexp"
)

func parseJiraUsernamesFromText(text string) []string {
	usernameMap := map[string]bool{}
	usernames := []string{}

	var re = regexp.MustCompile(`(?m)\[~([a-zA-Z0-9-_.]+)\]`)
	for _, match := range re.FindAllString(text, -1) {
		name := match[:len(match)-1]
		name = name[2:]
		if !usernameMap[name] {
			usernames = append(usernames, name)
			usernameMap[name] = true
		}
	}

	return usernames
}

func getIssueURL(i *Issue) string {
	u, _ := url.Parse(i.Self)
	return u.Scheme + "://" + u.Host + "/browse/" + i.Key
}

func getUserURL(i *Issue, user *User) string {
	u, _ := url.Parse(i.Self)
	return u.Scheme + "://" + u.Host + "/people/" + user.AccountId
}
