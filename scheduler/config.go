package scheduler

import (
	"log"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func GetConfig() *TypeConfig {

	conf := &TypeConfig{}

	data, errReadFile := ioutil.ReadFile("config/earnings-scheduler-config.yaml")
	if errReadFile != nil {
		log.Fatalf("error: %v", errReadFile)
	}

	errYaml := yaml.Unmarshal(data, &conf)
	if errYaml != nil {
		log.Fatalf("error: %v", errYaml)
	}

	return conf
}
