package nyb

import "time"

//Set the target year
var target = func() time.Time {
	tmp := time.Now().UTC()
	if tmp.Month() == time.January && tmp.Day() < 2 {
		return time.Date(tmp.Year(), time.January, 1, 0, 0, 0, 0, time.UTC)
	}
	//Debug target
	//return time.Date(tmp.Year(), time.October, 23, 6, 0, 0, 0, time.UTC)
	return time.Date(tmp.Year()+1, time.January, 1, 0, 0, 0, 0, time.UTC)
}()
