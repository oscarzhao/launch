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

// WithEnvMapping set environment variables for binary to start
func WithEnvMapping(envMapping map[string]string) Option {
	return func(options *defaultImpl) {
		for k, v := range envMapping {
			options.envs = append(options.envs, k+"="+v)
		}
	}
}

// WithWorkDir set working directory
func WithWorkDir(workDir string) Option {
	return func(options *defaultImpl) {
		options.workingDir = workDir
	}
}

// WithDaemon set the process as daemon process or open a new command window
func WithDaemon(isDaemon bool) Option {
	return func(options *defaultImpl) {
		options.isDaemon = isDaemon
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

	if instance.isDaemon || len(runVerbosePrefix) == 0 {
		instance.cmd = osexec.Command(binaryPath, args...)
	} else {
		newPath := runVerbosePrefix[0]
		newArgs := append(runVerbosePrefix[1:], binaryPath)
		newArgs = append(newArgs, args...)

		instance.cmd = osexec.Command(newPath, newArgs...)
	}

	instance.cmd.Env = instance.envs
	instance.cmd.Dir = instance.workingDir
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
