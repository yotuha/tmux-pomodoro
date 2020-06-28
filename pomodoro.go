package main

import (
    "fmt"
    "os"
    "time"

    "github.com/yotuha/pomodoro/util"
)

func main() {
    writePid()
    pid := readPid()
    fmt.Println(pid)

    t := time.Now()
    writeTime(t)
    rt := readTime()
    fmt.Println(rt)
}

func readTime() time.Time {
    t := util.ReadTime()
    return t
}
func writeTime(t time.Time) {
    util.WriteTime(t)
}
func readPid() int {
    pid := util.ReadPid()
    return pid
}
func writePid() {
    pid := os.Getpid()
    util.WritePid(pid)
}
