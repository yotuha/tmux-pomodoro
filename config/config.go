package config

import (
	"flag"
	"io/ioutil"
	"log"
	"os"

	"github.com/go-yaml/yaml"
)

type Config struct {
    SetNum int
    WorkSec int
    RestSec int
    AfterWorkRunCmd []string
    AfterRestRunCmd []string
}

func NewConfig(flagSet *flag.FlagSet) *Config {
    config := new(Config)

    setNum := flagSet.Int("set", 1, "Number of set.")
    workSec := flagSet.Int("work", 1500, "Number of work time[sec]")
    restSec := flagSet.Int("rest", 300, "Number of rest time[sec]")
    flagSet.Parse(os.Args[2:])
    config.SetNum = *setNum
    config.WorkSec = *workSec
    config.RestSec = *restSec

    m := readConfigFile()
    if _, ok := m["set"]; ok && !isFlagPassed(flagSet, "set") {
        config.SetNum = m["set"].(int)
    }
    if _, ok := m["work"]; ok && !isFlagPassed(flagSet, "work") {
        config.WorkSec = m["work"].(int)
    }
    if _, ok := m["rest"]; ok && !isFlagPassed(flagSet, "rest") {
        config.RestSec = m["rest"].(int)
    }
    if _, ok := m["afterWorkRunCmd"]; ok {
        arr := convertInterfaceArr2StrArr(m["afterWorkRunCmd"].([]interface{}))
        config.AfterWorkRunCmd = arr
    }
    if _, ok := m["afterRestRunCmd"]; ok {
        arr := convertInterfaceArr2StrArr(m["afterRestRunCmd"].([]interface{}))
        config.AfterRestRunCmd = arr
    }

    return config
}

func readConfigFile() map[interface{}]interface{} {
    m := make(map[interface{}]interface{})
    bytes, err := ioutil.ReadFile(os.Getenv("HOME") + "/.pomodoro.yaml")
    if err != nil {
        return m
    }
    err = yaml.Unmarshal(bytes, &m)
    if err != nil {
        log.Fatalln(err)
    }
    return m
}

func isFlagPassed(flagSet *flag.FlagSet, name string) bool {
    found := false
    flagSet.Visit(func(f *flag.Flag) {
        if f.Name == name {
            found = true
        }
    })
    return found
}

func convertInterfaceArr2StrArr(target []interface{}) []string {
    arr := make([]string, len(target))
    for i, v := range target {
        arr[i] = v.(string)
    }
    return arr;
}
