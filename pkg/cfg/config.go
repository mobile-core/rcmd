package cfg

import (
	"os"

	"gopkg.in/yaml.v2"
)

// getFileName gets the configuration file name.
func getFileName() string {
	const configFileName string = ".rcmd.yml"
	homeDir := GetHomedir()
	separate := GetSeparate()
	fileName := homeDir + separate + configFileName
	return fileName
}

// configLoad loads the configuration file.
func configLoad(fileName string) node {
	node := &node{}
	b, err := os.ReadFile(fileName)
	if err != nil {
		panic(err.Error())
	}

	err = yaml.Unmarshal(b, &node)
	if err != nil {
		panic(err.Error())
	}
	return *node
}
