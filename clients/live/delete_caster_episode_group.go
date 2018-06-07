package live

import (
	"github.com/morlay/goaliyun"
)

type DeleteCasterEpisodeGroupRequest struct {
	OwnerId   int64  `position:"Query" name:"OwnerId"`
	ProgramId string `position:"Query" name:"ProgramId"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *DeleteCasterEpisodeGroupRequest) Invoke(client goaliyun.Client) (*DeleteCasterEpisodeGroupResponse, error) {
	resp := &DeleteCasterEpisodeGroupResponse{}
	err := client.Request("live", "DeleteCasterEpisodeGroup", "2016-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteCasterEpisodeGroupResponse struct {
	RequestId goaliyun.String
}
