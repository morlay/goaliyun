package csb

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type GetServiceRequest struct {
	CsbId     int64  `position:"Query" name:"CsbId"`
	ServiceId int64  `position:"Query" name:"ServiceId"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *GetServiceRequest) Invoke(client goaliyun.Client) (*GetServiceResponse, error) {
	resp := &GetServiceResponse{}
	err := client.Request("csb", "GetService", "2017-11-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetServiceResponse struct {
	Code      goaliyun.Integer
	Message   goaliyun.String
	RequestId goaliyun.String
	Data      GetServiceData
}

type GetServiceData struct {
	Service GetServiceService
}

type GetServiceService struct {
	AccessParamsJSON       goaliyun.String
	Active                 bool
	Alias                  goaliyun.String
	AllVisiable            bool
	ApproveUserId          goaliyun.String
	CasTargets             goaliyun.String
	ConsumeTypesJSON       goaliyun.String
	CreateTime             goaliyun.Integer
	CsbId                  goaliyun.Integer
	ErrDefJSON             goaliyun.String
	Id                     goaliyun.Integer
	InterfaceName          goaliyun.String
	IpBlackStr             goaliyun.String
	IpWhiteStr             goaliyun.String
	ModelVersion           goaliyun.String
	ModifiedTime           goaliyun.Integer
	OldVersion             goaliyun.String
	OpenRestfulPath        goaliyun.String
	OttFlag                bool
	OwnerId                goaliyun.String
	PolicyHandler          goaliyun.String
	PrincipalName          goaliyun.String
	ProjectId              goaliyun.Integer
	ProjectName            goaliyun.String
	ProvideType            goaliyun.String
	RouteConfJson          goaliyun.String
	SSL                    bool
	Scope                  goaliyun.String
	ServiceName            goaliyun.String
	ServiceOpenRestfulPath goaliyun.String
	ServiceProviderType    goaliyun.String
	ServiceVersion         goaliyun.String
	SkipAuth               bool
	StatisticName          goaliyun.String
	Status                 goaliyun.Integer
	UserId                 goaliyun.String
	ValidConsumeTypes      bool
	ValidProvideType       bool
	ServiceVersionsList    GetServiceServiceVersionList
	VisiableGroupList      GetServiceVisiableGroupList
	CasServTargets         GetServiceCasServTargetList
	ConsumeTypes           GetServiceConsumeTypeList
	RouteConf              GetServiceRouteConf
}

type GetServiceServiceVersion struct {
	Active            bool
	AllVisiable       bool
	Id                goaliyun.Integer
	OldVersion        goaliyun.String
	OttFlag           bool
	SSL               bool
	Scope             goaliyun.String
	ServiceVersion    goaliyun.String
	SkipAuth          bool
	StatisticName     goaliyun.String
	Status            goaliyun.Integer
	ValidConsumeTypes bool
	ValidProvideType  bool
}

type GetServiceVisiableGroup struct {
	Id           goaliyun.Integer
	GroupId      goaliyun.Integer
	UserId       goaliyun.String
	ServiceId    goaliyun.Integer
	CreateTime   goaliyun.Integer
	ModifiedTime goaliyun.Integer
	Status       goaliyun.Integer
}

type GetServiceRouteConf struct {
	ServiceRouteStrategy goaliyun.String
	ImportConf           GetServiceImportConf
	ImportConfs          GetServiceImportConfs
}

type GetServiceImportConf struct {
	AccessEndpointJSON goaliyun.String
	ProvideType        goaliyun.String
	InputParameterMap  GetServiceInputParameterList
	OutputParameterMap GetServiceOutputParameterList
}

type GetServiceInputParameter struct {
	CatType      goaliyun.Integer
	Depth        goaliyun.Integer
	ExtType      goaliyun.Integer
	MapStyle     goaliyun.Integer
	Optional     bool
	OriginalName goaliyun.String
	ParamType    goaliyun.String
	PassMethod   goaliyun.String
	TargetName   goaliyun.String
}

type GetServiceOutputParameter struct {
	CatType      goaliyun.Integer
	Depth        goaliyun.Integer
	ExtType      goaliyun.Integer
	MapStyle     goaliyun.Integer
	Optional     bool
	OriginalName goaliyun.String
	ParamType    goaliyun.String
	PassMethod   goaliyun.String
	TargetName   goaliyun.String
}

type GetServiceImportConfs struct {
	AccessEndpointJSON  goaliyun.String
	ProvideType         goaliyun.String
	InputParameterMap1  GetServiceInputParameter3List
	OutputParameterMap2 GetServiceOutputParameter4List
}

type GetServiceInputParameter3 struct {
	CatType      goaliyun.Integer
	Depth        goaliyun.Integer
	ExtType      goaliyun.Integer
	MapStyle     goaliyun.Integer
	Optional     bool
	OriginalName goaliyun.String
	ParamType    goaliyun.String
	PassMethod   goaliyun.String
	TargetName   goaliyun.String
}

type GetServiceOutputParameter4 struct {
	CatType      goaliyun.Integer
	Depth        goaliyun.Integer
	ExtType      goaliyun.Integer
	MapStyle     goaliyun.Integer
	Optional     bool
	OriginalName goaliyun.String
	ParamType    goaliyun.String
	PassMethod   goaliyun.String
	TargetName   goaliyun.String
}

type GetServiceServiceVersionList []GetServiceServiceVersion

func (list *GetServiceServiceVersionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetServiceServiceVersion)
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

type GetServiceVisiableGroupList []GetServiceVisiableGroup

func (list *GetServiceVisiableGroupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetServiceVisiableGroup)
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

type GetServiceCasServTargetList []goaliyun.String

func (list *GetServiceCasServTargetList) UnmarshalJSON(data []byte) error {
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

type GetServiceConsumeTypeList []goaliyun.String

func (list *GetServiceConsumeTypeList) UnmarshalJSON(data []byte) error {
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

type GetServiceInputParameterList []GetServiceInputParameter

func (list *GetServiceInputParameterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetServiceInputParameter)
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

type GetServiceOutputParameterList []GetServiceOutputParameter

func (list *GetServiceOutputParameterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetServiceOutputParameter)
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

type GetServiceInputParameter3List []GetServiceInputParameter3

func (list *GetServiceInputParameter3List) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetServiceInputParameter3)
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

type GetServiceOutputParameter4List []GetServiceOutputParameter4

func (list *GetServiceOutputParameter4List) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetServiceOutputParameter4)
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
