package timer

import (
	"awesomeProject/app/task"
	"time"
)

func Timer() {
	ticker := time.NewTicker(time.Second)
	for i := 1; i == 1; i = 1 {
		t := <-ticker.C
		task.Task(t.Format(time.DateTime))
	}
}
