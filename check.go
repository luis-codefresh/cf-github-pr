package pullrequest

import (
	"fmt"
)

func Check(client *GithubClient) (string, error) {
	pulls, err := ListOpenPullRequests(client)

	if err != nil {
		return "hello", fmt.Errorf("failed to get list of open pull requests: %s", err)
	}

	return pulls, nil

}
