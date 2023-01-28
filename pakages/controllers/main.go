package controllers

type Controller struct {
}

//func (controller Controller) HomeControllerFactory(clientDB *gorm.DB) func(http.ResponseWriter, *http.Request) {
//	return func(response http.ResponseWriter, request *http.Request) {
//
//		response.Header().Set("Content-Type", "application/json")
//
//		_ = json.NewEncoder(response).Encode(shared.ApiResponse{
//			Ok:      true,
//			Status:  200,
//			Message: "",
//			Content: "Hello World!",
//		})
//	}
//}
