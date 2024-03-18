package main

import (
    "fmt"
    "log"
    "net"
    "net/http"

    "github.com/Mlstermass/task2/api/controller"
    "github.com/Mlstermass/task2/pkg/env"
    pb "github.com/Mlstermass/task2/pkg/proto"
    "github.com/Mlstermass/task2/storage"
    "github.com/Mlstermass/task2/storage/immudb"
    "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
    "google.golang.org/grpc"
)

func main() {
    // load env variables to the Config struct
    var conf env.Config
    env.LoadConfig(&conf)

    immuDBClient, err := immudb.NewImmuDBConn(conf)
    if err != nil {
        log.Fatalf("Failed to initialize immudb client: %v", err)
    }
    log.Printf("Connected to immudb at %s:%d", conf.ImmuDBAdress, conf.ImmuBDPort)

    immuDriver := immudb.NewImmu(immuDBClient, conf)
    ctl := newControllers(conf, immuDriver)

    // Create a new gRPC server
    grpcServer := grpc.NewServer()

    // Register the LogService with the gRPC server
    pb.RegisterLogServiceServer(grpcServer, ctl)

    // Listen on a port
    listener, err := net.Listen("tcp", fmt.Sprintf("%s", conf.AppHost))
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }
    log.Printf("Server started at %s", conf.AppHost)

    // Start the gRPC server in a goroutine so it doesn't block
    go func() {
        if err := grpcServer.Serve(listener); err != nil {
            log.Fatalf("Failed to serve: %v", err)
        }
    }()

    mux := runtime.NewServeMux()

    // Start the HTTP server
    http.ListenAndServe(":8081", mux)
}

func newControllers(
    config env.Config,
    immuDriver storage.DocumentActions,
) *controller.LogService {
    return controller.NewLogService(
        config, immuDriver)
}