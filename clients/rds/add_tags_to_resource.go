package rds

import (
	"github.com/morlay/goaliyun"
)

type AddTagsToResourceRequest struct {
	Tag4value            string `position:"Query" name:"Tag.4.value"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Tag2key              string `position:"Query" name:"Tag.2.key"`
	Tag5key              string `position:"Query" name:"Tag.5.key"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Tag3key              string `position:"Query" name:"Tag.3.key"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Tag5value            string `position:"Query" name:"Tag.5.value"`
	Tags                 string `position:"Query" name:"Tags"`
	Tag1key              string `position:"Query" name:"Tag.1.key"`
	Tag1value            string `position:"Query" name:"Tag.1.value"`
	Tag2value            string `position:"Query" name:"Tag.2.value"`
	Tag4key              string `position:"Query" name:"Tag.4.key"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	Tag3value            string `position:"Query" name:"Tag.3.value"`
	ProxyId              string `position:"Query" name:"ProxyId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *AddTagsToResourceRequest) Invoke(client goaliyun.Client) (*AddTagsToResourceResponse, error) {
	resp := &AddTagsToResourceResponse{}
	err := client.Request("rds", "AddTagsToResource", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AddTagsToResourceResponse struct {
	RequestId goaliyun.String
}
