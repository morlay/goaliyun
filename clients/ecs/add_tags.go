package ecs

import (
	"github.com/morlay/goaliyun"
)

type AddTagsRequest struct {
	Tag4Value            string `position:"Query" name:"Tag.4.Value"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceId           string `position:"Query" name:"ResourceId"`
	Tag2Key              string `position:"Query" name:"Tag.2.Key"`
	Tag5Key              string `position:"Query" name:"Tag.5.Key"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	Tag3Key              string `position:"Query" name:"Tag.3.Key"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ResourceType         string `position:"Query" name:"ResourceType"`
	Tag5Value            string `position:"Query" name:"Tag.5.Value"`
	Tag1Key              string `position:"Query" name:"Tag.1.Key"`
	Tag1Value            string `position:"Query" name:"Tag.1.Value"`
	Tag2Value            string `position:"Query" name:"Tag.2.Value"`
	Tag4Key              string `position:"Query" name:"Tag.4.Key"`
	Tag3Value            string `position:"Query" name:"Tag.3.Value"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *AddTagsRequest) Invoke(client goaliyun.Client) (*AddTagsResponse, error) {
	resp := &AddTagsResponse{}
	err := client.Request("ecs", "AddTags", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AddTagsResponse struct {
	RequestId goaliyun.String
}
