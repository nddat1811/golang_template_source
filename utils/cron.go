package utils

import (
	"fmt"
	"log"
	"time"

	"golang_template_source/config"
	"golang_template_source/domain"
	"golang_template_source/repository"

	"github.com/go-co-op/gocron/v2"
)

var (
	HOUR_SYNC_VGA      = 14
	MINUTE_SYNC_VGA    = 31
	HOUR_DELETE_LOG    = 2
	MINUTE_DELETE_LOG  = 0
)

func StartScheduler() {
	// Create scheduler
	scheduler, err := gocron.NewScheduler()
	if err != nil {
		log.Fatal("Failed to create scheduler:", err)
	}

	// Schedule jobs
	_, err = scheduler.NewJob(
		gocron.CronJob(buildCronExpression(14, 31), false),
		gocron.NewTask(syncUsersDaily),
	)
	if err != nil {
		log.Fatal("Failed to schedule job:", err) //ghi một thông báo log ra stderr (standard error) với mức độ nghiêm trọng cao, sau đó kết thúc chương trình ngay lập tức bằng os.Exit(1).
	}

	// Start scheduler
	scheduler.Start()
	log.Println("Scheduler started and jobs added.")
}

func buildCronExpression(hour, minute int) string {
	return fmt.Sprintf("%d %d * * *", minute, hour)
}

func syncUsersDaily() {
	log.Println("syncUsersDaily started")
	con := config.InitPostgreSQL()
	defer config.CloseConnectDB(con)

	repo := repository.NewSysLogRepository(con)

	log := domain.SysLog{
		ActionDatetime:   time.Now(),
		PathName:         "T",
		Method:           "c.",
		IP:               "realIP",
		StatusResponse:   1,
		Response:         "responseBody",
		Description:      "Request logged by middleware",
		RequestBody:      "requestBody",
		RequestQuery:     "string(requestQuery)",
		Duration:         0.12,
	}

	repo.InsertLog(&log)
}
