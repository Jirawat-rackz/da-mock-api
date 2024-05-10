package httpserve

import "github.com/gin-gonic/gin"

func (s Server) Use(middleware ...gin.HandlerFunc) gin.IRoutes {
	return s.Engine.Use(middleware...)
}

func (s Server) Group(relativePath string, handlers ...gin.HandlerFunc) *gin.RouterGroup {
	return s.Engine.Group(relativePath, handlers...)
}

func (s Server) GET(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return s.Engine.GET(relativePath, handlers...)
}

func (s Server) POST(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return s.Engine.POST(relativePath, handlers...)
}

func (s Server) PUT(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return s.Engine.PUT(relativePath, handlers...)
}

func (s Server) DELETE(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return s.Engine.DELETE(relativePath, handlers...)
}

func (s Server) PATCH(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return s.Engine.PATCH(relativePath, handlers...)
}
