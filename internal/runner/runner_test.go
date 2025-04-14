package runner_test

import (
	"os"
	"testing"

	put "github.com/mjdusa/my-go-template/internal/runner" // put - package under test
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

const (
	putRunner = "put_runner" // put - package under test
)

// Setup Suite.
type RunnerSuite struct {
	suite.Suite
}

func TestRunnerSuite(t *testing.T) {
	runnerSuite := RunnerSuite{} //nolint:exhaustruct  // This is then normal way to instantiate a suite
	suite.Run(t, &runnerSuite)
}

type testGetParameters struct {
	Description     string
	VersionFlag     bool
	DebugFlag       bool
	VerboseFlag     bool
	ExpectedVersion bool
	ExpectedDebug   bool
	ExpectedVerbose bool
}

func createGetParametersTestData() []testGetParameters {
	tests := []testGetParameters{
		{
			Description:     "All false",
			VersionFlag:     false,
			DebugFlag:       false,
			VerboseFlag:     false,
			ExpectedVersion: false,
			ExpectedDebug:   false,
			ExpectedVerbose: false,
		},
		{
			Description:     "All true",
			VersionFlag:     true,
			DebugFlag:       true,
			VerboseFlag:     true,
			ExpectedVersion: true,
			ExpectedDebug:   true,
			ExpectedVerbose: true,
		},
		{
			Description:     "Flip Flop",
			VersionFlag:     false,
			DebugFlag:       true,
			VerboseFlag:     false,
			ExpectedVersion: false,
			ExpectedDebug:   true,
			ExpectedVerbose: false,
		},
		{
			Description:     "Flop Flip",
			VersionFlag:     true,
			DebugFlag:       false,
			VerboseFlag:     true,
			ExpectedVersion: true,
			ExpectedDebug:   false,
			ExpectedVerbose: true,
		},
	}

	return tests
}

func (s *RunnerSuite) Test_GetParameters() {
	for _, test := range createGetParametersTestData() {
		os.Args = []string{putRunner}
		if test.VersionFlag {
			os.Args = append(os.Args, "-version")
		}
		if test.DebugFlag {
			os.Args = append(os.Args, "-debug")
		}
		if test.VerboseFlag {
			os.Args = append(os.Args, "-verbose")
		}

		actualVersion, actualDebug, actualVerbose := put.GetParameters()

		assert.Equal(s.T(), test.ExpectedVersion, actualVersion, test.Description)
		assert.Equal(s.T(), test.ExpectedDebug, actualDebug, test.Description)
		assert.Equal(s.T(), test.ExpectedVerbose, actualVerbose, test.Description)
	}
}
