package gorepo

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/skeptycal/util/gofile"
	"github.com/skeptycal/util/stringutils/ansi"
)

var (
	// byte code for ANSI red

	redText   = ansi.RedString
	whiteText = ansi.WhiteString
	reset     = ansi.Reset

	shellContext context.Context = context.TODO()
)

// type result struct {
// 	stdout string
// 	stderr string
// 	retval int
// 	err    error
// }

// gi returns a string response from the www.gitignore.io API containing
// standard .gitignore items for the args given.
//
//      default: "macos linux windows ssh vscode go zsh node vue nuxt python django"
//
// using: https://www.toptal.com/developers/gitignore/api/macos,linux,windows,ssh,vscode,go,zsh,node,vue,nuxt,python,django
func gi(args string) (string, error) {

	if len(args) == 0 {
		args = defaultGitignoreItems
	}

	args = strings.Join(strings.Split(args, " "), ",")

	url := fmt.Sprintf(defaultGitIgnoreAPIFmtString, args)

	return GetPage(url)
}

// Shell executes a command line string and returns the result.
func Shell(command string) string {

	cmd := cmdPrep(command)
	stdout, err := cmd.Output()

	if err != nil {
		return fmt.Errorf("%Terror: %v%T", redText, err, reset).Error()
	}

	return string(stdout)
}

// GetPage - return result from url
func GetPage(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("server returned error code: %v", resp.Status)
	}
	defer resp.Body.Close()

	buf, err := ioutil.ReadAll(bufio.NewReaderSize(resp.Body, gofile.InitialCapacity(resp.ContentLength)))

	if err != nil {
		return "", err
	}

	return string(buf), nil
}

// BufferURL - read result from url into sb
func BufferURL(url string) (string, error) {
	sb := strings.Builder{}
	defer sb.Reset()

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("server returned error code: %v", resp.Status)
	}

	defer resp.Body.Close()

	_, err = io.Copy(&sb, resp.Body)

	if err != nil {
		return "", err
	}

	return sb.String(), nil
}

// WriteFile creates the file 'fileName' and writes 'data' to it.
// It returns any error encountered. If the file already exists, it
// will be TRUNCATED and OVERWRITTEN.
func WriteFile(fileName string, data string) error {
	dataFile, err := OpenTrunc(fileName)
	if err != nil {
		log.Println(err)
		return err
	}
	defer dataFile.Close()

	n, err := dataFile.WriteString(data)
	if err != nil {
		log.Println(err)
		return err
	}
	if n != len(data) {
		log.Printf("incorrect string length written (wanted %d): %d\n", len(data), n)
		return fmt.Errorf("incorrect string length written (wanted %d): %d", len(data), n)
	}
	return nil
}

// OpenTrunc creates and opens the named file for writing. If successful, methods on
// the returned file can be used for writing; the associated file descriptor has mode
//      O_WRONLY|O_CREATE|O_TRUNC
// If the file does not exist, it is created with mode o644;
//
// If the file already exists, it is TRUNCATED and overwritten
//
// If there is an error, it will be of type *PathError.
func OpenTrunc(name string) (*os.File, error) {
	return os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 644)
}

// cmdPrep prepares a Cmd struct from a command line string.
func cmdPrep(command string) *exec.Cmd {
	commandSlice := strings.Split(command, " ")
	app := commandSlice[0]
	args := strings.Join(commandSlice[1:], " ")
	return exec.CommandContext(shellContext, app, args)
}

// getParentFolderName returns name of immediate parent folder; used to create repo name
func getParentFolderName() string {
	file, err := os.Getwd()
	if err != nil {
		log.Errorf("getParentFolderName could not locate parent folder %v", err)
		return ""
	}
	return filepath.Base(file)
}

// fileExists checks if a file exists and is not a directory
func fileExists(fileName string) bool {
	info, err := os.Stat(fileName)
	if errors.Is(err, os.ErrNotExist) {
		return false
	}
	return !info.IsDir()
}

func exists(file string) bool {
	_, err := os.Stat(file)
	return err == nil
}

// Notes: Cmd struct summary:
/*
type Cmd struct {
	Path            string
	Args            []string
	Env             []string
	Dir             string
	Stdin           io.Reader
	Stdout          io.Writer
	Stderr          io.Writer
	ExtraFiles      []*os.File
	SysProcAttr     *syscall.SysProcAttr
	Process         *os.Process
	ProcessState    *os.ProcessState
	ctx             context.Context // nil means none
	lookPathErr     error           // LookPath error, if any.
	finished        bool            // when Wait was called
	childFiles      []*os.File
	closeAfterStart []io.Closer
	closeAfterWait  []io.Closer
	goroutine       []func() error
	errch           chan error // one send per goroutine
	waitDone        chan struct{}
}
*/
