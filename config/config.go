package config

import (
	"os"
)

func CMSUrl () string {
	CmsUrl, cmsExists := os.LookupEnv("CMSUrl")
	if !cmsExists {
		CmsUrl = "default"
	}

	return CmsUrl
}

func CMSApiKey () string {
	CmsApiKey, keyExists := os.LookupEnv("CMSApiKey")
	if !keyExists {
		CmsApiKey = "default"
	}

	return CmsApiKey
}