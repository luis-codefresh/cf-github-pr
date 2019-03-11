package main

import (
    "os"
    "log"

    "github.com/lfurrea/cf-github-pr"
    "github.com/caarlos0/env"
)

type Gitrepo struct {
    Repository    string  `env:"GITHUB_REPOSITORY"`
    AccessToken   string  `env:"GITHUB_TOKEN"`
}

func main() {
    repo := Gitrepo{}
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

    //TODO: decide what the response is going to look like and return it
}
