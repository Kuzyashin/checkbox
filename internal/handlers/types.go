package handlers


type requestUri struct {
	RequestID uint64 `uri:"request_id" binding:"required"`
}
