package pullrequest

import (
    "fmt"

)

func Check(client GithubClient) {
    pulls, err := client.ListOpenPullRequests()
    if err != nil {
        return nil, fmt.Errorf("failed to get last commits: %s", err)
    }
}
