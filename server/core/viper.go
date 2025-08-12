package core

import (
	"flag"
	"fmt"
	"os"

	"blog/server/core/internal"
	"blog/server/global"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func getConfigPath() (config string) {
	flag.StringVar(&config, "c", "", "输入配置文件如 go run main.go -c config.yaml")
	flag.Parse() // 必须调用 flag.Parse()否则参数无法解析 命令行参数的顺序可以任意
	if config != "" {
		fmt.Println("配置路径： ", config)
		return
	}

	if env := os.Getenv(internal.ConfigEnv); env != "" {
		config = env
		fmt.Println("环境配置： ", config, internal.ConfigEnv)
		return
	}

	switch gin.Mode() {
	case gin.DebugMode:
		config = internal.ConfigDefaultFile
		//todo: 测试配置， 发布配置
	default:
		config = internal.ConfigDefaultFile
	}
	_, err := os.Stat(config)
	if err != nil || os.IsNotExist(err) {
		config = internal.ConfigDefaultFile
		fmt.Println("--- 使用默认配置  ---  ", config)
	}
	return
}

func Viper() *viper.Viper {
	config := getConfigPath()
	v := viper.New()        // viper.New()创建一个新的 Viper 实例
	v.SetConfigFile(config) // 设置配置文件路径
	v.SetConfigType("yaml") // 明确指定配置格式为 YAML
	err := v.ReadInConfig() // 读取配置文件内容
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()                           // 启用文件系统监听，当配置文件被修改时自动触发回调
	v.OnConfigChange(func(e fsnotify.Event) { // 注册变更回调函数。当文件变化时
		fmt.Println("config file changed:", e.Name)
		// 回调时执行
		if err = v.Unmarshal(&global.GVA_CONFIG); err != nil { // 重新解析配置到全局结构体 global.GVA_CONFIG
			fmt.Println(err)
		}
	})
	// 第一次加载时执行
	if err := v.Unmarshal(&global.GVA_CONFIG); err != nil { // 将配置文件内容解析到全局变量 global.GVA_CONFIG
		panic(fmt.Errorf("fatal error unmarshal config: %w", err))
	}
	return v
}
