package camunda

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"sync"

	"github.com/dimchansky/utfbom"
)

type Client struct {
	once sync.Once
	cli  http.Client
}

func (c *Client) init() {

}

func (c *Client) send(ctx context.Context, url, method string, payload io.Reader, out interface{}) error {
	c.once.Do(c.init)

	var content []byte
	var count int
	var err error
	/*
		token, err := c.authorize(ctx)

		if err != nil {
			return err
		}
	*/
	var req *http.Request
	var resp *http.Response

	req, err = http.NewRequestWithContext(ctx, method, url, payload)

	if err != nil {
		return err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	//req.Header.Set("Authorization", "Bearer "+token.AccessToken)

send:
	resp, err = c.cli.Do(req)
	count = count + 1

	if err != nil {
		return err
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		goto ok
	}

	if resp.StatusCode == 500 && count < 5 {
		goto send
	}

	//goland:noinspection GoUnhandledErrorResult
	defer resp.Body.Close()

	content, err = ioutil.ReadAll(utfbom.SkipOnly(resp.Body))

	if err != nil {
		return err
	}

	err = new(Error)

	if e := json.Unmarshal(content, err); e != nil {
		goto unknown
	}

	if v, ok := err.(*Error); ok {
		return v
	}

unknown:
	return fmt.Errorf("send failed, status code: %d", resp.StatusCode)

ok:
	//goland:noinspection GoUnhandledErrorResult
	defer resp.Body.Close()

	content, err = ioutil.ReadAll(utfbom.SkipOnly(resp.Body))

	if err != nil {
		return err
	}

	return json.Unmarshal(content, out)
}
