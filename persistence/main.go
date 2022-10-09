package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

type PeaceRepository struct {
	DB *mongo.Database
}

func DatabaseConnection(connectionURL string) (*mongo.Client, error) {
	opts := options.Client().ApplyURI(connectionURL)
	readPreference := options.Client().SetReadPreference(readpref.Primary())
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, opts, readPreference)
	if err != nil {
		return client, err
	}
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		return client, err
	}
	return client, err
}

func (d *PeaceRepository) NewDatabase(client *mongo.Client, databaseName string) *PeaceRepository {
	db := client.Database(databaseName)
	return &PeaceRepository{DB: db}
}

func (d *PeaceRepository) CreateIndex(ctx context.Context) error {
	indexOpts := options.CreateIndexes().SetMaxTime(time.Second * 10)
	userIndexModel := mongo.IndexModel{
		Options: options.Index().SetBackground(true),
		Keys: bsonx.Doc{{"_id", bsonx.Int32(1)}, {"email", bsonx.Int32(1)},
			{"city", bsonx.Int32(1)}},
	}
	// userLocationModel := mongo.IndexModel{}
	userIndex := d.DB.Collection("users").Indexes()
	_, err := userIndex.CreateMany(ctx, []mongo.IndexModel{userIndexModel}, indexOpts)
	if err != nil {
		return err
	}
	return nil
}
