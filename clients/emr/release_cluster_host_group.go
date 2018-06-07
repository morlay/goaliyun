package emr

import (
	"github.com/morlay/goaliyun"
)

type ReleaseClusterHostGroupRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	HostGroupId     string `position:"Query" name:"HostGroupId"`
	InstanceIdList  string `position:"Query" name:"InstanceIdList"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *ReleaseClusterHostGroupRequest) Invoke(client goaliyun.Client) (*ReleaseClusterHostGroupResponse, error) {
	resp := &ReleaseClusterHostGroupResponse{}
	err := client.Request("emr", "ReleaseClusterHostGroup", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ReleaseClusterHostGroupResponse struct {
	RequestId goaliyun.String
}
