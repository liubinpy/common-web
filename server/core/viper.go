package core

import (
	"common-web/server/global"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

func Viper() {
	v := viper.New()

	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	v.SetConfigFile(filepath.Join(dir, "config.yaml"))

	err = v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err = v.Unmarshal(&global.CommonConfig); err != nil {
			fmt.Println(err)
		}
	})

	if err = v.Unmarshal(&global.CommonConfig); err != nil {
		fmt.Println(err)
	}

	return

}
