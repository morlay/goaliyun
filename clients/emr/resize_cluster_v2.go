package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ResizeClusterV2Request struct {
	HostGroups *ResizeClusterV2HostGroupList `position:"Query" type:"Repeated" name:"HostGroup"`
	ClusterId  string                        `position:"Query" name:"ClusterId"`
	RegionId   string                        `position:"Query" name:"RegionId"`
}

func (req *ResizeClusterV2Request) Invoke(client goaliyun.Client) (*ResizeClusterV2Response, error) {
	resp := &ResizeClusterV2Response{}
	err := client.Request("emr", "ResizeClusterV2", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ResizeClusterV2HostGroup struct {
	ClusterId       string `name:"ClusterId"`
	HostGroupId     string `name:"HostGroupId"`
	HostGroupName   string `name:"HostGroupName"`
	HostGroupType   string `name:"HostGroupType"`
	Comment         string `name:"Comment"`
	CreateType      string `name:"CreateType"`
	ChargeType      string `name:"ChargeType"`
	Period          int64  `name:"Period"`
	NodeCount       int64  `name:"NodeCount"`
	InstanceType    string `name:"InstanceType"`
	DiskType        string `name:"DiskType"`
	DiskCapacity    int64  `name:"DiskCapacity"`
	DiskCount       int64  `name:"DiskCount"`
	SysDiskType     string `name:"SysDiskType"`
	SysDiskCapacity int64  `name:"SysDiskCapacity"`
	AutoRenew       string `name:"AutoRenew"`
	VswitchId       int64  `name:"VswitchId"`
}

type ResizeClusterV2Response struct {
	RequestId goaliyun.String
	ClusterId goaliyun.String
}

type ResizeClusterV2HostGroupList []ResizeClusterV2HostGroup

func (list *ResizeClusterV2HostGroupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ResizeClusterV2HostGroup)
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
