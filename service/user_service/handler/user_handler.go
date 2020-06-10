package handler

import (
	"fmt"
	"github.com/satori/go.uuid"
	"github.com/volatiletech/null"
	"time"

	"university_circles/service/user_service/logic"
	"university_circles/service/user_service/models"
	pb "university_circles/service/user_service/pb/user"

	"university_circles/service/user_service/utils/errcode"

	"university_circles/service/user_service/utils/common"

	"context"
	"go.uber.org/zap"
	"university_circles/service/user_service/utils/logger"
)

type UserHandler struct {
}

func (u *UserHandler) StudentRegister(ctx context.Context, req *pb.UserRegisterReq, resp *pb.Response) (err error) {
	var university *user.University
	var profession *user.UniversityCollegeProfession
	ul := &logic.UserLogic{}
	// 验证码对比
	uVerifyCode, err := ul.GetPhoneCode(req.Phone)
	if err != nil {
		resp.Success = -1
		resp.Msg = "register verify code failed"
		logger.Logger.Warn("register verify code failed", zap.String("phone", req.Phone), zap.Error(err))
		return err
	}

	if uVerifyCode != "" {
		if req.Code != uVerifyCode {
			resp.Success = errcode.ErrVerifyCodeCompareFailed.Code
			resp.Msg = "register user verify code is not equal failed"
			return nil
		}
	} else {
		resp.Success = errcode.ErrVerifyCodeNotExist.Code
		resp.Msg = "register user verify code is not exist"
		return nil
	}

	var student *user.UStudent
	var teacher *user.UTeacher
	if common.VerifyMobileFormat(req.Phone) {
		// 手机号已存在
		student, err = ul.FindOneStudentByPhone(req.Phone)
		if err != nil {
			resp.Success = -1
			resp.Msg = "select one student info by No failed"
			return err
		}
		if student != nil {
			resp.Success = errcode.ErrRegisterPhoneExist.Code
			resp.Msg = "user is exist"
			return nil
		}

		teacher, err = ul.FindOneTeacherByPhone(req.Phone)
		if err != nil {
			resp.Success = -1
			resp.Msg = "select one teacher info by No failed"
			return err
		}
		if teacher != nil {
			resp.Success = errcode.ErrRegisterPhoneExist.Code
			resp.Msg = "user is exist"
			return nil
		}

		// 用户名已经存在
		student, err = ul.FindOneStudentByScreenName(req.Username)
		if err != nil {
			resp.Success = -1
			resp.Msg = "select one student info by username failed"
			return err
		}
		if student != nil {
			resp.Success = errcode.ErrRegisterNicknameExist.Code
			resp.Msg = "user username is exist"
			return nil
		}

		teacher, err = ul.FindOneTeacherByScreenName(req.Username)
		if err != nil {
			resp.Success = -1
			resp.Msg = "select one teacher info by username failed"
			return err
		}
		if teacher != nil {
			resp.Success = errcode.ErrRegisterNicknameExist.Code
			resp.Msg = "user username is exist"
			return nil
		}

		// 学校不存在
		university, err = ul.FindUniversityByName(req.University)
		if err != nil {
			resp.Success = -1
			resp.Msg = "select one university info by No failed"
			return err
		}
		if university == nil {
			resp.Success = errcode.ErrRegisterUniversityNotExist.Code
			resp.Msg = "university is not exist"
			return nil
		}

		// 专业不存在
		profession, err = ul.FindProfessionByName(req.Profession)
		if err != nil {
			resp.Success = -1
			resp.Msg = "select one profession info by No failed"
			return err
		}
		if profession == nil {
			resp.Success = errcode.ErrRegisterProfessionNotExist.Code
			resp.Msg = "university is not exist"
			return nil
		}

		// 学号已经存在
		student, err = ul.FindOneStudentByNo(req.UserNo, university.ID)
		if err != nil {
			resp.Success = -1
			resp.Msg = "select student info by No failed"
			return err
		}
		if student != nil {
			resp.Success = errcode.ErrRegisterStuNoExist.Code
			resp.Msg = "student no is exist"
			return nil
		}

		// 身份证号已经存在
		student, err = ul.FindOneStudentByIDCardNumber(req.IdCardNumber)
		if err != nil {
			resp.Success = -1
			resp.Msg = "select one student info by id_card_number failed"
			return err
		}
		if student != nil {
			resp.Success = errcode.ErrRegisterIDCardNumberExist.Code
			resp.Msg = "user id_card_number is exist"
			return nil
		}

		teacher, err = ul.FindOneTeacherByIDCardNumber(req.IdCardNumber)
		if err != nil {
			resp.Success = -1
			resp.Msg = "select one teacher info by id_card_number failed"
			return err
		}
		if teacher != nil {
			resp.Success = errcode.ErrRegisterIDCardNumberExist.Code
			resp.Msg = "user id_card_number is exist"
			return nil
		}

		// 邮箱已经存在
		student, err = ul.FindOneStudentByEmail(req.Email)
		if err != nil {
			resp.Success = -1
			resp.Msg = "select one student info by email failed"
			return err
		}
		if student != nil {
			resp.Success = errcode.ErrRegisterEmailExist.Code
			resp.Msg = "user email is exist"
			return nil
		}

		// 邮箱已经存在
		teacher, err = ul.FindOneTeacherByEmail(req.Email)
		if err != nil {
			resp.Success = -1
			resp.Msg = "select one teacher info by email failed"
			return err
		}
		if teacher != nil {
			resp.Success = errcode.ErrRegisterEmailExist.Code
			resp.Msg = "user email is exist"
			return nil
		}

	} else {
		resp.Success = errcode.ErrLoginUserNameFailed.Code
		resp.Msg = errcode.ErrLoginUserNameFailed.Msg
		logger.Logger.Warn("phone invalid", zap.String("phone", req.Username), zap.Error(err))
		return err
	}

	// 加密密码
	hashPwd, err := common.GeneratePassword(req.Password, 1)
	if err != nil {
		resp.Success = -1
		resp.Msg = "student register generate password failed"
		logger.Logger.Warn("student register generate password failed", zap.Any("student", req), zap.Error(err))
		return err
	}

	// 身份证、手机号实名验证
	var userInfo map[string]interface{}
	code, userInfo, err := ul.CheckUserPhoneAndIdCard(req.Phone, req.IdCardNumber, req.RealName)
	if err != nil {
		resp.Success = -1
		resp.Msg = "check user phone and IdCard failed"
		logger.Logger.Warn("check user phone and IdCard failed", zap.Any("req", req), zap.Error(err))
		return err
	}

	if code != 0 {
		if code == errcode.ErrUserNameVerifyInvalid.Code {
			resp.Success = errcode.ErrUserNameVerifyInvalid.Code
			resp.Msg = errcode.ErrUserNameVerifyInvalid.Msg
			return nil
		} else if code == errcode.ErrUserIdCardVerifyInvalid.Code {
			resp.Success = errcode.ErrUserIdCardVerifyInvalid.Code
			resp.Msg = errcode.ErrUserIdCardVerifyInvalid.Msg
			return nil
		} else if code == errcode.ErrUserPhoneVerifyInvalid.Code {
			resp.Success = errcode.ErrUserPhoneVerifyInvalid.Code
			resp.Msg = errcode.ErrUserPhoneVerifyInvalid.Msg
			return nil
		} else {
			resp.Success = -1
			resp.Msg = "check user phone and IdCard failed"
			return nil
		}
	}

	uid := common.Md5V(uuid.NewV4().String())
	gender := "MALE"
	if req.Gender == 2 {
		gender = "FEMALE"
	}

	//mid, err := strconv.ParseInt(common.KRand(9, common.KC_RAND_KIND_NUM)+strconv.FormatInt(time.Now().Unix(), 10), 10, 64)
	mid, err := common.RandIncrId("im:peer")
	fmt.Println("student用户mid:::::", mid)
	if err != nil {
		resp.Success = -1
		resp.Msg = "student register generate  im id  failed"
		logger.Logger.Warn("student register generate  redis incr im id failed", zap.Any("student", req), zap.Error(err))
		return err
	}
	avatar := common.OSSFILEURLPREFIX +
		fmt.Sprintf(common.UPLOADFILEPATHPREFIX, common.USERAVATARFILE, req.ImageId+req.ImageId)

	verifyImage := common.OSSFILEURLPREFIX +
		fmt.Sprintf(common.UPLOADFILEPATHPREFIX, common.USERAVATARFILE, req.VerifyImageId+req.VerifyImageId)

	// 保存用户信息到mysql
	stu := &user.UStudent{
		UID:           uid,
		Password:      hashPwd,
		ScreenName:    req.Username,
		Phone:         req.Phone,
		Gender:        gender,
		Avatar:        avatar,
		Mid:           mid,
		ForbiStranger: int8(req.ForbiStranger),
		IDCardNumber:  req.IdCardNumber,
		StuNo:         req.UserNo,
		Email:         req.Email,
		UniversityID:  university.ID,
		ProfessionID:  profession.ID,
		VerifyImage:   verifyImage,
	}

	// 用户验证真实数据
	res, ok := userInfo["res"].(string)
	if res == "1" && ok {
		if userInfo["name"] != "" {
			stu.RealName = userInfo["name"].(string)
			stu.CreatedUser = null.String{String: userInfo["name"].(string), Valid: true}
		}
		stu.Birthday = userInfo["birthday"].(string)
		if userInfo["sex"].(string) != "男" {
			stu.Gender = "MALE"
		} else if userInfo["sex"].(string) != "女" {
			stu.Gender = "FEMALE"
		}
		stu.IDCardNumber = userInfo["idcard"].(string)
		stu.Address = userInfo["address"].(string)
	} else {
		resp.Success = -1
		resp.Msg = "register failed"
		return nil
	}

	var esAvatar *pb.Avatar
	if stu.Avatar != "" {
		esAvatar = &pb.Avatar{
			ThumbnailUrl: stu.Avatar + "?x-oss-process=image/resize,m_fixed,h_120,w_120",
			SmallPicUrl:  stu.Avatar + "?x-oss-process=image/resize,m_fixed,h_300,w_300",
			PicUrl:       stu.Avatar + "?x-oss-process=image/resize,m_fixed,h_800,w_800",
		}
	}

	// 保存动态的用户信息到es
	userReg := &pb.EsUserInfo{
		Id:            stu.UID,
		Username:      stu.ScreenName,
		ScreenName:    stu.ScreenName,
		UniversityId:  int64(university.ID),
		AvatarImage:   esAvatar,
		Mid:           stu.Mid,
		ForbiStranger: req.ForbiStranger,
		Phone:         stu.Phone,
		Email:         stu.Email,
	}

	if err = ul.SaveMsgUserInfoToEs(userReg); err != nil {
		resp.Success = -1
		resp.Msg = "user register failed"
		logger.Logger.Warn("insert student to es failed", zap.Any("student", stu), zap.Error(err))
		return err
	}

	if err = ul.InsertStudent(stu); err != nil {
		resp.Success = -1
		resp.Msg = "user register failed"
		logger.Logger.Warn("insert student to db failed", zap.Any("student", stu), zap.Error(err))
		return err
	}

	// 删除注册验证码
	ul.DelPhoneCode(req.Phone)

	resp.Success = 0
	resp.Msg = "success"
	return nil
}

func (u *UserHandler) TeacherRegister(ctx context.Context, req *pb.UserRegisterReq, resp *pb.Response) (err error) {
	var university *user.University
	var profession *user.UniversityCollegeProfession
	ul := &logic.UserLogic{}
	// 验证码对比
	uVerifyCode, err := ul.GetPhoneCode(req.Phone)
	if err != nil {
		resp.Success = -1
		resp.Msg = "register verify code failed"
		logger.Logger.Warn("register verify code failed", zap.String("phone", req.Phone), zap.Error(err))
		return err
	}

	if uVerifyCode != "" {
		if req.Code != uVerifyCode {
			resp.Success = errcode.ErrVerifyCodeCompareFailed.Code
			resp.Msg = "register user verify code is not equal failed"
			return nil
		}
	} else {
		resp.Success = errcode.ErrVerifyCodeNotExist.Code
		resp.Msg = "register user verify code is not exist"
		return nil
	}

	var student *user.UStudent
	var teacher *user.UTeacher
	if common.VerifyMobileFormat(req.Phone) {
		// 手机号已存在
		student, err = ul.FindOneStudentByPhone(req.Phone)
		if err != nil {
			resp.Success = -1
			resp.Msg = "select one student info by No failed"
			return err
		}
		if student != nil {
			resp.Success = errcode.ErrRegisterPhoneExist.Code
			resp.Msg = "user is exist"
			return nil
		}

		teacher, err = ul.FindOneTeacherByPhone(req.Phone)
		if err != nil {
			resp.Success = -1
			resp.Msg = "select one teacher info by No failed"
			return err
		}
		if teacher != nil {
			resp.Success = errcode.ErrRegisterPhoneExist.Code
			resp.Msg = "user is exist"
			return nil
		}

		// 用户名已经存在
		student, err = ul.FindOneStudentByScreenName(req.Username)
		if err != nil {
			resp.Success = -1
			resp.Msg = "select one student info by username failed"
			return err
		}
		if student != nil {
			resp.Success = errcode.ErrRegisterNicknameExist.Code
			resp.Msg = "user username is exist"
			return nil
		}

		teacher, err = ul.FindOneTeacherByScreenName(req.Username)
		if err != nil {
			resp.Success = -1
			resp.Msg = "select one teacher info by username failed"
			return err
		}
		if teacher != nil {
			resp.Success = errcode.ErrRegisterNicknameExist.Code
			resp.Msg = "user username is exist"
			return nil
		}

		// 学校不存在
		university, err = ul.FindUniversityByName(req.University)
		if err != nil {
			resp.Success = -1
			resp.Msg = "select one university info by No failed"
			return err
		}
		if university == nil {
			resp.Success = errcode.ErrRegisterUniversityNotExist.Code
			resp.Msg = "university is not exist"
			return nil
		}

		// 专业不存在
		profession, err = ul.FindProfessionByName(req.Profession)
		if err != nil {
			resp.Success = -1
			resp.Msg = "select one profession info by No failed"
			return err
		}
		if profession == nil {
			resp.Success = errcode.ErrRegisterProfessionNotExist.Code
			resp.Msg = "university is not exist"
			return nil
		}

		teacher, err = ul.FindOneTeacherByNo(req.UserNo, university.ID)
		if err != nil {
			resp.Success = -1
			resp.Msg = "select teacher info by No failed"
			return err
		}
		if teacher != nil {
			resp.Success = errcode.ErrRegisterStuNoExist.Code
			resp.Msg = "teacher no is exist"
			return nil
		}

		// 身份证号已经存在
		student, err = ul.FindOneStudentByIDCardNumber(req.IdCardNumber)
		if err != nil {
			resp.Success = -1
			resp.Msg = "select one student info by id_card_number failed"
			return err
		}
		if student != nil {
			resp.Success = errcode.ErrRegisterIDCardNumberExist.Code
			resp.Msg = "user id_card_number is exist"
			return nil
		}

		teacher, err = ul.FindOneTeacherByIDCardNumber(req.IdCardNumber)
		if err != nil {
			resp.Success = -1
			resp.Msg = "select one teacher info by id_card_number failed"
			return err
		}
		if teacher != nil {
			resp.Success = errcode.ErrRegisterIDCardNumberExist.Code
			resp.Msg = "user id_card_number is exist"
			return nil
		}

		// 邮箱已经存在
		student, err = ul.FindOneStudentByEmail(req.Email)
		if err != nil {
			resp.Success = -1
			resp.Msg = "select one student info by email failed"
			return err
		}
		if student != nil {
			resp.Success = errcode.ErrRegisterEmailExist.Code
			resp.Msg = "user email is exist"
			return nil
		}

		// 邮箱已经存在
		teacher, err = ul.FindOneTeacherByEmail(req.Email)
		if err != nil {
			resp.Success = -1
			resp.Msg = "select one teacher info by email failed"
			return err
		}
		if teacher != nil {
			resp.Success = errcode.ErrRegisterEmailExist.Code
			resp.Msg = "user email is exist"
			return nil
		}

	} else {
		resp.Success = errcode.ErrLoginUserNameFailed.Code
		resp.Msg = errcode.ErrLoginUserNameFailed.Msg
		logger.Logger.Warn("phone invalid", zap.String("phone", req.Username), zap.Error(err))
		return err
	}

	// 身份证、手机号实名验证
	var userInfo map[string]interface{}
	code, userInfo, err := ul.CheckUserPhoneAndIdCard(req.Phone, req.IdCardNumber, req.RealName)
	if err != nil {
		resp.Success = -1
		resp.Msg = "check user phone and IdCard failed"
		logger.Logger.Warn("check user phone and IdCard failed", zap.Any("req", req), zap.Error(err))
		return err
	}

	if code != 0 {
		if code == errcode.ErrUserNameVerifyInvalid.Code {
			resp.Success = errcode.ErrUserNameVerifyInvalid.Code
			resp.Msg = errcode.ErrUserNameVerifyInvalid.Msg
			return nil
		} else if code == errcode.ErrUserIdCardVerifyInvalid.Code {
			resp.Success = errcode.ErrUserIdCardVerifyInvalid.Code
			resp.Msg = errcode.ErrUserIdCardVerifyInvalid.Msg
			return nil
		} else if code == errcode.ErrUserPhoneVerifyInvalid.Code {
			resp.Success = errcode.ErrUserPhoneVerifyInvalid.Code
			resp.Msg = errcode.ErrUserPhoneVerifyInvalid.Msg
			return nil
		} else {
			resp.Success = -1
			resp.Msg = "check user phone and IdCard failed"
			return nil
		}
	}

	uid := common.Md5V(uuid.NewV4().String())
	gender := "MALE"
	if req.Gender == 2 {
		gender = "FEMALE"
	}

	//mid, err := strconv.ParseInt(common.KRand(9, common.KC_RAND_KIND_NUM)+strconv.FormatInt(time.Now().Unix(), 10), 10, 64)
	mid, err := common.RandIncrId("im:peer")
	if err != nil {
		resp.Success = -1
		resp.Msg = "teacher register generate  im id  failed"
		logger.Logger.Warn("teacher register generate  redis incr im id failed", zap.Any("teacher", req), zap.Error(err))
		return err
	}
	avatar := common.OSSFILEURLPREFIX +
		fmt.Sprintf(common.UPLOADFILEPATHPREFIX, common.USERAVATARFILE, req.ImageId+req.ImageId)

	verifyImage := common.OSSFILEURLPREFIX +
		fmt.Sprintf(common.UPLOADFILEPATHPREFIX, common.USERAVATARFILE, req.VerifyImageId+req.VerifyImageId)

	// 加密密码
	if req.Password == "" {
		resp.Success = -1
		resp.Msg = "teacher register password  null"
		logger.Logger.Warn("teacher register password  null", zap.Any("teacher", req), zap.Error(err))
		return err
	}
	hashPwd, err := common.GeneratePassword(req.Password, 1)
	if err != nil {
		resp.Success = -1
		resp.Msg = "student register generate password failed"
		logger.Logger.Warn("student register generate password failed", zap.Any("student", req), zap.Error(err))
		return err
	}

	// 保存用户信息到mysql
	teach := &user.UTeacher{
		UID:           uid,
		Password:      hashPwd,
		ScreenName:    req.Username,
		Phone:         req.Phone,
		Gender:        gender,
		Avatar:        avatar,
		Mid:           mid,
		ForbiStranger: int8(req.ForbiStranger),
		IDCardNumber:  req.IdCardNumber,
		TeachNo:       req.UserNo,
		Email:         req.Email,
		UniversityID:  university.ID,
		ProfessionID:  profession.ID,
		VerifyImage:   verifyImage,
	}

	// 用户验证真实数据
	res, ok := userInfo["res"].(string)
	if res == "1" && ok {
		if userInfo["name"] != "" {
			teach.RealName = userInfo["name"].(string)
			teach.CreatedUser = null.String{String: userInfo["name"].(string), Valid: true}
		}
		teach.Birthday = userInfo["birthday"].(string)
		if userInfo["sex"].(string) != "男" {
			teach.Gender = "MALE"
		} else if userInfo["sex"].(string) != "女" {
			teach.Gender = "FEMALE"
		}
		teach.IDCardNumber = userInfo["idcard"].(string)
		teach.Address = userInfo["address"].(string)
	} else {
		resp.Success = -1
		resp.Msg = "register failed"
		return nil
	}

	var esAvatar *pb.Avatar
	if teach.Avatar != "" {
		esAvatar = &pb.Avatar{
			ThumbnailUrl: teach.Avatar + "?x-oss-process=image/resize,m_fixed,h_120,w_120",
			SmallPicUrl:  teach.Avatar + "?x-oss-process=image/resize,m_fixed,h_300,w_300",
			PicUrl:       teach.Avatar + "?x-oss-process=image/resize,m_fixed,h_800,w_800",
		}
	}

	// 保存动态的用户信息到es
	userReg := &pb.EsUserInfo{
		Id:            teach.UID,
		Username:      teach.ScreenName,
		ScreenName:    teach.ScreenName,
		UniversityId:  int64(university.ID),
		AvatarImage:   esAvatar,
		Mid:           teach.Mid,
		ForbiStranger: req.ForbiStranger,
		Phone:         teach.Phone,
		Email:         teach.Email,
	}
	if err = ul.SaveMsgUserInfoToEs(userReg); err != nil {
		resp.Success = -1
		resp.Msg = "user register failed"
		logger.Logger.Warn("insert teacher to es failed", zap.Any("teacher", teach), zap.Error(err))
		return err
	}

	if err = ul.InsertTeacher(teach); err != nil {
		resp.Success = -1
		resp.Msg = "user register failed"
		logger.Logger.Warn("insert teacher to db failed", zap.Any("teacher", teacher), zap.Error(err))
		return err
	}

	// 删除注册验证码
	ul.DelPhoneCode(req.Phone)

	resp.Success = 0
	resp.Msg = "success"
	return nil
}

func (u *UserHandler) CheckVerifyCode(ctx context.Context, req *pb.VerifyCodeRegReq, resp *pb.Response) (err error) {
	var uVerifyCode string
	ul := &logic.UserLogic{}
	// 验证码对比
	uVerifyCode, err = ul.GetPhoneCode(req.Phone)
	if err != nil {
		resp.Success = -1
		resp.Msg = "register verify code failed"
		logger.Logger.Warn("register verify code failed", zap.String("phone", req.Phone), zap.Error(err))
		return err
	}

	if uVerifyCode != "" {
		if req.Code != uVerifyCode {
			resp.Success = errcode.ErrVerifyCodeCompareFailed.Code
			resp.Msg = "register user verify code is not equal failed"
			return nil
		}
	} else {
		resp.Success = errcode.ErrVerifyCodeNotExist.Code
		resp.Msg = "register user verify code is not exist"
		return nil
	}
	return nil
}

// 学生注册
//func (u *UserHandler) AddStudentInfo(ctx context.Context, req *pb.StudentRegInfo, resp *pb.Response) (err error) {
//	var student *user.UStudent
//	var university *user.University
//	var profession *user.UniversityCollegeProfession
//	ul := &logic.UserLogic{}
//
//	// 学校不存在
//	university, err = ul.FindUniversityByName(req.University)
//	if err != nil {
//		resp.Success = -1
//		resp.Msg = "select one university info by No failed"
//		return err
//	}
//	if university == nil {
//		resp.Success = errcode.ErrRegisterUniversityNotExist.Code
//		resp.Msg = "university is not exist"
//		return nil
//	}
//
//	// 专业不存在
//	profession, err = ul.FindProfessionByName(req.Profession)
//	if err != nil {
//		resp.Success = -1
//		resp.Msg = "select one profession info by No failed"
//		return err
//	}
//	if profession == nil {
//		resp.Success = errcode.ErrRegisterProfessionNotExist.Code
//		resp.Msg = "university is not exist"
//		return nil
//	}
//
//	// 学号已经存在
//	student, err = ul.FindOneStudentByNo(req.StuNo, university.ID)
//	if err != nil {
//		resp.Success = -1
//		resp.Msg = "select student info by No failed"
//		return err
//	}
//	if student != nil {
//		resp.Success = errcode.ErrRegisterStuNoExist.Code
//		resp.Msg = "user is not exist"
//		return nil
//	}
//
//	// 身份证号已经存在
//	student, err = ul.FindOneStudentByIDCardNumber(req.IdCardNumber)
//	if err != nil {
//		resp.Success = -1
//		resp.Msg = "select one student info by id_card_number failed"
//		return err
//	}
//	if student != nil {
//		resp.Success = errcode.ErrRegisterIDCardNumberExist.Code
//		resp.Msg = "user id_card_number is exist"
//		return nil
//	}
//
//	// 邮箱已经存在
//	student, err = ul.FindOneStudentByEmail(req.Email)
//	if err != nil {
//		resp.Success = -1
//		resp.Msg = "select one student info by email failed"
//		return err
//	}
//	if student != nil {
//		resp.Success = errcode.ErrRegisterEmailExist.Code
//		resp.Msg = "user email is exist"
//		return nil
//	}
//
//	// 手机号不存在
//	student, err = ul.FindOneStudentByPhone(req.Phone)
//	if err != nil {
//		resp.Success = -1
//		resp.Msg = "select one student info by phone failed"
//		return err
//	}
//	if student == nil {
//		resp.Success = errcode.ErrRegisterPhoneNotExist.Code
//		resp.Msg = "user phone is not exist"
//		return nil
//	}
//
//	// 身份证、手机号实名验证
//	code, err := ul.CheckUserPhoneAndIdCard(student, req.Phone, req.IdCardNumber, req.RealName)
//	if err != nil {
//		resp.Success = -1
//		resp.Msg = "check user phone and IdCard failed"
//		logger.Logger.Warn("check user phone and IdCard failed", zap.Any("req", req), zap.Error(err))
//		return err
//	}
//
//	if code != 0 {
//		if code == errcode.ErrUserNameVerifyInvalid.Code {
//			resp.Success = errcode.ErrUserNameVerifyInvalid.Code
//			resp.Msg = errcode.ErrUserNameVerifyInvalid.Msg
//			return nil
//		} else if code == errcode.ErrUserIdCardVerifyInvalid.Code {
//			resp.Success = errcode.ErrUserIdCardVerifyInvalid.Code
//			resp.Msg = errcode.ErrUserIdCardVerifyInvalid.Msg
//			return nil
//		} else if code == errcode.ErrUserPhoneVerifyInvalid.Code {
//			resp.Success = errcode.ErrUserPhoneVerifyInvalid.Code
//			resp.Msg = errcode.ErrUserPhoneVerifyInvalid.Msg
//			return nil
//		} else {
//			resp.Success = -1
//			resp.Msg = "check user phone and IdCard failed"
//			return nil
//		}
//	}
//
//	//entryDate, _ := strconv.ParseInt(req.EntryDate, 10, 64)
//	//graduationDate, _ := strconv.ParseInt(req.GraduationDate, 10, 64)
//
//	// 保存用户信息到mysql
//	student.StuNo = req.StuNo
//	student.Email = req.Email
//	student.UniversityID = university.ID
//	student.ProfessionID = profession.ID
//	student.IsVerified = 1
//	fmt.Println("更新stu", student)
//
//	if err = ul.UpdateStudentInfo(student); err != nil {
//		resp.Success = -1
//		resp.Msg = "user register failed"
//		logger.Logger.Warn("insert student to db failed", zap.Any("student", student), zap.Error(err))
//		return err
//	}
//
//	var avatar *pb.Avatar
//	if student.Avatar != "" {
//		avatar = &pb.Avatar{
//			ThumbnailUrl: student.Avatar + "?x-oss-process=image/resize,m_fixed,h_120,w_120",
//			SmallPicUrl:  student.Avatar + "?x-oss-process=image/resize,m_fixed,h_300,w_300",
//			PicUrl:       student.Avatar + "?x-oss-process=image/resize,m_fixed,h_800,w_800",
//		}
//	}
//
//	// 保存动态的用户信息到es
//	userReg := &pb.EsUserInfo{
//		Id:           student.UID,
//		Username:     student.ScreenName,
//		ScreenName:   student.ScreenName,
//		UniversityId: int64(university.ID),
//		AvatarImage:  avatar,
//	}
//	if err = ul.SaveMsgUserInfoToEs(userReg); err != nil {
//		if err = ul.DelStudent(student); err != nil {
//			logger.Logger.Warn("del user to from failed", zap.Any("student", student), zap.Error(err))
//		}
//		resp.Success = -1
//		resp.Msg = "user register failed"
//		logger.Logger.Warn("insert student to es failed", zap.Any("student", student), zap.Error(err))
//		return err
//	}
//
//	resp.Success = 0
//	resp.Msg = "success"
//	return nil
//}

func (u *UserHandler) UpdateStudentInfo(ctx context.Context, req *pb.UpdateStudentInfoReq, resp *pb.Response) (err error) {
	var student *user.UStudent
	ul := &logic.UserLogic{}

	student, err = ul.FindOneStudentById(req.Id)
	if err != nil {
		resp.Success = -1
		resp.Msg = "select one student info by No failed"
		return err
	}
	if student == nil {
		resp.Success = errcode.ErrUserNotExist.Code
		resp.Msg = "user No is exist"
		return nil
	}

	if req.ScreenName != "" {
		// 用户名已经存在
		stu, err := ul.FindOneStudentByScreenName(req.ScreenName)
		if err != nil {
			resp.Success = -1
			resp.Msg = "select one student info by screen_name failed"
			return err
		}
		if stu != nil {
			resp.Success = errcode.ErrRegisterNicknameExist.Code
			resp.Msg = "user screen_name is exist"
			return nil
		}

		student.ScreenName = req.ScreenName

	}

	if req.Birthday != "" {
		student.Birthday = req.Birthday
	}

	if req.Bio != "" {
		student.Bio = req.Bio
	}

	if req.Zodiac != "" {
		student.Zodiac = req.Zodiac
	}

	if req.Email != "" {
		student.Email = req.Email
	}

	student.UpdatedAt = time.Now().Local()

	isVerified := false
	if student.IsVerified == 1 {
		isVerified = true
	}

	// 更新动态的用户信息到es
	user := &pb.EsUserInfo{
		Id:            req.Id,
		Username:      student.ScreenName,
		ScreenName:    student.ScreenName,
		Bio:           student.Bio,
		IsVerified:    isVerified,
		VerifyMessage: student.VerifyMessage,
		Zodiac:        student.Zodiac,
		UniversityId:  int64(student.UniversityID),
	}
	if err = ul.UpdateMsgUserInfoToEs(user); err != nil {
		logger.Logger.Warn("insert student to es failed", zap.Any("user", user), zap.Error(err))
		resp.Success = -1
		resp.Msg = "user update failed"
		return err
	}

	// 更新用户基本信息
	if err = ul.UpdateStudentInfo(student); err != nil {
		resp.Success = -1
		resp.Msg = "user update failed"
		logger.Logger.Warn("update student to db failed", zap.Any("student", student), zap.Error(err))
		return
	}

	return nil
}

func (u *UserHandler) GetStudentInfoById(ctx context.Context, req *pb.GetStudentByIdReq, resp *pb.StudentInfoDetail) error {
	ul := &logic.UserLogic{}
	var err error
	var student *user.UStudent

	student, err = ul.FindOneStudentById(req.Id)
	if err != nil {
		logger.Logger.Warn("select one verify code by phone failed", zap.Any("uid", req.Id), zap.Error(err))
		return err
	}

	if student != nil {
		resp.StuNo = student.StuNo
		resp.ScreenName = student.ScreenName
		resp.Gender = student.Gender
		resp.Phone = student.Phone
		resp.Email = student.Email
		resp.Birthday = student.Birthday
		resp.Bio = student.Bio

		resp.AvatarImage.ThumbnailUrl = student.Avatar + "?x-oss-process=image/resize,m_fixed,h_120,w_120"
		resp.AvatarImage.SmallPicUrl = student.Avatar + "?x-oss-process=image/resize,m_fixed,h_300,w_300"
		resp.AvatarImage.PicUrl = student.Avatar + "?x-oss-process=image/resize,m_fixed,h_800,w_800"

		university, _ := ul.FindUniversityById(student.UniversityID)
		resp.University = university.Name

		college, _ := ul.FindCollegeById(student.CollegeID)
		resp.College = college.Name

		profession, _ := ul.FindProfessionById(student.ProfessionID)
		resp.Profession = profession.Name

		class, _ := ul.FindClassById(student.ClassID)
		resp.Class = class.Name

		resp.EntryDate = student.EntryDate.Format("2006-01-02 15:04:05")
		resp.GraduationDate = student.GraduationDate.Format("2006-01-02 15:04:05")

	}

	return nil
}

func (u *UserHandler) GetStudentInfoByUsername(ctx context.Context, req *pb.GetStudentByUsernameReq, resp *pb.StudentInfoDetail) error {
	ul := &logic.UserLogic{}
	var err error
	var student *user.UStudent

	student, err = ul.FindOneStudentByScreenName(req.Username)
	if err != nil {
		logger.Logger.Warn("select one verify code by phone failed", zap.Any("username", req.Username), zap.Error(err))
		return err
	}

	if student != nil {
		resp.StuNo = student.StuNo
		resp.ScreenName = student.ScreenName
		resp.Gender = student.Gender
		resp.Phone = student.Phone
		resp.Email = student.Email
		resp.Birthday = student.Birthday
		resp.Bio = student.Bio

		resp.AvatarImage.ThumbnailUrl = student.Avatar + "?x-oss-process=image/resize,m_fixed,h_120,w_120"
		resp.AvatarImage.SmallPicUrl = student.Avatar + "?x-oss-process=image/resize,m_fixed,h_300,w_300"
		resp.AvatarImage.PicUrl = student.Avatar + "?x-oss-process=image/resize,m_fixed,h_800,w_800"

		university, _ := ul.FindUniversityById(student.UniversityID)
		resp.University = university.Name

		college, _ := ul.FindCollegeById(student.CollegeID)
		resp.College = college.Name

		profession, _ := ul.FindProfessionById(student.ProfessionID)
		resp.Profession = profession.Name

		class, _ := ul.FindClassById(student.ClassID)
		resp.Class = class.Name

		resp.EntryDate = student.EntryDate.Format("2006-01-02 15:04:05")
		resp.GraduationDate = student.GraduationDate.Format("2006-01-02 15:04:05")

	}

	return nil
}

func (u *UserHandler) PwdLogin(ctx context.Context, req *pb.PwdLoginReq, resp *pb.LoginResponse) error {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("PwdLogin recover", err)
		}
	}()
	ul := &logic.UserLogic{}
	var err error
	// 老师：1   学生：2
	if common.VerifyMobileFormat(req.Username) {
		var student *user.UStudent
		student, err = ul.FindOneStudentByPhone(req.Username)
		if err != nil {
			resp.Success = -1
			resp.Msg = "select one student info by phone failed"
			return err
		}

		if student != nil {
			// 用户未验证
			//if student.IsVerified == 0 {
			//	resp.Success = errcode.ErrUserIsNotVerify.Code
			//	resp.Msg = errcode.ErrUserIsNotVerify.Msg
			//	logger.Logger.Warn("user is not verify", zap.String("phone", req.Username), zap.Error(err))
			//	return err
			//}

			validate := common.ValidatePassword(student.Password, req.Password)
			if !validate {
				resp.Success = errcode.ErrLoginUserPwdFailed.Code
				resp.Msg = "login password failed"
				return nil
			}

			resp.Uid = student.UID
			resp.Nickname = student.ScreenName
			resp.UniversityId = int64(student.UniversityID)
			resp.Mid = student.Mid
			resp.Avatar = student.Avatar + "?x-oss-process=image/resize,m_fixed,h_800,w_800"

		} else {
			var teacher *user.UTeacher
			teacher, err = ul.FindOneTeacherByPhone(req.Username)
			if err != nil {
				resp.Success = -1
				resp.Msg = "select one student info by phone failed"
				return err
			}

			if teacher == nil {
				resp.Success = errcode.ErrUserNotExist.Code
				resp.Msg = "user is not exist"
				return nil
			}

			validate := common.ValidatePassword(teacher.Password, req.Password)
			if !validate {
				resp.Success = errcode.ErrLoginUserPwdFailed.Code
				resp.Msg = "login password failed"
				return nil
			}

			resp.Uid = teacher.UID
			resp.Nickname = teacher.ScreenName
			resp.UniversityId = int64(teacher.UniversityID)
			resp.Mid = teacher.Mid
			resp.Avatar = teacher.Avatar + "?x-oss-process=image/resize,m_fixed,h_800,w_800"


		}
	} else {
		resp.Success = errcode.ErrLoginUserNameFailed.Code
		resp.Msg = errcode.ErrLoginUserNameFailed.Msg
		logger.Logger.Warn("phone invalid", zap.String("phone", req.Username), zap.Error(err))
		return err
	}

	resp.Success = 0
	resp.Msg = "success"
	return nil
}

func (u *UserHandler) VerifyCodeLogin(ctx context.Context, req *pb.VerifyCodeLoginReq, resp *pb.LoginResponse) error {
	ul := &logic.UserLogic{}
	var err error
	var uVerifyCode string

	uVerifyCode, err = ul.GetPhoneCode(req.Phone)
	if err != nil {
		resp.Success = -1
		resp.Msg = "login verify code failed"
		logger.Logger.Warn("login verify code failed", zap.String("phone", req.Phone), zap.Error(err))
		return err
	}

	if uVerifyCode != "" {
		if req.Code != uVerifyCode {
			resp.Success = errcode.ErrVerifyCodeCompareFailed.Code
			resp.Msg = "login user verify code is not equal failed"
			return nil
		}
	} else {
		resp.Success = errcode.ErrVerifyCodeNotExist.Code
		resp.Msg = "login user verify code is not exist"
		return nil
	}

	// 老师：1   学生：2
	if common.VerifyMobileFormat(req.Phone) {
		var student *user.UStudent

		student, err = ul.FindOneStudentByPhone(req.Phone)
		if err != nil {
			resp.Success = -1
			resp.Msg = "select one verify code by phone failed"
			logger.Logger.Warn("select one verify code by phone failed", zap.String("phone", req.Phone), zap.Error(err))
			return err
		}

		if student != nil {
			// 用户未验证
			//if student.IsVerified == 0 {
			//	resp.Success = errcode.ErrUserIsNotVerify.Code
			//	resp.Msg = errcode.ErrUserIsNotVerify.Msg
			//	logger.Logger.Warn("user is not verify", zap.String("phone", req.Phone), zap.Error(err))
			//	return err
			//}

			resp.Uid = student.UID
			resp.Nickname = student.ScreenName
			resp.UniversityId = int64(student.UniversityID)
			resp.Mid = student.Mid
			resp.Avatar = student.Avatar + "?x-oss-process=image/resize,m_fixed,h_800,w_800"

		} else {
			var teacher *user.UTeacher
			teacher, err = ul.FindOneTeacherByPhone(req.Phone)
			if err != nil {
				resp.Success = -1
				resp.Msg = "select one verify code by phone failed"
				logger.Logger.Warn("select one verify code by phone failed", zap.String("phone", req.Phone), zap.Error(err))
				return err
			}

			if teacher == nil {
				resp.Success = errcode.ErrUserNotExist.Code
				resp.Msg = "login user failed"
				return nil
			}

			// 用户未验证
			//if student.IsVerified == 0 {
			//	resp.Success = errcode.ErrUserIsNotVerify.Code
			//	resp.Msg = errcode.ErrUserIsNotVerify.Msg
			//	logger.Logger.Warn("user is not verify", zap.String("phone", req.Phone), zap.Error(err))
			//	return err
			//}

			resp.Uid = teacher.UID
			resp.Nickname = teacher.ScreenName
			resp.UniversityId = int64(teacher.UniversityID)
			resp.Mid = teacher.Mid
			resp.Avatar = teacher.Avatar + "?x-oss-process=image/resize,m_fixed,h_800,w_800"
		}
	} else {
		resp.Success = errcode.ErrLoginUserNameFailed.Code
		resp.Msg = errcode.ErrLoginUserNameFailed.Msg
		logger.Logger.Warn("phone invalid", zap.String("phone", req.Phone), zap.Error(err))
		return err

	}

	// 删除登录验证码
	go ul.DelPhoneCode(req.Phone)

	resp.Success = 0
	resp.Msg = "success"
	return nil
}

func (u *UserHandler) GetVerifyCode(ctx context.Context, req *pb.VerifyCodeReq, resp *pb.VerifyCodeResponse) error {
	code := common.GenValidateCode(6)

	ul := &logic.UserLogic{}

	tempCode, err := ul.GetPhoneCode(req.Phone)
	if err != nil {
		resp.Status = -1
		resp.Code = ""
		return err
	}
	if tempCode != "" {
		resp.Status = errcode.ErrBusinessMinuteLimitControlFailed.Code
		resp.Code = ""
		return err
	}

	err = ul.SavePhoneCode(req.Phone, code)
	if err != nil {
		resp.Status = -1
		resp.Code = ""
		return err
	}

	status, err := common.SendVerifyCode(req.Phone, code)
	if err != nil {
		resp.Status = -1
		resp.Code = ""
		logger.Logger.Warn("send Verify Code failed", zap.String("phone", req.Phone), zap.String("code", code), zap.Error(err))
		return err
	}

	resp.Status = status

	resp.Code = code
	return nil
}

func (u *UserHandler) Logout(ctx context.Context, req *pb.LogoutReq, resp *pb.Response) (err error) {
	ul := &logic.UserLogic{}
	if req.Type == 2 {

	}
	var student *user.UStudent

	student, err = ul.FindOneStudentById(req.Id)
	if err != nil {
		resp.Success = -1
		resp.Msg = "select one student info by id failed"
		return err
	}

	if student == nil {
		resp.Success = errcode.ErrUserNotExist.Code
		resp.Msg = "logout username failed"
		return nil
	}

	if err = ul.DelLoginSession(req.Id); err != nil {
		resp.Success = errcode.ErrLogoutFailed.Code
		resp.Msg = "user logout error"
		return err
	}

	resp.Success = 0
	resp.Msg = "success"
	return nil
}

func (u *UserHandler) GetUniversity(ctx context.Context, req *pb.GetUniversityReq, resp *pb.GetUniversityListResponse) error {
	ul := &logic.UserLogic{}
	universityList, err := ul.FindUniversityByLikeName(req.University)
	if err != nil {
		return err
	}
	for _, university := range universityList {
		getUniversityResponse := &pb.GetUniversityResponse{
			Id:   int64(university.ID),
			Name: university.Name,
		}
		resp.GetUniversityList = append(resp.GetUniversityList, getUniversityResponse)
	}

	return nil
}

func (u *UserHandler) GetUniversityList(ctx context.Context, req *pb.GetUniversityListReq, resp *pb.GetUniversityListResponse) (err error) {
	ul := &logic.UserLogic{}
	universityList, err := ul.FindUniversityList()
	if err != nil {
		return err
	}
	for _, university := range universityList {
		getUniversityResponse := &pb.GetUniversityResponse{
			Id:   int64(university.ID),
			Name: university.Name,
		}
		resp.GetUniversityList = append(resp.GetUniversityList, getUniversityResponse)
	}
	return nil
}

func (u *UserHandler) GetCollegeList(ctx context.Context, req *pb.GetCollegeListReq, resp *pb.GetCollegeListResponse) (err error) {
	ul := &logic.UserLogic{}
	collegeList, err := ul.FindCollegeListByUniId(req.UniversityId)
	if err != nil {
		return err
	}
	for _, college := range collegeList {
		getCollegeResponse := &pb.GetCollegeResponse{
			Id:   int64(college.ID),
			Name: college.Name,
		}
		resp.GetCollegeList = append(resp.GetCollegeList, getCollegeResponse)
	}
	return nil
}

func (u *UserHandler) GetProfessionList(ctx context.Context, req *pb.GetProfessionListReq, resp *pb.GetProfessionListResponse) error {
	ul := &logic.UserLogic{}
	scienceList, err := ul.FindProfessionListByCollegeId(req.UniversityId, req.CollegeId)
	if err != nil {
		return err
	}
	for _, science := range scienceList {
		getScienceResponse := &pb.GetProfessionResponse{
			Id:   int64(science.ID),
			Name: science.Name,
		}
		resp.GetProfessionList = append(resp.GetProfessionList, getScienceResponse)
	}
	return nil
}

func (u *UserHandler) GetClassList(ctx context.Context, req *pb.GetClassListReq, resp *pb.GetClassListResponse) error {
	ul := &logic.UserLogic{}
	classList, err := ul.FindClassListByProfessionId(req.UniversityId, req.CollegeId, req.ProfessionId)
	if err != nil {
		return err
	}
	for _, class := range classList {
		getClassResponse := &pb.GetClassResponse{
			Id:   int64(class.ID),
			Name: class.Name,
		}
		resp.GetClassList = append(resp.GetClassList, getClassResponse)
	}
	return nil
}

// 更新用户头像
func (u *UserHandler) UpdateUserAvatar(ctx context.Context, req *pb.UpdateUserAvatarReq, resp *pb.Response) error {
	ul := &logic.UserLogic{}

	if req.Type == 2 {
		student, err := ul.FindOneStudentById(req.Uid)
		if err != nil {
			resp.Success = -1
			resp.Msg = "select one student info by No failed"
			return err
		}
		if student == nil {
			resp.Success = errcode.ErrUserNotExist.Code
			resp.Msg = "user No is exist"
			return nil
		}

		if req.Avatar != "" {
			student.Avatar = req.Avatar

			if err = ul.UpdateStudentInfo(student); err != nil {
				resp.Success = -1
				resp.Msg = "user update avatar failed"
				logger.Logger.Warn("update student avatar to db failed", zap.Any("avatar", req.Avatar), zap.Error(err))
				return err
			}
		} else {
			resp.Success = -1
			resp.Msg = "params invalid"
			return nil
		}
	} else if req.Type == 1 {
		teacher, err := ul.FindOneTeacherById(req.Uid)
		if err != nil {
			resp.Success = -1
			resp.Msg = "select one teacher info by No failed"
			return err
		}
		if teacher == nil {
			resp.Success = errcode.ErrUserNotExist.Code
			resp.Msg = "user No is exist"
			return nil
		}

		if req.Avatar != "" {
			teacher.Avatar = req.Avatar

			if err = ul.UpdateTeacherInfo(teacher); err != nil {
				resp.Success = -1
				resp.Msg = "user update avatar failed"
				logger.Logger.Warn("update teacher avatar to db failed", zap.Any("avatar", req.Avatar), zap.Error(err))
				return err
			}
		} else {
			resp.Success = -1
			resp.Msg = "params invalid"
			return nil
		}
	} else {
		resp.Success = -1
		resp.Msg = "params invalid"
		return nil
	}

	resp.Success = 0
	resp.Msg = "update user avatar success"
	return nil
}

func (u *UserHandler) UpdateUserPhone(ctx context.Context, req *pb.UpdateUserPhoneReq, resp *pb.Response) (err error) {
	ul := &logic.UserLogic{}

	if req.Type == 2 {
		var student *user.UStudent
		student, err = ul.FindOneStudentById(req.Uid)
		if err != nil {
			resp.Success = -1
			resp.Msg = "select one student info by No failed"
			return err
		}
		if student == nil {
			resp.Success = errcode.ErrUserNotExist.Code
			resp.Msg = "user No is exist"
			return nil
		}

		if req.Phone != "" {
			student.ScreenName = req.Phone
		}

		uVerifyCode, err := ul.GetPhoneCode(req.Phone)
		if err != nil {
			resp.Success = -1
			resp.Msg = "verify code failed"
			logger.Logger.Warn("login verify code failed", zap.String("phone", req.Phone), zap.Error(err))
			return err
		}

		if uVerifyCode != "" {
			if req.Code == uVerifyCode {
				_ = ul.DelPhoneCode(req.Phone)
				resp.Success = 0
				resp.Msg = "success"
				return nil
			} else {
				resp.Success = errcode.ErrVerifyCodeCompareFailed.Code
				resp.Msg = "login user verify code is not equal failed"
				return nil
			}
		} else {
			resp.Success = errcode.ErrVerifyCodeNotExist.Code
			resp.Msg = " verify code is not exist"
			return nil
		}

		validate := common.ValidatePassword(student.Password, req.Password)
		if !validate {
			resp.Success = errcode.ErrLoginUserPwdFailed.Code
			resp.Msg = "login password failed"
			return nil
		}

		student.Phone = req.Phone

		if err = ul.UpdateStudentInfo(student); err != nil {
			resp.Success = -1
			resp.Msg = "user update phone failed"
			logger.Logger.Warn("user update phone failed", zap.Any("student", student), zap.Any("new phone", student), zap.Error(err))
		}
	} else if req.Type == 1 {
		var teacher *user.UTeacher
		teacher, err = ul.FindOneTeacherById(req.Uid)
		if err != nil {
			resp.Success = -1
			resp.Msg = "select one teacher info by No failed"
			return err
		}
		if teacher == nil {
			resp.Success = errcode.ErrUserNotExist.Code
			resp.Msg = "user No is exist"
			return nil
		}

		if req.Phone != "" {
			teacher.Phone = req.Phone
		}

		uVerifyCode, err := ul.GetPhoneCode(req.Phone)
		if err != nil {
			resp.Success = -1
			resp.Msg = "verify code failed"
			logger.Logger.Warn("login verify code failed", zap.String("phone", req.Phone), zap.Error(err))
			return err
		}

		if uVerifyCode != "" {
			if req.Code == uVerifyCode {
				_ = ul.DelPhoneCode(req.Phone)
				resp.Success = 0
				resp.Msg = "success"
				return nil
			} else {
				resp.Success = errcode.ErrVerifyCodeCompareFailed.Code
				resp.Msg = "login user verify code is not equal failed"
				return nil
			}
		} else {
			resp.Success = errcode.ErrVerifyCodeNotExist.Code
			resp.Msg = " verify code is not exist"
			return nil
		}

		validate := common.ValidatePassword(teacher.Password, req.Password)
		if !validate {
			resp.Success = errcode.ErrLoginUserPwdFailed.Code
			resp.Msg = "login password failed"
			return nil
		}

		teacher.Phone = req.Phone

		if err = ul.UpdateTeacherInfo(teacher); err != nil {
			resp.Success = -1
			resp.Msg = "user update phone failed"
			logger.Logger.Warn("user update phone failed", zap.Any("teacher", teacher), zap.Any("new phone", req), zap.Error(err))
		}
	} else {
		resp.Success = -1
		resp.Msg = "user update phone failed"
		return nil
	}

	resp.Success = 0
	resp.Msg = "success"
	return nil

}

func (u *UserHandler) UpdateUserPassword(ctx context.Context, req *pb.UpdateUserPasswordReq, resp *pb.Response) (err error) {
	ul := &logic.UserLogic{}

	if req.Type == 2 {
		var student *user.UStudent
		student, err = ul.FindOneStudentByPhone(req.Phone)
		if err != nil {
			resp.Success = -1
			resp.Msg = "select one student info by No failed"
			return err
		}
		if student == nil {
			resp.Success = errcode.ErrUserNotExist.Code
			resp.Msg = "user No is exist"
			return nil
		}

		// 用户未验证
		//if student.IsVerified == 0 {
		//	resp.Success = errcode.ErrUserIsNotVerify.Code
		//	resp.Msg = errcode.ErrUserIsNotVerify.Msg
		//	logger.Logger.Warn("user is not verify", zap.String("phone", req.Phone), zap.Error(err))
		//	return err
		//}

		uVerifyCode, err := ul.GetPhoneCode(req.Phone)
		if err != nil {
			resp.Success = -1
			resp.Msg = "login verify code failed"
			logger.Logger.Warn("login verify code failed", zap.String("phone", req.Phone), zap.Error(err))
			return err
		}

		if uVerifyCode != "" {
			if req.Code != uVerifyCode {
				resp.Success = errcode.ErrVerifyCodeCompareFailed.Code
				resp.Msg = "login user verify code is not equal failed"
				return nil
			}
		} else {
			resp.Success = errcode.ErrVerifyCodeNotExist.Code
			resp.Msg = "login user verify code is not exist"
			return nil
		}

		if req.OldPassword != "" {
			validate := common.ValidatePassword(student.Password, req.OldPassword)
			if !validate {
				resp.Success = errcode.ErrLoginUserPwdFailed.Code
				resp.Msg = "login password failed"
				return nil
			}
		}

		hashPwd, err := common.GeneratePassword(req.NewPassword, 1)
		if err != nil {
			resp.Success = -1
			resp.Msg = "generate user update password failed"
			logger.Logger.Warn("generate user update password failed", zap.Any("user", req), zap.Error(err))
			return err
		}
		student.Password = hashPwd
		if err = ul.UpdateStudentInfo(student); err != nil {
			resp.Success = -1
			resp.Msg = "user update password failed"
			logger.Logger.Warn("user update password failed", zap.Any("student", student), zap.Error(err))
			return err
		}

		// 删除登录的token信息，让用户重新登录
		if err = ul.DelLoginSession(student.UID); err != nil {
			logger.Logger.Warn("del loginSession for user update password failed", zap.Any("student", student), zap.Error(err))
			return err
		}
	} else if req.Type == 1 {
		var teacher *user.UTeacher
		teacher, err = ul.FindOneTeacherByPhone(req.Phone)
		if err != nil {
			resp.Success = -1
			resp.Msg = "select one teacher info by No failed"
			return err
		}
		if teacher == nil {
			resp.Success = errcode.ErrUserNotExist.Code
			resp.Msg = "user No is exist"
			return nil
		}

		// 用户未验证
		//if teacher.IsVerified == 0 {
		//	resp.Success = errcode.ErrUserIsNotVerify.Code
		//	resp.Msg = errcode.ErrUserIsNotVerify.Msg
		//	logger.Logger.Warn("user is not verify", zap.String("phone", req.Phone), zap.Error(err))
		//	return err
		//}

		uVerifyCode, err := ul.GetPhoneCode(req.Phone)
		if err != nil {
			resp.Success = -1
			resp.Msg = "login verify code failed"
			logger.Logger.Warn("login verify code failed", zap.String("phone", req.Phone), zap.Error(err))
			return err
		}

		if uVerifyCode != "" {
			if req.Code != uVerifyCode {
				resp.Success = errcode.ErrVerifyCodeCompareFailed.Code
				resp.Msg = "login user verify code is not equal failed"
				return nil
			}
		} else {
			resp.Success = errcode.ErrVerifyCodeNotExist.Code
			resp.Msg = "login user verify code is not exist"
			return nil
		}

		if req.OldPassword != "" {
			validate := common.ValidatePassword(teacher.Password, req.OldPassword)
			if !validate {
				resp.Success = errcode.ErrLoginUserPwdFailed.Code
				resp.Msg = "login password failed"
				return nil
			}
		}

		hashPwd, err := common.GeneratePassword(req.NewPassword, 1)
		if err != nil {
			resp.Success = -1
			resp.Msg = "generate user update password failed"
			logger.Logger.Warn("generate user update password failed", zap.Any("user", req), zap.Error(err))
			return err
		}
		teacher.Password = hashPwd
		if err = ul.UpdateTeacherInfo(teacher); err != nil {
			resp.Success = -1
			resp.Msg = "user update password failed"
			logger.Logger.Warn("user update password failed", zap.Any("teacher", teacher), zap.Error(err))
			return err
		}

		// 删除登录的token信息，让用户重新登录
		if err = ul.DelLoginSession(teacher.UID); err != nil {
			logger.Logger.Warn("del loginSession for user update password failed", zap.Any("teacher", teacher), zap.Error(err))
			return err
		}

	} else {
		resp.Success = -1
		resp.Msg = "user update password failed"
		return
	}

	resp.Success = 0
	resp.Msg = "success"

	go ul.DelPhoneCode(req.Phone)

	return nil
}

func (u *UserHandler) SaveUserFollow(ctx context.Context, req *pb.UserFollowOperateReq, resp *pb.Response) (err error) {
	ul := &logic.UserLogic{}
	err = ul.SaveUserFollow(req.Uid, req.FollowUid)
	if err != nil {
		logger.Logger.Warn("get home publish msg failed", zap.Error(err))
		return
	}

	resp.Msg = "success"
	resp.Success = 0
	return nil
}

func (u *UserHandler) CancelUserFollow(ctx context.Context, req *pb.UserFollowOperateReq, resp *pb.Response) (err error) {
	ul := &logic.UserLogic{}
	err = ul.SaveUserFollow(req.Uid, req.FollowUid)
	if err != nil {
		logger.Logger.Warn("get home publish msg failed", zap.Error(err))
		return
	}

	resp.Msg = "success"
	resp.Success = 0
	return nil
}

func (u *UserHandler) QueryUser(ctx context.Context, req *pb.QueryUserReq, resp *pb.QueryUserResp) error {
	ul := &logic.UserLogic{}
	var err error

	resp.UserInfos, err = ul.GetUserInfoFromEsByLike(req.Type, req.ReqStr)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserHandler) GetUserFollowList(ctx context.Context, req *pb.UserFollowListReq, resp *pb.UserFollowOperateResponse) error {
	return nil
}

func (u *UserHandler) GetUserFollowingList(ctx context.Context, req *pb.UserFollowListReq, resp *pb.UserFollowOperateResponse) error {
	return nil
}

func (u *UserHandler) UpdateTeacherInfo(ctx context.Context, req *pb.UpdateTeacherInfoReq, resp *pb.Response) (err error) {
	var teacher *user.UTeacher
	ul := &logic.UserLogic{}

	teacher, err = ul.FindOneTeacherById(req.Id)
	if err != nil {
		resp.Success = -1
		resp.Msg = "select one teacher info by No failed"
		return err
	}
	if teacher == nil {
		resp.Success = errcode.ErrUserNotExist.Code
		resp.Msg = "user No is exist"
		return nil
	}

	if req.ScreenName != "" {
		// 用户名已经存在
		teach, err := ul.FindOneTeacherByScreenName(req.ScreenName)
		if err != nil {
			resp.Success = -1
			resp.Msg = "select one student info by screen_name failed"
			return err
		}
		if teach != nil {
			resp.Success = errcode.ErrRegisterNicknameExist.Code
			resp.Msg = "user screen_name is exist"
			return nil
		}

		teacher.ScreenName = req.ScreenName

	}

	if req.Birthday != "" {
		teacher.Birthday = req.Birthday
	}

	if req.Bio != "" {
		teacher.Bio = req.Bio
	}

	if req.Zodiac != "" {
		teacher.Zodiac = req.Zodiac
	}

	if req.Email != "" {
		teacher.Email = req.Email
	}

	teacher.UpdatedAt = time.Now().Local()

	isVerified := false
	if teacher.IsVerified == 1 {
		isVerified = true
	}

	// 更新动态的用户信息到es
	user := &pb.EsUserInfo{
		Id:            req.Id,
		Username:      teacher.ScreenName,
		ScreenName:    teacher.ScreenName,
		Bio:           teacher.Bio,
		IsVerified:    isVerified,
		VerifyMessage: teacher.VerifyMessage,
		Zodiac:        teacher.Zodiac,
		UniversityId:  int64(teacher.UniversityID),
	}
	if err = ul.UpdateMsgUserInfoToEs(user); err != nil {
		logger.Logger.Warn("insert student to es failed", zap.Any("user", user), zap.Error(err))
		resp.Success = -1
		resp.Msg = "user update failed"
		return err
	}

	// 更新用户基本信息
	if err = ul.UpdateTeacherInfo(teacher); err != nil {
		resp.Success = -1
		resp.Msg = "user update failed"
		logger.Logger.Warn("update student to db failed", zap.Any("teacher", teacher), zap.Error(err))
		return
	}

	return nil
}

func (u *UserHandler) GetTeacherById(ctx context.Context, req *pb.GetTeacherByIdReq, resp *pb.TeacherInfoDetail) error {
	ul := &logic.UserLogic{}
	var err error
	var teacher *user.UTeacher

	teacher, err = ul.FindOneTeacherById(req.Id)
	if err != nil {
		logger.Logger.Warn("select one verify code by phone failed", zap.Any("uid", req.Id), zap.Error(err))
		return err
	}

	if teacher != nil {
		resp.TeachNo = teacher.TeachNo
		resp.ScreenName = teacher.ScreenName
		resp.Gender = teacher.Gender
		resp.Phone = teacher.Phone
		resp.Email = teacher.Email
		resp.Birthday = teacher.Birthday
		resp.Bio = teacher.Bio

		resp.AvatarImage.ThumbnailUrl = teacher.Avatar + "?x-oss-process=image/resize,m_fixed,h_120,w_120"
		resp.AvatarImage.SmallPicUrl = teacher.Avatar + "?x-oss-process=image/resize,m_fixed,h_300,w_300"
		resp.AvatarImage.PicUrl = teacher.Avatar + "?x-oss-process=image/resize,m_fixed,h_800,w_800"

		university, _ := ul.FindUniversityById(teacher.UniversityID)
		resp.University = university.Name

		college, _ := ul.FindCollegeById(teacher.CollegeID)
		resp.College = college.Name

		profession, _ := ul.FindProfessionById(teacher.ProfessionID)
		resp.Profession = profession.Name

	}
	return nil
}

func (u *UserHandler) GetTeacherInfoById(ctx context.Context, req *pb.GetTeacherByIdReq, resp *pb.TeacherInfoDetail) (er error) {
	ul := &logic.UserLogic{}
	var err error
	var teacher *user.UTeacher

	teacher, err = ul.FindOneTeacherById(req.Id)
	if err != nil {
		logger.Logger.Warn("select one verify code by phone failed", zap.Any("uid", req.Id), zap.Error(err))
		return err
	}

	if teacher != nil {
		resp.TeachNo = teacher.TeachNo
		resp.ScreenName = teacher.ScreenName
		resp.Gender = teacher.Gender
		resp.Phone = teacher.Phone
		resp.Email = teacher.Email
		resp.Birthday = teacher.Birthday
		resp.Bio = teacher.Bio

		resp.AvatarImage.ThumbnailUrl = teacher.Avatar + "?x-oss-process=image/resize,m_fixed,h_120,w_120"
		resp.AvatarImage.SmallPicUrl = teacher.Avatar + "?x-oss-process=image/resize,m_fixed,h_300,w_300"
		resp.AvatarImage.PicUrl = teacher.Avatar + "?x-oss-process=image/resize,m_fixed,h_800,w_800"

		university, _ := ul.FindUniversityById(teacher.UniversityID)
		resp.University = university.Name

		college, _ := ul.FindCollegeById(teacher.CollegeID)
		resp.College = college.Name

		profession, _ := ul.FindProfessionById(teacher.ProfessionID)
		resp.Profession = profession.Name
	}

	return nil
}

func (u *UserHandler) GetTeacherListByUniversityId(ctx context.Context, req *pb.GetTeacherListByUniIdReq, resp *pb.TeacherListResp) error {
	return nil
}

func (u *UserHandler) GetTeacherListByCollegeId(ctx context.Context, req *pb.GetTeacherListByCollegeIdReq, resp *pb.TeacherListResp) error {
	return nil
}

func (u *UserHandler) CheckUserIsExist(ctx context.Context, req *pb.CheckUserIsExistReq, resp *pb.Response) (err error) {
	ul := &logic.UserLogic{}
	var student *user.UStudent
	var teacher *user.UTeacher

	// 用户名已经存在
	student, err = ul.FindOneStudentByScreenName(req.Username)
	if err != nil {
		resp.Success = -1
		resp.Msg = "select one student info by username failed"
		return err
	}
	if student != nil {
		resp.Success = errcode.ErrRegisterNicknameExist.Code
		resp.Msg = "user username is exist"
		return nil
	}

	teacher, err = ul.FindOneTeacherByScreenName(req.Username)
	if err != nil {
		resp.Success = -1
		resp.Msg = "select one teacher info by username failed"
		return err
	}
	if teacher != nil {
		resp.Success = errcode.ErrRegisterNicknameExist.Code
		resp.Msg = "user username is exist"
		return nil
	}

	// 手机号已存在
	student, err = ul.FindOneStudentByPhone(req.Phone)
	if err != nil {
		resp.Success = -1
		resp.Msg = "select one student info by No failed"
		return err
	}
	if student != nil {
		resp.Success = errcode.ErrRegisterPhoneExist.Code
		resp.Msg = "user is exist"
		return nil
	}

	teacher, err = ul.FindOneTeacherByPhone(req.Phone)
	if err != nil {
		resp.Success = -1
		resp.Msg = "select one teacher info by No failed"
		return err
	}
	if teacher != nil {
		resp.Success = errcode.ErrRegisterPhoneExist.Code
		resp.Msg = "user is exist"
		return nil
	}

	resp.Success = 0
	resp.Msg = "success"
	return nil
}
