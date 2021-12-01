package database

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (_mongo *Mongo) AddData(collectionName string, data interface{}) (string, error) {
	result, err := _mongo.database.Collection(collectionName).InsertOne(*_mongo.context, data)
	if err != nil {
		return "", err
	}
	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (_mongo *Mongo) GetAllData(collectionName string) ([]bson.M, error) {
	cursor, err := _mongo.database.Collection(collectionName).Find(*_mongo.context, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(*_mongo.context)

	count, err := _mongo.database.Collection(collectionName).CountDocuments(*_mongo.context, bson.M{})
	if err != nil {
		return nil, err
	}

	allData := make([]bson.M, count)
	index := 0
	for cursor.Next(*instance.context) {
		var data bson.M
		err = cursor.Decode(&data)
		if err != nil {
			return nil, err
		}
		allData[index] = data
		index++
	}

	return allData, nil
}
