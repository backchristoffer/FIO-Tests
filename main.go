package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
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

func check_fio() {
	cmd := exec.Command(
		"fio",
		"--version",
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	output_str := string(output)

	if strings.Contains(output_str, "fio-") {
		fmt.Println("FIO is installed")
	} else {
		fmt.Println("FIO is not installed")
	}
}

func iops_and_bw_for_rand_reads() {
	cmd := exec.Command(
		"fio",
		"--filename=iops_and_bw_for_rand_reads.file", // name
		"--size=10GB",                        // size of test file
		"--direct=1",                         // use direct I/O mode (bypassing kernel cache)
		"--rw=randread",                      // random read operations during the test
		"--bs=4k",                            // block size of each I/O request (4 KB in this case)
		"--ioengine=libaio",                  // I/O engine used
		"--iodepth=256",                      // number of I/Os to keep in flight (queue depth)
		"--numjobs=8",                        // based on number of cpus
		"--time_based",                       // use time-based instead of size-based I/O
		"--group_reporting",                  // report the results for all jobs as a group
		"--name=iops_and_bw_for_rands_reads", // name
		"--runtime=10",                       // runtime in sec
		"--eta-newline=1",                    // print ETA on newline
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		panic(err)
	}

}

func iops_and_bw_for_rand_reads_and_writes() {
	cmd := exec.Command(
		"fio",
		"--filename=iops_and_bw_for_rand_reads_and_writes.file",
		"--size=500GB",
		"--direct=1",
		"--rw=randread",
		"--bs=4k",
		"--ioengine=libaio",
		"--iodepth=256",
		"--numjobs=8",
		"--time_based",
		"--group_reporting",
		"--name=iops_and_bw_for_rand_reads_and_writes",
		"--runtime=10",
		"--eta-newline=1",
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		panic(err)
	}

}

func main() {
	interval := 10 * time.Second

	for {
		iops_and_bw_for_rand_reads()
		time.Sleep(interval)
	}

}
