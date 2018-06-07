package cdn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeCdnServiceRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeCdnServiceRequest) Invoke(client goaliyun.Client) (*DescribeCdnServiceResponse, error) {
	resp := &DescribeCdnServiceResponse{}
	err := client.Request("cdn", "DescribeCdnService", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeCdnServiceResponse struct {
	RequestId          goaliyun.String
	InstanceId         goaliyun.String
	InternetChargeType goaliyun.String
	OpeningTime        goaliyun.String
	ChangingChargeType goaliyun.String
	ChangingAffectTime goaliyun.String
	OperationLocks     DescribeCdnServiceLockReasonList
}

type DescribeCdnServiceLockReason struct {
	LockReason goaliyun.String
}

type DescribeCdnServiceLockReasonList []DescribeCdnServiceLockReason

func (list *DescribeCdnServiceLockReasonList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCdnServiceLockReason)
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
