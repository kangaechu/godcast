package godcast

import (
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/dhowden/tag"
)

type AudioTag struct {
	Filename    string
	Title       string
	Artist      string
	Description string
	PubDate     time.Time
	Size        int64
}

func GetTag(filename string) (*AudioTag, error) {
	log.Println("Reading", filename)
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		return nil, err
	}

	m, err := tag.ReadFrom(file)
	if err != nil {
		log.Fatal(err)
	}
	stat, err := os.Stat(filename)
	if err != nil {
		log.Fatal("error on read file stat", filename)
	}
	comment := m.Comment()
	if m.Comment() == "" {
		comment = "no description"
	}
	audio := AudioTag{filename, m.Title(), m.Artist(), comment, stat.ModTime(), stat.Size()}
	return &audio, nil
}

func GetTagsInDir(dirname string) ([]*AudioTag, error) {
	dir, err := os.Open(dirname)
	defer dir.Close()
	if err != nil {
		return nil, err
	}
	fileInfos, err := dir.Readdir(0)
	if err != nil {
		return nil, err
	}
	var tags []*AudioTag

	sort.Slice(fileInfos, func(i, j int) bool {
		return fileInfos[i].ModTime().Unix() < fileInfos[j].ModTime().Unix()
	})
	for _, file := range fileInfos {
		if strings.HasSuffix(file.Name(), "mp3") ||
			strings.HasSuffix(file.Name(), "mp4") ||
			strings.HasSuffix(file.Name(), "m4a") {
			t, err := GetTag(filepath.Join(dirname, file.Name()))
			if err != nil {
				log.Fatal("error on ", file.Name())
				continue
			}
			tags = append(tags, t)
		}
	}
	return tags, nil
}
