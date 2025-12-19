package main

import (
	// "context"
	"fmt"
	"os/exec"
	"strings"
	// "github.com/nlpodyssey/openai-agents-go/agents"
)

func GetCodeDiff() (string, error) {
	mainBranch, err := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()
	if err != nil {
		return "", err
	}
	currentBranch, err := exec.Command("git", "branch", "--show-current").Output()
	if err != nil {
		return "", err
	}

	diffCommand := []string{"diff"}
	currentBranchName := strings.TrimSpace(string(currentBranch))
	mainBranchName := strings.TrimSpace(string(mainBranch))
	if currentBranchName != mainBranchName {
		diffCommand = append(diffCommand, mainBranchName)
	}

	diffString, err := exec.Command("git", diffCommand...).Output()
	if err != nil {
		return "", err
	}
	return string(diffString), nil
}

func main() {
	diff, err := GetCodeDiff()
	if err != nil {
		fmt.Println("Error getting code diff:", err)
		return
	}
	fmt.Println("Code Diff:\n\n", diff)
	/* agent := agents.New("Pull Request Agent").
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

	fmt.Println(result.FinalOutput) */
}
