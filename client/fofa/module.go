package fofa

type FoFaAccountInfo struct {
	Error   bool   `json:"error"`
	IsVip   bool   `json:"isvip"`
	Message string `json:"message"`
	ErrMsg  string `json:"errmsg"`
}

type Results struct {
	Error   bool       `json:"error"`
	Message string     `json:"message"`
	Size    int        `json:"size"`
	Query   string     `json:"query"`
	Page    int        `json:"page"`
	Results [][]string `json:"results"`
}
