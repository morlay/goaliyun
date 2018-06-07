package mts

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type QueryMediaListRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	IncludeSnapshotList  string `position:"Query" name:"IncludeSnapshotList"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	MediaIds             string `position:"Query" name:"MediaIds"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	IncludePlayList      string `position:"Query" name:"IncludePlayList"`
	IncludeMediaInfo     string `position:"Query" name:"IncludeMediaInfo"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *QueryMediaListRequest) Invoke(client goaliyun.Client) (*QueryMediaListResponse, error) {
	resp := &QueryMediaListResponse{}
	err := client.Request("mts", "QueryMediaList", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryMediaListResponse struct {
	RequestId        goaliyun.String
	MediaList        QueryMediaListMediaList
	NonExistMediaIds QueryMediaListNonExistMediaIdList
}

type QueryMediaListMedia struct {
	MediaId      goaliyun.String
	Title        goaliyun.String
	Description  goaliyun.String
	CoverURL     goaliyun.String
	CateId       goaliyun.Integer
	Duration     goaliyun.String
	Format       goaliyun.String
	Size         goaliyun.String
	Bitrate      goaliyun.String
	Width        goaliyun.String
	Height       goaliyun.String
	Fps          goaliyun.String
	PublishState goaliyun.String
	CreationTime goaliyun.String
	PlayList     QueryMediaListPlayList
	SnapshotList QueryMediaListSnapshotList
	Tags         QueryMediaListTagList
	RunIdList    QueryMediaListRunIdListList
	File         QueryMediaListFile
	MediaInfo    QueryMediaListMediaInfo
}

type QueryMediaListPlay struct {
	ActivityName      goaliyun.String
	MediaWorkflowId   goaliyun.String
	MediaWorkflowName goaliyun.String
	Duration          goaliyun.String
	Format            goaliyun.String
	Size              goaliyun.String
	Bitrate           goaliyun.String
	Width             goaliyun.String
	Height            goaliyun.String
	Fps               goaliyun.String
	Encryption        goaliyun.String
	File1             QueryMediaListFile1
}

type QueryMediaListFile1 struct {
	URL   goaliyun.String
	State goaliyun.String
}

type QueryMediaListSnapshot struct {
	Type              goaliyun.String
	MediaWorkflowId   goaliyun.String
	MediaWorkflowName goaliyun.String
	ActivityName      goaliyun.String
	Count             goaliyun.String
	File2             QueryMediaListFile2
}

type QueryMediaListFile2 struct {
	URL   goaliyun.String
	State goaliyun.String
}

type QueryMediaListFile struct {
	URL   goaliyun.String
	State goaliyun.String
}

type QueryMediaListMediaInfo struct {
	Streams QueryMediaListStreams
	Format  QueryMediaListFormat
}

type QueryMediaListStreams struct {
	VideoStreamList    QueryMediaListVideoStreamList
	AudioStreamList    QueryMediaListAudioStreamList
	SubtitleStreamList QueryMediaListSubtitleStreamList
}

type QueryMediaListVideoStream struct {
	Index          goaliyun.String
	CodecName      goaliyun.String
	CodecLongName  goaliyun.String
	Profile        goaliyun.String
	CodecTimeBase  goaliyun.String
	CodecTagString goaliyun.String
	CodecTag       goaliyun.String
	Width          goaliyun.String
	Height         goaliyun.String
	HasBFrames     goaliyun.String
	Sar            goaliyun.String
	Dar            goaliyun.String
	PixFmt         goaliyun.String
	Level          goaliyun.String
	Fps            goaliyun.String
	AvgFPS         goaliyun.String
	Timebase       goaliyun.String
	StartTime      goaliyun.String
	Duration       goaliyun.String
	Bitrate        goaliyun.String
	NumFrames      goaliyun.String
	Lang           goaliyun.String
	Rotate         goaliyun.String
	NetworkCost    QueryMediaListNetworkCost
}

type QueryMediaListNetworkCost struct {
	PreloadTime   goaliyun.String
	CostBandwidth goaliyun.String
	AvgBitrate    goaliyun.String
}

type QueryMediaListAudioStream struct {
	Index          goaliyun.String
	CodecName      goaliyun.String
	CodecTimeBase  goaliyun.String
	CodecLongName  goaliyun.String
	CodecTagString goaliyun.String
	CodecTag       goaliyun.String
	SampleFmt      goaliyun.String
	Samplerate     goaliyun.String
	Channels       goaliyun.String
	ChannelLayout  goaliyun.String
	Timebase       goaliyun.String
	StartTime      goaliyun.String
	Duration       goaliyun.String
	Bitrate        goaliyun.String
	NumFrames      goaliyun.String
	Lang           goaliyun.String
}

type QueryMediaListSubtitleStream struct {
	Index goaliyun.String
	Lang  goaliyun.String
}

type QueryMediaListFormat struct {
	NumStreams     goaliyun.String
	NumPrograms    goaliyun.String
	FormatName     goaliyun.String
	FormatLongName goaliyun.String
	StartTime      goaliyun.String
	Duration       goaliyun.String
	Size           goaliyun.String
	Bitrate        goaliyun.String
}

type QueryMediaListMediaList []QueryMediaListMedia

func (list *QueryMediaListMediaList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryMediaListMedia)
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

type QueryMediaListNonExistMediaIdList []goaliyun.String

func (list *QueryMediaListNonExistMediaIdList) UnmarshalJSON(data []byte) error {
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

type QueryMediaListPlayList []QueryMediaListPlay

func (list *QueryMediaListPlayList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryMediaListPlay)
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

type QueryMediaListSnapshotList []QueryMediaListSnapshot

func (list *QueryMediaListSnapshotList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryMediaListSnapshot)
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

type QueryMediaListTagList []goaliyun.String

func (list *QueryMediaListTagList) UnmarshalJSON(data []byte) error {
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

type QueryMediaListRunIdListList []goaliyun.String

func (list *QueryMediaListRunIdListList) UnmarshalJSON(data []byte) error {
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

type QueryMediaListVideoStreamList []QueryMediaListVideoStream

func (list *QueryMediaListVideoStreamList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryMediaListVideoStream)
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

type QueryMediaListAudioStreamList []QueryMediaListAudioStream

func (list *QueryMediaListAudioStreamList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryMediaListAudioStream)
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

type QueryMediaListSubtitleStreamList []QueryMediaListSubtitleStream

func (list *QueryMediaListSubtitleStreamList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryMediaListSubtitleStream)
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
