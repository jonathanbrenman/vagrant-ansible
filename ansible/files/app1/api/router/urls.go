package router

import "app1/api/controllers"

func (r routerImpl) routes() {
	factoryCtrl := controllers.NewCtrlFactory()

	r.router.GET("/ping", factoryCtrl.BuildPingController().Ping)
}
