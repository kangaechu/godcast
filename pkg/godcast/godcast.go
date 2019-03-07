package godcast

func Run(confFile string) {
	pc, err := ReadConfig(confFile)
	if err != nil {
		panic(err)
	}
	pc.Print()
}
