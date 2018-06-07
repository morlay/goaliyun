package mts

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListAllMediaBucketRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ListAllMediaBucketRequest) Invoke(client goaliyun.Client) (*ListAllMediaBucketResponse, error) {
	resp := &ListAllMediaBucketResponse{}
	err := client.Request("mts", "ListAllMediaBucket", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListAllMediaBucketResponse struct {
	RequestId       goaliyun.String
	MediaBucketList ListAllMediaBucketMediaBucketList
}

type ListAllMediaBucketMediaBucket struct {
	Bucket goaliyun.String
	Type   goaliyun.String
}

type ListAllMediaBucketMediaBucketList []ListAllMediaBucketMediaBucket

func (list *ListAllMediaBucketMediaBucketList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListAllMediaBucketMediaBucket)
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
