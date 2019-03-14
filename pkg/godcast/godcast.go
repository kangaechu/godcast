package godcast

func Run(confFile string) {
	pc, err := ReadConfig(confFile)
	if err != nil {
		panic(err)
	}
	tags, err := GetTagsInDir(pc.EpisodeDir)
	if err != nil {
		panic(err)
	}
	generatedXML, err := GeneratePodcastXML(pc, tags)
	if err != nil {
		panic(err)
	}
	print(generatedXML)
}
