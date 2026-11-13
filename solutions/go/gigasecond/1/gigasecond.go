package gigasecond

import "time"

// AddGigasecond returns t plus one gigasecond (1,000,000,000 seconds).
func AddGigasecond(t time.Time) time.Time {
	return t.Add(1_000_000_000 * time.Second)
}
