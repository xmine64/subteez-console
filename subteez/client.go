package subteez

import (
	"bytes"
	"encoding/json"
	"io"
	"mime"
	"net/http"
	"net/url"
	"time"
)

const endpointSearch = "/api/search"
const endpointDetails = "/api/details"
const endpointDownload = "/api/download"

type subteezClient struct {
	baseAddress string
	httpClient  http.Client
}

func NewClient(server string) ISubteezAPI {
	return &subteezClient{
		baseAddress: server,
		httpClient: http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (client subteezClient) sendRequest(endpoint string, request interface{}) ([]byte, error) {
	requestJson, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	httpRequest, err := http.NewRequest(
		http.MethodPost,
		client.baseAddress+endpoint,
		bytes.NewBuffer(requestJson),
	)
	if err != nil {
		return nil, err
	}
	httpRequest.Header.Add("Content-Type", "application/json")
	response, err := client.httpClient.Do(httpRequest)
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		switch response.StatusCode {
		case http.StatusNotFound:
			return nil, ErrNotFound
		case http.StatusBadRequest:
			return nil, ErrBadRequest
		case http.StatusInternalServerError:
			return nil, ErrServer
		}
		return nil, ErrUnhandledResponse(response.Status)
	}
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return responseBody, nil
}

func (client subteezClient) Search(request SearchRequest) (*SearchResult, error) {
	responseRaw, err := client.sendRequest(endpointSearch, request)
	if err != nil {
		return nil, err
	}
	var result SearchResult
	err = json.Unmarshal(responseRaw, &result)
	if err != nil {
		return nil, err
	}
	if result.Status != StatusOk {
		return nil, ErrUnhandledResponse(result.Status)
	}
	return &result, nil
}

func (client subteezClient) GetDetails(request SubtitleDetailsRequest) (*SubtitleDetails, error) {
	responseRaw, err := client.sendRequest(endpointDetails, request)
	if err != nil {
		return nil, err
	}
	var result SubtitleDetails
	err = json.Unmarshal(responseRaw, &result)
	if err != nil {
		return nil, err
	}
	if result.Status != StatusOk {
		return nil, ErrUnhandledResponse(result.Status)
	}
	return &result, nil
}

func (client subteezClient) Download(request SubtitleDownloadRequest) (string, []byte, error) {
	parameters := url.Values{"id": {request.ID}}.Encode()
	url := client.baseAddress + endpointDownload + "?" + parameters
	response, err := client.httpClient.Get(url)
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return "", nil, err
	}
	if response.StatusCode != http.StatusOK {
		switch response.StatusCode {
		case http.StatusNotFound:
			return "", nil, ErrNotFound
		case http.StatusBadRequest:
			return "", nil, ErrBadRequest
		case http.StatusInternalServerError:
			return "", nil, ErrServer
		}
		return "", nil, ErrUnhandledResponse(response.Status)
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return "", nil, err
	}

	_, params, err := mime.ParseMediaType(response.Header.Get("Content-Disposition"))
	if err != nil {
		return "", nil, err
	}
	return params["filename"], data, nil
}
