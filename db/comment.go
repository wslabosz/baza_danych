package db

import (
	"fmt"
	"time"

	fd "github.com/fauna/faunadb-go/v4/faunadb"
)

const commentColName = "Comments"

func (fclient *FaunaClient) CreateComment(comment Comment, userRef string, postRef string) error {
	timeNow := time.Now().String()
	comment.CreatedAt = timeNow
	comment.UpdatedAt = timeNow
	uRef, err := fclient.fc.Query(fd.Ref(fd.Collection("Users"), userRef))
	if err != nil {
		return fmt.Errorf("retrieving ref for %s failed: %v", userRef, err)
	}
	var userReference fd.RefV
	err = uRef.Get(&userReference)
	if err != nil {
		return fmt.Errorf("unpacking ref for %s failed: %v", userRef, err)
	}
	comment.UserRef = userReference

	pRef, err := fclient.fc.Query(fd.Ref(fd.Collection("Posts"), postRef))
	if err != nil {
		return fmt.Errorf("retrieving ref for %s failed: %v", postRef, err)
	}
	var postReference fd.RefV
	err = pRef.Get(&postReference)
	if err != nil {
		return fmt.Errorf("unpacking ref for %s failed: %v", postRef, err)
	}
	comment.PostRef = postReference
	return fclient.addRecord(commentColName, comment)
}

func (fclient *FaunaClient) UpdateComment(commentRef int, comment Comment) error {
	return fclient.updateRecord(commentColName, commentRef, comment)
}

func (fclient *FaunaClient) DeleteComment(commentRef int) error {
	return fclient.deleteRecord(commentColName, commentRef)
}

func (fclient *FaunaClient) RetrieveComment(commentRef int) (*Comment, error) {
	record, err := fclient.retrieveRecord(commentColName, commentRef)
	if err != nil {
		return nil, err
	}
	var comment Comment
	err = record.At(fd.ObjKey("data")).Get(&comment)
	if err != nil {
		return nil, err
	}
	return &comment, nil
}

func (fclient *FaunaClient) RetrieveCommentCollection() ([]Comment, error) {
	collection, err := fclient.retrieveCollection(commentColName)
	if err != nil {
		return nil, err
	}
	var commentList []Comment
	var comment Comment

	for _, item := range collection {
		err := item.At(fd.ObjKey("data")).Get(&comment)
		if err != nil {
			return nil, err
		}
		commentList = append(commentList, comment)
	}
	if err != nil {
		return nil, err
	}
	return commentList, nil
}
