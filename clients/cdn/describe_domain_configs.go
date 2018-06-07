package cdn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDomainConfigsRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	ConfigList    string `position:"Query" name:"ConfigList"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeDomainConfigsRequest) Invoke(client goaliyun.Client) (*DescribeDomainConfigsResponse, error) {
	resp := &DescribeDomainConfigsResponse{}
	err := client.Request("cdn", "DescribeDomainConfigs", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDomainConfigsResponse struct {
	RequestId     goaliyun.String
	DomainConfigs DescribeDomainConfigsDomainConfigs
}

type DescribeDomainConfigsDomainConfigs struct {
	CacheExpiredConfigs     DescribeDomainConfigsCacheExpiredConfigList
	HttpErrorPageConfigs    DescribeDomainConfigsHttpErrorPageConfigList
	HttpHeaderConfigs       DescribeDomainConfigsHttpHeaderConfigList
	DynamicConfigs          DescribeDomainConfigsDynamicConfigList
	ReqHeaderConfigs        DescribeDomainConfigsReqHeaderConfigList
	SetVarsConfigs          DescribeDomainConfigsSetVarsConfigList
	CcConfig                DescribeDomainConfigsCcConfig
	ErrorPageConfig         DescribeDomainConfigsErrorPageConfig
	OptimizeConfig          DescribeDomainConfigsOptimizeConfig
	PageCompressConfig      DescribeDomainConfigsPageCompressConfig
	IgnoreQueryStringConfig DescribeDomainConfigsIgnoreQueryStringConfig
	RangeConfig             DescribeDomainConfigsRangeConfig
	RefererConfig           DescribeDomainConfigsRefererConfig
	ReqAuthConfig           DescribeDomainConfigsReqAuthConfig
	SrcHostConfig           DescribeDomainConfigsSrcHostConfig
	VideoSeekConfig         DescribeDomainConfigsVideoSeekConfig
	WafConfig               DescribeDomainConfigsWafConfig
	NotifyUrlConfig         DescribeDomainConfigsNotifyUrlConfig
	RedirectTypeConfig      DescribeDomainConfigsRedirectTypeConfig
	ForwardSchemeConfig     DescribeDomainConfigsForwardSchemeConfig
	RemoveQueryStringConfig DescribeDomainConfigsRemoveQueryStringConfig
	L2OssKeyConfig          DescribeDomainConfigsL2OssKeyConfig
	MacServiceConfig        DescribeDomainConfigsMacServiceConfig
	GreenManagerConfig      DescribeDomainConfigsGreenManagerConfig
	HttpsOptionConfig       DescribeDomainConfigsHttpsOptionConfig
	AliBusinessConfig       DescribeDomainConfigsAliBusinessConfig
	IpAllowListConfig       DescribeDomainConfigsIpAllowListConfig
}

type DescribeDomainConfigsCacheExpiredConfig struct {
	ConfigId     goaliyun.String
	CacheType    goaliyun.String
	CacheContent goaliyun.String
	TTL          goaliyun.String
	Weight       goaliyun.String
	Status       goaliyun.String
}

type DescribeDomainConfigsHttpErrorPageConfig struct {
	ConfigId  goaliyun.String
	ErrorCode goaliyun.String
	PageUrl   goaliyun.String
	Status    goaliyun.String
}

type DescribeDomainConfigsHttpHeaderConfig struct {
	ConfigId    goaliyun.String
	HeaderKey   goaliyun.String
	HeaderValue goaliyun.String
	Status      goaliyun.String
}

type DescribeDomainConfigsDynamicConfig struct {
	ConfigId            goaliyun.String
	Enable              goaliyun.String
	DynamicOrigin       goaliyun.String
	StaticType          goaliyun.String
	StaticUri           goaliyun.String
	StaticPath          goaliyun.String
	DynamicCacheControl goaliyun.String
	Status              goaliyun.String
}

type DescribeDomainConfigsReqHeaderConfig struct {
	ConfigId goaliyun.String
	Key      goaliyun.String
	Value    goaliyun.String
	Status   goaliyun.String
}

type DescribeDomainConfigsSetVarsConfig struct {
	ConfigId goaliyun.String
	VarName  goaliyun.String
	VarValue goaliyun.String
	Status   goaliyun.String
}

type DescribeDomainConfigsCcConfig struct {
	ConfigId goaliyun.String
	Enable   goaliyun.String
	AllowIps goaliyun.String
	BlockIps goaliyun.String
	Status   goaliyun.String
}

type DescribeDomainConfigsErrorPageConfig struct {
	ConfigId      goaliyun.String
	ErrorCode     goaliyun.String
	PageType      goaliyun.String
	CustomPageUrl goaliyun.String
	Status        goaliyun.String
}

type DescribeDomainConfigsOptimizeConfig struct {
	ConfigId goaliyun.String
	Enable   goaliyun.String
	Status   goaliyun.String
}

type DescribeDomainConfigsPageCompressConfig struct {
	ConfigId goaliyun.String
	Enable   goaliyun.String
	Status   goaliyun.String
}

type DescribeDomainConfigsIgnoreQueryStringConfig struct {
	ConfigId    goaliyun.String
	HashKeyArgs goaliyun.String
	Enable      goaliyun.String
	Status      goaliyun.String
}

type DescribeDomainConfigsRangeConfig struct {
	ConfigId goaliyun.String
	Enable   goaliyun.String
	Status   goaliyun.String
}

type DescribeDomainConfigsRefererConfig struct {
	ConfigId   goaliyun.String
	ReferType  goaliyun.String
	ReferList  goaliyun.String
	AllowEmpty goaliyun.String
	DisableAst goaliyun.String
	Status     goaliyun.String
}

type DescribeDomainConfigsReqAuthConfig struct {
	ConfigId         goaliyun.String
	AuthType         goaliyun.String
	Key1             goaliyun.String
	Key2             goaliyun.String
	Status           goaliyun.String
	AliAuthWhiteList goaliyun.String
	AuthM3u8         goaliyun.String
	AuthAddr         goaliyun.String
	AuthRemoteDesc   goaliyun.String
	TimeOut          goaliyun.String
}

type DescribeDomainConfigsSrcHostConfig struct {
	ConfigId   goaliyun.String
	DomainName goaliyun.String
	Status     goaliyun.String
}

type DescribeDomainConfigsVideoSeekConfig struct {
	ConfigId goaliyun.String
	Enable   goaliyun.String
	Status   goaliyun.String
}

type DescribeDomainConfigsWafConfig struct {
	ConfigId goaliyun.String
	Enable   goaliyun.String
	Status   goaliyun.String
}

type DescribeDomainConfigsNotifyUrlConfig struct {
	Enable    goaliyun.String
	NotifyUrl goaliyun.String
}

type DescribeDomainConfigsRedirectTypeConfig struct {
	RedirectType goaliyun.String
}

type DescribeDomainConfigsForwardSchemeConfig struct {
	ConfigId         goaliyun.String
	Enable           goaliyun.String
	SchemeOrigin     goaliyun.String
	SchemeOriginPort goaliyun.String
	Status           goaliyun.String
}

type DescribeDomainConfigsRemoveQueryStringConfig struct {
	AliRemoveArgs goaliyun.String
	ConfigId      goaliyun.String
	Status        goaliyun.String
}

type DescribeDomainConfigsL2OssKeyConfig struct {
	PrivateOssAuth goaliyun.String
	ConfigId       goaliyun.String
	Status         goaliyun.String
}

type DescribeDomainConfigsMacServiceConfig struct {
	AppList       goaliyun.String
	Enabled       goaliyun.String
	ProcessResult goaliyun.String
	ConfigId      goaliyun.String
	Status        goaliyun.String
}

type DescribeDomainConfigsGreenManagerConfig struct {
	Enabled  goaliyun.String
	ConfigId goaliyun.String
	Status   goaliyun.String
}

type DescribeDomainConfigsHttpsOptionConfig struct {
	Http2    goaliyun.String
	ConfigId goaliyun.String
	Status   goaliyun.String
}

type DescribeDomainConfigsAliBusinessConfig struct {
	AliBusinessTable goaliyun.String
	AliBusinessType  goaliyun.String
	ConfigId         goaliyun.String
	Status           goaliyun.String
}

type DescribeDomainConfigsIpAllowListConfig struct {
	ConfigId  goaliyun.String
	IpList    goaliyun.String
	IpAclXfwd goaliyun.String
	Status    goaliyun.String
}

type DescribeDomainConfigsCacheExpiredConfigList []DescribeDomainConfigsCacheExpiredConfig

func (list *DescribeDomainConfigsCacheExpiredConfigList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainConfigsCacheExpiredConfig)
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

type DescribeDomainConfigsHttpErrorPageConfigList []DescribeDomainConfigsHttpErrorPageConfig

func (list *DescribeDomainConfigsHttpErrorPageConfigList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainConfigsHttpErrorPageConfig)
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

type DescribeDomainConfigsHttpHeaderConfigList []DescribeDomainConfigsHttpHeaderConfig

func (list *DescribeDomainConfigsHttpHeaderConfigList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainConfigsHttpHeaderConfig)
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

type DescribeDomainConfigsDynamicConfigList []DescribeDomainConfigsDynamicConfig

func (list *DescribeDomainConfigsDynamicConfigList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainConfigsDynamicConfig)
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

type DescribeDomainConfigsReqHeaderConfigList []DescribeDomainConfigsReqHeaderConfig

func (list *DescribeDomainConfigsReqHeaderConfigList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainConfigsReqHeaderConfig)
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

type DescribeDomainConfigsSetVarsConfigList []DescribeDomainConfigsSetVarsConfig

func (list *DescribeDomainConfigsSetVarsConfigList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainConfigsSetVarsConfig)
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
