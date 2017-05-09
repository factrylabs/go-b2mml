// WARNING: This package was auto generated.
// Changes to this file may cause incorrect behavior and
// will be lost if the code is regenerated.

package b2mml

import (
	"encoding/xml"
	"time"
)

type XmlNode struct {
	XmlName xml.Name
	Content []byte    `xml:",innerxml"`
	Nodes   []XmlNode `xml:",any"`
}

type TransResponseCodeType string

// ConfirmBODType was auto-generated. Do not change.
type ConfirmBODType struct {
	ApplicationArea TransApplicationAreaType `json:",omitempty" xml:",omitempty"`
	DataArea        ConfirmBODTypeDataArea   `json:",omitempty" xml:",omitempty"`
}

// TransApplicationAreaType was auto-generated. Do not change.
type TransApplicationAreaType struct {
	Sender           TransSenderType     `json:",omitempty" xml:",omitempty"`
	Receiver         []TransReceiverType `json:",omitempty" xml:",omitempty"`
	CreationDateTime DateTimeType        `json:",omitempty" xml:",omitempty"`
	Signature        TransSignatureType  `json:",omitempty" xml:",omitempty"`
	BODID            IdentifierType      `json:",omitempty" xml:",omitempty"`
	UserArea         TransUserAreaType   `json:",omitempty" xml:",omitempty"`
}

// TransSenderType was auto-generated. Do not change.
type TransSenderType struct {
	LogicalID        IdentifierType            `json:",omitempty" xml:",omitempty"`
	ComponentID      IdentifierType            `json:",omitempty" xml:",omitempty"`
	TaskID           IdentifierType            `json:",omitempty" xml:",omitempty"`
	ReferenceID      IdentifierType            `json:",omitempty" xml:",omitempty"`
	ConfirmationCode TransConfirmationCodeType `json:",omitempty" xml:",omitempty"`
	AuthorizationID  IdentifierType            `json:",omitempty" xml:",omitempty"`
}

// IdentifierType was auto-generated. Do not change.
type IdentifierType struct {
	SchemeID         string `json:",omitempty" xml:"schemeID,attr,omitempty"`
	SchemeName       string `json:",omitempty" xml:"schemeName,attr,omitempty"`
	SchemeAgencyID   string `json:",omitempty" xml:"schemeAgencyID,attr,omitempty"`
	SchemeAgencyName string `json:",omitempty" xml:"schemeAgencyName,attr,omitempty"`
	SchemeVersionID  string `json:",omitempty" xml:"schemeVersionID,attr,omitempty"`
	SchemeDataURI    string `json:",omitempty" xml:"schemeDataURI,attr,omitempty"`
	SchemeURI        string `json:",omitempty" xml:"schemeURI,attr,omitempty"`
	Value            string `json:",omitempty" xml:",chardata"`
}

// BODType was auto-generated. Do not change.
type BODType struct {
	OriginalApplicationArea TransApplicationAreaType `json:",omitempty" xml:",omitempty"`
	Description             []TextType               `json:",omitempty" xml:",omitempty"`
	Note                    []TextType               `json:",omitempty" xml:",omitempty"`
	UserArea                TransUserAreaType        `json:",omitempty" xml:",omitempty"`
}

// TextType was auto-generated. Do not change.
type TextType struct {
	LanguageID string `json:",omitempty" xml:"languageID,attr,omitempty"`
	Value      string `json:",omitempty" xml:",chardata"`
}

// ScheduleEntryNoteType was auto-generated. Do not change.
type ScheduleEntryNoteType struct {
	TextType
}

// NoteType was auto-generated. Do not change.
type NoteType struct {
	TextType
}

// EnumerationStringType was auto-generated. Do not change.
type EnumerationStringType struct {
	TextType
}

// DefaultValueType was auto-generated. Do not change.
type DefaultValueType struct {
	TextType
}

// JobOrderCommandRuleType was auto-generated. Do not change.
type JobOrderCommandRuleType struct {
	TextType
}

// DescriptionType was auto-generated. Do not change.
type DescriptionType struct {
	TextType
}

// TransUserAreaType was auto-generated. Do not change.
type TransUserAreaType struct {
	Any []XmlNode `json:",omitempty" xml:",omitempty"`
}

// TransStateChangeType was auto-generated. Do not change.
type TransStateChangeType struct {
	FromStateCode  CodeType          `json:",omitempty" xml:",omitempty"`
	ToStateCode    CodeType          `json:",omitempty" xml:",omitempty"`
	ChangeDateTime DateTimeType      `json:",omitempty" xml:",omitempty"`
	Description    []DescriptionType `json:",omitempty" xml:",omitempty"`
	Note           []TextType        `json:",omitempty" xml:",omitempty"`
	UserArea       TransUserAreaType `json:",omitempty" xml:",omitempty"`
}

// CodeType was auto-generated. Do not change.
type CodeType struct {
	ListID         string `json:",omitempty" xml:"listID,attr,omitempty"`
	ListAgencyID   string `json:",omitempty" xml:"listAgencyID,attr,omitempty"`
	ListAgencyName string `json:",omitempty" xml:"listAgencyName,attr,omitempty"`
	ListName       string `json:",omitempty" xml:"listName,attr,omitempty"`
	ListVersionID  string `json:",omitempty" xml:"listVersionID,attr,omitempty"`
	Name           string `json:",omitempty" xml:"name,attr,omitempty"`
	LanguageID     string `json:",omitempty" xml:"languageID,attr,omitempty"`
	ListURI        string `json:",omitempty" xml:"listURI,attr,omitempty"`
	ListSchemeURI  string `json:",omitempty" xml:"listSchemeURI,attr,omitempty"`
	Value          string `json:",omitempty" xml:",chardata"`
}

// SequenceOrderType1Type was auto-generated. Do not change.
type SequenceOrderType1Type struct {
	CodeType
}

// SequenceOrderType was auto-generated. Do not change.
type SequenceOrderType struct {
	SequenceOrderType1Type
	OtherValue string `json:",omitempty" xml:"OtherValue,attr,omitempty"`
}

// LifeCycleStateType1Type was auto-generated. Do not change.
type LifeCycleStateType1Type struct {
	CodeType
}

// LifeCycleStateType was auto-generated. Do not change.
type LifeCycleStateType struct {
	LifeCycleStateType1Type
	OtherValue string `json:",omitempty" xml:"OtherValue,attr,omitempty"`
}

// GRecipeType1Type was auto-generated. Do not change.
type GRecipeType1Type struct {
	CodeType
}

// GRecipeTypeType was auto-generated. Do not change.
type GRecipeTypeType struct {
	GRecipeType1Type
	OtherValue string `json:",omitempty" xml:"OtherValue,attr,omitempty"`
}

// ProcedureChartElementType1Type was auto-generated. Do not change.
type ProcedureChartElementType1Type struct {
	CodeType
}

// ProcedureChartElementTypeType was auto-generated. Do not change.
type ProcedureChartElementTypeType struct {
	ProcedureChartElementType1Type
	OtherValue string `json:",omitempty" xml:"OtherValue,attr,omitempty"`
}

// ProcessElementType1Type was auto-generated. Do not change.
type ProcessElementType1Type struct {
	CodeType
}

// ProcessElementTypeType was auto-generated. Do not change.
type ProcessElementTypeType struct {
	ProcessElementType1Type
	OtherValue string `json:",omitempty" xml:"OtherValue,attr,omitempty"`
}

// ProcessElementParameterType1Type was auto-generated. Do not change.
type ProcessElementParameterType1Type struct {
	CodeType
}

// ProcessElementParameterTypeType was auto-generated. Do not change.
type ProcessElementParameterTypeType struct {
	ProcessElementParameterType1Type
	OtherValue string `json:",omitempty" xml:"OtherValue,attr,omitempty"`
}

// MaterialsType1Type was auto-generated. Do not change.
type MaterialsType1Type struct {
	CodeType
}

// MaterialsTypeType was auto-generated. Do not change.
type MaterialsTypeType struct {
	MaterialsType1Type
	OtherValue string `json:",omitempty" xml:"OtherValue,attr,omitempty"`
}

// ConstraintType1Type was auto-generated. Do not change.
type ConstraintType1Type struct {
	CodeType
}

// ConstraintTypeType was auto-generated. Do not change.
type ConstraintTypeType struct {
	ConstraintType1Type
	OtherValue string `json:",omitempty" xml:"OtherValue,attr,omitempty"`
}

// EventSubType1Type was auto-generated. Do not change.
type EventSubType1Type struct {
	CodeType
}

// EventSubTypeType was auto-generated. Do not change.
type EventSubTypeType struct {
	EventSubType1Type
	OtherValue string `json:",omitempty" xml:"OtherValue,attr,omitempty"`
}

// EventType1Type was auto-generated. Do not change.
type EventType1Type struct {
	CodeType
}

// EventTypeType was auto-generated. Do not change.
type EventTypeType struct {
	EventType1Type
	OtherValue string `json:",omitempty" xml:"OtherValue,attr,omitempty"`
}

// RecordObjectType1Type was auto-generated. Do not change.
type RecordObjectType1Type struct {
	CodeType
}

// RecordObjectTypeType was auto-generated. Do not change.
type RecordObjectTypeType struct {
	RecordObjectType1Type
	OtherValue string `json:",omitempty" xml:"OtherValue,attr,omitempty"`
}

// ToType1Type was auto-generated. Do not change.
type ToType1Type struct {
	CodeType
}

// ToTypeType was auto-generated. Do not change.
type ToTypeType struct {
	ToType1Type
	OtherValue string `json:",omitempty" xml:"OtherValue,attr,omitempty"`
}

// BatchStatus1Type was auto-generated. Do not change.
type BatchStatus1Type struct {
	CodeType
}

// BatchStatusType was auto-generated. Do not change.
type BatchStatusType struct {
	BatchStatus1Type
	OtherValue string `json:",omitempty" xml:"OtherValue,attr,omitempty"`
}

// ScaledType was auto-generated. Do not change.
type ScaledType struct {
	CodeType
}

// RecipeElementType1Type was auto-generated. Do not change.
type RecipeElementType1Type struct {
	CodeType
}

// RecipeElementTypeType was auto-generated. Do not change.
type RecipeElementTypeType struct {
	RecipeElementType1Type
	OtherValue string `json:",omitempty" xml:"OtherValue,attr,omitempty"`
}

// ParameterType1Type was auto-generated. Do not change.
type ParameterType1Type struct {
	CodeType
}

// ParameterTypeType was auto-generated. Do not change.
type ParameterTypeType struct {
	ParameterType1Type
	OtherValue string `json:",omitempty" xml:"OtherValue,attr,omitempty"`
}

// Mode1Type was auto-generated. Do not change.
type Mode1Type struct {
	CodeType
}

// ModeType was auto-generated. Do not change.
type ModeType struct {
	Mode1Type
	OtherValue string `json:",omitempty" xml:"OtherValue,attr,omitempty"`
}

// LinkType1Type was auto-generated. Do not change.
type LinkType1Type struct {
	CodeType
}

// LinkTypeType was auto-generated. Do not change.
type LinkTypeType struct {
	LinkType1Type
	OtherValue string `json:",omitempty" xml:"OtherValue,attr,omitempty"`
}

// IDScope1Type was auto-generated. Do not change.
type IDScope1Type struct {
	CodeType
}

// IDScopeType was auto-generated. Do not change.
type IDScopeType struct {
	IDScope1Type
	OtherValue string `json:",omitempty" xml:"OtherValue,attr,omitempty"`
}

// FromType1Type was auto-generated. Do not change.
type FromType1Type struct {
	CodeType
}

// FromTypeType was auto-generated. Do not change.
type FromTypeType struct {
	FromType1Type
	OtherValue string `json:",omitempty" xml:"OtherValue,attr,omitempty"`
}

// EquipmentProceduralElementType1Type was auto-generated. Do not change.
type EquipmentProceduralElementType1Type struct {
	CodeType
}

// EquipmentProceduralElementTypeType was auto-generated. Do not change.
type EquipmentProceduralElementTypeType struct {
	EquipmentProceduralElementType1Type
	OtherValue string `json:",omitempty" xml:"OtherValue,attr,omitempty"`
}

// EquipmentElementType1Type was auto-generated. Do not change.
type EquipmentElementType1Type struct {
	CodeType
}

// EquipmentElementTypeType was auto-generated. Do not change.
type EquipmentElementTypeType struct {
	EquipmentElementType1Type
	OtherValue string `json:",omitempty" xml:"OtherValue,attr,omitempty"`
}

// Depiction1Type was auto-generated. Do not change.
type Depiction1Type struct {
	CodeType
}

// DepictionType was auto-generated. Do not change.
type DepictionType struct {
	Depiction1Type
	OtherValue string `json:",omitempty" xml:"OtherValue,attr,omitempty"`
}

// DataInterpretation1Type was auto-generated. Do not change.
type DataInterpretation1Type struct {
	CodeType
}

// DataInterpretationType was auto-generated. Do not change.
type DataInterpretationType struct {
	DataInterpretation1Type
	OtherValue string `json:",omitempty" xml:"OtherValue,attr,omitempty"`
}

// ConnectionType1Type was auto-generated. Do not change.
type ConnectionType1Type struct {
	CodeType
}

// ConnectionTypeType was auto-generated. Do not change.
type ConnectionTypeType struct {
	ConnectionType1Type
	OtherValue string `json:",omitempty" xml:"OtherValue,attr,omitempty"`
}

// BatchListEntryType1Type was auto-generated. Do not change.
type BatchListEntryType1Type struct {
	CodeType
}

// BatchListEntryTypeType was auto-generated. Do not change.
type BatchListEntryTypeType struct {
	BatchListEntryType1Type
	OtherValue string `json:",omitempty" xml:"OtherValue,attr,omitempty"`
}

// TransactionVerb1Type was auto-generated. Do not change.
type TransactionVerb1Type struct {
	CodeType
}

// TransactionVerbType was auto-generated. Do not change.
type TransactionVerbType struct {
	TransactionVerb1Type
	OtherValue string `json:",omitempty" xml:"OtherValue,attr,omitempty"`
}

// TransactionNoun1Type was auto-generated. Do not change.
type TransactionNoun1Type struct {
	CodeType
}

// TransactionNounType was auto-generated. Do not change.
type TransactionNounType struct {
	TransactionNoun1Type
	OtherValue string `json:",omitempty" xml:"OtherValue,attr,omitempty"`
}

// TransConfirmationCodeType was auto-generated. Do not change.
type TransConfirmationCodeType struct {
	CodeType
}

// WorkType1Type was auto-generated. Do not change.
type WorkType1Type struct {
	CodeType
}

// WorkTypeType was auto-generated. Do not change.
type WorkTypeType struct {
	WorkType1Type
	OtherValue string `json:",omitempty" xml:"OtherValue,attr,omitempty"`
}

// UnitOfMeasureType was auto-generated. Do not change.
type UnitOfMeasureType struct {
	CodeType
}

// StatusType was auto-generated. Do not change.
type StatusType struct {
	CodeType
}

// ResponseState1Type was auto-generated. Do not change.
type ResponseState1Type struct {
	CodeType
}

// ResponseStateType was auto-generated. Do not change.
type ResponseStateType struct {
	ResponseState1Type
	OtherValue string `json:",omitempty" xml:"OtherValue,attr,omitempty"`
}

// ResourcesType was auto-generated. Do not change.
type ResourcesType struct {
	CodeType
}

// ResourceReferenceType1Type was auto-generated. Do not change.
type ResourceReferenceType1Type struct {
	CodeType
}

// ResourceReferenceTypeType was auto-generated. Do not change.
type ResourceReferenceTypeType struct {
	ResourceReferenceType1Type
	OtherValue string `json:",omitempty" xml:"OtherValue,attr,omitempty"`
}

// RequiredByRequestedSegmentResponse1Type was auto-generated. Do not change.
type RequiredByRequestedSegmentResponse1Type struct {
	CodeType
}

// RequiredByRequestedSegmentResponseType was auto-generated. Do not change.
type RequiredByRequestedSegmentResponseType struct {
	RequiredByRequestedSegmentResponse1Type
	OtherValue string `json:",omitempty" xml:"OtherValue,attr,omitempty"`
}

// RequestState1Type was auto-generated. Do not change.
type RequestState1Type struct {
	CodeType
}

// RequestStateType was auto-generated. Do not change.
type RequestStateType struct {
	RequestState1Type
	OtherValue string `json:",omitempty" xml:"OtherValue,attr,omitempty"`
}

// RelationshipType1Type was auto-generated. Do not change.
type RelationshipType1Type struct {
	CodeType
}

// RelationshipTypeType was auto-generated. Do not change.
type RelationshipTypeType struct {
	RelationshipType1Type
	OtherValue string `json:",omitempty" xml:"OtherValue,attr,omitempty"`
}

// RelationshipForm1Type was auto-generated. Do not change.
type RelationshipForm1Type struct {
	CodeType
}

// RelationshipFormType was auto-generated. Do not change.
type RelationshipFormType struct {
	RelationshipForm1Type
	OtherValue string `json:",omitempty" xml:"OtherValue,attr,omitempty"`
}

// ReasonType was auto-generated. Do not change.
type ReasonType struct {
	CodeType
}

// ProblemType was auto-generated. Do not change.
type ProblemType struct {
	CodeType
}

// PhysicalAssetUseType was auto-generated. Do not change.
type PhysicalAssetUseType struct {
	CodeType
}

// PersonnelUseType was auto-generated. Do not change.
type PersonnelUseType struct {
	CodeType
}

// OtherDependencyType was auto-generated. Do not change.
type OtherDependencyType struct {
	CodeType
}

// OperationsType1Type was auto-generated. Do not change.
type OperationsType1Type struct {
	CodeType
}

// OperationsTypeType was auto-generated. Do not change.
type OperationsTypeType struct {
	OperationsType1Type
	OtherValue string `json:",omitempty" xml:"OtherValue,attr,omitempty"`
}

// MaterialUse1Type was auto-generated. Do not change.
type MaterialUse1Type struct {
	CodeType
}

// MaterialUseType was auto-generated. Do not change.
type MaterialUseType struct {
	MaterialUse1Type
	OtherValue string `json:",omitempty" xml:"OtherValue,attr,omitempty"`
}

// JobOrderDispatchStatusType was auto-generated. Do not change.
type JobOrderDispatchStatusType struct {
	CodeType
}

// JobOrderCommand1Type was auto-generated. Do not change.
type JobOrderCommand1Type struct {
	CodeType
}

// JobOrderCommandType was auto-generated. Do not change.
type JobOrderCommandType struct {
	JobOrderCommand1Type
	OtherValue string `json:",omitempty" xml:"OtherValue,attr,omitempty"`
}

// EquipmentUseType was auto-generated. Do not change.
type EquipmentUseType struct {
	CodeType
}

// EquipmentElementLevel1Type was auto-generated. Do not change.
type EquipmentElementLevel1Type struct {
	CodeType
}

// EquipmentElementLevelType was auto-generated. Do not change.
type EquipmentElementLevelType struct {
	EquipmentElementLevel1Type
	OtherValue string `json:",omitempty" xml:"OtherValue,attr,omitempty"`
}

// Dependency1Type was auto-generated. Do not change.
type Dependency1Type struct {
	CodeType
}

// DependencyType was auto-generated. Do not change.
type DependencyType struct {
	Dependency1Type
	OtherValue string `json:",omitempty" xml:"OtherValue,attr,omitempty"`
}

// DataType1Type was auto-generated. Do not change.
type DataType1Type struct {
	CodeType
}

// DataTypeType was auto-generated. Do not change.
type DataTypeType struct {
	DataType1Type
	OtherValue string `json:",omitempty" xml:"OtherValue,attr,omitempty"`
}

// CorrectionType was auto-generated. Do not change.
type CorrectionType struct {
	CodeType
}

// CauseType was auto-generated. Do not change.
type CauseType struct {
	CodeType
}

// CapabilityType1Type was auto-generated. Do not change.
type CapabilityType1Type struct {
	CodeType
}

// CapabilityTypeType was auto-generated. Do not change.
type CapabilityTypeType struct {
	CapabilityType1Type
	OtherValue string `json:",omitempty" xml:"OtherValue,attr,omitempty"`
}

// AssemblyType1Type was auto-generated. Do not change.
type AssemblyType1Type struct {
	CodeType
}

// AssemblyTypeType was auto-generated. Do not change.
type AssemblyTypeType struct {
	AssemblyType1Type
	OtherValue string `json:",omitempty" xml:"OtherValue,attr,omitempty"`
}

// AssemblyRelationship1Type was auto-generated. Do not change.
type AssemblyRelationship1Type struct {
	CodeType
}

// AssemblyRelationshipType was auto-generated. Do not change.
type AssemblyRelationshipType struct {
	AssemblyRelationship1Type
	OtherValue string `json:",omitempty" xml:"OtherValue,attr,omitempty"`
}

// DateTimeType was auto-generated. Do not change.
type DateTimeType struct {
	Format string    `json:",omitempty" xml:"format,attr,omitempty"`
	Value  time.Time `json:",omitempty" xml:",chardata"`
}

// VersionDateType was auto-generated. Do not change.
type VersionDateType struct {
	DateTimeType
}

// StartConditionType was auto-generated. Do not change.
type StartConditionType struct {
	DateTimeType
}

// ScheduleStartTimeType was auto-generated. Do not change.
type ScheduleStartTimeType struct {
	DateTimeType
}

// ScheduleEndTimeType was auto-generated. Do not change.
type ScheduleEndTimeType struct {
	DateTimeType
}

// RequestedStartTimeType was auto-generated. Do not change.
type RequestedStartTimeType struct {
	DateTimeType
}

// RequestedEndTimeType was auto-generated. Do not change.
type RequestedEndTimeType struct {
	DateTimeType
}

// ModifiedDateType was auto-generated. Do not change.
type ModifiedDateType struct {
	DateTimeType
}

// FinalApprovalDateType was auto-generated. Do not change.
type FinalApprovalDateType struct {
	DateTimeType
}

// ExpirationDateType was auto-generated. Do not change.
type ExpirationDateType struct {
	DateTimeType
}

// EffectiveDateType was auto-generated. Do not change.
type EffectiveDateType struct {
	DateTimeType
}

// CreateDateType was auto-generated. Do not change.
type CreateDateType struct {
	DateTimeType
}

// ApprovalDateType was auto-generated. Do not change.
type ApprovalDateType struct {
	DateTimeType
}

// TestDateTimeType was auto-generated. Do not change.
type TestDateTimeType struct {
	DateTimeType
}

// StatusTimeType was auto-generated. Do not change.
type StatusTimeType struct {
	DateTimeType
}

// StartTimeType was auto-generated. Do not change.
type StartTimeType struct {
	DateTimeType
}

// RequestedCompletionDateType was auto-generated. Do not change.
type RequestedCompletionDateType struct {
	DateTimeType
}

// PublishedDateType was auto-generated. Do not change.
type PublishedDateType struct {
	DateTimeType
}

// PlannedStartTimeType was auto-generated. Do not change.
type PlannedStartTimeType struct {
	DateTimeType
}

// PlannedFinishTimeType was auto-generated. Do not change.
type PlannedFinishTimeType struct {
	DateTimeType
}

// LatestEndTimeType was auto-generated. Do not change.
type LatestEndTimeType struct {
	DateTimeType
}

// ExpirationTimeType was auto-generated. Do not change.
type ExpirationTimeType struct {
	DateTimeType
}

// EndTimeType was auto-generated. Do not change.
type EndTimeType struct {
	DateTimeType
}

// EarliestStartTimeType was auto-generated. Do not change.
type EarliestStartTimeType struct {
	DateTimeType
}

// ActualStartTimeType was auto-generated. Do not change.
type ActualStartTimeType struct {
	DateTimeType
}

// ActualFinishTimeType was auto-generated. Do not change.
type ActualFinishTimeType struct {
	DateTimeType
}

// ActualEndTimeType was auto-generated. Do not change.
type ActualEndTimeType struct {
	DateTimeType
}

// TransChangeStatusType was auto-generated. Do not change.
type TransChangeStatusType struct {
	Code              CodeType               `json:",omitempty" xml:",omitempty"`
	Description       DescriptionType        `json:",omitempty" xml:",omitempty"`
	EffectiveDateTime DateTimeType           `json:",omitempty" xml:",omitempty"`
	ReasonCode        CodeType               `json:",omitempty" xml:",omitempty"`
	Reason            []CodeType             `json:",omitempty" xml:",omitempty"`
	StateChange       []TransStateChangeType `json:",omitempty" xml:",omitempty"`
	UserArea          TransUserAreaType      `json:",omitempty" xml:",omitempty"`
}

// TransExpression1Type was auto-generated. Do not change.
type TransExpression1Type struct {
	Value string `json:",omitempty" xml:",chardata"`
}

// TransExpressionType was auto-generated. Do not change.
type TransExpressionType struct {
	TransExpression1Type
	ActionCode         string `json:",omitempty" xml:"actionCode,attr,omitempty"`
	ExpressionLanguage string `json:",omitempty" xml:"expressionLanguage,attr,omitempty"`
}

// TransResponseCriteriaType was auto-generated. Do not change.
type TransResponseCriteriaType struct {
	ResponseExpression TransExpressionType   `json:",omitempty" xml:",omitempty"`
	ChangeStatus       TransChangeStatusType `json:",omitempty" xml:",omitempty"`
}

// TransConfirmType was auto-generated. Do not change.
type TransConfirmType struct {
	OriginalApplicationArea TransApplicationAreaType    `json:",omitempty" xml:",omitempty"`
	ResponseCriteria        []TransResponseCriteriaType `json:",omitempty" xml:",omitempty"`
}

// TransSignatureType was auto-generated. Do not change.
type TransSignatureType struct {
	Any                []XmlNode `json:",omitempty" xml:",omitempty"`
	QualifyingAgencyID string    `json:",omitempty" xml:"qualifyingAgencyID,attr,omitempty"`
}

// TransReceiverType was auto-generated. Do not change.
type TransReceiverType struct {
	LogicalID   IdentifierType   `json:",omitempty" xml:",omitempty"`
	ComponentID IdentifierType   `json:",omitempty" xml:",omitempty"`
	ID          []IdentifierType `json:",omitempty" xml:",omitempty"`
}

// ToEquipmentIDType was auto-generated. Do not change.
type ToEquipmentIDType struct {
	IdentifierType
}

// RecipeVersionType was auto-generated. Do not change.
type RecipeVersionType struct {
	IdentifierType
}

// RecipeIDType was auto-generated. Do not change.
type RecipeIDType struct {
	IdentifierType
}

// RecipeElementVersionType was auto-generated. Do not change.
type RecipeElementVersionType struct {
	IdentifierType
}

// RecipeElementIDType was auto-generated. Do not change.
type RecipeElementIDType struct {
	IdentifierType
}

// ProductNameType was auto-generated. Do not change.
type ProductNameType struct {
	IdentifierType
}

// ProductIDType was auto-generated. Do not change.
type ProductIDType struct {
	IdentifierType
}

// ParameterSubTypeType was auto-generated. Do not change.
type ParameterSubTypeType struct {
	IdentifierType
}

// OriginType was auto-generated. Do not change.
type OriginType struct {
	IdentifierType
}

// OrderIDType was auto-generated. Do not change.
type OrderIDType struct {
	IdentifierType
}

// MemberEquipmentIDType was auto-generated. Do not change.
type MemberEquipmentIDType struct {
	IdentifierType
}

// LotIDType was auto-generated. Do not change.
type LotIDType struct {
	IdentifierType
}

// IDType was auto-generated. Do not change.
type IDType struct {
	IdentifierType
}

// FromEquipmentIDType was auto-generated. Do not change.
type FromEquipmentIDType struct {
	IdentifierType
}

// ExternalIDType was auto-generated. Do not change.
type ExternalIDType struct {
	IdentifierType
}

// EquipmentProceduralElementClassIDType was auto-generated. Do not change.
type EquipmentProceduralElementClassIDType struct {
	IdentifierType
}

// EquipmentElementIDType was auto-generated. Do not change.
type EquipmentElementIDType struct {
	IdentifierType
}

// BatchEquipmentClassIDType was auto-generated. Do not change.
type BatchEquipmentClassIDType struct {
	IdentifierType
}

// EnumerationSetIDType was auto-generated. Do not change.
type EnumerationSetIDType struct {
	IdentifierType
}

// ConditionAnnotationType was auto-generated. Do not change.
type ConditionAnnotationType struct {
	IdentifierType
}

// ConditionType was auto-generated. Do not change.
type ConditionType struct {
	IdentifierType
}

// ClassEquipmentIDType was auto-generated. Do not change.
type ClassEquipmentIDType struct {
	IdentifierType
}

// CampaignIDType was auto-generated. Do not change.
type CampaignIDType struct {
	IdentifierType
}

// BuildingBlockElementVersionType was auto-generated. Do not change.
type BuildingBlockElementVersionType struct {
	IdentifierType
}

// BuildingBlockElementIDType was auto-generated. Do not change.
type BuildingBlockElementIDType struct {
	IdentifierType
}

// BatchIDType was auto-generated. Do not change.
type BatchIDType struct {
	IdentifierType
}

// AuthorType was auto-generated. Do not change.
type AuthorType struct {
	IdentifierType
}

// ActualProductProducedType was auto-generated. Do not change.
type ActualProductProducedType struct {
	IdentifierType
}

// ActualEquipmentIDType was auto-generated. Do not change.
type ActualEquipmentIDType struct {
	IdentifierType
}

// WorkScheduleIDType was auto-generated. Do not change.
type WorkScheduleIDType struct {
	IdentifierType
}

// WorkRequestIDType was auto-generated. Do not change.
type WorkRequestIDType struct {
	IdentifierType
}

// VersionType was auto-generated. Do not change.
type VersionType struct {
	IdentifierType
}

// StorageHierarchyScopeType was auto-generated. Do not change.
type StorageHierarchyScopeType struct {
	IdentifierType
}

// StorageLocationType was auto-generated. Do not change.
type StorageLocationType struct {
	IdentifierType
}

// ResourceIDType was auto-generated. Do not change.
type ResourceIDType struct {
	IdentifierType
}

// ResourceNetworkConnectionIDType was auto-generated. Do not change.
type ResourceNetworkConnectionIDType struct {
	IdentifierType
}

// QualificationTestSpecificationIDType was auto-generated. Do not change.
type QualificationTestSpecificationIDType struct {
	IdentifierType
}

// PropertyIDType was auto-generated. Do not change.
type PropertyIDType struct {
	IdentifierType
}

// ProductSegmentIDType was auto-generated. Do not change.
type ProductSegmentIDType struct {
	IdentifierType
}

// ProductionRequestIDType was auto-generated. Do not change.
type ProductionRequestIDType struct {
	IdentifierType
}

// ProductProductionRuleType was auto-generated. Do not change.
type ProductProductionRuleType struct {
	IdentifierType
}

// ProductProductionRuleIDType was auto-generated. Do not change.
type ProductProductionRuleIDType struct {
	IdentifierType
}

// ProductionScheduleIDType was auto-generated. Do not change.
type ProductionScheduleIDType struct {
	IdentifierType
}

// ProcessSegmentIDType was auto-generated. Do not change.
type ProcessSegmentIDType struct {
	IdentifierType
}

// PhysicalAssetIDType was auto-generated. Do not change.
type PhysicalAssetIDType struct {
	IdentifierType
}

// PhysicalAssetClassIDType was auto-generated. Do not change.
type PhysicalAssetClassIDType struct {
	IdentifierType
}

// PhysicalAssetCapabilityTestSpecificationIDType was auto-generated. Do not change.
type PhysicalAssetCapabilityTestSpecificationIDType struct {
	IdentifierType
}

// PhysicalAssetActualIDType was auto-generated. Do not change.
type PhysicalAssetActualIDType struct {
	IdentifierType
}

// PersonnelClassIDType was auto-generated. Do not change.
type PersonnelClassIDType struct {
	IdentifierType
}

// PersonNameType was auto-generated. Do not change.
type PersonNameType struct {
	IdentifierType
}

// PersonIDType was auto-generated. Do not change.
type PersonIDType struct {
	IdentifierType
}

// ParameterIDType was auto-generated. Do not change.
type ParameterIDType struct {
	IdentifierType
}

// OperationsSegmentIDType was auto-generated. Do not change.
type OperationsSegmentIDType struct {
	IdentifierType
}

// OperationsScheduleIDType was auto-generated. Do not change.
type OperationsScheduleIDType struct {
	IdentifierType
}

// OperationsRequestIDType was auto-generated. Do not change.
type OperationsRequestIDType struct {
	IdentifierType
}

// OperationsDefinitionIDType was auto-generated. Do not change.
type OperationsDefinitionIDType struct {
	IdentifierType
}

// MaterialTestSpecificationIDType was auto-generated. Do not change.
type MaterialTestSpecificationIDType struct {
	IdentifierType
}

// MaterialSubLotIDType was auto-generated. Do not change.
type MaterialSubLotIDType struct {
	IdentifierType
}

// MaterialSpecificationIDType was auto-generated. Do not change.
type MaterialSpecificationIDType struct {
	IdentifierType
}

// MaterialRequirementIDType was auto-generated. Do not change.
type MaterialRequirementIDType struct {
	IdentifierType
}

// MaterialLotIDType was auto-generated. Do not change.
type MaterialLotIDType struct {
	IdentifierType
}

// MaterialDefinitionIDType was auto-generated. Do not change.
type MaterialDefinitionIDType struct {
	IdentifierType
}

// MaterialClassIDType was auto-generated. Do not change.
type MaterialClassIDType struct {
	IdentifierType
}

// MaterialCapabilityIDType was auto-generated. Do not change.
type MaterialCapabilityIDType struct {
	IdentifierType
}

// MaterialActualIDType was auto-generated. Do not change.
type MaterialActualIDType struct {
	IdentifierType
}

// ManufacturingBillIDType was auto-generated. Do not change.
type ManufacturingBillIDType struct {
	IdentifierType
}

// EquipmentIDType was auto-generated. Do not change.
type EquipmentIDType struct {
	IdentifierType
}

// EquipmentClassIDType was auto-generated. Do not change.
type EquipmentClassIDType struct {
	IdentifierType
}

// EquipmentCapabilityTestSpecificationIDType was auto-generated. Do not change.
type EquipmentCapabilityTestSpecificationIDType struct {
	IdentifierType
}

// ConfidenceFactorType was auto-generated. Do not change.
type ConfidenceFactorType struct {
	IdentifierType
}

// CertificateOfAnalysisReferenceType was auto-generated. Do not change.
type CertificateOfAnalysisReferenceType struct {
	IdentifierType
}

// BillOfResourcesIDType was auto-generated. Do not change.
type BillOfResourcesIDType struct {
	IdentifierType
}

// BillOfMaterialsIDType was auto-generated. Do not change.
type BillOfMaterialsIDType struct {
	IdentifierType
}

// BillOfMaterialIDType was auto-generated. Do not change.
type BillOfMaterialIDType struct {
	IdentifierType
}

// ConfirmBODTypeDataArea was auto-generated. Do not change.
type ConfirmBODTypeDataArea struct {
	Confirm TransConfirmType `json:",omitempty" xml:",omitempty"`
	BOD     []BODType        `json:",omitempty" xml:",omitempty"`
}

// EquipmentInformationType was auto-generated. Do not change.
type EquipmentInformationType struct {
	ID                                   IdentifierType                             `json:",omitempty" xml:",omitempty"`
	Description                          []DescriptionType                          `json:",omitempty" xml:",omitempty"`
	Location                             LocationType                               `json:",omitempty" xml:",omitempty"`
	HierarchyScope                       HierarchyScopeType                         `json:",omitempty" xml:",omitempty"`
	PublishedDate                        PublishedDateType                          `json:",omitempty" xml:",omitempty"`
	Equipment                            []EquipmentType                            `json:",omitempty" xml:",omitempty"`
	EquipmentClass                       []EquipmentClassType                       `json:",omitempty" xml:",omitempty"`
	EquipmentCapabilityTestSpecification []EquipmentCapabilityTestSpecificationType `json:",omitempty" xml:",omitempty"`
}

// LocationType was auto-generated. Do not change.
type LocationType struct {
	EquipmentID           EquipmentIDType           `json:",omitempty" xml:",omitempty"`
	EquipmentElementLevel EquipmentElementLevelType `json:",omitempty" xml:",omitempty"`
	Location              *LocationType             `json:",omitempty" xml:",omitempty"`
}

// HierarchyScopeType was auto-generated. Do not change.
type HierarchyScopeType struct {
	EquipmentID           EquipmentIDType           `json:",omitempty" xml:",omitempty"`
	EquipmentElementLevel EquipmentElementLevelType `json:",omitempty" xml:",omitempty"`
	HierarchyScope        *HierarchyScopeType       `json:",omitempty" xml:",omitempty"`
}

// EquipmentType was auto-generated. Do not change.
type EquipmentType struct {
	ID                                     IdentifierType                               `json:",omitempty" xml:",omitempty"`
	Description                            []DescriptionType                            `json:",omitempty" xml:",omitempty"`
	Location                               LocationType                                 `json:",omitempty" xml:",omitempty"`
	HierarchyScope                         HierarchyScopeType                           `json:",omitempty" xml:",omitempty"`
	EquipmentLevel                         HierarchyScopeType                           `json:",omitempty" xml:",omitempty"`
	EquipmentAssetMapping                  []EquipmentAssetMappingType                  `json:",omitempty" xml:",omitempty"`
	EquipmentProperty                      []EquipmentPropertyType                      `json:",omitempty" xml:",omitempty"`
	Equipment                              []EquipmentType                              `json:",omitempty" xml:",omitempty"`
	EquipmentClassID                       []EquipmentClassIDType                       `json:",omitempty" xml:",omitempty"`
	EquipmentCapabilityTestSpecificationID []EquipmentCapabilityTestSpecificationIDType `json:",omitempty" xml:",omitempty"`
}

// EquipmentAssetMappingType was auto-generated. Do not change.
type EquipmentAssetMappingType struct {
	EquipmentID     EquipmentIDType     `json:",omitempty" xml:",omitempty"`
	PhysicalAssetID PhysicalAssetIDType `json:",omitempty" xml:",omitempty"`
	StartTime       DateTimeType        `json:",omitempty" xml:",omitempty"`
	EndTime         DateTimeType        `json:",omitempty" xml:",omitempty"`
}

// EquipmentPropertyType was auto-generated. Do not change.
type EquipmentPropertyType struct {
	ID                                     IdentifierType                               `json:",omitempty" xml:",omitempty"`
	Description                            []DescriptionType                            `json:",omitempty" xml:",omitempty"`
	Value                                  []ValueType                                  `json:",omitempty" xml:",omitempty"`
	EquipmentProperty                      []EquipmentPropertyType                      `json:",omitempty" xml:",omitempty"`
	EquipmentCapabilityTestSpecificationID []EquipmentCapabilityTestSpecificationIDType `json:",omitempty" xml:",omitempty"`
	TestResult                             []TestResultType                             `json:",omitempty" xml:",omitempty"`
}

// ValueType was auto-generated. Do not change.
type ValueType struct {
	ValueString   ValueStringType   `json:",omitempty" xml:",omitempty"`
	DataType      DataTypeType      `json:",omitempty" xml:",omitempty"`
	UnitOfMeasure UnitOfMeasureType `json:",omitempty" xml:",omitempty"`
	Key           IdentifierType    `json:",omitempty" xml:",omitempty"`
}

// ValueStringType was auto-generated. Do not change.
type ValueStringType struct {
	AnyGenericValueType
}

// AnyGenericValueType was auto-generated. Do not change.
type AnyGenericValueType struct {
	CurrencyID                string `json:",omitempty" xml:"currencyID,attr,omitempty"`
	CurrencyCodeListVersionID string `json:",omitempty" xml:"currencyCodeListVersionID,attr,omitempty"`
	EncodingCode              string `json:",omitempty" xml:"encodingCode,attr,omitempty"`
	Format                    string `json:",omitempty" xml:"format,attr,omitempty"`
	CharacterSetCode          string `json:",omitempty" xml:"characterSetCode,attr,omitempty"`
	ListID                    string `json:",omitempty" xml:"listID,attr,omitempty"`
	ListAgencyID              string `json:",omitempty" xml:"listAgencyID,attr,omitempty"`
	ListAgencyName            string `json:",omitempty" xml:"listAgencyName,attr,omitempty"`
	ListName                  string `json:",omitempty" xml:"listName,attr,omitempty"`
	ListVersionID             string `json:",omitempty" xml:"listVersionID,attr,omitempty"`
	LanguageID                string `json:",omitempty" xml:"languageID,attr,omitempty"`
	LanguageLocaleID          string `json:",omitempty" xml:"languageLocaleID,attr,omitempty"`
	ListURI                   string `json:",omitempty" xml:"listURI,attr,omitempty"`
	ListSchemaURI             string `json:",omitempty" xml:"listSchemaURI,attr,omitempty"`
	MimeCode                  string `json:",omitempty" xml:"mimeCode,attr,omitempty"`
	Name                      string `json:",omitempty" xml:"name,attr,omitempty"`
	SchemaID                  string `json:",omitempty" xml:"schemaID,attr,omitempty"`
	SchemaName                string `json:",omitempty" xml:"schemaName,attr,omitempty"`
	SchemaAgencyID            string `json:",omitempty" xml:"schemaAgencyID,attr,omitempty"`
	SchemaAgencyName          string `json:",omitempty" xml:"schemaAgencyName,attr,omitempty"`
	SchemaVersionID           string `json:",omitempty" xml:"schemaVersionID,attr,omitempty"`
	SchemaDataURI             string `json:",omitempty" xml:"schemaDataURI,attr,omitempty"`
	SchemaURI                 string `json:",omitempty" xml:"schemaURI,attr,omitempty"`
	UnitCode                  string `json:",omitempty" xml:"unitCode,attr,omitempty"`
	UnitCodeListID            string `json:",omitempty" xml:"unitCodeListID,attr,omitempty"`
	UnitCodeListAgencyID      string `json:",omitempty" xml:"unitCodeListAgencyID,attr,omitempty"`
	UnitCodeListAgencyName    string `json:",omitempty" xml:"unitCodeListAgencyName,attr,omitempty"`
	UnitCodeListVersionID     string `json:",omitempty" xml:"unitCodeListVersionID,attr,omitempty"`
	Filename                  string `json:",omitempty" xml:"filename,attr,omitempty"`
	Uri                       string `json:",omitempty" xml:"uri,attr,omitempty"`
	Value                     string `json:",omitempty" xml:",chardata"`
}

// QuantityStringType was auto-generated. Do not change.
type QuantityStringType struct {
	AnyGenericValueType
}

// TestResultType was auto-generated. Do not change.
type TestResultType struct {
	ID             IdentifierType     `json:",omitempty" xml:",omitempty"`
	Description    []DescriptionType  `json:",omitempty" xml:",omitempty"`
	TestDateTime   TestDateTimeType   `json:",omitempty" xml:",omitempty"`
	Result         []ResultType       `json:",omitempty" xml:",omitempty"`
	ExpirationTime ExpirationTimeType `json:",omitempty" xml:",omitempty"`
}

// ResultType was auto-generated. Do not change.
type ResultType struct {
	ValueString   ValueStringType   `json:",omitempty" xml:",omitempty"`
	DataType      DataTypeType      `json:",omitempty" xml:",omitempty"`
	UnitOfMeasure UnitOfMeasureType `json:",omitempty" xml:",omitempty"`
	Key           IdentifierType    `json:",omitempty" xml:",omitempty"`
}

// EquipmentClassType was auto-generated. Do not change.
type EquipmentClassType struct {
	ID                                     IdentifierType                               `json:",omitempty" xml:",omitempty"`
	Description                            []DescriptionType                            `json:",omitempty" xml:",omitempty"`
	Location                               LocationType                                 `json:",omitempty" xml:",omitempty"`
	HierarchyScope                         HierarchyScopeType                           `json:",omitempty" xml:",omitempty"`
	EquipmentLevel                         HierarchyScopeType                           `json:",omitempty" xml:",omitempty"`
	EquipmentClassProperty                 []EquipmentClassPropertyType                 `json:",omitempty" xml:",omitempty"`
	EquipmentID                            []EquipmentIDType                            `json:",omitempty" xml:",omitempty"`
	EquipmentCapabilityTestSpecificationID []EquipmentCapabilityTestSpecificationIDType `json:",omitempty" xml:",omitempty"`
}

// EquipmentClassPropertyType was auto-generated. Do not change.
type EquipmentClassPropertyType struct {
	ID                                     IdentifierType                               `json:",omitempty" xml:",omitempty"`
	Description                            []DescriptionType                            `json:",omitempty" xml:",omitempty"`
	Value                                  []ValueType                                  `json:",omitempty" xml:",omitempty"`
	EquipmentClassProperty                 []EquipmentClassPropertyType                 `json:",omitempty" xml:",omitempty"`
	EquipmentCapabilityTestSpecificationID []EquipmentCapabilityTestSpecificationIDType `json:",omitempty" xml:",omitempty"`
}

// EquipmentCapabilityTestSpecificationType was auto-generated. Do not change.
type EquipmentCapabilityTestSpecificationType struct {
	Name                         NameType                           `json:",omitempty" xml:",omitempty"`
	Description                  []DescriptionType                  `json:",omitempty" xml:",omitempty"`
	Version                      VersionType                        `json:",omitempty" xml:",omitempty"`
	Location                     LocationType                       `json:",omitempty" xml:",omitempty"`
	HierarchyScope               HierarchyScopeType                 `json:",omitempty" xml:",omitempty"`
	TestedEquipmentProperty      []TestedEquipmentPropertyType      `json:",omitempty" xml:",omitempty"`
	TestedEquipmentClassProperty []TestedEquipmentClassPropertyType `json:",omitempty" xml:",omitempty"`
}

// NameType was auto-generated. Do not change.
type NameType struct {
	LanguageID string `json:",omitempty" xml:"languageID,attr,omitempty"`
	Value      string `json:",omitempty" xml:",chardata"`
}

// ApprovedByType was auto-generated. Do not change.
type ApprovedByType struct {
	NameType
}

// TestedEquipmentPropertyType was auto-generated. Do not change.
type TestedEquipmentPropertyType struct {
	EquipmentID EquipmentIDType `json:",omitempty" xml:",omitempty"`
	PropertyID  PropertyIDType  `json:",omitempty" xml:",omitempty"`
}

// TestedEquipmentClassPropertyType was auto-generated. Do not change.
type TestedEquipmentClassPropertyType struct {
	EquipmentClassID EquipmentClassIDType `json:",omitempty" xml:",omitempty"`
	PropertyID       PropertyIDType       `json:",omitempty" xml:",omitempty"`
}

// GetEquipmentInformationType was auto-generated. Do not change.
type GetEquipmentInformationType struct {
	ApplicationArea TransApplicationAreaType            `json:",omitempty" xml:",omitempty"`
	DataArea        GetEquipmentInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                              `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                              `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// GetEquipmentInformationTypeDataArea was auto-generated. Do not change.
type GetEquipmentInformationTypeDataArea struct {
	Get                  []string                   `json:",omitempty" xml:",omitempty"`
	EquipmentInformation []EquipmentInformationType `json:",omitempty" xml:",omitempty"`
}

// ShowEquipmentInformationType was auto-generated. Do not change.
type ShowEquipmentInformationType struct {
	ApplicationArea TransApplicationAreaType             `json:",omitempty" xml:",omitempty"`
	DataArea        ShowEquipmentInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                               `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                               `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ShowEquipmentInformationTypeDataArea was auto-generated. Do not change.
type ShowEquipmentInformationTypeDataArea struct {
	Show                 TransShowType              `json:",omitempty" xml:",omitempty"`
	EquipmentInformation []EquipmentInformationType `json:",omitempty" xml:",omitempty"`
}

// TransShowType was auto-generated. Do not change.
type TransShowType struct {
	OriginalApplicationArea TransApplicationAreaType    `json:",omitempty" xml:",omitempty"`
	ResponseCriteria        []TransResponseCriteriaType `json:",omitempty" xml:",omitempty"`
}

// ProcessEquipmentInformationType was auto-generated. Do not change.
type ProcessEquipmentInformationType struct {
	ApplicationArea TransApplicationAreaType                `json:",omitempty" xml:",omitempty"`
	DataArea        ProcessEquipmentInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                  `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                  `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ProcessEquipmentInformationTypeDataArea was auto-generated. Do not change.
type ProcessEquipmentInformationTypeDataArea struct {
	Process              TransProcessType           `json:",omitempty" xml:",omitempty"`
	EquipmentInformation []EquipmentInformationType `json:",omitempty" xml:",omitempty"`
}

// TransProcessType was auto-generated. Do not change.
type TransProcessType struct {
	ActionCriteria           []TransActionCriteriaType `json:",omitempty" xml:",omitempty"`
	AcknowledgeCode          TransResponseCodeType     `json:",omitempty" xml:"acknowledgeCode,attr,omitempty"`
	AcknowledgeCodeSpecified bool                      `json:",omitempty" xml:",omitempty"`
}

// TransActionCriteriaType was auto-generated. Do not change.
type TransActionCriteriaType struct {
	ActionExpression []TransExpressionType `json:",omitempty" xml:",omitempty"`
	ChangeStatus     TransChangeStatusType `json:",omitempty" xml:",omitempty"`
}

// AcknowledgeEquipmentInformationType was auto-generated. Do not change.
type AcknowledgeEquipmentInformationType struct {
	ApplicationArea TransApplicationAreaType                    `json:",omitempty" xml:",omitempty"`
	DataArea        AcknowledgeEquipmentInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                      `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                      `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// AcknowledgeEquipmentInformationTypeDataArea was auto-generated. Do not change.
type AcknowledgeEquipmentInformationTypeDataArea struct {
	Acknowledge          TransAcknowledgeType       `json:",omitempty" xml:",omitempty"`
	EquipmentInformation []EquipmentInformationType `json:",omitempty" xml:",omitempty"`
}

// TransAcknowledgeType was auto-generated. Do not change.
type TransAcknowledgeType struct {
	OriginalApplicationArea TransApplicationAreaType    `json:",omitempty" xml:",omitempty"`
	ResponseCriteria        []TransResponseCriteriaType `json:",omitempty" xml:",omitempty"`
}

// ChangeEquipmentInformationType was auto-generated. Do not change.
type ChangeEquipmentInformationType struct {
	ApplicationArea TransApplicationAreaType               `json:",omitempty" xml:",omitempty"`
	DataArea        ChangeEquipmentInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                 `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                 `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ChangeEquipmentInformationTypeDataArea was auto-generated. Do not change.
type ChangeEquipmentInformationTypeDataArea struct {
	Change               TransChangeType            `json:",omitempty" xml:",omitempty"`
	EquipmentInformation []EquipmentInformationType `json:",omitempty" xml:",omitempty"`
}

// TransChangeType was auto-generated. Do not change.
type TransChangeType struct {
	ActionCriteria        []TransActionCriteriaType `json:",omitempty" xml:",omitempty"`
	ResponseCode          TransResponseCodeType     `json:",omitempty" xml:"responseCode,attr,omitempty"`
	ResponseCodeSpecified bool                      `json:",omitempty" xml:",omitempty"`
}

// RespondEquipmentInformationType was auto-generated. Do not change.
type RespondEquipmentInformationType struct {
	ApplicationArea TransApplicationAreaType                `json:",omitempty" xml:",omitempty"`
	DataArea        RespondEquipmentInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                  `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                  `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// RespondEquipmentInformationTypeDataArea was auto-generated. Do not change.
type RespondEquipmentInformationTypeDataArea struct {
	Respond              TransRespondType           `json:",omitempty" xml:",omitempty"`
	EquipmentInformation []EquipmentInformationType `json:",omitempty" xml:",omitempty"`
}

// TransRespondType was auto-generated. Do not change.
type TransRespondType struct {
	OriginalApplicationArea TransApplicationAreaType    `json:",omitempty" xml:",omitempty"`
	ResponseCriteria        []TransResponseCriteriaType `json:",omitempty" xml:",omitempty"`
}

// CancelEquipmentInformationType was auto-generated. Do not change.
type CancelEquipmentInformationType struct {
	ApplicationArea TransApplicationAreaType               `json:",omitempty" xml:",omitempty"`
	DataArea        CancelEquipmentInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                 `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                 `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// CancelEquipmentInformationTypeDataArea was auto-generated. Do not change.
type CancelEquipmentInformationTypeDataArea struct {
	Cancel               []TransActionCriteriaType  `json:",omitempty" xml:",omitempty"`
	EquipmentInformation []EquipmentInformationType `json:",omitempty" xml:",omitempty"`
}

// SyncEquipmentInformationType was auto-generated. Do not change.
type SyncEquipmentInformationType struct {
	ApplicationArea TransApplicationAreaType             `json:",omitempty" xml:",omitempty"`
	DataArea        SyncEquipmentInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                               `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                               `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// SyncEquipmentInformationTypeDataArea was auto-generated. Do not change.
type SyncEquipmentInformationTypeDataArea struct {
	Sync                 []TransActionCriteriaType  `json:",omitempty" xml:",omitempty"`
	EquipmentInformation []EquipmentInformationType `json:",omitempty" xml:",omitempty"`
}

// GetEquipmentClassType was auto-generated. Do not change.
type GetEquipmentClassType struct {
	ApplicationArea TransApplicationAreaType      `json:",omitempty" xml:",omitempty"`
	DataArea        GetEquipmentClassTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                        `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                        `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// GetEquipmentClassTypeDataArea was auto-generated. Do not change.
type GetEquipmentClassTypeDataArea struct {
	Get            []string             `json:",omitempty" xml:",omitempty"`
	EquipmentClass []EquipmentClassType `json:",omitempty" xml:",omitempty"`
}

// ShowEquipmentClassType was auto-generated. Do not change.
type ShowEquipmentClassType struct {
	ApplicationArea TransApplicationAreaType       `json:",omitempty" xml:",omitempty"`
	DataArea        ShowEquipmentClassTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                         `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                         `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ShowEquipmentClassTypeDataArea was auto-generated. Do not change.
type ShowEquipmentClassTypeDataArea struct {
	Show           TransShowType        `json:",omitempty" xml:",omitempty"`
	EquipmentClass []EquipmentClassType `json:",omitempty" xml:",omitempty"`
}

// ProcessEquipmentClassType was auto-generated. Do not change.
type ProcessEquipmentClassType struct {
	ApplicationArea TransApplicationAreaType          `json:",omitempty" xml:",omitempty"`
	DataArea        ProcessEquipmentClassTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                            `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                            `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ProcessEquipmentClassTypeDataArea was auto-generated. Do not change.
type ProcessEquipmentClassTypeDataArea struct {
	Process        TransProcessType     `json:",omitempty" xml:",omitempty"`
	EquipmentClass []EquipmentClassType `json:",omitempty" xml:",omitempty"`
}

// AcknowledgeEquipmentClassType was auto-generated. Do not change.
type AcknowledgeEquipmentClassType struct {
	ApplicationArea TransApplicationAreaType              `json:",omitempty" xml:",omitempty"`
	DataArea        AcknowledgeEquipmentClassTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// AcknowledgeEquipmentClassTypeDataArea was auto-generated. Do not change.
type AcknowledgeEquipmentClassTypeDataArea struct {
	Acknowledge    TransAcknowledgeType `json:",omitempty" xml:",omitempty"`
	EquipmentClass []EquipmentClassType `json:",omitempty" xml:",omitempty"`
}

// ChangeEquipmentClassType was auto-generated. Do not change.
type ChangeEquipmentClassType struct {
	ApplicationArea TransApplicationAreaType         `json:",omitempty" xml:",omitempty"`
	DataArea        ChangeEquipmentClassTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                           `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                           `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ChangeEquipmentClassTypeDataArea was auto-generated. Do not change.
type ChangeEquipmentClassTypeDataArea struct {
	Change         TransChangeType      `json:",omitempty" xml:",omitempty"`
	EquipmentClass []EquipmentClassType `json:",omitempty" xml:",omitempty"`
}

// RespondEquipmentClassType was auto-generated. Do not change.
type RespondEquipmentClassType struct {
	ApplicationArea TransApplicationAreaType          `json:",omitempty" xml:",omitempty"`
	DataArea        RespondEquipmentClassTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                            `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                            `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// RespondEquipmentClassTypeDataArea was auto-generated. Do not change.
type RespondEquipmentClassTypeDataArea struct {
	Respond        TransRespondType     `json:",omitempty" xml:",omitempty"`
	EquipmentClass []EquipmentClassType `json:",omitempty" xml:",omitempty"`
}

// CancelEquipmentClassType was auto-generated. Do not change.
type CancelEquipmentClassType struct {
	ApplicationArea TransApplicationAreaType         `json:",omitempty" xml:",omitempty"`
	DataArea        CancelEquipmentClassTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                           `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                           `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// CancelEquipmentClassTypeDataArea was auto-generated. Do not change.
type CancelEquipmentClassTypeDataArea struct {
	Cancel         []TransActionCriteriaType `json:",omitempty" xml:",omitempty"`
	EquipmentClass []EquipmentClassType      `json:",omitempty" xml:",omitempty"`
}

// SyncEquipmentClassType was auto-generated. Do not change.
type SyncEquipmentClassType struct {
	ApplicationArea TransApplicationAreaType       `json:",omitempty" xml:",omitempty"`
	DataArea        SyncEquipmentClassTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                         `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                         `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// SyncEquipmentClassTypeDataArea was auto-generated. Do not change.
type SyncEquipmentClassTypeDataArea struct {
	Sync           []TransActionCriteriaType `json:",omitempty" xml:",omitempty"`
	EquipmentClass []EquipmentClassType      `json:",omitempty" xml:",omitempty"`
}

// GetEquipmentType was auto-generated. Do not change.
type GetEquipmentType struct {
	ApplicationArea TransApplicationAreaType `json:",omitempty" xml:",omitempty"`
	DataArea        GetEquipmentTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                   `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                   `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// GetEquipmentTypeDataArea was auto-generated. Do not change.
type GetEquipmentTypeDataArea struct {
	Get       []string        `json:",omitempty" xml:",omitempty"`
	Equipment []EquipmentType `json:",omitempty" xml:",omitempty"`
}

// ShowEquipmentType was auto-generated. Do not change.
type ShowEquipmentType struct {
	ApplicationArea TransApplicationAreaType  `json:",omitempty" xml:",omitempty"`
	DataArea        ShowEquipmentTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                    `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                    `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ShowEquipmentTypeDataArea was auto-generated. Do not change.
type ShowEquipmentTypeDataArea struct {
	Show      TransShowType   `json:",omitempty" xml:",omitempty"`
	Equipment []EquipmentType `json:",omitempty" xml:",omitempty"`
}

// ProcessEquipmentType was auto-generated. Do not change.
type ProcessEquipmentType struct {
	ApplicationArea TransApplicationAreaType     `json:",omitempty" xml:",omitempty"`
	DataArea        ProcessEquipmentTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                       `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                       `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ProcessEquipmentTypeDataArea was auto-generated. Do not change.
type ProcessEquipmentTypeDataArea struct {
	Process   TransProcessType `json:",omitempty" xml:",omitempty"`
	Equipment []EquipmentType  `json:",omitempty" xml:",omitempty"`
}

// AcknowledgeEquipmentType was auto-generated. Do not change.
type AcknowledgeEquipmentType struct {
	ApplicationArea TransApplicationAreaType         `json:",omitempty" xml:",omitempty"`
	DataArea        AcknowledgeEquipmentTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                           `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                           `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// AcknowledgeEquipmentTypeDataArea was auto-generated. Do not change.
type AcknowledgeEquipmentTypeDataArea struct {
	Acknowledge TransAcknowledgeType `json:",omitempty" xml:",omitempty"`
	Equipment   []EquipmentType      `json:",omitempty" xml:",omitempty"`
}

// ChangeEquipmentType was auto-generated. Do not change.
type ChangeEquipmentType struct {
	ApplicationArea TransApplicationAreaType    `json:",omitempty" xml:",omitempty"`
	DataArea        ChangeEquipmentTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                      `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                      `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ChangeEquipmentTypeDataArea was auto-generated. Do not change.
type ChangeEquipmentTypeDataArea struct {
	Change    TransChangeType `json:",omitempty" xml:",omitempty"`
	Equipment []EquipmentType `json:",omitempty" xml:",omitempty"`
}

// RespondEquipmentType was auto-generated. Do not change.
type RespondEquipmentType struct {
	ApplicationArea TransApplicationAreaType     `json:",omitempty" xml:",omitempty"`
	DataArea        RespondEquipmentTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                       `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                       `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// RespondEquipmentTypeDataArea was auto-generated. Do not change.
type RespondEquipmentTypeDataArea struct {
	Respond   TransRespondType `json:",omitempty" xml:",omitempty"`
	Equipment []EquipmentType  `json:",omitempty" xml:",omitempty"`
}

// CancelEquipmentType was auto-generated. Do not change.
type CancelEquipmentType struct {
	ApplicationArea TransApplicationAreaType    `json:",omitempty" xml:",omitempty"`
	DataArea        CancelEquipmentTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                      `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                      `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// CancelEquipmentTypeDataArea was auto-generated. Do not change.
type CancelEquipmentTypeDataArea struct {
	Cancel    []TransActionCriteriaType `json:",omitempty" xml:",omitempty"`
	Equipment []EquipmentType           `json:",omitempty" xml:",omitempty"`
}

// SyncEquipmentType was auto-generated. Do not change.
type SyncEquipmentType struct {
	ApplicationArea TransApplicationAreaType  `json:",omitempty" xml:",omitempty"`
	DataArea        SyncEquipmentTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                    `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                    `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// SyncEquipmentTypeDataArea was auto-generated. Do not change.
type SyncEquipmentTypeDataArea struct {
	Sync      []TransActionCriteriaType `json:",omitempty" xml:",omitempty"`
	Equipment []EquipmentType           `json:",omitempty" xml:",omitempty"`
}

// GetEquipmentCapabilityTestSpecType was auto-generated. Do not change.
type GetEquipmentCapabilityTestSpecType struct {
	ApplicationArea TransApplicationAreaType                   `json:",omitempty" xml:",omitempty"`
	DataArea        GetEquipmentCapabilityTestSpecTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                     `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                     `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// GetEquipmentCapabilityTestSpecTypeDataArea was auto-generated. Do not change.
type GetEquipmentCapabilityTestSpecTypeDataArea struct {
	Get                         []string                                   `json:",omitempty" xml:",omitempty"`
	EquipmentCapabilityTestSpec []EquipmentCapabilityTestSpecificationType `json:",omitempty" xml:",omitempty"`
}

// ShowEquipmentCapabilityTestSpecType was auto-generated. Do not change.
type ShowEquipmentCapabilityTestSpecType struct {
	ApplicationArea TransApplicationAreaType                    `json:",omitempty" xml:",omitempty"`
	DataArea        ShowEquipmentCapabilityTestSpecTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                      `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                      `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ShowEquipmentCapabilityTestSpecTypeDataArea was auto-generated. Do not change.
type ShowEquipmentCapabilityTestSpecTypeDataArea struct {
	Show                        TransShowType                              `json:",omitempty" xml:",omitempty"`
	EquipmentCapabilityTestSpec []EquipmentCapabilityTestSpecificationType `json:",omitempty" xml:",omitempty"`
}

// ProcessEquipmentCapabilityTestSpecType was auto-generated. Do not change.
type ProcessEquipmentCapabilityTestSpecType struct {
	ApplicationArea TransApplicationAreaType                       `json:",omitempty" xml:",omitempty"`
	DataArea        ProcessEquipmentCapabilityTestSpecTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                         `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                         `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ProcessEquipmentCapabilityTestSpecTypeDataArea was auto-generated. Do not change.
type ProcessEquipmentCapabilityTestSpecTypeDataArea struct {
	Process                     TransProcessType                           `json:",omitempty" xml:",omitempty"`
	EquipmentCapabilityTestSpec []EquipmentCapabilityTestSpecificationType `json:",omitempty" xml:",omitempty"`
}

// AcknowledgeEquipmentCapabilityTestSpecType was auto-generated. Do not change.
type AcknowledgeEquipmentCapabilityTestSpecType struct {
	ApplicationArea TransApplicationAreaType                           `json:",omitempty" xml:",omitempty"`
	DataArea        AcknowledgeEquipmentCapabilityTestSpecTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                             `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                             `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// AcknowledgeEquipmentCapabilityTestSpecTypeDataArea was auto-generated. Do not change.
type AcknowledgeEquipmentCapabilityTestSpecTypeDataArea struct {
	Acknowledge                 TransAcknowledgeType                       `json:",omitempty" xml:",omitempty"`
	EquipmentCapabilityTestSpec []EquipmentCapabilityTestSpecificationType `json:",omitempty" xml:",omitempty"`
}

// ChangeEquipmentCapabilityTestSpecType was auto-generated. Do not change.
type ChangeEquipmentCapabilityTestSpecType struct {
	ApplicationArea TransApplicationAreaType                      `json:",omitempty" xml:",omitempty"`
	DataArea        ChangeEquipmentCapabilityTestSpecTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                        `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                        `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ChangeEquipmentCapabilityTestSpecTypeDataArea was auto-generated. Do not change.
type ChangeEquipmentCapabilityTestSpecTypeDataArea struct {
	Change                      TransChangeType                            `json:",omitempty" xml:",omitempty"`
	EquipmentCapabilityTestSpec []EquipmentCapabilityTestSpecificationType `json:",omitempty" xml:",omitempty"`
}

// RespondEquipmentCapabilityTestSpecType was auto-generated. Do not change.
type RespondEquipmentCapabilityTestSpecType struct {
	ApplicationArea TransApplicationAreaType                       `json:",omitempty" xml:",omitempty"`
	DataArea        RespondEquipmentCapabilityTestSpecTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                         `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                         `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// RespondEquipmentCapabilityTestSpecTypeDataArea was auto-generated. Do not change.
type RespondEquipmentCapabilityTestSpecTypeDataArea struct {
	Respond                     TransRespondType                           `json:",omitempty" xml:",omitempty"`
	EquipmentCapabilityTestSpec []EquipmentCapabilityTestSpecificationType `json:",omitempty" xml:",omitempty"`
}

// CancelEquipmentCapabilityTestSpecType was auto-generated. Do not change.
type CancelEquipmentCapabilityTestSpecType struct {
	ApplicationArea TransApplicationAreaType                      `json:",omitempty" xml:",omitempty"`
	DataArea        CancelEquipmentCapabilityTestSpecTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                        `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                        `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// CancelEquipmentCapabilityTestSpecTypeDataArea was auto-generated. Do not change.
type CancelEquipmentCapabilityTestSpecTypeDataArea struct {
	Cancel                      []TransActionCriteriaType                  `json:",omitempty" xml:",omitempty"`
	EquipmentCapabilityTestSpec []EquipmentCapabilityTestSpecificationType `json:",omitempty" xml:",omitempty"`
}

// SyncEquipmentCapabilityTestSpecType was auto-generated. Do not change.
type SyncEquipmentCapabilityTestSpecType struct {
	ApplicationArea TransApplicationAreaType                    `json:",omitempty" xml:",omitempty"`
	DataArea        SyncEquipmentCapabilityTestSpecTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                      `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                      `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// SyncEquipmentCapabilityTestSpecTypeDataArea was auto-generated. Do not change.
type SyncEquipmentCapabilityTestSpecTypeDataArea struct {
	Sync                        []TransActionCriteriaType                  `json:",omitempty" xml:",omitempty"`
	EquipmentCapabilityTestSpec []EquipmentCapabilityTestSpecificationType `json:",omitempty" xml:",omitempty"`
}

// MaterialInformationType was auto-generated. Do not change.
type MaterialInformationType struct {
	ID                        IdentifierType                  `json:",omitempty" xml:",omitempty"`
	Description               []DescriptionType               `json:",omitempty" xml:",omitempty"`
	Location                  LocationType                    `json:",omitempty" xml:",omitempty"`
	HierarchyScope            HierarchyScopeType              `json:",omitempty" xml:",omitempty"`
	PublishedDate             PublishedDateType               `json:",omitempty" xml:",omitempty"`
	MaterialClass             []MaterialClassType             `json:",omitempty" xml:",omitempty"`
	MaterialDefinition        []MaterialDefinitionType        `json:",omitempty" xml:",omitempty"`
	MaterialLot               []MaterialLotType               `json:",omitempty" xml:",omitempty"`
	MaterialSubLot            []MaterialSubLotType            `json:",omitempty" xml:",omitempty"`
	MaterialTestSpecification []MaterialTestSpecificationType `json:",omitempty" xml:",omitempty"`
}

// MaterialClassType was auto-generated. Do not change.
type MaterialClassType struct {
	ID                          IdentifierType                    `json:",omitempty" xml:",omitempty"`
	Description                 []DescriptionType                 `json:",omitempty" xml:",omitempty"`
	Location                    LocationType                      `json:",omitempty" xml:",omitempty"`
	HierarchyScope              HierarchyScopeType                `json:",omitempty" xml:",omitempty"`
	MaterialClassProperty       []MaterialClassPropertyType       `json:",omitempty" xml:",omitempty"`
	MaterialDefinitionID        []MaterialDefinitionIDType        `json:",omitempty" xml:",omitempty"`
	MaterialTestSpecificationID []MaterialTestSpecificationIDType `json:",omitempty" xml:",omitempty"`
	AssemblyClassID             []MaterialClassIDType             `json:",omitempty" xml:",omitempty"`
	AssemblyType                AssemblyTypeType                  `json:",omitempty" xml:",omitempty"`
	AssemblyRelationship        AssemblyRelationshipType          `json:",omitempty" xml:",omitempty"`
}

// MaterialClassPropertyType was auto-generated. Do not change.
type MaterialClassPropertyType struct {
	ID                          IdentifierType                    `json:",omitempty" xml:",omitempty"`
	Description                 []DescriptionType                 `json:",omitempty" xml:",omitempty"`
	Value                       []ValueType                       `json:",omitempty" xml:",omitempty"`
	MaterialClassProperty       []MaterialClassPropertyType       `json:",omitempty" xml:",omitempty"`
	MaterialTestSpecificationID []MaterialTestSpecificationIDType `json:",omitempty" xml:",omitempty"`
}

// MaterialDefinitionType was auto-generated. Do not change.
type MaterialDefinitionType struct {
	ID                          IdentifierType                    `json:",omitempty" xml:",omitempty"`
	Description                 []DescriptionType                 `json:",omitempty" xml:",omitempty"`
	Location                    LocationType                      `json:",omitempty" xml:",omitempty"`
	HierarchyScope              HierarchyScopeType                `json:",omitempty" xml:",omitempty"`
	MaterialDefinitionProperty  []MaterialDefinitionPropertyType  `json:",omitempty" xml:",omitempty"`
	MaterialClassID             []MaterialClassIDType             `json:",omitempty" xml:",omitempty"`
	MaterialLotID               []MaterialLotIDType               `json:",omitempty" xml:",omitempty"`
	MaterialTestSpecificationID []MaterialTestSpecificationIDType `json:",omitempty" xml:",omitempty"`
	AssemblylDefinitionID       []MaterialDefinitionIDType        `json:",omitempty" xml:",omitempty"`
	AssemblyType                AssemblyTypeType                  `json:",omitempty" xml:",omitempty"`
	AssemblyRelationship        AssemblyRelationshipType          `json:",omitempty" xml:",omitempty"`
}

// MaterialDefinitionPropertyType was auto-generated. Do not change.
type MaterialDefinitionPropertyType struct {
	ID                          IdentifierType                    `json:",omitempty" xml:",omitempty"`
	Description                 []DescriptionType                 `json:",omitempty" xml:",omitempty"`
	Value                       []ValueType                       `json:",omitempty" xml:",omitempty"`
	MaterialDefinitionProperty  []MaterialDefinitionPropertyType  `json:",omitempty" xml:",omitempty"`
	MaterialTestSpecificationID []MaterialTestSpecificationIDType `json:",omitempty" xml:",omitempty"`
}

// MaterialLotType was auto-generated. Do not change.
type MaterialLotType struct {
	ID                          IdentifierType                    `json:",omitempty" xml:",omitempty"`
	Description                 []DescriptionType                 `json:",omitempty" xml:",omitempty"`
	Location                    LocationType                      `json:",omitempty" xml:",omitempty"`
	HierarchyScope              HierarchyScopeType                `json:",omitempty" xml:",omitempty"`
	MaterialDefinitionID        MaterialDefinitionIDType          `json:",omitempty" xml:",omitempty"`
	Status                      StatusType                        `json:",omitempty" xml:",omitempty"`
	MaterialLotProperty         []MaterialLotPropertyType         `json:",omitempty" xml:",omitempty"`
	MaterialSubLot              []MaterialSubLotType              `json:",omitempty" xml:",omitempty"`
	StorageLocation             StorageHierarchyScopeType         `json:",omitempty" xml:",omitempty"`
	Quantity                    []QuantityValueType               `json:",omitempty" xml:",omitempty"`
	MaterialTestSpecificationID []MaterialTestSpecificationIDType `json:",omitempty" xml:",omitempty"`
	AssemblyLotID               []MaterialLotType                 `json:",omitempty" xml:",omitempty"`
	AssemblySubLotID            []MaterialSubLotType              `json:",omitempty" xml:",omitempty"`
	AssemblyType                AssemblyTypeType                  `json:",omitempty" xml:",omitempty"`
	AssemblyRelationship        AssemblyRelationshipType          `json:",omitempty" xml:",omitempty"`
}

// MaterialLotPropertyType was auto-generated. Do not change.
type MaterialLotPropertyType struct {
	ID                          IdentifierType                    `json:",omitempty" xml:",omitempty"`
	Description                 []DescriptionType                 `json:",omitempty" xml:",omitempty"`
	Value                       []ValueType                       `json:",omitempty" xml:",omitempty"`
	MaterialLotProperty         []MaterialLotPropertyType         `json:",omitempty" xml:",omitempty"`
	MaterialTestSpecificationID []MaterialTestSpecificationIDType `json:",omitempty" xml:",omitempty"`
	TestResult                  []TestResultType                  `json:",omitempty" xml:",omitempty"`
}

// MaterialSubLotType was auto-generated. Do not change.
type MaterialSubLotType struct {
	ID                     IdentifierType            `json:",omitempty" xml:",omitempty"`
	Description            []DescriptionType         `json:",omitempty" xml:",omitempty"`
	Location               LocationType              `json:",omitempty" xml:",omitempty"`
	HierarchyScope         HierarchyScopeType        `json:",omitempty" xml:",omitempty"`
	Status                 StatusType                `json:",omitempty" xml:",omitempty"`
	MaterialSublotProperty []MaterialLotPropertyType `json:",omitempty" xml:",omitempty"`
	StorageLocation        StorageHierarchyScopeType `json:",omitempty" xml:",omitempty"`
	Quantity               []QuantityValueType       `json:",omitempty" xml:",omitempty"`
	MaterialSubLot         []MaterialSubLotType      `json:",omitempty" xml:",omitempty"`
	MaterialLotID          MaterialLotIDType         `json:",omitempty" xml:",omitempty"`
	AssemblyLotID          []MaterialLotType         `json:",omitempty" xml:",omitempty"`
	AssemblySubLotID       []MaterialSubLotType      `json:",omitempty" xml:",omitempty"`
	AssemblyType           AssemblyTypeType          `json:",omitempty" xml:",omitempty"`
	AssemblyRelationship   AssemblyRelationshipType  `json:",omitempty" xml:",omitempty"`
}

// QuantityValueType was auto-generated. Do not change.
type QuantityValueType struct {
	QuantityString QuantityStringType `json:",omitempty" xml:",omitempty"`
	DataType       DataTypeType       `json:",omitempty" xml:",omitempty"`
	UnitOfMeasure  UnitOfMeasureType  `json:",omitempty" xml:",omitempty"`
	Key            IdentifierType     `json:",omitempty" xml:",omitempty"`
}

// MaterialTestSpecificationType was auto-generated. Do not change.
type MaterialTestSpecificationType struct {
	Name                             NameType                               `json:",omitempty" xml:",omitempty"`
	Description                      []DescriptionType                      `json:",omitempty" xml:",omitempty"`
	Version                          VersionType                            `json:",omitempty" xml:",omitempty"`
	Location                         LocationType                           `json:",omitempty" xml:",omitempty"`
	HierarchyScope                   HierarchyScopeType                     `json:",omitempty" xml:",omitempty"`
	TestedMaterialClassProperty      []TestedMaterialClassPropertyType      `json:",omitempty" xml:",omitempty"`
	TestedMaterialDefinitionProperty []TestedMaterialDefinitionPropertyType `json:",omitempty" xml:",omitempty"`
	TestedMaterialLotProperty        []TestedMaterialLotPropertyType        `json:",omitempty" xml:",omitempty"`
}

// TestedMaterialClassPropertyType was auto-generated. Do not change.
type TestedMaterialClassPropertyType struct {
	MaterialClassID MaterialClassIDType `json:",omitempty" xml:",omitempty"`
	PropertyID      PropertyIDType      `json:",omitempty" xml:",omitempty"`
}

// TestedMaterialDefinitionPropertyType was auto-generated. Do not change.
type TestedMaterialDefinitionPropertyType struct {
	MaterialDefinitionID MaterialDefinitionIDType `json:",omitempty" xml:",omitempty"`
	PropertyID           PropertyIDType           `json:",omitempty" xml:",omitempty"`
}

// TestedMaterialLotPropertyType was auto-generated. Do not change.
type TestedMaterialLotPropertyType struct {
	MaterialLotID MaterialLotIDType `json:",omitempty" xml:",omitempty"`
	PropertyID    PropertyIDType    `json:",omitempty" xml:",omitempty"`
}

// GetMaterialInformationType was auto-generated. Do not change.
type GetMaterialInformationType struct {
	ApplicationArea TransApplicationAreaType           `json:",omitempty" xml:",omitempty"`
	DataArea        GetMaterialInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                             `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                             `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// GetMaterialInformationTypeDataArea was auto-generated. Do not change.
type GetMaterialInformationTypeDataArea struct {
	Get                 []string                  `json:",omitempty" xml:",omitempty"`
	MaterialInformation []MaterialInformationType `json:",omitempty" xml:",omitempty"`
}

// ShowMaterialInformationType was auto-generated. Do not change.
type ShowMaterialInformationType struct {
	ApplicationArea TransApplicationAreaType            `json:",omitempty" xml:",omitempty"`
	DataArea        ShowMaterialInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                              `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                              `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ShowMaterialInformationTypeDataArea was auto-generated. Do not change.
type ShowMaterialInformationTypeDataArea struct {
	Show                TransShowType             `json:",omitempty" xml:",omitempty"`
	MaterialInformation []MaterialInformationType `json:",omitempty" xml:",omitempty"`
}

// ProcessMaterialInformationType was auto-generated. Do not change.
type ProcessMaterialInformationType struct {
	ApplicationArea TransApplicationAreaType               `json:",omitempty" xml:",omitempty"`
	DataArea        ProcessMaterialInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                 `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                 `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ProcessMaterialInformationTypeDataArea was auto-generated. Do not change.
type ProcessMaterialInformationTypeDataArea struct {
	Process             TransProcessType          `json:",omitempty" xml:",omitempty"`
	MaterialInformation []MaterialInformationType `json:",omitempty" xml:",omitempty"`
}

// AcknowledgeMaterialInformationType was auto-generated. Do not change.
type AcknowledgeMaterialInformationType struct {
	ApplicationArea TransApplicationAreaType                   `json:",omitempty" xml:",omitempty"`
	DataArea        AcknowledgeMaterialInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                     `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                     `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// AcknowledgeMaterialInformationTypeDataArea was auto-generated. Do not change.
type AcknowledgeMaterialInformationTypeDataArea struct {
	Acknowledge         TransAcknowledgeType      `json:",omitempty" xml:",omitempty"`
	MaterialInformation []MaterialInformationType `json:",omitempty" xml:",omitempty"`
}

// ChangeMaterialInformationType was auto-generated. Do not change.
type ChangeMaterialInformationType struct {
	ApplicationArea TransApplicationAreaType              `json:",omitempty" xml:",omitempty"`
	DataArea        ChangeMaterialInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ChangeMaterialInformationTypeDataArea was auto-generated. Do not change.
type ChangeMaterialInformationTypeDataArea struct {
	Change              TransChangeType           `json:",omitempty" xml:",omitempty"`
	MaterialInformation []MaterialInformationType `json:",omitempty" xml:",omitempty"`
}

// RespondMaterialInformationType was auto-generated. Do not change.
type RespondMaterialInformationType struct {
	ApplicationArea TransApplicationAreaType               `json:",omitempty" xml:",omitempty"`
	DataArea        RespondMaterialInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                 `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                 `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// RespondMaterialInformationTypeDataArea was auto-generated. Do not change.
type RespondMaterialInformationTypeDataArea struct {
	Respond             TransRespondType          `json:",omitempty" xml:",omitempty"`
	MaterialInformation []MaterialInformationType `json:",omitempty" xml:",omitempty"`
}

// CancelMaterialInformationType was auto-generated. Do not change.
type CancelMaterialInformationType struct {
	ApplicationArea TransApplicationAreaType              `json:",omitempty" xml:",omitempty"`
	DataArea        CancelMaterialInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// CancelMaterialInformationTypeDataArea was auto-generated. Do not change.
type CancelMaterialInformationTypeDataArea struct {
	Cancel              []TransActionCriteriaType `json:",omitempty" xml:",omitempty"`
	MaterialInformation []MaterialInformationType `json:",omitempty" xml:",omitempty"`
}

// SyncMaterialInformationType was auto-generated. Do not change.
type SyncMaterialInformationType struct {
	ApplicationArea TransApplicationAreaType            `json:",omitempty" xml:",omitempty"`
	DataArea        SyncMaterialInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                              `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                              `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// SyncMaterialInformationTypeDataArea was auto-generated. Do not change.
type SyncMaterialInformationTypeDataArea struct {
	Sync                []TransActionCriteriaType `json:",omitempty" xml:",omitempty"`
	MaterialInformation []MaterialInformationType `json:",omitempty" xml:",omitempty"`
}

// GetMaterialClassType was auto-generated. Do not change.
type GetMaterialClassType struct {
	ApplicationArea TransApplicationAreaType     `json:",omitempty" xml:",omitempty"`
	DataArea        GetMaterialClassTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                       `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                       `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// GetMaterialClassTypeDataArea was auto-generated. Do not change.
type GetMaterialClassTypeDataArea struct {
	Get           []string            `json:",omitempty" xml:",omitempty"`
	MaterialClass []MaterialClassType `json:",omitempty" xml:",omitempty"`
}

// ShowMaterialClassType was auto-generated. Do not change.
type ShowMaterialClassType struct {
	ApplicationArea TransApplicationAreaType      `json:",omitempty" xml:",omitempty"`
	DataArea        ShowMaterialClassTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                        `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                        `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ShowMaterialClassTypeDataArea was auto-generated. Do not change.
type ShowMaterialClassTypeDataArea struct {
	Show          TransShowType       `json:",omitempty" xml:",omitempty"`
	MaterialClass []MaterialClassType `json:",omitempty" xml:",omitempty"`
}

// ProcessMaterialClassType was auto-generated. Do not change.
type ProcessMaterialClassType struct {
	ApplicationArea TransApplicationAreaType         `json:",omitempty" xml:",omitempty"`
	DataArea        ProcessMaterialClassTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                           `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                           `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ProcessMaterialClassTypeDataArea was auto-generated. Do not change.
type ProcessMaterialClassTypeDataArea struct {
	Process       TransProcessType    `json:",omitempty" xml:",omitempty"`
	MaterialClass []MaterialClassType `json:",omitempty" xml:",omitempty"`
}

// AcknowledgeMaterialClassType was auto-generated. Do not change.
type AcknowledgeMaterialClassType struct {
	ApplicationArea TransApplicationAreaType             `json:",omitempty" xml:",omitempty"`
	DataArea        AcknowledgeMaterialClassTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                               `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                               `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// AcknowledgeMaterialClassTypeDataArea was auto-generated. Do not change.
type AcknowledgeMaterialClassTypeDataArea struct {
	Acknowledge   TransAcknowledgeType `json:",omitempty" xml:",omitempty"`
	MaterialClass []MaterialClassType  `json:",omitempty" xml:",omitempty"`
}

// ChangeMaterialClassType was auto-generated. Do not change.
type ChangeMaterialClassType struct {
	ApplicationArea TransApplicationAreaType        `json:",omitempty" xml:",omitempty"`
	DataArea        ChangeMaterialClassTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                          `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                          `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ChangeMaterialClassTypeDataArea was auto-generated. Do not change.
type ChangeMaterialClassTypeDataArea struct {
	Change        TransChangeType     `json:",omitempty" xml:",omitempty"`
	MaterialClass []MaterialClassType `json:",omitempty" xml:",omitempty"`
}

// RespondMaterialClassType was auto-generated. Do not change.
type RespondMaterialClassType struct {
	ApplicationArea TransApplicationAreaType         `json:",omitempty" xml:",omitempty"`
	DataArea        RespondMaterialClassTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                           `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                           `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// RespondMaterialClassTypeDataArea was auto-generated. Do not change.
type RespondMaterialClassTypeDataArea struct {
	Respond       TransRespondType    `json:",omitempty" xml:",omitempty"`
	MaterialClass []MaterialClassType `json:",omitempty" xml:",omitempty"`
}

// CancelMaterialClassType was auto-generated. Do not change.
type CancelMaterialClassType struct {
	ApplicationArea TransApplicationAreaType        `json:",omitempty" xml:",omitempty"`
	DataArea        CancelMaterialClassTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                          `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                          `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// CancelMaterialClassTypeDataArea was auto-generated. Do not change.
type CancelMaterialClassTypeDataArea struct {
	Cancel        []TransActionCriteriaType `json:",omitempty" xml:",omitempty"`
	MaterialClass []MaterialClassType       `json:",omitempty" xml:",omitempty"`
}

// SyncMaterialClassType was auto-generated. Do not change.
type SyncMaterialClassType struct {
	ApplicationArea TransApplicationAreaType      `json:",omitempty" xml:",omitempty"`
	DataArea        SyncMaterialClassTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                        `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                        `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// SyncMaterialClassTypeDataArea was auto-generated. Do not change.
type SyncMaterialClassTypeDataArea struct {
	Sync          []TransActionCriteriaType `json:",omitempty" xml:",omitempty"`
	MaterialClass []MaterialClassType       `json:",omitempty" xml:",omitempty"`
}

// GetMaterialDefinitionType was auto-generated. Do not change.
type GetMaterialDefinitionType struct {
	ApplicationArea TransApplicationAreaType          `json:",omitempty" xml:",omitempty"`
	DataArea        GetMaterialDefinitionTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                            `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                            `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// GetMaterialDefinitionTypeDataArea was auto-generated. Do not change.
type GetMaterialDefinitionTypeDataArea struct {
	Get                []string                 `json:",omitempty" xml:",omitempty"`
	MaterialDefinition []MaterialDefinitionType `json:",omitempty" xml:",omitempty"`
}

// ShowMaterialDefinitionType was auto-generated. Do not change.
type ShowMaterialDefinitionType struct {
	ApplicationArea TransApplicationAreaType           `json:",omitempty" xml:",omitempty"`
	DataArea        ShowMaterialDefinitionTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                             `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                             `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ShowMaterialDefinitionTypeDataArea was auto-generated. Do not change.
type ShowMaterialDefinitionTypeDataArea struct {
	Show               TransShowType            `json:",omitempty" xml:",omitempty"`
	MaterialDefinition []MaterialDefinitionType `json:",omitempty" xml:",omitempty"`
}

// ProcessMaterialDefinitionType was auto-generated. Do not change.
type ProcessMaterialDefinitionType struct {
	ApplicationArea TransApplicationAreaType              `json:",omitempty" xml:",omitempty"`
	DataArea        ProcessMaterialDefinitionTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ProcessMaterialDefinitionTypeDataArea was auto-generated. Do not change.
type ProcessMaterialDefinitionTypeDataArea struct {
	Process            TransProcessType         `json:",omitempty" xml:",omitempty"`
	MaterialDefinition []MaterialDefinitionType `json:",omitempty" xml:",omitempty"`
}

// AcknowledgeMaterialDefinitionType was auto-generated. Do not change.
type AcknowledgeMaterialDefinitionType struct {
	ApplicationArea TransApplicationAreaType                  `json:",omitempty" xml:",omitempty"`
	DataArea        AcknowledgeMaterialDefinitionTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                    `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                    `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// AcknowledgeMaterialDefinitionTypeDataArea was auto-generated. Do not change.
type AcknowledgeMaterialDefinitionTypeDataArea struct {
	Acknowledge        TransAcknowledgeType     `json:",omitempty" xml:",omitempty"`
	MaterialDefinition []MaterialDefinitionType `json:",omitempty" xml:",omitempty"`
}

// ChangeMaterialDefinitionType was auto-generated. Do not change.
type ChangeMaterialDefinitionType struct {
	ApplicationArea TransApplicationAreaType             `json:",omitempty" xml:",omitempty"`
	DataArea        ChangeMaterialDefinitionTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                               `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                               `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ChangeMaterialDefinitionTypeDataArea was auto-generated. Do not change.
type ChangeMaterialDefinitionTypeDataArea struct {
	Change             TransChangeType          `json:",omitempty" xml:",omitempty"`
	MaterialDefinition []MaterialDefinitionType `json:",omitempty" xml:",omitempty"`
}

// RespondMaterialDefinitionType was auto-generated. Do not change.
type RespondMaterialDefinitionType struct {
	ApplicationArea TransApplicationAreaType              `json:",omitempty" xml:",omitempty"`
	DataArea        RespondMaterialDefinitionTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// RespondMaterialDefinitionTypeDataArea was auto-generated. Do not change.
type RespondMaterialDefinitionTypeDataArea struct {
	Respond            TransRespondType         `json:",omitempty" xml:",omitempty"`
	MaterialDefinition []MaterialDefinitionType `json:",omitempty" xml:",omitempty"`
}

// CancelMaterialDefinitionType was auto-generated. Do not change.
type CancelMaterialDefinitionType struct {
	ApplicationArea TransApplicationAreaType             `json:",omitempty" xml:",omitempty"`
	DataArea        CancelMaterialDefinitionTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                               `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                               `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// CancelMaterialDefinitionTypeDataArea was auto-generated. Do not change.
type CancelMaterialDefinitionTypeDataArea struct {
	Cancel             []TransActionCriteriaType `json:",omitempty" xml:",omitempty"`
	MaterialDefinition []MaterialDefinitionType  `json:",omitempty" xml:",omitempty"`
}

// SyncMaterialDefinitionType was auto-generated. Do not change.
type SyncMaterialDefinitionType struct {
	ApplicationArea TransApplicationAreaType           `json:",omitempty" xml:",omitempty"`
	DataArea        SyncMaterialDefinitionTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                             `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                             `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// SyncMaterialDefinitionTypeDataArea was auto-generated. Do not change.
type SyncMaterialDefinitionTypeDataArea struct {
	Sync               []TransActionCriteriaType `json:",omitempty" xml:",omitempty"`
	MaterialDefinition []MaterialDefinitionType  `json:",omitempty" xml:",omitempty"`
}

// GetMaterialLotType was auto-generated. Do not change.
type GetMaterialLotType struct {
	ApplicationArea TransApplicationAreaType   `json:",omitempty" xml:",omitempty"`
	DataArea        GetMaterialLotTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                     `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                     `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// GetMaterialLotTypeDataArea was auto-generated. Do not change.
type GetMaterialLotTypeDataArea struct {
	Get         []string          `json:",omitempty" xml:",omitempty"`
	MaterialLot []MaterialLotType `json:",omitempty" xml:",omitempty"`
}

// ShowMaterialLotType was auto-generated. Do not change.
type ShowMaterialLotType struct {
	ApplicationArea TransApplicationAreaType    `json:",omitempty" xml:",omitempty"`
	DataArea        ShowMaterialLotTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                      `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                      `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ShowMaterialLotTypeDataArea was auto-generated. Do not change.
type ShowMaterialLotTypeDataArea struct {
	Show        TransShowType     `json:",omitempty" xml:",omitempty"`
	MaterialLot []MaterialLotType `json:",omitempty" xml:",omitempty"`
}

// ProcessMaterialLotType was auto-generated. Do not change.
type ProcessMaterialLotType struct {
	ApplicationArea TransApplicationAreaType       `json:",omitempty" xml:",omitempty"`
	DataArea        ProcessMaterialLotTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                         `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                         `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ProcessMaterialLotTypeDataArea was auto-generated. Do not change.
type ProcessMaterialLotTypeDataArea struct {
	Process     TransProcessType  `json:",omitempty" xml:",omitempty"`
	MaterialLot []MaterialLotType `json:",omitempty" xml:",omitempty"`
}

// AcknowledgeMaterialLotType was auto-generated. Do not change.
type AcknowledgeMaterialLotType struct {
	ApplicationArea TransApplicationAreaType           `json:",omitempty" xml:",omitempty"`
	DataArea        AcknowledgeMaterialLotTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                             `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                             `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// AcknowledgeMaterialLotTypeDataArea was auto-generated. Do not change.
type AcknowledgeMaterialLotTypeDataArea struct {
	Acknowledge TransAcknowledgeType `json:",omitempty" xml:",omitempty"`
	MaterialLot []MaterialLotType    `json:",omitempty" xml:",omitempty"`
}

// ChangeMaterialLotType was auto-generated. Do not change.
type ChangeMaterialLotType struct {
	ApplicationArea TransApplicationAreaType      `json:",omitempty" xml:",omitempty"`
	DataArea        ChangeMaterialLotTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                        `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                        `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ChangeMaterialLotTypeDataArea was auto-generated. Do not change.
type ChangeMaterialLotTypeDataArea struct {
	Change      TransChangeType   `json:",omitempty" xml:",omitempty"`
	MaterialLot []MaterialLotType `json:",omitempty" xml:",omitempty"`
}

// RespondMaterialLotType was auto-generated. Do not change.
type RespondMaterialLotType struct {
	ApplicationArea TransApplicationAreaType       `json:",omitempty" xml:",omitempty"`
	DataArea        RespondMaterialLotTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                         `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                         `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// RespondMaterialLotTypeDataArea was auto-generated. Do not change.
type RespondMaterialLotTypeDataArea struct {
	Respond     TransRespondType  `json:",omitempty" xml:",omitempty"`
	MaterialLot []MaterialLotType `json:",omitempty" xml:",omitempty"`
}

// CancelMaterialLotType was auto-generated. Do not change.
type CancelMaterialLotType struct {
	ApplicationArea TransApplicationAreaType      `json:",omitempty" xml:",omitempty"`
	DataArea        CancelMaterialLotTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                        `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                        `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// CancelMaterialLotTypeDataArea was auto-generated. Do not change.
type CancelMaterialLotTypeDataArea struct {
	Cancel      []TransActionCriteriaType `json:",omitempty" xml:",omitempty"`
	MaterialLot []MaterialLotType         `json:",omitempty" xml:",omitempty"`
}

// SyncMaterialLotType was auto-generated. Do not change.
type SyncMaterialLotType struct {
	ApplicationArea TransApplicationAreaType    `json:",omitempty" xml:",omitempty"`
	DataArea        SyncMaterialLotTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                      `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                      `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// SyncMaterialLotTypeDataArea was auto-generated. Do not change.
type SyncMaterialLotTypeDataArea struct {
	Sync        []TransActionCriteriaType `json:",omitempty" xml:",omitempty"`
	MaterialLot []MaterialLotType         `json:",omitempty" xml:",omitempty"`
}

// GetMaterialSubLotType was auto-generated. Do not change.
type GetMaterialSubLotType struct {
	ApplicationArea TransApplicationAreaType      `json:",omitempty" xml:",omitempty"`
	DataArea        GetMaterialSubLotTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                        `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                        `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// GetMaterialSubLotTypeDataArea was auto-generated. Do not change.
type GetMaterialSubLotTypeDataArea struct {
	Get            []string             `json:",omitempty" xml:",omitempty"`
	MaterialSubLot []MaterialSubLotType `json:",omitempty" xml:",omitempty"`
}

// ShowMaterialSubLotType was auto-generated. Do not change.
type ShowMaterialSubLotType struct {
	ApplicationArea TransApplicationAreaType       `json:",omitempty" xml:",omitempty"`
	DataArea        ShowMaterialSubLotTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                         `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                         `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ShowMaterialSubLotTypeDataArea was auto-generated. Do not change.
type ShowMaterialSubLotTypeDataArea struct {
	Show           TransShowType        `json:",omitempty" xml:",omitempty"`
	MaterialSubLot []MaterialSubLotType `json:",omitempty" xml:",omitempty"`
}

// ProcessMaterialSubLotType was auto-generated. Do not change.
type ProcessMaterialSubLotType struct {
	ApplicationArea TransApplicationAreaType          `json:",omitempty" xml:",omitempty"`
	DataArea        ProcessMaterialSubLotTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                            `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                            `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ProcessMaterialSubLotTypeDataArea was auto-generated. Do not change.
type ProcessMaterialSubLotTypeDataArea struct {
	Process        TransProcessType     `json:",omitempty" xml:",omitempty"`
	MaterialSubLot []MaterialSubLotType `json:",omitempty" xml:",omitempty"`
}

// AcknowledgeMaterialSubLotType was auto-generated. Do not change.
type AcknowledgeMaterialSubLotType struct {
	ApplicationArea TransApplicationAreaType              `json:",omitempty" xml:",omitempty"`
	DataArea        AcknowledgeMaterialSubLotTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// AcknowledgeMaterialSubLotTypeDataArea was auto-generated. Do not change.
type AcknowledgeMaterialSubLotTypeDataArea struct {
	Acknowledge    TransAcknowledgeType `json:",omitempty" xml:",omitempty"`
	MaterialSubLot []MaterialSubLotType `json:",omitempty" xml:",omitempty"`
}

// ChangeMaterialSubLotType was auto-generated. Do not change.
type ChangeMaterialSubLotType struct {
	ApplicationArea TransApplicationAreaType         `json:",omitempty" xml:",omitempty"`
	DataArea        ChangeMaterialSubLotTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                           `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                           `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ChangeMaterialSubLotTypeDataArea was auto-generated. Do not change.
type ChangeMaterialSubLotTypeDataArea struct {
	Change         TransChangeType      `json:",omitempty" xml:",omitempty"`
	MaterialSubLot []MaterialSubLotType `json:",omitempty" xml:",omitempty"`
}

// RespondMaterialSubLotType was auto-generated. Do not change.
type RespondMaterialSubLotType struct {
	ApplicationArea TransApplicationAreaType          `json:",omitempty" xml:",omitempty"`
	DataArea        RespondMaterialSubLotTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                            `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                            `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// RespondMaterialSubLotTypeDataArea was auto-generated. Do not change.
type RespondMaterialSubLotTypeDataArea struct {
	Respond        TransRespondType     `json:",omitempty" xml:",omitempty"`
	MaterialSubLot []MaterialSubLotType `json:",omitempty" xml:",omitempty"`
}

// CancelMaterialSubLotType was auto-generated. Do not change.
type CancelMaterialSubLotType struct {
	ApplicationArea TransApplicationAreaType         `json:",omitempty" xml:",omitempty"`
	DataArea        CancelMaterialSubLotTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                           `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                           `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// CancelMaterialSubLotTypeDataArea was auto-generated. Do not change.
type CancelMaterialSubLotTypeDataArea struct {
	Cancel         []TransActionCriteriaType `json:",omitempty" xml:",omitempty"`
	MaterialSubLot []MaterialSubLotType      `json:",omitempty" xml:",omitempty"`
}

// SyncMaterialSubLotType was auto-generated. Do not change.
type SyncMaterialSubLotType struct {
	ApplicationArea TransApplicationAreaType       `json:",omitempty" xml:",omitempty"`
	DataArea        SyncMaterialSubLotTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                         `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                         `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// SyncMaterialSubLotTypeDataArea was auto-generated. Do not change.
type SyncMaterialSubLotTypeDataArea struct {
	Sync           []TransActionCriteriaType `json:",omitempty" xml:",omitempty"`
	MaterialSubLot []MaterialSubLotType      `json:",omitempty" xml:",omitempty"`
}

// GetMaterialTestSpecType was auto-generated. Do not change.
type GetMaterialTestSpecType struct {
	ApplicationArea TransApplicationAreaType        `json:",omitempty" xml:",omitempty"`
	DataArea        GetMaterialTestSpecTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                          `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                          `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// GetMaterialTestSpecTypeDataArea was auto-generated. Do not change.
type GetMaterialTestSpecTypeDataArea struct {
	Get              []string                        `json:",omitempty" xml:",omitempty"`
	MaterialTestSpec []MaterialTestSpecificationType `json:",omitempty" xml:",omitempty"`
}

// ShowMaterialTestSpecType was auto-generated. Do not change.
type ShowMaterialTestSpecType struct {
	ApplicationArea TransApplicationAreaType         `json:",omitempty" xml:",omitempty"`
	DataArea        ShowMaterialTestSpecTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                           `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                           `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ShowMaterialTestSpecTypeDataArea was auto-generated. Do not change.
type ShowMaterialTestSpecTypeDataArea struct {
	Show             TransShowType                   `json:",omitempty" xml:",omitempty"`
	MaterialTestSpec []MaterialTestSpecificationType `json:",omitempty" xml:",omitempty"`
}

// ProcessMaterialTestSpecType was auto-generated. Do not change.
type ProcessMaterialTestSpecType struct {
	ApplicationArea TransApplicationAreaType            `json:",omitempty" xml:",omitempty"`
	DataArea        ProcessMaterialTestSpecTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                              `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                              `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ProcessMaterialTestSpecTypeDataArea was auto-generated. Do not change.
type ProcessMaterialTestSpecTypeDataArea struct {
	Process          TransProcessType                `json:",omitempty" xml:",omitempty"`
	MaterialTestSpec []MaterialTestSpecificationType `json:",omitempty" xml:",omitempty"`
}

// AcknowledgeMaterialTestSpecType was auto-generated. Do not change.
type AcknowledgeMaterialTestSpecType struct {
	ApplicationArea TransApplicationAreaType                `json:",omitempty" xml:",omitempty"`
	DataArea        AcknowledgeMaterialTestSpecTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                  `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                  `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// AcknowledgeMaterialTestSpecTypeDataArea was auto-generated. Do not change.
type AcknowledgeMaterialTestSpecTypeDataArea struct {
	Acknowledge      TransAcknowledgeType            `json:",omitempty" xml:",omitempty"`
	MaterialTestSpec []MaterialTestSpecificationType `json:",omitempty" xml:",omitempty"`
}

// ChangeMaterialTestSpecType was auto-generated. Do not change.
type ChangeMaterialTestSpecType struct {
	ApplicationArea TransApplicationAreaType           `json:",omitempty" xml:",omitempty"`
	DataArea        ChangeMaterialTestSpecTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                             `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                             `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ChangeMaterialTestSpecTypeDataArea was auto-generated. Do not change.
type ChangeMaterialTestSpecTypeDataArea struct {
	Change           TransChangeType                 `json:",omitempty" xml:",omitempty"`
	MaterialTestSpec []MaterialTestSpecificationType `json:",omitempty" xml:",omitempty"`
}

// RespondMaterialTestSpecType was auto-generated. Do not change.
type RespondMaterialTestSpecType struct {
	ApplicationArea TransApplicationAreaType            `json:",omitempty" xml:",omitempty"`
	DataArea        RespondMaterialTestSpecTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                              `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                              `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// RespondMaterialTestSpecTypeDataArea was auto-generated. Do not change.
type RespondMaterialTestSpecTypeDataArea struct {
	Respond          TransRespondType                `json:",omitempty" xml:",omitempty"`
	MaterialTestSpec []MaterialTestSpecificationType `json:",omitempty" xml:",omitempty"`
}

// CancelMaterialTestSpecType was auto-generated. Do not change.
type CancelMaterialTestSpecType struct {
	ApplicationArea TransApplicationAreaType           `json:",omitempty" xml:",omitempty"`
	DataArea        CancelMaterialTestSpecTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                             `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                             `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// CancelMaterialTestSpecTypeDataArea was auto-generated. Do not change.
type CancelMaterialTestSpecTypeDataArea struct {
	Cancel           []TransActionCriteriaType       `json:",omitempty" xml:",omitempty"`
	MaterialTestSpec []MaterialTestSpecificationType `json:",omitempty" xml:",omitempty"`
}

// SyncMaterialTestSpecType was auto-generated. Do not change.
type SyncMaterialTestSpecType struct {
	ApplicationArea TransApplicationAreaType         `json:",omitempty" xml:",omitempty"`
	DataArea        SyncMaterialTestSpecTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                           `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                           `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// SyncMaterialTestSpecTypeDataArea was auto-generated. Do not change.
type SyncMaterialTestSpecTypeDataArea struct {
	Sync             []TransActionCriteriaType       `json:",omitempty" xml:",omitempty"`
	MaterialTestSpec []MaterialTestSpecificationType `json:",omitempty" xml:",omitempty"`
}

// OperationsCapabilityInformationType was auto-generated. Do not change.
type OperationsCapabilityInformationType struct {
	ID                   IdentifierType             `json:",omitempty" xml:",omitempty"`
	Description          []DescriptionType          `json:",omitempty" xml:",omitempty"`
	HierarchyScope       HierarchyScopeType         `json:",omitempty" xml:",omitempty"`
	PublishedDate        PublishedDateType          `json:",omitempty" xml:",omitempty"`
	OperationsCapability []OperationsCapabilityType `json:",omitempty" xml:",omitempty"`
}

// OperationsCapabilityType was auto-generated. Do not change.
type OperationsCapabilityType struct {
	ID                       IdentifierType                   `json:",omitempty" xml:",omitempty"`
	Description              []DescriptionType                `json:",omitempty" xml:",omitempty"`
	HierarchyScope           HierarchyScopeType               `json:",omitempty" xml:",omitempty"`
	CapabilityType           CapabilityTypeType               `json:",omitempty" xml:",omitempty"`
	Reason                   ReasonType                       `json:",omitempty" xml:",omitempty"`
	ConfidenceFactor         ConfidenceFactorType             `json:",omitempty" xml:",omitempty"`
	StartTime                StartTimeType                    `json:",omitempty" xml:",omitempty"`
	EndTime                  EndTimeType                      `json:",omitempty" xml:",omitempty"`
	PublishedDate            PublishedDateType                `json:",omitempty" xml:",omitempty"`
	PersonnelCapability      []OpPersonnelCapabilityType      `json:",omitempty" xml:",omitempty"`
	EquipmentCapability      []OpEquipmentCapabilityType      `json:",omitempty" xml:",omitempty"`
	PhysicalAssetCapability  []OpPhysicalAssetCapabilityType  `json:",omitempty" xml:",omitempty"`
	MaterialCapability       []OpMaterialCapabilityType       `json:",omitempty" xml:",omitempty"`
	ProcessSegmentCapability []OpProcessSegmentCapabilityType `json:",omitempty" xml:",omitempty"`
}

// OpPersonnelCapabilityType was auto-generated. Do not change.
type OpPersonnelCapabilityType struct {
	PersonnelClassID            []PersonnelClassIDType              `json:",omitempty" xml:",omitempty"`
	PersonID                    []PersonIDType                      `json:",omitempty" xml:",omitempty"`
	Description                 []DescriptionType                   `json:",omitempty" xml:",omitempty"`
	CapabilityType              CapabilityTypeType                  `json:",omitempty" xml:",omitempty"`
	Reason                      ReasonType                          `json:",omitempty" xml:",omitempty"`
	ConfidenceFactor            ConfidenceFactorType                `json:",omitempty" xml:",omitempty"`
	HierarchyScope              HierarchyScopeType                  `json:",omitempty" xml:",omitempty"`
	PersonnelUse                PersonnelUseType                    `json:",omitempty" xml:",omitempty"`
	StartTime                   StartTimeType                       `json:",omitempty" xml:",omitempty"`
	EndTime                     EndTimeType                         `json:",omitempty" xml:",omitempty"`
	Quantity                    []QuantityValueType                 `json:",omitempty" xml:",omitempty"`
	PersonnelCapabilityProperty []OpPersonnelCapabilityPropertyType `json:",omitempty" xml:",omitempty"`
}

// OpPersonnelCapabilityPropertyType was auto-generated. Do not change.
type OpPersonnelCapabilityPropertyType struct {
	ID          IdentifierType      `json:",omitempty" xml:",omitempty"`
	Description []DescriptionType   `json:",omitempty" xml:",omitempty"`
	Value       []ValueType         `json:",omitempty" xml:",omitempty"`
	Quantity    []QuantityValueType `json:",omitempty" xml:",omitempty"`
}

// OpEquipmentCapabilityType was auto-generated. Do not change.
type OpEquipmentCapabilityType struct {
	EquipmentClassID            []EquipmentClassIDType              `json:",omitempty" xml:",omitempty"`
	EquipmentID                 []EquipmentIDType                   `json:",omitempty" xml:",omitempty"`
	Description                 []DescriptionType                   `json:",omitempty" xml:",omitempty"`
	CapabilityType              CapabilityTypeType                  `json:",omitempty" xml:",omitempty"`
	Reason                      ReasonType                          `json:",omitempty" xml:",omitempty"`
	ConfidenceFactor            ConfidenceFactorType                `json:",omitempty" xml:",omitempty"`
	HierarchyScope              HierarchyScopeType                  `json:",omitempty" xml:",omitempty"`
	EquipmentUse                EquipmentUseType                    `json:",omitempty" xml:",omitempty"`
	StartTime                   StartTimeType                       `json:",omitempty" xml:",omitempty"`
	EndTime                     EndTimeType                         `json:",omitempty" xml:",omitempty"`
	Quantity                    []QuantityValueType                 `json:",omitempty" xml:",omitempty"`
	EquipmentCapabilityProperty []OpEquipmentCapabilityPropertyType `json:",omitempty" xml:",omitempty"`
}

// OpEquipmentCapabilityPropertyType was auto-generated. Do not change.
type OpEquipmentCapabilityPropertyType struct {
	ID          IdentifierType      `json:",omitempty" xml:",omitempty"`
	Description []DescriptionType   `json:",omitempty" xml:",omitempty"`
	Value       []ValueType         `json:",omitempty" xml:",omitempty"`
	Quantity    []QuantityValueType `json:",omitempty" xml:",omitempty"`
}

// OpPhysicalAssetCapabilityType was auto-generated. Do not change.
type OpPhysicalAssetCapabilityType struct {
	PhysicalAssetClassID            []PhysicalAssetClassIDType              `json:",omitempty" xml:",omitempty"`
	PhysicalAssetID                 []PhysicalAssetIDType                   `json:",omitempty" xml:",omitempty"`
	Description                     []DescriptionType                       `json:",omitempty" xml:",omitempty"`
	CapabilityType                  CapabilityTypeType                      `json:",omitempty" xml:",omitempty"`
	Reason                          ReasonType                              `json:",omitempty" xml:",omitempty"`
	ConfidenceFactor                ConfidenceFactorType                    `json:",omitempty" xml:",omitempty"`
	HierarchyScope                  HierarchyScopeType                      `json:",omitempty" xml:",omitempty"`
	PhysicalAssetUse                PhysicalAssetUseType                    `json:",omitempty" xml:",omitempty"`
	StartTime                       StartTimeType                           `json:",omitempty" xml:",omitempty"`
	EndTime                         EndTimeType                             `json:",omitempty" xml:",omitempty"`
	Quantity                        []QuantityValueType                     `json:",omitempty" xml:",omitempty"`
	PhysicalAssetCapabilityProperty []OpPhysicalAssetCapabilityPropertyType `json:",omitempty" xml:",omitempty"`
}

// OpPhysicalAssetCapabilityPropertyType was auto-generated. Do not change.
type OpPhysicalAssetCapabilityPropertyType struct {
	ID          IdentifierType      `json:",omitempty" xml:",omitempty"`
	Description []DescriptionType   `json:",omitempty" xml:",omitempty"`
	Value       []ValueType         `json:",omitempty" xml:",omitempty"`
	Quantity    []QuantityValueType `json:",omitempty" xml:",omitempty"`
}

// OpMaterialCapabilityType was auto-generated. Do not change.
type OpMaterialCapabilityType struct {
	MaterialClassID            []MaterialClassIDType              `json:",omitempty" xml:",omitempty"`
	MaterialDefinitionID       []MaterialDefinitionIDType         `json:",omitempty" xml:",omitempty"`
	MaterialLotID              []MaterialLotIDType                `json:",omitempty" xml:",omitempty"`
	MaterialSubLotID           []MaterialSubLotIDType             `json:",omitempty" xml:",omitempty"`
	Description                []DescriptionType                  `json:",omitempty" xml:",omitempty"`
	CapabilityType             CapabilityTypeType                 `json:",omitempty" xml:",omitempty"`
	Reason                     ReasonType                         `json:",omitempty" xml:",omitempty"`
	ConfidenceFactor           ConfidenceFactorType               `json:",omitempty" xml:",omitempty"`
	HierarchyScope             HierarchyScopeType                 `json:",omitempty" xml:",omitempty"`
	MaterialUse                MaterialUseType                    `json:",omitempty" xml:",omitempty"`
	StartTime                  StartTimeType                      `json:",omitempty" xml:",omitempty"`
	EndTime                    EndTimeType                        `json:",omitempty" xml:",omitempty"`
	AssemblyCapability         []OpMaterialCapabilityType         `json:",omitempty" xml:",omitempty"`
	AssemblyType               AssemblyTypeType                   `json:",omitempty" xml:",omitempty"`
	AssemblyRelationship       AssemblyRelationshipType           `json:",omitempty" xml:",omitempty"`
	Quantity                   []QuantityValueType                `json:",omitempty" xml:",omitempty"`
	MaterialCapabilityProperty []OpMaterialCapabilityPropertyType `json:",omitempty" xml:",omitempty"`
}

// OpMaterialCapabilityPropertyType was auto-generated. Do not change.
type OpMaterialCapabilityPropertyType struct {
	ID          IdentifierType      `json:",omitempty" xml:",omitempty"`
	Description []DescriptionType   `json:",omitempty" xml:",omitempty"`
	Value       []ValueType         `json:",omitempty" xml:",omitempty"`
	Quantity    []QuantityValueType `json:",omitempty" xml:",omitempty"`
	MaterialUse MaterialUseType     `json:",omitempty" xml:",omitempty"`
}

// OpProcessSegmentCapabilityType was auto-generated. Do not change.
type OpProcessSegmentCapabilityType struct {
	ID                       IdentifierType                   `json:",omitempty" xml:",omitempty"`
	Description              []DescriptionType                `json:",omitempty" xml:",omitempty"`
	ProcessSegmentID         []ProcessSegmentIDType           `json:",omitempty" xml:",omitempty"`
	CapabilityType           CapabilityTypeType               `json:",omitempty" xml:",omitempty"`
	Reason                   []ReasonType                     `json:",omitempty" xml:",omitempty"`
	HierarchyScope           []HierarchyScopeType             `json:",omitempty" xml:",omitempty"`
	EquipmentElementLevel    []EquipmentElementLevelType      `json:",omitempty" xml:",omitempty"`
	StartTime                StartTimeType                    `json:",omitempty" xml:",omitempty"`
	EndTime                  EndTimeType                      `json:",omitempty" xml:",omitempty"`
	PersonnelCapability      []OpPersonnelCapabilityType      `json:",omitempty" xml:",omitempty"`
	EquipmentCapability      []OpEquipmentCapabilityType      `json:",omitempty" xml:",omitempty"`
	PhysicalAssetCapability  []OpPhysicalAssetCapabilityType  `json:",omitempty" xml:",omitempty"`
	MaterialCapability       []OpMaterialCapabilityType       `json:",omitempty" xml:",omitempty"`
	ProcessSegmentCapability []OpProcessSegmentCapabilityType `json:",omitempty" xml:",omitempty"`
}

// GetOperationsCapabilityType was auto-generated. Do not change.
type GetOperationsCapabilityType struct {
	ApplicationArea TransApplicationAreaType            `json:",omitempty" xml:",omitempty"`
	DataArea        GetOperationsCapabilityTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                              `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                              `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// GetOperationsCapabilityTypeDataArea was auto-generated. Do not change.
type GetOperationsCapabilityTypeDataArea struct {
	Get                  []string                   `json:",omitempty" xml:",omitempty"`
	OperationsCapability []OperationsCapabilityType `json:",omitempty" xml:",omitempty"`
}

// ShowOperationsCapabilityType was auto-generated. Do not change.
type ShowOperationsCapabilityType struct {
	ApplicationArea TransApplicationAreaType             `json:",omitempty" xml:",omitempty"`
	DataArea        ShowOperationsCapabilityTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                               `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                               `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ShowOperationsCapabilityTypeDataArea was auto-generated. Do not change.
type ShowOperationsCapabilityTypeDataArea struct {
	Show                 TransShowType              `json:",omitempty" xml:",omitempty"`
	OperationsCapability []OperationsCapabilityType `json:",omitempty" xml:",omitempty"`
}

// ProcessOperationsCapabilityType was auto-generated. Do not change.
type ProcessOperationsCapabilityType struct {
	ApplicationArea TransApplicationAreaType                `json:",omitempty" xml:",omitempty"`
	DataArea        ProcessOperationsCapabilityTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                  `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                  `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ProcessOperationsCapabilityTypeDataArea was auto-generated. Do not change.
type ProcessOperationsCapabilityTypeDataArea struct {
	Process              TransProcessType           `json:",omitempty" xml:",omitempty"`
	OperationsCapability []OperationsCapabilityType `json:",omitempty" xml:",omitempty"`
}

// AcknowledgeOperationsCapabilityType was auto-generated. Do not change.
type AcknowledgeOperationsCapabilityType struct {
	ApplicationArea TransApplicationAreaType                    `json:",omitempty" xml:",omitempty"`
	DataArea        AcknowledgeOperationsCapabilityTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                      `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                      `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// AcknowledgeOperationsCapabilityTypeDataArea was auto-generated. Do not change.
type AcknowledgeOperationsCapabilityTypeDataArea struct {
	Acknowledge          TransAcknowledgeType       `json:",omitempty" xml:",omitempty"`
	OperationsCapability []OperationsCapabilityType `json:",omitempty" xml:",omitempty"`
}

// ChangeOperationsCapabilityType was auto-generated. Do not change.
type ChangeOperationsCapabilityType struct {
	ApplicationArea TransApplicationAreaType               `json:",omitempty" xml:",omitempty"`
	DataArea        ChangeOperationsCapabilityTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                 `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                 `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ChangeOperationsCapabilityTypeDataArea was auto-generated. Do not change.
type ChangeOperationsCapabilityTypeDataArea struct {
	Change               TransChangeType            `json:",omitempty" xml:",omitempty"`
	OperationsCapability []OperationsCapabilityType `json:",omitempty" xml:",omitempty"`
}

// RespondOperationsCapabilityType was auto-generated. Do not change.
type RespondOperationsCapabilityType struct {
	ApplicationArea TransApplicationAreaType                `json:",omitempty" xml:",omitempty"`
	DataArea        RespondOperationsCapabilityTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                  `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                  `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// RespondOperationsCapabilityTypeDataArea was auto-generated. Do not change.
type RespondOperationsCapabilityTypeDataArea struct {
	Respond              TransRespondType           `json:",omitempty" xml:",omitempty"`
	OperationsCapability []OperationsCapabilityType `json:",omitempty" xml:",omitempty"`
}

// CancelOperationsCapabilityType was auto-generated. Do not change.
type CancelOperationsCapabilityType struct {
	ApplicationArea TransApplicationAreaType               `json:",omitempty" xml:",omitempty"`
	DataArea        CancelOperationsCapabilityTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                 `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                 `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// CancelOperationsCapabilityTypeDataArea was auto-generated. Do not change.
type CancelOperationsCapabilityTypeDataArea struct {
	Cancel               []TransActionCriteriaType  `json:",omitempty" xml:",omitempty"`
	OperationsCapability []OperationsCapabilityType `json:",omitempty" xml:",omitempty"`
}

// SyncOperationsCapabilityType was auto-generated. Do not change.
type SyncOperationsCapabilityType struct {
	ApplicationArea TransApplicationAreaType             `json:",omitempty" xml:",omitempty"`
	DataArea        SyncOperationsCapabilityTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                               `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                               `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// SyncOperationsCapabilityTypeDataArea was auto-generated. Do not change.
type SyncOperationsCapabilityTypeDataArea struct {
	Sync                 []TransActionCriteriaType  `json:",omitempty" xml:",omitempty"`
	OperationsCapability []OperationsCapabilityType `json:",omitempty" xml:",omitempty"`
}

// GetOperationsCapabilityInformationType was auto-generated. Do not change.
type GetOperationsCapabilityInformationType struct {
	ApplicationArea TransApplicationAreaType                       `json:",omitempty" xml:",omitempty"`
	DataArea        GetOperationsCapabilityInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                         `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                         `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// GetOperationsCapabilityInformationTypeDataArea was auto-generated. Do not change.
type GetOperationsCapabilityInformationTypeDataArea struct {
	Get                             []string                              `json:",omitempty" xml:",omitempty"`
	OperationsCapabilityInformation []OperationsCapabilityInformationType `json:",omitempty" xml:",omitempty"`
}

// ShowOperationsCapabilityInformationType was auto-generated. Do not change.
type ShowOperationsCapabilityInformationType struct {
	ApplicationArea TransApplicationAreaType                        `json:",omitempty" xml:",omitempty"`
	DataArea        ShowOperationsCapabilityInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                          `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                          `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ShowOperationsCapabilityInformationTypeDataArea was auto-generated. Do not change.
type ShowOperationsCapabilityInformationTypeDataArea struct {
	Show                            TransShowType                         `json:",omitempty" xml:",omitempty"`
	OperationsCapabilityInformation []OperationsCapabilityInformationType `json:",omitempty" xml:",omitempty"`
}

// ProcessOperationsCapabilityInformationType was auto-generated. Do not change.
type ProcessOperationsCapabilityInformationType struct {
	ApplicationArea TransApplicationAreaType                           `json:",omitempty" xml:",omitempty"`
	DataArea        ProcessOperationsCapabilityInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                             `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                             `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ProcessOperationsCapabilityInformationTypeDataArea was auto-generated. Do not change.
type ProcessOperationsCapabilityInformationTypeDataArea struct {
	Process                         TransProcessType                      `json:",omitempty" xml:",omitempty"`
	OperationsCapabilityInformation []OperationsCapabilityInformationType `json:",omitempty" xml:",omitempty"`
}

// AcknowledgeOperationsCapabilityInformationType was auto-generated. Do not change.
type AcknowledgeOperationsCapabilityInformationType struct {
	ApplicationArea TransApplicationAreaType                               `json:",omitempty" xml:",omitempty"`
	DataArea        AcknowledgeOperationsCapabilityInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                                 `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                                 `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// AcknowledgeOperationsCapabilityInformationTypeDataArea was auto-generated. Do not change.
type AcknowledgeOperationsCapabilityInformationTypeDataArea struct {
	Acknowledge                     TransAcknowledgeType                  `json:",omitempty" xml:",omitempty"`
	OperationsCapabilityInformation []OperationsCapabilityInformationType `json:",omitempty" xml:",omitempty"`
}

// ChangeOperationsCapabilityInformationType was auto-generated. Do not change.
type ChangeOperationsCapabilityInformationType struct {
	ApplicationArea TransApplicationAreaType                          `json:",omitempty" xml:",omitempty"`
	DataArea        ChangeOperationsCapabilityInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                            `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                            `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ChangeOperationsCapabilityInformationTypeDataArea was auto-generated. Do not change.
type ChangeOperationsCapabilityInformationTypeDataArea struct {
	Change                          TransChangeType                       `json:",omitempty" xml:",omitempty"`
	OperationsCapabilityInformation []OperationsCapabilityInformationType `json:",omitempty" xml:",omitempty"`
}

// RespondOperationsCapabilityInformationType was auto-generated. Do not change.
type RespondOperationsCapabilityInformationType struct {
	ApplicationArea TransApplicationAreaType                           `json:",omitempty" xml:",omitempty"`
	DataArea        RespondOperationsCapabilityInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                             `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                             `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// RespondOperationsCapabilityInformationTypeDataArea was auto-generated. Do not change.
type RespondOperationsCapabilityInformationTypeDataArea struct {
	Respond                         TransRespondType                      `json:",omitempty" xml:",omitempty"`
	OperationsCapabilityInformation []OperationsCapabilityInformationType `json:",omitempty" xml:",omitempty"`
}

// CancelOperationsCapabilityInformationType was auto-generated. Do not change.
type CancelOperationsCapabilityInformationType struct {
	ApplicationArea TransApplicationAreaType                          `json:",omitempty" xml:",omitempty"`
	DataArea        CancelOperationsCapabilityInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                            `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                            `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// CancelOperationsCapabilityInformationTypeDataArea was auto-generated. Do not change.
type CancelOperationsCapabilityInformationTypeDataArea struct {
	Cancel                          []TransActionCriteriaType             `json:",omitempty" xml:",omitempty"`
	OperationsCapabilityInformation []OperationsCapabilityInformationType `json:",omitempty" xml:",omitempty"`
}

// SyncOperationsCapabilityInformationType was auto-generated. Do not change.
type SyncOperationsCapabilityInformationType struct {
	ApplicationArea TransApplicationAreaType                        `json:",omitempty" xml:",omitempty"`
	DataArea        SyncOperationsCapabilityInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                          `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                          `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// SyncOperationsCapabilityInformationTypeDataArea was auto-generated. Do not change.
type SyncOperationsCapabilityInformationTypeDataArea struct {
	Sync                            []TransActionCriteriaType             `json:",omitempty" xml:",omitempty"`
	OperationsCapabilityInformation []OperationsCapabilityInformationType `json:",omitempty" xml:",omitempty"`
}

// OperationsDefinitionInformationType was auto-generated. Do not change.
type OperationsDefinitionInformationType struct {
	ID                   IdentifierType             `json:",omitempty" xml:",omitempty"`
	Description          []DescriptionType          `json:",omitempty" xml:",omitempty"`
	HierarchyScope       HierarchyScopeType         `json:",omitempty" xml:",omitempty"`
	PublishedDate        PublishedDateType          `json:",omitempty" xml:",omitempty"`
	OperationsDefinition []OperationsDefinitionType `json:",omitempty" xml:",omitempty"`
}

// OperationsDefinitionType was auto-generated. Do not change.
type OperationsDefinitionType struct {
	ID                     IdentifierType               `json:",omitempty" xml:",omitempty"`
	Version                VersionType                  `json:",omitempty" xml:",omitempty"`
	Description            []DescriptionType            `json:",omitempty" xml:",omitempty"`
	HierarchyScope         HierarchyScopeType           `json:",omitempty" xml:",omitempty"`
	OperationsType         OperationsTypeType           `json:",omitempty" xml:",omitempty"`
	PublishedDate          PublishedDateType            `json:",omitempty" xml:",omitempty"`
	BillOfMaterialsID      BillOfMaterialsIDType        `json:",omitempty" xml:",omitempty"`
	WorkDefinitionID       IdentifierType               `json:",omitempty" xml:",omitempty"`
	BillOfResourcesID      BillOfResourcesIDType        `json:",omitempty" xml:",omitempty"`
	OperationsMaterialBill []OperationsMaterialBillType `json:",omitempty" xml:",omitempty"`
	OperationsSegment      []OperationsSegmentType      `json:",omitempty" xml:",omitempty"`
}

// OperationsMaterialBillType was auto-generated. Do not change.
type OperationsMaterialBillType struct {
	ID                         IdentifierType                   `json:",omitempty" xml:",omitempty"`
	Description                DescriptionType                  `json:",omitempty" xml:",omitempty"`
	OperationsMaterialBillItem []OperationsMaterialBillItemType `json:",omitempty" xml:",omitempty"`
}

// OperationsMaterialBillItemType was auto-generated. Do not change.
type OperationsMaterialBillItemType struct {
	ID                         IdentifierType                   `json:",omitempty" xml:",omitempty"`
	Description                DescriptionType                  `json:",omitempty" xml:",omitempty"`
	MaterialClassID            []MaterialClassIDType            `json:",omitempty" xml:",omitempty"`
	MaterialDefinitionID       []MaterialDefinitionIDType       `json:",omitempty" xml:",omitempty"`
	UseType                    CodeType                         `json:",omitempty" xml:",omitempty"`
	AssemblyBillOfMaterialItem []OperationsMaterialBillItemType `json:",omitempty" xml:",omitempty"`
	AssemblyType               AssemblyTypeType                 `json:",omitempty" xml:",omitempty"`
	AssemblyRelationship       AssemblyRelationshipType         `json:",omitempty" xml:",omitempty"`
	MaterialSpecificationID    []IdentifierType                 `json:",omitempty" xml:",omitempty"`
	Quantity                   []QuantityValueType              `json:",omitempty" xml:",omitempty"`
}

// OperationsSegmentType was auto-generated. Do not change.
type OperationsSegmentType struct {
	ID                         IdentifierType                     `json:",omitempty" xml:",omitempty"`
	Description                DescriptionType                    `json:",omitempty" xml:",omitempty"`
	OperationsType             OperationsTypeType                 `json:",omitempty" xml:",omitempty"`
	HierarchyScope             HierarchyScopeType                 `json:",omitempty" xml:",omitempty"`
	Duration                   string                             `json:",omitempty" xml:",omitempty"`
	ProcessSegmentID           []ProcessSegmentIDType             `json:",omitempty" xml:",omitempty"`
	Parameter                  []ParameterType                    `json:",omitempty" xml:",omitempty"`
	PersonnelSpecification     []OpPersonnelSpecificationType     `json:",omitempty" xml:",omitempty"`
	EquipmentSpecification     []OpEquipmentSpecificationType     `json:",omitempty" xml:",omitempty"`
	PhysicalAssetSpecification []OpPhysicalAssetSpecificationType `json:",omitempty" xml:",omitempty"`
	MaterialSpecification      []OpMaterialSpecificationType      `json:",omitempty" xml:",omitempty"`
	SegmentDependency          []SegmentDependencyType            `json:",omitempty" xml:",omitempty"`
	OperationsSegment          []OperationsSegmentType            `json:",omitempty" xml:",omitempty"`
}

// ParameterType was auto-generated. Do not change.
type ParameterType struct {
	ID          IdentifierType    `json:",omitempty" xml:",omitempty"`
	Value       []ValueType       `json:",omitempty" xml:",omitempty"`
	Description []DescriptionType `json:",omitempty" xml:",omitempty"`
	Parameter   []ParameterType   `json:",omitempty" xml:",omitempty"`
}

// OpPersonnelSpecificationType was auto-generated. Do not change.
type OpPersonnelSpecificationType struct {
	PersonnelClassID               []PersonnelClassIDType                 `json:",omitempty" xml:",omitempty"`
	PersonID                       []PersonIDType                         `json:",omitempty" xml:",omitempty"`
	Description                    []DescriptionType                      `json:",omitempty" xml:",omitempty"`
	PersonnelUse                   PersonnelUseType                       `json:",omitempty" xml:",omitempty"`
	Quantity                       []QuantityValueType                    `json:",omitempty" xml:",omitempty"`
	PersonnelSpecificationProperty []OpPersonnelSpecificationPropertyType `json:",omitempty" xml:",omitempty"`
}

// OpPersonnelSpecificationPropertyType was auto-generated. Do not change.
type OpPersonnelSpecificationPropertyType struct {
	ID          IdentifierType      `json:",omitempty" xml:",omitempty"`
	Description []DescriptionType   `json:",omitempty" xml:",omitempty"`
	Value       []ValueType         `json:",omitempty" xml:",omitempty"`
	Quantity    []QuantityValueType `json:",omitempty" xml:",omitempty"`
}

// OpEquipmentSpecificationType was auto-generated. Do not change.
type OpEquipmentSpecificationType struct {
	EquipmentClassID               []EquipmentClassIDType                 `json:",omitempty" xml:",omitempty"`
	EquipmentID                    []EquipmentIDType                      `json:",omitempty" xml:",omitempty"`
	Description                    []DescriptionType                      `json:",omitempty" xml:",omitempty"`
	EquipmentUse                   EquipmentUseType                       `json:",omitempty" xml:",omitempty"`
	Quantity                       []QuantityValueType                    `json:",omitempty" xml:",omitempty"`
	EquipmentSpecificationProperty []OpEquipmentSpecificationPropertyType `json:",omitempty" xml:",omitempty"`
}

// OpEquipmentSpecificationPropertyType was auto-generated. Do not change.
type OpEquipmentSpecificationPropertyType struct {
	ID          IdentifierType      `json:",omitempty" xml:",omitempty"`
	Description []DescriptionType   `json:",omitempty" xml:",omitempty"`
	Value       []ValueType         `json:",omitempty" xml:",omitempty"`
	Quantity    []QuantityValueType `json:",omitempty" xml:",omitempty"`
}

// OpPhysicalAssetSpecificationType was auto-generated. Do not change.
type OpPhysicalAssetSpecificationType struct {
	PhysicalAssetClassID               []PhysicalAssetClassIDType                 `json:",omitempty" xml:",omitempty"`
	PhysicalAssetID                    []PhysicalAssetIDType                      `json:",omitempty" xml:",omitempty"`
	Description                        []DescriptionType                          `json:",omitempty" xml:",omitempty"`
	PhysicalAssetUse                   PhysicalAssetUseType                       `json:",omitempty" xml:",omitempty"`
	Quantity                           []QuantityValueType                        `json:",omitempty" xml:",omitempty"`
	PhysicalAssetSpecificationProperty []OpPhysicalAssetSpecificationPropertyType `json:",omitempty" xml:",omitempty"`
}

// OpPhysicalAssetSpecificationPropertyType was auto-generated. Do not change.
type OpPhysicalAssetSpecificationPropertyType struct {
	ID          IdentifierType      `json:",omitempty" xml:",omitempty"`
	Description []DescriptionType   `json:",omitempty" xml:",omitempty"`
	Value       []ValueType         `json:",omitempty" xml:",omitempty"`
	Quantity    []QuantityValueType `json:",omitempty" xml:",omitempty"`
}

// OpMaterialSpecificationType was auto-generated. Do not change.
type OpMaterialSpecificationType struct {
	ID                            IdentifierType                        `json:",omitempty" xml:",omitempty"`
	MaterialClassID               []MaterialClassIDType                 `json:",omitempty" xml:",omitempty"`
	MaterialDefinitionID          []MaterialDefinitionIDType            `json:",omitempty" xml:",omitempty"`
	Description                   []DescriptionType                     `json:",omitempty" xml:",omitempty"`
	MaterialUse                   MaterialUseType                       `json:",omitempty" xml:",omitempty"`
	Quantity                      []QuantityValueType                   `json:",omitempty" xml:",omitempty"`
	AssemblySpecification         []OpMaterialSpecificationType         `json:",omitempty" xml:",omitempty"`
	AssemblyType                  AssemblyTypeType                      `json:",omitempty" xml:",omitempty"`
	AssemblyRelationship          AssemblyRelationshipType              `json:",omitempty" xml:",omitempty"`
	MaterialSpecificationProperty []OpMaterialSpecificationPropertyType `json:",omitempty" xml:",omitempty"`
}

// OpMaterialSpecificationPropertyType was auto-generated. Do not change.
type OpMaterialSpecificationPropertyType struct {
	ID          IdentifierType      `json:",omitempty" xml:",omitempty"`
	Description []DescriptionType   `json:",omitempty" xml:",omitempty"`
	Value       []ValueType         `json:",omitempty" xml:",omitempty"`
	Quantity    []QuantityValueType `json:",omitempty" xml:",omitempty"`
}

// SegmentDependencyType was auto-generated. Do not change.
type SegmentDependencyType struct {
	ID           IdentifierType    `json:",omitempty" xml:",omitempty"`
	Description  []DescriptionType `json:",omitempty" xml:",omitempty"`
	Dependency   DependencyType    `json:",omitempty" xml:",omitempty"`
	TimingFactor []ValueType       `json:",omitempty" xml:",omitempty"`
	Items        []IdentifierType  `json:",omitempty" xml:",omitempty"`
}

// GetOperationsDefinitionInformationType was auto-generated. Do not change.
type GetOperationsDefinitionInformationType struct {
	ApplicationArea TransApplicationAreaType                       `json:",omitempty" xml:",omitempty"`
	DataArea        GetOperationsDefinitionInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                         `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                         `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// GetOperationsDefinitionInformationTypeDataArea was auto-generated. Do not change.
type GetOperationsDefinitionInformationTypeDataArea struct {
	Get                             []string                              `json:",omitempty" xml:",omitempty"`
	OperationsDefinitionInformation []OperationsDefinitionInformationType `json:",omitempty" xml:",omitempty"`
}

// ShowOperationsDefinitionInformationType was auto-generated. Do not change.
type ShowOperationsDefinitionInformationType struct {
	ApplicationArea TransApplicationAreaType                        `json:",omitempty" xml:",omitempty"`
	DataArea        ShowOperationsDefinitionInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                          `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                          `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ShowOperationsDefinitionInformationTypeDataArea was auto-generated. Do not change.
type ShowOperationsDefinitionInformationTypeDataArea struct {
	Show                            TransShowType                         `json:",omitempty" xml:",omitempty"`
	OperationsDefinitionInformation []OperationsDefinitionInformationType `json:",omitempty" xml:",omitempty"`
}

// ProcessOperationsDefinitionInformationType was auto-generated. Do not change.
type ProcessOperationsDefinitionInformationType struct {
	ApplicationArea TransApplicationAreaType                           `json:",omitempty" xml:",omitempty"`
	DataArea        ProcessOperationsDefinitionInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                             `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                             `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ProcessOperationsDefinitionInformationTypeDataArea was auto-generated. Do not change.
type ProcessOperationsDefinitionInformationTypeDataArea struct {
	Process                         TransProcessType                      `json:",omitempty" xml:",omitempty"`
	OperationsDefinitionInformation []OperationsDefinitionInformationType `json:",omitempty" xml:",omitempty"`
}

// AcknowledgeOperationsDefinitionInformationType was auto-generated. Do not change.
type AcknowledgeOperationsDefinitionInformationType struct {
	ApplicationArea TransApplicationAreaType                               `json:",omitempty" xml:",omitempty"`
	DataArea        AcknowledgeOperationsDefinitionInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                                 `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                                 `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// AcknowledgeOperationsDefinitionInformationTypeDataArea was auto-generated. Do not change.
type AcknowledgeOperationsDefinitionInformationTypeDataArea struct {
	Acknowledge                     TransAcknowledgeType                  `json:",omitempty" xml:",omitempty"`
	OperationsDefinitionInformation []OperationsDefinitionInformationType `json:",omitempty" xml:",omitempty"`
}

// ChangeOperationsDefinitionInformationType was auto-generated. Do not change.
type ChangeOperationsDefinitionInformationType struct {
	ApplicationArea TransApplicationAreaType                          `json:",omitempty" xml:",omitempty"`
	DataArea        ChangeOperationsDefinitionInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                            `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                            `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ChangeOperationsDefinitionInformationTypeDataArea was auto-generated. Do not change.
type ChangeOperationsDefinitionInformationTypeDataArea struct {
	Change                          TransChangeType                       `json:",omitempty" xml:",omitempty"`
	OperationsDefinitionInformation []OperationsDefinitionInformationType `json:",omitempty" xml:",omitempty"`
}

// RespondOperationsDefinitionInformationType was auto-generated. Do not change.
type RespondOperationsDefinitionInformationType struct {
	ApplicationArea TransApplicationAreaType                           `json:",omitempty" xml:",omitempty"`
	DataArea        RespondOperationsDefinitionInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                             `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                             `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// RespondOperationsDefinitionInformationTypeDataArea was auto-generated. Do not change.
type RespondOperationsDefinitionInformationTypeDataArea struct {
	Respond                         TransRespondType                      `json:",omitempty" xml:",omitempty"`
	OperationsDefinitionInformation []OperationsDefinitionInformationType `json:",omitempty" xml:",omitempty"`
}

// CancelOperationsDefinitionInformationType was auto-generated. Do not change.
type CancelOperationsDefinitionInformationType struct {
	ApplicationArea TransApplicationAreaType                          `json:",omitempty" xml:",omitempty"`
	DataArea        CancelOperationsDefinitionInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                            `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                            `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// CancelOperationsDefinitionInformationTypeDataArea was auto-generated. Do not change.
type CancelOperationsDefinitionInformationTypeDataArea struct {
	Cancel                          []TransActionCriteriaType             `json:",omitempty" xml:",omitempty"`
	OperationsDefinitionInformation []OperationsDefinitionInformationType `json:",omitempty" xml:",omitempty"`
}

// SyncOperationsDefinitionInformationType was auto-generated. Do not change.
type SyncOperationsDefinitionInformationType struct {
	ApplicationArea TransApplicationAreaType                        `json:",omitempty" xml:",omitempty"`
	DataArea        SyncOperationsDefinitionInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                          `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                          `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// SyncOperationsDefinitionInformationTypeDataArea was auto-generated. Do not change.
type SyncOperationsDefinitionInformationTypeDataArea struct {
	Sync                            []TransActionCriteriaType             `json:",omitempty" xml:",omitempty"`
	OperationsDefinitionInformation []OperationsDefinitionInformationType `json:",omitempty" xml:",omitempty"`
}

// GetOperationsDefinitionType was auto-generated. Do not change.
type GetOperationsDefinitionType struct {
	ApplicationArea TransApplicationAreaType            `json:",omitempty" xml:",omitempty"`
	DataArea        GetOperationsDefinitionTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                              `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                              `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// GetOperationsDefinitionTypeDataArea was auto-generated. Do not change.
type GetOperationsDefinitionTypeDataArea struct {
	Get                  []string                   `json:",omitempty" xml:",omitempty"`
	OperationsDefinition []OperationsDefinitionType `json:",omitempty" xml:",omitempty"`
}

// ShowOperationsDefinitionType was auto-generated. Do not change.
type ShowOperationsDefinitionType struct {
	ApplicationArea TransApplicationAreaType             `json:",omitempty" xml:",omitempty"`
	DataArea        ShowOperationsDefinitionTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                               `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                               `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ShowOperationsDefinitionTypeDataArea was auto-generated. Do not change.
type ShowOperationsDefinitionTypeDataArea struct {
	Show                 TransShowType              `json:",omitempty" xml:",omitempty"`
	OperationsDefinition []OperationsDefinitionType `json:",omitempty" xml:",omitempty"`
}

// ProcessOperationsDefinitionType was auto-generated. Do not change.
type ProcessOperationsDefinitionType struct {
	ApplicationArea TransApplicationAreaType                `json:",omitempty" xml:",omitempty"`
	DataArea        ProcessOperationsDefinitionTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                  `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                  `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ProcessOperationsDefinitionTypeDataArea was auto-generated. Do not change.
type ProcessOperationsDefinitionTypeDataArea struct {
	Process              TransProcessType           `json:",omitempty" xml:",omitempty"`
	OperationsDefinition []OperationsDefinitionType `json:",omitempty" xml:",omitempty"`
}

// AcknowledgeOperationsDefinitionType was auto-generated. Do not change.
type AcknowledgeOperationsDefinitionType struct {
	ApplicationArea TransApplicationAreaType                    `json:",omitempty" xml:",omitempty"`
	DataArea        AcknowledgeOperationsDefinitionTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                      `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                      `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// AcknowledgeOperationsDefinitionTypeDataArea was auto-generated. Do not change.
type AcknowledgeOperationsDefinitionTypeDataArea struct {
	Acknowledge          TransAcknowledgeType       `json:",omitempty" xml:",omitempty"`
	OperationsDefinition []OperationsDefinitionType `json:",omitempty" xml:",omitempty"`
}

// ChangeOperationsDefinitionType was auto-generated. Do not change.
type ChangeOperationsDefinitionType struct {
	ApplicationArea TransApplicationAreaType               `json:",omitempty" xml:",omitempty"`
	DataArea        ChangeOperationsDefinitionTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                 `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                 `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ChangeOperationsDefinitionTypeDataArea was auto-generated. Do not change.
type ChangeOperationsDefinitionTypeDataArea struct {
	Change               TransChangeType            `json:",omitempty" xml:",omitempty"`
	OperationsDefinition []OperationsDefinitionType `json:",omitempty" xml:",omitempty"`
}

// RespondOperationsDefinitionType was auto-generated. Do not change.
type RespondOperationsDefinitionType struct {
	ApplicationArea TransApplicationAreaType                `json:",omitempty" xml:",omitempty"`
	DataArea        RespondOperationsDefinitionTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                  `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                  `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// RespondOperationsDefinitionTypeDataArea was auto-generated. Do not change.
type RespondOperationsDefinitionTypeDataArea struct {
	Respond              TransRespondType           `json:",omitempty" xml:",omitempty"`
	OperationsDefinition []OperationsDefinitionType `json:",omitempty" xml:",omitempty"`
}

// CancelOperationsDefinitionType was auto-generated. Do not change.
type CancelOperationsDefinitionType struct {
	ApplicationArea TransApplicationAreaType               `json:",omitempty" xml:",omitempty"`
	DataArea        CancelOperationsDefinitionTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                 `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                 `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// CancelOperationsDefinitionTypeDataArea was auto-generated. Do not change.
type CancelOperationsDefinitionTypeDataArea struct {
	Cancel               []TransActionCriteriaType  `json:",omitempty" xml:",omitempty"`
	OperationsDefinition []OperationsDefinitionType `json:",omitempty" xml:",omitempty"`
}

// SyncOperationsDefinitionType was auto-generated. Do not change.
type SyncOperationsDefinitionType struct {
	ApplicationArea TransApplicationAreaType             `json:",omitempty" xml:",omitempty"`
	DataArea        SyncOperationsDefinitionTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                               `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                               `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// SyncOperationsDefinitionTypeDataArea was auto-generated. Do not change.
type SyncOperationsDefinitionTypeDataArea struct {
	Sync                 []TransActionCriteriaType  `json:",omitempty" xml:",omitempty"`
	OperationsDefinition []OperationsDefinitionType `json:",omitempty" xml:",omitempty"`
}

// OperationsPerformanceType was auto-generated. Do not change.
type OperationsPerformanceType struct {
	ID                   IdentifierType           `json:",omitempty" xml:",omitempty"`
	Description          []DescriptionType        `json:",omitempty" xml:",omitempty"`
	HierarchyScope       HierarchyScopeType       `json:",omitempty" xml:",omitempty"`
	OperationsType       OperationsTypeType       `json:",omitempty" xml:",omitempty"`
	OperationsScheduleID OperationsScheduleIDType `json:",omitempty" xml:",omitempty"`
	StartTime            StartTimeType            `json:",omitempty" xml:",omitempty"`
	EndTime              EndTimeType              `json:",omitempty" xml:",omitempty"`
	PerformanceState     ResponseStateType        `json:",omitempty" xml:",omitempty"`
	PublishedDate        PublishedDateType        `json:",omitempty" xml:",omitempty"`
	OperationsResponse   []OperationsResponseType `json:",omitempty" xml:",omitempty"`
}

// OperationsResponseType was auto-generated. Do not change.
type OperationsResponseType struct {
	ID                     IdentifierType               `json:",omitempty" xml:",omitempty"`
	Description            []DescriptionType            `json:",omitempty" xml:",omitempty"`
	HierarchyScope         HierarchyScopeType           `json:",omitempty" xml:",omitempty"`
	OperationsType         OperationsTypeType           `json:",omitempty" xml:",omitempty"`
	OperationsRequestID    []OperationsRequestIDType    `json:",omitempty" xml:",omitempty"`
	StartTime              StartTimeType                `json:",omitempty" xml:",omitempty"`
	EndTime                EndTimeType                  `json:",omitempty" xml:",omitempty"`
	OperationsDefinitionID []OperationsDefinitionIDType `json:",omitempty" xml:",omitempty"`
	ResponseState          ResponseStateType            `json:",omitempty" xml:",omitempty"`
	SegmentResponse        []OpSegmentResponseType      `json:",omitempty" xml:",omitempty"`
}

// OpSegmentResponseType was auto-generated. Do not change.
type OpSegmentResponseType struct {
	ID                                 IdentifierType                         `json:",omitempty" xml:",omitempty"`
	Description                        []DescriptionType                      `json:",omitempty" xml:",omitempty"`
	HierarchyScope                     HierarchyScopeType                     `json:",omitempty" xml:",omitempty"`
	OperationsType                     OperationsTypeType                     `json:",omitempty" xml:",omitempty"`
	ProcessSegmentID                   []ProcessSegmentIDType                 `json:",omitempty" xml:",omitempty"`
	ActualStartTime                    ActualStartTimeType                    `json:",omitempty" xml:",omitempty"`
	ActualEndTime                      ActualEndTimeType                      `json:",omitempty" xml:",omitempty"`
	OperationsDefinitionID             []OperationsDefinitionIDType           `json:",omitempty" xml:",omitempty"`
	SegmentState                       ResponseStateType                      `json:",omitempty" xml:",omitempty"`
	SegmentData                        []OpSegmentDataType                    `json:",omitempty" xml:",omitempty"`
	PersonnelActual                    []OpPersonnelActualType                `json:",omitempty" xml:",omitempty"`
	EquipmentActual                    []OpEquipmentActualType                `json:",omitempty" xml:",omitempty"`
	PhysicalAssetActual                []OpPhysicalAssetActualType            `json:",omitempty" xml:",omitempty"`
	MaterialActual                     []OpMaterialActualType                 `json:",omitempty" xml:",omitempty"`
	SegmentResponse                    []OpSegmentResponseType                `json:",omitempty" xml:",omitempty"`
	RequiredByRequestedSegmentResponse RequiredByRequestedSegmentResponseType `json:",omitempty" xml:",omitempty"`
}

// OpSegmentDataType was auto-generated. Do not change.
type OpSegmentDataType struct {
	ID                                 IdentifierType                         `json:",omitempty" xml:",omitempty"`
	Description                        []DescriptionType                      `json:",omitempty" xml:",omitempty"`
	Value                              []ValueType                            `json:",omitempty" xml:",omitempty"`
	SegmentData                        []OpSegmentDataType                    `json:",omitempty" xml:",omitempty"`
	RequiredByRequestedSegmentResponse RequiredByRequestedSegmentResponseType `json:",omitempty" xml:",omitempty"`
}

// OpPersonnelActualType was auto-generated. Do not change.
type OpPersonnelActualType struct {
	PersonnelClassID                   []PersonnelClassIDType                 `json:",omitempty" xml:",omitempty"`
	PersonID                           []PersonIDType                         `json:",omitempty" xml:",omitempty"`
	Description                        []DescriptionType                      `json:",omitempty" xml:",omitempty"`
	PersonnelUse                       PersonnelUseType                       `json:",omitempty" xml:",omitempty"`
	Quantity                           []QuantityValueType                    `json:",omitempty" xml:",omitempty"`
	HierarchyScope                     HierarchyScopeType                     `json:",omitempty" xml:",omitempty"`
	PersonnelActualProperty            []OpPersonnelActualPropertyType        `json:",omitempty" xml:",omitempty"`
	RequiredByRequestedSegmentResponse RequiredByRequestedSegmentResponseType `json:",omitempty" xml:",omitempty"`
}

// OpPersonnelActualPropertyType was auto-generated. Do not change.
type OpPersonnelActualPropertyType struct {
	ID                                 IdentifierType                         `json:",omitempty" xml:",omitempty"`
	Description                        []DescriptionType                      `json:",omitempty" xml:",omitempty"`
	Value                              []ValueType                            `json:",omitempty" xml:",omitempty"`
	Quantity                           []QuantityValueType                    `json:",omitempty" xml:",omitempty"`
	RequiredByRequestedSegmentResponse RequiredByRequestedSegmentResponseType `json:",omitempty" xml:",omitempty"`
}

// OpEquipmentActualType was auto-generated. Do not change.
type OpEquipmentActualType struct {
	EquipmentClassID                   []EquipmentClassIDType                 `json:",omitempty" xml:",omitempty"`
	EquipmentID                        []EquipmentIDType                      `json:",omitempty" xml:",omitempty"`
	Description                        []DescriptionType                      `json:",omitempty" xml:",omitempty"`
	EquipmentUse                       EquipmentUseType                       `json:",omitempty" xml:",omitempty"`
	Quantity                           []QuantityValueType                    `json:",omitempty" xml:",omitempty"`
	HierarchyScope                     HierarchyScopeType                     `json:",omitempty" xml:",omitempty"`
	EquipmentActualProperty            []OpEquipmentActualPropertyType        `json:",omitempty" xml:",omitempty"`
	RequiredByRequestedSegmentResponse RequiredByRequestedSegmentResponseType `json:",omitempty" xml:",omitempty"`
}

// OpEquipmentActualPropertyType was auto-generated. Do not change.
type OpEquipmentActualPropertyType struct {
	ID                                 IdentifierType                         `json:",omitempty" xml:",omitempty"`
	Description                        []DescriptionType                      `json:",omitempty" xml:",omitempty"`
	Value                              []ValueType                            `json:",omitempty" xml:",omitempty"`
	Quantity                           []QuantityValueType                    `json:",omitempty" xml:",omitempty"`
	RequiredByRequestedSegmentResponse RequiredByRequestedSegmentResponseType `json:",omitempty" xml:",omitempty"`
}

// OpPhysicalAssetActualType was auto-generated. Do not change.
type OpPhysicalAssetActualType struct {
	PhysicalAssetClassID               []PhysicalAssetClassIDType             `json:",omitempty" xml:",omitempty"`
	PhysicalAssetID                    []PhysicalAssetIDType                  `json:",omitempty" xml:",omitempty"`
	Description                        []DescriptionType                      `json:",omitempty" xml:",omitempty"`
	Quantity                           []QuantityValueType                    `json:",omitempty" xml:",omitempty"`
	HierarchyScope                     HierarchyScopeType                     `json:",omitempty" xml:",omitempty"`
	PhysicalAssetActualProperty        []OpPhysicalAssetActualPropertyType    `json:",omitempty" xml:",omitempty"`
	RequiredByRequestedSegmentResponse RequiredByRequestedSegmentResponseType `json:",omitempty" xml:",omitempty"`
}

// OpPhysicalAssetActualPropertyType was auto-generated. Do not change.
type OpPhysicalAssetActualPropertyType struct {
	ID                                 IdentifierType                         `json:",omitempty" xml:",omitempty"`
	Description                        []DescriptionType                      `json:",omitempty" xml:",omitempty"`
	Value                              []ValueType                            `json:",omitempty" xml:",omitempty"`
	Quantity                           []QuantityValueType                    `json:",omitempty" xml:",omitempty"`
	RequiredByRequestedSegmentResponse RequiredByRequestedSegmentResponseType `json:",omitempty" xml:",omitempty"`
}

// OpMaterialActualType was auto-generated. Do not change.
type OpMaterialActualType struct {
	MaterialClassID                    []MaterialClassIDType                  `json:",omitempty" xml:",omitempty"`
	MaterialDefinitionID               []MaterialDefinitionIDType             `json:",omitempty" xml:",omitempty"`
	MaterialLotID                      []MaterialLotIDType                    `json:",omitempty" xml:",omitempty"`
	MaterialSubLotID                   []MaterialSubLotIDType                 `json:",omitempty" xml:",omitempty"`
	Description                        []DescriptionType                      `json:",omitempty" xml:",omitempty"`
	MaterialUse                        MaterialUseType                        `json:",omitempty" xml:",omitempty"`
	StorageLocation                    StorageLocationType                    `json:",omitempty" xml:",omitempty"`
	Quantity                           []QuantityValueType                    `json:",omitempty" xml:",omitempty"`
	AssemblyActual                     []OpMaterialActualType                 `json:",omitempty" xml:",omitempty"`
	AssemblyType                       AssemblyTypeType                       `json:",omitempty" xml:",omitempty"`
	AssemblyRelationship               AssemblyRelationshipType               `json:",omitempty" xml:",omitempty"`
	HierarchyScope                     HierarchyScopeType                     `json:",omitempty" xml:",omitempty"`
	MaterialActualProperty             []OpMaterialActualPropertyType         `json:",omitempty" xml:",omitempty"`
	RequiredByRequestedSegmentResponse RequiredByRequestedSegmentResponseType `json:",omitempty" xml:",omitempty"`
}

// OpMaterialActualPropertyType was auto-generated. Do not change.
type OpMaterialActualPropertyType struct {
	ID                                 IdentifierType                         `json:",omitempty" xml:",omitempty"`
	Description                        []DescriptionType                      `json:",omitempty" xml:",omitempty"`
	Value                              []ValueType                            `json:",omitempty" xml:",omitempty"`
	Quantity                           []QuantityValueType                    `json:",omitempty" xml:",omitempty"`
	MaterialActualProperty             []OpMaterialActualPropertyType         `json:",omitempty" xml:",omitempty"`
	RequiredByRequestedSegmentResponse RequiredByRequestedSegmentResponseType `json:",omitempty" xml:",omitempty"`
}

// GetOperationsPerformanceType was auto-generated. Do not change.
type GetOperationsPerformanceType struct {
	ApplicationArea TransApplicationAreaType             `json:",omitempty" xml:",omitempty"`
	DataArea        GetOperationsPerformanceTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                               `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                               `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// GetOperationsPerformanceTypeDataArea was auto-generated. Do not change.
type GetOperationsPerformanceTypeDataArea struct {
	Get                   []string                    `json:",omitempty" xml:",omitempty"`
	OperationsPerformance []OperationsPerformanceType `json:",omitempty" xml:",omitempty"`
}

// ShowOperationsPerformanceType was auto-generated. Do not change.
type ShowOperationsPerformanceType struct {
	ApplicationArea TransApplicationAreaType              `json:",omitempty" xml:",omitempty"`
	DataArea        ShowOperationsPerformanceTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ShowOperationsPerformanceTypeDataArea was auto-generated. Do not change.
type ShowOperationsPerformanceTypeDataArea struct {
	Show                  TransShowType               `json:",omitempty" xml:",omitempty"`
	OperationsPerformance []OperationsPerformanceType `json:",omitempty" xml:",omitempty"`
}

// ProcessOperationsPerformanceType was auto-generated. Do not change.
type ProcessOperationsPerformanceType struct {
	ApplicationArea TransApplicationAreaType                 `json:",omitempty" xml:",omitempty"`
	DataArea        ProcessOperationsPerformanceTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                   `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                   `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ProcessOperationsPerformanceTypeDataArea was auto-generated. Do not change.
type ProcessOperationsPerformanceTypeDataArea struct {
	Process               TransProcessType            `json:",omitempty" xml:",omitempty"`
	OperationsPerformance []OperationsPerformanceType `json:",omitempty" xml:",omitempty"`
}

// AcknowledgeOperationsPerformanceType was auto-generated. Do not change.
type AcknowledgeOperationsPerformanceType struct {
	ApplicationArea TransApplicationAreaType                     `json:",omitempty" xml:",omitempty"`
	DataArea        AcknowledgeOperationsPerformanceTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                       `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                       `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// AcknowledgeOperationsPerformanceTypeDataArea was auto-generated. Do not change.
type AcknowledgeOperationsPerformanceTypeDataArea struct {
	Acknowledge           TransAcknowledgeType        `json:",omitempty" xml:",omitempty"`
	OperationsPerformance []OperationsPerformanceType `json:",omitempty" xml:",omitempty"`
}

// ChangeOperationsPerformanceType was auto-generated. Do not change.
type ChangeOperationsPerformanceType struct {
	ApplicationArea TransApplicationAreaType                `json:",omitempty" xml:",omitempty"`
	DataArea        ChangeOperationsPerformanceTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                  `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                  `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ChangeOperationsPerformanceTypeDataArea was auto-generated. Do not change.
type ChangeOperationsPerformanceTypeDataArea struct {
	Change                TransChangeType             `json:",omitempty" xml:",omitempty"`
	OperationsPerformance []OperationsPerformanceType `json:",omitempty" xml:",omitempty"`
}

// RespondOperationsPerformanceType was auto-generated. Do not change.
type RespondOperationsPerformanceType struct {
	ApplicationArea TransApplicationAreaType                 `json:",omitempty" xml:",omitempty"`
	DataArea        RespondOperationsPerformanceTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                   `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                   `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// RespondOperationsPerformanceTypeDataArea was auto-generated. Do not change.
type RespondOperationsPerformanceTypeDataArea struct {
	Respond               TransRespondType            `json:",omitempty" xml:",omitempty"`
	OperationsPerformance []OperationsPerformanceType `json:",omitempty" xml:",omitempty"`
}

// CancelOperationsPerformanceType was auto-generated. Do not change.
type CancelOperationsPerformanceType struct {
	ApplicationArea TransApplicationAreaType                `json:",omitempty" xml:",omitempty"`
	DataArea        CancelOperationsPerformanceTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                  `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                  `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// CancelOperationsPerformanceTypeDataArea was auto-generated. Do not change.
type CancelOperationsPerformanceTypeDataArea struct {
	Cancel                []TransActionCriteriaType   `json:",omitempty" xml:",omitempty"`
	OperationsPerformance []OperationsPerformanceType `json:",omitempty" xml:",omitempty"`
}

// SyncOperationsPerformanceType was auto-generated. Do not change.
type SyncOperationsPerformanceType struct {
	ApplicationArea TransApplicationAreaType              `json:",omitempty" xml:",omitempty"`
	DataArea        SyncOperationsPerformanceTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// SyncOperationsPerformanceTypeDataArea was auto-generated. Do not change.
type SyncOperationsPerformanceTypeDataArea struct {
	Sync                  []TransActionCriteriaType   `json:",omitempty" xml:",omitempty"`
	OperationsPerformance []OperationsPerformanceType `json:",omitempty" xml:",omitempty"`
}

// OperationsScheduleType was auto-generated. Do not change.
type OperationsScheduleType struct {
	ID                IdentifierType          `json:",omitempty" xml:",omitempty"`
	Description       []DescriptionType       `json:",omitempty" xml:",omitempty"`
	HierarchyScope    HierarchyScopeType      `json:",omitempty" xml:",omitempty"`
	OperationsType    OperationsTypeType      `json:",omitempty" xml:",omitempty"`
	StartTime         StartTimeType           `json:",omitempty" xml:",omitempty"`
	EndTime           EndTimeType             `json:",omitempty" xml:",omitempty"`
	ScheduleState     RequestStateType        `json:",omitempty" xml:",omitempty"`
	PublishedDate     PublishedDateType       `json:",omitempty" xml:",omitempty"`
	OperationsRequest []OperationsRequestType `json:",omitempty" xml:",omitempty"`
}

// OperationsRequestType was auto-generated. Do not change.
type OperationsRequestType struct {
	ID                     IdentifierType             `json:",omitempty" xml:",omitempty"`
	Description            []DescriptionType          `json:",omitempty" xml:",omitempty"`
	HierarchyScope         HierarchyScopeType         `json:",omitempty" xml:",omitempty"`
	OperationsType         OperationsTypeType         `json:",omitempty" xml:",omitempty"`
	StartTime              StartTimeType              `json:",omitempty" xml:",omitempty"`
	EndTime                EndTimeType                `json:",omitempty" xml:",omitempty"`
	Priority               PriorityType               `json:",omitempty" xml:",omitempty"`
	OperationsDefinitionID OperationsDefinitionIDType `json:",omitempty" xml:",omitempty"`
	RequestState           RequestStateType           `json:",omitempty" xml:",omitempty"`
	SegmentRequirement     []OpSegmentRequirementType `json:",omitempty" xml:",omitempty"`
	SegmentResponse        []OpSegmentResponseType    `json:",omitempty" xml:",omitempty"`
}

// PriorityType was auto-generated. Do not change.
type PriorityType struct {
	NumericType
}

// NumericType was auto-generated. Do not change.
type NumericType struct {
	Format string  `json:",omitempty" xml:"format,attr,omitempty"`
	Value  float64 `json:",omitempty" xml:",chardata"`
}

// EvaluationOrderType was auto-generated. Do not change.
type EvaluationOrderType struct {
	NumericType
}

// EnumerationNumberType was auto-generated. Do not change.
type EnumerationNumberType struct {
	NumericType
}

// BatchPriorityType was auto-generated. Do not change.
type BatchPriorityType struct {
	NumericType
}

// ActualBatchSizeType was auto-generated. Do not change.
type ActualBatchSizeType struct {
	NumericType
}

// RequestedPriorityType was auto-generated. Do not change.
type RequestedPriorityType struct {
	NumericType
}

// OpSegmentRequirementType was auto-generated. Do not change.
type OpSegmentRequirementType struct {
	ID                                 IdentifierType                         `json:",omitempty" xml:",omitempty"`
	Description                        []DescriptionType                      `json:",omitempty" xml:",omitempty"`
	HierarchyScope                     HierarchyScopeType                     `json:",omitempty" xml:",omitempty"`
	OperationsType                     OperationsTypeType                     `json:",omitempty" xml:",omitempty"`
	ProcessSegmentID                   ProcessSegmentIDType                   `json:",omitempty" xml:",omitempty"`
	EarliestStartTime                  EarliestStartTimeType                  `json:",omitempty" xml:",omitempty"`
	LatestEndTime                      LatestEndTimeType                      `json:",omitempty" xml:",omitempty"`
	Duration                           string                                 `json:",omitempty" xml:",omitempty"`
	OperationsDefinitionID             OperationsDefinitionIDType             `json:",omitempty" xml:",omitempty"`
	SegmentState                       RequestStateType                       `json:",omitempty" xml:",omitempty"`
	SegmentParameter                   []ParameterType                        `json:",omitempty" xml:",omitempty"`
	PersonnelRequirement               []OpPersonnelRequirementType           `json:",omitempty" xml:",omitempty"`
	EquipmentRequirement               []OpEquipmentRequirementType           `json:",omitempty" xml:",omitempty"`
	PhysicalAssetRequirement           []OpPhysicalAssetRequirementType       `json:",omitempty" xml:",omitempty"`
	MaterialRequirement                []OpMaterialRequirementType            `json:",omitempty" xml:",omitempty"`
	SegmentRequirement                 []OpSegmentRequirementType             `json:",omitempty" xml:",omitempty"`
	RequiredByRequestedSegmentResponse RequiredByRequestedSegmentResponseType `json:",omitempty" xml:",omitempty"`
}

// OpPersonnelRequirementType was auto-generated. Do not change.
type OpPersonnelRequirementType struct {
	PersonnelClassID                   []PersonnelClassIDType                 `json:",omitempty" xml:",omitempty"`
	PersonID                           []PersonIDType                         `json:",omitempty" xml:",omitempty"`
	Description                        []DescriptionType                      `json:",omitempty" xml:",omitempty"`
	PersonnelUse                       PersonnelUseType                       `json:",omitempty" xml:",omitempty"`
	Quantity                           []QuantityValueType                    `json:",omitempty" xml:",omitempty"`
	HierarchyScope                     HierarchyScopeType                     `json:",omitempty" xml:",omitempty"`
	PersonnelRequirementProperty       []OpPersonnelRequirementPropertyType   `json:",omitempty" xml:",omitempty"`
	RequiredByRequestedSegmentResponse RequiredByRequestedSegmentResponseType `json:",omitempty" xml:",omitempty"`
}

// OpPersonnelRequirementPropertyType was auto-generated. Do not change.
type OpPersonnelRequirementPropertyType struct {
	ID          IdentifierType      `json:",omitempty" xml:",omitempty"`
	Description []DescriptionType   `json:",omitempty" xml:",omitempty"`
	Value       []ValueType         `json:",omitempty" xml:",omitempty"`
	Quantity    []QuantityValueType `json:",omitempty" xml:",omitempty"`
}

// OpEquipmentRequirementType was auto-generated. Do not change.
type OpEquipmentRequirementType struct {
	EquipmentClassID                   []EquipmentClassIDType                 `json:",omitempty" xml:",omitempty"`
	EquipmentID                        []EquipmentIDType                      `json:",omitempty" xml:",omitempty"`
	Description                        []DescriptionType                      `json:",omitempty" xml:",omitempty"`
	EquipmentUse                       EquipmentUseType                       `json:",omitempty" xml:",omitempty"`
	Quantity                           []QuantityValueType                    `json:",omitempty" xml:",omitempty"`
	HierarchyScope                     HierarchyScopeType                     `json:",omitempty" xml:",omitempty"`
	EquipmentLevel                     HierarchyScopeType                     `json:",omitempty" xml:",omitempty"`
	EquipmentRequirementProperty       []OpEquipmentRequirementPropertyType   `json:",omitempty" xml:",omitempty"`
	RequiredByRequestedSegmentResponse RequiredByRequestedSegmentResponseType `json:",omitempty" xml:",omitempty"`
}

// OpEquipmentRequirementPropertyType was auto-generated. Do not change.
type OpEquipmentRequirementPropertyType struct {
	ID          IdentifierType      `json:",omitempty" xml:",omitempty"`
	Description []DescriptionType   `json:",omitempty" xml:",omitempty"`
	Value       []ValueType         `json:",omitempty" xml:",omitempty"`
	Quantity    []QuantityValueType `json:",omitempty" xml:",omitempty"`
}

// OpPhysicalAssetRequirementType was auto-generated. Do not change.
type OpPhysicalAssetRequirementType struct {
	PhysicalAssetClassID               []PhysicalAssetClassIDType               `json:",omitempty" xml:",omitempty"`
	PhysicalAssetID                    []PhysicalAssetIDType                    `json:",omitempty" xml:",omitempty"`
	Description                        []DescriptionType                        `json:",omitempty" xml:",omitempty"`
	PhysicalAssetUse                   PhysicalAssetUseType                     `json:",omitempty" xml:",omitempty"`
	Quantity                           []QuantityValueType                      `json:",omitempty" xml:",omitempty"`
	HierarchyScope                     HierarchyScopeType                       `json:",omitempty" xml:",omitempty"`
	EquipmentLevel                     HierarchyScopeType                       `json:",omitempty" xml:",omitempty"`
	PhysicalAssetRequirementProperty   []OpPhysicalAssetRequirementPropertyType `json:",omitempty" xml:",omitempty"`
	RequiredByRequestedSegmentResponse RequiredByRequestedSegmentResponseType   `json:",omitempty" xml:",omitempty"`
}

// OpPhysicalAssetRequirementPropertyType was auto-generated. Do not change.
type OpPhysicalAssetRequirementPropertyType struct {
	ID          IdentifierType      `json:",omitempty" xml:",omitempty"`
	Description []DescriptionType   `json:",omitempty" xml:",omitempty"`
	Value       []ValueType         `json:",omitempty" xml:",omitempty"`
	Quantity    []QuantityValueType `json:",omitempty" xml:",omitempty"`
}

// OpMaterialRequirementType was auto-generated. Do not change.
type OpMaterialRequirementType struct {
	MaterialClassID                    []MaterialClassIDType                  `json:",omitempty" xml:",omitempty"`
	MaterialDefinitionID               []MaterialDefinitionIDType             `json:",omitempty" xml:",omitempty"`
	MaterialLotID                      []MaterialLotIDType                    `json:",omitempty" xml:",omitempty"`
	MaterialSubLotID                   []MaterialSubLotIDType                 `json:",omitempty" xml:",omitempty"`
	Description                        []DescriptionType                      `json:",omitempty" xml:",omitempty"`
	MaterialUse                        MaterialUseType                        `json:",omitempty" xml:",omitempty"`
	StorageLocation                    StorageLocationType                    `json:",omitempty" xml:",omitempty"`
	Quantity                           []QuantityValueType                    `json:",omitempty" xml:",omitempty"`
	AssemblyRequirement                []OpMaterialRequirementType            `json:",omitempty" xml:",omitempty"`
	AssemblyType                       AssemblyTypeType                       `json:",omitempty" xml:",omitempty"`
	AssemblyRelationship               AssemblyRelationshipType               `json:",omitempty" xml:",omitempty"`
	HierarchyScope                     HierarchyScopeType                     `json:",omitempty" xml:",omitempty"`
	MaterialRequirementProperty        []OpMaterialRequirementPropertyType    `json:",omitempty" xml:",omitempty"`
	RequiredByRequestedSegmentResponse RequiredByRequestedSegmentResponseType `json:",omitempty" xml:",omitempty"`
}

// OpMaterialRequirementPropertyType was auto-generated. Do not change.
type OpMaterialRequirementPropertyType struct {
	ID                          IdentifierType                      `json:",omitempty" xml:",omitempty"`
	Description                 []DescriptionType                   `json:",omitempty" xml:",omitempty"`
	Value                       []ValueType                         `json:",omitempty" xml:",omitempty"`
	Quantity                    []QuantityValueType                 `json:",omitempty" xml:",omitempty"`
	MaterialRequirementProperty []OpMaterialRequirementPropertyType `json:",omitempty" xml:",omitempty"`
}

// GetOperationsScheduleType was auto-generated. Do not change.
type GetOperationsScheduleType struct {
	ApplicationArea TransApplicationAreaType          `json:",omitempty" xml:",omitempty"`
	DataArea        GetOperationsScheduleTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                            `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                            `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// GetOperationsScheduleTypeDataArea was auto-generated. Do not change.
type GetOperationsScheduleTypeDataArea struct {
	Get                []string                 `json:",omitempty" xml:",omitempty"`
	OperationsSchedule []OperationsScheduleType `json:",omitempty" xml:",omitempty"`
}

// ShowOperationsScheduleType was auto-generated. Do not change.
type ShowOperationsScheduleType struct {
	ApplicationArea TransApplicationAreaType           `json:",omitempty" xml:",omitempty"`
	DataArea        ShowOperationsScheduleTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                             `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                             `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ShowOperationsScheduleTypeDataArea was auto-generated. Do not change.
type ShowOperationsScheduleTypeDataArea struct {
	Show               TransShowType            `json:",omitempty" xml:",omitempty"`
	OperationsSchedule []OperationsScheduleType `json:",omitempty" xml:",omitempty"`
}

// ProcessOperationsScheduleType was auto-generated. Do not change.
type ProcessOperationsScheduleType struct {
	ApplicationArea TransApplicationAreaType              `json:",omitempty" xml:",omitempty"`
	DataArea        ProcessOperationsScheduleTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ProcessOperationsScheduleTypeDataArea was auto-generated. Do not change.
type ProcessOperationsScheduleTypeDataArea struct {
	Process            TransProcessType         `json:",omitempty" xml:",omitempty"`
	OperationsSchedule []OperationsScheduleType `json:",omitempty" xml:",omitempty"`
}

// AcknowledgeOperationsScheduleType was auto-generated. Do not change.
type AcknowledgeOperationsScheduleType struct {
	ApplicationArea TransApplicationAreaType                  `json:",omitempty" xml:",omitempty"`
	DataArea        AcknowledgeOperationsScheduleTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                    `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                    `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// AcknowledgeOperationsScheduleTypeDataArea was auto-generated. Do not change.
type AcknowledgeOperationsScheduleTypeDataArea struct {
	Acknowledge        TransAcknowledgeType     `json:",omitempty" xml:",omitempty"`
	OperationsSchedule []OperationsScheduleType `json:",omitempty" xml:",omitempty"`
}

// ChangeOperationsScheduleType was auto-generated. Do not change.
type ChangeOperationsScheduleType struct {
	ApplicationArea TransApplicationAreaType             `json:",omitempty" xml:",omitempty"`
	DataArea        ChangeOperationsScheduleTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                               `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                               `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ChangeOperationsScheduleTypeDataArea was auto-generated. Do not change.
type ChangeOperationsScheduleTypeDataArea struct {
	Change             TransChangeType          `json:",omitempty" xml:",omitempty"`
	OperationsSchedule []OperationsScheduleType `json:",omitempty" xml:",omitempty"`
}

// RespondOperationsScheduleType was auto-generated. Do not change.
type RespondOperationsScheduleType struct {
	ApplicationArea TransApplicationAreaType              `json:",omitempty" xml:",omitempty"`
	DataArea        RespondOperationsScheduleTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// RespondOperationsScheduleTypeDataArea was auto-generated. Do not change.
type RespondOperationsScheduleTypeDataArea struct {
	Respond            TransRespondType         `json:",omitempty" xml:",omitempty"`
	OperationsSchedule []OperationsScheduleType `json:",omitempty" xml:",omitempty"`
}

// CancelOperationsScheduleType was auto-generated. Do not change.
type CancelOperationsScheduleType struct {
	ApplicationArea TransApplicationAreaType             `json:",omitempty" xml:",omitempty"`
	DataArea        CancelOperationsScheduleTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                               `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                               `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// CancelOperationsScheduleTypeDataArea was auto-generated. Do not change.
type CancelOperationsScheduleTypeDataArea struct {
	Cancel             []TransActionCriteriaType `json:",omitempty" xml:",omitempty"`
	OperationsSchedule []OperationsScheduleType  `json:",omitempty" xml:",omitempty"`
}

// SyncOperationsScheduleType was auto-generated. Do not change.
type SyncOperationsScheduleType struct {
	ApplicationArea TransApplicationAreaType           `json:",omitempty" xml:",omitempty"`
	DataArea        SyncOperationsScheduleTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                             `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                             `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// SyncOperationsScheduleTypeDataArea was auto-generated. Do not change.
type SyncOperationsScheduleTypeDataArea struct {
	Sync               []TransActionCriteriaType `json:",omitempty" xml:",omitempty"`
	OperationsSchedule []OperationsScheduleType  `json:",omitempty" xml:",omitempty"`
}

// PersonnelInformationType was auto-generated. Do not change.
type PersonnelInformationType struct {
	ID                             IdentifierType                       `json:",omitempty" xml:",omitempty"`
	Description                    []DescriptionType                    `json:",omitempty" xml:",omitempty"`
	Location                       LocationType                         `json:",omitempty" xml:",omitempty"`
	HierarchyScope                 HierarchyScopeType                   `json:",omitempty" xml:",omitempty"`
	PublishedDate                  PublishedDateType                    `json:",omitempty" xml:",omitempty"`
	Person                         []PersonType                         `json:",omitempty" xml:",omitempty"`
	PersonnelClass                 []PersonnelClassType                 `json:",omitempty" xml:",omitempty"`
	QualificationTestSpecification []QualificationTestSpecificationType `json:",omitempty" xml:",omitempty"`
}

// PersonType was auto-generated. Do not change.
type PersonType struct {
	ID                               IdentifierType                         `json:",omitempty" xml:",omitempty"`
	Description                      []DescriptionType                      `json:",omitempty" xml:",omitempty"`
	Location                         LocationType                           `json:",omitempty" xml:",omitempty"`
	HierarchyScope                   HierarchyScopeType                     `json:",omitempty" xml:",omitempty"`
	PersonName                       PersonNameType                         `json:",omitempty" xml:",omitempty"`
	PersonProperty                   []PersonPropertyType                   `json:",omitempty" xml:",omitempty"`
	PersonnelClassID                 []PersonnelClassIDType                 `json:",omitempty" xml:",omitempty"`
	QualificationTestSpecificationID []QualificationTestSpecificationIDType `json:",omitempty" xml:",omitempty"`
}

// PersonPropertyType was auto-generated. Do not change.
type PersonPropertyType struct {
	ID                               IdentifierType                         `json:",omitempty" xml:",omitempty"`
	Description                      []DescriptionType                      `json:",omitempty" xml:",omitempty"`
	Value                            []ValueType                            `json:",omitempty" xml:",omitempty"`
	PersonProperty                   []PersonPropertyType                   `json:",omitempty" xml:",omitempty"`
	QualificationTestSpecificationID []QualificationTestSpecificationIDType `json:",omitempty" xml:",omitempty"`
	TestResult                       []TestResultType                       `json:",omitempty" xml:",omitempty"`
}

// PersonnelClassType was auto-generated. Do not change.
type PersonnelClassType struct {
	ID                               IdentifierType                         `json:",omitempty" xml:",omitempty"`
	Description                      []DescriptionType                      `json:",omitempty" xml:",omitempty"`
	Location                         LocationType                           `json:",omitempty" xml:",omitempty"`
	HierarchyScope                   HierarchyScopeType                     `json:",omitempty" xml:",omitempty"`
	PersonnelClassProperty           []PersonnelClassPropertyType           `json:",omitempty" xml:",omitempty"`
	PersonID                         []PersonIDType                         `json:",omitempty" xml:",omitempty"`
	QualificationTestSpecificationID []QualificationTestSpecificationIDType `json:",omitempty" xml:",omitempty"`
}

// PersonnelClassPropertyType was auto-generated. Do not change.
type PersonnelClassPropertyType struct {
	ID                               IdentifierType                         `json:",omitempty" xml:",omitempty"`
	Description                      []DescriptionType                      `json:",omitempty" xml:",omitempty"`
	Value                            []ValueType                            `json:",omitempty" xml:",omitempty"`
	PersonnelClassProperty           []PersonnelClassPropertyType           `json:",omitempty" xml:",omitempty"`
	QualificationTestSpecificationID []QualificationTestSpecificationIDType `json:",omitempty" xml:",omitempty"`
}

// QualificationTestSpecificationType was auto-generated. Do not change.
type QualificationTestSpecificationType struct {
	ID                           IdentifierType                     `json:",omitempty" xml:",omitempty"`
	Description                  []DescriptionType                  `json:",omitempty" xml:",omitempty"`
	Version                      VersionType                        `json:",omitempty" xml:",omitempty"`
	Location                     LocationType                       `json:",omitempty" xml:",omitempty"`
	HierarchyScope               HierarchyScopeType                 `json:",omitempty" xml:",omitempty"`
	TestedPersonProperty         []TestedPersonPropertyType         `json:",omitempty" xml:",omitempty"`
	TestedPersonnelClassProperty []TestedPersonnelClassPropertyType `json:",omitempty" xml:",omitempty"`
}

// TestedPersonPropertyType was auto-generated. Do not change.
type TestedPersonPropertyType struct {
	PersonID   PersonIDType   `json:",omitempty" xml:",omitempty"`
	PropertyID PropertyIDType `json:",omitempty" xml:",omitempty"`
}

// TestedPersonnelClassPropertyType was auto-generated. Do not change.
type TestedPersonnelClassPropertyType struct {
	PersonnelClassID PersonnelClassIDType `json:",omitempty" xml:",omitempty"`
	PropertyID       PropertyIDType       `json:",omitempty" xml:",omitempty"`
}

// GetPersonnelInformationType was auto-generated. Do not change.
type GetPersonnelInformationType struct {
	ApplicationArea TransApplicationAreaType            `json:",omitempty" xml:",omitempty"`
	DataArea        GetPersonnelInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                              `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                              `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// GetPersonnelInformationTypeDataArea was auto-generated. Do not change.
type GetPersonnelInformationTypeDataArea struct {
	Get                  []string                   `json:",omitempty" xml:",omitempty"`
	PersonnelInformation []PersonnelInformationType `json:",omitempty" xml:",omitempty"`
}

// ShowPersonnelInformationType was auto-generated. Do not change.
type ShowPersonnelInformationType struct {
	ApplicationArea TransApplicationAreaType             `json:",omitempty" xml:",omitempty"`
	DataArea        ShowPersonnelInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                               `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                               `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ShowPersonnelInformationTypeDataArea was auto-generated. Do not change.
type ShowPersonnelInformationTypeDataArea struct {
	Show                 TransShowType              `json:",omitempty" xml:",omitempty"`
	PersonnelInformation []PersonnelInformationType `json:",omitempty" xml:",omitempty"`
}

// ProcessPersonnelInformationType was auto-generated. Do not change.
type ProcessPersonnelInformationType struct {
	ApplicationArea TransApplicationAreaType                `json:",omitempty" xml:",omitempty"`
	DataArea        ProcessPersonnelInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                  `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                  `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ProcessPersonnelInformationTypeDataArea was auto-generated. Do not change.
type ProcessPersonnelInformationTypeDataArea struct {
	Process              TransProcessType           `json:",omitempty" xml:",omitempty"`
	PersonnelInformation []PersonnelInformationType `json:",omitempty" xml:",omitempty"`
}

// AcknowledgePersonnelInformationType was auto-generated. Do not change.
type AcknowledgePersonnelInformationType struct {
	ApplicationArea TransApplicationAreaType                    `json:",omitempty" xml:",omitempty"`
	DataArea        AcknowledgePersonnelInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                      `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                      `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// AcknowledgePersonnelInformationTypeDataArea was auto-generated. Do not change.
type AcknowledgePersonnelInformationTypeDataArea struct {
	Acknowledge          TransAcknowledgeType       `json:",omitempty" xml:",omitempty"`
	PersonnelInformation []PersonnelInformationType `json:",omitempty" xml:",omitempty"`
}

// ChangePersonnelInformationType was auto-generated. Do not change.
type ChangePersonnelInformationType struct {
	ApplicationArea TransApplicationAreaType               `json:",omitempty" xml:",omitempty"`
	DataArea        ChangePersonnelInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                 `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                 `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ChangePersonnelInformationTypeDataArea was auto-generated. Do not change.
type ChangePersonnelInformationTypeDataArea struct {
	Change               TransChangeType            `json:",omitempty" xml:",omitempty"`
	PersonnelInformation []PersonnelInformationType `json:",omitempty" xml:",omitempty"`
}

// RespondPersonnelInformationType was auto-generated. Do not change.
type RespondPersonnelInformationType struct {
	ApplicationArea TransApplicationAreaType                `json:",omitempty" xml:",omitempty"`
	DataArea        RespondPersonnelInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                  `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                  `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// RespondPersonnelInformationTypeDataArea was auto-generated. Do not change.
type RespondPersonnelInformationTypeDataArea struct {
	Respond              TransRespondType           `json:",omitempty" xml:",omitempty"`
	PersonnelInformation []PersonnelInformationType `json:",omitempty" xml:",omitempty"`
}

// CancelPersonnelInformationType was auto-generated. Do not change.
type CancelPersonnelInformationType struct {
	ApplicationArea TransApplicationAreaType               `json:",omitempty" xml:",omitempty"`
	DataArea        CancelPersonnelInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
}

// CancelPersonnelInformationTypeDataArea was auto-generated. Do not change.
type CancelPersonnelInformationTypeDataArea struct {
	Cancel               []TransActionCriteriaType  `json:",omitempty" xml:",omitempty"`
	PersonnelInformation []PersonnelInformationType `json:",omitempty" xml:",omitempty"`
}

// SyncPersonnelInformationType was auto-generated. Do not change.
type SyncPersonnelInformationType struct {
	ApplicationArea TransApplicationAreaType             `json:",omitempty" xml:",omitempty"`
	DataArea        SyncPersonnelInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                               `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                               `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// SyncPersonnelInformationTypeDataArea was auto-generated. Do not change.
type SyncPersonnelInformationTypeDataArea struct {
	Sync                 []TransActionCriteriaType  `json:",omitempty" xml:",omitempty"`
	PersonnelInformation []PersonnelInformationType `json:",omitempty" xml:",omitempty"`
}

// GetPersonnelClassType was auto-generated. Do not change.
type GetPersonnelClassType struct {
	ApplicationArea TransApplicationAreaType      `json:",omitempty" xml:",omitempty"`
	DataArea        GetPersonnelClassTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                        `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                        `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// GetPersonnelClassTypeDataArea was auto-generated. Do not change.
type GetPersonnelClassTypeDataArea struct {
	Get            []string             `json:",omitempty" xml:",omitempty"`
	PersonnelClass []PersonnelClassType `json:",omitempty" xml:",omitempty"`
}

// ShowPersonnelClassType was auto-generated. Do not change.
type ShowPersonnelClassType struct {
	ApplicationArea TransApplicationAreaType       `json:",omitempty" xml:",omitempty"`
	DataArea        ShowPersonnelClassTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                         `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                         `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ShowPersonnelClassTypeDataArea was auto-generated. Do not change.
type ShowPersonnelClassTypeDataArea struct {
	Show           TransShowType        `json:",omitempty" xml:",omitempty"`
	PersonnelClass []PersonnelClassType `json:",omitempty" xml:",omitempty"`
}

// ProcessPersonnelClassType was auto-generated. Do not change.
type ProcessPersonnelClassType struct {
	ApplicationArea TransApplicationAreaType          `json:",omitempty" xml:",omitempty"`
	DataArea        ProcessPersonnelClassTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                            `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                            `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ProcessPersonnelClassTypeDataArea was auto-generated. Do not change.
type ProcessPersonnelClassTypeDataArea struct {
	Process        TransProcessType     `json:",omitempty" xml:",omitempty"`
	PersonnelClass []PersonnelClassType `json:",omitempty" xml:",omitempty"`
}

// AcknowledgePersonnelClassType was auto-generated. Do not change.
type AcknowledgePersonnelClassType struct {
	ApplicationArea TransApplicationAreaType              `json:",omitempty" xml:",omitempty"`
	DataArea        AcknowledgePersonnelClassTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// AcknowledgePersonnelClassTypeDataArea was auto-generated. Do not change.
type AcknowledgePersonnelClassTypeDataArea struct {
	Acknowledge    TransAcknowledgeType `json:",omitempty" xml:",omitempty"`
	PersonnelClass []PersonnelClassType `json:",omitempty" xml:",omitempty"`
}

// ChangePersonnelClassType was auto-generated. Do not change.
type ChangePersonnelClassType struct {
	ApplicationArea TransApplicationAreaType         `json:",omitempty" xml:",omitempty"`
	DataArea        ChangePersonnelClassTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                           `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                           `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ChangePersonnelClassTypeDataArea was auto-generated. Do not change.
type ChangePersonnelClassTypeDataArea struct {
	Change         TransChangeType      `json:",omitempty" xml:",omitempty"`
	PersonnelClass []PersonnelClassType `json:",omitempty" xml:",omitempty"`
}

// RespondPersonnelClassType was auto-generated. Do not change.
type RespondPersonnelClassType struct {
	ApplicationArea TransApplicationAreaType          `json:",omitempty" xml:",omitempty"`
	DataArea        RespondPersonnelClassTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                            `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                            `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// RespondPersonnelClassTypeDataArea was auto-generated. Do not change.
type RespondPersonnelClassTypeDataArea struct {
	Respond        TransRespondType     `json:",omitempty" xml:",omitempty"`
	PersonnelClass []PersonnelClassType `json:",omitempty" xml:",omitempty"`
}

// CancelPersonnelClassType was auto-generated. Do not change.
type CancelPersonnelClassType struct {
	ApplicationArea TransApplicationAreaType         `json:",omitempty" xml:",omitempty"`
	DataArea        CancelPersonnelClassTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                           `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                           `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// CancelPersonnelClassTypeDataArea was auto-generated. Do not change.
type CancelPersonnelClassTypeDataArea struct {
	Cancel         []TransActionCriteriaType `json:",omitempty" xml:",omitempty"`
	PersonnelClass []PersonnelClassType      `json:",omitempty" xml:",omitempty"`
}

// SyncPersonnelClassType was auto-generated. Do not change.
type SyncPersonnelClassType struct {
	ApplicationArea TransApplicationAreaType       `json:",omitempty" xml:",omitempty"`
	DataArea        SyncPersonnelClassTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                         `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                         `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// SyncPersonnelClassTypeDataArea was auto-generated. Do not change.
type SyncPersonnelClassTypeDataArea struct {
	Sync           []TransActionCriteriaType `json:",omitempty" xml:",omitempty"`
	PersonnelClass []PersonnelClassType      `json:",omitempty" xml:",omitempty"`
}

// GetPersonType was auto-generated. Do not change.
type GetPersonType struct {
	ApplicationArea TransApplicationAreaType `json:",omitempty" xml:",omitempty"`
	DataArea        GetPersonTypeDataArea    `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                   `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                   `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// GetPersonTypeDataArea was auto-generated. Do not change.
type GetPersonTypeDataArea struct {
	Get    []string     `json:",omitempty" xml:",omitempty"`
	Person []PersonType `json:",omitempty" xml:",omitempty"`
}

// ShowPersonType was auto-generated. Do not change.
type ShowPersonType struct {
	ApplicationArea TransApplicationAreaType `json:",omitempty" xml:",omitempty"`
	DataArea        ShowPersonTypeDataArea   `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                   `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                   `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ShowPersonTypeDataArea was auto-generated. Do not change.
type ShowPersonTypeDataArea struct {
	Show   TransShowType `json:",omitempty" xml:",omitempty"`
	Person []PersonType  `json:",omitempty" xml:",omitempty"`
}

// ProcessPersonType was auto-generated. Do not change.
type ProcessPersonType struct {
	ApplicationArea TransApplicationAreaType  `json:",omitempty" xml:",omitempty"`
	DataArea        ProcessPersonTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                    `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                    `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ProcessPersonTypeDataArea was auto-generated. Do not change.
type ProcessPersonTypeDataArea struct {
	Process TransProcessType `json:",omitempty" xml:",omitempty"`
	Person  []PersonType     `json:",omitempty" xml:",omitempty"`
}

// AcknowledgePersonType was auto-generated. Do not change.
type AcknowledgePersonType struct {
	ApplicationArea TransApplicationAreaType      `json:",omitempty" xml:",omitempty"`
	DataArea        AcknowledgePersonTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                        `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                        `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// AcknowledgePersonTypeDataArea was auto-generated. Do not change.
type AcknowledgePersonTypeDataArea struct {
	Acknowledge TransAcknowledgeType `json:",omitempty" xml:",omitempty"`
	Person      []PersonType         `json:",omitempty" xml:",omitempty"`
}

// ChangePersonType was auto-generated. Do not change.
type ChangePersonType struct {
	ApplicationArea TransApplicationAreaType `json:",omitempty" xml:",omitempty"`
	DataArea        ChangePersonTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                   `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                   `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ChangePersonTypeDataArea was auto-generated. Do not change.
type ChangePersonTypeDataArea struct {
	Change TransChangeType `json:",omitempty" xml:",omitempty"`
	Person []PersonType    `json:",omitempty" xml:",omitempty"`
}

// RespondPersonType was auto-generated. Do not change.
type RespondPersonType struct {
	ApplicationArea TransApplicationAreaType  `json:",omitempty" xml:",omitempty"`
	DataArea        RespondPersonTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                    `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                    `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// RespondPersonTypeDataArea was auto-generated. Do not change.
type RespondPersonTypeDataArea struct {
	Respond TransRespondType `json:",omitempty" xml:",omitempty"`
	Person  []PersonType     `json:",omitempty" xml:",omitempty"`
}

// CancelPersonType was auto-generated. Do not change.
type CancelPersonType struct {
	ApplicationArea TransApplicationAreaType `json:",omitempty" xml:",omitempty"`
	DataArea        CancelPersonTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                   `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                   `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// CancelPersonTypeDataArea was auto-generated. Do not change.
type CancelPersonTypeDataArea struct {
	Cancel []TransActionCriteriaType `json:",omitempty" xml:",omitempty"`
	Person []PersonType              `json:",omitempty" xml:",omitempty"`
}

// SyncPersonType was auto-generated. Do not change.
type SyncPersonType struct {
	ApplicationArea TransApplicationAreaType `json:",omitempty" xml:",omitempty"`
	DataArea        SyncPersonTypeDataArea   `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                   `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                   `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// SyncPersonTypeDataArea was auto-generated. Do not change.
type SyncPersonTypeDataArea struct {
	Sync   []TransActionCriteriaType `json:",omitempty" xml:",omitempty"`
	Person []PersonType              `json:",omitempty" xml:",omitempty"`
}

// GetQualificationTestSpecificationType was auto-generated. Do not change.
type GetQualificationTestSpecificationType struct {
	ApplicationArea TransApplicationAreaType                      `json:",omitempty" xml:",omitempty"`
	DataArea        GetQualificationTestSpecificationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                        `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                        `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// GetQualificationTestSpecificationTypeDataArea was auto-generated. Do not change.
type GetQualificationTestSpecificationTypeDataArea struct {
	Get                            []string                             `json:",omitempty" xml:",omitempty"`
	QualificationTestSpecification []QualificationTestSpecificationType `json:",omitempty" xml:",omitempty"`
}

// ShowQualificationTestSpecificationType was auto-generated. Do not change.
type ShowQualificationTestSpecificationType struct {
	ApplicationArea TransApplicationAreaType                       `json:",omitempty" xml:",omitempty"`
	DataArea        ShowQualificationTestSpecificationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                         `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                         `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ShowQualificationTestSpecificationTypeDataArea was auto-generated. Do not change.
type ShowQualificationTestSpecificationTypeDataArea struct {
	Show                           TransShowType                        `json:",omitempty" xml:",omitempty"`
	QualificationTestSpecification []QualificationTestSpecificationType `json:",omitempty" xml:",omitempty"`
}

// ProcessQualificationTestSpecificationType was auto-generated. Do not change.
type ProcessQualificationTestSpecificationType struct {
	ApplicationArea TransApplicationAreaType                          `json:",omitempty" xml:",omitempty"`
	DataArea        ProcessQualificationTestSpecificationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                            `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                            `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ProcessQualificationTestSpecificationTypeDataArea was auto-generated. Do not change.
type ProcessQualificationTestSpecificationTypeDataArea struct {
	Process                        TransProcessType                     `json:",omitempty" xml:",omitempty"`
	QualificationTestSpecification []QualificationTestSpecificationType `json:",omitempty" xml:",omitempty"`
}

// AcknowledgeQualificationTestSpecificationType was auto-generated. Do not change.
type AcknowledgeQualificationTestSpecificationType struct {
	ApplicationArea TransApplicationAreaType                              `json:",omitempty" xml:",omitempty"`
	DataArea        AcknowledgeQualificationTestSpecificationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                                `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                                `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// AcknowledgeQualificationTestSpecificationTypeDataArea was auto-generated. Do not change.
type AcknowledgeQualificationTestSpecificationTypeDataArea struct {
	Acknowledge                    TransAcknowledgeType                 `json:",omitempty" xml:",omitempty"`
	QualificationTestSpecification []QualificationTestSpecificationType `json:",omitempty" xml:",omitempty"`
}

// ChangeQualificationTestSpecificationType was auto-generated. Do not change.
type ChangeQualificationTestSpecificationType struct {
	ApplicationArea TransApplicationAreaType                         `json:",omitempty" xml:",omitempty"`
	DataArea        ChangeQualificationTestSpecificationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                           `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                           `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ChangeQualificationTestSpecificationTypeDataArea was auto-generated. Do not change.
type ChangeQualificationTestSpecificationTypeDataArea struct {
	Change                         TransChangeType                      `json:",omitempty" xml:",omitempty"`
	QualificationTestSpecification []QualificationTestSpecificationType `json:",omitempty" xml:",omitempty"`
}

// RespondQualificationTestSpecificationType was auto-generated. Do not change.
type RespondQualificationTestSpecificationType struct {
	ApplicationArea TransApplicationAreaType                          `json:",omitempty" xml:",omitempty"`
	DataArea        RespondQualificationTestSpecificationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                            `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                            `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// RespondQualificationTestSpecificationTypeDataArea was auto-generated. Do not change.
type RespondQualificationTestSpecificationTypeDataArea struct {
	Respond                        TransRespondType                     `json:",omitempty" xml:",omitempty"`
	QualificationTestSpecification []QualificationTestSpecificationType `json:",omitempty" xml:",omitempty"`
}

// CancelQualificationTestSpecificationType was auto-generated. Do not change.
type CancelQualificationTestSpecificationType struct {
	ApplicationArea TransApplicationAreaType                         `json:",omitempty" xml:",omitempty"`
	DataArea        CancelQualificationTestSpecificationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                           `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                           `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// CancelQualificationTestSpecificationTypeDataArea was auto-generated. Do not change.
type CancelQualificationTestSpecificationTypeDataArea struct {
	Cancel                         []TransActionCriteriaType            `json:",omitempty" xml:",omitempty"`
	QualificationTestSpecification []QualificationTestSpecificationType `json:",omitempty" xml:",omitempty"`
}

// SyncQualificationTestSpecificationType was auto-generated. Do not change.
type SyncQualificationTestSpecificationType struct {
	ApplicationArea TransApplicationAreaType                       `json:",omitempty" xml:",omitempty"`
	DataArea        SyncQualificationTestSpecificationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                         `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                         `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// SyncQualificationTestSpecificationTypeDataArea was auto-generated. Do not change.
type SyncQualificationTestSpecificationTypeDataArea struct {
	Sync                           []TransActionCriteriaType            `json:",omitempty" xml:",omitempty"`
	QualificationTestSpecification []QualificationTestSpecificationType `json:",omitempty" xml:",omitempty"`
}

// PhysicalAssetInformationType was auto-generated. Do not change.
type PhysicalAssetInformationType struct {
	ID                                       IdentifierType                                 `json:",omitempty" xml:",omitempty"`
	Description                              []DescriptionType                              `json:",omitempty" xml:",omitempty"`
	HierarchyScope                           HierarchyScopeType                             `json:",omitempty" xml:",omitempty"`
	PublishedDate                            PublishedDateType                              `json:",omitempty" xml:",omitempty"`
	PhysicalAsset                            []PhysicalAssetType                            `json:",omitempty" xml:",omitempty"`
	PhysicalAssetClass                       []PhysicalAssetClassType                       `json:",omitempty" xml:",omitempty"`
	PhysicalAssetCapabilityTestSpecification []PhysicalAssetCapabilityTestSpecificationType `json:",omitempty" xml:",omitempty"`
}

// PhysicalAssetType was auto-generated. Do not change.
type PhysicalAssetType struct {
	ID                                         IdentifierType                                   `json:",omitempty" xml:",omitempty"`
	Description                                []DescriptionType                                `json:",omitempty" xml:",omitempty"`
	HierarchyScope                             HierarchyScopeType                               `json:",omitempty" xml:",omitempty"`
	PhysicalLocation                           IdentifierType                                   `json:",omitempty" xml:",omitempty"`
	FixedAssetID                               IdentifierType                                   `json:",omitempty" xml:",omitempty"`
	VendorID                                   IdentifierType                                   `json:",omitempty" xml:",omitempty"`
	EquipmentLevel                             HierarchyScopeType                               `json:",omitempty" xml:",omitempty"`
	EquipmentAssetMapping                      []EquipmentAssetMappingType                      `json:",omitempty" xml:",omitempty"`
	PhysicalAssetProperty                      []PhysicalAssetPropertyType                      `json:",omitempty" xml:",omitempty"`
	PhysicalAsset                              []PhysicalAssetType                              `json:",omitempty" xml:",omitempty"`
	PhysicalAssetClassID                       []PhysicalAssetClassIDType                       `json:",omitempty" xml:",omitempty"`
	PhysicalAssetCapabilityTestSpecificationID []PhysicalAssetCapabilityTestSpecificationIDType `json:",omitempty" xml:",omitempty"`
}

// PhysicalAssetPropertyType was auto-generated. Do not change.
type PhysicalAssetPropertyType struct {
	ID                                         IdentifierType                                   `json:",omitempty" xml:",omitempty"`
	Description                                []DescriptionType                                `json:",omitempty" xml:",omitempty"`
	Value                                      []ValueType                                      `json:",omitempty" xml:",omitempty"`
	PhysicalAssetProperty                      []PhysicalAssetPropertyType                      `json:",omitempty" xml:",omitempty"`
	PhysicalAssetCapabilityTestSpecificationID []PhysicalAssetCapabilityTestSpecificationIDType `json:",omitempty" xml:",omitempty"`
	TestResult                                 []TestResultType                                 `json:",omitempty" xml:",omitempty"`
}

// PhysicalAssetClassType was auto-generated. Do not change.
type PhysicalAssetClassType struct {
	ID                                         IdentifierType                                   `json:",omitempty" xml:",omitempty"`
	Description                                []DescriptionType                                `json:",omitempty" xml:",omitempty"`
	HierarchyScope                             HierarchyScopeType                               `json:",omitempty" xml:",omitempty"`
	Manufacturer                               []NameType                                       `json:",omitempty" xml:",omitempty"`
	PhysicalAssetClassProperty                 []PhysicalAssetClassPropertyType                 `json:",omitempty" xml:",omitempty"`
	PhysicalAssetID                            []PhysicalAssetIDType                            `json:",omitempty" xml:",omitempty"`
	PhysicalAssetCapabilityTestSpecificationID []PhysicalAssetCapabilityTestSpecificationIDType `json:",omitempty" xml:",omitempty"`
}

// PhysicalAssetClassPropertyType was auto-generated. Do not change.
type PhysicalAssetClassPropertyType struct {
	ID                                         IdentifierType                                   `json:",omitempty" xml:",omitempty"`
	Description                                []DescriptionType                                `json:",omitempty" xml:",omitempty"`
	Value                                      []ValueType                                      `json:",omitempty" xml:",omitempty"`
	PhysicalAssetClassProperty                 []PhysicalAssetClassPropertyType                 `json:",omitempty" xml:",omitempty"`
	PhysicalAssetCapabilityTestSpecificationID []PhysicalAssetCapabilityTestSpecificationIDType `json:",omitempty" xml:",omitempty"`
}

// PhysicalAssetCapabilityTestSpecificationType was auto-generated. Do not change.
type PhysicalAssetCapabilityTestSpecificationType struct {
	Name                             NameType                               `json:",omitempty" xml:",omitempty"`
	Description                      []DescriptionType                      `json:",omitempty" xml:",omitempty"`
	Version                          VersionType                            `json:",omitempty" xml:",omitempty"`
	HierarchyScope                   HierarchyScopeType                     `json:",omitempty" xml:",omitempty"`
	TestedPhysicalAssetProperty      []TestedPhysicalAssetPropertyType      `json:",omitempty" xml:",omitempty"`
	TestedPhysicalAssetClassProperty []TestedPhysicalAssetClassPropertyType `json:",omitempty" xml:",omitempty"`
}

// TestedPhysicalAssetPropertyType was auto-generated. Do not change.
type TestedPhysicalAssetPropertyType struct {
	PhysicalAssetID PhysicalAssetIDType `json:",omitempty" xml:",omitempty"`
	PropertyID      PropertyIDType      `json:",omitempty" xml:",omitempty"`
}

// TestedPhysicalAssetClassPropertyType was auto-generated. Do not change.
type TestedPhysicalAssetClassPropertyType struct {
	PhysicalAssetClassID PhysicalAssetClassIDType `json:",omitempty" xml:",omitempty"`
	PropertyID           PropertyIDType           `json:",omitempty" xml:",omitempty"`
}

// GetPhysicalAssetInformationType was auto-generated. Do not change.
type GetPhysicalAssetInformationType struct {
	ApplicationArea TransApplicationAreaType                `json:",omitempty" xml:",omitempty"`
	DataArea        GetPhysicalAssetInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                  `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                  `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// GetPhysicalAssetInformationTypeDataArea was auto-generated. Do not change.
type GetPhysicalAssetInformationTypeDataArea struct {
	Get                      []string                       `json:",omitempty" xml:",omitempty"`
	PhysicalAssetInformation []PhysicalAssetInformationType `json:",omitempty" xml:",omitempty"`
}

// ShowPhysicalAssetInformationType was auto-generated. Do not change.
type ShowPhysicalAssetInformationType struct {
	ApplicationArea TransApplicationAreaType                 `json:",omitempty" xml:",omitempty"`
	DataArea        ShowPhysicalAssetInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                   `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                   `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ShowPhysicalAssetInformationTypeDataArea was auto-generated. Do not change.
type ShowPhysicalAssetInformationTypeDataArea struct {
	Show                     TransShowType                  `json:",omitempty" xml:",omitempty"`
	PhysicalAssetInformation []PhysicalAssetInformationType `json:",omitempty" xml:",omitempty"`
}

// ProcessPhysicalAssetInformationType was auto-generated. Do not change.
type ProcessPhysicalAssetInformationType struct {
	ApplicationArea TransApplicationAreaType                    `json:",omitempty" xml:",omitempty"`
	DataArea        ProcessPhysicalAssetInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                      `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                      `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ProcessPhysicalAssetInformationTypeDataArea was auto-generated. Do not change.
type ProcessPhysicalAssetInformationTypeDataArea struct {
	Process                  TransProcessType               `json:",omitempty" xml:",omitempty"`
	PhysicalAssetInformation []PhysicalAssetInformationType `json:",omitempty" xml:",omitempty"`
}

// AcknowledgePhysicalAssetInformationType was auto-generated. Do not change.
type AcknowledgePhysicalAssetInformationType struct {
	ApplicationArea TransApplicationAreaType                        `json:",omitempty" xml:",omitempty"`
	DataArea        AcknowledgePhysicalAssetInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                          `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                          `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// AcknowledgePhysicalAssetInformationTypeDataArea was auto-generated. Do not change.
type AcknowledgePhysicalAssetInformationTypeDataArea struct {
	Acknowledge              TransAcknowledgeType           `json:",omitempty" xml:",omitempty"`
	PhysicalAssetInformation []PhysicalAssetInformationType `json:",omitempty" xml:",omitempty"`
}

// ChangePhysicalAssetInformationType was auto-generated. Do not change.
type ChangePhysicalAssetInformationType struct {
	ApplicationArea TransApplicationAreaType                   `json:",omitempty" xml:",omitempty"`
	DataArea        ChangePhysicalAssetInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                     `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                     `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ChangePhysicalAssetInformationTypeDataArea was auto-generated. Do not change.
type ChangePhysicalAssetInformationTypeDataArea struct {
	Change                   TransChangeType                `json:",omitempty" xml:",omitempty"`
	PhysicalAssetInformation []PhysicalAssetInformationType `json:",omitempty" xml:",omitempty"`
}

// RespondPhysicalAssetInformationType was auto-generated. Do not change.
type RespondPhysicalAssetInformationType struct {
	ApplicationArea TransApplicationAreaType                    `json:",omitempty" xml:",omitempty"`
	DataArea        RespondPhysicalAssetInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                      `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                      `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// RespondPhysicalAssetInformationTypeDataArea was auto-generated. Do not change.
type RespondPhysicalAssetInformationTypeDataArea struct {
	Respond                  TransRespondType               `json:",omitempty" xml:",omitempty"`
	PhysicalAssetInformation []PhysicalAssetInformationType `json:",omitempty" xml:",omitempty"`
}

// CancelPhysicalAssetInformationType was auto-generated. Do not change.
type CancelPhysicalAssetInformationType struct {
	ApplicationArea TransApplicationAreaType                   `json:",omitempty" xml:",omitempty"`
	DataArea        CancelPhysicalAssetInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                     `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                     `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// CancelPhysicalAssetInformationTypeDataArea was auto-generated. Do not change.
type CancelPhysicalAssetInformationTypeDataArea struct {
	Cancel                   []TransActionCriteriaType      `json:",omitempty" xml:",omitempty"`
	PhysicalAssetInformation []PhysicalAssetInformationType `json:",omitempty" xml:",omitempty"`
}

// SyncPhysicalAssetInformationType was auto-generated. Do not change.
type SyncPhysicalAssetInformationType struct {
	ApplicationArea TransApplicationAreaType                 `json:",omitempty" xml:",omitempty"`
	DataArea        SyncPhysicalAssetInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                   `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                   `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// SyncPhysicalAssetInformationTypeDataArea was auto-generated. Do not change.
type SyncPhysicalAssetInformationTypeDataArea struct {
	Sync                     []TransActionCriteriaType      `json:",omitempty" xml:",omitempty"`
	PhysicalAssetInformation []PhysicalAssetInformationType `json:",omitempty" xml:",omitempty"`
}

// GetPhysicalAssetClassType was auto-generated. Do not change.
type GetPhysicalAssetClassType struct {
	ApplicationArea TransApplicationAreaType          `json:",omitempty" xml:",omitempty"`
	DataArea        GetPhysicalAssetClassTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                            `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                            `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// GetPhysicalAssetClassTypeDataArea was auto-generated. Do not change.
type GetPhysicalAssetClassTypeDataArea struct {
	Get                []string                 `json:",omitempty" xml:",omitempty"`
	PhysicalAssetClass []PhysicalAssetClassType `json:",omitempty" xml:",omitempty"`
}

// ShowPhysicalAssetClassType was auto-generated. Do not change.
type ShowPhysicalAssetClassType struct {
	ApplicationArea TransApplicationAreaType           `json:",omitempty" xml:",omitempty"`
	DataArea        ShowPhysicalAssetClassTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                             `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                             `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ShowPhysicalAssetClassTypeDataArea was auto-generated. Do not change.
type ShowPhysicalAssetClassTypeDataArea struct {
	Show               TransShowType            `json:",omitempty" xml:",omitempty"`
	PhysicalAssetClass []PhysicalAssetClassType `json:",omitempty" xml:",omitempty"`
}

// ProcessPhysicalAssetClassType was auto-generated. Do not change.
type ProcessPhysicalAssetClassType struct {
	ApplicationArea TransApplicationAreaType              `json:",omitempty" xml:",omitempty"`
	DataArea        ProcessPhysicalAssetClassTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ProcessPhysicalAssetClassTypeDataArea was auto-generated. Do not change.
type ProcessPhysicalAssetClassTypeDataArea struct {
	Process            TransProcessType         `json:",omitempty" xml:",omitempty"`
	PhysicalAssetClass []PhysicalAssetClassType `json:",omitempty" xml:",omitempty"`
}

// AcknowledgePhysicalAssetClassType was auto-generated. Do not change.
type AcknowledgePhysicalAssetClassType struct {
	ApplicationArea TransApplicationAreaType                  `json:",omitempty" xml:",omitempty"`
	DataArea        AcknowledgePhysicalAssetClassTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                    `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                    `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// AcknowledgePhysicalAssetClassTypeDataArea was auto-generated. Do not change.
type AcknowledgePhysicalAssetClassTypeDataArea struct {
	Acknowledge        TransAcknowledgeType     `json:",omitempty" xml:",omitempty"`
	PhysicalAssetClass []PhysicalAssetClassType `json:",omitempty" xml:",omitempty"`
}

// ChangePhysicalAssetClassType was auto-generated. Do not change.
type ChangePhysicalAssetClassType struct {
	ApplicationArea TransApplicationAreaType             `json:",omitempty" xml:",omitempty"`
	DataArea        ChangePhysicalAssetClassTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                               `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                               `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ChangePhysicalAssetClassTypeDataArea was auto-generated. Do not change.
type ChangePhysicalAssetClassTypeDataArea struct {
	Change             TransChangeType          `json:",omitempty" xml:",omitempty"`
	PhysicalAssetClass []PhysicalAssetClassType `json:",omitempty" xml:",omitempty"`
}

// RespondPhysicalAssetClassType was auto-generated. Do not change.
type RespondPhysicalAssetClassType struct {
	ApplicationArea TransApplicationAreaType              `json:",omitempty" xml:",omitempty"`
	DataArea        RespondPhysicalAssetClassTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// RespondPhysicalAssetClassTypeDataArea was auto-generated. Do not change.
type RespondPhysicalAssetClassTypeDataArea struct {
	Respond            TransRespondType         `json:",omitempty" xml:",omitempty"`
	PhysicalAssetClass []PhysicalAssetClassType `json:",omitempty" xml:",omitempty"`
}

// CancelPhysicalAssetClassType was auto-generated. Do not change.
type CancelPhysicalAssetClassType struct {
	ApplicationArea TransApplicationAreaType             `json:",omitempty" xml:",omitempty"`
	DataArea        CancelPhysicalAssetClassTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                               `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                               `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// CancelPhysicalAssetClassTypeDataArea was auto-generated. Do not change.
type CancelPhysicalAssetClassTypeDataArea struct {
	Cancel             []TransActionCriteriaType `json:",omitempty" xml:",omitempty"`
	PhysicalAssetClass []PhysicalAssetClassType  `json:",omitempty" xml:",omitempty"`
}

// SyncPhysicalAssetClassType was auto-generated. Do not change.
type SyncPhysicalAssetClassType struct {
	ApplicationArea TransApplicationAreaType           `json:",omitempty" xml:",omitempty"`
	DataArea        SyncPhysicalAssetClassTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                             `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                             `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// SyncPhysicalAssetClassTypeDataArea was auto-generated. Do not change.
type SyncPhysicalAssetClassTypeDataArea struct {
	Sync               []TransActionCriteriaType `json:",omitempty" xml:",omitempty"`
	PhysicalAssetClass []PhysicalAssetClassType  `json:",omitempty" xml:",omitempty"`
}

// GetPhysicalAssetType was auto-generated. Do not change.
type GetPhysicalAssetType struct {
	ApplicationArea TransApplicationAreaType     `json:",omitempty" xml:",omitempty"`
	DataArea        GetPhysicalAssetTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                       `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                       `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// GetPhysicalAssetTypeDataArea was auto-generated. Do not change.
type GetPhysicalAssetTypeDataArea struct {
	Get           []string            `json:",omitempty" xml:",omitempty"`
	PhysicalAsset []PhysicalAssetType `json:",omitempty" xml:",omitempty"`
}

// ShowPhysicalAssetType was auto-generated. Do not change.
type ShowPhysicalAssetType struct {
	ApplicationArea TransApplicationAreaType      `json:",omitempty" xml:",omitempty"`
	DataArea        ShowPhysicalAssetTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                        `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                        `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ShowPhysicalAssetTypeDataArea was auto-generated. Do not change.
type ShowPhysicalAssetTypeDataArea struct {
	Show          TransShowType       `json:",omitempty" xml:",omitempty"`
	PhysicalAsset []PhysicalAssetType `json:",omitempty" xml:",omitempty"`
}

// ProcessPhysicalAssetType was auto-generated. Do not change.
type ProcessPhysicalAssetType struct {
	ApplicationArea TransApplicationAreaType         `json:",omitempty" xml:",omitempty"`
	DataArea        ProcessPhysicalAssetTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                           `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                           `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ProcessPhysicalAssetTypeDataArea was auto-generated. Do not change.
type ProcessPhysicalAssetTypeDataArea struct {
	Process       TransProcessType    `json:",omitempty" xml:",omitempty"`
	PhysicalAsset []PhysicalAssetType `json:",omitempty" xml:",omitempty"`
}

// AcknowledgePhysicalAssetType was auto-generated. Do not change.
type AcknowledgePhysicalAssetType struct {
	ApplicationArea TransApplicationAreaType             `json:",omitempty" xml:",omitempty"`
	DataArea        AcknowledgePhysicalAssetTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                               `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                               `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// AcknowledgePhysicalAssetTypeDataArea was auto-generated. Do not change.
type AcknowledgePhysicalAssetTypeDataArea struct {
	Acknowledge   TransAcknowledgeType `json:",omitempty" xml:",omitempty"`
	PhysicalAsset []PhysicalAssetType  `json:",omitempty" xml:",omitempty"`
}

// ChangePhysicalAssetType was auto-generated. Do not change.
type ChangePhysicalAssetType struct {
	ApplicationArea TransApplicationAreaType        `json:",omitempty" xml:",omitempty"`
	DataArea        ChangePhysicalAssetTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                          `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                          `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ChangePhysicalAssetTypeDataArea was auto-generated. Do not change.
type ChangePhysicalAssetTypeDataArea struct {
	Change        TransChangeType     `json:",omitempty" xml:",omitempty"`
	PhysicalAsset []PhysicalAssetType `json:",omitempty" xml:",omitempty"`
}

// RespondPhysicalAssetType was auto-generated. Do not change.
type RespondPhysicalAssetType struct {
	ApplicationArea TransApplicationAreaType         `json:",omitempty" xml:",omitempty"`
	DataArea        RespondPhysicalAssetTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                           `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                           `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// RespondPhysicalAssetTypeDataArea was auto-generated. Do not change.
type RespondPhysicalAssetTypeDataArea struct {
	Respond       TransRespondType    `json:",omitempty" xml:",omitempty"`
	PhysicalAsset []PhysicalAssetType `json:",omitempty" xml:",omitempty"`
}

// CancelPhysicalAssetType was auto-generated. Do not change.
type CancelPhysicalAssetType struct {
	ApplicationArea TransApplicationAreaType        `json:",omitempty" xml:",omitempty"`
	DataArea        CancelPhysicalAssetTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                          `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                          `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// CancelPhysicalAssetTypeDataArea was auto-generated. Do not change.
type CancelPhysicalAssetTypeDataArea struct {
	Cancel        []TransActionCriteriaType `json:",omitempty" xml:",omitempty"`
	PhysicalAsset []PhysicalAssetType       `json:",omitempty" xml:",omitempty"`
}

// SyncPhysicalAssetType was auto-generated. Do not change.
type SyncPhysicalAssetType struct {
	ApplicationArea TransApplicationAreaType      `json:",omitempty" xml:",omitempty"`
	DataArea        SyncPhysicalAssetTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                        `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                        `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// SyncPhysicalAssetTypeDataArea was auto-generated. Do not change.
type SyncPhysicalAssetTypeDataArea struct {
	Sync          []TransActionCriteriaType `json:",omitempty" xml:",omitempty"`
	PhysicalAsset []PhysicalAssetType       `json:",omitempty" xml:",omitempty"`
}

// GetPhysicalAssetCapabilityTestSpecType was auto-generated. Do not change.
type GetPhysicalAssetCapabilityTestSpecType struct {
	ApplicationArea TransApplicationAreaType                       `json:",omitempty" xml:",omitempty"`
	DataArea        GetPhysicalAssetCapabilityTestSpecTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                         `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                         `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// GetPhysicalAssetCapabilityTestSpecTypeDataArea was auto-generated. Do not change.
type GetPhysicalAssetCapabilityTestSpecTypeDataArea struct {
	Get                             []string                                       `json:",omitempty" xml:",omitempty"`
	PhysicalAssetCapabilityTestSpec []PhysicalAssetCapabilityTestSpecificationType `json:",omitempty" xml:",omitempty"`
}

// ShowPhysicalAssetCapabilityTestSpecType was auto-generated. Do not change.
type ShowPhysicalAssetCapabilityTestSpecType struct {
	ApplicationArea TransApplicationAreaType                        `json:",omitempty" xml:",omitempty"`
	DataArea        ShowPhysicalAssetCapabilityTestSpecTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                          `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                          `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ShowPhysicalAssetCapabilityTestSpecTypeDataArea was auto-generated. Do not change.
type ShowPhysicalAssetCapabilityTestSpecTypeDataArea struct {
	Show                            TransShowType                                  `json:",omitempty" xml:",omitempty"`
	PhysicalAssetCapabilityTestSpec []PhysicalAssetCapabilityTestSpecificationType `json:",omitempty" xml:",omitempty"`
}

// ProcessPhysicalAssetCapabilityTestSpecType was auto-generated. Do not change.
type ProcessPhysicalAssetCapabilityTestSpecType struct {
	ApplicationArea TransApplicationAreaType                           `json:",omitempty" xml:",omitempty"`
	DataArea        ProcessPhysicalAssetCapabilityTestSpecTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                             `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                             `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ProcessPhysicalAssetCapabilityTestSpecTypeDataArea was auto-generated. Do not change.
type ProcessPhysicalAssetCapabilityTestSpecTypeDataArea struct {
	Process                         TransProcessType                               `json:",omitempty" xml:",omitempty"`
	PhysicalAssetCapabilityTestSpec []PhysicalAssetCapabilityTestSpecificationType `json:",omitempty" xml:",omitempty"`
}

// AcknowledgePhysicalAssetCapabilityTestSpecType was auto-generated. Do not change.
type AcknowledgePhysicalAssetCapabilityTestSpecType struct {
	ApplicationArea TransApplicationAreaType                               `json:",omitempty" xml:",omitempty"`
	DataArea        AcknowledgePhysicalAssetCapabilityTestSpecTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                                 `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                                 `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// AcknowledgePhysicalAssetCapabilityTestSpecTypeDataArea was auto-generated. Do not change.
type AcknowledgePhysicalAssetCapabilityTestSpecTypeDataArea struct {
	Acknowledge                     TransAcknowledgeType                           `json:",omitempty" xml:",omitempty"`
	PhysicalAssetCapabilityTestSpec []PhysicalAssetCapabilityTestSpecificationType `json:",omitempty" xml:",omitempty"`
}

// ChangePhysicalAssetCapabilityTestSpecType was auto-generated. Do not change.
type ChangePhysicalAssetCapabilityTestSpecType struct {
	ApplicationArea TransApplicationAreaType                          `json:",omitempty" xml:",omitempty"`
	DataArea        ChangePhysicalAssetCapabilityTestSpecTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                            `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                            `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ChangePhysicalAssetCapabilityTestSpecTypeDataArea was auto-generated. Do not change.
type ChangePhysicalAssetCapabilityTestSpecTypeDataArea struct {
	Change                          TransChangeType                                `json:",omitempty" xml:",omitempty"`
	PhysicalAssetCapabilityTestSpec []PhysicalAssetCapabilityTestSpecificationType `json:",omitempty" xml:",omitempty"`
}

// RespondPhysicalAssetCapabilityTestSpecType was auto-generated. Do not change.
type RespondPhysicalAssetCapabilityTestSpecType struct {
	ApplicationArea TransApplicationAreaType                           `json:",omitempty" xml:",omitempty"`
	DataArea        RespondPhysicalAssetCapabilityTestSpecTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                             `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                             `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// RespondPhysicalAssetCapabilityTestSpecTypeDataArea was auto-generated. Do not change.
type RespondPhysicalAssetCapabilityTestSpecTypeDataArea struct {
	Respond                         TransRespondType                               `json:",omitempty" xml:",omitempty"`
	PhysicalAssetCapabilityTestSpec []PhysicalAssetCapabilityTestSpecificationType `json:",omitempty" xml:",omitempty"`
}

// CancelPhysicalAssetCapabilityTestSpecType was auto-generated. Do not change.
type CancelPhysicalAssetCapabilityTestSpecType struct {
	ApplicationArea TransApplicationAreaType                          `json:",omitempty" xml:",omitempty"`
	DataArea        CancelPhysicalAssetCapabilityTestSpecTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                            `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                            `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// CancelPhysicalAssetCapabilityTestSpecTypeDataArea was auto-generated. Do not change.
type CancelPhysicalAssetCapabilityTestSpecTypeDataArea struct {
	Cancel                          []TransActionCriteriaType                      `json:",omitempty" xml:",omitempty"`
	PhysicalAssetCapabilityTestSpec []PhysicalAssetCapabilityTestSpecificationType `json:",omitempty" xml:",omitempty"`
}

// SyncPhysicalAssetCapabilityTestSpecType was auto-generated. Do not change.
type SyncPhysicalAssetCapabilityTestSpecType struct {
	ApplicationArea TransApplicationAreaType                        `json:",omitempty" xml:",omitempty"`
	DataArea        SyncPhysicalAssetCapabilityTestSpecTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                          `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                          `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// SyncPhysicalAssetCapabilityTestSpecTypeDataArea was auto-generated. Do not change.
type SyncPhysicalAssetCapabilityTestSpecTypeDataArea struct {
	Sync                            []TransActionCriteriaType                      `json:",omitempty" xml:",omitempty"`
	PhysicalAssetCapabilityTestSpec []PhysicalAssetCapabilityTestSpecificationType `json:",omitempty" xml:",omitempty"`
}

// ProcessSegmentInformationType was auto-generated. Do not change.
type ProcessSegmentInformationType struct {
	ID             IdentifierType       `json:",omitempty" xml:",omitempty"`
	Description    []DescriptionType    `json:",omitempty" xml:",omitempty"`
	Location       LocationType         `json:",omitempty" xml:",omitempty"`
	HierarchyScope HierarchyScopeType   `json:",omitempty" xml:",omitempty"`
	PublishedDate  PublishedDateType    `json:",omitempty" xml:",omitempty"`
	ProcessSegment []ProcessSegmentType `json:",omitempty" xml:",omitempty"`
}

// ProcessSegmentType was auto-generated. Do not change.
type ProcessSegmentType struct {
	ID                                IdentifierType                          `json:",omitempty" xml:",omitempty"`
	Description                       []DescriptionType                       `json:",omitempty" xml:",omitempty"`
	OperationsType                    OperationsTypeType                      `json:",omitempty" xml:",omitempty"`
	Location                          LocationType                            `json:",omitempty" xml:",omitempty"`
	HierarchyScope                    HierarchyScopeType                      `json:",omitempty" xml:",omitempty"`
	PublishedDate                     PublishedDateType                       `json:",omitempty" xml:",omitempty"`
	Duration                          string                                  `json:",omitempty" xml:",omitempty"`
	PersonnelSegmentSpecification     []PersonnelSegmentSpecificationType     `json:",omitempty" xml:",omitempty"`
	EquipmentSegmentSpecification     []EquipmentSegmentSpecificationType     `json:",omitempty" xml:",omitempty"`
	PhysicalAssetSegmentSpecification []PhysicalAssetSegmentSpecificationType `json:",omitempty" xml:",omitempty"`
	MaterialSegmentSpecification      []MaterialSegmentSpecificationType      `json:",omitempty" xml:",omitempty"`
	Parameter                         []ParameterType                         `json:",omitempty" xml:",omitempty"`
	SegmentDependency                 []SegmentDependencyType                 `json:",omitempty" xml:",omitempty"`
	ProcessSegment                    []ProcessSegmentType                    `json:",omitempty" xml:",omitempty"`
}

// PersonnelSegmentSpecificationType was auto-generated. Do not change.
type PersonnelSegmentSpecificationType struct {
	PersonnelClassID                      PersonnelClassIDType                        `json:",omitempty" xml:",omitempty"`
	PersonID                              PersonIDType                                `json:",omitempty" xml:",omitempty"`
	Description                           []DescriptionType                           `json:",omitempty" xml:",omitempty"`
	PersonnelUse                          CodeType                                    `json:",omitempty" xml:",omitempty"`
	Quantity                              []QuantityValueType                         `json:",omitempty" xml:",omitempty"`
	PersonnelSegmentSpecificationProperty []PersonnelSegmentSpecificationPropertyType `json:",omitempty" xml:",omitempty"`
}

// PersonnelSegmentSpecificationPropertyType was auto-generated. Do not change.
type PersonnelSegmentSpecificationPropertyType struct {
	ID          IdentifierType      `json:",omitempty" xml:",omitempty"`
	Description []DescriptionType   `json:",omitempty" xml:",omitempty"`
	Value       []ValueType         `json:",omitempty" xml:",omitempty"`
	Quantity    []QuantityValueType `json:",omitempty" xml:",omitempty"`
}

// EquipmentSegmentSpecificationType was auto-generated. Do not change.
type EquipmentSegmentSpecificationType struct {
	EquipmentClassID                      EquipmentClassIDType                        `json:",omitempty" xml:",omitempty"`
	EquipmentID                           EquipmentIDType                             `json:",omitempty" xml:",omitempty"`
	Description                           []DescriptionType                           `json:",omitempty" xml:",omitempty"`
	EquipmentUse                          CodeType                                    `json:",omitempty" xml:",omitempty"`
	Quantity                              []QuantityValueType                         `json:",omitempty" xml:",omitempty"`
	EquipmentSegmentSpecificationProperty []EquipmentSegmentSpecificationPropertyType `json:",omitempty" xml:",omitempty"`
}

// EquipmentSegmentSpecificationPropertyType was auto-generated. Do not change.
type EquipmentSegmentSpecificationPropertyType struct {
	ID          IdentifierType      `json:",omitempty" xml:",omitempty"`
	Description []DescriptionType   `json:",omitempty" xml:",omitempty"`
	Value       []ValueType         `json:",omitempty" xml:",omitempty"`
	Quantity    []QuantityValueType `json:",omitempty" xml:",omitempty"`
}

// PhysicalAssetSegmentSpecificationType was auto-generated. Do not change.
type PhysicalAssetSegmentSpecificationType struct {
	PhysicalAssetClassID                      PhysicalAssetClassIDType                        `json:",omitempty" xml:",omitempty"`
	PhysicalAssetID                           PhysicalAssetIDType                             `json:",omitempty" xml:",omitempty"`
	Description                               []DescriptionType                               `json:",omitempty" xml:",omitempty"`
	PhysicalAssetUse                          CodeType                                        `json:",omitempty" xml:",omitempty"`
	Quantity                                  []QuantityValueType                             `json:",omitempty" xml:",omitempty"`
	PhysicalAssetSegmentSpecificationProperty []PhysicalAssetSegmentSpecificationPropertyType `json:",omitempty" xml:",omitempty"`
}

// PhysicalAssetSegmentSpecificationPropertyType was auto-generated. Do not change.
type PhysicalAssetSegmentSpecificationPropertyType struct {
	ID          IdentifierType      `json:",omitempty" xml:",omitempty"`
	Description []DescriptionType   `json:",omitempty" xml:",omitempty"`
	Value       []ValueType         `json:",omitempty" xml:",omitempty"`
	Quantity    []QuantityValueType `json:",omitempty" xml:",omitempty"`
}

// MaterialSegmentSpecificationType was auto-generated. Do not change.
type MaterialSegmentSpecificationType struct {
	ID                                   IdentifierType                             `json:",omitempty" xml:",omitempty"`
	MaterialClassID                      MaterialClassIDType                        `json:",omitempty" xml:",omitempty"`
	MaterialDefinitionID                 MaterialDefinitionIDType                   `json:",omitempty" xml:",omitempty"`
	Description                          []DescriptionType                          `json:",omitempty" xml:",omitempty"`
	AssemblyType                         AssemblyTypeType                           `json:",omitempty" xml:",omitempty"`
	AssemblyRelationship                 AssemblyRelationshipType                   `json:",omitempty" xml:",omitempty"`
	AssemblySpecificationID              []IdentifierType                           `json:",omitempty" xml:",omitempty"`
	MaterialUse                          MaterialUseType                            `json:",omitempty" xml:",omitempty"`
	Quantity                             []QuantityValueType                        `json:",omitempty" xml:",omitempty"`
	MaterialSegmentSpecificationProperty []MaterialSegmentSpecificationPropertyType `json:",omitempty" xml:",omitempty"`
}

// MaterialSegmentSpecificationPropertyType was auto-generated. Do not change.
type MaterialSegmentSpecificationPropertyType struct {
	ID          IdentifierType      `json:",omitempty" xml:",omitempty"`
	Description []DescriptionType   `json:",omitempty" xml:",omitempty"`
	Value       []ValueType         `json:",omitempty" xml:",omitempty"`
	Quantity    []QuantityValueType `json:",omitempty" xml:",omitempty"`
}

// GetProcessSegmentInformationType was auto-generated. Do not change.
type GetProcessSegmentInformationType struct {
	ApplicationArea TransApplicationAreaType                 `json:",omitempty" xml:",omitempty"`
	DataArea        GetProcessSegmentInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                   `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                   `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// GetProcessSegmentInformationTypeDataArea was auto-generated. Do not change.
type GetProcessSegmentInformationTypeDataArea struct {
	Get                       []string                        `json:",omitempty" xml:",omitempty"`
	ProcessSegmentInformation []ProcessSegmentInformationType `json:",omitempty" xml:",omitempty"`
}

// ShowProcessSegmentInformationType was auto-generated. Do not change.
type ShowProcessSegmentInformationType struct {
	ApplicationArea TransApplicationAreaType                  `json:",omitempty" xml:",omitempty"`
	DataArea        ShowProcessSegmentInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                    `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                    `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ShowProcessSegmentInformationTypeDataArea was auto-generated. Do not change.
type ShowProcessSegmentInformationTypeDataArea struct {
	Show                      TransShowType                   `json:",omitempty" xml:",omitempty"`
	ProcessSegmentInformation []ProcessSegmentInformationType `json:",omitempty" xml:",omitempty"`
}

// ProcessProcessSegmentInformationType was auto-generated. Do not change.
type ProcessProcessSegmentInformationType struct {
	ApplicationArea TransApplicationAreaType                     `json:",omitempty" xml:",omitempty"`
	DataArea        ProcessProcessSegmentInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                       `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                       `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ProcessProcessSegmentInformationTypeDataArea was auto-generated. Do not change.
type ProcessProcessSegmentInformationTypeDataArea struct {
	Process                   TransProcessType                `json:",omitempty" xml:",omitempty"`
	ProcessSegmentInformation []ProcessSegmentInformationType `json:",omitempty" xml:",omitempty"`
}

// AcknowledgeProcessSegmentInformationType was auto-generated. Do not change.
type AcknowledgeProcessSegmentInformationType struct {
	ApplicationArea TransApplicationAreaType                         `json:",omitempty" xml:",omitempty"`
	DataArea        AcknowledgeProcessSegmentInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                           `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                           `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// AcknowledgeProcessSegmentInformationTypeDataArea was auto-generated. Do not change.
type AcknowledgeProcessSegmentInformationTypeDataArea struct {
	Acknowledge               TransAcknowledgeType            `json:",omitempty" xml:",omitempty"`
	ProcessSegmentInformation []ProcessSegmentInformationType `json:",omitempty" xml:",omitempty"`
}

// ChangeProcessSegmentInformationType was auto-generated. Do not change.
type ChangeProcessSegmentInformationType struct {
	ApplicationArea TransApplicationAreaType                    `json:",omitempty" xml:",omitempty"`
	DataArea        ChangeProcessSegmentInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                      `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                      `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ChangeProcessSegmentInformationTypeDataArea was auto-generated. Do not change.
type ChangeProcessSegmentInformationTypeDataArea struct {
	Change                    TransChangeType                 `json:",omitempty" xml:",omitempty"`
	ProcessSegmentInformation []ProcessSegmentInformationType `json:",omitempty" xml:",omitempty"`
}

// RespondProcessSegmentInformationType was auto-generated. Do not change.
type RespondProcessSegmentInformationType struct {
	ApplicationArea TransApplicationAreaType                     `json:",omitempty" xml:",omitempty"`
	DataArea        RespondProcessSegmentInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                       `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                       `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// RespondProcessSegmentInformationTypeDataArea was auto-generated. Do not change.
type RespondProcessSegmentInformationTypeDataArea struct {
	Respond                   TransRespondType                `json:",omitempty" xml:",omitempty"`
	ProcessSegmentInformation []ProcessSegmentInformationType `json:",omitempty" xml:",omitempty"`
}

// CancelProcessSegmentInformationType was auto-generated. Do not change.
type CancelProcessSegmentInformationType struct {
	ApplicationArea TransApplicationAreaType                    `json:",omitempty" xml:",omitempty"`
	DataArea        CancelProcessSegmentInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                      `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                      `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// CancelProcessSegmentInformationTypeDataArea was auto-generated. Do not change.
type CancelProcessSegmentInformationTypeDataArea struct {
	Cancel                    []TransActionCriteriaType       `json:",omitempty" xml:",omitempty"`
	ProcessSegmentInformation []ProcessSegmentInformationType `json:",omitempty" xml:",omitempty"`
}

// SyncProcessSegmentInformationType was auto-generated. Do not change.
type SyncProcessSegmentInformationType struct {
	ApplicationArea TransApplicationAreaType                  `json:",omitempty" xml:",omitempty"`
	DataArea        SyncProcessSegmentInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                    `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                    `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// SyncProcessSegmentInformationTypeDataArea was auto-generated. Do not change.
type SyncProcessSegmentInformationTypeDataArea struct {
	Sync                      []TransActionCriteriaType       `json:",omitempty" xml:",omitempty"`
	ProcessSegmentInformation []ProcessSegmentInformationType `json:",omitempty" xml:",omitempty"`
}

// GetProcessSegmentType was auto-generated. Do not change.
type GetProcessSegmentType struct {
	ApplicationArea TransApplicationAreaType      `json:",omitempty" xml:",omitempty"`
	DataArea        GetProcessSegmentTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                        `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                        `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// GetProcessSegmentTypeDataArea was auto-generated. Do not change.
type GetProcessSegmentTypeDataArea struct {
	Get            []string             `json:",omitempty" xml:",omitempty"`
	ProcessSegment []ProcessSegmentType `json:",omitempty" xml:",omitempty"`
}

// ShowProcessSegmentType was auto-generated. Do not change.
type ShowProcessSegmentType struct {
	ApplicationArea TransApplicationAreaType       `json:",omitempty" xml:",omitempty"`
	DataArea        ShowProcessSegmentTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                         `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                         `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ShowProcessSegmentTypeDataArea was auto-generated. Do not change.
type ShowProcessSegmentTypeDataArea struct {
	Show           TransShowType        `json:",omitempty" xml:",omitempty"`
	ProcessSegment []ProcessSegmentType `json:",omitempty" xml:",omitempty"`
}

// ProcessProcessSegmentType was auto-generated. Do not change.
type ProcessProcessSegmentType struct {
	ApplicationArea TransApplicationAreaType          `json:",omitempty" xml:",omitempty"`
	DataArea        ProcessProcessSegmentTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                            `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                            `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ProcessProcessSegmentTypeDataArea was auto-generated. Do not change.
type ProcessProcessSegmentTypeDataArea struct {
	Process        TransProcessType     `json:",omitempty" xml:",omitempty"`
	ProcessSegment []ProcessSegmentType `json:",omitempty" xml:",omitempty"`
}

// AcknowledgeProcessSegmentType was auto-generated. Do not change.
type AcknowledgeProcessSegmentType struct {
	ApplicationArea TransApplicationAreaType              `json:",omitempty" xml:",omitempty"`
	DataArea        AcknowledgeProcessSegmentTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// AcknowledgeProcessSegmentTypeDataArea was auto-generated. Do not change.
type AcknowledgeProcessSegmentTypeDataArea struct {
	Acknowledge    TransAcknowledgeType `json:",omitempty" xml:",omitempty"`
	ProcessSegment []ProcessSegmentType `json:",omitempty" xml:",omitempty"`
}

// ChangeProcessSegmentType was auto-generated. Do not change.
type ChangeProcessSegmentType struct {
	ApplicationArea TransApplicationAreaType         `json:",omitempty" xml:",omitempty"`
	DataArea        ChangeProcessSegmentTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                           `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                           `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ChangeProcessSegmentTypeDataArea was auto-generated. Do not change.
type ChangeProcessSegmentTypeDataArea struct {
	Change         TransChangeType      `json:",omitempty" xml:",omitempty"`
	ProcessSegment []ProcessSegmentType `json:",omitempty" xml:",omitempty"`
}

// RespondProcessSegmentType was auto-generated. Do not change.
type RespondProcessSegmentType struct {
	ApplicationArea TransApplicationAreaType          `json:",omitempty" xml:",omitempty"`
	DataArea        RespondProcessSegmentTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                            `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                            `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// RespondProcessSegmentTypeDataArea was auto-generated. Do not change.
type RespondProcessSegmentTypeDataArea struct {
	Respond        TransRespondType     `json:",omitempty" xml:",omitempty"`
	ProcessSegment []ProcessSegmentType `json:",omitempty" xml:",omitempty"`
}

// CancelProcessSegmentType was auto-generated. Do not change.
type CancelProcessSegmentType struct {
	ApplicationArea TransApplicationAreaType         `json:",omitempty" xml:",omitempty"`
	DataArea        CancelProcessSegmentTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                           `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                           `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// CancelProcessSegmentTypeDataArea was auto-generated. Do not change.
type CancelProcessSegmentTypeDataArea struct {
	Cancel         []TransActionCriteriaType `json:",omitempty" xml:",omitempty"`
	ProcessSegment []ProcessSegmentType      `json:",omitempty" xml:",omitempty"`
}

// SyncProcessSegmentType was auto-generated. Do not change.
type SyncProcessSegmentType struct {
	ApplicationArea TransApplicationAreaType       `json:",omitempty" xml:",omitempty"`
	DataArea        SyncProcessSegmentTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                         `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                         `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// SyncProcessSegmentTypeDataArea was auto-generated. Do not change.
type SyncProcessSegmentTypeDataArea struct {
	Sync           []TransActionCriteriaType `json:",omitempty" xml:",omitempty"`
	ProcessSegment []ProcessSegmentType      `json:",omitempty" xml:",omitempty"`
}

// ProductInformationType was auto-generated. Do not change.
type ProductInformationType struct {
	ID                IdentifierType          `json:",omitempty" xml:",omitempty"`
	Description       []DescriptionType       `json:",omitempty" xml:",omitempty"`
	Location          LocationType            `json:",omitempty" xml:",omitempty"`
	HierarchyScope    HierarchyScopeType      `json:",omitempty" xml:",omitempty"`
	PublishedDate     PublishedDateType       `json:",omitempty" xml:",omitempty"`
	ProductDefinition []ProductDefinitionType `json:",omitempty" xml:",omitempty"`
}

// ProductDefinitionType was auto-generated. Do not change.
type ProductDefinitionType struct {
	ID                    IdentifierType            `json:",omitempty" xml:",omitempty"`
	Version               VersionType               `json:",omitempty" xml:",omitempty"`
	Description           []DescriptionType         `json:",omitempty" xml:",omitempty"`
	Location              LocationType              `json:",omitempty" xml:",omitempty"`
	HierarchyScope        HierarchyScopeType        `json:",omitempty" xml:",omitempty"`
	PublishedDate         PublishedDateType         `json:",omitempty" xml:",omitempty"`
	ProductProductionRule ProductProductionRuleType `json:",omitempty" xml:",omitempty"`
	BillOfMaterialsID     BillOfMaterialsIDType     `json:",omitempty" xml:",omitempty"`
	BillOfResourcesID     BillOfResourcesIDType     `json:",omitempty" xml:",omitempty"`
	ManufacturingBill     []ManufacturingBillType   `json:",omitempty" xml:",omitempty"`
	ProductSegment        []ProductSegmentType      `json:",omitempty" xml:",omitempty"`
}

// ManufacturingBillType was auto-generated. Do not change.
type ManufacturingBillType struct {
	ID                        IdentifierType           `json:",omitempty" xml:",omitempty"`
	Description               DescriptionType          `json:",omitempty" xml:",omitempty"`
	MaterialClassID           MaterialClassIDType      `json:",omitempty" xml:",omitempty"`
	MaterialDefinitionID      MaterialDefinitionIDType `json:",omitempty" xml:",omitempty"`
	Quantity                  []QuantityValueType      `json:",omitempty" xml:",omitempty"`
	AssemblyManufacturingBill []ManufacturingBillType  `json:",omitempty" xml:",omitempty"`
	AssemblyType              AssemblyTypeType         `json:",omitempty" xml:",omitempty"`
	AssemblyRelationship      AssemblyRelationshipType `json:",omitempty" xml:",omitempty"`
	BillOfMaterialID          BillOfMaterialIDType     `json:",omitempty" xml:",omitempty"`
}

// ProductSegmentType was auto-generated. Do not change.
type ProductSegmentType struct {
	ID                         IdentifierType                   `json:",omitempty" xml:",omitempty"`
	Description                DescriptionType                  `json:",omitempty" xml:",omitempty"`
	Duration                   string                           `json:",omitempty" xml:",omitempty"`
	ProcessSegmentID           []ProcessSegmentIDType           `json:",omitempty" xml:",omitempty"`
	Parameter                  []ParameterType                  `json:",omitempty" xml:",omitempty"`
	PersonnelSpecification     []PersonnelSpecificationType     `json:",omitempty" xml:",omitempty"`
	EquipmentSpecification     []EquipmentSpecificationType     `json:",omitempty" xml:",omitempty"`
	PhysicalAssetSpecification []PhysicalAssetSpecificationType `json:",omitempty" xml:",omitempty"`
	MaterialSpecification      []MaterialSpecificationType      `json:",omitempty" xml:",omitempty"`
	SegmentDependency          []SegmentDependencyType          `json:",omitempty" xml:",omitempty"`
	ProductSegment             []ProductSegmentType             `json:",omitempty" xml:",omitempty"`
}

// PersonnelSpecificationType was auto-generated. Do not change.
type PersonnelSpecificationType struct {
	PersonnelClassID               PersonnelClassIDType                 `json:",omitempty" xml:",omitempty"`
	PersonID                       PersonIDType                         `json:",omitempty" xml:",omitempty"`
	Description                    []DescriptionType                    `json:",omitempty" xml:",omitempty"`
	Quantity                       []QuantityValueType                  `json:",omitempty" xml:",omitempty"`
	PersonnelSpecificationProperty []PersonnelSpecificationPropertyType `json:",omitempty" xml:",omitempty"`
}

// PersonnelSpecificationPropertyType was auto-generated. Do not change.
type PersonnelSpecificationPropertyType struct {
	ID          IdentifierType      `json:",omitempty" xml:",omitempty"`
	Description []DescriptionType   `json:",omitempty" xml:",omitempty"`
	Value       []ValueType         `json:",omitempty" xml:",omitempty"`
	Quantity    []QuantityValueType `json:",omitempty" xml:",omitempty"`
}

// EquipmentSpecificationType was auto-generated. Do not change.
type EquipmentSpecificationType struct {
	EquipmentClassID               EquipmentClassIDType                 `json:",omitempty" xml:",omitempty"`
	EquipmentID                    EquipmentIDType                      `json:",omitempty" xml:",omitempty"`
	Description                    []DescriptionType                    `json:",omitempty" xml:",omitempty"`
	Quantity                       []QuantityValueType                  `json:",omitempty" xml:",omitempty"`
	EquipmentSpecificationProperty []EquipmentSpecificationPropertyType `json:",omitempty" xml:",omitempty"`
}

// EquipmentSpecificationPropertyType was auto-generated. Do not change.
type EquipmentSpecificationPropertyType struct {
	ID          IdentifierType      `json:",omitempty" xml:",omitempty"`
	Description []DescriptionType   `json:",omitempty" xml:",omitempty"`
	Value       []ValueType         `json:",omitempty" xml:",omitempty"`
	Quantity    []QuantityValueType `json:",omitempty" xml:",omitempty"`
}

// PhysicalAssetSpecificationType was auto-generated. Do not change.
type PhysicalAssetSpecificationType struct {
	PhysicalAssetClassID               PhysicalAssetClassIDType                 `json:",omitempty" xml:",omitempty"`
	PhysicalAssetID                    PhysicalAssetIDType                      `json:",omitempty" xml:",omitempty"`
	Description                        []DescriptionType                        `json:",omitempty" xml:",omitempty"`
	Quantity                           []QuantityValueType                      `json:",omitempty" xml:",omitempty"`
	PhysicalAssetSpecificationProperty []PhysicalAssetSpecificationPropertyType `json:",omitempty" xml:",omitempty"`
}

// PhysicalAssetSpecificationPropertyType was auto-generated. Do not change.
type PhysicalAssetSpecificationPropertyType struct {
	ID          IdentifierType      `json:",omitempty" xml:",omitempty"`
	Description []DescriptionType   `json:",omitempty" xml:",omitempty"`
	Value       []ValueType         `json:",omitempty" xml:",omitempty"`
	Quantity    []QuantityValueType `json:",omitempty" xml:",omitempty"`
}

// MaterialSpecificationType was auto-generated. Do not change.
type MaterialSpecificationType struct {
	MaterialClassID               MaterialClassIDType                 `json:",omitempty" xml:",omitempty"`
	MaterialDefinitionID          MaterialDefinitionIDType            `json:",omitempty" xml:",omitempty"`
	Description                   []DescriptionType                   `json:",omitempty" xml:",omitempty"`
	MaterialUse                   MaterialUseType                     `json:",omitempty" xml:",omitempty"`
	Quantity                      []QuantityValueType                 `json:",omitempty" xml:",omitempty"`
	AssemblySpecification         []MaterialSpecificationType         `json:",omitempty" xml:",omitempty"`
	AssemblyType                  AssemblyTypeType                    `json:",omitempty" xml:",omitempty"`
	AssemblyRelationship          AssemblyRelationshipType            `json:",omitempty" xml:",omitempty"`
	MaterialSpecificationProperty []MaterialSpecificationPropertyType `json:",omitempty" xml:",omitempty"`
}

// MaterialSpecificationPropertyType was auto-generated. Do not change.
type MaterialSpecificationPropertyType struct {
	ID          IdentifierType      `json:",omitempty" xml:",omitempty"`
	Description []DescriptionType   `json:",omitempty" xml:",omitempty"`
	Value       []ValueType         `json:",omitempty" xml:",omitempty"`
	Quantity    []QuantityValueType `json:",omitempty" xml:",omitempty"`
}

// GetProductInformationType was auto-generated. Do not change.
type GetProductInformationType struct {
	ApplicationArea TransApplicationAreaType          `json:",omitempty" xml:",omitempty"`
	DataArea        GetProductInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                            `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                            `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// GetProductInformationTypeDataArea was auto-generated. Do not change.
type GetProductInformationTypeDataArea struct {
	Get                []string                 `json:",omitempty" xml:",omitempty"`
	ProductInformation []ProductInformationType `json:",omitempty" xml:",omitempty"`
}

// ShowProductInformationType was auto-generated. Do not change.
type ShowProductInformationType struct {
	ApplicationArea TransApplicationAreaType           `json:",omitempty" xml:",omitempty"`
	DataArea        ShowProductInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                             `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                             `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ShowProductInformationTypeDataArea was auto-generated. Do not change.
type ShowProductInformationTypeDataArea struct {
	Show               TransShowType            `json:",omitempty" xml:",omitempty"`
	ProductInformation []ProductInformationType `json:",omitempty" xml:",omitempty"`
}

// ProcessProductInformationType was auto-generated. Do not change.
type ProcessProductInformationType struct {
	ApplicationArea TransApplicationAreaType              `json:",omitempty" xml:",omitempty"`
	DataArea        ProcessProductInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ProcessProductInformationTypeDataArea was auto-generated. Do not change.
type ProcessProductInformationTypeDataArea struct {
	Process            TransProcessType         `json:",omitempty" xml:",omitempty"`
	ProductInformation []ProductInformationType `json:",omitempty" xml:",omitempty"`
}

// AcknowledgeProductInformationType was auto-generated. Do not change.
type AcknowledgeProductInformationType struct {
	ApplicationArea TransApplicationAreaType                  `json:",omitempty" xml:",omitempty"`
	DataArea        AcknowledgeProductInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                    `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                    `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// AcknowledgeProductInformationTypeDataArea was auto-generated. Do not change.
type AcknowledgeProductInformationTypeDataArea struct {
	Acknowledge        TransAcknowledgeType     `json:",omitempty" xml:",omitempty"`
	ProductInformation []ProductInformationType `json:",omitempty" xml:",omitempty"`
}

// ChangeProductInformationType was auto-generated. Do not change.
type ChangeProductInformationType struct {
	ApplicationArea TransApplicationAreaType             `json:",omitempty" xml:",omitempty"`
	DataArea        ChangeProductInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                               `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                               `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ChangeProductInformationTypeDataArea was auto-generated. Do not change.
type ChangeProductInformationTypeDataArea struct {
	Change             TransChangeType          `json:",omitempty" xml:",omitempty"`
	ProductInformation []ProductInformationType `json:",omitempty" xml:",omitempty"`
}

// RespondProductInformationType was auto-generated. Do not change.
type RespondProductInformationType struct {
	ApplicationArea TransApplicationAreaType              `json:",omitempty" xml:",omitempty"`
	DataArea        RespondProductInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// RespondProductInformationTypeDataArea was auto-generated. Do not change.
type RespondProductInformationTypeDataArea struct {
	Respond            TransRespondType         `json:",omitempty" xml:",omitempty"`
	ProductInformation []ProductInformationType `json:",omitempty" xml:",omitempty"`
}

// CancelProductInformationType was auto-generated. Do not change.
type CancelProductInformationType struct {
	ApplicationArea TransApplicationAreaType             `json:",omitempty" xml:",omitempty"`
	DataArea        CancelProductInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                               `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                               `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// CancelProductInformationTypeDataArea was auto-generated. Do not change.
type CancelProductInformationTypeDataArea struct {
	Cancel             []TransActionCriteriaType `json:",omitempty" xml:",omitempty"`
	ProductInformation []ProductInformationType  `json:",omitempty" xml:",omitempty"`
}

// SyncProductInformationType was auto-generated. Do not change.
type SyncProductInformationType struct {
	ApplicationArea TransApplicationAreaType           `json:",omitempty" xml:",omitempty"`
	DataArea        SyncProductInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                             `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                             `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// SyncProductInformationTypeDataArea was auto-generated. Do not change.
type SyncProductInformationTypeDataArea struct {
	Sync               []TransActionCriteriaType `json:",omitempty" xml:",omitempty"`
	ProductInformation []ProductInformationType  `json:",omitempty" xml:",omitempty"`
}

// GetProductDefinitionType was auto-generated. Do not change.
type GetProductDefinitionType struct {
	ApplicationArea TransApplicationAreaType         `json:",omitempty" xml:",omitempty"`
	DataArea        GetProductDefinitionTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                           `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                           `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// GetProductDefinitionTypeDataArea was auto-generated. Do not change.
type GetProductDefinitionTypeDataArea struct {
	Get               []string                `json:",omitempty" xml:",omitempty"`
	ProductDefinition []ProductDefinitionType `json:",omitempty" xml:",omitempty"`
}

// ShowProductDefinitionType was auto-generated. Do not change.
type ShowProductDefinitionType struct {
	ApplicationArea TransApplicationAreaType          `json:",omitempty" xml:",omitempty"`
	DataArea        ShowProductDefinitionTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                            `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                            `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ShowProductDefinitionTypeDataArea was auto-generated. Do not change.
type ShowProductDefinitionTypeDataArea struct {
	Show              TransShowType           `json:",omitempty" xml:",omitempty"`
	ProductDefinition []ProductDefinitionType `json:",omitempty" xml:",omitempty"`
}

// ProcessProductDefinitionType was auto-generated. Do not change.
type ProcessProductDefinitionType struct {
	ApplicationArea TransApplicationAreaType             `json:",omitempty" xml:",omitempty"`
	DataArea        ProcessProductDefinitionTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                               `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                               `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ProcessProductDefinitionTypeDataArea was auto-generated. Do not change.
type ProcessProductDefinitionTypeDataArea struct {
	Process           TransProcessType        `json:",omitempty" xml:",omitempty"`
	ProductDefinition []ProductDefinitionType `json:",omitempty" xml:",omitempty"`
}

// AcknowledgeProductDefinitionType was auto-generated. Do not change.
type AcknowledgeProductDefinitionType struct {
	ApplicationArea TransApplicationAreaType                 `json:",omitempty" xml:",omitempty"`
	DataArea        AcknowledgeProductDefinitionTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                   `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                   `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// AcknowledgeProductDefinitionTypeDataArea was auto-generated. Do not change.
type AcknowledgeProductDefinitionTypeDataArea struct {
	Acknowledge       TransAcknowledgeType    `json:",omitempty" xml:",omitempty"`
	ProductDefinition []ProductDefinitionType `json:",omitempty" xml:",omitempty"`
}

// ChangeProductDefinitionType was auto-generated. Do not change.
type ChangeProductDefinitionType struct {
	ApplicationArea TransApplicationAreaType            `json:",omitempty" xml:",omitempty"`
	DataArea        ChangeProductDefinitionTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                              `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                              `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ChangeProductDefinitionTypeDataArea was auto-generated. Do not change.
type ChangeProductDefinitionTypeDataArea struct {
	Change            TransChangeType         `json:",omitempty" xml:",omitempty"`
	ProductDefinition []ProductDefinitionType `json:",omitempty" xml:",omitempty"`
}

// RespondProductDefinitionType was auto-generated. Do not change.
type RespondProductDefinitionType struct {
	ApplicationArea TransApplicationAreaType             `json:",omitempty" xml:",omitempty"`
	DataArea        RespondProductDefinitionTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                               `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                               `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// RespondProductDefinitionTypeDataArea was auto-generated. Do not change.
type RespondProductDefinitionTypeDataArea struct {
	Respond           TransRespondType        `json:",omitempty" xml:",omitempty"`
	ProductDefinition []ProductDefinitionType `json:",omitempty" xml:",omitempty"`
}

// CancelProductDefinitionType was auto-generated. Do not change.
type CancelProductDefinitionType struct {
	ApplicationArea TransApplicationAreaType            `json:",omitempty" xml:",omitempty"`
	DataArea        CancelProductDefinitionTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                              `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                              `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// CancelProductDefinitionTypeDataArea was auto-generated. Do not change.
type CancelProductDefinitionTypeDataArea struct {
	Cancel            []TransActionCriteriaType `json:",omitempty" xml:",omitempty"`
	ProductDefinition []ProductDefinitionType   `json:",omitempty" xml:",omitempty"`
}

// SyncProductDefinitionType was auto-generated. Do not change.
type SyncProductDefinitionType struct {
	ApplicationArea TransApplicationAreaType          `json:",omitempty" xml:",omitempty"`
	DataArea        SyncProductDefinitionTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                            `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                            `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// SyncProductDefinitionTypeDataArea was auto-generated. Do not change.
type SyncProductDefinitionTypeDataArea struct {
	Sync              []TransActionCriteriaType `json:",omitempty" xml:",omitempty"`
	ProductDefinition []ProductDefinitionType   `json:",omitempty" xml:",omitempty"`
}

// ProductionCapabilityType was auto-generated. Do not change.
type ProductionCapabilityType struct {
	ID                       IdentifierType                 `json:",omitempty" xml:",omitempty"`
	Description              []DescriptionType              `json:",omitempty" xml:",omitempty"`
	Location                 LocationType                   `json:",omitempty" xml:",omitempty"`
	HierarchyScope           HierarchyScopeType             `json:",omitempty" xml:",omitempty"`
	PublishedDate            PublishedDateType              `json:",omitempty" xml:",omitempty"`
	CapabilityType           CapabilityTypeType             `json:",omitempty" xml:",omitempty"`
	Reason                   ReasonType                     `json:",omitempty" xml:",omitempty"`
	EquipmentElementLevel    []EquipmentElementLevelType    `json:",omitempty" xml:",omitempty"`
	StartTime                StartTimeType                  `json:",omitempty" xml:",omitempty"`
	EndTime                  EndTimeType                    `json:",omitempty" xml:",omitempty"`
	PersonnelCapability      []PersonnelCapabilityType      `json:",omitempty" xml:",omitempty"`
	EquipmentCapability      []EquipmentCapabilityType      `json:",omitempty" xml:",omitempty"`
	PhysicalAssetCapability  []PhysicalAssetCapabilityType  `json:",omitempty" xml:",omitempty"`
	MaterialCapability       []MaterialCapabilityType       `json:",omitempty" xml:",omitempty"`
	ProcessSegmentCapability []ProcessSegmentCapabilityType `json:",omitempty" xml:",omitempty"`
}

// PersonnelCapabilityType was auto-generated. Do not change.
type PersonnelCapabilityType struct {
	PersonnelClassID            PersonnelClassIDType              `json:",omitempty" xml:",omitempty"`
	PersonID                    PersonIDType                      `json:",omitempty" xml:",omitempty"`
	Description                 []DescriptionType                 `json:",omitempty" xml:",omitempty"`
	CapabilityType              CapabilityTypeType                `json:",omitempty" xml:",omitempty"`
	Reason                      ReasonType                        `json:",omitempty" xml:",omitempty"`
	EquipmentElementLevel       []EquipmentElementLevelType       `json:",omitempty" xml:",omitempty"`
	StartTime                   StartTimeType                     `json:",omitempty" xml:",omitempty"`
	EndTime                     EndTimeType                       `json:",omitempty" xml:",omitempty"`
	Location                    LocationType                      `json:",omitempty" xml:",omitempty"`
	HierarchyScope              HierarchyScopeType                `json:",omitempty" xml:",omitempty"`
	Quantity                    []QuantityValueType               `json:",omitempty" xml:",omitempty"`
	PersonnelCapabilityProperty []PersonnelCapabilityPropertyType `json:",omitempty" xml:",omitempty"`
}

// PersonnelCapabilityPropertyType was auto-generated. Do not change.
type PersonnelCapabilityPropertyType struct {
	ID          IdentifierType      `json:",omitempty" xml:",omitempty"`
	Description []DescriptionType   `json:",omitempty" xml:",omitempty"`
	Value       []ValueType         `json:",omitempty" xml:",omitempty"`
	Quantity    []QuantityValueType `json:",omitempty" xml:",omitempty"`
}

// EquipmentCapabilityType was auto-generated. Do not change.
type EquipmentCapabilityType struct {
	EquipmentClassID            EquipmentClassIDType              `json:",omitempty" xml:",omitempty"`
	EquipmentID                 EquipmentIDType                   `json:",omitempty" xml:",omitempty"`
	Description                 []DescriptionType                 `json:",omitempty" xml:",omitempty"`
	CapabilityType              CapabilityTypeType                `json:",omitempty" xml:",omitempty"`
	Reason                      ReasonType                        `json:",omitempty" xml:",omitempty"`
	EquipmentElementLevel       []EquipmentElementLevelType       `json:",omitempty" xml:",omitempty"`
	StartTime                   StartTimeType                     `json:",omitempty" xml:",omitempty"`
	EndTime                     EndTimeType                       `json:",omitempty" xml:",omitempty"`
	Location                    LocationType                      `json:",omitempty" xml:",omitempty"`
	HierarchyScope              HierarchyScopeType                `json:",omitempty" xml:",omitempty"`
	Quantity                    []QuantityValueType               `json:",omitempty" xml:",omitempty"`
	EquipmentCapabilityProperty []EquipmentCapabilityPropertyType `json:",omitempty" xml:",omitempty"`
}

// EquipmentCapabilityPropertyType was auto-generated. Do not change.
type EquipmentCapabilityPropertyType struct {
	ID          IdentifierType      `json:",omitempty" xml:",omitempty"`
	Description []DescriptionType   `json:",omitempty" xml:",omitempty"`
	Value       []ValueType         `json:",omitempty" xml:",omitempty"`
	Quantity    []QuantityValueType `json:",omitempty" xml:",omitempty"`
}

// PhysicalAssetCapabilityType was auto-generated. Do not change.
type PhysicalAssetCapabilityType struct {
	PhysicalAssetClassID            PhysicalAssetClassIDType              `json:",omitempty" xml:",omitempty"`
	PhysicalAssetID                 PhysicalAssetIDType                   `json:",omitempty" xml:",omitempty"`
	Description                     []DescriptionType                     `json:",omitempty" xml:",omitempty"`
	CapabilityType                  CapabilityTypeType                    `json:",omitempty" xml:",omitempty"`
	Reason                          ReasonType                            `json:",omitempty" xml:",omitempty"`
	EquipmentElementLevel           []EquipmentElementLevelType           `json:",omitempty" xml:",omitempty"`
	StartTime                       StartTimeType                         `json:",omitempty" xml:",omitempty"`
	EndTime                         EndTimeType                           `json:",omitempty" xml:",omitempty"`
	Location                        LocationType                          `json:",omitempty" xml:",omitempty"`
	HierarchyScope                  HierarchyScopeType                    `json:",omitempty" xml:",omitempty"`
	Quantity                        []QuantityValueType                   `json:",omitempty" xml:",omitempty"`
	PhysicalAssetCapabilityProperty []PhysicalAssetCapabilityPropertyType `json:",omitempty" xml:",omitempty"`
}

// PhysicalAssetCapabilityPropertyType was auto-generated. Do not change.
type PhysicalAssetCapabilityPropertyType struct {
	ID          IdentifierType      `json:",omitempty" xml:",omitempty"`
	Description []DescriptionType   `json:",omitempty" xml:",omitempty"`
	Value       []ValueType         `json:",omitempty" xml:",omitempty"`
	Quantity    []QuantityValueType `json:",omitempty" xml:",omitempty"`
}

// MaterialCapabilityType was auto-generated. Do not change.
type MaterialCapabilityType struct {
	MaterialClassID            MaterialClassIDType              `json:",omitempty" xml:",omitempty"`
	MaterialDefinitionID       MaterialDefinitionIDType         `json:",omitempty" xml:",omitempty"`
	MaterialLotID              MaterialLotIDType                `json:",omitempty" xml:",omitempty"`
	MaterialSubLotID           MaterialSubLotIDType             `json:",omitempty" xml:",omitempty"`
	Description                []DescriptionType                `json:",omitempty" xml:",omitempty"`
	CapabilityType             CapabilityTypeType               `json:",omitempty" xml:",omitempty"`
	Reason                     ReasonType                       `json:",omitempty" xml:",omitempty"`
	EquipmentElementLevel      []EquipmentElementLevelType      `json:",omitempty" xml:",omitempty"`
	MaterialUse                MaterialUseType                  `json:",omitempty" xml:",omitempty"`
	StartTime                  StartTimeType                    `json:",omitempty" xml:",omitempty"`
	EndTime                    EndTimeType                      `json:",omitempty" xml:",omitempty"`
	Location                   LocationType                     `json:",omitempty" xml:",omitempty"`
	HierarchyScope             HierarchyScopeType               `json:",omitempty" xml:",omitempty"`
	Quantity                   []QuantityValueType              `json:",omitempty" xml:",omitempty"`
	AssemblyCapability         []MaterialCapabilityType         `json:",omitempty" xml:",omitempty"`
	AssemblyType               AssemblyTypeType                 `json:",omitempty" xml:",omitempty"`
	AssemblyRelationship       AssemblyRelationshipType         `json:",omitempty" xml:",omitempty"`
	MaterialCapabilityProperty []MaterialCapabilityPropertyType `json:",omitempty" xml:",omitempty"`
}

// MaterialCapabilityPropertyType was auto-generated. Do not change.
type MaterialCapabilityPropertyType struct {
	ID          IdentifierType      `json:",omitempty" xml:",omitempty"`
	Description []DescriptionType   `json:",omitempty" xml:",omitempty"`
	Value       []ValueType         `json:",omitempty" xml:",omitempty"`
	Quantity    []QuantityValueType `json:",omitempty" xml:",omitempty"`
	MaterialUse MaterialUseType     `json:",omitempty" xml:",omitempty"`
}

// ProcessSegmentCapabilityType was auto-generated. Do not change.
type ProcessSegmentCapabilityType struct {
	ID                       IdentifierType                 `json:",omitempty" xml:",omitempty"`
	Description              []DescriptionType              `json:",omitempty" xml:",omitempty"`
	ProcessSegmentID         ProcessSegmentIDType           `json:",omitempty" xml:",omitempty"`
	CapabilityType           CapabilityTypeType             `json:",omitempty" xml:",omitempty"`
	Reason                   ReasonType                     `json:",omitempty" xml:",omitempty"`
	Location                 []LocationType                 `json:",omitempty" xml:",omitempty"`
	HierarchyScope           HierarchyScopeType             `json:",omitempty" xml:",omitempty"`
	EquipmentElementLevel    []EquipmentElementLevelType    `json:",omitempty" xml:",omitempty"`
	StartTime                StartTimeType                  `json:",omitempty" xml:",omitempty"`
	EndTime                  EndTimeType                    `json:",omitempty" xml:",omitempty"`
	PersonnelCapability      []PersonnelCapabilityType      `json:",omitempty" xml:",omitempty"`
	EquipmentCapability      []EquipmentCapabilityType      `json:",omitempty" xml:",omitempty"`
	MaterialCapability       []MaterialCapabilityType       `json:",omitempty" xml:",omitempty"`
	ProcessSegmentCapability []ProcessSegmentCapabilityType `json:",omitempty" xml:",omitempty"`
}

// GetProductionCapabilityType was auto-generated. Do not change.
type GetProductionCapabilityType struct {
	ApplicationArea TransApplicationAreaType            `json:",omitempty" xml:",omitempty"`
	DataArea        GetProductionCapabilityTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                              `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                              `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// GetProductionCapabilityTypeDataArea was auto-generated. Do not change.
type GetProductionCapabilityTypeDataArea struct {
	Get                  []string                   `json:",omitempty" xml:",omitempty"`
	ProductionCapability []ProductionCapabilityType `json:",omitempty" xml:",omitempty"`
}

// ShowProductionCapabilityType was auto-generated. Do not change.
type ShowProductionCapabilityType struct {
	ApplicationArea TransApplicationAreaType             `json:",omitempty" xml:",omitempty"`
	DataArea        ShowProductionCapabilityTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                               `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                               `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ShowProductionCapabilityTypeDataArea was auto-generated. Do not change.
type ShowProductionCapabilityTypeDataArea struct {
	Show                 TransShowType              `json:",omitempty" xml:",omitempty"`
	ProductionCapability []ProductionCapabilityType `json:",omitempty" xml:",omitempty"`
}

// ProcessProductionCapabilityType was auto-generated. Do not change.
type ProcessProductionCapabilityType struct {
	ApplicationArea TransApplicationAreaType                `json:",omitempty" xml:",omitempty"`
	DataArea        ProcessProductionCapabilityTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                  `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                  `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ProcessProductionCapabilityTypeDataArea was auto-generated. Do not change.
type ProcessProductionCapabilityTypeDataArea struct {
	Process              TransProcessType           `json:",omitempty" xml:",omitempty"`
	ProductionCapability []ProductionCapabilityType `json:",omitempty" xml:",omitempty"`
}

// AcknowledgeProductionCapabilityType was auto-generated. Do not change.
type AcknowledgeProductionCapabilityType struct {
	ApplicationArea TransApplicationAreaType                    `json:",omitempty" xml:",omitempty"`
	DataArea        AcknowledgeProductionCapabilityTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                      `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                      `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// AcknowledgeProductionCapabilityTypeDataArea was auto-generated. Do not change.
type AcknowledgeProductionCapabilityTypeDataArea struct {
	Acknowledge          TransAcknowledgeType       `json:",omitempty" xml:",omitempty"`
	ProductionCapability []ProductionCapabilityType `json:",omitempty" xml:",omitempty"`
}

// ChangeProductionCapabilityType was auto-generated. Do not change.
type ChangeProductionCapabilityType struct {
	ApplicationArea TransApplicationAreaType               `json:",omitempty" xml:",omitempty"`
	DataArea        ChangeProductionCapabilityTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                 `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                 `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ChangeProductionCapabilityTypeDataArea was auto-generated. Do not change.
type ChangeProductionCapabilityTypeDataArea struct {
	Change               TransChangeType            `json:",omitempty" xml:",omitempty"`
	ProductionCapability []ProductionCapabilityType `json:",omitempty" xml:",omitempty"`
}

// RespondProductionCapabilityType was auto-generated. Do not change.
type RespondProductionCapabilityType struct {
	ApplicationArea TransApplicationAreaType                `json:",omitempty" xml:",omitempty"`
	DataArea        RespondProductionCapabilityTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                  `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                  `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// RespondProductionCapabilityTypeDataArea was auto-generated. Do not change.
type RespondProductionCapabilityTypeDataArea struct {
	Respond              TransRespondType           `json:",omitempty" xml:",omitempty"`
	ProductionCapability []ProductionCapabilityType `json:",omitempty" xml:",omitempty"`
}

// CancelProductionCapabilityType was auto-generated. Do not change.
type CancelProductionCapabilityType struct {
	ApplicationArea TransApplicationAreaType               `json:",omitempty" xml:",omitempty"`
	DataArea        CancelProductionCapabilityTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                 `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                 `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// CancelProductionCapabilityTypeDataArea was auto-generated. Do not change.
type CancelProductionCapabilityTypeDataArea struct {
	Cancel               []TransActionCriteriaType  `json:",omitempty" xml:",omitempty"`
	ProductionCapability []ProductionCapabilityType `json:",omitempty" xml:",omitempty"`
}

// SyncProductionCapabilityType was auto-generated. Do not change.
type SyncProductionCapabilityType struct {
	ApplicationArea TransApplicationAreaType             `json:",omitempty" xml:",omitempty"`
	DataArea        SyncProductionCapabilityTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                               `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                               `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// SyncProductionCapabilityTypeDataArea was auto-generated. Do not change.
type SyncProductionCapabilityTypeDataArea struct {
	Sync                 []TransActionCriteriaType  `json:",omitempty" xml:",omitempty"`
	ProductionCapability []ProductionCapabilityType `json:",omitempty" xml:",omitempty"`
}

// ProductionPerformanceType was auto-generated. Do not change.
type ProductionPerformanceType struct {
	ID                    IdentifierType            `json:",omitempty" xml:",omitempty"`
	Description           []DescriptionType         `json:",omitempty" xml:",omitempty"`
	Location              LocationType              `json:",omitempty" xml:",omitempty"`
	HierarchyScope        HierarchyScopeType        `json:",omitempty" xml:",omitempty"`
	PublishedDate         PublishedDateType         `json:",omitempty" xml:",omitempty"`
	ProductionScheduleID  ProductionScheduleIDType  `json:",omitempty" xml:",omitempty"`
	StartTime             StartTimeType             `json:",omitempty" xml:",omitempty"`
	EndTime               EndTimeType               `json:",omitempty" xml:",omitempty"`
	EquipmentElementLevel EquipmentElementLevelType `json:",omitempty" xml:",omitempty"`
	ProductionResponse    []ProductionResponseType  `json:",omitempty" xml:",omitempty"`
	PerformanceState      ResponseStateType         `json:",omitempty" xml:",omitempty"`
}

// ProductionResponseType was auto-generated. Do not change.
type ProductionResponseType struct {
	ID                      IdentifierType                `json:",omitempty" xml:",omitempty"`
	ProductionRequestID     ProductionRequestIDType       `json:",omitempty" xml:",omitempty"`
	ProductProductionRuleID []ProductProductionRuleIDType `json:",omitempty" xml:",omitempty"`
	Version                 []VersionType                 `json:",omitempty" xml:",omitempty"`
	Location                LocationType                  `json:",omitempty" xml:",omitempty"`
	HierarchyScope          HierarchyScopeType            `json:",omitempty" xml:",omitempty"`
	StartTime               StartTimeType                 `json:",omitempty" xml:",omitempty"`
	EndTime                 EndTimeType                   `json:",omitempty" xml:",omitempty"`
	SegmentResponse         []SegmentResponseType         `json:",omitempty" xml:",omitempty"`
	ResponseState           ResponseStateType             `json:",omitempty" xml:",omitempty"`
}

// SegmentResponseType was auto-generated. Do not change.
type SegmentResponseType struct {
	ID                                 IdentifierType                         `json:",omitempty" xml:",omitempty"`
	ProcessSegmentID                   []ProcessSegmentIDType                 `json:",omitempty" xml:",omitempty"`
	ProductSegmentID                   []ProductSegmentIDType                 `json:",omitempty" xml:",omitempty"`
	Description                        []DescriptionType                      `json:",omitempty" xml:",omitempty"`
	Location                           LocationType                           `json:",omitempty" xml:",omitempty"`
	HierarchyScope                     HierarchyScopeType                     `json:",omitempty" xml:",omitempty"`
	ActualStartTime                    ActualStartTimeType                    `json:",omitempty" xml:",omitempty"`
	ActualEndTime                      ActualEndTimeType                      `json:",omitempty" xml:",omitempty"`
	ProductionData                     []ProductionDataType                   `json:",omitempty" xml:",omitempty"`
	PersonnelActual                    []PersonnelActualType                  `json:",omitempty" xml:",omitempty"`
	EquipmentActual                    []EquipmentActualType                  `json:",omitempty" xml:",omitempty"`
	PhysicalAssetActual                []PhysicalAssetActualType              `json:",omitempty" xml:",omitempty"`
	MaterialActual                     []MaterialActualType                   `json:",omitempty" xml:",omitempty"`
	MaterialProducedActual             []MaterialProducedActualType           `json:",omitempty" xml:",omitempty"`
	MaterialConsumedActual             []MaterialConsumedActualType           `json:",omitempty" xml:",omitempty"`
	ConsumableActual                   []ConsumableActualType                 `json:",omitempty" xml:",omitempty"`
	SegmentResponse                    []SegmentResponseType                  `json:",omitempty" xml:",omitempty"`
	RequiredByRequestedSegmentResponse RequiredByRequestedSegmentResponseType `json:",omitempty" xml:",omitempty"`
	SegmentState                       ResponseStateType                      `json:",omitempty" xml:",omitempty"`
}

// ProductionDataType was auto-generated. Do not change.
type ProductionDataType struct {
	ID                                 IdentifierType                         `json:",omitempty" xml:",omitempty"`
	Description                        []DescriptionType                      `json:",omitempty" xml:",omitempty"`
	Value                              []ValueType                            `json:",omitempty" xml:",omitempty"`
	RequiredByRequestedSegmentResponse RequiredByRequestedSegmentResponseType `json:",omitempty" xml:",omitempty"`
	ProductionData                     []ProductionDataType                   `json:",omitempty" xml:",omitempty"`
}

// PersonnelActualType was auto-generated. Do not change.
type PersonnelActualType struct {
	PersonnelClassID                   []PersonnelClassIDType                 `json:",omitempty" xml:",omitempty"`
	PersonID                           []PersonIDType                         `json:",omitempty" xml:",omitempty"`
	Description                        []DescriptionType                      `json:",omitempty" xml:",omitempty"`
	Location                           LocationType                           `json:",omitempty" xml:",omitempty"`
	HierarchyScope                     HierarchyScopeType                     `json:",omitempty" xml:",omitempty"`
	Quantity                           []QuantityValueType                    `json:",omitempty" xml:",omitempty"`
	PersonnelActualProperty            []PersonnelActualPropertyType          `json:",omitempty" xml:",omitempty"`
	RequiredByRequestedSegmentResponse RequiredByRequestedSegmentResponseType `json:",omitempty" xml:",omitempty"`
}

// PersonnelActualPropertyType was auto-generated. Do not change.
type PersonnelActualPropertyType struct {
	ID                                 IdentifierType                         `json:",omitempty" xml:",omitempty"`
	Description                        []DescriptionType                      `json:",omitempty" xml:",omitempty"`
	Value                              []ValueType                            `json:",omitempty" xml:",omitempty"`
	Quantity                           []QuantityValueType                    `json:",omitempty" xml:",omitempty"`
	RequiredByRequestedSegmentResponse RequiredByRequestedSegmentResponseType `json:",omitempty" xml:",omitempty"`
}

// EquipmentActualType was auto-generated. Do not change.
type EquipmentActualType struct {
	EquipmentClassID                   []EquipmentClassIDType                 `json:",omitempty" xml:",omitempty"`
	EquipmentID                        []EquipmentIDType                      `json:",omitempty" xml:",omitempty"`
	Description                        []DescriptionType                      `json:",omitempty" xml:",omitempty"`
	Location                           LocationType                           `json:",omitempty" xml:",omitempty"`
	HierarchyScope                     HierarchyScopeType                     `json:",omitempty" xml:",omitempty"`
	Quantity                           []QuantityValueType                    `json:",omitempty" xml:",omitempty"`
	EquipmentActualProperty            []EquipmentActualPropertyType          `json:",omitempty" xml:",omitempty"`
	RequiredByRequestedSegmentResponse RequiredByRequestedSegmentResponseType `json:",omitempty" xml:",omitempty"`
}

// EquipmentActualPropertyType was auto-generated. Do not change.
type EquipmentActualPropertyType struct {
	ID                                 IdentifierType                         `json:",omitempty" xml:",omitempty"`
	Description                        []DescriptionType                      `json:",omitempty" xml:",omitempty"`
	Value                              []ValueType                            `json:",omitempty" xml:",omitempty"`
	Quantity                           []QuantityValueType                    `json:",omitempty" xml:",omitempty"`
	RequiredByRequestedSegmentResponse RequiredByRequestedSegmentResponseType `json:",omitempty" xml:",omitempty"`
}

// PhysicalAssetActualType was auto-generated. Do not change.
type PhysicalAssetActualType struct {
	PhysicalAssetClassID               []PhysicalAssetClassIDType             `json:",omitempty" xml:",omitempty"`
	PhysicalAssetID                    []PhysicalAssetIDType                  `json:",omitempty" xml:",omitempty"`
	Description                        []DescriptionType                      `json:",omitempty" xml:",omitempty"`
	Location                           LocationType                           `json:",omitempty" xml:",omitempty"`
	HierarchyScope                     HierarchyScopeType                     `json:",omitempty" xml:",omitempty"`
	Quantity                           []QuantityValueType                    `json:",omitempty" xml:",omitempty"`
	PhysicalAssetActualProperty        []PhysicalAssetActualPropertyType      `json:",omitempty" xml:",omitempty"`
	RequiredByRequestedSegmentResponse RequiredByRequestedSegmentResponseType `json:",omitempty" xml:",omitempty"`
}

// PhysicalAssetActualPropertyType was auto-generated. Do not change.
type PhysicalAssetActualPropertyType struct {
	ID                                 IdentifierType                         `json:",omitempty" xml:",omitempty"`
	Description                        []DescriptionType                      `json:",omitempty" xml:",omitempty"`
	Value                              []ValueType                            `json:",omitempty" xml:",omitempty"`
	Quantity                           []QuantityValueType                    `json:",omitempty" xml:",omitempty"`
	RequiredByRequestedSegmentResponse RequiredByRequestedSegmentResponseType `json:",omitempty" xml:",omitempty"`
}

// MaterialActualType was auto-generated. Do not change.
type MaterialActualType struct {
	MaterialClassID                    []MaterialClassIDType                  `json:",omitempty" xml:",omitempty"`
	MaterialDefinitionID               []MaterialDefinitionIDType             `json:",omitempty" xml:",omitempty"`
	MaterialLotID                      []MaterialLotIDType                    `json:",omitempty" xml:",omitempty"`
	MaterialSubLotID                   []MaterialSubLotIDType                 `json:",omitempty" xml:",omitempty"`
	Description                        []DescriptionType                      `json:",omitempty" xml:",omitempty"`
	Location                           LocationType                           `json:",omitempty" xml:",omitempty"`
	HierarchyScope                     HierarchyScopeType                     `json:",omitempty" xml:",omitempty"`
	MaterialUse                        MaterialUseType                        `json:",omitempty" xml:",omitempty"`
	Quantity                           []QuantityValueType                    `json:",omitempty" xml:",omitempty"`
	AssemblyActual                     []MaterialActualType                   `json:",omitempty" xml:",omitempty"`
	AssemblyType                       AssemblyTypeType                       `json:",omitempty" xml:",omitempty"`
	AssemblyRelationship               AssemblyRelationshipType               `json:",omitempty" xml:",omitempty"`
	MaterialActualProperty             []MaterialActualPropertyType           `json:",omitempty" xml:",omitempty"`
	RequiredByRequestedSegmentResponse RequiredByRequestedSegmentResponseType `json:",omitempty" xml:",omitempty"`
}

// MaterialActualPropertyType was auto-generated. Do not change.
type MaterialActualPropertyType struct {
	ID                                 IdentifierType                         `json:",omitempty" xml:",omitempty"`
	Description                        []DescriptionType                      `json:",omitempty" xml:",omitempty"`
	Value                              []ValueType                            `json:",omitempty" xml:",omitempty"`
	Quantity                           []QuantityValueType                    `json:",omitempty" xml:",omitempty"`
	RequiredByRequestedSegmentResponse RequiredByRequestedSegmentResponseType `json:",omitempty" xml:",omitempty"`
}

// MaterialProducedActualType was auto-generated. Do not change.
type MaterialProducedActualType struct {
	MaterialClassID                    []MaterialClassIDType                  `json:",omitempty" xml:",omitempty"`
	MaterialDefinitionID               []MaterialDefinitionIDType             `json:",omitempty" xml:",omitempty"`
	MaterialLotID                      []MaterialLotIDType                    `json:",omitempty" xml:",omitempty"`
	MaterialSubLotID                   []MaterialSubLotIDType                 `json:",omitempty" xml:",omitempty"`
	Description                        []DescriptionType                      `json:",omitempty" xml:",omitempty"`
	Location                           LocationType                           `json:",omitempty" xml:",omitempty"`
	HierarchyScope                     HierarchyScopeType                     `json:",omitempty" xml:",omitempty"`
	Quantity                           []QuantityValueType                    `json:",omitempty" xml:",omitempty"`
	MaterialProducedActualProperty     []MaterialProducedActualPropertyType   `json:",omitempty" xml:",omitempty"`
	RequiredByRequestedSegmentResponse RequiredByRequestedSegmentResponseType `json:",omitempty" xml:",omitempty"`
}

// MaterialProducedActualPropertyType was auto-generated. Do not change.
type MaterialProducedActualPropertyType struct {
	ID                                 IdentifierType                         `json:",omitempty" xml:",omitempty"`
	Description                        []DescriptionType                      `json:",omitempty" xml:",omitempty"`
	Value                              []ValueType                            `json:",omitempty" xml:",omitempty"`
	Quantity                           []QuantityValueType                    `json:",omitempty" xml:",omitempty"`
	RequiredByRequestedSegmentResponse RequiredByRequestedSegmentResponseType `json:",omitempty" xml:",omitempty"`
}

// MaterialConsumedActualType was auto-generated. Do not change.
type MaterialConsumedActualType struct {
	MaterialClassID                    []MaterialClassIDType                  `json:",omitempty" xml:",omitempty"`
	MaterialDefinitionID               []MaterialDefinitionIDType             `json:",omitempty" xml:",omitempty"`
	MaterialLotID                      []MaterialLotIDType                    `json:",omitempty" xml:",omitempty"`
	MaterialSubLotID                   []MaterialSubLotIDType                 `json:",omitempty" xml:",omitempty"`
	Description                        []DescriptionType                      `json:",omitempty" xml:",omitempty"`
	Location                           LocationType                           `json:",omitempty" xml:",omitempty"`
	HierarchyScope                     HierarchyScopeType                     `json:",omitempty" xml:",omitempty"`
	Quantity                           []QuantityValueType                    `json:",omitempty" xml:",omitempty"`
	MaterialConsumedActualProperty     []MaterialConsumedActualPropertyType   `json:",omitempty" xml:",omitempty"`
	RequiredByRequestedSegmentResponse RequiredByRequestedSegmentResponseType `json:",omitempty" xml:",omitempty"`
}

// MaterialConsumedActualPropertyType was auto-generated. Do not change.
type MaterialConsumedActualPropertyType struct {
	ID                                 IdentifierType                         `json:",omitempty" xml:",omitempty"`
	Description                        []DescriptionType                      `json:",omitempty" xml:",omitempty"`
	Value                              []ValueType                            `json:",omitempty" xml:",omitempty"`
	Quantity                           []QuantityValueType                    `json:",omitempty" xml:",omitempty"`
	RequiredByRequestedSegmentResponse RequiredByRequestedSegmentResponseType `json:",omitempty" xml:",omitempty"`
}

// ConsumableActualType was auto-generated. Do not change.
type ConsumableActualType struct {
	MaterialClassID                    []MaterialClassIDType                  `json:",omitempty" xml:",omitempty"`
	MaterialDefinitionID               []MaterialDefinitionIDType             `json:",omitempty" xml:",omitempty"`
	Description                        []DescriptionType                      `json:",omitempty" xml:",omitempty"`
	Location                           LocationType                           `json:",omitempty" xml:",omitempty"`
	HierarchyScope                     HierarchyScopeType                     `json:",omitempty" xml:",omitempty"`
	Quantity                           []QuantityValueType                    `json:",omitempty" xml:",omitempty"`
	ConsumableActualProperty           []ConsumableActualPropertyType         `json:",omitempty" xml:",omitempty"`
	RequiredByRequestedSegmentResponse RequiredByRequestedSegmentResponseType `json:",omitempty" xml:",omitempty"`
}

// ConsumableActualPropertyType was auto-generated. Do not change.
type ConsumableActualPropertyType struct {
	ID                                 IdentifierType                         `json:",omitempty" xml:",omitempty"`
	Description                        []DescriptionType                      `json:",omitempty" xml:",omitempty"`
	Value                              []ValueType                            `json:",omitempty" xml:",omitempty"`
	Quantity                           []QuantityValueType                    `json:",omitempty" xml:",omitempty"`
	RequiredByRequestedSegmentResponse RequiredByRequestedSegmentResponseType `json:",omitempty" xml:",omitempty"`
}

// GetProductionPerformanceType was auto-generated. Do not change.
type GetProductionPerformanceType struct {
	ApplicationArea TransApplicationAreaType             `json:",omitempty" xml:",omitempty"`
	DataArea        GetProductionPerformanceTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                               `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                               `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// GetProductionPerformanceTypeDataArea was auto-generated. Do not change.
type GetProductionPerformanceTypeDataArea struct {
	Get                   []string                    `json:",omitempty" xml:",omitempty"`
	ProductionPerformance []ProductionPerformanceType `json:",omitempty" xml:",omitempty"`
}

// ShowProductionPerformanceType was auto-generated. Do not change.
type ShowProductionPerformanceType struct {
	ApplicationArea TransApplicationAreaType              `json:",omitempty" xml:",omitempty"`
	DataArea        ShowProductionPerformanceTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ShowProductionPerformanceTypeDataArea was auto-generated. Do not change.
type ShowProductionPerformanceTypeDataArea struct {
	Show                  TransShowType               `json:",omitempty" xml:",omitempty"`
	ProductionPerformance []ProductionPerformanceType `json:",omitempty" xml:",omitempty"`
}

// ProcessProductionPerformanceType was auto-generated. Do not change.
type ProcessProductionPerformanceType struct {
	ApplicationArea TransApplicationAreaType                 `json:",omitempty" xml:",omitempty"`
	DataArea        ProcessProductionPerformanceTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                   `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                   `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ProcessProductionPerformanceTypeDataArea was auto-generated. Do not change.
type ProcessProductionPerformanceTypeDataArea struct {
	Process               TransProcessType            `json:",omitempty" xml:",omitempty"`
	ProductionPerformance []ProductionPerformanceType `json:",omitempty" xml:",omitempty"`
}

// AcknowledgeProductionPerformanceType was auto-generated. Do not change.
type AcknowledgeProductionPerformanceType struct {
	ApplicationArea TransApplicationAreaType                     `json:",omitempty" xml:",omitempty"`
	DataArea        AcknowledgeProductionPerformanceTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                       `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                       `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// AcknowledgeProductionPerformanceTypeDataArea was auto-generated. Do not change.
type AcknowledgeProductionPerformanceTypeDataArea struct {
	Acknowledge           TransAcknowledgeType        `json:",omitempty" xml:",omitempty"`
	ProductionPerformance []ProductionPerformanceType `json:",omitempty" xml:",omitempty"`
}

// ChangeProductionPerformanceType was auto-generated. Do not change.
type ChangeProductionPerformanceType struct {
	ApplicationArea TransApplicationAreaType                `json:",omitempty" xml:",omitempty"`
	DataArea        ChangeProductionPerformanceTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                  `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                  `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ChangeProductionPerformanceTypeDataArea was auto-generated. Do not change.
type ChangeProductionPerformanceTypeDataArea struct {
	Change                TransChangeType             `json:",omitempty" xml:",omitempty"`
	ProductionPerformance []ProductionPerformanceType `json:",omitempty" xml:",omitempty"`
}

// RespondProductionPerformanceType was auto-generated. Do not change.
type RespondProductionPerformanceType struct {
	ApplicationArea TransApplicationAreaType                 `json:",omitempty" xml:",omitempty"`
	DataArea        RespondProductionPerformanceTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                   `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                   `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// RespondProductionPerformanceTypeDataArea was auto-generated. Do not change.
type RespondProductionPerformanceTypeDataArea struct {
	Respond               TransRespondType            `json:",omitempty" xml:",omitempty"`
	ProductionPerformance []ProductionPerformanceType `json:",omitempty" xml:",omitempty"`
}

// CancelProductionPerformanceType was auto-generated. Do not change.
type CancelProductionPerformanceType struct {
	ApplicationArea TransApplicationAreaType                `json:",omitempty" xml:",omitempty"`
	DataArea        CancelProductionPerformanceTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                  `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                  `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// CancelProductionPerformanceTypeDataArea was auto-generated. Do not change.
type CancelProductionPerformanceTypeDataArea struct {
	Cancel                []TransActionCriteriaType   `json:",omitempty" xml:",omitempty"`
	ProductionPerformance []ProductionPerformanceType `json:",omitempty" xml:",omitempty"`
}

// SyncProductionPerformanceType was auto-generated. Do not change.
type SyncProductionPerformanceType struct {
	ApplicationArea TransApplicationAreaType              `json:",omitempty" xml:",omitempty"`
	DataArea        SyncProductionPerformanceTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// SyncProductionPerformanceTypeDataArea was auto-generated. Do not change.
type SyncProductionPerformanceTypeDataArea struct {
	Sync                  []TransActionCriteriaType   `json:",omitempty" xml:",omitempty"`
	ProductionPerformance []ProductionPerformanceType `json:",omitempty" xml:",omitempty"`
}

// ProductionScheduleType was auto-generated. Do not change.
type ProductionScheduleType struct {
	ID                    IdentifierType            `json:",omitempty" xml:",omitempty"`
	Description           []DescriptionType         `json:",omitempty" xml:",omitempty"`
	Location              LocationType              `json:",omitempty" xml:",omitempty"`
	HierarchyScope        HierarchyScopeType        `json:",omitempty" xml:",omitempty"`
	PublishedDate         PublishedDateType         `json:",omitempty" xml:",omitempty"`
	StartTime             StartTimeType             `json:",omitempty" xml:",omitempty"`
	EndTime               EndTimeType               `json:",omitempty" xml:",omitempty"`
	EquipmentElementLevel EquipmentElementLevelType `json:",omitempty" xml:",omitempty"`
	ProductionRequest     []ProductionRequestType   `json:",omitempty" xml:",omitempty"`
	ScheduleState         RequestStateType          `json:",omitempty" xml:",omitempty"`
}

// ProductionRequestType was auto-generated. Do not change.
type ProductionRequestType struct {
	ID                      IdentifierType                `json:",omitempty" xml:",omitempty"`
	Description             []DescriptionType             `json:",omitempty" xml:",omitempty"`
	ProductProductionRuleID []ProductProductionRuleIDType `json:",omitempty" xml:",omitempty"`
	Version                 []VersionType                 `json:",omitempty" xml:",omitempty"`
	Location                LocationType                  `json:",omitempty" xml:",omitempty"`
	HierarchyScope          HierarchyScopeType            `json:",omitempty" xml:",omitempty"`
	StartTime               StartTimeType                 `json:",omitempty" xml:",omitempty"`
	EndTime                 EndTimeType                   `json:",omitempty" xml:",omitempty"`
	Priority                PriorityType                  `json:",omitempty" xml:",omitempty"`
	SegmentRequirement      []SegmentRequirementType      `json:",omitempty" xml:",omitempty"`
	SegmentResponse         []SegmentResponseType         `json:",omitempty" xml:",omitempty"`
	RequestState            RequestStateType              `json:",omitempty" xml:",omitempty"`
}

// SegmentRequirementType was auto-generated. Do not change.
type SegmentRequirementType struct {
	ID                                 IdentifierType                         `json:",omitempty" xml:",omitempty"`
	ProductSegmentID                   ProductSegmentIDType                   `json:",omitempty" xml:",omitempty"`
	ProcessSegmentID                   ProcessSegmentIDType                   `json:",omitempty" xml:",omitempty"`
	Description                        []DescriptionType                      `json:",omitempty" xml:",omitempty"`
	Location                           LocationType                           `json:",omitempty" xml:",omitempty"`
	HierarchyScope                     HierarchyScopeType                     `json:",omitempty" xml:",omitempty"`
	EarliestStartTime                  EarliestStartTimeType                  `json:",omitempty" xml:",omitempty"`
	LatestEndTime                      LatestEndTimeType                      `json:",omitempty" xml:",omitempty"`
	Duration                           string                                 `json:",omitempty" xml:",omitempty"`
	ProductionParameter                []ProductionParameterType              `json:",omitempty" xml:",omitempty"`
	PersonnelRequirement               []PersonnelRequirementType             `json:",omitempty" xml:",omitempty"`
	EquipmentRequirement               []EquipmentRequirementType             `json:",omitempty" xml:",omitempty"`
	PhysicalAssetRequirement           []PhysicalAssetRequirementType         `json:",omitempty" xml:",omitempty"`
	MaterialRequirement                []MaterialRequirementType              `json:",omitempty" xml:",omitempty"`
	MaterialProducedRequirement        []MaterialProducedRequirementType      `json:",omitempty" xml:",omitempty"`
	MaterialConsumedRequirement        []MaterialConsumedRequirementType      `json:",omitempty" xml:",omitempty"`
	ConsumableExpectedRequirement      []ConsumableExpectedRequirementType    `json:",omitempty" xml:",omitempty"`
	SegmentRequirement                 []SegmentRequirementType               `json:",omitempty" xml:",omitempty"`
	RequiredByRequestedSegmentResponse RequiredByRequestedSegmentResponseType `json:",omitempty" xml:",omitempty"`
	SegmentState                       RequestStateType                       `json:",omitempty" xml:",omitempty"`
}

// ProductionParameterType was auto-generated. Do not change.
type ProductionParameterType struct {
	ProductSegmentID ProductSegmentIDType `json:",omitempty" xml:",omitempty"`
	ProcessSegmentID ProcessSegmentIDType `json:",omitempty" xml:",omitempty"`
	Parameter        ParameterType        `json:",omitempty" xml:",omitempty"`
}

// PersonnelRequirementType was auto-generated. Do not change.
type PersonnelRequirementType struct {
	PersonnelClassID                   []PersonnelClassIDType                 `json:",omitempty" xml:",omitempty"`
	PersonID                           []PersonIDType                         `json:",omitempty" xml:",omitempty"`
	Description                        []DescriptionType                      `json:",omitempty" xml:",omitempty"`
	Location                           LocationType                           `json:",omitempty" xml:",omitempty"`
	HierarchyScope                     HierarchyScopeType                     `json:",omitempty" xml:",omitempty"`
	Quantity                           []QuantityValueType                    `json:",omitempty" xml:",omitempty"`
	PersonnelRequirementProperty       []PersonnelRequirementPropertyType     `json:",omitempty" xml:",omitempty"`
	RequiredByRequestedSegmentResponse RequiredByRequestedSegmentResponseType `json:",omitempty" xml:",omitempty"`
}

// PersonnelRequirementPropertyType was auto-generated. Do not change.
type PersonnelRequirementPropertyType struct {
	ID          IdentifierType      `json:",omitempty" xml:",omitempty"`
	Description []DescriptionType   `json:",omitempty" xml:",omitempty"`
	Value       []ValueType         `json:",omitempty" xml:",omitempty"`
	Quantity    []QuantityValueType `json:",omitempty" xml:",omitempty"`
}

// EquipmentRequirementType was auto-generated. Do not change.
type EquipmentRequirementType struct {
	EquipmentClassID                   []EquipmentClassIDType                 `json:",omitempty" xml:",omitempty"`
	EquipmentID                        []EquipmentIDType                      `json:",omitempty" xml:",omitempty"`
	Description                        []DescriptionType                      `json:",omitempty" xml:",omitempty"`
	Location                           LocationType                           `json:",omitempty" xml:",omitempty"`
	HierarchyScope                     HierarchyScopeType                     `json:",omitempty" xml:",omitempty"`
	Quantity                           []QuantityValueType                    `json:",omitempty" xml:",omitempty"`
	EquipmentRequirementProperty       []EquipmentRequirementPropertyType     `json:",omitempty" xml:",omitempty"`
	RequiredByRequestedSegmentResponse RequiredByRequestedSegmentResponseType `json:",omitempty" xml:",omitempty"`
}

// EquipmentRequirementPropertyType was auto-generated. Do not change.
type EquipmentRequirementPropertyType struct {
	ID          IdentifierType      `json:",omitempty" xml:",omitempty"`
	Description []DescriptionType   `json:",omitempty" xml:",omitempty"`
	Value       []ValueType         `json:",omitempty" xml:",omitempty"`
	Quantity    []QuantityValueType `json:",omitempty" xml:",omitempty"`
}

// PhysicalAssetRequirementType was auto-generated. Do not change.
type PhysicalAssetRequirementType struct {
	PhysicalAssetClassID               []PhysicalAssetClassIDType             `json:",omitempty" xml:",omitempty"`
	PhysicalAssetID                    []PhysicalAssetIDType                  `json:",omitempty" xml:",omitempty"`
	Description                        []DescriptionType                      `json:",omitempty" xml:",omitempty"`
	Location                           LocationType                           `json:",omitempty" xml:",omitempty"`
	HierarchyScope                     HierarchyScopeType                     `json:",omitempty" xml:",omitempty"`
	Quantity                           []QuantityValueType                    `json:",omitempty" xml:",omitempty"`
	PhysicalAssetRequirementProperty   []PhysicalAssetRequirementPropertyType `json:",omitempty" xml:",omitempty"`
	RequiredByRequestedSegmentResponse RequiredByRequestedSegmentResponseType `json:",omitempty" xml:",omitempty"`
}

// PhysicalAssetRequirementPropertyType was auto-generated. Do not change.
type PhysicalAssetRequirementPropertyType struct {
	ID          IdentifierType      `json:",omitempty" xml:",omitempty"`
	Description []DescriptionType   `json:",omitempty" xml:",omitempty"`
	Value       []ValueType         `json:",omitempty" xml:",omitempty"`
	Quantity    []QuantityValueType `json:",omitempty" xml:",omitempty"`
}

// MaterialRequirementType was auto-generated. Do not change.
type MaterialRequirementType struct {
	MaterialClassID                    []MaterialClassIDType                  `json:",omitempty" xml:",omitempty"`
	MaterialDefinitionID               []MaterialDefinitionIDType             `json:",omitempty" xml:",omitempty"`
	MaterialLotID                      []MaterialLotIDType                    `json:",omitempty" xml:",omitempty"`
	MaterialSubLotID                   []MaterialSubLotIDType                 `json:",omitempty" xml:",omitempty"`
	Description                        []DescriptionType                      `json:",omitempty" xml:",omitempty"`
	Location                           LocationType                           `json:",omitempty" xml:",omitempty"`
	HierarchyScope                     HierarchyScopeType                     `json:",omitempty" xml:",omitempty"`
	MaterialUse                        MaterialUseType                        `json:",omitempty" xml:",omitempty"`
	Quantity                           []QuantityValueType                    `json:",omitempty" xml:",omitempty"`
	AssemblyRequirement                []MaterialRequirementType              `json:",omitempty" xml:",omitempty"`
	AssemblyType                       AssemblyTypeType                       `json:",omitempty" xml:",omitempty"`
	AssemblyRelationship               AssemblyRelationshipType               `json:",omitempty" xml:",omitempty"`
	MaterialRequirementProperty        []MaterialRequirementPropertyType      `json:",omitempty" xml:",omitempty"`
	RequiredByRequestedSegmentResponse RequiredByRequestedSegmentResponseType `json:",omitempty" xml:",omitempty"`
}

// MaterialRequirementPropertyType was auto-generated. Do not change.
type MaterialRequirementPropertyType struct {
	ID          IdentifierType      `json:",omitempty" xml:",omitempty"`
	Description []DescriptionType   `json:",omitempty" xml:",omitempty"`
	Value       []ValueType         `json:",omitempty" xml:",omitempty"`
	Quantity    []QuantityValueType `json:",omitempty" xml:",omitempty"`
}

// MaterialProducedRequirementType was auto-generated. Do not change.
type MaterialProducedRequirementType struct {
	MaterialClassID                     []MaterialClassIDType                     `json:",omitempty" xml:",omitempty"`
	MaterialDefinitionID                []MaterialDefinitionIDType                `json:",omitempty" xml:",omitempty"`
	MaterialLotID                       []MaterialLotIDType                       `json:",omitempty" xml:",omitempty"`
	MaterialSubLotID                    []MaterialSubLotIDType                    `json:",omitempty" xml:",omitempty"`
	Description                         []DescriptionType                         `json:",omitempty" xml:",omitempty"`
	Location                            LocationType                              `json:",omitempty" xml:",omitempty"`
	HierarchyScope                      HierarchyScopeType                        `json:",omitempty" xml:",omitempty"`
	Quantity                            []QuantityValueType                       `json:",omitempty" xml:",omitempty"`
	MaterialProducedRequirementProperty []MaterialProducedRequirementPropertyType `json:",omitempty" xml:",omitempty"`
	RequiredByRequestedSegmentResponse  RequiredByRequestedSegmentResponseType    `json:",omitempty" xml:",omitempty"`
}

// MaterialProducedRequirementPropertyType was auto-generated. Do not change.
type MaterialProducedRequirementPropertyType struct {
	ID          IdentifierType      `json:",omitempty" xml:",omitempty"`
	Description []DescriptionType   `json:",omitempty" xml:",omitempty"`
	Value       []ValueType         `json:",omitempty" xml:",omitempty"`
	Quantity    []QuantityValueType `json:",omitempty" xml:",omitempty"`
}

// MaterialConsumedRequirementType was auto-generated. Do not change.
type MaterialConsumedRequirementType struct {
	MaterialClassID                     []MaterialClassIDType                     `json:",omitempty" xml:",omitempty"`
	MaterialDefinitionID                []MaterialDefinitionIDType                `json:",omitempty" xml:",omitempty"`
	MaterialLotID                       []MaterialLotIDType                       `json:",omitempty" xml:",omitempty"`
	MaterialSubLotID                    []MaterialSubLotIDType                    `json:",omitempty" xml:",omitempty"`
	Description                         []DescriptionType                         `json:",omitempty" xml:",omitempty"`
	Location                            LocationType                              `json:",omitempty" xml:",omitempty"`
	HierarchyScope                      HierarchyScopeType                        `json:",omitempty" xml:",omitempty"`
	Quantity                            []QuantityValueType                       `json:",omitempty" xml:",omitempty"`
	MaterialConsumedRequirementProperty []MaterialConsumedRequirementPropertyType `json:",omitempty" xml:",omitempty"`
	RequiredByRequestedSegmentResponse  RequiredByRequestedSegmentResponseType    `json:",omitempty" xml:",omitempty"`
}

// MaterialConsumedRequirementPropertyType was auto-generated. Do not change.
type MaterialConsumedRequirementPropertyType struct {
	ID          IdentifierType      `json:",omitempty" xml:",omitempty"`
	Description []DescriptionType   `json:",omitempty" xml:",omitempty"`
	Value       []ValueType         `json:",omitempty" xml:",omitempty"`
	Quantity    []QuantityValueType `json:",omitempty" xml:",omitempty"`
}

// ConsumableExpectedRequirementType was auto-generated. Do not change.
type ConsumableExpectedRequirementType struct {
	MaterialClassID                       []MaterialClassIDType                       `json:",omitempty" xml:",omitempty"`
	MaterialDefinitionID                  []MaterialDefinitionIDType                  `json:",omitempty" xml:",omitempty"`
	Description                           []DescriptionType                           `json:",omitempty" xml:",omitempty"`
	Location                              LocationType                                `json:",omitempty" xml:",omitempty"`
	HierarchyScope                        HierarchyScopeType                          `json:",omitempty" xml:",omitempty"`
	Quantity                              []QuantityValueType                         `json:",omitempty" xml:",omitempty"`
	ConsumableExpectedRequirementProperty []ConsumableExpectedRequirementPropertyType `json:",omitempty" xml:",omitempty"`
	RequiredByRequestedSegmentResponse    RequiredByRequestedSegmentResponseType      `json:",omitempty" xml:",omitempty"`
}

// ConsumableExpectedRequirementPropertyType was auto-generated. Do not change.
type ConsumableExpectedRequirementPropertyType struct {
	ID          IdentifierType      `json:",omitempty" xml:",omitempty"`
	Description []DescriptionType   `json:",omitempty" xml:",omitempty"`
	Value       []ValueType         `json:",omitempty" xml:",omitempty"`
	Quantity    []QuantityValueType `json:",omitempty" xml:",omitempty"`
}

// GetProductionScheduleType was auto-generated. Do not change.
type GetProductionScheduleType struct {
	ApplicationArea TransApplicationAreaType          `json:",omitempty" xml:",omitempty"`
	DataArea        GetProductionScheduleTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                            `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                            `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// GetProductionScheduleTypeDataArea was auto-generated. Do not change.
type GetProductionScheduleTypeDataArea struct {
	Get                []string                 `json:",omitempty" xml:",omitempty"`
	ProductionSchedule []ProductionScheduleType `json:",omitempty" xml:",omitempty"`
}

// ShowProductionScheduleType was auto-generated. Do not change.
type ShowProductionScheduleType struct {
	ApplicationArea TransApplicationAreaType           `json:",omitempty" xml:",omitempty"`
	DataArea        ShowProductionScheduleTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                             `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                             `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ShowProductionScheduleTypeDataArea was auto-generated. Do not change.
type ShowProductionScheduleTypeDataArea struct {
	Show               TransShowType            `json:",omitempty" xml:",omitempty"`
	ProductionSchedule []ProductionScheduleType `json:",omitempty" xml:",omitempty"`
}

// ProcessProductionScheduleType was auto-generated. Do not change.
type ProcessProductionScheduleType struct {
	ApplicationArea TransApplicationAreaType              `json:",omitempty" xml:",omitempty"`
	DataArea        ProcessProductionScheduleTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ProcessProductionScheduleTypeDataArea was auto-generated. Do not change.
type ProcessProductionScheduleTypeDataArea struct {
	Process            TransProcessType         `json:",omitempty" xml:",omitempty"`
	ProductionSchedule []ProductionScheduleType `json:",omitempty" xml:",omitempty"`
}

// AcknowledgeProductionScheduleType was auto-generated. Do not change.
type AcknowledgeProductionScheduleType struct {
	ApplicationArea TransApplicationAreaType                  `json:",omitempty" xml:",omitempty"`
	DataArea        AcknowledgeProductionScheduleTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                    `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                    `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// AcknowledgeProductionScheduleTypeDataArea was auto-generated. Do not change.
type AcknowledgeProductionScheduleTypeDataArea struct {
	Acknowledge        TransAcknowledgeType     `json:",omitempty" xml:",omitempty"`
	ProductionSchedule []ProductionScheduleType `json:",omitempty" xml:",omitempty"`
}

// ChangeProductionScheduleType was auto-generated. Do not change.
type ChangeProductionScheduleType struct {
	ApplicationArea TransApplicationAreaType             `json:",omitempty" xml:",omitempty"`
	DataArea        ChangeProductionScheduleTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                               `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                               `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ChangeProductionScheduleTypeDataArea was auto-generated. Do not change.
type ChangeProductionScheduleTypeDataArea struct {
	Change             TransChangeType          `json:",omitempty" xml:",omitempty"`
	ProductionSchedule []ProductionScheduleType `json:",omitempty" xml:",omitempty"`
}

// RespondProductionScheduleType was auto-generated. Do not change.
type RespondProductionScheduleType struct {
	ApplicationArea TransApplicationAreaType              `json:",omitempty" xml:",omitempty"`
	DataArea        RespondProductionScheduleTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// RespondProductionScheduleTypeDataArea was auto-generated. Do not change.
type RespondProductionScheduleTypeDataArea struct {
	Respond            TransRespondType         `json:",omitempty" xml:",omitempty"`
	ProductionSchedule []ProductionScheduleType `json:",omitempty" xml:",omitempty"`
}

// CancelProductionScheduleType was auto-generated. Do not change.
type CancelProductionScheduleType struct {
	ApplicationArea TransApplicationAreaType             `json:",omitempty" xml:",omitempty"`
	DataArea        CancelProductionScheduleTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                               `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                               `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// CancelProductionScheduleTypeDataArea was auto-generated. Do not change.
type CancelProductionScheduleTypeDataArea struct {
	Cancel             []TransActionCriteriaType `json:",omitempty" xml:",omitempty"`
	ProductionSchedule []ProductionScheduleType  `json:",omitempty" xml:",omitempty"`
}

// SyncProductionScheduleType was auto-generated. Do not change.
type SyncProductionScheduleType struct {
	ApplicationArea TransApplicationAreaType           `json:",omitempty" xml:",omitempty"`
	DataArea        SyncProductionScheduleTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                             `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                             `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// SyncProductionScheduleTypeDataArea was auto-generated. Do not change.
type SyncProductionScheduleTypeDataArea struct {
	Sync               []TransActionCriteriaType `json:",omitempty" xml:",omitempty"`
	ProductionSchedule []ProductionScheduleType  `json:",omitempty" xml:",omitempty"`
}

// ResourceRelationshipNetworkType was auto-generated. Do not change.
type ResourceRelationshipNetworkType struct {
	ID                        IdentifierType                  `json:",omitempty" xml:",omitempty"`
	Description               []DescriptionType               `json:",omitempty" xml:",omitempty"`
	HierarchyScope            HierarchyScopeType              `json:",omitempty" xml:",omitempty"`
	RelationshipType          RelationshipTypeType            `json:",omitempty" xml:",omitempty"`
	RelationshipForm          RelationshipFormType            `json:",omitempty" xml:",omitempty"`
	PublishedDate             PublishedDateType               `json:",omitempty" xml:",omitempty"`
	ResourceNetworkConnection []ResourceNetworkConnectionType `json:",omitempty" xml:",omitempty"`
}

// ResourceNetworkConnectionType was auto-generated. Do not change.
type ResourceNetworkConnectionType struct {
	ID                          IdentifierType                  `json:",omitempty" xml:",omitempty"`
	Description                 []DescriptionType               `json:",omitempty" xml:",omitempty"`
	ResourceNetworkConnectionID ResourceNetworkConnectionIDType `json:",omitempty" xml:",omitempty"`
	FromResourceReference       ResourceReferenceType           `json:",omitempty" xml:",omitempty"`
	ToResourceReference         ResourceReferenceType           `json:",omitempty" xml:",omitempty"`
	ConnectionProperty          []ResourcePropertyType          `json:",omitempty" xml:",omitempty"`
}

// ResourceReferenceType was auto-generated. Do not change.
type ResourceReferenceType struct {
	ID               IdentifierType            `json:",omitempty" xml:",omitempty"`
	ResourceID       ResourceIDType            `json:",omitempty" xml:",omitempty"`
	ResourceType     ResourceReferenceTypeType `json:",omitempty" xml:",omitempty"`
	ResourceProperty []ResourcePropertyType    `json:",omitempty" xml:",omitempty"`
}

// ResourcePropertyType was auto-generated. Do not change.
type ResourcePropertyType struct {
	ID          IdentifierType    `json:",omitempty" xml:",omitempty"`
	Description []DescriptionType `json:",omitempty" xml:",omitempty"`
	Value       []ValueType       `json:",omitempty" xml:",omitempty"`
}

// ResourceNetworkConnectionInformationType was auto-generated. Do not change.
type ResourceNetworkConnectionInformationType struct {
	ID                            IdentifierType                      `json:",omitempty" xml:",omitempty"`
	Description                   []DescriptionType                   `json:",omitempty" xml:",omitempty"`
	HierarchyScope                HierarchyScopeType                  `json:",omitempty" xml:",omitempty"`
	ResourceNetworkConnectionType []ResourceNetworkConnectionTypeType `json:",omitempty" xml:",omitempty"`
}

// ResourceNetworkConnectionTypeType was auto-generated. Do not change.
type ResourceNetworkConnectionTypeType struct {
	ID                          IdentifierType                  `json:",omitempty" xml:",omitempty"`
	Description                 []DescriptionType               `json:",omitempty" xml:",omitempty"`
	ResourceNetworkConnectionID ResourceNetworkConnectionIDType `json:",omitempty" xml:",omitempty"`
	ConnectionProperty          []ResourcePropertyType          `json:",omitempty" xml:",omitempty"`
}

// GetResourceRelationshipNetworkType was auto-generated. Do not change.
type GetResourceRelationshipNetworkType struct {
	ApplicationArea TransApplicationAreaType                   `json:",omitempty" xml:",omitempty"`
	DataArea        GetResourceRelationshipNetworkTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                     `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                     `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// GetResourceRelationshipNetworkTypeDataArea was auto-generated. Do not change.
type GetResourceRelationshipNetworkTypeDataArea struct {
	Get                         []string                          `json:",omitempty" xml:",omitempty"`
	ResourceRelationshipNetwork []ResourceRelationshipNetworkType `json:",omitempty" xml:",omitempty"`
}

// ShowResourceRelationshipNetworkType was auto-generated. Do not change.
type ShowResourceRelationshipNetworkType struct {
	ApplicationArea TransApplicationAreaType                    `json:",omitempty" xml:",omitempty"`
	DataArea        ShowResourceRelationshipNetworkTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                      `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                      `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ShowResourceRelationshipNetworkTypeDataArea was auto-generated. Do not change.
type ShowResourceRelationshipNetworkTypeDataArea struct {
	Show                        TransShowType                     `json:",omitempty" xml:",omitempty"`
	ResourceRelationshipNetwork []ResourceRelationshipNetworkType `json:",omitempty" xml:",omitempty"`
}

// ProcessResourceRelationshipNetworkType was auto-generated. Do not change.
type ProcessResourceRelationshipNetworkType struct {
	ApplicationArea TransApplicationAreaType                       `json:",omitempty" xml:",omitempty"`
	DataArea        ProcessResourceRelationshipNetworkTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                         `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                         `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ProcessResourceRelationshipNetworkTypeDataArea was auto-generated. Do not change.
type ProcessResourceRelationshipNetworkTypeDataArea struct {
	Process                     TransProcessType                  `json:",omitempty" xml:",omitempty"`
	ResourceRelationshipNetwork []ResourceRelationshipNetworkType `json:",omitempty" xml:",omitempty"`
}

// AcknowledgeResourceRelationshipNetworkType was auto-generated. Do not change.
type AcknowledgeResourceRelationshipNetworkType struct {
	ApplicationArea TransApplicationAreaType                           `json:",omitempty" xml:",omitempty"`
	DataArea        AcknowledgeResourceRelationshipNetworkTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                             `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                             `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// AcknowledgeResourceRelationshipNetworkTypeDataArea was auto-generated. Do not change.
type AcknowledgeResourceRelationshipNetworkTypeDataArea struct {
	Acknowledge                 TransAcknowledgeType              `json:",omitempty" xml:",omitempty"`
	ResourceRelationshipNetwork []ResourceRelationshipNetworkType `json:",omitempty" xml:",omitempty"`
}

// ChangeResourceRelationshipNetworkType was auto-generated. Do not change.
type ChangeResourceRelationshipNetworkType struct {
	ApplicationArea TransApplicationAreaType                      `json:",omitempty" xml:",omitempty"`
	DataArea        ChangeResourceRelationshipNetworkTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                        `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                        `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ChangeResourceRelationshipNetworkTypeDataArea was auto-generated. Do not change.
type ChangeResourceRelationshipNetworkTypeDataArea struct {
	Change                      TransChangeType                   `json:",omitempty" xml:",omitempty"`
	ResourceRelationshipNetwork []ResourceRelationshipNetworkType `json:",omitempty" xml:",omitempty"`
}

// RespondResourceRelationshipNetworkType was auto-generated. Do not change.
type RespondResourceRelationshipNetworkType struct {
	ApplicationArea TransApplicationAreaType                       `json:",omitempty" xml:",omitempty"`
	DataArea        RespondResourceRelationshipNetworkTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                         `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                         `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// RespondResourceRelationshipNetworkTypeDataArea was auto-generated. Do not change.
type RespondResourceRelationshipNetworkTypeDataArea struct {
	Respond                     TransRespondType                  `json:",omitempty" xml:",omitempty"`
	ResourceRelationshipNetwork []ResourceRelationshipNetworkType `json:",omitempty" xml:",omitempty"`
}

// CancelResourceRelationshipNetworkType was auto-generated. Do not change.
type CancelResourceRelationshipNetworkType struct {
	ApplicationArea TransApplicationAreaType                      `json:",omitempty" xml:",omitempty"`
	DataArea        CancelResourceRelationshipNetworkTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                        `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                        `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// CancelResourceRelationshipNetworkTypeDataArea was auto-generated. Do not change.
type CancelResourceRelationshipNetworkTypeDataArea struct {
	Cancel                      []TransActionCriteriaType         `json:",omitempty" xml:",omitempty"`
	ResourceRelationshipNetwork []ResourceRelationshipNetworkType `json:",omitempty" xml:",omitempty"`
}

// SyncResourceRelationshipNetworkType was auto-generated. Do not change.
type SyncResourceRelationshipNetworkType struct {
	ApplicationArea TransApplicationAreaType                    `json:",omitempty" xml:",omitempty"`
	DataArea        SyncResourceRelationshipNetworkTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                      `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                      `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// SyncResourceRelationshipNetworkTypeDataArea was auto-generated. Do not change.
type SyncResourceRelationshipNetworkTypeDataArea struct {
	Sync                        []TransActionCriteriaType         `json:",omitempty" xml:",omitempty"`
	ResourceRelationshipNetwork []ResourceRelationshipNetworkType `json:",omitempty" xml:",omitempty"`
}

// GetResourceNetworkConnectionInformationType was auto-generated. Do not change.
type GetResourceNetworkConnectionInformationType struct {
	ApplicationArea TransApplicationAreaType                            `json:",omitempty" xml:",omitempty"`
	DataArea        GetResourceNetworkConnectionInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                              `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                              `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// GetResourceNetworkConnectionInformationTypeDataArea was auto-generated. Do not change.
type GetResourceNetworkConnectionInformationTypeDataArea struct {
	Get                                  []string                                   `json:",omitempty" xml:",omitempty"`
	ResourceNetworkConnectionInformation []ResourceNetworkConnectionInformationType `json:",omitempty" xml:",omitempty"`
}

// ShowResourceNetworkConnectionInformationType was auto-generated. Do not change.
type ShowResourceNetworkConnectionInformationType struct {
	ApplicationArea TransApplicationAreaType                             `json:",omitempty" xml:",omitempty"`
	DataArea        ShowResourceNetworkConnectionInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                               `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                               `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ShowResourceNetworkConnectionInformationTypeDataArea was auto-generated. Do not change.
type ShowResourceNetworkConnectionInformationTypeDataArea struct {
	Show                                 TransShowType                              `json:",omitempty" xml:",omitempty"`
	ResourceNetworkConnectionInformation []ResourceNetworkConnectionInformationType `json:",omitempty" xml:",omitempty"`
}

// ProcessResourceNetworkConnectionInformationType was auto-generated. Do not change.
type ProcessResourceNetworkConnectionInformationType struct {
	ApplicationArea TransApplicationAreaType                                `json:",omitempty" xml:",omitempty"`
	DataArea        ProcessResourceNetworkConnectionInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                                  `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                                  `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ProcessResourceNetworkConnectionInformationTypeDataArea was auto-generated. Do not change.
type ProcessResourceNetworkConnectionInformationTypeDataArea struct {
	Process                              TransProcessType                           `json:",omitempty" xml:",omitempty"`
	ResourceNetworkConnectionInformation []ResourceNetworkConnectionInformationType `json:",omitempty" xml:",omitempty"`
}

// AcknowledgeResourceNetworkConnectionInformationType was auto-generated. Do not change.
type AcknowledgeResourceNetworkConnectionInformationType struct {
	ApplicationArea TransApplicationAreaType                                    `json:",omitempty" xml:",omitempty"`
	DataArea        AcknowledgeResourceNetworkConnectionInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                                      `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                                      `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// AcknowledgeResourceNetworkConnectionInformationTypeDataArea was auto-generated. Do not change.
type AcknowledgeResourceNetworkConnectionInformationTypeDataArea struct {
	Acknowledge                          TransAcknowledgeType                       `json:",omitempty" xml:",omitempty"`
	ResourceNetworkConnectionInformation []ResourceNetworkConnectionInformationType `json:",omitempty" xml:",omitempty"`
}

// ChangeResourceNetworkConnectionInformationType was auto-generated. Do not change.
type ChangeResourceNetworkConnectionInformationType struct {
	ApplicationArea TransApplicationAreaType                               `json:",omitempty" xml:",omitempty"`
	DataArea        ChangeResourceNetworkConnectionInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                                 `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                                 `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ChangeResourceNetworkConnectionInformationTypeDataArea was auto-generated. Do not change.
type ChangeResourceNetworkConnectionInformationTypeDataArea struct {
	Change                               TransChangeType                            `json:",omitempty" xml:",omitempty"`
	ResourceNetworkConnectionInformation []ResourceNetworkConnectionInformationType `json:",omitempty" xml:",omitempty"`
}

// RespondResourceNetworkConnectionInformationType was auto-generated. Do not change.
type RespondResourceNetworkConnectionInformationType struct {
	ApplicationArea TransApplicationAreaType                                `json:",omitempty" xml:",omitempty"`
	DataArea        RespondResourceNetworkConnectionInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                                  `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                                  `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// RespondResourceNetworkConnectionInformationTypeDataArea was auto-generated. Do not change.
type RespondResourceNetworkConnectionInformationTypeDataArea struct {
	Respond                              TransRespondType                           `json:",omitempty" xml:",omitempty"`
	ResourceNetworkConnectionInformation []ResourceNetworkConnectionInformationType `json:",omitempty" xml:",omitempty"`
}

// CancelResourceNetworkConnectionInformationType was auto-generated. Do not change.
type CancelResourceNetworkConnectionInformationType struct {
	ApplicationArea TransApplicationAreaType                               `json:",omitempty" xml:",omitempty"`
	DataArea        CancelResourceNetworkConnectionInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                                 `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                                 `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// CancelResourceNetworkConnectionInformationTypeDataArea was auto-generated. Do not change.
type CancelResourceNetworkConnectionInformationTypeDataArea struct {
	Cancel                               []TransActionCriteriaType                  `json:",omitempty" xml:",omitempty"`
	ResourceNetworkConnectionInformation []ResourceNetworkConnectionInformationType `json:",omitempty" xml:",omitempty"`
}

// SyncResourceNetworkConnectionInformationType was auto-generated. Do not change.
type SyncResourceNetworkConnectionInformationType struct {
	ApplicationArea TransApplicationAreaType                             `json:",omitempty" xml:",omitempty"`
	DataArea        SyncResourceNetworkConnectionInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                               `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                               `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// SyncResourceNetworkConnectionInformationTypeDataArea was auto-generated. Do not change.
type SyncResourceNetworkConnectionInformationTypeDataArea struct {
	Sync                                 []TransActionCriteriaType                  `json:",omitempty" xml:",omitempty"`
	ResourceNetworkConnectionInformation []ResourceNetworkConnectionInformationType `json:",omitempty" xml:",omitempty"`
}

// TransactionProfileType was auto-generated. Do not change.
type TransactionProfileType struct {
	ID              IdentifierType        `json:",omitempty" xml:",omitempty"`
	Description     []DescriptionType     `json:",omitempty" xml:",omitempty"`
	Location        LocationType          `json:",omitempty" xml:",omitempty"`
	HierarchyScope  HierarchyScopeType    `json:",omitempty" xml:",omitempty"`
	PublishedDate   PublishedDateType     `json:",omitempty" xml:",omitempty"`
	SupportedAction []SupportedActionType `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// SupportedActionType was auto-generated. Do not change.
type SupportedActionType struct {
	ID                                 IdentifierType      `json:",omitempty" xml:",omitempty"`
	Description                        []DescriptionType   `json:",omitempty" xml:",omitempty"`
	TransactionVerb                    TransactionVerbType `json:",omitempty" xml:",omitempty"`
	TransactionNoun                    TransactionNounType `json:",omitempty" xml:",omitempty"`
	InformationUser                    bool                `json:",omitempty" xml:",omitempty"`
	InformationUserSpecified           bool                `json:",omitempty" xml:",omitempty"`
	InformationProvider                bool                `json:",omitempty" xml:",omitempty"`
	InformationProviderSpecified       bool                `json:",omitempty" xml:",omitempty"`
	InformationSender                  bool                `json:",omitempty" xml:",omitempty"`
	InformationSenderSpecified         bool                `json:",omitempty" xml:",omitempty"`
	InformationReceiver                bool                `json:",omitempty" xml:",omitempty"`
	InformationReceiverSpecified       bool                `json:",omitempty" xml:",omitempty"`
	ObjectWildcardSupported            bool                `json:",omitempty" xml:",omitempty"`
	ObjectWildcardSupportedSpecified   bool                `json:",omitempty" xml:",omitempty"`
	PropertyWildcardSupported          bool                `json:",omitempty" xml:",omitempty"`
	PropertyWildcardSupportedSpecified bool                `json:",omitempty" xml:",omitempty"`
	ReleaseID                          string              `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID                          string              `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// GetTransactionProfileType was auto-generated. Do not change.
type GetTransactionProfileType struct {
	ApplicationArea TransApplicationAreaType          `json:",omitempty" xml:",omitempty"`
	DataArea        GetTransactionProfileTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                            `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                            `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// GetTransactionProfileTypeDataArea was auto-generated. Do not change.
type GetTransactionProfileTypeDataArea struct {
	Get                []string                 `json:",omitempty" xml:",omitempty"`
	TransactionProfile []TransactionProfileType `json:",omitempty" xml:",omitempty"`
}

// ShowTransactionProfileType was auto-generated. Do not change.
type ShowTransactionProfileType struct {
	ApplicationArea TransApplicationAreaType           `json:",omitempty" xml:",omitempty"`
	DataArea        ShowTransactionProfileTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                             `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                             `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ShowTransactionProfileTypeDataArea was auto-generated. Do not change.
type ShowTransactionProfileTypeDataArea struct {
	Show               TransShowType            `json:",omitempty" xml:",omitempty"`
	TransactionProfile []TransactionProfileType `json:",omitempty" xml:",omitempty"`
}

// WorkAlertInformationType was auto-generated. Do not change.
type WorkAlertInformationType struct {
	ID                  IdentifierType            `json:",omitempty" xml:",omitempty"`
	Description         []DescriptionType         `json:",omitempty" xml:",omitempty"`
	HierarchyScope      HierarchyScopeType        `json:",omitempty" xml:",omitempty"`
	PublishedDate       PublishedDateType         `json:",omitempty" xml:",omitempty"`
	WorkAlertDefinition []WorkAlertDefinitionType `json:",omitempty" xml:",omitempty"`
	WorkAlert           []WorkAlertType           `json:",omitempty" xml:",omitempty"`
}

// WorkAlertDefinitionType was auto-generated. Do not change.
type WorkAlertDefinitionType struct {
	ID             IdentifierType          `json:",omitempty" xml:",omitempty"`
	Description    []DescriptionType       `json:",omitempty" xml:",omitempty"`
	HierarchyScope HierarchyScopeType      `json:",omitempty" xml:",omitempty"`
	Priority       []PriorityType          `json:",omitempty" xml:",omitempty"`
	Category       []IdentifierType        `json:",omitempty" xml:",omitempty"`
	Property       []WorkAlertPropertyType `json:",omitempty" xml:",omitempty"`
}

// WorkAlertPropertyType was auto-generated. Do not change.
type WorkAlertPropertyType struct {
	ID          IdentifierType    `json:",omitempty" xml:",omitempty"`
	Description []DescriptionType `json:",omitempty" xml:",omitempty"`
	Value       []ValueType       `json:",omitempty" xml:",omitempty"`
}

// WorkAlertType was auto-generated. Do not change.
type WorkAlertType struct {
	ID             IdentifierType          `json:",omitempty" xml:",omitempty"`
	MessageText    []DescriptionType       `json:",omitempty" xml:",omitempty"`
	HierarchyScope HierarchyScopeType      `json:",omitempty" xml:",omitempty"`
	TimeStamp      StartTimeType           `json:",omitempty" xml:",omitempty"`
	Priority       PriorityType            `json:",omitempty" xml:",omitempty"`
	Category       IdentifierType          `json:",omitempty" xml:",omitempty"`
	Property       []WorkAlertPropertyType `json:",omitempty" xml:",omitempty"`
}

// GetWorkAlertInformationType was auto-generated. Do not change.
type GetWorkAlertInformationType struct {
	ApplicationArea TransApplicationAreaType            `json:",omitempty" xml:",omitempty"`
	DataArea        GetWorkAlertInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                              `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                              `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// GetWorkAlertInformationTypeDataArea was auto-generated. Do not change.
type GetWorkAlertInformationTypeDataArea struct {
	Get                  []string                   `json:",omitempty" xml:",omitempty"`
	WorkAlertInformation []WorkAlertInformationType `json:",omitempty" xml:",omitempty"`
}

// ShowWorkAlertInformationType was auto-generated. Do not change.
type ShowWorkAlertInformationType struct {
	ApplicationArea TransApplicationAreaType             `json:",omitempty" xml:",omitempty"`
	DataArea        ShowWorkAlertInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                               `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                               `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ShowWorkAlertInformationTypeDataArea was auto-generated. Do not change.
type ShowWorkAlertInformationTypeDataArea struct {
	Show                 TransShowType              `json:",omitempty" xml:",omitempty"`
	WorkAlertInformation []WorkAlertInformationType `json:",omitempty" xml:",omitempty"`
}

// ProcessWorkAlertInformationType was auto-generated. Do not change.
type ProcessWorkAlertInformationType struct {
	ApplicationArea TransApplicationAreaType                `json:",omitempty" xml:",omitempty"`
	DataArea        ProcessWorkAlertInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                  `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                  `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ProcessWorkAlertInformationTypeDataArea was auto-generated. Do not change.
type ProcessWorkAlertInformationTypeDataArea struct {
	Process              TransProcessType           `json:",omitempty" xml:",omitempty"`
	WorkAlertInformation []WorkAlertInformationType `json:",omitempty" xml:",omitempty"`
}

// AcknowledgeWorkAlertInformationType was auto-generated. Do not change.
type AcknowledgeWorkAlertInformationType struct {
	ApplicationArea TransApplicationAreaType                    `json:",omitempty" xml:",omitempty"`
	DataArea        AcknowledgeWorkAlertInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                      `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                      `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// AcknowledgeWorkAlertInformationTypeDataArea was auto-generated. Do not change.
type AcknowledgeWorkAlertInformationTypeDataArea struct {
	Acknowledge          TransAcknowledgeType       `json:",omitempty" xml:",omitempty"`
	WorkAlertInformation []WorkAlertInformationType `json:",omitempty" xml:",omitempty"`
}

// ChangeWorkAlertInformationType was auto-generated. Do not change.
type ChangeWorkAlertInformationType struct {
	ApplicationArea TransApplicationAreaType               `json:",omitempty" xml:",omitempty"`
	DataArea        ChangeWorkAlertInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                 `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                 `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ChangeWorkAlertInformationTypeDataArea was auto-generated. Do not change.
type ChangeWorkAlertInformationTypeDataArea struct {
	Change               TransChangeType            `json:",omitempty" xml:",omitempty"`
	WorkAlertInformation []WorkAlertInformationType `json:",omitempty" xml:",omitempty"`
}

// RespondWorkAlertInformationType was auto-generated. Do not change.
type RespondWorkAlertInformationType struct {
	ApplicationArea TransApplicationAreaType                `json:",omitempty" xml:",omitempty"`
	DataArea        RespondWorkAlertInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                  `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                  `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// RespondWorkAlertInformationTypeDataArea was auto-generated. Do not change.
type RespondWorkAlertInformationTypeDataArea struct {
	Respond              TransRespondType           `json:",omitempty" xml:",omitempty"`
	WorkAlertInformation []WorkAlertInformationType `json:",omitempty" xml:",omitempty"`
}

// CancelWorkAlertInformationType was auto-generated. Do not change.
type CancelWorkAlertInformationType struct {
	ApplicationArea TransApplicationAreaType               `json:",omitempty" xml:",omitempty"`
	DataArea        CancelWorkAlertInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                 `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                 `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// CancelWorkAlertInformationTypeDataArea was auto-generated. Do not change.
type CancelWorkAlertInformationTypeDataArea struct {
	Cancel               []TransActionCriteriaType  `json:",omitempty" xml:",omitempty"`
	WorkAlertInformation []WorkAlertInformationType `json:",omitempty" xml:",omitempty"`
}

// SyncWorkAlertInformationType was auto-generated. Do not change.
type SyncWorkAlertInformationType struct {
	ApplicationArea TransApplicationAreaType             `json:",omitempty" xml:",omitempty"`
	DataArea        SyncWorkAlertInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                               `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                               `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// SyncWorkAlertInformationTypeDataArea was auto-generated. Do not change.
type SyncWorkAlertInformationTypeDataArea struct {
	Sync                 []TransActionCriteriaType  `json:",omitempty" xml:",omitempty"`
	WorkAlertInformation []WorkAlertInformationType `json:",omitempty" xml:",omitempty"`
}

// GetWorkAlertDefinitionType was auto-generated. Do not change.
type GetWorkAlertDefinitionType struct {
	ApplicationArea TransApplicationAreaType           `json:",omitempty" xml:",omitempty"`
	DataArea        GetWorkAlertDefinitionTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                             `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                             `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// GetWorkAlertDefinitionTypeDataArea was auto-generated. Do not change.
type GetWorkAlertDefinitionTypeDataArea struct {
	Get                 []string                  `json:",omitempty" xml:",omitempty"`
	WorkAlertDefinition []WorkAlertDefinitionType `json:",omitempty" xml:",omitempty"`
}

// ShowWorkAlertDefinitionType was auto-generated. Do not change.
type ShowWorkAlertDefinitionType struct {
	ApplicationArea TransApplicationAreaType            `json:",omitempty" xml:",omitempty"`
	DataArea        ShowWorkAlertDefinitionTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                              `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                              `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ShowWorkAlertDefinitionTypeDataArea was auto-generated. Do not change.
type ShowWorkAlertDefinitionTypeDataArea struct {
	Show                TransShowType             `json:",omitempty" xml:",omitempty"`
	WorkAlertDefinition []WorkAlertDefinitionType `json:",omitempty" xml:",omitempty"`
}

// ProcessWorkAlertDefinitionType was auto-generated. Do not change.
type ProcessWorkAlertDefinitionType struct {
	ApplicationArea TransApplicationAreaType               `json:",omitempty" xml:",omitempty"`
	DataArea        ProcessWorkAlertDefinitionTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                 `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                 `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ProcessWorkAlertDefinitionTypeDataArea was auto-generated. Do not change.
type ProcessWorkAlertDefinitionTypeDataArea struct {
	Process             TransProcessType          `json:",omitempty" xml:",omitempty"`
	WorkAlertDefinition []WorkAlertDefinitionType `json:",omitempty" xml:",omitempty"`
}

// AcknowledgeWorkAlertDefinitionType was auto-generated. Do not change.
type AcknowledgeWorkAlertDefinitionType struct {
	ApplicationArea TransApplicationAreaType                   `json:",omitempty" xml:",omitempty"`
	DataArea        AcknowledgeWorkAlertDefinitionTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                     `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                     `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// AcknowledgeWorkAlertDefinitionTypeDataArea was auto-generated. Do not change.
type AcknowledgeWorkAlertDefinitionTypeDataArea struct {
	Acknowledge         TransAcknowledgeType      `json:",omitempty" xml:",omitempty"`
	WorkAlertDefinition []WorkAlertDefinitionType `json:",omitempty" xml:",omitempty"`
}

// ChangeWorkAlertDefinitionType was auto-generated. Do not change.
type ChangeWorkAlertDefinitionType struct {
	ApplicationArea TransApplicationAreaType              `json:",omitempty" xml:",omitempty"`
	DataArea        ChangeWorkAlertDefinitionTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ChangeWorkAlertDefinitionTypeDataArea was auto-generated. Do not change.
type ChangeWorkAlertDefinitionTypeDataArea struct {
	Change              TransChangeType           `json:",omitempty" xml:",omitempty"`
	WorkAlertDefinition []WorkAlertDefinitionType `json:",omitempty" xml:",omitempty"`
}

// RespondWorkAlertDefinitionType was auto-generated. Do not change.
type RespondWorkAlertDefinitionType struct {
	ApplicationArea TransApplicationAreaType               `json:",omitempty" xml:",omitempty"`
	DataArea        RespondWorkAlertDefinitionTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                 `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                 `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// RespondWorkAlertDefinitionTypeDataArea was auto-generated. Do not change.
type RespondWorkAlertDefinitionTypeDataArea struct {
	Respond             TransRespondType          `json:",omitempty" xml:",omitempty"`
	WorkAlertDefinition []WorkAlertDefinitionType `json:",omitempty" xml:",omitempty"`
}

// CancelWorkAlertDefinitionType was auto-generated. Do not change.
type CancelWorkAlertDefinitionType struct {
	ApplicationArea TransApplicationAreaType              `json:",omitempty" xml:",omitempty"`
	DataArea        CancelWorkAlertDefinitionTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// CancelWorkAlertDefinitionTypeDataArea was auto-generated. Do not change.
type CancelWorkAlertDefinitionTypeDataArea struct {
	Cancel              []TransActionCriteriaType `json:",omitempty" xml:",omitempty"`
	WorkAlertDefinition []WorkAlertDefinitionType `json:",omitempty" xml:",omitempty"`
}

// SyncWorkAlertDefinitionType was auto-generated. Do not change.
type SyncWorkAlertDefinitionType struct {
	ApplicationArea TransApplicationAreaType            `json:",omitempty" xml:",omitempty"`
	DataArea        SyncWorkAlertDefinitionTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                              `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                              `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// SyncWorkAlertDefinitionTypeDataArea was auto-generated. Do not change.
type SyncWorkAlertDefinitionTypeDataArea struct {
	Sync                []TransActionCriteriaType `json:",omitempty" xml:",omitempty"`
	WorkAlertDefinition []WorkAlertDefinitionType `json:",omitempty" xml:",omitempty"`
}

// GetWorkAlertType was auto-generated. Do not change.
type GetWorkAlertType struct {
	ApplicationArea TransApplicationAreaType `json:",omitempty" xml:",omitempty"`
	DataArea        GetWorkAlertTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                   `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                   `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// GetWorkAlertTypeDataArea was auto-generated. Do not change.
type GetWorkAlertTypeDataArea struct {
	Get       []string        `json:",omitempty" xml:",omitempty"`
	WorkAlert []WorkAlertType `json:",omitempty" xml:",omitempty"`
}

// ShowWorkAlertType was auto-generated. Do not change.
type ShowWorkAlertType struct {
	ApplicationArea TransApplicationAreaType  `json:",omitempty" xml:",omitempty"`
	DataArea        ShowWorkAlertTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                    `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                    `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ShowWorkAlertTypeDataArea was auto-generated. Do not change.
type ShowWorkAlertTypeDataArea struct {
	Show      TransShowType   `json:",omitempty" xml:",omitempty"`
	WorkAlert []WorkAlertType `json:",omitempty" xml:",omitempty"`
}

// ProcessWorkAlertType was auto-generated. Do not change.
type ProcessWorkAlertType struct {
	ApplicationArea TransApplicationAreaType     `json:",omitempty" xml:",omitempty"`
	DataArea        ProcessWorkAlertTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                       `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                       `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ProcessWorkAlertTypeDataArea was auto-generated. Do not change.
type ProcessWorkAlertTypeDataArea struct {
	Process   TransProcessType `json:",omitempty" xml:",omitempty"`
	WorkAlert []WorkAlertType  `json:",omitempty" xml:",omitempty"`
}

// AcknowledgeWorkAlertType was auto-generated. Do not change.
type AcknowledgeWorkAlertType struct {
	ApplicationArea TransApplicationAreaType         `json:",omitempty" xml:",omitempty"`
	DataArea        AcknowledgeWorkAlertTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                           `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                           `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// AcknowledgeWorkAlertTypeDataArea was auto-generated. Do not change.
type AcknowledgeWorkAlertTypeDataArea struct {
	Acknowledge TransAcknowledgeType `json:",omitempty" xml:",omitempty"`
	WorkAlert   []WorkAlertType      `json:",omitempty" xml:",omitempty"`
}

// ChangeWorkAlertType was auto-generated. Do not change.
type ChangeWorkAlertType struct {
	ApplicationArea TransApplicationAreaType    `json:",omitempty" xml:",omitempty"`
	DataArea        ChangeWorkAlertTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                      `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                      `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ChangeWorkAlertTypeDataArea was auto-generated. Do not change.
type ChangeWorkAlertTypeDataArea struct {
	Change    TransChangeType `json:",omitempty" xml:",omitempty"`
	WorkAlert []WorkAlertType `json:",omitempty" xml:",omitempty"`
}

// RespondWorkAlertType was auto-generated. Do not change.
type RespondWorkAlertType struct {
	ApplicationArea TransApplicationAreaType     `json:",omitempty" xml:",omitempty"`
	DataArea        RespondWorkAlertTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                       `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                       `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// RespondWorkAlertTypeDataArea was auto-generated. Do not change.
type RespondWorkAlertTypeDataArea struct {
	Respond   TransRespondType `json:",omitempty" xml:",omitempty"`
	WorkAlert []WorkAlertType  `json:",omitempty" xml:",omitempty"`
}

// CancelWorkAlertType was auto-generated. Do not change.
type CancelWorkAlertType struct {
	ApplicationArea TransApplicationAreaType    `json:",omitempty" xml:",omitempty"`
	DataArea        CancelWorkAlertTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                      `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                      `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// CancelWorkAlertTypeDataArea was auto-generated. Do not change.
type CancelWorkAlertTypeDataArea struct {
	Cancel    []TransActionCriteriaType `json:",omitempty" xml:",omitempty"`
	WorkAlert []WorkAlertType           `json:",omitempty" xml:",omitempty"`
}

// SyncWorkAlertType was auto-generated. Do not change.
type SyncWorkAlertType struct {
	ApplicationArea TransApplicationAreaType  `json:",omitempty" xml:",omitempty"`
	DataArea        SyncWorkAlertTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                    `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                    `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// SyncWorkAlertTypeDataArea was auto-generated. Do not change.
type SyncWorkAlertTypeDataArea struct {
	Sync      []TransActionCriteriaType `json:",omitempty" xml:",omitempty"`
	WorkAlert []WorkAlertType           `json:",omitempty" xml:",omitempty"`
}

// WorkCapabilityInformationType was auto-generated. Do not change.
type WorkCapabilityInformationType struct {
	ID             IdentifierType       `json:",omitempty" xml:",omitempty"`
	Description    []DescriptionType    `json:",omitempty" xml:",omitempty"`
	HierarchyScope HierarchyScopeType   `json:",omitempty" xml:",omitempty"`
	PublishedDate  PublishedDateType    `json:",omitempty" xml:",omitempty"`
	WorkCapability []WorkCapabilityType `json:",omitempty" xml:",omitempty"`
}

// WorkCapabilityType was auto-generated. Do not change.
type WorkCapabilityType struct {
	ID                      IdentifierType                  `json:",omitempty" xml:",omitempty"`
	Description             []DescriptionType               `json:",omitempty" xml:",omitempty"`
	HierarchyScope          HierarchyScopeType              `json:",omitempty" xml:",omitempty"`
	CapabilityType          CapabilityTypeType              `json:",omitempty" xml:",omitempty"`
	Reason                  ReasonType                      `json:",omitempty" xml:",omitempty"`
	ConfidenceFactor        ConfidenceFactorType            `json:",omitempty" xml:",omitempty"`
	StartTime               StartTimeType                   `json:",omitempty" xml:",omitempty"`
	EndTime                 EndTimeType                     `json:",omitempty" xml:",omitempty"`
	PublishedDate           PublishedDateType               `json:",omitempty" xml:",omitempty"`
	PersonnelCapability     []OpPersonnelCapabilityType     `json:",omitempty" xml:",omitempty"`
	EquipmentCapability     []OpEquipmentCapabilityType     `json:",omitempty" xml:",omitempty"`
	PhysicalAssetCapability []OpPhysicalAssetCapabilityType `json:",omitempty" xml:",omitempty"`
	MaterialCapability      []OpMaterialCapabilityType      `json:",omitempty" xml:",omitempty"`
	WorkMasterCapability    []WorkMasterCapabilityType      `json:",omitempty" xml:",omitempty"`
}

// WorkMasterCapabilityType was auto-generated. Do not change.
type WorkMasterCapabilityType struct {
	ID                      IdentifierType                  `json:",omitempty" xml:",omitempty"`
	Description             []DescriptionType               `json:",omitempty" xml:",omitempty"`
	WorkMasterID            []IdentifierType                `json:",omitempty" xml:",omitempty"`
	CapabilityType          CapabilityTypeType              `json:",omitempty" xml:",omitempty"`
	Reason                  []ReasonType                    `json:",omitempty" xml:",omitempty"`
	HierarchyScope          []HierarchyScopeType            `json:",omitempty" xml:",omitempty"`
	StartTime               StartTimeType                   `json:",omitempty" xml:",omitempty"`
	EndTime                 EndTimeType                     `json:",omitempty" xml:",omitempty"`
	PersonnelCapability     []OpPersonnelCapabilityType     `json:",omitempty" xml:",omitempty"`
	EquipmentCapability     []OpEquipmentCapabilityType     `json:",omitempty" xml:",omitempty"`
	PhysicalAssetCapability []OpPhysicalAssetCapabilityType `json:",omitempty" xml:",omitempty"`
	MaterialCapability      []OpMaterialCapabilityType      `json:",omitempty" xml:",omitempty"`
}

// GetWorkCapabilityType was auto-generated. Do not change.
type GetWorkCapabilityType struct {
	ApplicationArea TransApplicationAreaType      `json:",omitempty" xml:",omitempty"`
	DataArea        GetWorkCapabilityTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                        `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                        `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// GetWorkCapabilityTypeDataArea was auto-generated. Do not change.
type GetWorkCapabilityTypeDataArea struct {
	Get            []string             `json:",omitempty" xml:",omitempty"`
	WorkCapability []WorkCapabilityType `json:",omitempty" xml:",omitempty"`
}

// ShowWorkCapabilityType was auto-generated. Do not change.
type ShowWorkCapabilityType struct {
	ApplicationArea TransApplicationAreaType       `json:",omitempty" xml:",omitempty"`
	DataArea        ShowWorkCapabilityTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                         `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                         `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ShowWorkCapabilityTypeDataArea was auto-generated. Do not change.
type ShowWorkCapabilityTypeDataArea struct {
	Show           TransShowType        `json:",omitempty" xml:",omitempty"`
	WorkCapability []WorkCapabilityType `json:",omitempty" xml:",omitempty"`
}

// ProcessWorkCapabilityType was auto-generated. Do not change.
type ProcessWorkCapabilityType struct {
	ApplicationArea TransApplicationAreaType          `json:",omitempty" xml:",omitempty"`
	DataArea        ProcessWorkCapabilityTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                            `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                            `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ProcessWorkCapabilityTypeDataArea was auto-generated. Do not change.
type ProcessWorkCapabilityTypeDataArea struct {
	Process        TransProcessType     `json:",omitempty" xml:",omitempty"`
	WorkCapability []WorkCapabilityType `json:",omitempty" xml:",omitempty"`
}

// AcknowledgeWorkCapabilityType was auto-generated. Do not change.
type AcknowledgeWorkCapabilityType struct {
	ApplicationArea TransApplicationAreaType              `json:",omitempty" xml:",omitempty"`
	DataArea        AcknowledgeWorkCapabilityTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// AcknowledgeWorkCapabilityTypeDataArea was auto-generated. Do not change.
type AcknowledgeWorkCapabilityTypeDataArea struct {
	Acknowledge    TransAcknowledgeType `json:",omitempty" xml:",omitempty"`
	WorkCapability []WorkCapabilityType `json:",omitempty" xml:",omitempty"`
}

// ChangeWorkCapabilityType was auto-generated. Do not change.
type ChangeWorkCapabilityType struct {
	ApplicationArea TransApplicationAreaType         `json:",omitempty" xml:",omitempty"`
	DataArea        ChangeWorkCapabilityTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                           `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                           `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ChangeWorkCapabilityTypeDataArea was auto-generated. Do not change.
type ChangeWorkCapabilityTypeDataArea struct {
	Change         TransChangeType      `json:",omitempty" xml:",omitempty"`
	WorkCapability []WorkCapabilityType `json:",omitempty" xml:",omitempty"`
}

// RespondWorkCapabilityType was auto-generated. Do not change.
type RespondWorkCapabilityType struct {
	ApplicationArea TransApplicationAreaType          `json:",omitempty" xml:",omitempty"`
	DataArea        RespondWorkCapabilityTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                            `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                            `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// RespondWorkCapabilityTypeDataArea was auto-generated. Do not change.
type RespondWorkCapabilityTypeDataArea struct {
	Respond        TransRespondType     `json:",omitempty" xml:",omitempty"`
	WorkCapability []WorkCapabilityType `json:",omitempty" xml:",omitempty"`
}

// CancelWorkCapabilityType was auto-generated. Do not change.
type CancelWorkCapabilityType struct {
	ApplicationArea TransApplicationAreaType         `json:",omitempty" xml:",omitempty"`
	DataArea        CancelWorkCapabilityTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                           `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                           `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// CancelWorkCapabilityTypeDataArea was auto-generated. Do not change.
type CancelWorkCapabilityTypeDataArea struct {
	Cancel         []TransActionCriteriaType `json:",omitempty" xml:",omitempty"`
	WorkCapability []WorkCapabilityType      `json:",omitempty" xml:",omitempty"`
}

// SyncWorkCapabilityType was auto-generated. Do not change.
type SyncWorkCapabilityType struct {
	ApplicationArea TransApplicationAreaType       `json:",omitempty" xml:",omitempty"`
	DataArea        SyncWorkCapabilityTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                         `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                         `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// SyncWorkCapabilityTypeDataArea was auto-generated. Do not change.
type SyncWorkCapabilityTypeDataArea struct {
	Sync           []TransActionCriteriaType `json:",omitempty" xml:",omitempty"`
	WorkCapability []WorkCapabilityType      `json:",omitempty" xml:",omitempty"`
}

// GetWorkCapabilityInformationType was auto-generated. Do not change.
type GetWorkCapabilityInformationType struct {
	ApplicationArea TransApplicationAreaType                 `json:",omitempty" xml:",omitempty"`
	DataArea        GetWorkCapabilityInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                   `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                   `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// GetWorkCapabilityInformationTypeDataArea was auto-generated. Do not change.
type GetWorkCapabilityInformationTypeDataArea struct {
	Get                       []string                        `json:",omitempty" xml:",omitempty"`
	WorkCapabilityInformation []WorkCapabilityInformationType `json:",omitempty" xml:",omitempty"`
}

// ShowWorkCapabilityInformationType was auto-generated. Do not change.
type ShowWorkCapabilityInformationType struct {
	ApplicationArea TransApplicationAreaType                  `json:",omitempty" xml:",omitempty"`
	DataArea        ShowWorkCapabilityInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                    `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                    `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ShowWorkCapabilityInformationTypeDataArea was auto-generated. Do not change.
type ShowWorkCapabilityInformationTypeDataArea struct {
	Show                      TransShowType                   `json:",omitempty" xml:",omitempty"`
	WorkCapabilityInformation []WorkCapabilityInformationType `json:",omitempty" xml:",omitempty"`
}

// ProcessWorkCapabilityInformationType was auto-generated. Do not change.
type ProcessWorkCapabilityInformationType struct {
	ApplicationArea TransApplicationAreaType                     `json:",omitempty" xml:",omitempty"`
	DataArea        ProcessWorkCapabilityInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                       `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                       `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ProcessWorkCapabilityInformationTypeDataArea was auto-generated. Do not change.
type ProcessWorkCapabilityInformationTypeDataArea struct {
	Process                   TransProcessType                `json:",omitempty" xml:",omitempty"`
	WorkCapabilityInformation []WorkCapabilityInformationType `json:",omitempty" xml:",omitempty"`
}

// AcknowledgeWorkCapabilityInformationType was auto-generated. Do not change.
type AcknowledgeWorkCapabilityInformationType struct {
	ApplicationArea TransApplicationAreaType                         `json:",omitempty" xml:",omitempty"`
	DataArea        AcknowledgeWorkCapabilityInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                           `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                           `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// AcknowledgeWorkCapabilityInformationTypeDataArea was auto-generated. Do not change.
type AcknowledgeWorkCapabilityInformationTypeDataArea struct {
	Acknowledge               TransAcknowledgeType            `json:",omitempty" xml:",omitempty"`
	WorkCapabilityInformation []WorkCapabilityInformationType `json:",omitempty" xml:",omitempty"`
}

// ChangeWorkCapabilityInformationType was auto-generated. Do not change.
type ChangeWorkCapabilityInformationType struct {
	ApplicationArea TransApplicationAreaType                    `json:",omitempty" xml:",omitempty"`
	DataArea        ChangeWorkCapabilityInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                      `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                      `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ChangeWorkCapabilityInformationTypeDataArea was auto-generated. Do not change.
type ChangeWorkCapabilityInformationTypeDataArea struct {
	Change                    TransChangeType                 `json:",omitempty" xml:",omitempty"`
	WorkCapabilityInformation []WorkCapabilityInformationType `json:",omitempty" xml:",omitempty"`
}

// RespondWorkCapabilityInformationType was auto-generated. Do not change.
type RespondWorkCapabilityInformationType struct {
	ApplicationArea TransApplicationAreaType                     `json:",omitempty" xml:",omitempty"`
	DataArea        RespondWorkCapabilityInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                       `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                       `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// RespondWorkCapabilityInformationTypeDataArea was auto-generated. Do not change.
type RespondWorkCapabilityInformationTypeDataArea struct {
	Respond                   TransRespondType                `json:",omitempty" xml:",omitempty"`
	WorkCapabilityInformation []WorkCapabilityInformationType `json:",omitempty" xml:",omitempty"`
}

// CancelWorkCapabilityInformationType was auto-generated. Do not change.
type CancelWorkCapabilityInformationType struct {
	ApplicationArea TransApplicationAreaType                    `json:",omitempty" xml:",omitempty"`
	DataArea        CancelWorkCapabilityInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                      `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                      `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// CancelWorkCapabilityInformationTypeDataArea was auto-generated. Do not change.
type CancelWorkCapabilityInformationTypeDataArea struct {
	Cancel                    []TransActionCriteriaType       `json:",omitempty" xml:",omitempty"`
	WorkCapabilityInformation []WorkCapabilityInformationType `json:",omitempty" xml:",omitempty"`
}

// SyncWorkCapabilityInformationType was auto-generated. Do not change.
type SyncWorkCapabilityInformationType struct {
	ApplicationArea TransApplicationAreaType                  `json:",omitempty" xml:",omitempty"`
	DataArea        SyncWorkCapabilityInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                    `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                    `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// SyncWorkCapabilityInformationTypeDataArea was auto-generated. Do not change.
type SyncWorkCapabilityInformationTypeDataArea struct {
	Sync                      []TransActionCriteriaType       `json:",omitempty" xml:",omitempty"`
	WorkCapabilityInformation []WorkCapabilityInformationType `json:",omitempty" xml:",omitempty"`
}

// WorkflowSpecificationInformationType was auto-generated. Do not change.
type WorkflowSpecificationInformationType struct {
	ID                        IdentifierType                  `json:",omitempty" xml:",omitempty"`
	Description               []DescriptionType               `json:",omitempty" xml:",omitempty"`
	HierarchyScope            HierarchyScopeType              `json:",omitempty" xml:",omitempty"`
	PublishedDate             PublishedDateType               `json:",omitempty" xml:",omitempty"`
	WorkflowSpecification     []WorkflowSpecificationType     `json:",omitempty" xml:",omitempty"`
	WorkflowSpecificationType []WorkflowSpecificationTypeType `json:",omitempty" xml:",omitempty"`
}

// WorkflowSpecificationType was auto-generated. Do not change.
type WorkflowSpecificationType struct {
	ID          IdentifierType                        `json:",omitempty" xml:",omitempty"`
	Version     VersionType                           `json:",omitempty" xml:",omitempty"`
	Description []DescriptionType                     `json:",omitempty" xml:",omitempty"`
	Node        []WorkflowSpecificationNodeType       `json:",omitempty" xml:",omitempty"`
	Connection  []WorkflowSpecificationConnectionType `json:",omitempty" xml:",omitempty"`
}

// WorkflowSpecificationNodeType was auto-generated. Do not change.
type WorkflowSpecificationNodeType struct {
	ID                    IdentifierType                      `json:",omitempty" xml:",omitempty"`
	Description           []DescriptionType                   `json:",omitempty" xml:",omitempty"`
	NodeType              IdentifierType                      `json:",omitempty" xml:",omitempty"`
	WorkflowSpecification WorkflowSpecificationType           `json:",omitempty" xml:",omitempty"`
	Property              []WorkflowSpecificationPropertyType `json:",omitempty" xml:",omitempty"`
}

// WorkflowSpecificationPropertyType was auto-generated. Do not change.
type WorkflowSpecificationPropertyType struct {
	ID           IdentifierType                      `json:",omitempty" xml:",omitempty"`
	Description  []DescriptionType                   `json:",omitempty" xml:",omitempty"`
	PropertyType IdentifierType                      `json:",omitempty" xml:",omitempty"`
	Value        []ValueType                         `json:",omitempty" xml:",omitempty"`
	Property     []WorkflowSpecificationPropertyType `json:",omitempty" xml:",omitempty"`
}

// WorkflowSpecificationConnectionType was auto-generated. Do not change.
type WorkflowSpecificationConnectionType struct {
	ID             IdentifierType                      `json:",omitempty" xml:",omitempty"`
	Description    []DescriptionType                   `json:",omitempty" xml:",omitempty"`
	ConnectionType IdentifierType                      `json:",omitempty" xml:",omitempty"`
	FromNodeID     []DescriptionType                   `json:",omitempty" xml:",omitempty"`
	ToNodeID       []DescriptionType                   `json:",omitempty" xml:",omitempty"`
	Property       []WorkflowSpecificationPropertyType `json:",omitempty" xml:",omitempty"`
}

// WorkflowSpecificationTypeType was auto-generated. Do not change.
type WorkflowSpecificationTypeType struct {
	ID             IdentifierType                            `json:",omitempty" xml:",omitempty"`
	Version        VersionType                               `json:",omitempty" xml:",omitempty"`
	Description    []DescriptionType                         `json:",omitempty" xml:",omitempty"`
	NodeType       []WorkflowSpecificationNodeTypeType       `json:",omitempty" xml:",omitempty"`
	ConnectionType []WorkflowSpecificationConnectionTypeType `json:",omitempty" xml:",omitempty"`
}

// WorkflowSpecificationNodeTypeType was auto-generated. Do not change.
type WorkflowSpecificationNodeTypeType struct {
	ID          IdentifierType                      `json:",omitempty" xml:",omitempty"`
	Description []DescriptionType                   `json:",omitempty" xml:",omitempty"`
	Property    []WorkflowSpecificationPropertyType `json:",omitempty" xml:",omitempty"`
}

// WorkflowSpecificationConnectionTypeType was auto-generated. Do not change.
type WorkflowSpecificationConnectionTypeType struct {
	ID          IdentifierType                      `json:",omitempty" xml:",omitempty"`
	Description []DescriptionType                   `json:",omitempty" xml:",omitempty"`
	Property    []WorkflowSpecificationPropertyType `json:",omitempty" xml:",omitempty"`
}

// GetWorkflowSpecificationInformationType was auto-generated. Do not change.
type GetWorkflowSpecificationInformationType struct {
	ApplicationArea TransApplicationAreaType                        `json:",omitempty" xml:",omitempty"`
	DataArea        GetWorkflowSpecificationInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                          `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                          `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// GetWorkflowSpecificationInformationTypeDataArea was auto-generated. Do not change.
type GetWorkflowSpecificationInformationTypeDataArea struct {
	Get                              []string                               `json:",omitempty" xml:",omitempty"`
	WorkflowSpecificationInformation []WorkflowSpecificationInformationType `json:",omitempty" xml:",omitempty"`
}

// ShowWorkflowSpecificationInformationType was auto-generated. Do not change.
type ShowWorkflowSpecificationInformationType struct {
	ApplicationArea TransApplicationAreaType                         `json:",omitempty" xml:",omitempty"`
	DataArea        ShowWorkflowSpecificationInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                           `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                           `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ShowWorkflowSpecificationInformationTypeDataArea was auto-generated. Do not change.
type ShowWorkflowSpecificationInformationTypeDataArea struct {
	Show                             TransShowType                          `json:",omitempty" xml:",omitempty"`
	WorkflowSpecificationInformation []WorkflowSpecificationInformationType `json:",omitempty" xml:",omitempty"`
}

// ProcessWorkflowSpecificationInformationType was auto-generated. Do not change.
type ProcessWorkflowSpecificationInformationType struct {
	ApplicationArea TransApplicationAreaType                            `json:",omitempty" xml:",omitempty"`
	DataArea        ProcessWorkflowSpecificationInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                              `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                              `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ProcessWorkflowSpecificationInformationTypeDataArea was auto-generated. Do not change.
type ProcessWorkflowSpecificationInformationTypeDataArea struct {
	Process                          TransProcessType                       `json:",omitempty" xml:",omitempty"`
	WorkflowSpecificationInformation []WorkflowSpecificationInformationType `json:",omitempty" xml:",omitempty"`
}

// AcknowledgeWorkflowSpecificationInformationType was auto-generated. Do not change.
type AcknowledgeWorkflowSpecificationInformationType struct {
	ApplicationArea TransApplicationAreaType                                `json:",omitempty" xml:",omitempty"`
	DataArea        AcknowledgeWorkflowSpecificationInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                                  `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                                  `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// AcknowledgeWorkflowSpecificationInformationTypeDataArea was auto-generated. Do not change.
type AcknowledgeWorkflowSpecificationInformationTypeDataArea struct {
	Acknowledge                      TransAcknowledgeType                   `json:",omitempty" xml:",omitempty"`
	WorkflowSpecificationInformation []WorkflowSpecificationInformationType `json:",omitempty" xml:",omitempty"`
}

// ChangeWorkflowSpecificationInformationType was auto-generated. Do not change.
type ChangeWorkflowSpecificationInformationType struct {
	ApplicationArea TransApplicationAreaType                           `json:",omitempty" xml:",omitempty"`
	DataArea        ChangeWorkflowSpecificationInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                             `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                             `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ChangeWorkflowSpecificationInformationTypeDataArea was auto-generated. Do not change.
type ChangeWorkflowSpecificationInformationTypeDataArea struct {
	Change                           TransChangeType                        `json:",omitempty" xml:",omitempty"`
	WorkflowSpecificationInformation []WorkflowSpecificationInformationType `json:",omitempty" xml:",omitempty"`
}

// RespondWorkflowSpecificationInformationType was auto-generated. Do not change.
type RespondWorkflowSpecificationInformationType struct {
	ApplicationArea TransApplicationAreaType                            `json:",omitempty" xml:",omitempty"`
	DataArea        RespondWorkflowSpecificationInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                              `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                              `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// RespondWorkflowSpecificationInformationTypeDataArea was auto-generated. Do not change.
type RespondWorkflowSpecificationInformationTypeDataArea struct {
	Respond                          TransRespondType                       `json:",omitempty" xml:",omitempty"`
	WorkflowSpecificationInformation []WorkflowSpecificationInformationType `json:",omitempty" xml:",omitempty"`
}

// CancelWorkflowSpecificationInformationType was auto-generated. Do not change.
type CancelWorkflowSpecificationInformationType struct {
	ApplicationArea TransApplicationAreaType                           `json:",omitempty" xml:",omitempty"`
	DataArea        CancelWorkflowSpecificationInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                             `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                             `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// CancelWorkflowSpecificationInformationTypeDataArea was auto-generated. Do not change.
type CancelWorkflowSpecificationInformationTypeDataArea struct {
	Cancel                           []TransActionCriteriaType              `json:",omitempty" xml:",omitempty"`
	WorkflowSpecificationInformation []WorkflowSpecificationInformationType `json:",omitempty" xml:",omitempty"`
}

// SyncWorkflowSpecificationInformationType was auto-generated. Do not change.
type SyncWorkflowSpecificationInformationType struct {
	ApplicationArea TransApplicationAreaType                         `json:",omitempty" xml:",omitempty"`
	DataArea        SyncWorkflowSpecificationInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                           `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                           `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// SyncWorkflowSpecificationInformationTypeDataArea was auto-generated. Do not change.
type SyncWorkflowSpecificationInformationTypeDataArea struct {
	Sync                             []TransActionCriteriaType              `json:",omitempty" xml:",omitempty"`
	WorkflowSpecificationInformation []WorkflowSpecificationInformationType `json:",omitempty" xml:",omitempty"`
}

// GetWorkflowSpecificationType was auto-generated. Do not change.
type GetWorkflowSpecificationType struct {
	ApplicationArea TransApplicationAreaType             `json:",omitempty" xml:",omitempty"`
	DataArea        GetWorkflowSpecificationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                               `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                               `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// GetWorkflowSpecificationTypeDataArea was auto-generated. Do not change.
type GetWorkflowSpecificationTypeDataArea struct {
	Get                   []string                    `json:",omitempty" xml:",omitempty"`
	WorkflowSpecification []WorkflowSpecificationType `json:",omitempty" xml:",omitempty"`
}

// ShowWorkflowSpecificationType was auto-generated. Do not change.
type ShowWorkflowSpecificationType struct {
	ApplicationArea TransApplicationAreaType              `json:",omitempty" xml:",omitempty"`
	DataArea        ShowWorkflowSpecificationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ShowWorkflowSpecificationTypeDataArea was auto-generated. Do not change.
type ShowWorkflowSpecificationTypeDataArea struct {
	Show                  TransShowType               `json:",omitempty" xml:",omitempty"`
	WorkflowSpecification []WorkflowSpecificationType `json:",omitempty" xml:",omitempty"`
}

// ProcessWorkflowSpecificationType was auto-generated. Do not change.
type ProcessWorkflowSpecificationType struct {
	ApplicationArea TransApplicationAreaType                 `json:",omitempty" xml:",omitempty"`
	DataArea        ProcessWorkflowSpecificationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                   `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                   `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ProcessWorkflowSpecificationTypeDataArea was auto-generated. Do not change.
type ProcessWorkflowSpecificationTypeDataArea struct {
	Process               TransProcessType            `json:",omitempty" xml:",omitempty"`
	WorkflowSpecification []WorkflowSpecificationType `json:",omitempty" xml:",omitempty"`
}

// AcknowledgeWorkflowSpecificationType was auto-generated. Do not change.
type AcknowledgeWorkflowSpecificationType struct {
	ApplicationArea TransApplicationAreaType                     `json:",omitempty" xml:",omitempty"`
	DataArea        AcknowledgeWorkflowSpecificationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                       `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                       `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// AcknowledgeWorkflowSpecificationTypeDataArea was auto-generated. Do not change.
type AcknowledgeWorkflowSpecificationTypeDataArea struct {
	Acknowledge           TransAcknowledgeType        `json:",omitempty" xml:",omitempty"`
	WorkflowSpecification []WorkflowSpecificationType `json:",omitempty" xml:",omitempty"`
}

// ChangeWorkflowSpecificationType was auto-generated. Do not change.
type ChangeWorkflowSpecificationType struct {
	ApplicationArea TransApplicationAreaType                `json:",omitempty" xml:",omitempty"`
	DataArea        ChangeWorkflowSpecificationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                  `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                  `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ChangeWorkflowSpecificationTypeDataArea was auto-generated. Do not change.
type ChangeWorkflowSpecificationTypeDataArea struct {
	Change                TransChangeType             `json:",omitempty" xml:",omitempty"`
	WorkflowSpecification []WorkflowSpecificationType `json:",omitempty" xml:",omitempty"`
}

// RespondWorkflowSpecificationType was auto-generated. Do not change.
type RespondWorkflowSpecificationType struct {
	ApplicationArea TransApplicationAreaType                 `json:",omitempty" xml:",omitempty"`
	DataArea        RespondWorkflowSpecificationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                   `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                   `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// RespondWorkflowSpecificationTypeDataArea was auto-generated. Do not change.
type RespondWorkflowSpecificationTypeDataArea struct {
	Respond               TransRespondType            `json:",omitempty" xml:",omitempty"`
	WorkflowSpecification []WorkflowSpecificationType `json:",omitempty" xml:",omitempty"`
}

// CancelWorkflowSpecificationType was auto-generated. Do not change.
type CancelWorkflowSpecificationType struct {
	ApplicationArea TransApplicationAreaType                `json:",omitempty" xml:",omitempty"`
	DataArea        CancelWorkflowSpecificationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                  `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                  `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// CancelWorkflowSpecificationTypeDataArea was auto-generated. Do not change.
type CancelWorkflowSpecificationTypeDataArea struct {
	Cancel                []TransActionCriteriaType   `json:",omitempty" xml:",omitempty"`
	WorkflowSpecification []WorkflowSpecificationType `json:",omitempty" xml:",omitempty"`
}

// SyncWorkflowSpecificationType was auto-generated. Do not change.
type SyncWorkflowSpecificationType struct {
	ApplicationArea TransApplicationAreaType              `json:",omitempty" xml:",omitempty"`
	DataArea        SyncWorkflowSpecificationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// SyncWorkflowSpecificationTypeDataArea was auto-generated. Do not change.
type SyncWorkflowSpecificationTypeDataArea struct {
	Sync                  []TransActionCriteriaType   `json:",omitempty" xml:",omitempty"`
	WorkflowSpecification []WorkflowSpecificationType `json:",omitempty" xml:",omitempty"`
}

// GetWorkflowSpecificationTypeType was auto-generated. Do not change.
type GetWorkflowSpecificationTypeType struct {
	ApplicationArea TransApplicationAreaType                 `json:",omitempty" xml:",omitempty"`
	DataArea        GetWorkflowSpecificationTypeTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                   `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                   `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// GetWorkflowSpecificationTypeTypeDataArea was auto-generated. Do not change.
type GetWorkflowSpecificationTypeTypeDataArea struct {
	Get                       []string                        `json:",omitempty" xml:",omitempty"`
	WorkflowSpecificationType []WorkflowSpecificationTypeType `json:",omitempty" xml:",omitempty"`
}

// ShowWorkflowSpecificationTypeType was auto-generated. Do not change.
type ShowWorkflowSpecificationTypeType struct {
	ApplicationArea TransApplicationAreaType                  `json:",omitempty" xml:",omitempty"`
	DataArea        ShowWorkflowSpecificationTypeTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                    `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                    `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ShowWorkflowSpecificationTypeTypeDataArea was auto-generated. Do not change.
type ShowWorkflowSpecificationTypeTypeDataArea struct {
	Show                      TransShowType                   `json:",omitempty" xml:",omitempty"`
	WorkflowSpecificationType []WorkflowSpecificationTypeType `json:",omitempty" xml:",omitempty"`
}

// ProcessWorkflowSpecificationTypeType was auto-generated. Do not change.
type ProcessWorkflowSpecificationTypeType struct {
	ApplicationArea TransApplicationAreaType                     `json:",omitempty" xml:",omitempty"`
	DataArea        ProcessWorkflowSpecificationTypeTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                       `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                       `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ProcessWorkflowSpecificationTypeTypeDataArea was auto-generated. Do not change.
type ProcessWorkflowSpecificationTypeTypeDataArea struct {
	Process                   TransProcessType                `json:",omitempty" xml:",omitempty"`
	WorkflowSpecificationType []WorkflowSpecificationTypeType `json:",omitempty" xml:",omitempty"`
}

// AcknowledgeWorkflowSpecificationTypeType was auto-generated. Do not change.
type AcknowledgeWorkflowSpecificationTypeType struct {
	ApplicationArea TransApplicationAreaType                         `json:",omitempty" xml:",omitempty"`
	DataArea        AcknowledgeWorkflowSpecificationTypeTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                           `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                           `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// AcknowledgeWorkflowSpecificationTypeTypeDataArea was auto-generated. Do not change.
type AcknowledgeWorkflowSpecificationTypeTypeDataArea struct {
	Acknowledge               TransAcknowledgeType            `json:",omitempty" xml:",omitempty"`
	WorkflowSpecificationType []WorkflowSpecificationTypeType `json:",omitempty" xml:",omitempty"`
}

// ChangeWorkflowSpecificationTypeType was auto-generated. Do not change.
type ChangeWorkflowSpecificationTypeType struct {
	ApplicationArea TransApplicationAreaType                    `json:",omitempty" xml:",omitempty"`
	DataArea        ChangeWorkflowSpecificationTypeTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                      `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                      `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ChangeWorkflowSpecificationTypeTypeDataArea was auto-generated. Do not change.
type ChangeWorkflowSpecificationTypeTypeDataArea struct {
	Change                    TransChangeType                 `json:",omitempty" xml:",omitempty"`
	WorkflowSpecificationType []WorkflowSpecificationTypeType `json:",omitempty" xml:",omitempty"`
}

// RespondWorkflowSpecificationTypeType was auto-generated. Do not change.
type RespondWorkflowSpecificationTypeType struct {
	ApplicationArea TransApplicationAreaType                     `json:",omitempty" xml:",omitempty"`
	DataArea        RespondWorkflowSpecificationTypeTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                       `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                       `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// RespondWorkflowSpecificationTypeTypeDataArea was auto-generated. Do not change.
type RespondWorkflowSpecificationTypeTypeDataArea struct {
	Respond                   TransRespondType                `json:",omitempty" xml:",omitempty"`
	WorkflowSpecificationType []WorkflowSpecificationTypeType `json:",omitempty" xml:",omitempty"`
}

// CancelWorkflowSpecificationTypeType was auto-generated. Do not change.
type CancelWorkflowSpecificationTypeType struct {
	ApplicationArea TransApplicationAreaType                    `json:",omitempty" xml:",omitempty"`
	DataArea        CancelWorkflowSpecificationTypeTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                      `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                      `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// CancelWorkflowSpecificationTypeTypeDataArea was auto-generated. Do not change.
type CancelWorkflowSpecificationTypeTypeDataArea struct {
	Cancel                    []TransActionCriteriaType       `json:",omitempty" xml:",omitempty"`
	WorkflowSpecificationType []WorkflowSpecificationTypeType `json:",omitempty" xml:",omitempty"`
}

// SyncWorkflowSpecificationTypeType was auto-generated. Do not change.
type SyncWorkflowSpecificationTypeType struct {
	ApplicationArea TransApplicationAreaType                  `json:",omitempty" xml:",omitempty"`
	DataArea        SyncWorkflowSpecificationTypeTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                    `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                    `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// SyncWorkflowSpecificationTypeTypeDataArea was auto-generated. Do not change.
type SyncWorkflowSpecificationTypeTypeDataArea struct {
	Sync                      []TransActionCriteriaType       `json:",omitempty" xml:",omitempty"`
	WorkflowSpecificationType []WorkflowSpecificationTypeType `json:",omitempty" xml:",omitempty"`
}

// WorkDefinitionInformationType was auto-generated. Do not change.
type WorkDefinitionInformationType struct {
	ID             IdentifierType      `json:",omitempty" xml:",omitempty"`
	Description    []DescriptionType   `json:",omitempty" xml:",omitempty"`
	HierarchyScope HierarchyScopeType  `json:",omitempty" xml:",omitempty"`
	PublishedDate  PublishedDateType   `json:",omitempty" xml:",omitempty"`
	WorkMaster     []WorkMasterType    `json:",omitempty" xml:",omitempty"`
	WorkDirective  []WorkDirectiveType `json:",omitempty" xml:",omitempty"`
}

// WorkMasterType was auto-generated. Do not change.
type WorkMasterType struct {
	ID                         IdentifierType                     `json:",omitempty" xml:",omitempty"`
	Version                    VersionType                        `json:",omitempty" xml:",omitempty"`
	Description                []DescriptionType                  `json:",omitempty" xml:",omitempty"`
	HierarchyScope             HierarchyScopeType                 `json:",omitempty" xml:",omitempty"`
	WorkType                   WorkTypeType                       `json:",omitempty" xml:",omitempty"`
	Duration                   string                             `json:",omitempty" xml:",omitempty"`
	PublishedDate              PublishedDateType                  `json:",omitempty" xml:",omitempty"`
	OperationsDefinitionID     OperationsDefinitionIDType         `json:",omitempty" xml:",omitempty"`
	Parameter                  []ParameterType                    `json:",omitempty" xml:",omitempty"`
	PersonnelSpecification     []OpPersonnelSpecificationType     `json:",omitempty" xml:",omitempty"`
	EquipmentSpecification     []OpEquipmentSpecificationType     `json:",omitempty" xml:",omitempty"`
	PhysicalAssetSpecification []OpPhysicalAssetSpecificationType `json:",omitempty" xml:",omitempty"`
	MaterialSpecification      []OpMaterialSpecificationType      `json:",omitempty" xml:",omitempty"`
	WorkflowSpecification      []WorkflowSpecificationType        `json:",omitempty" xml:",omitempty"`
	WorkMaster                 []WorkMasterType                   `json:",omitempty" xml:",omitempty"`
}

// WorkDirectiveType was auto-generated. Do not change.
type WorkDirectiveType struct {
	ID                         IdentifierType                     `json:",omitempty" xml:",omitempty"`
	Version                    VersionType                        `json:",omitempty" xml:",omitempty"`
	Description                []DescriptionType                  `json:",omitempty" xml:",omitempty"`
	HierarchyScope             HierarchyScopeType                 `json:",omitempty" xml:",omitempty"`
	WorkType                   WorkTypeType                       `json:",omitempty" xml:",omitempty"`
	Duration                   string                             `json:",omitempty" xml:",omitempty"`
	PublishedDate              PublishedDateType                  `json:",omitempty" xml:",omitempty"`
	OperationsDefinitionID     OperationsDefinitionIDType         `json:",omitempty" xml:",omitempty"`
	Parameter                  []ParameterType                    `json:",omitempty" xml:",omitempty"`
	PersonnelSpecification     []OpPersonnelSpecificationType     `json:",omitempty" xml:",omitempty"`
	EquipmentSpecification     []OpEquipmentSpecificationType     `json:",omitempty" xml:",omitempty"`
	PhysicalAssetSpecification []OpPhysicalAssetSpecificationType `json:",omitempty" xml:",omitempty"`
	MaterialSpecification      []OpMaterialSpecificationType      `json:",omitempty" xml:",omitempty"`
	WorkflowSpecification      []WorkflowSpecificationType        `json:",omitempty" xml:",omitempty"`
	WorkMasterID               IdentifierType                     `json:",omitempty" xml:",omitempty"`
	WorkDirective              []WorkDirectiveType                `json:",omitempty" xml:",omitempty"`
}

// GetWorkDefinitionInformationType was auto-generated. Do not change.
type GetWorkDefinitionInformationType struct {
	ApplicationArea TransApplicationAreaType                 `json:",omitempty" xml:",omitempty"`
	DataArea        GetWorkDefinitionInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                   `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                   `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// GetWorkDefinitionInformationTypeDataArea was auto-generated. Do not change.
type GetWorkDefinitionInformationTypeDataArea struct {
	Get                       []string                        `json:",omitempty" xml:",omitempty"`
	WorkDefinitionInformation []WorkDefinitionInformationType `json:",omitempty" xml:",omitempty"`
}

// ShowWorkDefinitionInformationType was auto-generated. Do not change.
type ShowWorkDefinitionInformationType struct {
	ApplicationArea TransApplicationAreaType                  `json:",omitempty" xml:",omitempty"`
	DataArea        ShowWorkDefinitionInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                    `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                    `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ShowWorkDefinitionInformationTypeDataArea was auto-generated. Do not change.
type ShowWorkDefinitionInformationTypeDataArea struct {
	Show                      TransShowType                   `json:",omitempty" xml:",omitempty"`
	WorkDefinitionInformation []WorkDefinitionInformationType `json:",omitempty" xml:",omitempty"`
}

// ProcessWorkDefinitionInformationType was auto-generated. Do not change.
type ProcessWorkDefinitionInformationType struct {
	ApplicationArea TransApplicationAreaType                     `json:",omitempty" xml:",omitempty"`
	DataArea        ProcessWorkDefinitionInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                       `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                       `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ProcessWorkDefinitionInformationTypeDataArea was auto-generated. Do not change.
type ProcessWorkDefinitionInformationTypeDataArea struct {
	Process                   TransProcessType                `json:",omitempty" xml:",omitempty"`
	WorkDefinitionInformation []WorkDefinitionInformationType `json:",omitempty" xml:",omitempty"`
}

// AcknowledgeWorkDefinitionInformationType was auto-generated. Do not change.
type AcknowledgeWorkDefinitionInformationType struct {
	ApplicationArea TransApplicationAreaType                         `json:",omitempty" xml:",omitempty"`
	DataArea        AcknowledgeWorkDefinitionInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                           `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                           `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// AcknowledgeWorkDefinitionInformationTypeDataArea was auto-generated. Do not change.
type AcknowledgeWorkDefinitionInformationTypeDataArea struct {
	Acknowledge               TransAcknowledgeType            `json:",omitempty" xml:",omitempty"`
	WorkDefinitionInformation []WorkDefinitionInformationType `json:",omitempty" xml:",omitempty"`
}

// ChangeWorkDefinitionInformationType was auto-generated. Do not change.
type ChangeWorkDefinitionInformationType struct {
	ApplicationArea TransApplicationAreaType                    `json:",omitempty" xml:",omitempty"`
	DataArea        ChangeWorkDefinitionInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                      `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                      `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ChangeWorkDefinitionInformationTypeDataArea was auto-generated. Do not change.
type ChangeWorkDefinitionInformationTypeDataArea struct {
	Change                    TransChangeType                 `json:",omitempty" xml:",omitempty"`
	WorkDefinitionInformation []WorkDefinitionInformationType `json:",omitempty" xml:",omitempty"`
}

// RespondWorkDefinitionInformationType was auto-generated. Do not change.
type RespondWorkDefinitionInformationType struct {
	ApplicationArea TransApplicationAreaType                     `json:",omitempty" xml:",omitempty"`
	DataArea        RespondWorkDefinitionInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                       `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                       `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// RespondWorkDefinitionInformationTypeDataArea was auto-generated. Do not change.
type RespondWorkDefinitionInformationTypeDataArea struct {
	Respond                   TransRespondType                `json:",omitempty" xml:",omitempty"`
	WorkDefinitionInformation []WorkDefinitionInformationType `json:",omitempty" xml:",omitempty"`
}

// CancelWorkDefinitionInformationType was auto-generated. Do not change.
type CancelWorkDefinitionInformationType struct {
	ApplicationArea TransApplicationAreaType                    `json:",omitempty" xml:",omitempty"`
	DataArea        CancelWorkDefinitionInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                      `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                      `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// CancelWorkDefinitionInformationTypeDataArea was auto-generated. Do not change.
type CancelWorkDefinitionInformationTypeDataArea struct {
	Cancel                    []TransActionCriteriaType       `json:",omitempty" xml:",omitempty"`
	WorkDefinitionInformation []WorkDefinitionInformationType `json:",omitempty" xml:",omitempty"`
}

// SyncWorkDefinitionInformationType was auto-generated. Do not change.
type SyncWorkDefinitionInformationType struct {
	ApplicationArea TransApplicationAreaType                  `json:",omitempty" xml:",omitempty"`
	DataArea        SyncWorkDefinitionInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                    `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                    `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// SyncWorkDefinitionInformationTypeDataArea was auto-generated. Do not change.
type SyncWorkDefinitionInformationTypeDataArea struct {
	Sync                      []TransActionCriteriaType       `json:",omitempty" xml:",omitempty"`
	WorkDefinitionInformation []WorkDefinitionInformationType `json:",omitempty" xml:",omitempty"`
}

// GetWorkMasterType was auto-generated. Do not change.
type GetWorkMasterType struct {
	ApplicationArea TransApplicationAreaType  `json:",omitempty" xml:",omitempty"`
	DataArea        GetWorkMasterTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                    `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                    `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// GetWorkMasterTypeDataArea was auto-generated. Do not change.
type GetWorkMasterTypeDataArea struct {
	Get        []string         `json:",omitempty" xml:",omitempty"`
	WorkMaster []WorkMasterType `json:",omitempty" xml:",omitempty"`
}

// ShowWorkMasterType was auto-generated. Do not change.
type ShowWorkMasterType struct {
	ApplicationArea TransApplicationAreaType   `json:",omitempty" xml:",omitempty"`
	DataArea        ShowWorkMasterTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                     `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                     `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ShowWorkMasterTypeDataArea was auto-generated. Do not change.
type ShowWorkMasterTypeDataArea struct {
	Show       TransShowType    `json:",omitempty" xml:",omitempty"`
	WorkMaster []WorkMasterType `json:",omitempty" xml:",omitempty"`
}

// ProcessWorkMasterType was auto-generated. Do not change.
type ProcessWorkMasterType struct {
	ApplicationArea TransApplicationAreaType      `json:",omitempty" xml:",omitempty"`
	DataArea        ProcessWorkMasterTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                        `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                        `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ProcessWorkMasterTypeDataArea was auto-generated. Do not change.
type ProcessWorkMasterTypeDataArea struct {
	Process    TransProcessType `json:",omitempty" xml:",omitempty"`
	WorkMaster []WorkMasterType `json:",omitempty" xml:",omitempty"`
}

// AcknowledgeWorkMasterType was auto-generated. Do not change.
type AcknowledgeWorkMasterType struct {
	ApplicationArea TransApplicationAreaType          `json:",omitempty" xml:",omitempty"`
	DataArea        AcknowledgeWorkMasterTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                            `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                            `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// AcknowledgeWorkMasterTypeDataArea was auto-generated. Do not change.
type AcknowledgeWorkMasterTypeDataArea struct {
	Acknowledge TransAcknowledgeType `json:",omitempty" xml:",omitempty"`
	WorkMaster  []WorkMasterType     `json:",omitempty" xml:",omitempty"`
}

// ChangeWorkMasterType was auto-generated. Do not change.
type ChangeWorkMasterType struct {
	ApplicationArea TransApplicationAreaType     `json:",omitempty" xml:",omitempty"`
	DataArea        ChangeWorkMasterTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                       `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                       `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ChangeWorkMasterTypeDataArea was auto-generated. Do not change.
type ChangeWorkMasterTypeDataArea struct {
	Change     TransChangeType  `json:",omitempty" xml:",omitempty"`
	WorkMaster []WorkMasterType `json:",omitempty" xml:",omitempty"`
}

// RespondWorkMasterType was auto-generated. Do not change.
type RespondWorkMasterType struct {
	ApplicationArea TransApplicationAreaType      `json:",omitempty" xml:",omitempty"`
	DataArea        RespondWorkMasterTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                        `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                        `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// RespondWorkMasterTypeDataArea was auto-generated. Do not change.
type RespondWorkMasterTypeDataArea struct {
	Respond    TransRespondType `json:",omitempty" xml:",omitempty"`
	WorkMaster []WorkMasterType `json:",omitempty" xml:",omitempty"`
}

// CancelWorkMasterType was auto-generated. Do not change.
type CancelWorkMasterType struct {
	ApplicationArea TransApplicationAreaType     `json:",omitempty" xml:",omitempty"`
	DataArea        CancelWorkMasterTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                       `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                       `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// CancelWorkMasterTypeDataArea was auto-generated. Do not change.
type CancelWorkMasterTypeDataArea struct {
	Cancel     []TransActionCriteriaType `json:",omitempty" xml:",omitempty"`
	WorkMaster []WorkMasterType          `json:",omitempty" xml:",omitempty"`
}

// SyncWorkMasterType was auto-generated. Do not change.
type SyncWorkMasterType struct {
	ApplicationArea TransApplicationAreaType   `json:",omitempty" xml:",omitempty"`
	DataArea        SyncWorkMasterTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                     `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                     `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// SyncWorkMasterTypeDataArea was auto-generated. Do not change.
type SyncWorkMasterTypeDataArea struct {
	Sync       []TransActionCriteriaType `json:",omitempty" xml:",omitempty"`
	WorkMaster []WorkMasterType          `json:",omitempty" xml:",omitempty"`
}

// GetWorkDirectiveType was auto-generated. Do not change.
type GetWorkDirectiveType struct {
	ApplicationArea TransApplicationAreaType     `json:",omitempty" xml:",omitempty"`
	DataArea        GetWorkDirectiveTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                       `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                       `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// GetWorkDirectiveTypeDataArea was auto-generated. Do not change.
type GetWorkDirectiveTypeDataArea struct {
	Get           []string            `json:",omitempty" xml:",omitempty"`
	WorkDirective []WorkDirectiveType `json:",omitempty" xml:",omitempty"`
}

// ShowWorkDirectiveType was auto-generated. Do not change.
type ShowWorkDirectiveType struct {
	ApplicationArea TransApplicationAreaType      `json:",omitempty" xml:",omitempty"`
	DataArea        ShowWorkDirectiveTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                        `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                        `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ShowWorkDirectiveTypeDataArea was auto-generated. Do not change.
type ShowWorkDirectiveTypeDataArea struct {
	Show          TransShowType       `json:",omitempty" xml:",omitempty"`
	WorkDirective []WorkDirectiveType `json:",omitempty" xml:",omitempty"`
}

// ProcessWorkDirectiveType was auto-generated. Do not change.
type ProcessWorkDirectiveType struct {
	ApplicationArea TransApplicationAreaType         `json:",omitempty" xml:",omitempty"`
	DataArea        ProcessWorkDirectiveTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                           `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                           `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ProcessWorkDirectiveTypeDataArea was auto-generated. Do not change.
type ProcessWorkDirectiveTypeDataArea struct {
	Process       TransProcessType    `json:",omitempty" xml:",omitempty"`
	WorkDirective []WorkDirectiveType `json:",omitempty" xml:",omitempty"`
}

// AcknowledgeWorkDirectiveType was auto-generated. Do not change.
type AcknowledgeWorkDirectiveType struct {
	ApplicationArea TransApplicationAreaType             `json:",omitempty" xml:",omitempty"`
	DataArea        AcknowledgeWorkDirectiveTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                               `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                               `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// AcknowledgeWorkDirectiveTypeDataArea was auto-generated. Do not change.
type AcknowledgeWorkDirectiveTypeDataArea struct {
	Acknowledge   TransAcknowledgeType `json:",omitempty" xml:",omitempty"`
	WorkDirective []WorkDirectiveType  `json:",omitempty" xml:",omitempty"`
}

// ChangeWorkDirectiveType was auto-generated. Do not change.
type ChangeWorkDirectiveType struct {
	ApplicationArea TransApplicationAreaType        `json:",omitempty" xml:",omitempty"`
	DataArea        ChangeWorkDirectiveTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                          `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                          `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ChangeWorkDirectiveTypeDataArea was auto-generated. Do not change.
type ChangeWorkDirectiveTypeDataArea struct {
	Change        TransChangeType     `json:",omitempty" xml:",omitempty"`
	WorkDirective []WorkDirectiveType `json:",omitempty" xml:",omitempty"`
}

// RespondWorkDirectiveType was auto-generated. Do not change.
type RespondWorkDirectiveType struct {
	ApplicationArea TransApplicationAreaType         `json:",omitempty" xml:",omitempty"`
	DataArea        RespondWorkDirectiveTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                           `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                           `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// RespondWorkDirectiveTypeDataArea was auto-generated. Do not change.
type RespondWorkDirectiveTypeDataArea struct {
	Respond       TransRespondType    `json:",omitempty" xml:",omitempty"`
	WorkDirective []WorkDirectiveType `json:",omitempty" xml:",omitempty"`
}

// CancelWorkDirectiveType was auto-generated. Do not change.
type CancelWorkDirectiveType struct {
	ApplicationArea TransApplicationAreaType        `json:",omitempty" xml:",omitempty"`
	DataArea        CancelWorkDirectiveTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                          `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                          `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// CancelWorkDirectiveTypeDataArea was auto-generated. Do not change.
type CancelWorkDirectiveTypeDataArea struct {
	Cancel        []TransActionCriteriaType `json:",omitempty" xml:",omitempty"`
	WorkDirective []WorkDirectiveType       `json:",omitempty" xml:",omitempty"`
}

// SyncWorkDirectiveType was auto-generated. Do not change.
type SyncWorkDirectiveType struct {
	ApplicationArea TransApplicationAreaType      `json:",omitempty" xml:",omitempty"`
	DataArea        SyncWorkDirectiveTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                        `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                        `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// SyncWorkDirectiveTypeDataArea was auto-generated. Do not change.
type SyncWorkDirectiveTypeDataArea struct {
	Sync          []TransActionCriteriaType `json:",omitempty" xml:",omitempty"`
	WorkDirective []WorkDirectiveType       `json:",omitempty" xml:",omitempty"`
}

// WorkPerformanceType was auto-generated. Do not change.
type WorkPerformanceType struct {
	ID             IdentifierType     `json:",omitempty" xml:",omitempty"`
	Description    []DescriptionType  `json:",omitempty" xml:",omitempty"`
	HierarchyScope HierarchyScopeType `json:",omitempty" xml:",omitempty"`
	WorkType       OperationsTypeType `json:",omitempty" xml:",omitempty"`
	WorkScheduleID WorkScheduleIDType `json:",omitempty" xml:",omitempty"`
	StartTime      StartTimeType      `json:",omitempty" xml:",omitempty"`
	EndTime        EndTimeType        `json:",omitempty" xml:",omitempty"`
	WorkState      ResponseStateType  `json:",omitempty" xml:",omitempty"`
	PublishedDate  PublishedDateType  `json:",omitempty" xml:",omitempty"`
	WorkResponse   []WorkResponseType `json:",omitempty" xml:",omitempty"`
}

// WorkResponseType was auto-generated. Do not change.
type WorkResponseType struct {
	ID             IdentifierType     `json:",omitempty" xml:",omitempty"`
	Description    []DescriptionType  `json:",omitempty" xml:",omitempty"`
	HierarchyScope HierarchyScopeType `json:",omitempty" xml:",omitempty"`
	WorkType       OperationsTypeType `json:",omitempty" xml:",omitempty"`
	WorkRequestID  WorkRequestIDType  `json:",omitempty" xml:",omitempty"`
	StartTime      StartTimeType      `json:",omitempty" xml:",omitempty"`
	EndTime        EndTimeType        `json:",omitempty" xml:",omitempty"`
	ResponseState  ResponseStateType  `json:",omitempty" xml:",omitempty"`
	JobResponse    []JobResponseType  `json:",omitempty" xml:",omitempty"`
}

// JobResponseType was auto-generated. Do not change.
type JobResponseType struct {
	ID                   IdentifierType              `json:",omitempty" xml:",omitempty"`
	Description          []DescriptionType           `json:",omitempty" xml:",omitempty"`
	HierarchyScope       HierarchyScopeType          `json:",omitempty" xml:",omitempty"`
	WorkType             OperationsTypeType          `json:",omitempty" xml:",omitempty"`
	JobOrderID           IdentifierType              `json:",omitempty" xml:",omitempty"`
	WorkDirectiveID      IdentifierType              `json:",omitempty" xml:",omitempty"`
	WorkDirectiveVersion VersionType                 `json:",omitempty" xml:",omitempty"`
	StartTime            ActualStartTimeType         `json:",omitempty" xml:",omitempty"`
	EndTime              ActualEndTimeType           `json:",omitempty" xml:",omitempty"`
	JobState             ResponseStateType           `json:",omitempty" xml:",omitempty"`
	JobResponsetData     []OpSegmentDataType         `json:",omitempty" xml:",omitempty"`
	PersonnelActual      []OpPersonnelActualType     `json:",omitempty" xml:",omitempty"`
	EquipmentActual      []OpEquipmentActualType     `json:",omitempty" xml:",omitempty"`
	PhysicalAssetActual  []OpPhysicalAssetActualType `json:",omitempty" xml:",omitempty"`
	MaterialActual       []OpMaterialActualType      `json:",omitempty" xml:",omitempty"`
	JobResponse          []JobResponseType           `json:",omitempty" xml:",omitempty"`
}

// GetWorkPerformanceType was auto-generated. Do not change.
type GetWorkPerformanceType struct {
	ApplicationArea TransApplicationAreaType       `json:",omitempty" xml:",omitempty"`
	DataArea        GetWorkPerformanceTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                         `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                         `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// GetWorkPerformanceTypeDataArea was auto-generated. Do not change.
type GetWorkPerformanceTypeDataArea struct {
	Get             []string              `json:",omitempty" xml:",omitempty"`
	WorkPerformance []WorkPerformanceType `json:",omitempty" xml:",omitempty"`
}

// ShowWorkPerformanceType was auto-generated. Do not change.
type ShowWorkPerformanceType struct {
	ApplicationArea TransApplicationAreaType        `json:",omitempty" xml:",omitempty"`
	DataArea        ShowWorkPerformanceTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                          `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                          `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ShowWorkPerformanceTypeDataArea was auto-generated. Do not change.
type ShowWorkPerformanceTypeDataArea struct {
	Show            TransShowType         `json:",omitempty" xml:",omitempty"`
	WorkPerformance []WorkPerformanceType `json:",omitempty" xml:",omitempty"`
}

// ProcessWorkPerformanceType was auto-generated. Do not change.
type ProcessWorkPerformanceType struct {
	ApplicationArea TransApplicationAreaType           `json:",omitempty" xml:",omitempty"`
	DataArea        ProcessWorkPerformanceTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                             `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                             `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ProcessWorkPerformanceTypeDataArea was auto-generated. Do not change.
type ProcessWorkPerformanceTypeDataArea struct {
	Process         TransProcessType      `json:",omitempty" xml:",omitempty"`
	WorkPerformance []WorkPerformanceType `json:",omitempty" xml:",omitempty"`
}

// AcknowledgeWorkPerformanceType was auto-generated. Do not change.
type AcknowledgeWorkPerformanceType struct {
	ApplicationArea TransApplicationAreaType               `json:",omitempty" xml:",omitempty"`
	DataArea        AcknowledgeWorkPerformanceTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                 `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                 `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// AcknowledgeWorkPerformanceTypeDataArea was auto-generated. Do not change.
type AcknowledgeWorkPerformanceTypeDataArea struct {
	Acknowledge     TransAcknowledgeType  `json:",omitempty" xml:",omitempty"`
	WorkPerformance []WorkPerformanceType `json:",omitempty" xml:",omitempty"`
}

// ChangeWorkPerformanceType was auto-generated. Do not change.
type ChangeWorkPerformanceType struct {
	ApplicationArea TransApplicationAreaType          `json:",omitempty" xml:",omitempty"`
	DataArea        ChangeWorkPerformanceTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                            `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                            `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ChangeWorkPerformanceTypeDataArea was auto-generated. Do not change.
type ChangeWorkPerformanceTypeDataArea struct {
	Change          TransChangeType       `json:",omitempty" xml:",omitempty"`
	WorkPerformance []WorkPerformanceType `json:",omitempty" xml:",omitempty"`
}

// RespondWorkPerformanceType was auto-generated. Do not change.
type RespondWorkPerformanceType struct {
	ApplicationArea TransApplicationAreaType           `json:",omitempty" xml:",omitempty"`
	DataArea        RespondWorkPerformanceTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                             `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                             `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// RespondWorkPerformanceTypeDataArea was auto-generated. Do not change.
type RespondWorkPerformanceTypeDataArea struct {
	Respond         TransRespondType      `json:",omitempty" xml:",omitempty"`
	WorkPerformance []WorkPerformanceType `json:",omitempty" xml:",omitempty"`
}

// CancelWorkPerformanceType was auto-generated. Do not change.
type CancelWorkPerformanceType struct {
	ApplicationArea TransApplicationAreaType          `json:",omitempty" xml:",omitempty"`
	DataArea        CancelWorkPerformanceTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                            `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                            `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// CancelWorkPerformanceTypeDataArea was auto-generated. Do not change.
type CancelWorkPerformanceTypeDataArea struct {
	Cancel          []TransActionCriteriaType `json:",omitempty" xml:",omitempty"`
	WorkPerformance []WorkPerformanceType     `json:",omitempty" xml:",omitempty"`
}

// SyncWorkPerformanceType was auto-generated. Do not change.
type SyncWorkPerformanceType struct {
	ApplicationArea TransApplicationAreaType        `json:",omitempty" xml:",omitempty"`
	DataArea        SyncWorkPerformanceTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                          `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                          `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// SyncWorkPerformanceTypeDataArea was auto-generated. Do not change.
type SyncWorkPerformanceTypeDataArea struct {
	Sync            []TransActionCriteriaType `json:",omitempty" xml:",omitempty"`
	WorkPerformance []WorkPerformanceType     `json:",omitempty" xml:",omitempty"`
}

// GetJobResponseType was auto-generated. Do not change.
type GetJobResponseType struct {
	ApplicationArea TransApplicationAreaType   `json:",omitempty" xml:",omitempty"`
	DataArea        GetJobResponseTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                     `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                     `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// GetJobResponseTypeDataArea was auto-generated. Do not change.
type GetJobResponseTypeDataArea struct {
	Get         []string          `json:",omitempty" xml:",omitempty"`
	JobResponse []JobResponseType `json:",omitempty" xml:",omitempty"`
}

// ShowJobResponseType was auto-generated. Do not change.
type ShowJobResponseType struct {
	ApplicationArea TransApplicationAreaType    `json:",omitempty" xml:",omitempty"`
	DataArea        ShowJobResponseTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                      `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                      `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ShowJobResponseTypeDataArea was auto-generated. Do not change.
type ShowJobResponseTypeDataArea struct {
	Show        TransShowType     `json:",omitempty" xml:",omitempty"`
	JobResponse []JobResponseType `json:",omitempty" xml:",omitempty"`
}

// ProcessJobResponseType was auto-generated. Do not change.
type ProcessJobResponseType struct {
	ApplicationArea TransApplicationAreaType       `json:",omitempty" xml:",omitempty"`
	DataArea        ProcessJobResponseTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                         `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                         `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ProcessJobResponseTypeDataArea was auto-generated. Do not change.
type ProcessJobResponseTypeDataArea struct {
	Process     TransProcessType  `json:",omitempty" xml:",omitempty"`
	JobResponse []JobResponseType `json:",omitempty" xml:",omitempty"`
}

// AcknowledgeJobResponseType was auto-generated. Do not change.
type AcknowledgeJobResponseType struct {
	ApplicationArea TransApplicationAreaType           `json:",omitempty" xml:",omitempty"`
	DataArea        AcknowledgeJobResponseTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                             `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                             `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// AcknowledgeJobResponseTypeDataArea was auto-generated. Do not change.
type AcknowledgeJobResponseTypeDataArea struct {
	Acknowledge TransAcknowledgeType `json:",omitempty" xml:",omitempty"`
	JobResponse []JobResponseType    `json:",omitempty" xml:",omitempty"`
}

// ChangeJobResponseType was auto-generated. Do not change.
type ChangeJobResponseType struct {
	ApplicationArea TransApplicationAreaType      `json:",omitempty" xml:",omitempty"`
	DataArea        ChangeJobResponseTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                        `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                        `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ChangeJobResponseTypeDataArea was auto-generated. Do not change.
type ChangeJobResponseTypeDataArea struct {
	Change      TransChangeType   `json:",omitempty" xml:",omitempty"`
	JobResponse []JobResponseType `json:",omitempty" xml:",omitempty"`
}

// RespondJobResponseType was auto-generated. Do not change.
type RespondJobResponseType struct {
	ApplicationArea TransApplicationAreaType       `json:",omitempty" xml:",omitempty"`
	DataArea        RespondJobResponseTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                         `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                         `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// RespondJobResponseTypeDataArea was auto-generated. Do not change.
type RespondJobResponseTypeDataArea struct {
	Respond     TransRespondType  `json:",omitempty" xml:",omitempty"`
	JobResponse []JobResponseType `json:",omitempty" xml:",omitempty"`
}

// CancelJobResponseType was auto-generated. Do not change.
type CancelJobResponseType struct {
	ApplicationArea TransApplicationAreaType      `json:",omitempty" xml:",omitempty"`
	DataArea        CancelJobResponseTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                        `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                        `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// CancelJobResponseTypeDataArea was auto-generated. Do not change.
type CancelJobResponseTypeDataArea struct {
	Cancel      []TransActionCriteriaType `json:",omitempty" xml:",omitempty"`
	JobResponse []JobResponseType         `json:",omitempty" xml:",omitempty"`
}

// SyncJobResponseType was auto-generated. Do not change.
type SyncJobResponseType struct {
	ApplicationArea TransApplicationAreaType    `json:",omitempty" xml:",omitempty"`
	DataArea        SyncJobResponseTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                      `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                      `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// SyncJobResponseTypeDataArea was auto-generated. Do not change.
type SyncJobResponseTypeDataArea struct {
	Sync        []TransActionCriteriaType `json:",omitempty" xml:",omitempty"`
	JobResponse []JobResponseType         `json:",omitempty" xml:",omitempty"`
}

// WorkScheduleType was auto-generated. Do not change.
type WorkScheduleType struct {
	ID             IdentifierType     `json:",omitempty" xml:",omitempty"`
	Description    []DescriptionType  `json:",omitempty" xml:",omitempty"`
	HierarchyScope HierarchyScopeType `json:",omitempty" xml:",omitempty"`
	WorkType       OperationsTypeType `json:",omitempty" xml:",omitempty"`
	StartTime      StartTimeType      `json:",omitempty" xml:",omitempty"`
	EndTime        EndTimeType        `json:",omitempty" xml:",omitempty"`
	ScheduleState  RequestStateType   `json:",omitempty" xml:",omitempty"`
	PublishedDate  PublishedDateType  `json:",omitempty" xml:",omitempty"`
	WorkSchedule   []WorkScheduleType `json:",omitempty" xml:",omitempty"`
	WorkRequest    []WorkRequestType  `json:",omitempty" xml:",omitempty"`
}

// WorkRequestType was auto-generated. Do not change.
type WorkRequestType struct {
	ID             IdentifierType     `json:",omitempty" xml:",omitempty"`
	Description    []DescriptionType  `json:",omitempty" xml:",omitempty"`
	HierarchyScope HierarchyScopeType `json:",omitempty" xml:",omitempty"`
	Workype        OperationsTypeType `json:",omitempty" xml:",omitempty"`
	StartTime      StartTimeType      `json:",omitempty" xml:",omitempty"`
	EndTime        EndTimeType        `json:",omitempty" xml:",omitempty"`
	Priority       PriorityType       `json:",omitempty" xml:",omitempty"`
	WorkRequest    []WorkRequestType  `json:",omitempty" xml:",omitempty"`
	JobOrder       []JobOrderType     `json:",omitempty" xml:",omitempty"`
}

// JobOrderType was auto-generated. Do not change.
type JobOrderType struct {
	ID                       IdentifierType                   `json:",omitempty" xml:",omitempty"`
	Description              []DescriptionType                `json:",omitempty" xml:",omitempty"`
	HierarchyScope           HierarchyScopeType               `json:",omitempty" xml:",omitempty"`
	WorkType                 OperationsTypeType               `json:",omitempty" xml:",omitempty"`
	WorkMasterID             IdentifierType                   `json:",omitempty" xml:",omitempty"`
	WorkMasterVersion        VersionType                      `json:",omitempty" xml:",omitempty"`
	StartTime                StartTimeType                    `json:",omitempty" xml:",omitempty"`
	EndTime                  EndTimeType                      `json:",omitempty" xml:",omitempty"`
	Priority                 PriorityType                     `json:",omitempty" xml:",omitempty"`
	Command                  JobOrderCommandType              `json:",omitempty" xml:",omitempty"`
	CommandRule              JobOrderCommandRuleType          `json:",omitempty" xml:",omitempty"`
	DispatchStatus           JobOrderDispatchStatusType       `json:",omitempty" xml:",omitempty"`
	JobOrderParameter        []ParameterType                  `json:",omitempty" xml:",omitempty"`
	PersonnelRequirement     []OpPersonnelRequirementType     `json:",omitempty" xml:",omitempty"`
	EquipmentRequirement     []OpEquipmentRequirementType     `json:",omitempty" xml:",omitempty"`
	PhysicalAssetRequirement []OpPhysicalAssetRequirementType `json:",omitempty" xml:",omitempty"`
	MaterialRequirement      []OpMaterialRequirementType      `json:",omitempty" xml:",omitempty"`
}

// JobListType was auto-generated. Do not change.
type JobListType struct {
	ID             IdentifierType     `json:",omitempty" xml:",omitempty"`
	Description    []DescriptionType  `json:",omitempty" xml:",omitempty"`
	HierarchyScope HierarchyScopeType `json:",omitempty" xml:",omitempty"`
	Workype        OperationsTypeType `json:",omitempty" xml:",omitempty"`
	StartTime      StartTimeType      `json:",omitempty" xml:",omitempty"`
	EndTime        EndTimeType        `json:",omitempty" xml:",omitempty"`
	JobOrder       []JobOrderType     `json:",omitempty" xml:",omitempty"`
}

// GetWorkScheduleType was auto-generated. Do not change.
type GetWorkScheduleType struct {
	ApplicationArea TransApplicationAreaType    `json:",omitempty" xml:",omitempty"`
	DataArea        GetWorkScheduleTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                      `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                      `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// GetWorkScheduleTypeDataArea was auto-generated. Do not change.
type GetWorkScheduleTypeDataArea struct {
	Get          []string           `json:",omitempty" xml:",omitempty"`
	WorkSchedule []WorkScheduleType `json:",omitempty" xml:",omitempty"`
}

// ShowWorkScheduleType was auto-generated. Do not change.
type ShowWorkScheduleType struct {
	ApplicationArea TransApplicationAreaType     `json:",omitempty" xml:",omitempty"`
	DataArea        ShowWorkScheduleTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                       `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                       `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ShowWorkScheduleTypeDataArea was auto-generated. Do not change.
type ShowWorkScheduleTypeDataArea struct {
	Show         TransShowType      `json:",omitempty" xml:",omitempty"`
	WorkSchedule []WorkScheduleType `json:",omitempty" xml:",omitempty"`
}

// ProcessWorkScheduleType was auto-generated. Do not change.
type ProcessWorkScheduleType struct {
	ApplicationArea TransApplicationAreaType        `json:",omitempty" xml:",omitempty"`
	DataArea        ProcessWorkScheduleTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                          `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                          `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ProcessWorkScheduleTypeDataArea was auto-generated. Do not change.
type ProcessWorkScheduleTypeDataArea struct {
	Process      TransProcessType   `json:",omitempty" xml:",omitempty"`
	WorkSchedule []WorkScheduleType `json:",omitempty" xml:",omitempty"`
}

// AcknowledgeWorkScheduleType was auto-generated. Do not change.
type AcknowledgeWorkScheduleType struct {
	ApplicationArea TransApplicationAreaType            `json:",omitempty" xml:",omitempty"`
	DataArea        AcknowledgeWorkScheduleTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                              `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                              `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// AcknowledgeWorkScheduleTypeDataArea was auto-generated. Do not change.
type AcknowledgeWorkScheduleTypeDataArea struct {
	Acknowledge  TransAcknowledgeType `json:",omitempty" xml:",omitempty"`
	WorkSchedule []WorkScheduleType   `json:",omitempty" xml:",omitempty"`
}

// ChangeWorkScheduleType was auto-generated. Do not change.
type ChangeWorkScheduleType struct {
	ApplicationArea TransApplicationAreaType       `json:",omitempty" xml:",omitempty"`
	DataArea        ChangeWorkScheduleTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                         `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                         `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ChangeWorkScheduleTypeDataArea was auto-generated. Do not change.
type ChangeWorkScheduleTypeDataArea struct {
	Change       TransChangeType    `json:",omitempty" xml:",omitempty"`
	WorkSchedule []WorkScheduleType `json:",omitempty" xml:",omitempty"`
}

// RespondWorkScheduleType was auto-generated. Do not change.
type RespondWorkScheduleType struct {
	ApplicationArea TransApplicationAreaType        `json:",omitempty" xml:",omitempty"`
	DataArea        RespondWorkScheduleTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                          `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                          `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// RespondWorkScheduleTypeDataArea was auto-generated. Do not change.
type RespondWorkScheduleTypeDataArea struct {
	Respond      TransRespondType   `json:",omitempty" xml:",omitempty"`
	WorkSchedule []WorkScheduleType `json:",omitempty" xml:",omitempty"`
}

// CancelWorkScheduleType was auto-generated. Do not change.
type CancelWorkScheduleType struct {
	ApplicationArea TransApplicationAreaType       `json:",omitempty" xml:",omitempty"`
	DataArea        CancelWorkScheduleTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                         `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                         `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// CancelWorkScheduleTypeDataArea was auto-generated. Do not change.
type CancelWorkScheduleTypeDataArea struct {
	Cancel       []TransActionCriteriaType `json:",omitempty" xml:",omitempty"`
	WorkSchedule []WorkScheduleType        `json:",omitempty" xml:",omitempty"`
}

// SyncWorkScheduleType was auto-generated. Do not change.
type SyncWorkScheduleType struct {
	ApplicationArea TransApplicationAreaType     `json:",omitempty" xml:",omitempty"`
	DataArea        SyncWorkScheduleTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                       `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                       `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// SyncWorkScheduleTypeDataArea was auto-generated. Do not change.
type SyncWorkScheduleTypeDataArea struct {
	Sync         []TransActionCriteriaType `json:",omitempty" xml:",omitempty"`
	WorkSchedule []WorkScheduleType        `json:",omitempty" xml:",omitempty"`
}

// GetJobListType was auto-generated. Do not change.
type GetJobListType struct {
	ApplicationArea TransApplicationAreaType `json:",omitempty" xml:",omitempty"`
	DataArea        GetJobListTypeDataArea   `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                   `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                   `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// GetJobListTypeDataArea was auto-generated. Do not change.
type GetJobListTypeDataArea struct {
	Get     []string      `json:",omitempty" xml:",omitempty"`
	JobList []JobListType `json:",omitempty" xml:",omitempty"`
}

// ShowJobListType was auto-generated. Do not change.
type ShowJobListType struct {
	ApplicationArea TransApplicationAreaType `json:",omitempty" xml:",omitempty"`
	DataArea        ShowJobListTypeDataArea  `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                   `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                   `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ShowJobListTypeDataArea was auto-generated. Do not change.
type ShowJobListTypeDataArea struct {
	Show    TransShowType `json:",omitempty" xml:",omitempty"`
	JobList []JobListType `json:",omitempty" xml:",omitempty"`
}

// ProcessJobListType was auto-generated. Do not change.
type ProcessJobListType struct {
	ApplicationArea TransApplicationAreaType   `json:",omitempty" xml:",omitempty"`
	DataArea        ProcessJobListTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                     `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                     `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ProcessJobListTypeDataArea was auto-generated. Do not change.
type ProcessJobListTypeDataArea struct {
	Process TransProcessType `json:",omitempty" xml:",omitempty"`
	JobList []JobListType    `json:",omitempty" xml:",omitempty"`
}

// AcknowledgeJobListType was auto-generated. Do not change.
type AcknowledgeJobListType struct {
	ApplicationArea TransApplicationAreaType       `json:",omitempty" xml:",omitempty"`
	DataArea        AcknowledgeJobListTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                         `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                         `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// AcknowledgeJobListTypeDataArea was auto-generated. Do not change.
type AcknowledgeJobListTypeDataArea struct {
	Acknowledge TransAcknowledgeType `json:",omitempty" xml:",omitempty"`
	JobList     []JobListType        `json:",omitempty" xml:",omitempty"`
}

// ChangeJobListType was auto-generated. Do not change.
type ChangeJobListType struct {
	ApplicationArea TransApplicationAreaType  `json:",omitempty" xml:",omitempty"`
	DataArea        ChangeJobListTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                    `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                    `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ChangeJobListTypeDataArea was auto-generated. Do not change.
type ChangeJobListTypeDataArea struct {
	Change  TransChangeType `json:",omitempty" xml:",omitempty"`
	JobList []JobListType   `json:",omitempty" xml:",omitempty"`
}

// RespondJobListType was auto-generated. Do not change.
type RespondJobListType struct {
	ApplicationArea TransApplicationAreaType   `json:",omitempty" xml:",omitempty"`
	DataArea        RespondJobListTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                     `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                     `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// RespondJobListTypeDataArea was auto-generated. Do not change.
type RespondJobListTypeDataArea struct {
	Respond TransRespondType `json:",omitempty" xml:",omitempty"`
	JobList []JobListType    `json:",omitempty" xml:",omitempty"`
}

// CancelJobListType was auto-generated. Do not change.
type CancelJobListType struct {
	ApplicationArea TransApplicationAreaType  `json:",omitempty" xml:",omitempty"`
	DataArea        CancelJobListTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                    `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                    `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// CancelJobListTypeDataArea was auto-generated. Do not change.
type CancelJobListTypeDataArea struct {
	Cancel  []TransActionCriteriaType `json:",omitempty" xml:",omitempty"`
	JobList []JobListType             `json:",omitempty" xml:",omitempty"`
}

// SyncJobListType was auto-generated. Do not change.
type SyncJobListType struct {
	ApplicationArea TransApplicationAreaType `json:",omitempty" xml:",omitempty"`
	DataArea        SyncJobListTypeDataArea  `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                   `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                   `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// SyncJobListTypeDataArea was auto-generated. Do not change.
type SyncJobListTypeDataArea struct {
	Sync    []TransActionCriteriaType `json:",omitempty" xml:",omitempty"`
	JobList []JobListType             `json:",omitempty" xml:",omitempty"`
}

// BatchInformationType was auto-generated. Do not change.
type BatchInformationType struct {
	ListHeader          []ListHeaderType          `json:",omitempty" xml:",omitempty"`
	Description         []DescriptionType         `json:",omitempty" xml:",omitempty"`
	MasterRecipe        []MasterRecipeType        `json:",omitempty" xml:",omitempty"`
	ControlRecipe       []ControlRecipeType       `json:",omitempty" xml:",omitempty"`
	RecipeBuildingBlock []RecipeBuildingBlockType `json:",omitempty" xml:",omitempty"`
	EquipmentElement    []EquipmentElementType    `json:",omitempty" xml:",omitempty"`
	BatchList           []BatchListType           `json:",omitempty" xml:",omitempty"`
	EnumerationSet      []EnumerationSetType      `json:",omitempty" xml:",omitempty"`
}

// ListHeaderType was auto-generated. Do not change.
type ListHeaderType struct {
	ID              IDType                `json:",omitempty" xml:",omitempty"`
	Version         VersionType           `json:",omitempty" xml:",omitempty"`
	Description     []DescriptionType     `json:",omitempty" xml:",omitempty"`
	Origin          OriginType            `json:",omitempty" xml:",omitempty"`
	CreateDate      CreateDateType        `json:",omitempty" xml:",omitempty"`
	ModificationLog []ModificationLogType `json:",omitempty" xml:",omitempty"`
}

// ModificationLogType was auto-generated. Do not change.
type ModificationLogType struct {
	ModifiedDate ModifiedDateType  `json:",omitempty" xml:",omitempty"`
	Description  []DescriptionType `json:",omitempty" xml:",omitempty"`
	Author       AuthorType        `json:",omitempty" xml:",omitempty"`
}

// MasterRecipeType was auto-generated. Do not change.
type MasterRecipeType struct {
	ID                   IDType                          `json:",omitempty" xml:",omitempty"`
	Version              VersionType                     `json:",omitempty" xml:",omitempty"`
	VersionDate          VersionDateType                 `json:",omitempty" xml:",omitempty"`
	Description          []DescriptionType               `json:",omitempty" xml:",omitempty"`
	Header               HeaderType                      `json:",omitempty" xml:",omitempty"`
	EquipmentRequirement []BatchEquipmentRequirementType `json:",omitempty" xml:",omitempty"`
	Formula              FormulaType                     `json:",omitempty" xml:",omitempty"`
	ProcedureLogic       ProcedureLogicType              `json:",omitempty" xml:",omitempty"`
	RecipeElement        []RecipeElementType             `json:",omitempty" xml:",omitempty"`
	OtherInformation     []OtherInformationType          `json:",omitempty" xml:",omitempty"`
}

// HeaderType was auto-generated. Do not change.
type HeaderType struct {
	ModificationLog       []ModificationLogType       `json:",omitempty" xml:",omitempty"`
	ApprovalHistory       []ApprovalHistoryType       `json:",omitempty" xml:",omitempty"`
	EffectiveDate         EffectiveDateType           `json:",omitempty" xml:",omitempty"`
	ExpirationDate        ExpirationDateType          `json:",omitempty" xml:",omitempty"`
	ProductID             ProductIDType               `json:",omitempty" xml:",omitempty"`
	ProductName           ProductNameType             `json:",omitempty" xml:",omitempty"`
	BatchSize             BatchSizeType               `json:",omitempty" xml:",omitempty"`
	ActualProductProduced []ActualProductProducedType `json:",omitempty" xml:",omitempty"`
	Status                BatchStatusType             `json:",omitempty" xml:",omitempty"`
}

// ApprovalHistoryType was auto-generated. Do not change.
type ApprovalHistoryType struct {
	FinalApprovalDate  FinalApprovalDateType    `json:",omitempty" xml:",omitempty"`
	Version            VersionType              `json:",omitempty" xml:",omitempty"`
	Description        []DescriptionType        `json:",omitempty" xml:",omitempty"`
	IndividualApproval []IndividualApprovalType `json:",omitempty" xml:",omitempty"`
}

// IndividualApprovalType was auto-generated. Do not change.
type IndividualApprovalType struct {
	ApprovedBy   ApprovedByType    `json:",omitempty" xml:",omitempty"`
	ApprovalDate ApprovalDateType  `json:",omitempty" xml:",omitempty"`
	Description  []DescriptionType `json:",omitempty" xml:",omitempty"`
}

// BatchSizeType was auto-generated. Do not change.
type BatchSizeType struct {
	Nominal        NominalType        `json:",omitempty" xml:",omitempty"`
	Min            MinType            `json:",omitempty" xml:",omitempty"`
	Max            MaxType            `json:",omitempty" xml:",omitempty"`
	ScaleReference ScaleReferenceType `json:",omitempty" xml:",omitempty"`
	ScaledSize     ScaledSizeType     `json:",omitempty" xml:",omitempty"`
	UnitOfMeasure  UnitOfMeasureType  `json:",omitempty" xml:",omitempty"`
}

// NominalType was auto-generated. Do not change.
type NominalType struct {
	MeasureType
}

// MeasureType was auto-generated. Do not change.
type MeasureType struct {
	UnitCode              string  `json:",omitempty" xml:"unitCode,attr,omitempty"`
	UnitCodeListVersionID string  `json:",omitempty" xml:"unitCodeListVersionID,attr,omitempty"`
	Value                 float64 `json:",omitempty" xml:",chardata"`
}

// MinType was auto-generated. Do not change.
type MinType struct {
	MeasureType
}

// MaxType was auto-generated. Do not change.
type MaxType struct {
	MeasureType
}

// ScaleReferenceType was auto-generated. Do not change.
type ScaleReferenceType struct {
	MeasureType
}

// ScaledSizeType was auto-generated. Do not change.
type ScaledSizeType struct {
	MeasureType
}

// BatchEquipmentRequirementType was auto-generated. Do not change.
type BatchEquipmentRequirementType struct {
	ID          IDType           `json:",omitempty" xml:",omitempty"`
	Constraint  []ConstraintType `json:",omitempty" xml:",omitempty"`
	Description DescriptionType  `json:",omitempty" xml:",omitempty"`
}

// ConstraintType was auto-generated. Do not change.
type ConstraintType struct {
	ID        IDType        `json:",omitempty" xml:",omitempty"`
	Condition ConditionType `json:",omitempty" xml:",omitempty"`
}

// FormulaType was auto-generated. Do not change.
type FormulaType struct {
	Parameter []BatchParameterType `json:",omitempty" xml:",omitempty"`
}

// BatchParameterType was auto-generated. Do not change.
type BatchParameterType struct {
	ID               IDType                 `json:",omitempty" xml:",omitempty"`
	Description      DescriptionType        `json:",omitempty" xml:",omitempty"`
	ParameterType    ParameterTypeType      `json:",omitempty" xml:",omitempty"`
	ParameterSubType []ParameterSubTypeType `json:",omitempty" xml:",omitempty"`
	Value            []BatchValueType       `json:",omitempty" xml:",omitempty"`
	Scaled           ScaledType             `json:",omitempty" xml:",omitempty"`
	ScaleReference   ScaleReferenceType     `json:",omitempty" xml:",omitempty"`
	Parameter        []BatchParameterType   `json:",omitempty" xml:",omitempty"`
}

// BatchValueType was auto-generated. Do not change.
type BatchValueType struct {
	ValueString        []ValueStringType      `json:",omitempty" xml:",omitempty"`
	DataInterpretation DataInterpretationType `json:",omitempty" xml:",omitempty"`
	DataType           DataTypeType           `json:",omitempty" xml:",omitempty"`
	UnitOfMeasure      UnitOfMeasureType      `json:",omitempty" xml:",omitempty"`
	EnumerationSetID   []EnumerationSetIDType `json:",omitempty" xml:",omitempty"`
	Text               []string               `json:",omitempty" xml:",chardata"`
}

// ProcedureLogicType was auto-generated. Do not change.
type ProcedureLogicType struct {
	Link       []LinkType       `json:",omitempty" xml:",omitempty"`
	Step       []StepType       `json:",omitempty" xml:",omitempty"`
	Transition []TransitionType `json:",omitempty" xml:",omitempty"`
}

// LinkType was auto-generated. Do not change.
type LinkType struct {
	ID              IDType              `json:",omitempty" xml:",omitempty"`
	FromID          []FromIDType        `json:",omitempty" xml:",omitempty"`
	ToID            []ToIDType          `json:",omitempty" xml:",omitempty"`
	LinkType1       LinkTypeType        `json:",omitempty" xml:",omitempty"`
	Depiction       DepictionType       `json:",omitempty" xml:",omitempty"`
	EvaluationOrder EvaluationOrderType `json:",omitempty" xml:",omitempty"`
	Description     []DescriptionType   `json:",omitempty" xml:",omitempty"`
}

// FromIDType was auto-generated. Do not change.
type FromIDType struct {
	FromIDValue string       `json:",omitempty" xml:",omitempty"`
	FromType    FromTypeType `json:",omitempty" xml:",omitempty"`
	IDScope     IDScopeType  `json:",omitempty" xml:",omitempty"`
}

// ToIDType was auto-generated. Do not change.
type ToIDType struct {
	ToIDValue string      `json:",omitempty" xml:",omitempty"`
	ToType    ToTypeType  `json:",omitempty" xml:",omitempty"`
	IDScope   IDScopeType `json:",omitempty" xml:",omitempty"`
}

// StepType was auto-generated. Do not change.
type StepType struct {
	ID                   IDType                   `json:",omitempty" xml:",omitempty"`
	RecipeElementID      RecipeElementIDType      `json:",omitempty" xml:",omitempty"`
	RecipeElementVersion RecipeElementVersionType `json:",omitempty" xml:",omitempty"`
	Description          []DescriptionType        `json:",omitempty" xml:",omitempty"`
}

// TransitionType was auto-generated. Do not change.
type TransitionType struct {
	ID                  IDType                  `json:",omitempty" xml:",omitempty"`
	Condition           ConditionType           `json:",omitempty" xml:",omitempty"`
	ConditionAnnotation ConditionAnnotationType `json:",omitempty" xml:",omitempty"`
	Description         []DescriptionType       `json:",omitempty" xml:",omitempty"`
}

// RecipeElementType was auto-generated. Do not change.
type RecipeElementType struct {
	ID                          IDType                          `json:",omitempty" xml:",omitempty"`
	Version                     VersionType                     `json:",omitempty" xml:",omitempty"`
	VersionDate                 VersionDateType                 `json:",omitempty" xml:",omitempty"`
	Description                 []DescriptionType               `json:",omitempty" xml:",omitempty"`
	RecipeElementType1          RecipeElementTypeType           `json:",omitempty" xml:",omitempty"`
	BuildingBlockElementID      BuildingBlockElementIDType      `json:",omitempty" xml:",omitempty"`
	BuildingBlockElementVersion BuildingBlockElementVersionType `json:",omitempty" xml:",omitempty"`
	ActualEquipmentID           []ActualEquipmentIDType         `json:",omitempty" xml:",omitempty"`
	Header                      HeaderType                      `json:",omitempty" xml:",omitempty"`
	EquipmentRequirement        []BatchEquipmentRequirementType `json:",omitempty" xml:",omitempty"`
	Parameter                   []BatchParameterType            `json:",omitempty" xml:",omitempty"`
	ProcedureLogic              ProcedureLogicType              `json:",omitempty" xml:",omitempty"`
	RecipeElement               []RecipeElementType             `json:",omitempty" xml:",omitempty"`
	OtherInformation            []OtherInformationType          `json:",omitempty" xml:",omitempty"`
}

// OtherInformationType was auto-generated. Do not change.
type OtherInformationType struct {
	ID          IDType            `json:",omitempty" xml:",omitempty"`
	Value       []BatchValueType  `json:",omitempty" xml:",omitempty"`
	Description []DescriptionType `json:",omitempty" xml:",omitempty"`
}

// ControlRecipeType was auto-generated. Do not change.
type ControlRecipeType struct {
	ID                   IDType                          `json:",omitempty" xml:",omitempty"`
	Version              VersionType                     `json:",omitempty" xml:",omitempty"`
	VersionDate          VersionDateType                 `json:",omitempty" xml:",omitempty"`
	Description          []DescriptionType               `json:",omitempty" xml:",omitempty"`
	BatchID              BatchIDType                     `json:",omitempty" xml:",omitempty"`
	Header               HeaderType                      `json:",omitempty" xml:",omitempty"`
	EquipmentRequirement []BatchEquipmentRequirementType `json:",omitempty" xml:",omitempty"`
	Formula              FormulaType                     `json:",omitempty" xml:",omitempty"`
	ProcedureLogic       ProcedureLogicType              `json:",omitempty" xml:",omitempty"`
	RecipeElement        []RecipeElementType             `json:",omitempty" xml:",omitempty"`
	OtherInformation     []OtherInformationType          `json:",omitempty" xml:",omitempty"`
}

// RecipeBuildingBlockType was auto-generated. Do not change.
type RecipeBuildingBlockType struct {
	Description   []DescriptionType   `json:",omitempty" xml:",omitempty"`
	RecipeElement []RecipeElementType `json:",omitempty" xml:",omitempty"`
}

// EquipmentElementType was auto-generated. Do not change.
type EquipmentElementType struct {
	ID                              IDType                                `json:",omitempty" xml:",omitempty"`
	Description                     []DescriptionType                     `json:",omitempty" xml:",omitempty"`
	EquipmentElementType1           EquipmentElementTypeType              `json:",omitempty" xml:",omitempty"`
	EquipmentElementLevel           EquipmentElementLevelType             `json:",omitempty" xml:",omitempty"`
	ClassInstanceAssociation        []ClassInstanceAssociationType        `json:",omitempty" xml:",omitempty"`
	Property                        []EquipmentElementPropertyType        `json:",omitempty" xml:",omitempty"`
	EquipmentProceduralElementClass []EquipmentProceduralElementClassType `json:",omitempty" xml:",omitempty"`
	EquipmentProceduralElement      []EquipmentProceduralElementType      `json:",omitempty" xml:",omitempty"`
	EquipmentConnection             []EquipmentConnectionType             `json:",omitempty" xml:",omitempty"`
	Items                           []interface{}                         `json:",omitempty" xml:",omitempty"`
}

// ClassInstanceAssociationType was auto-generated. Do not change.
type ClassInstanceAssociationType struct {
	ClassEquipmentID  ClassEquipmentIDType  `json:",omitempty" xml:",omitempty"`
	MemberEquipmentID MemberEquipmentIDType `json:",omitempty" xml:",omitempty"`
	Description       []DescriptionType     `json:",omitempty" xml:",omitempty"`
}

// EquipmentElementPropertyType was auto-generated. Do not change.
type EquipmentElementPropertyType struct {
	ID          IDType                         `json:",omitempty" xml:",omitempty"`
	Description []DescriptionType              `json:",omitempty" xml:",omitempty"`
	Value       BatchValueType                 `json:",omitempty" xml:",omitempty"`
	Property    []EquipmentElementPropertyType `json:",omitempty" xml:",omitempty"`
}

// EquipmentProceduralElementClassType was auto-generated. Do not change.
type EquipmentProceduralElementClassType struct {
	ID                             IDType                             `json:",omitempty" xml:",omitempty"`
	Description                    []DescriptionType                  `json:",omitempty" xml:",omitempty"`
	EquipmentProceduralElementType EquipmentProceduralElementTypeType `json:",omitempty" xml:",omitempty"`
	Parameter                      []BatchParameterType               `json:",omitempty" xml:",omitempty"`
}

// EquipmentProceduralElementType was auto-generated. Do not change.
type EquipmentProceduralElementType struct {
	ID                              IDType                             `json:",omitempty" xml:",omitempty"`
	Description                     []DescriptionType                  `json:",omitempty" xml:",omitempty"`
	EquipmentProceduralElementType1 EquipmentProceduralElementTypeType `json:",omitempty" xml:",omitempty"`
	Items                           []interface{}                      `json:",omitempty" xml:",omitempty"`
}

// EquipmentConnectionType was auto-generated. Do not change.
type EquipmentConnectionType struct {
	ID              IDType                `json:",omitempty" xml:",omitempty"`
	Description     []DescriptionType     `json:",omitempty" xml:",omitempty"`
	ConnectionType  ConnectionTypeType    `json:",omitempty" xml:",omitempty"`
	FromEquipmentID []FromEquipmentIDType `json:",omitempty" xml:",omitempty"`
	ToEquipmentID   []ToEquipmentIDType   `json:",omitempty" xml:",omitempty"`
}

// BatchListType was auto-generated. Do not change.
type BatchListType struct {
	ListHeader     ListHeaderType       `json:",omitempty" xml:",omitempty"`
	Description    []DescriptionType    `json:",omitempty" xml:",omitempty"`
	BatchListEntry []BatchListEntryType `json:",omitempty" xml:",omitempty"`
}

// BatchListEntryType was auto-generated. Do not change.
type BatchListEntryType struct {
	ID                  IDType                  `json:",omitempty" xml:",omitempty"`
	Description         []DescriptionType       `json:",omitempty" xml:",omitempty"`
	BatchListEntryType1 BatchListEntryTypeType  `json:",omitempty" xml:",omitempty"`
	Status              BatchStatusType         `json:",omitempty" xml:",omitempty"`
	Mode                ModeType                `json:",omitempty" xml:",omitempty"`
	ExternalID          ExternalIDType          `json:",omitempty" xml:",omitempty"`
	RecipeID            RecipeIDType            `json:",omitempty" xml:",omitempty"`
	RecipeVersion       RecipeVersionType       `json:",omitempty" xml:",omitempty"`
	BatchID             BatchIDType             `json:",omitempty" xml:",omitempty"`
	LotID               LotIDType               `json:",omitempty" xml:",omitempty"`
	CampaignID          CampaignIDType          `json:",omitempty" xml:",omitempty"`
	ProductID           ProductIDType           `json:",omitempty" xml:",omitempty"`
	OrderID             OrderIDType             `json:",omitempty" xml:",omitempty"`
	StartCondition      StartConditionType      `json:",omitempty" xml:",omitempty"`
	RequestedStartTime  RequestedStartTimeType  `json:",omitempty" xml:",omitempty"`
	ActualStartTime     ActualStartTimeType     `json:",omitempty" xml:",omitempty"`
	RequestedEndTime    RequestedEndTimeType    `json:",omitempty" xml:",omitempty"`
	ActualEndTime       ActualEndTimeType       `json:",omitempty" xml:",omitempty"`
	BatchPriority       BatchPriorityType       `json:",omitempty" xml:",omitempty"`
	RequestedBatchSize  RequestedBatchSizeType  `json:",omitempty" xml:",omitempty"`
	ActualBatchSize     ActualBatchSizeType     `json:",omitempty" xml:",omitempty"`
	UnitOfMeasure       UnitOfMeasureType       `json:",omitempty" xml:",omitempty"`
	Note                []NoteType              `json:",omitempty" xml:",omitempty"`
	Parameter           []BatchParameterType    `json:",omitempty" xml:",omitempty"`
	Items               []interface{}           `json:",omitempty" xml:",omitempty"`
	ActualEquipmentID   []ActualEquipmentIDType `json:",omitempty" xml:",omitempty"`
	BatchListEntry      []BatchListEntryType    `json:",omitempty" xml:",omitempty"`
}

// RequestedBatchSizeType was auto-generated. Do not change.
type RequestedBatchSizeType struct {
	MeasureType
}

// BatchEquipmentIDType was auto-generated. Do not change.
type BatchEquipmentIDType struct {
	Property  EquipmentElementPropertyType `json:",omitempty" xml:",omitempty"`
	Condition ConditionType                `json:",omitempty" xml:",omitempty"`
	Value     BatchValueType               `json:",omitempty" xml:",omitempty"`
}

// EnumerationSetType was auto-generated. Do not change.
type EnumerationSetType struct {
	ID          IDType            `json:",omitempty" xml:",omitempty"`
	Description []DescriptionType `json:",omitempty" xml:",omitempty"`
	Enumeration []EnumerationType `json:",omitempty" xml:",omitempty"`
}

// EnumerationType was auto-generated. Do not change.
type EnumerationType struct {
	EnumerationNumber EnumerationNumberType `json:",omitempty" xml:",omitempty"`
	EnumerationString EnumerationStringType `json:",omitempty" xml:",omitempty"`
	Description       []DescriptionType     `json:",omitempty" xml:",omitempty"`
}

// BatchProductionRecordType was auto-generated. Do not change.
type BatchProductionRecordType struct {
	ID                        IdentifierType                        `json:",omitempty" xml:",omitempty"`
	EntryID                   IdentifierType                        `json:",omitempty" xml:",omitempty"`
	ObjectType                RecordObjectTypeType                  `json:",omitempty" xml:",omitempty"`
	TimeStamp                 DateTimeType                          `json:",omitempty" xml:",omitempty"`
	ExternalReference         IdentifierType                        `json:",omitempty" xml:",omitempty"`
	Description               []DescriptionType                     `json:",omitempty" xml:",omitempty"`
	EquipmentScope            IdentifierType                        `json:",omitempty" xml:",omitempty"`
	PublishedDate             DateTimeType                          `json:",omitempty" xml:",omitempty"`
	CreationDate              DateTimeType                          `json:",omitempty" xml:",omitempty"`
	BatchID                   IdentifierType                        `json:",omitempty" xml:",omitempty"`
	BatchProductionRecordSpec IdentifierType                        `json:",omitempty" xml:",omitempty"`
	CampaignID                IdentifierType                        `json:",omitempty" xml:",omitempty"`
	ChangeIndication          string                                `json:",omitempty" xml:",omitempty"`
	Delimiter                 TextType                              `json:",omitempty" xml:",omitempty"`
	EquipmentID               IdentifierType                        `json:",omitempty" xml:",omitempty"`
	ExpirationDate            DateTimeType                          `json:",omitempty" xml:",omitempty"`
	Language                  CodeType                              `json:",omitempty" xml:",omitempty"`
	LastChangedDate           DateTimeType                          `json:",omitempty" xml:",omitempty"`
	LotID                     IdentifierType                        `json:",omitempty" xml:",omitempty"`
	MaterialDefinitionID      IdentifierType                        `json:",omitempty" xml:",omitempty"`
	PhysicalAssetID           IdentifierType                        `json:",omitempty" xml:",omitempty"`
	RecordStatus              CodeType                              `json:",omitempty" xml:",omitempty"`
	Version                   IdentifierType                        `json:",omitempty" xml:",omitempty"`
	ChangeHistory             []ChangeType                          `json:",omitempty" xml:",omitempty"`
	Comments                  []CommentType                         `json:",omitempty" xml:",omitempty"`
	ControlRecipes            []ControlRecipeRecordType             `json:",omitempty" xml:",omitempty"`
	DataSets                  []DataSetType                         `json:",omitempty" xml:",omitempty"`
	Events                    []SingleEventType                     `json:",omitempty" xml:",omitempty"`
	MasterRecipes             []MasterRecipeRecordType              `json:",omitempty" xml:",omitempty"`
	PersonnelIdentification   []PersonnelIdentificationManifestType `json:",omitempty" xml:",omitempty"`
	OperationsDefinitions     []OperationsDefinitionRecordType      `json:",omitempty" xml:",omitempty"`
	OperationsPerformances    []OperationsPerformanceRecordType     `json:",omitempty" xml:",omitempty"`
	OperationsSchedules       []OperationsScheduleRecordType        `json:",omitempty" xml:",omitempty"`
	ProductDefinitions        []ProductDefinitionRecordType         `json:",omitempty" xml:",omitempty"`
	ProductionPerformances    []ProductionPerformanceRecordType     `json:",omitempty" xml:",omitempty"`
	ProductionSchedules       []ProductionScheduleRecordType        `json:",omitempty" xml:",omitempty"`
	RecipeElements            []RecipeElementRecordType             `json:",omitempty" xml:",omitempty"`
	ResourceQualifications    []ResourceQualificationsManifestType  `json:",omitempty" xml:",omitempty"`
	Samples                   []SampleType                          `json:",omitempty" xml:",omitempty"`
	WorkDirectives            []WorkDirectiveRecordType             `json:",omitempty" xml:",omitempty"`
	WorkMasters               []WorkMasterRecordType                `json:",omitempty" xml:",omitempty"`
	WorkPerformances          []WorkPerformanceRecordType           `json:",omitempty" xml:",omitempty"`
	WorkSchedules             []WorkScheduleRecordType              `json:",omitempty" xml:",omitempty"`
	BatchProductionRecord     *BatchProductionRecordType            `json:",omitempty" xml:",omitempty"`
}

// ChangeType was auto-generated. Do not change.
type ChangeType struct {
	EntryID           IdentifierType       `json:",omitempty" xml:",omitempty"`
	ObjectType        RecordObjectTypeType `json:",omitempty" xml:",omitempty"`
	TimeStamp         DateTimeType         `json:",omitempty" xml:",omitempty"`
	ExternalReference IdentifierType       `json:",omitempty" xml:",omitempty"`
	Description       []DescriptionType    `json:",omitempty" xml:",omitempty"`
	RecordReference   IdentifierType       `json:",omitempty" xml:",omitempty"`
	PrechangeData     []ValueType          `json:",omitempty" xml:",omitempty"`
	Reason            []TextType           `json:",omitempty" xml:",omitempty"`
}

// CommentType was auto-generated. Do not change.
type CommentType struct {
	EntryID           IdentifierType       `json:",omitempty" xml:",omitempty"`
	ObjectType        RecordObjectTypeType `json:",omitempty" xml:",omitempty"`
	TimeStamp         DateTimeType         `json:",omitempty" xml:",omitempty"`
	ExternalReference IdentifierType       `json:",omitempty" xml:",omitempty"`
	Description       []DescriptionType    `json:",omitempty" xml:",omitempty"`
	RecordReference   IdentifierType       `json:",omitempty" xml:",omitempty"`
	CommentText       []TextType           `json:",omitempty" xml:",omitempty"`
	PersonID          NameType             `json:",omitempty" xml:",omitempty"`
}

// ControlRecipeRecordType was auto-generated. Do not change.
type ControlRecipeRecordType struct {
	EntryID           IdentifierType       `json:",omitempty" xml:",omitempty"`
	ObjectType        RecordObjectTypeType `json:",omitempty" xml:",omitempty"`
	TimeStamp         DateTimeType         `json:",omitempty" xml:",omitempty"`
	ExternalReference IdentifierType       `json:",omitempty" xml:",omitempty"`
	Description       []DescriptionType    `json:",omitempty" xml:",omitempty"`
	ControlRecipe     ControlRecipeType    `json:",omitempty" xml:",omitempty"`
}

// DataSetType was auto-generated. Do not change.
type DataSetType struct {
	EntryID              IdentifierType         `json:",omitempty" xml:",omitempty"`
	ObjectType           RecordObjectTypeType   `json:",omitempty" xml:",omitempty"`
	TimeStamp            DateTimeType           `json:",omitempty" xml:",omitempty"`
	ExternalReference    IdentifierType         `json:",omitempty" xml:",omitempty"`
	Description          []DescriptionType      `json:",omitempty" xml:",omitempty"`
	TrendSystemReference IdentifierType         `json:",omitempty" xml:",omitempty"`
	StartTime            DateTimeType           `json:",omitempty" xml:",omitempty"`
	EndTime              DateTimeType           `json:",omitempty" xml:",omitempty"`
	TimeSpecification    TimeSpecificationType  `json:",omitempty" xml:",omitempty"`
	TagSpecification     []TagSpecificationType `json:",omitempty" xml:",omitempty"`
	Items                []interface{}          `json:",omitempty" xml:",omitempty"`
}

// TimeSpecificationType was auto-generated. Do not change.
type TimeSpecificationType struct {
	Relative          bool         `json:",omitempty" xml:",omitempty"`
	RelativeSpecified bool         `json:",omitempty" xml:",omitempty"`
	OffsetTime        DateTimeType `json:",omitempty" xml:",omitempty"`
}

// TagSpecificationType was auto-generated. Do not change.
type TagSpecificationType struct {
	TagIndex                   NumericType       `json:",omitempty" xml:",omitempty"`
	DataType                   DataTypeType      `json:",omitempty" xml:",omitempty"`
	UnitOfMeasure              UnitOfMeasureType `json:",omitempty" xml:",omitempty"`
	DataSource                 []IdentifierType  `json:",omitempty" xml:",omitempty"`
	Alias                      []IdentifierType  `json:",omitempty" xml:",omitempty"`
	Description                []DescriptionType `json:",omitempty" xml:",omitempty"`
	EquipmentID                []IdentifierType  `json:",omitempty" xml:",omitempty"`
	PhysicalAssetID            []IdentifierType  `json:",omitempty" xml:",omitempty"`
	ProceduralElementReference []IdentifierType  `json:",omitempty" xml:",omitempty"`
	Deadband                   []ValueType       `json:",omitempty" xml:",omitempty"`
	SignificantDigits          []IdentifierType  `json:",omitempty" xml:",omitempty"`
	DataCompression            []IdentifierType  `json:",omitempty" xml:",omitempty"`
	SamplingType               []IdentifierType  `json:",omitempty" xml:",omitempty"`
}

// DelimitedDataBlockType was auto-generated. Do not change.
type DelimitedDataBlockType struct {
	TagDelimiter   string `json:",omitempty" xml:",omitempty"`
	OrderDelimiter string `json:",omitempty" xml:",omitempty"`
	DelimitedData  string `json:",omitempty" xml:",omitempty"`
}

// OrderedDataType was auto-generated. Do not change.
type OrderedDataType struct {
	OrderIndex NumericType     `json:",omitempty" xml:",omitempty"`
	TimeValue  DateTimeType    `json:",omitempty" xml:",omitempty"`
	DataValue  []DataValueType `json:",omitempty" xml:",omitempty"`
}

// DataValueType was auto-generated. Do not change.
type DataValueType struct {
	TagIndex NumericType    `json:",omitempty" xml:",omitempty"`
	Value    []string       `json:",omitempty" xml:",omitempty"`
	Quality  IdentifierType `json:",omitempty" xml:",omitempty"`
}

// SingleEventType was auto-generated. Do not change.
type SingleEventType struct {
	EntryID                    IdentifierType       `json:",omitempty" xml:",omitempty"`
	ObjectType                 RecordObjectTypeType `json:",omitempty" xml:",omitempty"`
	TimeStamp                  DateTimeType         `json:",omitempty" xml:",omitempty"`
	ExternalReference          IdentifierType       `json:",omitempty" xml:",omitempty"`
	Description                []DescriptionType    `json:",omitempty" xml:",omitempty"`
	EventType                  EventTypeType        `json:",omitempty" xml:",omitempty"`
	EventSubType               EventSubTypeType     `json:",omitempty" xml:",omitempty"`
	EquipmentID                []IdentifierType     `json:",omitempty" xml:",omitempty"`
	Value                      []ValueType          `json:",omitempty" xml:",omitempty"`
	PreviousValue              []ValueType          `json:",omitempty" xml:",omitempty"`
	MessageText                []TextType           `json:",omitempty" xml:",omitempty"`
	PersonID                   []NameType           `json:",omitempty" xml:",omitempty"`
	ComputerID                 []IdentifierType     `json:",omitempty" xml:",omitempty"`
	PhysicalAssetID            []IdentifierType     `json:",omitempty" xml:",omitempty"`
	ProceduralElementReference []IdentifierType     `json:",omitempty" xml:",omitempty"`
	Category                   []IdentifierType     `json:",omitempty" xml:",omitempty"`
	AlarmData                  []AlarmDataType      `json:",omitempty" xml:",omitempty"`
	AssociatedEventID          []IdentifierType     `json:",omitempty" xml:",omitempty"`
	UserAttribute              []UserAttributeType  `json:",omitempty" xml:",omitempty"`
}

// AlarmDataType was auto-generated. Do not change.
type AlarmDataType struct {
	AlarmEvent CodeType         `json:",omitempty" xml:",omitempty"`
	AlarmType  CodeType         `json:",omitempty" xml:",omitempty"`
	AlarmLimit []ValueType      `json:",omitempty" xml:",omitempty"`
	Priority   []IdentifierType `json:",omitempty" xml:",omitempty"`
}

// UserAttributeType was auto-generated. Do not change.
type UserAttributeType struct {
	AttributeID CodeType    `json:",omitempty" xml:",omitempty"`
	Description []TextType  `json:",omitempty" xml:",omitempty"`
	Value       []ValueType `json:",omitempty" xml:",omitempty"`
}

// MasterRecipeRecordType was auto-generated. Do not change.
type MasterRecipeRecordType struct {
	EntryID           IdentifierType       `json:",omitempty" xml:",omitempty"`
	ObjectType        RecordObjectTypeType `json:",omitempty" xml:",omitempty"`
	TimeStamp         DateTimeType         `json:",omitempty" xml:",omitempty"`
	ExternalReference IdentifierType       `json:",omitempty" xml:",omitempty"`
	Description       []DescriptionType    `json:",omitempty" xml:",omitempty"`
	MasterRecipe      MasterRecipeType     `json:",omitempty" xml:",omitempty"`
}

// PersonnelIdentificationManifestType was auto-generated. Do not change.
type PersonnelIdentificationManifestType struct {
	EntryID           IdentifierType       `json:",omitempty" xml:",omitempty"`
	ObjectType        RecordObjectTypeType `json:",omitempty" xml:",omitempty"`
	TimeStamp         DateTimeType         `json:",omitempty" xml:",omitempty"`
	ExternalReference IdentifierType       `json:",omitempty" xml:",omitempty"`
	Description       []DescriptionType    `json:",omitempty" xml:",omitempty"`
	RecordReference   IdentifierType       `json:",omitempty" xml:",omitempty"`
	Name              []NameType           `json:",omitempty" xml:",omitempty"`
	ChangeIndication  string               `json:",omitempty" xml:",omitempty"`
	Reason            TextType             `json:",omitempty" xml:",omitempty"`
}

// OperationsDefinitionRecordType was auto-generated. Do not change.
type OperationsDefinitionRecordType struct {
	EntryID              IdentifierType           `json:",omitempty" xml:",omitempty"`
	ObjectType           RecordObjectTypeType     `json:",omitempty" xml:",omitempty"`
	TimeStamp            DateTimeType             `json:",omitempty" xml:",omitempty"`
	ExternalReference    IdentifierType           `json:",omitempty" xml:",omitempty"`
	Description          []DescriptionType        `json:",omitempty" xml:",omitempty"`
	OperationsDefinition OperationsDefinitionType `json:",omitempty" xml:",omitempty"`
}

// OperationsPerformanceRecordType was auto-generated. Do not change.
type OperationsPerformanceRecordType struct {
	EntryID               IdentifierType            `json:",omitempty" xml:",omitempty"`
	ObjectType            RecordObjectTypeType      `json:",omitempty" xml:",omitempty"`
	TimeStamp             DateTimeType              `json:",omitempty" xml:",omitempty"`
	ExternalReference     IdentifierType            `json:",omitempty" xml:",omitempty"`
	Description           []DescriptionType         `json:",omitempty" xml:",omitempty"`
	OperationsPerformance OperationsPerformanceType `json:",omitempty" xml:",omitempty"`
}

// OperationsScheduleRecordType was auto-generated. Do not change.
type OperationsScheduleRecordType struct {
	EntryID            IdentifierType         `json:",omitempty" xml:",omitempty"`
	ObjectType         RecordObjectTypeType   `json:",omitempty" xml:",omitempty"`
	TimeStamp          DateTimeType           `json:",omitempty" xml:",omitempty"`
	ExternalReference  IdentifierType         `json:",omitempty" xml:",omitempty"`
	Description        []DescriptionType      `json:",omitempty" xml:",omitempty"`
	OperationsSchedule OperationsScheduleType `json:",omitempty" xml:",omitempty"`
}

// ProductDefinitionRecordType was auto-generated. Do not change.
type ProductDefinitionRecordType struct {
	EntryID           IdentifierType        `json:",omitempty" xml:",omitempty"`
	ObjectType        RecordObjectTypeType  `json:",omitempty" xml:",omitempty"`
	TimeStamp         DateTimeType          `json:",omitempty" xml:",omitempty"`
	ExternalReference IdentifierType        `json:",omitempty" xml:",omitempty"`
	Description       []DescriptionType     `json:",omitempty" xml:",omitempty"`
	ProductDefinition ProductDefinitionType `json:",omitempty" xml:",omitempty"`
}

// ProductionPerformanceRecordType was auto-generated. Do not change.
type ProductionPerformanceRecordType struct {
	EntryID                IdentifierType            `json:",omitempty" xml:",omitempty"`
	ObjectType             RecordObjectTypeType      `json:",omitempty" xml:",omitempty"`
	TimeStamp              DateTimeType              `json:",omitempty" xml:",omitempty"`
	ExternalReference      IdentifierType            `json:",omitempty" xml:",omitempty"`
	Description            []DescriptionType         `json:",omitempty" xml:",omitempty"`
	ProductionPerformaance ProductionPerformanceType `json:",omitempty" xml:",omitempty"`
}

// ProductionScheduleRecordType was auto-generated. Do not change.
type ProductionScheduleRecordType struct {
	EntryID            IdentifierType         `json:",omitempty" xml:",omitempty"`
	ObjectType         RecordObjectTypeType   `json:",omitempty" xml:",omitempty"`
	TimeStamp          DateTimeType           `json:",omitempty" xml:",omitempty"`
	ExternalReference  IdentifierType         `json:",omitempty" xml:",omitempty"`
	Description        []DescriptionType      `json:",omitempty" xml:",omitempty"`
	ProductionSchedule ProductionScheduleType `json:",omitempty" xml:",omitempty"`
}

// RecipeElementRecordType was auto-generated. Do not change.
type RecipeElementRecordType struct {
	EntryID           IdentifierType       `json:",omitempty" xml:",omitempty"`
	ObjectType        RecordObjectTypeType `json:",omitempty" xml:",omitempty"`
	TimeStamp         DateTimeType         `json:",omitempty" xml:",omitempty"`
	ExternalReference IdentifierType       `json:",omitempty" xml:",omitempty"`
	Description       []DescriptionType    `json:",omitempty" xml:",omitempty"`
	RecipeElementType RecipeElementType    `json:",omitempty" xml:",omitempty"`
}

// ResourceQualificationsManifestType was auto-generated. Do not change.
type ResourceQualificationsManifestType struct {
	EntryID             IdentifierType       `json:",omitempty" xml:",omitempty"`
	ObjectType          RecordObjectTypeType `json:",omitempty" xml:",omitempty"`
	TimeStamp           DateTimeType         `json:",omitempty" xml:",omitempty"`
	ExternalReference   IdentifierType       `json:",omitempty" xml:",omitempty"`
	Description         []DescriptionType    `json:",omitempty" xml:",omitempty"`
	RecordReference     IdentifierType       `json:",omitempty" xml:",omitempty"`
	ResourceID          IdentifierType       `json:",omitempty" xml:",omitempty"`
	ResourceUse         IdentifierType       `json:",omitempty" xml:",omitempty"`
	ResourceType        IdentifierType       `json:",omitempty" xml:",omitempty"`
	QualificationStatus CodeType             `json:",omitempty" xml:",omitempty"`
	EffectiveTimeStamp  DateTimeType         `json:",omitempty" xml:",omitempty"`
	ExpirationTimeStamp DateTimeType         `json:",omitempty" xml:",omitempty"`
}

// SampleType was auto-generated. Do not change.
type SampleType struct {
	EntryID                    IdentifierType       `json:",omitempty" xml:",omitempty"`
	ObjectType                 RecordObjectTypeType `json:",omitempty" xml:",omitempty"`
	TimeStamp                  DateTimeType         `json:",omitempty" xml:",omitempty"`
	ExternalReference          IdentifierType       `json:",omitempty" xml:",omitempty"`
	Description                []DescriptionType    `json:",omitempty" xml:",omitempty"`
	SampleSourceID             IdentifierType       `json:",omitempty" xml:",omitempty"`
	SampleSize                 []QuantityValueType  `json:",omitempty" xml:",omitempty"`
	SampleType1                []CodeType           `json:",omitempty" xml:",omitempty"`
	SamplePullReason           []TextType           `json:",omitempty" xml:",omitempty"`
	SampleExpiration           DateTimeType         `json:",omitempty" xml:",omitempty"`
	EquipmentID                []IdentifierType     `json:",omitempty" xml:",omitempty"`
	PhysicalAssetID            []IdentifierType     `json:",omitempty" xml:",omitempty"`
	ProceduralElementReference []IdentifierType     `json:",omitempty" xml:",omitempty"`
	SOPReference               []IdentifierType     `json:",omitempty" xml:",omitempty"`
	SampleTest                 []SampleTestType     `json:",omitempty" xml:",omitempty"`
}

// SampleTestType was auto-generated. Do not change.
type SampleTestType struct {
	EntryID           IdentifierType         `json:",omitempty" xml:",omitempty"`
	ObjectType        RecordObjectTypeType   `json:",omitempty" xml:",omitempty"`
	TimeStamp         DateTimeType           `json:",omitempty" xml:",omitempty"`
	ExternalReference IdentifierType         `json:",omitempty" xml:",omitempty"`
	Description       []DescriptionType      `json:",omitempty" xml:",omitempty"`
	TestCode          CodeType               `json:",omitempty" xml:",omitempty"`
	TestName          IdentifierType         `json:",omitempty" xml:",omitempty"`
	SampleTestResult  []SampleTestResultType `json:",omitempty" xml:",omitempty"`
}

// SampleTestResultType was auto-generated. Do not change.
type SampleTestResultType struct {
	EntryID           IdentifierType       `json:",omitempty" xml:",omitempty"`
	ObjectType        RecordObjectTypeType `json:",omitempty" xml:",omitempty"`
	TimeStamp         DateTimeType         `json:",omitempty" xml:",omitempty"`
	ExternalReference IdentifierType       `json:",omitempty" xml:",omitempty"`
	Description       []DescriptionType    `json:",omitempty" xml:",omitempty"`
	TestDisposition   IdentifierType       `json:",omitempty" xml:",omitempty"`
	EquipmentID       []IdentifierType     `json:",omitempty" xml:",omitempty"`
	PhysicalAssetID   []IdentifierType     `json:",omitempty" xml:",omitempty"`
	AnalysisUsed      []CodeType           `json:",omitempty" xml:",omitempty"`
	Expiration        DateTimeType         `json:",omitempty" xml:",omitempty"`
	Results           []ValueType          `json:",omitempty" xml:",omitempty"`
	ExpectedResults   []ValueType          `json:",omitempty" xml:",omitempty"`
}

// WorkDirectiveRecordType was auto-generated. Do not change.
type WorkDirectiveRecordType struct {
	EntryID           IdentifierType       `json:",omitempty" xml:",omitempty"`
	ObjectType        RecordObjectTypeType `json:",omitempty" xml:",omitempty"`
	TimeStamp         DateTimeType         `json:",omitempty" xml:",omitempty"`
	ExternalReference IdentifierType       `json:",omitempty" xml:",omitempty"`
	Description       []DescriptionType    `json:",omitempty" xml:",omitempty"`
	WorkDirective     WorkDirectiveType    `json:",omitempty" xml:",omitempty"`
}

// WorkMasterRecordType was auto-generated. Do not change.
type WorkMasterRecordType struct {
	EntryID           IdentifierType       `json:",omitempty" xml:",omitempty"`
	ObjectType        RecordObjectTypeType `json:",omitempty" xml:",omitempty"`
	TimeStamp         DateTimeType         `json:",omitempty" xml:",omitempty"`
	ExternalReference IdentifierType       `json:",omitempty" xml:",omitempty"`
	Description       []DescriptionType    `json:",omitempty" xml:",omitempty"`
	WorkMaster        WorkMasterType       `json:",omitempty" xml:",omitempty"`
}

// WorkPerformanceRecordType was auto-generated. Do not change.
type WorkPerformanceRecordType struct {
	EntryID           IdentifierType       `json:",omitempty" xml:",omitempty"`
	ObjectType        RecordObjectTypeType `json:",omitempty" xml:",omitempty"`
	TimeStamp         DateTimeType         `json:",omitempty" xml:",omitempty"`
	ExternalReference IdentifierType       `json:",omitempty" xml:",omitempty"`
	Description       []DescriptionType    `json:",omitempty" xml:",omitempty"`
	WorkPerformaance  WorkPerformanceType  `json:",omitempty" xml:",omitempty"`
}

// WorkScheduleRecordType was auto-generated. Do not change.
type WorkScheduleRecordType struct {
	EntryID           IdentifierType       `json:",omitempty" xml:",omitempty"`
	ObjectType        RecordObjectTypeType `json:",omitempty" xml:",omitempty"`
	TimeStamp         DateTimeType         `json:",omitempty" xml:",omitempty"`
	ExternalReference IdentifierType       `json:",omitempty" xml:",omitempty"`
	Description       []DescriptionType    `json:",omitempty" xml:",omitempty"`
	WorkSchedule      WorkScheduleType     `json:",omitempty" xml:",omitempty"`
}

// GetBatchProductionRecordType was auto-generated. Do not change.
type GetBatchProductionRecordType struct {
	ApplicationArea TransApplicationAreaType             `json:",omitempty" xml:",omitempty"`
	DataArea        GetBatchProductionRecordTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                               `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                               `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// GetBatchProductionRecordTypeDataArea was auto-generated. Do not change.
type GetBatchProductionRecordTypeDataArea struct {
	Get                   []string                    `json:",omitempty" xml:",omitempty"`
	BatchProductionRecord []BatchProductionRecordType `json:",omitempty" xml:",omitempty"`
}

// ShowBatchProductionRecordType was auto-generated. Do not change.
type ShowBatchProductionRecordType struct {
	ApplicationArea TransApplicationAreaType              `json:",omitempty" xml:",omitempty"`
	DataArea        ShowBatchProductionRecordTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ShowBatchProductionRecordTypeDataArea was auto-generated. Do not change.
type ShowBatchProductionRecordTypeDataArea struct {
	Show                  TransShowType               `json:",omitempty" xml:",omitempty"`
	BatchProductionRecord []BatchProductionRecordType `json:",omitempty" xml:",omitempty"`
}

// ProcessBatchProductionRecordType was auto-generated. Do not change.
type ProcessBatchProductionRecordType struct {
	ApplicationArea TransApplicationAreaType                 `json:",omitempty" xml:",omitempty"`
	DataArea        ProcessBatchProductionRecordTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                   `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                   `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ProcessBatchProductionRecordTypeDataArea was auto-generated. Do not change.
type ProcessBatchProductionRecordTypeDataArea struct {
	Process               TransProcessType            `json:",omitempty" xml:",omitempty"`
	BatchProductionRecord []BatchProductionRecordType `json:",omitempty" xml:",omitempty"`
}

// AcknowledgeBatchProductionRecordType was auto-generated. Do not change.
type AcknowledgeBatchProductionRecordType struct {
	ApplicationArea TransApplicationAreaType                     `json:",omitempty" xml:",omitempty"`
	DataArea        AcknowledgeBatchProductionRecordTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                       `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                       `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// AcknowledgeBatchProductionRecordTypeDataArea was auto-generated. Do not change.
type AcknowledgeBatchProductionRecordTypeDataArea struct {
	Acknowledge           TransAcknowledgeType        `json:",omitempty" xml:",omitempty"`
	BatchProductionRecord []BatchProductionRecordType `json:",omitempty" xml:",omitempty"`
}

// ChangeBatchProductionRecordType was auto-generated. Do not change.
type ChangeBatchProductionRecordType struct {
	ApplicationArea TransApplicationAreaType                `json:",omitempty" xml:",omitempty"`
	DataArea        ChangeBatchProductionRecordTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                  `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                  `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ChangeBatchProductionRecordTypeDataArea was auto-generated. Do not change.
type ChangeBatchProductionRecordTypeDataArea struct {
	Change                TransChangeType             `json:",omitempty" xml:",omitempty"`
	BatchProductionRecord []BatchProductionRecordType `json:",omitempty" xml:",omitempty"`
}

// RespondBatchProductionRecordType was auto-generated. Do not change.
type RespondBatchProductionRecordType struct {
	ApplicationArea TransApplicationAreaType                 `json:",omitempty" xml:",omitempty"`
	DataArea        RespondBatchProductionRecordTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                   `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                   `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// RespondBatchProductionRecordTypeDataArea was auto-generated. Do not change.
type RespondBatchProductionRecordTypeDataArea struct {
	Respond               TransRespondType            `json:",omitempty" xml:",omitempty"`
	BatchProductionRecord []BatchProductionRecordType `json:",omitempty" xml:",omitempty"`
}

// CancelBatchProductionRecordType was auto-generated. Do not change.
type CancelBatchProductionRecordType struct {
	ApplicationArea TransApplicationAreaType                `json:",omitempty" xml:",omitempty"`
	DataArea        CancelBatchProductionRecordTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                  `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                  `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// CancelBatchProductionRecordTypeDataArea was auto-generated. Do not change.
type CancelBatchProductionRecordTypeDataArea struct {
	Cancel                []TransActionCriteriaType   `json:",omitempty" xml:",omitempty"`
	BatchProductionRecord []BatchProductionRecordType `json:",omitempty" xml:",omitempty"`
}

// SyncBatchProductionRecordType was auto-generated. Do not change.
type SyncBatchProductionRecordType struct {
	ApplicationArea TransApplicationAreaType              `json:",omitempty" xml:",omitempty"`
	DataArea        SyncBatchProductionRecordTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// SyncBatchProductionRecordTypeDataArea was auto-generated. Do not change.
type SyncBatchProductionRecordTypeDataArea struct {
	Sync                  []TransActionCriteriaType   `json:",omitempty" xml:",omitempty"`
	BatchProductionRecord []BatchProductionRecordType `json:",omitempty" xml:",omitempty"`
}

// GRecipeInformationType was auto-generated. Do not change.
type GRecipeInformationType struct {
	ID             IdentifierType     `json:",omitempty" xml:",omitempty"`
	Description    []DescriptionType  `json:",omitempty" xml:",omitempty"`
	HierarchyScope HierarchyScopeType `json:",omitempty" xml:",omitempty"`
	PublishedDate  PublishedDateType  `json:",omitempty" xml:",omitempty"`
	Recipe         []GRecipeType      `json:",omitempty" xml:",omitempty"`
}

// GRecipeType was auto-generated. Do not change.
type GRecipeType struct {
	ID                 IdentifierType           `json:",omitempty" xml:",omitempty"`
	Description        []DescriptionType        `json:",omitempty" xml:",omitempty"`
	GRecipeType1       GRecipeTypeType          `json:",omitempty" xml:",omitempty"`
	LifeCycleState     LifeCycleStateType       `json:",omitempty" xml:",omitempty"`
	Header             GRecipeHeaderType        `json:",omitempty" xml:",omitempty"`
	Formula            GRecipeFormulaType       `json:",omitempty" xml:",omitempty"`
	ProcessProcedure   ProcessElementType       `json:",omitempty" xml:",omitempty"`
	ResourceConstraint []ResourceConstraintType `json:",omitempty" xml:",omitempty"`
	OtherInformation   []GROtherInformationType `json:",omitempty" xml:",omitempty"`
}

// GRecipeHeaderType was auto-generated. Do not change.
type GRecipeHeaderType struct {
	Description    []DescriptionType  `json:",omitempty" xml:",omitempty"`
	DerivedFromID  IdentifierType     `json:",omitempty" xml:",omitempty"`
	ProductID      []IdentifierType   `json:",omitempty" xml:",omitempty"`
	Description1   []DescriptionType  `json:",omitempty" xml:",omitempty"`
	ProductName    []NameType         `json:",omitempty" xml:",omitempty"`
	BatchSize      []ValueType        `json:",omitempty" xml:",omitempty"`
	EffectiveDate  DateTimeType       `json:",omitempty" xml:",omitempty"`
	ExpirationDate DateTimeType       `json:",omitempty" xml:",omitempty"`
	HeaderProperty HeaderPropertyType `json:",omitempty" xml:",omitempty"`
}

// HeaderPropertyType was auto-generated. Do not change.
type HeaderPropertyType struct {
	ID          IdentifierType    `json:",omitempty" xml:",omitempty"`
	Description []DescriptionType `json:",omitempty" xml:",omitempty"`
	Value       []ValueType       `json:",omitempty" xml:",omitempty"`
}

// GRecipeFormulaType was auto-generated. Do not change.
type GRecipeFormulaType struct {
	Description             []DescriptionType             `json:",omitempty" xml:",omitempty"`
	ProcessInputs           GRecipeMaterialsType          `json:",omitempty" xml:",omitempty"`
	ProcessOutputs          GRecipeMaterialsType          `json:",omitempty" xml:",omitempty"`
	ProcessIntermediates    GRecipeMaterialsType          `json:",omitempty" xml:",omitempty"`
	ProcessElementParameter []ProcessElementParameterType `json:",omitempty" xml:",omitempty"`
}

// GRecipeMaterialsType was auto-generated. Do not change.
type GRecipeMaterialsType struct {
	ID            IdentifierType        `json:",omitempty" xml:",omitempty"`
	Description   []DescriptionType     `json:",omitempty" xml:",omitempty"`
	MaterialsType MaterialsTypeType     `json:",omitempty" xml:",omitempty"`
	Material      []GRecipeMaterialType `json:",omitempty" xml:",omitempty"`
}

// GRecipeMaterialType was auto-generated. Do not change.
type GRecipeMaterialType struct {
	ID          IdentifierType    `json:",omitempty" xml:",omitempty"`
	Description []DescriptionType `json:",omitempty" xml:",omitempty"`
	MaterialID  IdentifierType    `json:",omitempty" xml:",omitempty"`
	Order       NumericType       `json:",omitempty" xml:",omitempty"`
	Amount      QuantityValueType `json:",omitempty" xml:",omitempty"`
}

// ProcessElementParameterType was auto-generated. Do not change.
type ProcessElementParameterType struct {
	ID                           IdentifierType                    `json:",omitempty" xml:",omitempty"`
	Description                  []DescriptionType                 `json:",omitempty" xml:",omitempty"`
	ProcessElementParameterType1 []ProcessElementParameterTypeType `json:",omitempty" xml:",omitempty"`
	Value                        []ValueType                       `json:",omitempty" xml:",omitempty"`
}

// ProcessElementType was auto-generated. Do not change.
type ProcessElementType struct {
	ID                      IdentifierType                `json:",omitempty" xml:",omitempty"`
	Description             []DescriptionType             `json:",omitempty" xml:",omitempty"`
	ProcessElementType1     ProcessElementTypeType        `json:",omitempty" xml:",omitempty"`
	LifeCycleState          LifeCycleStateType            `json:",omitempty" xml:",omitempty"`
	SequenceOrder           SequenceOrderType             `json:",omitempty" xml:",omitempty"`
	SequencePath            NumericType                   `json:",omitempty" xml:",omitempty"`
	Materials               []GRecipeMaterialsType        `json:",omitempty" xml:",omitempty"`
	DirectedLink            []DirectedLinkType            `json:",omitempty" xml:",omitempty"`
	ProcedureChartElement   []ProcedureChartElementType   `json:",omitempty" xml:",omitempty"`
	ProcessElement          []ProcessElementType          `json:",omitempty" xml:",omitempty"`
	ProcessElementParameter []ProcessElementParameterType `json:",omitempty" xml:",omitempty"`
	ResourceConstraint      []ResourceConstraintType      `json:",omitempty" xml:",omitempty"`
	OtherInformation        []GROtherInformationType      `json:",omitempty" xml:",omitempty"`
}

// DirectedLinkType was auto-generated. Do not change.
type DirectedLinkType struct {
	ID          IdentifierType    `json:",omitempty" xml:",omitempty"`
	Description []DescriptionType `json:",omitempty" xml:",omitempty"`
	FromID      IdentifierType    `json:",omitempty" xml:",omitempty"`
	ToID        IdentifierType    `json:",omitempty" xml:",omitempty"`
}

// ProcedureChartElementType was auto-generated. Do not change.
type ProcedureChartElementType struct {
	ID                         IdentifierType                `json:",omitempty" xml:",omitempty"`
	Label                      TextType                      `json:",omitempty" xml:",omitempty"`
	Description                []DescriptionType             `json:",omitempty" xml:",omitempty"`
	ProcedureChartElementType1 ProcedureChartElementTypeType `json:",omitempty" xml:",omitempty"`
}

// ResourceConstraintType was auto-generated. Do not change.
type ResourceConstraintType struct {
	ConstraintID               IdentifierType                   `json:",omitempty" xml:",omitempty"`
	Description                []DescriptionType                `json:",omitempty" xml:",omitempty"`
	ConstraintType             []ConstraintTypeType             `json:",omitempty" xml:",omitempty"`
	LifeCycleState             LifeCycleStateType               `json:",omitempty" xml:",omitempty"`
	Range                      []ValueType                      `json:",omitempty" xml:",omitempty"`
	ResourceConstraintProperty []ResourceConstraintPropertyType `json:",omitempty" xml:",omitempty"`
}

// ResourceConstraintPropertyType was auto-generated. Do not change.
type ResourceConstraintPropertyType struct {
	ID          IdentifierType    `json:",omitempty" xml:",omitempty"`
	Description []DescriptionType `json:",omitempty" xml:",omitempty"`
	Value       []ValueType       `json:",omitempty" xml:",omitempty"`
}

// GROtherInformationType was auto-generated. Do not change.
type GROtherInformationType struct {
	OtherInfoID IdentifierType    `json:",omitempty" xml:",omitempty"`
	Description []DescriptionType `json:",omitempty" xml:",omitempty"`
	OtherValue  []ValueType       `json:",omitempty" xml:",omitempty"`
}

// ResourceConstraintLibraryType was auto-generated. Do not change.
type ResourceConstraintLibraryType struct {
	ID                     IdentifierType           `json:",omitempty" xml:",omitempty"`
	Description            []DescriptionType        `json:",omitempty" xml:",omitempty"`
	HierarchyScope         HierarchyScopeType       `json:",omitempty" xml:",omitempty"`
	PublishedDate          PublishedDateType        `json:",omitempty" xml:",omitempty"`
	ResourceConstraintSpec []ResourceConstraintType `json:",omitempty" xml:",omitempty"`
}

// ProcessElementLibraryType was auto-generated. Do not change.
type ProcessElementLibraryType struct {
	ID                 IdentifierType       `json:",omitempty" xml:",omitempty"`
	Description        []DescriptionType    `json:",omitempty" xml:",omitempty"`
	HierarchyScope     HierarchyScopeType   `json:",omitempty" xml:",omitempty"`
	PublishedDate      PublishedDateType    `json:",omitempty" xml:",omitempty"`
	ProcessElementSpec []ProcessElementType `json:",omitempty" xml:",omitempty"`
}

// GetGRecipeInformationType was auto-generated. Do not change.
type GetGRecipeInformationType struct {
	ApplicationArea TransApplicationAreaType          `json:",omitempty" xml:",omitempty"`
	DataArea        GetGRecipeInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                            `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                            `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// GetGRecipeInformationTypeDataArea was auto-generated. Do not change.
type GetGRecipeInformationTypeDataArea struct {
	Get                []string                 `json:",omitempty" xml:",omitempty"`
	GRecipeInformation []GRecipeInformationType `json:",omitempty" xml:",omitempty"`
}

// ShowGRecipeInformationType was auto-generated. Do not change.
type ShowGRecipeInformationType struct {
	ApplicationArea TransApplicationAreaType           `json:",omitempty" xml:",omitempty"`
	DataArea        ShowGRecipeInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                             `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                             `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ShowGRecipeInformationTypeDataArea was auto-generated. Do not change.
type ShowGRecipeInformationTypeDataArea struct {
	Show               TransShowType            `json:",omitempty" xml:",omitempty"`
	GRecipeInformation []GRecipeInformationType `json:",omitempty" xml:",omitempty"`
}

// ProcessGRecipeInformationType was auto-generated. Do not change.
type ProcessGRecipeInformationType struct {
	ApplicationArea TransApplicationAreaType              `json:",omitempty" xml:",omitempty"`
	DataArea        ProcessGRecipeInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ProcessGRecipeInformationTypeDataArea was auto-generated. Do not change.
type ProcessGRecipeInformationTypeDataArea struct {
	Process            TransProcessType         `json:",omitempty" xml:",omitempty"`
	GRecipeInformation []GRecipeInformationType `json:",omitempty" xml:",omitempty"`
}

// AcknowledgeGRecipeInformationType was auto-generated. Do not change.
type AcknowledgeGRecipeInformationType struct {
	ApplicationArea TransApplicationAreaType                  `json:",omitempty" xml:",omitempty"`
	DataArea        AcknowledgeGRecipeInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                    `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                    `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// AcknowledgeGRecipeInformationTypeDataArea was auto-generated. Do not change.
type AcknowledgeGRecipeInformationTypeDataArea struct {
	Acknowledge        TransAcknowledgeType     `json:",omitempty" xml:",omitempty"`
	GRecipeInformation []GRecipeInformationType `json:",omitempty" xml:",omitempty"`
}

// ChangeGRecipeInformationType was auto-generated. Do not change.
type ChangeGRecipeInformationType struct {
	ApplicationArea TransApplicationAreaType             `json:",omitempty" xml:",omitempty"`
	DataArea        ChangeGRecipeInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                               `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                               `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ChangeGRecipeInformationTypeDataArea was auto-generated. Do not change.
type ChangeGRecipeInformationTypeDataArea struct {
	Change             TransChangeType          `json:",omitempty" xml:",omitempty"`
	GRecipeInformation []GRecipeInformationType `json:",omitempty" xml:",omitempty"`
}

// RespondGRecipeInformationType was auto-generated. Do not change.
type RespondGRecipeInformationType struct {
	ApplicationArea TransApplicationAreaType              `json:",omitempty" xml:",omitempty"`
	DataArea        RespondGRecipeInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// RespondGRecipeInformationTypeDataArea was auto-generated. Do not change.
type RespondGRecipeInformationTypeDataArea struct {
	Respond            TransRespondType         `json:",omitempty" xml:",omitempty"`
	GRecipeInformation []GRecipeInformationType `json:",omitempty" xml:",omitempty"`
}

// CancelGRecipeInformationType was auto-generated. Do not change.
type CancelGRecipeInformationType struct {
	ApplicationArea TransApplicationAreaType             `json:",omitempty" xml:",omitempty"`
	DataArea        CancelGRecipeInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                               `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                               `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// CancelGRecipeInformationTypeDataArea was auto-generated. Do not change.
type CancelGRecipeInformationTypeDataArea struct {
	Cancel             []TransActionCriteriaType `json:",omitempty" xml:",omitempty"`
	GRecipeInformation []GRecipeInformationType  `json:",omitempty" xml:",omitempty"`
}

// SyncGRecipeInformationType was auto-generated. Do not change.
type SyncGRecipeInformationType struct {
	ApplicationArea TransApplicationAreaType           `json:",omitempty" xml:",omitempty"`
	DataArea        SyncGRecipeInformationTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                             `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                             `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// SyncGRecipeInformationTypeDataArea was auto-generated. Do not change.
type SyncGRecipeInformationTypeDataArea struct {
	Sync               []TransActionCriteriaType `json:",omitempty" xml:",omitempty"`
	GRecipeInformation []GRecipeInformationType  `json:",omitempty" xml:",omitempty"`
}

// GetResourceConstraintLibraryType was auto-generated. Do not change.
type GetResourceConstraintLibraryType struct {
	ApplicationArea TransApplicationAreaType                 `json:",omitempty" xml:",omitempty"`
	DataArea        GetResourceConstraintLibraryTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                   `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                   `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// GetResourceConstraintLibraryTypeDataArea was auto-generated. Do not change.
type GetResourceConstraintLibraryTypeDataArea struct {
	Get                       []string                        `json:",omitempty" xml:",omitempty"`
	ResourceConstraintLibrary []ResourceConstraintLibraryType `json:",omitempty" xml:",omitempty"`
}

// ShowResourceConstraintLibraryType was auto-generated. Do not change.
type ShowResourceConstraintLibraryType struct {
	ApplicationArea TransApplicationAreaType                  `json:",omitempty" xml:",omitempty"`
	DataArea        ShowResourceConstraintLibraryTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                    `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                    `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ShowResourceConstraintLibraryTypeDataArea was auto-generated. Do not change.
type ShowResourceConstraintLibraryTypeDataArea struct {
	Show                      TransShowType                   `json:",omitempty" xml:",omitempty"`
	ResourceConstraintLibrary []ResourceConstraintLibraryType `json:",omitempty" xml:",omitempty"`
}

// ProcessResourceConstraintLibraryType was auto-generated. Do not change.
type ProcessResourceConstraintLibraryType struct {
	ApplicationArea TransApplicationAreaType                     `json:",omitempty" xml:",omitempty"`
	DataArea        ProcessResourceConstraintLibraryTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                       `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                       `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ProcessResourceConstraintLibraryTypeDataArea was auto-generated. Do not change.
type ProcessResourceConstraintLibraryTypeDataArea struct {
	Process                   TransProcessType                `json:",omitempty" xml:",omitempty"`
	ResourceConstraintLibrary []ResourceConstraintLibraryType `json:",omitempty" xml:",omitempty"`
}

// AcknowledgeResourceConstraintLibraryType was auto-generated. Do not change.
type AcknowledgeResourceConstraintLibraryType struct {
	ApplicationArea TransApplicationAreaType                         `json:",omitempty" xml:",omitempty"`
	DataArea        AcknowledgeResourceConstraintLibraryTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                           `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                           `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// AcknowledgeResourceConstraintLibraryTypeDataArea was auto-generated. Do not change.
type AcknowledgeResourceConstraintLibraryTypeDataArea struct {
	Acknowledge               TransAcknowledgeType            `json:",omitempty" xml:",omitempty"`
	ResourceConstraintLibrary []ResourceConstraintLibraryType `json:",omitempty" xml:",omitempty"`
}

// ChangeResourceConstraintLibraryType was auto-generated. Do not change.
type ChangeResourceConstraintLibraryType struct {
	ApplicationArea TransApplicationAreaType                    `json:",omitempty" xml:",omitempty"`
	DataArea        ChangeResourceConstraintLibraryTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                      `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                      `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ChangeResourceConstraintLibraryTypeDataArea was auto-generated. Do not change.
type ChangeResourceConstraintLibraryTypeDataArea struct {
	Change                    TransChangeType                 `json:",omitempty" xml:",omitempty"`
	ResourceConstraintLibrary []ResourceConstraintLibraryType `json:",omitempty" xml:",omitempty"`
}

// RespondResourceConstraintLibraryType was auto-generated. Do not change.
type RespondResourceConstraintLibraryType struct {
	ApplicationArea TransApplicationAreaType                     `json:",omitempty" xml:",omitempty"`
	DataArea        RespondResourceConstraintLibraryTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                       `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                       `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// RespondResourceConstraintLibraryTypeDataArea was auto-generated. Do not change.
type RespondResourceConstraintLibraryTypeDataArea struct {
	Respond                   TransRespondType                `json:",omitempty" xml:",omitempty"`
	ResourceConstraintLibrary []ResourceConstraintLibraryType `json:",omitempty" xml:",omitempty"`
}

// CancelResourceConstraintLibraryType was auto-generated. Do not change.
type CancelResourceConstraintLibraryType struct {
	ApplicationArea TransApplicationAreaType                    `json:",omitempty" xml:",omitempty"`
	DataArea        CancelResourceConstraintLibraryTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                      `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                      `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// CancelResourceConstraintLibraryTypeDataArea was auto-generated. Do not change.
type CancelResourceConstraintLibraryTypeDataArea struct {
	Cancel                    []TransActionCriteriaType       `json:",omitempty" xml:",omitempty"`
	ResourceConstraintLibrary []ResourceConstraintLibraryType `json:",omitempty" xml:",omitempty"`
}

// SyncResourceConstraintLibraryType was auto-generated. Do not change.
type SyncResourceConstraintLibraryType struct {
	ApplicationArea TransApplicationAreaType                  `json:",omitempty" xml:",omitempty"`
	DataArea        SyncResourceConstraintLibraryTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                    `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                    `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// SyncResourceConstraintLibraryTypeDataArea was auto-generated. Do not change.
type SyncResourceConstraintLibraryTypeDataArea struct {
	Sync                      []TransActionCriteriaType       `json:",omitempty" xml:",omitempty"`
	ResourceConstraintLibrary []ResourceConstraintLibraryType `json:",omitempty" xml:",omitempty"`
}

// GetProcessElementLibraryType was auto-generated. Do not change.
type GetProcessElementLibraryType struct {
	ApplicationArea TransApplicationAreaType             `json:",omitempty" xml:",omitempty"`
	DataArea        GetProcessElementLibraryTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                               `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                               `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// GetProcessElementLibraryTypeDataArea was auto-generated. Do not change.
type GetProcessElementLibraryTypeDataArea struct {
	Get                   []string                    `json:",omitempty" xml:",omitempty"`
	ProcessElementLibrary []ProcessElementLibraryType `json:",omitempty" xml:",omitempty"`
}

// ShowProcessElementLibraryType was auto-generated. Do not change.
type ShowProcessElementLibraryType struct {
	ApplicationArea TransApplicationAreaType              `json:",omitempty" xml:",omitempty"`
	DataArea        ShowProcessElementLibraryTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ShowProcessElementLibraryTypeDataArea was auto-generated. Do not change.
type ShowProcessElementLibraryTypeDataArea struct {
	Show                  TransShowType               `json:",omitempty" xml:",omitempty"`
	ProcessElementLibrary []ProcessElementLibraryType `json:",omitempty" xml:",omitempty"`
}

// ProcessProcessElementLibraryType was auto-generated. Do not change.
type ProcessProcessElementLibraryType struct {
	ApplicationArea TransApplicationAreaType                 `json:",omitempty" xml:",omitempty"`
	DataArea        ProcessProcessElementLibraryTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                   `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                   `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ProcessProcessElementLibraryTypeDataArea was auto-generated. Do not change.
type ProcessProcessElementLibraryTypeDataArea struct {
	Process               TransProcessType            `json:",omitempty" xml:",omitempty"`
	ProcessElementLibrary []ProcessElementLibraryType `json:",omitempty" xml:",omitempty"`
}

// AcknowledgeProcessElementLibraryType was auto-generated. Do not change.
type AcknowledgeProcessElementLibraryType struct {
	ApplicationArea TransApplicationAreaType                     `json:",omitempty" xml:",omitempty"`
	DataArea        AcknowledgeProcessElementLibraryTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                       `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                       `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// AcknowledgeProcessElementLibraryTypeDataArea was auto-generated. Do not change.
type AcknowledgeProcessElementLibraryTypeDataArea struct {
	Acknowledge           TransAcknowledgeType        `json:",omitempty" xml:",omitempty"`
	ProcessElementLibrary []ProcessElementLibraryType `json:",omitempty" xml:",omitempty"`
}

// ChangeProcessElementLibraryType was auto-generated. Do not change.
type ChangeProcessElementLibraryType struct {
	ApplicationArea TransApplicationAreaType                `json:",omitempty" xml:",omitempty"`
	DataArea        ChangeProcessElementLibraryTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                  `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                  `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// ChangeProcessElementLibraryTypeDataArea was auto-generated. Do not change.
type ChangeProcessElementLibraryTypeDataArea struct {
	Change                TransChangeType             `json:",omitempty" xml:",omitempty"`
	ProcessElementLibrary []ProcessElementLibraryType `json:",omitempty" xml:",omitempty"`
}

// RespondProcessElementLibraryType was auto-generated. Do not change.
type RespondProcessElementLibraryType struct {
	ApplicationArea TransApplicationAreaType                 `json:",omitempty" xml:",omitempty"`
	DataArea        RespondProcessElementLibraryTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                   `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                   `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// RespondProcessElementLibraryTypeDataArea was auto-generated. Do not change.
type RespondProcessElementLibraryTypeDataArea struct {
	Respond               TransRespondType            `json:",omitempty" xml:",omitempty"`
	ProcessElementLibrary []ProcessElementLibraryType `json:",omitempty" xml:",omitempty"`
}

// CancelProcessElementLibraryType was auto-generated. Do not change.
type CancelProcessElementLibraryType struct {
	ApplicationArea TransApplicationAreaType                `json:",omitempty" xml:",omitempty"`
	DataArea        CancelProcessElementLibraryTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                  `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                  `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// CancelProcessElementLibraryTypeDataArea was auto-generated. Do not change.
type CancelProcessElementLibraryTypeDataArea struct {
	Cancel                []TransActionCriteriaType   `json:",omitempty" xml:",omitempty"`
	ProcessElementLibrary []ProcessElementLibraryType `json:",omitempty" xml:",omitempty"`
}

// SyncProcessElementLibraryType was auto-generated. Do not change.
type SyncProcessElementLibraryType struct {
	ApplicationArea TransApplicationAreaType              `json:",omitempty" xml:",omitempty"`
	DataArea        SyncProcessElementLibraryTypeDataArea `json:",omitempty" xml:",omitempty"`
	ReleaseID       string                                `json:",omitempty" xml:"releaseID,attr,omitempty"`
	VersionID       string                                `json:",omitempty" xml:"versionID,attr,omitempty"`
}

// SyncProcessElementLibraryTypeDataArea was auto-generated. Do not change.
type SyncProcessElementLibraryTypeDataArea struct {
	Sync                  []TransActionCriteriaType   `json:",omitempty" xml:",omitempty"`
	ProcessElementLibrary []ProcessElementLibraryType `json:",omitempty" xml:",omitempty"`
}
