package example_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

const (
	defTimeoutSecond = 10
	keyEnvDebug      = "CI_DEBUG"
	keyEnvCiNum      = "CI_NUMBER"
	keyEnvCiKey      = "CI_KEY"
	keyEnvCiKeys     = "CI_KEYS"
)

var (
	// testBaseFolderPath
	//  test base dir will auto get by package init()
	testBaseFolderPath = ""

	envDebug = false

	envCiNum  = 0
	envCiKey  = ""
	envCiKeys []string

	strData []string
)

func init() {
	testBaseFolderPath, _ = getCurrentFolderPath()
	envDebug = fetchOsEnvBool(keyEnvDebug, false)
	envCiNum = fetchOsEnvInt(keyEnvCiNum, 1)
	envCiKey = fetchOsEnvStr(keyEnvCiKey, "")
	envCiKeys = fetchOsEnvArray(keyEnvCiKeys)
	for i := 0; i < 200; i++ {
		strData = append(strData, randomStr(300))
	}
}

// test case file tools start

// goldenDataSaveFast
//
//	save data to golden file
//	style as: "TestFuncName/extraName.golden"
func goldenDataSaveFast(t *testing.T, data interface{}, extraName string) error {
	marshal, errJson := json.Marshal(data)
	if errJson != nil {
		t.Fatal(errJson)
	}
	return goldenDataSave(t, marshal, extraName, os.FileMode(0766))
}

// goldenDataSave
//
//	save data to golden file
//	style as: "TestFuncName/extraName.golden"
func goldenDataSave(t *testing.T, data []byte, extraName string, fileMod fs.FileMode) error {
	testDataFolderFullPath, err := getOrCreateTestDataFolderFullPath()
	if err != nil {
		return fmt.Errorf("try goldenDataSave err: %v", err)
	}
	testDataFolder := filepath.Join(testDataFolderFullPath, t.Name())
	if !pathExistsFast(testDataFolder) {
		errMk := mkdir(testDataFolder)
		if errMk != nil {
			t.Fatal(errMk)
		}
	}
	savePath := filepath.Join(testDataFolderFullPath, t.Name(), fmt.Sprintf("%s.golden", extraName))
	var str bytes.Buffer
	err = json.Indent(&str, data, "", "  ")
	if err != nil {
		return err
	}
	err = writeFileByByte(savePath, str.Bytes(), fileMod, true)
	if err != nil {
		return fmt.Errorf("try goldenDataSave at path: %s err: %v", savePath, err)
	}
	return nil
}

// goldenDataReadAsByte
//
//	read golden file as byte
//	style as: "TestFuncName/extraName.golden"
func goldenDataReadAsByte(t *testing.T, extraName string) ([]byte, error) {
	testDataFolderFullPath, err := getOrCreateTestDataFolderFullPath()
	if err != nil {
		return nil, fmt.Errorf("try goldenDataReadAsByte err: %v", err)
	}

	savePath := filepath.Join(testDataFolderFullPath, t.Name(), fmt.Sprintf("%s.golden", extraName))

	fileAsByte, err := readFileAsByte(savePath)
	if err != nil {
		return nil, fmt.Errorf("try goldenDataReadAsByte err: %v", err)
	}
	return fileAsByte, nil
}

// goldenDataReadAsType
//
//	read golden file as type
//	style as: "TestFuncName/extraName.golden"
func goldenDataReadAsType(t *testing.T, extraName string, v interface{}) error {
	readAsByte, err := goldenDataReadAsByte(t, extraName)
	if err != nil {
		t.Fatal(err)
	}
	return json.Unmarshal(readAsByte, v)
}

var currentTestDataFolderAbsPath = ""

// getOrCreateTestDataFullPath
//
//	get or create test data full path will under this package testdata
//	this function will create dir for return full path
func getOrCreateTestDataFullPath(elem ...string) (string, error) {
	if elem == nil || len(elem) < 1 {
		return "", fmt.Errorf("must has one elem")
	}
	dataFolderFullPath, err := getOrCreateTestDataFolderFullPath()
	if err != nil {
		return "", err
	}
	fullPath := filepath.Join(dataFolderFullPath, elem[0])
	if len(elem) > 1 {
		for i := 1; i < len(elem); i++ {
			fullPath = filepath.Join(fullPath, elem[i])
		}
	}
	baseDir := filepath.Dir(fullPath)
	if !pathExistsFast(baseDir) {
		errMkdir := mkdir(baseDir)
		if errMkdir != nil {
			return fullPath, errMkdir
		}
	}
	if !pathIsDir(baseDir) {
		return "", fmt.Errorf("getOrCreateTestDataFullPath exist file, and can not create dir at path: %s", baseDir)
	}

	return fullPath, nil
}

// getOrCreateTestDataFolderFullPath
//
//	will create testdata folder under this package named `testdata`
func getOrCreateTestDataFolderFullPath() (string, error) {
	if currentTestDataFolderAbsPath != "" {
		return currentTestDataFolderAbsPath, nil
	}
	if testBaseFolderPath == "" {
		return "", fmt.Errorf("can not get testBaseFolderPath")
	}
	currentTestDataFolderAbsPath = filepath.Join(testBaseFolderPath, "testdata")
	if !pathExistsFast(currentTestDataFolderAbsPath) {
		err := mkdir(currentTestDataFolderAbsPath)
		if err != nil {
			currentTestDataFolderAbsPath = ""
			return "", err
		}
	}
	return currentTestDataFolderAbsPath, nil
}

// getCurrentFolderPath
//
//	can get run path this golang dir
func getCurrentFolderPath() (string, error) {
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		return "", errors.New("can not get current file info")
	}
	return filepath.Dir(file), nil
}

// pathExists
//
//	path exists
func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// pathExistsFast
//
//	path exists fast
func pathExistsFast(path string) bool {
	exists, _ := pathExists(path)
	return exists
}

// pathIsDir
//
//	path is dir
func pathIsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// rmDir
//
//	remove dir by path
//
//nolint:golint,unused
func rmDir(path string, force bool) error {
	if force {
		return os.RemoveAll(path)
	}
	exists, err := pathExists(path)
	if err != nil {
		return err
	}
	if exists {
		return os.RemoveAll(path)
	}
	return fmt.Errorf("remove dir not exist path: %s , use force can cover this err", path)
}

// mkdir
//
//	will use FileMode 0766
func mkdir(path string) error {
	err := os.MkdirAll(path, os.FileMode(0766))
	if err != nil {
		return fmt.Errorf("fail MkdirAll at path: %s , err: %v", path, err)
	}
	return nil
}

// readFileAsByte
//
//	read file as byte array
func readFileAsByte(path string) ([]byte, error) {
	exists, err := pathExists(path)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, fmt.Errorf("path not exist %v", path)
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("path: %s read err: %v", path, err)
	}

	return data, nil
}

// readFileAsJson
//
//	read file as json
func readFileAsJson(path string, v interface{}) error {
	fileAsByte, errRead := readFileAsByte(path)
	if errRead != nil {
		return fmt.Errorf("path: %s , read file as err: %v", path, errRead)
	}
	err := json.Unmarshal(fileAsByte, v)
	if err != nil {
		return fmt.Errorf("path: %s , read file as json err: %v", path, err)
	}
	return nil
}

// writeFileByByte
//
//	write bytes to file
//	path most use Abs Path
//	data []byte
//	fileMod os.FileMode(0666) or os.FileMode(0644)
//	coverage true will coverage old
func writeFileByByte(path string, data []byte, fileMod fs.FileMode, coverage bool) error {
	if !coverage {
		exists, err := pathExists(path)
		if err != nil {
			return err
		}
		if exists {
			return fmt.Errorf("not coverage, which path exist %v", path)
		}
	}
	parentPath := filepath.Dir(path)
	if !pathExistsFast(parentPath) {
		err := os.MkdirAll(parentPath, fileMod)
		if err != nil {
			return fmt.Errorf("can not WriteFileByByte at new dir at mode: %v , at parent path: %v", fileMod, parentPath)
		}
	}
	err := os.WriteFile(path, data, fileMod)
	if err != nil {
		return fmt.Errorf("write data at path: %v, err: %v", path, err)
	}
	return nil
}

// writeFileAsJson write json file
//
//	path most use Abs Path
//	v data
//	fileMod os.FileMode(0666) or os.FileMode(0644)
//	coverage true will coverage old
//	beauty will format json when write
func writeFileAsJson(path string, v interface{}, fileMod fs.FileMode, coverage, beauty bool) error {
	if !coverage {
		exists, err := pathExists(path)
		if err != nil {
			return err
		}
		if exists {
			return fmt.Errorf("not coverage, which path exist %v", path)
		}
	}
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	if beauty {
		var str bytes.Buffer
		err := json.Indent(&str, data, "", "  ")
		if err != nil {
			return err
		}
		return writeFileByByte(path, str.Bytes(), fileMod, coverage)
	}
	return writeFileByByte(path, data, fileMod, coverage)
}

// writeFileAsJsonBeauty
//
//	write json file as 0766 and beauty
func writeFileAsJsonBeauty(path string, v interface{}, coverage bool) error {
	return writeFileAsJson(path, v, os.FileMode(0766), coverage, true)
}

// test case file tools end
