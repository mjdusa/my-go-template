package runner

import (
	"flag"
	"fmt"
	"os"
	"runtime/debug"

	"github.com/mjdusa/my-go-template/internal/version"
)

func GetParameters() (bool, bool, bool) {
	flagSet := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)

	flagSet.SetOutput(os.Stderr)

	var dbg bool
	var verbose bool
	var version bool

	// add flags
	flagSet.BoolVar(&dbg, "debug", false, "Log Debug")
	flagSet.BoolVar(&verbose, "verbose", false, "Show Verbose Logging")
	flagSet.BoolVar(&version, "version", false, "Show Version")

	// Parse the flags
	if err := flagSet.Parse(os.Args[1:]); err != nil {
		flagSet.Usage()
		panic(err)
	}

	if dbg {
		buildInfo, ok := debug.ReadBuildInfo()
		if ok {
			fmt.Println(buildInfo.String())
		}
	}

	return version, dbg, verbose
}

func Run() int {
	// ctx := context.Background()
	versionFlag, debugFlag, verboseFlag := GetParameters()

	fmt.Println("Debug Flag: ", debugFlag)
	fmt.Println("Verbose Flag: ", verboseFlag)
	fmt.Println("Version Flag: ", versionFlag)

	if versionFlag {
		fmt.Println(version.GetVersion())
	}

	// go something usefull here

	return 0
}
