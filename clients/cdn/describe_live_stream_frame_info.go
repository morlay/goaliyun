package cdn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeLiveStreamFrameInfoRequest struct {
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	StreamName    string `position:"Query" name:"StreamName"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeLiveStreamFrameInfoRequest) Invoke(client goaliyun.Client) (*DescribeLiveStreamFrameInfoResponse, error) {
	resp := &DescribeLiveStreamFrameInfoResponse{}
	err := client.Request("cdn", "DescribeLiveStreamFrameInfo", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeLiveStreamFrameInfoResponse struct {
	RequestId      goaliyun.String
	FrameDataInfos DescribeLiveStreamFrameInfoFrameDataModelList
}

type DescribeLiveStreamFrameInfoFrameDataModel struct {
	Time       goaliyun.String
	Stream     goaliyun.String
	ClientAddr goaliyun.String
	Server     goaliyun.String
	AudioRate  goaliyun.Float
	AudioByte  goaliyun.Float
	FrameRate  goaliyun.Float
	FrameByte  goaliyun.Float
}

type DescribeLiveStreamFrameInfoFrameDataModelList []DescribeLiveStreamFrameInfoFrameDataModel

func (list *DescribeLiveStreamFrameInfoFrameDataModelList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveStreamFrameInfoFrameDataModel)
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
