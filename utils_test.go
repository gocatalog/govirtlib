package govirtlib

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"libvirt.org/go/libvirt"
)

func TestConvertLibvirtVersion(t *testing.T) {
	hypervisorVersion := convertLibvirtVersion(8000000)
	assert.Equal(t, hypervisorVersion, "8.8000.0", "they should be equal")
}

type StateToStatusTestSuite struct {
	suite.Suite
}

// Make sure that VariableThatShouldStartAtFive is set to five
// before each test
// func (suite *StateToStatusTestSuite) SetupTest() {
//     suite.VariableThatShouldStartAtFive = 5
// }

// All methods that begin with "Test" are run as tests within a
// suite.
func (suite *StateToStatusTestSuite) TestNoState() {
	assert.Equal(suite.T(), stateToStatus(libvirt.DOMAIN_NOSTATE), nostate)
}

func (suite *StateToStatusTestSuite) TestRunning() {
	assert.Equal(suite.T(), stateToStatus(libvirt.DOMAIN_RUNNING), running)
}

func (suite *StateToStatusTestSuite) TestNoBlocked() {
	assert.Equal(suite.T(), stateToStatus(libvirt.DOMAIN_BLOCKED), blocked)
}

func (suite *StateToStatusTestSuite) TestPaused() {
	assert.Equal(suite.T(), stateToStatus(libvirt.DOMAIN_PAUSED), paused)
}

func (suite *StateToStatusTestSuite) TestNoShutdown() {
	assert.Equal(suite.T(), stateToStatus(libvirt.DOMAIN_SHUTDOWN), shutdown)
}

func (suite *StateToStatusTestSuite) TestNoCrash() {
	assert.Equal(suite.T(), stateToStatus(libvirt.DOMAIN_CRASHED), crashed)
}

func (suite *StateToStatusTestSuite) TestNoPms() {
	assert.Equal(suite.T(), stateToStatus(libvirt.DOMAIN_PMSUSPENDED), pmSuspended)
}

func (suite *StateToStatusTestSuite) TestShutOff() {
	assert.Equal(suite.T(), stateToStatus(libvirt.DOMAIN_SHUTOFF), shutOff)
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestStateToStatusTestSuite(t *testing.T) {
	suite.Run(t, new(StateToStatusTestSuite))
}

// func TestStateToStatusNoState(t *testing.T) {
// 	hypervisorVersion := convertLibvirtVersion(8000000)
// 	assert.Equal(t, hypervisorVersion, "8.8000.0", "they should be equal")
// }
