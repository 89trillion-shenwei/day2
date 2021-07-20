package ctrl

import (
	"day2/internal"
	"day2/internal/handler"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

func CountApi(c *gin.Context) (string, error) {
	s := c.PostForm("string")
	s = strings.Replace(s, " ", "", -1)
	if s == "" {
		return "", internal.IsEmptyError("表达式不能为空")
	}

	e1, e2, err := handler.Analy(s)
	if err != nil {
		return "", internal.IllegalExpressionError(err.Error())
	}
	str := strconv.Itoa(handler.Couter(e1, e2))
	return str, nil
}

type Api func(c *gin.Context) (string, error)

func ReturnData(api Api) gin.HandlerFunc {
	return func(c *gin.Context) {
		data, err := api(c)
		if err != nil {
			globalError := err.(internal.GlobalError)
			c.JSON(globalError.Status, globalError)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"查询结果": data,
		})
	}
}
