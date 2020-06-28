package main

import (
	"flag"
    "fmt"
	"log"
	"os"
	"time"

	"github.com/sevlyar/go-daemon"
	"github.com/yotuha/tmux-pomodoro/config"
	"github.com/yotuha/tmux-pomodoro/pstate"
)

const COMMAND_START = "start"
const COMMAND_STOP = "stop"

func main() {
    start := flag.NewFlagSet("start", flag.ExitOnError)
	cntxt := &daemon.Context{
		PidFileName: os.Getenv("HOME") + "/.pomodoro.pid",
		PidFilePerm: 0644,
		LogFileName: os.Getenv("HOME") + "/.pomodoro.log",
		LogFilePerm: 0640,
	}

    if len(os.Args) == 1 {
        printHelp()
        return
    }
    switch os.Args[1] {
    case "start":
        con := config.NewConfig(start)

        d, err := cntxt.Reborn()
        if err != nil {
            log.Fatalln(err)
        }
        if d != nil {
            return
        }
        defer cntxt.Release()

        run(con)
    case "stop":
		d, err := cntxt.Search()
		if err != nil {
			log.Fatalf("Error: %s", err.Error())
		}
        d.Kill()
        os.Remove(os.Getenv("HOME") + "/.pomodoro.pid")
        pstate.Clear()
    case "help":
        printHelp()
    default:
        fmt.Printf("%q is not valid command.\n", os.Args[1])
        printHelp()
    }
}

func run(con *config.Config) {
    ps := pstate.NewPomodoroState(con)
	for range time.Tick(time.Second) {
        ps.UpdateState()
        ps.WriteState()
        if ps.IsDoneAllSet() {
            pstate.Clear()
            return
        }
	}
}

func printHelp() {
    fmt.Println("usage: tmux-pomodoro <command> [<args>]")
    fmt.Println("command: start   Start pomodoro set.")
    fmt.Println("         stop    Stop pomodoro set.")
    fmt.Println("         help    Print help.")
    fmt.Println("args: Only start can be specified as an args.")
    fmt.Println("      set   Number of set.")
    fmt.Println("      work  Number of work time[sec].")
    fmt.Println("      rest  Number of rest time[sec].")
    fmt.Println("The configuration file can for more detailed control.")
    fmt.Println("More details at github.com/yotuha/tmux-pomodoro.")
}
