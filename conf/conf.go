package conf

import (
	"github.com/device-listener/util"
	"github.com/syndtr/goleveldb/leveldb"
	"os"
)

type Config struct {
	db   *leveldb.DB
	ip   string
	port string
}

var Conf = &Config{}

func (conf *Config) NewConf(db *leveldb.DB) {
	Conf.db = db
	Conf.ip = util.GetIntranetIp()
	Conf.port = os.Getenv("PORT")
}

func (conf *Config) GetDB() *leveldb.DB {
	return Conf.db
}

func (conf *Config) GetIP() string {
	return Conf.ip
}

func (conf *Config) GetPORT() string {
	return Conf.port
}
