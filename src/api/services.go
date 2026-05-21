package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"io"

	"github.com/yashodhanketkar/arsg/src/db"
)

func contentList(content_type string, DB *sql.DB) ([]byte, error) {
	var output_list []db.Rating

	if !validContentTypes[content_type] {
		return nil, errors.New("Invalid content type: " + content_type)
	}

	for _, rating := range db.ListRatings(DB, content_type) {
		output_list = append(output_list, rating)
	}

	jsonData, err := json.Marshal(output_list)
	if err != nil {
		return nil, errors.New("Invalid input format: " + err.Error())
	}

	return jsonData, nil
}

func calculateResult(body io.ReadCloser) ([]byte, error) {
	var data calcStruct

	if err := json.NewDecoder(body).Decode(&data); err != nil {
		return nil, errors.New("Invalid input format: " + err.Error())
	}

	vals, err := data.parseAndValidate()
	if err != nil {
		return nil, errors.New("Invalid input values: " + err.Error())
	}

	res, err := data.getCalc(vals)
	if err != nil {
		return nil, errors.New("Failed to process: " + err.Error())
	}

	output := data.formatOutput(res)

	return []byte(output), nil
}

func addRating(content_type string, body io.ReadCloser, DB *sql.DB) (db.Rating, error) {
	var rating db.Rating
	if !validContentTypes[content_type] {
		return rating, errors.New("Invalid content type: " + content_type)
	}

	jsonBody, err := calculateResult(body)
	if err != nil {
		return rating, err
	}

	if len(jsonBody) == 0 {
		return rating, errors.New("Failed to generate output")
	}

	var rawRatings map[string]interface{}
	if err := json.Unmarshal(jsonBody, &rawRatings); err != nil {
		return rating, errors.New("Invalid input format: " + err.Error())
	}

	dbRating := convertRating(rawRatings)
	db.AddRatings(DB, dbRating, content_type)

	return dbRating, nil
}
