// Gerbo - Rodent and data extractor
// https://github.com/luk4z7/gerbo for the canonical source repository
// Copyright Lucas Alves 2017

package operation

import (
	"gerbo/core/drive/sqlite"
	"gerbo/services/operation"
	"gerbo/lib/logs"
	"os"
	"sync"
)

func Start() {
	if len(os.Args) == 2 {
		if os.Args[1] == "--sync" {
			logs.INFO.Println("Checking to sync...")

			db := sqlite.GetDB()
			defer db.Close()

			done := make(chan struct{})
			mu := &sync.Mutex{}

			go func() {
				operation.CheckSync(db, mu)
				done <- struct{}{}
			}()
			<-done
		}
	}
}