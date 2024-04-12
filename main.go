package main

import (
	"fmt"
	"os/exec"
	"time"
)

type fioConfig struct {
	Filename       string // filename saved
	Size           string // size of test file
	Direct         bool   // use direct I/O mode (bypassing kernel cache)
	RW             string // read operations during the test. randwrite/randrw/write/rw/randtrim
	BS             string // block size of each I/O request
	IOEngine       string // I/O engine used
	IODepth        int    // number of I/Os to keep in flight (queue depth)
	NumJobs        int    // based on number of cpus
	TimeBased      bool   // use time-based instead of size-based I/O
	GroupReporting bool   // report the results for all jobs as a group
	Name           string // name
	Runtime        int    // runtime in sec
	ETANewline     int    // print ETA on newline
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func runFio(config fioConfig) (string, error) {
	args := []string{
		"--filename=" + config.Filename,
		"--size=" + config.Size,
		"--direct=" + fmt.Sprintf("%d", boolToInt(config.Direct)),
		"--rw=" + config.RW,
		"--bs=" + config.BS,
		"--ioengine=" + config.IOEngine,
		"--iodepth=" + fmt.Sprintf("%d", config.IODepth),
		"--numbjobs=" + fmt.Sprintf("%d", config.NumJobs),
		"--time_based=" + fmt.Sprintf("%d", boolToInt(config.TimeBased)),
		"--group_reporting=" + fmt.Sprintf("%d", boolToInt(config.GroupReporting)),
		"--name=" + config.Name,
		"--runtime=" + fmt.Sprintf("%d", config.Runtime),
		"--eta-newline=" + fmt.Sprintf("%d", config.ETANewline),
	}

	cmd := exec.Command("fio", args...)
	output, err := cmd.CombinedOutput()
	return string(output), err
}

func runTests(configs []fioConfig) {
	for _, config := range configs {
		output, err := runFio(config)
		if err != nil {
			fmt.Println("Error running test for %s: %v\n", config.Name, err)
			continue
		}
		fmt.Println("Output for %s:\n%s\n", config.Name, output)
	}
}

func main() {
	interval := 10 * time.Second
	for {
		runTests(fioConfigs)
		time.Sleep(interval)
	}
}
