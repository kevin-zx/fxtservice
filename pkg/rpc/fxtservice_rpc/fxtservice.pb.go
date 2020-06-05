// Code generated by protoc-gen-go. DO NOT EDIT.
// source: fxtservice.proto

package fxtservice_rpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type User struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserName             string   `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	AgentBusiness        string   `protobuf:"bytes,3,opt,name=agent_business,json=agentBusiness,proto3" json:"agent_business,omitempty"`
	Guanwang             bool     `protobuf:"varint,4,opt,name=guanwang,proto3" json:"guanwang,omitempty"`
	Guanbao              bool     `protobuf:"varint,5,opt,name=guanbao,proto3" json:"guanbao,omitempty"`
	Cibao                bool     `protobuf:"varint,6,opt,name=cibao,proto3" json:"cibao,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_551186ad8c086063, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *User) GetAgentBusiness() string {
	if m != nil {
		return m.AgentBusiness
	}
	return ""
}

func (m *User) GetGuanwang() bool {
	if m != nil {
		return m.Guanwang
	}
	return false
}

func (m *User) GetGuanbao() bool {
	if m != nil {
		return m.Guanbao
	}
	return false
}

func (m *User) GetCibao() bool {
	if m != nil {
		return m.Cibao
	}
	return false
}

type Site struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId               uint32   `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	SiteName             string   `protobuf:"bytes,3,opt,name=site_name,json=siteName,proto3" json:"site_name,omitempty"`
	SiteUrl              string   `protobuf:"bytes,4,opt,name=site_url,json=siteUrl,proto3" json:"site_url,omitempty"`
	SiteType             int32    `protobuf:"varint,5,opt,name=site_type,json=siteType,proto3" json:"site_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Site) Reset()         { *m = Site{} }
func (m *Site) String() string { return proto.CompactTextString(m) }
func (*Site) ProtoMessage()    {}
func (*Site) Descriptor() ([]byte, []int) {
	return fileDescriptor_551186ad8c086063, []int{1}
}

func (m *Site) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Site.Unmarshal(m, b)
}
func (m *Site) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Site.Marshal(b, m, deterministic)
}
func (m *Site) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Site.Merge(m, src)
}
func (m *Site) XXX_Size() int {
	return xxx_messageInfo_Site.Size(m)
}
func (m *Site) XXX_DiscardUnknown() {
	xxx_messageInfo_Site.DiscardUnknown(m)
}

var xxx_messageInfo_Site proto.InternalMessageInfo

func (m *Site) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Site) GetUserId() uint32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *Site) GetSiteName() string {
	if m != nil {
		return m.SiteName
	}
	return ""
}

func (m *Site) GetSiteUrl() string {
	if m != nil {
		return m.SiteUrl
	}
	return ""
}

func (m *Site) GetSiteType() int32 {
	if m != nil {
		return m.SiteType
	}
	return 0
}

type SiteKeyword struct {
	Id                   uint32             `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	SiteId               uint32             `protobuf:"varint,2,opt,name=site_id,json=siteId,proto3" json:"site_id,omitempty"`
	KeywordName          string             `protobuf:"bytes,3,opt,name=keyword_name,json=keywordName,proto3" json:"keyword_name,omitempty"`
	Platform             string             `protobuf:"bytes,4,opt,name=platform,proto3" json:"platform,omitempty"`
	JingzhunStatus       int32              `protobuf:"varint,5,opt,name=jingzhun_status,json=jingzhunStatus,proto3" json:"jingzhun_status,omitempty"`
	ExecutionTime        int64              `protobuf:"varint,6,opt,name=execution_time,json=executionTime,proto3" json:"execution_time,omitempty"`
	SpecialReason        string             `protobuf:"bytes,7,opt,name=special_reason,json=specialReason,proto3" json:"special_reason,omitempty"`
	SiteKeywordRanks     []*SiteKeywordRank `protobuf:"bytes,8,rep,name=site_keyword_ranks,json=siteKeywordRanks,proto3" json:"site_keyword_ranks,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *SiteKeyword) Reset()         { *m = SiteKeyword{} }
func (m *SiteKeyword) String() string { return proto.CompactTextString(m) }
func (*SiteKeyword) ProtoMessage()    {}
func (*SiteKeyword) Descriptor() ([]byte, []int) {
	return fileDescriptor_551186ad8c086063, []int{2}
}

func (m *SiteKeyword) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SiteKeyword.Unmarshal(m, b)
}
func (m *SiteKeyword) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SiteKeyword.Marshal(b, m, deterministic)
}
func (m *SiteKeyword) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SiteKeyword.Merge(m, src)
}
func (m *SiteKeyword) XXX_Size() int {
	return xxx_messageInfo_SiteKeyword.Size(m)
}
func (m *SiteKeyword) XXX_DiscardUnknown() {
	xxx_messageInfo_SiteKeyword.DiscardUnknown(m)
}

var xxx_messageInfo_SiteKeyword proto.InternalMessageInfo

func (m *SiteKeyword) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SiteKeyword) GetSiteId() uint32 {
	if m != nil {
		return m.SiteId
	}
	return 0
}

func (m *SiteKeyword) GetKeywordName() string {
	if m != nil {
		return m.KeywordName
	}
	return ""
}

func (m *SiteKeyword) GetPlatform() string {
	if m != nil {
		return m.Platform
	}
	return ""
}

func (m *SiteKeyword) GetJingzhunStatus() int32 {
	if m != nil {
		return m.JingzhunStatus
	}
	return 0
}

func (m *SiteKeyword) GetExecutionTime() int64 {
	if m != nil {
		return m.ExecutionTime
	}
	return 0
}

func (m *SiteKeyword) GetSpecialReason() string {
	if m != nil {
		return m.SpecialReason
	}
	return ""
}

func (m *SiteKeyword) GetSiteKeywordRanks() []*SiteKeywordRank {
	if m != nil {
		return m.SiteKeywordRanks
	}
	return nil
}

type SiteKeywordRank struct {
	SiteKeywordId        uint32   `protobuf:"varint,1,opt,name=site_keyword_id,json=siteKeywordId,proto3" json:"site_keyword_id,omitempty"`
	Rank                 int32    `protobuf:"varint,2,opt,name=rank,proto3" json:"rank,omitempty"`
	Date                 int64    `protobuf:"varint,3,opt,name=date,proto3" json:"date,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SiteKeywordRank) Reset()         { *m = SiteKeywordRank{} }
func (m *SiteKeywordRank) String() string { return proto.CompactTextString(m) }
func (*SiteKeywordRank) ProtoMessage()    {}
func (*SiteKeywordRank) Descriptor() ([]byte, []int) {
	return fileDescriptor_551186ad8c086063, []int{3}
}

func (m *SiteKeywordRank) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SiteKeywordRank.Unmarshal(m, b)
}
func (m *SiteKeywordRank) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SiteKeywordRank.Marshal(b, m, deterministic)
}
func (m *SiteKeywordRank) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SiteKeywordRank.Merge(m, src)
}
func (m *SiteKeywordRank) XXX_Size() int {
	return xxx_messageInfo_SiteKeywordRank.Size(m)
}
func (m *SiteKeywordRank) XXX_DiscardUnknown() {
	xxx_messageInfo_SiteKeywordRank.DiscardUnknown(m)
}

var xxx_messageInfo_SiteKeywordRank proto.InternalMessageInfo

func (m *SiteKeywordRank) GetSiteKeywordId() uint32 {
	if m != nil {
		return m.SiteKeywordId
	}
	return 0
}

func (m *SiteKeywordRank) GetRank() int32 {
	if m != nil {
		return m.Rank
	}
	return 0
}

func (m *SiteKeywordRank) GetDate() int64 {
	if m != nil {
		return m.Date
	}
	return 0
}

type GuanwangProject struct {
	User                 *User          `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Site                 *Site          `protobuf:"bytes,2,opt,name=site,proto3" json:"site,omitempty"`
	SiteKeywords         []*SiteKeyword `protobuf:"bytes,3,rep,name=site_keywords,json=siteKeywords,proto3" json:"site_keywords,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GuanwangProject) Reset()         { *m = GuanwangProject{} }
func (m *GuanwangProject) String() string { return proto.CompactTextString(m) }
func (*GuanwangProject) ProtoMessage()    {}
func (*GuanwangProject) Descriptor() ([]byte, []int) {
	return fileDescriptor_551186ad8c086063, []int{4}
}

func (m *GuanwangProject) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GuanwangProject.Unmarshal(m, b)
}
func (m *GuanwangProject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GuanwangProject.Marshal(b, m, deterministic)
}
func (m *GuanwangProject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GuanwangProject.Merge(m, src)
}
func (m *GuanwangProject) XXX_Size() int {
	return xxx_messageInfo_GuanwangProject.Size(m)
}
func (m *GuanwangProject) XXX_DiscardUnknown() {
	xxx_messageInfo_GuanwangProject.DiscardUnknown(m)
}

var xxx_messageInfo_GuanwangProject proto.InternalMessageInfo

func (m *GuanwangProject) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *GuanwangProject) GetSite() *Site {
	if m != nil {
		return m.Site
	}
	return nil
}

func (m *GuanwangProject) GetSiteKeywords() []*SiteKeyword {
	if m != nil {
		return m.SiteKeywords
	}
	return nil
}

type GuanbaoWord struct {
	Id                   uint32             `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	PackageId            uint32             `protobuf:"varint,2,opt,name=package_id,json=packageId,proto3" json:"package_id,omitempty"`
	Name                 string             `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	GuanbaoWordRanks     []*GuanbaoWordRank `protobuf:"bytes,4,rep,name=guanbao_word_ranks,json=guanbaoWordRanks,proto3" json:"guanbao_word_ranks,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *GuanbaoWord) Reset()         { *m = GuanbaoWord{} }
func (m *GuanbaoWord) String() string { return proto.CompactTextString(m) }
func (*GuanbaoWord) ProtoMessage()    {}
func (*GuanbaoWord) Descriptor() ([]byte, []int) {
	return fileDescriptor_551186ad8c086063, []int{5}
}

func (m *GuanbaoWord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GuanbaoWord.Unmarshal(m, b)
}
func (m *GuanbaoWord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GuanbaoWord.Marshal(b, m, deterministic)
}
func (m *GuanbaoWord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GuanbaoWord.Merge(m, src)
}
func (m *GuanbaoWord) XXX_Size() int {
	return xxx_messageInfo_GuanbaoWord.Size(m)
}
func (m *GuanbaoWord) XXX_DiscardUnknown() {
	xxx_messageInfo_GuanbaoWord.DiscardUnknown(m)
}

var xxx_messageInfo_GuanbaoWord proto.InternalMessageInfo

func (m *GuanbaoWord) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *GuanbaoWord) GetPackageId() uint32 {
	if m != nil {
		return m.PackageId
	}
	return 0
}

func (m *GuanbaoWord) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GuanbaoWord) GetGuanbaoWordRanks() []*GuanbaoWordRank {
	if m != nil {
		return m.GuanbaoWordRanks
	}
	return nil
}

type GuanbaoWordRank struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	WordId               uint32   `protobuf:"varint,2,opt,name=word_id,json=wordId,proto3" json:"word_id,omitempty"`
	Date                 int64    `protobuf:"varint,3,opt,name=date,proto3" json:"date,omitempty"`
	Rank                 int32    `protobuf:"varint,4,opt,name=rank,proto3" json:"rank,omitempty"`
	Engine               string   `protobuf:"bytes,5,opt,name=engine,proto3" json:"engine,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GuanbaoWordRank) Reset()         { *m = GuanbaoWordRank{} }
func (m *GuanbaoWordRank) String() string { return proto.CompactTextString(m) }
func (*GuanbaoWordRank) ProtoMessage()    {}
func (*GuanbaoWordRank) Descriptor() ([]byte, []int) {
	return fileDescriptor_551186ad8c086063, []int{6}
}

func (m *GuanbaoWordRank) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GuanbaoWordRank.Unmarshal(m, b)
}
func (m *GuanbaoWordRank) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GuanbaoWordRank.Marshal(b, m, deterministic)
}
func (m *GuanbaoWordRank) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GuanbaoWordRank.Merge(m, src)
}
func (m *GuanbaoWordRank) XXX_Size() int {
	return xxx_messageInfo_GuanbaoWordRank.Size(m)
}
func (m *GuanbaoWordRank) XXX_DiscardUnknown() {
	xxx_messageInfo_GuanbaoWordRank.DiscardUnknown(m)
}

var xxx_messageInfo_GuanbaoWordRank proto.InternalMessageInfo

func (m *GuanbaoWordRank) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *GuanbaoWordRank) GetWordId() uint32 {
	if m != nil {
		return m.WordId
	}
	return 0
}

func (m *GuanbaoWordRank) GetDate() int64 {
	if m != nil {
		return m.Date
	}
	return 0
}

func (m *GuanbaoWordRank) GetRank() int32 {
	if m != nil {
		return m.Rank
	}
	return 0
}

func (m *GuanbaoWordRank) GetEngine() string {
	if m != nil {
		return m.Engine
	}
	return ""
}

type GuanbaoProject struct {
	User                 *User          `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Site                 *Site          `protobuf:"bytes,2,opt,name=site,proto3" json:"site,omitempty"`
	PackageId            uint32         `protobuf:"varint,3,opt,name=package_id,json=packageId,proto3" json:"package_id,omitempty"`
	Status               int32          `protobuf:"varint,4,opt,name=status,proto3" json:"status,omitempty"`
	OptimizeConfirmedAt  int64          `protobuf:"varint,5,opt,name=optimize_confirmed_at,json=optimizeConfirmedAt,proto3" json:"optimize_confirmed_at,omitempty"`
	GuanbaoWords         []*GuanbaoWord `protobuf:"bytes,6,rep,name=guanbao_words,json=guanbaoWords,proto3" json:"guanbao_words,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GuanbaoProject) Reset()         { *m = GuanbaoProject{} }
func (m *GuanbaoProject) String() string { return proto.CompactTextString(m) }
func (*GuanbaoProject) ProtoMessage()    {}
func (*GuanbaoProject) Descriptor() ([]byte, []int) {
	return fileDescriptor_551186ad8c086063, []int{7}
}

func (m *GuanbaoProject) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GuanbaoProject.Unmarshal(m, b)
}
func (m *GuanbaoProject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GuanbaoProject.Marshal(b, m, deterministic)
}
func (m *GuanbaoProject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GuanbaoProject.Merge(m, src)
}
func (m *GuanbaoProject) XXX_Size() int {
	return xxx_messageInfo_GuanbaoProject.Size(m)
}
func (m *GuanbaoProject) XXX_DiscardUnknown() {
	xxx_messageInfo_GuanbaoProject.DiscardUnknown(m)
}

var xxx_messageInfo_GuanbaoProject proto.InternalMessageInfo

func (m *GuanbaoProject) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *GuanbaoProject) GetSite() *Site {
	if m != nil {
		return m.Site
	}
	return nil
}

func (m *GuanbaoProject) GetPackageId() uint32 {
	if m != nil {
		return m.PackageId
	}
	return 0
}

func (m *GuanbaoProject) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *GuanbaoProject) GetOptimizeConfirmedAt() int64 {
	if m != nil {
		return m.OptimizeConfirmedAt
	}
	return 0
}

func (m *GuanbaoProject) GetGuanbaoWords() []*GuanbaoWord {
	if m != nil {
		return m.GuanbaoWords
	}
	return nil
}

type ProjectRequest struct {
	SiteUrl              string   `protobuf:"bytes,1,opt,name=site_url,json=siteUrl,proto3" json:"site_url,omitempty"`
	ClientName           string   `protobuf:"bytes,2,opt,name=client_name,json=clientName,proto3" json:"client_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProjectRequest) Reset()         { *m = ProjectRequest{} }
func (m *ProjectRequest) String() string { return proto.CompactTextString(m) }
func (*ProjectRequest) ProtoMessage()    {}
func (*ProjectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_551186ad8c086063, []int{8}
}

func (m *ProjectRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProjectRequest.Unmarshal(m, b)
}
func (m *ProjectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProjectRequest.Marshal(b, m, deterministic)
}
func (m *ProjectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProjectRequest.Merge(m, src)
}
func (m *ProjectRequest) XXX_Size() int {
	return xxx_messageInfo_ProjectRequest.Size(m)
}
func (m *ProjectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProjectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProjectRequest proto.InternalMessageInfo

func (m *ProjectRequest) GetSiteUrl() string {
	if m != nil {
		return m.SiteUrl
	}
	return ""
}

func (m *ProjectRequest) GetClientName() string {
	if m != nil {
		return m.ClientName
	}
	return ""
}

type GuanbaoProjectResponse struct {
	GuanbaoProjects      []*GuanbaoProject `protobuf:"bytes,1,rep,name=guanbao_projects,json=guanbaoProjects,proto3" json:"guanbao_projects,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *GuanbaoProjectResponse) Reset()         { *m = GuanbaoProjectResponse{} }
func (m *GuanbaoProjectResponse) String() string { return proto.CompactTextString(m) }
func (*GuanbaoProjectResponse) ProtoMessage()    {}
func (*GuanbaoProjectResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_551186ad8c086063, []int{9}
}

func (m *GuanbaoProjectResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GuanbaoProjectResponse.Unmarshal(m, b)
}
func (m *GuanbaoProjectResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GuanbaoProjectResponse.Marshal(b, m, deterministic)
}
func (m *GuanbaoProjectResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GuanbaoProjectResponse.Merge(m, src)
}
func (m *GuanbaoProjectResponse) XXX_Size() int {
	return xxx_messageInfo_GuanbaoProjectResponse.Size(m)
}
func (m *GuanbaoProjectResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GuanbaoProjectResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GuanbaoProjectResponse proto.InternalMessageInfo

func (m *GuanbaoProjectResponse) GetGuanbaoProjects() []*GuanbaoProject {
	if m != nil {
		return m.GuanbaoProjects
	}
	return nil
}

type GuanwangProjectResponse struct {
	GuanwangProjects     []*GuanwangProject `protobuf:"bytes,1,rep,name=guanwang_projects,json=guanwangProjects,proto3" json:"guanwang_projects,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *GuanwangProjectResponse) Reset()         { *m = GuanwangProjectResponse{} }
func (m *GuanwangProjectResponse) String() string { return proto.CompactTextString(m) }
func (*GuanwangProjectResponse) ProtoMessage()    {}
func (*GuanwangProjectResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_551186ad8c086063, []int{10}
}

func (m *GuanwangProjectResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GuanwangProjectResponse.Unmarshal(m, b)
}
func (m *GuanwangProjectResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GuanwangProjectResponse.Marshal(b, m, deterministic)
}
func (m *GuanwangProjectResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GuanwangProjectResponse.Merge(m, src)
}
func (m *GuanwangProjectResponse) XXX_Size() int {
	return xxx_messageInfo_GuanwangProjectResponse.Size(m)
}
func (m *GuanwangProjectResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GuanwangProjectResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GuanwangProjectResponse proto.InternalMessageInfo

func (m *GuanwangProjectResponse) GetGuanwangProjects() []*GuanwangProject {
	if m != nil {
		return m.GuanwangProjects
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "fxtservice_rpc.User")
	proto.RegisterType((*Site)(nil), "fxtservice_rpc.Site")
	proto.RegisterType((*SiteKeyword)(nil), "fxtservice_rpc.SiteKeyword")
	proto.RegisterType((*SiteKeywordRank)(nil), "fxtservice_rpc.SiteKeywordRank")
	proto.RegisterType((*GuanwangProject)(nil), "fxtservice_rpc.GuanwangProject")
	proto.RegisterType((*GuanbaoWord)(nil), "fxtservice_rpc.GuanbaoWord")
	proto.RegisterType((*GuanbaoWordRank)(nil), "fxtservice_rpc.GuanbaoWordRank")
	proto.RegisterType((*GuanbaoProject)(nil), "fxtservice_rpc.GuanbaoProject")
	proto.RegisterType((*ProjectRequest)(nil), "fxtservice_rpc.ProjectRequest")
	proto.RegisterType((*GuanbaoProjectResponse)(nil), "fxtservice_rpc.GuanbaoProjectResponse")
	proto.RegisterType((*GuanwangProjectResponse)(nil), "fxtservice_rpc.GuanwangProjectResponse")
}

func init() { proto.RegisterFile("fxtservice.proto", fileDescriptor_551186ad8c086063) }

var fileDescriptor_551186ad8c086063 = []byte{
	// 780 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0xcb, 0x6e, 0xdb, 0x38,
	0x14, 0x1d, 0xf9, 0xed, 0xeb, 0xd8, 0x4e, 0x38, 0x99, 0x44, 0x93, 0x60, 0x26, 0x1e, 0x01, 0x93,
	0x78, 0x95, 0x85, 0xfb, 0x03, 0x7d, 0x00, 0x0d, 0x8c, 0xa6, 0x45, 0xc1, 0x24, 0xe8, 0xaa, 0x10,
	0x18, 0x89, 0x51, 0x19, 0xdb, 0x92, 0x2a, 0x52, 0xcd, 0x63, 0xdd, 0x1f, 0xe8, 0x17, 0x74, 0xd7,
	0x7e, 0x44, 0x7f, 0xa1, 0x1f, 0x55, 0xf0, 0x21, 0x85, 0x72, 0xdc, 0x66, 0xd5, 0x1d, 0xef, 0x11,
	0xc9, 0x7b, 0x78, 0xce, 0x21, 0x05, 0xeb, 0x17, 0xd7, 0x82, 0xd3, 0xec, 0x03, 0x0b, 0xe8, 0x61,
	0x9a, 0x25, 0x22, 0x41, 0x83, 0x3b, 0xc4, 0xcf, 0xd2, 0xc0, 0xfb, 0xe2, 0x40, 0xe3, 0x8c, 0xd3,
	0x0c, 0x0d, 0xa0, 0xc6, 0x42, 0xd7, 0x19, 0x39, 0xe3, 0x3e, 0xae, 0xb1, 0x10, 0xed, 0x42, 0x37,
	0xe7, 0x34, 0xf3, 0x63, 0xb2, 0xa0, 0x6e, 0x6d, 0xe4, 0x8c, 0xbb, 0xb8, 0x23, 0x81, 0x57, 0x64,
	0x41, 0xd1, 0xff, 0x30, 0x20, 0x11, 0x8d, 0x85, 0x7f, 0x9e, 0x73, 0x16, 0x53, 0xce, 0xdd, 0xba,
	0x9a, 0xd1, 0x57, 0xe8, 0x53, 0x03, 0xa2, 0x1d, 0xe8, 0x44, 0x39, 0x89, 0xaf, 0x48, 0x1c, 0xb9,
	0x8d, 0x91, 0x33, 0xee, 0xe0, 0xb2, 0x46, 0x2e, 0xb4, 0xe5, 0xf8, 0x9c, 0x24, 0x6e, 0x53, 0x7d,
	0x2a, 0x4a, 0xb4, 0x09, 0xcd, 0x80, 0x49, 0xbc, 0xa5, 0x70, 0x5d, 0x78, 0x1f, 0x1d, 0x68, 0x9c,
	0x30, 0x41, 0xef, 0x11, 0xdd, 0x86, 0xb6, 0x22, 0xca, 0x42, 0x45, 0xb3, 0x8f, 0x5b, 0xb2, 0x9c,
	0xaa, 0x13, 0x70, 0x26, 0xa8, 0x3e, 0x81, 0xe6, 0xd7, 0x91, 0x80, 0x3a, 0xc1, 0xdf, 0xa0, 0xc6,
	0x7e, 0x9e, 0xcd, 0x15, 0xb5, 0x2e, 0x6e, 0xcb, 0xfa, 0x2c, 0x9b, 0x97, 0xeb, 0xc4, 0x4d, 0x4a,
	0x15, 0xb7, 0xa6, 0x5e, 0x77, 0x7a, 0x93, 0x52, 0xef, 0x5b, 0x0d, 0x7a, 0x92, 0xc6, 0x0b, 0x7a,
	0x73, 0x95, 0x64, 0xe1, 0x2a, 0x36, 0x6a, 0xf1, 0x1d, 0x1b, 0x59, 0x4e, 0x43, 0xf4, 0x1f, 0xac,
	0xcd, 0xf4, 0x1a, 0x9b, 0x50, 0xcf, 0x60, 0x8a, 0xd3, 0x0e, 0x74, 0xd2, 0x39, 0x11, 0x17, 0x49,
	0xb6, 0x30, 0x9c, 0xca, 0x1a, 0x1d, 0xc0, 0xf0, 0x92, 0xc5, 0xd1, 0xed, 0xbb, 0x3c, 0xf6, 0xb9,
	0x20, 0x22, 0xe7, 0x86, 0xda, 0xa0, 0x80, 0x4f, 0x14, 0x2a, 0xad, 0xa1, 0xd7, 0x34, 0xc8, 0x05,
	0x4b, 0x62, 0x5f, 0xb0, 0x05, 0x55, 0x32, 0xd6, 0x71, 0xbf, 0x44, 0x4f, 0x99, 0x76, 0x90, 0xa7,
	0x34, 0x60, 0x64, 0xee, 0x67, 0x94, 0xf0, 0x24, 0x76, 0xdb, 0xda, 0x41, 0x83, 0x62, 0x05, 0xa2,
	0x97, 0x80, 0xd4, 0x71, 0x0a, 0xea, 0x19, 0x89, 0x67, 0xdc, 0xed, 0x8c, 0xea, 0xe3, 0xde, 0x64,
	0xef, 0xb0, 0x9a, 0xa5, 0x43, 0x4b, 0x17, 0x4c, 0xe2, 0x19, 0x5e, 0xe7, 0x55, 0x80, 0x7b, 0x04,
	0x86, 0x4b, 0x93, 0xd0, 0x3e, 0x0c, 0x2b, 0x1d, 0x4a, 0x35, 0xfb, 0xd6, 0xea, 0x69, 0x88, 0x10,
	0x34, 0x64, 0x73, 0xa5, 0x6a, 0x13, 0xab, 0xb1, 0xc4, 0x42, 0x22, 0xb4, 0x96, 0x75, 0xac, 0xc6,
	0xde, 0x57, 0x07, 0x86, 0x47, 0x26, 0x64, 0xaf, 0xb3, 0xe4, 0x92, 0x06, 0x02, 0x8d, 0xa1, 0x21,
	0x33, 0xa1, 0x36, 0xee, 0x4d, 0x36, 0x97, 0x79, 0xcb, 0xfc, 0x63, 0x35, 0x43, 0xce, 0x94, 0x6d,
	0x55, 0x97, 0x15, 0x33, 0x25, 0x79, 0xac, 0x66, 0xa0, 0xc7, 0xd0, 0xb7, 0x79, 0xcb, 0x1b, 0x20,
	0x45, 0xd9, 0xfd, 0x95, 0x28, 0x6b, 0xd6, 0x91, 0xb8, 0xf7, 0xd9, 0x81, 0xde, 0x91, 0xce, 0xfc,
	0x9b, 0x55, 0x51, 0xfa, 0x07, 0x20, 0x25, 0xc1, 0x8c, 0x44, 0x56, 0x9a, 0xba, 0x06, 0xd1, 0x82,
	0x58, 0x41, 0x52, 0x63, 0x69, 0x97, 0xb9, 0x45, 0xbe, 0x65, 0x57, 0x63, 0xb5, 0x5d, 0x56, 0x6f,
	0x6d, 0x57, 0x54, 0x05, 0xb8, 0x77, 0xab, 0xa5, 0xb4, 0xb0, 0x55, 0x79, 0x2f, 0x6c, 0x33, 0x79,
	0xbf, 0xf3, 0x6b, 0xd9, 0x9b, 0xd2, 0xc3, 0x86, 0xe5, 0xe1, 0x16, 0xb4, 0x68, 0x1c, 0xb1, 0x58,
	0x5f, 0xb5, 0x2e, 0x36, 0x95, 0xf7, 0xa9, 0x06, 0x03, 0xd3, 0xfc, 0x77, 0xda, 0x58, 0x15, 0xb9,
	0xbe, 0x2c, 0xf2, 0x16, 0xb4, 0xcc, 0x6d, 0xd3, 0x9c, 0x4d, 0x85, 0x26, 0xf0, 0x57, 0x92, 0x0a,
	0xb6, 0x60, 0xb7, 0xd4, 0x0f, 0x92, 0xf8, 0x82, 0x65, 0x0b, 0x1a, 0xfa, 0x44, 0xa8, 0x43, 0xd4,
	0xf1, 0x9f, 0xc5, 0xc7, 0x67, 0xc5, 0xb7, 0x27, 0x42, 0x26, 0xc6, 0x36, 0x87, 0xbb, 0xad, 0xd5,
	0x89, 0xb1, 0x25, 0x5f, 0xb3, 0x3c, 0xe1, 0xde, 0x31, 0x0c, 0x8c, 0x16, 0x98, 0xbe, 0xcf, 0x29,
	0x17, 0x95, 0x67, 0xcc, 0xa9, 0x3e, 0x63, 0x7b, 0xd0, 0x0b, 0xe6, 0x4c, 0x3e, 0xd2, 0xd6, 0x13,
	0x0e, 0x1a, 0x92, 0xcf, 0x8d, 0x17, 0xc0, 0x56, 0x55, 0x60, 0x4c, 0x79, 0x9a, 0xc4, 0x9c, 0xa2,
	0x29, 0x14, 0x59, 0xf0, 0x53, 0xfd, 0x89, 0xbb, 0x8e, 0x22, 0xfb, 0xef, 0x4f, 0xc8, 0x16, 0x3b,
	0x0c, 0xa3, 0x4a, 0xcd, 0xbd, 0x08, 0xb6, 0x97, 0x6e, 0x63, 0xd9, 0xe5, 0x18, 0x36, 0x8a, 0xbf,
	0xc1, 0x72, 0x9b, 0x95, 0x59, 0xb5, 0xf7, 0x58, 0x8f, 0xaa, 0x00, 0x9f, 0x7c, 0x77, 0x00, 0x9e,
	0x5f, 0x8b, 0x13, 0xbd, 0x08, 0xbd, 0x85, 0x8d, 0x23, 0x2a, 0x96, 0x02, 0x74, 0x8f, 0x7d, 0x55,
	0xcd, 0x9d, 0xfd, 0x07, 0x4e, 0x67, 0x98, 0x7b, 0x7f, 0x20, 0x1f, 0x90, 0xd9, 0xde, 0x7e, 0x67,
	0x1e, 0xda, 0xff, 0xe0, 0xa1, 0x63, 0x95, 0x0d, 0xce, 0x5b, 0xea, 0x77, 0xfd, 0xe8, 0x47, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xb2, 0x52, 0xa9, 0xe1, 0xc2, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FxtServiceClient is the client API for FxtService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FxtServiceClient interface {
	GetGuanbaoProject(ctx context.Context, in *ProjectRequest, opts ...grpc.CallOption) (*GuanbaoProjectResponse, error)
	GetGuanwangProject(ctx context.Context, in *ProjectRequest, opts ...grpc.CallOption) (*GuanwangProjectResponse, error)
}

type fxtServiceClient struct {
	cc *grpc.ClientConn
}

func NewFxtServiceClient(cc *grpc.ClientConn) FxtServiceClient {
	return &fxtServiceClient{cc}
}

func (c *fxtServiceClient) GetGuanbaoProject(ctx context.Context, in *ProjectRequest, opts ...grpc.CallOption) (*GuanbaoProjectResponse, error) {
	out := new(GuanbaoProjectResponse)
	err := c.cc.Invoke(ctx, "/fxtservice_rpc.FxtService/GetGuanbaoProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fxtServiceClient) GetGuanwangProject(ctx context.Context, in *ProjectRequest, opts ...grpc.CallOption) (*GuanwangProjectResponse, error) {
	out := new(GuanwangProjectResponse)
	err := c.cc.Invoke(ctx, "/fxtservice_rpc.FxtService/GetGuanwangProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FxtServiceServer is the server API for FxtService service.
type FxtServiceServer interface {
	GetGuanbaoProject(context.Context, *ProjectRequest) (*GuanbaoProjectResponse, error)
	GetGuanwangProject(context.Context, *ProjectRequest) (*GuanwangProjectResponse, error)
}

// UnimplementedFxtServiceServer can be embedded to have forward compatible implementations.
type UnimplementedFxtServiceServer struct {
}

func (*UnimplementedFxtServiceServer) GetGuanbaoProject(ctx context.Context, req *ProjectRequest) (*GuanbaoProjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGuanbaoProject not implemented")
}
func (*UnimplementedFxtServiceServer) GetGuanwangProject(ctx context.Context, req *ProjectRequest) (*GuanwangProjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGuanwangProject not implemented")
}

func RegisterFxtServiceServer(s *grpc.Server, srv FxtServiceServer) {
	s.RegisterService(&_FxtService_serviceDesc, srv)
}

func _FxtService_GetGuanbaoProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FxtServiceServer).GetGuanbaoProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fxtservice_rpc.FxtService/GetGuanbaoProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FxtServiceServer).GetGuanbaoProject(ctx, req.(*ProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FxtService_GetGuanwangProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FxtServiceServer).GetGuanwangProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fxtservice_rpc.FxtService/GetGuanwangProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FxtServiceServer).GetGuanwangProject(ctx, req.(*ProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FxtService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fxtservice_rpc.FxtService",
	HandlerType: (*FxtServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGuanbaoProject",
			Handler:    _FxtService_GetGuanbaoProject_Handler,
		},
		{
			MethodName: "GetGuanwangProject",
			Handler:    _FxtService_GetGuanwangProject_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fxtservice.proto",
}