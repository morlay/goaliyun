package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type GetJobRunningTimeStatisticInfoRequest struct {
	FromDatetime    string `position:"Query" name:"FromDatetime"`
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	PageSize        int64  `position:"Query" name:"PageSize"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	ToDatetime      string `position:"Query" name:"ToDatetime"`
	PageNumber      int64  `position:"Query" name:"PageNumber"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *GetJobRunningTimeStatisticInfoRequest) Invoke(client goaliyun.Client) (*GetJobRunningTimeStatisticInfoResponse, error) {
	resp := &GetJobRunningTimeStatisticInfoResponse{}
	err := client.Request("emr", "GetJobRunningTimeStatisticInfo", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetJobRunningTimeStatisticInfoResponse struct {
	RequestId       goaliyun.String
	Total           goaliyun.Integer
	PageNumber      goaliyun.Integer
	PageSize        goaliyun.Integer
	RunningTimeList GetJobRunningTimeStatisticInfoClusterStatJobRunningTimeList
}

type GetJobRunningTimeStatisticInfoClusterStatJobRunningTime struct {
	ApplicationId goaliyun.String
	JobId         goaliyun.String
	StartTime     goaliyun.Integer
	FinishTime    goaliyun.Integer
	Name          goaliyun.String
	Queue         goaliyun.String
	User          goaliyun.String
	State         goaliyun.String
	RunningTime   goaliyun.Integer
}

type GetJobRunningTimeStatisticInfoClusterStatJobRunningTimeList []GetJobRunningTimeStatisticInfoClusterStatJobRunningTime

func (list *GetJobRunningTimeStatisticInfoClusterStatJobRunningTimeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetJobRunningTimeStatisticInfoClusterStatJobRunningTime)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}
