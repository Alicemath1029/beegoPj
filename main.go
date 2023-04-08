package main

import "github.com/astaxie/beego"

func main() {
	beego.Info("hello beego")
	beego.Run("localhost:80")
}
