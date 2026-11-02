package main
import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"encoding/csv"
	
)

type Question struct {
	Text 	string
	Options []string
	Answer  int
}

type GameState struct {
	Name  	string
	Points  string
	Questions []Question
}

func (g* GameState) Init()	{
	fmt.Println("Bem vindo ao Quiz Game!")
	fmt.Println("Digite seu nome:")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	
	if err != nil {
		panic("Erro ao ler o nome do jogador")
	}

	g.Name = strings.TrimSpace(name)
	
	fmt.Printf("Olá, %s! Vamos começar o quiz!\n", g.Name)

}

func (g*GameState) ProccessCSV(q Question) {
	f, err := os.Open("quizgo.csv")
	if err != nil {
		panic("Erro ao abrir o arquivo do quiz")
	}
	defer f.Close()

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		panic("Erro ao ler o arquivo do quiz")
	}
	
	for _, record := range records {
		fmt.Println(record)
	}
}

func main( ) {
	game := GameState{}
	// game.Init()
	game.ProccessCSV(Question{})
}