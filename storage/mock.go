package storage

import (
    "context"

    pb "github.com/Mlstermass/task2/pkg/proto"
)

type MockDocumentActions struct {
    StoreLogFunc      func(ctx context.Context, log *pb.StoreLogRequest) error
    GetLogsFunc       func(ctx context.Context, in *pb.GetLogsRequest) (*pb.GetLogsResponse, error)
    GetLogCountFunc   func(ctx context.Context, in *pb.GetLogCountRequest) (*pb.GetLogCountResponse, error)
}

func (m *MockDocumentActions) StoreLog(ctx context.Context, log *pb.StoreLogRequest) error {
    return m.StoreLogFunc(ctx, log)
}

func (m *MockDocumentActions) GetLogs(ctx context.Context, in *pb.GetLogsRequest) (*pb.GetLogsResponse, error) {
    return m.GetLogsFunc(ctx, in)
}

func (m *MockDocumentActions) GetLogCount(ctx context.Context, in *pb.GetLogCountRequest) (*pb.GetLogCountResponse, error) {
    return m.GetLogCountFunc(ctx, in)
}
