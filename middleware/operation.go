package middleware

import (
	"bytes"
	"fmt"
	"github.com/codycoding/goDuck/global"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"time"
)

//
// OperationRecordData
//  @Description: 数据记录结构
//
type OperationRecordData struct {
	Id           int64     `json:"id" form:"id" gorm:"column:id;primary"`
	Ip           string    `json:"ip" form:"ip" gorm:"column:ip;comment:请求ip"`                                   // 请求ip
	CreatedTime  time.Time `json:"createdTime" from:"createdTime" gorm:"column:created_time;comment:创建时间"`       // 记录创建时间
	Method       string    `json:"method" form:"method" gorm:"column:method;comment:请求方法"`                       // 请求方法
	Path         string    `json:"path" form:"path" gorm:"column:path;comment:请求路径"`                             // 请求路径
	Status       int       `json:"status" form:"status" gorm:"column:status;comment:请求状态"`                       // 请求状态
	Latency      int64     `json:"latency" form:"latency" gorm:"column:latency;comment:延迟" swaggertype:"string"` // 延迟
	Agent        string    `json:"agent" form:"agent" gorm:"column:agent;comment:代理"`                            // 代理
	ErrorMessage string    `json:"error_message" form:"error_message" gorm:"column:error_message;comment:错误信息"`  // 错误信息
	Body         string    `json:"body" form:"body" gorm:"type:longtext;column:body;comment:请求Body"`             // 请求Body
	Resp         string    `json:"resp" form:"resp" gorm:"type:longtext;column:resp;comment:响应Body"`             // 响应Body
	UserID       int64     `json:"user_id" form:"user_id" gorm:"column:user_id;comment:用户id"`                    // 用户id
}

func (OperationRecordData) TableName() string {
	return "authority_api_operation_records"
}

func OperationRecord() gin.HandlerFunc {
	return func(c *gin.Context) {
		var body []byte
		var userId int64
		if c.Request.Method != http.MethodGet {
			var err error
			body, err = ioutil.ReadAll(c.Request.Body)
			if err != nil {
				global.Log.Error("read body from request error:", zap.Any("err", err))
			} else {
				c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
			}
		}
		if claims, ok := c.Get("claims"); ok {
			waitUse := claims.(*CustomClaims)
			userId = waitUse.UserInfo.AccountId
		} else {
			userId = 0
		}
		record := OperationRecordData{
			Ip:     c.ClientIP(),
			Method: c.Request.Method,
			Path:   c.Request.URL.Path,
			Agent:  c.Request.UserAgent(),
			Body:   string(body),
			UserID: userId,
		}
		// 存在某些未知错误 TODO
		//values := c.Request.Header.Values("content-type")
		//if len(values) >0 && strings.Contains(values[0], "boundary") {
		//	record.Body = "file"
		//}
		writer := responseBodyWriter{
			ResponseWriter: c.Writer,
			body:           &bytes.Buffer{},
		}
		c.Writer = writer
		now := time.Now()

		c.Next()

		latency := time.Now().Sub(now).Milliseconds()
		record.ErrorMessage = c.Errors.ByType(gin.ErrorTypePrivate).String()
		record.Status = c.Writer.Status()
		record.Latency = latency
		fmt.Println(latency)
		record.Resp = writer.body.String()

		if err := global.PostgresDb.Create(&record).Error; err != nil {
			global.Log.Error("create operation record error:", zap.Any("err", err))
		}
	}
}

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}
