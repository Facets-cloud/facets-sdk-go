// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// FacetsAuditLogResponse FacetsAuditLogResponse
//
// swagger:model FacetsAuditLogResponse
type FacetsAuditLogResponse struct {

	// cluster Id
	ClusterID string `json:"clusterId,omitempty"`

	// cluster name
	ClusterName string `json:"clusterName,omitempty"`

	// entity
	// Enum: ["BLUEPRINT","ENVIRONMENT","RESOURCE","CI_INTEGRATION","CI_RULE","CI_CD","PROMOTION_WORKFLOW","RELEASE_STREAM","DELIVERY_PIPELINE","TRASH","RELEASE_SCHEDULE","AVAILABILITY_RULES","TEMPLATE_INPUT","SETTINGS","ACCOUNT","CONTAINER_REGISTRY","NOTIFICATION_CHANNEL","NOTIFICATION_SUBSCRIPTION","OAUTH_INTEGRATION","GUARDRAIL_POLICY","USER","USER_GROUP","CUSTOM_ROLE","RESOURCE_GROUP","PROJECT_TYPE","MODULE","WEB_COMPONENT","UNKNOWN"]
	Entity string `json:"entity,omitempty"`

	// entity action
	// Enum: ["BLUEPRINT_CREATE","BLUEPRINT_DELETE","BLUEPRINT_UPDATE","BLUEPRINT_RESOURCE_CREATE","BLUEPRINT_RESOURCE_UPDATE","BLUEPRINT_RESOURCE_RENAME","BLUEPRINT_RESOURCE_DELETE","BLUEPRINT_RESOURCE_ENABLE","BLUEPRINT_RESOURCE_DISABLE","BLUEPRINT_ADD_TEMPLATE","BLUEPRINT_SECRETS_VARIABLES_CREATE","BLUEPRINT_SECRETS_VARIABLES_UPDATE","BLUEPRINT_SECRETS_VARIABLES_DELETE","PROJECT_TYPE_CREATE","PROJECT_TYPE_UPDATE","PROJECT_TYPE_DELETE","MODULE_UPLOAD","MODULE_MARKED_AS_PUBLISHED","MODULE_DELETE","ENVIRONMENT_CREATE","ENVIRONMENT_DELETE","ENVIRONMENT_UPDATE","CI_INTEGRATION_CREATE","CI_INTEGRATION_UPDATE","CI_INTEGRATION_DELETE","CI_CD_CONFIGURE","CI_RULE_CREATE","CI_RULE_UPDATE","CI_RULE_DELETE","PROMOTION_WORKFLOW_CREATE","PROMOTION_WORKFLOW_UPDATE","PROMOTION_WORKFLOW_DELETE","RELEASE_STREAM_CREATE","RELEASE_STREAM_DELETE","DELIVERY_PIPELINE_UPDATE","OVERRIDE_UPDATE","OVERRIDE_VERSION_ROLLBACK","SECRETS_VARIABLES_UPDATE","TEMPLATE_INPUT_CREATE","TEMPLATE_INPUT_DELETE","IAC_VERSION_UPDATE","RELEASES_PAUSE","RELEASES_RESUME","RELEASE_SCHEDULE_CREATE","RELEASE_SCHEDULE_UPDATE","RELEASE_SCHEDULE_DELETE","AVAILABILITY_RULES_CREATE","AVAILABILITY_RULES_UPDATE","AVAILABILITY_RULES_DELETE","MAINTENANCE_WINDOW_UPDATE","ENVIRONMENT_SETTINGS_UPDATE","GENERAL_SETTINGS_UPDATE","ACCOUNT_CREATE","ACCOUNT_UPDATE","ACCOUNT_DELETE","CONTAINER_REGISTRY_CREATE","CONTAINER_REGISTRY_UPDATE","CONTAINER_REGISTRY_DELETE","NOTIFICATION_SUBSCRIPTION_CREATE","NOTIFICATION_SUBSCRIPTION_UPDATE","NOTIFICATION_SUBSCRIPTION_DELETE","NOTIFICATION_CHANNEL_CREATE","NOTIFICATION_CHANNEL_UPDATE","NOTIFICATION_CHANNEL_DELETE","OAUTH_INTEGRATION_CREATE","OAUTH_INTEGRATION_UPDATE","OAUTH_INTEGRATION_DELETE","ENVIRONMENT_RESOURCE_DISABLE","ENVIRONMENT_RESOURCE_ENABLE","GUARDRAIL_POLICY_CREATE","GUARDRAIL_POLICY_UPDATE","GUARDRAIL_POLICY_DELETE","USER_CREATE","USER_UPDATE","USER_DELETE","USER_GROUP_CREATE","USER_GROUP_UPDATE","USER_GROUP_DELETE","RESOURCE_GROUP_CREATE","RESOURCE_GROUP_UPDATE","RESOURCE_GROUP_DELETE","CUSTOM_ROLE_CREATE","CUSTOM_ROLE_UPDATE","CUSTOM_ROLE_DELETE","TRASH_RESTORE","TRASH_DELETE","TRASH_EMPTY","WEB_COMPONENT_CREATE","WEB_COMPONENT_UPDATE","WEB_COMPONENT_DELETE","UNKNOWN"]
	EntityAction string `json:"entityAction,omitempty"`

	// entity action label
	EntityActionLabel string `json:"entityActionLabel,omitempty"`

	// entity label
	EntityLabel string `json:"entityLabel,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// metadata
	Metadata interface{} `json:"metadata,omitempty"`

	// performed at
	// Format: date-time
	PerformedAt strfmt.DateTime `json:"performedAt,omitempty"`

	// performed by
	PerformedBy string `json:"performedBy,omitempty"`

	// stack name
	StackName string `json:"stackName,omitempty"`

	// target
	Target string `json:"target,omitempty"`
}

// Validate validates this facets audit log response
func (m *FacetsAuditLogResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEntity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePerformedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var facetsAuditLogResponseTypeEntityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["BLUEPRINT","ENVIRONMENT","RESOURCE","CI_INTEGRATION","CI_RULE","CI_CD","PROMOTION_WORKFLOW","RELEASE_STREAM","DELIVERY_PIPELINE","TRASH","RELEASE_SCHEDULE","AVAILABILITY_RULES","TEMPLATE_INPUT","SETTINGS","ACCOUNT","CONTAINER_REGISTRY","NOTIFICATION_CHANNEL","NOTIFICATION_SUBSCRIPTION","OAUTH_INTEGRATION","GUARDRAIL_POLICY","USER","USER_GROUP","CUSTOM_ROLE","RESOURCE_GROUP","PROJECT_TYPE","MODULE","WEB_COMPONENT","UNKNOWN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		facetsAuditLogResponseTypeEntityPropEnum = append(facetsAuditLogResponseTypeEntityPropEnum, v)
	}
}

const (

	// FacetsAuditLogResponseEntityBLUEPRINT captures enum value "BLUEPRINT"
	FacetsAuditLogResponseEntityBLUEPRINT string = "BLUEPRINT"

	// FacetsAuditLogResponseEntityENVIRONMENT captures enum value "ENVIRONMENT"
	FacetsAuditLogResponseEntityENVIRONMENT string = "ENVIRONMENT"

	// FacetsAuditLogResponseEntityRESOURCE captures enum value "RESOURCE"
	FacetsAuditLogResponseEntityRESOURCE string = "RESOURCE"

	// FacetsAuditLogResponseEntityCIINTEGRATION captures enum value "CI_INTEGRATION"
	FacetsAuditLogResponseEntityCIINTEGRATION string = "CI_INTEGRATION"

	// FacetsAuditLogResponseEntityCIRULE captures enum value "CI_RULE"
	FacetsAuditLogResponseEntityCIRULE string = "CI_RULE"

	// FacetsAuditLogResponseEntityCICD captures enum value "CI_CD"
	FacetsAuditLogResponseEntityCICD string = "CI_CD"

	// FacetsAuditLogResponseEntityPROMOTIONWORKFLOW captures enum value "PROMOTION_WORKFLOW"
	FacetsAuditLogResponseEntityPROMOTIONWORKFLOW string = "PROMOTION_WORKFLOW"

	// FacetsAuditLogResponseEntityRELEASESTREAM captures enum value "RELEASE_STREAM"
	FacetsAuditLogResponseEntityRELEASESTREAM string = "RELEASE_STREAM"

	// FacetsAuditLogResponseEntityDELIVERYPIPELINE captures enum value "DELIVERY_PIPELINE"
	FacetsAuditLogResponseEntityDELIVERYPIPELINE string = "DELIVERY_PIPELINE"

	// FacetsAuditLogResponseEntityTRASH captures enum value "TRASH"
	FacetsAuditLogResponseEntityTRASH string = "TRASH"

	// FacetsAuditLogResponseEntityRELEASESCHEDULE captures enum value "RELEASE_SCHEDULE"
	FacetsAuditLogResponseEntityRELEASESCHEDULE string = "RELEASE_SCHEDULE"

	// FacetsAuditLogResponseEntityAVAILABILITYRULES captures enum value "AVAILABILITY_RULES"
	FacetsAuditLogResponseEntityAVAILABILITYRULES string = "AVAILABILITY_RULES"

	// FacetsAuditLogResponseEntityTEMPLATEINPUT captures enum value "TEMPLATE_INPUT"
	FacetsAuditLogResponseEntityTEMPLATEINPUT string = "TEMPLATE_INPUT"

	// FacetsAuditLogResponseEntitySETTINGS captures enum value "SETTINGS"
	FacetsAuditLogResponseEntitySETTINGS string = "SETTINGS"

	// FacetsAuditLogResponseEntityACCOUNT captures enum value "ACCOUNT"
	FacetsAuditLogResponseEntityACCOUNT string = "ACCOUNT"

	// FacetsAuditLogResponseEntityCONTAINERREGISTRY captures enum value "CONTAINER_REGISTRY"
	FacetsAuditLogResponseEntityCONTAINERREGISTRY string = "CONTAINER_REGISTRY"

	// FacetsAuditLogResponseEntityNOTIFICATIONCHANNEL captures enum value "NOTIFICATION_CHANNEL"
	FacetsAuditLogResponseEntityNOTIFICATIONCHANNEL string = "NOTIFICATION_CHANNEL"

	// FacetsAuditLogResponseEntityNOTIFICATIONSUBSCRIPTION captures enum value "NOTIFICATION_SUBSCRIPTION"
	FacetsAuditLogResponseEntityNOTIFICATIONSUBSCRIPTION string = "NOTIFICATION_SUBSCRIPTION"

	// FacetsAuditLogResponseEntityOAUTHINTEGRATION captures enum value "OAUTH_INTEGRATION"
	FacetsAuditLogResponseEntityOAUTHINTEGRATION string = "OAUTH_INTEGRATION"

	// FacetsAuditLogResponseEntityGUARDRAILPOLICY captures enum value "GUARDRAIL_POLICY"
	FacetsAuditLogResponseEntityGUARDRAILPOLICY string = "GUARDRAIL_POLICY"

	// FacetsAuditLogResponseEntityUSER captures enum value "USER"
	FacetsAuditLogResponseEntityUSER string = "USER"

	// FacetsAuditLogResponseEntityUSERGROUP captures enum value "USER_GROUP"
	FacetsAuditLogResponseEntityUSERGROUP string = "USER_GROUP"

	// FacetsAuditLogResponseEntityCUSTOMROLE captures enum value "CUSTOM_ROLE"
	FacetsAuditLogResponseEntityCUSTOMROLE string = "CUSTOM_ROLE"

	// FacetsAuditLogResponseEntityRESOURCEGROUP captures enum value "RESOURCE_GROUP"
	FacetsAuditLogResponseEntityRESOURCEGROUP string = "RESOURCE_GROUP"

	// FacetsAuditLogResponseEntityPROJECTTYPE captures enum value "PROJECT_TYPE"
	FacetsAuditLogResponseEntityPROJECTTYPE string = "PROJECT_TYPE"

	// FacetsAuditLogResponseEntityMODULE captures enum value "MODULE"
	FacetsAuditLogResponseEntityMODULE string = "MODULE"

	// FacetsAuditLogResponseEntityWEBCOMPONENT captures enum value "WEB_COMPONENT"
	FacetsAuditLogResponseEntityWEBCOMPONENT string = "WEB_COMPONENT"

	// FacetsAuditLogResponseEntityUNKNOWN captures enum value "UNKNOWN"
	FacetsAuditLogResponseEntityUNKNOWN string = "UNKNOWN"
)

// prop value enum
func (m *FacetsAuditLogResponse) validateEntityEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, facetsAuditLogResponseTypeEntityPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *FacetsAuditLogResponse) validateEntity(formats strfmt.Registry) error {
	if swag.IsZero(m.Entity) { // not required
		return nil
	}

	// value enum
	if err := m.validateEntityEnum("entity", "body", m.Entity); err != nil {
		return err
	}

	return nil
}

var facetsAuditLogResponseTypeEntityActionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["BLUEPRINT_CREATE","BLUEPRINT_DELETE","BLUEPRINT_UPDATE","BLUEPRINT_RESOURCE_CREATE","BLUEPRINT_RESOURCE_UPDATE","BLUEPRINT_RESOURCE_RENAME","BLUEPRINT_RESOURCE_DELETE","BLUEPRINT_RESOURCE_ENABLE","BLUEPRINT_RESOURCE_DISABLE","BLUEPRINT_ADD_TEMPLATE","BLUEPRINT_SECRETS_VARIABLES_CREATE","BLUEPRINT_SECRETS_VARIABLES_UPDATE","BLUEPRINT_SECRETS_VARIABLES_DELETE","PROJECT_TYPE_CREATE","PROJECT_TYPE_UPDATE","PROJECT_TYPE_DELETE","MODULE_UPLOAD","MODULE_MARKED_AS_PUBLISHED","MODULE_DELETE","ENVIRONMENT_CREATE","ENVIRONMENT_DELETE","ENVIRONMENT_UPDATE","CI_INTEGRATION_CREATE","CI_INTEGRATION_UPDATE","CI_INTEGRATION_DELETE","CI_CD_CONFIGURE","CI_RULE_CREATE","CI_RULE_UPDATE","CI_RULE_DELETE","PROMOTION_WORKFLOW_CREATE","PROMOTION_WORKFLOW_UPDATE","PROMOTION_WORKFLOW_DELETE","RELEASE_STREAM_CREATE","RELEASE_STREAM_DELETE","DELIVERY_PIPELINE_UPDATE","OVERRIDE_UPDATE","OVERRIDE_VERSION_ROLLBACK","SECRETS_VARIABLES_UPDATE","TEMPLATE_INPUT_CREATE","TEMPLATE_INPUT_DELETE","IAC_VERSION_UPDATE","RELEASES_PAUSE","RELEASES_RESUME","RELEASE_SCHEDULE_CREATE","RELEASE_SCHEDULE_UPDATE","RELEASE_SCHEDULE_DELETE","AVAILABILITY_RULES_CREATE","AVAILABILITY_RULES_UPDATE","AVAILABILITY_RULES_DELETE","MAINTENANCE_WINDOW_UPDATE","ENVIRONMENT_SETTINGS_UPDATE","GENERAL_SETTINGS_UPDATE","ACCOUNT_CREATE","ACCOUNT_UPDATE","ACCOUNT_DELETE","CONTAINER_REGISTRY_CREATE","CONTAINER_REGISTRY_UPDATE","CONTAINER_REGISTRY_DELETE","NOTIFICATION_SUBSCRIPTION_CREATE","NOTIFICATION_SUBSCRIPTION_UPDATE","NOTIFICATION_SUBSCRIPTION_DELETE","NOTIFICATION_CHANNEL_CREATE","NOTIFICATION_CHANNEL_UPDATE","NOTIFICATION_CHANNEL_DELETE","OAUTH_INTEGRATION_CREATE","OAUTH_INTEGRATION_UPDATE","OAUTH_INTEGRATION_DELETE","ENVIRONMENT_RESOURCE_DISABLE","ENVIRONMENT_RESOURCE_ENABLE","GUARDRAIL_POLICY_CREATE","GUARDRAIL_POLICY_UPDATE","GUARDRAIL_POLICY_DELETE","USER_CREATE","USER_UPDATE","USER_DELETE","USER_GROUP_CREATE","USER_GROUP_UPDATE","USER_GROUP_DELETE","RESOURCE_GROUP_CREATE","RESOURCE_GROUP_UPDATE","RESOURCE_GROUP_DELETE","CUSTOM_ROLE_CREATE","CUSTOM_ROLE_UPDATE","CUSTOM_ROLE_DELETE","TRASH_RESTORE","TRASH_DELETE","TRASH_EMPTY","WEB_COMPONENT_CREATE","WEB_COMPONENT_UPDATE","WEB_COMPONENT_DELETE","UNKNOWN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		facetsAuditLogResponseTypeEntityActionPropEnum = append(facetsAuditLogResponseTypeEntityActionPropEnum, v)
	}
}

const (

	// FacetsAuditLogResponseEntityActionBLUEPRINTCREATE captures enum value "BLUEPRINT_CREATE"
	FacetsAuditLogResponseEntityActionBLUEPRINTCREATE string = "BLUEPRINT_CREATE"

	// FacetsAuditLogResponseEntityActionBLUEPRINTDELETE captures enum value "BLUEPRINT_DELETE"
	FacetsAuditLogResponseEntityActionBLUEPRINTDELETE string = "BLUEPRINT_DELETE"

	// FacetsAuditLogResponseEntityActionBLUEPRINTUPDATE captures enum value "BLUEPRINT_UPDATE"
	FacetsAuditLogResponseEntityActionBLUEPRINTUPDATE string = "BLUEPRINT_UPDATE"

	// FacetsAuditLogResponseEntityActionBLUEPRINTRESOURCECREATE captures enum value "BLUEPRINT_RESOURCE_CREATE"
	FacetsAuditLogResponseEntityActionBLUEPRINTRESOURCECREATE string = "BLUEPRINT_RESOURCE_CREATE"

	// FacetsAuditLogResponseEntityActionBLUEPRINTRESOURCEUPDATE captures enum value "BLUEPRINT_RESOURCE_UPDATE"
	FacetsAuditLogResponseEntityActionBLUEPRINTRESOURCEUPDATE string = "BLUEPRINT_RESOURCE_UPDATE"

	// FacetsAuditLogResponseEntityActionBLUEPRINTRESOURCERENAME captures enum value "BLUEPRINT_RESOURCE_RENAME"
	FacetsAuditLogResponseEntityActionBLUEPRINTRESOURCERENAME string = "BLUEPRINT_RESOURCE_RENAME"

	// FacetsAuditLogResponseEntityActionBLUEPRINTRESOURCEDELETE captures enum value "BLUEPRINT_RESOURCE_DELETE"
	FacetsAuditLogResponseEntityActionBLUEPRINTRESOURCEDELETE string = "BLUEPRINT_RESOURCE_DELETE"

	// FacetsAuditLogResponseEntityActionBLUEPRINTRESOURCEENABLE captures enum value "BLUEPRINT_RESOURCE_ENABLE"
	FacetsAuditLogResponseEntityActionBLUEPRINTRESOURCEENABLE string = "BLUEPRINT_RESOURCE_ENABLE"

	// FacetsAuditLogResponseEntityActionBLUEPRINTRESOURCEDISABLE captures enum value "BLUEPRINT_RESOURCE_DISABLE"
	FacetsAuditLogResponseEntityActionBLUEPRINTRESOURCEDISABLE string = "BLUEPRINT_RESOURCE_DISABLE"

	// FacetsAuditLogResponseEntityActionBLUEPRINTADDTEMPLATE captures enum value "BLUEPRINT_ADD_TEMPLATE"
	FacetsAuditLogResponseEntityActionBLUEPRINTADDTEMPLATE string = "BLUEPRINT_ADD_TEMPLATE"

	// FacetsAuditLogResponseEntityActionBLUEPRINTSECRETSVARIABLESCREATE captures enum value "BLUEPRINT_SECRETS_VARIABLES_CREATE"
	FacetsAuditLogResponseEntityActionBLUEPRINTSECRETSVARIABLESCREATE string = "BLUEPRINT_SECRETS_VARIABLES_CREATE"

	// FacetsAuditLogResponseEntityActionBLUEPRINTSECRETSVARIABLESUPDATE captures enum value "BLUEPRINT_SECRETS_VARIABLES_UPDATE"
	FacetsAuditLogResponseEntityActionBLUEPRINTSECRETSVARIABLESUPDATE string = "BLUEPRINT_SECRETS_VARIABLES_UPDATE"

	// FacetsAuditLogResponseEntityActionBLUEPRINTSECRETSVARIABLESDELETE captures enum value "BLUEPRINT_SECRETS_VARIABLES_DELETE"
	FacetsAuditLogResponseEntityActionBLUEPRINTSECRETSVARIABLESDELETE string = "BLUEPRINT_SECRETS_VARIABLES_DELETE"

	// FacetsAuditLogResponseEntityActionPROJECTTYPECREATE captures enum value "PROJECT_TYPE_CREATE"
	FacetsAuditLogResponseEntityActionPROJECTTYPECREATE string = "PROJECT_TYPE_CREATE"

	// FacetsAuditLogResponseEntityActionPROJECTTYPEUPDATE captures enum value "PROJECT_TYPE_UPDATE"
	FacetsAuditLogResponseEntityActionPROJECTTYPEUPDATE string = "PROJECT_TYPE_UPDATE"

	// FacetsAuditLogResponseEntityActionPROJECTTYPEDELETE captures enum value "PROJECT_TYPE_DELETE"
	FacetsAuditLogResponseEntityActionPROJECTTYPEDELETE string = "PROJECT_TYPE_DELETE"

	// FacetsAuditLogResponseEntityActionMODULEUPLOAD captures enum value "MODULE_UPLOAD"
	FacetsAuditLogResponseEntityActionMODULEUPLOAD string = "MODULE_UPLOAD"

	// FacetsAuditLogResponseEntityActionMODULEMARKEDASPUBLISHED captures enum value "MODULE_MARKED_AS_PUBLISHED"
	FacetsAuditLogResponseEntityActionMODULEMARKEDASPUBLISHED string = "MODULE_MARKED_AS_PUBLISHED"

	// FacetsAuditLogResponseEntityActionMODULEDELETE captures enum value "MODULE_DELETE"
	FacetsAuditLogResponseEntityActionMODULEDELETE string = "MODULE_DELETE"

	// FacetsAuditLogResponseEntityActionENVIRONMENTCREATE captures enum value "ENVIRONMENT_CREATE"
	FacetsAuditLogResponseEntityActionENVIRONMENTCREATE string = "ENVIRONMENT_CREATE"

	// FacetsAuditLogResponseEntityActionENVIRONMENTDELETE captures enum value "ENVIRONMENT_DELETE"
	FacetsAuditLogResponseEntityActionENVIRONMENTDELETE string = "ENVIRONMENT_DELETE"

	// FacetsAuditLogResponseEntityActionENVIRONMENTUPDATE captures enum value "ENVIRONMENT_UPDATE"
	FacetsAuditLogResponseEntityActionENVIRONMENTUPDATE string = "ENVIRONMENT_UPDATE"

	// FacetsAuditLogResponseEntityActionCIINTEGRATIONCREATE captures enum value "CI_INTEGRATION_CREATE"
	FacetsAuditLogResponseEntityActionCIINTEGRATIONCREATE string = "CI_INTEGRATION_CREATE"

	// FacetsAuditLogResponseEntityActionCIINTEGRATIONUPDATE captures enum value "CI_INTEGRATION_UPDATE"
	FacetsAuditLogResponseEntityActionCIINTEGRATIONUPDATE string = "CI_INTEGRATION_UPDATE"

	// FacetsAuditLogResponseEntityActionCIINTEGRATIONDELETE captures enum value "CI_INTEGRATION_DELETE"
	FacetsAuditLogResponseEntityActionCIINTEGRATIONDELETE string = "CI_INTEGRATION_DELETE"

	// FacetsAuditLogResponseEntityActionCICDCONFIGURE captures enum value "CI_CD_CONFIGURE"
	FacetsAuditLogResponseEntityActionCICDCONFIGURE string = "CI_CD_CONFIGURE"

	// FacetsAuditLogResponseEntityActionCIRULECREATE captures enum value "CI_RULE_CREATE"
	FacetsAuditLogResponseEntityActionCIRULECREATE string = "CI_RULE_CREATE"

	// FacetsAuditLogResponseEntityActionCIRULEUPDATE captures enum value "CI_RULE_UPDATE"
	FacetsAuditLogResponseEntityActionCIRULEUPDATE string = "CI_RULE_UPDATE"

	// FacetsAuditLogResponseEntityActionCIRULEDELETE captures enum value "CI_RULE_DELETE"
	FacetsAuditLogResponseEntityActionCIRULEDELETE string = "CI_RULE_DELETE"

	// FacetsAuditLogResponseEntityActionPROMOTIONWORKFLOWCREATE captures enum value "PROMOTION_WORKFLOW_CREATE"
	FacetsAuditLogResponseEntityActionPROMOTIONWORKFLOWCREATE string = "PROMOTION_WORKFLOW_CREATE"

	// FacetsAuditLogResponseEntityActionPROMOTIONWORKFLOWUPDATE captures enum value "PROMOTION_WORKFLOW_UPDATE"
	FacetsAuditLogResponseEntityActionPROMOTIONWORKFLOWUPDATE string = "PROMOTION_WORKFLOW_UPDATE"

	// FacetsAuditLogResponseEntityActionPROMOTIONWORKFLOWDELETE captures enum value "PROMOTION_WORKFLOW_DELETE"
	FacetsAuditLogResponseEntityActionPROMOTIONWORKFLOWDELETE string = "PROMOTION_WORKFLOW_DELETE"

	// FacetsAuditLogResponseEntityActionRELEASESTREAMCREATE captures enum value "RELEASE_STREAM_CREATE"
	FacetsAuditLogResponseEntityActionRELEASESTREAMCREATE string = "RELEASE_STREAM_CREATE"

	// FacetsAuditLogResponseEntityActionRELEASESTREAMDELETE captures enum value "RELEASE_STREAM_DELETE"
	FacetsAuditLogResponseEntityActionRELEASESTREAMDELETE string = "RELEASE_STREAM_DELETE"

	// FacetsAuditLogResponseEntityActionDELIVERYPIPELINEUPDATE captures enum value "DELIVERY_PIPELINE_UPDATE"
	FacetsAuditLogResponseEntityActionDELIVERYPIPELINEUPDATE string = "DELIVERY_PIPELINE_UPDATE"

	// FacetsAuditLogResponseEntityActionOVERRIDEUPDATE captures enum value "OVERRIDE_UPDATE"
	FacetsAuditLogResponseEntityActionOVERRIDEUPDATE string = "OVERRIDE_UPDATE"

	// FacetsAuditLogResponseEntityActionOVERRIDEVERSIONROLLBACK captures enum value "OVERRIDE_VERSION_ROLLBACK"
	FacetsAuditLogResponseEntityActionOVERRIDEVERSIONROLLBACK string = "OVERRIDE_VERSION_ROLLBACK"

	// FacetsAuditLogResponseEntityActionSECRETSVARIABLESUPDATE captures enum value "SECRETS_VARIABLES_UPDATE"
	FacetsAuditLogResponseEntityActionSECRETSVARIABLESUPDATE string = "SECRETS_VARIABLES_UPDATE"

	// FacetsAuditLogResponseEntityActionTEMPLATEINPUTCREATE captures enum value "TEMPLATE_INPUT_CREATE"
	FacetsAuditLogResponseEntityActionTEMPLATEINPUTCREATE string = "TEMPLATE_INPUT_CREATE"

	// FacetsAuditLogResponseEntityActionTEMPLATEINPUTDELETE captures enum value "TEMPLATE_INPUT_DELETE"
	FacetsAuditLogResponseEntityActionTEMPLATEINPUTDELETE string = "TEMPLATE_INPUT_DELETE"

	// FacetsAuditLogResponseEntityActionIACVERSIONUPDATE captures enum value "IAC_VERSION_UPDATE"
	FacetsAuditLogResponseEntityActionIACVERSIONUPDATE string = "IAC_VERSION_UPDATE"

	// FacetsAuditLogResponseEntityActionRELEASESPAUSE captures enum value "RELEASES_PAUSE"
	FacetsAuditLogResponseEntityActionRELEASESPAUSE string = "RELEASES_PAUSE"

	// FacetsAuditLogResponseEntityActionRELEASESRESUME captures enum value "RELEASES_RESUME"
	FacetsAuditLogResponseEntityActionRELEASESRESUME string = "RELEASES_RESUME"

	// FacetsAuditLogResponseEntityActionRELEASESCHEDULECREATE captures enum value "RELEASE_SCHEDULE_CREATE"
	FacetsAuditLogResponseEntityActionRELEASESCHEDULECREATE string = "RELEASE_SCHEDULE_CREATE"

	// FacetsAuditLogResponseEntityActionRELEASESCHEDULEUPDATE captures enum value "RELEASE_SCHEDULE_UPDATE"
	FacetsAuditLogResponseEntityActionRELEASESCHEDULEUPDATE string = "RELEASE_SCHEDULE_UPDATE"

	// FacetsAuditLogResponseEntityActionRELEASESCHEDULEDELETE captures enum value "RELEASE_SCHEDULE_DELETE"
	FacetsAuditLogResponseEntityActionRELEASESCHEDULEDELETE string = "RELEASE_SCHEDULE_DELETE"

	// FacetsAuditLogResponseEntityActionAVAILABILITYRULESCREATE captures enum value "AVAILABILITY_RULES_CREATE"
	FacetsAuditLogResponseEntityActionAVAILABILITYRULESCREATE string = "AVAILABILITY_RULES_CREATE"

	// FacetsAuditLogResponseEntityActionAVAILABILITYRULESUPDATE captures enum value "AVAILABILITY_RULES_UPDATE"
	FacetsAuditLogResponseEntityActionAVAILABILITYRULESUPDATE string = "AVAILABILITY_RULES_UPDATE"

	// FacetsAuditLogResponseEntityActionAVAILABILITYRULESDELETE captures enum value "AVAILABILITY_RULES_DELETE"
	FacetsAuditLogResponseEntityActionAVAILABILITYRULESDELETE string = "AVAILABILITY_RULES_DELETE"

	// FacetsAuditLogResponseEntityActionMAINTENANCEWINDOWUPDATE captures enum value "MAINTENANCE_WINDOW_UPDATE"
	FacetsAuditLogResponseEntityActionMAINTENANCEWINDOWUPDATE string = "MAINTENANCE_WINDOW_UPDATE"

	// FacetsAuditLogResponseEntityActionENVIRONMENTSETTINGSUPDATE captures enum value "ENVIRONMENT_SETTINGS_UPDATE"
	FacetsAuditLogResponseEntityActionENVIRONMENTSETTINGSUPDATE string = "ENVIRONMENT_SETTINGS_UPDATE"

	// FacetsAuditLogResponseEntityActionGENERALSETTINGSUPDATE captures enum value "GENERAL_SETTINGS_UPDATE"
	FacetsAuditLogResponseEntityActionGENERALSETTINGSUPDATE string = "GENERAL_SETTINGS_UPDATE"

	// FacetsAuditLogResponseEntityActionACCOUNTCREATE captures enum value "ACCOUNT_CREATE"
	FacetsAuditLogResponseEntityActionACCOUNTCREATE string = "ACCOUNT_CREATE"

	// FacetsAuditLogResponseEntityActionACCOUNTUPDATE captures enum value "ACCOUNT_UPDATE"
	FacetsAuditLogResponseEntityActionACCOUNTUPDATE string = "ACCOUNT_UPDATE"

	// FacetsAuditLogResponseEntityActionACCOUNTDELETE captures enum value "ACCOUNT_DELETE"
	FacetsAuditLogResponseEntityActionACCOUNTDELETE string = "ACCOUNT_DELETE"

	// FacetsAuditLogResponseEntityActionCONTAINERREGISTRYCREATE captures enum value "CONTAINER_REGISTRY_CREATE"
	FacetsAuditLogResponseEntityActionCONTAINERREGISTRYCREATE string = "CONTAINER_REGISTRY_CREATE"

	// FacetsAuditLogResponseEntityActionCONTAINERREGISTRYUPDATE captures enum value "CONTAINER_REGISTRY_UPDATE"
	FacetsAuditLogResponseEntityActionCONTAINERREGISTRYUPDATE string = "CONTAINER_REGISTRY_UPDATE"

	// FacetsAuditLogResponseEntityActionCONTAINERREGISTRYDELETE captures enum value "CONTAINER_REGISTRY_DELETE"
	FacetsAuditLogResponseEntityActionCONTAINERREGISTRYDELETE string = "CONTAINER_REGISTRY_DELETE"

	// FacetsAuditLogResponseEntityActionNOTIFICATIONSUBSCRIPTIONCREATE captures enum value "NOTIFICATION_SUBSCRIPTION_CREATE"
	FacetsAuditLogResponseEntityActionNOTIFICATIONSUBSCRIPTIONCREATE string = "NOTIFICATION_SUBSCRIPTION_CREATE"

	// FacetsAuditLogResponseEntityActionNOTIFICATIONSUBSCRIPTIONUPDATE captures enum value "NOTIFICATION_SUBSCRIPTION_UPDATE"
	FacetsAuditLogResponseEntityActionNOTIFICATIONSUBSCRIPTIONUPDATE string = "NOTIFICATION_SUBSCRIPTION_UPDATE"

	// FacetsAuditLogResponseEntityActionNOTIFICATIONSUBSCRIPTIONDELETE captures enum value "NOTIFICATION_SUBSCRIPTION_DELETE"
	FacetsAuditLogResponseEntityActionNOTIFICATIONSUBSCRIPTIONDELETE string = "NOTIFICATION_SUBSCRIPTION_DELETE"

	// FacetsAuditLogResponseEntityActionNOTIFICATIONCHANNELCREATE captures enum value "NOTIFICATION_CHANNEL_CREATE"
	FacetsAuditLogResponseEntityActionNOTIFICATIONCHANNELCREATE string = "NOTIFICATION_CHANNEL_CREATE"

	// FacetsAuditLogResponseEntityActionNOTIFICATIONCHANNELUPDATE captures enum value "NOTIFICATION_CHANNEL_UPDATE"
	FacetsAuditLogResponseEntityActionNOTIFICATIONCHANNELUPDATE string = "NOTIFICATION_CHANNEL_UPDATE"

	// FacetsAuditLogResponseEntityActionNOTIFICATIONCHANNELDELETE captures enum value "NOTIFICATION_CHANNEL_DELETE"
	FacetsAuditLogResponseEntityActionNOTIFICATIONCHANNELDELETE string = "NOTIFICATION_CHANNEL_DELETE"

	// FacetsAuditLogResponseEntityActionOAUTHINTEGRATIONCREATE captures enum value "OAUTH_INTEGRATION_CREATE"
	FacetsAuditLogResponseEntityActionOAUTHINTEGRATIONCREATE string = "OAUTH_INTEGRATION_CREATE"

	// FacetsAuditLogResponseEntityActionOAUTHINTEGRATIONUPDATE captures enum value "OAUTH_INTEGRATION_UPDATE"
	FacetsAuditLogResponseEntityActionOAUTHINTEGRATIONUPDATE string = "OAUTH_INTEGRATION_UPDATE"

	// FacetsAuditLogResponseEntityActionOAUTHINTEGRATIONDELETE captures enum value "OAUTH_INTEGRATION_DELETE"
	FacetsAuditLogResponseEntityActionOAUTHINTEGRATIONDELETE string = "OAUTH_INTEGRATION_DELETE"

	// FacetsAuditLogResponseEntityActionENVIRONMENTRESOURCEDISABLE captures enum value "ENVIRONMENT_RESOURCE_DISABLE"
	FacetsAuditLogResponseEntityActionENVIRONMENTRESOURCEDISABLE string = "ENVIRONMENT_RESOURCE_DISABLE"

	// FacetsAuditLogResponseEntityActionENVIRONMENTRESOURCEENABLE captures enum value "ENVIRONMENT_RESOURCE_ENABLE"
	FacetsAuditLogResponseEntityActionENVIRONMENTRESOURCEENABLE string = "ENVIRONMENT_RESOURCE_ENABLE"

	// FacetsAuditLogResponseEntityActionGUARDRAILPOLICYCREATE captures enum value "GUARDRAIL_POLICY_CREATE"
	FacetsAuditLogResponseEntityActionGUARDRAILPOLICYCREATE string = "GUARDRAIL_POLICY_CREATE"

	// FacetsAuditLogResponseEntityActionGUARDRAILPOLICYUPDATE captures enum value "GUARDRAIL_POLICY_UPDATE"
	FacetsAuditLogResponseEntityActionGUARDRAILPOLICYUPDATE string = "GUARDRAIL_POLICY_UPDATE"

	// FacetsAuditLogResponseEntityActionGUARDRAILPOLICYDELETE captures enum value "GUARDRAIL_POLICY_DELETE"
	FacetsAuditLogResponseEntityActionGUARDRAILPOLICYDELETE string = "GUARDRAIL_POLICY_DELETE"

	// FacetsAuditLogResponseEntityActionUSERCREATE captures enum value "USER_CREATE"
	FacetsAuditLogResponseEntityActionUSERCREATE string = "USER_CREATE"

	// FacetsAuditLogResponseEntityActionUSERUPDATE captures enum value "USER_UPDATE"
	FacetsAuditLogResponseEntityActionUSERUPDATE string = "USER_UPDATE"

	// FacetsAuditLogResponseEntityActionUSERDELETE captures enum value "USER_DELETE"
	FacetsAuditLogResponseEntityActionUSERDELETE string = "USER_DELETE"

	// FacetsAuditLogResponseEntityActionUSERGROUPCREATE captures enum value "USER_GROUP_CREATE"
	FacetsAuditLogResponseEntityActionUSERGROUPCREATE string = "USER_GROUP_CREATE"

	// FacetsAuditLogResponseEntityActionUSERGROUPUPDATE captures enum value "USER_GROUP_UPDATE"
	FacetsAuditLogResponseEntityActionUSERGROUPUPDATE string = "USER_GROUP_UPDATE"

	// FacetsAuditLogResponseEntityActionUSERGROUPDELETE captures enum value "USER_GROUP_DELETE"
	FacetsAuditLogResponseEntityActionUSERGROUPDELETE string = "USER_GROUP_DELETE"

	// FacetsAuditLogResponseEntityActionRESOURCEGROUPCREATE captures enum value "RESOURCE_GROUP_CREATE"
	FacetsAuditLogResponseEntityActionRESOURCEGROUPCREATE string = "RESOURCE_GROUP_CREATE"

	// FacetsAuditLogResponseEntityActionRESOURCEGROUPUPDATE captures enum value "RESOURCE_GROUP_UPDATE"
	FacetsAuditLogResponseEntityActionRESOURCEGROUPUPDATE string = "RESOURCE_GROUP_UPDATE"

	// FacetsAuditLogResponseEntityActionRESOURCEGROUPDELETE captures enum value "RESOURCE_GROUP_DELETE"
	FacetsAuditLogResponseEntityActionRESOURCEGROUPDELETE string = "RESOURCE_GROUP_DELETE"

	// FacetsAuditLogResponseEntityActionCUSTOMROLECREATE captures enum value "CUSTOM_ROLE_CREATE"
	FacetsAuditLogResponseEntityActionCUSTOMROLECREATE string = "CUSTOM_ROLE_CREATE"

	// FacetsAuditLogResponseEntityActionCUSTOMROLEUPDATE captures enum value "CUSTOM_ROLE_UPDATE"
	FacetsAuditLogResponseEntityActionCUSTOMROLEUPDATE string = "CUSTOM_ROLE_UPDATE"

	// FacetsAuditLogResponseEntityActionCUSTOMROLEDELETE captures enum value "CUSTOM_ROLE_DELETE"
	FacetsAuditLogResponseEntityActionCUSTOMROLEDELETE string = "CUSTOM_ROLE_DELETE"

	// FacetsAuditLogResponseEntityActionTRASHRESTORE captures enum value "TRASH_RESTORE"
	FacetsAuditLogResponseEntityActionTRASHRESTORE string = "TRASH_RESTORE"

	// FacetsAuditLogResponseEntityActionTRASHDELETE captures enum value "TRASH_DELETE"
	FacetsAuditLogResponseEntityActionTRASHDELETE string = "TRASH_DELETE"

	// FacetsAuditLogResponseEntityActionTRASHEMPTY captures enum value "TRASH_EMPTY"
	FacetsAuditLogResponseEntityActionTRASHEMPTY string = "TRASH_EMPTY"

	// FacetsAuditLogResponseEntityActionWEBCOMPONENTCREATE captures enum value "WEB_COMPONENT_CREATE"
	FacetsAuditLogResponseEntityActionWEBCOMPONENTCREATE string = "WEB_COMPONENT_CREATE"

	// FacetsAuditLogResponseEntityActionWEBCOMPONENTUPDATE captures enum value "WEB_COMPONENT_UPDATE"
	FacetsAuditLogResponseEntityActionWEBCOMPONENTUPDATE string = "WEB_COMPONENT_UPDATE"

	// FacetsAuditLogResponseEntityActionWEBCOMPONENTDELETE captures enum value "WEB_COMPONENT_DELETE"
	FacetsAuditLogResponseEntityActionWEBCOMPONENTDELETE string = "WEB_COMPONENT_DELETE"

	// FacetsAuditLogResponseEntityActionUNKNOWN captures enum value "UNKNOWN"
	FacetsAuditLogResponseEntityActionUNKNOWN string = "UNKNOWN"
)

// prop value enum
func (m *FacetsAuditLogResponse) validateEntityActionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, facetsAuditLogResponseTypeEntityActionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *FacetsAuditLogResponse) validateEntityAction(formats strfmt.Registry) error {
	if swag.IsZero(m.EntityAction) { // not required
		return nil
	}

	// value enum
	if err := m.validateEntityActionEnum("entityAction", "body", m.EntityAction); err != nil {
		return err
	}

	return nil
}

func (m *FacetsAuditLogResponse) validatePerformedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.PerformedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("performedAt", "body", "date-time", m.PerformedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this facets audit log response based on context it is used
func (m *FacetsAuditLogResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FacetsAuditLogResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FacetsAuditLogResponse) UnmarshalBinary(b []byte) error {
	var res FacetsAuditLogResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
