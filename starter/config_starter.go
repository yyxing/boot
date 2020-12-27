package starter

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/yyxing/boot"
	"github.com/yyxing/boot/context"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type ConfigStarter struct {
	ConfigPath string
	AbstractStarter
}

const (
	GlobalConfigKey = "SystemConfig"
)

var (
	localConfig viper.Viper
)

func pathExists(path string) (bool, int) {
	_, err := os.Stat(path)
	if err == nil {
		files, _ := ioutil.ReadDir(path)
		return true, len(files)
	}
	if os.IsNotExist(err) {
		return false, 0
	}
	return false, 0

}
func (config *ConfigStarter) Init(context context.ApplicationContext) {
	configPath := "./resource"
	exists, fileCount := pathExists(configPath)
	if !exists {
		logrus.Warnf("You should store the configuration in the .resource directory")
		return
	}
	if fileCount <= 0 {
		logrus.Warnf("You should store a configuration file in the .resource directory")
		return
	}
	v := viper.New()
	if config.ConfigPath != "" && len(config.ConfigPath) > 0 {
		if strings.Contains(config.ConfigPath, string(filepath.Separator)) {
			v.SetConfigFile(config.ConfigPath)
		} else {
			v.AddConfigPath(configPath)
			v.SetConfigName(config.ConfigPath)
		}
	} else {
		files, err := ioutil.ReadDir(configPath)
		if err != nil {
			log.Println("Find config files failed use default config")
		}
		v.AddConfigPath(configPath)
		suffix := ""
		for _, file := range files {
			if !file.IsDir() {
				suffix = path.Ext(file.Name())[1:]
				switch suffix {
				case "yml", "yaml", "properties", "ini":
					v.SetConfigName(file.Name())
					v.SetConfigType(suffix)
				default:
					logrus.Errorf("config %s type can not parse", file.Name())
				}
			}
		}
	}
	if err := v.ReadInConfig(); err != nil {
		panic("read config failed error message:" + err.Error())
	}
	context.Set(GlobalConfigKey, *v)
	localConfig = *v
	log.Println("config init success")
}
func GetConfig() viper.Viper {
	return localConfig
}
func (config *ConfigStarter) Finalize(context context.ApplicationContext) {
	context.Remove(GlobalConfigKey)
}

func (config *ConfigStarter) GetOrder() int {
	return boot.Int32Min
}
