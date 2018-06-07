package cdn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDomainCnameRequest struct {
	DomainName string `position:"Query" name:"DomainName"`
	OwnerId    int64  `position:"Query" name:"OwnerId"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *DescribeDomainCnameRequest) Invoke(client goaliyun.Client) (*DescribeDomainCnameResponse, error) {
	resp := &DescribeDomainCnameResponse{}
	err := client.Request("cdn", "DescribeDomainCname", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDomainCnameResponse struct {
	RequestId  goaliyun.String
	CnameDatas DescribeDomainCnameDataList
}

type DescribeDomainCnameData struct {
	Domain goaliyun.String
	Cname  goaliyun.String
	Status goaliyun.Integer
}

type DescribeDomainCnameDataList []DescribeDomainCnameData

func (list *DescribeDomainCnameDataList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainCnameData)
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
