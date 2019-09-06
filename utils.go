package gruz

import "time"

func ruzDate(t *time.Time) string {
    return t.Format("2006.01.02")
}
