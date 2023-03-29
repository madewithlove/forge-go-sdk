/*
 * Laravel Forge
 *
 * The Forge API allows you to create and interact with servers and sites on Laravel Forge through a simple REST API.
 *
 * API version: 0.0.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package forge

type RecipesBody struct {
	Name   string `json:"name,omitempty"`
	User   string `json:"user,omitempty"`
	Script string `json:"script,omitempty"`
}
