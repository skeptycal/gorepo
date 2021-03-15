package gorepo

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

const (
	defaultUserFileNameFmtString = `.gorepo_user_%s`
	copyrightFmtString           = `Copyright (c) %s`
	defaultGitignoreItems        = "macos linux windows ssh vscode go zsh node vue nuxt python django"
	defaultGitIgnoreAPIFmtString = `https://www.toptal.com/developers/gitignore/api/%s`
)

type User interface {
	Name() string
	UserName() string
}

type user struct {
	name                  string
	userName              string
	configFile            string
	DefaultGitignoreItems string
	DefaultLicense        string
	DefaultCopyrightYear  int
}

func (u *user) GitHub() string {
	return fmt.Sprintf(gitHubRemoteFmtString, u.UserName())
}

func (u *user) UserName() string {
	return strings.TrimSpace(u.userName)
}

func (u *user) Name() string {
	return strings.TrimSpace(u.name)
}

func (u *user) ConfigPath() string {
	if u.configFile == "" {
		u.configFile = fmt.Sprintf(defaultUserFileNameFmtString, u.UserName())
	}
	return u.configFile
}

func (u *user) Copyright(year int) (yearString string) {
	if year == 0 || year > time.Now().Year() {
		year = u.DefaultCopyrightYear
	}

	if u.DefaultCopyrightYear != time.Now().Year() {
		yearString = fmt.Sprintf("%d-", u.DefaultCopyrightYear)
	}

	yearString += fmt.Sprint("%d", time.Now().Year())

	return
}

func (u *user) MarshalJSON() ([]byte, error) {
	return json.Marshal(u)
}

func (u *user) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, u)
}

func (u *user) Save() error {

	b, err := u.MarshalJSON()
	if err != nil {
		return err
	}

	return os.WriteFile(u.ConfigPath(), b, 0644)
}

func (u *user) Load() error {
	b, err := os.ReadFile(u.ConfigPath())
	if err != nil {
		return err
	}

	return u.UnmarshalJSON(b)
}

var defaultUser User = &user{
	name:                  "Michael Treanor",
	userName:              "skeptycal",
	configFile:            "",
	DefaultGitignoreItems: defaultGitignoreItems,
	DefaultLicense:        "MIT",
	DefaultCopyrightYear:  time.Now().Year(),
}
