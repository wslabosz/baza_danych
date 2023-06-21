package db

import (
	"fmt"

	fd "github.com/fauna/faunadb-go/v4/faunadb"
)

type PostDeep struct {
	Post     Post
	Comments []Comment
	Votes    []Vote
}

type UserFollow struct {
	User        User
	Follows     []Follow
	FollowingBy []Follow
}

type UserDeep struct {
	User     User
	Posts    []Post
	Comments []Comment
	Votes    []Vote
}

func (fclient *FaunaClient) GetPostsDeep() ([]PostDeep, error) {
	list, err := fclient.fc.Query(
		fd.Map(
			fd.Paginate(
				fd.Documents(
					fd.Collection("Posts"),
				),
			),
			fd.Lambda("p",
				fd.Let().
					Bind("post", fd.Get(fd.Var("p"))).
					Bind("comments", fd.Map(
						fd.Paginate(fd.MatchTerm(fd.Index("comments_by_post"), fd.Var("p"))),
						fd.Lambda(fd.Arr{"post_ref"}, fd.Get(fd.Var("post_ref"))),
					)).
					Bind("votes", fd.Map(
						fd.Paginate(fd.MatchTerm(fd.Index("votes_by_post"), fd.Var("p"))),
						fd.Lambda(fd.Arr{"post_ref"}, fd.Get(fd.Var("post_ref"))),
					)).
					In(fd.Obj{
						"post":     fd.Var("post"),
						"comments": fd.Var("comments"),
						"votes":    fd.Var("votes"),
					}),
			),
		),
	)
	if err != nil {
		return nil, err
	}
	data := list.At(fd.ObjKey("data"))
	if data == nil {
		return nil, fmt.Errorf("failed to retrieve data from result")
	}

	var collection fd.ArrayV
	err = data.Get(&collection)
	if err != nil {
		return nil, err
	}

	var postList []PostDeep
	var post Post

	for _, item := range collection {
		var singleEntry fd.ObjectV
		var manyEntries fd.ArrayV

		commentsArray := item.At(fd.ObjKey("comments"))
		err = commentsArray.Get(&singleEntry)
		if err != nil {
			return nil, err
		}
		commentsObj := singleEntry.At(fd.ObjKey("data"))
		err = commentsObj.Get(&manyEntries)
		if err != nil {
			return nil, err
		}
		var comments []Comment
		var comment Comment
		for _, item := range manyEntries {
			err := item.At(fd.ObjKey("data")).Get(&comment)
			if err != nil {
				return nil, err
			}
			comments = append(comments, comment)
		}

		votesArray := item.At(fd.ObjKey("votes"))
		err = votesArray.Get(&singleEntry)
		if err != nil {
			return nil, err
		}
		votesObj := singleEntry.At(fd.ObjKey("data"))
		err = votesObj.Get(&manyEntries)
		if err != nil {
			return nil, err
		}
		var votes []Vote
		var vote Vote
		for _, item := range manyEntries {
			err := item.At(fd.ObjKey("data")).Get(&vote)
			if err != nil {
				return nil, err
			}
			votes = append(votes, vote)
		}

		postArray := item.At(fd.ObjKey("post"))
		err = postArray.Get(&singleEntry)
		if err != nil {
			return nil, err
		}
		postObj := singleEntry.At(fd.ObjKey("data"))
		err = postObj.Get(&post)
		if err != nil {
			return nil, err
		}

		postEntry := PostDeep{
			Post:     post,
			Comments: comments,
			Votes:    votes,
		}
		postList = append(postList, postEntry)
	}
	return postList, nil
}

func (fclient *FaunaClient) GetUserFollowersAndFollowing() (any, error) {
	list, err := fclient.fc.Query(
		fd.Map(
			fd.Paginate(
				fd.Documents(
					fd.Collection("Users"),
				),
			),
			fd.Lambda("u",
				fd.Let().
					Bind("user", fd.Get(fd.Var("u"))).
					Bind("followers", fd.Map(
						fd.Paginate(fd.MatchTerm(fd.Index("followers_by_user"), fd.Var("u"))),
						fd.Lambda(fd.Arr{"follower_ref"}, fd.Get(fd.Var("follower_ref"))),
					)).
					Bind("following", fd.Map(
						fd.Paginate(fd.MatchTerm(fd.Index("following_by_user"), fd.Var("u"))),
						fd.Lambda(fd.Arr{"following_ref"}, fd.Get(fd.Var("following_ref"))),
					)).
					In(fd.Obj{
						"user":      fd.Var("user"),
						"followers": fd.Var("followers"),
						"following": fd.Var("following"),
					}),
			),
		),
	)
	if err != nil {
		return nil, err
	}
	data := list.At(fd.ObjKey("data"))
	if data == nil {
		return nil, fmt.Errorf("failed to retrieve data from result")
	}

	var collection fd.ArrayV
	err = data.Get(&collection)
	if err != nil {
		return nil, err
	}

	var userList []UserFollow
	var user User

	for _, item := range collection {
		var singleEntry fd.ObjectV
		var manyEntries fd.ArrayV

		followersArray := item.At(fd.ObjKey("followers"))
		err = followersArray.Get(&singleEntry)
		if err != nil {
			return nil, err
		}
		followersObj := singleEntry.At(fd.ObjKey("data"))
		err = followersObj.Get(&manyEntries)
		if err != nil {
			return nil, err
		}
		var followers []Follow
		var follow Follow
		for _, item := range manyEntries {
			err := item.At(fd.ObjKey("data")).Get(&follow)
			if err != nil {
				return nil, err
			}
			followers = append(followers, follow)
		}

		followingArray := item.At(fd.ObjKey("following"))
		err = followingArray.Get(&singleEntry)
		if err != nil {
			return nil, err
		}
		followingObj := singleEntry.At(fd.ObjKey("data"))
		err = followingObj.Get(&manyEntries)
		if err != nil {
			return nil, err
		}
		var following []Follow
		for _, item := range manyEntries {
			err := item.At(fd.ObjKey("data")).Get(&follow)
			if err != nil {
				return nil, err
			}
			following = append(following, follow)
		}

		userArray := item.At(fd.ObjKey("user"))
		err = userArray.Get(&singleEntry)
		if err != nil {
			return nil, err
		}
		userObj := singleEntry.At(fd.ObjKey("data"))
		err = userObj.Get(&user)
		if err != nil {
			return nil, err
		}

		userFollowEntry := UserFollow{
			User:        user,
			Follows:     following,
			FollowingBy: followers,
		}
		userList = append(userList, userFollowEntry)
	}
	fmt.Printf("%+v", userList)
	return userList, nil
}

func (fclient *FaunaClient) GetUserDeep() ([]UserDeep, error) {
	list, err := fclient.fc.Query(
		fd.Map(
			fd.Paginate(
				fd.Documents(
					fd.Collection("Users"),
				),
			),
			fd.Lambda("u",
				fd.Let().
					Bind("user", fd.Get(fd.Var("u"))).
					Bind("posts", fd.Map(
						fd.Paginate(fd.MatchTerm(fd.Index("posts_by_user"), fd.Var("u"))),
						fd.Lambda(fd.Arr{"user_ref"}, fd.Get(fd.Var("user_ref"))),
					)).
					Bind("comments", fd.Map(
						fd.Paginate(fd.MatchTerm(fd.Index("comments_by_user"), fd.Var("u"))),
						fd.Lambda(fd.Arr{"user_ref"}, fd.Get(fd.Var("user_ref"))),
					)).
					Bind("votes", fd.Map(
						fd.Paginate(fd.MatchTerm(fd.Index("votes_by_user"), fd.Var("u"))),
						fd.Lambda(fd.Arr{"user_ref"}, fd.Get(fd.Var("user_ref"))),
					)).
					In(fd.Obj{
						"user":     fd.Var("user"),
						"posts":    fd.Var("posts"),
						"comments": fd.Var("comments"),
						"votes":    fd.Var("votes"),
					}),
			),
		),
	)
	if err != nil {
		return nil, err
	}
	data := list.At(fd.ObjKey("data"))
	if data == nil {
		return nil, fmt.Errorf("failed to retrieve data from result")
	}

	var collection fd.ArrayV
	err = data.Get(&collection)
	if err != nil {
		return nil, err
	}

	var userList []UserDeep
	var user User

	for _, item := range collection {
		var singleEntry fd.ObjectV
		var manyEntries fd.ArrayV

		postsArray := item.At(fd.ObjKey("posts"))
		err = postsArray.Get(&singleEntry)
		if err != nil {
			return nil, err
		}
		postsObj := singleEntry.At(fd.ObjKey("data"))
		err = postsObj.Get(&manyEntries)
		if err != nil {
			return nil, err
		}
		var posts []Post
		var post Post
		for _, item := range manyEntries {
			err := item.At(fd.ObjKey("data")).Get(&post)
			if err != nil {
				return nil, err
			}
			posts = append(posts, post)
		}

		commentsArray := item.At(fd.ObjKey("comments"))
		err = commentsArray.Get(&singleEntry)
		if err != nil {
			return nil, err
		}
		commentsObj := singleEntry.At(fd.ObjKey("data"))
		err = commentsObj.Get(&manyEntries)
		if err != nil {
			return nil, err
		}
		var comments []Comment
		var comment Comment
		for _, item := range manyEntries {
			err := item.At(fd.ObjKey("data")).Get(&comment)
			if err != nil {
				return nil, err
			}
			comments = append(comments, comment)
		}

		votesArray := item.At(fd.ObjKey("votes"))
		err = votesArray.Get(&singleEntry)
		if err != nil {
			return nil, err
		}
		votesObj := singleEntry.At(fd.ObjKey("data"))
		err = votesObj.Get(&manyEntries)
		if err != nil {
			return nil, err
		}
		var votes []Vote
		var vote Vote
		for _, item := range manyEntries {
			err := item.At(fd.ObjKey("data")).Get(&vote)
			if err != nil {
				return nil, err
			}
			votes = append(votes, vote)
		}

		userArray := item.At(fd.ObjKey("user"))
		err = userArray.Get(&singleEntry)
		if err != nil {
			return nil, err
		}
		userObj := singleEntry.At(fd.ObjKey("data"))
		err = userObj.Get(&user)
		if err != nil {
			return nil, err
		}

		userEntry := UserDeep{
			User:     user,
			Posts:    posts,
			Comments: comments,
			Votes:    votes,
		}
		userList = append(userList, userEntry)
	}

	fmt.Printf("%+v", userList)
	return userList, nil
}
