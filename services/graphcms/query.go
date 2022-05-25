package graphcms

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/impwellbeing/config"
	"github.com/machinebox/graphql"
)

func LandingCards(ctx *gin.Context) CardResponse {
	graphqlClient := graphql.NewClient(config.CMSUrl())
	graphqlRequest := graphql.NewRequest(`
		query MyQuery {
			landingPageCardsPlural {
			title
			content
			content2
			content3
			image {
				url
			}
			}
		}
	`)

	graphqlRequest.Header.Set("Authorization", fmt.Sprintf("Bearer %s", config.CMSApiKey()))
	var graphqlResponse CardResponse
	if err := graphqlClient.Run(ctx, graphqlRequest, &graphqlResponse); err != nil {
		panic(err)
	}

	return graphqlResponse
}

func ResourceCards(ctx *gin.Context) ResourceResponse {
	graphqlClient := graphql.NewClient(config.CMSUrl())
	graphqlRequest := graphql.NewRequest(`
		query MyQuery {
			resourcesPlural {
				title
				content1
				content2
				content3
				document {
				  fileName
				  url
				}
				image {
				  url
				}
				link1
				link2
				link3
				hasLink
			  }
		}
	`)

	graphqlRequest.Header.Set("Authorization", fmt.Sprintf("Bearer %s", config.CMSApiKey()))
	var graphqlResponse ResourceResponse
	if err := graphqlClient.Run(ctx, graphqlRequest, &graphqlResponse); err != nil {
		panic(err)
	}

	return graphqlResponse
}

func GetAllBlogs(ctx *gin.Context) BlogResponse {
	graphqlClient := graphql.NewClient(config.CMSUrl())
	graphqlRequest := graphql.NewRequest(`
		query MyQuery {
			blogs(orderBy: publishedAt_ASC) {
			id
			title
			dateCreated
			titleImage {
				url
			}
			}
		}
	`)

	graphqlRequest.Header.Set("Authorization", fmt.Sprintf("Bearer %s", config.CMSApiKey()))
	var graphqlResponse BlogResponse
	if err := graphqlClient.Run(ctx, graphqlRequest, &graphqlResponse); err != nil {
		panic(err)
	}

	return graphqlResponse
}
