package config

type Dbconfig struct{
	DB_USERNAME string `yaml:"DB_USERNAME"`
	DB_PASSWORD string `yaml:"DB_PASSWORD"`
	DB_HOST     string `yaml:"DB_HOST"`
	DB_NAME     string `yaml:"DB_NAME"`
}