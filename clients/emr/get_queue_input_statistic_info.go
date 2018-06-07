package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type GetQueueInputStatisticInfoRequest struct {
	FromDatetime    string `position:"Query" name:"FromDatetime"`
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	ToDatetime      string `position:"Query" name:"ToDatetime"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *GetQueueInputStatisticInfoRequest) Invoke(client goaliyun.Client) (*GetQueueInputStatisticInfoResponse, error) {
	resp := &GetQueueInputStatisticInfoResponse{}
	err := client.Request("emr", "GetQueueInputStatisticInfo", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetQueueInputStatisticInfoResponse struct {
	RequestId      goaliyun.String
	QueueInputList GetQueueInputStatisticInfoClusterStatQueueInputList
}

type GetQueueInputStatisticInfoClusterStatQueueInput struct {
	Queue      goaliyun.String
	BytesInput goaliyun.Integer
}

type GetQueueInputStatisticInfoClusterStatQueueInputList []GetQueueInputStatisticInfoClusterStatQueueInput

func (list *GetQueueInputStatisticInfoClusterStatQueueInputList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetQueueInputStatisticInfoClusterStatQueueInput)
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
