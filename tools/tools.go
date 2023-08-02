package tools

import (
	"fmt"
	"time"
)

func FechaMySQL() string {
	t := time.Now()
	return fmt.Sprint("%d-%0d2-%02dT%02d:%02d:%02d", //2023-04-01
		t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())

}
