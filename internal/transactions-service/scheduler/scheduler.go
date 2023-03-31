package scheduler

import (
	"time"

	"github.com/go-co-op/gocron"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type TransactionSchedulerService struct {
	db *gorm.DB

	logger    *zerolog.Logger
	scheduler *gocron.Scheduler
}

func NewTransactionSchedulerService(
	db *gorm.DB,
	logger *zerolog.Logger,
) *TransactionSchedulerService {
	return &TransactionSchedulerService{
		db:        db,
		logger:    logger,
		scheduler: gocron.NewScheduler(time.Local),
	}
}

func (receiver *TransactionSchedulerService) RunSchedule(logger zerolog.Logger) {
	_, err := receiver.scheduler.Every(3).Seconds().WaitForSchedule().SingletonMode().Do(receiver.depositJob)
	if err != nil {
		logger.Error().Err(err).Msg("error shedule")
		return
	}

	_, err = receiver.scheduler.Every(5).Seconds().WaitForSchedule().SingletonMode().Do(receiver.withdrawalJob)
	if err != nil {
		logger.Error().Err(err).Msg("error shedule")
		return
	}

	receiver.scheduler.StartAsync()
}
