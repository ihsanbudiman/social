package thread_handler_rest

type CreateThreadRequest struct {
	Content string `json:"content" form:"content"`
	ReplyTo uint   `json:"reply_to" form:"reply_to"`
}

type GetThreadRequest struct {
	ThreadID uint `json:"thread_id" query:"thread_id"`
}

type GetRepliesRequest struct {
	ThreadID uint `json:"thread_id" query:"thread_id"`
	Page     int  `json:"page" query:"page"`
}
