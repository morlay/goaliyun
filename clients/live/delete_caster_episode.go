package live

import (
	"github.com/morlay/goaliyun"
)

type DeleteCasterEpisodeRequest struct {
	CasterId  string `position:"Query" name:"CasterId"`
	OwnerId   int64  `position:"Query" name:"OwnerId"`
	EpisodeId string `position:"Query" name:"EpisodeId"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *DeleteCasterEpisodeRequest) Invoke(client goaliyun.Client) (*DeleteCasterEpisodeResponse, error) {
	resp := &DeleteCasterEpisodeResponse{}
	err := client.Request("live", "DeleteCasterEpisode", "2016-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteCasterEpisodeResponse struct {
	RequestId goaliyun.String
	CasterId  goaliyun.String
	EpisodeId goaliyun.String
}
