package config

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

const DashDir = "gh-actions-dashboard"
const ConfigFileName = "config.yml"
const DEFAULT_XDG_CONFIG_DIRNAME = ".config"

type Config struct {
	RepoPaths []string `yaml:"repoPaths"`
}

type configError struct {
	configDir string
	parser    ConfigParser
	err       error
}

type ConfigParser struct{}

func (parser ConfigParser) getDefaultConfig() Config {
	return Config{
		RepoPaths: []string{},
	}
}

func (parser ConfigParser) getDefaultConfigYamlContents() string {
	defaultConfig := parser.getDefaultConfig()
	yaml, _ := yaml.Marshal(defaultConfig)

	return string(yaml)
}

func (e configError) Error() string {
	return fmt.Sprintf(
		`Couldn't find a config.yml configuration file.
Create one under: %s

Example of a config.yml file:
%s

For more info, go to https://github.com/morzzz007/gh-actions-dashboard
press q to exit.

Original error: %v`,
		path.Join(e.configDir, DashDir, ConfigFileName),
		string(e.parser.getDefaultConfigYamlContents()),
		e.err,
	)
}

func (parser ConfigParser) writeDefaultConfigContents(newConfigFile *os.File) error {
	_, err := newConfigFile.WriteString(parser.getDefaultConfigYamlContents())

	if err != nil {
		return err
	}

	return nil
}

func (parser ConfigParser) createConfigFileIfMissing(configFilePath string) error {
	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		newConfigFile, err := os.OpenFile(configFilePath, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)
		if err != nil {
			return err
		}

		defer newConfigFile.Close()
		return parser.writeDefaultConfigContents(newConfigFile)
	}

	return nil
}

func (parser ConfigParser) getExistingConfigFile() (*string, error) {
	var err error
	var dashConfigFile string
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	xdgConfigDir := os.Getenv("XDG_CONFIG_HOME")
	if xdgConfigDir == "" {
		xdgConfigDir = filepath.Join(homeDir, DEFAULT_XDG_CONFIG_DIRNAME)
	}

	dashConfigFile = filepath.Join(xdgConfigDir, DashDir, ConfigFileName)
	if _, err := os.Stat(dashConfigFile); err == nil {
		return &dashConfigFile, nil
	}

	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		return nil, err
	}

	dashConfigFile = filepath.Join(userConfigDir, DashDir, ConfigFileName)
	if _, err := os.Stat(dashConfigFile); err == nil {
		return &dashConfigFile, nil
	}

	return nil, nil
}

func (parser ConfigParser) getConfigFileOrCreateIfMissing() (*string, error) {
	var err error

	existingConfigFile, err := parser.getExistingConfigFile()
	if err != nil {
		return nil, err
	}
	if existingConfigFile != nil {
		return existingConfigFile, nil
	}

	configDir := os.Getenv("XDG_CONFIG_HOME")
	if configDir == "" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return nil, err
		}
		configDir = filepath.Join(homeDir, DEFAULT_XDG_CONFIG_DIRNAME)
	}

	dashConfigDir := filepath.Join(configDir, DashDir)
	err = os.MkdirAll(dashConfigDir, os.ModePerm)
	if err != nil {
		return nil, configError{parser: parser, configDir: configDir, err: err}
	}

	configFilePath := filepath.Join(dashConfigDir, ConfigFileName)
	err = parser.createConfigFileIfMissing(configFilePath)
	if err != nil {
		return nil, configError{parser: parser, configDir: configDir, err: err}
	}

	return &configFilePath, nil
}

type parsingError struct {
	err error
}

func (e parsingError) Error() string {
	return fmt.Sprintf("failed parsing config.yml: %v", e.err)
}

func (parser ConfigParser) readConfigFile(path string) (Config, error) {
	config := parser.getDefaultConfig()
	data, err := os.ReadFile(path)
	if err != nil {
		return config, configError{parser: parser, configDir: path, err: err}
	}

	err = yaml.Unmarshal([]byte(data), &config)
	if err != nil {
		return config, err
	}
	return config, err
}

func initParser() ConfigParser {
	return ConfigParser{}
}

func ParseConfig() (Config, error) {
	parser := initParser()

	var config Config
	var err error

	configFilePath, err := parser.getConfigFileOrCreateIfMissing()
	if err != nil {
		return config, parsingError{err: err}
	}

	config, err = parser.readConfigFile(*configFilePath)
	if err != nil {
		return config, parsingError{err: err}
	}

	return config, nil
}
