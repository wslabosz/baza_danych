package db

import (
	"github.com/fauna/faunadb-go/v4/faunadb"
)

const UserColName = "Users"

func (fclient *FaunaClient) CreateUser(user User) error {
	return fclient.addRecord(UserColName, user)
}

func (fclient *FaunaClient) UpdateUser(userRef int, user User) error {
	return fclient.updateRecord(UserColName, userRef, user)
}

func (fclient *FaunaClient) DeleteUser(userRef int) error {
	return fclient.deleteRecord(UserColName, userRef)
}

func (fclient *FaunaClient) RetrieveUser(userRef int) (*User, error) {
	record, err := fclient.retrieveRecord(UserColName, userRef)
	if err != nil {
		return nil, err
	}
	var user User
	err = record.At(faunadb.ObjKey("data")).Get(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (fclient *FaunaClient) RetrieveUserCollection() ([]User, error) {
	collection, err := fclient.retrieveCollection(UserColName)
	if err != nil {
		return nil, err
	}
	var userList []User
	var user User

	for _, item := range collection {
		err := item.At(faunadb.ObjKey("data")).Get(&user)
		if err != nil {
			return nil, err
		}
		userList = append(userList, user)
	}
	if err != nil {
		return nil, err
	}
	return userList, nil
}
