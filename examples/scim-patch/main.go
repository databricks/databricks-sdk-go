package main

// A simple example of fetching a user by ID and using the Patch API to update
// their name.
//
// Usage:
//   go run main.go <host> <token> <user-id>

import (
	"context"
	"fmt"
	"os"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/logger"
	"github.com/databricks/databricks-sdk-go/service/iam"
)

func fetchAndPrintUserName(c *databricks.AccountClient, id string) error {
	user, err := c.Users.GetById(context.Background(), id)
	if err != nil {
		return fmt.Errorf("error fetching user: %s", err)
	}
	fmt.Println("User name:", user.DisplayName)
	fmt.Println("User roles:", user.Roles)

	return nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <user name>\n", os.Args[0])
		os.Exit(1)
	}
	name := os.Args[1]

	// Login via databricks auth login https://accounts.cloud.databricks.com --account-id ...

	// Create a client
	cfg := databricks.Config{
		Host: "https://accounts.cloud.databricks.com",
		// Specify Account ID via the DATABRICKS_ACCOUNT_ID environment variable
	}
	logger.DefaultLogger = &logger.SimpleLogger{
		Level: logger.LevelTrace,
	}
	c, err := databricks.NewAccountClient(&cfg)
	if err != nil {
		fmt.Printf("Error creating client: %s\n", err)
		os.Exit(1)
	}

	// Fetch the user ID
	users, err := c.Users.ListAll(context.Background(), iam.ListAccountUsersRequest{
		Filter: "displayName sw " + name,
	})
	if err != nil {
		fmt.Printf("Error fetching user: %s\n", err)
		os.Exit(1)
	}
	userID := users[0].Id

	// Fetch the user
	err = fetchAndPrintUserName(c, userID)
	if err != nil {
		fmt.Printf("Error fetching user: %s\n", err)
		os.Exit(1)
	}

	// Make non admin
	pu := iam.PartialUpdate{
		Id: userID,
		// Note that this request works without schemas
		Operations: []iam.Patch{
			{
				Op:   "remove",
				Path: "roles",
				Value: []map[string]string{
					{"value": "account_admin"},
				},
			},
		},
	}
	err = c.Users.Patch(context.Background(), pu)
	if err != nil {
		fmt.Printf("Error updating user: %s\n", err)
		os.Exit(1)
	}

	// Fetch user and print its name
	err = fetchAndPrintUserName(c, userID)
	if err != nil {
		fmt.Printf("Error fetching user: %s\n", err)
		os.Exit(1)
	}

	// restore the old name
	pu = iam.PartialUpdate{
		Id: userID,
		// It also works with schemas
		Schema: []iam.PatchSchema{
			iam.PatchSchemaUrnIetfParamsScimApiMessagesPatchop,
		},
		Operations: []iam.Patch{
			{
				Op: "add",
				Value: map[string][]map[string]string{
					"roles": {
						{"value": "account_admin"},
					},
				},
			},
		},
	}
	err = c.Users.Patch(context.Background(), pu)
	if err != nil {
		fmt.Printf("Error updating user: %s\n", err)
		os.Exit(1)
	}

	// Fetch user and print its name
	err = fetchAndPrintUserName(c, userID)
	if err != nil {
		fmt.Printf("Error fetching user: %s\n", err)
		os.Exit(1)
	}
}
