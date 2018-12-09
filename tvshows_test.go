package showrss

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func NewTestClient(fn TestRoundTripFunc) *Client {
	internal := &http.Client{Transport: fn}
	return &Client{internal: internal}

}

type TestRoundTripFunc func(req *http.Request) *http.Response

func (rt TestRoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return rt(req), nil
}

func TestClient_GetTVShowEpisodes(t *testing.T) {
	c := NewTestClient(func(r *http.Request) *http.Response {
		assert.Equal(t, "https://showrss.info/show/1000.rss", r.URL.String())
		assert.Equal(t, "GET", r.Method)
		xml, _ := ioutil.ReadFile("fixtures/tvshow_1000.xml")
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(xml)),
			Header:     map[string][]string{"Content-Type": {"text/xml; charset=UTF-8"}},
		}
	})

	episodes, err := c.GetTVShowEpisodes(1000)
	assert.NoError(t, err)

	expected, _ := ioutil.ReadFile("fixtures/tvshow_1000.expected")
	builder := strings.Builder{}
	for _, e := range episodes {
		builder.WriteString(fmt.Sprintf("%+v\n", e))
	}

	assert.Equal(t, string(expected), builder.String())
}

func TestClient_GetPersonalEpisodes(t *testing.T) {
	c := NewTestClient(func(r *http.Request) *http.Response {
		assert.Equal(t, "http://showrss.info/user/123123.rss?magnets=true&namespaces=true&name=null&quality=null&re=null", r.URL.String())
		assert.Equal(t, "GET", r.Method)
		xml, _ := ioutil.ReadFile("fixtures/myfeed.xml")
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(xml)),
			Header:     map[string][]string{"Content-Type": {"text/xml; charset=UTF-8"}},
		}
	})

	episodes, err := c.GetPersonalEpisodes("http://showrss.info/user/123123.rss?magnets=true&namespaces=true&name=null&quality=null&re=null")
	assert.NoError(t, err)

	expected, _ := ioutil.ReadFile("fixtures/myfeed.expected")
	builder := strings.Builder{}
	for _, e := range episodes {
		builder.WriteString(fmt.Sprintf("%+v\n", e))
	}

	assert.Equal(t, string(expected), builder.String())
}
