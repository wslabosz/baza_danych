package db

import fd "github.com/fauna/faunadb-go/v4/faunadb"

func (fclient *FaunaClient) CreateIndexes() error {
	arrayOfIndexCreateFunc := []func() error{
		fclient.createPostsByUserIndex, fclient.createCommentsByUserIndex, fclient.createVotesByUserIndex,
		fclient.createCommentsByPostIndex, fclient.createVotesByPostIndex, fclient.createFollowersByUserIndex,
		fclient.createFollowingByUserIndex,
	}
	for _, f := range arrayOfIndexCreateFunc {
		err := f()
		if err != nil {
			return err
		}
	}
	return nil
}

func (fclient *FaunaClient) createPostsByUserIndex() error {
	_, err := fclient.fc.Query(
		fd.If(
			fd.Exists(fd.Index("posts_by_user")),
			nil,
			fd.CreateIndex(
				fd.Obj{
					"name":   "posts_by_user",
					"source": fd.Collection("Posts"),
					"terms": fd.Arr{
						fd.Obj{
							"field": fd.Arr{"data", "user_ref"},
						},
					},
				},
			),
		),
	)
	if err != nil {
		return err
	}
	return nil
}

func (fclient *FaunaClient) createCommentsByUserIndex() error {
	_, err := fclient.fc.Query(
		fd.If(
			fd.Exists(fd.Index("comments_by_user")),
			nil,
			fd.CreateIndex(
				fd.Obj{
					"name":   "comments_by_user",
					"source": fd.Collection("Comments"),
					"terms": fd.Arr{
						fd.Obj{
							"field": fd.Arr{"data", "user_ref"},
						},
					},
				},
			),
		),
	)
	if err != nil {
		return err
	}
	return nil
}

func (fclient *FaunaClient) createVotesByUserIndex() error {
	_, err := fclient.fc.Query(
		fd.If(
			fd.Exists(fd.Index("votes_by_user")),
			nil,
			fd.CreateIndex(
				fd.Obj{
					"name":   "votes_by_user",
					"source": fd.Collection("Votes"),
					"terms": fd.Arr{
						fd.Obj{
							"field": fd.Arr{"data", "user_ref"},
						},
					},
				},
			),
		),
	)
	if err != nil {
		return err
	}
	return nil
}

func (fclient *FaunaClient) createCommentsByPostIndex() error {
	_, err := fclient.fc.Query(
		fd.If(
			fd.Exists(fd.Index("comments_by_post")),
			nil,
			fd.CreateIndex(
				fd.Obj{
					"name":   "comments_by_post",
					"source": fd.Collection("Comments"),
					"terms": fd.Arr{
						fd.Obj{
							"field": fd.Arr{"data", "post_ref"},
						},
					},
				},
			),
		),
	)
	if err != nil {
		return err
	}
	return nil
}

func (fclient *FaunaClient) createVotesByPostIndex() error {
	_, err := fclient.fc.Query(
		fd.If(
			fd.Exists(fd.Index("votes_by_post")),
			nil,
			fd.CreateIndex(
				fd.Obj{
					"name":   "votes_by_post",
					"source": fd.Collection("Votes"),
					"terms": fd.Arr{
						fd.Obj{
							"field": fd.Arr{"data", "post_ref"},
						},
					},
				},
			),
		),
	)
	if err != nil {
		return err
	}
	return nil
}

func (fclient *FaunaClient) createFollowersByUserIndex() error {
	_, err := fclient.fc.Query(
		fd.If(
			fd.Exists(fd.Index("followers_by_user")),
			nil,
			fd.CreateIndex(
				fd.Obj{
					"name":   "followers_by_user",
					"source": fd.Collection("Follows"),
					"terms": fd.Arr{
						fd.Obj{
							"field": fd.Arr{"data", "follower_ref"},
						},
					},
				},
			),
		),
	)
	if err != nil {
		return err
	}
	return nil
}

func (fclient *FaunaClient) createFollowingByUserIndex() error {
	_, err := fclient.fc.Query(
		fd.If(
			fd.Exists(fd.Index("following_by_user")),
			nil,
			fd.CreateIndex(
				fd.Obj{
					"name":   "following_by_user",
					"source": fd.Collection("Follows"),
					"terms": fd.Arr{
						fd.Obj{
							"field": fd.Arr{"data", "following_ref"},
						},
					},
				},
			),
		),
	)
	if err != nil {
		return err
	}
	return nil
}
