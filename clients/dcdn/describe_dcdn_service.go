package dcdn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDcdnServiceRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeDcdnServiceRequest) Invoke(client goaliyun.Client) (*DescribeDcdnServiceResponse, error) {
	resp := &DescribeDcdnServiceResponse{}
	err := client.Request("dcdn", "DescribeDcdnService", "2018-01-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDcdnServiceResponse struct {
	RequestId          goaliyun.String
	InstanceId         goaliyun.String
	InternetChargeType goaliyun.String
	OpeningTime        goaliyun.String
	ChangingChargeType goaliyun.String
	ChangingAffectTime goaliyun.String
	OperationLocks     DescribeDcdnServiceLockReasonList
}

type DescribeDcdnServiceLockReason struct {
	LockReason goaliyun.String
}

type DescribeDcdnServiceLockReasonList []DescribeDcdnServiceLockReason

func (list *DescribeDcdnServiceLockReasonList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDcdnServiceLockReason)
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
