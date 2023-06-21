package db

import (
	"fmt"

	fd "github.com/fauna/faunadb-go/v4/faunadb"
)

type User struct {
	Username  string `fauna:"username"`
	Email     string `fauna:"email"`
	CreatedAt string `fauna:"created_at"`
	UpdatedAt string `fauna:"updated_at"`
}

type Post struct {
	Title     string  `fauna:"title"`
	Content   string  `fauna:"content"`
	CreatedAt string  `fauna:"created_at"`
	UpdatedAt string  `fauna:"updated_at"`
	UserRef   fd.RefV `fauna:"user_ref"`
}

type Comment struct {
	Content   string  `fauna:"content"`
	CreatedAt string  `fauna:"created_at"`
	UpdatedAt string  `fauna:"updated_at"`
	UserRef   fd.RefV `fauna:"user_ref"`
	PostRef   fd.RefV `fauna:"post_ref"`
}

type Vote struct {
	Value     int     `fauna:"value"`
	CreatedAt string  `fauna:"created_at"`
	UpdatedAt string  `fauna:"updated_at"`
	UserRef   fd.RefV `fauna:"user_ref"`
	PostRef   fd.RefV `fauna:"post_ref"`
}

type Follow struct {
	CreatedAt    string  `fauna:"created_at"`
	UpdatedAt    string  `fauna:"updated_at"`
	FollowerRef  fd.RefV `fauna:"follower_ref"`
	FollowingRef fd.RefV `fauna:"following_ref"`
}

func (fclient *FaunaClient) InitalizeDatabase() error {
	// Create Users collection
	_, err := fclient.fc.Query(
		fd.If(
			fd.Exists(fd.Collection("Users")),
			nil,
			fd.CreateCollection(fd.Obj{"name": "Users"}),
		),
	)
	if err != nil {
		return fmt.Errorf("users collection creation failed: %v", err)
	}

	// Create Posts collection
	_, err = fclient.fc.Query(
		fd.If(
			fd.Exists(fd.Collection("Posts")),
			nil,
			fd.CreateCollection(fd.Obj{"name": "Posts"}),
		),
	)
	if err != nil {
		return fmt.Errorf("posts collection creation failed: %v", err)
	}

	// Create Comments collection
	_, err = fclient.fc.Query(
		fd.If(
			fd.Exists(fd.Collection("Comments")),
			nil,
			fd.CreateCollection(fd.Obj{"name": "Comments"}),
		),
	)
	if err != nil {
		return fmt.Errorf("comments collection creation failed: %v", err)
	}

	// Create Votes collection
	_, err = fclient.fc.Query(
		fd.If(
			fd.Exists(fd.Collection("Votes")),
			nil,
			fd.CreateCollection(fd.Obj{"name": "Votes"}),
		),
	)
	if err != nil {
		return fmt.Errorf("votes collection creation failed: %v", err)
	}

	// Create Follows collection
	_, err = fclient.fc.Query(
		fd.If(
			fd.Exists(fd.Collection("Follows")),
			nil,
			fd.CreateCollection(fd.Obj{"name": "Follows"}),
		),
	)
	if err != nil {
		return fmt.Errorf("follows collection creation failed: %v", err)
	}

	return nil
}
