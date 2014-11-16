package routers

import (
	"github.com/fromruijin/order/src/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
