package main

import (
	"context"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/google/uuid"
	"github.com/hay-kot/git-web-template/backend/app/cli/reader"
	"github.com/hay-kot/git-web-template/backend/internal/types"
	"github.com/hay-kot/git-web-template/backend/pkgs/hasher"
	"github.com/urfave/cli/v2"
)

func (a *app) UserCreate(c *cli.Context) error {
	var defaultValidators = []reader.StringValidator{
		reader.StringRequired,
		reader.StringNoLeadingOrTrailingWhitespace,
	}
	// Get Flags
	name := reader.ReadString("Name: ",
		defaultValidators...,
	)
	password := reader.ReadString("Password: ",
		defaultValidators...,
	)

	email := reader.ReadString("Email: ",
		reader.StringRequired,
		reader.StringNoLeadingOrTrailingWhitespace,
		reader.StringContainsAt,
	)
	isSuper := reader.ReadBool("Is Superuser?")

	pwHash, err := hasher.HashPassword(password)
	if err != nil {
		return err
	}

	usr := types.UserCreate{
		Name:        name,
		Email:       email,
		Password:    pwHash,
		IsSuperuser: isSuper,
	}

	_, err = a.repos.Users.Create(context.Background(), usr)

	if err == nil {
		fmt.Println("Super user created")
	}
	return err
}

func (a *app) UserDelete(c *cli.Context) error {
	// Get Flags
	id := c.String("id")
	uid := uuid.MustParse(id)

	fmt.Printf("Deleting user with id: %s\n", id)

	// Confirm Action
	fmt.Printf("Are you sure you want to delete this user? (y/n) ")
	var answer string
	_, err := fmt.Scanln(&answer)
	if answer != "y" || err != nil {
		fmt.Println("Aborting")
		return nil
	}

	err = a.repos.Users.Delete(context.Background(), uid)

	if err == nil {
		fmt.Printf("%v User(s) deleted (id=%v)\n", 1, id)
	}
	return err
}

func (a *app) UserList(c *cli.Context) error {
	fmt.Println("Superuser List")

	users, err := a.repos.Users.GetAll(context.Background())

	if err != nil {
		return err
	}

	tabWriter := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	defer func(tabWriter *tabwriter.Writer) {
		_ = tabWriter.Flush()
	}(tabWriter)

	_, err = fmt.Fprintln(tabWriter, "Id\tName\tEmail\tIsSuper")

	if err != nil {
		return err
	}

	for _, u := range users {
		_, _ = fmt.Fprintf(tabWriter, "%v\t%s\t%s\t%v\n", u.ID, u.Name, u.Email, u.IsSuperuser)
	}

	return nil
}
