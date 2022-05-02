package db

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"github.com/golang/glog"
	"github.com/husong998/db_updater/app"
	"strings"
	"sync"
)

//go:embed template.sql
var template string

type Adapter struct {
	DB *sql.DB
}

func (o *Adapter) Upsert(ctx context.Context, records []app.Item) (err error) {
	tx, err := o.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	wg := sync.WaitGroup{}
	wg.Add((len(records) + 9999) / 10000)
	var rollback bool
	for i := 0; i < len(records); i += 10000 {
		r := i + 10000
		if len(records) < r {
			r = len(records)
		}
		l := i
		go func() {
			err = o.execOn(ctx, records[l:r])
			if err != nil {
				glog.Error(err)
				rollback = true
			}
			wg.Done()
		}()
	}
	wg.Wait()
	if rollback {
		err = tx.Rollback()
	} else {
		err = tx.Commit()
	}
	return
}

func (o *Adapter) execOn(ctx context.Context, records []app.Item) (err error) {
	var (
		sb   strings.Builder
		args []interface{}
	)
	args = make([]interface{}, 0, 3*len(records))
	for _, record := range records {
		sb.WriteString(template)
		args = append(args, record.ID, record.Price, record.Stock)
	}
	statements := strings.ReplaceAll(sb.String(), "?", "%v")
	statements = fmt.Sprintf(statements, args...)
	_, err = o.DB.ExecContext(ctx, statements)
	return
}
