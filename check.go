package pullrequest

import (
	"fmt"
)

func Check(client *GithubClient) ([]*PullRequest, error) {
	pulls, err := ListOpenPullRequests(client)

	if err != nil {
		return nil, fmt.Errorf("failed to get list of open pull requests: %s", err)
	}

	var response []*PullRequest

	for _, p := range pulls {
		if p.HeadRefName == client.WebhookBranch {
			response = append(response, p)
		}
	}
	return response, nil
}
