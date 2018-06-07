package cdn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDomainBpsDataByTimeStampRequest struct {
	IspNames      string `position:"Query" name:"IspNames"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	LocationNames string `position:"Query" name:"LocationNames"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	TimePoint     string `position:"Query" name:"TimePoint"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeDomainBpsDataByTimeStampRequest) Invoke(client goaliyun.Client) (*DescribeDomainBpsDataByTimeStampResponse, error) {
	resp := &DescribeDomainBpsDataByTimeStampResponse{}
	err := client.Request("cdn", "DescribeDomainBpsDataByTimeStamp", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDomainBpsDataByTimeStampResponse struct {
	RequestId   goaliyun.String
	DomainName  goaliyun.String
	TimeStamp   goaliyun.String
	BpsDataList DescribeDomainBpsDataByTimeStampBpsDataModelList
}

type DescribeDomainBpsDataByTimeStampBpsDataModel struct {
	LocationName goaliyun.String
	IspName      goaliyun.String
	Bps          goaliyun.Integer
}

type DescribeDomainBpsDataByTimeStampBpsDataModelList []DescribeDomainBpsDataByTimeStampBpsDataModel

func (list *DescribeDomainBpsDataByTimeStampBpsDataModelList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainBpsDataByTimeStampBpsDataModel)
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
