// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codecommit

type ApprovalState string

// Enum values for ApprovalState
const (
	ApprovalStateApprove ApprovalState = "APPROVE"
	ApprovalStateRevoke  ApprovalState = "REVOKE"
)

func (enum ApprovalState) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ApprovalState) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ChangeTypeEnum string

// Enum values for ChangeTypeEnum
const (
	ChangeTypeEnumA ChangeTypeEnum = "A"
	ChangeTypeEnumM ChangeTypeEnum = "M"
	ChangeTypeEnumD ChangeTypeEnum = "D"
)

func (enum ChangeTypeEnum) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ChangeTypeEnum) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ConflictDetailLevelTypeEnum string

// Enum values for ConflictDetailLevelTypeEnum
const (
	ConflictDetailLevelTypeEnumFileLevel ConflictDetailLevelTypeEnum = "FILE_LEVEL"
	ConflictDetailLevelTypeEnumLineLevel ConflictDetailLevelTypeEnum = "LINE_LEVEL"
)

func (enum ConflictDetailLevelTypeEnum) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ConflictDetailLevelTypeEnum) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ConflictResolutionStrategyTypeEnum string

// Enum values for ConflictResolutionStrategyTypeEnum
const (
	ConflictResolutionStrategyTypeEnumNone              ConflictResolutionStrategyTypeEnum = "NONE"
	ConflictResolutionStrategyTypeEnumAcceptSource      ConflictResolutionStrategyTypeEnum = "ACCEPT_SOURCE"
	ConflictResolutionStrategyTypeEnumAcceptDestination ConflictResolutionStrategyTypeEnum = "ACCEPT_DESTINATION"
	ConflictResolutionStrategyTypeEnumAutomerge         ConflictResolutionStrategyTypeEnum = "AUTOMERGE"
)

func (enum ConflictResolutionStrategyTypeEnum) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ConflictResolutionStrategyTypeEnum) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type FileModeTypeEnum string

// Enum values for FileModeTypeEnum
const (
	FileModeTypeEnumExecutable FileModeTypeEnum = "EXECUTABLE"
	FileModeTypeEnumNormal     FileModeTypeEnum = "NORMAL"
	FileModeTypeEnumSymlink    FileModeTypeEnum = "SYMLINK"
)

func (enum FileModeTypeEnum) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum FileModeTypeEnum) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type MergeOptionTypeEnum string

// Enum values for MergeOptionTypeEnum
const (
	MergeOptionTypeEnumFastForwardMerge MergeOptionTypeEnum = "FAST_FORWARD_MERGE"
	MergeOptionTypeEnumSquashMerge      MergeOptionTypeEnum = "SQUASH_MERGE"
	MergeOptionTypeEnumThreeWayMerge    MergeOptionTypeEnum = "THREE_WAY_MERGE"
)

func (enum MergeOptionTypeEnum) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum MergeOptionTypeEnum) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ObjectTypeEnum string

// Enum values for ObjectTypeEnum
const (
	ObjectTypeEnumFile         ObjectTypeEnum = "FILE"
	ObjectTypeEnumDirectory    ObjectTypeEnum = "DIRECTORY"
	ObjectTypeEnumGitLink      ObjectTypeEnum = "GIT_LINK"
	ObjectTypeEnumSymbolicLink ObjectTypeEnum = "SYMBOLIC_LINK"
)

func (enum ObjectTypeEnum) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ObjectTypeEnum) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type OrderEnum string

// Enum values for OrderEnum
const (
	OrderEnumAscending  OrderEnum = "ascending"
	OrderEnumDescending OrderEnum = "descending"
)

func (enum OrderEnum) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum OrderEnum) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type OverrideStatus string

// Enum values for OverrideStatus
const (
	OverrideStatusOverride OverrideStatus = "OVERRIDE"
	OverrideStatusRevoke   OverrideStatus = "REVOKE"
)

func (enum OverrideStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum OverrideStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type PullRequestEventType string

// Enum values for PullRequestEventType
const (
	PullRequestEventTypePullRequestCreated                PullRequestEventType = "PULL_REQUEST_CREATED"
	PullRequestEventTypePullRequestStatusChanged          PullRequestEventType = "PULL_REQUEST_STATUS_CHANGED"
	PullRequestEventTypePullRequestSourceReferenceUpdated PullRequestEventType = "PULL_REQUEST_SOURCE_REFERENCE_UPDATED"
	PullRequestEventTypePullRequestMergeStateChanged      PullRequestEventType = "PULL_REQUEST_MERGE_STATE_CHANGED"
	PullRequestEventTypePullRequestApprovalRuleCreated    PullRequestEventType = "PULL_REQUEST_APPROVAL_RULE_CREATED"
	PullRequestEventTypePullRequestApprovalRuleUpdated    PullRequestEventType = "PULL_REQUEST_APPROVAL_RULE_UPDATED"
	PullRequestEventTypePullRequestApprovalRuleDeleted    PullRequestEventType = "PULL_REQUEST_APPROVAL_RULE_DELETED"
	PullRequestEventTypePullRequestApprovalRuleOverridden PullRequestEventType = "PULL_REQUEST_APPROVAL_RULE_OVERRIDDEN"
	PullRequestEventTypePullRequestApprovalStateChanged   PullRequestEventType = "PULL_REQUEST_APPROVAL_STATE_CHANGED"
)

func (enum PullRequestEventType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum PullRequestEventType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type PullRequestStatusEnum string

// Enum values for PullRequestStatusEnum
const (
	PullRequestStatusEnumOpen   PullRequestStatusEnum = "OPEN"
	PullRequestStatusEnumClosed PullRequestStatusEnum = "CLOSED"
)

func (enum PullRequestStatusEnum) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum PullRequestStatusEnum) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type RelativeFileVersionEnum string

// Enum values for RelativeFileVersionEnum
const (
	RelativeFileVersionEnumBefore RelativeFileVersionEnum = "BEFORE"
	RelativeFileVersionEnumAfter  RelativeFileVersionEnum = "AFTER"
)

func (enum RelativeFileVersionEnum) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum RelativeFileVersionEnum) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ReplacementTypeEnum string

// Enum values for ReplacementTypeEnum
const (
	ReplacementTypeEnumKeepBase        ReplacementTypeEnum = "KEEP_BASE"
	ReplacementTypeEnumKeepSource      ReplacementTypeEnum = "KEEP_SOURCE"
	ReplacementTypeEnumKeepDestination ReplacementTypeEnum = "KEEP_DESTINATION"
	ReplacementTypeEnumUseNewContent   ReplacementTypeEnum = "USE_NEW_CONTENT"
)

func (enum ReplacementTypeEnum) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ReplacementTypeEnum) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type RepositoryTriggerEventEnum string

// Enum values for RepositoryTriggerEventEnum
const (
	RepositoryTriggerEventEnumAll             RepositoryTriggerEventEnum = "all"
	RepositoryTriggerEventEnumUpdateReference RepositoryTriggerEventEnum = "updateReference"
	RepositoryTriggerEventEnumCreateReference RepositoryTriggerEventEnum = "createReference"
	RepositoryTriggerEventEnumDeleteReference RepositoryTriggerEventEnum = "deleteReference"
)

func (enum RepositoryTriggerEventEnum) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum RepositoryTriggerEventEnum) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type SortByEnum string

// Enum values for SortByEnum
const (
	SortByEnumRepositoryName   SortByEnum = "repositoryName"
	SortByEnumLastModifiedDate SortByEnum = "lastModifiedDate"
)

func (enum SortByEnum) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum SortByEnum) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}