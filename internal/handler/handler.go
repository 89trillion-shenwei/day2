package handler

import (
	util2 "day2/internal/ctrl"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

func CountApi(c *gin.Context)  {
	c.Request.URL.RawQuery = strings.ReplaceAll(c.Request.URL.RawQuery, "+", "%2B")
	c.Request.URL.RawQuery = strings.ReplaceAll(c.Request.URL.RawQuery, " ", "%20")
	s :=c.Query("string")
	str:=strconv.Itoa(util2.Couter(util2.Analy(s)))
	c.String(http.StatusOK,"结果为："+str)
}
