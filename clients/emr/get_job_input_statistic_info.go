package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type GetJobInputStatisticInfoRequest struct {
	FromDatetime    string `position:"Query" name:"FromDatetime"`
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	PageSize        int64  `position:"Query" name:"PageSize"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	ToDatetime      string `position:"Query" name:"ToDatetime"`
	PageNumber      int64  `position:"Query" name:"PageNumber"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *GetJobInputStatisticInfoRequest) Invoke(client goaliyun.Client) (*GetJobInputStatisticInfoResponse, error) {
	resp := &GetJobInputStatisticInfoResponse{}
	err := client.Request("emr", "GetJobInputStatisticInfo", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetJobInputStatisticInfoResponse struct {
	RequestId    goaliyun.String
	Total        goaliyun.Integer
	PageNumber   goaliyun.Integer
	PageSize     goaliyun.Integer
	JobInputList GetJobInputStatisticInfoClusterStatJobInputList
}

type GetJobInputStatisticInfoClusterStatJobInput struct {
	ApplicationId goaliyun.String
	JobId         goaliyun.String
	StartTime     goaliyun.Integer
	FinishTime    goaliyun.Integer
	Name          goaliyun.String
	Queue         goaliyun.String
	User          goaliyun.String
	State         goaliyun.String
	BytesInput    goaliyun.Integer
}

type GetJobInputStatisticInfoClusterStatJobInputList []GetJobInputStatisticInfoClusterStatJobInput

func (list *GetJobInputStatisticInfoClusterStatJobInputList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetJobInputStatisticInfoClusterStatJobInput)
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
