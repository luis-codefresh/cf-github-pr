package pullrequest

import (
    "context"
    "crypto/tls"
    "errors"
    "fmt"
    "strings"

    "golang.org/x/oauth2"
    "github.com/shurcooL/githubv4"
)

type GithubClient struct {
    V4         *github.Client
    Repository string
    Owner      string
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

func NewGithubClient (r *Gitrepo) (*GithubClient, error) {
    owner, repository, error := parseRepository(r.Repository)
    if err != nil {
		    return nil, err
	  }

    var ctx context.Context
    ctx = context.TODO()

    client := oauth2.NewClient(ctx, oauth2.StaticTokenSource(
          &oauth2.Token{AccessToken: r.AccessToken}
      ))

    var v4 *githubv4.Client
    v4 = githubv4.NewClient(client)

    return &GithubClient{
        V4:         v4,
        Owner:      owner,
        Repository: repository,
    } , nil

}

func ListOpenPullRequests (c *GithubClient) ([]*PullRequest, error) {
    var query struct{
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

    for {
      if err := c.V4.Query(context.TODO(), &query, vars); err != nil {
          return nil, err
      }

      for _, p := range query.Repository.PullRequests.Edges {
          for _, c := range p.Node {
              response = append(response, &PullRequest{
                  ID: p.Node.Id
                  BaseRefName: p.Node.BaseRefName
                  HeadRefName: p.Node.HeadRefName
              })
              log.Printf("Branch Name: %s", p.Node.HeadRefName)
          }
      }
    }
}


func parseRepository (s string) (string, string, error) {
    parts := string.Split(s, "/")

    if len(parts) != 2 {
        return "", "", errors.New("malformed repository")
    }

    return parts[0], parts[1], nil

}
