package main

import (
	"context"
	"fmt"
	"github.com/nlpodyssey/openai-agents-go/agents"
	"os"
	"os/exec"
	"strings"
)

const MODEL = "gpt-4o"
const GIT = "git"

func GetCodeDiff() (string, error) {
	mainBranch, err := exec.Command(GIT, "rev-parse", "--abbrev-ref", "HEAD").Output()
	if err != nil {
		return "", err
	}
	currentBranch, err := exec.Command(GIT, "branch", "--show-current").Output()
	if err != nil {
		return "", err
	}

	diffCommand := []string{"diff"}
	currentBranchName := strings.TrimSpace(string(currentBranch))
	mainBranchName := strings.TrimSpace(string(mainBranch))
	if currentBranchName != mainBranchName {
		diffCommand = append(diffCommand, mainBranchName)
	}

	diffString, err := exec.Command(GIT, diffCommand...).Output()
	if err != nil {
		return "", err
	}
	return string(diffString), nil
}

func GetAgentInstructions() (string, error) {
	instructionsBytes, err := os.ReadFile("instructions.md")
	if err != nil {
		return "", err
	}
	return string(instructionsBytes), nil
}

func main() {
	diff, err := GetCodeDiff()
	if err != nil {
		fmt.Println("Error getting code diff:", err)
		return
	}

	instructions, err := GetAgentInstructions()
	if err != nil {
		fmt.Println("Error getting agent instructions:", err)
		return
	}
	agent := agents.New("Pull Request Agent").
		WithInstructions(instructions).
		WithModel(MODEL)

	result, err := agents.Run(context.Background(), agent, diff)
	if err != nil {
		panic(err)
	}

	fmt.Println(result.FinalOutput)
}
