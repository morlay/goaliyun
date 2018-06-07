package live

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type AddCasterEpisodeRequest struct {
	ResourceId   string                           `position:"Query" name:"ResourceId"`
	ComponentIds *AddCasterEpisodeComponentIdList `position:"Query" type:"Repeated" name:"ComponentId"`
	SwitchType   string                           `position:"Query" name:"SwitchType"`
	CasterId     string                           `position:"Query" name:"CasterId"`
	EpisodeType  string                           `position:"Query" name:"EpisodeType"`
	EpisodeName  string                           `position:"Query" name:"EpisodeName"`
	EndTime      string                           `position:"Query" name:"EndTime"`
	StartTime    string                           `position:"Query" name:"StartTime"`
	OwnerId      int64                            `position:"Query" name:"OwnerId"`
	RegionId     string                           `position:"Query" name:"RegionId"`
}

func (req *AddCasterEpisodeRequest) Invoke(client goaliyun.Client) (*AddCasterEpisodeResponse, error) {
	resp := &AddCasterEpisodeResponse{}
	err := client.Request("live", "AddCasterEpisode", "2016-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AddCasterEpisodeResponse struct {
	RequestId goaliyun.String
	EpisodeId goaliyun.String
}

type AddCasterEpisodeComponentIdList []string

func (list *AddCasterEpisodeComponentIdList) UnmarshalJSON(data []byte) error {
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
