package config_reader_test

import (
	"github.com/emsal1863/resolvplox_alt/config_reader"
	"reflect"
	"testing"
)

func TestValidateConfigBlank(t *testing.T) {
	test_config := config_reader.Config{}
	if config_reader.ValidateConfig(test_config) == nil {
		t.Error("Blank config shouldn't pass validation")
	}
}

func TestValidateConfigEmptyString(t *testing.T) {
	test_config := config_reader.Config{
		Default: "",
	}
	if config_reader.ValidateConfig(test_config) == nil {
		t.Error("Config with empty default shouldn't pass validation")
	}
}

func TestValidateConfigValid(t *testing.T) {
	test_config := config_reader.Config{
		Directory: make(map[string]string),
		Default:   "a",
	}
	if config_reader.ValidateConfig(test_config) != nil {
		t.Error("Config with valid fields should pass validation")
	}
}

func TestGetServerNoDirectory(t *testing.T) {
	test_config := config_reader.Config{
		Directory: make(map[string]string),
		Default:   "a",
	}

	if test_config.GetServer("a") != test_config.Default {
		t.Error("Should have resulted in Default")
	}
}

func TestGetServerWithDirectory(t *testing.T) {
	test_config := config_reader.Config{
		Directory: make(map[string]string),
		Default:   "a",
	}
	test_config.Directory["google.com"] = "8.8.8.8"

	if test_config.GetServer("google.com") != "8.8.8.8" {
		t.Error("Should have resulted in directory value")
	}
}

func TestFromFile(t *testing.T) {
	path := "reader_testconfig.plox"
	conf, err := config_reader.FromFile(path)

	expectedConf := config_reader.Config{
		Directory: map[string]string{
			"google.com": "8.8.8.4",
		},
		Default: "8.8.8.8",
	}

	t.Log(conf)
	if err != nil {
		t.Error("Error parsing test file: " + err.Error())
	}

	if !reflect.DeepEqual(conf, expectedConf) {
		t.Error("reader_testconfig.plox wasn't parsed correctly")
	}
}

func TestFromFileBadWhitespace(t *testing.T) {
	path := "reader_testconfigbadwhitespace.plox"
	conf, err := config_reader.FromFile(path)

	expectedConf := config_reader.Config{
		Directory: map[string]string{
			"google.com": "8.8.8.4",
		},
		Default: "8.8.8.8",
	}

	t.Log(conf)
	if err != nil {
		t.Error("Error parsing test file: " + err.Error())
	}

	if !reflect.DeepEqual(conf, expectedConf) {
		t.Error("reader_testconfig.plox wasn't parsed correctly")
	}
}
