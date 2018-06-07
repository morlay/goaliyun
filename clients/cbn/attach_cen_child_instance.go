package cbn

import (
	"github.com/morlay/goaliyun"
)

type AttachCenChildInstanceRequest struct {
	ChildInstanceId       string `position:"Query" name:"ChildInstanceId"`
	ResourceOwnerId       int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount  string `position:"Query" name:"ResourceOwnerAccount"`
	CenId                 string `position:"Query" name:"CenId"`
	OwnerAccount          string `position:"Query" name:"OwnerAccount"`
	OwnerId               int64  `position:"Query" name:"OwnerId"`
	ChildInstanceType     string `position:"Query" name:"ChildInstanceType"`
	ChildInstanceOwnerId  int64  `position:"Query" name:"ChildInstanceOwnerId"`
	ChildInstanceRegionId string `position:"Query" name:"ChildInstanceRegionId"`
	RegionId              string `position:"Query" name:"RegionId"`
}

func (req *AttachCenChildInstanceRequest) Invoke(client goaliyun.Client) (*AttachCenChildInstanceResponse, error) {
	resp := &AttachCenChildInstanceResponse{}
	err := client.Request("cbn", "AttachCenChildInstance", "2017-09-12", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AttachCenChildInstanceResponse struct {
	RequestId goaliyun.String
}
