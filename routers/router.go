package routers

import (
	"github.com/astaxie/beego"

	"EZChain/controllers"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}