package audit

import (
	"context"
	"fmt"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

func (s Service) indexDB() {
	commonIndex := []mongo.IndexModel{
		newIndex("source", "target", "targetId"),
	}

	// Index all allowed sources
	for _, target := range s.Targets {
		// Index
		indexCol(s.DB.Collection(getColName(s.Source, target)), commonIndex)
	}
}

// NewIndex ...
func newIndex(key ...string) mongo.IndexModel {
	var doc bsonx.Doc
	for _, s := range key {
		e := bsonx.Elem{
			Key:   s,
			Value: bsonx.Int32(1),
		}
		if strings.HasPrefix(s, "-") {
			e = bsonx.Elem{
				Key:   strings.Replace(s, "-", "", 1),
				Value: bsonx.Int32(-1),
			}
		}
		doc = append(doc, e)
	}

	return mongo.IndexModel{Keys: doc}
}

// index collection
func indexCol(col *mongo.Collection, indexes []mongo.IndexModel) {
	opts := options.CreateIndexes().SetMaxTime(time.Minute * 10)
	_, err := col.Indexes().CreateMany(context.Background(), indexes, opts)
	if err != nil {
		fmt.Printf("Index collection %s err: %v", col.Name(), err)
	}
}
