package main

import (
	c "github.com/LeyouHong/device-listener/conf"
	"github.com/LeyouHong/device-listener/database"
	s "github.com/LeyouHong/device-listener/rpc/server"
)

func main() {
	db := database.NewDataBase()
	c.Conf.NewConf(db)
	defer database.CloseDataBase(db)

	s.Start()
}
