package btmtasks

import (
	"log"
	"os"
)

func Logfile() (*os.File, error) {

	// Custom File Writer for logging
	file, err := os.OpenFile("open-btm.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
		return nil, err
	}
	return file, nil

}
