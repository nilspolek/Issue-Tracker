package repo

import "github.com/google/uuid"

type Repo interface {
	CreateIssue(Issue) error
	GetIssues() ([]Issue, error)
	GetIssue(uuid.UUID) (Issue, error)
	PutIssue(uuid.UUID, Issue) error
	PatchIssue(uuid.UUID, Issue) error
	DeleteIssue(uuid.UUID) error
}

type Issue struct {
	Id    uuid.UUID `bson:"id" json:"id"`
	Title string    `bson:"title" json:"title"`
	Owner string    `bson:"owner" json:"owner"`
}
