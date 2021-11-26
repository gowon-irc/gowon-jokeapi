package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	jokeapiURLTmpl = "https://v2.jokeapi.dev/joke/Any?blacklistFlags=%s"
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

func jokeapi(blacklist string) (msg string, err error) {
	url := fmt.Sprintf(jokeapiURLTmpl, blacklist)

	res, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	log.Print(string(body))

	j := JSON{}
	err = json.Unmarshal(body, &j)

	if err != nil {
		return "", err
	}

	return j.String(), nil
}
