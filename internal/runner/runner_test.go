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
	suite.Run(t, &RunnerSuite{})
}

type testGetParameters struct {
	Description     string
	DebugFlag       bool
	VerboseFlag     bool
	ExpectedDebug   bool
	ExpectedVerbose bool
}

func createGetParametersTestData() []testGetParameters {
	tests := []testGetParameters{
		{
			Description:     "All false",
			DebugFlag:       false,
			VerboseFlag:     false,
			ExpectedDebug:   false,
			ExpectedVerbose: false,
		},
		{
			Description:     "All true",
			DebugFlag:       true,
			VerboseFlag:     true,
			ExpectedDebug:   true,
			ExpectedVerbose: true,
		},
		{
			Description:     "Flip Flop",
			DebugFlag:       true,
			VerboseFlag:     false,
			ExpectedDebug:   true,
			ExpectedVerbose: false,
		},
		{
			Description:     "Flop Flip",
			DebugFlag:       false,
			VerboseFlag:     true,
			ExpectedDebug:   false,
			ExpectedVerbose: true,
		},
	}

	return tests
}

func (s *RunnerSuite) Test_GetParameters() {
	for _, test := range createGetParametersTestData() {
		os.Args = []string{putRunner}
		if test.DebugFlag {
			os.Args = append(os.Args, "-debug")
		}
		if test.VerboseFlag {
			os.Args = append(os.Args, "-verbose")
		}

		actualDebug, actualVerbose := put.GetParameters()

		assert.Equal(s.T(), test.ExpectedDebug, actualDebug, test.Description)
		assert.Equal(s.T(), test.ExpectedVerbose, actualVerbose, test.Description)
	}
}
