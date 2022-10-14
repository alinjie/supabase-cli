// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.11.0 DO NOT EDIT.
package api

const (
	ApiKeyHeaderScopes = "apiKeyHeader.Scopes"
	ApiKeyParamScopes  = "apiKeyParam.Scopes"
	BearerScopes       = "bearer.Scopes"
)

// Defines values for CreateProjectBodyPlan.
const (
	Free CreateProjectBodyPlan = "free"
	Pro  CreateProjectBodyPlan = "pro"
)

// Defines values for CreateProjectBodyRegion.
const (
	ApNortheast1 CreateProjectBodyRegion = "ap-northeast-1"
	ApNortheast2 CreateProjectBodyRegion = "ap-northeast-2"
	ApSouth1     CreateProjectBodyRegion = "ap-south-1"
	ApSoutheast1 CreateProjectBodyRegion = "ap-southeast-1"
	ApSoutheast2 CreateProjectBodyRegion = "ap-southeast-2"
	CaCentral1   CreateProjectBodyRegion = "ca-central-1"
	EuCentral1   CreateProjectBodyRegion = "eu-central-1"
	EuWest1      CreateProjectBodyRegion = "eu-west-1"
	EuWest2      CreateProjectBodyRegion = "eu-west-2"
	SaEast1      CreateProjectBodyRegion = "sa-east-1"
	UsEast1      CreateProjectBodyRegion = "us-east-1"
	UsWest1      CreateProjectBodyRegion = "us-west-1"
)

// Defines values for FunctionResponseStatus.
const (
	FunctionResponseStatusACTIVE    FunctionResponseStatus = "ACTIVE"
	FunctionResponseStatusREMOVED   FunctionResponseStatus = "REMOVED"
	FunctionResponseStatusTHROTTLED FunctionResponseStatus = "THROTTLED"
)

// Defines values for FunctionSlugResponseStatus.
const (
	FunctionSlugResponseStatusACTIVE    FunctionSlugResponseStatus = "ACTIVE"
	FunctionSlugResponseStatusREMOVED   FunctionSlugResponseStatus = "REMOVED"
	FunctionSlugResponseStatusTHROTTLED FunctionSlugResponseStatus = "THROTTLED"
)

// Defines values for UpdateCustomHostnameResponseStatus.
const (
	N1NotStarted           UpdateCustomHostnameResponseStatus = "1_not_started"
	N2Initiated            UpdateCustomHostnameResponseStatus = "2_initiated"
	N3ChallengeVerified    UpdateCustomHostnameResponseStatus = "3_challenge_verified"
	N4OriginSetupCompleted UpdateCustomHostnameResponseStatus = "4_origin_setup_completed"
	N5ServicesReconfigured UpdateCustomHostnameResponseStatus = "5_services_reconfigured"
)

// CreateFunctionBody defines model for CreateFunctionBody.
type CreateFunctionBody struct {
	Body      string `json:"body"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	VerifyJwt *bool  `json:"verify_jwt,omitempty"`
}

// CreateOrganizationBody defines model for CreateOrganizationBody.
type CreateOrganizationBody struct {
	Name string `json:"name"`
}

// CreateProjectBody defines model for CreateProjectBody.
type CreateProjectBody struct {
	DbPass         string                  `json:"db_pass"`
	KpsEnabled     *bool                   `json:"kps_enabled,omitempty"`
	Name           string                  `json:"name"`
	OrganizationId string                  `json:"organization_id"`
	Plan           CreateProjectBodyPlan   `json:"plan"`
	Region         CreateProjectBodyRegion `json:"region"`
}

// CreateProjectBodyPlan defines model for CreateProjectBody.Plan.
type CreateProjectBodyPlan string

// CreateProjectBodyRegion defines model for CreateProjectBody.Region.
type CreateProjectBodyRegion string

// CreateSecretBody defines model for CreateSecretBody.
type CreateSecretBody struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// FunctionResponse defines model for FunctionResponse.
type FunctionResponse struct {
	CreatedAt float32                `json:"created_at"`
	Id        string                 `json:"id"`
	Name      string                 `json:"name"`
	Slug      string                 `json:"slug"`
	Status    FunctionResponseStatus `json:"status"`
	UpdatedAt float32                `json:"updated_at"`
	VerifyJwt *bool                  `json:"verify_jwt,omitempty"`
	Version   float32                `json:"version"`
}

// FunctionResponseStatus defines model for FunctionResponse.Status.
type FunctionResponseStatus string

// FunctionSlugResponse defines model for FunctionSlugResponse.
type FunctionSlugResponse struct {
	Body      *string                    `json:"body,omitempty"`
	CreatedAt float32                    `json:"created_at"`
	Id        string                     `json:"id"`
	Name      string                     `json:"name"`
	Slug      string                     `json:"slug"`
	Status    FunctionSlugResponseStatus `json:"status"`
	UpdatedAt float32                    `json:"updated_at"`
	VerifyJwt *bool                      `json:"verify_jwt,omitempty"`
	Version   float32                    `json:"version"`
}

// FunctionSlugResponseStatus defines model for FunctionSlugResponse.Status.
type FunctionSlugResponseStatus string

// OrganizationResponse defines model for OrganizationResponse.
type OrganizationResponse struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

// ProjectResponse defines model for ProjectResponse.
type ProjectResponse struct {
	CreatedAt      string `json:"created_at"`
	Id             string `json:"id"`
	Name           string `json:"name"`
	OrganizationId string `json:"organization_id"`
	Region         string `json:"region"`
}

// SecretResponse defines model for SecretResponse.
type SecretResponse struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// TypescriptResponse defines model for TypescriptResponse.
type TypescriptResponse struct {
	Types string `json:"types"`
}

// UpdateCustomHostnameBody defines model for UpdateCustomHostnameBody.
type UpdateCustomHostnameBody struct {
	CustomHostname string `json:"custom_hostname"`
}

// UpdateCustomHostnameResponse defines model for UpdateCustomHostnameResponse.
type UpdateCustomHostnameResponse struct {
	CustomHostname string                             `json:"custom_hostname"`
	Data           map[string]interface{}             `json:"data"`
	Status         UpdateCustomHostnameResponseStatus `json:"status"`
}

// UpdateCustomHostnameResponseStatus defines model for UpdateCustomHostnameResponse.Status.
type UpdateCustomHostnameResponseStatus string

// UpdateFunctionBody defines model for UpdateFunctionBody.
type UpdateFunctionBody struct {
	Body      *string `json:"body,omitempty"`
	Name      *string `json:"name,omitempty"`
	VerifyJwt *bool   `json:"verify_jwt,omitempty"`
}

// UpdatePgsodiumConfigBody defines model for UpdatePgsodiumConfigBody.
type UpdatePgsodiumConfigBody struct {
	RootKey string `json:"root_key"`
}

// CreateOrganizationJSONBody defines parameters for CreateOrganization.
type CreateOrganizationJSONBody = CreateOrganizationBody

// CreateProjectJSONBody defines parameters for CreateProject.
type CreateProjectJSONBody = CreateProjectBody

// CreateCustomHostnameConfigJSONBody defines parameters for CreateCustomHostnameConfig.
type CreateCustomHostnameConfigJSONBody = UpdateCustomHostnameBody

// CreateFunctionJSONBody defines parameters for CreateFunction.
type CreateFunctionJSONBody = CreateFunctionBody

// GetFunctionParams defines parameters for GetFunction.
type GetFunctionParams struct {
	IncludeBody *bool `form:"include_body,omitempty" json:"include_body,omitempty"`
}

// UpdateFunctionJSONBody defines parameters for UpdateFunction.
type UpdateFunctionJSONBody = UpdateFunctionBody

// UpdateConfigJSONBody defines parameters for UpdateConfig.
type UpdateConfigJSONBody = UpdatePgsodiumConfigBody

// DeleteSecretsJSONBody defines parameters for DeleteSecrets.
type DeleteSecretsJSONBody = []string

// CreateSecretsJSONBody defines parameters for CreateSecrets.
type CreateSecretsJSONBody = []CreateSecretBody

// GetTypescriptTypesParams defines parameters for GetTypescriptTypes.
type GetTypescriptTypesParams struct {
	IncludedSchemas *string `form:"included_schemas,omitempty" json:"included_schemas,omitempty"`
}

// CreateOrganizationJSONRequestBody defines body for CreateOrganization for application/json ContentType.
type CreateOrganizationJSONRequestBody = CreateOrganizationJSONBody

// CreateProjectJSONRequestBody defines body for CreateProject for application/json ContentType.
type CreateProjectJSONRequestBody = CreateProjectJSONBody

// CreateCustomHostnameConfigJSONRequestBody defines body for CreateCustomHostnameConfig for application/json ContentType.
type CreateCustomHostnameConfigJSONRequestBody = CreateCustomHostnameConfigJSONBody

// CreateFunctionJSONRequestBody defines body for CreateFunction for application/json ContentType.
type CreateFunctionJSONRequestBody = CreateFunctionJSONBody

// UpdateFunctionJSONRequestBody defines body for UpdateFunction for application/json ContentType.
type UpdateFunctionJSONRequestBody = UpdateFunctionJSONBody

// UpdateConfigJSONRequestBody defines body for UpdateConfig for application/json ContentType.
type UpdateConfigJSONRequestBody = UpdateConfigJSONBody

// DeleteSecretsJSONRequestBody defines body for DeleteSecrets for application/json ContentType.
type DeleteSecretsJSONRequestBody = DeleteSecretsJSONBody

// CreateSecretsJSONRequestBody defines body for CreateSecrets for application/json ContentType.
type CreateSecretsJSONRequestBody = CreateSecretsJSONBody
