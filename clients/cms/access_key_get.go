package cms

import (
	"github.com/morlay/goaliyun"
)

type AccessKeyGetRequest struct {
	UserId   int64  `position:"Query" name:"UserId"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *AccessKeyGetRequest) Invoke(client goaliyun.Client) (*AccessKeyGetResponse, error) {
	resp := &AccessKeyGetResponse{}
	err := client.Request("cms", "AccessKeyGet", "2018-03-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AccessKeyGetResponse struct {
	ErrorCode    goaliyun.Integer
	ErrorMessage goaliyun.String
	Success      bool
	RequestId    goaliyun.String
	UserId       goaliyun.Integer
	AccessKey    goaliyun.String
	SecretKey    goaliyun.String
}
