package vod

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type GetVideoInfoRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	VideoId              string `position:"Query" name:"VideoId"`
	ResultTypes          string `position:"Query" name:"ResultTypes"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *GetVideoInfoRequest) Invoke(client goaliyun.Client) (*GetVideoInfoResponse, error) {
	resp := &GetVideoInfoResponse{}
	err := client.Request("vod", "GetVideoInfo", "2017-03-21", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetVideoInfoResponse struct {
	RequestId goaliyun.String
	AI        goaliyun.String
	Video     GetVideoInfoVideo
}

type GetVideoInfoVideo struct {
	VideoId          goaliyun.String
	Title            goaliyun.String
	Tags             goaliyun.String
	Status           goaliyun.String
	Size             goaliyun.Integer
	Duration         goaliyun.Float
	Description      goaliyun.String
	CreateTime       goaliyun.String
	CreationTime     goaliyun.String
	ModifyTime       goaliyun.String
	CoverURL         goaliyun.String
	CateId           goaliyun.Integer
	CateName         goaliyun.String
	PreprocessStatus goaliyun.String
	Snapshots        GetVideoInfoSnapshotList
}

type GetVideoInfoSnapshotList []goaliyun.String

func (list *GetVideoInfoSnapshotList) UnmarshalJSON(data []byte) error {
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
