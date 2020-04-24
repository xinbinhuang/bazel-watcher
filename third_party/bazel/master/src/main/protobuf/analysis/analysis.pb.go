// Code generated by protoc-gen-go. DO NOT EDIT.
// source: third_party/bazel/master/src/main/protobuf/analysis/analysis.proto

package analysis

import (
	fmt "fmt"
	blaze_query "github.com/bazelbuild/bazel-watcher/third_party/bazel/master/src/main/protobuf/blaze_query"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ActionGraphContainer struct {
	Artifacts            []*Artifact         `protobuf:"bytes,1,rep,name=artifacts,proto3" json:"artifacts,omitempty"`
	Actions              []*Action           `protobuf:"bytes,2,rep,name=actions,proto3" json:"actions,omitempty"`
	Targets              []*Target           `protobuf:"bytes,3,rep,name=targets,proto3" json:"targets,omitempty"`
	DepSetOfFiles        []*DepSetOfFiles    `protobuf:"bytes,4,rep,name=dep_set_of_files,json=depSetOfFiles,proto3" json:"dep_set_of_files,omitempty"`
	Configuration        []*Configuration    `protobuf:"bytes,5,rep,name=configuration,proto3" json:"configuration,omitempty"`
	AspectDescriptors    []*AspectDescriptor `protobuf:"bytes,6,rep,name=aspect_descriptors,json=aspectDescriptors,proto3" json:"aspect_descriptors,omitempty"`
	RuleClasses          []*RuleClass        `protobuf:"bytes,7,rep,name=rule_classes,json=ruleClasses,proto3" json:"rule_classes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ActionGraphContainer) Reset()         { *m = ActionGraphContainer{} }
func (m *ActionGraphContainer) String() string { return proto.CompactTextString(m) }
func (*ActionGraphContainer) ProtoMessage()    {}
func (*ActionGraphContainer) Descriptor() ([]byte, []int) {
	return fileDescriptor_b577e5ef8e9d5380, []int{0}
}

func (m *ActionGraphContainer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ActionGraphContainer.Unmarshal(m, b)
}
func (m *ActionGraphContainer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ActionGraphContainer.Marshal(b, m, deterministic)
}
func (m *ActionGraphContainer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActionGraphContainer.Merge(m, src)
}
func (m *ActionGraphContainer) XXX_Size() int {
	return xxx_messageInfo_ActionGraphContainer.Size(m)
}
func (m *ActionGraphContainer) XXX_DiscardUnknown() {
	xxx_messageInfo_ActionGraphContainer.DiscardUnknown(m)
}

var xxx_messageInfo_ActionGraphContainer proto.InternalMessageInfo

func (m *ActionGraphContainer) GetArtifacts() []*Artifact {
	if m != nil {
		return m.Artifacts
	}
	return nil
}

func (m *ActionGraphContainer) GetActions() []*Action {
	if m != nil {
		return m.Actions
	}
	return nil
}

func (m *ActionGraphContainer) GetTargets() []*Target {
	if m != nil {
		return m.Targets
	}
	return nil
}

func (m *ActionGraphContainer) GetDepSetOfFiles() []*DepSetOfFiles {
	if m != nil {
		return m.DepSetOfFiles
	}
	return nil
}

func (m *ActionGraphContainer) GetConfiguration() []*Configuration {
	if m != nil {
		return m.Configuration
	}
	return nil
}

func (m *ActionGraphContainer) GetAspectDescriptors() []*AspectDescriptor {
	if m != nil {
		return m.AspectDescriptors
	}
	return nil
}

func (m *ActionGraphContainer) GetRuleClasses() []*RuleClass {
	if m != nil {
		return m.RuleClasses
	}
	return nil
}

type Artifact struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ExecPath             string   `protobuf:"bytes,2,opt,name=exec_path,json=execPath,proto3" json:"exec_path,omitempty"`
	IsTreeArtifact       bool     `protobuf:"varint,3,opt,name=is_tree_artifact,json=isTreeArtifact,proto3" json:"is_tree_artifact,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Artifact) Reset()         { *m = Artifact{} }
func (m *Artifact) String() string { return proto.CompactTextString(m) }
func (*Artifact) ProtoMessage()    {}
func (*Artifact) Descriptor() ([]byte, []int) {
	return fileDescriptor_b577e5ef8e9d5380, []int{1}
}

func (m *Artifact) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Artifact.Unmarshal(m, b)
}
func (m *Artifact) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Artifact.Marshal(b, m, deterministic)
}
func (m *Artifact) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Artifact.Merge(m, src)
}
func (m *Artifact) XXX_Size() int {
	return xxx_messageInfo_Artifact.Size(m)
}
func (m *Artifact) XXX_DiscardUnknown() {
	xxx_messageInfo_Artifact.DiscardUnknown(m)
}

var xxx_messageInfo_Artifact proto.InternalMessageInfo

func (m *Artifact) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Artifact) GetExecPath() string {
	if m != nil {
		return m.ExecPath
	}
	return ""
}

func (m *Artifact) GetIsTreeArtifact() bool {
	if m != nil {
		return m.IsTreeArtifact
	}
	return false
}

type Action struct {
	TargetId             string          `protobuf:"bytes,1,opt,name=target_id,json=targetId,proto3" json:"target_id,omitempty"`
	AspectDescriptorIds  []string        `protobuf:"bytes,2,rep,name=aspect_descriptor_ids,json=aspectDescriptorIds,proto3" json:"aspect_descriptor_ids,omitempty"`
	ActionKey            string          `protobuf:"bytes,3,opt,name=action_key,json=actionKey,proto3" json:"action_key,omitempty"`
	Mnemonic             string          `protobuf:"bytes,4,opt,name=mnemonic,proto3" json:"mnemonic,omitempty"`
	ConfigurationId      string          `protobuf:"bytes,5,opt,name=configuration_id,json=configurationId,proto3" json:"configuration_id,omitempty"`
	Arguments            []string        `protobuf:"bytes,6,rep,name=arguments,proto3" json:"arguments,omitempty"`
	EnvironmentVariables []*KeyValuePair `protobuf:"bytes,7,rep,name=environment_variables,json=environmentVariables,proto3" json:"environment_variables,omitempty"`
	InputDepSetIds       []string        `protobuf:"bytes,8,rep,name=input_dep_set_ids,json=inputDepSetIds,proto3" json:"input_dep_set_ids,omitempty"`
	OutputIds            []string        `protobuf:"bytes,9,rep,name=output_ids,json=outputIds,proto3" json:"output_ids,omitempty"`
	DiscoversInputs      bool            `protobuf:"varint,10,opt,name=discovers_inputs,json=discoversInputs,proto3" json:"discovers_inputs,omitempty"`
	ExecutionInfo        []*KeyValuePair `protobuf:"bytes,11,rep,name=execution_info,json=executionInfo,proto3" json:"execution_info,omitempty"`
	ParamFiles           []*ParamFile    `protobuf:"bytes,12,rep,name=param_files,json=paramFiles,proto3" json:"param_files,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Action) Reset()         { *m = Action{} }
func (m *Action) String() string { return proto.CompactTextString(m) }
func (*Action) ProtoMessage()    {}
func (*Action) Descriptor() ([]byte, []int) {
	return fileDescriptor_b577e5ef8e9d5380, []int{2}
}

func (m *Action) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Action.Unmarshal(m, b)
}
func (m *Action) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Action.Marshal(b, m, deterministic)
}
func (m *Action) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Action.Merge(m, src)
}
func (m *Action) XXX_Size() int {
	return xxx_messageInfo_Action.Size(m)
}
func (m *Action) XXX_DiscardUnknown() {
	xxx_messageInfo_Action.DiscardUnknown(m)
}

var xxx_messageInfo_Action proto.InternalMessageInfo

func (m *Action) GetTargetId() string {
	if m != nil {
		return m.TargetId
	}
	return ""
}

func (m *Action) GetAspectDescriptorIds() []string {
	if m != nil {
		return m.AspectDescriptorIds
	}
	return nil
}

func (m *Action) GetActionKey() string {
	if m != nil {
		return m.ActionKey
	}
	return ""
}

func (m *Action) GetMnemonic() string {
	if m != nil {
		return m.Mnemonic
	}
	return ""
}

func (m *Action) GetConfigurationId() string {
	if m != nil {
		return m.ConfigurationId
	}
	return ""
}

func (m *Action) GetArguments() []string {
	if m != nil {
		return m.Arguments
	}
	return nil
}

func (m *Action) GetEnvironmentVariables() []*KeyValuePair {
	if m != nil {
		return m.EnvironmentVariables
	}
	return nil
}

func (m *Action) GetInputDepSetIds() []string {
	if m != nil {
		return m.InputDepSetIds
	}
	return nil
}

func (m *Action) GetOutputIds() []string {
	if m != nil {
		return m.OutputIds
	}
	return nil
}

func (m *Action) GetDiscoversInputs() bool {
	if m != nil {
		return m.DiscoversInputs
	}
	return false
}

func (m *Action) GetExecutionInfo() []*KeyValuePair {
	if m != nil {
		return m.ExecutionInfo
	}
	return nil
}

func (m *Action) GetParamFiles() []*ParamFile {
	if m != nil {
		return m.ParamFiles
	}
	return nil
}

type Target struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Label                string   `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
	RuleClassId          string   `protobuf:"bytes,3,opt,name=rule_class_id,json=ruleClassId,proto3" json:"rule_class_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Target) Reset()         { *m = Target{} }
func (m *Target) String() string { return proto.CompactTextString(m) }
func (*Target) ProtoMessage()    {}
func (*Target) Descriptor() ([]byte, []int) {
	return fileDescriptor_b577e5ef8e9d5380, []int{3}
}

func (m *Target) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Target.Unmarshal(m, b)
}
func (m *Target) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Target.Marshal(b, m, deterministic)
}
func (m *Target) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Target.Merge(m, src)
}
func (m *Target) XXX_Size() int {
	return xxx_messageInfo_Target.Size(m)
}
func (m *Target) XXX_DiscardUnknown() {
	xxx_messageInfo_Target.DiscardUnknown(m)
}

var xxx_messageInfo_Target proto.InternalMessageInfo

func (m *Target) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Target) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *Target) GetRuleClassId() string {
	if m != nil {
		return m.RuleClassId
	}
	return ""
}

type RuleClass struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RuleClass) Reset()         { *m = RuleClass{} }
func (m *RuleClass) String() string { return proto.CompactTextString(m) }
func (*RuleClass) ProtoMessage()    {}
func (*RuleClass) Descriptor() ([]byte, []int) {
	return fileDescriptor_b577e5ef8e9d5380, []int{4}
}

func (m *RuleClass) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RuleClass.Unmarshal(m, b)
}
func (m *RuleClass) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RuleClass.Marshal(b, m, deterministic)
}
func (m *RuleClass) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RuleClass.Merge(m, src)
}
func (m *RuleClass) XXX_Size() int {
	return xxx_messageInfo_RuleClass.Size(m)
}
func (m *RuleClass) XXX_DiscardUnknown() {
	xxx_messageInfo_RuleClass.DiscardUnknown(m)
}

var xxx_messageInfo_RuleClass proto.InternalMessageInfo

func (m *RuleClass) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *RuleClass) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type AspectDescriptor struct {
	Id                   string          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string          `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Parameters           []*KeyValuePair `protobuf:"bytes,3,rep,name=parameters,proto3" json:"parameters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *AspectDescriptor) Reset()         { *m = AspectDescriptor{} }
func (m *AspectDescriptor) String() string { return proto.CompactTextString(m) }
func (*AspectDescriptor) ProtoMessage()    {}
func (*AspectDescriptor) Descriptor() ([]byte, []int) {
	return fileDescriptor_b577e5ef8e9d5380, []int{5}
}

func (m *AspectDescriptor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AspectDescriptor.Unmarshal(m, b)
}
func (m *AspectDescriptor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AspectDescriptor.Marshal(b, m, deterministic)
}
func (m *AspectDescriptor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AspectDescriptor.Merge(m, src)
}
func (m *AspectDescriptor) XXX_Size() int {
	return xxx_messageInfo_AspectDescriptor.Size(m)
}
func (m *AspectDescriptor) XXX_DiscardUnknown() {
	xxx_messageInfo_AspectDescriptor.DiscardUnknown(m)
}

var xxx_messageInfo_AspectDescriptor proto.InternalMessageInfo

func (m *AspectDescriptor) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *AspectDescriptor) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AspectDescriptor) GetParameters() []*KeyValuePair {
	if m != nil {
		return m.Parameters
	}
	return nil
}

type DepSetOfFiles struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	TransitiveDepSetIds  []string `protobuf:"bytes,2,rep,name=transitive_dep_set_ids,json=transitiveDepSetIds,proto3" json:"transitive_dep_set_ids,omitempty"`
	DirectArtifactIds    []string `protobuf:"bytes,3,rep,name=direct_artifact_ids,json=directArtifactIds,proto3" json:"direct_artifact_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DepSetOfFiles) Reset()         { *m = DepSetOfFiles{} }
func (m *DepSetOfFiles) String() string { return proto.CompactTextString(m) }
func (*DepSetOfFiles) ProtoMessage()    {}
func (*DepSetOfFiles) Descriptor() ([]byte, []int) {
	return fileDescriptor_b577e5ef8e9d5380, []int{6}
}

func (m *DepSetOfFiles) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DepSetOfFiles.Unmarshal(m, b)
}
func (m *DepSetOfFiles) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DepSetOfFiles.Marshal(b, m, deterministic)
}
func (m *DepSetOfFiles) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DepSetOfFiles.Merge(m, src)
}
func (m *DepSetOfFiles) XXX_Size() int {
	return xxx_messageInfo_DepSetOfFiles.Size(m)
}
func (m *DepSetOfFiles) XXX_DiscardUnknown() {
	xxx_messageInfo_DepSetOfFiles.DiscardUnknown(m)
}

var xxx_messageInfo_DepSetOfFiles proto.InternalMessageInfo

func (m *DepSetOfFiles) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DepSetOfFiles) GetTransitiveDepSetIds() []string {
	if m != nil {
		return m.TransitiveDepSetIds
	}
	return nil
}

func (m *DepSetOfFiles) GetDirectArtifactIds() []string {
	if m != nil {
		return m.DirectArtifactIds
	}
	return nil
}

type Configuration struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Mnemonic             string   `protobuf:"bytes,2,opt,name=mnemonic,proto3" json:"mnemonic,omitempty"`
	PlatformName         string   `protobuf:"bytes,3,opt,name=platform_name,json=platformName,proto3" json:"platform_name,omitempty"`
	Checksum             string   `protobuf:"bytes,4,opt,name=checksum,proto3" json:"checksum,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Configuration) Reset()         { *m = Configuration{} }
func (m *Configuration) String() string { return proto.CompactTextString(m) }
func (*Configuration) ProtoMessage()    {}
func (*Configuration) Descriptor() ([]byte, []int) {
	return fileDescriptor_b577e5ef8e9d5380, []int{7}
}

func (m *Configuration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Configuration.Unmarshal(m, b)
}
func (m *Configuration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Configuration.Marshal(b, m, deterministic)
}
func (m *Configuration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Configuration.Merge(m, src)
}
func (m *Configuration) XXX_Size() int {
	return xxx_messageInfo_Configuration.Size(m)
}
func (m *Configuration) XXX_DiscardUnknown() {
	xxx_messageInfo_Configuration.DiscardUnknown(m)
}

var xxx_messageInfo_Configuration proto.InternalMessageInfo

func (m *Configuration) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Configuration) GetMnemonic() string {
	if m != nil {
		return m.Mnemonic
	}
	return ""
}

func (m *Configuration) GetPlatformName() string {
	if m != nil {
		return m.PlatformName
	}
	return ""
}

func (m *Configuration) GetChecksum() string {
	if m != nil {
		return m.Checksum
	}
	return ""
}

type KeyValuePair struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeyValuePair) Reset()         { *m = KeyValuePair{} }
func (m *KeyValuePair) String() string { return proto.CompactTextString(m) }
func (*KeyValuePair) ProtoMessage()    {}
func (*KeyValuePair) Descriptor() ([]byte, []int) {
	return fileDescriptor_b577e5ef8e9d5380, []int{8}
}

func (m *KeyValuePair) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyValuePair.Unmarshal(m, b)
}
func (m *KeyValuePair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyValuePair.Marshal(b, m, deterministic)
}
func (m *KeyValuePair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyValuePair.Merge(m, src)
}
func (m *KeyValuePair) XXX_Size() int {
	return xxx_messageInfo_KeyValuePair.Size(m)
}
func (m *KeyValuePair) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyValuePair.DiscardUnknown(m)
}

var xxx_messageInfo_KeyValuePair proto.InternalMessageInfo

func (m *KeyValuePair) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *KeyValuePair) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type ConfiguredTarget struct {
	Target               *blaze_query.Target `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
	Configuration        *Configuration      `protobuf:"bytes,2,opt,name=configuration,proto3" json:"configuration,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ConfiguredTarget) Reset()         { *m = ConfiguredTarget{} }
func (m *ConfiguredTarget) String() string { return proto.CompactTextString(m) }
func (*ConfiguredTarget) ProtoMessage()    {}
func (*ConfiguredTarget) Descriptor() ([]byte, []int) {
	return fileDescriptor_b577e5ef8e9d5380, []int{9}
}

func (m *ConfiguredTarget) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfiguredTarget.Unmarshal(m, b)
}
func (m *ConfiguredTarget) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfiguredTarget.Marshal(b, m, deterministic)
}
func (m *ConfiguredTarget) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfiguredTarget.Merge(m, src)
}
func (m *ConfiguredTarget) XXX_Size() int {
	return xxx_messageInfo_ConfiguredTarget.Size(m)
}
func (m *ConfiguredTarget) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfiguredTarget.DiscardUnknown(m)
}

var xxx_messageInfo_ConfiguredTarget proto.InternalMessageInfo

func (m *ConfiguredTarget) GetTarget() *blaze_query.Target {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *ConfiguredTarget) GetConfiguration() *Configuration {
	if m != nil {
		return m.Configuration
	}
	return nil
}

type CqueryResult struct {
	Results              []*ConfiguredTarget `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *CqueryResult) Reset()         { *m = CqueryResult{} }
func (m *CqueryResult) String() string { return proto.CompactTextString(m) }
func (*CqueryResult) ProtoMessage()    {}
func (*CqueryResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_b577e5ef8e9d5380, []int{10}
}

func (m *CqueryResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CqueryResult.Unmarshal(m, b)
}
func (m *CqueryResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CqueryResult.Marshal(b, m, deterministic)
}
func (m *CqueryResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CqueryResult.Merge(m, src)
}
func (m *CqueryResult) XXX_Size() int {
	return xxx_messageInfo_CqueryResult.Size(m)
}
func (m *CqueryResult) XXX_DiscardUnknown() {
	xxx_messageInfo_CqueryResult.DiscardUnknown(m)
}

var xxx_messageInfo_CqueryResult proto.InternalMessageInfo

func (m *CqueryResult) GetResults() []*ConfiguredTarget {
	if m != nil {
		return m.Results
	}
	return nil
}

type ParamFile struct {
	ExecPath             string   `protobuf:"bytes,1,opt,name=exec_path,json=execPath,proto3" json:"exec_path,omitempty"`
	Arguments            []string `protobuf:"bytes,2,rep,name=arguments,proto3" json:"arguments,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ParamFile) Reset()         { *m = ParamFile{} }
func (m *ParamFile) String() string { return proto.CompactTextString(m) }
func (*ParamFile) ProtoMessage()    {}
func (*ParamFile) Descriptor() ([]byte, []int) {
	return fileDescriptor_b577e5ef8e9d5380, []int{11}
}

func (m *ParamFile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ParamFile.Unmarshal(m, b)
}
func (m *ParamFile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ParamFile.Marshal(b, m, deterministic)
}
func (m *ParamFile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ParamFile.Merge(m, src)
}
func (m *ParamFile) XXX_Size() int {
	return xxx_messageInfo_ParamFile.Size(m)
}
func (m *ParamFile) XXX_DiscardUnknown() {
	xxx_messageInfo_ParamFile.DiscardUnknown(m)
}

var xxx_messageInfo_ParamFile proto.InternalMessageInfo

func (m *ParamFile) GetExecPath() string {
	if m != nil {
		return m.ExecPath
	}
	return ""
}

func (m *ParamFile) GetArguments() []string {
	if m != nil {
		return m.Arguments
	}
	return nil
}

func init() {
	proto.RegisterType((*ActionGraphContainer)(nil), "analysis.ActionGraphContainer")
	proto.RegisterType((*Artifact)(nil), "analysis.Artifact")
	proto.RegisterType((*Action)(nil), "analysis.Action")
	proto.RegisterType((*Target)(nil), "analysis.Target")
	proto.RegisterType((*RuleClass)(nil), "analysis.RuleClass")
	proto.RegisterType((*AspectDescriptor)(nil), "analysis.AspectDescriptor")
	proto.RegisterType((*DepSetOfFiles)(nil), "analysis.DepSetOfFiles")
	proto.RegisterType((*Configuration)(nil), "analysis.Configuration")
	proto.RegisterType((*KeyValuePair)(nil), "analysis.KeyValuePair")
	proto.RegisterType((*ConfiguredTarget)(nil), "analysis.ConfiguredTarget")
	proto.RegisterType((*CqueryResult)(nil), "analysis.CqueryResult")
	proto.RegisterType((*ParamFile)(nil), "analysis.ParamFile")
}

func init() {
	proto.RegisterFile("third_party/bazel/master/src/main/protobuf/analysis/analysis.proto", fileDescriptor_b577e5ef8e9d5380)
}

var fileDescriptor_b577e5ef8e9d5380 = []byte{
	// 919 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0x5f, 0x8b, 0xe3, 0x36,
	0x10, 0x27, 0xc9, 0x6e, 0x36, 0x99, 0x4d, 0x72, 0x59, 0xed, 0xde, 0xd5, 0xa4, 0x2d, 0x2c, 0x2e,
	0x94, 0xdd, 0x16, 0x92, 0x72, 0x77, 0x2c, 0x7d, 0x39, 0xe8, 0xfe, 0xe1, 0x4a, 0x58, 0x68, 0x17,
	0xf7, 0xb8, 0x57, 0xa3, 0xd8, 0x93, 0x44, 0x9c, 0x6d, 0xb9, 0x92, 0x1c, 0x9a, 0xa3, 0xf4, 0xa9,
	0x1f, 0xa1, 0x1f, 0xa1, 0x1f, 0xb4, 0x48, 0xb2, 0x1c, 0xc7, 0xa1, 0xe5, 0xfa, 0x26, 0xcd, 0xfc,
	0xe6, 0x8f, 0x66, 0x7e, 0x33, 0x82, 0x3b, 0xb5, 0x66, 0x22, 0x0e, 0x73, 0x2a, 0xd4, 0x76, 0xb6,
	0xa0, 0x1f, 0x31, 0x99, 0xa5, 0x54, 0x2a, 0x14, 0x33, 0x29, 0xa2, 0x59, 0x4a, 0x59, 0x36, 0xcb,
	0x05, 0x57, 0x7c, 0x51, 0x2c, 0x67, 0x34, 0xa3, 0xc9, 0x56, 0x32, 0x59, 0x1d, 0xa6, 0x46, 0x45,
	0x7a, 0xee, 0x3e, 0xf9, 0x3f, 0xde, 0x16, 0x09, 0xfd, 0x88, 0xe1, 0xaf, 0x05, 0x8a, 0xed, 0x6c,
	0x51, 0xb0, 0x24, 0xb6, 0xde, 0xfc, 0xbf, 0x3b, 0x70, 0x71, 0x1b, 0x29, 0xc6, 0xb3, 0x1f, 0x05,
	0xcd, 0xd7, 0xf7, 0x3c, 0x53, 0x94, 0x65, 0x28, 0xc8, 0x77, 0xd0, 0xa7, 0x42, 0xb1, 0x25, 0x8d,
	0x94, 0xf4, 0x5a, 0x97, 0x9d, 0xab, 0xd3, 0x97, 0x64, 0x5a, 0xa5, 0x72, 0x5b, 0xaa, 0x82, 0x1d,
	0x88, 0x7c, 0x03, 0x27, 0xd4, 0x78, 0x92, 0x5e, 0xdb, 0xe0, 0xc7, 0x35, 0xbc, 0x51, 0x04, 0x0e,
	0xa0, 0xb1, 0x8a, 0x8a, 0x15, 0x2a, 0xe9, 0x75, 0x9a, 0xd8, 0x77, 0x46, 0x11, 0x38, 0x00, 0xf9,
	0x01, 0xc6, 0x31, 0xe6, 0xa1, 0x44, 0x15, 0xf2, 0x65, 0xb8, 0x64, 0x09, 0x4a, 0xef, 0xc8, 0x18,
	0x7d, 0xb6, 0x33, 0x7a, 0xc0, 0xfc, 0x17, 0x54, 0x3f, 0x2f, 0xdf, 0x6a, 0x75, 0x30, 0x8c, 0xeb,
	0x57, 0xf2, 0x06, 0x86, 0x11, 0xcf, 0x96, 0x6c, 0x55, 0x08, 0xaa, 0xe3, 0x7b, 0xc7, 0x4d, 0xf3,
	0xfb, 0xba, 0x3a, 0xd8, 0x47, 0x93, 0x39, 0x10, 0x2a, 0x73, 0x8c, 0x54, 0x18, 0xa3, 0x8c, 0x04,
	0xcb, 0x15, 0x17, 0xd2, 0xeb, 0x1a, 0x1f, 0x93, 0xda, 0x1b, 0x0d, 0xe6, 0xa1, 0x82, 0x04, 0x67,
	0xb4, 0x21, 0x91, 0xe4, 0x06, 0x06, 0xa2, 0x48, 0x30, 0x8c, 0x12, 0x2a, 0x25, 0x4a, 0xef, 0xc4,
	0x38, 0x39, 0xdf, 0x39, 0x09, 0x8a, 0x04, 0xef, 0xb5, 0x32, 0x38, 0x15, 0xee, 0x88, 0xd2, 0xa7,
	0xd0, 0x73, 0x25, 0x27, 0x23, 0x68, 0xb3, 0xd8, 0x6b, 0x5d, 0xb6, 0xae, 0xfa, 0x41, 0x9b, 0xc5,
	0xe4, 0x73, 0xe8, 0xe3, 0x6f, 0x18, 0x85, 0x39, 0x55, 0x6b, 0xaf, 0x6d, 0xc4, 0x3d, 0x2d, 0x78,
	0xa2, 0x6a, 0x4d, 0xae, 0x60, 0xcc, 0x64, 0xa8, 0x04, 0x62, 0xe8, 0x3a, 0xe5, 0x75, 0x2e, 0x5b,
	0x57, 0xbd, 0x60, 0xc4, 0xe4, 0x3b, 0x81, 0xe8, 0xdc, 0xfa, 0x7f, 0x1d, 0x41, 0xd7, 0xb6, 0x49,
	0x7b, 0xb4, 0xc5, 0x0f, 0xab, 0x40, 0x3d, 0x2b, 0x98, 0xc7, 0xe4, 0x25, 0x3c, 0x3f, 0xa8, 0x46,
	0xc8, 0x62, 0xdb, 0xf4, 0x7e, 0x70, 0xde, 0x7c, 0xf4, 0x3c, 0x96, 0xe4, 0x4b, 0x00, 0xdb, 0xf9,
	0xf0, 0x03, 0x6e, 0x4d, 0xfc, 0x7e, 0xd0, 0xb7, 0x92, 0x47, 0xdc, 0x92, 0x09, 0xf4, 0xd2, 0x0c,
	0x53, 0x9e, 0xb1, 0xc8, 0x3b, 0xb2, 0xe1, 0xdc, 0x9d, 0x5c, 0xc3, 0x78, 0xaf, 0x1b, 0x3a, 0xa5,
	0x63, 0x83, 0x79, 0xb6, 0x27, 0x9f, 0xc7, 0xe4, 0x0b, 0x4d, 0xd9, 0x55, 0x91, 0x62, 0xa6, 0x6c,
	0x7b, 0x74, 0x10, 0x27, 0x20, 0x8f, 0xf0, 0x1c, 0xb3, 0x0d, 0x13, 0x3c, 0xd3, 0xf7, 0x70, 0x43,
	0x05, 0xa3, 0x8b, 0xa4, 0xea, 0xc1, 0x8b, 0x5d, 0x0f, 0x1e, 0x71, 0xfb, 0x9e, 0x26, 0x05, 0x3e,
	0x51, 0x26, 0x82, 0x8b, 0x9a, 0xd1, 0x7b, 0x67, 0x43, 0xae, 0xe1, 0x8c, 0x65, 0x79, 0xa1, 0x6b,
	0x60, 0x99, 0xa9, 0x0b, 0xd0, 0x33, 0x21, 0x47, 0x46, 0x61, 0xf9, 0x58, 0xbe, 0x9d, 0x17, 0x4a,
	0x63, 0x35, 0xa6, 0x6f, 0xd3, 0xb2, 0x12, 0xad, 0xbe, 0x86, 0x71, 0xcc, 0x64, 0xc4, 0x37, 0x28,
	0x64, 0x68, 0x4c, 0xa5, 0x07, 0xa6, 0x41, 0xcf, 0x2a, 0xf9, 0xdc, 0x88, 0xc9, 0x1b, 0x18, 0xe9,
	0xbe, 0x16, 0xb6, 0x0c, 0xd9, 0x92, 0x7b, 0xa7, 0xff, 0x99, 0xfa, 0xb0, 0x42, 0xcf, 0xb3, 0x25,
	0x27, 0xaf, 0xe1, 0x34, 0xa7, 0x82, 0xa6, 0xe5, 0x08, 0x0d, 0x9a, 0xd4, 0x7b, 0xd2, 0x4a, 0x3d,
	0x30, 0x01, 0xe4, 0xee, 0x28, 0xfd, 0x00, 0xba, 0x76, 0x20, 0x0f, 0x78, 0x77, 0x01, 0xc7, 0x09,
	0x5d, 0x60, 0x52, 0x72, 0xce, 0x5e, 0x88, 0x0f, 0xc3, 0x1d, 0xc3, 0x75, 0xb3, 0x6c, 0xb7, 0x77,
	0x6c, 0x9e, 0xc7, 0xfe, 0x0c, 0xfa, 0x15, 0xcf, 0x0f, 0xdc, 0x12, 0x38, 0xca, 0x68, 0x8a, 0xa5,
	0x57, 0x73, 0xf6, 0x33, 0x18, 0x37, 0xa7, 0xeb, 0x53, 0xec, 0xc8, 0x0d, 0xd8, 0xa7, 0xa0, 0x42,
	0xe1, 0x36, 0xcd, 0xbf, 0x55, 0xab, 0x86, 0xf4, 0xff, 0x6c, 0xc1, 0x70, 0x6f, 0xa3, 0x1c, 0x44,
	0x7b, 0x05, 0x2f, 0x94, 0xa0, 0x99, 0x64, 0x8a, 0x6d, 0x70, 0x8f, 0x05, 0xe5, 0x18, 0xec, 0xb4,
	0x3b, 0x2a, 0x4c, 0xe1, 0x3c, 0x66, 0x42, 0x8f, 0x8e, 0x9b, 0x45, 0x63, 0xd1, 0x31, 0x16, 0x67,
	0x56, 0xe5, 0xe6, 0x71, 0x1e, 0x4b, 0xff, 0x77, 0x18, 0xee, 0x2d, 0xa6, 0x83, 0x2c, 0xea, 0x83,
	0xd3, 0x6e, 0x0c, 0xce, 0x57, 0x30, 0xcc, 0x13, 0xaa, 0x96, 0x5c, 0xa4, 0xa1, 0x29, 0x8c, 0x6d,
	0xc4, 0xc0, 0x09, 0x7f, 0xd2, 0x05, 0x9a, 0x40, 0x2f, 0x5a, 0x63, 0xf4, 0x41, 0x16, 0xa9, 0x9b,
	0x3c, 0x77, 0xf7, 0x6f, 0x60, 0x50, 0x2f, 0x10, 0x19, 0x43, 0x47, 0x4f, 0xaf, 0x8d, 0xae, 0x8f,
	0x9a, 0x01, 0x1b, 0xad, 0x76, 0x0c, 0x30, 0x17, 0xff, 0x0f, 0x18, 0xbb, 0xac, 0x31, 0x2e, 0xb9,
	0xf3, 0x2d, 0x74, 0xed, 0x02, 0x31, 0xe6, 0x9a, 0x76, 0xb5, 0x0f, 0xc9, 0x6d, 0xfc, 0x12, 0x72,
	0xb8, 0xae, 0xdb, 0xc6, 0xe6, 0x13, 0xd7, 0xb5, 0xff, 0x00, 0x83, 0x7b, 0xe3, 0x37, 0x40, 0x59,
	0x24, 0x8a, 0xbc, 0x86, 0x13, 0x61, 0x4e, 0xee, 0x1f, 0x9b, 0x1c, 0x3a, 0x72, 0x89, 0x06, 0x0e,
	0xea, 0xbf, 0x85, 0x7e, 0x35, 0x10, 0xfb, 0x2b, 0xb6, 0xd5, 0x58, 0xb1, 0x7b, 0x6b, 0xa7, 0xdd,
	0x58, 0x3b, 0x77, 0xdf, 0xc3, 0xd7, 0x11, 0x4f, 0xa7, 0x2b, 0xce, 0x57, 0x09, 0x4e, 0x63, 0xdc,
	0x28, 0xce, 0x13, 0x39, 0xb5, 0x7f, 0x70, 0xc2, 0x16, 0x55, 0x2e, 0x77, 0xa3, 0xdb, 0xf2, 0xf4,
	0xa4, 0x7f, 0x66, 0xb9, 0xe8, 0x9a, 0x1f, 0xfa, 0xd5, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xc0,
	0xc3, 0x2c, 0x88, 0x35, 0x08, 0x00, 0x00,
}