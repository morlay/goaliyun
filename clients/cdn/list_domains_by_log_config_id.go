package cdn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListDomainsByLogConfigIdRequest struct {
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	Version       string `position:"Query" name:"Version"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	ConfigId      string `position:"Query" name:"ConfigId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *ListDomainsByLogConfigIdRequest) Invoke(client goaliyun.Client) (*ListDomainsByLogConfigIdResponse, error) {
	resp := &ListDomainsByLogConfigIdResponse{}
	err := client.Request("cdn", "ListDomainsByLogConfigId", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListDomainsByLogConfigIdResponse struct {
	RequestId goaliyun.String
	Domains   ListDomainsByLogConfigIdDomainList
}

type ListDomainsByLogConfigIdDomainList []goaliyun.String

func (list *ListDomainsByLogConfigIdDomainList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]goaliyun.String)
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
