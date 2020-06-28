package util

import (
    "io/ioutil"
    "strconv"
    "strings"
    "time"
    "os"
    "log"
)

func ReadPid() int {
	bytes, err := ioutil.ReadFile(pidFilePath())
	if err != nil {
		log.Fatal(err)
	}
	str := string(bytes[:])
    pid, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}
    return pid
}
func WritePid(pid int) {
    str := strconv.Itoa(pid)
    bytes := []byte(str)
	err := ioutil.WriteFile(pidFilePath(), bytes, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
func pidFilePath() string {
    return os.Getenv("HOME") + "/.pomodoro.pid"
}

func ReadTime() time.Time {
    bytes, err := ioutil.ReadFile(timeFilePath())
    if err != nil {
        log.Fatal()
    }
	str := string(bytes[:])
	str = strings.TrimSpace(str)

	t, err := time.Parse(TimeFormat, str)
    if err != nil {
        log.Fatal()
    }
    return t;
}
func WriteTime(t time.Time) {
    bytes := []byte(t.Format(TimeFormat))
	err := ioutil.WriteFile(timeFilePath(), bytes, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
func timeFilePath() string {
    return os.Getenv("HOME") + "/.pomodoro"
}
