package util

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
	"bytes"
	"io"
)

// 获得项目根目录
func GetProjectRoot() string {
	binDir, err := executableDir()
	if err != nil {
		return ""
	}
	return path.Dir(binDir)
}

// 获得可执行程序所在目录
func executableDir() (string, error) {
	pathAbs, err := filepath.Abs(os.Args[0])
	if err != nil {
		return "", err
	}
	return filepath.Dir(pathAbs), nil
}

func Welcome() {
	fmt.Println("***********************************")
	fmt.Println("*******欢迎来到Go语言中文网*******")
	fmt.Println("***********************************")
}

// strings.Index的UTF-8版本
// 即 Utf8Index("Go语言中文网", "学习") 返回 4，而不是strings.Index的 8
func Utf8Index(str, substr string) int {
	asciiPos := strings.Index(str, substr)
	if asciiPos == -1 || asciiPos == 0 {
		return asciiPos
	}
	pos := 0
	totalSize := 0
	reader := strings.NewReader(str)
	for _, size, err := reader.ReadRune(); err == nil; _, size, err = reader.ReadRune() {
		totalSize += size
		pos++
		// 匹配到
		if totalSize == asciiPos {
			return pos
		}
	}
	return pos
}
func ByteRWerExample() {
FOREND:
	for {
		fmt.Println("请输入要通过WriteByte写入的一个ASCII字符（b：返回上级；q：退出）：")
		var ch byte
		fmt.Scanf("%c\n", &ch)
		switch ch {
		case 'b':
			fmt.Println("返回上级菜单！")
			break FOREND
		case 'q':
			fmt.Println("程序退出！")
			os.Exit(0)
		default:
			buffer := new(bytes.Buffer)
			err := buffer.WriteByte(ch)
			if err == nil {
				fmt.Println("写入一个字节成功！准备读取该字节……")
				newCh, _ := buffer.ReadByte()
				fmt.Printf("读取的字节：%c\n", newCh)
			} else {
				fmt.Println("写入错误")
			}
		}

	}
}

func ReadFrom(reader io.Reader, num int) ([]byte, error) {
	p := make([]byte, num)
	n, err := reader.Read(p)
	if n > 0 {
		return p[:n], nil
	}
	return p, err
}

func CommandOutput() {

}
