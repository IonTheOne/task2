package immudb

import (
	"context"
	"log"

	pb "github.com/Mlstermass/task2/pkg/proto"
	"google.golang.org/protobuf/proto"
)

func (i *Immu) StoreLog(ctx context.Context, in *pb.StoreLogRequest) error {
	// Store the log line in the immudb database
	_, err := i.client.Set(ctx, []byte(in.Id), []byte(in.LogLine))
	if err != nil {
		log.Printf("failed to store log line: %v", err)
	}

	return nil
}

func (i *Immu) GetLogs(ctx context.Context, in *pb.GetLogsRequest) (*pb.GetLogsResponse, error) {
	var logs []*pb.StoreLogRequest

	// Get all keys from the immudb database
	keys, err := i.client.GetAll(ctx, [][]byte{})
	if err != nil {
		return nil, err
	}

	// Iterate over the keys
	for _, key := range keys.Entries {
		// Get the value for the key from the immudb database
		value, err := i.client.Get(ctx, key.GetKey())
		if err != nil {
			return nil, err
		}

		// Unmarshal the value into a Log
		var log pb.StoreLogRequest
		if err := proto.Unmarshal(value.Value, &log); err != nil {
			return nil, err
		}

		// Add the log to the logs slice
		logs = append(logs, &log)
	}

	// If LastX is greater than 0, return only the last X logs
	if in.LastX > 0 && len(logs) > int(in.LastX) {
		logs = logs[len(logs)-int(in.LastX):]
	}

	return &pb.GetLogsResponse{Logs: logs}, nil
}

func (i *Immu) GetLogCount(ctx context.Context, in *pb.GetLogCountRequest) (*pb.GetLogCountResponse, error) {
    var count int64

    // Get all keys from the immudb database
    keys, err := i.client.GetAll(ctx, [][]byte{})
    if err != nil {
        return nil, err
    }

    // Iterate over the keys
    for _, key := range keys.Entries {
        // Get the value for the key from the immudb database
        value, err := i.client.Get(ctx, key.GetKey())
        if err != nil {
            return nil, err
        }

        // Unmarshal the value into a Log
        var log pb.StoreLogRequest
        if err := proto.Unmarshal(value.Value, &log); err != nil {
            return nil, err
        }

        // If the bucket and token are not set, or if they match the log's bucket and token, increment the count
        if (in.Bucket == "" && in.Token == "") || (log.Bucket == in.Bucket && log.Token == in.Token) {
            count++
        }
    }

    return &pb.GetLogCountResponse{Count: int32(count)}, nil
}
