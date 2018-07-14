package exec

import (
	"io"
	"os"
	osexec "os/exec"
	"syscall"
	"time"

	"github.com/oscarzhao/launcher/logging"
	// k8sexec "k8s.io/utils/exec"
)

// defaultImpl contain all information you needs to run a command
type defaultImpl struct {
	name       string
	binaryPath string
	args       []string
	envs       []string // in the format `key=val`
	workingDir string
	isDaemon   bool
	autoStart  bool

	processLogger io.Writer
	logger        logging.Logger
	cmd           *osexec.Cmd
}

func newDefaultImpl(name, binaryPath string, args []string) *defaultImpl {
	return &defaultImpl{
		name:       name,
		binaryPath: binaryPath,
		args:       args,
		isDaemon:   true,
		autoStart:  false,

		processLogger: os.Stdout,
		logger:        logging.NewStdoutLogger(logging.LevelInfo),
	}
}

// Option is used to modify extra configs
type Option func(*defaultImpl)

// WithEnv set environment variables for binary to start
func WithEnv(envs []string) Option {
	return func(options *defaultImpl) {
		options.envs = append(options.envs, envs...)
	}
}

// WithWorkDir set working directory
func WithWorkDir(workDir string) Option {
	return func(options *defaultImpl) {
		options.workingDir = workDir
	}
}

// WithRuntimeLogger set the logger to store process manager's logs
func WithRuntimeLogger(logger logging.Logger) Option {
	return func(options *defaultImpl) {
		options.logger = logger
	}
}

// WithProcessLogger set the place the sub-process will print log to
func WithProcessLogger(writer io.Writer) Option {
	return func(options *defaultImpl) {
		options.processLogger = writer
	}
}

// New creates a new cmd instance
func New(name, binaryPath string, args []string, opts ...Option) (Execer, error) {
	instance := newDefaultImpl(name, binaryPath, args)
	for _, opt := range opts {
		opt(instance)
	}

	instance.cmd = osexec.Command(binaryPath, args...)

	instance.envs = append(instance.envs, os.Environ()...)
	instance.cmd.Env = instance.envs
	instance.cmd.Dir = instance.workingDir
	instance.cmd.Stdin = os.Stdin
	instance.cmd.Stdout = instance.processLogger
	instance.cmd.Stderr = instance.processLogger

	return instance, nil
}

// Run run a command, and wait for it to complete
func (ei *defaultImpl) Start() error {
	if err := ei.cmd.Start(); err != nil {
		handledErr := handleError(err)
		ei.logger.Error("defaultImpl.Start", "fails to start a sub process, name=%s, err=%s", ei.name, handledErr)
		return handledErr
	}
	if err := ei.cmd.Wait(); err != nil {
		handledErr := handleError(err)
		ei.logger.Error("defaultImpl.Start", "fails to Wait a sub process, name=%s, err=%s", ei.name, handledErr)
		return handledErr
	}
	return nil
}

// Stop try to kill sub commandline process by sending signals
func (ei *defaultImpl) Stop() {
	if ei.cmd.Process == nil {
		return
	}

	ei.cmd.Process.Signal(syscall.SIGTERM)

	_ = time.AfterFunc(10*time.Second, func() {
		if !ei.cmd.ProcessState.Exited() {
			ei.cmd.Process.Signal(syscall.SIGKILL)
		}
	})
}
