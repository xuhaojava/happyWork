package main


import (
	"fmt"
	"time"

	"github.com/go-vgo/robotgo"
)


func main() {
	var check int
	fmt.Print("秒级请输入1，分钟级请输入2，小时级别请输入3：")
	_,err := fmt.Scan(&check)
	if err != nil {
		fmt.Println("输入错误:",err)
		return
	}
	var timeNum int
	fmt.Print("请输入时间：")
	_,err = fmt.Scan(&timeNum)
	if err != nil {
		fmt.Println("输入错误:",err)
		return
	}
	var clickNum int
	switch check {
	case 1:
		clickNum =  timeNum / 30
	case 2:
		clickNum = timeNum * 60 / 30
	case 3:
		clickNum = timeNum * 60 * 60 / 30
	default:
		fmt.Println("时间类型输入错误")
		return
	}
	if clickNum < 0 {
		clickNum = 1
	}
	fmt.Println("begin , clickNum : ",clickNum)
	for times := 0;times<clickNum;times++{
		fmt.Println(times," click")
		robotgo.MoveMouse(0, 0)
		time.Sleep(1*time.Second)
		robotgo.Click()
		time.Sleep(1*time.Second)
		robotgo.MoveMouse(8000, 8000)
		time.Sleep(1*time.Second)
		robotgo.Click()
		for i:=30;i>0;i-- {
			fmt.Println("倒计时 ",i,"秒")
			time.Sleep(1 * time.Second)
		}
	}
}
