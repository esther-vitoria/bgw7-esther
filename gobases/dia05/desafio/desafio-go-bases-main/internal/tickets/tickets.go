package tickets

import (
	"encoding/csv"
	"errors"
	"os"
	"strconv"
	"strings"
)

// struct para representar o arquivo csv
type Ticket struct {
	Id      int
	Nome    string
	Email   string
	Destino string
	Horario string
	Preco   float64
}

// Carrega os dados que estão no arquivo tickets.csv, caso não ocorra erros faz verificacões e converte alguns dados (id, preco) depois armazena tudo em um array
func LoadTickets() ([]Ticket, error) {
	file, err := os.Open("/Users/enbalbino/Documents/GitHub/bgw7-esther/gobases/dia05/desafio/desafio-go-bases-main/tickets.csv")
	if err != nil {
		return nil, err
	}

	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var tickets []Ticket
	for _, record := range records {
		id, _ := strconv.Atoi(record[0])
		preco, _ := strconv.ParseFloat(record[5], 64)
		tickets = append(tickets, Ticket{
			Id:      id,
			Nome:    record[1],
			Email:   record[2],
			Destino: record[3],
			Horario: record[4],
			Preco:   preco,
		})

	}

	return tickets, nil

}

// Recebe o destino, verifica no array quais dados contem o destino e devolve o total e um erro
func GetTotalTickets(destination string) (int, error) {
	tickets, err := LoadTickets()
	if err != nil {
		return 0, err
	}

	count := 0

	for _, ticket := range tickets {
		if strings.EqualFold(ticket.Destino, destination) {
			count++
		}
	}
	return count, nil
}

// Recebe um periodo (madrugada, manhã, tarde, noite), faz a leitura do arquivo CSV, analisa o campo horário de cada ticket e conta quantos se encaixam naquele período
func GetCountByPeriod(period string) (int, error) {
	tickets, err := LoadTickets()
	if err != nil {
		return 0, err
	}

	count := 0
	var start, end int

	switch strings.ToLower(period) {
	case "madrugada":
		start, end = 0, 6
	case "manhã":
		start, end = 1, 12
	case "tarde":
		start, end = 13, 19
	case "noite":
		start, end = 20, 23
	default:
		return 0, errors.New("periodo invalido")
	}

	// Para cada ticket, verifica se o horário está no período
	for _, ticket := range tickets {
		// horário está assim: "08:00" por exemplo
		parts := strings.Split(ticket.Horario, ":")
		if len(parts) < 1 {
			continue // se não der para dividir, pula esse ticket
		}
		hour, _ := strconv.Atoi(parts[0]) // pega a primeira parte (a hora)
		if hour >= start && hour <= end {
			count++
		}
	}
	return count, nil
}

// Recebe um destino e total de itens e devolve a porcentagem de tickets com destino inserido
func AverageDestination(destination string, total int) (float64, error) {
	if total == 0 {
		return 0, errors.New("total não pode ser 0")
	}

	tickets, err := LoadTickets()
	if err != nil {
		return 0, err
	}

	destCount := 0
	for _, ticket := range tickets {
		if strings.EqualFold(ticket.Destino, destination) {
			destCount++
		}
	}
	percent := (float64(destCount) / float64(total)) * 100
	return percent, nil
}
