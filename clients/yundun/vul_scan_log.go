package yundun

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type VulScanLogRequest struct {
	JstOwnerId int64  `position:"Query" name:"JstOwnerId"`
	PageNumber int64  `position:"Query" name:"PageNumber"`
	PageSize   int64  `position:"Query" name:"PageSize"`
	InstanceId string `position:"Query" name:"InstanceId"`
	VulStatus  int64  `position:"Query" name:"VulStatus"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *VulScanLogRequest) Invoke(client goaliyun.Client) (*VulScanLogResponse, error) {
	resp := &VulScanLogResponse{}
	err := client.Request("yundun", "VulScanLog", "2015-04-16", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type VulScanLogResponse struct {
	RequestId  goaliyun.String
	StartTime  goaliyun.String
	EndTime    goaliyun.String
	PageNumber goaliyun.Integer
	PageSize   goaliyun.Integer
	TotalCount goaliyun.Integer
	LogList    VulScanLogVulScanLogList
}

type VulScanLogVulScanLog struct {
	Id           goaliyun.Integer
	Type         goaliyun.String
	Url          goaliyun.String
	HelpAddress  goaliyun.String
	VulParameter goaliyun.String
	Status       goaliyun.Integer
}

type VulScanLogVulScanLogList []VulScanLogVulScanLog

func (list *VulScanLogVulScanLogList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]VulScanLogVulScanLog)
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
