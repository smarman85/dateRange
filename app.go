//package main
package dateRange

import (
        "fmt"
        //"strings"
        "time"
)

// since=2018-12-03T10:00&until=2018-12-10T10:00

func stringifyTime(rawTime time.Time) string {
        // stringTime := rawTime.String()
        // return strings.Split(stringTime, " ")[0]
        return rawTime.Format("2018-12-01")
}

func timeInfo(day int) time.Time {

        var lastMonday time.Time
        switch day {
        case 0:
                lastMonday = time.Now().AddDate(0, 0, -6)
        case 1: // monday
                lastMonday = time.Now().AddDate(0, 0, -7)
        case 2:
                lastMonday = time.Now().AddDate(0, 0, -1)
        case 3:
                lastMonday = time.Now().AddDate(0, 0, -2)
        case 4:
                lastMonday = time.Now().AddDate(0, 0, -3)
        case 5:
                lastMonday = time.Now().AddDate(0, 0, -4)
        case 6:
                lastMonday = time.Now().AddDate(0, 0, -5)
        }
        return lastMonday
}

func queryString(isMonday bool, lastMonday string) string {

        var searchParam string
        if isMonday == true {
                today := stringifyTime(time.Now())
                searchParam = fmt.Sprintf("since=%sT10:00&until=%sT10:00", lastMonday, today)
        } else {
                searchParam = fmt.Sprintf("since=%sT10:00", lastMonday)
        }
        return searchParam
}

func Run() string {
        day := int(time.Now().Weekday())
        rawMonday := timeInfo(day)
        lastMonday := stringifyTime(rawMonday)
        var searchString string
        if day == 1 {
                searchString = queryString(true, lastMonday)
        } else {
                searchString = queryString(false, lastMonday)
        }
        //fmt.Println(searchString)
        return searchString

}
