package subteez

const StatusOk = "ok"
const StatusNotFound = "not found"
const StatusBadRequest = "bad request"
const StatusServerError = "server error"

type Language string

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

type SearchResultItem struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Count int    `json:"count"`
}

type SearchResult struct {
	Status string             `json:"status"`
	Result []SearchResultItem `json:"result"`
}

type SubtitleFile struct {
	ID       string   `json:"id"`
	Language Language `json:"lang"`
	Name     string   `json:"name"`
	Author   string   `json:"author"`
	Comment  string   `json:"comment"`
	Title    string   `json:"title"`
}

type SubtitleDetails struct {
	Status string         `json:"status"`
	Name   string         `json:"name"`
	Year   string         `json:"year"`
	Banner interface{}    `json:"posterUrl"`
	Files  []SubtitleFile `json:"files"`
}
