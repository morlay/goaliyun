package qualitycheck

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type GetScoreInfoRequest struct {
	JsonStr  string `position:"Query" name:"JsonStr"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *GetScoreInfoRequest) Invoke(client goaliyun.Client) (*GetScoreInfoResponse, error) {
	resp := &GetScoreInfoResponse{}
	err := client.Request("qualitycheck", "GetScoreInfo", "2016-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetScoreInfoResponse struct {
	RequestId goaliyun.String
	Success   bool
	Code      goaliyun.String
	Message   goaliyun.String
	Data      GetScoreInfoScorePoList
}

type GetScoreInfoScorePo struct {
	ScoreId    goaliyun.Integer
	ScoreName  goaliyun.String
	ScoreInfos GetScoreInfoScoreParamList
}

type GetScoreInfoScoreParam struct {
	ScoreNum     goaliyun.Integer
	ScoreSubId   goaliyun.Integer
	ScoreSubName goaliyun.String
	ScoreType    goaliyun.Integer
}

type GetScoreInfoScorePoList []GetScoreInfoScorePo

func (list *GetScoreInfoScorePoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetScoreInfoScorePo)
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

type GetScoreInfoScoreParamList []GetScoreInfoScoreParam

func (list *GetScoreInfoScoreParamList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetScoreInfoScoreParam)
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
