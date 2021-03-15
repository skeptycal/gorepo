package gorepo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	log "github.com/sirupsen/logrus"
)

func getConfigFileName() (string, error) {

	repoJSONFileName := getParentFolderName() + ".json"
	log.Info(repoJSONFileName)

	fi, err := os.Stat(repoJSONFileName)
	if err == nil {
		return filepath.Abs(fi.Name())
	}

	// check for default file name
	repoJSONFileName = configFileDefaultName
	log.Info(repoJSONFileName)

	fi, err = os.Stat(repoJSONFileName)
	if err == nil {
		return fi.Name(), nil
	}

	return "", errors.New("no configuration file found")

}

func GetConfig() (*softwareInfo, error) {

	configFileName := configFileDefaultName

	configFile, err := os.Open(configFileName)
	if err != nil {
		log.Errorf("error opening config file %s: %v", configFileName, err)
		return nil, err
	}

	defer configFile.Close()

	n := new(softwareInfo)

	jsonParser := json.NewDecoder(configFile)

	err = jsonParser.Decode(n)
	if err != nil {
		log.Errorf("error parsing config file %s: %v", configFileName, err)
	}
	return n, err
}

type softwareInfo struct {
	Name        string `json:"Name"`
	Author      string `json:"Author"`
	Description string `json:"Description"`
	License     string `json:"License"`
	StartYear   int    `json:"StartYear"`
}

func (s *softwareInfo) WriteConfigFile() error {
	return nil
}

func (s *softwareInfo) Copyright() string {
	year := time.Now().Year()
	yearString := fmt.Sprintf("Copyright (C) %d", s.StartYear)
	if year != s.StartYear {
		return fmt.Sprintf("%s-%d %s", yearString, year, s.Author)
	}
	return fmt.Sprintf("%s %s", yearString, s.Author)
}

func (s *softwareInfo) LicenseInfo() string {
	return fmt.Sprintf("%s\n\n%s %s", s.Copyright(), s.Name, cliLicenseInfoStub)
}

func (s *softwareInfo) LicenseHeader() string {
	return fmt.Sprintf("%s\n%s", s.Copyright(), licenseHeaderStub)
}

const (
	configFileDefaultName = "gorepo.json"
	cliLicenseInfoStub    = `comes with ABSOLUTELY NO WARRANTY.
This is free software, and you are welcome to redistribute it
under certain conditions;  for details type 'tree -help'.`

	licenseHeaderStub = `
All Rights reserved

This program is free software; you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation; either version 2 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program; if not, write to the Free Software
Foundation, Inc., 59 Temple Place, Suite 330, Boston, MA  02111-1307  USA`
)
