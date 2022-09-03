package data

import (
	"fmt"
	"sort"
	"time"

	"github.com/cli/go-gh"
)

type Repository struct {
	Name string `json:"name"`
}

type HeadCommit struct {
	Message string `json:"message"`
}

type WorkflowRun struct {
	Name       string     `json:"name"`
	Status     string     `json:"status"`
	Event      string     `json:"event"`
	HeadCommit HeadCommit `json:"head_commit"`
	RunNumber  int        `json:"run_number"`
	Conclusion string     `json:"conclusion"`
	HeadSHA    string     `json:"head_sha"`
	Repository Repository `json:"repository"`
	CreatedAt  time.Time  `json:"created_at"`
}

type response struct {
	WorkflowRuns []WorkflowRun `json:"workflow_runs"`
	TotalCount   int32         `json:"total_count"`
}

func GetWorkflowRuns(repoPaths []string) []WorkflowRun {
	client, err := gh.RESTClient(nil)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	wFlows := make([]WorkflowRun, 0)

	for _, repoPath := range repoPaths {
		var resp response
		url := fmt.Sprintf("repos/%s/actions/runs?per_page=5", repoPath)
		err = client.Get(url, &resp)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		wFlows = append(wFlows, resp.WorkflowRuns...)
	}

	sort.Slice(wFlows, func(i, j int) bool {
		return wFlows[i].CreatedAt.After(wFlows[j].CreatedAt)
	})

	return wFlows
}
