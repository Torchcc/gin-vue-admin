package v1

import (
	"fmt"
	"strconv"

	"gin-vue-admin/global/response"
	"gin-vue-admin/service"
	"github.com/gin-gonic/gin"
)

func GetRegionsByParentId(c *gin.Context) {
	parentId, err := strconv.ParseInt(c.Query("parent_id"), 10, 64)
	if err != nil {
		parentId = int64(0)
	}
	output, err := service.RetrieveRegionsByParentId(parentId)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
	}
	response.OkWithData(output, c)
}
