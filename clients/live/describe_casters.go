package live

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeCastersRequest struct {
	CasterName string `position:"Query" name:"CasterName"`
	CasterId   string `position:"Query" name:"CasterId"`
	PageSize   int64  `position:"Query" name:"PageSize"`
	EndTime    string `position:"Query" name:"EndTime"`
	StartTime  string `position:"Query" name:"StartTime"`
	OwnerId    int64  `position:"Query" name:"OwnerId"`
	PageNum    int64  `position:"Query" name:"PageNum"`
	Status     int64  `position:"Query" name:"Status"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *DescribeCastersRequest) Invoke(client goaliyun.Client) (*DescribeCastersResponse, error) {
	resp := &DescribeCastersResponse{}
	err := client.Request("live", "DescribeCasters", "2016-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeCastersResponse struct {
	RequestId  goaliyun.String
	Total      goaliyun.Integer
	CasterList DescribeCastersCasterList
}

type DescribeCastersCaster struct {
	Status         goaliyun.Integer
	NormType       goaliyun.Integer
	CasterId       goaliyun.String
	CasterName     goaliyun.String
	CreateTime     goaliyun.String
	StartTime      goaliyun.String
	PurchaseTime   goaliyun.String
	ExpireTime     goaliyun.String
	ChargeType     goaliyun.String
	CasterTemplate goaliyun.String
}

type DescribeCastersCasterList []DescribeCastersCaster

func (list *DescribeCastersCasterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCastersCaster)
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
