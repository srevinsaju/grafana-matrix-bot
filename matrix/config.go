package matrix

import (
	"encoding/json"
	"io/ioutil"
)

func ConfigFromFile(filepath string) (Config, error) {
	rawData, err := ioutil.ReadFile(filepath)
	var cfg Config
	err = json.Unmarshal(rawData, &cfg)
	if err != nil {
		logger.Fatal(err)
		return Config{}, err
	}
	return cfg, nil

}