package yundun

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DeleteCNameWafRequest struct {
	InstanceId   string `position:"Query" name:"InstanceId"`
	Domain       string `position:"Query" name:"Domain"`
	CnameId      int64  `position:"Query" name:"CnameId"`
	InstanceType string `position:"Query" name:"InstanceType"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *DeleteCNameWafRequest) Invoke(client goaliyun.Client) (*DeleteCNameWafResponse, error) {
	resp := &DeleteCNameWafResponse{}
	err := client.Request("yundun", "DeleteCNameWaf", "2015-04-16", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteCNameWafResponse struct {
	RequestId   goaliyun.String
	WafInfoList DeleteCNameWafWafInfoList
}

type DeleteCNameWafWafInfo struct {
	Id     goaliyun.Integer
	Domain goaliyun.String
	Cname  goaliyun.String
	Status goaliyun.Integer
}

type DeleteCNameWafWafInfoList []DeleteCNameWafWafInfo

func (list *DeleteCNameWafWafInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DeleteCNameWafWafInfo)
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
