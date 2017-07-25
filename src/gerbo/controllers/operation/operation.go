// Gerbo - Rodent and data extractor
// https://github.com/luk4z7/gerbo for the canonical source repository
// Copyright Lucas Alves 2017

package operation

import (
	"database/sql"
	"time"
	"gerbo/services/operation"
)

func Start(db *sql.DB) struct{} {
	for {
		select {
			case <-time.After(time.Second * 3):
				operation.CheckSync(db)
		}
	}
}