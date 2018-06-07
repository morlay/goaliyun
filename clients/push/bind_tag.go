package push

import (
	"github.com/morlay/goaliyun"
)

type BindTagRequest struct {
	TagName   string `position:"Query" name:"TagName"`
	ClientKey string `position:"Query" name:"ClientKey"`
	AppKey    int64  `position:"Query" name:"AppKey"`
	KeyType   string `position:"Query" name:"KeyType"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *BindTagRequest) Invoke(client goaliyun.Client) (*BindTagResponse, error) {
	resp := &BindTagResponse{}
	err := client.Request("push", "BindTag", "2016-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type BindTagResponse struct {
	RequestId goaliyun.String
}
