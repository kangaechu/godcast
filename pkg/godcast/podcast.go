package godcast

import (
	"github.com/eduncan911/podcast"
	"log"
	"net/url"
	"path"
	"path/filepath"
	"time"
)

func GeneratePodcastXML(pc *PodcastConfig, tags []*AudioTag) (string, error) {
	now := time.Now()
	p := podcast.New(pc.Title, pc.Link, pc.Description, &now, &now)
	p.Language = "ja-jp"

	for _, t := range tags {
		item := podcast.Item{}
		item.Title = t.Title
		item.Description = "no description"
		link, err := pathToUrl(t.Filename, pc.EpisodeDir, pc.Link)
		if err != nil {
			log.Fatal("link create error", err, t)
			continue
		}
		item.Link = link

		item.PubDate = &t.PubDate

		_, err = p.AddItem(item)
		if err != nil {
			log.Fatal("could not write item:", t, err)
		}
	}
	return p.String(), nil
}

func pathToUrl(soundFilePath string, basePath string, urlBase string) (string, error) {
	relativePath, err := filepath.Rel(basePath, soundFilePath)
	if err != nil {
		return "", err
	}
	baseUrl, err := url.Parse(urlBase)
	if err != nil {
		return "", err
	}
	baseUrl.Path = path.Join(baseUrl.Path, relativePath)
	return baseUrl.String(), nil
}
