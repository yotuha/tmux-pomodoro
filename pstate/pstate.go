package pstate

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/yotuha/tmux-pomodoro/config"
)

const MODE_WORK = 0
const MODE_REST = 1

type PomodoroState struct {
    config *config.Config
    workDoneNum int
    setDoneNum int
    mode int
    endTime time.Time
}

func NewPomodoroState(con *config.Config) *PomodoroState {
    ps := new(PomodoroState)
    ps.config = con
    ps.workDoneNum = 0
    ps.setDoneNum = 0
    ps.mode = MODE_WORK
    ps.endTime = time.Now().Add(time.Duration(con.WorkSec) * time.Second)
    return ps
}

func (ps PomodoroState) WriteState() {
    remain := -time.Since(ps.endTime)
    stateStr := ""
    if ps.mode == MODE_WORK {
        stateStr += "🍅 "
    } else if ps.mode == MODE_REST {
        stateStr += "🍫 "
    }
    if remain < 0 {
        stateStr += "00:00"
    } else {
        stateStr += remainTime2Str(remain)
    }
    stateStr += fmt.Sprintf(" [%d/%d]", ps.workDoneNum, ps.config.SetNum)
    writeState(stateStr)
}

func (ps *PomodoroState) UpdateState() {
    remain := -time.Since(ps.endTime)
    if remain > 0 {
        return
    }
    if ps.mode == MODE_REST {
        ps.setDoneNum++
        ps.mode = MODE_WORK
        addTime := time.Duration(ps.config.WorkSec) * time.Second
        ps.endTime = time.Now().Add(addTime)
        runCmds(ps.config.AfterRestRunCmd)
    } else {
        ps.workDoneNum++
        ps.mode = MODE_REST
        addTime := time.Duration(ps.config.RestSec) * time.Second
        ps.endTime = time.Now().Add(addTime)
        runCmds(ps.config.AfterWorkRunCmd)
    }
}

func (ps PomodoroState) IsDoneAllSet() bool {
    return ps.config.SetNum <= ps.setDoneNum
}

func Clear() {
    writeState("")
}

func runCmds(cmds []string) {
    for _, cmd := range cmds {
        err := exec.Command("zsh", "-c", cmd).Run()
        if err != nil {
            log.Println(err)
        }
    }
}

func writeState(str string) {
    path := os.Getenv("HOME") + "/.pomodoro"
    bytes := []byte("    " + str + "    \r")
    err := ioutil.WriteFile(path, bytes, 0644)
    if err != nil {
        log.Fatalln(err)
    }
}

func remainTime2Str(td time.Duration) string {
	min := int(td.Minutes())
	sec := int(td.Seconds()) % 60
	return fmt.Sprintf("%02d:%02d", min, sec)
}
