// Package lumberjack provides a rolling logger.
//
// Note that this is v2.0 of lumberjack, and should be imported using gopkg.in
// thusly:
//
//   import "gopkg.in/natefinch/lumberjack.v2"
//
// The package name remains simply lumberjack, and the code resides at
// https://github.com/natefinch/lumberjack under the v2.0 branch.
//
// Lumberjack is intended to be one part of a logging infrastructure.
// It is not an all-in-one solution, but instead is a pluggable
// component at the bottom of the logging stack that simply controls the files
// to which logs are written.
//
// Lumberjack plays well with any logging package that can write to an
// io.Writer, including the standard library's log package.
//
// Lumberjack assumes that only one process is writing to the output files.
// Using the same lumberjack configuration from multiple processes on the same
// machine will result in improper behavior.
package zaplogger

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

type WriteSyncLogger struct {
	Filename string
	file     *os.File
	mu       sync.Mutex
}

func (l *WriteSyncLogger) Write(p []byte) (n int, err error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.file == nil {
		if err = l.openExistingOrNew(len(p)); err != nil {
			return 0, err
		}
	}

	n, err = l.file.Write(p)

	return n, err
}

func (l *WriteSyncLogger) openExistingOrNew(writeLen int) error {
	filename := l.Filename
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return l.openNew()
	}
	if err != nil {
		return fmt.Errorf("error getting log file info: %s", err)
	}

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		// if we fail to open the old log file for some reason, just ignore
		// it and open a new log file.
		return l.openNew()
	}
	l.file = file
	return nil
}

func (l *WriteSyncLogger) filename() string {
	if l.Filename != "" {
		return l.Filename
	}
	name := filepath.Base(os.Args[0]) + "-lumberjack.log"
	return filepath.Join(os.TempDir(), name)
}

func (l *WriteSyncLogger) dir() string {
	return filepath.Dir(l.filename())
}

func (l *WriteSyncLogger) openNew() error {
	err := os.MkdirAll(l.dir(), 0744)
	if err != nil {
		return fmt.Errorf("can't make directories for new logfile: %s", err)
	}

	name := l.filename()
	mode := os.FileMode(0644)

	// we use truncate here because this should only get called when we've moved
	// the file ourselves. if someone else creates the file in the meantime,
	// just wipe out the contents.
	f, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, mode)
	if err != nil {
		return fmt.Errorf("can't open new logfile: %s", err)
	}
	l.file = f
	return nil
}

// Rotate causes Logger to close the existing log file and immediately create a
// new one.  This is a helper function for applications that want to initiate
// rotations outside of the normal rotation rules, such as in response to
// SIGHUP.  After rotating, this initiates compression and removal of old log
// files according to the configuration.
func (l *WriteSyncLogger) Rotate() error {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.rotate()
}

// rotate closes the current file, moves it aside with a timestamp in the name,
// (if it exists), opens a new file with the original filename, and then runs
// post-rotation processing and removal.
func (l *WriteSyncLogger) rotate() error {
	if err := l.close(); err != nil {
		return err
	}
	if err := l.openNew(); err != nil {
		return err
	}
	return nil
}

func (l *WriteSyncLogger) Close() error {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.close()
}

// close closes the file if it is open.
func (l *WriteSyncLogger) close() error {
	if l.file == nil {
		return nil
	}
	err := l.file.Close()
	l.file = nil
	return err
}

func (l *WriteSyncLogger) Sync() error {
	return l.file.Sync()
}
