package controller

import (
	"context"
	"log"

	"github.com/Mlstermass/task2/pkg/env"
	pb "github.com/Mlstermass/task2/pkg/proto"
	"github.com/Mlstermass/task2/storage"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type LogService struct {
	pb.UnimplementedLogServiceServer
	config  env.Config
	storage storage.DocumentActions
	tokens  map[string]bool
}

func NewLogService(
	config env.Config,
	storage storage.DocumentActions,
) *LogService {
	return &LogService{
		config:  config,
		storage: storage,
		tokens:  make(map[string]bool),
	}
}

func (ls *LogService) StoreLog(ctx context.Context, in *pb.StoreLogRequest) (*emptypb.Empty, error) {
	// Check the token
	if !ls.tokens[in.Token] {
		return nil, status.Error(codes.Unauthenticated, "invalid token")
	}
	err := ls.storage.StoreLog(ctx, in)
	if err != nil {
		log.Printf("Error storing logs: %v", err)
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (ls *LogService) StoreLogBatch(ctx context.Context, in *pb.StoreLogBatchRequest) (*emptypb.Empty, error) {
	// Check the token
	if !ls.tokens[in.Token] {
		return nil, status.Error(codes.Unauthenticated, "invalid token")
	}
	for _, logLine := range in.LogLines {
		err := ls.storage.StoreLog(ctx, &pb.StoreLogRequest{
			Id:        logLine.Id,
			LogLine:   logLine.LogLine,
			Timestamp: logLine.Timestamp,
			Bucket:    in.Bucket,
			Token:     in.Token,
		})
		if err != nil {
			log.Printf("Error storing log line: %v", err)
			return nil, err
		}
	}
	return &emptypb.Empty{}, nil
}

func (ls *LogService) GetLogs(ctx context.Context, in *pb.GetLogsRequest) (*pb.GetLogsResponse, error) {
	// Check the token
	if !ls.tokens[in.Token] {
		return nil, status.Error(codes.Unauthenticated, "invalid token")
	}

	logs, err := ls.storage.GetLogs(ctx, in)
	if err != nil {
		log.Printf("Error getting logs: %v", err)
		return nil, err
	}
	return logs, nil
}

func (ls *LogService) GetLogCount(ctx context.Context, in *pb.GetLogCountRequest) (*pb.GetLogCountResponse, error) {
	// Check the token
	if !ls.tokens[in.Token] {
		return nil, status.Error(codes.Unauthenticated, "invalid token")
	}
	count, err := ls.storage.GetLogCount(ctx, in)
	if err != nil {
		log.Printf("Error getting log count: %v", err)
		return nil, err
	}
	return count, nil
}

func (ls *LogService) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	// Check the username and password
	if in.Username != "admin" || in.Password != "admin123" {
		return nil, status.Error(codes.Unauthenticated, "invalid username or password")
	}

	// Generate a token
	token := uuid.New().String()

	// Store the token
	ls.tokens[token] = true

	return &pb.LoginResponse{Token: token}, nil
}
