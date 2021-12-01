package api

import (
	"errors"
	"strings"

	"github.com/Octy-ai/octy-cli/pkg/globals"
)

//
// Public methods
//

func (api Application) VersionAssesment(osInfo map[string]string) (bool, []string, error) {
	version, errs := api.rest.GetVersionInfo()
	if errs != nil {
		return false, nil, errs
	}

	currentVersionAssetURLS := make([]string, 0)
	for _, asset := range version.Assets {
		currentVersionAssetURLS = append(currentVersionAssetURLS, asset.BrowserDownloadURL)
	}

	// assess if current cli version is up-to-date
	if globals.CliVersion != version.VersionName {
		switch osInfo["os"] { // switch top level os
		case "darwin":
			// macos architectures
			switch osInfo["arch"] {
			case "amd64":
				return true, identifyAssetURLS("darwin", "amd64", currentVersionAssetURLS), nil
			case "arm64":
				return true, identifyAssetURLS("darwin", "arm64", currentVersionAssetURLS), nil
			}
		case "linux":
			// linux architectures
			switch osInfo["arch"] {
			case "386":
				return true, identifyAssetURLS("linux", "386", currentVersionAssetURLS), nil
			case "amd64":
				return true, identifyAssetURLS("linux", "amd64", currentVersionAssetURLS), nil
			case "arm64":
				return true, identifyAssetURLS("linux", "arm64", currentVersionAssetURLS), nil
			case "armv6":
				return true, identifyAssetURLS("linux", "armv6", currentVersionAssetURLS), nil
			case "armv7":
				return true, identifyAssetURLS("linux", "armv7", currentVersionAssetURLS), nil

			}
		case "windows":
			// windows architectures
			switch osInfo["arch"] {
			case "386":
				return true, identifyAssetURLS("windows", "386", currentVersionAssetURLS), nil
			case "amd64":
				return true, identifyAssetURLS("windows", "amd64", currentVersionAssetURLS), nil
			case "armv6":
				return true, identifyAssetURLS("windows", "armv6", currentVersionAssetURLS), nil
			case "armv7":
				return true, identifyAssetURLS("windows", "armv7", currentVersionAssetURLS), nil
			}
		default:
			return false, nil, errors.New("could not identify operating system")
		}
	}

	// app does not require update
	return false, nil, nil

}

//
// Private functions
//

func identifyAssetURLS(os string, arch string, assetURLS []string) []string {
	archAssetURLS := make([]string, 0)
	for _, url := range assetURLS {
		if strings.Contains(url, arch) && strings.Contains(url, os) {
			archAssetURLS = append(archAssetURLS, url)
		}
	}
	return archAssetURLS
}
