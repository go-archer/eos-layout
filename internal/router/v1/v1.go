package v1

import (
	"eos-layout/internal/handler"

	"github.com/gin-gonic/gin"
)

type Router struct {
	root *gin.RouterGroup
}

func Register(r *gin.RouterGroup) *Router {
	v1 := r.Group("/v1")
	return &Router{root: v1}
}

func (r *Router) AreaRouter(h handler.AreaHandler) *Router {
	g := r.root.Group("/areas")
	g.GET("/provinces", h.ProvinceList)
	g.GET("/provinces/:id", h.Province)
	g.GET("/provinces/:id/cities", h.CityList)
	g.GET("/cities/:id", h.City)
	g.GET("/cities/:id/districts", h.DistrictList)
	g.GET("/districts/:id", h.District)
	g.GET("/districts/:id/streets", h.StreetList)
	g.GET("/streets/:id", h.Street)
	g.GET("/streets/:id/committees", h.CommitteeList)
	g.GET("/committees/:id", h.Committee)
	return r
}
