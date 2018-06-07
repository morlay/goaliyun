package cdn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDomainCCAttackInfoRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeDomainCCAttackInfoRequest) Invoke(client goaliyun.Client) (*DescribeDomainCCAttackInfoResponse, error) {
	resp := &DescribeDomainCCAttackInfoResponse{}
	err := client.Request("cdn", "DescribeDomainCCAttackInfo", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDomainCCAttackInfoResponse struct {
	RequestId           goaliyun.String
	DomainName          goaliyun.String
	StartTime           goaliyun.String
	EndTime             goaliyun.String
	AttackIpDataList    DescribeDomainCCAttackInfoAttackIpDatasList
	AttackedUrlDataList DescribeDomainCCAttackInfoAttackedUrlDatasList
}

type DescribeDomainCCAttackInfoAttackIpDatas struct {
	Ip          goaliyun.String
	AttackCount goaliyun.String
	Result      goaliyun.String
}

type DescribeDomainCCAttackInfoAttackedUrlDatas struct {
	Url         goaliyun.String
	AttackCount goaliyun.String
	Result      goaliyun.String
}

type DescribeDomainCCAttackInfoAttackIpDatasList []DescribeDomainCCAttackInfoAttackIpDatas

func (list *DescribeDomainCCAttackInfoAttackIpDatasList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainCCAttackInfoAttackIpDatas)
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

type DescribeDomainCCAttackInfoAttackedUrlDatasList []DescribeDomainCCAttackInfoAttackedUrlDatas

func (list *DescribeDomainCCAttackInfoAttackedUrlDatasList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainCCAttackInfoAttackedUrlDatas)
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
