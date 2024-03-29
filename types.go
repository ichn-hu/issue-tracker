package main

import (
	"time"

	"github.com/shurcooL/githubv4"
)

type Repository struct {
	Name  githubv4.String
	Owner struct {
		Login githubv4.String
	}
}

type PullRequest struct {
	PullRequestWithoutTimelineItems
	TimelineItems struct {
		Edges []struct {
			Node struct {
				Typename             string `graphql:"__typename"`
				CrossReferencedEvent struct {
					Source struct {
						PullRequest PullRequestWithoutTimelineItems `graphql:"... on PullRequest"`
					}
				} `graphql:"... on CrossReferencedEvent"`
				IssueComment struct {
					Author struct {
						Login githubv4.String
					}
					Body githubv4.String
				} `graphql:"... on IssueComment"`
			}
		}
	} `graphql:"timelineItems(first: 15, itemTypes: [CROSS_REFERENCED_EVENT, ISSUE_COMMENT] )"`
}

type CloserPRInfo struct {
	Repository  string
	Number      int
	Title       string
	Url         string
	State       string
	MergeTarget string
	MergedAt    time.Time
	Commit      string
	Refs        []string
}

func (pr *PullRequest) getCloserInfo() *CloserPRInfo {
	by := &CloserPRInfo{}
	by.Title = string(pr.Title)
	by.Number = int(pr.Number)
	by.Repository = string(pr.Repository.Name)
	by.Url = string(pr.Url)
	by.State = string(pr.State)
	by.MergeTarget = string(pr.BaseRefName)
	by.MergedAt = pr.MergedAt.Time
	by.Commit = string(pr.MergeCommit.OID)
	if by.Commit != "" {
	}
	return by
}

type PullRequestWithoutTimelineItems struct {
	ID          githubv4.ID
	State       githubv4.PullRequestState
	Merged      githubv4.Boolean
	MergedAt    githubv4.DateTime
	MergeCommit struct {
		OID           githubv4.GitObjectID
		CommittedDate githubv4.DateTime
	}
	Author struct {
		Login githubv4.String
	}
	CreatedAt githubv4.DateTime
	UpdatedAt githubv4.DateTime
	Title     githubv4.String
	Url       githubv4.String
	Number    githubv4.Int
	Labels    struct {
		Nodes []struct {
			Name githubv4.String
		}
	} `graphql:"labels(last: 15)"`
	Repository  Repository
	BaseRefName githubv4.String
	HeadRefName githubv4.String
}

type IssueNode struct {
	Title  githubv4.String
	State  githubv4.IssueState
	ID     githubv4.ID
	Number githubv4.Int
	Url    githubv4.String
	Author struct {
		Login githubv4.String
	}
	Body       githubv4.String
	ClosedAt   githubv4.DateTime
	CreatedAt  githubv4.DateTime
	UpdatedAt  githubv4.DateTime
	Repository struct {
		Name  githubv4.String
		Owner struct {
			Login githubv4.String
		}
	}
	Labels struct {
		Nodes []struct {
			Name githubv4.String
		}
	} `graphql:"labels(last: 15)"`
	Assignees struct {
		Nodes []struct {
			Login     githubv4.String
			CreatedAt githubv4.DateTime
		}
	} `graphql:"assignees(last: 5)"`
	TimelineItems struct {
		Edges []struct {
			Node struct {
				Typename             string `graphql:"__typename"`
				CrossReferencedEvent struct {
					WillCloseTarget githubv4.Boolean
					Source          struct {
						PullRequest PullRequest `graphql:"... on PullRequest"`
					}
				} `graphql:"... on CrossReferencedEvent"`
				ClosedEvent struct {
					Actor struct {
						Login githubv4.String
					}
					Closer struct {
						PullRequest PullRequest `graphql:"... on PullRequest"`
					}
				} `graphql:"... on ClosedEvent"`
			}
		}
	} `graphql:"timelineItems(first: 20, itemTypes: [CROSS_REFERENCED_EVENT, CLOSED_EVENT] )"`
}

type ClosedIssueInfo struct {
	Number             int
	Title              string
	Url                string
	ClosedAt           time.Time
	Severity           string
	AffectedVersions   []string
	ClosedByPR         *CloserPRInfo
	CloserCherryPicked []*CloserPRInfo
}
