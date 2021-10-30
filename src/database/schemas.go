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
	"required":             []string{"value", "contact", "datetime"},
	"additionalProperties": false,
	"properties": bson.M{
		"_id": bson.M{
			"bsonType": "objectId",
		},
		"value": bson.M{
			"bsonType": "double",
		},
		"contact": bson.M{
			"bsonType":             "object",
			"required":             []string{"name", "accountNumber"},
			"additionalProperties": false,
			"properties": bson.M{
				"name": bson.M{
					"bsonType": "string",
				},
				"accountNumber": bson.M{
					"bsonType": "int",
				},
			},
		},
		"datetime": bson.M{
			"bsonType": "date",
		},
	},
}
