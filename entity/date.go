package entity

import (
	"fmt"
	"strconv"
)

// Date .
type Date struct {
	Year, Month, Day, Hour, Minute int
}

func (mdate Date) init(tyear, tmonth, tday, thour, tminute int) {
	mdate.Year = tyear
	mdate.Month = tmonth
	mdate.Day = tday
	mdate.Hour = thour
	mdate.Minute = tminute
}

// GetYear .
func (mdate Date) GetYear() int {
	return mdate.Year
}

// SetYear .
func (mdate *Date) SetYear(tyear int) {
	mdate.Year = tyear
}

// GetMonth .
func (mdate Date) GetMonth() int {
	return mdate.Month
}

func (mdate *Date) SetMonth(tmonth int) {
	mdate.Month = tmonth
}

func (mdate Date) GetDay() int {
	return mdate.Day
}

func (mdate *Date) SetDay(tday int) {
	mdate.Day = tday
}

func (mdate Date) GetHour() int {
	return mdate.Hour
}

func (mdate *Date) SetHour(thour int) {
	mdate.Hour = thour
}

func (mdate Date) GetMinute() int {
	return mdate.Minute
}

func (mdate *Date) SetMinute(tminute int) {
	mdate.Minute = tminute
}

func StringToInt(s string) int {
	result, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("String to Int fail")
	}
	return result
}

func IsValid(tdate Date) bool {
	currentYear := tdate.GetYear()
	currentMonth := tdate.GetMonth()
	currentDay := tdate.GetDay()
	currentHour := tdate.GetHour()
	currentMinute := tdate.GetMinute()
	day := []int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if currentYear > 9999 || currentYear < 1000 {
		return false
	}
	if currentMonth > 12 || currentMonth < 1 {
		return false
	}
	if (currentYear%4 == 0 && currentYear%100 != 0) || currentYear%400 == 0 {
		day[2] = 29
	}
	if currentDay > day[currentMonth] || currentDay < 1 {
		return false
	}
	if currentHour > 23 || currentHour < 0 {
		return false
	}
	if currentMinute > 59 || currentMinute < 0 {
		return false
	}
	return true
}

func StringToDate(t_dateString string) Date {
	var t_date = Date{0, 0, 0, 0, 0}
	if len(t_dateString) != 16 || t_dateString[4] != '-' || t_dateString[7] != '-' || t_dateString[10] != '/' || t_dateString[13] != ':' {
		fmt.Println("the form of the Date is wrong!")
		return t_date
	}
	for i := 0; i < 16; i++ {
		if i == 4 || i == 7 || i == 10 || i == 13 {
			continue
		} else {
			if t_dateString[i] < '0' || t_dateString[i] > '9' {
				fmt.Println("the form of the Date is wrong!")
				return t_date
			}
		}
	}

	t_year, _ := strconv.Atoi(t_dateString[0:4])
	t_month, _ := strconv.Atoi(t_dateString[5:7])
	t_day, _ := strconv.Atoi(t_dateString[8:10])
	t_hour, _ := strconv.Atoi(t_dateString[11:13])
	t_min, _ := strconv.Atoi(t_dateString[14:16])
	t_date.SetYear(t_year)
	t_date.SetMonth(t_month)
	t_date.SetDay(t_day)
	t_date.SetHour(t_hour)
	t_date.SetMinute(t_min)

	return t_date
}

func IntToString(a int) string {
	resultstring := strconv.Itoa(a)
	return resultstring
}

func DateToString(tdate Date) string {
	initTime := "0000-00-00/00:00"
	if !IsValid(tdate) {
		dateString := initTime
		return dateString
	}
	dateString := IntToString(tdate.GetYear()) + "-" + IntToString(tdate.GetMonth()) +
		"-" + IntToString(tdate.GetDay()) + "/" + IntToString(tdate.GetHour()) + ":" + IntToString(tdate.GetMinute())
	return dateString
}

func (mdate Date) CopyDate(tdate Date) Date {
	mdate.SetYear(tdate.GetYear())
	mdate.SetMonth(tdate.GetMonth())
	mdate.SetDay(tdate.GetDay())
	mdate.SetMinute(tdate.GetMinute())
	mdate.SetHour(tdate.GetHour())
	return mdate
}

func (mdate Date) IsSameDate(tdate Date) bool {
	return tdate.GetYear() == mdate.GetYear() &&
		tdate.GetMonth() == mdate.GetMonth() &&
		tdate.GetDay() == mdate.GetDay() &&
		tdate.GetHour() == mdate.GetHour() &&
		tdate.GetMinute() == mdate.GetMinute()
}

func (mdate Date) MoreThan(tdate Date) bool {
	if mdate.Year > tdate.GetYear() {
		return true
	}
	if mdate.Year < tdate.GetYear() {
		return false
	}
	if mdate.Month > tdate.GetMonth() {
		return true
	}
	if mdate.Month < tdate.GetMonth() {
		return false
	}
	if mdate.Day > tdate.GetDay() {
		return true
	}
	if mdate.Day < tdate.GetDay() {
		return false
	}
	if mdate.Hour > tdate.GetHour() {
		return true
	}
	if mdate.Hour < tdate.GetHour() {
		return false
	}
	if mdate.Minute > tdate.GetMinute() {
		return true
	}
	if mdate.Minute < tdate.GetMinute() {
		return false
	}
	return false
}

func (mdate Date) LessThan(tdate Date) bool {
	if mdate.IsSameDate(tdate) == false && !mdate.MoreThan(tdate) == false {
		return false
	}
	return true
}

func (mdate Date) MoreOrEqual(tdate Date) bool {
	return mdate.IsSameDate(tdate) || mdate.MoreThan(tdate)
}

func (mdate Date) LessOrEqual(tdate Date) bool {
	return !mdate.MoreThan(tdate)
}
