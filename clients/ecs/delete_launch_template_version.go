package ecs

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DeleteLaunchTemplateVersionRequest struct {
	LaunchTemplateName   string                                        `position:"Query" name:"LaunchTemplateName"`
	ResourceOwnerId      int64                                         `position:"Query" name:"ResourceOwnerId"`
	DeleteVersions       *DeleteLaunchTemplateVersionDeleteVersionList `position:"Query" type:"Repeated" name:"DeleteVersion"`
	LaunchTemplateId     string                                        `position:"Query" name:"LaunchTemplateId"`
	ResourceOwnerAccount string                                        `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string                                        `position:"Query" name:"OwnerAccount"`
	OwnerId              int64                                         `position:"Query" name:"OwnerId"`
	RegionId             string                                        `position:"Query" name:"RegionId"`
}

func (req *DeleteLaunchTemplateVersionRequest) Invoke(client goaliyun.Client) (*DeleteLaunchTemplateVersionResponse, error) {
	resp := &DeleteLaunchTemplateVersionResponse{}
	err := client.Request("ecs", "DeleteLaunchTemplateVersion", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteLaunchTemplateVersionResponse struct {
	RequestId goaliyun.String
}

type DeleteLaunchTemplateVersionDeleteVersionList []int64

func (list *DeleteLaunchTemplateVersionDeleteVersionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]int64)
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
