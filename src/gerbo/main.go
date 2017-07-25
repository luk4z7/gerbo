// Gerbo - Rodent and data extractor
// https://github.com/luk4z7/gerbo for the canonical source repository
// Copyright Lucas Alves 2017

package main

import (
	"gerbo/lib/logs"
	"gerbo/controllers/operation"
	"runtime"
	"gerbo/core/drive/sqlite"
	"database/sql"
	"sync"
	"time"
)

type Instance struct {
	sync.Mutex
	db *sql.DB
}

func main() {
	logs.Start()
	logs.INFO.Println("Running!")
	runtime.GOMAXPROCS(runtime.NumCPU())

	instance := Instance{}
	instance.db = sqlite.GetDB()
	defer instance.db.Close()

	done := make(chan struct{})
	go func() {
		done <- operation.Start(instance.db)
	}()
	time.Sleep(time.Second)
	instance.Lock()
	defer instance.Unlock()

	<-done
}