package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	if len(os.Args) == 2 {
		file := os.Args[1]
		fmt.Println("config file is " + file)
		config.Load(file)
	}

	g_aof.init(config.AsyncMode, config.AsyncFile)
	go g_aof.run()

	fmt.Println("start init dbs")
	dbmgr.init(config.DbNum)

	register_cmds()

	fmt.Println("start listen port: ", config.Port)
	l := &Listener{}
	l.init(config.Port)

    go l.run()

    go dbmgr.run()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	<-sigs

    l.shutdown()

    dbmgr.shutdown()
}
