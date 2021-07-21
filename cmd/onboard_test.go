package cmd

import (
	"os"
	"path/filepath"

	"github.com/operate-first/opfcli/constants"
	"github.com/stretchr/testify/require"
)

func (suite *commandTestSuite) TestOnboardBasic() {
	assert := require.New(suite.T())

	// Should fail with too few args
	cmd := NewCmdOnboard(suite.api)
	cmd.SetArgs([]string{})
	err := cmd.Execute()
	assert.EqualError(err, "accepts 5 arg(s), received 0")

	// ---

	// Should fail with too many args
	cmd = NewCmdOnboard(suite.api)
	cmd.SetArgs([]string{"arg1", "arg2", "arg3"})
	err = cmd.Execute()
	assert.EqualError(err, "accepts 5 arg(s), received 3")

	// ---

	// Should fail with unknown option
	cmd = NewCmdOnboard(suite.api)
	cmd.SetArgs([]string{"--failure", "arg1", "arg2", "arg3", "arg4", "arg5"})
	err = cmd.Execute()
	assert.EqualError(err, "unknown flag: --failure")

	// ---

	// Should succeed
	cmd = NewCmdOnboard(suite.api)
	cmd.SetArgs([]string{"arg1", "arg2", "arg3", "arg4", "arg5"})
	err = cmd.Execute()
	assert.Nil(err)

	// ---

	// Should succeed if group already exists
	cmd = NewCmdOnboard(suite.api)
	cmd.SetArgs([]string{"arg1", "arg2", "arg3", "arg4", "arg5"})
	err = cmd.Execute()
	assert.Nil(err)
}

func (suite *commandTestSuite) TestOnboardQuota() {
	assert := require.New(suite.T())

	// Should fail because quota does not exist
	cmd := NewCmdOnboard(suite.api)
	cmd.SetArgs([]string{"-q", "testquota", "arg1", "arg2", "arg3", "arg4", "arg5"})
	err := cmd.Execute()
	assert.EqualError(err, "quota testquota does not exist")

	// ---

	err = os.MkdirAll(filepath.Join(
		suite.api.RepoDirectory,
		suite.api.AppName,
		constants.ComponentPath,
		"resourcequotas",
		"testquota",
	), 0755)
	assert.Nil(err)

	// Should succeed
	cmd = NewCmdOnboard(suite.api)
	cmd.SetArgs([]string{"-q", "testquota", "arg1", "arg2", "arg3", "arg4", "arg5"})
	err = cmd.Execute()
	assert.Nil(err)
}
