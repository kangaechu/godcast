package godcast

import "os"

func Run(confFile string) {
	// 設定を読み込む
	pc, err := ReadConfig(confFile)
	if err != nil {
		panic(err)
	}

	//Podcastのファイルからタグを検索
	tags, err := GetTagsInDir(pc.EpisodeDir)
	if err != nil {
		panic(err)
	}

	// XMLの生成
	generatedXML, err := GeneratePodcastXML(pc, tags)
	if err != nil {
		panic(err)
	}

	// ファイルに出力
	file, err := os.Create(pc.PodcastFilePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.Write(([]byte)(generatedXML))
	if err != nil {
		panic(err)
	}
}
