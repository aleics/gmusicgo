package gmusicjson

import (
	"errors"
	"io/ioutil"
	"encoding/json"
)

type Gmusicjson struct{
}

func Deserialize(t interface{}) ([]byte, error){
	return json.Marshal(t)
	return json.Marshal(nil)
}

func Import(path string, t interface{}) (error) {
	if path != "" {
		dat, err := ioutil.ReadFile(path)
		if err != nil{
			return errors.New("Error importing JSON: reading file")
		}
		err = json.Unmarshal(dat, &t)
		if err != nil {
			return errors.New("Error importing JSON: Unmarshaling file")
		}
		return nil
	}
	return errors.New("Error importing JSON: path parameter empty")
}

func Export(input interface{}, path string) (string, error) {
	output, err := Deserialize(input)
	if err != nil {
		return string(output), errors.New("Error exporting JSON: deserializing struct")
	}

	if path != "" {
		err = ioutil.WriteFile(path, output, 0644)
		if err != nil {
		return string(output), errors.New("Error exporting JSON: writing file")
		}
	}

	if string(output) == "" {
	return string(output), errors.New("Error exporting JSON: return empty")
	} else {
		return string(output), nil
	}
}
