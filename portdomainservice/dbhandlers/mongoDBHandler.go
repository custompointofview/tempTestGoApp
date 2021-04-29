package dbhandlers

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	srv "portdomainservice/service"
)

type BPort struct {
	Id          string    `bson:"_id"`
	Name        string    `bson:"name"`
	City        string    `bson:"city"`
	Country     string    `bson:"country"`
	Alias       []string  `bson:"alias"`
	Regions     []string  `bson:"regions"`
	Coordinates []float32 `bson:"coordinates"`
	Province    string    `bson:"province"`
	Timezone    string    `bson:"timezone"`
	Unlocs      []string  `bson:"ulocs"`
	Code        string    `bson:"code"`
}

type MongoDBHandler struct {
	client         *mongo.Client
	databaseName   string
	collectionName string
}

func NewMongoDBHandler(dbConfig *DBConfig) srv.DBHandlerInterf {
	return &MongoDBHandler{
		databaseName:   dbConfig.database,
		collectionName: dbConfig.collection,
	}
}

func (dbh *MongoDBHandler) Connect(ctx context.Context, endpoint string) error {
	// set client options
	clientOptions := options.Client().ApplyURI(endpoint)
	done := time.After(30 * time.Second)
	tick := time.NewTicker(2 * time.Second)

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-done:
			return fmt.Errorf("connection to endpoint was not established: %v", endpoint)
		case <-tick.C:
			// connect to MongoDB
			client, err := mongo.Connect(ctx, clientOptions)
			if err != nil {
				break
			}
			// check the connection
			err = client.Ping(ctx, nil)
			if err != nil {
				break
			}
			dbh.client = client
			log.Printf("Connected to DB endpoint @: %s", endpoint)
			return nil
		}
	}
}

func (dbh *MongoDBHandler) GetPort(ctx context.Context, port *srv.Port) (*srv.Port, error) {
	retPort := &srv.Port{}
	// setup query
	filter := bson.D{{"_id", port.Id}}
	// execute query
	coll := dbh.client.Database(dbh.databaseName).Collection(dbh.collectionName)
	err := coll.FindOne(ctx, filter).Decode(retPort)
	if err != nil {
		log.Printf("ERROR finding one: %v", err)
		return nil, err
	}
	return retPort, nil
}
func (dbh *MongoDBHandler) GetAllPorts(ctx context.Context) (map[string]*srv.Port, error) {
	return nil, nil
}
func (dbh *MongoDBHandler) CreateOrUpdatePort(ctx context.Context, port *srv.Port) (*srv.Port, error) {
	bport := dbh.convertPort2BPort(port)
	// setup query
	opts := options.Update().SetUpsert(true)
	filter := bson.D{{"_id", bport.Id}}
	update := bson.D{{"$set", bport}}

	// execute query
	coll := dbh.client.Database(dbh.databaseName).Collection(dbh.collectionName)
	_, err := coll.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return nil, err
	}
	return port, nil
}
func (dbh *MongoDBHandler) DeletePort(ctx context.Context, port *srv.Port) error {
	return nil
}

func (dbh *MongoDBHandler) convertPort2BPort(p *srv.Port) *BPort {
	return &BPort{
		Id:          p.Id,
		Name:        p.Name,
		City:        p.City,
		Country:     p.Country,
		Alias:       p.Alias,
		Regions:     p.Regions,
		Coordinates: p.Coordinates,
		Province:    p.Province,
		Timezone:    p.Timezone,
		Unlocs:      p.Unlocs,
		Code:        p.Code,
	}
}
