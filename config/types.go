package config

type DbConfig struct{
	DBUSERNAME string `yaml:"DBUSERNAME"`
	DBPASSWORD string `yaml:"DBPASSWORD"`
	DBHOST     string `yaml:"DBHOST"`
	DBNAME     string `yaml:"DBNAME"`
}