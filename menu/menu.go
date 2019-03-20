package menu

import (
	"fmt"
	"strings"
	"os"
	"beego_test/util"
)

func MainMenu() {
MAINFOR:
	for {
		fmt.Println("")
		fmt.Println("*******请选择：*********")
		fmt.Println("1 输入命令")
		fmt.Println("q 退出")
		fmt.Println("***********************************")

		var ch string
		fmt.Scanln(&ch)

		switch ch {
		case "1":
			CommandMenu()
			//case "2":
			//	ByteRWerExample()
		case "q":
			fmt.Println("程序退出！")
			break MAINFOR
		default:
			fmt.Println("输入错误！")
			continue
		}
	}
}

func CommandMenu() {
FOREND:
	for {
		readerMenu()

		var ch,cmd string
		fmt.Scanln(&ch)
		var (
			_ []byte
			err  error
		)
		switch strings.ToLower(ch) {
		case "1":
			fmt.Println("请输入命令，以回车结束：")
			fmt.Scanln(&cmd)
		case "2":
			file, err := os.Open(util.GetProjectRoot() + "/src/chapter01/io/01.txt")
			if err != nil {
				fmt.Println("打开文件 01.txt 错误:", err)
				continue
			}
			//data, err = util.ReadFrom(file, 9)
			file.Close()
		case "b":
			fmt.Println("返回上级菜单！")
			break FOREND
		case "q":
			fmt.Println("程序退出！")
			os.Exit(0)
		default:
			fmt.Println("输入错误！")
			continue
		}

		if err != nil {
			fmt.Println("数据读取失败，可以试试从其他输入源读取！")
		} else {
			fmt.Printf("读取到的数据是：%s\n", cmd)
		}
	}
}

func readerMenu() {
	fmt.Println("")
	fmt.Println("*******执行命令的不同方式*********")
	fmt.Println("*******请选择类型：*********")
	fmt.Println("1 直接执行脚本")
	//fmt.Println("2 读取并执行文件内容")
	fmt.Println("b 返回上级菜单")
	fmt.Println("q 退出")
	fmt.Println("***********************************")
}
