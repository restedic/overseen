package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

type reportInfo struct {
	duration time.Duration
	info     []string
}

var reportMap = map[string]reportInfo{
	"err": reportInfo{
		duration: time.Duration(0)
		info: []string{}
	}
}

func init() {
	os.OpenFile("/var/lock/overseen")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			// handle ctrl-c/ctrl-z exits
		}
	}()
}

func main() {
	switch os.Args[1] {
	case "start":
		p := ""
		cmd := exec.Command("xdotool", "getwindowfocus", "getwindowname")
		t := time.Now()
		for true {
			out, err := cmd.Output()
			if err != nil {
				append(reportMap["err"].info, err.Error())
			}
			np := string(out)
			if np != p {
				if val, ok := reportMap[p]; ok {
					reportMap[p].duration.Add(time.Now().Sub(t))
				} else {
					reportMap[p].duration = time.Now().Sub(t)
				}
				t := time.Now()
			}

		}
	case "stop":
	default:
	}
}

func reportMapString() (out string) {
	invMap := map[int64]string{}
	for k, v := range reportMap {
		invMap[int64(v.duration)] = k
	}
	intkeys := int64{}
	for k := range invMap {
		intkeys = append(keys, int64(k))
	}
	nmap := map[string]string{}
	for k, v := range reportMap {
		if v == "err" {
			continue
		} else {
			nmap[k] = fmt.Sprintf("%.0fh%.0fm%.0fs", v.duration.Hours(), v.duration.Minutes(), v.duration.Seconds())
		}
	}
	var maxLenKey int
    for k, _ := range nmap {
        if len(k) > maxLenKey {
            maxLenKey = len(k)
        }
    }

	out = ""
	for i := range intkeys {
        out = out + fmt.Sprintf("%*s: %s\n", maxLenKey, invMap[i], nmap[invMap[i]])
    }
}
