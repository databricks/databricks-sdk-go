# NEXT CHANGELOG

## Release v0.86.0

### New Features and Improvements

### Bug Fixes

### Documentation

### Internal Changes

### API Changes
* Add `UpdateNotifications` method for [w.ServingEndpoints](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpointsAPI) workspace-level service.
* Add `Parameters` field for [dashboards.GenieQueryAttachment](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieQueryAttachment).
* Add `DatabaseInstanceName` field for [database.CreateDatabaseInstanceRoleRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/database#CreateDatabaseInstanceRoleRequest).
* Add `EffectiveAttributes` and `InstanceName` fields for [database.DatabaseInstanceRole](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/database#DatabaseInstanceRole).
* Add `KeyRegion` field for [provisioning.CreateAwsKeyInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#CreateAwsKeyInfo).
* Add `ExternalId` field for [provisioning.CreateCredentialStsRole](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#CreateCredentialStsRole).
* Add `GcpServiceAccount` field for [provisioning.CreateGcpKeyInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#CreateGcpKeyInfo).
* Add `AzureKeyInfo` field for [provisioning.CustomerManagedKey](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#CustomerManagedKey).
* Add `OptOutExternalCustomerTosWorkflow`, `TosAcceptedByEmail`, `TosAcceptedByFullName` and `TosAcceptedTimestamp` fields for [provisioning.ExternalCustomerInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#ExternalCustomerInfo).
* Add `GcpServiceAccount` field for [provisioning.GcpKeyInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#GcpKeyInfo).
* Add `AwsNetworkInfo` field for [provisioning.Network](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#Network).
* [Breaking] Add `CustomerFacingPrivateAccessSettings` field for [provisioning.ReplacePrivateAccessSettingsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#ReplacePrivateAccessSettingsRequest).
* Add `RoleArn` field for [provisioning.StorageConfiguration](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#StorageConfiguration).
* [Breaking] Add `CustomerFacingWorkspace` field for [provisioning.UpdateWorkspaceRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#UpdateWorkspaceRequest).
* Add `UpdateMask` field for [provisioning.UpdateWorkspaceRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#UpdateWorkspaceRequest).
* Add `ComputeMode`, `Network`, `NetworkConnectivityConfigId` and `StorageMode` fields for [provisioning.Workspace](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#Workspace).
* Add `StreamNative` enum value for [catalog.SystemType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SystemType).
* Add `ExceededMaxTokenLengthException` enum value for [dashboards.MessageErrorType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#MessageErrorType).
* Add `UnknownAccessLevel` and `Any` enum values for [provisioning.PrivateAccessLevel](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#PrivateAccessLevel).
* [Breaking] Change `Create` method for [a.Credentials](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#CredentialsAPI) account-level service with new required argument order.
* [Breaking] Change `Delete` method for [a.Credentials](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#CredentialsAPI) account-level service to start returning [provisioning.Credential](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#Credential).
* [Breaking] Change `Create` method for [a.EncryptionKeys](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#EncryptionKeysAPI) account-level service with new required argument order.
* [Breaking] Change `Delete` method for [a.EncryptionKeys](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#EncryptionKeysAPI) account-level service to start returning [provisioning.CustomerManagedKey](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#CustomerManagedKey).
* [Breaking] Change `Create` method for [a.Networks](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#NetworksAPI) account-level service with new required argument order.
* [Breaking] Change `Delete` method for [a.Networks](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#NetworksAPI) account-level service to start returning [provisioning.Network](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#Network).
* [Breaking] Change `Create` and `Replace` methods for [a.PrivateAccess](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#PrivateAccessAPI) account-level service with new required argument order.
* [Breaking] Change `Delete` and `Replace` methods for [a.PrivateAccess](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#PrivateAccessAPI) account-level service to start returning [provisioning.PrivateAccessSettings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#PrivateAccessSettings).
* [Breaking] Change `Create` method for [a.Storage](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#StorageAPI) account-level service with new required argument order.
* [Breaking] Change `Delete` method for [a.Storage](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#StorageAPI) account-level service to start returning [provisioning.StorageConfiguration](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#StorageConfiguration).
* [Breaking] Change `Create` method for [a.VpcEndpoints](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#VpcEndpointsAPI) account-level service with new required argument order.
* [Breaking] Change `Delete` method for [a.VpcEndpoints](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#VpcEndpointsAPI) account-level service to start returning [provisioning.VpcEndpoint](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#VpcEndpoint).
* [Breaking] Change `Create` and `Update` methods for [a.Workspaces](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#WorkspacesAPI) account-level service with new required argument order.
* [Breaking] Change `Delete` and `Update` methods for [a.Workspaces](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#WorkspacesAPI) account-level service to start returning [provisioning.Workspace](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#Workspace).
* [Breaking] Change `Name` field for [database.DatabaseInstanceRole](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/database#DatabaseInstanceRole) to be required.
* Change `Name` field for [database.DatabaseInstanceRole](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/database#DatabaseInstanceRole) to be required.
* Change `AwsCredentials` and `CredentialsName` fields for [provisioning.CreateCredentialRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#CreateCredentialRequest) to no longer be required.
* Change `UseCases` field for [provisioning.CreateCustomerManagedKeyRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#CreateCustomerManagedKeyRequest) to no longer be required.
* Change `NetworkName` field for [provisioning.CreateNetworkRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#CreateNetworkRequest) to no longer be required.
* Change `PrivateAccessSettingsName` and `Region` fields for [provisioning.CreatePrivateAccessSettingsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#CreatePrivateAccessSettingsRequest) to no longer be required.
* Change `RootBucketInfo` and `StorageConfigurationName` fields for [provisioning.CreateStorageConfigurationRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#CreateStorageConfigurationRequest) to no longer be required.
* Change `VpcEndpointName` field for [provisioning.CreateVpcEndpointRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#CreateVpcEndpointRequest) to no longer be required.
* Change `WorkspaceName` field for [provisioning.CreateWorkspaceRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#CreateWorkspaceRequest) to no longer be required.
* Change `DataplaneRelay` and `RestApi` fields for [provisioning.NetworkVpcEndpoints](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#NetworkVpcEndpoints) to no longer be required.
* [Breaking] Change `DataplaneRelay` and `RestApi` fields for [provisioning.NetworkVpcEndpoints](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#NetworkVpcEndpoints) to no longer be required.
* [Breaking] Change waiter for [WorkspacesAPI.Update](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#WorkspacesAPI.Update).
* [Breaking] Remove `AllowedVpcEndpointIds`, `PrivateAccessLevel`, `PrivateAccessSettingsName`, `PublicAccessEnabled` and `Region` fields for [provisioning.ReplacePrivateAccessSettingsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#ReplacePrivateAccessSettingsRequest).
* [Breaking] Remove `AwsRegion`, `CredentialsId`, `CustomTags`, `ManagedServicesCustomerManagedKeyId`, `NetworkConnectivityConfigId`, `NetworkId`, `PrivateAccessSettingsId`, `StorageConfigurationId` and `StorageCustomerManagedKeyId` fields for [provisioning.UpdateWorkspaceRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#UpdateWorkspaceRequest).