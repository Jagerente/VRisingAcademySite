package api

import (
	"VRisingAcademySite/api/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterApiHandlers(r *gin.Engine) {
	items := r.Group("/api/items")
	{
		controllers.HandleItemsRequest(items)
	}

}
