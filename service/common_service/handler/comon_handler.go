package handler

import (
	"context"
	"go.uber.org/zap"
	"university_circles/service/common_service/logic"
	pb "university_circles/service/common_service/pb/common"
	"university_circles/service/common_service/utils/logger"
)

type CommonHandler struct {
}

func (h *CommonHandler) Report(ctx context.Context, req *pb.ReportReq, resp *pb.Response) (err error) {
	hl := &logic.CommonLogic{}

	// 通过反馈接口  1、反馈  2、举报首页消息 3、举报商品
	if req.Type == 1 {
		err = hl.SaveUserReportToES(req)
	} else if req.Type == 2 {
		err = hl.SaveHomeMsgReportToES(req)
	} else if req.Type == 3 {
		err = hl.SaveGoodsReportToES(req)
	}

	if err != nil {
		resp.Success = -1
		resp.Msg = "user report failed"
		logger.Logger.Warn("user report msg failed", zap.Any("report", req), zap.Error(err))
		return
	}

	resp.Success = 0
	resp.Msg = "user operate success"

	return nil
}
