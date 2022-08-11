package date

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"testing"
	"time"
)

func TestWorkday(t *testing.T) {
	InitDate()
	date, _ := time.ParseInLocation("2006-01-02", "2022-01-01", time.Local)
	loopOut, _ := time.ParseInLocation("2006-01-02", "2022-12-31", time.Local)
	i := 1

	f := excelize.NewFile()
	// 这里设置表头
	f.SetCellValue("Sheet1", "A1", "序号")
	f.SetCellValue("Sheet1", "B1", "日期")
	f.SetCellValue("Sheet1", "C1", "是否工作日")

	for {
		weekday := IsWeekday(date)
		fmt.Printf("%d %s %d\n", i, date.Format("2006-01-02"), weekday)
		f.SetCellValue("Sheet1", fmt.Sprintf("A%d", i), i)
		f.SetCellValue("Sheet1", fmt.Sprintf("B%d", i), date.Format("2006-01-02"))
		f.SetCellValue("Sheet1", fmt.Sprintf("C%d", i), weekday)
		if date.Equal(loopOut) {
			break
		}
		i++
		date = date.AddDate(0, 0, 1)
	}

	if err := f.SaveAs("out.xlsx"); err != nil {
		fmt.Println(err)
	}
}
