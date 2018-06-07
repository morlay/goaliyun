package r_kvstore

import (
	"github.com/morlay/goaliyun"
)

type SwitchNetworkRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	VSwitchId            string `position:"Query" name:"VSwitchId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	TargetNetworkType    string `position:"Query" name:"TargetNetworkType"`
	RetainClassic        string `position:"Query" name:"RetainClassic"`
	ClassicExpiredDays   string `position:"Query" name:"ClassicExpiredDays"`
	VpcId                string `position:"Query" name:"VpcId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *SwitchNetworkRequest) Invoke(client goaliyun.Client) (*SwitchNetworkResponse, error) {
	resp := &SwitchNetworkResponse{}
	err := client.Request("r-kvstore", "SwitchNetwork", "2015-01-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SwitchNetworkResponse struct {
	RequestId goaliyun.String
	TaskId    goaliyun.String
}
