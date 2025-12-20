package main

import (
    "errors"
    "testing"
)

func TestGetCodeDiff_SameBranch_NoDiff(t *testing.T) {
    origRun := runCommand
    defer func() { runCommand = origRun }()

    runCommand = func(name string, args ...string) ([]byte, error) {
        if len(args) >= 1 && args[0] == "rev-parse" {
            return []byte("main\n"), nil
        }
        if len(args) >= 1 && args[0] == "branch" {
            return []byte("main\n"), nil
        }
        if len(args) >= 1 && args[0] == "diff" {
            return []byte(""), nil
        }
        return nil, errors.New("unexpected git command")
    }

    diff, err := GetCodeDiff()
    if err != nil {
        t.Fatalf("expected no error, got %v", err)
    }
    if diff != "" {
        t.Fatalf("expected empty diff, got %q", diff)
    }
}

func TestGetCodeDiff_DifferentBranch_WithDiff(t *testing.T) {
    origRun := runCommand
    defer func() { runCommand = origRun }()

    runCommand = func(name string, args ...string) ([]byte, error) {
        if len(args) >= 1 && args[0] == "rev-parse" {
            return []byte("main\n"), nil
        }
        if len(args) >= 1 && args[0] == "branch" {
            return []byte("feature-1\n"), nil
        }
        if len(args) >= 1 && args[0] == "diff" {
            return []byte("diff --git a/foo b/foo\n"), nil
        }
        return nil, errors.New("unexpected git command")
    }

    diff, err := GetCodeDiff()
    if err != nil {
        t.Fatalf("expected no error, got %v", err)
    }
    if diff == "" {
        t.Fatalf("expected non-empty diff")
    }
}

func TestGetCodeDiff_CommandError(t *testing.T) {
    origRun := runCommand
    defer func() { runCommand = origRun }()

    runCommand = func(name string, args ...string) ([]byte, error) {
        return nil, errors.New("git failed")
    }

    _, err := GetCodeDiff()
    if err == nil {
        t.Fatalf("expected error, got nil")
    }
}

func TestGetAgentInstructions_Success(t *testing.T) {
    origRead := readFile
    defer func() { readFile = origRead }()

    readFile = func(path string) ([]byte, error) {
        if path != "instructions.md" {
            return nil, errors.New("unexpected path")
        }
        return []byte("Do X\n"), nil
    }

    s, err := GetAgentInstructions()
    if err != nil {
        t.Fatalf("expected no error, got %v", err)
    }
    if s != "Do X\n" {
        t.Fatalf("unexpected instructions: %q", s)
    }
}

func TestGetAgentInstructions_FileError(t *testing.T) {
    origRead := readFile
    defer func() { readFile = origRead }()

    readFile = func(path string) ([]byte, error) {
        return nil, errors.New("read failed")
    }

    _, err := GetAgentInstructions()
    if err == nil {
        t.Fatalf("expected error, got nil")
    }
}
