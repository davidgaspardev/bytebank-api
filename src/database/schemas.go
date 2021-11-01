// Schemas to data
package database

import "go.mongodb.org/mongo-driver/bson"

func getSchema(name string) bson.M {
	var schema bson.M
	switch name {
	case schemaList[0]:
		schema = transferSchema
		break
	default:
		return nil
	}
	return bson.M{
		"$jsonSchema": schema,
	}
}

// All schema to the database (Mongo db)
var schemaList = []string{
	"transfer",
}

var transferSchema = bson.M{
	"bsonType":             "object",
	"required":             []string{"value", "contact", "dateTime"},
	"description":          "Must be a Tranfer model object",
	"additionalProperties": false,
	"properties": bson.M{
		"_id": bson.M{
			"bsonType": "objectId",
		},
		"value": bson.M{
			"bsonType":    "double",
			"description": "Must be a double",
		},
		"contact": bson.M{
			"bsonType":             "object",
			"required":             []string{"fullname", "accountNumber"},
			"additionalProperties": false,
			"properties": bson.M{
				"fullname": bson.M{
					"bsonType":    "string",
					"description": "Must be a string",
				},
				"accountNumber": bson.M{
					"bsonType":    "int",
					"description": "Must be a integer",
				},
			},
		},
		"dateTime": bson.M{
			"bsonType":    "date",
			"description": "Must be a date",
		},
	},
}
