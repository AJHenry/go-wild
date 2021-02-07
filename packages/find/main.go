package main

import (
	"errors"
	"math"
	"os"
	goPath "path"
	"path/filepath"
)

// Options are arguments that change the behavior of Up
type Options struct {
	depth int
	cwd string
	isDir bool
}

// Up finds a file/directory by walking up parent directories and can be passed the following options
// 
//	depth int - how far the path should walk (default: infinite)
//	cwd string - starting path (default: cwd)
//	isDir bool - whether to allow directories (default: false)
func Up(file string, options Options) (string, error) {
	if file == "" {
		return "", errors.New("File name is blank")
	}

	depth := options.depth; if options.depth == 0 {
		depth = -1
	}

	isDir := options.isDir

	cwd := options.cwd; if options.cwd == "" {
		dir, err := os.Getwd(); if err != nil {
			return "", err
		}
		cwd = dir
	}

	return search(file, cwd, depth, isDir)
}

func search(file string, path string, depth int, allowDir bool) (string, error) {

	if depth != -1 {
		return searchDepth(file, path, 0, depth, allowDir)
	}
	
	return searchDepth(file, path, -1, math.MaxInt16, allowDir)
}

func searchDepth(file string, currPath string, currDepth int, maxDepth int, allowDir bool) (string, error) {
	if currDepth > maxDepth {
		return "", errors.New("File not found in any parent directories")
	}

	switch {
	case allowDir: 
		if DirExists(file, currPath) {
			return currPath, nil
		}
	case !allowDir:
		if FileExists(file, currPath) {
			return currPath, nil
		}
	}

	// recursively search the parents
	parent := filepath.Dir(currPath)
	if parent == currPath {
		return "", errors.New("File not found in any parent directories")
	}

	return searchDepth(file, parent, currDepth + 1, maxDepth, allowDir)
}


// FileExists returns true if a file is found within the path
func FileExists(file string, path string) bool {
	filePath := goPath.Join(path, file)

	if data, err := os.Stat(filePath); err == nil {
		// path/to/whatever exists
		return !data.IsDir()
	  } else if os.IsNotExist(err) {
		// path/to/whatever does *not* exist
		return false
	  } else {
		// Schrodinger: file may or may not exist. See err for details.
	  
		// Therefore, do *NOT* use !os.IsNotExist(err) to test for file existence
	  	return false
	  }
}

// DirExists returns true if a directory is found within the path
func DirExists(file string, path string) bool {
	filePath := goPath.Join(path, file)

	if data, err := os.Stat(filePath); err == nil {
		// path/to/whatever exists
		return data.IsDir()
	  } else if os.IsNotExist(err) {
		// path/to/whatever does *not* exist
		return false
	  } else {
		// Schrodinger: file may or may not exist. See err for details.
	  
		// Therefore, do *NOT* use !os.IsNotExist(err) to test for file existence
	  	return false
	  }
}
