package config

import(
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

func confg() *DbConfig{
	configFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
        log.Printf("Error in loading yaml file#%v ", err)
    }

	config := &DbConfig{}
	err = yaml.Unmarshal(configFile,config)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return config

}

var (Config = confg())