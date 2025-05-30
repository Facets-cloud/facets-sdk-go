// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/Facets-cloud/facets-sdk-go/facets/client/application_controller"
	"github.com/Facets-cloud/facets-sdk-go/facets/client/artifact_controller"
	"github.com/Facets-cloud/facets-sdk-go/facets/client/artifactory_controller"
	"github.com/Facets-cloud/facets-sdk-go/facets/client/aws_cluster_controller"
	"github.com/Facets-cloud/facets-sdk-go/facets/client/azure_cluster_controller"
	"github.com/Facets-cloud/facets-sdk-go/facets/client/build_controller"
	"github.com/Facets-cloud/facets-sdk-go/facets/client/callback_controller"
	"github.com/Facets-cloud/facets-sdk-go/facets/client/capillary_cloud_callback_controller"
	"github.com/Facets-cloud/facets-sdk-go/facets/client/common_cluster_controller"
	"github.com/Facets-cloud/facets-sdk-go/facets/client/deployment_controller"
	"github.com/Facets-cloud/facets-sdk-go/facets/client/gcp_cluster_controller"
	"github.com/Facets-cloud/facets-sdk-go/facets/client/meta_controller"
	"github.com/Facets-cloud/facets-sdk-go/facets/client/public_ap_is"
	"github.com/Facets-cloud/facets-sdk-go/facets/client/stack_controller"
	"github.com/Facets-cloud/facets-sdk-go/facets/client/ui_accounts_controller"
	"github.com/Facets-cloud/facets-sdk-go/facets/client/ui_alerts_controller"
	"github.com/Facets-cloud/facets-sdk-go/facets/client/ui_application_controller"
	"github.com/Facets-cloud/facets-sdk-go/facets/client/ui_artifact_ci_controller"
	"github.com/Facets-cloud/facets-sdk-go/facets/client/ui_artifact_hub_controller"
	"github.com/Facets-cloud/facets-sdk-go/facets/client/ui_artifact_routing_rule_controller"
	"github.com/Facets-cloud/facets-sdk-go/facets/client/ui_artifactory_controller"
	"github.com/Facets-cloud/facets-sdk-go/facets/client/ui_artifacts_controller"
	"github.com/Facets-cloud/facets-sdk-go/facets/client/ui_assistant_controller"
	"github.com/Facets-cloud/facets-sdk-go/facets/client/ui_audit_logs_controller"
	"github.com/Facets-cloud/facets-sdk-go/facets/client/ui_aws_cluster_controller"
	"github.com/Facets-cloud/facets-sdk-go/facets/client/ui_azure_cluster_controller"
	"github.com/Facets-cloud/facets-sdk-go/facets/client/ui_billing_controller"
	"github.com/Facets-cloud/facets-sdk-go/facets/client/ui_blueprint_designer_controller"
	"github.com/Facets-cloud/facets-sdk-go/facets/client/ui_chat_gpt_controller"
	"github.com/Facets-cloud/facets-sdk-go/facets/client/ui_ci_cd_controller"
	"github.com/Facets-cloud/facets-sdk-go/facets/client/ui_cloud_cost_explorer_controller"
	"github.com/Facets-cloud/facets-sdk-go/facets/client/ui_coder_controller"
	"github.com/Facets-cloud/facets-sdk-go/facets/client/ui_common_cluster_controller"
	"github.com/Facets-cloud/facets-sdk-go/facets/client/ui_custom_content_controller"
	"github.com/Facets-cloud/facets-sdk-go/facets/client/ui_custom_role_controller"
	"github.com/Facets-cloud/facets-sdk-go/facets/client/ui_delivery_pipeline_controller"
	"github.com/Facets-cloud/facets-sdk-go/facets/client/ui_deployment_controller"
	"github.com/Facets-cloud/facets-sdk-go/facets/client/ui_domain_mapping_controller"
	"github.com/Facets-cloud/facets-sdk-go/facets/client/ui_dropdowns_controller"
	"github.com/Facets-cloud/facets-sdk-go/facets/client/ui_gcp_cluster_controller"
	"github.com/Facets-cloud/facets-sdk-go/facets/client/ui_iac_repo_controller"
	"github.com/Facets-cloud/facets-sdk-go/facets/client/ui_intent_controller"
	"github.com/Facets-cloud/facets-sdk-go/facets/client/ui_k_8s_cluster_controller"
	"github.com/Facets-cloud/facets-sdk-go/facets/client/ui_kubernetes_explorer_controller"
	"github.com/Facets-cloud/facets-sdk-go/facets/client/ui_local_cluster_controller"
	"github.com/Facets-cloud/facets-sdk-go/facets/client/ui_maintenance_window_controller"
	"github.com/Facets-cloud/facets-sdk-go/facets/client/ui_meta_controller"
	"github.com/Facets-cloud/facets-sdk-go/facets/client/ui_no_auth_user_controller"
	"github.com/Facets-cloud/facets-sdk-go/facets/client/ui_notification_controller"
	"github.com/Facets-cloud/facets-sdk-go/facets/client/ui_o_auth_controller"
	"github.com/Facets-cloud/facets-sdk-go/facets/client/ui_one_time_webhook_controller"
	"github.com/Facets-cloud/facets-sdk-go/facets/client/ui_opa_controller"
	"github.com/Facets-cloud/facets-sdk-go/facets/client/ui_project_type_controller"
	"github.com/Facets-cloud/facets-sdk-go/facets/client/ui_promotion_workflow_controller"
	"github.com/Facets-cloud/facets-sdk-go/facets/client/ui_provided_resources_controller"
	"github.com/Facets-cloud/facets-sdk-go/facets/client/ui_queued_release_controller"
	"github.com/Facets-cloud/facets-sdk-go/facets/client/ui_release_stream_controller"
	"github.com/Facets-cloud/facets-sdk-go/facets/client/ui_resource_group_controller"
	"github.com/Facets-cloud/facets-sdk-go/facets/client/ui_resource_status_controller"
	"github.com/Facets-cloud/facets-sdk-go/facets/client/ui_settings_controller"
	"github.com/Facets-cloud/facets-sdk-go/facets/client/ui_stack_controller"
	"github.com/Facets-cloud/facets-sdk-go/facets/client/ui_team_controller"
	"github.com/Facets-cloud/facets-sdk-go/facets/client/ui_tf_module_controller"
	"github.com/Facets-cloud/facets-sdk-go/facets/client/ui_tf_output_controller"
	"github.com/Facets-cloud/facets-sdk-go/facets/client/ui_tf_version_controller"
	"github.com/Facets-cloud/facets-sdk-go/facets/client/ui_theme_file_controller"
	"github.com/Facets-cloud/facets-sdk-go/facets/client/ui_user_controller"
	"github.com/Facets-cloud/facets-sdk-go/facets/client/ui_user_group_controller"
	"github.com/Facets-cloud/facets-sdk-go/facets/client/ui_user_vcs_token_controller"
	"github.com/Facets-cloud/facets-sdk-go/facets/client/ui_versioning_controller"
	"github.com/Facets-cloud/facets-sdk-go/facets/client/ui_web_component_controller"
)

// Default facets HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "facetsdemo.console.facets.cloud"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http"}

// NewHTTPClient creates a new facets HTTP client.
func NewHTTPClient(formats strfmt.Registry) *Facets {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new facets HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *Facets {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new facets client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Facets {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(Facets)
	cli.Transport = transport
	cli.ApplicationController = application_controller.New(transport, formats)
	cli.ArtifactController = artifact_controller.New(transport, formats)
	cli.ArtifactoryController = artifactory_controller.New(transport, formats)
	cli.AwsClusterController = aws_cluster_controller.New(transport, formats)
	cli.AzureClusterController = azure_cluster_controller.New(transport, formats)
	cli.BuildController = build_controller.New(transport, formats)
	cli.CallbackController = callback_controller.New(transport, formats)
	cli.CapillaryCloudCallbackController = capillary_cloud_callback_controller.New(transport, formats)
	cli.CommonClusterController = common_cluster_controller.New(transport, formats)
	cli.DeploymentController = deployment_controller.New(transport, formats)
	cli.GcpClusterController = gcp_cluster_controller.New(transport, formats)
	cli.MetaController = meta_controller.New(transport, formats)
	cli.PublicApIs = public_ap_is.New(transport, formats)
	cli.StackController = stack_controller.New(transport, formats)
	cli.UIAccountsController = ui_accounts_controller.New(transport, formats)
	cli.UIAlertsController = ui_alerts_controller.New(transport, formats)
	cli.UIApplicationController = ui_application_controller.New(transport, formats)
	cli.UIArtifactCiController = ui_artifact_ci_controller.New(transport, formats)
	cli.UIArtifactHubController = ui_artifact_hub_controller.New(transport, formats)
	cli.UIArtifactRoutingRuleController = ui_artifact_routing_rule_controller.New(transport, formats)
	cli.UIArtifactoryController = ui_artifactory_controller.New(transport, formats)
	cli.UIArtifactsController = ui_artifacts_controller.New(transport, formats)
	cli.UIAssistantController = ui_assistant_controller.New(transport, formats)
	cli.UIAuditLogsController = ui_audit_logs_controller.New(transport, formats)
	cli.UIAwsClusterController = ui_aws_cluster_controller.New(transport, formats)
	cli.UIAzureClusterController = ui_azure_cluster_controller.New(transport, formats)
	cli.UIBillingController = ui_billing_controller.New(transport, formats)
	cli.UIBlueprintDesignerController = ui_blueprint_designer_controller.New(transport, formats)
	cli.UIChatGptController = ui_chat_gpt_controller.New(transport, formats)
	cli.UICiCdController = ui_ci_cd_controller.New(transport, formats)
	cli.UICloudCostExplorerController = ui_cloud_cost_explorer_controller.New(transport, formats)
	cli.UICoderController = ui_coder_controller.New(transport, formats)
	cli.UICommonClusterController = ui_common_cluster_controller.New(transport, formats)
	cli.UICustomContentController = ui_custom_content_controller.New(transport, formats)
	cli.UICustomRoleController = ui_custom_role_controller.New(transport, formats)
	cli.UIDeliveryPipelineController = ui_delivery_pipeline_controller.New(transport, formats)
	cli.UIDeploymentController = ui_deployment_controller.New(transport, formats)
	cli.UIDomainMappingController = ui_domain_mapping_controller.New(transport, formats)
	cli.UIDropdownsController = ui_dropdowns_controller.New(transport, formats)
	cli.UIGcpClusterController = ui_gcp_cluster_controller.New(transport, formats)
	cli.UIIacRepoController = ui_iac_repo_controller.New(transport, formats)
	cli.UIIntentController = ui_intent_controller.New(transport, formats)
	cli.UIk8sClusterController = ui_k_8s_cluster_controller.New(transport, formats)
	cli.UIKubernetesExplorerController = ui_kubernetes_explorer_controller.New(transport, formats)
	cli.UILocalClusterController = ui_local_cluster_controller.New(transport, formats)
	cli.UIMaintenanceWindowController = ui_maintenance_window_controller.New(transport, formats)
	cli.UIMetaController = ui_meta_controller.New(transport, formats)
	cli.UINoAuthUserController = ui_no_auth_user_controller.New(transport, formats)
	cli.UINotificationController = ui_notification_controller.New(transport, formats)
	cli.UIoAuthController = ui_o_auth_controller.New(transport, formats)
	cli.UIOneTimeWebhookController = ui_one_time_webhook_controller.New(transport, formats)
	cli.UIOpaController = ui_opa_controller.New(transport, formats)
	cli.UIProjectTypeController = ui_project_type_controller.New(transport, formats)
	cli.UIPromotionWorkflowController = ui_promotion_workflow_controller.New(transport, formats)
	cli.UIProvidedResourcesController = ui_provided_resources_controller.New(transport, formats)
	cli.UIQueuedReleaseController = ui_queued_release_controller.New(transport, formats)
	cli.UIReleaseStreamController = ui_release_stream_controller.New(transport, formats)
	cli.UIResourceGroupController = ui_resource_group_controller.New(transport, formats)
	cli.UIResourceStatusController = ui_resource_status_controller.New(transport, formats)
	cli.UISettingsController = ui_settings_controller.New(transport, formats)
	cli.UIStackController = ui_stack_controller.New(transport, formats)
	cli.UITeamController = ui_team_controller.New(transport, formats)
	cli.UITfModuleController = ui_tf_module_controller.New(transport, formats)
	cli.UITfOutputController = ui_tf_output_controller.New(transport, formats)
	cli.UITfVersionController = ui_tf_version_controller.New(transport, formats)
	cli.UIThemeFileController = ui_theme_file_controller.New(transport, formats)
	cli.UIUserController = ui_user_controller.New(transport, formats)
	cli.UIUserGroupController = ui_user_group_controller.New(transport, formats)
	cli.UIUserVcsTokenController = ui_user_vcs_token_controller.New(transport, formats)
	cli.UIVersioningController = ui_versioning_controller.New(transport, formats)
	cli.UIWebComponentController = ui_web_component_controller.New(transport, formats)
	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// Facets is a client for facets
type Facets struct {
	ApplicationController application_controller.ClientService

	ArtifactController artifact_controller.ClientService

	ArtifactoryController artifactory_controller.ClientService

	AwsClusterController aws_cluster_controller.ClientService

	AzureClusterController azure_cluster_controller.ClientService

	BuildController build_controller.ClientService

	CallbackController callback_controller.ClientService

	CapillaryCloudCallbackController capillary_cloud_callback_controller.ClientService

	CommonClusterController common_cluster_controller.ClientService

	DeploymentController deployment_controller.ClientService

	GcpClusterController gcp_cluster_controller.ClientService

	MetaController meta_controller.ClientService

	PublicApIs public_ap_is.ClientService

	StackController stack_controller.ClientService

	UIAccountsController ui_accounts_controller.ClientService

	UIAlertsController ui_alerts_controller.ClientService

	UIApplicationController ui_application_controller.ClientService

	UIArtifactCiController ui_artifact_ci_controller.ClientService

	UIArtifactHubController ui_artifact_hub_controller.ClientService

	UIArtifactRoutingRuleController ui_artifact_routing_rule_controller.ClientService

	UIArtifactoryController ui_artifactory_controller.ClientService

	UIArtifactsController ui_artifacts_controller.ClientService

	UIAssistantController ui_assistant_controller.ClientService

	UIAuditLogsController ui_audit_logs_controller.ClientService

	UIAwsClusterController ui_aws_cluster_controller.ClientService

	UIAzureClusterController ui_azure_cluster_controller.ClientService

	UIBillingController ui_billing_controller.ClientService

	UIBlueprintDesignerController ui_blueprint_designer_controller.ClientService

	UIChatGptController ui_chat_gpt_controller.ClientService

	UICiCdController ui_ci_cd_controller.ClientService

	UICloudCostExplorerController ui_cloud_cost_explorer_controller.ClientService

	UICoderController ui_coder_controller.ClientService

	UICommonClusterController ui_common_cluster_controller.ClientService

	UICustomContentController ui_custom_content_controller.ClientService

	UICustomRoleController ui_custom_role_controller.ClientService

	UIDeliveryPipelineController ui_delivery_pipeline_controller.ClientService

	UIDeploymentController ui_deployment_controller.ClientService

	UIDomainMappingController ui_domain_mapping_controller.ClientService

	UIDropdownsController ui_dropdowns_controller.ClientService

	UIGcpClusterController ui_gcp_cluster_controller.ClientService

	UIIacRepoController ui_iac_repo_controller.ClientService

	UIIntentController ui_intent_controller.ClientService

	UIk8sClusterController ui_k_8s_cluster_controller.ClientService

	UIKubernetesExplorerController ui_kubernetes_explorer_controller.ClientService

	UILocalClusterController ui_local_cluster_controller.ClientService

	UIMaintenanceWindowController ui_maintenance_window_controller.ClientService

	UIMetaController ui_meta_controller.ClientService

	UINoAuthUserController ui_no_auth_user_controller.ClientService

	UINotificationController ui_notification_controller.ClientService

	UIoAuthController ui_o_auth_controller.ClientService

	UIOneTimeWebhookController ui_one_time_webhook_controller.ClientService

	UIOpaController ui_opa_controller.ClientService

	UIProjectTypeController ui_project_type_controller.ClientService

	UIPromotionWorkflowController ui_promotion_workflow_controller.ClientService

	UIProvidedResourcesController ui_provided_resources_controller.ClientService

	UIQueuedReleaseController ui_queued_release_controller.ClientService

	UIReleaseStreamController ui_release_stream_controller.ClientService

	UIResourceGroupController ui_resource_group_controller.ClientService

	UIResourceStatusController ui_resource_status_controller.ClientService

	UISettingsController ui_settings_controller.ClientService

	UIStackController ui_stack_controller.ClientService

	UITeamController ui_team_controller.ClientService

	UITfModuleController ui_tf_module_controller.ClientService

	UITfOutputController ui_tf_output_controller.ClientService

	UITfVersionController ui_tf_version_controller.ClientService

	UIThemeFileController ui_theme_file_controller.ClientService

	UIUserController ui_user_controller.ClientService

	UIUserGroupController ui_user_group_controller.ClientService

	UIUserVcsTokenController ui_user_vcs_token_controller.ClientService

	UIVersioningController ui_versioning_controller.ClientService

	UIWebComponentController ui_web_component_controller.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *Facets) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.ApplicationController.SetTransport(transport)
	c.ArtifactController.SetTransport(transport)
	c.ArtifactoryController.SetTransport(transport)
	c.AwsClusterController.SetTransport(transport)
	c.AzureClusterController.SetTransport(transport)
	c.BuildController.SetTransport(transport)
	c.CallbackController.SetTransport(transport)
	c.CapillaryCloudCallbackController.SetTransport(transport)
	c.CommonClusterController.SetTransport(transport)
	c.DeploymentController.SetTransport(transport)
	c.GcpClusterController.SetTransport(transport)
	c.MetaController.SetTransport(transport)
	c.PublicApIs.SetTransport(transport)
	c.StackController.SetTransport(transport)
	c.UIAccountsController.SetTransport(transport)
	c.UIAlertsController.SetTransport(transport)
	c.UIApplicationController.SetTransport(transport)
	c.UIArtifactCiController.SetTransport(transport)
	c.UIArtifactHubController.SetTransport(transport)
	c.UIArtifactRoutingRuleController.SetTransport(transport)
	c.UIArtifactoryController.SetTransport(transport)
	c.UIArtifactsController.SetTransport(transport)
	c.UIAssistantController.SetTransport(transport)
	c.UIAuditLogsController.SetTransport(transport)
	c.UIAwsClusterController.SetTransport(transport)
	c.UIAzureClusterController.SetTransport(transport)
	c.UIBillingController.SetTransport(transport)
	c.UIBlueprintDesignerController.SetTransport(transport)
	c.UIChatGptController.SetTransport(transport)
	c.UICiCdController.SetTransport(transport)
	c.UICloudCostExplorerController.SetTransport(transport)
	c.UICoderController.SetTransport(transport)
	c.UICommonClusterController.SetTransport(transport)
	c.UICustomContentController.SetTransport(transport)
	c.UICustomRoleController.SetTransport(transport)
	c.UIDeliveryPipelineController.SetTransport(transport)
	c.UIDeploymentController.SetTransport(transport)
	c.UIDomainMappingController.SetTransport(transport)
	c.UIDropdownsController.SetTransport(transport)
	c.UIGcpClusterController.SetTransport(transport)
	c.UIIacRepoController.SetTransport(transport)
	c.UIIntentController.SetTransport(transport)
	c.UIk8sClusterController.SetTransport(transport)
	c.UIKubernetesExplorerController.SetTransport(transport)
	c.UILocalClusterController.SetTransport(transport)
	c.UIMaintenanceWindowController.SetTransport(transport)
	c.UIMetaController.SetTransport(transport)
	c.UINoAuthUserController.SetTransport(transport)
	c.UINotificationController.SetTransport(transport)
	c.UIoAuthController.SetTransport(transport)
	c.UIOneTimeWebhookController.SetTransport(transport)
	c.UIOpaController.SetTransport(transport)
	c.UIProjectTypeController.SetTransport(transport)
	c.UIPromotionWorkflowController.SetTransport(transport)
	c.UIProvidedResourcesController.SetTransport(transport)
	c.UIQueuedReleaseController.SetTransport(transport)
	c.UIReleaseStreamController.SetTransport(transport)
	c.UIResourceGroupController.SetTransport(transport)
	c.UIResourceStatusController.SetTransport(transport)
	c.UISettingsController.SetTransport(transport)
	c.UIStackController.SetTransport(transport)
	c.UITeamController.SetTransport(transport)
	c.UITfModuleController.SetTransport(transport)
	c.UITfOutputController.SetTransport(transport)
	c.UITfVersionController.SetTransport(transport)
	c.UIThemeFileController.SetTransport(transport)
	c.UIUserController.SetTransport(transport)
	c.UIUserGroupController.SetTransport(transport)
	c.UIUserVcsTokenController.SetTransport(transport)
	c.UIVersioningController.SetTransport(transport)
	c.UIWebComponentController.SetTransport(transport)
}
