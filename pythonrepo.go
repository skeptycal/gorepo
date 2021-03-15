package gorepo

import "github.com/skeptycal/zsh"

func RepoInitPython() {
    zsh.Sh("poetry version")
}
