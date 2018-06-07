package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type GetQueueSubmissionStatisticInfoRequest struct {
	FromDatetime    string `position:"Query" name:"FromDatetime"`
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	ToDatetime      string `position:"Query" name:"ToDatetime"`
	ApplicationType string `position:"Query" name:"ApplicationType"`
	FinalStatus     string `position:"Query" name:"FinalStatus"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *GetQueueSubmissionStatisticInfoRequest) Invoke(client goaliyun.Client) (*GetQueueSubmissionStatisticInfoResponse, error) {
	resp := &GetQueueSubmissionStatisticInfoResponse{}
	err := client.Request("emr", "GetQueueSubmissionStatisticInfo", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetQueueSubmissionStatisticInfoResponse struct {
	RequestId           goaliyun.String
	QueueSubmissionList GetQueueSubmissionStatisticInfoClusterStatQueueSubmissionList
}

type GetQueueSubmissionStatisticInfoClusterStatQueueSubmission struct {
	Queue      goaliyun.String
	Submission goaliyun.Integer
}

type GetQueueSubmissionStatisticInfoClusterStatQueueSubmissionList []GetQueueSubmissionStatisticInfoClusterStatQueueSubmission

func (list *GetQueueSubmissionStatisticInfoClusterStatQueueSubmissionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetQueueSubmissionStatisticInfoClusterStatQueueSubmission)
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
