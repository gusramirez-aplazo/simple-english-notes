package controllers

type Controller struct {
}

var controller = Controller{}

func GetController() *Controller {
	return &controller
}
