package ram

import (
	"github.com/morlay/goaliyun"
)

type GetGroupRequest struct {
	GroupName string `position:"Query" name:"GroupName"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *GetGroupRequest) Invoke(client goaliyun.Client) (*GetGroupResponse, error) {
	resp := &GetGroupResponse{}
	err := client.Request("ram", "GetGroup", "2015-05-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetGroupResponse struct {
	RequestId goaliyun.String
	Group     GetGroupGroup
}

type GetGroupGroup struct {
	GroupName  goaliyun.String
	Comments   goaliyun.String
	CreateDate goaliyun.String
	UpdateDate goaliyun.String
}
