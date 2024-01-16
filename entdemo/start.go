package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"entdemo/ent"
	"entdemo/ent/user"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("no args are given")
	}
	// client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	client, err := ent.Open("sqlite3", "file:entgo.db?_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	// Run the auto migration tool.
	switch os.Args[1] {
	case "migrate":
		log.Println("migrate")
		if err := client.Schema.Create(ctx); err != nil {
			log.Fatalf("failed creating schema resources: %v", err)
		}
	case "create":
		if len(os.Args) < 4 {
			log.Fatalf("go run start.go <user name> <age>")
		}
		log.Printf("create user: %s", os.Args[2])
		age, err := strconv.Atoi(os.Args[3])
    if err != nil {
			log.Fatalln(err)
    }
		if _, err = CreateUser(ctx, client, os.Args[2], age); err != nil {
			log.Fatalln(err)
		}
	case "query":
		if len(os.Args) < 3 {
			log.Fatalf("go run start.go <user name>")
		}
		log.Printf("query user: %s", os.Args[2])
		if _, err = QueryUser(ctx, client, os.Args[2]); err != nil {
			log.Fatalln(err)
		}
	}
}

func CreateUser(ctx context.Context, client *ent.Client, name string, age int) (*ent.User, error) {
	u, err := client.User.
		Create().
		SetAge(age).
		SetName(name).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("user was created: ", u)
	return u, nil
}

func QueryUser(ctx context.Context, client *ent.Client, name string) (*ent.User, error) {
	u, err := client.User.
		Query().
		Where(user.Name(name)).
		// `Only` fails if no user found,
		// or more than 1 user returned.
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}
	log.Println("user returned: ", u)
	return u, nil
}
