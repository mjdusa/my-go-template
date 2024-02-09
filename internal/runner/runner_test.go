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

// Setup Suite
type RunnerSuite struct {
	suite.Suite
}

func Test_RunnerSuite(t *testing.T) {
	suite.Run(t, &RunnerSuite{})
}

type testGetParameters struct {
	Description     string
	DebugFlag       bool
	VerboseFlag     bool
	ExpectedDebug   bool
	ExpectedVerbose bool
}

func get_testGetParameters_data() []testGetParameters {
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
	put.PanicOnExit = true

	defer func() {
		if r := recover(); r == nil {
			s.T().Logf("The code did not panic")
		} else {
			s.T().Errorf("Recovered in %v", r)
		}
	}()

	for _, test := range get_testGetParameters_data() {
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
