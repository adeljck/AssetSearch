package quake

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
	Data    []struct {
		Host     string `json:"host"`
		Domain   string `json:"domain"`
		Time     string `json:"time"`
		Hostname string `json:"hostname"`
		Port     int    `json:"port"`
		IP       string `json:"ip"`
		Service  struct {
			Http struct {
				Icp struct {
					License     string `json:"license"`
					UpdateTime  string `json:"update_time"`
					IsExpired   bool   `json:"is_expired"`
					Domain      string `json:"domain"`
					LeaderName  string `json:"leader_name"`
					MainLicense struct {
						License string `json:"license"`
						Unit    string `json:"unit"`
						Nature  string `json:"nature"`
					} `json:"main_license"`
				} `json:"icp"`
			} `json:"http"`
		} `json:"service"`
	} `json:"data"`
	Meta struct {
		Pagination struct {
			Count     int `json:"count"`
			PageIndex int `json:"page_index"`
			PageSize  int `json:"page_size"`
			Total     int `json:"total"`
		} `json:"pagination"`
	} `json:"meta"`
}
