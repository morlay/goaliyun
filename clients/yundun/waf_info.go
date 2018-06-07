package yundun

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type WafInfoRequest struct {
	InstanceId   string `position:"Query" name:"InstanceId"`
	InstanceType string `position:"Query" name:"InstanceType"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *WafInfoRequest) Invoke(client goaliyun.Client) (*WafInfoResponse, error) {
	resp := &WafInfoResponse{}
	err := client.Request("yundun", "WafInfo", "2015-04-16", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type WafInfoResponse struct {
	RequestId    goaliyun.String
	WafDomainNum goaliyun.Integer
	WafInfos     WafInfoWafInfoList
}

type WafInfoWafInfo struct {
	Id     goaliyun.Integer
	Domain goaliyun.String
	Cname  goaliyun.String
	Status goaliyun.Integer
}

type WafInfoWafInfoList []WafInfoWafInfo

func (list *WafInfoWafInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]WafInfoWafInfo)
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
