package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PddController(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pdd",
	})
	// args := {
	// 	client_id: "",
	// 	data_type: "JSON",
	// 	sort_type: 0,
	// 	type: "pdd.ddk.goods.search",
	// 	pid: "",
	// 	page: 1,
	// 	block_cat_packages: ["1", "2"],
	// }
}
