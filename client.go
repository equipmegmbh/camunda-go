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

func (c *Client) authorize(ctx context.Context) (*Token, error) {
	return new(Token), nil
}

func (c *Client) send(ctx context.Context, url, method, ct string, payload io.Reader, out interface{}) error {
	c.once.Do(c.init)

	var content []byte
	var err error

	var req *http.Request
	var resp *http.Response

	req, err = http.NewRequestWithContext(ctx, method, url, payload)

	if err != nil {
		return err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", ct)

	resp, err = c.cli.Do(req)

	if err != nil {
		return err
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		goto ok
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

	if resp.StatusCode == 204 {
		return nil
	}

	content, err = ioutil.ReadAll(utfbom.SkipOnly(resp.Body))

	if err != nil {
		return err
	}

	return json.Unmarshal(content, out)
}
