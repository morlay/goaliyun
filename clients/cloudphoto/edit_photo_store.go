package cloudphoto

import (
	"github.com/morlay/goaliyun"
)

type EditPhotoStoreRequest struct {
	AutoCleanEnabled  string `position:"Query" name:"AutoCleanEnabled"`
	DefaultTrashQuota int64  `position:"Query" name:"DefaultTrashQuota"`
	StoreName         string `position:"Query" name:"StoreName"`
	Remark            string `position:"Query" name:"Remark"`
	DefaultQuota      int64  `position:"Query" name:"DefaultQuota"`
	AutoCleanDays     int64  `position:"Query" name:"AutoCleanDays"`
	RegionId          string `position:"Query" name:"RegionId"`
}

func (req *EditPhotoStoreRequest) Invoke(client goaliyun.Client) (*EditPhotoStoreResponse, error) {
	resp := &EditPhotoStoreResponse{}
	err := client.Request("cloudphoto", "EditPhotoStore", "2017-07-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type EditPhotoStoreResponse struct {
	Code      goaliyun.String
	Message   goaliyun.String
	RequestId goaliyun.String
	Action    goaliyun.String
}
