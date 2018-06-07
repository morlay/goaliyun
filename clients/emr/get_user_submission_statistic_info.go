package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type GetUserSubmissionStatisticInfoRequest struct {
	FromDatetime    string `position:"Query" name:"FromDatetime"`
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	ToDatetime      string `position:"Query" name:"ToDatetime"`
	ApplicationType string `position:"Query" name:"ApplicationType"`
	FinalStatus     string `position:"Query" name:"FinalStatus"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *GetUserSubmissionStatisticInfoRequest) Invoke(client goaliyun.Client) (*GetUserSubmissionStatisticInfoResponse, error) {
	resp := &GetUserSubmissionStatisticInfoResponse{}
	err := client.Request("emr", "GetUserSubmissionStatisticInfo", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetUserSubmissionStatisticInfoResponse struct {
	RequestId          goaliyun.String
	UserSubmissionList GetUserSubmissionStatisticInfoClusterStatUserSubmissionList
}

type GetUserSubmissionStatisticInfoClusterStatUserSubmission struct {
	User       goaliyun.String
	Submission goaliyun.Integer
}

type GetUserSubmissionStatisticInfoClusterStatUserSubmissionList []GetUserSubmissionStatisticInfoClusterStatUserSubmission

func (list *GetUserSubmissionStatisticInfoClusterStatUserSubmissionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetUserSubmissionStatisticInfoClusterStatUserSubmission)
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
