package extras

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	"github.com/hassan-algo/pc-gamification/types"
	"github.com/tealeg/xlsx"
)

func UpdateField(obj interface{}, fieldName string, newValue interface{}) {
	objValue := reflect.ValueOf(obj).Elem()
	fieldValue := objValue.FieldByName(fieldName)

	if fieldValue.IsValid() {
		if fieldValue.CanSet() {
			newValueValue := reflect.ValueOf(newValue)
			fieldValue.Set(newValueValue)
		}
	}
}

func Contains(slice []string, value string) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}

func ContainsInt(slice []int, value int) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}

func Max(arr []float64) float64 {
	if len(arr) == 0 {
		// Return an appropriate default value if the array is empty
		return 0
	}

	maxValue := arr[0]
	for _, value := range arr {
		if value > maxValue {
			maxValue = value
		}
	}
	return maxValue
}

func SheetToJson(sheet *xlsx.Sheet) []map[string]interface{} {

	var jsonData []map[string]interface{}

	colToKey := make(map[int]string)

	// Iterate through each row in the sheet
	for rowIndex, row := range sheet.Rows {
		if rowIndex == 0 {
			// The first row contains the keys for the JSON objects
			for cellIndex, cell := range row.Cells {
				colToKey[cellIndex] = cell.String()
			}
			continue
		}

		// Create a map to store the data for the current row
		rowData := make(map[string]interface{})

		// Populate the map using the keys from the first row
		for cellIndex, cell := range row.Cells {
			key := colToKey[cellIndex]
			if key == "Min Value" || key == "PCC Reward" || key == "Max Value" {
				rowData[key], _ = cell.Float()
			} else if key == "Started" {
				rowData[key], _ = cell.Int()
			} else {
				rowData[key] = cell.String()
			}
		}

		// Append the row data to the slice
		jsonData = append(jsonData, rowData)
	}
	// printJSON(jsonData)
	return jsonData
}

func GetField(nft types.PlayerStats, field string) interface{} {
	r := reflect.ValueOf(nft)
	f := reflect.Indirect(r).FieldByName(field)
	return f.Interface()
}

func printJSON(data interface{}) {
	jsonBytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsonBytes))
}
