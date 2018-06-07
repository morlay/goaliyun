package yundun

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type AddCNameWafRequest struct {
	InstanceId   string `position:"Query" name:"InstanceId"`
	InstanceType string `position:"Query" name:"InstanceType"`
	Domain       string `position:"Query" name:"Domain"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *AddCNameWafRequest) Invoke(client goaliyun.Client) (*AddCNameWafResponse, error) {
	resp := &AddCNameWafResponse{}
	err := client.Request("yundun", "AddCNameWaf", "2015-04-16", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AddCNameWafResponse struct {
	RequestId   goaliyun.String
	WafInfoList AddCNameWafWafInfoList
}

type AddCNameWafWafInfo struct {
	Id     goaliyun.Integer
	Domain goaliyun.String
	Cname  goaliyun.String
	Status goaliyun.Integer
}

type AddCNameWafWafInfoList []AddCNameWafWafInfo

func (list *AddCNameWafWafInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]AddCNameWafWafInfo)
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
