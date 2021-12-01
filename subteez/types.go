package subteez

// possible status values
const (
	StatusOk          = "ok"
	StatusNotFound    = "not found"
	StatusBadRequest  = "bad request"
	StatusServerError = "server error"
)

type SearchResultItem struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Count int    `json:"count"`
}

type SubtitleFile struct {
	ID       string   `json:"id"`
	Language Language `json:"lang"`
	Name     string   `json:"name"`
	Author   string   `json:"author"`
	Comment  string   `json:"comment"`
	Title    string   `json:"title"`
}

type SearchRequest struct {
	Query           string     `json:"query"`
	LanguageFilters []Language `json:"lang"`
}

type SubtitleDetailsRequest struct {
	ID              string     `json:"id"`
	LanguageFilters []Language `json:"lang"`
}

type SubtitleDownloadRequest struct {
	ID string `json:"id" form:"id"`
}

type SearchResultResponse struct {
	Status string             `json:"status"`
	Result []SearchResultItem `json:"result"`
}

type SubtitleDetailsResponse struct {
	Status string         `json:"status"`
	Name   string         `json:"name"`
	Year   string         `json:"year"`
	Banner interface{}    `json:"posterUrl"`
	Files  []SubtitleFile `json:"files"`
}
