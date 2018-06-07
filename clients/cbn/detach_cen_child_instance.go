package cbn

import (
	"github.com/morlay/goaliyun"
)

type DetachCenChildInstanceRequest struct {
	ChildInstanceId       string `position:"Query" name:"ChildInstanceId"`
	ResourceOwnerId       int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount  string `position:"Query" name:"ResourceOwnerAccount"`
	CenId                 string `position:"Query" name:"CenId"`
	OwnerAccount          string `position:"Query" name:"OwnerAccount"`
	CenOwnerId            int64  `position:"Query" name:"CenOwnerId"`
	OwnerId               int64  `position:"Query" name:"OwnerId"`
	ChildInstanceType     string `position:"Query" name:"ChildInstanceType"`
	ChildInstanceOwnerId  int64  `position:"Query" name:"ChildInstanceOwnerId"`
	ChildInstanceRegionId string `position:"Query" name:"ChildInstanceRegionId"`
	RegionId              string `position:"Query" name:"RegionId"`
}

func (req *DetachCenChildInstanceRequest) Invoke(client goaliyun.Client) (*DetachCenChildInstanceResponse, error) {
	resp := &DetachCenChildInstanceResponse{}
	err := client.Request("cbn", "DetachCenChildInstance", "2017-09-12", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DetachCenChildInstanceResponse struct {
	RequestId goaliyun.String
}
