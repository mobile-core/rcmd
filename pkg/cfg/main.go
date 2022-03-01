package cfg

func Load() (nodes, error) {
	fileName := getFileName()
	cfg, err := configLoad(fileName)

	if err != nil {
		return cfg, err
	}
	return cfg, nil
}
