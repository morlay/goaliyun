package live

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ModifyCasterEpisodeRequest struct {
	ResourceId   string                              `position:"Query" name:"ResourceId"`
	ComponentIds *ModifyCasterEpisodeComponentIdList `position:"Query" type:"Repeated" name:"ComponentId"`
	SwitchType   string                              `position:"Query" name:"SwitchType"`
	CasterId     string                              `position:"Query" name:"CasterId"`
	EpisodeName  string                              `position:"Query" name:"EpisodeName"`
	EndTime      string                              `position:"Query" name:"EndTime"`
	StartTime    string                              `position:"Query" name:"StartTime"`
	OwnerId      int64                               `position:"Query" name:"OwnerId"`
	EpisodeId    string                              `position:"Query" name:"EpisodeId"`
	RegionId     string                              `position:"Query" name:"RegionId"`
}

func (req *ModifyCasterEpisodeRequest) Invoke(client goaliyun.Client) (*ModifyCasterEpisodeResponse, error) {
	resp := &ModifyCasterEpisodeResponse{}
	err := client.Request("live", "ModifyCasterEpisode", "2016-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyCasterEpisodeResponse struct {
	RequestId goaliyun.String
	CasterId  goaliyun.String
	EpisodeId goaliyun.String
}

type ModifyCasterEpisodeComponentIdList []string

func (list *ModifyCasterEpisodeComponentIdList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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
