package entity

import (
	"Agenda/loghelper"
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

// UserFilter .
type UserFilter func(*User) bool

// MeetingFilter .
type MeetingFilter func(*Meeting) bool

var userinfoPath = "data/userinfo"
var metinfoPath = "data/meetinginfo"
var curUserPath = "data/curUser.txt"

var curUserName *string

var dirty bool

var uData []User
var mData []Meeting

var errLog *log.Logger
var loginLog *log.Logger

func init() {
	errLog = loghelper.Error
	loginLog = loghelper.Login
	dirty = false
	if err := readFromFile(); err != nil {
		fmt.Println("readFromFile fail:", err)
		errLog.Println("readFromFile fail: ", err)
	}
}

// Logout .
func Logout() error {
	return Sync()
}

// Sync .
func Sync() error {
	if err := writeToFile(); err != nil {
		fmt.Println("write to file fail,", err)
		errLog.Println("write to file fail,", err)
		return err
	}
	return nil
}

func readFromFile() error {
	var e []error
	str, err1 := readString(curUserPath)
	//str, err1 := readString("data/curUser.txt")
	if err1 != nil {
		e = append(e, err1)
	}
	curUserName = str
	if err := readUser(); err != nil {
		e = append(e, err)
	}
	if err := readMet(); err != nil {
		e = append(e, err)
	}
	if len(e) == 0 {
		return nil
	}
	er := e[0]
	for i := 1; i < len(e); i++ {
		er = errors.New(er.Error() + e[i].Error())
	}
	return er
}

func writeToFile() error {
	var e []error
	if err := writeString(curUserPath, curUserName); err != nil {
		e = append(e, err)
	}
	if dirty {
		if err := writeUser(); err != nil {
			e = append(e, err)
		}
		if err := writeMet(); err != nil {
			e = append(e, err)
		}
	}
	if len(e) == 0 {
		return nil
	}
	er := e[0]
	for i := 1; i < len(e); i++ {
		er = errors.New(er.Error() + e[i].Error())
	}
	return er
}

func readUser() error {
	file, err := os.Open(userinfoPath)
	if err != nil {
		fmt.Println("Open File Fail:", userinfoPath, err)
		errLog.Println("Open File Fail:", userinfoPath, err)
		return err
	}
	defer file.Close()

	dec := json.NewDecoder(file)
	switch err := dec.Decode(&uData); err {
	case nil, io.EOF:
		return nil
	default:
		fmt.Println("Decode User Fail:", err)
		errLog.Println("Decode User Fail:", err)
		return err
	}
}

func readMet() error {
	file, err := os.Open(metinfoPath)
	if err != nil {
		fmt.Println("Open File Fail:", metinfoPath, err)
		errLog.Println("Open File Fail:", metinfoPath, err)
		return err
	}
	defer file.Close()

	dec := json.NewDecoder(file)
	switch err := dec.Decode(&mData); err {
	case nil, io.EOF:
		return nil
	default:
		fmt.Println("Decode Met Fail:", err)
		errLog.Println("Decode Met Fail:", err)
		return err
	}
}

func writeUser() error {
	file, err := os.Create(userinfoPath)
	if err != nil {
		fmt.Println("writeUser Fail:", err)
		errLog.Println("writeUser Fail:", err)
		return err
	}
	defer file.Close()

	enc := json.NewEncoder(file)

	if err := enc.Encode(&uData); err != nil {
		fmt.Println("Encode User Fail:", err)
		errLog.Println("Encode User Fail:", err)
		return err
	}
	return nil
}

func writeMet() error {
	file, err := os.Create(metinfoPath)
	if err != nil {
		// fmt.Println("writeMet Fail:", err)
		errLog.Println("writeMet Fail:", err)
		return err
	}
	defer file.Close()

	enc := json.NewEncoder(file)

	if err := enc.Encode(&mData); err != nil {
		// fmt.Println("Encode Met Fail:", err)
		errLog.Println("Encode Met Fail:", err)
		return err
	}
	return nil
}

func writeString(path string, data *string) error {
	file, err := os.Create(path)
	if err != nil {
		// println("Log Error: Create file error:", path)
		errLog.Println("Log Error: Create file error:", path)
		return err
	}
	defer file.Close()

	if data == nil {
		//fmt.Println("curUser:", err)
		errLog.Println("curUser:", err)
		return err
	}
	writer := bufio.NewWriter(file)
	if _, err := writer.WriteString(*data); err != nil {
		// fmt.Println("Log: Write File fail,", err)
		errLog.Println("Log: Write File fail,", err)
		return err
	}
	if err := writer.Flush(); err != nil { //Flush 是将缓存中的所有数据都写进wirter里面
		// fmt.Println("Log: Flush File Fail,", err)
		errLog.Println("Log: Flush File Fail,", err)
		return err
	}
	return nil
}

func readString(path string) (*string, error) {
	file, err := os.Open(path)
	if err != nil {
		//fmt.Println("Log: Open file fail,", err)
		errLog.Println("Log: Open file fail,", err)
		return nil, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	str, err := reader.ReadString('\n') //识别换行符结束读取
	if err != nil && err != io.EOF {
		// fmt.Println("Log: Read file fail,", path, err)
		errLog.Println("Log: Read file fail,", path, err)
		return nil, err
	}
	return &str, nil
}

// CreateUser : create a user
// @param a user object
func CreateUser(v *User) {
	uData = append(uData, *v)
	dirty = true
	Sync()
	//loginLog.Println("Create ", &v, " sucessfully")
}

// QueryUser : query users
// @param a lambda function as the filter
// @return a list of fitted users
func QueryUser(filter UserFilter) []User {
	var user []User
	for _, v := range uData {
		if filter(&v) {
			user = append(user, v)
		}
	}
	return user
}

// UpdateUser : update users
// @param a lambda function as the filter
// @param a lambda function as the method to update the user
// @return the number of updated users
func UpdateUser(filter UserFilter, switcher func(*User)) int {
	count := 0
	for i := 0; i < len(uData); i++ {
		if v := &uData[i]; filter(v) {
			switcher(v)
			count++
		}
	}
	if count > 0 {
		dirty = true
	}
	return count
}

// DeleteUser : delete users
// @param a lambda function as the filter
// @return the number of deleted users
func DeleteUser(filter UserFilter) int {
	count := 0
	for i, v := range uData {
		if filter(&v) {
			// uData[i] = uData[len(uData)-1]
			// uData = uData[:len(uData)-1]
			uData = append(uData[:i], uData[i+1:]...)
			count++
		}
	}
	if count > 0 {
		dirty = true
	}
	return count
}

// CreateMeeting : create a meeting
// @param a meeting object
func CreateMeeting(v *Meeting) {
	//fmt.Println(v)
	mData = append(mData, *v)
	dirty = true
	Sync()
}

// QueryMeeting : query meetings
// @param a lambda function as the filter
// @return a list of fitted meetings
func QueryMeeting(filter MeetingFilter) []Meeting {
	var met []Meeting
	for _, v := range mData {
		if filter(&v) {
			met = append(met, v)
		}
	}
	return met
}

// UpdateMeeting : update meetings
// @param a lambda function as the filter
// @param a lambda function as the method to update the meeting
// @return the number of updated meetings
func UpdateMeeting(filter MeetingFilter, switcher func(*Meeting)) int {
	count := 0
	for i := 0; i < len(mData); i++ {
		if v := &mData[i]; filter(v) {
			switcher(v)
			count++
		}
	}
	if count > 0 {
		dirty = true
	}
	Sync()
	return count
}

// DeleteMeeting : delete meetings
// @param a lambda function as the filter
// @return the number of deleted meetings
func DeleteMeeting(filter MeetingFilter) int {
	count := 0
	for i, v := range mData {
		if filter(&v) {
			// mData[i] = mData[len(mData)-1]
			// mData = mData[:len(mData)-1]
			mData = append(mData[:i], mData[i+1:]...)
			count++
		}
	}
	if count > 0 {
		dirty = true
	}
	Sync()
	return count
}

// SetCurUser .
func SetCurUser(u *User) {
	if curUserName == nil {
		p := u.Name
		curUserName = &p
	} else {
		*curUserName = u.Name
	}

	//写入文件
	Sync()
}

// GetCurUser .
func GetCurUser() (User, error) {
	if curUserName == nil {
		return User{}, errors.New("Current user does not exist")
	}
	for _, v := range uData {
		if v.Name == *curUserName {
			return v, nil
		}
	}
	return User{}, errors.New("Current user does not exist")
}
