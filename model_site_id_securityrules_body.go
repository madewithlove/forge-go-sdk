/*
 * Laravel Forge
 *
 * The Forge API allows you to create and interact with servers and sites on Laravel Forge through a simple REST API.
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package forge

type SiteIdSecurityrulesBody struct {
	Name string `json:"name,omitempty"`
	Path string `json:"path,omitempty"`
	Credentials string `json:"credentials,omitempty"`
}
