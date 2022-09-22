package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"testing"
)

func Test(t *testing.T) {
	var config OtherConfig
	yamlFile, err := ioutil.ReadFile("D:/GithubLibrary/gocybee/QQbot/config/other_conf.yml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", config)
}
