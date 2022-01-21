package cfg

func Load() (node, error) {
	fileName := getFileName()
	cfg, err := configLoad(fileName)

	if err != nil {
		return cfg, err
	}
	return cfg, nil
}
