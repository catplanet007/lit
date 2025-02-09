package encoding

import (
	"encoding/json"
	"os"
)

func MustJson(v any) []byte {
	bs, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return bs
}

func LoadJsonFromFile(filename string, result any) error {
	bs, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(bs, result); err != nil {
		return err
	}
	return nil
}

func DumpJsonToFile(filename string, v any) error {
	bs, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, bs, 0644)
}
