package godcast

import (
	"errors"
	"github.com/eduncan911/podcast"
	"log"
	"net/url"
	"path"
	"path/filepath"
	"strings"
	"time"
)

func GeneratePodcastXML(pc *PodcastConfig, tags []*AudioTag) (string, error) {
	now := time.Now()
	p := podcast.New(pc.Title, pc.Link, pc.Description, &now, &now)
	p.Language = "ja-jp"

	for _, t := range tags {
		item := podcast.Item{}
		item.Title = t.Title
		item.Description = t.Description
		link, err := pathToUrl(t.Filename, pc.EpisodeDir, pc.Link)
		if err != nil {
			log.Fatal("link create error", err, t)
		}
		item.Link = link

		item.PubDate = &t.PubDate

		enclosure := podcast.Enclosure{}
		enclosure.URL = item.Link
		enclosure.Length = t.Size

		et, err := fileNameToEnclosureType(t.Filename)
		if err != nil {
			return "", err
		}
		enclosure.Type = et
		item.Enclosure = &enclosure

		_, err = p.AddItem(item)
		if err != nil {
			// エラーとなったファイルは出力しない
			log.Print("could not write item:", t, err)
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

func fileNameToEnclosureType(filename string) (podcast.EnclosureType, error) {
	switch strings.ToUpper(filepath.Ext(filename)) {
	case ".M4A":
		return podcast.M4A, nil
	case ".MP4":
		return podcast.MP4, nil
	case ".MP3":
		return podcast.MP3, nil
	}
	return podcast.EPUB, errors.New("could not convert filename into EnclosureType")
}
