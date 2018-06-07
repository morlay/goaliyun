package cdn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeUserVipsByDomainRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	Available     string `position:"Query" name:"Available"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeUserVipsByDomainRequest) Invoke(client goaliyun.Client) (*DescribeUserVipsByDomainResponse, error) {
	resp := &DescribeUserVipsByDomainResponse{}
	err := client.Request("cdn", "DescribeUserVipsByDomain", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeUserVipsByDomainResponse struct {
	RequestId  goaliyun.String
	DomainName goaliyun.Integer
	Vips       DescribeUserVipsByDomainVipList
}

type DescribeUserVipsByDomainVipList []goaliyun.String

func (list *DescribeUserVipsByDomainVipList) UnmarshalJSON(data []byte) error {
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
