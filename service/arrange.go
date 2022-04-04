package service

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"psy-consult-backend/api"
	"psy-consult-backend/constant"
	"psy-consult-backend/database"
	"psy-consult-backend/exception"
	"psy-consult-backend/utils"
	"psy-consult-backend/utils/helper"
	"psy-consult-backend/utils/sessions"
	"time"
)

func PutArrange(c *gin.Context) {
	// 判断有无管理员权限
	if info := sessions.GetCounsellorInfoBySession(c); info == nil || info.Role != 0 {
		c.Error(exception.AuthError())
		return
	}
	params := make(map[string]interface{})
	c.BindJSON(&params)
	arrangeTime := time.Unix(int64(params["arrangeTime"].(float64)), 0)
	tmpIDs := params["users"].([]interface{})
	counsellorIDs := make([]string, 0)
	for _, counsellorID := range tmpIDs {
		counsellorIDs = append(counsellorIDs, counsellorID.(string))
	}
	counsellorInfoMap, err := database.GetCounsellorUsersByCounsellorIDs(counsellorIDs)
	if err != nil {
		logrus.Error(constant.Service+"PutArrange Failed, err= %v", err)
		c.Error(exception.ServerError())
		return
	}
	success := make([]string, 0)
	for _, counsellorID := range counsellorIDs {
		counsellorInfo, ok := counsellorInfoMap[counsellorID]
		if ok {
			if database.CheckUnique(counsellorInfo.CounsellorID, arrangeTime) == false {
				continue
			}
			err = database.AddArrangement(counsellorInfo.CounsellorID, counsellorInfo.Role, counsellorInfo.Name, arrangeTime)
			if err != nil {
				logrus.Errorf(constant.Service+"PutArrange Failed, err= %v", err)
				c.Error(exception.ServerError())
				return
			}
			success = append(success, counsellorID)
		}
	}
	c.JSON(http.StatusOK, utils.GenSuccessResponse(0, "OK", success))
}

func GetArrange(c *gin.Context) {
	arrangeTime := helper.S2I64(c.Query("arrangeTime"))
	if arrangeTime < 0 || arrangeTime > 3752668800 {
		c.Error(exception.ParameterError())
		return
	}
	arragements, err := database.GetArrangementsByArrangeTime(time.Unix(arrangeTime, 0))
	if err != nil {
		logrus.Errorf(constant.Service+"GetArrange Failed, err= %v", err)
		c.Error(exception.ServerError())
		return
	}
	counsellorIDs := make([]string, 0)
	for _, arragement := range arragements {
		counsellorIDs = append(counsellorIDs, arragement.CounsellorID)
	}
	counsellorID2InfoMap := make(map[string]*database.CounsellorUser, 0)
	if counsellorID2InfoMap, err = database.GetCounsellorUsersByCounsellorIDs(counsellorIDs); err != nil {
		logrus.Errorf(constant.Service+"GetArrange Failed, err= %v", err)
		c.Error(exception.ServerError())
		return
	}
	counsellors := make([]*api.ArrangeResponse, 0)
	supervisors := make([]*api.ArrangeResponse, 0)
	for _, arragement := range arragements {
		counsellor, _ := counsellorID2InfoMap[arragement.CounsellorID]
		var name string
		if counsellor != nil {
			name = counsellor.Name
		}
		t := &api.ArrangeResponse{
			ArrangeID:    arragement.ArrangeID,
			ArrangeTime:  arragement.ArrangeTime,
			Role:         arragement.Role,
			CounsellorID: arragement.CounsellorID,
			Name:         name,
		}
		// 咨询师
		if arragement.Role == 1 {
			counsellors = append(counsellors, t)
		}
		// 督导
		if arragement.Role == 2 {
			supervisors = append(supervisors, t)
		}
	}
	resp := &api.GetArrangeResponse{
		Counsellors: counsellors,
		Supervisors: supervisors,
	}
	c.JSON(http.StatusOK, utils.GenSuccessResponse(0, "OK", resp))
}

func DeleteArrange(c *gin.Context) {
	// 判断有无管理员权限
	if info := sessions.GetCounsellorInfoBySession(c); info == nil || info.Role != 0 {
		c.Error(exception.AuthError())
		return
	}
	params := make(map[string]interface{})
	c.BindJSON(&params)
	arrangeID := int64(params["arrangeID"].(float64))
	err := database.DeleteArrangement(arrangeID)
	if err != nil {
		c.Error(exception.ServerError())
		logrus.Errorf(constant.Service+"DeleteArrange Failed, err= %v", err)
		return
	}
	c.JSON(http.StatusOK, utils.GenSuccessResponse(0, "OK", nil))
}

func GetBatchArrange(c *gin.Context) {
	params := make(map[string]interface{})
	c.BindJSON(&params)
	arrangeTimeList := params["arrangeTimeList"].([]interface{})
	allArrange := make([]*api.GetArrangeResponse, 0)
	req := make([]time.Time, 0)
	for _, v := range arrangeTimeList {
		arrangeTime := int64(v.(float64))
		req = append(req, time.Unix(arrangeTime, 0))
	}
	ArrangeTimeMap, err := database.GetArrangementsByArrangeTimeList(req)
	if err != nil {
		c.Error(exception.ServerError())
		logrus.Errorf(constant.Service+"GetBatchArrange Failed, err= %v", err)
		return
	}

	counsellorID2InfoMap := make(map[string]*database.CounsellorUser)
	counsellorUserList, err := database.GetCounsellorUserList(0, 99999, 0, "")
	if err != nil {
		c.Error(exception.ServerError())
		logrus.Errorf(constant.Service+"GetBatchArrange Failed, err= %v", err)
		return
	}
	for _, c := range counsellorUserList {
		counsellorID2InfoMap[c.CounsellorID] = c
	}

	for _, v := range arrangeTimeList {
		arrangeTime := int64(v.(float64))
		arrangements := ArrangeTimeMap[arrangeTime]
		counsellors := make([]*api.ArrangeResponse, 0)
		supervisors := make([]*api.ArrangeResponse, 0)
		for _, arrangement := range arrangements {
			counsellor, _ := counsellorID2InfoMap[arrangement.CounsellorID]
			var name string
			if counsellor != nil {
				name = counsellor.Name
			}
			t := &api.ArrangeResponse{
				ArrangeID:    arrangement.ArrangeID,
				ArrangeTime:  arrangement.ArrangeTime,
				Role:         arrangement.Role,
				CounsellorID: arrangement.CounsellorID,
				Name:         name,
			}
			// 咨询师
			if arrangement.Role == 1 {
				counsellors = append(counsellors, t)
			}
			// 督导
			if arrangement.Role == 2 {
				supervisors = append(supervisors, t)
			}
		}
		resp := &api.GetArrangeResponse{
			Counsellors: counsellors,
			Supervisors: supervisors,
		}
		allArrange = append(allArrange, resp)
	}
	//for _, v := range arrangeTimeList {
	//	arrangeTime := int64(v.(float64))
	//	arragements, err := database.GetArrangementsByArrangeTime(time.Unix(arrangeTime, 0))
	//	if err != nil {
	//		logrus.Errorf(constant.Service+"GetArrange Failed, err= %v", err)
	//		c.Error(exception.ServerError())
	//		return
	//	}
	//	counsellorIDs := make([]string, 0)
	//	for _, arragement := range arragements {
	//		counsellorIDs = append(counsellorIDs, arragement.CounsellorID)
	//	}
	//	counsellorID2InfoMap := make(map[string]*database.CounsellorUser, 0)
	//	if counsellorID2InfoMap, err = database.GetCounsellorUsersByCounsellorIDs(counsellorIDs); err != nil {
	//		logrus.Errorf(constant.Service+"GetArrange Failed, err= %v", err)
	//		c.Error(exception.ServerError())
	//		return
	//	}
	//	counsellors := make([]*api.ArrangeResponse, 0)
	//	supervisors := make([]*api.ArrangeResponse, 0)
	//	for _, arragement := range arragements {
	//		counsellor, _ := counsellorID2InfoMap[arragement.CounsellorID]
	//		var name string
	//		if counsellor != nil {
	//			name = counsellor.Name
	//		}
	//		t := &api.ArrangeResponse{
	//			ArrangeID:    arragement.ArrangeID,
	//			ArrangeTime:  arragement.ArrangeTime,
	//			Role:         arragement.Role,
	//			CounsellorID: arragement.CounsellorID,
	//			Name:         name,
	//		}
	//		// 咨询师
	//		if arragement.Role == 1 {
	//			counsellors = append(counsellors, t)
	//		}
	//		// 督导
	//		if arragement.Role == 2 {
	//			supervisors = append(supervisors, t)
	//		}
	//	}
	//	resp := &api.GetArrangeResponse{
	//		Counsellors: counsellors,
	//		Supervisors: supervisors,
	//	}
	//	allArrange = append(allArrange, resp)
	//}
	c.JSON(http.StatusOK, utils.GenSuccessResponse(0, "OK", allArrange))
}
