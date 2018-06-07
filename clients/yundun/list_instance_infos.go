package yundun

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListInstanceInfosRequest struct {
	JstOwnerId   int64  `position:"Query" name:"JstOwnerId"`
	PageNumber   int64  `position:"Query" name:"PageNumber"`
	PageSize     int64  `position:"Query" name:"PageSize"`
	Region       string `position:"Query" name:"Region"`
	EventType    string `position:"Query" name:"EventType"`
	InstanceName string `position:"Query" name:"InstanceName"`
	InstanceType string `position:"Query" name:"InstanceType"`
	InstanceIds  string `position:"Query" name:"InstanceIds"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *ListInstanceInfosRequest) Invoke(client goaliyun.Client) (*ListInstanceInfosResponse, error) {
	resp := &ListInstanceInfosResponse{}
	err := client.Request("yundun", "ListInstanceInfos", "2015-04-16", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListInstanceInfosResponse struct {
	RequestId  goaliyun.String
	PageNumber goaliyun.Integer
	PageSize   goaliyun.Integer
	TotalCount goaliyun.Integer
	InfosList  ListInstanceInfosInstanceInfoList
}

type ListInstanceInfosInstanceInfo struct {
	Region       goaliyun.String
	RegionName   goaliyun.String
	RegionEnName goaliyun.String
	InstanceName goaliyun.String
	InstanceId   goaliyun.String
	Ip           goaliyun.String
	InternetIp   goaliyun.String
	IntranetIp   goaliyun.String
	Ddos         goaliyun.Integer
	HostEvent    goaliyun.Integer
	SecureCheck  goaliyun.Integer
	AegisStatus  goaliyun.Integer
	Waf          goaliyun.Integer
	IsLock       bool
	LockType     goaliyun.String
	UnLockTimes  goaliyun.Integer
	TriggerTime  goaliyun.String
}

type ListInstanceInfosInstanceInfoList []ListInstanceInfosInstanceInfo

func (list *ListInstanceInfosInstanceInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListInstanceInfosInstanceInfo)
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
