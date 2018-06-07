package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type GetQueueOutputStatisticInfoRequest struct {
	FromDatetime    string `position:"Query" name:"FromDatetime"`
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	ToDatetime      string `position:"Query" name:"ToDatetime"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *GetQueueOutputStatisticInfoRequest) Invoke(client goaliyun.Client) (*GetQueueOutputStatisticInfoResponse, error) {
	resp := &GetQueueOutputStatisticInfoResponse{}
	err := client.Request("emr", "GetQueueOutputStatisticInfo", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetQueueOutputStatisticInfoResponse struct {
	RequestId       goaliyun.String
	QueueOutputList GetQueueOutputStatisticInfoClusterStatQueueOutputList
}

type GetQueueOutputStatisticInfoClusterStatQueueOutput struct {
	Queue       goaliyun.String
	BytesOutput goaliyun.Integer
}

type GetQueueOutputStatisticInfoClusterStatQueueOutputList []GetQueueOutputStatisticInfoClusterStatQueueOutput

func (list *GetQueueOutputStatisticInfoClusterStatQueueOutputList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetQueueOutputStatisticInfoClusterStatQueueOutput)
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
