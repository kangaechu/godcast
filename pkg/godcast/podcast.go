package godcast

import (
	"github.com/eduncan911/podcast"
	"log"
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
		item.Link = t.Filename
		item.PubDate = &t.PubDate

		_, err := p.AddItem(item)
		if err != nil {
			log.Fatal("could not write item:", t, err)
		}
	}
	return p.String(), nil
}
