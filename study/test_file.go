package main

import (
	"fmt"
	"os"
)

func main() {
	//err := createFile("test.txt")
	//if err != nil {
	//	fmt.Printf("create file error: %v\n", err)
	//}
	//
	//err = createDir("study/new")
	//if err != nil {
	//	fmt.Printf("create dir error: %v\n", err)
	//}
	//
	//err = removeFile("test2.txt")
	//if err != nil {
	//	fmt.Printf("remove file error: %v\n", err)
	//}
	//
	//err = removeDir("study/new")
	//if err != nil {
	//	fmt.Printf("remove dir error: %v\n", err)
	//}
	getWd()
}

func createFile(path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	} else {
		fmt.Printf("file is %v\n", *f)
		return nil
	}
}

func createDir(path string) error {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func removeFile(path string) error {
	err := os.RemoveAll(path)
	if err != nil {
		return err
	}
	return nil
}

func removeDir(path string) error {
	err := os.RemoveAll(path)
	if err != nil {
		return err
	}
	return nil
}

func getWd() {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Printf("get wd error: %v\n", err)
	}
	fmt.Printf("wd: %v\n", wd)
}
