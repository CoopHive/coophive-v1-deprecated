package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"path/filepath"

	"github.com/CoopHive/coophive/pkg/data"
	"github.com/CoopHive/coophive/pkg/saas/filestore"
	"github.com/CoopHive/coophive/pkg/saas/job"
	"github.com/CoopHive/coophive/pkg/saas/types"
)

func (apiServer *CoopSaaSAPIServer) status(res http.ResponseWriter, req *http.Request) (types.UserStatus, error) {
	return apiServer.Controller.GetStatus(apiServer.getRequestContext(req))
}

func (apiServer *CoopSaaSAPIServer) getJobs(res http.ResponseWriter, req *http.Request) ([]*types.Job, error) {
	return apiServer.Controller.GetJobs(apiServer.getRequestContext(req))
}

func (apiServer *CoopSaaSAPIServer) getTransactions(res http.ResponseWriter, req *http.Request) ([]*types.BalanceTransfer, error) {
	return apiServer.Controller.GetTransactions(apiServer.getRequestContext(req))
}

func (apiServer *CoopSaaSAPIServer) getModules(res http.ResponseWriter, req *http.Request) ([]types.Module, error) {
	return job.GetModules()
}

type RunJobResults struct {
	JobOffer data.JobOfferContainer `json:"job_offer"`
	Result   data.Result            `json:"result"`
}

func (apiServer *CoopSaaSAPIServer) createJobSync(res http.ResponseWriter, req *http.Request) (*RunJobResults, error) {
	request := types.JobSpec{}
	bs, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bs, &request)
	if err != nil {
		return nil, err
	}
	result, err := apiServer.Controller.CreateJobSync(apiServer.getRequestContext(req), request)
	if err != nil {
		return nil, err
	}
	return &RunJobResults{
		JobOffer: result.JobOffer,
		Result:   result.Result,
	}, nil
}

func (apiServer *CoopSaaSAPIServer) createJobAsync(res http.ResponseWriter, req *http.Request) (string, error) {
	request := types.JobSpec{}
	bs, err := io.ReadAll(req.Body)
	if err != nil {
		return "", err
	}
	err = json.Unmarshal(bs, &request)
	if err != nil {
		return "", err
	}
	return apiServer.Controller.CreateJobAsync(apiServer.getRequestContext(req), request)
}

func (apiServer *CoopSaaSAPIServer) filestoreConfig(res http.ResponseWriter, req *http.Request) (filestore.FilestoreConfig, error) {
	return apiServer.Controller.FilestoreConfig(apiServer.getRequestContext(req))
}

func (apiServer *CoopSaaSAPIServer) filestoreList(res http.ResponseWriter, req *http.Request) ([]filestore.FileStoreItem, error) {
	return apiServer.Controller.FilestoreList(apiServer.getRequestContext(req), req.URL.Query().Get("path"))
}

func (apiServer *CoopSaaSAPIServer) filestoreGet(res http.ResponseWriter, req *http.Request) (filestore.FileStoreItem, error) {
	return apiServer.Controller.FilestoreGet(apiServer.getRequestContext(req), req.URL.Query().Get("path"))
}

func (apiServer *CoopSaaSAPIServer) filestoreCreateFolder(res http.ResponseWriter, req *http.Request) (filestore.FileStoreItem, error) {
	return apiServer.Controller.FilestoreCreateFolder(apiServer.getRequestContext(req), req.URL.Query().Get("path"))
}

func (apiServer *CoopSaaSAPIServer) filestoreRename(res http.ResponseWriter, req *http.Request) (filestore.FileStoreItem, error) {
	return apiServer.Controller.FilestoreRename(apiServer.getRequestContext(req), req.URL.Query().Get("path"), req.URL.Query().Get("new_path"))
}

func (apiServer *CoopSaaSAPIServer) filestoreDelete(res http.ResponseWriter, req *http.Request) (string, error) {
	path := req.URL.Query().Get("path")
	err := apiServer.Controller.FilestoreDelete(apiServer.getRequestContext(req), path)
	return path, err
}

func (apiServer *CoopSaaSAPIServer) filestoreUpload(res http.ResponseWriter, req *http.Request) (bool, error) {
	path := req.URL.Query().Get("path")
	err := req.ParseMultipartForm(10 << 20)
	if err != nil {
		return false, err
	}

	files := req.MultipartForm.File["files"]
	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			return false, fmt.Errorf("Unable to open file")
		}
		defer file.Close()
		_, err = apiServer.Controller.FilestoreUpload(apiServer.getRequestContext(req), filepath.Join(path, fileHeader.Filename), file)
		if err != nil {
			return false, fmt.Errorf("Unable to upload file: %s", err.Error())
		}
	}

	return true, nil
}

func (apiServer *CoopSaaSAPIServer) createAPIKey(res http.ResponseWriter, req *http.Request) (string, error) {
	name := req.URL.Query().Get("name")
	apiKey, err := apiServer.Controller.CreateAPIKey(apiServer.getRequestContext(req), name)
	if err != nil {
		return "", err
	}
	return apiKey, nil
}

func (apiServer *CoopSaaSAPIServer) getAPIKeys(res http.ResponseWriter, req *http.Request) ([]*types.ApiKey, error) {
	apiKeys, err := apiServer.Controller.GetAPIKeys(apiServer.getRequestContext(req))
	if err != nil {
		return nil, err
	}
	return apiKeys, nil
}

func (apiServer *CoopSaaSAPIServer) deleteAPIKey(res http.ResponseWriter, req *http.Request) (string, error) {
	apiKey := req.URL.Query().Get("key")
	err := apiServer.Controller.DeleteAPIKey(apiServer.getRequestContext(req), apiKey)
	if err != nil {
		return "", err
	}
	return "", nil
}

func (apiServer *CoopSaaSAPIServer) checkAPIKey(res http.ResponseWriter, req *http.Request) (*types.ApiKey, error) {
	apiKey := req.URL.Query().Get("key")
	key, err := apiServer.Controller.CheckAPIKey(apiServer.getRequestContext(req).Ctx, apiKey)
	if err != nil {
		return nil, err
	}
	return key, nil
}
