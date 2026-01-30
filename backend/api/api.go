package api

type API struct {
	OutputDir       string
	ServersFilePath string
}

// Constructor
func NewAPI(outputDir, serversFilePath string) *API {
	return &API{
		OutputDir:       outputDir,
		ServersFilePath: serversFilePath,
	}
}
