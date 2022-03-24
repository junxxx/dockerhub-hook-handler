package internal

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Version  string
	Services []struct {
		RepoName string `yaml:"repo-name"`
		Handler  string `yaml:"handler"`
		Script   string `yaml:"script"`
	}
}

var config *Config

func GetConf() *Config {
	if config == nil {
		data, err := ioutil.ReadFile("./configs/config.yaml")
		type Config struct {
			Version  string
			Services []struct {
				RepoName string `yaml:"repo-name"`
				Handler  string `yaml:"handler"`
				Script   string `yaml:"script"`
			}
		}

		err = yaml.Unmarshal([]byte(data), &config)
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Printf("t:%v\n", config)
	return config
}
