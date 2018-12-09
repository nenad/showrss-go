package showrss

import (
	"encoding/xml"
	"fmt"
)

const (
	TVShowUrl = ApiBaseUrl + "/show/%d.rss"
)

func (c *Client) GetTVShowEpisodes(id int) (episodes []Episode, err error) {
	resp, err := c.internal.Get(fmt.Sprintf(TVShowUrl, id))
	feed := Rss{}
	err = xml.NewDecoder(resp.Body).Decode(&feed)
	episodes = extractEpisodes(feed)
	return episodes, err
}

func (c *Client) GetPersonalEpisodes(url string) (episodes []Episode, err error) {
	resp, err := c.internal.Get(url)
	feed := Rss{}
	err = xml.NewDecoder(resp.Body).Decode(&feed)
	episodes = extractEpisodes(feed)
	return episodes, err
}
