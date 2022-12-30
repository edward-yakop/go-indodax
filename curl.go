package indodax

import (
	"context"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// callPublic call public API with specific method and parameters.
// On success, it will return response body.
// On fail, it will return an empty body with an error.
func (cl *Client) curlPublic(ctx context.Context, urlPath string) (body []byte, err error) {
	httpUrl, err := url.Parse(cl.env.BaseHostPublic + urlPath)
	if err != nil {
		return nil, fmt.Errorf("curlPublic: " + err.Error())
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, httpUrl.String(), http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("curlPublic: " + err.Error())
	}
	req.Header = http.Header{
		"Content-Type": []string{
			"application/x-www-form-urlencoded",
		},
	}

	res, err := cl.conn.Do(req)
	if err != nil {
		return nil, fmt.Errorf("curlPublic: " + err.Error())
	}

	body, err = io.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("curlPublic: " + err.Error())
	}

	printDebug(string(body))

	return body, nil
}

// callPrivate call private API with specific method and parameters.
// On success, it will return response body.
// On fail, it will return an empty body with an error.
func (cl *Client) curlPrivate(ctx context.Context, method string, params url.Values) (
	body []byte, err error,
) {
	req, err := cl.newPrivateRequest(ctx, method, params)
	if err != nil {
		return nil, fmt.Errorf("curlPrivate: " + err.Error())
	}

	res, err := cl.conn.Do(req)
	if err != nil {
		return nil, fmt.Errorf("curlPrivate: " + err.Error())
	}

	body, err = io.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("curlPrivate: " + err.Error())
	}

	printDebug(string(body))

	return body, nil
}

// newPrivateRequest is method to generate authentication for private API.
// On success, it will return http request.
// On fail, it will return an error.
func (cl *Client) newPrivateRequest(
	ctx context.Context, apiMethod string, params url.Values,
) (req *http.Request, err error) {
	query := url.Values{
		"timestamp": []string{
			timestampAsString(),
		},
		"method": []string{
			apiMethod,
		},
	}

	virtualParams := map[string][]string(params)
	for k, v := range virtualParams {
		if len(v) > 0 {
			query.Set(k, v[0])
		}
	}

	reqBody := query.Encode()

	printDebug(fmt.Sprintf("newPrivateRequest >> request body:%s", reqBody))

	httpUrl, err := url.Parse(cl.env.BaseHostPrivate)
	if err != nil {
		err = fmt.Errorf("newPrivateRequest: " + err.Error())
		return nil, err
	}

	req, err = http.NewRequestWithContext(ctx, http.MethodPost, httpUrl.String(), io.NopCloser(strings.NewReader(reqBody)))
	sign := cl.encodeToHmac512(reqBody)
	req.Header = http.Header{
		"Content-Type": []string{
			"application/x-www-form-urlencoded",
		},
		"Key": []string{
			cl.env.apiKey,
		},
		"Sign": []string{
			sign,
		},
	}

	return req, nil
}

func (cl *Client) encodeToHmac512(param string) string {
	sign := hmac.New(sha512.New, []byte(cl.env.apiSecret))

	sign.Write([]byte(param))

	return hex.EncodeToString(sign.Sum(nil))
}
