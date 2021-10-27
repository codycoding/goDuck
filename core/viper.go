package core

import (
	"flag"
	"fmt"
	"github.com/codycoding/mini-go-app/global"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

//
// Viper
//  @Description: 根据命令行参数读取指定配置或读取默认配置   默认配置文件名称: config.yaml
//  @return *viper.Viper
//
func Viper() *viper.Viper {
	var config string
	flag.StringVar(&config, "c", "", "choose config file.")
	flag.Parse()
	if config == "" { // 优先级: 命令行 > 默认值
		config = "config.yaml"
		fmt.Printf("您正在使用config的默认值,config的路径为%v\n", config)

	} else {
		fmt.Printf("您正在使用命令行的-c参数传递的值,config的路径为%v\n", config)
	}

	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&global.Config); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&global.Config); err != nil {
		fmt.Println(err)
	}
	return v
}
