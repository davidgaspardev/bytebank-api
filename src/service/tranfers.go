package service

import (
	database "bytebank-api/src/database"
	model "bytebank-api/src/models"
	"encoding/json"
	"log"
)

func AddTransfer(transfer *model.Transfer) error {
	var object_id, err = database.LoadMongo().AddData("transfer", *transfer)
	if err != nil {
		return err
	}
	transfer.Id = object_id
	return nil
}

func GetAllTransfers() []model.Transfer {
	data, err := database.LoadMongo().GetAllData("transfer")
	if err != nil {
		return []model.Transfer{}
	}

	var result = make([]model.Transfer, len(data))

	for i, d := range data {
		var dataJson, err = json.Marshal(d)
		if err != nil {
			log.Fatal(err)
		}

		if err := json.Unmarshal(dataJson, &result[i]); err != nil {
			log.Fatal(err)
		}
	}

	return result
}
