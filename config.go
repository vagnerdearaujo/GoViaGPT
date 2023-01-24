package config

import (
    "encoding/json"
    "os"
)

// DBConfig contains the information necessary to connect to the database
type DBConfig struct {
    Type     string `json:"type"`
    Host     string `json:"host"`
    Port     string `json:"port"`
    User     string `json:"user"`
    Password string `json:"password"`
    Database string `json:"database"`
}

// LoadDBConfig loads the database configuration from the config.json file
func LoadDBConfig() (DBConfig, error) {
    var dbConfig DBConfig
    file, err := os.Open("config.json")
    if err != nil {
        return dbConfig, err
    }
    defer file.Close()
    decoder := json.NewDecoder(file)
    err = decoder.Decode(&dbConfig)
    if err != nil {
        return dbConfig, err
    }
    return dbConfig, nil
}
