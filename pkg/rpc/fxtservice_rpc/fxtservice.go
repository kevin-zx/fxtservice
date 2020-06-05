package fxtservice_rpc

import (
	context "context"
	"github.com/kevin-zx/fxtservice/pkg/storage/mysql"
)

func NewFxtRPCService(storage mysql.FxtStorage) FxtServiceServer {
	return &fxtService{s: storage}
}

type fxtService struct {
	s mysql.FxtStorage
}

func (f *fxtService) GetGuanbaoProject(ctx context.Context, request *ProjectRequest) (*GuanbaoProjectResponse, error) {
	gps, err := f.s.GetGuanbaoProjectBySiteURLAndClientName(request.SiteUrl, request.ClientName)
	if err != nil {
		return nil, err
	}
	gr := &GuanbaoProjectResponse{}
	for _, gp := range gps {
		gr.GuanbaoProjects = append(gr.GuanbaoProjects, convertGuanbaoProject2RPC(gp))
	}
	return gr, nil
}

func convertUser2RPC(user *mysql.User) *User {
	return &User{
		Id:            uint32(user.ID),
		UserName:      user.UserName,
		AgentBusiness: user.AgentBusiness,
		Guanwang:      user.Guanwang,
		Guanbao:       user.Guanbao,
		Cibao:         user.Cibao,
	}
}

func convertSite2RPC(site *mysql.Site) *Site {
	return &Site{
		Id:       uint32(site.ID),
		UserId:   uint32(site.UserID),
		SiteName: site.SiteName,
		SiteUrl:  site.SiteURL,
		SiteType: int32(site.SiteType),
	}
}

func convertGuanbaoProject2RPC(project *mysql.GuanbaoProject) *GuanbaoProject {
	ot := int64(0)
	if project.OptimizeConfirmedAt != nil {
		ot = project.OptimizeConfirmedAt.Unix()
	}
	gpRPC := &GuanbaoProject{
		User:                convertUser2RPC(project.User),
		Site:                convertSite2RPC(project.Site),
		PackageId:           uint32(project.PackageID),
		Status:              int32(project.Status),
		OptimizeConfirmedAt: ot,
		GuanbaoWords:        nil,
	}
	for _, word := range project.GuanbaoWords {
		gpRPC.GuanbaoWords = append(gpRPC.GuanbaoWords, convertGuanbaoWord2RPC(word))
	}
	return gpRPC

}

func convertGuanbaoWord2RPC(word *mysql.GuanbaoWord) *GuanbaoWord {
	var gbwrs []*GuanbaoWordRank
	for _, rank := range word.GuanbaoWordRanks {
		gbwrs = append(gbwrs, convertGuanbaoWordRank2RPC(rank))
	}
	gbwRPC := GuanbaoWord{
		Id:               uint32(word.ID),
		PackageId:        uint32(word.PackageID),
		Name:             word.Name,
		GuanbaoWordRanks: gbwrs,
	}
	return &gbwRPC
}

func convertGuanbaoWordRank2RPC(rank *mysql.GuanbaoWordRank) *GuanbaoWordRank {
	var d int64 = 0
	if rank.Date != nil {
		d = rank.Date.Unix()
	}
	gbRpc := GuanbaoWordRank{
		Id:     uint32(rank.ID),
		WordId: uint32(rank.WordID),
		Date:   d,
		Rank:   int32(rank.Rank),
		Engine: rank.Engine,
	}
	return &gbRpc
}

func (f *fxtService) GetGuanwangProject(ctx context.Context, request *ProjectRequest) (*GuanwangProjectResponse, error) {
	gws, err := f.s.GetGuanwangProjectByURL(request.SiteUrl, request.ClientName)
	if err != nil {
		return nil, err
	}
	gpr := GuanwangProjectResponse{}

	for _, gw := range gws {
		gpr.GuanwangProjects = append(gpr.GuanwangProjects, convertGuanwangProject2RPC(gw))
	}
	return &gpr, nil
}

func convertGuanwangProject2RPC(project *mysql.GuanwangProject) *GuanwangProject {
	var sks []*SiteKeyword
	for _, keyword := range project.SiteKeywords {
		sks = append(sks, convertGuanwangSiteKeyword2RPC(keyword))
	}
	return &GuanwangProject{
		User:         convertUser2RPC(project.User),
		Site:         convertSite2RPC(project.Site),
		SiteKeywords: sks,
	}
}

func convertGuanwangSiteKeyword2RPC(keyword *mysql.SiteKeyword) *SiteKeyword {
	et := int64(0)
	if keyword.ExecutionTime != nil {
		et = keyword.ExecutionTime.Unix()
	}
	var skrs []*SiteKeywordRank
	for _, rank := range keyword.SiteKeywordRanks {
		skrs = append(skrs, convertGuanwangSiteKeywordRank2RPC(rank))
	}
	return &SiteKeyword{
		Id:               uint32(keyword.ID),
		SiteId:           uint32(keyword.SiteID),
		KeywordName:      keyword.KeywordName,
		Platform:         keyword.Platform,
		JingzhunStatus:   int32(keyword.JingzhunStatus),
		ExecutionTime:    et,
		SpecialReason:    keyword.SpecialReason,
		SiteKeywordRanks: skrs,
	}
}

func convertGuanwangSiteKeywordRank2RPC(rank *mysql.SiteKeywordRank) *SiteKeywordRank {
	d := int64(0)
	if rank.Date != nil {
		d = rank.Date.Unix()
	}
	return &SiteKeywordRank{
		SiteKeywordId: uint32(rank.SiteKeywordID),
		Rank:          int32(rank.Rank),
		Date:          d,
	}
}
