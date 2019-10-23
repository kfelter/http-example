package main

func (s *server) setRoutes() {
	s.router.GET("/status", s.handleStatus())
	s.router.POST("/data", s.handleData())
}
