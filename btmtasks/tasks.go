package btmtasks

import (
	"fmt"
	"strconv"
	"time"

	"github.com/madflojo/tasks"
	"open-btm.com/configs"
	"open-btm.com/database"
	"open-btm.com/models"
)

func ScheduledTasks() *tasks.Scheduler {

	//  initalizing scheduler for regullarly running tasks
	scheduler := tasks.New()

	// JWT signature salt will be updated at the IAM at provided interval
	// this is what the peice of code below does, using golangs task scheduler
	//  the below task runs the login to app twice between that time interval
	clear_run, _ := strconv.Atoi(configs.AppConfig.Get("JWT_SALT_LIFE_TIME"))
	clear_run = int(clear_run) / 2
	jwt_update_interval := time.Minute * time.Duration(clear_run)
	//  Task 2 for testing Make random heartbeat call
	if _, err := scheduler.Add(&tasks.Task{
		Interval: jwt_update_interval,
		TaskFunc: func() error {
			models.LoginBlueAdmin()
			return nil
		},
	}); err != nil {
		fmt.Println(err)
	}

	// // Add a task to move to Logs Directory Every Interval, Interval to Be Provided From Configuration File
	gormLoggerfile, _ := database.GormLoggerFile()
	//  App should not start
	log_file, _ := Logfile()
	if _, err := scheduler.Add(&tasks.Task{
		Interval: time.Duration(1 * time.Minute),
		TaskFunc: func() error {

			gormLoggerfile.Truncate(0)
			log_file.Truncate(0)
			return nil
		},
	}); err != nil {
		fmt.Println(err)

	}

	return scheduler
}
