package db

import (
	"fmt"
	"time"

	fd "github.com/fauna/faunadb-go/v4/faunadb"
)

const followColName = "Follows"

func (fclient *FaunaClient) CreateFollow(followingRef, followerRef string) error {
	follow := Follow{}
	timeNow := time.Now().String()
	follow.CreatedAt = timeNow
	follow.UpdatedAt = timeNow
	followerFaunaRef, err := fclient.fc.Query(fd.Ref(fd.Collection("Users"), followingRef))
	if err != nil {
		return fmt.Errorf("retrieving ref for %s failed: %v", followingRef, err)
	}
	var followerReference fd.RefV
	err = followerFaunaRef.Get(&followerReference)
	if err != nil {
		return fmt.Errorf("unpacking ref for %s failed: %v", followingRef, err)
	}
	follow.FollowerRef = followerReference

	followingFaunaRef, err := fclient.fc.Query(fd.Ref(fd.Collection("Users"), followerRef))
	if err != nil {
		return fmt.Errorf("retrieving ref for %s failed: %v", followerRef, err)
	}
	var followingReference fd.RefV
	err = followingFaunaRef.Get(&followingReference)
	if err != nil {
		return fmt.Errorf("unpacking ref for %s failed: %v", followerRef, err)
	}
	follow.FollowingRef = followingReference
	return fclient.addRecord(followColName, follow)
}

func (fclient *FaunaClient) UpdateFollow(followRef int, follow Follow) error {
	return fclient.updateRecord(followColName, followRef, follow)
}

func (fclient *FaunaClient) DeleteFollow(followRef int) error {
	return fclient.deleteRecord(followColName, followRef)
}

func (fclient *FaunaClient) RetrieveFollow(followRef int) (*Follow, error) {
	record, err := fclient.retrieveRecord(followColName, followRef)
	if err != nil {
		return nil, err
	}
	var follow Follow
	err = record.At(fd.ObjKey("data")).Get(&follow)
	if err != nil {
		return nil, err
	}
	return &follow, nil
}

func (fclient *FaunaClient) RetrieveFollowCollection() ([]Follow, error) {
	collection, err := fclient.retrieveCollection(followColName)
	if err != nil {
		return nil, err
	}
	var followList []Follow
	var follow Follow

	for _, item := range collection {
		err := item.At(fd.ObjKey("data")).Get(&follow)
		if err != nil {
			return nil, err
		}
		followList = append(followList, follow)
	}
	if err != nil {
		return nil, err
	}
	return followList, nil
}
