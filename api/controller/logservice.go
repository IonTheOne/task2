package controller

import (
	"github.com/Mlstermass/task2/pkg/env"
	"github.com/Mlstermass/task2/storage"
)

type LogService struct {
	config  env.Config
	storage storage.DocumentActions
}

func NewLogService(
	config env.Config,
	storage storage.DocumentActions,
) LogService {
	return LogService{
		config:  config,
		storage: storage,
	}
}


func (ls *LogService) GetLogs() {

}

func (ls *LogService) GetLog() {

}

