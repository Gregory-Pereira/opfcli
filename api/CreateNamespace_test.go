package api

import "github.com/stretchr/testify/require"

func (suite *apiTestSuite) TestCreateNamespace() {
	assert := require.New(suite.T())

	err := suite.api.CreateNamespace(
		"testproject",
		"testgroup",
		"test description",
		"",
		false,
		false,
	)
	assert.Nil(err)

	expectedPaths := []string{
		"cluster-scope/base/core/namespaces/testproject/kustomization.yaml",
		"cluster-scope/base/core/namespaces/testproject/namespace.yaml",
	}

	compareWithExpected(assert, "testdata/CreateNamespace", suite.dir, expectedPaths)
}

func (suite *apiTestSuite) TestCreateNamespaceExistsOk() {
	assert := require.New(suite.T())

	err := suite.api.CreateNamespace(
		"testproject",
		"testgroup",
		"test description",
		"",
		false,
		false,
	)
	assert.Nil(err)

	err = suite.api.CreateNamespace(
		"testproject",
		"testgroup",
		"test description",
		"",
		false,
		true,
	)
	assert.Nil(err)

	expectedPaths := []string{
		"cluster-scope/base/core/namespaces/testproject/kustomization.yaml",
		"cluster-scope/base/core/namespaces/testproject/namespace.yaml",
	}

	compareWithExpected(assert, "testdata/CreateNamespace", suite.dir, expectedPaths)
}

func (suite *apiTestSuite) TestCreateNamespaceExistsNotOk() {
	assert := require.New(suite.T())

	err := suite.api.CreateNamespace(
		"testproject",
		"testgroup",
		"test description",
		"",
		false,
		false,
	)
	assert.Nil(err)

	err = suite.api.CreateNamespace(
		"testproject",
		"testgroup",
		"test description",
		"",
		false,
		false,
	)
	assert.EqualError(err, "namespace testproject already exists")

	expectedPaths := []string{
		"cluster-scope/base/core/namespaces/testproject/kustomization.yaml",
		"cluster-scope/base/core/namespaces/testproject/namespace.yaml",
	}

	compareWithExpected(assert, "testdata/CreateNamespace", suite.dir, expectedPaths)
}

func (suite *apiTestSuite) TestCreateNamespaceQuota() {
	assert := require.New(suite.T())

	// Should fail if quota doesn't exist
	err := suite.api.CreateNamespace(
		"testproject",
		"testgroup",
		"test description",
		"testquota",
		false,
		false,
	)
	assert.Nil(err)

	expectedPaths := []string{
		"cluster-scope/base/core/namespaces/testproject/kustomization.yaml",
	}

	compareWithExpected(assert, "testdata/CreateNamespaceQuota", suite.dir, expectedPaths)
}

func (suite *apiTestSuite) TestCreateNamespaceNoLimitrange() {
	assert := require.New(suite.T())

	// Should fail if quota doesn't exist
	err := suite.api.CreateNamespace(
		"testproject",
		"testgroup",
		"test description",
		"",
		true,
		false,
	)
	assert.Nil(err)

	expectedPaths := []string{
		"cluster-scope/base/core/namespaces/testproject/kustomization.yaml",
	}

	compareWithExpected(assert, "testdata/CreateNamespaceNoLimitrange", suite.dir, expectedPaths)
}
