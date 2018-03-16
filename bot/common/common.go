package common

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

// type Strategy interface {
// 	Type
// }

// Config strategy and auth options
type Config struct {
	Auth     map[string]map[string]string
	Strategy struct {
		Type    string
		Options map[string]string
	}
}

// LoadFrom Load configuration file
func (c *Config) LoadFrom(path string) *Config {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}
