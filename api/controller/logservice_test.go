package controller

import (
    "context"
    "testing"

    "github.com/Mlstermass/task2/pkg/env"
    pb "github.com/Mlstermass/task2/pkg/proto"
    "github.com/Mlstermass/task2/storage"
    "github.com/stretchr/testify/assert"
)

func TestStoreLog(t *testing.T) {
    // Create a new LogService with a mock storage
    ls := NewLogService(env.Config{}, &storage.MockDocumentActions{
        StoreLogFunc: func(ctx context.Context, log *pb.StoreLogRequest) error {
            // Mock implementation
            return nil
        },
        // Implement other functions similarly
    })

    // Call the Login method with valid credentials
    resp, err := ls.Login(context.Background(), &pb.LoginRequest{
        Username: "admin",
        Password: "admin123",
    })

    // Check that no error was returned
    assert.NoError(t, err)

    // Check that a token was returned
    assert.NotEmpty(t, resp.Token)

    // Call the Login method with invalid credentials
    resp, err = ls.Login(context.Background(), &pb.LoginRequest{
        Username: "invalid",
        Password: "invalid",
    })

    // Check that an error was returned
    assert.Error(t, err)

}