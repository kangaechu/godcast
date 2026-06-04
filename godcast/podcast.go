package godcast

import (
	"errors"
	"log"
	"net/url"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/eduncan911/podcast"
)

func GeneratePodcastXML(pc *PodcastConfig, tags []*AudioTag) (string, error) {
	now := time.Now()
	p := podcast.New(pc.Title, pc.Link, pc.Description, &now, &now)
	p.Language = "ja-jp"
	if pc.ImageURL != "" {
		p.AddImage(pc.ImageURL)
	}

	// Apple Podcasts はフィード登録に itunes:category と itunes:explicit を必須とする。
	// 設定が未指定の場合はデフォルト値を補完する。
	category := pc.Category
	if category == "" {
		category = "Leisure"
	}
	p.AddCategory(category, nil)

	explicit := pc.Explicit
	if explicit == "" {
		explicit = "false"
	}
	p.IExplicit = explicit

	if pc.Author != "" {
		p.IAuthor = pc.Author
	}

	for _, t := range tags {
		item := podcast.Item{}
		item.Title = t.Title
		item.Description = t.Description
		link, err := pathToURL(t.Filename, pc.EpisodeDir, pc.Link)
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

func pathToURL(soundFilePath string, basePath string, urlBase string) (string, error) {
	relativePath, err := filepath.Rel(basePath, soundFilePath)
	if err != nil {
		return "", err
	}
	baseURL, err := url.Parse(urlBase)
	if err != nil {
		return "", err
	}
	baseURL.Path = path.Join(baseURL.Path, relativePath)
	return baseURL.String(), nil
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
