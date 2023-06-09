/*
 * Laravel Forge
 *
 * The Forge API allows you to create and interact with servers and sites on Laravel Forge through a simple REST API.
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package forge

type CreateSite struct {
	Domain string `json:"domain,omitempty"`
	ProjectType string `json:"project_type,omitempty"`
	Aliases string `json:"aliases,omitempty"`
	Directory string `json:"directory,omitempty"`
	Isolated string `json:"isolated,omitempty"`
	Username string `json:"username,omitempty"`
	Database string `json:"database,omitempty"`
	PhpVersion string `json:"php_version,omitempty"`
	NginxTemplate string `json:"nginx_template,omitempty"`
}
