package datastore

import (
	"context"

	"fmt"
	"github.com/pkg/errors"
	"gopkg.in/olivere/elastic.v6"
)

var client *elastic.Client
var AutoBulk BulkIndexer

var mappings = make(map[string]string)

func Setup(ctx context.Context) error {

	var err error
	client, err = elastic.NewClient()
	if err != nil {
		return errors.Wrap(err, "datastore.setup.newClient")
	}

	for index, mapping := range mappings {
		err := createIndex(ctx, index, mapping)
		if err != nil {
			return errors.Wrap(err, fmt.Sprint("datastore.setup.createIndex", index))
		}
	}

	AutoBulk = BeginBulkIndexer()

	return nil
}

func RegisterMapping(index string, mapping string) error {
	mappings[index] = mapping
	if client != nil {
		return createIndex(context.TODO(), index, mapping)
	}
	return nil
}

func createIndex(ctx context.Context, index string, mapping string) error {
	exists, err := client.IndexExists(index).Do(ctx)
	if err != nil {
		return errors.Wrap(err, "index existence check failure")
	}

	if !exists {
		createIndex, err := client.CreateIndex(index).BodyString(mapping).Do(ctx)
		if err != nil {
			return errors.Wrap(err, "create index failed")
		}
		if !createIndex.Acknowledged {
			return errors.New("create index not acknowledged")
		}
	}

	return nil
}

func Client() *elastic.Client {
	return client
}
