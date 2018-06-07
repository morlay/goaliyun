package live

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeCasterStreamUrlRequest struct {
	CasterId string `position:"Query" name:"CasterId"`
	OwnerId  int64  `position:"Query" name:"OwnerId"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *DescribeCasterStreamUrlRequest) Invoke(client goaliyun.Client) (*DescribeCasterStreamUrlResponse, error) {
	resp := &DescribeCasterStreamUrlResponse{}
	err := client.Request("live", "DescribeCasterStreamUrl", "2016-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeCasterStreamUrlResponse struct {
	RequestId     goaliyun.String
	CasterId      goaliyun.String
	Total         goaliyun.Integer
	CasterStreams DescribeCasterStreamUrlCasterStreamList
}

type DescribeCasterStreamUrlCasterStream struct {
	SceneId     goaliyun.String
	StreamUrl   goaliyun.String
	OutputType  goaliyun.Integer
	StreamInfos DescribeCasterStreamUrlStreamInfoList
}

type DescribeCasterStreamUrlStreamInfo struct {
	TranscodeConfig goaliyun.String
	VideoFormat     goaliyun.String
	OutputStreamUrl goaliyun.String
}

type DescribeCasterStreamUrlCasterStreamList []DescribeCasterStreamUrlCasterStream

func (list *DescribeCasterStreamUrlCasterStreamList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCasterStreamUrlCasterStream)
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

type DescribeCasterStreamUrlStreamInfoList []DescribeCasterStreamUrlStreamInfo

func (list *DescribeCasterStreamUrlStreamInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCasterStreamUrlStreamInfo)
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
