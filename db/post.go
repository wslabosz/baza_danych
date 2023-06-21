package db

import (
	"fmt"
	"time"

	fd "github.com/fauna/faunadb-go/v4/faunadb"
)

const postColName = "Posts"

func (fclient *FaunaClient) CreatePost(post Post, userRef string) error {
	timeNow := time.Now().String()
	post.CreatedAt = timeNow
	post.UpdatedAt = timeNow
	ref, err := fclient.fc.Query(fd.Ref(fd.Collection("Users"), userRef))
	if err != nil {
		return fmt.Errorf("retrieving ref for %s failed: %v", userRef, err)
	}
	var reference fd.RefV
	err = ref.Get(&reference)
	if err != nil {
		return fmt.Errorf("unpacking ref for %s failed: %v", userRef, err)
	}
	post.UserRef = reference
	return fclient.addRecord(postColName, post)
}

func (fclient *FaunaClient) UpdatePost(postRef int, post Post) error {
	return fclient.updateRecord(postColName, postRef, post)
}

func (fclient *FaunaClient) DeletePost(postRef int) error {
	return fclient.deleteRecord(postColName, postRef)
}

func (fclient *FaunaClient) RetrievePost(postRef int) (*Post, error) {
	record, err := fclient.retrieveRecord(postColName, postRef)
	if err != nil {
		return nil, err
	}
	var post Post
	err = record.At(fd.ObjKey("data")).Get(&post)
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (fclient *FaunaClient) RetrievePostCollection() ([]Post, error) {
	collection, err := fclient.retrieveCollection(postColName)
	if err != nil {
		return nil, err
	}
	var postList []Post
	var post Post

	for _, item := range collection {
		err := item.At(fd.ObjKey("data")).Get(&post)
		if err != nil {
			return nil, err
		}
		postList = append(postList, post)
	}
	if err != nil {
		return nil, err
	}
	return postList, nil
}
