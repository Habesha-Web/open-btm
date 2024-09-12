package utils

import (
	"encoding/json"
	"fmt"
	"math/rand"
)

const (
	charsetLen = 26
	charset    = "abcdefghijklmnopqrstuvwxyz"
)

func GenerateRandomString(length int) (string, error) {
	result := make([]byte, length)

	for i := 0; i < length; i++ {
		index := rand.Intn(charsetLen)
		result[i] = charset[index]
	}

	return string(result), nil
}

func PopulateStructFromMap(dest interface{}, data map[string]interface{}) (interface{}, error) {

	// Marshal map to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON to struct
	err = json.Unmarshal(jsonData, &dest)
	if err != nil {
		return nil, err
	}

	fmt.Println(dest)
	return dest, nil
}
