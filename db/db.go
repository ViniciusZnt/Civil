package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	DBHost     = "192.168.0.10"
	DBPort     = 3306
	DBUser     = "root"
	DBPassword = "123Mudar"
	DBName     = "banco_temperatura"
)

// Estrutura para armazenar os dados do banco de dados
type Dados struct {
	ID                    int    `json:"id"`
	TemperaturaAmbiente   string `json:"temperatura_ambiente"`
	TemperaturaRecipiente string `json:"temperatura_recipiente"`
	Humidade              string `json:"humidade"`
	TempoAtual            string `json:"tempo_atual"`
}

func Db_handler() ([]Dados, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", DBUser, DBPassword, DBHost, DBPort, DBName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err // retornar o erro em vez de usar panic ou log.Fatalf
	}
	defer db.Close()

	query := "SELECT temperatura_ambiente, temperatura_recipiente, humidade_atual, tempo_atual,id FROM temperatura ORDER BY tempo_atual DESC LIMIT 100;"

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Variáveis para armazenar os resultados da consulta
	var dados []Dados

	//Função para ler cada linha mapeando todas as colunas com variaveis
	for rows.Next() {
		var temperaturaAmbiente, temperaturaRecipiente, humidade, hora []uint8
		var id int32
		err := rows.Scan(&temperaturaAmbiente, &temperaturaRecipiente, &humidade, &hora, &id) //A função scan mapeia as variaveis,então devem estar na ordem da querry
		if err != nil {
			return nil, err
		}

		// Criar uma nova instância de Dados e adicionar ao slice
		dado := Dados{TempoAtual: string(hora), TemperaturaAmbiente: string(temperaturaAmbiente), TemperaturaRecipiente: string(temperaturaRecipiente), Humidade: string(humidade), ID: int(id)} // Incluindo humidade
		dados = append([]Dados{dado}, dados...)                                                                                                                                                  //Para adicionar os dados novos ao inicio da slice PILHA                                                                                                                                    // <-Isso aqui é maluquice
	}

	// Verificar erros após o loop
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return dados, nil
}
