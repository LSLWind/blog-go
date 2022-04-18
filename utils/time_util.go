package utils

import "time"

/**
时间处理工具类
*/

type TimeUtil struct {
}

var Year_MONTH_DAY_FORMATTER = "2006-01-02"

func (t TimeUtil) GetYearMonthDay(time time.Time) string {
	return time.Format(Year_MONTH_DAY_FORMATTER)
}
