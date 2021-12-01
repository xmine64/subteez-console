package subteez

type ISubteezAPI interface {
	Search(SearchRequest) (*SearchResultResponse, error)
	GetDetails(SubtitleDetailsRequest) (*SubtitleDetailsResponse, error)
	Download(SubtitleDownloadRequest) (fileName string, data []byte, err error)
}
