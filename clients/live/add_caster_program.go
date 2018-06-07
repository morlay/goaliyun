package live

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type AddCasterProgramRequest struct {
	CasterId string                       `position:"Query" name:"CasterId"`
	Episodes *AddCasterProgramEpisodeList `position:"Query" type:"Repeated" name:"Episode"`
	OwnerId  int64                        `position:"Query" name:"OwnerId"`
	RegionId string                       `position:"Query" name:"RegionId"`
}

func (req *AddCasterProgramRequest) Invoke(client goaliyun.Client) (*AddCasterProgramResponse, error) {
	resp := &AddCasterProgramResponse{}
	err := client.Request("live", "AddCasterProgram", "2016-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AddCasterProgramEpisode struct {
	EpisodeType  string                           `name:"EpisodeType"`
	EpisodeName  string                           `name:"EpisodeName"`
	ResourceId   string                           `name:"ResourceId"`
	ComponentIds *AddCasterProgramComponentIdList `type:"Repeated" name:"ComponentId"`
	StartTime    string                           `name:"StartTime"`
	EndTime      string                           `name:"EndTime"`
	SwitchType   string                           `name:"SwitchType"`
}

type AddCasterProgramResponse struct {
	RequestId  goaliyun.String
	EpisodeIds AddCasterProgramEpisodeIdList
}

type AddCasterProgramEpisodeId struct {
	EpisodeId goaliyun.String
}

type AddCasterProgramEpisodeList []AddCasterProgramEpisode

func (list *AddCasterProgramEpisodeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]AddCasterProgramEpisode)
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

type AddCasterProgramComponentIdList []string

func (list *AddCasterProgramComponentIdList) UnmarshalJSON(data []byte) error {
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

type AddCasterProgramEpisodeIdList []AddCasterProgramEpisodeId

func (list *AddCasterProgramEpisodeIdList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]AddCasterProgramEpisodeId)
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
