package cfg

import (
	"os"

	"github.com/mobile-core/rcmd/pkg/fileutil"
	"gopkg.in/yaml.v2"
)

// getFileName gets the configuration file name.
func getFileName() string {
	const configFileName string = ".rcmd.yml"
	homeDir := fileutil.GetHomedir()
	separate := fileutil.GetSeparate()
	fileName := homeDir + separate + configFileName
	return fileName
}

// configLoad loads the configuration file.
func configLoad(fileName string) (node, error) {
	node := &node{}
	b, _ := os.ReadFile(fileName)

	if err := yaml.Unmarshal(b, &node); err != nil {
		return *node, err
	}
	return *node, nil
}
