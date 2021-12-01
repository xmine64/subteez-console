// Subteez Client Implementation

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

const searchEndpoint = "/api/search"
const detailsEndpoint = "/api/details"
const downloadEndpoint = "/api/download"

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

// send given request to given endpoint
func (client subteezClient) sendRequest(endpoint string, request interface{}) ([]byte, error) {
	// serialize request to json string
	requestJson, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	// create new http request
	httpRequest, err := http.NewRequest(
		http.MethodPost,
		client.baseAddress+endpoint,
		bytes.NewBuffer(requestJson),
	)
	if err != nil {
		return nil, err
	}
	httpRequest.Header.Add("Content-Type", "application/json")

	// send http request
	response, err := client.httpClient.Do(httpRequest)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// return error when status code is not ok
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

	// read response body and return it
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return responseBody, nil
}

func (client subteezClient) Search(request SearchRequest) (*SearchResultResponse, error) {
	rawResponse, err := client.sendRequest(searchEndpoint, request)
	if err != nil {
		return nil, err
	}

	// deserialize response body
	var result SearchResultResponse
	if json.Unmarshal(rawResponse, &result); err != nil {
		return nil, err
	}

	if result.Status != StatusOk {
		return nil, ErrUnhandledResponse(result.Status)
	}
	return &result, nil
}

func (client subteezClient) GetDetails(request SubtitleDetailsRequest) (*SubtitleDetailsResponse, error) {
	rawResponse, err := client.sendRequest(detailsEndpoint, request)
	if err != nil {
		return nil, err
	}

	// deserialize response body
	var result SubtitleDetailsResponse
	if err = json.Unmarshal(rawResponse, &result); err != nil {
		return nil, err
	}
	if result.Status != StatusOk {
		return nil, ErrUnhandledResponse(result.Status)
	}
	return &result, nil
}

func (client subteezClient) Download(request SubtitleDownloadRequest) (string, []byte, error) {
	// generate download link
	parameters := url.Values{"id": {request.ID}}.Encode()
	url := client.baseAddress + downloadEndpoint + "?" + parameters

	// download file
	response, err := client.httpClient.Get(url)
	if err != nil {
		return "", nil, err
	}
	defer response.Body.Close()

	// return error code it's not ok
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

	// find file name by parsing "Content-Disposition" header
	_, params, err := mime.ParseMediaType(response.Header.Get("Content-Disposition"))
	if err != nil {
		return "", nil, err
	}
	return params["filename"], data, nil
}
