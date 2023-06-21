package db

import (
	"fmt"
	"time"

	fd "github.com/fauna/faunadb-go/v4/faunadb"
)

const VoteColName = "Votes"

func (fclient *FaunaClient) CreateVote(vote Vote, userRef, postRef string) error {
	timeNow := time.Now().String()
	vote.CreatedAt = timeNow
	vote.UpdatedAt = timeNow

	uRef, err := fclient.fc.Query(fd.Ref(fd.Collection("Users"), userRef))
	if err != nil {
		return fmt.Errorf("retrieving ref for %s failed: %v", userRef, err)
	}
	var userReference fd.RefV
	err = uRef.Get(&userReference)
	if err != nil {
		return fmt.Errorf("unpacking ref for %s failed: %v", userRef, err)
	}
	vote.UserRef = userReference

	pRef, err := fclient.fc.Query(fd.Ref(fd.Collection("Posts"), postRef))
	if err != nil {
		return fmt.Errorf("retrieving ref for %s failed: %v", postRef, err)
	}
	var postReference fd.RefV
	err = pRef.Get(&postReference)
	if err != nil {
		return fmt.Errorf("unpacking ref for %s failed: %v", postRef, err)
	}
	vote.PostRef = postReference
	return fclient.addRecord(VoteColName, vote)
}

func (fclient *FaunaClient) UpdateVote(voteRef int, vote Vote) error {
	return fclient.updateRecord(VoteColName, voteRef, vote)
}

func (fclient *FaunaClient) DeleteVote(voteRef int) error {
	return fclient.deleteRecord(VoteColName, voteRef)
}

func (fclient *FaunaClient) RetrieveVote(voteRef int) (*Vote, error) {
	record, err := fclient.retrieveRecord(VoteColName, voteRef)
	if err != nil {
		return nil, err
	}
	var vote Vote
	err = record.At(fd.ObjKey("data")).Get(&vote)
	if err != nil {
		return nil, err
	}
	return &vote, nil
}

func (fclient *FaunaClient) RetrieveVoteCollection() ([]Vote, error) {
	collection, err := fclient.retrieveCollection(VoteColName)
	if err != nil {
		return nil, err
	}
	var voteList []Vote
	var vote Vote

	for _, item := range collection {
		err := item.At(fd.ObjKey("data")).Get(&vote)
		if err != nil {
			return nil, err
		}
		voteList = append(voteList, vote)
	}
	if err != nil {
		return nil, err
	}
	return voteList, nil
}
