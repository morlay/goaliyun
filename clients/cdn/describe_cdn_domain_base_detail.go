package cdn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeCdnDomainBaseDetailRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeCdnDomainBaseDetailRequest) Invoke(client goaliyun.Client) (*DescribeCdnDomainBaseDetailResponse, error) {
	resp := &DescribeCdnDomainBaseDetailResponse{}
	err := client.Request("cdn", "DescribeCdnDomainBaseDetail", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeCdnDomainBaseDetailResponse struct {
	RequestId             goaliyun.String
	DomainBaseDetailModel DescribeCdnDomainBaseDetailDomainBaseDetailModel
}

type DescribeCdnDomainBaseDetailDomainBaseDetailModel struct {
	Cname        goaliyun.String
	CdnType      goaliyun.String
	DomainStatus goaliyun.String
	SourceType   goaliyun.String
	Region       goaliyun.String
	DomainName   goaliyun.String
	Remark       goaliyun.String
	GmtModified  goaliyun.String
	GmtCreated   goaliyun.String
	Sources      DescribeCdnDomainBaseDetailSourceList
}

type DescribeCdnDomainBaseDetailSourceList []goaliyun.String

func (list *DescribeCdnDomainBaseDetailSourceList) UnmarshalJSON(data []byte) error {
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
