package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type GetUserOutputStatisticInfoRequest struct {
	FromDatetime    string `position:"Query" name:"FromDatetime"`
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	ToDatetime      string `position:"Query" name:"ToDatetime"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *GetUserOutputStatisticInfoRequest) Invoke(client goaliyun.Client) (*GetUserOutputStatisticInfoResponse, error) {
	resp := &GetUserOutputStatisticInfoResponse{}
	err := client.Request("emr", "GetUserOutputStatisticInfo", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetUserOutputStatisticInfoResponse struct {
	RequestId      goaliyun.String
	UserOutputList GetUserOutputStatisticInfoClusterStatUserOutputList
}

type GetUserOutputStatisticInfoClusterStatUserOutput struct {
	User        goaliyun.String
	BytesOutput goaliyun.Integer
}

type GetUserOutputStatisticInfoClusterStatUserOutputList []GetUserOutputStatisticInfoClusterStatUserOutput

func (list *GetUserOutputStatisticInfoClusterStatUserOutputList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetUserOutputStatisticInfoClusterStatUserOutput)
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
