package file

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	"os"
)

func PathExists(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err == nil {
		if fi.IsDir() {
			return true, nil
		}
		return false, errors.New("存在同名文件")
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func CreateDir(dirs ...string) (err error) {
	for _, v := range dirs {
		exist, err := PathExists(v)
		if err != nil {
			return err
		}
		if !exist {
			fmt.Print("create directory" + v)
			if err := os.MkdirAll(v, os.ModePerm); err != nil {
				fmt.Printf("create directory"+v, zap.Any(" error:", err))
				return err
			}
		}
	}
	return err
}
