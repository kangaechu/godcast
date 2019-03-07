package godcast

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type PodcastConfig struct {
	Title           string `yaml:"title"`
	Link            string `yaml:"link"`
	Description     string `yaml:"description"`
	EpisodeDir      string `yaml:"episodes_directory"`
	PodcastFilePath string `yaml:"podcast_file_path"`
}

func ReadConfig(confFile string) (pc *PodcastConfig, err error) {
	file, err := os.Open(confFile)
	defer file.Close()
	if err != nil {
		// ファイルがない場合は空のDownloadedProgramsを返す
		return nil, nil
	}
	readBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(readBytes, &pc)
	if err != nil {
		return nil, err
	}
	return pc, nil
}

func (pc PodcastConfig) Print() {
	fmt.Println("title:            ", pc.Title)
	fmt.Println("link:             ", pc.Link)
	fmt.Println("description:      ", pc.Description)
	fmt.Println("episodes_dir:     ", pc.EpisodeDir)
	fmt.Println("podcast_file_path:", pc.PodcastFilePath)
}
