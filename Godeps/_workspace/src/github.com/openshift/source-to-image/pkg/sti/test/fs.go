package test

import (
	"bytes"
	"io"
	"os"
	"sync"
)

// FakeFileSystem provides a fake filesystem structure for testing
type FakeFileSystem struct {
	ChmodFile  []string
	ChmodMode  os.FileMode
	ChmodError map[string]error

	RenameFrom  string
	RenameTo    string
	RenameError error

	MkdirAllDir   []string
	MkdirAllError error

	MkdirDir   string
	MkdirError error

	ExistsFile   []string
	ExistsResult map[string]bool

	CopySource string
	CopyDest   string
	CopyError  error

	RemoveDirName  string
	RemoveDirError error

	WorkingDirCalled bool
	WorkingDirResult string
	WorkingDirError  error

	OpenFile       string
	OpenFileResult *FakeReadCloser
	OpenContent    string
	OpenError      error
	OpenCloseError error

	WriteFileName    string
	WriteFileError   error
	WriteFileContent string

	mutex sync.Mutex
}

// FakeReadCloser provider a fake ReadCloser
type FakeReadCloser struct {
	*bytes.Buffer
	CloseCalled bool
	CloseError  error
}

// Close closes the fake ReadCloser
func (f *FakeReadCloser) Close() error {
	f.CloseCalled = true
	return f.CloseError
}

// Chmod manipulates permissions on the fake filesystem
func (f *FakeFileSystem) Chmod(file string, mode os.FileMode) error {
	f.mutex.Lock()
	defer f.mutex.Unlock()
	f.ChmodFile = append(f.ChmodFile, file)
	f.ChmodMode = mode

	return f.ChmodError[file]
}

// Rename renames files on the fake filesystem
func (f *FakeFileSystem) Rename(from, to string) error {
	f.RenameFrom = from
	f.RenameTo = to
	return f.RenameError
}

// MkdirAll creates a new directories on the fake filesystem
func (f *FakeFileSystem) MkdirAll(dirname string) error {
	f.MkdirAllDir = append(f.MkdirAllDir, dirname)
	return f.MkdirAllError
}

// Mkdir creates a new directory on the fake filesystem
func (f *FakeFileSystem) Mkdir(dirname string) error {
	f.MkdirDir = dirname
	return f.MkdirError
}

// Exists checks if the file exists in fake filesystem
func (f *FakeFileSystem) Exists(file string) bool {
	f.ExistsFile = append(f.ExistsFile, file)
	return f.ExistsResult[file]
}

// Copy copies files on the fake filesystem
func (f *FakeFileSystem) Copy(sourcePath, targetPath string) error {
	f.CopySource = sourcePath
	f.CopyDest = targetPath
	return f.CopyError
}

// RemoveDirectory removes a directory in the fake filesystem
func (f *FakeFileSystem) RemoveDirectory(dir string) error {
	f.RemoveDirName = dir
	return f.RemoveDirError
}

// CreateWorkingDirectory creates a fake working directory
func (f *FakeFileSystem) CreateWorkingDirectory() (string, error) {
	f.WorkingDirCalled = true
	return f.WorkingDirResult, f.WorkingDirError
}

// Open opens a file
func (f *FakeFileSystem) Open(file string) (io.ReadCloser, error) {
	f.OpenFile = file
	buf := bytes.NewBufferString(f.OpenContent)
	f.OpenFileResult = &FakeReadCloser{
		Buffer:     buf,
		CloseError: f.OpenCloseError,
	}
	return f.OpenFileResult, f.OpenError
}

// Write writes a file
func (f *FakeFileSystem) WriteFile(file string, data []byte) error {
	f.WriteFileName = file
	f.WriteFileContent = string(data)
	return f.WriteFileError
}
