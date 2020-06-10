package logic

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/olivere/elastic/v7"
	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"net/url"
	"strconv"
	"time"
	"university_circles/service/user_service/databases/es"
	"university_circles/service/user_service/databases/mysql"
	"university_circles/service/user_service/models"
	"university_circles/service/user_service/utils"
	"university_circles/service/user_service/utils/errcode"
	"university_circles/service/user_service/utils/logger"

	myRedis "university_circles/service/user_service/databases/redis"
	"university_circles/service/user_service/utils/common"

	pb "university_circles/service/user_service/pb/user"

	"go.uber.org/zap"

	"github.com/volatiletech/sqlboiler/queries/qm"
)

type UserLogic struct{}

// 根据学号获取学生信息
func (u *UserLogic) FindOneStudentByNo(stuNo string, universityId int) (student *user.UStudent, err error) {
	db, err := mysql.NewMySQL("db_university_circles")
	if err != nil {
		logger.Logger.Warn("new db_university_circles mysql failed", zap.Error(err))
		return
	}
	defer db.Close()

	qmWhere := qm.Where("deleted = 0 and stu_no = ? and university_id = ?", stuNo, universityId)
	student, err = user.UStudents(qmWhere).One(context.Background(), db)
	if errors.Cause(err) == sql.ErrNoRows {
		return student, nil
	} else if err != nil {
		logger.Logger.Warn("select one student info by No failed", zap.Any("stu_no", stuNo), zap.Any("university_id", universityId), zap.Error(err))
		return
	}

	return
}

// 根据身份证号获取学生信息
func (u *UserLogic) FindOneStudentByIDCardNumber(idCardNumber string) (student *user.UStudent, err error) {
	db, err := mysql.NewMySQL("db_university_circles")
	if err != nil {
		logger.Logger.Warn("new db_university_circles mysql failed", zap.Error(err))
		return
	}
	defer db.Close()

	qmWhere := qm.Where("deleted = 0 and id_card_number = ?", idCardNumber)
	student, err = user.UStudents(qmWhere).One(context.Background(), db)
	if errors.Cause(err) == sql.ErrNoRows {
		return student, nil
	} else if err != nil {
		logger.Logger.Warn("select one student info by id_card_number failed", zap.Any("id_card_number", idCardNumber), zap.Error(err))
		return
	}

	return
}

// 根据邮箱获取学生信息
func (u *UserLogic) FindOneStudentByEmail(email string) (student *user.UStudent, err error) {
	db, err := mysql.NewMySQL("db_university_circles")
	if err != nil {
		logger.Logger.Warn("new db_university_circles mysql failed", zap.Error(err))
		return
	}
	defer db.Close()

	qmWhere := qm.Where("deleted = 0 and email = ? ", email)
	student, err = user.UStudents(qmWhere).One(context.Background(), db)
	if errors.Cause(err) == sql.ErrNoRows {
		return student, nil
	} else if err != nil {
		logger.Logger.Warn("select one student info by email failed", zap.Any("email", email), zap.Error(err))
		return
	}

	return student, nil
}

// 根据用户名获取学生信息
func (u *UserLogic) FindOneStudentByScreenName(nickname string) (student *user.UStudent, err error) {
	db, err := mysql.NewMySQL("db_university_circles")
	if err != nil {
		logger.Logger.Warn("new db_university_circles mysql failed", zap.Error(err))
		return
	}
	defer db.Close()

	qmWhere := qm.Where("deleted = 0 and screen_name = ?", nickname)
	student, err = user.UStudents(qmWhere).One(context.Background(), db)
	if errors.Cause(err) == sql.ErrNoRows {
		return student, nil
	} else if err != nil {
		logger.Logger.Warn("select one student info by screen_name failed", zap.Any("username", nickname), zap.Error(err))
		return
	}
	return
}

// 根据id获取学生信息
func (u *UserLogic) FindOneStudentById(uid string) (student *user.UStudent, err error) {
	db, err := mysql.NewMySQL("db_university_circles")
	if err != nil {
		logger.Logger.Warn("new db_university_circles mysql failed", zap.Error(err))
		return
	}
	defer db.Close()

	qmWhere := qm.Where("deleted = 0 and uid = ?", uid)
	fmt.Println(qmWhere)
	student, err = user.UStudents(qmWhere).One(context.Background(), db)
	fmt.Println("query student, err ", student, err)
	if errors.Cause(err) == sql.ErrNoRows {
		return student, nil
	} else if err != nil {
		logger.Logger.Warn("select one student info by id failed", zap.Any("uid", uid), zap.Error(err))
		return
	}
	return
}

func (u *UserLogic) FindOneStudentByPhone(phone string) (student *user.UStudent, err error) {
	db, err := mysql.NewMySQL("db_university_circles")
	if err != nil {
		logger.Logger.Warn("new db_university_circles mysql failed", zap.Error(err))
		return
	}
	defer db.Close()

	qmWhere := qm.Where("deleted = 0 and phone = ?", phone)
	student, err = user.UStudents(qmWhere).One(context.Background(), db)
	if errors.Cause(err) == sql.ErrNoRows {
		return student, nil
	} else if err != nil {
		logger.Logger.Warn("select one student info by phone failed", zap.Any("phone", phone), zap.Error(err))
		return
	}
	return
}

func (u *UserLogic) InsertStudent(student *user.UStudent) (err error) {
	db, err := mysql.NewMySQL("db_university_circles")
	if err != nil {
		logger.Logger.Warn("new db_university_circles mysql failed", zap.Error(err))
		return
	}
	defer db.Close()

	if err = student.Insert(context.Background(), db, boil.Infer()); err != nil {
		logger.Logger.Warn("insert student failed", zap.Any("student", student), zap.Error(err))
		return
	}
	return
}

func (u *UserLogic) DelStudent(student *user.UStudent) (err error) {
	db, err := mysql.NewMySQL("db_university_circles")
	if err != nil {
		logger.Logger.Warn("new db_university_circles mysql failed", zap.Error(err))
		return
	}
	defer db.Close()

	var exec boil.ContextExecutor
	if _, err = student.Delete(context.Background(), exec); err != nil {
		logger.Logger.Warn("del student failed", zap.Any("student", student), zap.Error(err))
		return
	}
	return
}

func (u *UserLogic) InsertVerifyCode(verifyCode *user.UVerifyCode) (err error) {
	db, err := mysql.NewMySQL("db_university_circles")
	if err != nil {
		logger.Logger.Warn("new db_university_circles mysql failed", zap.Error(err))
		return
	}
	defer db.Close()

	if err = verifyCode.Insert(context.Background(), db, boil.Infer()); err != nil {
		logger.Logger.Warn("insert verify code failed", zap.Any("verifyCode", verifyCode), zap.Error(err))
		return
	}
	return
}

func (u *UserLogic) UpdateStudentInfo(student *user.UStudent) (err error) {
	db, err := mysql.NewMySQL("db_university_circles")
	if err != nil {
		logger.Logger.Warn("new db_university_circles mysql failed", zap.Error(err))
		return
	}
	defer db.Close()

	if _, err = student.Update(context.Background(), db, boil.Infer()); err != nil {
		logger.Logger.Warn("update verify code failed", zap.Any("student", student), zap.Error(err))
		return
	}
	return
}

// 更新某个用户Es的头像
func (u *UserLogic) UpdateUserAvatarInfo(student *user.UStudent) (err error) {
	esClient, err := es.NewElasticSearch()
	if err != nil {
		logger.Logger.Warn("UpdateUserAvatarInfo New ElasticSearch failed", zap.Error(err))
		return
	}

	var avatar *pb.Avatar
	if student.Avatar != "" {
		avatar = &pb.Avatar{
			ThumbnailUrl: student.Avatar + "?x-oss-process=image/resize,m_fixed,h_120,w_120",
			SmallPicUrl:  student.Avatar + "?x-oss-process=image/resize,m_fixed,h_300,w_300",
			PicUrl:       student.Avatar + "?x-oss-process=image/resize,m_fixed,h_800,w_800",
		}
	}

	// 更新动态里的用户头像信息到es
	_, err = esClient.Update().
		Index(common.HOMEPUBLISHMSGUSER).
		Id(student.UID).
		Doc(map[string]interface{}{"avatarImage": avatar}).
		Do(context.Background())

	fmt.Println("avatar", avatar, err)
	if err != nil {
		logger.Logger.Warn("insert Msg User Info To Es failed", zap.Any("avatar", avatar), zap.Error(err))
		return
	}

	return
}

func (u *UserLogic) FindVerifyCodeByPhone(phone string) (verifyCode *user.UVerifyCode, err error) {
	db, err := mysql.NewMySQL("db_university_circles")
	if err != nil {
		logger.Logger.Warn("new db_university_circles mysql failed", zap.Error(err))
		return
	}
	defer db.Close()

	qmWhere := qm.Where(" phone = ? ", phone)
	verifyCode, err = user.UVerifyCodes(qmWhere, qm.OrderBy(" updated desc ")).One(context.Background(), db)
	if errors.Cause(err) == sql.ErrNoRows {
		return verifyCode, nil
	} else if err != nil {
		logger.Logger.Warn("select one verifyCode by phone failed", zap.Any("phone", phone), zap.Error(err))
		return
	}
	return
}

func (u *UserLogic) FindUniversityByName(name string) (university *user.University, err error) {
	db, err := mysql.NewMySQL("db_university_circles")
	if err != nil {
		logger.Logger.Warn("new db_university_circles mysql failed", zap.Error(err))
		return
	}
	defer db.Close()

	qmWhere := qm.Where(" deleted = 0 and name = ?", name)
	university, err = user.Universitys(qmWhere).One(context.Background(), db)
	if errors.Cause(err) == sql.ErrNoRows {
		return university, nil
	} else if err != nil {
		logger.Logger.Warn("select one university by id failed", zap.Any("university", name), zap.Error(err))
		return
	}
	return
}

func (u *UserLogic) FindUniversityByLikeName(name string) (universitys []*user.University, err error) {
	db, err := mysql.NewMySQL("db_university_circles")
	if err != nil {
		logger.Logger.Warn("new db_university_circles mysql failed", zap.Error(err))
		return
	}
	defer db.Close()

	// %放后面会使用到索引
	qmWhere := qm.Where(" deleted = 0 and name like ?%", name)
	universitys, err = user.Universitys(qmWhere).All(context.Background(), db)
	if errors.Cause(err) == sql.ErrNoRows {
		return universitys, nil
	} else if err != nil {
		logger.Logger.Warn("select one university by id failed", zap.Any("university", name), zap.Error(err))
		return
	}
	return
}

func (u *UserLogic) FindUniversityById(id int) (university *user.University, err error) {
	db, err := mysql.NewMySQL("db_university_circles")
	if err != nil {
		logger.Logger.Warn("new db_university_circles mysql failed", zap.Error(err))
		return
	}
	defer db.Close()

	qmWhere := qm.Where("deleted = 0 and id = ? ", id)
	university, err = user.Universitys(qmWhere).One(context.Background(), db)
	if errors.Cause(err) == sql.ErrNoRows {
		return university, nil
	} else if err != nil {
		logger.Logger.Warn("select one university by id failed", zap.Any("id", id), zap.Error(err))
		return
	}
	return
}

func (u *UserLogic) FindUniversityList() (universityList []*user.University, err error) {
	db, err := mysql.NewMySQL("db_university_circles")
	if err != nil {
		logger.Logger.Warn("new db_university_circles mysql failed", zap.Error(err))
		return
	}
	defer db.Close()

	qmWhere := qm.Where(" deleted = 0 ")
	universityList, err = user.Universitys(qmWhere).All(context.Background(), db)
	if errors.Cause(err) == sql.ErrNoRows {
		return universityList, nil
	} else if err != nil {
		logger.Logger.Warn("select all university failed", zap.Error(err))
		return
	}
	return
}

func (u *UserLogic) FindCollegeById(id int) (universityCollege *user.UniversityCollege, err error) {
	db, err := mysql.NewMySQL("db_university_circles")
	if err != nil {
		logger.Logger.Warn("new db_university_circles mysql failed", zap.Error(err))
		return
	}
	defer db.Close()

	qmWhere := qm.Where(" deleted = 0 and id = ? ", id)
	universityCollege, err = user.UniversityColleges(qmWhere).One(context.Background(), db)
	if errors.Cause(err) == sql.ErrNoRows {
		return universityCollege, nil
	} else if err != nil {
		logger.Logger.Warn("select one universityCollege by id failed", zap.Any("id", id), zap.Error(err))
		return
	}
	return
}

func (u *UserLogic) FindCollegeListByUniId(universityId string) (collegeList []*user.UniversityCollege, err error) {
	db, err := mysql.NewMySQL("db_university_circles")
	if err != nil {
		logger.Logger.Warn("new db_university_circles mysql failed", zap.Error(err))
		return
	}
	defer db.Close()

	qmWhere := qm.Where(" deleted = 0 and university_id = ?", universityId)
	collegeList, err = user.UniversityColleges(qmWhere).All(context.Background(), db)
	if errors.Cause(err) == sql.ErrNoRows {
		return collegeList, nil
	} else if err != nil {
		logger.Logger.Warn("select one universityCollege by id failed", zap.Any("university_id", universityId), zap.Error(err))
		return
	}
	return
}

func (u *UserLogic) FindProfessionByName(name string) (profession *user.UniversityCollegeProfession, err error) {
	db, err := mysql.NewMySQL("db_university_circles")
	if err != nil {
		logger.Logger.Warn("new db_university_circles mysql failed", zap.Error(err))
		return
	}
	defer db.Close()

	qmWhere := qm.Where(" deleted = 0 and name = ? ", name)
	profession, err = user.UniversityCollegeProfessions(qmWhere).One(context.Background(), db)
	if errors.Cause(err) == sql.ErrNoRows {
		return profession, nil
	} else if err != nil {
		logger.Logger.Warn("select one universityCollegeProfession by id failed", zap.Any("name", name), zap.Error(err))
		return
	}
	return
}

func (u *UserLogic) FindProfessionById(id int) (collegeScience *user.UniversityCollegeProfession, err error) {
	db, err := mysql.NewMySQL("db_university_circles")
	if err != nil {
		logger.Logger.Warn("new db_university_circles mysql failed", zap.Error(err))
		return
	}
	defer db.Close()

	qmWhere := qm.Where(" deleted = 0 and id = ? ", id)
	collegeScience, err = user.UniversityCollegeProfessions(qmWhere).One(context.Background(), db)
	if errors.Cause(err) == sql.ErrNoRows {
		return collegeScience, nil
	} else if err != nil {
		logger.Logger.Warn("select one universityCollegeProfession by id failed", zap.Any("id", id), zap.Error(err))
		return
	}
	return
}

func (u *UserLogic) FindProfessionListByCollegeId(universityId, collegeId string) (scienceList []*user.UniversityCollegeProfession, err error) {
	db, err := mysql.NewMySQL("db_university_circles")
	if err != nil {
		logger.Logger.Warn("new db_university_circles mysql failed", zap.Error(err))
		return
	}
	defer db.Close()

	qmWhere := qm.Where(" deleted = 0 and university_id = ? and college_id = ? ", universityId, collegeId)
	scienceList, err = user.UniversityCollegeProfessions(qmWhere).All(context.Background(), db)
	if errors.Cause(err) == sql.ErrNoRows {
		return scienceList, nil
	} else if err != nil {
		logger.Logger.Warn("select one universityCollegeScience by id failed", zap.Any("university_id", universityId), zap.Any("university_id", universityId), zap.Error(err))
		return
	}
	return
}

func (u *UserLogic) FindClassById(id int) (universityCollegeScienceClass *user.UniversityCollegeProfessionClass, err error) {
	db, err := mysql.NewMySQL("db_university_circles")
	if err != nil {
		logger.Logger.Warn("new db_university_circles mysql failed", zap.Error(err))
		return
	}
	defer db.Close()

	qmWhere := qm.Where(" deleted = 0 and id = ? ", id)
	universityCollegeScienceClass, err = user.UniversityCollegeProfessionClasses(qmWhere).One(context.Background(), db)
	if errors.Cause(err) == sql.ErrNoRows {
		return universityCollegeScienceClass, nil
	} else if err != nil {
		logger.Logger.Warn("select one universityCollegeScienceClass by id failed", zap.Any("id", id), zap.Error(err))
		return
	}
	return
}

func (u *UserLogic) FindClassListByProfessionId(universityId, collegeId, scienceId string) (classList []*user.UniversityCollegeProfessionClass, err error) {
	db, err := mysql.NewMySQL("db_university_circles")
	if err != nil {
		logger.Logger.Warn("new db_university_circles mysql failed", zap.Error(err))
		return
	}
	defer db.Close()

	qmWhere := qm.Where("  deleted = 0 and university_id = ? and college_id = ?  and profession_id = ? ", universityId, collegeId, scienceId)
	classList, err = user.UniversityCollegeProfessionClasses(qmWhere).All(context.Background(), db)
	if errors.Cause(err) == sql.ErrNoRows {
		return classList, nil
	} else if err != nil {
		logger.Logger.Warn("select one universityCollegeScienceClass by id failed", zap.Any("university_id", universityId), zap.Any("college_id", collegeId), zap.Any("profession_id", scienceId), zap.Error(err))
		return
	}
	return
}

func (u *UserLogic) SavePhoneCode(phone, code string) (err error) {
	rd := myRedis.DefaultRedisPool.Get()
	defer rd.Close()

	cacheKey := common.LOGINPHONECODECACHEPREFIX + phone
	// 设置手机验证码5分钟有效期
	_, err = rd.Do("Set", cacheKey, code, "EX", 300)
	if err != nil {
		logger.Logger.Warn("set phone code into redis failed", zap.Any("cacheKey", cacheKey), zap.Any("code", code), zap.Error(err))
		return
	}

	return nil
}

func (u *UserLogic) GetPhoneCode(phone string) (code string, err error) {
	rd := myRedis.DefaultRedisPool.Get()
	defer rd.Close()

	cacheKey := common.LOGINPHONECODECACHEPREFIX + phone
	code, err = redis.String(rd.Do("Get", cacheKey))
	fmt.Println("redis验证码：：", code, err)
	if err != nil {
		if err.Error() != "redigo: nil returned" {
			logger.Logger.Warn("get phone code into redis failed", zap.Any("cacheKey", cacheKey), zap.Any("code", code), zap.Error(err))
			return
		}
	}

	return code, nil
}

func (u *UserLogic) DelPhoneCode(phone string) (err error) {
	rd := myRedis.DefaultRedisPool.Get()
	defer rd.Close()

	cacheKey := common.LOGINTOKENCACHEPREFIX + phone
	_, err = rd.Do("Del", cacheKey)
	if err != nil {
		if err.Error() != "redigo: nil returned" {
			logger.Logger.Warn("del phone code into redis failed", zap.Any("cacheKey", cacheKey), zap.Error(err))
			return
		}
	}

	return nil
}

func (u *UserLogic) DelLoginSession(uid string) (err error) {
	rd := myRedis.DefaultRedisPool.Get()
	defer rd.Close()

	cacheKey := common.LOGINTOKENCACHEPREFIX + uid
	// 设置缓存
	_, err = rd.Do("Del", cacheKey)
	if err != nil {
		logger.Logger.Warn("set user token into redis failed", zap.Any("uid", uid), zap.Any("key", cacheKey), zap.Error(err))
		return
	}

	return nil

}

func (u *UserLogic) SaveLoginSession(student *user.UStudent) (err error) {
	rd := myRedis.DefaultRedisPool.Get()
	defer rd.Close()

	id := strconv.FormatUint(uint64(student.ID), 10)
	cacheKey := common.STUDENTLOGINSESSIONPREKEY + id

	stu, err := json.Marshal(&student)
	// 设置用户信息
	_, err = rd.Do("Set", cacheKey, stu)
	if err != nil {
		logger.Logger.Warn("set student into redis failed", zap.Any("cacheKey", cacheKey), zap.Any("student", student), zap.Error(err))
		return
	}
	return nil
}

func (u *UserLogic) GetUserInfoFromEs(uid string) (esUserInfo *pb.EsUserInfo, err error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("GetUserInfoFromEs recover", err)
		}
	}()
	esClient, err := es.NewElasticSearch()
	if err != nil {
		logger.Logger.Warn("New ElasticSearch failed", zap.Error(err))
		return nil, err
	}

	// 查询动态里的用户信息
	getUser, err := esClient.Get().Index(common.HOMEPUBLISHMSGUSER).Id(uid).Do(context.Background())
	if err != nil {
		logger.Logger.Warn("home publish msg user byte convert string failed", zap.Any("uid", uid), zap.Error(err))
		return nil, err
	}

	sourceUser, err := getUser.Source.MarshalJSON()
	if err != nil {
		logger.Logger.Warn("home publish msg user byte convert string failed", zap.Any("uid", uid), zap.Error(err))
		return nil, err
	}

	fmt.Println("sourceUser", string(sourceUser))

	esUserInfo = &pb.EsUserInfo{}
	if err = json.Unmarshal(sourceUser, esUserInfo); err != nil {
		logger.Logger.Warn("json Unmarshal home publish msg user failed", zap.Any("esUserInfo", string(sourceUser)), zap.Error(err))
		return nil, err
	}

	return esUserInfo, nil

}

// 多条件模糊查询用户
func (u *UserLogic) GetUserInfoFromEsByLike(userType int64, reqStr string) (userInfos []*pb.QueryUser, err error) {
	esClient, err := es.NewElasticSearch()
	if err != nil {
		logger.Logger.Warn("New ElasticSearch failed", zap.Error(err))
		return nil, err
	}

	// 查询动态里的用户信息
	phoneQuery := elastic.NewMatchQuery("phone", reqStr)
	emailQuery := elastic.NewMatchQuery("email", reqStr)
	screenNameQuery := elastic.NewMatchQuery("screenName", reqStr)
	//typeQuery := elastic.NewMatchQuery("type", userType)
	userQuery := elastic.NewBoolQuery().Should(phoneQuery, emailQuery, screenNameQuery)

	getUser, err := esClient.Search(common.HOMEPUBLISHMSGUSER).Query(userQuery).Do(context.Background())
	if err != nil {
		logger.Logger.Warn("get user from es by like failed", zap.Any("req string", reqStr), zap.Error(err))
		return nil, err
	}

	if len(getUser.Hits.Hits) > 0 {
		for _, hit := range getUser.Hits.Hits {
			sourceUser, err := hit.Source.MarshalJSON()
			if err != nil {
				logger.Logger.Warn("get user from es by like convert string failed", zap.Any("req string", reqStr), zap.Error(err))
				return nil, err
			}

			esUserInfo := &pb.EsUserInfo{}
			if err = json.Unmarshal(sourceUser, esUserInfo); err != nil {
				logger.Logger.Warn("json Unmarshal get user from es by like failed", zap.Any("esUserInfo", string(sourceUser)), zap.Error(err))
				return nil, err
			}
			queryUser := &pb.QueryUser{
				Uid: esUserInfo.Mid,
				ImageUrl: esUserInfo.AvatarImage.ThumbnailUrl,
				Nickname: esUserInfo.ScreenName,
			}
			userInfos = append(userInfos, queryUser)
		}
	}

	return userInfos, nil

}

func (u *UserLogic) SaveMsgUserInfoToEs(user *pb.EsUserInfo) (err error) {
	esClient, err := es.NewElasticSearch()
	if err != nil {
		fmt.Println(err)
		logger.Logger.Warn("New ElasticSearch failed", zap.Any("user", user), zap.Error(err))
		return
	}

	// 添加动态里的用户信息到es
	_, err = esClient.Index().
		Index(common.HOMEPUBLISHMSGUSER).
		Id(user.Id).
		BodyJson(user).
		Do(context.Background())
	fmt.Println("HOMEPUBLISHMSGUSER", err)
	if err != nil {
		logger.Logger.Warn("insert Msg User Info To Es failed", zap.Any("stuseru", user), zap.Error(err))
		return
	}

	return
}

func (u *UserLogic) UpdateMsgUserInfoToEs(user *pb.EsUserInfo) (err error) {
	esClient, err := es.NewElasticSearch()
	if err != nil {
		logger.Logger.Warn("New ElasticSearch failed", zap.Any("user", user), zap.Error(err))
		return
	}

	// 更新动态里的用户信息到es
	_, err = esClient.Update().
		Index(common.HOMEPUBLISHMSGUSER).
		Id(user.Id).
		Doc(map[string]interface{}{
			"username":      user.ScreenName,
			"screenName":    user.ScreenName,
			"bio":           user.Bio,
			"isVerified":    user.IsVerified,
			"zodiac":        user.Zodiac,
			"university_id": user.UniversityId,
		}).
		Do(context.Background())
	fmt.Println("UpdateMsgUserInfoToEs", user, err)
	if err != nil {
		logger.Logger.Warn("insert Msg User Info To Es failed", zap.Any("user", user), zap.Error(err))
		return
	}

	return
}

func (u *UserLogic) SaveUserFollow(uid, followUid string) (err error) {
	rd := myRedis.DefaultRedisPool.Get()
	defer rd.Close()

	// 用户关注
	followCacheKey := common.USERFOLLOWList + uid
	_, err = rd.Do("zAdd", followCacheKey, time.Now().Local().Unix(), followUid)
	if err != nil {
		logger.Logger.Warn("save user follow into redis failed", zap.Any("followCacheKey", followCacheKey), zap.Any("followUid", followUid), zap.Error(err))
		return
	}

	// 对应被关注用户的关系
	followingCacheKey := common.USERFOLLOWINGList + followUid
	_, err = rd.Do("zAdd", followingCacheKey, time.Now().Local().Unix(), uid)
	if err != nil {
		logger.Logger.Warn("save user following into redis failed", zap.Any("followingCacheKey", followingCacheKey), zap.Any("uid", uid), zap.Error(err))
		return
	}
	return nil
}

func (u *UserLogic) CancelUserFollow(uid, followUid string) (err error) {
	rd := myRedis.DefaultRedisPool.Get()
	defer rd.Close()

	// 用户关注
	followCacheKey := common.USERFOLLOWList + uid
	_, err = rd.Do("zRem", followCacheKey, followUid)
	if err != nil {
		logger.Logger.Warn("save user follow into redis failed", zap.Any("followCacheKey", followCacheKey), zap.Any("followUid", followUid), zap.Error(err))
		return
	}

	// 对应被关注用户的关系
	followingCacheKey := common.USERFOLLOWINGList + followUid
	_, err = rd.Do("zRem", followingCacheKey, uid)
	if err != nil {
		logger.Logger.Warn("save user following into redis failed", zap.Any("followingCacheKey", followingCacheKey), zap.Any("uid", uid), zap.Error(err))
		return
	}
	return nil
}

// 查询用户是否已提交验证信息
func (u *UserLogic) CheckUserIsVerify(phone string) (isVerify int8, err error) {
	db, err := mysql.NewMySQL("db_university_circles")
	if err != nil {
		logger.Logger.Warn("new db_university_circles mysql failed", zap.Error(err))
		return
	}
	defer db.Close()

	qmWhere := qm.Where(" deleted = 0 and phone = ?", phone)
	student, err := user.UStudents(qmWhere).One(context.Background(), db)
	if err != nil {
		logger.Logger.Warn("select user by phone failed", zap.Any("phone", phone), zap.Error(err))
		return 0, err
	}

	if student != nil {
		return student.IsVerified, nil
	} else {
		return 0, nil
	}

}

func (u *UserLogic) CheckUserPhoneAndIdCard(phone, idCard, username string) (code int64, userInfo map[string]interface{}, err error) {
	db, err := mysql.NewMySQL("db_university_circles")
	if err != nil {
		logger.Logger.Warn("new db_university_circles mysql failed", zap.Error(err))
		return
	}
	defer db.Close()

	apiUrl := "https://mobile3elements.shumaidata.com/mobile/verify_real_name"
	header := make(map[string]string)

	body := url.Values{}
	body.Add("idcard", idCard)
	body.Add("mobile", phone)
	body.Add("name", username)

	data := body.Encode()
	header["Authorization"] = "APPCODE 658b00b65e7b40bd86a427349c50ceb4"
	header["Content-Type"] = "application/x-www-form-urlencoded; charset=UTF-8"
	header["Content-Length"] = strconv.Itoa(len(data))

	var resp map[string]interface{}
	//resp["code"] = 0
	//resp["message"] = "成功"
	//resp["result"] = map[string]interface{}{
	//	"address": "广东省云浮市新兴县",
	//	"birthday":"19930608",
	//	"description": "一致",
	//	"idcard": "445321199306082519",
	//	"name": "崔伟栋",
	//	"res":1,
	//	"sex":"男",
	//}
	resp, err = utils.DoRequest("POST", apiUrl, header, data)
	fmt.Println("DoRequest CheckUserPhoneAndIdCard", resp, header, body.Encode(), err)
	if err != nil {
		logger.Logger.Warn("Check User Phone And IdCard failed", zap.Any("body", body.Encode()), zap.Error(err))
		return -1, nil, err
	}

	logger.Logger.Info("Check User Phone And IdCard ", zap.Any("body", body.Encode()), zap.Any("response", resp), zap.Error(err))

	newCode, ok := resp["code"].(string)
	fmt.Printf(" type:%T\n", resp["code"])
	fmt.Println("code, ok  int64", code, ok)
	if !ok {
		return -1, nil, nil
	}
	if newCode == "0" {
		userInfo, ok = resp["result"].(map[string]interface{})
		fmt.Printf("result type:%T\n", resp["result"])
		fmt.Println("userInfo, ok  result", userInfo, ok)
		if !ok {
			return -1, nil, nil
		}
	} else if newCode == "20310" {
		return errcode.ErrUserNameVerifyInvalid.Code, nil, nil
	} else if newCode == "20010" {
		return errcode.ErrUserIdCardVerifyInvalid.Code, nil, nil
	} else if newCode == "20004" {
		return errcode.ErrUserPhoneVerifyInvalid.Code, nil, nil
	} else {
		return -1, nil, nil
	}

	return 0, userInfo, nil
}

// 根据邮箱获取学生信息
func (u *UserLogic) FindOneTeacherByEmail(email string) (teacher *user.UTeacher, err error) {
	db, err := mysql.NewMySQL("db_university_circles")
	if err != nil {
		logger.Logger.Warn("new db_university_circles mysql failed", zap.Error(err))
		return
	}
	defer db.Close()

	qmWhere := qm.Where("deleted = 0 and email = ? ", email)
	teacher, err = user.UTeachers(qmWhere).One(context.Background(), db)
	if errors.Cause(err) == sql.ErrNoRows {
		return teacher, nil
	} else if err != nil {
		logger.Logger.Warn("select one teacher info by email failed", zap.Any("email", email), zap.Error(err))
		return
	}

	return
}

// 根据用户名获取教师信息
func (u *UserLogic) FindOneTeacherByScreenName(nickname string) (teacher *user.UTeacher, err error) {
	db, err := mysql.NewMySQL("db_university_circles")
	if err != nil {
		logger.Logger.Warn("new db_university_circles mysql failed", zap.Error(err))
		return
	}
	defer db.Close()

	qmWhere := qm.Where("deleted = 0 and screen_name = ?", nickname)
	teacher, err = user.UTeachers(qmWhere).One(context.Background(), db)
	if errors.Cause(err) == sql.ErrNoRows {
		return teacher, nil
	} else if err != nil {
		logger.Logger.Warn("select one teacher info by screen_name failed", zap.Any("username", nickname), zap.Error(err))
		return
	}
	return
}

// 根据id获取教师信息
func (u *UserLogic) FindOneTeacherById(uid string) (teacher *user.UTeacher, err error) {
	db, err := mysql.NewMySQL("db_university_circles")
	if err != nil {
		logger.Logger.Warn("new db_university_circles mysql failed", zap.Error(err))
		return
	}
	defer db.Close()

	qmWhere := qm.Where("deleted = 0 and uid = ?", uid)
	fmt.Println(qmWhere)
	teacher, err = user.UTeachers(qmWhere).One(context.Background(), db)
	fmt.Println("query student, err ", teacher, err)
	if errors.Cause(err) == sql.ErrNoRows {
		return teacher, nil
	} else if err != nil {
		logger.Logger.Warn("select one teacher info by id failed", zap.Any("uid", uid), zap.Error(err))
		return
	}
	return
}

func (u *UserLogic) FindOneTeacherByPhone(phone string) (teacher *user.UTeacher, err error) {
	db, err := mysql.NewMySQL("db_university_circles")
	if err != nil {
		logger.Logger.Warn("new db_university_circles mysql failed", zap.Error(err))
		return
	}
	defer db.Close()

	qmWhere := qm.Where("deleted = 0 and phone = ?", phone)
	teacher, err = user.UTeachers(qmWhere).One(context.Background(), db)
	if errors.Cause(err) == sql.ErrNoRows {
		return teacher, nil
	} else if err != nil {
		logger.Logger.Warn("select one teacher info by phone failed", zap.Any("phone", phone), zap.Error(err))
		return
	}
	return
}

func (u *UserLogic) InsertTeacher(teacher *user.UTeacher) (err error) {
	db, err := mysql.NewMySQL("db_university_circles")
	if err != nil {
		logger.Logger.Warn("new db_university_circles mysql failed", zap.Error(err))
		return
	}
	defer db.Close()

	if err = teacher.Insert(context.Background(), db, boil.Infer()); err != nil {
		logger.Logger.Warn("insert teacher failed", zap.Any("teacher", teacher), zap.Error(err))
		return
	}
	return
}

func (u *UserLogic) UpdateTeacherInfo(teacher *user.UTeacher) (err error) {
	db, err := mysql.NewMySQL("db_university_circles")
	if err != nil {
		logger.Logger.Warn("new db_university_circles mysql failed", zap.Error(err))
		return
	}
	defer db.Close()

	if _, err = teacher.Update(context.Background(), db, boil.Infer()); err != nil {
		logger.Logger.Warn("update teacher info failed", zap.Any("teacher", teacher), zap.Error(err))
		return
	}
	return
}

// 根据学号获取教师信息
func (u *UserLogic) FindOneTeacherByNo(teacherNo string, universityId int) (teacher *user.UTeacher, err error) {
	db, err := mysql.NewMySQL("db_university_circles")
	if err != nil {
		logger.Logger.Warn("new db_university_circles mysql failed", zap.Error(err))
		return
	}
	defer db.Close()

	qmWhere := qm.Where("deleted = 0 and teach_no = ? and university_id = ?", teacherNo, universityId)
	teacher, err = user.UTeachers(qmWhere).One(context.Background(), db)
	if errors.Cause(err) == sql.ErrNoRows {
		return teacher, nil
	} else if err != nil {
		logger.Logger.Warn("select one teacher info by No failed", zap.Any("teach_no", teacherNo), zap.Any("university_id", universityId), zap.Error(err))
		return
	}

	return
}

// 根据身份证号获取教师信息
func (u *UserLogic) FindOneTeacherByIDCardNumber(idCardNumber string) (teacher *user.UTeacher, err error) {
	db, err := mysql.NewMySQL("db_university_circles")
	if err != nil {
		logger.Logger.Warn("new db_university_circles mysql failed", zap.Error(err))
		return
	}
	defer db.Close()

	qmWhere := qm.Where("deleted = 0 and id_card_number = ?", idCardNumber)
	teacher, err = user.UTeachers(qmWhere).One(context.Background(), db)
	if errors.Cause(err) == sql.ErrNoRows {
		return teacher, nil
	} else if err != nil {
		logger.Logger.Warn("select one teacher info by id_card_number failed", zap.Any("id_card_number", idCardNumber), zap.Error(err))
		return
	}

	return
}

func (u *UserLogic) DelTeacher(teacher *user.UTeacher) (err error) {
	db, err := mysql.NewMySQL("db_university_circles")
	if err != nil {
		logger.Logger.Warn("new db_university_circles mysql failed", zap.Error(err))
		return
	}
	defer db.Close()

	var exec boil.ContextExecutor
	if _, err = teacher.Delete(context.Background(), exec); err != nil {
		logger.Logger.Warn("del student failed", zap.Any("teacher", teacher), zap.Error(err))
		return
	}
	return
}
