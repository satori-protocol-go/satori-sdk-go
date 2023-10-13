package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/dezhishen/satori-sdk-go/pkg/config"
)

type HttpChannel struct {
	cli         http.Client
	endpoint    string
	accessToken string
	// todo add secret support
}

type accessTokenTransport struct {
	http.RoundTripper
	accessToken string
}

func (ct *accessTokenTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", ct.accessToken))
	return ct.RoundTripper.RoundTrip(req)
}

func NewHttpApiTemplate(conf *config.SatoriApiConfig) (ApiTemplate, error) {
	if strings.HasSuffix("/", conf.Endpoint) {
		conf.Endpoint = strings.TrimSuffix(conf.Endpoint, "/")
	}
	var httpCli http.Client
	if conf.AccessToken != "" {
		httpCli = http.Client{
			Transport: &accessTokenTransport{http.DefaultTransport, conf.AccessToken},
		}
	} else {
		httpCli = http.Client{}
	}
	return &HttpChannel{
		cli:         httpCli,
		endpoint:    conf.Endpoint,
		accessToken: conf.AccessToken,
	}, nil
}

func concatUrl(basePath, url string) string {
	return basePath + "/" + url
}

func (c *HttpChannel) PostByRequestForResult(url string, req, result interface{}) error {
	requestBody, _ := json.Marshal(req)
	resp, err := c.cli.Post(
		concatUrl(c.endpoint, url),
		"application/json",
		bytes.NewBuffer(requestBody),
	)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	respBodyContent, _ := io.ReadAll(resp.Body)
	json.Unmarshal(respBodyContent, &result)
	return nil
}

func (c *HttpChannel) PostForResult(url string, result interface{}) error {
	resp, err := c.cli.Post(
		concatUrl(c.endpoint, url),
		"application/json",
		bytes.NewBuffer([]byte("{}")),
	)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	respBodyContent, _ := io.ReadAll(resp.Body)
	json.Unmarshal(respBodyContent, &result)
	return nil
}

func (c *HttpChannel) PostByRequest(url string, req interface{}) error {
	requestBody, _ := json.Marshal(req)
	_, err := c.cli.Post(
		concatUrl(c.endpoint, url),
		"application/json",
		bytes.NewBuffer(requestBody),
	)
	return err
}

func (c *HttpChannel) Post(url string) error {
	_, err := c.cli.Post(
		concatUrl(c.endpoint, url),
		"application/json",
		bytes.NewBuffer([]byte("{}")),
	)
	return err
}
