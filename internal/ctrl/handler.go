package ctrl

import (
	"day2/internal/handler"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

func CountApi(c *gin.Context) {
	//参数出现加号处理
	c.Request.URL.RawQuery = strings.ReplaceAll(c.Request.URL.RawQuery, "+", "%2B")
	//参数出现空格处理
	c.Request.URL.RawQuery = strings.ReplaceAll(c.Request.URL.RawQuery, " ", "%20")
	s := c.Query("string")
	str := strconv.Itoa(handler.Couter(handler.Analy(s)))
	c.String(http.StatusOK, "结果为："+str)
}
