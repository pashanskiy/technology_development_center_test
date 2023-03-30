package scheduler

import (
	"github.com/go-co-op/gocron"
	// "github.com/rs/zerolog"
	"gorm.io/gorm"
)

type TransactionSchedulerService struct {
	db *gorm.DB

	scheduler *gocron.Scheduler
}

func NewWalletTransactionService(
	db *gorm.DB,
	scheduler *gocron.Scheduler,

) *TransactionSchedulerService {
	return &TransactionSchedulerService{
		db:        db,
		scheduler: scheduler,
	}
}

// func (receiver *TransactionSchedulerService) runSchedule(logger zerolog.Logger) {
// 	_, err := receiver.scheduler.Every(5).Seconds().WaitForSchedule().SingletonMode().Do(receiver.depositJob(), logger)
// 	if err != nil {
// 		return
// 	}

// 	receiver.scheduler.StartAsync()
// }
