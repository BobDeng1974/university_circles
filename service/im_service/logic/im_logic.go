package logic

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"strconv"
	"university_circles/service/im_service/databases/es"
	pb "university_circles/service/im_service/pb/im"
	"university_circles/service/im_service/utils"
	"university_circles/service/im_service/utils/common"
	"university_circles/service/im_service/utils/errcode"

	"github.com/volatiletech/sqlboiler/boil"
	"go.uber.org/zap"
	im "university_circles/service/im_service/models"
	myRedis "university_circles/service/user_service/databases/redis"
	//"university_circles/service/im_service/databases/es"
	"university_circles/service/im_service/utils/logger"

	"university_circles/service/im_service/databases/mysql"
	//pb "university_circles/service/im_service/pb/im"
)

const (
	USERREPORT    = "user_report"
	HOMEMSGREPORT = "home_msg_report"
	GOODSREPORT   = "goods_report"

	POSTPEERMESSAGEURL = "http://127.0.0.1:6666/post_peer_message?"
	POSTGroupMESSAGEURL = "http://127.0.0.1:6666/post_group_message?"
)

type ImLogic struct{}

//func (il *ImLogic) SaveUserReportToES(req *pb.ReportReq) (err error) {
//	esClient, err := es.NewElasticSearch()
//	if err != nil {
//		logger.Logger.Warn("New ElasticSearch failed", zap.Error(err))
//	}
//
//	commId := uuid.NewV4().String()
//	_, err = esClient.Index().
//		Index(USERREPORT).
//		Id(commId).
//		BodyJson(req).
//		Do(context.Background())
//	if err != nil {
//		logger.Logger.Warn("save user report failed", zap.Any("report", req), zap.Error(err))
//		return err
//	}
//
//	return nil
//}



func (il *ImLogic) SendPeerMsg(appid, sender, receiver int64, chatType, method, msgType, content, token, status, username string) (err error) {
	apiUrl := POSTPEERMESSAGEURL +
		"appid=" + strconv.FormatInt(appid, 10) +
		"&sender=" + strconv.FormatInt(sender, 10) +
		"&receiver=" + strconv.FormatInt(receiver, 10)
	header := make(map[string]string)

	data := "{ \"message_type\": \""+msgType+"\", \"message\": \""+content+"\", " +
		"\"type\": \""+chatType+"\", \"method\": \""+method+"\", \"status\": \""+status+"\", \"sender_name\": \""+username+"\"}"
	header["Content-Type"] = "application/json; charset=UTF-8"
	header["Authorization"] = "Bearer " + token

	var resp map[string]interface{}

	resp, err = utils.DoRequest("POST", apiUrl, header, data)
	fmt.Println("DoRequest SendPeerMsg", resp, header, data, err)
	if err.Error() == errcode.ErrHttpResponse.Error() {
		return nil
	}
	if err != nil {
		logger.Logger.Warn("SendPeerMsg failed", zap.Any("body", data), zap.Error(err))
		return err
	}

	return nil
}

func (il *ImLogic) SendGroupMsg(appid, sender, receiver int64, chatType, method, msgType, content, token, status, username string) (err error) {
	apiUrl := POSTGroupMESSAGEURL +
		"appid=" + strconv.FormatInt(appid, 10) +
		"&sender=" + strconv.FormatInt(sender, 10) +
		"&receiver=" + strconv.FormatInt(receiver, 10)
	header := make(map[string]string)

	data := "{ \"message_type\": \""+msgType+"\", \"message\": \""+content+"\", " +
		"\"type\": \""+chatType+"\", \"method\": \""+method+"\", \"status\": \""+status+"\", \"sender_name\": \""+username+"\"}"
	header["Content-Type"] = "application/json; charset=UTF-8"
	header["Authorization"] = "Bearer " + token

	var resp map[string]interface{}

	resp, err = utils.DoRequest("POST", apiUrl, header, data)
	fmt.Println("DoRequest SendGroupMsg", resp, header, data, err)
	if err.Error() == errcode.ErrHttpResponse.Error() {
		return nil
	}
	if err != nil {
		logger.Logger.Warn("SendGroupMsg failed", zap.Any("body", data), zap.Error(err))
		return err
	}

	return nil
}

func (il *ImLogic) AddFriend(friend *im.Friend) (err error) {
	db, err := mysql.NewMySQL("db_university_circles")
	if err != nil {
		logger.Logger.Warn("new db_university_circles mysql failed", zap.Error(err))
		return
	}
	defer db.Close()

	if err = friend.Insert(context.Background(), db, boil.Infer()); err != nil {
		logger.Logger.Warn("AddFriend failed", zap.Any("friend", friend), zap.Error(err))
		return
	}
	fmt.Println("friend logic", friend)

	return nil
}

func (il *ImLogic) DelFriend(friend *im.Friend) (err error) {
	var exec boil.ContextExecutor
	if _, err = friend.Delete(context.Background(), exec); err != nil {
		logger.Logger.Warn("delete friend failed", zap.Any("friend", friend), zap.Error(err))
		return
	}
	fmt.Println("friend logic", friend)

	return nil
}

func (il *ImLogic) UpdateFriendRemark(friend *im.Friend) (err error) {
	db, err := mysql.NewMySQL("db_university_circles")
	if err != nil {
		logger.Logger.Warn("new db_university_circles mysql failed", zap.Error(err))
		return
	}
	defer db.Close()

	if _, err = friend.Update(context.Background(), db, boil.Infer()); err != nil {
		logger.Logger.Warn("update friend remark failed", zap.Any("friend", friend), zap.Error(err))
		return
	}

	return nil
}

// 获取一个好友
func (il *ImLogic) GetFriendById(uid, friendUid, appid int64) (friend *im.Friend, err error) {
	db, err := mysql.NewMySQL("db_university_circles")
	if err != nil {
		logger.Logger.Warn("new db_university_circles mysql failed", zap.Error(err))
		return
	}
	defer db.Close()

	qmWhere := qm.Where("uid = ? and friend_uid = ? and appid = ?", uid, friendUid, appid)
	friend, err = im.Friends(qmWhere).One(context.Background(), db)
	if errors.Cause(err) == sql.ErrNoRows {
		return friend, nil
	} else if err != nil {
		logger.Logger.Warn("GetFriendById failed", zap.Any("uid", uid), zap.Any("friend_uid", friendUid), zap.Any("appid", appid), zap.Error(err))
		return
	}

	return
}

// 获取好友列表
func (il *ImLogic) GetAllFriend(uid, appid int64) (friends []*im.Friend, err error) {
	db, err := mysql.NewMySQL("db_university_circles")
	if err != nil {
		logger.Logger.Warn("new db_university_circles mysql failed", zap.Error(err))
		return
	}
	defer db.Close()

	qmWhere := qm.Where("uid = ? and appid = ?", uid, appid)
	friends, err = im.Friends(qmWhere).All(context.Background(), db)
	if errors.Cause(err) == sql.ErrNoRows {
		return make([]*im.Friend, 0), nil
	} else if err != nil {
		logger.Logger.Warn("GetAllFriend failed", zap.Any("uid", uid), zap.Any("appid", appid), zap.Error(err))
		return
	}

	return
}

func (il *ImLogic) FindOneStudentByPhone(phone string) (student *im.UStudent, err error) {
	db, err := mysql.NewMySQL("db_university_circles")
	if err != nil {
		logger.Logger.Warn("new db_university_circles mysql failed", zap.Error(err))
		return
	}
	defer db.Close()

	qmWhere := qm.Where("deleted = 0 and phone = ?", phone)
	student, err = im.UStudents(qmWhere).One(context.Background(), db)
	if errors.Cause(err) == sql.ErrNoRows {
		return student, nil
	} else if err != nil {
		logger.Logger.Warn("select one student info by phone failed", zap.Any("phone", phone), zap.Error(err))
		return
	}
	return
}

func (il *ImLogic) FindOneStudentByMid(mid int64) (student *im.UStudent, err error) {
	db, err := mysql.NewMySQL("db_university_circles")
	if err != nil {
		logger.Logger.Warn("new db_university_circles mysql failed", zap.Error(err))
		return
	}
	defer db.Close()

	qmWhere := qm.Where("deleted = 0 and mid = ?", mid)
	student, err = im.UStudents(qmWhere).One(context.Background(), db)
	if errors.Cause(err) == sql.ErrNoRows {
		return student, nil
	} else if err != nil {
		logger.Logger.Warn("select one student info by mid failed", zap.Any("mid", mid), zap.Error(err))
		return
	}
	return
}

func (il *ImLogic) FindOneTeacherByMid(mid int64) (teacher *im.UTeacher, err error) {
	db, err := mysql.NewMySQL("db_university_circles")
	if err != nil {
		logger.Logger.Warn("new db_university_circles mysql failed", zap.Error(err))
		return
	}
	defer db.Close()

	qmWhere := qm.Where("deleted = 0 and mid = ?", mid)
	teacher, err = im.UTeachers(qmWhere).One(context.Background(), db)
	if errors.Cause(err) == sql.ErrNoRows {
		return teacher, nil
	} else if err != nil {
		logger.Logger.Warn("select one teacher info by mid failed", zap.Any("mid", mid), zap.Error(err))
		return
	}
	return
}

func (il *ImLogic) GetEsUserInfo(uid string) (esUserInfo *pb.EsUserInfo, err error) {
	esClient, err := es.NewElasticSearch()
	if err != nil {
		logger.Logger.Warn("New ElasticSearch failed", zap.Error(err))
	}

	// 查询动态里的用户信息
	getUser, err := esClient.Get().Index(common.HOMEPUBLISHMSGUSER).Id(uid).Do(context.Background())
	if err != nil {
		logger.Logger.Warn("get es user byte convert string failed", zap.Any("uid", uid), zap.Error(err))
	}

	sourceUser, err := getUser.Source.MarshalJSON()
	if err != nil {
		logger.Logger.Warn("get es user byte convert string failed", zap.Any("uid", uid), zap.Error(err))
	}

	fmt.Println("sourceUser", string(sourceUser))

	esUserInfo = &pb.EsUserInfo{}
	if err = json.Unmarshal(sourceUser, esUserInfo); err != nil {
		logger.Logger.Warn("json Unmarshal es user failed", zap.Any("esUserInfo", string(sourceUser)), zap.Error(err))
	}

	return
}

// 添加好友到黑名单
func (il *ImLogic) AddFriendBlackList(blacklist *im.Blacklist) (err error) {
	db, err := mysql.NewMySQL("db_university_circles")
	if err != nil {
		logger.Logger.Warn("new db_university_circles mysql failed", zap.Error(err))
		return
	}
	defer db.Close()

	if err = blacklist.Insert(context.Background(), db, boil.Infer()); err != nil {
		logger.Logger.Warn("AddFriendBlackList failed", zap.Any("blacklist", blacklist), zap.Error(err))
		return
	}
	fmt.Println("friend logic", blacklist)

	return nil
}

// 获取一个好友是否在黑名单
func (il *ImLogic) GetBlackListById(uid, friendUid, appid int64) (blacklist *im.Blacklist, err error) {
	db, err := mysql.NewMySQL("db_university_circles")
	if err != nil {
		logger.Logger.Warn("new db_university_circles mysql failed", zap.Error(err))
		return
	}
	defer db.Close()

	qmWhere := qm.Where("uid = ? and friend_uid = ? and appid = ?", uid, friendUid, appid)
	blacklist, err = im.Blacklists(qmWhere).One(context.Background(), db)
	if errors.Cause(err) == sql.ErrNoRows {
		return blacklist, nil
	} else if err != nil {
		logger.Logger.Warn("GetBlackListById failed", zap.Any("uid", uid), zap.Any("friend_uid", friendUid), zap.Any("appid", appid), zap.Error(err))
		return
	}

	return
}

func (il *ImLogic) DelFriendBlackList(blacklist *im.Blacklist) (err error) {
	var exec boil.ContextExecutor
	if _, err = blacklist.Delete(context.Background(), exec); err != nil {
		logger.Logger.Warn("DelFriendBlackList failed", zap.Any("blacklist", blacklist), zap.Error(err))
		return
	}
	fmt.Println("friend logic", blacklist)

	return nil
}

// 获取黑名单列表
func (il *ImLogic) GetBlackList(uid, appid int64) (blacklist []*im.Blacklist, err error) {
	db, err := mysql.NewMySQL("db_university_circles")
	if err != nil {
		logger.Logger.Warn("new db_university_circles mysql failed", zap.Error(err))
		return
	}
	defer db.Close()

	qmWhere := qm.Where("uid = ? and appid = ?", uid, appid)
	blacklist, err = im.Blacklists(qmWhere).All(context.Background(), db)
	if errors.Cause(err) == sql.ErrNoRows {
		return blacklist, nil
	} else if err != nil {
		logger.Logger.Warn("GetBlackList failed", zap.Any("uid", uid), zap.Any("appid", appid), zap.Error(err))
		return
	}

	return
}

func (il *ImLogic) CreateGroup(group *im.Group) (err error) {
	db, err := mysql.NewMySQL("db_university_circles")
	if err != nil {
		logger.Logger.Warn("new db_university_circles mysql failed", zap.Error(err))
		return
	}
	defer db.Close()

	if err = group.Insert(context.Background(), db, boil.Infer()); err != nil {
		logger.Logger.Warn("CreateGroup failed", zap.Any("group", group), zap.Error(err))
		return
	}
	fmt.Println("group logic", group)

	return nil
}

// 根据ID获取群组
func (il *ImLogic) GetGroupById(groupId, appid int64) (group *im.Group, err error) {
	db, err := mysql.NewMySQL("db_university_circles")
	if err != nil {
		logger.Logger.Warn("new db_university_circles mysql failed", zap.Error(err))
		return
	}
	defer db.Close()

	qmWhere := qm.Where("id = ? and appid = ?", groupId, appid)
	group, err = im.Groups(qmWhere).One(context.Background(), db)
	if errors.Cause(err) == sql.ErrNoRows {
		return group, nil
	} else if err != nil {
		logger.Logger.Warn("GetGroupById failed", zap.Any("groupId", groupId), zap.Any("appid", appid), zap.Error(err))
		return
	}

	return
}

// 根据名称获取群组
func (il *ImLogic) GetGroupByName(name string, appid int64) (group *im.Group, err error) {
	db, err := mysql.NewMySQL("db_university_circles")
	if err != nil {
		logger.Logger.Warn("new db_university_circles mysql failed", zap.Error(err))
		return
	}
	defer db.Close()

	qmWhere := qm.Where("name = ? and appid = ?", name, appid)
	group, err = im.Groups(qmWhere).One(context.Background(), db)
	if errors.Cause(err) == sql.ErrNoRows {
		return group, nil
	} else if err != nil {
		logger.Logger.Warn("GetGroupByName failed", zap.Any("name", name), zap.Any("appid", appid), zap.Error(err))
		return
	}

	return
}

// 根据ID获取群组
func (il *ImLogic) GetGroupListByUid(uid, appid int64) (group []*im.Group, err error) {
	db, err := mysql.NewMySQL("db_university_circles")
	if err != nil {
		logger.Logger.Warn("new db_university_circles mysql failed", zap.Error(err))
		return
	}
	defer db.Close()

	group, err = im.Groups(
		qm.Select("id", "master", "super", "name", "notice"),
		qm.InnerJoin("group_member g on g.group_id = id"),
		qm.Where("g.uid = ? and appid = ?", uid, appid),
		//qm.AndIn("g.kind in ?", "visa", "mastercard"),
		//qm.Or("email like ?", `%aol.com%`),
		//qm.GroupBy("id", "name"),
		//qm.Having("count(id) > ?", 2),
		//qm.Limit(5),
		//qm.Offset(6),
	).All(context.Background(), db)

	if errors.Cause(err) == sql.ErrNoRows {
		return group, nil
	} else if err != nil {
		logger.Logger.Warn("GetGroupListByUid failed", zap.Any("uid", uid), zap.Any("appid", appid), zap.Error(err))
		return
	}

	return
}

func (il *ImLogic) DelGroup(group *im.Group) (err error) {
	var exec boil.ContextExecutor
	if _, err = group.Delete(context.Background(), exec); err != nil {
		logger.Logger.Warn("delete group failed", zap.Any("group", group), zap.Error(err))
		return
	}
	fmt.Println("group logic", group)

	return nil
}

func (il *ImLogic) DelGroupMember(groupMember *im.GroupMember) (err error) {
	var exec boil.ContextExecutor
	if _, err = groupMember.Delete(context.Background(), exec); err != nil {
		logger.Logger.Warn("delete group member failed", zap.Any("groupMember", groupMember), zap.Error(err))
		return
	}

	return nil
}

// 获取群组成员
func (il *ImLogic) GetGroupMemberList(groupId, appid int64) (groupMember []*im.GroupMember, err error) {
	db, err := mysql.NewMySQL("db_university_circles")
	if err != nil {
		logger.Logger.Warn("new db_university_circles mysql failed", zap.Error(err))
		return
	}
	defer db.Close()

	qmWhere := qm.Where("group_id = ? and appid = ?", groupId, appid)
	groupMember, err = im.GroupMembers(qmWhere).All(context.Background(), db)
	if errors.Cause(err) == sql.ErrNoRows {
		return groupMember, nil
	} else if err != nil {
		logger.Logger.Warn("GetGroupMemberList failed", zap.Any("group_id", groupId), zap.Any("appid", appid), zap.Error(err))
		return
	}

	return
}

// 获取群组成员人数
func (il *ImLogic) GetGroupMemberNum(groupId, appid int64) (num int64, err error) {
	db, err := mysql.NewMySQL("db_university_circles")
	if err != nil {
		logger.Logger.Warn("new db_university_circles mysql failed", zap.Error(err))
		return
	}
	defer db.Close()

	qmWhere := qm.Where("group_id = ? and appid = ?", groupId, appid)
	num, err = im.GroupMembers(qmWhere).Count(context.Background(), db)
	if errors.Cause(err) == sql.ErrNoRows {
		return num, nil
	} else if err != nil {
		logger.Logger.Warn("GetGroupMemberNum failed", zap.Any("group_id", groupId), zap.Any("appid", appid), zap.Error(err))
		return
	}

	return
}


// 判断是否群组成员
func (il *ImLogic) GetGroupMember(uid, groupId int64) (groupMember *im.GroupMember, err error) {
	db, err := mysql.NewMySQL("db_university_circles")
	if err != nil {
		logger.Logger.Warn("new db_university_circles mysql failed", zap.Error(err))
		return
	}
	defer db.Close()

	qmWhere := qm.Where("uid = ? and group_id = ?", uid, groupId)
	groupMember, err = im.GroupMembers(qmWhere).One(context.Background(), db)
	if errors.Cause(err) == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		logger.Logger.Warn("GetBlackList failed", zap.Any("group_id", groupId), zap.Error(err))
		return
	}

	return
}

func (il *ImLogic) JoinGroup(member *im.GroupMember) (err error) {
	db, err := mysql.NewMySQL("db_university_circles")
	if err != nil {
		logger.Logger.Warn("new db_university_circles mysql failed", zap.Error(err))
		return
	}
	defer db.Close()

	if err = member.Insert(context.Background(), db, boil.Infer()); err != nil {
		logger.Logger.Warn("join group failed", zap.Any("member", member), zap.Error(err))
		return
	}
	fmt.Println("group logic", member)

	return nil
}

func (il *ImLogic) GetUserUnReadCount(appid, uid int64) (unReadCount int64, err error) {
	rd := myRedis.DefaultRedisPool.Get()
	defer rd.Close()

	cacheKey := fmt.Sprintf("users_%d_%d", appid, uid)
	fmt.Println("GetUserUnReadCount cacheKey",cacheKey)
	unReadNum, err := redis.String(rd.Do("HGet", cacheKey, "unread"))
	fmt.Println("GetUserUnReadCount unReadNum",unReadNum, err)
	if err != nil {
		logger.Logger.Warn("get user unread count failed", zap.Any("appid", appid), zap.Any("uid", uid), zap.Any("cacheKey", cacheKey), zap.Error(err))
		return
	}
	
	unReadCount, err = strconv.ParseInt(unReadNum, 10, 64)
	if err != nil {
		logger.Logger.Warn("get user unread count failed", zap.Any("appid", appid), zap.Any("uid", uid), zap.Any("cacheKey", cacheKey), zap.Error(err))
		return
	}

	return
}

func (il *ImLogic) UpdateGroup(group *im.Group) (err error) {
	db, err := mysql.NewMySQL("db_university_circles")
	if err != nil {
		logger.Logger.Warn("new db_university_circles mysql failed", zap.Error(err))
		return
	}
	defer db.Close()

	if _, err = group.Update(context.Background(), db, boil.Infer()); err != nil {
		logger.Logger.Warn("update group avatar failed", zap.Any("group", group), zap.Error(err))
		return
	}

	return nil
}

