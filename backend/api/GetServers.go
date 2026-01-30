package api

import servermanager "mvp/backend/server_manager"

func (a *API) GetServers() ([]servermanager.Server, error) {
	servers, err := servermanager.LoadServers(a.ServersFilePath)
	if err != nil {
		return nil, err
	}

	return servers, nil
}
