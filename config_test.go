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

	site1 := Site{Config: Config{Interval: 15000000000}}
	expectedSites := map[string]*Site{"https://site1/": &site1}

	tmpfile, err := createTempConfig(configFile)
	defer os.Remove(tmpfile.Name())
	assert.Nil(err, "create temp file error")

	sitest := new(Sitest)
	sitest.ConfigFile = tmpfile.Name()
	sitest.LoadConfig()

	assert.Equal(sitest.Sites, expectedSites, "assert sites")
}
func TestLoadConfigWithDefault(t *testing.T) {

	assert := assert.New(t)

	configFile := `default:
  interval: 30s
sites:
  "https://site2/":
    interval: 15s
  "https://site3/": {}`

	site2 := Site{Config: Config{Interval: 15000000000}}
	site3 := Site{Config: Config{Interval: 30000000000}}
	expectedSites := map[string]*Site{"https://site2/": &site2, "https://site3/": &site3}

	tmpfile, err := createTempConfig(configFile)
	defer os.Remove(tmpfile.Name())
	assert.Nil(err, "create temp file error")

	sitest := new(Sitest)
	sitest.ConfigFile = tmpfile.Name()
	sitest.LoadConfig()

	assert.Equal(sitest.Sites, expectedSites, "assert sites")

}
