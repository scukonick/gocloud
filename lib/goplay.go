package lib

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type GoplayComputer struct {
	httpClient http.Client
}

func NewGoplay(httpClient *http.Client) *GoplayComputer {
	if httpClient == nil {
		httpClient = &http.Client{
			Timeout: 5 * time.Second,
		}
	}
	return &GoplayComputer{httpClient: *httpClient}
}

func (c *GoplayComputer) Run(code string) (string, error) {
	codeReader := bytes.NewReader([]byte(code))
	resp, err := c.httpClient.Post("https://goplay.space/compile", "text/plain", codeReader)
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
			Delay   int    `json:"Delay"`
		} `json:"Events"`
		Status      int  `json:"Status"`
		IsTest      bool `json:"IsTest"`
		TestsFailed int  `json:"TestsFailed"`
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
