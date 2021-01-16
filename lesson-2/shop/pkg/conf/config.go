package conf

import (
	"flag"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

// Config struct for application config
type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"server"`
	Telegram struct {
		Token  string `yaml:"bot-token"`
		ChatId int64  `yaml:"chat-id"`
	} `yaml:"telegram"`
	Smtp struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
		From string `yaml:"from"`
		Pass string `yaml:"pass"`
	} `yaml:"smtp"`
}

// Will parse config flag and returns the path to config file
// or error if file doesn't exist.
func ParseConfigFlag() (string, error) {
	var configPath string

	flag.StringVar(&configPath, "c", "./config.yaml", "path to yaml config file")
	flag.Parse()

	_, err := os.Stat(configPath)
	if err != nil {
		return "", err
	}

	return configPath, nil
}

// NewConfig returns a new decoded Config struct.
func NewConfig(configPath string) (*Config, error) {
	// Create config structure
	config := &Config{}

	// Open config file
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(data, config)
	if err != nil {
		return nil, err
	}

	return config, err
}
