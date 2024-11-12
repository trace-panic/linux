package jsonfs

import (
	"encoding/json"
)

// MarshalJSON marshals the provided data into a JSON byte slice.
// It returns the marshaled byte slice or an error if marshaling fails.
func MarshalJSON(data interface{}) ([]byte, error) {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}

// UnmarshalJSON unmarshals the provided JSON byte slice into the target data structure.
// It returns an error if unmarshaling fails.
func UnmarshalJSON(data []byte, target interface{}) error {
	// Unmarshal the JSON byte slice into the target variable
	if err := json.Unmarshal(data, target); err != nil {
		return err
	}
	return nil
}
