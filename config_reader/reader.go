package config_reader

import (
	"bytes"
	"errors"
	"io/ioutil"
	"regexp"
)

type Config struct {
	Directory map[string]string
	Default   string
}

func ValidateConfig(config Config) (err error) {
	if config.Default == "" {
		return errors.New("Can't have a configuration with no default server!")
	}
	return nil
}

func FromFile(filename string) (config Config, err error) {
	ret := Config{
		Directory: make(map[string]string),
		Default:   "",
	}

	matcher, err := regexp.Compile(`([\w|\*|\.]+)\s+([\w|\.]+)`)
	if err != nil {
		return ret, err
	}

	file_content, err := ioutil.ReadFile(filename)
	if err != nil {
		return ret, err
	}
	lines := bytes.Split(file_content, []byte{'\n'})

	for _, line := range lines {
		match := matcher.FindStringSubmatch(string(line))
		if match == nil {
			continue
		} else {
			if match[1] == "*" {
				ret.Default = match[2]
			} else {
				ret.Directory[match[1]] = match[2]
			}
		}
	}

	err = ValidateConfig(ret)
	if err != nil {
		return ret, err
	}

	return ret, nil
}

func (c *Config) GetServer(name string) string {
	if val, ok := c.Directory[name]; ok {
		return val
	} else {
		return c.Default
	}
}
