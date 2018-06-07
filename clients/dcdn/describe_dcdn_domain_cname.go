package dcdn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDcdnDomainCnameRequest struct {
	DomainName string `position:"Query" name:"DomainName"`
	OwnerId    int64  `position:"Query" name:"OwnerId"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *DescribeDcdnDomainCnameRequest) Invoke(client goaliyun.Client) (*DescribeDcdnDomainCnameResponse, error) {
	resp := &DescribeDcdnDomainCnameResponse{}
	err := client.Request("dcdn", "DescribeDcdnDomainCname", "2018-01-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDcdnDomainCnameResponse struct {
	RequestId  goaliyun.String
	CnameDatas DescribeDcdnDomainCnameDataList
}

type DescribeDcdnDomainCnameData struct {
	Domain goaliyun.String
	Cname  goaliyun.String
	Status goaliyun.Integer
}

type DescribeDcdnDomainCnameDataList []DescribeDcdnDomainCnameData

func (list *DescribeDcdnDomainCnameDataList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDcdnDomainCnameData)
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
