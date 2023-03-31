package controllers

type ControllerBuilder interface {
	BuildPingController() PingController
}

type controllerBuildImpl struct {}

func NewCtrlFactory() ControllerBuilder {
	return &controllerBuildImpl{}
}

func (ctrlFactory *controllerBuildImpl) BuildPingController() PingController {
	return NewPingController()
}
