package api

import "fmt"

type authEndpoints struct {
	register     string
	login        string
	refreshToken string
	logout       string
}
type generateEndpoints struct {
	boilerplate string
	graphql     string
}
type projectEndpoints struct {
	list   string
	create string
	deploy func(projectID string) string
	export func(projectID string) string
	test   func(projectID string) string
	delete func(projectID string) string
}
type dataEndpoints struct {
	list string
}
type networkEndpoints struct {
	list string
}

var endpoints = struct {
	auth     authEndpoints
	generate generateEndpoints
	project  projectEndpoints
	data     dataEndpoints
	network  networkEndpoints
}{
	auth: authEndpoints{
		register:     "auth/register",
		login:        "auth/login",
		refreshToken: "auth/refresh",
		logout:       "auth/logout",
	},
	generate: generateEndpoints{
		boilerplate: "generate/boilerplate",
		graphql:     "generate/graphql",
	},
	project: projectEndpoints{
		list:   "projects",
		create: "projects/create",
		deploy: func(projectID string) string {
			return fmt.Sprintf("projects/%s/deploy", projectID)
		},
		export: func(projectID string) string {
			return fmt.Sprintf("projects/%s/export", projectID)
		},
		test: func(projectID string) string {
			return fmt.Sprintf("projects/%s/test", projectID)
		},
		delete: func(projectID string) string {
			return fmt.Sprintf("projects/%s", projectID)
		},
	},
	data: dataEndpoints{
		list: "project-data",
	},
	network: networkEndpoints{
		list: "networks",
	},
}
