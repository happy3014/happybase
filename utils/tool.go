package utils

import (
	"fmt"
	"os"
	"strings"
)

// CheckFileExists 检查文件是否存在
func CheckFileExists(filepath string) (bool, error) {
	_, err := os.Stat(filepath)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

// CreateDirIfNotExists 如果目录不存在，则创建目录
func CreateDirIfNotExists(d string) error {
	exists, err := CheckFileExists(d)
	if err != nil {
		return fmt.Errorf("failed to check %s exists: %v", d, err)
	}
	if !exists {
		err = os.MkdirAll(d, 0644)
		if err != nil {
			return fmt.Errorf("failed to create %s: %v", d, err)
		}
	}
	return nil
}

// StrCenter 将字符串居中对齐，并用指定字符填充空白部分扩展字符串到指定长度
func StrCenter(s string, length int, padChar string) string {
	if len(padChar) != 1 {
		panic(fmt.Errorf("pad char should be a single character: \"%s\" is invalid", padChar))
	}
	if len(s) >= length {
		return s
	}

	extraLen := length - len(s)
	leftLen := extraLen / 2
	rightLen := extraLen - leftLen

	return strings.Repeat(padChar, leftLen) + s + strings.Repeat(padChar, rightLen)

}
