package main

import (
	_ "gett2/routers"
	"github.com/astaxie/beego"
	_ "github.com/lib/pq"
)

func main() {
	beego.Run()
}
