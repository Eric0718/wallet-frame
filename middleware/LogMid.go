package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"strings"
	"wallet-frame/logger"
	"wallet-frame/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func LoggerMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//请求方式
		method := ctx.Request.Method
		//请求路由
		reqUrl := ctx.Request.RequestURI
		//请求路由
		reqUrlList := strings.Split(ctx.Request.URL.String(), "?")
		//状态码
		statusCode := ctx.Writer.Status()
		//请求ip
		clientIP := utils.GetCurrentIP(ctx)
		// 打印日志
		//loggerMap := map[string]interface{} {
		//	"status_code":statusCode,
		//	"client_ip": clientIP,
		//	"req_method":method,
		//	"req_uri": reqUrl,
		//}
		/*--------------TODO 日志可以存储 获取请求体数据----------------*/
		var data map[string]interface{}
		body, err := io.ReadAll(ctx.Request.Body)
		// 等于拷贝一份往下传递,否则下接口的方法中拿不到请求体数据
		ctx.Request.Body = io.NopCloser(bytes.NewBuffer(body))
		if err != nil {
			fmt.Println(err, "????")
		}
		err = json.Unmarshal(body, &data)
		message := utils.MapToJson(data)
		fmt.Println("当前请求POST数据:", message)
		fmt.Println("当前请求GET参数:", reqUrlList)
		fmt.Println(ctx.Request.URL.Query(), "<---请求地址")
		/*--------------TODO 日志可以存储 获取请求体数据----------------*/
		//marshal, _ := json.Marshal(loggerMap)
		loggerStr := fmt.Sprintf("status_code:%s,client_ip:%s,req_method:%s,req_uri:%s", strconv.Itoa(statusCode), clientIP, method, reqUrl)
		logger.Logs.Info("中间件本次请求", zap.String("http", loggerStr))
		ctx.Next()
	}
}
