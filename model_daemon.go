/*
 * Laravel Forge
 *
 * The Forge API allows you to create and interact with servers and sites on Laravel Forge through a simple REST API.
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package forge

type Daemon struct {
	Id int32 `json:"id,omitempty"`
	Command string `json:"command,omitempty"`
	User string `json:"user,omitempty"`
	Directory string `json:"directory,omitempty"`
	Processes int32 `json:"processes,omitempty"`
	Startsecs int32 `json:"startsecs,omitempty"`
	Stopwaitsecs int32 `json:"stopwaitsecs,omitempty"`
	Stopsignal string `json:"stopsignal,omitempty"`
	Status string `json:"status,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
}