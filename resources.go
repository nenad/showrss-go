package showrss

import (
	"encoding/xml"
)

type Item struct {
	Text  string `xml:",chardata"`
	Title string `xml:"title"`
	Link  string `xml:"link"`
	Guid  struct {
		Text        string `xml:",chardata"`
		IsPermalink string `xml:"isPermaLink,attr"`
	} `xml:"guid"`
	PubDate     string `xml:"pubDate"`
	Description string `xml:"description"`
	ShowID      string `xml:"show_id"`
	ExternalID  string `xml:"external_id"`
	ShowName    string `xml:"show_name"`
	EpisodeID   string `xml:"episode_id"`
	RawTitle    string `xml:"raw_title"`
	InfoHash    string `xml:"info_hash"`
	Enclosure   struct {
		Text   string `xml:",chardata"`
		URL    string `xml:"url,attr"`
		Length string `xml:"length,attr"`
		Type   string `xml:"type,attr"`
	} `xml:"enclosure"`
}

type Rss struct {
	XMLName xml.Name `xml:"rss"`
	Text    string   `xml:",chardata"`
	Version string   `xml:"version,attr"`
	Tv      string   `xml:"tv,attr"`
	Channel struct {
		Text        string `xml:",chardata"`
		Title       string `xml:"title"`
		Link        string `xml:"link"`
		Ttl         string `xml:"ttl"`
		Description string `xml:"description"`
		Items       []Item `xml:"item"`
	} `xml:"channel"`
}

type Episode struct {
	ShowName string
	Name     string
	Episode  int
	Season   int
	Magnet   string
	Quality  string
}
