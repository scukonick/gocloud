package lib

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"time"
)

type Computer struct {
}

func NewComputer() *Computer {
	return &Computer{}
}

func (c *Computer) Run(code string) (string, error) {
	values := url.Values{}
	values.Add("version", "2")
	values.Add("body", code)

	httpClient := http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := httpClient.PostForm("https://play.golang.org/compile", values)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return "", errors.New("calculation failed")
	}

	type calcResult struct {
		Errors string `json:"Errors"`
		Events []struct {
			Message string `json:"Message"`
			Kind    string `json:"Kind"`
		} `json:"Events"`
	}

	r := &calcResult{}

	unmarshaller := json.NewDecoder(resp.Body)
	err = unmarshaller.Decode(r)
	if err != nil {
		return "", errors.New("calculation failed")
	}

	if r.Errors != "" {
		return "", errors.New(r.Errors)
	}

	if len(r.Events) != 1 {
		return "", errors.New("calculation failed")
	}

	return r.Events[0].Message, nil
}
