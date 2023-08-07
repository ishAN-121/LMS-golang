package config

import(
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

func confg() *Dbconfig{
	configfile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
        log.Printf("Error in loading yaml file#%v ", err)
    }

	config := &Dbconfig{}
	err = yaml.Unmarshal(configfile,config)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return config

}

var (Config = confg())