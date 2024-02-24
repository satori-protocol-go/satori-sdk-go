package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/satori-protocol-go/satori-sdk-go/pkg/config"
)

type HttpChannel struct {
	cli      http.Client
	endpoint string
}

type normalTransport struct {
	http.RoundTripper
	conf config.SatoriApiConfig
}

func (ct *normalTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("X-Platform", ct.conf.Platform)
	req.Header.Add("X-Self-ID", ct.conf.SelfId)
	return ct.RoundTripper.RoundTrip(req)
}

type accessTokenTransport struct {
	http.RoundTripper
	conf config.SatoriApiConfig
}

func (ct *accessTokenTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", ct.conf.AccessToken))
	req.Header.Add("X-Platform", ct.conf.Platform)
	req.Header.Add("X-Self-ID", ct.conf.SelfId)
	return ct.RoundTripper.RoundTrip(req)
}
func NewHttpApiTemplate(conf config.SatoriApiConfig) (ApiTemplate, error) {
	if strings.HasSuffix("/", conf.Endpoint) {
		conf.Endpoint = strings.TrimSuffix(conf.Endpoint, "/")
	}
	var httpCli http.Client
	if conf.AccessToken == "" {
		httpCli = http.Client{
			Transport: &normalTransport{
				http.DefaultTransport,
				conf},
		}
	} else {
		httpCli = http.Client{
			Transport: &accessTokenTransport{
				http.DefaultTransport,
				conf},
		}
	}
	version := conf.Version
	if version == "" {
		version = "v1"
	}
	return &HttpChannel{
		cli:      httpCli,
		endpoint: fmt.Sprintf("%s/%s", conf.Endpoint, version),
	}, nil
}

func concatUrl(basePath, url string) string {
	return basePath + url
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
