package db

import (
	f "github.com/fauna/faunadb-go/v4/faunadb"
)

type FaunaClient struct {
	fc *f.FaunaClient
}

func NewFaunaClient(c *f.FaunaClient) *FaunaClient {
	return &FaunaClient{
		fc: c,
	}
}
