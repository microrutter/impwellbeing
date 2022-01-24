package graphcms

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/impwellbeing/config"
	"github.com/machinebox/graphql"
)

type CardResponse struct {
	Cards []Card `json:"landingPageCardsPlural"`
}

type Card struct {
	Content string `json:"content"`
	Image map[string]string `json:"image"`
}

func LandingCards (ctx *gin.Context) CardResponse {
	CmsUrl, cmsExists := os.LookupEnv("CMSUrl")
	if !cmsExists {
		CmsUrl = config.CmsUrl
	}
	CmsApiKey, keyExists := os.LookupEnv("CMSApiKey")
	if !keyExists {
		CmsApiKey = config.CmsApiKey
	}
	graphqlClient := graphql.NewClient(CmsUrl)
	graphqlRequest := graphql.NewRequest(`
		query MyQuery {
			landingPageCardsPlural {
			content
			image {
				url
			}
			}
		}
	`)

	graphqlRequest.Header.Set("Authorization", fmt.Sprintf("Bearer %s", CmsApiKey))
	var graphqlResponse CardResponse
	if err := graphqlClient.Run(ctx, graphqlRequest, &graphqlResponse); err != nil {
        panic(err)
    }

	return graphqlResponse
}