package config

import (
	"QQbot/global"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"testing"
)

func Test(t *testing.T) {
	var config OtherConfig
	yamlFile, err := ioutil.ReadFile(global.URLTOOTHERConfig)
	if err != nil {
		fmt.Println("err when read file:", err)
	}
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		fmt.Println("err when unmarshal file:", err)
	}
	fmt.Printf("%#v\n", config)
}
