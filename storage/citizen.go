package storage

import (
	"context"
	"github.com/decodeworms/citizenvisas/types"
	"time"
)

type DataStore struct {
	client *Client
	ctx context.Context
}

func NewDataStore(ctx context.Context, c *Client) DataStore {
	return DataStore{
		client: c,
		ctx: ctx,
	}
}

func (store DataStore) Create(data *types.Citizen) error {
	// NOTE: if you ever need to use the context, if you access it using this: store.ctx
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()

	citData := &types.Citizen{
		//ID:      data.ID,     DO NOT do this; we have already told postgres engine to increment this field automatically when we created the table
		Name:    data.Name,
		Gender:  data.Gender,
		Age:     data.Age,
		Country: data.Country,
	}
	return store.client.db.Debug().Create(&citData).Error  // we use Debug() to see how gorm executes SQL queries: it's meant for debugging only
}

func (store DataStore) Citizens() ([]*types.Citizen, error) {
	// NOTE: if you ever need to use the context, if you access it using this: store.ctx

	var res []*types.Citizen
	return res, store.client.db.Debug().Find(&res).Error
}

func (store DataStore) Citizen(id string) (*types.Citizen, error) {
	// NOTE: if you ever need to use the context, if you access it using this: store.ctx

	var res types.Citizen
	return &res, store.client.db.Debug().Where("id =?", id).Find(&res).Error
}

func (store DataStore) Update(data types.Citizen) error {
	// NOTE: if you ever need to use the context, if you access it using this: store.ctx

	data.UpdatedAt = time.Now()
	return store.client.db.Debug().Where("id = ?", data.ID).Update("name", data.Name).Error
}

func (store DataStore) Delete(id string) error {
	// NOTE: if you ever need to use the context, if you access it using this: store.ctx

	var data types.Citizen
	return store.client.db.Debug().Where("id = ?", id).Unscoped().Delete(&data).Error
}
