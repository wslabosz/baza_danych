package db

import (
	"fmt"

	"github.com/fauna/faunadb-go/v4/faunadb"
)

func (fclient *FaunaClient) addRecord(collectionName string, item interface{}) error {
	_, err := fclient.fc.Query(
		faunadb.Create(faunadb.Collection(collectionName), faunadb.Obj{"data": item}),
	)
	if err != nil {
		return fmt.Errorf("adding %v failed: %v", item, err)
	}

	return nil
}

func (fclient *FaunaClient) updateRecord(collectionName string, recordID int, item interface{}) error {
	_, err := fclient.fc.Query(
		faunadb.Update(faunadb.RefCollection(faunadb.Collection(collectionName), recordID), faunadb.Obj{"data": item}),
	)
	if err != nil {
		return fmt.Errorf("updating %v failed: %v", item, err)
	}
	return nil
}

func (fclient *FaunaClient) deleteRecord(collectionName string, recordID int) error {
	_, err := fclient.fc.Query(
		faunadb.Delete(faunadb.RefCollection(faunadb.Collection(collectionName), recordID)),
	)
	if err != nil {
		return fmt.Errorf("deleting %d failed: %v", recordID, err)
	}
	return nil
}

func (fclient *FaunaClient) retrieveRecord(collectionName string, recordID int) (faunadb.Value, error) {
	record, err := fclient.fc.Query(
		faunadb.Get(
			faunadb.RefCollection(faunadb.Collection(collectionName), recordID),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("retrieving %d failed: %v", recordID, err)
	}
	return record, nil
}

func (fclient *FaunaClient) retrieveCollection(collectionName string) (faunadb.ArrayV, error) {
	collection, err := fclient.fc.Query(
		faunadb.Map(
			faunadb.Paginate(faunadb.Documents(faunadb.Collection(collectionName))),
			faunadb.Lambda("data", faunadb.Get(faunadb.Var("data"))),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("retrieving collection %s failed: %v", collectionName, err)
	}

	data := collection.At(faunadb.ObjKey("data"))
	if data == nil {
		return nil, fmt.Errorf("failed to retrieve data from result")
	}

	var faunaArray faunadb.ArrayV
	err = collection.Get(&faunaArray)
	if err != nil {
		return nil, err
	}

	return faunaArray, nil
}
