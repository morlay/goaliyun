package vpc

import (
	"github.com/morlay/goaliyun"
)

type CreateGlobalAccelerationInstanceRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	BandwidthType        string `position:"Query" name:"BandwidthType"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ServiceLocation      string `position:"Query" name:"ServiceLocation"`
	Bandwidth            string `position:"Query" name:"Bandwidth"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Description          string `position:"Query" name:"Description"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	InternetChargeType   string `position:"Query" name:"InternetChargeType"`
	Name                 string `position:"Query" name:"Name"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *CreateGlobalAccelerationInstanceRequest) Invoke(client goaliyun.Client) (*CreateGlobalAccelerationInstanceResponse, error) {
	resp := &CreateGlobalAccelerationInstanceResponse{}
	err := client.Request("vpc", "CreateGlobalAccelerationInstance", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateGlobalAccelerationInstanceResponse struct {
	RequestId                    goaliyun.String
	GlobalAccelerationInstanceId goaliyun.String
	IpAddress                    goaliyun.String
}
