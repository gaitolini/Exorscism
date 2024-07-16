package giga_segundos

import "time"

const GIGASEGUNDO = 1000000000

func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * GIGASEGUNDO)
}
