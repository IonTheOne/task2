package storage

import (
	"context"

	pb "github.com/Mlstermass/task2/pkg/proto"
)
type DocumentActions interface {
	StoreLog(ctx context.Context, log *pb.StoreLogRequest) error
	GetLogs(ctx context.Context, in *pb.GetLogsRequest) (*pb.GetLogsResponse, error)
	GetLogCount(ctx context.Context, in *pb.GetLogCountRequest) (*pb.GetLogCountResponse, error)
	
}
