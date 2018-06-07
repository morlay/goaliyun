package live

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ModifyCasterProgramRequest struct {
	CasterId string                          `position:"Query" name:"CasterId"`
	Episodes *ModifyCasterProgramEpisodeList `position:"Query" type:"Repeated" name:"Episode"`
	OwnerId  int64                           `position:"Query" name:"OwnerId"`
	RegionId string                          `position:"Query" name:"RegionId"`
}

func (req *ModifyCasterProgramRequest) Invoke(client goaliyun.Client) (*ModifyCasterProgramResponse, error) {
	resp := &ModifyCasterProgramResponse{}
	err := client.Request("live", "ModifyCasterProgram", "2016-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyCasterProgramEpisode struct {
	EpisodeId    string                              `name:"EpisodeId"`
	EpisodeType  string                              `name:"EpisodeType"`
	EpisodeName  string                              `name:"EpisodeName"`
	ResourceId   string                              `name:"ResourceId"`
	ComponentIds *ModifyCasterProgramComponentIdList `type:"Repeated" name:"ComponentId"`
	StartTime    string                              `name:"StartTime"`
	EndTime      string                              `name:"EndTime"`
	SwitchType   string                              `name:"SwitchType"`
}

type ModifyCasterProgramResponse struct {
	RequestId goaliyun.String
	CasterId  goaliyun.String
}

type ModifyCasterProgramEpisodeList []ModifyCasterProgramEpisode

func (list *ModifyCasterProgramEpisodeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ModifyCasterProgramEpisode)
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

type ModifyCasterProgramComponentIdList []string

func (list *ModifyCasterProgramComponentIdList) UnmarshalJSON(data []byte) error {
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
