package camunda

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	endpoint = "http://127.0.0.1:18080"
)

var (
	client Client
)

func Configure() {
	// ...
}

func FetchAndLock(ctx context.Context, request *Fetch) ([]*Task, error) {
	var uri string
	var err error

	payload, err := json.Marshal(request)

	if err != nil {
		return nil, err
	}

	result := make([]*Task, 0)

	uri = fmt.Sprintf("%s/engine-rest/external-task/fetchAndLock", endpoint)
	err = client.send(ctx, uri, http.MethodPost, bytes.NewReader(payload), &result)

	if err != nil {
		return nil, err
	}

	return result, nil
}
