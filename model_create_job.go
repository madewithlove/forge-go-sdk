/*
 * Laravel Forge
 *
 * The Forge API allows you to create and interact with servers and sites on Laravel Forge through a simple REST API.
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package forge

type CreateJob struct {
	Frequency string `json:"frequency,omitempty"`
	Minute string `json:"minute,omitempty"`
	Hour string `json:"hour,omitempty"`
	Day string `json:"day,omitempty"`
	Month string `json:"month,omitempty"`
	Weekday string `json:"weekday,omitempty"`
}
