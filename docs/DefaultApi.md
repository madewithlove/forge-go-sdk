# {{classname}}

All URIs are relative to *https://forge.laravel.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateCertificate**](DefaultApi.md#ActivateCertificate) | **Post** /servers/{serverId}/sites/{siteId}/certificates/{id}/activate | Activate Certificate
[**AddSiteAliases**](DefaultApi.md#AddSiteAliases) | **Put** /servers/{serverId}/sites/{siteId}/aliases | Add Site Aliases
[**ChangeSitePHPVersion**](DefaultApi.md#ChangeSitePHPVersion) | **Put** /servers/{serverId}/sites/{siteId}/php | Change Site PHP Version
[**CloningAnExistingCertificate**](DefaultApi.md#CloningAnExistingCertificate) | **Post** /servers/{serverId}/sites/{siteId}/certificates | Cloning An Existing Certificate
[**Create**](DefaultApi.md#Create) | **Post** /servers/{server_id}/sites/{site_id}/webhooks | Create
[**CreateBackupConfiguration**](DefaultApi.md#CreateBackupConfiguration) | **Post** /servers/{serverId}/backup-configs | Create Backup Configuration
[**CreateDaemon**](DefaultApi.md#CreateDaemon) | **Post** /servers/{serverId}/daemons | Create Daemon
[**CreateDatabase**](DefaultApi.md#CreateDatabase) | **Post** /servers/{serverId}/databases | Create Database
[**CreateDeployKey**](DefaultApi.md#CreateDeployKey) | **Post** /servers/{serverId}/sites/{siteId}/deploy-key | Create Deploy Key
[**CreateFirewallRule**](DefaultApi.md#CreateFirewallRule) | **Post** /servers/{serverId}/firewall-rules | Create Firewall Rule
[**CreateJob**](DefaultApi.md#CreateJob) | **Post** /servers/{serverId}/jobs | Create Job
[**CreateKey**](DefaultApi.md#CreateKey) | **Post** /servers/{serverId}/keys | Create Key
[**CreateMonitor**](DefaultApi.md#CreateMonitor) | **Post** /servers/{serverId}/monitors | Create Monitor
[**CreateRecipe**](DefaultApi.md#CreateRecipe) | **Post** /recipes | Create Recipe
[**CreateRedirectRule**](DefaultApi.md#CreateRedirectRule) | **Post** /servers/{serverId}/sites/{siteId}/redirect-rules | Create Redirect Rule
[**CreateSecurityRule**](DefaultApi.md#CreateSecurityRule) | **Post** /servers/{serverId}/sites/{siteId}/security-rules | Create Security Rule
[**CreateServer**](DefaultApi.md#CreateServer) | **Post** /servers | Create Server
[**CreateSite**](DefaultApi.md#CreateSite) | **Post** /servers/{serverId}/sites | Create Site
[**CreateTemplate**](DefaultApi.md#CreateTemplate) | **Post** /servers/{serverId}/nginx/templates | Create Template
[**CreateUser**](DefaultApi.md#CreateUser) | **Post** /servers/{serverId}/database-users | Create User
[**CreateWorker**](DefaultApi.md#CreateWorker) | **Post** /servers/{serverId}/sites/{siteId}/workers | Create Worker
[**DeleteBackup**](DefaultApi.md#DeleteBackup) | **Delete** /servers/{serverId}/backup-configs/{backupConfigurationId}/backups/{backupId} | Delete Backup
[**DeleteBackupConfiguration**](DefaultApi.md#DeleteBackupConfiguration) | **Delete** /servers/{serverId}/backup-configs/{backupConfigurationId} | Delete Backup Configuration
[**DeleteCertificate**](DefaultApi.md#DeleteCertificate) | **Delete** /servers/{serverId}/sites/{siteId}/certificates/{id} | Delete Certificate
[**DeleteDaemon**](DefaultApi.md#DeleteDaemon) | **Delete** /servers/{serverId}/daemons/{daemonId} | Delete Daemon
[**DeleteDatabase**](DefaultApi.md#DeleteDatabase) | **Delete** /servers/{serverId}/databases/{databaseId} | Delete Database
[**DeleteDeployKey**](DefaultApi.md#DeleteDeployKey) | **Delete** /servers/{serverId}/sites/{siteId}/deploy-key | Delete Deploy Key
[**DeleteFirewallRule**](DefaultApi.md#DeleteFirewallRule) | **Delete** /servers/{serverId}/firewall-rules/{ruleId} | Delete Firewall Rule
[**DeleteJob**](DefaultApi.md#DeleteJob) | **Delete** /servers/{serverId}/jobs/{jobId} | Delete Job
[**DeleteKey**](DefaultApi.md#DeleteKey) | **Delete** /servers/{serverId}/keys/{keyId} | Delete Key
[**DeleteMonitor**](DefaultApi.md#DeleteMonitor) | **Delete** /servers/{serverId}/monitors/{monitorId} | Delete Monitor
[**DeleteNginxTemplate**](DefaultApi.md#DeleteNginxTemplate) | **Delete** /servers/{serverId}/nginx/templates/{templateId} | Delete Nginx Template
[**DeleteRecipe**](DefaultApi.md#DeleteRecipe) | **Delete** /recipes/{recipeId} | Delete Recipe
[**DeleteRedirectRule**](DefaultApi.md#DeleteRedirectRule) | **Delete** /servers/{serverId}/sites/{siteId}/redirect-rules/{id} | Delete Redirect Rule
[**DeleteSecurityRule**](DefaultApi.md#DeleteSecurityRule) | **Delete** /servers/{serverId}/sites/{siteId}/security-rules/{id} | Delete Security Rule
[**DeleteServer**](DefaultApi.md#DeleteServer) | **Delete** /servers/{id} | Delete Server
[**DeleteSite**](DefaultApi.md#DeleteSite) | **Delete** /servers/{serverId}/sites/{siteId} | Delete Site
[**DeleteUser**](DefaultApi.md#DeleteUser) | **Delete** /servers/{serverId}/database-users/{userId} | Delete User
[**DeleteWebhook**](DefaultApi.md#DeleteWebhook) | **Delete** /servers/{server_id}/sites/{site_id}/webhooks/{id} | Delete Webhook
[**DeleteWorker**](DefaultApi.md#DeleteWorker) | **Delete** /servers/{serverId}/sites/{siteId}/workers/{id} | Delete Worker
[**DeployNow**](DefaultApi.md#DeployNow) | **Post** /servers/{serverId}/sites/{siteId}/deployment/deploy | Deploy Now
[**DisableOPCache**](DefaultApi.md#DisableOPCache) | **Delete** /servers/{serverId}/php/opcache | Disable OPCache
[**DisableQuickDeployment**](DefaultApi.md#DisableQuickDeployment) | **Delete** /servers/{serverId}/sites/{siteId}/deployment | Disable Quick Deployment
[**EnableOPCache**](DefaultApi.md#EnableOPCache) | **Post** /servers/{serverId}/php/opcache | Enable OPCache
[**EnableQuickDeployment**](DefaultApi.md#EnableQuickDeployment) | **Post** /servers/{serverId}/sites/{siteId}/deployment | Enable Quick Deployment
[**ExecuteCommand**](DefaultApi.md#ExecuteCommand) | **Post** /servers/{serverId}/sites/{siteId}/commands | Execute Command
[**GetBackupConfiguration**](DefaultApi.md#GetBackupConfiguration) | **Get** /servers/{serverId}/backup-configs/{backupConfigurationId} | Get Backup Configuration
[**GetCertificate**](DefaultApi.md#GetCertificate) | **Get** /servers/{serverId}/sites/{siteId}/certificates/{id} | Get Certificate
[**GetCommand**](DefaultApi.md#GetCommand) | **Get** /servers/{serverId}/sites/{siteId}/commands/{commandId} | Get Command
[**GetCredentials**](DefaultApi.md#GetCredentials) | **Get** /credentials | Get Credentials
[**GetDaemon**](DefaultApi.md#GetDaemon) | **Get** /servers/{serverId}/daemons/{daemonId} | Get Daemon
[**GetDatabase**](DefaultApi.md#GetDatabase) | **Get** /servers/{serverId}/databases/{databaseId} | Get Database
[**GetDeployment**](DefaultApi.md#GetDeployment) | **Get** /servers/{serverId}/sites/{siteId}/deployment-history/{deploymentId} | Get Deployment
[**GetDeploymentLog**](DefaultApi.md#GetDeploymentLog) | **Get** /servers/{serverId}/sites/{siteId}/deployment/log | Get Deployment Log
[**GetDeploymentOutput**](DefaultApi.md#GetDeploymentOutput) | **Get** /servers/{serverId}/sites/{siteId}/deployment-history/{deploymentId}/output | Get Deployment Output
[**GetDeploymentScript**](DefaultApi.md#GetDeploymentScript) | **Get** /servers/{serverId}/sites/{siteId}/deployment/script | Get Deployment Script
[**GetEnvFile**](DefaultApi.md#GetEnvFile) | **Get** /servers/{serverId}/sites/{siteId}/env | Get .env File
[**GetFirewallRule**](DefaultApi.md#GetFirewallRule) | **Get** /servers/{serverId}/firewall-rules/{ruleId} | Get Firewall Rule
[**GetJob**](DefaultApi.md#GetJob) | **Get** /servers/{serverId}/jobs/{jobId} | Get Job
[**GetKey**](DefaultApi.md#GetKey) | **Get** /servers/{serverId}/keys/{keyId} | Get Key
[**GetLog**](DefaultApi.md#GetLog) | **Get** /servers/{serverId}/logs | Get Log
[**GetMonitor**](DefaultApi.md#GetMonitor) | **Get** /servers/{serverId}/monitors/{monitorId} | Get Monitor
[**GetNginxConfiguration**](DefaultApi.md#GetNginxConfiguration) | **Get** /servers/{serverId}/sites/{siteId}/nginx | Get Nginx Configuration
[**GetNginxTemplate**](DefaultApi.md#GetNginxTemplate) | **Get** /servers/{serverId}/nginx/templates/{templateId} | Get Nginx Template
[**GetRecentEvents**](DefaultApi.md#GetRecentEvents) | **Get** /servers/events | Get Recent Events
[**GetRecipe**](DefaultApi.md#GetRecipe) | **Get** /recipes/{recipeId} | Get Recipe
[**GetRedirectRule**](DefaultApi.md#GetRedirectRule) | **Get** /servers/{serverId}/sites/{siteId}/redirect-rules/{id} | Get Redirect Rule
[**GetRegions**](DefaultApi.md#GetRegions) | **Get** /regions | Get Regions
[**GetSecurityRule**](DefaultApi.md#GetSecurityRule) | **Get** /servers/{serverId}/sites/{siteId}/security-rules/{id} | Get Security Rule
[**GetServer**](DefaultApi.md#GetServer) | **Get** /servers/{id} | Get Server
[**GetSigningRequest**](DefaultApi.md#GetSigningRequest) | **Get** /servers/{serverId}/sites/{siteId}/certificates/{id}/csr | Get Signing Request
[**GetSite**](DefaultApi.md#GetSite) | **Get** /servers/{serverId}/sites/{siteId} | Get Site
[**GetUser**](DefaultApi.md#GetUser) | **Get** /servers/{serverId}/database-users/{userId} | Get User
[**GetWebhook**](DefaultApi.md#GetWebhook) | **Get** /servers/{server_id}/sites/{site_id}/webhooks/{id} | Get Webhook
[**GetWebhooks**](DefaultApi.md#GetWebhooks) | **Get** /servers/{server_id}/sites/{site_id}/webhooks | Get Webhooks
[**GetWorker**](DefaultApi.md#GetWorker) | **Get** /servers/{serverId}/sites/{siteId}/workers/{id} | Get Worker
[**InstallBlackfire**](DefaultApi.md#InstallBlackfire) | **Post** /servers/{id}/blackfire/install | Install Blackfire
[**InstallCertificate**](DefaultApi.md#InstallCertificate) | **Post** /servers/{serverId}/sites/{siteId}/certificates/{id}/install | Install Certificate
[**InstallNewGitProject**](DefaultApi.md#InstallNewGitProject) | **Post** /servers/{serverId}/sites/{siteId}/git | Install New Git Project
[**InstallPHPVersion**](DefaultApi.md#InstallPHPVersion) | **Post** /servers/{serverId}/php | Install PHP Version
[**InstallPapertrail**](DefaultApi.md#InstallPapertrail) | **Post** /servers/{id}/papertrail/install | Install Papertrail
[**InstallPhpMyAdmin**](DefaultApi.md#InstallPhpMyAdmin) | **Post** /servers/{serverId}/sites/{siteId}/phpmyadmin | Install phpMyAdmin
[**InstallWordPress**](DefaultApi.md#InstallWordPress) | **Post** /servers/{serverId}/sites/{siteId}/wordpress | Install WordPress
[**ListBackupConfigurations**](DefaultApi.md#ListBackupConfigurations) | **Get** /servers/{serverId}/backup-configs | List Backup Configurations
[**ListCertificates**](DefaultApi.md#ListCertificates) | **Get** /servers/{serverId}/sites/{siteId}/certificates | List Certificates
[**ListCommandHistory**](DefaultApi.md#ListCommandHistory) | **Get** /servers/{serverId}/sites/{siteId}/commands | List Command History
[**ListDaemons**](DefaultApi.md#ListDaemons) | **Get** /servers/{serverId}/daemons | List Daemons
[**ListDatabases**](DefaultApi.md#ListDatabases) | **Get** /servers/{serverId}/databases | List Databases
[**ListDeployments**](DefaultApi.md#ListDeployments) | **Get** /servers/{serverId}/sites/{siteId}/deployment-history | List Deployments
[**ListFirewallRules**](DefaultApi.md#ListFirewallRules) | **Get** /servers/{serverId}/firewall-rules | List Firewall Rules
[**ListJobs**](DefaultApi.md#ListJobs) | **Get** /servers/{serverId}/jobs | List Jobs
[**ListKeys**](DefaultApi.md#ListKeys) | **Get** /servers/{serverId}/keys | List Keys
[**ListMonitors**](DefaultApi.md#ListMonitors) | **Get** /servers/{serverId}/monitors | List Monitors
[**ListNginxTemplates**](DefaultApi.md#ListNginxTemplates) | **Get** /servers/{serverId}/nginx/templates/default | List Nginx Templates
[**ListPHPVersions**](DefaultApi.md#ListPHPVersions) | **Get** /servers/{serverId}/php | List PHP Versions
[**ListRecipes**](DefaultApi.md#ListRecipes) | **Get** /recipes | List Recipes
[**ListRedirectRules**](DefaultApi.md#ListRedirectRules) | **Get** /servers/{serverId}/sites/{siteId}/redirect-rules | List Redirect Rules
[**ListSecurityRules**](DefaultApi.md#ListSecurityRules) | **Get** /servers/{serverId}/sites/{siteId}/security-rules | List Security Rules
[**ListServers**](DefaultApi.md#ListServers) | **Get** /servers | List Servers
[**ListSites**](DefaultApi.md#ListSites) | **Get** /servers/{serverId}/sites | List Sites
[**ListUsers**](DefaultApi.md#ListUsers) | **Get** /servers/{serverId}/database-users | List Users
[**ListWorkers**](DefaultApi.md#ListWorkers) | **Get** /servers/{serverId}/sites/{siteId}/workers | List Workers
[**LoadBalancing**](DefaultApi.md#LoadBalancing) | **Get** /servers/{serverId}/sites/{siteId}/balancing | Load Balancing
[**ObtainALetsEncryptCertificate**](DefaultApi.md#ObtainALetsEncryptCertificate) | **Post** /servers/{serverId}/sites/{siteId}/certificates/letsencrypt | Obtain A LetsEncrypt Certificate
[**ReactivateRevokedServer**](DefaultApi.md#ReactivateRevokedServer) | **Post** /servers/{id}/reactivate | Reactivate revoked server
[**RebootMySQL**](DefaultApi.md#RebootMySQL) | **Post** /servers/{id}/mysql/reboot | Reboot MySQL
[**RebootNginx**](DefaultApi.md#RebootNginx) | **Post** /servers/{id}/nginx/reboot | Reboot Nginx
[**RebootPHP**](DefaultApi.md#RebootPHP) | **Post** /servers/{id}/php/reboot | Reboot PHP
[**RebootPostgres**](DefaultApi.md#RebootPostgres) | **Post** /servers/{id}/postgres/reboot | Reboot Postgres
[**RebootServer**](DefaultApi.md#RebootServer) | **Post** /servers/{id}/reboot | Reboot Server
[**ReconnectRevokedServer**](DefaultApi.md#ReconnectRevokedServer) | **Post** /servers/{id}/reconnect | Reconnect revoked server
[**RemoveBlackfire**](DefaultApi.md#RemoveBlackfire) | **Delete** /servers/{id}/blackfire/remove | Remove Blackfire
[**RemovePapertrail**](DefaultApi.md#RemovePapertrail) | **Delete** /servers/{id}/papertrail/remove | Remove Papertrail
[**RemoveProject**](DefaultApi.md#RemoveProject) | **Delete** /servers/{serverId}/sites/{siteId}/git | Remove Project
[**ResetDeploymentStatus**](DefaultApi.md#ResetDeploymentStatus) | **Post** /servers/{serverId}/sites/{siteId}/deployment/reset | Reset Deployment Status
[**RestartDaemon**](DefaultApi.md#RestartDaemon) | **Post** /servers/{serverId}/daemons/{daemonId}/restart | Restart Daemon
[**RestartService**](DefaultApi.md#RestartService) | **Post** /servers/{id}/services/restart | Restart Service
[**RestartWorker**](DefaultApi.md#RestartWorker) | **Post** /servers/{serverId}/sites/{siteId}/workers/{id}/restart | Restart Worker
[**RestoreBackup**](DefaultApi.md#RestoreBackup) | **Post** /servers/{serverId}/backup-configs/{backupConfigurationId}/backups/{backupId} | Restore Backup
[**RevokeForgeAccessToServer**](DefaultApi.md#RevokeForgeAccessToServer) | **Post** /servers/{id}/revoke | Revoke Forge access to server
[**RunBackupConfiguration**](DefaultApi.md#RunBackupConfiguration) | **Post** /servers/{serverId}/backup-configs/{backupConfigurationId} | Run Backup Configuration
[**RunRecipe**](DefaultApi.md#RunRecipe) | **Post** /recipes/{recipeId}/run | Run Recipe
[**Show**](DefaultApi.md#Show) | **Get** /user | show
[**SiteLog**](DefaultApi.md#SiteLog) | **Get** /servers/{serverId}/sites/{siteId}/logs | Site Log
[**StartService**](DefaultApi.md#StartService) | **Post** /servers/{id}/services/start | Start Service
[**StopMySQL**](DefaultApi.md#StopMySQL) | **Post** /servers/{id}/mysql/stop | Stop MySQL
[**StopNginx**](DefaultApi.md#StopNginx) | **Post** /servers/{id}/nginx/stop | Stop Nginx
[**StopPostgres**](DefaultApi.md#StopPostgres) | **Post** /servers/{id}/postgres/stop | Stop Postgres
[**StopService**](DefaultApi.md#StopService) | **Post** /servers/{id}/services/stop | Stop Service
[**SyncDatabase**](DefaultApi.md#SyncDatabase) | **Post** /servers/{serverId}/databases/sync | Sync Database
[**TestNginx**](DefaultApi.md#TestNginx) | **Get** /servers/{id}/nginx/test | Test Nginx
[**UninstallPhpMyAdmin**](DefaultApi.md#UninstallPhpMyAdmin) | **Delete** /servers/{serverId}/sites/{siteId}/phpmyadmin | Uninstall phpMyAdmin
[**UninstallWordPress**](DefaultApi.md#UninstallWordPress) | **Delete** /servers/{serverId}/sites/{siteId}/wordpress | Uninstall WordPress
[**UpdateBackupConfiguration**](DefaultApi.md#UpdateBackupConfiguration) | **Put** /servers/{serverId}/backup-configs/{backupConfigurationId} | Update Backup Configuration
[**UpdateDatabasePassword**](DefaultApi.md#UpdateDatabasePassword) | **Put** /servers/{serverId}/database-password | Update Database Password
[**UpdateDeploymentScript**](DefaultApi.md#UpdateDeploymentScript) | **Put** /servers/{serverId}/sites/{siteId}/deployment/script | Update Deployment Script
[**UpdateEnvFile**](DefaultApi.md#UpdateEnvFile) | **Put** /servers/{serverId}/sites/{siteId}/env | Update .env File
[**UpdateLoadBalancing**](DefaultApi.md#UpdateLoadBalancing) | **Put** /servers/{serverId}/sites/{siteId}/balancing | Update Load Balancing
[**UpdateNginxConfiguration**](DefaultApi.md#UpdateNginxConfiguration) | **Put** /servers/{serverId}/sites/{siteId}/nginx | Update Nginx Configuration
[**UpdateNginxTemplate**](DefaultApi.md#UpdateNginxTemplate) | **Put** /servers/{serverId}/nginx/templates/{templateId} | Update Nginx Template
[**UpdateRecipe**](DefaultApi.md#UpdateRecipe) | **Put** /recipes/{recipeId} | Update Recipe
[**UpdateRepository**](DefaultApi.md#UpdateRepository) | **Put** /servers/{serverId}/sites/{siteId}/git | Update Repository
[**UpdateServer**](DefaultApi.md#UpdateServer) | **Put** /servers/{id} | Update Server
[**UpdateSite**](DefaultApi.md#UpdateSite) | **Put** /servers/{serverId}/sites/{siteId} | Update Site
[**UpdateUser**](DefaultApi.md#UpdateUser) | **Put** /servers/{serverId}/database-users/{userId} | Update User
[**UpgradePHPPatchVersion**](DefaultApi.md#UpgradePHPPatchVersion) | **Post** /servers/{serverId}/php/update | Upgrade PHP Patch Version

# **ActivateCertificate**
> ActivateCertificate(ctx, serverId, siteId, id)
Activate Certificate

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **int32**| The ID of the server. | 
  **siteId** | **int32**| The ID of the site. | 
  **id** | **int32**| The ID of the resource. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddSiteAliases**
> Site AddSiteAliases(ctx, body, serverId, siteId)
Add Site Aliases

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SiteIdAliasesBody**](SiteIdAliasesBody.md)|  | 
  **serverId** | **int32**| The ID of the server. | 
  **siteId** | **int32**| The ID of the site. | 

### Return type

[**Site**](Site.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChangeSitePHPVersion**
> ChangeSitePHPVersion(ctx, body, serverId, siteId)
Change Site PHP Version

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SiteIdPhpBody**](SiteIdPhpBody.md)|  | 
  **serverId** | **int32**| The ID of the server. | 
  **siteId** | **int32**| The ID of the site. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CloningAnExistingCertificate**
> Certificate CloningAnExistingCertificate(ctx, body, serverId, siteId)
Cloning An Existing Certificate

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SiteIdCertificatesBody**](SiteIdCertificatesBody.md)|  | 
  **serverId** | **int32**| The ID of the server. | 
  **siteId** | **int32**| The ID of the site. | 

### Return type

[**Certificate**](Certificate.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Create**
> Create(ctx, serverId, siteId)
Create

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **int32**| The ID of the server. | 
  **siteId** | **int32**| The ID of the site. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateBackupConfiguration**
> Backup CreateBackupConfiguration(ctx, body, serverId)
Create Backup Configuration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServerIdBackupconfigsBody**](ServerIdBackupconfigsBody.md)|  | 
  **serverId** | **int32**| The ID of the server. | 

### Return type

[**Backup**](Backup.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateDaemon**
> Daemon CreateDaemon(ctx, body, serverId)
Create Daemon

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServerIdDaemonsBody**](ServerIdDaemonsBody.md)|  | 
  **serverId** | **int32**| The ID of the server. | 

### Return type

[**Daemon**](Daemon.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateDatabase**
> Database CreateDatabase(ctx, body, serverId)
Create Database

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServerIdDatabasesBody**](ServerIdDatabasesBody.md)|  | 
  **serverId** | **int32**| The ID of the server. | 

### Return type

[**Database**](Database.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateDeployKey**
> CreateDeployKey(ctx, serverId, siteId)
Create Deploy Key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **int32**| The ID of the server. | 
  **siteId** | **int32**| The ID of the site. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateFirewallRule**
> Rule CreateFirewallRule(ctx, body, serverId)
Create Firewall Rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServerIdFirewallrulesBody**](ServerIdFirewallrulesBody.md)|  | 
  **serverId** | **int32**| The ID of the server. | 

### Return type

[**Rule**](Rule.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateJob**
> Job CreateJob(ctx, body, serverId)
Create Job

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServerIdJobsBody**](ServerIdJobsBody.md)|  | 
  **serverId** | **int32**| The ID of the server. | 

### Return type

[**Job**](Job.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateKey**
> Key CreateKey(ctx, body, serverId)
Create Key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServerIdKeysBody**](ServerIdKeysBody.md)|  | 
  **serverId** | **int32**| The ID of the server. | 

### Return type

[**Key**](Key.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMonitor**
> Monitor CreateMonitor(ctx, body, serverId)
Create Monitor

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServerIdMonitorsBody**](ServerIdMonitorsBody.md)|  | 
  **serverId** | **int32**| The ID of the server. | 

### Return type

[**Monitor**](Monitor.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRecipe**
> Recipe CreateRecipe(ctx, body)
Create Recipe

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RecipesBody**](RecipesBody.md)|  | 

### Return type

[**Recipe**](Recipe.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRedirectRule**
> RedirectRule CreateRedirectRule(ctx, body, serverId, siteId)
Create Redirect Rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SiteIdRedirectrulesBody**](SiteIdRedirectrulesBody.md)|  | 
  **serverId** | **int32**| The ID of the server. | 
  **siteId** | **int32**| The ID of the site. | 

### Return type

[**RedirectRule**](RedirectRule.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateSecurityRule**
> SecurityRule CreateSecurityRule(ctx, body, serverId, siteId)
Create Security Rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SiteIdSecurityrulesBody**](SiteIdSecurityrulesBody.md)|  | 
  **serverId** | **int32**| The ID of the server. | 
  **siteId** | **int32**| The ID of the site. | 

### Return type

[**SecurityRule**](SecurityRule.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateServer**
> Server CreateServer(ctx, body)
Create Server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServersBody**](ServersBody.md)|  | 

### Return type

[**Server**](Server.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateSite**
> Site CreateSite(ctx, body, serverId)
Create Site

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServerIdSitesBody**](ServerIdSitesBody.md)|  | 
  **serverId** | **int32**| The ID of the server. | 

### Return type

[**Site**](Site.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateTemplate**
> Template CreateTemplate(ctx, body, serverId)
Create Template

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**NginxTemplatesBody**](NginxTemplatesBody.md)|  | 
  **serverId** | **int32**| The ID of the server. | 

### Return type

[**Template**](Template.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateUser**
> User CreateUser(ctx, body, serverId)
Create User

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServerIdDatabaseusersBody**](ServerIdDatabaseusersBody.md)|  | 
  **serverId** | **int32**| The ID of the server. | 

### Return type

[**User**](User.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateWorker**
> Worker CreateWorker(ctx, body, serverId, siteId)
Create Worker

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SiteIdWorkersBody**](SiteIdWorkersBody.md)|  | 
  **serverId** | **int32**| The ID of the server. | 
  **siteId** | **int32**| The ID of the site. | 

### Return type

[**Worker**](Worker.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteBackup**
> DeleteBackup(ctx, serverId, backupConfigurationId, backupId)
Delete Backup

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **int32**| The ID of the server. | 
  **backupConfigurationId** | **int32**| The ID of the backup configuration. | 
  **backupId** | **int32**| The ID of the backup. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteBackupConfiguration**
> DeleteBackupConfiguration(ctx, serverId, backupConfigurationId)
Delete Backup Configuration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **int32**| The ID of the server. | 
  **backupConfigurationId** | **int32**| The ID of the backup configuration. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCertificate**
> DeleteCertificate(ctx, serverId, siteId, id)
Delete Certificate

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **int32**| The ID of the server. | 
  **siteId** | **int32**| The ID of the site. | 
  **id** | **int32**| The ID of the resource. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDaemon**
> DeleteDaemon(ctx, serverId, daemonId)
Delete Daemon

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **int32**| The ID of the server. | 
  **daemonId** | **int32**| The ID of the daemon. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDatabase**
> DeleteDatabase(ctx, serverId, databaseId)
Delete Database

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **int32**| The ID of the server. | 
  **databaseId** | **int32**| The ID of the database. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDeployKey**
> DeleteDeployKey(ctx, serverId, siteId)
Delete Deploy Key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **int32**| The ID of the server. | 
  **siteId** | **int32**| The ID of the site. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteFirewallRule**
> DeleteFirewallRule(ctx, serverId, ruleId)
Delete Firewall Rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **int32**| The ID of the server. | 
  **ruleId** | **int32**| The ID of the rule. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteJob**
> DeleteJob(ctx, serverId, jobId)
Delete Job

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **int32**| The ID of the server. | 
  **jobId** | **int32**| The ID of the job. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteKey**
> DeleteKey(ctx, serverId, keyId)
Delete Key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **int32**| The ID of the server. | 
  **keyId** | **int32**| The ID of the key. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMonitor**
> DeleteMonitor(ctx, serverId, monitorId)
Delete Monitor

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **int32**| The ID of the server. | 
  **monitorId** | **int32**| The ID of the monitor. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteNginxTemplate**
> DeleteNginxTemplate(ctx, serverId, templateId)
Delete Nginx Template

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **int32**| The ID of the server. | 
  **templateId** | **int32**| The ID of the template. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRecipe**
> DeleteRecipe(ctx, recipeId)
Delete Recipe

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **recipeId** | **int32**| The ID of the recipe. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRedirectRule**
> DeleteRedirectRule(ctx, serverId, siteId, id)
Delete Redirect Rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **int32**| The ID of the server. | 
  **siteId** | **int32**| The ID of the site. | 
  **id** | **int32**| The ID of the resource. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSecurityRule**
> DeleteSecurityRule(ctx, serverId, siteId, id)
Delete Security Rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **int32**| The ID of the server. | 
  **siteId** | **int32**| The ID of the site. | 
  **id** | **int32**| The ID of the resource. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteServer**
> DeleteServer(ctx, id)
Delete Server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the resource. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSite**
> DeleteSite(ctx, serverId, siteId)
Delete Site

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **int32**| The ID of the server. | 
  **siteId** | **int32**| The ID of the site. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteUser**
> DeleteUser(ctx, serverId, userId)
Delete User

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **int32**| The ID of the server. | 
  **userId** | **int32**| The ID of the user. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteWebhook**
> DeleteWebhook(ctx, serverId, siteId, id)
Delete Webhook

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **int32**| The ID of the server. | 
  **siteId** | **int32**| The ID of the site. | 
  **id** | **int32**| The ID of the resource. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteWorker**
> DeleteWorker(ctx, serverId, siteId, id)
Delete Worker

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **int32**| The ID of the server. | 
  **siteId** | **int32**| The ID of the site. | 
  **id** | **int32**| The ID of the resource. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeployNow**
> DeployNow(ctx, serverId, siteId)
Deploy Now

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **int32**| The ID of the server. | 
  **siteId** | **int32**| The ID of the site. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DisableOPCache**
> DisableOPCache(ctx, serverId)
Disable OPCache

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **int32**| The ID of the server. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DisableQuickDeployment**
> DisableQuickDeployment(ctx, serverId, siteId)
Disable Quick Deployment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **int32**| The ID of the server. | 
  **siteId** | **int32**| The ID of the site. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnableOPCache**
> EnableOPCache(ctx, serverId)
Enable OPCache

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **int32**| The ID of the server. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnableQuickDeployment**
> EnableQuickDeployment(ctx, serverId, siteId)
Enable Quick Deployment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **int32**| The ID of the server. | 
  **siteId** | **int32**| The ID of the site. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExecuteCommand**
> ExecuteCommand(ctx, body, serverId, siteId)
Execute Command

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SiteIdCommandsBody**](SiteIdCommandsBody.md)|  | 
  **serverId** | **int32**| The ID of the server. | 
  **siteId** | **int32**| The ID of the site. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBackupConfiguration**
> Backup GetBackupConfiguration(ctx, serverId, backupConfigurationId)
Get Backup Configuration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **int32**| The ID of the server. | 
  **backupConfigurationId** | **int32**| The ID of the backup configuration. | 

### Return type

[**Backup**](Backup.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCertificate**
> Certificate GetCertificate(ctx, serverId, siteId, id)
Get Certificate

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **int32**| The ID of the server. | 
  **siteId** | **int32**| The ID of the site. | 
  **id** | **int32**| The ID of the resource. | 

### Return type

[**Certificate**](Certificate.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCommand**
> Command GetCommand(ctx, serverId, siteId, commandId)
Get Command

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **int32**| The ID of the server. | 
  **siteId** | **int32**| The ID of the site. | 
  **commandId** | **int32**| The ID of the command. | 

### Return type

[**Command**](Command.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCredentials**
> []Credential GetCredentials(ctx, )
Get Credentials

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Credential**](array.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDaemon**
> Daemon GetDaemon(ctx, serverId, daemonId)
Get Daemon

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **int32**| The ID of the server. | 
  **daemonId** | **int32**| The ID of the daemon. | 

### Return type

[**Daemon**](Daemon.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDatabase**
> Database GetDatabase(ctx, serverId, databaseId)
Get Database

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **int32**| The ID of the server. | 
  **databaseId** | **int32**| The ID of the database. | 

### Return type

[**Database**](Database.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDeployment**
> Deployment GetDeployment(ctx, serverId, siteId, deploymentId)
Get Deployment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **int32**| The ID of the server. | 
  **siteId** | **int32**| The ID of the site. | 
  **deploymentId** | **int32**| The ID of the deployment. | 

### Return type

[**Deployment**](Deployment.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDeploymentLog**
> GetDeploymentLog(ctx, serverId, siteId)
Get Deployment Log

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **int32**| The ID of the server. | 
  **siteId** | **int32**| The ID of the site. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDeploymentOutput**
> GetDeploymentOutput(ctx, serverId, siteId, deploymentId)
Get Deployment Output

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **int32**| The ID of the server. | 
  **siteId** | **int32**| The ID of the site. | 
  **deploymentId** | **int32**| The ID of the deployment. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDeploymentScript**
> GetDeploymentScript(ctx, serverId, siteId)
Get Deployment Script

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **int32**| The ID of the server. | 
  **siteId** | **int32**| The ID of the site. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEnvFile**
> GetEnvFile(ctx, serverId, siteId)
Get .env File

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **int32**| The ID of the server. | 
  **siteId** | **int32**| The ID of the site. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFirewallRule**
> Rule GetFirewallRule(ctx, serverId, ruleId)
Get Firewall Rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **int32**| The ID of the server. | 
  **ruleId** | **int32**| The ID of the rule. | 

### Return type

[**Rule**](Rule.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetJob**
> Job GetJob(ctx, serverId, jobId)
Get Job

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **int32**| The ID of the server. | 
  **jobId** | **int32**| The ID of the job. | 

### Return type

[**Job**](Job.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetKey**
> Key GetKey(ctx, serverId, keyId)
Get Key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **int32**| The ID of the server. | 
  **keyId** | **int32**| The ID of the key. | 

### Return type

[**Key**](Key.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLog**
> GetLog(ctx, serverId)
Get Log

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **int32**| The ID of the server. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMonitor**
> GetMonitor(ctx, serverId, monitorId)
Get Monitor

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **int32**| The ID of the server. | 
  **monitorId** | **int32**| The ID of the monitor. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNginxConfiguration**
> GetNginxConfiguration(ctx, serverId, siteId)
Get Nginx Configuration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **int32**| The ID of the server. | 
  **siteId** | **int32**| The ID of the site. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNginxTemplate**
> Template GetNginxTemplate(ctx, serverId, templateId)
Get Nginx Template

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **int32**| The ID of the server. | 
  **templateId** | **int32**| The ID of the template. | 

### Return type

[**Template**](Template.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRecentEvents**
> []InlineResponse200 GetRecentEvents(ctx, body)
Get Recent Events

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServersEventsBody**](ServersEventsBody.md)|  | 

### Return type

[**[]InlineResponse200**](inline_response_200.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRecipe**
> Recipe GetRecipe(ctx, recipeId)
Get Recipe

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **recipeId** | **int32**| The ID of the recipe. | 

### Return type

[**Recipe**](Recipe.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRedirectRule**
> RedirectRule GetRedirectRule(ctx, serverId, siteId, id)
Get Redirect Rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **int32**| The ID of the server. | 
  **siteId** | **int32**| The ID of the site. | 
  **id** | **int32**| The ID of the resource. | 

### Return type

[**RedirectRule**](RedirectRule.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRegions**
> Regions GetRegions(ctx, )
Get Regions

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**Regions**](Regions.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSecurityRule**
> SecurityRule GetSecurityRule(ctx, serverId, siteId, id)
Get Security Rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **int32**| The ID of the server. | 
  **siteId** | **int32**| The ID of the site. | 
  **id** | **int32**| The ID of the resource. | 

### Return type

[**SecurityRule**](SecurityRule.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServer**
> Server GetServer(ctx, id)
Get Server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the resource. | 

### Return type

[**Server**](Server.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSigningRequest**
> GetSigningRequest(ctx, serverId, siteId, id)
Get Signing Request

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **int32**| The ID of the server. | 
  **siteId** | **int32**| The ID of the site. | 
  **id** | **int32**| The ID of the resource. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSite**
> Site GetSite(ctx, serverId, siteId)
Get Site

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **int32**| The ID of the server. | 
  **siteId** | **int32**| The ID of the site. | 

### Return type

[**Site**](Site.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUser**
> User GetUser(ctx, serverId, userId)
Get User

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **int32**| The ID of the server. | 
  **userId** | **int32**| The ID of the user. | 

### Return type

[**User**](User.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWebhook**
> Webhook GetWebhook(ctx, serverId, siteId, id)
Get Webhook

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **int32**| The ID of the server. | 
  **siteId** | **int32**| The ID of the site. | 
  **id** | **int32**| The ID of the resource. | 

### Return type

[**Webhook**](Webhook.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWebhooks**
> []Webhook GetWebhooks(ctx, serverId, siteId)
Get Webhooks

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **int32**| The ID of the server. | 
  **siteId** | **int32**| The ID of the site. | 

### Return type

[**[]Webhook**](array.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWorker**
> Worker GetWorker(ctx, serverId, siteId, id)
Get Worker

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **int32**| The ID of the server. | 
  **siteId** | **int32**| The ID of the site. | 
  **id** | **int32**| The ID of the resource. | 

### Return type

[**Worker**](Worker.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InstallBlackfire**
> InstallBlackfire(ctx, body, id)
Install Blackfire

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BlackfireInstallBody**](BlackfireInstallBody.md)|  | 
  **id** | **int32**| The ID of the resource. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InstallCertificate**
> InstallCertificate(ctx, body, serverId, siteId, id)
Install Certificate

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IdInstallBody**](IdInstallBody.md)|  | 
  **serverId** | **int32**| The ID of the server. | 
  **siteId** | **int32**| The ID of the site. | 
  **id** | **int32**| The ID of the resource. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InstallNewGitProject**
> InstallNewGitProject(ctx, body, serverId, siteId)
Install New Git Project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SiteIdGitBody1**](SiteIdGitBody1.md)|  | 
  **serverId** | **int32**| The ID of the server. | 
  **siteId** | **int32**| The ID of the site. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InstallPHPVersion**
> InstallPHPVersion(ctx, serverId)
Install PHP Version

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **int32**| The ID of the server. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InstallPapertrail**
> InstallPapertrail(ctx, body, id)
Install Papertrail

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PapertrailInstallBody**](PapertrailInstallBody.md)|  | 
  **id** | **int32**| The ID of the resource. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InstallPhpMyAdmin**
> InstallPhpMyAdmin(ctx, body, serverId, siteId)
Install phpMyAdmin

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SiteIdPhpmyadminBody**](SiteIdPhpmyadminBody.md)|  | 
  **serverId** | **int32**| The ID of the server. | 
  **siteId** | **int32**| The ID of the site. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InstallWordPress**
> InstallWordPress(ctx, body, serverId, siteId)
Install WordPress

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SiteIdWordpressBody**](SiteIdWordpressBody.md)|  | 
  **serverId** | **int32**| The ID of the server. | 
  **siteId** | **int32**| The ID of the site. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListBackupConfigurations**
> ListBackupConfigurations(ctx, serverId)
List Backup Configurations

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **int32**| The ID of the server. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCertificates**
> []Certificate ListCertificates(ctx, serverId, siteId)
List Certificates

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **int32**| The ID of the server. | 
  **siteId** | **int32**| The ID of the site. | 

### Return type

[**[]Certificate**](array.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCommandHistory**
> []Command ListCommandHistory(ctx, serverId, siteId)
List Command History

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **int32**| The ID of the server. | 
  **siteId** | **int32**| The ID of the site. | 

### Return type

[**[]Command**](array.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDaemons**
> []Daemon ListDaemons(ctx, serverId)
List Daemons

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **int32**| The ID of the server. | 

### Return type

[**[]Daemon**](array.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDatabases**
> []Database ListDatabases(ctx, serverId)
List Databases

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **int32**| The ID of the server. | 

### Return type

[**[]Database**](array.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDeployments**
> []Deployment ListDeployments(ctx, serverId, siteId)
List Deployments

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **int32**| The ID of the server. | 
  **siteId** | **int32**| The ID of the site. | 

### Return type

[**[]Deployment**](array.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListFirewallRules**
> []Rule ListFirewallRules(ctx, serverId)
List Firewall Rules

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **int32**| The ID of the server. | 

### Return type

[**[]Rule**](array.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListJobs**
> []Job ListJobs(ctx, serverId)
List Jobs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **int32**| The ID of the server. | 

### Return type

[**[]Job**](array.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListKeys**
> []Key ListKeys(ctx, serverId)
List Keys

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **int32**| The ID of the server. | 

### Return type

[**[]Key**](array.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListMonitors**
> ListMonitors(ctx, serverId)
List Monitors

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **int32**| The ID of the server. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListNginxTemplates**
> []Template ListNginxTemplates(ctx, serverId)
List Nginx Templates

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **int32**| The ID of the server. | 

### Return type

[**[]Template**](array.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPHPVersions**
> []InlineResponse2001 ListPHPVersions(ctx, serverId)
List PHP Versions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **int32**| The ID of the server. | 

### Return type

[**[]InlineResponse2001**](inline_response_200_1.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListRecipes**
> []Recipe ListRecipes(ctx, )
List Recipes

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Recipe**](array.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListRedirectRules**
> []RedirectRule ListRedirectRules(ctx, serverId, siteId)
List Redirect Rules

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **int32**| The ID of the server. | 
  **siteId** | **int32**| The ID of the site. | 

### Return type

[**[]RedirectRule**](array.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSecurityRules**
> []SecurityRule ListSecurityRules(ctx, serverId, siteId)
List Security Rules

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **int32**| The ID of the server. | 
  **siteId** | **int32**| The ID of the site. | 

### Return type

[**[]SecurityRule**](array.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListServers**
> []Server ListServers(ctx, )
List Servers

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Server**](array.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSites**
> []Site ListSites(ctx, serverId)
List Sites

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **int32**| The ID of the server. | 

### Return type

[**[]Site**](array.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListUsers**
> []User ListUsers(ctx, serverId)
List Users

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **int32**| The ID of the server. | 

### Return type

[**[]User**](array.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListWorkers**
> []Worker ListWorkers(ctx, serverId, siteId)
List Workers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **int32**| The ID of the server. | 
  **siteId** | **int32**| The ID of the site. | 

### Return type

[**[]Worker**](array.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LoadBalancing**
> []Node LoadBalancing(ctx, serverId, siteId)
Load Balancing

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **int32**| The ID of the server. | 
  **siteId** | **int32**| The ID of the site. | 

### Return type

[**[]Node**](array.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ObtainALetsEncryptCertificate**
> Certificate ObtainALetsEncryptCertificate(ctx, serverId, siteId)
Obtain A LetsEncrypt Certificate

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **int32**| The ID of the server. | 
  **siteId** | **int32**| The ID of the site. | 

### Return type

[**Certificate**](Certificate.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReactivateRevokedServer**
> ReactivateRevokedServer(ctx, id)
Reactivate revoked server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the resource. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RebootMySQL**
> RebootMySQL(ctx, id)
Reboot MySQL

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the resource. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RebootNginx**
> RebootNginx(ctx, id)
Reboot Nginx

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the resource. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RebootPHP**
> RebootPHP(ctx, body, id)
Reboot PHP

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PhpRebootBody**](PhpRebootBody.md)|  | 
  **id** | **int32**| The ID of the resource. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RebootPostgres**
> RebootPostgres(ctx, id)
Reboot Postgres

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the resource. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RebootServer**
> RebootServer(ctx, id)
Reboot Server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the resource. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReconnectRevokedServer**
> ReconnectRevokedServer(ctx, id)
Reconnect revoked server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the resource. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveBlackfire**
> RemoveBlackfire(ctx, id)
Remove Blackfire

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the resource. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemovePapertrail**
> RemovePapertrail(ctx, id)
Remove Papertrail

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the resource. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveProject**
> RemoveProject(ctx, serverId, siteId)
Remove Project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **int32**| The ID of the server. | 
  **siteId** | **int32**| The ID of the site. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResetDeploymentStatus**
> ResetDeploymentStatus(ctx, serverId, siteId)
Reset Deployment Status

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **int32**| The ID of the server. | 
  **siteId** | **int32**| The ID of the site. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestartDaemon**
> RestartDaemon(ctx, serverId, daemonId)
Restart Daemon

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **int32**| The ID of the server. | 
  **daemonId** | **int32**| The ID of the daemon. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestartService**
> RestartService(ctx, id)
Restart Service

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the resource. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestartWorker**
> RestartWorker(ctx, serverId, siteId, id)
Restart Worker

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **int32**| The ID of the server. | 
  **siteId** | **int32**| The ID of the site. | 
  **id** | **int32**| The ID of the resource. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestoreBackup**
> RestoreBackup(ctx, body, serverId, backupConfigurationId, backupId)
Restore Backup

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BackupsBackupIdBody**](BackupsBackupIdBody.md)|  | 
  **serverId** | **int32**| The ID of the server. | 
  **backupConfigurationId** | **int32**| The ID of the backup configuration. | 
  **backupId** | **int32**| The ID of the backup. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RevokeForgeAccessToServer**
> RevokeForgeAccessToServer(ctx, id)
Revoke Forge access to server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the resource. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RunBackupConfiguration**
> RunBackupConfiguration(ctx, serverId, backupConfigurationId)
Run Backup Configuration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **int32**| The ID of the server. | 
  **backupConfigurationId** | **int32**| The ID of the backup configuration. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RunRecipe**
> RunRecipe(ctx, body, recipeId)
Run Recipe

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RecipeIdRunBody**](RecipeIdRunBody.md)|  | 
  **recipeId** | **int32**| The ID of the recipe. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Show**
> User Show(ctx, )
show

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**User**](User.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SiteLog**
> SiteLog(ctx, serverId, siteId)
Site Log

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **int32**| The ID of the server. | 
  **siteId** | **int32**| The ID of the site. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StartService**
> StartService(ctx, id)
Start Service

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the resource. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StopMySQL**
> StopMySQL(ctx, id)
Stop MySQL

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the resource. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StopNginx**
> StopNginx(ctx, id)
Stop Nginx

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the resource. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StopPostgres**
> StopPostgres(ctx, id)
Stop Postgres

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the resource. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StopService**
> StopService(ctx, id)
Stop Service

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the resource. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SyncDatabase**
> SyncDatabase(ctx, serverId)
Sync Database

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **int32**| The ID of the server. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TestNginx**
> TestNginx(ctx, id)
Test Nginx

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The ID of the resource. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UninstallPhpMyAdmin**
> UninstallPhpMyAdmin(ctx, serverId, siteId)
Uninstall phpMyAdmin

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **int32**| The ID of the server. | 
  **siteId** | **int32**| The ID of the site. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UninstallWordPress**
> UninstallWordPress(ctx, serverId, siteId)
Uninstall WordPress

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **int32**| The ID of the server. | 
  **siteId** | **int32**| The ID of the site. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateBackupConfiguration**
> Backup UpdateBackupConfiguration(ctx, body, serverId, backupConfigurationId)
Update Backup Configuration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BackupconfigsBackupConfigurationIdBody**](BackupconfigsBackupConfigurationIdBody.md)|  | 
  **serverId** | **int32**| The ID of the server. | 
  **backupConfigurationId** | **int32**| The ID of the backup configuration. | 

### Return type

[**Backup**](Backup.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDatabasePassword**
> UpdateDatabasePassword(ctx, body, serverId)
Update Database Password

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServerIdDatabasepasswordBody**](ServerIdDatabasepasswordBody.md)|  | 
  **serverId** | **int32**| The ID of the server. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDeploymentScript**
> UpdateDeploymentScript(ctx, body, serverId, siteId)
Update Deployment Script

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DeploymentScriptBody**](DeploymentScriptBody.md)|  | 
  **serverId** | **int32**| The ID of the server. | 
  **siteId** | **int32**| The ID of the site. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateEnvFile**
> UpdateEnvFile(ctx, body, serverId, siteId)
Update .env File

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SiteIdEnvBody**](SiteIdEnvBody.md)|  | 
  **serverId** | **int32**| The ID of the server. | 
  **siteId** | **int32**| The ID of the site. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateLoadBalancing**
> UpdateLoadBalancing(ctx, body, serverId, siteId)
Update Load Balancing

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SiteIdBalancingBody**](SiteIdBalancingBody.md)|  | 
  **serverId** | **int32**| The ID of the server. | 
  **siteId** | **int32**| The ID of the site. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNginxConfiguration**
> UpdateNginxConfiguration(ctx, body, serverId, siteId)
Update Nginx Configuration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SiteIdNginxBody**](SiteIdNginxBody.md)|  | 
  **serverId** | **int32**| The ID of the server. | 
  **siteId** | **int32**| The ID of the site. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNginxTemplate**
> Template UpdateNginxTemplate(ctx, serverId, templateId)
Update Nginx Template

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **int32**| The ID of the server. | 
  **templateId** | **int32**| The ID of the template. | 

### Return type

[**Template**](Template.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRecipe**
> Recipe UpdateRecipe(ctx, body, recipeId)
Update Recipe

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RecipesRecipeIdBody**](RecipesRecipeIdBody.md)|  | 
  **recipeId** | **int32**| The ID of the recipe. | 

### Return type

[**Recipe**](Recipe.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRepository**
> UpdateRepository(ctx, body, serverId, siteId)
Update Repository

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SiteIdGitBody**](SiteIdGitBody.md)|  | 
  **serverId** | **int32**| The ID of the server. | 
  **siteId** | **int32**| The ID of the site. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateServer**
> Server UpdateServer(ctx, body, id)
Update Server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServersIdBody**](ServersIdBody.md)|  | 
  **id** | **int32**| The ID of the resource. | 

### Return type

[**Server**](Server.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSite**
> Site UpdateSite(ctx, body, serverId, siteId)
Update Site

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SitesSiteIdBody**](SitesSiteIdBody.md)|  | 
  **serverId** | **int32**| The ID of the server. | 
  **siteId** | **int32**| The ID of the site. | 

### Return type

[**Site**](Site.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateUser**
> User UpdateUser(ctx, body, serverId, userId)
Update User

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DatabaseusersUserIdBody**](DatabaseusersUserIdBody.md)|  | 
  **serverId** | **int32**| The ID of the server. | 
  **userId** | **int32**| The ID of the user. | 

### Return type

[**User**](User.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpgradePHPPatchVersion**
> UpgradePHPPatchVersion(ctx, serverId)
Upgrade PHP Patch Version

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **int32**| The ID of the server. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

