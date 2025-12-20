# Pull Request Agent 🤖

A really basic pull request agent that will take the diff of the current branch against the previous commit on the main branch. It uses this diff to write a brief pull request title and description.

## Prerequisites

Ensure you have a valid `OPENAI_API_KEY` variable set in your environment.

## Building and Running

TBC

## Running tests

Run the unit tests with:

![CI](https://github.com/theoboldalex/pull_request_agent/actions/workflows/go.yml/badge.svg)

# Pull Request Agent 🤖

A really basic pull request agent that takes the diff of the current branch against the main branch and uses that diff to generate a short pull request title and description.

## Prerequisites

Ensure you have a valid `OPENAI_API_KEY` environment variable set.

## Building and running

TBC

## Running tests

Run the unit tests with:

```sh
go test ./...
```

The tests mock external git and file IO, so they are deterministic and do not require a git repository or an `instructions.md` file to be present.
