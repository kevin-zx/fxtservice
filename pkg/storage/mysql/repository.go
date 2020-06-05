package mysql

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/kevin-zx/sqlbase"
)

type FxtStorage interface {
	GetGuanbaoProjectBySiteURLAndClientName(siteURL string, clientName string) ([]*GuanbaoProject, error)
	GetGuanwangProjectByURL(siteURL string, clientName string) ([]*GuanwangProject, error)
	Close()
}

func NewFxtStorage(engine string, port int, dbName string,
	user string, host string, passWd string) (FxtStorage, error) {
	storage, err := sqlbase.NewStorage(engine, port, dbName,
		user, host, passWd)
	if err != nil {
		return nil, err
	}
	return &fxtStorage{s: storage}, nil
}

type fxtStorage struct {
	s *sqlbase.Storage
}

func (f *fxtStorage) Close() {
	f.s.Close()
}

const userInfoBaseSql = `SELECT
	u.id id,
	u.username user_name,
	uc.username agent_business,
	ups.guanwang_or_xiala_active guanwang,
	ups.guanbao_active guanbao,
	ups.cibao_active cibao
FROM
	users u
left join users ua on
	u.parent_id = ua.id
left join users uc on
	ua.parent_id = uc.id
left join user_product_status as ups on
	ups.user_id = u.id
WHERE
	u.is_active = 1
	AND u.is_terminate = 0 
	AND u.deleted_at is null 
`

func (f *fxtStorage) GetUserByName(name string) (*User, error) {
	us, err := f.GetUser(` AND u.username LIKE ? LIMIT 1`, "%"+name+"%")
	if err != nil {
		return nil, err
	}
	if len(us) == 0 {
		return nil, fmt.Errorf("can't find specify user,name: %s", name)
	}
	return us[0], nil

}

func (f *fxtStorage) GetUserByID(id uint) (*User, error) {
	us, err := f.GetUser(` AND u.id = ? LIMIT 1`, id)
	if err != nil {
		return nil, err
	}
	if len(us) == 0 {
		return nil, fmt.Errorf("can't find specify user,id: %d", id)
	}
	return us[0], nil

}

func (f *fxtStorage) GetUser(limitPart string, values ...interface{}) ([]*User, error) {
	var users []*User

	var err error
	err = f.rawScan(userInfoBaseSql, limitPart, func(rows *sql.Rows) error {
		usr := &User{}
		err := rows.Scan(&usr.ID, &usr.UserName, &usr.AgentBusiness, &usr.Guanwang, &usr.Guanbao, &usr.Cibao)
		if err != nil {
			return err
		}
		users = append(users, usr)
		return nil
	}, values...)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (f *fxtStorage) GetSitesByUserID(userID uint) ([]*Site, error) {
	ss, err := f.GetSites(" AND user_id = ?", userID)
	if err != nil {
		return nil, err
	}
	if len(ss) == 0 {
		return nil, fmt.Errorf("can't find site by user_id:%d", userID)
	}
	return ss, nil
}

func (f *fxtStorage) GetSitesBySiteURL(siteURL string) ([]*Site, error) {
	ss, err := f.GetSites(" AND site_url like ?", "%"+siteURL+"%")
	if err != nil {
		return nil, err
	}
	if len(ss) == 0 {
		return nil, fmt.Errorf("can't find site by site_url:%s", siteURL)
	}
	return ss, nil
}

var siteBaseSql = "SELECT s.id,s.name as site_name,s.site_url,s.site_type,s.user_id,COALESCE(tsne.content,'') as special_reason FROM sites s left join tag_site_not_examine tsne on tsne.site_id = s.id WHERE deleted_at is null"
func (f *fxtStorage) GetSites(limitPart string, values ...interface{}) ([]*Site, error) {
	var sites []*Site
	var err error
	err = f.rawScan(siteBaseSql, limitPart, func(rows *sql.Rows) error {
		s := Site{}
		err := rows.Scan(&s.ID, &s.SiteName, &s.SiteURL, &s.SiteType, &s.UserID, &s.SpecialReason)
		if err != nil {
			return err
		}
		sites = append(sites, &s)
		return nil
	}, values...)
	if err != nil {
		return nil, err
	}
	return sites, nil

}

func (f *fxtStorage) GetGuanbaoProjectBySiteURLAndClientName(siteURL string, clientName string) ([]*GuanbaoProject, error) {
	u, err := f.GetUserByName(clientName)
	if err != nil {
		return nil, err
	}
	if !u.Guanbao {
		return nil, nil
	}
	sites, err := f.GetSitesBySiteURL(siteURL)
	if err != nil {
		return nil, err
	}
	var projects []*GuanbaoProject
	for _, site := range sites {
		if site.UserID == u.ID {
			gp, err := f.GetGuanbaoPackageInfo(site.UserID, site.ID)
			if err != nil {
				return nil, err
			}
			gp.User = u
			gp.Site = site
			gp.GuanbaoWords, err = f.GetGuanbaoWordByPackageID(gp.PackageID)
			if err != nil {
				return nil, err
			}
			err = f.loadGuanbaoWordRanks(gp.GuanbaoWords, 2)
			if err != nil {
				return nil, err
			}
			projects = append(projects, gp)
		}
	}
	return projects, nil
}

var guanbaoPackageInfoBaseSql = `SELECT
	gp.id as package_id,
	gp.status ,
	g.optimize_confirmed_at 
FROM
	guanbaos g
	left join guanbao_packages gp on g.remote_package_id = gp.id 
WHERE
	g.deleted_at IS NULL
	
`

func (f *fxtStorage) GetGuanbaoPackageInfo(userID uint, SiteID uint) (*GuanbaoProject, error) {
	gp := GuanbaoProject{}
	err := f.rawScan(guanbaoPackageInfoBaseSql, " AND g.user_id = ? AND g.site_id = ? ", func(rows *sql.Rows) error {
		return rows.Scan(&gp.PackageID, &gp.Status, &gp.OptimizeConfirmedAt)

	}, userID, SiteID)
	if err != nil {
		return nil, err
	}
	return &gp, nil
}

var guanbaoWordBaseSql = "SELECT id,package_id,name FROM guanbao_words gw WHERE 1 = 1 "

func (f *fxtStorage) GetGuanbaoWord(limitPart string, values ...interface{}) ([]*GuanbaoWord, error) {
	var gws []*GuanbaoWord
	err := f.rawScan(guanbaoWordBaseSql, limitPart, func(rows *sql.Rows) error {
		gw := &GuanbaoWord{}
		err := rows.Scan(&gw.ID, &gw.PackageID, &gw.Name)
		if err != nil {
			return err
		}
		gws = append(gws, gw)
		return nil
	}, values...)
	return gws, err
}

func (f *fxtStorage) GetGuanbaoWordByPackageID(packageID uint) ([]*GuanbaoWord, error) {
	return f.GetGuanbaoWord(" AND package_id = ?", packageID)
}

var guanbaoWordRankBaseSql = "SELECT id,word_id,rank,date,engine FROM guanbao_word_ranks gwr WHERE 1=1 "

func (f *fxtStorage) GetGuanbaoWordRanks(limitPart string, values ...interface{}) ([]*GuanbaoWordRank, error) {
	var grs []*GuanbaoWordRank
	err := f.rawScan(guanbaoWordRankBaseSql, limitPart, func(rows *sql.Rows) error {
		var gr = GuanbaoWordRank{}
		err := rows.Scan(&gr.ID, &gr.WordID, &gr.Rank, &gr.Date, &gr.Engine)
		if err != nil {
			return err
		}
		grs = append(grs, &gr)
		return nil
	}, values...)
	return grs, err
}

func (f *fxtStorage) GetGuanbaoWordRanksByWordIDs(wordIDs []uint, recentDays int) ([]*GuanbaoWordRank, error) {
	return f.GetGuanbaoWordRanks(" AND word_id in (?) AND TIMESTAMPDIFF(DAY,date,NOW()) <= ? ORDER BY date", wordIDs, recentDays)
}

func (f *fxtStorage) loadGuanbaoWordRanks(gws []*GuanbaoWord, recentDays int) error {
	var gwMap = make(map[uint]*GuanbaoWord)
	var gwIds []uint
	for _, gw := range gws {
		gwIds = append(gwIds, gw.ID)
		gwMap[gw.ID] = gw
	}
	if len(gwIds) == 0 {
		return nil
	}
	gwrs, err := f.GetGuanbaoWordRanksByWordIDs(gwIds, recentDays)
	if err != nil {
		return err
	}

	for _, gwr := range gwrs {
		if gw, ok := gwMap[gwr.WordID]; ok {
			gw.GuanbaoWordRanks = append(gw.GuanbaoWordRanks, gwr)
		}
	}

	return nil
}

func (f *fxtStorage) GetGuanwangProjectByURL(siteURL string, clientName string) ([]*GuanwangProject, error) {
	u, err := f.GetUserByName(clientName)
	if err != nil {
		return nil, err
	}
	if !u.Guanwang {
		return nil, nil
	}
	sites, err := f.GetSitesBySiteURL(siteURL)
	if err != nil {
		return nil, err
	}
	var projects []*GuanwangProject
	for _, site := range sites {
		if site.UserID == u.ID {
			gp := GuanwangProject{
				Project: Project{
					User: u,
					Site: site,
				},
				SiteKeywords: nil,
			}
			sks, err := f.GetGuanwangKeywordsAndRecentRank(site.ID, 2)
			if err != nil {
				return nil, err
			}
			gp.SiteKeywords = sks
			projects = append(projects, &gp)
		}
	}
	return projects, nil
}

func (f *fxtStorage) GetGuanwangKeywordsAndRecentRank(siteID uint, recentDays int) ([]*SiteKeyword, error) {
	sks, err := f.GetSiteKeywordsBySiteID(siteID)
	if err != nil {
		return nil, err
	}
	err = f.loadSiteKeywordsRank(sks, recentDays)
	if err != nil {
		return nil, err
	}
	return sks, err
}

// 这个方法有 side effect 所以作为私有方法来用
func (f *fxtStorage) loadSiteKeywordsRank(sks []*SiteKeyword, recentDays int) error {
	var sksMap = make(map[uint]*SiteKeyword)
	var skids []uint
	for _, sk := range sks {
		//if sk.SiteKeywordRanks == nil {
		//	
		//}
		skids = append(skids, sk.ID)
		sksMap[sk.ID] = sk
	}
	if len(skids) == 0 {
		return nil
	}
	skrs, err := f.GetSiteKeywordRankByKeywordIDsAndRecent(skids, recentDays)
	if err != nil {
		return err
	}
	for _, skr := range skrs {
		if sk, ok := sksMap[skr.SiteKeywordID]; ok {
			sk.SiteKeywordRanks = append(sk.SiteKeywordRanks, skr)
		}
	}
	return nil
}

func (f *fxtStorage) GetSiteKeywordsBySiteID(siteID uint) ([]*SiteKeyword, error) {
	return f.GetSiteKeyword(" AND sk.site_id = ?", siteID)
}

var siteKeywordBaseSql = `SELECT
	sk.id,
	sk.site_id,
	sk.keyword_name,
	sk.keyword_platform,
	COALESCE(xp.status,0) AS jingzhun_status,
	sk.execution_time
FROM
	site_keywords sk
left join sites s on s.id = sk.site_id 
left join jingzhun_wrapper.xiaowei_paimings xp on xp.domain = s.site_url AND xp.product = "guanwang" AND xp.word = sk.keyword_name AND xp.engine = sk.keyword_platform

where
	sk.deleted_at IS NULL 
`

func (f *fxtStorage) GetSiteKeyword(limitPart string, values ...interface{}) ([]*SiteKeyword, error) {
	var sks []*SiteKeyword
	var err error
	err = f.rawScan(siteKeywordBaseSql, limitPart, func(rows *sql.Rows) error {
		sk := SiteKeyword{}
		err := rows.Scan(&sk.ID, &sk.SiteID, &sk.KeywordName, &sk.Platform, &sk.JingzhunStatus, &sk.ExecutionTime)
		if err != nil {
			return err
		}
		sks = append(sks, &sk)
		return nil
	}, values...)
	if err != nil {
		return nil, err
	}
	return sks, nil
}

var siteKeywordRankBaseSql = "SELECT site_keyword_id,`rank`,`date` FROM site_keyword_ranks skr WHERE 1 = 1 "

func (f *fxtStorage) GetSiteKeywordRank(limitPart string, values ...interface{}) ([]*SiteKeywordRank, error) {
	var skrs []*SiteKeywordRank
	var err error
	err = f.rawScan(siteKeywordRankBaseSql, limitPart, func(rows *sql.Rows) error {
		var skr = SiteKeywordRank{}
		err := rows.Scan(&skr.SiteKeywordID, &skr.Rank, &skr.Date)
		if err != nil {
			return err
		}
		skrs = append(skrs, &skr)
		return nil
	}, values...)
	return skrs, err
}

func (f *fxtStorage) GetSiteKeywordRankByKeywordIDsAndRecent(skIDs []uint, recent int) ([]*SiteKeywordRank, error) {
	// 限制一下别崩了
	if recent <= 0 || recent > 180 {
		recent = 30
	}
	return f.GetSiteKeywordRank(" AND site_keyword_id in (?) AND TIMESTAMPDIFF(DAY,date,NOW()) <= ? ORDER BY `date`", skIDs, recent)
}

func (f *fxtStorage) rawScan(baseSql string, limitPart string, scan func(rows *sql.Rows) error, values ...interface{}) error {
	rows, err := f.s.DB.Raw(baseSql+limitPart, values...).Rows()
	if err != nil {
		return err
	}

	defer rows.Close()
	exist := false
	for rows.Next() {
		exist = true
		err = scan(rows)
		if err != nil {
			return err
		}
	}
	if !exist {
		return gorm.ErrRecordNotFound
	}
	return nil
}
