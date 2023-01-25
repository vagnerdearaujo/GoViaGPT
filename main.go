package main

import (
	"Go_via_GPT/config"
	"Go_via_GPT/db"
	"fmt"
)

func main() {
	//Carrega a configuração
	dbConfig, err := config.LoadDBConfig()
	if err != nil {
		fmt.Println("Não foi possível abrir o arquivo de configuração:", dbConfig.Database, "\n", err)
		return
	}

	db, err := db.Connect(dbConfig)
	if err != nil {
		fmt.Println("Error opening database:", dbConfig.Database, "\n", err)
		return
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("Error pinging database:", err)
		return
	}
	fmt.Println("Successfully connected to: ", dbConfig.Database)
}
