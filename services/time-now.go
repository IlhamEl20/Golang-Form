package services

import "time"

func TimeNow() time.Time {

	currentTime := time.Now().Local()
	// now := currentTime.Local().Format("2 Jan 2006 15:04")

	return currentTime
}