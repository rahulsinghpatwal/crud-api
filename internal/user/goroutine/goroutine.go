package goroutine

import (
	"crud/internal/config"
	"crud/pkg/db"
	"crud/utils"
	"fmt"
	"os"
	"strings"
	"time"

	"go.uber.org/zap"
)

func Check(tk *time.Ticker) {

	duration := os.Getenv("DURATION")

	for range tk.C {
		CheckForSoftDelete(duration)
	}

}

func CheckForSoftDelete(duration string) {
	config := config.Load()

	d, err := time.ParseDuration("-" + duration)
	if err != nil {
		fmt.Println(err)
		return
	}
	timeduration := time.Now().Add(d)
	td := timeduration.String()
	stringtime := strings.Split(td, " ")
	last_accessed := stringtime[0] + " " + stringtime[1]
	db, err := db.CreateConnection(config)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var usersDetails []utils.SoftDeletedRecord

	rows, err := db.Query(`Select id, firstname from users where archived = true and last_accessed <= $1`, last_accessed)
	if err != nil {
		zap.S().Error(err)
		return
	}
	for rows.Next() {
		var userdata utils.SoftDeletedRecord
		err := rows.Scan(&userdata.Id, &userdata.Firstname)
		if err != nil {
			zap.S().Error(err)
			return

		}

		usersDetails = append(usersDetails, userdata)
	}

	_, err = db.Exec(`delete from users where archived = true and last_accessed <=$1`, last_accessed)
	if err != nil {
		zap.S().Error(err)
	}

	zap.S().Info(usersDetails)

}
