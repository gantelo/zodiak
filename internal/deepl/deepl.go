package deepl

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	httpi "zodiak/internal/http"
)

const (
	// V2 is the base url for v2 of the deepl API.
	V2 = "https://api-free.deepl.com/v2"
)

// A Client is a deepl client.
type Client struct {
	client       httpi.Client
	authKey      string
	baseURL      string
	translateURL string
}

// A ClientOption configures a Client.
type ClientOption func(*Client)

// A TranslateOption configures a translation request.
type TranslateOption func(url.Values)

// Error is a DeepL error.
type Error struct {
	// The HTTP error code, returned by the DeepL API.
	Code int

	Body []byte
}

type translateResponse struct {
	Translations []Translation `json:"translations"`
}

// Translation is a translation result from deepl.
type Translation struct {
	DetectedSourceLanguage string `json:"detected_source_language"`
	Text                   string `json:"text"`
}

// BaseURL returns a ClientOption that sets the base url for requests.
func BaseURL(url string) ClientOption {
	return func(c *Client) {
		c.baseURL = url
		c.translateURL = fmt.Sprintf("%s/translate", c.baseURL)
	}
}

// HTTPClient returns a ClientOption that specifies the http.Client that's used
// when making requests.
func HTTPClient(client httpi.Client) ClientOption {
	return func(c *Client) {
		c.client = client
	}
}

// SourceLang returns a ClientOption that specifies the source language of the
// input text. If SourceLang is not used, DeepL automatically figures out the
// source language.
func SourceLang(lang Language) TranslateOption {
	return func(vals url.Values) {
		vals.Set("source_lang", string(lang))
	}
}

// PreserveFormatting returns a TranslateOption that sets the
// `preserve_formatting` DeepL option.
func PreserveFormatting(preserve bool) TranslateOption {
	return func(vals url.Values) {
		vals.Set("preserve_formatting", boolString(preserve))
	}
}

// New returns a Client that uses authKey as the DeepL authentication key.
func New(authKey string, opts ...ClientOption) *Client {
	c := Client{
		authKey: authKey,
		client:  http.DefaultClient,
	}

	// default base url
	BaseURL(V2)(&c)

	for _, opt := range opts {
		opt(&c)
	}

	return &c
}

// HTTPClient returns the underlying http.Client.
func (c *Client) HTTPClient() httpi.Client {
	return c.client
}

// BaseURL returns the configured base url for requests.
func (c *Client) BaseURL() string {
	return c.baseURL
}

// AuthKey returns the DeepL authentication key.
func (c *Client) AuthKey() string {
	return c.authKey
}

func (c *Client) Translate(ctx context.Context, text string, targetLang Language, opts ...TranslateOption) (string, error) {
	translation, err := c.translate(ctx, text, targetLang, opts...)
	if err != nil {
		return "", fmt.Errorf("translate: %w", err)
	}

	return translation, nil
}

func (c *Client) translate(ctx context.Context, text string, targetLang Language, opts ...TranslateOption) (string, error) {
	vals := make(url.Values)
	vals.Set("text", text)
	vals.Add("target_lang", string(targetLang))

	for _, opt := range opts {
		opt(vals)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", c.translateURL, strings.NewReader(vals.Encode()))
	if err != nil {
		return "", fmt.Errorf("build request: %w", err)
	}

	req.Header.Add("Authorization", "DeepL-Auth-Key "+c.authKey)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("do request: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", Error{Code: resp.StatusCode}
	}

	var response translateResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", fmt.Errorf("decode deepl response: %w", err)
	}

	return response.Translations[0].Text, nil
}

func (err Error) Error() string {
	switch err.Code {
	case 456:
		return "Quota exceeded. The character limit has been reached."
	default:
		if len(err.Body) > 0 {
			return fmt.Sprintf("unexpected HTTP status %s (%s)",
				http.StatusText(err.Code),
				strings.TrimSpace(string(err.Body)))
		}
		return fmt.Sprintf("unexpected HTTP status %s",
			http.StatusText(err.Code))
	}
}

func boolString(b bool) string {
	if b {
		return "1"
	}
	return "0"
}
