package task

import "fmt"

func Task(date string) {
	d := date[5:10]
	//fmt.Println(d)
	if d != "07-31" {
		return
	}
	mail() // 发送邮件
}

func mail() {
	fmt.Println("mail-success")
}
