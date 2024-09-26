package tools

import (
	"encoding/json"
	"os"
)

var Conf Config

func ReadConfJson(confPath string) error {
    // Read config file
    confReader, err := os.Open(confPath)
    if err != nil {
        return err
    }

    if err := json.NewDecoder(confReader).Decode(&Conf); err != nil {
        panic(err)
    }

    return nil
}
