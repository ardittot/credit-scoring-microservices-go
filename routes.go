package main

func initializeRoutes() {

  // Handle the index route
  router.GET("/crs", GetStatus)
  router.GET("/crs/:id", GetStatusSingle)
  router.POST("/crs", CreateStatus)
  router.DELETE("/crs/:id", DeleteStatus)

}
