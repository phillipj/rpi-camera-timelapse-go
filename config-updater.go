package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type configUpdaterFn func(config)

func fetchAndApplyRemoteConfig(remoteConfigUrl string, updateConfig configUpdaterFn) {
	if remoteConfigUrl == "" {
		return
	}

	config, err := downloadConfig(remoteConfigUrl)
	if err != nil {
		log.Printf("*** Not able to download remote config: %v", err)
		return
	}

	updateConfig(config)
}

func downloadConfig(url string) (config, error) {
	response, err := http.Get(url)

	if err != nil {
		return config{}, fmt.Errorf("downloading remote config: %v", err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return config{}, fmt.Errorf("reading contents of response: %v", err)
	}

	var conf config

	if err := json.Unmarshal(body, &conf); err == nil {
		return conf, nil
	} else {
		return config{}, fmt.Errorf("deserialising response: %v", err)
	}
}
