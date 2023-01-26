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
		fmt.Println("Erro ao tentar abrir o banco de dados:", dbConfig.Database, "\n", err)
		return
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("Erro ao tentar conectar no banco de dados:", err)
		return
	}
	fmt.Println("Conectado com successo ao: ", dbConfig.Database)
}
