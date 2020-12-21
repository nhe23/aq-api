package aqdb

import (
	"context"
	"log"
	"time"

	"github.com/nhe23/aq-api/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Collections mongodb collections used in service
type Collections struct {
	MeasurementsCol *mongo.Collection
	CitiesCol       *mongo.Collection
	CountriesCol    *mongo.Collection
}

var measurementsCol *mongo.Collection
var citiesCol *mongo.Collection
var countriesCol *mongo.Collection

//
func SetCollections(cols *Collections) {
	countriesCol = cols.CountriesCol
	citiesCol = cols.CitiesCol
	measurementsCol = cols.MeasurementsCol
}

// func (db *DB) Save(input *model.) *model.Dog {
// 	collection := db.client.Database("animals").Collection("dogs")
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()
// 	res, err := collection.InsertOne(ctx, input)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return &model.Dog{
// 		ID:        res.InsertedID.(primitive.ObjectID).Hex(),
// 		Name:      input.Name,
// 		IsGoodBoi: input.IsGoodBoi,
// 	}
// }

// func (db *DB) FindByID(ID string) *model.Dog {
// 	ObjectID, err := primitive.ObjectIDFromHex(ID)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	collection := db.client.Database("animals").Collection("dogs")
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()
// 	res := collection.FindOne(ctx, bson.M{"_id": ObjectID})
// 	dog := model.Dog{}
// 	res.Decode(&dog)
// 	return &dog
// }

func LocationResults(take *int, after *string) []*model.LocationResult {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	options := options.Find()

	if take != nil {
		options.SetLimit(int64(*take))
	}

	var filter primitive.M
	if after != nil {
		docID, _ := primitive.ObjectIDFromHex(*after)
		filter = bson.M{"_id": bson.M{"$gt": docID}}
	} else {
		filter = bson.M{}
	}

	cur, err := measurementsCol.Find(ctx, filter, options)
	if err != nil {
		log.Fatal(err)
	}
	var locResults []*model.LocationResult
	for cur.Next(ctx) {
		var locRes *model.LocationResult
		err := cur.Decode(&locRes)
		if err != nil {
			log.Fatal(err)
		}
		locResults = append(locResults, locRes)
	}
	return locResults
}
