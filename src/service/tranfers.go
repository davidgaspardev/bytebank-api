package service

import (
	database "bytebank-api/src/database"
	model "bytebank-api/src/models"
	"encoding/json"
	"log"
)

func AddTransfer(transfer model.Transfer) (string, error) {
	var object_id, err = database.AddData("transfer", transfer)
	if err != nil {
		return "", err
	}
	return object_id, nil
}

func GetAllTransfers() []model.Transfer {
	var data = database.GetAllData("transfer")

	var result = make([]model.Transfer, len(data))

	for i, d := range data {
		var dataJson, err = json.Marshal(d)
		if err != nil {
			log.Fatal(err)
		}

		json.Unmarshal(dataJson, &result[i])
	}

	return result
}
