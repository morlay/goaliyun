package live

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeCasterProgramRequest struct {
	CasterId    string `position:"Query" name:"CasterId"`
	EpisodeType string `position:"Query" name:"EpisodeType"`
	PageSize    int64  `position:"Query" name:"PageSize"`
	EndTime     string `position:"Query" name:"EndTime"`
	StartTime   string `position:"Query" name:"StartTime"`
	OwnerId     int64  `position:"Query" name:"OwnerId"`
	EpisodeId   string `position:"Query" name:"EpisodeId"`
	PageNum     int64  `position:"Query" name:"PageNum"`
	Status      int64  `position:"Query" name:"Status"`
	RegionId    string `position:"Query" name:"RegionId"`
}

func (req *DescribeCasterProgramRequest) Invoke(client goaliyun.Client) (*DescribeCasterProgramResponse, error) {
	resp := &DescribeCasterProgramResponse{}
	err := client.Request("live", "DescribeCasterProgram", "2016-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeCasterProgramResponse struct {
	RequestId     goaliyun.String
	CasterId      goaliyun.String
	ProgramName   goaliyun.String
	ProgramEffect goaliyun.Integer
	Total         goaliyun.Integer
	Episodes      DescribeCasterProgramEpisodeList
}

type DescribeCasterProgramEpisode struct {
	EpisodeId    goaliyun.String
	EpisodeType  goaliyun.String
	EpisodeName  goaliyun.String
	ResourceId   goaliyun.String
	StartTime    goaliyun.String
	EndTime      goaliyun.String
	SwitchType   goaliyun.String
	Status       goaliyun.Integer
	ComponentIds DescribeCasterProgramComponentIdList
}

type DescribeCasterProgramEpisodeList []DescribeCasterProgramEpisode

func (list *DescribeCasterProgramEpisodeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCasterProgramEpisode)
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

type DescribeCasterProgramComponentIdList []goaliyun.String

func (list *DescribeCasterProgramComponentIdList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]goaliyun.String)
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
