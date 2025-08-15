package third

import (
	"flag"
	"fmt"
	"os"
)

func getConfigPath() (config string) {
	flag.StringVar(&config, "c", "", "命令： -c /xx/config.yaml")
	flag.Parse()
	if config != "" {
		fmt.Println("config: --> ", config)
		return
	}
	// 加载系统环境变量：CONFIG_PATH
	if env := os.Getenv("CONFIG_PATH"); env != "" {
		config = env
		fmt.Println("config: --> ", config)
		return
	}
	// todo: gin.Mode()
	// switch gin.Model() {}
	config = "config.yaml"
	_, err := os.Stat(config)
	if err != nil || os.IsNotExist(err) {
		// todo:
	}
	return
}

func Vipper() *viper.Viper {

}
