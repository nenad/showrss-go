package showrss

import (
	"regexp"
	"strconv"
	"strings"
)

func extractEpisodes(feed Rss) []Episode {
	episodes := make([]Episode, len(feed.Channel.Items))
	for i, item := range feed.Channel.Items {
		episodes[i].Season, episodes[i].Episode = extractEpisodeSeason(item)
		episodes[i].Quality = extractQuality(item)
		episodes[i].ShowName = item.ShowName
		episodes[i].Magnet = item.Link
		episodes[i].Name = extractName(item)
	}

	return episodes
}

var episodeSeasonRegexes = []*regexp.Regexp{
	regexp.MustCompile("[sS]([0-9]{1,2})[eE]([0-9]{1,2})"),
	regexp.MustCompile("([0-9]{1,2})x([0-9]{1,2})"),
}

func extractEpisodeSeason(item Item) (season int, episode int) {
	for _, regex := range episodeSeasonRegexes {
		matches := regex.FindAllStringSubmatch(item.Title, -1)
		if len(matches) != 1 {
			continue
		}

		season, _ = strconv.Atoi(matches[0][1])
		episode, _ = strconv.Atoi(matches[0][2])
		return season, episode
	}

	return 0, 0
}

var qualityRegex = regexp.MustCompile("1080p|720p")

func extractQuality(item Item) string {
	matches := qualityRegex.FindAllStringSubmatch(item.Title, -1)
	if len(matches) != 1 {
		return "SD"
	}

	return matches[0][0]
}

func extractName(item Item) string {
	name := item.Title
	for _, regex := range episodeSeasonRegexes {
		name = regex.ReplaceAllString(name, "")
	}
	name = qualityRegex.ReplaceAllString(name, "")
	name = strings.Replace(name, "REPACK", "", -1)
	name = strings.Replace(name, item.ShowName, "", 1)
	name = strings.TrimSpace(name)

	if name == "" {
		return "-"
	}
	return name
}
