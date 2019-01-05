package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

func prepareSiteconfig(config *Config, defConfig Config) {
	if config.Interval == 0 {
		config.Interval = defConfig.Interval
	}
}

// LoadConfig reads configuration from yaml file
func (sitest *Sitest) LoadConfig() {

	source, err := ioutil.ReadFile(sitest.ConfigFile)
	if err != nil {
		log.Fatal(err)
	}

	type ConfigStuct struct {
		Default Config
		Sites   map[string]Config
	}

	var config ConfigStuct

	err = yaml.UnmarshalStrict(source, &config)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Config loaded:")

	sitest.Sites = make(map[string]*Site)
	for k, v := range config.Sites {
		prepareSiteconfig(&v, config.Default)
		log.Printf("- site: %v params: %v", k, v)
		sitest.Sites[k] = &Site{Config: v}
	}

}
