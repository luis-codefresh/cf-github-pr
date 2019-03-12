package main

import (
	"log"

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
	log.Printf("Branches: %", response)

	if err != nil {
		log.Fatalf("failed to check repository for pullrequests: %s", err)
	}

	//TODO: decide what the response is going to look like and return it
}
