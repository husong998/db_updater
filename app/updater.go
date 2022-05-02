package app

import (
	"context"
	"fmt"
)

type DB interface {
	Upsert(ctx context.Context, records []Item) error
}

type Upserter struct {
	DB DB
}

type Item struct {
	ID    uint
	Price float64
	Stock uint
}

func (o *Upserter) Upsert(ctx context.Context, records []Item) (err error) {
	exist := make(map[uint]struct{})
	for _, item := range records {
		if _, ok := exist[item.ID]; ok {
			err = fmt.Errorf("duplicated product id %v", item.ID)
			return
		}
		exist[item.ID] = struct{}{}
	}
	err = o.DB.Upsert(ctx, records)
	return
}
