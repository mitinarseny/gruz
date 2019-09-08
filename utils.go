package gruz

import (
    "sync"
    "time"
)

var (
    once sync.Once
    location *time.Location
)

func ruzDate(t *time.Time) string {
    return t.Format("2006.01.02")
}

func getLocation() *time.Location {
    once.Do(func() {
        location, _ = time.LoadLocation("Europe/Moscow")
    })
    return location
}
