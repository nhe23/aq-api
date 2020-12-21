package aqdb

import (
	"context"
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

func LocationResults(take *int, after *string) ([]*model.LocationResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	cur, err := getPaginatedResult(measurementsCol, take, after)
	if err != nil {
		return nil, err
	}
	var resSlice []*model.LocationResult
	cur.All(ctx, &resSlice)

	return resSlice, nil
}

func GetCities(take *int, after *string) ([]*model.City, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cur, err := getPaginatedResult(citiesCol, take, after)
	if err != nil {
		return nil, err
	}
	var cities []*model.City
	cur.All(ctx, &cities)

	return cities, nil
}

func GetCountries() ([]*model.Country, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cur, err := getPaginatedResult(countriesCol, nil, nil)
	if err != nil {
		return nil, err
	}
	var countries []*model.Country
	cur.All(ctx, &countries)

	return countries, nil
}

func getPaginatedResult(col *mongo.Collection, take *int, after *string) (*mongo.Cursor, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
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
	cur, err := col.Find(ctx, filter, options)
	return cur, err
}
