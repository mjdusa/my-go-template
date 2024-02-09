package runner

import (
	"flag"
	"fmt"
	"os"
	"runtime/debug"

	"github.com/mjdusa/my-go-template/internal/version"
)

func GetParameters() (bool, bool) {
	flagSet := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)

	flagSet.SetOutput(os.Stderr)

	var dbg bool
	var verbose bool

	// add flags
	flagSet.BoolVar(&dbg, "debug", false, "Log Debug")
	flagSet.BoolVar(&verbose, "verbose", false, "Show Verbose Logging")

	// Parse the flags
	if err := flagSet.Parse(os.Args[1:]); err != nil {
		flagSet.Usage()
		panic(err)
	}

	if verbose {
		fmt.Println(version.GetVersion())
	}

	if dbg {
		buildInfo, ok := debug.ReadBuildInfo()
		if ok {
			fmt.Println(buildInfo.String())
		}
	}

	return dbg, verbose
}

func Run() int {
	// ctx := context.Background()
	debugFlag, verboseFlag := GetParameters()

	fmt.Println("Debug Flag: ", debugFlag)
	fmt.Println("Verbose Flag: ", verboseFlag)

	// go something usefull here

	return 0
}
