package runner_test

import (
	"flag"
	"fmt"
	"os"
	"testing"

	main "github.com/mjdusa/my-go-template/cmd/my-go-template"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// Setup Suite
type RunnerSuite struct {
	suite.Suite
}

func Test_RunnerSuite(t *testing.T) {
	suite.Run(t, &RunnerSuite{})
}

type TestGetParameters struct {
	Description     string
	AuthFlag        *string
	DebugFlag       bool
	VerboseFlag     bool
	ExpectedAuth    string
	ExpectedDebug   bool
	ExpectedVerbose bool
}

func Call_GetParameters(s *RunnerSuite) {
	os.Args = []string{"mainTest"}
	arg := fmt.Sprintf("-auth=%s", os.Getenv("GITHUB_AUTH"))
	os.Args = append(os.Args, arg)
	os.Args = append(os.Args, "-debug")
	os.Args = append(os.Args, "-verbose")

	actualAuth, actualDebug, actualVerbose := main.GetParameters()
	main.GetParameters()

	fmt.Println("inside")

	fmt.Fprintf(os.Stdout, "actualAuth=[%s]\n", actualAuth)
	fmt.Fprintf(os.Stdout, "actualDebug=[%t]\n", actualDebug)
	fmt.Fprintf(os.Stdout, "actualVerbose=[%t]\n", actualVerbose)
}

func (s *RunnerSuite) Test_GetParameters_AuthFlag_Empty() {
	os.Args = []string{"mainTest"}

	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	os.Args = append(os.Args, "-auth=")
	os.Args = append(os.Args, "-debug")
	os.Args = append(os.Args, "-verbose")

	main.PanicOnExit = true

	defer func() {
		if r := recover(); r == nil {
			s.T().Errorf("The code did not panic")
		} else {
			s.T().Logf("Recovered in %v", r)
		}
	}()

	main.GetParameters()

	assert.Fail(s.T(), "Test_GetParameters_AuthFlag_Empty expected Panic to fire")
}

func (s *RunnerSuite) Test_GetParameters_FlagParse() {
	os.Args = []string{"mainTest"}

	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	os.Args = append(os.Args, "-auth=''")
	os.Args = append(os.Args, "-debug")
	os.Args = append(os.Args, "-verbose")
	os.Args = append(os.Args, "-panic")

	main.PanicOnExit = true

	defer func() {
		if r := recover(); r == nil {
			s.T().Errorf("The code did not panic")
		} else {
			s.T().Logf("Recovered in %v", r)
		}
	}()

	main.GetParameters()

	assert.Fail(s.T(), "Test_GetParameters_FlagParse expected Panic to fire")
}
