package main

import (
	"fmt"

	f "github.com/fauna/faunadb-go/v4/faunadb"
	"github.com/wslabosz/baza_danych/db"
)

func main() {
	client := db.NewFaunaClient(f.NewFaunaClient(
		// dont worry its inactive
		"fnAFGP5D-2AAzT_Aifq-3Uv-jiz9ngL9A4NPKfiv",
		f.Endpoint("https://db.fauna.com/"),
	))
	err := client.InitalizeDatabase()
	if err != nil {
		fmt.Println(err)
	}
	err = client.CreateIndexes()
	if err != nil {
		fmt.Println(err)
	}
	// user := db.User{
	// 	Username:  "wojciech",
	// 	Email:     "test@example.com",
	// 	CreatedAt: time.Now().String(),
	// }
	// err = client.CreateUser(user)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// post := db.Post{
	// 	Title:   "third post",
	// 	Content: "very long content maybe 50 minutes",
	// }
	// client.CreatePost(post, "367330693205721292")

	// comment := db.Comment{
	// 	Content: "I don't like it",
	// }
	// err = client.CreateComment(comment, "367327539706724556", "368093581020233932")

	// vote := db.Vote{
	// 	Value: 1,
	// }
	// err = client.CreateVote(vote, "367327539706724556", "368093581020233932")

	// err = client.CreateFollow("367327539706724556", "367330693205721292")

	// test, err := client.RetrieveUserCollection()
	// if err != nil {
	// 	fmt.Println(err)
	// }


	// result, err := client.GetPostsDeep()
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// _, err = client.GetUserFollowersAndFollowing()
	// if err != nil {
	// 	fmt.Println(err)
	// }

	_, err = client.GetUserDeep()
	if err != nil {
		fmt.Println(err)
	}
}
