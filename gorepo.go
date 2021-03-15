// Package gorepo is an implementation of a Git repository
// linked to a remote GitHub repository.
package gorepo

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/skeptycal/zsh"
)

const (
	defaultVersion        = "0.1.0"
	gitHubRemoteFmtString = `https://www.github.com/%s/%s`
	gitHubDocsFmtString   = `https://%s.github.com/%s`
)

// GitRepo represents a git repository and its associated GitHub repository.
type GitRepo struct {
	Name           string
	User           User
	LocalPath      string
	Version        string
	remote         string
	GitignoreItems string
}

func (repo *GitRepo) RemoteURL() string {
	return fmt.Sprintf(gitHubRemoteFmtString, repo.User.UserName, repo.remote)
}

func (repo *GitRepo) DocsURL() string {
	return fmt.Sprintf(gitHubDocsFmtString, repo.User.UserName, repo.remote)
}

func init() {
	pwd, _ := os.Getwd()
	pwd = path.Clean(pwd)
	_, repoName := path.Split(pwd)

	// todo - setup user config save file
	var user User = defaultUser

	repo := &GitRepo{
		Name:      repoName,
		User:      user,
		LocalPath: pwd,
		Version:   defaultVersion,
		remote:    repoName,
	}
}

// Setup initializes the repo, creates files, prompts as needed, creates the
// github.com repository, and pushes the initial commit.
func Setup() error {
	err := gitRepoSetup()
	if err != nil {
		log.Printf("gitRepoSetup failed with %v", err)
		return err
	}
	err = createAutomatedFiles()
	if err != nil {
		log.Printf("createAutomatedFiles failed with %v", err)
		return err
	}
	return nil
}

// gitRepoSetup initializes the repo, prompts as needed, creates the
// github.com repository, and pushes the initial commit.
func gitRepoSetup() error {
	err := gitInit()
	if err != nil {
		return err
	}
	// todo - stuff
	return err
}

// CreateAutomatedFiles creates the automated files.
func createAutomatedFiles() error {
	zsh.Sh("touch main.go")
	return nil
}

// gitIgnore writes a .gitignore file, including default items followed by the response from
// the www.gitignore.io API containing standard .gitignore items for the args given.
//
//      default: "macos linux windows ssh vscode go zsh node vue nuxt python django"
//
// using: https://www.toptal.com/developers/gitignore/api/macos,linux,windows,ssh,vscode,go,zsh,node,vue,nuxt,python,django
func gitIgnore(args string) error {
	// notes - .gitignore header
	/*
	   # gorepo - .gitignore file

	   # generic secure items:
	   *private*
	   *secret*
	   *bak

	   # repo specific items
	   coverage.txt
	   profile.out
	*/

	var sb strings.Builder
	defer sb.Reset()

	sb.WriteString(fmt.Sprintf("# %s - .gitignore file\n\n", repoName))

	sb.WriteString("# generic secure items:\n")
	sb.WriteString("*private*\n*secret*\n*bak\n\n")

	sb.WriteString("# repo specific items:\n")
	sb.WriteString("coverage.txt\nprofile.out\n\n")

	// add .gitignore contents from gitignore.io API
	gitignore, err := gi(args)
	if err != nil {
		return err
	}
	sb.WriteString(gitignore)

	return WriteFile(".gitignore", sb.String())
}

// GitInit initializes the Git environment
func gitInit() error {
	if !fileExists(".gitignore") {
		gitIgnore("")
	}

	Shell("git init")
	GitCommitAll("initial commit")
	return nil
}

// GoMod creates and initializes the repo go.mod file.
func GoMod() error {
	Shell("go mod init")
	Shell("go mod tidy")
	Shell("go mod download")
	GitCommitAll("go mod setup")
	return nil
}

// GitCommit creates a commit with message
func GitCommitAll(message string) error {
	Shell("git add --all")
	Shell("git commit -m '" + message + "'")
	return nil
}

// GoSum creates the go.sum file.
func GoSum() error {
	return nil
}

// GoTestSh creates the go.test.sh script.
func GoTestSh() error {
	return nil
}

// GoDoc creates the go.doc file.
func GoDoc() error {
	return nil
}

// BugReportMd creates the .github/ISSUE_TEMPLATE/bug_report.md file.
func BugReportMd() error {
	return nil
}

// FeatureRequestMd creates the .github/ISSUE_TEMPLATE/feature_request.md file.
func FeatureRequestMd() error {
	return nil
}

// GitWorkflows creates initial .github/workflows/... files:
// codeql-analysis.yml go.yml greetings.yml label.yml stale.yml
func GitWorkflows() error {
	return nil
}

// CodeCovYml creates the initial .codecov.yml file.
func CodeCovYml() error {
	return nil
}

// FundingYml creates the initial FUNDING.yml file.
func FundingYml() error {
	return nil
}

// PreCommitYaml creates the initial .pre-commit-config.yaml file.
func PreCommitYaml() error {
	return nil
}

// TravisYml creates the initial .travis.yml file.
func TravisYml() error {
	return nil
}

// DocGo creates the initial doc.go file.
func DocGo() error {
	return nil
}

// ReadMeMd creates the initial README.md file.
func ReadMeMd() error {
	return nil
}

// SecurityMd creates the initial SECURITY.md file.
func SecurityMd() error {
	return nil
}

// CodeOfConduct creates the initial CODE_OF_CONDUCT.md file.
func CodeOfConduct() error {
	return nil
}

// License creates the initial LICENSE file.
func License(license string) error {
	return nil
}
