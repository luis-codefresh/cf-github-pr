package pullrequest

import (
	"context"
	"errors"
	"log"
	"strings"

	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

type Gitrepo struct {
	Repository    string `env:"GITHUB_REPOSITORY"`
	AccessToken   string `env:"GITHUB_TOKEN"`
	WebhookBranch string `env:"CF_BRANCH"`
}

type GithubClient struct {
	V4            *githubv4.Client
	Repository    string
	Owner         string
	WebhookBranch string
}

type PullRequest struct {
	ID          string
	Number      int
	Title       string
	URL         string
	BaseRefName string
	HeadRefName string
	Repository  struct {
		URL string
	}
	IsCrossRepository bool
}

func NewGithubClient(r *Gitrepo) (*GithubClient, error) {
	owner, repository, err := parseRepository(r.Repository)
	if err != nil {
		return nil, err
	}

	var ctx context.Context
	ctx = context.TODO()

	client := oauth2.NewClient(ctx, oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: r.AccessToken},
	))

	var v4 *githubv4.Client
	v4 = githubv4.NewClient(client)

	return &GithubClient{
		V4:            v4,
		Owner:         owner,
		Repository:    repository,
		WebhookBranch: r.WebhookBranch,
	}, nil

}

func ListOpenPullRequests(c *GithubClient) ([]*PullRequest, error) {
	//func ListOpenPullRequests(c *GithubClient) (string, error) {
	var query struct {
		Repository struct {
			PullRequests struct {
				Edges []struct {
					Node struct {
						PullRequest
					}
				}
			} `graphql:"pullRequests(first:$prFirst,states:$prStates)"`
		} `graphql:"repository(owner:$repositoryOwner,name:$repositoryName)"`
	}

	vars := map[string]interface{}{
		"repositoryOwner": githubv4.String(c.Owner),
		"repositoryName":  githubv4.String(c.Repository),
		"prFirst":         githubv4.Int(100),
		"prStates":        []githubv4.PullRequestState{githubv4.PullRequestStateOpen},
		//		    "prCursor":        (*githubv4.String)(nil),
	}

	var response []*PullRequest

	if err := c.V4.Query(context.TODO(), &query, vars); err != nil {
		return nil, err
	}

	for _, p := range query.Repository.PullRequests.Edges {
		response = append(response, &PullRequest{
			ID:          p.Node.ID,
			BaseRefName: p.Node.BaseRefName,
			HeadRefName: p.Node.HeadRefName,
		})
		branch := p.Node.HeadRefName
		id := p.Node.ID
		log.Printf("Found branch %s with an open PR->id: %s", branch, id)
	}

	return response, nil
}

func parseRepository(s string) (string, string, error) {
	parts := strings.Split(s, "/")

	if len(parts) != 2 {
		return "", "", errors.New("malformed repository")
	}

	return parts[0], parts[1], nil

}
