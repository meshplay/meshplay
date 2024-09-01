package constants

const (
	// Meshplay Repository Location
	meshplayGitHubOrg  string = "meshplay"
	meshplayGitHubRepo string = "meshplay"
	CallbackURLENV     string = "MESHPLAY_SERVER_CALLBACK_URL"
	ProviderENV        string = "PROVIDER"
	ProviderURLsENV    string = "PROVIDER_BASE_URLS"
)

// GetMeshplayGitHubOrg retrieves the name of the GitHub organization under which the Meshplay repository resides.
func GetMeshplayGitHubOrg() string {
	return meshplayGitHubOrg
}

// GetMeshplayGitHubRepo retrieves the name of the Meshplay repository
func GetMeshplayGitHubRepo() string {
	return meshplayGitHubRepo
}
