package version_test

import (
	"fmt"
	"os"
	"testing"

	put "github.com/mjdusa/my-go-template/internal/version" // put - package under test
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// Setup Suite.
type VersionSuite struct {
	suite.Suite
}

func TestVersionSuite(t *testing.T) {
	versionSuite := VersionSuite{} //nolint:exhaustivestruct  // This is then normal way to instantiate a suite
	suite.Run(t, &versionSuite)
}

type testGetVersion struct {
	Description string
	AppVersion  string
	Branch      string
	BuildTime   string
	Commit      string
	GoVersion   string
}

func createGetVersionTestData() []testGetVersion {
	tests := []testGetVersion{
		{
			Description: "All are empty strings",
			AppVersion:  "",
			Branch:      "",
			BuildTime:   "",
			Commit:      "",
			GoVersion:   "",
		},
		{
			Description: "All have values",
			AppVersion:  "AppVersion",
			Branch:      "Branch",
			BuildTime:   "BuildTime",
			Commit:      "Commit",
			GoVersion:   "GoVersion",
		},
	}

	return tests
}

func (s *VersionSuite) TestGetVersion() {
	for _, tst := range createGetVersionTestData() {
		expected := fmt.Sprintf(
			"%s version: [%s]\n- Branch:     [%s]\n- Build Time: [%s]\n- Commit:     [%s]\n- Go Version: [%s]\n",
			os.Args[0], tst.AppVersion, tst.Branch, tst.BuildTime, tst.Commit, tst.GoVersion)

		put.AppVersion = tst.AppVersion
		put.Branch = tst.Branch
		put.BuildTime = tst.BuildTime
		put.Commit = tst.Commit
		put.GoVersion = tst.GoVersion

		actual := put.GetVersion()

		assert.Equal(s.T(), expected, actual, tst.Description+fmt.Sprintf(" expected '%s', actual '%s'", expected, actual))
	}
}
