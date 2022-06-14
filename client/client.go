// This file is auto-generated, don't edit it. Thanks.
package client

import (
	rpcutil "github.com/alibabacloud-go/tea-rpc-utils/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
	antchainutil "github.com/antchain-openapi-sdk-go/antchain-util/service"
)

/**
 * Model for initing client
 */
type Config struct {
	// accesskey id
	AccessKeyId *string `json:"accessKeyId,omitempty" xml:"accessKeyId,omitempty"`
	// accesskey secret
	AccessKeySecret *string `json:"accessKeySecret,omitempty" xml:"accessKeySecret,omitempty"`
	// security token
	SecurityToken *string `json:"securityToken,omitempty" xml:"securityToken,omitempty"`
	// http protocol
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// read timeout
	ReadTimeout *int `json:"readTimeout,omitempty" xml:"readTimeout,omitempty"`
	// connect timeout
	ConnectTimeout *int `json:"connectTimeout,omitempty" xml:"connectTimeout,omitempty"`
	// http proxy
	HttpProxy *string `json:"httpProxy,omitempty" xml:"httpProxy,omitempty"`
	// https proxy
	HttpsProxy *string `json:"httpsProxy,omitempty" xml:"httpsProxy,omitempty"`
	// endpoint
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// proxy white list
	NoProxy *string `json:"noProxy,omitempty" xml:"noProxy,omitempty"`
	// max idle conns
	MaxIdleConns *int `json:"maxIdleConns,omitempty" xml:"maxIdleConns,omitempty"`
	// user agent
	UserAgent *string `json:"userAgent,omitempty" xml:"userAgent,omitempty"`
	// socks5 proxy
	Socks5Proxy *string `json:"socks5Proxy,omitempty" xml:"socks5Proxy,omitempty"`
	// socks5 network
	Socks5NetWork *string `json:"socks5NetWork,omitempty" xml:"socks5NetWork,omitempty"`
	// 长链接最大空闲时长
	MaxIdleTimeMillis *int `json:"maxIdleTimeMillis,omitempty" xml:"maxIdleTimeMillis,omitempty"`
	// 长链接最大连接时长
	KeepAliveDurationMillis *int `json:"keepAliveDurationMillis,omitempty" xml:"keepAliveDurationMillis,omitempty"`
	// 最大连接数（长链接最大总数）
	MaxRequests *int `json:"maxRequests,omitempty" xml:"maxRequests,omitempty"`
	// 每个目标主机的最大连接数（分主机域名的长链接最大总数
	MaxRequestsPerHost *int `json:"maxRequestsPerHost,omitempty" xml:"maxRequestsPerHost,omitempty"`
}

func (s Config) String() string {
	return tea.Prettify(s)
}

func (s Config) GoString() string {
	return s.String()
}

func (s *Config) SetAccessKeyId(v string) *Config {
	s.AccessKeyId = &v
	return s
}

func (s *Config) SetAccessKeySecret(v string) *Config {
	s.AccessKeySecret = &v
	return s
}

func (s *Config) SetSecurityToken(v string) *Config {
	s.SecurityToken = &v
	return s
}

func (s *Config) SetProtocol(v string) *Config {
	s.Protocol = &v
	return s
}

func (s *Config) SetReadTimeout(v int) *Config {
	s.ReadTimeout = &v
	return s
}

func (s *Config) SetConnectTimeout(v int) *Config {
	s.ConnectTimeout = &v
	return s
}

func (s *Config) SetHttpProxy(v string) *Config {
	s.HttpProxy = &v
	return s
}

func (s *Config) SetHttpsProxy(v string) *Config {
	s.HttpsProxy = &v
	return s
}

func (s *Config) SetEndpoint(v string) *Config {
	s.Endpoint = &v
	return s
}

func (s *Config) SetNoProxy(v string) *Config {
	s.NoProxy = &v
	return s
}

func (s *Config) SetMaxIdleConns(v int) *Config {
	s.MaxIdleConns = &v
	return s
}

func (s *Config) SetUserAgent(v string) *Config {
	s.UserAgent = &v
	return s
}

func (s *Config) SetSocks5Proxy(v string) *Config {
	s.Socks5Proxy = &v
	return s
}

func (s *Config) SetSocks5NetWork(v string) *Config {
	s.Socks5NetWork = &v
	return s
}

func (s *Config) SetMaxIdleTimeMillis(v int) *Config {
	s.MaxIdleTimeMillis = &v
	return s
}

func (s *Config) SetKeepAliveDurationMillis(v int) *Config {
	s.KeepAliveDurationMillis = &v
	return s
}

func (s *Config) SetMaxRequests(v int) *Config {
	s.MaxRequests = &v
	return s
}

func (s *Config) SetMaxRequestsPerHost(v int) *Config {
	s.MaxRequestsPerHost = &v
	return s
}

// 场所信息
type SceneInfo struct {
	// 场所名称
	SceneName *string `json:"scene_name,omitempty" xml:"scene_name,omitempty"`
	// 场所码
	SceneCode *string `json:"scene_code,omitempty" xml:"scene_code,omitempty"`
}

func (s SceneInfo) String() string {
	return tea.Prettify(s)
}

func (s SceneInfo) GoString() string {
	return s.String()
}

func (s *SceneInfo) SetSceneName(v string) *SceneInfo {
	s.SceneName = &v
	return s
}

func (s *SceneInfo) SetSceneCode(v string) *SceneInfo {
	s.SceneCode = &v
	return s
}

// 疫苗接种信息
type VaccinationInfo struct {
	// 疫苗接种信息：
	// 0:查询失败
	// 1:未接种
	// 2:已接种一针
	// 3:完成接种
	VaccinationCode *string `json:"vaccination_code,omitempty" xml:"vaccination_code,omitempty"`
	// 疫苗接种信息
	VaccinationDesc *string `json:"vaccination_desc,omitempty" xml:"vaccination_desc,omitempty"`
	// 疫苗接种时间戳（单位：ms）
	VaccinationTimestamp *int64 `json:"vaccination_timestamp,omitempty" xml:"vaccination_timestamp,omitempty"`
}

func (s VaccinationInfo) String() string {
	return tea.Prettify(s)
}

func (s VaccinationInfo) GoString() string {
	return s.String()
}

func (s *VaccinationInfo) SetVaccinationCode(v string) *VaccinationInfo {
	s.VaccinationCode = &v
	return s
}

func (s *VaccinationInfo) SetVaccinationDesc(v string) *VaccinationInfo {
	s.VaccinationDesc = &v
	return s
}

func (s *VaccinationInfo) SetVaccinationTimestamp(v int64) *VaccinationInfo {
	s.VaccinationTimestamp = &v
	return s
}

// 行程信息
type TravelInfo struct {
	// 行程码信息
	// 1:没去过疫情区，绿码;
	// 2:去过疫情区，红码;
	// 3:其他，黄码;
	// 0:行程信息不全;
	// -1:查询失败
	TravelCode *string `json:"travel_code,omitempty" xml:"travel_code,omitempty"`
	// 行程码异常原因
	TravelFactor *string `json:"travel_factor,omitempty" xml:"travel_factor,omitempty"`
}

func (s TravelInfo) String() string {
	return tea.Prettify(s)
}

func (s TravelInfo) GoString() string {
	return s.String()
}

func (s *TravelInfo) SetTravelCode(v string) *TravelInfo {
	s.TravelCode = &v
	return s
}

func (s *TravelInfo) SetTravelFactor(v string) *TravelInfo {
	s.TravelFactor = &v
	return s
}

// 抗原信息
type AntigenInfo struct {
	// 抗原检测状态
	AntigenResult *string `json:"antigen_result,omitempty" xml:"antigen_result,omitempty"`
	// 抗原检测时间戳
	AntigenTimestamp *int64 `json:"antigen_timestamp,omitempty" xml:"antigen_timestamp,omitempty"`
}

func (s AntigenInfo) String() string {
	return tea.Prettify(s)
}

func (s AntigenInfo) GoString() string {
	return s.String()
}

func (s *AntigenInfo) SetAntigenResult(v string) *AntigenInfo {
	s.AntigenResult = &v
	return s
}

func (s *AntigenInfo) SetAntigenTimestamp(v int64) *AntigenInfo {
	s.AntigenTimestamp = &v
	return s
}

// 健康码信息
type HealthInfo struct {
	// 健康码编码：
	// 1:绿色 ，
	// 2:黄色 ，
	// 3:红色 ，
	// 4.灰码 。
	HealthCode *string `json:"health_code,omitempty" xml:"health_code,omitempty"`
	// 健康码红色原因
	HealthFactor *string `json:"health_factor,omitempty" xml:"health_factor,omitempty"`
}

func (s HealthInfo) String() string {
	return tea.Prettify(s)
}

func (s HealthInfo) GoString() string {
	return s.String()
}

func (s *HealthInfo) SetHealthCode(v string) *HealthInfo {
	s.HealthCode = &v
	return s
}

func (s *HealthInfo) SetHealthFactor(v string) *HealthInfo {
	s.HealthFactor = &v
	return s
}

// 核酸信息
type NucleicAcidInfo struct {
	// 检测类型
	ReportType *string `json:"report_type,omitempty" xml:"report_type,omitempty"`
	// 检测结果
	ReportResult *string `json:"report_result,omitempty" xml:"report_result,omitempty"`
	// 检测机构
	ReportOrganization *string `json:"report_organization,omitempty" xml:"report_organization,omitempty"`
	// 检测时间戳(单位: ms)
	ReportTimestamp *int64 `json:"report_timestamp,omitempty" xml:"report_timestamp,omitempty"`
}

func (s NucleicAcidInfo) String() string {
	return tea.Prettify(s)
}

func (s NucleicAcidInfo) GoString() string {
	return s.String()
}

func (s *NucleicAcidInfo) SetReportType(v string) *NucleicAcidInfo {
	s.ReportType = &v
	return s
}

func (s *NucleicAcidInfo) SetReportResult(v string) *NucleicAcidInfo {
	s.ReportResult = &v
	return s
}

func (s *NucleicAcidInfo) SetReportOrganization(v string) *NucleicAcidInfo {
	s.ReportOrganization = &v
	return s
}

func (s *NucleicAcidInfo) SetReportTimestamp(v int64) *NucleicAcidInfo {
	s.ReportTimestamp = &v
	return s
}

type QueryHealthinfoRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 设备SN
	SerialNo *string `json:"serial_no,omitempty" xml:"serial_no,omitempty" require:"true"`
	// 设备厂商
	CorpName *string `json:"corp_name,omitempty" xml:"corp_name,omitempty" require:"true"`
	// 被查询人姓名
	Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
	// 被查询人身份证号
	CertNo *string `json:"cert_no,omitempty" xml:"cert_no,omitempty" require:"true"`
	// 防疫信息类型,枚举值如下：
	// 健康信息 health ,
	// 疫苗信息 vaccination ,
	// 核酸信息 nucleicAcid ,
	// 行程信息 travel ,
	// 抗原信息 antigen ,
	// 场景信息  scene .
	// 需要有多项信息类型时，用"|"分隔开。eg: "health|nucleicAcid"
	HealthTypes *string `json:"health_types,omitempty" xml:"health_types,omitempty" require:"true"`
	// 通行记录ID
	PassId *string `json:"pass_id,omitempty" xml:"pass_id,omitempty"`
}

func (s QueryHealthinfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryHealthinfoRequest) GoString() string {
	return s.String()
}

func (s *QueryHealthinfoRequest) SetAuthToken(v string) *QueryHealthinfoRequest {
	s.AuthToken = &v
	return s
}

func (s *QueryHealthinfoRequest) SetProductInstanceId(v string) *QueryHealthinfoRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *QueryHealthinfoRequest) SetSerialNo(v string) *QueryHealthinfoRequest {
	s.SerialNo = &v
	return s
}

func (s *QueryHealthinfoRequest) SetCorpName(v string) *QueryHealthinfoRequest {
	s.CorpName = &v
	return s
}

func (s *QueryHealthinfoRequest) SetName(v string) *QueryHealthinfoRequest {
	s.Name = &v
	return s
}

func (s *QueryHealthinfoRequest) SetCertNo(v string) *QueryHealthinfoRequest {
	s.CertNo = &v
	return s
}

func (s *QueryHealthinfoRequest) SetHealthTypes(v string) *QueryHealthinfoRequest {
	s.HealthTypes = &v
	return s
}

func (s *QueryHealthinfoRequest) SetPassId(v string) *QueryHealthinfoRequest {
	s.PassId = &v
	return s
}

type QueryHealthinfoResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 健康信息
	HealthInfo *HealthInfo `json:"health_info,omitempty" xml:"health_info,omitempty"`
	// 核酸信息
	NucleicAcidInfo *NucleicAcidInfo `json:"nucleic_acid_info,omitempty" xml:"nucleic_acid_info,omitempty"`
	// 行程信息
	//
	TravelInfo *TravelInfo `json:"travel_info,omitempty" xml:"travel_info,omitempty"`
	// 疫苗接种信息
	VaccinationInfo *VaccinationInfo `json:"vaccination_info,omitempty" xml:"vaccination_info,omitempty"`
	// 抗原信息
	AntigenInfo *AntigenInfo `json:"antigen_info,omitempty" xml:"antigen_info,omitempty"`
	// 场所信息
	SceneInfo *SceneInfo `json:"scene_info,omitempty" xml:"scene_info,omitempty"`
	// 通行记录ID
	PassId *string `json:"pass_id,omitempty" xml:"pass_id,omitempty"`
}

func (s QueryHealthinfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryHealthinfoResponse) GoString() string {
	return s.String()
}

func (s *QueryHealthinfoResponse) SetReqMsgId(v string) *QueryHealthinfoResponse {
	s.ReqMsgId = &v
	return s
}

func (s *QueryHealthinfoResponse) SetResultCode(v string) *QueryHealthinfoResponse {
	s.ResultCode = &v
	return s
}

func (s *QueryHealthinfoResponse) SetResultMsg(v string) *QueryHealthinfoResponse {
	s.ResultMsg = &v
	return s
}

func (s *QueryHealthinfoResponse) SetHealthInfo(v *HealthInfo) *QueryHealthinfoResponse {
	s.HealthInfo = v
	return s
}

func (s *QueryHealthinfoResponse) SetNucleicAcidInfo(v *NucleicAcidInfo) *QueryHealthinfoResponse {
	s.NucleicAcidInfo = v
	return s
}

func (s *QueryHealthinfoResponse) SetTravelInfo(v *TravelInfo) *QueryHealthinfoResponse {
	s.TravelInfo = v
	return s
}

func (s *QueryHealthinfoResponse) SetVaccinationInfo(v *VaccinationInfo) *QueryHealthinfoResponse {
	s.VaccinationInfo = v
	return s
}

func (s *QueryHealthinfoResponse) SetAntigenInfo(v *AntigenInfo) *QueryHealthinfoResponse {
	s.AntigenInfo = v
	return s
}

func (s *QueryHealthinfoResponse) SetSceneInfo(v *SceneInfo) *QueryHealthinfoResponse {
	s.SceneInfo = v
	return s
}

func (s *QueryHealthinfoResponse) SetPassId(v string) *QueryHealthinfoResponse {
	s.PassId = &v
	return s
}

type PushHealthinfologRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 通行记录ID
	PassId *string `json:"pass_id,omitempty" xml:"pass_id,omitempty" require:"true"`
	// 设备SN
	SerialNo *string `json:"serial_no,omitempty" xml:"serial_no,omitempty" require:"true"`
	// 设备厂商
	//
	CorpName *string `json:"corp_name,omitempty" xml:"corp_name,omitempty" require:"true"`
	// 通行时长(ms)
	CostTime *int64 `json:"cost_time,omitempty" xml:"cost_time,omitempty" require:"true"`
	// 通行结果
	PassResult *string `json:"pass_result,omitempty" xml:"pass_result,omitempty"`
	// 被查询人姓名
	//
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 被查询人身份证号
	CertNo *string `json:"cert_no,omitempty" xml:"cert_no,omitempty"`
	// 异常类型
	ErrorType *string `json:"error_type,omitempty" xml:"error_type,omitempty"`
	// 错误码
	ErrorCode *string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg *string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 体温
	Temperature *string `json:"temperature,omitempty" xml:"temperature,omitempty"`
	// 健康信息
	HealthInfo *HealthInfo `json:"health_info,omitempty" xml:"health_info,omitempty"`
	// 核酸信息
	NucleicAcidInfo *NucleicAcidInfo `json:"nucleic_acid_info,omitempty" xml:"nucleic_acid_info,omitempty"`
	// 行程信息
	//
	TravelInfo *TravelInfo `json:"travel_info,omitempty" xml:"travel_info,omitempty"`
	// 疫苗接种信息
	VaccinationInfo *VaccinationInfo `json:"vaccination_info,omitempty" xml:"vaccination_info,omitempty"`
	// 抗原信息
	AntigenInfo *AntigenInfo `json:"antigen_info,omitempty" xml:"antigen_info,omitempty"`
	// 场所信息
	//
	SceneInfo *SceneInfo `json:"scene_info,omitempty" xml:"scene_info,omitempty"`
}

func (s PushHealthinfologRequest) String() string {
	return tea.Prettify(s)
}

func (s PushHealthinfologRequest) GoString() string {
	return s.String()
}

func (s *PushHealthinfologRequest) SetAuthToken(v string) *PushHealthinfologRequest {
	s.AuthToken = &v
	return s
}

func (s *PushHealthinfologRequest) SetProductInstanceId(v string) *PushHealthinfologRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *PushHealthinfologRequest) SetPassId(v string) *PushHealthinfologRequest {
	s.PassId = &v
	return s
}

func (s *PushHealthinfologRequest) SetSerialNo(v string) *PushHealthinfologRequest {
	s.SerialNo = &v
	return s
}

func (s *PushHealthinfologRequest) SetCorpName(v string) *PushHealthinfologRequest {
	s.CorpName = &v
	return s
}

func (s *PushHealthinfologRequest) SetCostTime(v int64) *PushHealthinfologRequest {
	s.CostTime = &v
	return s
}

func (s *PushHealthinfologRequest) SetPassResult(v string) *PushHealthinfologRequest {
	s.PassResult = &v
	return s
}

func (s *PushHealthinfologRequest) SetName(v string) *PushHealthinfologRequest {
	s.Name = &v
	return s
}

func (s *PushHealthinfologRequest) SetCertNo(v string) *PushHealthinfologRequest {
	s.CertNo = &v
	return s
}

func (s *PushHealthinfologRequest) SetErrorType(v string) *PushHealthinfologRequest {
	s.ErrorType = &v
	return s
}

func (s *PushHealthinfologRequest) SetErrorCode(v string) *PushHealthinfologRequest {
	s.ErrorCode = &v
	return s
}

func (s *PushHealthinfologRequest) SetErrorMsg(v string) *PushHealthinfologRequest {
	s.ErrorMsg = &v
	return s
}

func (s *PushHealthinfologRequest) SetTemperature(v string) *PushHealthinfologRequest {
	s.Temperature = &v
	return s
}

func (s *PushHealthinfologRequest) SetHealthInfo(v *HealthInfo) *PushHealthinfologRequest {
	s.HealthInfo = v
	return s
}

func (s *PushHealthinfologRequest) SetNucleicAcidInfo(v *NucleicAcidInfo) *PushHealthinfologRequest {
	s.NucleicAcidInfo = v
	return s
}

func (s *PushHealthinfologRequest) SetTravelInfo(v *TravelInfo) *PushHealthinfologRequest {
	s.TravelInfo = v
	return s
}

func (s *PushHealthinfologRequest) SetVaccinationInfo(v *VaccinationInfo) *PushHealthinfologRequest {
	s.VaccinationInfo = v
	return s
}

func (s *PushHealthinfologRequest) SetAntigenInfo(v *AntigenInfo) *PushHealthinfologRequest {
	s.AntigenInfo = v
	return s
}

func (s *PushHealthinfologRequest) SetSceneInfo(v *SceneInfo) *PushHealthinfologRequest {
	s.SceneInfo = v
	return s
}

type PushHealthinfologResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 通行记录ID
	PassId *string `json:"pass_id,omitempty" xml:"pass_id,omitempty"`
}

func (s PushHealthinfologResponse) String() string {
	return tea.Prettify(s)
}

func (s PushHealthinfologResponse) GoString() string {
	return s.String()
}

func (s *PushHealthinfologResponse) SetReqMsgId(v string) *PushHealthinfologResponse {
	s.ReqMsgId = &v
	return s
}

func (s *PushHealthinfologResponse) SetResultCode(v string) *PushHealthinfologResponse {
	s.ResultCode = &v
	return s
}

func (s *PushHealthinfologResponse) SetResultMsg(v string) *PushHealthinfologResponse {
	s.ResultMsg = &v
	return s
}

func (s *PushHealthinfologResponse) SetPassId(v string) *PushHealthinfologResponse {
	s.PassId = &v
	return s
}

type GetHealthinfoRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 设备SN
	//
	SerialNo *string `json:"serial_no,omitempty" xml:"serial_no,omitempty" require:"true"`
	// 设备厂商
	//
	CorpName *string `json:"corp_name,omitempty" xml:"corp_name,omitempty" require:"true"`
	// 健康身份二维码
	QrCode *string `json:"qr_code,omitempty" xml:"qr_code,omitempty" require:"true"`
	// 防疫信息类型,枚举值如下： 健康信息 health , 疫苗信息 vaccination , 核酸信息 nucleicAcid , 行程信息 travel , 抗原信息 antigen , 场景信息 scene . 需要有多项信息类型时，用"|"分隔开。eg: "health|nucleicAcid"
	HealthTypes *string `json:"health_types,omitempty" xml:"health_types,omitempty" require:"true"`
	// 通行记录ID
	//
	PassId *string `json:"pass_id,omitempty" xml:"pass_id,omitempty"`
}

func (s GetHealthinfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHealthinfoRequest) GoString() string {
	return s.String()
}

func (s *GetHealthinfoRequest) SetAuthToken(v string) *GetHealthinfoRequest {
	s.AuthToken = &v
	return s
}

func (s *GetHealthinfoRequest) SetProductInstanceId(v string) *GetHealthinfoRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *GetHealthinfoRequest) SetSerialNo(v string) *GetHealthinfoRequest {
	s.SerialNo = &v
	return s
}

func (s *GetHealthinfoRequest) SetCorpName(v string) *GetHealthinfoRequest {
	s.CorpName = &v
	return s
}

func (s *GetHealthinfoRequest) SetQrCode(v string) *GetHealthinfoRequest {
	s.QrCode = &v
	return s
}

func (s *GetHealthinfoRequest) SetHealthTypes(v string) *GetHealthinfoRequest {
	s.HealthTypes = &v
	return s
}

func (s *GetHealthinfoRequest) SetPassId(v string) *GetHealthinfoRequest {
	s.PassId = &v
	return s
}

type GetHealthinfoResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 健康信息
	//
	HealthInfo *HealthInfo `json:"health_info,omitempty" xml:"health_info,omitempty"`
	// 核酸信息
	//
	NucleicAcidInfo *NucleicAcidInfo `json:"nucleic_acid_info,omitempty" xml:"nucleic_acid_info,omitempty"`
	// 行程信息
	//
	TravelInfo *TravelInfo `json:"travel_info,omitempty" xml:"travel_info,omitempty"`
	// 疫苗接种信息
	//
	VaccinationInfo *VaccinationInfo `json:"vaccination_info,omitempty" xml:"vaccination_info,omitempty"`
	// 抗原信息
	//
	AntigenInfo *AntigenInfo `json:"antigen_info,omitempty" xml:"antigen_info,omitempty"`
	// 场所信息
	//
	SceneInfo *SceneInfo `json:"scene_info,omitempty" xml:"scene_info,omitempty"`
	// 被查询人姓名
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 被查询人身份证号
	//
	CertNo *string `json:"cert_no,omitempty" xml:"cert_no,omitempty"`
	// 通行记录ID
	PassId *string `json:"pass_id,omitempty" xml:"pass_id,omitempty"`
}

func (s GetHealthinfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetHealthinfoResponse) GoString() string {
	return s.String()
}

func (s *GetHealthinfoResponse) SetReqMsgId(v string) *GetHealthinfoResponse {
	s.ReqMsgId = &v
	return s
}

func (s *GetHealthinfoResponse) SetResultCode(v string) *GetHealthinfoResponse {
	s.ResultCode = &v
	return s
}

func (s *GetHealthinfoResponse) SetResultMsg(v string) *GetHealthinfoResponse {
	s.ResultMsg = &v
	return s
}

func (s *GetHealthinfoResponse) SetHealthInfo(v *HealthInfo) *GetHealthinfoResponse {
	s.HealthInfo = v
	return s
}

func (s *GetHealthinfoResponse) SetNucleicAcidInfo(v *NucleicAcidInfo) *GetHealthinfoResponse {
	s.NucleicAcidInfo = v
	return s
}

func (s *GetHealthinfoResponse) SetTravelInfo(v *TravelInfo) *GetHealthinfoResponse {
	s.TravelInfo = v
	return s
}

func (s *GetHealthinfoResponse) SetVaccinationInfo(v *VaccinationInfo) *GetHealthinfoResponse {
	s.VaccinationInfo = v
	return s
}

func (s *GetHealthinfoResponse) SetAntigenInfo(v *AntigenInfo) *GetHealthinfoResponse {
	s.AntigenInfo = v
	return s
}

func (s *GetHealthinfoResponse) SetSceneInfo(v *SceneInfo) *GetHealthinfoResponse {
	s.SceneInfo = v
	return s
}

func (s *GetHealthinfoResponse) SetName(v string) *GetHealthinfoResponse {
	s.Name = &v
	return s
}

func (s *GetHealthinfoResponse) SetCertNo(v string) *GetHealthinfoResponse {
	s.CertNo = &v
	return s
}

func (s *GetHealthinfoResponse) SetPassId(v string) *GetHealthinfoResponse {
	s.PassId = &v
	return s
}

type Client struct {
	Endpoint                *string
	RegionId                *string
	AccessKeyId             *string
	AccessKeySecret         *string
	Protocol                *string
	UserAgent               *string
	ReadTimeout             *int
	ConnectTimeout          *int
	HttpProxy               *string
	HttpsProxy              *string
	Socks5Proxy             *string
	Socks5NetWork           *string
	NoProxy                 *string
	MaxIdleConns            *int
	SecurityToken           *string
	MaxIdleTimeMillis       *int
	KeepAliveDurationMillis *int
	MaxRequests             *int
	MaxRequestsPerHost      *int
}

/**
 * Init client with Config
 * @param config config contains the necessary information to create a client
 */
func NewClient(config *Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *Config) (_err error) {
	if tea.BoolValue(util.IsUnset(tea.ToMap(config))) {
		_err = tea.NewSDKError(map[string]interface{}{
			"code":    "ParameterMissing",
			"message": "'config' can not be unset",
		})
		return _err
	}

	client.AccessKeyId = config.AccessKeyId
	client.AccessKeySecret = config.AccessKeySecret
	client.SecurityToken = config.SecurityToken
	client.Endpoint = config.Endpoint
	client.Protocol = config.Protocol
	client.UserAgent = config.UserAgent
	client.ReadTimeout = util.DefaultNumber(config.ReadTimeout, tea.Int(20000))
	client.ConnectTimeout = util.DefaultNumber(config.ConnectTimeout, tea.Int(20000))
	client.HttpProxy = config.HttpProxy
	client.HttpsProxy = config.HttpsProxy
	client.NoProxy = config.NoProxy
	client.Socks5Proxy = config.Socks5Proxy
	client.Socks5NetWork = config.Socks5NetWork
	client.MaxIdleConns = util.DefaultNumber(config.MaxIdleConns, tea.Int(60000))
	client.MaxIdleTimeMillis = util.DefaultNumber(config.MaxIdleTimeMillis, tea.Int(5))
	client.KeepAliveDurationMillis = util.DefaultNumber(config.KeepAliveDurationMillis, tea.Int(5000))
	client.MaxRequests = util.DefaultNumber(config.MaxRequests, tea.Int(100))
	client.MaxRequestsPerHost = util.DefaultNumber(config.MaxRequestsPerHost, tea.Int(100))
	return nil
}

/**
 * Encapsulate the request and invoke the network
 * @param action api name
 * @param protocol http or https
 * @param method e.g. GET
 * @param pathname pathname of every api
 * @param request which contains request params
 * @param runtime which controls some details of call api, such as retry times
 * @return the response
 */
func (client *Client) DoRequest(version *string, action *string, protocol *string, method *string, pathname *string, request map[string]interface{}, headers map[string]*string, runtime *util.RuntimeOptions) (_result map[string]interface{}, _err error) {
	_err = tea.Validate(runtime)
	if _err != nil {
		return _result, _err
	}
	_runtime := map[string]interface{}{
		"timeouted":               "retry",
		"readTimeout":             tea.IntValue(util.DefaultNumber(runtime.ReadTimeout, client.ReadTimeout)),
		"connectTimeout":          tea.IntValue(util.DefaultNumber(runtime.ConnectTimeout, client.ConnectTimeout)),
		"httpProxy":               tea.StringValue(util.DefaultString(runtime.HttpProxy, client.HttpProxy)),
		"httpsProxy":              tea.StringValue(util.DefaultString(runtime.HttpsProxy, client.HttpsProxy)),
		"noProxy":                 tea.StringValue(util.DefaultString(runtime.NoProxy, client.NoProxy)),
		"maxIdleConns":            tea.IntValue(util.DefaultNumber(runtime.MaxIdleConns, client.MaxIdleConns)),
		"maxIdleTimeMillis":       tea.IntValue(client.MaxIdleTimeMillis),
		"keepAliveDurationMillis": tea.IntValue(client.KeepAliveDurationMillis),
		"maxRequests":             tea.IntValue(client.MaxRequests),
		"maxRequestsPerHost":      tea.IntValue(client.MaxRequestsPerHost),
		"retry": map[string]interface{}{
			"retryable":   tea.BoolValue(runtime.Autoretry),
			"maxAttempts": tea.IntValue(util.DefaultNumber(runtime.MaxAttempts, tea.Int(3))),
		},
		"backoff": map[string]interface{}{
			"policy": tea.StringValue(util.DefaultString(runtime.BackoffPolicy, tea.String("no"))),
			"period": tea.IntValue(util.DefaultNumber(runtime.BackoffPeriod, tea.Int(1))),
		},
		"ignoreSSL": tea.BoolValue(runtime.IgnoreSSL),
	}

	_resp := make(map[string]interface{})
	for _retryTimes := 0; tea.BoolValue(tea.AllowRetry(_runtime["retry"], tea.Int(_retryTimes))); _retryTimes++ {
		if _retryTimes > 0 {
			_backoffTime := tea.GetBackoffTime(_runtime["backoff"], tea.Int(_retryTimes))
			if tea.IntValue(_backoffTime) > 0 {
				tea.Sleep(_backoffTime)
			}
		}

		_resp, _err = func() (map[string]interface{}, error) {
			request_ := tea.NewRequest()
			request_.Protocol = util.DefaultString(client.Protocol, protocol)
			request_.Method = method
			request_.Pathname = pathname
			request_.Query = map[string]*string{
				"method":           action,
				"version":          version,
				"sign_type":        tea.String("HmacSHA1"),
				"req_time":         antchainutil.GetTimestamp(),
				"req_msg_id":       antchainutil.GetNonce(),
				"access_key":       client.AccessKeyId,
				"base_sdk_version": tea.String("TeaSDK-2.0"),
				"sdk_version":      tea.String("1.0.3"),
			}
			if !tea.BoolValue(util.Empty(client.SecurityToken)) {
				request_.Query["security_token"] = client.SecurityToken
			}

			request_.Headers = tea.Merge(map[string]*string{
				"host":       util.DefaultString(client.Endpoint, tea.String("openapi.antchain.antgroup.com")),
				"user-agent": util.GetUserAgent(client.UserAgent),
			}, headers)
			tmp := util.AnyifyMapValue(rpcutil.Query(request))
			request_.Body = tea.ToReader(util.ToFormString(tmp))
			request_.Headers["content-type"] = tea.String("application/x-www-form-urlencoded")
			signedParam := tea.Merge(request_.Query,
				rpcutil.Query(request))
			request_.Query["sign"] = antchainutil.GetSignature(signedParam, client.AccessKeySecret)
			response_, _err := tea.DoRequest(request_, _runtime)
			if _err != nil {
				return _result, _err
			}
			raw, _err := util.ReadAsString(response_.Body)
			if _err != nil {
				return _result, _err
			}

			obj := util.ParseJSON(raw)
			res := util.AssertAsMap(obj)
			resp := util.AssertAsMap(res["response"])
			if tea.BoolValue(antchainutil.HasError(raw, client.AccessKeySecret)) {
				_err = tea.NewSDKError(map[string]interface{}{
					"message": resp["result_msg"],
					"data":    resp,
					"code":    resp["result_code"],
				})
				return _result, _err
			}

			_result = resp
			return _result, _err
		}()
		if !tea.BoolValue(tea.Retryable(_err)) {
			break
		}
	}

	return _resp, _err
}

/**
 * Description: 查询防疫健康信息
 * Summary: 查询防疫健康信息
 */
func (client *Client) QueryHealthinfo(request *QueryHealthinfoRequest) (_result *QueryHealthinfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryHealthinfoResponse{}
	_body, _err := client.QueryHealthinfoEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 查询防疫健康信息
 * Summary: 查询防疫健康信息
 */
func (client *Client) QueryHealthinfoEx(request *QueryHealthinfoRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QueryHealthinfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &QueryHealthinfoResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("antchain.antim.healthinfo.query"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 上报通行记录
 * Summary: 上报通行记录
 */
func (client *Client) PushHealthinfolog(request *PushHealthinfologRequest) (_result *PushHealthinfologResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PushHealthinfologResponse{}
	_body, _err := client.PushHealthinfologEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 上报通行记录
 * Summary: 上报通行记录
 */
func (client *Client) PushHealthinfologEx(request *PushHealthinfologRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PushHealthinfologResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &PushHealthinfologResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("antchain.antim.healthinfolog.push"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 健康身份二维码反查防疫健康信息
 * Summary: 健康身份二维码反查防疫健康信息
 */
func (client *Client) GetHealthinfo(request *GetHealthinfoRequest) (_result *GetHealthinfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetHealthinfoResponse{}
	_body, _err := client.GetHealthinfoEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 健康身份二维码反查防疫健康信息
 * Summary: 健康身份二维码反查防疫健康信息
 */
func (client *Client) GetHealthinfoEx(request *GetHealthinfoRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetHealthinfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetHealthinfoResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("antchain.antim.healthinfo.get"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}