package httpUtils

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"
)

// https://www.useragentlist.net/
const UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36 Edg/113.0.1774.35"

const (
	RetryMaxAttempts = 3
	RetryBaseDelay   = time.Second
)

type HTTPStatusError struct {
	StatusCode int
	URL        string
}

func (e *HTTPStatusError) Error() string {
	return fmt.Sprintf("request to %s failed with status %d", e.URL, e.StatusCode)
}

func SendRequest[T any](ctx context.Context, client *http.Client, method string, url string, body io.Reader) (t T, err error) {
	req, err := http.NewRequestWithContext(ctx, method, url, body)
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

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusMultipleChoices {
		return t, &HTTPStatusError{
			StatusCode: res.StatusCode,
			URL:        url,
		}
	}

	err = json.NewDecoder(res.Body).Decode(&t)
	if err != nil {
		return t, err
	}

	return t, nil
}

func SendRequestWithRetry[T any](ctx context.Context, client *http.Client, method string, url string, body []byte) (t T, err error) {
	reqBody := getRequestBody(body)

	for attempt := range RetryMaxAttempts {
		t, err = SendRequest[T](ctx, client, method, url, reqBody)
		if err == nil {
			return t, nil
		}

		if !shouldRetry(err) || attempt == RetryMaxAttempts-1 {
			return t, err
		}

		backoff := RetryBaseDelay * time.Duration(1<<attempt) // exponential backoff
		timer := time.NewTimer(backoff)
		select {
		case <-ctx.Done():
			timer.Stop()
			return t, ctx.Err()
		case <-timer.C:
		}
	}

	return t, err
}

func getRequestBody(body []byte) io.Reader {
	if body == nil {
		return nil
	}

	return bytes.NewReader(body)
}

func shouldRetry(err error) bool {
	if err == nil {
		return false
	}

	if errors.Is(err, context.Canceled) {
		return false
	}

	if errors.Is(err, context.DeadlineExceeded) {
		return true
	}

	var netErr net.Error
	if errors.As(err, &netErr) {
		return netErr.Timeout()
	}

	var statusErr *HTTPStatusError
	if errors.As(err, &statusErr) {
		return statusErr.StatusCode == http.StatusTooManyRequests || statusErr.StatusCode >= http.StatusInternalServerError
	}

	return false
}
