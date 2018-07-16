package exec

import (
	"bytes"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/oscarzhao/launch/logging"
)

func TestNewExec_Simple(t *testing.T) {
	execer, err := New("test", "echo", []string{`"Hello, world"`})
	require.Nil(t, err)
	require.NotNil(t, execer)
}

func TestNewExec_WithOptions(t *testing.T) {
	execer, err := New("test", "echo", []string{`"Hello, world"`},
		WithEnv([]string{"JAVA_HOME=D:/Java/jdk1.9"}),
		WithProcessLogger(os.Stdout),
		WithRuntimeLogger(logging.NewNoop()),
	)
	require.Nil(t, err)
	require.NotNil(t, execer)
}

func TestStart(t *testing.T) {
	var buff bytes.Buffer
	execer, err := New("test", "ls", []string{"./"},
		WithProcessLogger(&buff),
	)
	require.Nil(t, err)

	err = execer.Start()
	require.Nil(t, err)
	require.Contains(t, buff.String(), "impl_test.go")
	require.Contains(t, buff.String(), "iface.go")
	require.Contains(t, buff.String(), "impl.go")
}

func TestStart_Binary_NotFound(t *testing.T) {
	var buff bytes.Buffer
	execer, err := New("test", "wtf", []string{},
		WithProcessLogger(&buff),
		WithRuntimeLogger(logging.NewNoop()),
	)
	require.Nil(t, err)

	err = execer.Start()
	require.NotNil(t, err)
	require.Contains(t, ErrExecutableNotFound.Error(), err.Error())
}

// TestStart_Stop has race-conditions here, @todo
// Detail: Start() write `cmd.Process`, while Stop() read it
func TestStart_Stop(t *testing.T) {
	var buff bytes.Buffer
	var execer Execer
	var err error
	execer, err = New("test", "sleep", []string{"15s"},
		WithProcessLogger(&buff),
		WithRuntimeLogger(logging.NewNoop()),
	)
	require.Nil(t, err)

	go func() {
		startErr := execer.Start()
		require.Nil(t, startErr)
	}()

	time.Sleep(2 * time.Second)
	execer.Stop()
}
