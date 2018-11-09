package main

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func createTempConfig(content string) (tmpfile *os.File, err error) {
	tmpfile, err = ioutil.TempFile("", "sitest-test-config")
	defer tmpfile.Close()
	if err == nil {
		_, err = tmpfile.Write([]byte(content))
	}
	return
}

func TestLoadConfig(t *testing.T) {

	assert := assert.New(t)

	configFile := `
sites:
  "https://site1/":
    interval: 15s`

	expectedConfig := map[string]Config{
		"https://site1/": Config{Interval: 15000000000},
	}

	tmpfile, err := createTempConfig(configFile)
	defer os.Remove(tmpfile.Name())
	assert.Nil(err, "create temp file error")

	sitest := new(Sitest)
	sitest.ConfigFile = tmpfile.Name()
	sitest.LoadConfig()

	assert.Equal(sitest.Sites, expectedConfig, "assert length")
}
func TestLoadConfigWithDefault(t *testing.T) {

	assert := assert.New(t)

	configFile := `default:
  interval: 30s
sites:
  "https://site2/":
    interval: 15s
  "https://site3/": {}`

	expectedConfig := map[string]Config{
		"https://site2/": Config{Interval: 15000000000},
		"https://site3/": Config{Interval: 30000000000},
	}

	tmpfile, err := createTempConfig(configFile)
	defer os.Remove(tmpfile.Name())
	assert.Nil(err, "create temp file error")

	sitest := new(Sitest)
	sitest.ConfigFile = tmpfile.Name()
	sitest.LoadConfig()

	assert.Equal(sitest.Sites, expectedConfig, "assert length")
}
