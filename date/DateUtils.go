package date

import "time"

func FormatDateYmd(t time.Time) string {
	return t.Format("2006-01-02")
}

var (
	start = "2022-01-01"

	newYear              = []string{"2022-01-01", "2022-01-02", "2022-01-03"}
	newYearWorkday       []string
	springNewYear        = []string{"2022-01-31", "2022-02-01", "2022-02-02", "2022-02-03", "2022-02-04", "2022-02-05", "2022-02-06"}
	springNewYearWorkday = []string{"2022-01-29", "2022-01-30"}
	qingming             = []string{"2022-04-03", "2022-04-04", "2022-04-05"}
	qingmingWorkday      = []string{"2022-04-02"}
	laodong              = []string{"2022-05-01", "2022-05-02", "2022-05-03", "2022-05-04"}
	laodongWorkday       = []string{"2022-04-30", "2022-05-07"}
	duanwu               = []string{"2022-06-03", "2022-06-04", "2022-06-05"}
	duanwuWorkday        []string
	zhnogqiu             = []string{"2022-09-10", "2022-09-11", "2022-09-12"}
	zhongqiuWorkday      []string
	guoqing              = []string{"2022-10-01", "2022-10-02", "2022-10-03", "2022-10-04", "2022-10-05", "2022-10-06", "2022-10-07"}
	guoqingWorkday       = []string{"2022-10-08", "2022-10-09"}

	holiday []string
	workday []string

	isWorkdayMap = make(map[string]int)

	result = make(map[string]int)
)

func InitDate() {
	holiday = append(holiday, newYear...)
	holiday = append(holiday, springNewYear...)
	holiday = append(holiday, qingming...)
	holiday = append(holiday, laodong...)
	holiday = append(holiday, duanwu...)
	holiday = append(holiday, zhnogqiu...)
	holiday = append(holiday, guoqing...)
	for _, v := range holiday {
		isWorkdayMap[v] = 0
	}

	workday = append(workday, newYearWorkday...)
	workday = append(workday, springNewYearWorkday...)
	workday = append(workday, qingmingWorkday...)
	workday = append(workday, laodongWorkday...)
	workday = append(workday, duanwuWorkday...)
	workday = append(workday, zhongqiuWorkday...)
	workday = append(workday, guoqingWorkday...)

	for _, v := range workday {
		isWorkdayMap[v] = 1
	}
}

func IsWeekday(t time.Time) int {
	b, ok := isWorkdayMap[t.Format("2006-01-02")]
	if ok {
		return b
	} else {
		if t.Weekday() != time.Saturday && t.Weekday() != time.Sunday {
			return 1
		} else {
			return 0
		}

	}
}
