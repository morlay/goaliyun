package cloudphoto

import (
	"github.com/morlay/goaliyun"
)

type GetLibraryRequest struct {
	LibraryId string `position:"Query" name:"LibraryId"`
	StoreName string `position:"Query" name:"StoreName"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *GetLibraryRequest) Invoke(client goaliyun.Client) (*GetLibraryResponse, error) {
	resp := &GetLibraryResponse{}
	err := client.Request("cloudphoto", "GetLibrary", "2017-07-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetLibraryResponse struct {
	Code      goaliyun.String
	Message   goaliyun.String
	RequestId goaliyun.String
	Action    goaliyun.String
	Library   GetLibraryLibrary
}

type GetLibraryLibrary struct {
	Ctime           goaliyun.Integer
	Quota           GetLibraryQuota
	AutoCleanConfig GetLibraryAutoCleanConfig
}

type GetLibraryQuota struct {
	TotalQuota      goaliyun.Integer
	TotalTrashQuota goaliyun.Integer
	FacesCount      goaliyun.Integer
	PhotosCount     goaliyun.Integer
	UsedQuota       goaliyun.Integer
	VideosCount     goaliyun.Integer
	ActiveSize      goaliyun.Integer
	InactiveSize    goaliyun.Integer
}

type GetLibraryAutoCleanConfig struct {
	AutoCleanEnabled bool
	AutoCleanDays    goaliyun.Integer
}
