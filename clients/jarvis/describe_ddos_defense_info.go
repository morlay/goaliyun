package jarvis

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDdosDefenseInfoRequest struct {
	SourceIp   string `position:"Query" name:"SourceIp"`
	Lang       string `position:"Query" name:"Lang"`
	SrcUid     int64  `position:"Query" name:"SrcUid"`
	SourceCode string `position:"Query" name:"SourceCode"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *DescribeDdosDefenseInfoRequest) Invoke(client goaliyun.Client) (*DescribeDdosDefenseInfoResponse, error) {
	resp := &DescribeDdosDefenseInfoResponse{}
	err := client.Request("jarvis", "DescribeDdosDefenseInfo", "2018-02-06", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDdosDefenseInfoResponse struct {
	RequestId            goaliyun.String
	Module               goaliyun.String
	BlackTimes           goaliyun.Integer
	Duration             goaliyun.Integer
	BgpPkgState          goaliyun.String
	DdosDefenseThreshold DescribeDdosDefenseInfoDdosDefenseThresholdItemList
}

type DescribeDdosDefenseInfoDdosDefenseThresholdItem struct {
	RegionId           goaliyun.String
	CurrThreshold      goaliyun.Integer
	RecommendThreshold goaliyun.Integer
}

type DescribeDdosDefenseInfoDdosDefenseThresholdItemList []DescribeDdosDefenseInfoDdosDefenseThresholdItem

func (list *DescribeDdosDefenseInfoDdosDefenseThresholdItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDdosDefenseInfoDdosDefenseThresholdItem)
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
