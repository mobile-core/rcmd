package cfg

func Load() node {
	fileName := getFileName()
	return configLoad(fileName)
}
