package util

import (
	"context"
	"eduFair/domain"
	"time"
)

func ResponseInterceptor(ctx context.Context, resp *domain.ApiResponse) {
	traceIdInf := ctx.Value("requestid")
	traceId := ""
	if traceIdInf != nil {
		traceId = traceIdInf.(string)
	}

	resp.Timestamp = time.Now()
	resp.TraceID = traceId
}
