package main

func (s *server) setRoutes() {
	s.router.GET("/status", s.handleStatus())
}
