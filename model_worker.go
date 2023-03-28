/*
 * Laravel Forge
 *
 * The Forge API allows you to create and interact with servers and sites on Laravel Forge through a simple REST API.
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package forge

type Worker struct {
	Id int32 `json:"id,omitempty"`
	Connection string `json:"connection,omitempty"`
	Command string `json:"command,omitempty"`
	Queue string `json:"queue,omitempty"`
	Timeout int32 `json:"timeout,omitempty"`
	Sleep int32 `json:"sleep,omitempty"`
	Tries string `json:"tries,omitempty"`
	Processes int32 `json:"processes,omitempty"`
	Stopwaitsecs string `json:"stopwaitsecs,omitempty"`
	Environment string `json:"environment,omitempty"`
	PhpVersion string `json:"php_version,omitempty"`
	Daemon int32 `json:"daemon,omitempty"`
	Force int32 `json:"force,omitempty"`
	Status string `json:"status,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
}