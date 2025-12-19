package main

import (
	"context"
	"fmt"

	"github.com/nlpodyssey/openai-agents-go/agents"
)

func GetCodeDiff() (string, error) {

}

func main() {
	agent := agents.New("Pull Request Agent").
		WithInstructions("You are a helpful assistant that creates pull requests.").
		WithModel("gpt-4o")

	diff, err := GetCodeDiff()
	if err != nil {
		fmt.Println("Error getting code diff:", err)
	}
	result, err := agents.Run(context.Background(), agent, diff)
	if err != nil {
		panic(err)
	}

	fmt.Println(result.FinalOutput)
}
