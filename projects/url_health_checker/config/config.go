package config

import (
	"errors"
	"os"
	"strconv"
)

//*********************************************************

func Usage() error {
	return errors.New("\nURL Health Checker Usage:\n" +
		"\t--help:     print usage\n" +
		"\t--path:     path to URL list\n" +
		"\t--num-jobs: number of jobs\n" +
		"\t--timeout:  time waiting for http request before timeout\n")
}

//*********************************************************

type AppConfig struct {
	Path    string
	NumJobs uint32
	Timeout uint32
}

//*********************************************************

func (a *AppConfig) ProcessCmdLine(args []string) error {
	if len(args) == 1 && (args[0] == "--help" || args[0] == "-h") {
		return Usage()
	} else if len(args) < 4 {
		return errors.New("too few args for configuring URL health checker")
	}

	// Process the arguments over a ranged loop -- no need for correctness at the moment
	// Non-path or jobs args discarded
	for ii := range args {
		switch args[ii] {
		case "--path":
			a.Path = args[ii+1]
		case "--num-jobs":
			v, err := strconv.Atoi(args[ii+1])
			if err != nil {
				return err
			}
			a.NumJobs = uint32(v)
		case "--timeout":
			v, err := strconv.Atoi(args[ii+1])
			if err != nil {
				return err
			}
			a.Timeout = uint32(v)
		}
	}

	// Now sanity check
	_, err := os.Stat(a.Path) // Path must exist
	if err != nil {
		return err
	}

	if a.NumJobs < 1 { // Number of jobs must be positive non-zero
		return errors.New("no jobs requested for URL health check")
	}

	return nil
}

//*********************************************************
