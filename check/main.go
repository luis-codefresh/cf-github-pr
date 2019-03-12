package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/caarlos0/env"
	"github.com/lfurrea/cf-github-pr"
)

func main() {
	repo := pullrequest.Gitrepo{}
	err := env.Parse(&repo)

	if err != nil {
		log.Fatalf("invalid source repository configuration: %s", err)
	}

	github, err := pullrequest.NewGithubClient(&repo)

	if err != nil {
		log.Fatalf("failed to create github client: %s", err)
	}

	response, err := pullrequest.Check(github)

	if err != nil {
		log.Fatalf("failed to check repository for pullrequests: %s", err)
	}

	if err := json.NewEncoder(os.Stdout).Encode(response); err != nil {
		log.Fatalf("failed to marshal response: %s", err)
	}
}
