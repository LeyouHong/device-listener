package main

import (
	c "github.com/device-listener/conf"
	"github.com/device-listener/database"
	s "github.com/device-listener/rpc/server"
)

func main() {
	db := database.NewDataBase()
	c.Conf.NewConf(db)
	defer database.CloseDataBase(db)

	s.Start()
}
