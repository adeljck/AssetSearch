package quake

import "time"

type QuakeUserInfo struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		ID   string `json:"id"`
		User struct {
			ID       string `json:"id"`
			Username string `json:"username"`
			Fullname string `json:"fullname"`
			Email    string `json:"email"`
		} `json:"user"`
		Baned            bool   `json:"baned"`
		BanStatus        string `json:"ban_status"`
		Credit           int    `json:"credit"`
		PersistentCredit int    `json:"persistent_credit"`
		Token            string `json:"token"`
		MobilePhone      string `json:"mobile_phone"`
		Source           string `json:"source"`
		PrivacyLog       struct {
			Status bool        `json:"status"`
			Time   interface{} `json:"time"`
		} `json:"privacy_log"`
		EnterpriseInformation struct {
			Name   interface{} `json:"name"`
			Email  interface{} `json:"email"`
			Status string      `json:"status"`
		} `json:"enterprise_information"`
		PersonalInformationStatus bool `json:"personal_information_status"`
		Role                      []struct {
			Fullname string `json:"fullname"`
			Priority int    `json:"priority"`
			Credit   int    `json:"credit"`
		} `json:"role"`
	} `json:"data"`
	Meta struct {
	} `json:"meta"`
}

type Result struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    []Data `json:"data"`
	Meta    struct {
		Pagination struct {
			Count     int `json:"count"`
			PageIndex int `json:"page_index"`
			PageSize  int `json:"page_size"`
			Total     int `json:"total"`
		} `json:"pagination"`
	} `json:"meta"`
}

type Data struct {
	Hostname string `json:"hostname"`
	Port     int    `json:"port"`
	Service  struct {
		HTTP struct {
			MetaKeywords string `json:"meta_keywords"`
			Title        string `json:"title"`
			Icp          struct {
				Licence     string    `json:"licence"`
				UpdateTime  time.Time `json:"update_time"`
				IsExpired   bool      `json:"is_expired"`
				LeaderName  string    `json:"leader_name"`
				Domain      string    `json:"domain"`
				MainLicence struct {
					Licence string `json:"licence"`
					Unit    string `json:"unit"`
					Nature  string `json:"nature"`
				} `json:"main_licence"`
				ContentTypeName string `json:"content_type_name"`
				LimitAccess     bool   `json:"limit_access"`
			} `json:"icp"`
			Host string `json:"host"`
		} `json:"http"`
	} `json:"service"`
	IP     string    `json:"ip"`
	Domain string    `json:"domain"`
	Time   time.Time `json:"time"`
}
