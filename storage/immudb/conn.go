package immudb

import (
	"context"
	"fmt"

	"github.com/Mlstermass/task2/pkg/env"
	"github.com/Mlstermass/task2/storage"
	"github.com/codenotary/immudb/pkg/client"
)

func NewImmuDBConn(conf env.Config) (client.ImmuClient, error) {
	// Establish connection to immudb
	client, err := connect(
		conf.ImmuDBAdress, conf.ImmuBDPort, conf.ImmuDBUser, conf.ImmuDBPassword)
	if err != nil {
		fmt.Printf("Error connecting to immudb: %v\n", err)
		return nil, err
	}
	defer client.Disconnect()
	return client, nil
}

// Function to establish a connection to immudb
func connect(host string, port int, username string, password string) (client.ImmuClient, error) {
	// Create a new client instance
	clt := client.NewClient()
	clt = clt.WithOptions(client.DefaultOptions())
	// Connect to immudb server
	err := clt.OpenSession(context.Background(), []byte(username), []byte(password), "defaultdb")
	if err != nil {
		return nil, fmt.Errorf("failed to connect to immudb: %v", err)
	}
	return clt, nil
}

type Immu struct {
	client client.ImmuClient
	conf   env.Config
}

func NewImmu(client client.ImmuClient, conf env.Config) storage.DocumentActions {
	return &Immu{client: client, conf: conf}
}
