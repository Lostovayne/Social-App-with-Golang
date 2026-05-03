package db

import (
	"context"

	"github.com/Elevate-Techworks/social/internal/store"
)

var userNames = []string{
	"Alejandro", "Beatriz", "Carlos", "Daniela", "Eduardo", "Fernanda", "Gabriel", "Helena", "Ignacio", "Julia",
	"Kevin", "Lucia", "Manuel", "Natalia", "Oscar", "Patricia", "Ricardo", "Sofia", "Tomas", "Ursula",
	"Victor", "Wendy", "Xavier", "Yolanda", "Zacarias", "Adriana", "Bernardo", "Carmen", "David", "Elena",
	"Felipe", "Gloria", "Hugo", "Isabel", "Javier", "Karla", "Luis", "Marta", "Nicolas", "Olivia",
	"Pablo", "Raquel", "Santiago", "Teresa", "Ulises", "Valeria", "William", "Ximena", "Yosef", "Zoe",
}

func Seed(store store.Storage) error {
	ctx := context.Background()
	users := generateUsers(100)
	return nil
}

func generateUsers(num int) []*store.User {
	users := make([]*store.User, num)
	for i := 0; i < num; i++ {
		users[i] = &store.User{
			Username: userNames[i],
			Email:    userNames[i] + string(i+1) + "@example.com",
			Password: "123123",
		}
	}
	return users
}

func generatePosts(num int, users []*store.User) []*store.Post {
	posts := make([]*store.Post, num)
	for i := 0; i < num; i++ {
		posts[i] = &store.Post{
			Title:   "Post " + string(i+1),
			Content: "This is the content of post " + string(i+1),
		}
	}
	return posts
}
