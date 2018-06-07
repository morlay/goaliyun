package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ResetAutoRenewalRequest struct {
	ResourceOwnerId      int64                                    `position:"Query" name:"ResourceOwnerId"`
	ClusterId            string                                   `position:"Query" name:"ClusterId"`
	EcsResetAutoRenewDos *ResetAutoRenewalEcsResetAutoRenewDoList `position:"Query" type:"Repeated" name:"EcsResetAutoRenewDo"`
	RegionId             string                                   `position:"Query" name:"RegionId"`
}

func (req *ResetAutoRenewalRequest) Invoke(client goaliyun.Client) (*ResetAutoRenewalResponse, error) {
	resp := &ResetAutoRenewalResponse{}
	err := client.Request("emr", "ResetAutoRenewal", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ResetAutoRenewalEcsResetAutoRenewDo struct {
	EcsId                string `name:"EcsId"`
	EcsAutoRenew         string `name:"EcsAutoRenew"`
	EcsAutoRenewDuration string `name:"EcsAutoRenewDuration"`
	EmrAutoRenew         string `name:"EmrAutoRenew"`
	EmrAutoRenewDuration string `name:"EmrAutoRenewDuration"`
}

type ResetAutoRenewalResponse struct {
	RequestId goaliyun.String
}

type ResetAutoRenewalEcsResetAutoRenewDoList []ResetAutoRenewalEcsResetAutoRenewDo

func (list *ResetAutoRenewalEcsResetAutoRenewDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ResetAutoRenewalEcsResetAutoRenewDo)
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
