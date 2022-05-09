package service

import (
	"fmt"
	"io/ioutil"
	"os"
)

type ServiceDefinition struct {
	Name    string
	Image   string
	Ports   []string
	Env     []string
	Volumes []string
}

func GetDefinition(name string) (string, error) {
	userDir, err := os.UserHomeDir()
	check(err)
	content, err := ioutil.ReadFile(fmt.Sprintf("%s/.loko/data/services/%s.yml", userDir, name))
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
