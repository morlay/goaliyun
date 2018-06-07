package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type GetHdfsCapacityStatisticInfoRequest struct {
	FromDatetime    string `position:"Query" name:"FromDatetime"`
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	ToDatetime      string `position:"Query" name:"ToDatetime"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *GetHdfsCapacityStatisticInfoRequest) Invoke(client goaliyun.Client) (*GetHdfsCapacityStatisticInfoResponse, error) {
	resp := &GetHdfsCapacityStatisticInfoResponse{}
	err := client.Request("emr", "GetHdfsCapacityStatisticInfo", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetHdfsCapacityStatisticInfoResponse struct {
	RequestId        goaliyun.String
	HdfsCapacityList GetHdfsCapacityStatisticInfoClusterStatHdfsCapacityList
}

type GetHdfsCapacityStatisticInfoClusterStatHdfsCapacity struct {
	CapacityTotal       goaliyun.Integer
	CapacityTotalGB     goaliyun.Integer
	CapacityUsed        goaliyun.Integer
	CapacityUsedGB      goaliyun.Integer
	CapacityRemaining   goaliyun.Integer
	CapacityRemainingGB goaliyun.Integer
	CapacityUsedNonDfs  goaliyun.Integer
	ClusterBizId        goaliyun.String
	DateTime            goaliyun.String
}

type GetHdfsCapacityStatisticInfoClusterStatHdfsCapacityList []GetHdfsCapacityStatisticInfoClusterStatHdfsCapacity

func (list *GetHdfsCapacityStatisticInfoClusterStatHdfsCapacityList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetHdfsCapacityStatisticInfoClusterStatHdfsCapacity)
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
