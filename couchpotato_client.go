package go_couchpotato

import (
	"net/http"
	"net/url"
	"encoding/json"
	"fmt"
	"errors"
	"strings"
	"bytes"
	"io/ioutil"
	"github.com/Sirupsen/logrus"
)

type CouchpotatoClient struct {
	address *url.URL
	apiKey string
	HttpClient *http.Client
}

func NewCouchpotatoClient(address string, apiKey string) (*CouchpotatoClient, error) {

	if address == "" {
		return nil, errors.New("No address specified")
	}

	addressUrl, err := url.Parse(address)

	path := addressUrl.Path
	//
	if !strings.HasSuffix(path, "/") {
		path += "/"
	}

	if !strings.HasSuffix(path, "api/") {
		path += "api/"
	}

	addressUrl.Path = path

	if err != nil {
		return nil, err
	}

	return &CouchpotatoClient{
		address:addressUrl,
		apiKey:apiKey,
		HttpClient:http.DefaultClient,
	}, nil
}

func (cc *CouchpotatoClient) DoRequest(action, path string, params map[string]string, reqData, resData interface{}) error {
	lookupUrl := *cc.address

	parameters := url.Values{}

	if params != nil {
		for k, v := range params {
			parameters.Add(k, v)
		}
	}

	lookupUrl.RawQuery = parameters.Encode()
	lookupUrl.Path += cc.apiKey + "/" + path

	jsonValue, err := json.Marshal(reqData)

	if err != nil {
		return err
	}

	logrus.Debugf("Calling Couchpotato at %v", lookupUrl.String())

	req, err := http.NewRequest(action, lookupUrl.String(), bytes.NewBuffer(jsonValue))

	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	response, err := cc.HttpClient.Do(req)

	if err != nil {
		return err
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		bodyBytes, err := ioutil.ReadAll(response.Body)

		if err == nil {
			logrus.Debugf("Failing (%v) call returned:\n%v", response.StatusCode, string(bodyBytes))
		}

		return errors.New(fmt.Sprintf("Status code %v", response.StatusCode))
	}

	body, err := ioutil.ReadAll(response.Body)
	err = json.NewDecoder(bytes.NewBuffer(body)).Decode(resData)

	if err != nil {
		return err
	}

	return nil
}