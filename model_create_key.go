/*
 * Laravel Forge
 *
 * The Forge API allows you to create and interact with servers and sites on Laravel Forge through a simple REST API.
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package forge

type CreateKey struct {
	Name string `json:"name,omitempty"`
	Key string `json:"key,omitempty"`
	Username string `json:"username,omitempty"`
}
