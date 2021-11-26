package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	jokeapiURL = "https://v2.jokeapi.dev/joke/Any"
)

// JSON represents the return from quotes.rest
type JSON struct {
	Type     string
	Joke     string
	Setup    string
	Delivery string
}

func (j JSON) String() string {
	if j.Type == "single" {
		return j.Joke
	}

	return fmt.Sprintf("%s\n%s", j.Setup, j.Delivery)
}

func jokeapi() (msg string, err error) {
	res, err := http.Get(jokeapiURL)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	j := JSON{}
	err = json.Unmarshal(body, &j)

	if err != nil {
		return "", err
	}

	return j.String(), nil
}
