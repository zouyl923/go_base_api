package file

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func LoadConfig(confDir string, conf interface{}) (interface{}, error) {
	yamlFile, err := os.ReadFile(confDir)
	if err != nil {
		fmt.Println("解析yaml文件失败：", err)
		return nil, err
	}
	yaml.Unmarshal(yamlFile, &conf)
	return conf, nil
}
