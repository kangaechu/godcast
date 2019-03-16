package godcast

import (
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)
import "github.com/dhowden/tag"

type AudioTag struct {
	Filename string
	Title    string
	Artist   string
	PubDate  time.Time
	Size     int64
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

	audio := AudioTag{filename, m.Title(), m.Artist(), stat.ModTime(), stat.Size()}
	return &audio, nil
}

type FileInfos []os.FileInfo
type ByModDate struct{ FileInfos }

func (fi ByModDate) Len() int {
	return len(fi.FileInfos)
}
func (fi ByModDate) Swap(i, j int) {
	fi.FileInfos[i], fi.FileInfos[j] = fi.FileInfos[j], fi.FileInfos[i]
}
func (fi ByModDate) Less(i, j int) bool {
	return fi.FileInfos[j].ModTime().Unix() < fi.FileInfos[i].ModTime().Unix()
}

func GetTagsInDir(dirname string) ([]*AudioTag, error) {
	dir, err := os.Open(dirname)
	defer dir.Close()
	if err != nil {
		return nil, err
	}
	files, err := dir.Readdir(0)
	if err != nil {
		return nil, err
	}
	var tags []*AudioTag

	sort.Sort(ByModDate{FileInfos{}})
	for _, file := range files {
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
