package handlers

import (
	"net/http"
	"runtime"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"hospital-queue/service"
)

// MainHandler 渲染主页面
func MainHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title":    "医院叫号系统",
		"datetime": time.Now().Format("2006-01-02 15:04:05"),
		"os":       runtime.GOOS,
		"arch":     runtime.GOARCH,
	})
}

// IndexHandler 渲染主页面
func IndexHandler(c *gin.Context) {
	c.Request.URL.Path = "/"
	MainHandler(c)
}

// GetAllQueuesHandler 获取队列信息
func GetAllQueuesHandler(c *gin.Context) {
	patients, err := service.ReadAllQueues()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	queues := make([]map[string]any, 0, len(patients))
	for _, patient := range patients {
		remaining := 0
		for _, p := range patients {
			if p.Department == patient.Department && p.Status == 0 && p.Number < patient.Number {
				remaining++
			}
		}

		queues = append(queues, map[string]any{
			"id":         patient.ID,
			"number":     patient.Number,
			"name":       patient.Name,
			"phone":      patient.Phone,
			"department": patient.Department,
			"status":     patient.Status,
			"datetime":   patient.CreateAt.Format("2006-01-02 15:04:05"),
			"remaining":  remaining,
		})
	}

	c.JSON(http.StatusOK, queues)
}

// CreateQueueHandler 新增排队号码
func CreateQueueHandler(c *gin.Context) {
	department := c.PostForm("department")
	if department == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "诊室号不能为空"})
		return
	}
	departmentNum, err := strconv.Atoi(department)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if departmentNum <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "诊室号码必须为正整数",
		})
	}

	name := c.PostForm("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "患者姓名不能为空"})
		return
	}

	phone := c.PostForm("phone")
	newQueue, err := service.CreateNewQueue(name, phone, uint(departmentNum))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, newQueue)
}
