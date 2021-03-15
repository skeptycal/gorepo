package gorepo

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"github.com/skeptycal/util/devtools/gogit"
	gofile "github.com/skeptycal/util/gofile/bak"
	"github.com/skeptycal/zsh"
)

// To escape a text segment, bracket it with Escape characters.
// For instance, the tab in this string "Ignore this tab: \xff\t\xff"
// does not terminate a cell and constitutes a single character of
// width one for formatting purposes.
//
// The value 0xff was chosen because it cannot appear in a valid UTF-8 sequence.
//
// Ref: https://golang.org/pkg/text/tabwriter/
const Escape = '\xff'

type GitHubRepo struct {
	name    string
	url     string
	license string
	tag     string // most recent (highest) version tag
}

// gitIgnore writes a .gitignore file, including default items followed by the response from
// the www.gitignore.io API containing standard .gitignore items for the args given.
//
//      default: "macos linux windows ssh vscode go zsh node vue nuxt python django"
//
// using: https://www.toptal.com/developers/gitignore/api/macos,linux,windows,ssh,vscode,go,zsh,node,vue,nuxt,python,django
func (g *GitHubRepo) gitIgnore(args string) error {
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

	// func GitIgnore(reponame string, personalItems string, repoSpecific string, args string) error
	return GitIgnore(g.name, "", "", "")

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
	err := gogit.GitInit()
	if err != nil {
		return err
	}
	GitIgnore("", "", "", "")

	// todo - stuff
	return nil
}

// GetPageBody - returns the body of the page at url.
func GetPageBody(url string) (*bytes.Buffer, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return bytes.NewBuffer(body), nil
}

// CreateAutomatedFiles creates the automated files.
func createAutomatedFiles() error {
	zsh.Sh("touch main.go")
	return nil
}

// GoTestSh creates the go.test.sh script.
func GoTestSh() error {
	gofile.WriteFile("go.test.sh", goTestTemplate)
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

func NewGitHubRepo(name string, dir string) (*GitHubRepo, error) {
	wd := dir
	if wd == "" {
		wd, _ = os.Getwd()

	}
	wd, err := filepath.Abs(dir)
	if err == os.ErrNotExist {
		wd, _ = os.Getwd()

	}

	if err != nil {
		return nil, err
	}

	wd, _ = os.Getwd()
	pwd, _ := filepath.Abs(wd)
	default_path, default_name := path.Split(pwd)

	r := new(GitHubRepo)
	r.name = default_name
	r.url = default_path
	return r, nil
}

// GitHubRepoSetup initializes the repo, creates files, prompts as needed, creates the
// github.com repository, and pushes the initial commit.
func GitHubRepoSetup() (err error) {
	if GitRepoSetup() != nil {
		log.Info("GitRepoSetup failed with %v", err)
		return
	}
	if CreateAutomatedFiles() != nil {
		log.Info("CreateAutomatedFiles failed with %v", err)
		return
	}
	return nil
}

// GitRepoSetup initializes the repo, prompts as needed, creates the
// github.com repository, and pushes the initial commit.
func GitRepoSetup() error {
	return nil
}

// CreateAutomatedFiles creates the automated files.
func CreateAutomatedFiles() error {
	return nil
}

// GoMod creates and initializes the repo go.mod file.
func GoMod() error {
	zsh.Sh("go mod init")
	zsh.Sh("go mod tidy")
	zsh.Sh("go mod download")
	gogit.Commit("go mod setup")
	return nil
}
