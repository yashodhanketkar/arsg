package api

import (
	"errors"
	"fmt"

	"github.com/yashodhanketkar/arsg/src/db"
	"github.com/yashodhanketkar/arsg/src/util"
)

type calcStruct struct {
	Title   string `json:"title"`
	Comment string `json:"comments"`
	Art     string `json:"art"`
	Cast    string `json:"cast"`
	Plot    string `json:"plot"`
	Bias    string `json:"bias"`
	Rating  string `json:"rating"`
}

var (
	// INFO: Will replace util.DefaultParams in future updates
	apiParams         = []string{"Art", "Cast", "Plot", "Bias"}
	validContentTypes = map[string]bool{"anime": true, "manga": true, "lightnovel": true}
)

// convert data to a list for code reduction
func (bc calcStruct) metrics() []string { return []string{bc.Art, bc.Cast, bc.Plot, bc.Bias} }

// check if content type is valid
func validateContentType(content_type string) (string, error) {
	if !validContentTypes[content_type] {
		return "", errors.New("Invalid content type: " + content_type)
	}

	return content_type, nil
}

// handle input validation as per rquirements
func (bc calcStruct) parseAndValidate() ([4]float32, error) {
	var vals [4]float32

	if bc.Title == "" {
		return vals, errors.New("Missing title")
	}

	for i, str := range bc.metrics() {
		if str == "" {
			return vals, errors.New("Missing parameter - " + apiParams[i])
		}
		vals[i] = util.FloatParser(str)
	}

	return vals, nil
}

// get calculated rating
func (bc calcStruct) getCalc(v [4]float32) (float32, error) {
	return util.Calculator(&util.ConfigType{}, v[0], v[1], v[2], v[3])
}

// fomrat output for response
// will be converted to `{ message: "", data: {}}` format in future
func (bc calcStruct) formatOutput(rating float32) string {
	output := fmt.Sprintf(`{"Title": "%s",`, bc.Title)

	if bc.Comment != "" {
		output += fmt.Sprintf(`"Comments": "%s",`, bc.Comment)
	}

	for i, v := range bc.metrics() {
		output += fmt.Sprintf(`"%s": %s,`, apiParams[i], v)
	}

	output += fmt.Sprintf(`"Rating": %.1f}`, rating)
	return output
}

// handle conversion for database entry
// might be removed in future with UI package changes
func convertRating(rating map[string]interface{}) db.Rating {
	var dbRating db.Rating

	dbRating.Name = extractStringValue(rating, "Title")
	dbRating.Comments = extractStringValue(rating, "Comments")
	dbRating.Rating = extractStringValue(rating, "Rating")
	dbRating.Art = extractFloatValue(rating, "Art")
	dbRating.Support = extractFloatValue(rating, "Cast")
	dbRating.Plot = extractFloatValue(rating, "Plot")
	dbRating.Bias = extractFloatValue(rating, "Bias")

	return dbRating
}
