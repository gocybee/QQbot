package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"testing"
)

func Test(t *testing.T) {
	var config Config
	yamlFile, err := ioutil.ReadFile("D:/GithubLibrary/gocybee/QQbot/config/config.yml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v", config)
}
