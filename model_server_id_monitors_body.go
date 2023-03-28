/*
 * Laravel Forge
 *
 * The Forge API allows you to create and interact with servers and sites on Laravel Forge through a simple REST API.
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package forge

type ServerIdMonitorsBody struct {
	Type_ string `json:"type,omitempty"`
	Operator string `json:"operator,omitempty"`
	Threshold string `json:"threshold,omitempty"`
	Minutes string `json:"minutes,omitempty"`
	Notify string `json:"notify,omitempty"`
}
