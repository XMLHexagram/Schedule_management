//读取初始化数据，并储存在serve实例中
package init

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"main/deal_error"
)

type DBConfig struct {
	Address  string
	User     string
	Password string
	DBName   string
}

type webConfig struct {
	Port string
}

type Conf struct {
	DB  DBConfig
	Web webConfig
}

func (s *Serve) ConfigInit() {
	confPath := "../SMS_config/conf.toml"

	//c := new(Conf)

	_, err := toml.DecodeFile(confPath, &s.Config)
	dealError.DealError(err)
	fmt.Println(s.Config)

	//s.Config = *c
}
