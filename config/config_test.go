package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"testing"
)

func Test(t *testing.T) {
	var config Config
	yamlFile, err := ioutil.ReadFile("config.yml")
	if err != nil {
		fmt.Println("err when read file:", err)
	}
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		fmt.Println("err when unmarshal file:", err)
	}
	fmt.Printf("%#v\n", config)
}
