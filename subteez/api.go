package subteez

type ISubteezAPI interface {
	Search(SearchRequest) (*SearchResult, error)
	GetDetails(SubtitleDetailsRequest) (*SubtitleDetails, error)
	Download(SubtitleDownloadRequest) (fileName string, data []byte, err error)
}
