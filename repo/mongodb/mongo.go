package mongodb

import (
	"context"

	"Hausuebung-I/repo"

	"github.com/google/uuid"
	"github.com/nilspolek/goLog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongorepo struct {
	IssueCollection mongo.Collection
}

func New() repo.Repo {
	var mr mongorepo

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		goLog.Error(err.Error())
	}

	mr.IssueCollection = *client.Database("issuesdb").Collection("issues")

	return &mr
}

func (mr *mongorepo) CreateIssue(issue repo.Issue) error {
	_, err := mr.IssueCollection.InsertOne(context.TODO(), issue)
	return err
}

func (mr *mongorepo) GetIssues() ([]repo.Issue, error) {
	var (
		issues []repo.Issue
		ctx    = context.Background()
	)
	cursor, err := mr.IssueCollection.Find(ctx, bson.D{})
	defer cursor.Close(ctx)
	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var result repo.Issue
		if err := cursor.Decode(&result); err != nil {
			return nil, err
		}
		issues = append(issues, result)
	}
	return issues, nil
}

func (mr *mongorepo) GetIssue(id uuid.UUID) (issue repo.Issue, err error) {
	var (
		ctx = context.Background()
	)

	err = mr.IssueCollection.FindOne(ctx, bson.M{"id": id}).Decode(&issue)
	return
}

func (mr *mongorepo) PutIssue(id uuid.UUID, issue repo.Issue) (err error) {
	ctx := context.Background()

	_, err = mr.IssueCollection.UpdateOne(ctx, bson.M{"id": id}, bson.M{"$set": issue})
	return
}

func (mr *mongorepo) PatchIssue(id uuid.UUID, issue repo.Issue) (err error) {
	ctx := context.Background()
	update := bson.M{"$set": bson.M{}}

	if issue.Title != "" {
		update["$set"].(bson.M)["title"] = issue.Title
	}
	if issue.Owner != "" {
		update["$set"].(bson.M)["status"] = issue.Owner
	}
	if issue.Id.String() != "" {
		update["$set"].(bson.M)["assignee"] = issue.Id
	}

	_, err = mr.IssueCollection.UpdateOne(ctx, bson.M{"id": id}, update)
	return
}

func (mr *mongorepo) DeleteIssue(id uuid.UUID) (err error) {
	var (
		ctx = context.Background()
	)

	_, err = mr.IssueCollection.DeleteOne(ctx, bson.M{"id": id})
	return
}
