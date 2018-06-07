package ccc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type GetServiceExtensionsRequest struct {
	ServiceType string `position:"Query" name:"ServiceType"`
	InstanceId  string `position:"Query" name:"InstanceId"`
	RegionId    string `position:"Query" name:"RegionId"`
}

func (req *GetServiceExtensionsRequest) Invoke(client goaliyun.Client) (*GetServiceExtensionsResponse, error) {
	resp := &GetServiceExtensionsResponse{}
	err := client.Request("ccc", "GetServiceExtensions", "2017-07-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetServiceExtensionsResponse struct {
	RequestId         goaliyun.String
	Success           bool
	Code              goaliyun.String
	Message           goaliyun.String
	HttpStatusCode    goaliyun.Integer
	ServiceExtensions GetServiceExtensionsServiceExtensionList
}

type GetServiceExtensionsServiceExtension struct {
	Name   goaliyun.String
	Number goaliyun.String
}

type GetServiceExtensionsServiceExtensionList []GetServiceExtensionsServiceExtension

func (list *GetServiceExtensionsServiceExtensionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetServiceExtensionsServiceExtension)
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
