package date

import "time"

func FormatDateYmd(t time.Time) string {
	return t.Format("2006-01-02")
}
