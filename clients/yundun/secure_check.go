package yundun

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type SecureCheckRequest struct {
	JstOwnerId  int64  `position:"Query" name:"JstOwnerId"`
	InstanceIds string `position:"Query" name:"InstanceIds"`
	RegionId    string `position:"Query" name:"RegionId"`
}

func (req *SecureCheckRequest) Invoke(client goaliyun.Client) (*SecureCheckResponse, error) {
	resp := &SecureCheckResponse{}
	err := client.Request("yundun", "SecureCheck", "2015-04-16", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SecureCheckResponse struct {
	RequestId        goaliyun.String
	RecentInstanceId goaliyun.String
	ProblemList      SecureCheckInfoList
	NoProblemList    SecureCheckInfoList
	NoScanList       SecureCheckInfoList
	ScanningList     SecureCheckInfoList
	InnerIpList      SecureCheckInfoList
}

type SecureCheckInfo struct {
	Ip         goaliyun.String
	Status     goaliyun.String
	VulNum     goaliyun.String
	InstanceId goaliyun.String
}

type SecureCheckInfoList []SecureCheckInfo

func (list *SecureCheckInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SecureCheckInfo)
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
