package conf

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/util"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

type config struct {
	Mongodb *mongodb `yaml:"Mongodb"`
	Redis   *redis   `yaml:"Redis"`
	Mysql   *mysql   `yaml:"Mysql"`
	Wexin   *wexin   `yaml:"Wexin"`
	Neo4j   *neo4j   `yaml:"Neo4j"`
}

type neo4j struct {
	DbUri    string
	UserName string
	Password string
}

type mongodb struct {
	Uri            string `yaml:"Uri"`
	Database       string `yaml:"Database"`
	ConnectTimeout int64  `yaml:"ConnectTimeout"`
}

type redis struct {
	Addr     string `yaml:"Addr"`
	Password string `yaml:"Password"`
	DB       string `yaml:"DB"`
}

type wexin struct {
	AppId                       string `yaml:"AppId"`
	Secret                      string `yaml:"Secret"`
	GetUserOpenIdApiTemplateUrl string `yaml:"GetUserOpenIdApiTemplateUrl"`
}

type mysql struct {
	User     string `yaml:"User"`
	Password string `yaml:"Password"`
	Host     string `yaml:"Host"`
	Port     string `yaml:"Port"`
	DbName   string `yaml:"DbName"`
}

var Config *config

// InitViper 初始化Viper
func InitConfig() {
	ctx := context.Background()
	viper.SetConfigType("yaml")
	runEnv := os.Getenv("RUN_ENV")
	confPath := util.GetAbsolutePath() + "../../conf"
	if runEnv == consts.RUN_ENV_LOCAL {
		viper.SetConfigFile(filepath.Join(confPath, "config.local.yaml"))
	} else if runEnv == consts.RUN_ENV_PROD {
		viper.SetConfigFile(filepath.Join(confPath, "config.prod.yaml"))
	} else {
		viper.SetConfigFile(filepath.Join(confPath, "config.dev.yaml"))
	}
	hlog.CtxInfof(ctx, "[Config] env:%s", runEnv)

	if err := viper.ReadInConfig(); err != nil {
		hlog.CtxErrorf(ctx, "[Config] ReadInConfig failed, err: %v", err)
	}

	if err := viper.Unmarshal(&Config); err != nil {
		hlog.CtxErrorf(ctx, "[Config] Unmarshal failed, err: %v", err)
	}

	hlog.CtxInfof(ctx, "[Config] config.Mongodb: %#v", Config.Mongodb)
	hlog.CtxInfof(ctx, "[Config] config.Mysql: %#v", Config.Mysql)
	hlog.CtxInfof(ctx, "[Config] config.Redis: %#v", Config.Redis)
}

func GetConfig() *config {
	return Config
}
