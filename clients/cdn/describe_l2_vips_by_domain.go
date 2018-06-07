package cdn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeL2VipsByDomainRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeL2VipsByDomainRequest) Invoke(client goaliyun.Client) (*DescribeL2VipsByDomainResponse, error) {
	resp := &DescribeL2VipsByDomainResponse{}
	err := client.Request("cdn", "DescribeL2VipsByDomain", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeL2VipsByDomainResponse struct {
	RequestId  goaliyun.String
	DomainName goaliyun.String
	Vips       DescribeL2VipsByDomainVipList
}

type DescribeL2VipsByDomainVipList []goaliyun.String

func (list *DescribeL2VipsByDomainVipList) UnmarshalJSON(data []byte) error {
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
