package mysql

import "time"

type User struct {
	ID            uint   `json:"id"`
	UserName      string `json:"user_name"`
	AgentBusiness string `json:"agent_business"`
	Guanwang      bool   `json:"guanwang"`
	Guanbao       bool   `json:"guanbao"`
	Cibao         bool   `json:"cibao"`
}

type Project struct {
	User *User `json:"user"`
	Site *Site `json:"site"`
}

type GuanwangProject struct {
	Project
	SiteKeywords []*SiteKeyword `json:"site_keywords"`
}

type Site struct {
	ID            uint   `json:"id"`
	UserID        uint   `json:"user_id"`
	SiteName      string `json:"site_name"`
	SiteURL       string `json:"site_url"`
	SiteType      int    `json:"site_type"`
	SpecialReason string `json:"special_reason"`
}

type SiteKeyword struct {
	ID               uint               `json:"id"`
	SiteID           uint               `json:"site_id"`
	KeywordName      string             `json:"keyword_name"`
	Platform         string             `json:"platform"`
	JingzhunStatus   int                `json:"jingzhun_status"` // 1 开启精准  2 删除精准
	ExecutionTime    *time.Time         `json:"execution_time"`
	SiteKeywordRanks []*SiteKeywordRank `json:"site_keyword_ranks"`
}

type SiteKeywordRank struct {
	SiteKeywordID uint       `json:"site_keyword_id"`
	Rank          int        `json:"rank"`
	Date          *time.Time `json:"date"`
}

type GuanbaoProject struct {
	Project
	PackageID           uint           `json:"package_id"`
	Status              int            `json:"status"`
	OptimizeConfirmedAt *time.Time     `json:"optimize_confirmed_at"`
	GuanbaoWords        []*GuanbaoWord `json:"guanbao_words"`
}

type GuanbaoWord struct {
	ID               uint               `json:"id"`
	PackageID        uint               `json:"package_id"`
	Name             string             `json:"name"`
	GuanbaoWordRanks []*GuanbaoWordRank `json:"guanbao_word_ranks"`
}

type GuanbaoWordRank struct {
	ID     uint       `json:"id"`
	WordID uint       `json:"word_id"`
	Date   *time.Time `json:"date"`
	Rank   int        `json:"rank"`
	Engine string     `json:"engine"`
}
