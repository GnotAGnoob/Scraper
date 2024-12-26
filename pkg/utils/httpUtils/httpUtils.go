package httpUtils

import (
	"encoding/json"
	"io"
	"net/http"
)

// https://www.useragentlist.net/
const UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36 Edg/113.0.1774.35"

func SendRequest[T any](client *http.Client, method string, url string, body io.Reader) (t T, err error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return t, err
	}

	req.Header.Set("User-Agent", UserAgent)
	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return t, err
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return t, err
	}

	err = json.Unmarshal(resBody, &t)
	if err != nil {
		return t, err
	}

	return t, nil
}
