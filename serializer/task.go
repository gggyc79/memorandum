package serializer

// swagger:response Resp
type Task struct {
	ID        uint   `json:"id" example:"1"`       // 任务ID
	Title     string `json:"title" example:"吃饭"`   // 题目
	Content   string `json:"content" example:"睡觉"` // 内容
	View      uint64 `json:"view" example:"32"`    // 浏览量
	Status    int    `json:"status" example:"0"`   // 状态(0未完成，1已完成)
	CreatedAt int64  `json:"created_at"`
	StartTime int64  `json:"start_time"`
	EndTime   int64  `json:"end_time"`
}

// 创建任务的服务
type CreateTaskService struct {
	Title   string `form:"title" json:"title" binding:"required,min=2,max=100"`
	Content string `form:"content" json:"content" binding:"max=1000"`
	Status  int    `form:"status" json:"status"` // 0 待办   1已完成
}
