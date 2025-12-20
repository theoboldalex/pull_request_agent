package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/nlpodyssey/openai-agents-go/agents"
)

const MODEL = "gpt-4o"
const GIT = "git"

// runCommand is a package-level variable so it can be mocked in tests.
var runCommand = func(name string, args ...string) ([]byte, error) {
	return exec.Command(name, args...).Output()
}

// readFile is a package-level variable wrapper for os.ReadFile so it can be mocked in tests.
var readFile = os.ReadFile

func GetCodeDiff() (string, error) {
	mainBranch, err := runCommand(GIT, "rev-parse", "--abbrev-ref", "HEAD")
	if err != nil {
		return "", err
	}
	currentBranch, err := runCommand(GIT, "branch", "--show-current")
	if err != nil {
		return "", err
	}

	diffCommand := []string{"diff"}
	currentBranchName := strings.TrimSpace(string(currentBranch))
	mainBranchName := strings.TrimSpace(string(mainBranch))
	if currentBranchName != mainBranchName {
		diffCommand = append(diffCommand, mainBranchName)
	}

	diffString, err := runCommand(GIT, diffCommand...)
	if err != nil {
		return "", err
	}
	return string(diffString), nil
}

func GetAgentInstructions() (string, error) {
	instructionsBytes, err := readFile("instructions.md")
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
