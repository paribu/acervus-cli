package api

import "fmt"

type authEndpoints struct {
	register        string
	login           string
	refreshToken    string
	recoverPassword string
	resetPassword   string
	logout          string
}
type generateEndpoints struct {
	boilerplate func(projectID string) string
	graphql     func(projectID string) string
}
type projectEndpoints struct {
	list   string
	create string
	deploy func(projectID string) string
	export func(projectID string) string
	test   func(projectID string) string
	pause  func(projectID string) string
	resume func(projectID string) string
	delete func(projectID string) string
}
type dataEndpoints struct {
	list string
}
type logEndpoints struct {
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
	log      logEndpoints
	network  networkEndpoints
}{
	auth: authEndpoints{
		register:        "auth/register",
		login:           "auth/login",
		refreshToken:    "auth/refresh",
		recoverPassword: "auth/recover-password",
		resetPassword:   "auth/reset-password",
		logout:          "auth/logout",
	},
	generate: generateEndpoints{
		boilerplate: func(projectID string) string {
			return fmt.Sprintf("generate/%s/boilerplate", projectID)
		},
		graphql: func(projectID string) string {
			return fmt.Sprintf("generate/%s/graphql", projectID)
		},
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
		pause: func(projectID string) string {
			return fmt.Sprintf("projects/%s/pause", projectID)
		},
		resume: func(projectID string) string {
			return fmt.Sprintf("projects/%s/resume", projectID)
		},
		delete: func(projectID string) string {
			return fmt.Sprintf("projects/%s", projectID)
		},
	},
	data: dataEndpoints{
		list: "project-data",
	},
	log: logEndpoints{
		list: "project-logs",
	},
	network: networkEndpoints{
		list: "networks",
	},
}
