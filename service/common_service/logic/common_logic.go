package logic

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"university_circles/service/common_service/databases/es"
	"university_circles/service/common_service/utils/logger"

	pb "university_circles/service/common_service/pb/common"
)

const (
	USERREPORT    = "user_report"
	HOMEMSGREPORT = "home_msg_report"
	GOODSREPORT   = "goods_report"
)

type CommonLogic struct{}

func (cl *CommonLogic) SaveUserReportToES(req *pb.ReportReq) (err error) {
	esClient, err := es.NewElasticSearch()
	if err != nil {
		logger.Logger.Warn("New ElasticSearch failed", zap.Error(err))
	}

	commId := uuid.NewV4().String()
	_, err = esClient.Index().
		Index(USERREPORT).
		Id(commId).
		BodyJson(req).
		Do(context.Background())
	if err != nil {
		logger.Logger.Warn("save user report failed", zap.Any("report", req), zap.Error(err))
		return err
	}

	return nil
}

func (cl *CommonLogic) SaveHomeMsgReportToES(req *pb.ReportReq) (err error) {
	esClient, err := es.NewElasticSearch()
	if err != nil {
		logger.Logger.Warn("New ElasticSearch failed", zap.Error(err))
	}

	commId := uuid.NewV4().String()
	_, err = esClient.Index().
		Index(HOMEMSGREPORT).
		Id(commId).
		BodyJson(req).
		Do(context.Background())
	if err != nil {
		logger.Logger.Warn("save home msg report failed", zap.Any("report", req), zap.Error(err))
		return err
	}

	return nil
}

func (cl *CommonLogic) SaveGoodsReportToES(req *pb.ReportReq) (err error) {
	esClient, err := es.NewElasticSearch()
	if err != nil {
		logger.Logger.Warn("New ElasticSearch failed", zap.Error(err))
	}

	commId := uuid.NewV4().String()
	_, err = esClient.Index().
		Index(GOODSREPORT).
		Id(commId).
		BodyJson(req).
		Do(context.Background())
	if err != nil {
		logger.Logger.Warn("save goods msg report failed", zap.Any("report", req), zap.Error(err))
		return err
	}

	return nil
}
