package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type GetUserInputStatisticInfoRequest struct {
	FromDatetime    string `position:"Query" name:"FromDatetime"`
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	ToDatetime      string `position:"Query" name:"ToDatetime"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *GetUserInputStatisticInfoRequest) Invoke(client goaliyun.Client) (*GetUserInputStatisticInfoResponse, error) {
	resp := &GetUserInputStatisticInfoResponse{}
	err := client.Request("emr", "GetUserInputStatisticInfo", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetUserInputStatisticInfoResponse struct {
	RequestId     goaliyun.String
	UserInputList GetUserInputStatisticInfoClusterStatUserInputList
}

type GetUserInputStatisticInfoClusterStatUserInput struct {
	User       goaliyun.String
	BytesInput goaliyun.Integer
}

type GetUserInputStatisticInfoClusterStatUserInputList []GetUserInputStatisticInfoClusterStatUserInput

func (list *GetUserInputStatisticInfoClusterStatUserInputList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetUserInputStatisticInfoClusterStatUserInput)
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
