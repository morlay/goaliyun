package ehpc

import (
	"github.com/morlay/goaliyun"
)

type ListCurrentClientVersionRequest struct {
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *ListCurrentClientVersionRequest) Invoke(client goaliyun.Client) (*ListCurrentClientVersionResponse, error) {
	resp := &ListCurrentClientVersionResponse{}
	err := client.Request("ehpc", "ListCurrentClientVersion", "2018-04-12", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListCurrentClientVersionResponse struct {
	RequestId     goaliyun.String
	ClientVersion goaliyun.String
}
