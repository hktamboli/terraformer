// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glue

type CatalogEncryptionMode string

// Enum values for CatalogEncryptionMode
const (
	CatalogEncryptionModeDisabled CatalogEncryptionMode = "DISABLED"
	CatalogEncryptionModeSseKms   CatalogEncryptionMode = "SSE-KMS"
)

func (enum CatalogEncryptionMode) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum CatalogEncryptionMode) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type CloudWatchEncryptionMode string

// Enum values for CloudWatchEncryptionMode
const (
	CloudWatchEncryptionModeDisabled CloudWatchEncryptionMode = "DISABLED"
	CloudWatchEncryptionModeSseKms   CloudWatchEncryptionMode = "SSE-KMS"
)

func (enum CloudWatchEncryptionMode) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum CloudWatchEncryptionMode) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type Comparator string

// Enum values for Comparator
const (
	ComparatorEquals            Comparator = "EQUALS"
	ComparatorGreaterThan       Comparator = "GREATER_THAN"
	ComparatorLessThan          Comparator = "LESS_THAN"
	ComparatorGreaterThanEquals Comparator = "GREATER_THAN_EQUALS"
	ComparatorLessThanEquals    Comparator = "LESS_THAN_EQUALS"
)

func (enum Comparator) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum Comparator) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ConnectionPropertyKey string

// Enum values for ConnectionPropertyKey
const (
	ConnectionPropertyKeyHost                         ConnectionPropertyKey = "HOST"
	ConnectionPropertyKeyPort                         ConnectionPropertyKey = "PORT"
	ConnectionPropertyKeyUsername                     ConnectionPropertyKey = "USERNAME"
	ConnectionPropertyKeyPassword                     ConnectionPropertyKey = "PASSWORD"
	ConnectionPropertyKeyEncryptedPassword            ConnectionPropertyKey = "ENCRYPTED_PASSWORD"
	ConnectionPropertyKeyJdbcDriverJarUri             ConnectionPropertyKey = "JDBC_DRIVER_JAR_URI"
	ConnectionPropertyKeyJdbcDriverClassName          ConnectionPropertyKey = "JDBC_DRIVER_CLASS_NAME"
	ConnectionPropertyKeyJdbcEngine                   ConnectionPropertyKey = "JDBC_ENGINE"
	ConnectionPropertyKeyJdbcEngineVersion            ConnectionPropertyKey = "JDBC_ENGINE_VERSION"
	ConnectionPropertyKeyConfigFiles                  ConnectionPropertyKey = "CONFIG_FILES"
	ConnectionPropertyKeyInstanceId                   ConnectionPropertyKey = "INSTANCE_ID"
	ConnectionPropertyKeyJdbcConnectionUrl            ConnectionPropertyKey = "JDBC_CONNECTION_URL"
	ConnectionPropertyKeyJdbcEnforceSsl               ConnectionPropertyKey = "JDBC_ENFORCE_SSL"
	ConnectionPropertyKeyCustomJdbcCert               ConnectionPropertyKey = "CUSTOM_JDBC_CERT"
	ConnectionPropertyKeySkipCustomJdbcCertValidation ConnectionPropertyKey = "SKIP_CUSTOM_JDBC_CERT_VALIDATION"
	ConnectionPropertyKeyCustomJdbcCertString         ConnectionPropertyKey = "CUSTOM_JDBC_CERT_STRING"
	ConnectionPropertyKeyConnectionUrl                ConnectionPropertyKey = "CONNECTION_URL"
	ConnectionPropertyKeyKafkaBootstrapServers        ConnectionPropertyKey = "KAFKA_BOOTSTRAP_SERVERS"
)

func (enum ConnectionPropertyKey) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ConnectionPropertyKey) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ConnectionType string

// Enum values for ConnectionType
const (
	ConnectionTypeJdbc    ConnectionType = "JDBC"
	ConnectionTypeSftp    ConnectionType = "SFTP"
	ConnectionTypeMongodb ConnectionType = "MONGODB"
	ConnectionTypeKafka   ConnectionType = "KAFKA"
)

func (enum ConnectionType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ConnectionType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type CrawlState string

// Enum values for CrawlState
const (
	CrawlStateRunning   CrawlState = "RUNNING"
	CrawlStateSucceeded CrawlState = "SUCCEEDED"
	CrawlStateCancelled CrawlState = "CANCELLED"
	CrawlStateFailed    CrawlState = "FAILED"
)

func (enum CrawlState) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum CrawlState) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type CrawlerState string

// Enum values for CrawlerState
const (
	CrawlerStateReady    CrawlerState = "READY"
	CrawlerStateRunning  CrawlerState = "RUNNING"
	CrawlerStateStopping CrawlerState = "STOPPING"
)

func (enum CrawlerState) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum CrawlerState) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type CsvHeaderOption string

// Enum values for CsvHeaderOption
const (
	CsvHeaderOptionUnknown CsvHeaderOption = "UNKNOWN"
	CsvHeaderOptionPresent CsvHeaderOption = "PRESENT"
	CsvHeaderOptionAbsent  CsvHeaderOption = "ABSENT"
)

func (enum CsvHeaderOption) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum CsvHeaderOption) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type DeleteBehavior string

// Enum values for DeleteBehavior
const (
	DeleteBehaviorLog                 DeleteBehavior = "LOG"
	DeleteBehaviorDeleteFromDatabase  DeleteBehavior = "DELETE_FROM_DATABASE"
	DeleteBehaviorDeprecateInDatabase DeleteBehavior = "DEPRECATE_IN_DATABASE"
)

func (enum DeleteBehavior) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum DeleteBehavior) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ExistCondition string

// Enum values for ExistCondition
const (
	ExistConditionMustExist ExistCondition = "MUST_EXIST"
	ExistConditionNotExist  ExistCondition = "NOT_EXIST"
	ExistConditionNone      ExistCondition = "NONE"
)

func (enum ExistCondition) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ExistCondition) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type JobBookmarksEncryptionMode string

// Enum values for JobBookmarksEncryptionMode
const (
	JobBookmarksEncryptionModeDisabled JobBookmarksEncryptionMode = "DISABLED"
	JobBookmarksEncryptionModeCseKms   JobBookmarksEncryptionMode = "CSE-KMS"
)

func (enum JobBookmarksEncryptionMode) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum JobBookmarksEncryptionMode) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type JobRunState string

// Enum values for JobRunState
const (
	JobRunStateStarting  JobRunState = "STARTING"
	JobRunStateRunning   JobRunState = "RUNNING"
	JobRunStateStopping  JobRunState = "STOPPING"
	JobRunStateStopped   JobRunState = "STOPPED"
	JobRunStateSucceeded JobRunState = "SUCCEEDED"
	JobRunStateFailed    JobRunState = "FAILED"
	JobRunStateTimeout   JobRunState = "TIMEOUT"
)

func (enum JobRunState) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum JobRunState) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type Language string

// Enum values for Language
const (
	LanguagePython Language = "PYTHON"
	LanguageScala  Language = "SCALA"
)

func (enum Language) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum Language) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type LastCrawlStatus string

// Enum values for LastCrawlStatus
const (
	LastCrawlStatusSucceeded LastCrawlStatus = "SUCCEEDED"
	LastCrawlStatusCancelled LastCrawlStatus = "CANCELLED"
	LastCrawlStatusFailed    LastCrawlStatus = "FAILED"
)

func (enum LastCrawlStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum LastCrawlStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type Logical string

// Enum values for Logical
const (
	LogicalAnd Logical = "AND"
	LogicalAny Logical = "ANY"
)

func (enum Logical) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum Logical) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type LogicalOperator string

// Enum values for LogicalOperator
const (
	LogicalOperatorEquals LogicalOperator = "EQUALS"
)

func (enum LogicalOperator) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum LogicalOperator) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type NodeType string

// Enum values for NodeType
const (
	NodeTypeCrawler NodeType = "CRAWLER"
	NodeTypeJob     NodeType = "JOB"
	NodeTypeTrigger NodeType = "TRIGGER"
)

func (enum NodeType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum NodeType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type Permission string

// Enum values for Permission
const (
	PermissionAll                Permission = "ALL"
	PermissionSelect             Permission = "SELECT"
	PermissionAlter              Permission = "ALTER"
	PermissionDrop               Permission = "DROP"
	PermissionDelete             Permission = "DELETE"
	PermissionInsert             Permission = "INSERT"
	PermissionCreateDatabase     Permission = "CREATE_DATABASE"
	PermissionCreateTable        Permission = "CREATE_TABLE"
	PermissionDataLocationAccess Permission = "DATA_LOCATION_ACCESS"
)

func (enum Permission) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum Permission) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type PrincipalType string

// Enum values for PrincipalType
const (
	PrincipalTypeUser  PrincipalType = "USER"
	PrincipalTypeRole  PrincipalType = "ROLE"
	PrincipalTypeGroup PrincipalType = "GROUP"
)

func (enum PrincipalType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum PrincipalType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ResourceType string

// Enum values for ResourceType
const (
	ResourceTypeJar     ResourceType = "JAR"
	ResourceTypeFile    ResourceType = "FILE"
	ResourceTypeArchive ResourceType = "ARCHIVE"
)

func (enum ResourceType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ResourceType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type S3EncryptionMode string

// Enum values for S3EncryptionMode
const (
	S3EncryptionModeDisabled S3EncryptionMode = "DISABLED"
	S3EncryptionModeSseKms   S3EncryptionMode = "SSE-KMS"
	S3EncryptionModeSseS3    S3EncryptionMode = "SSE-S3"
)

func (enum S3EncryptionMode) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum S3EncryptionMode) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ScheduleState string

// Enum values for ScheduleState
const (
	ScheduleStateScheduled     ScheduleState = "SCHEDULED"
	ScheduleStateNotScheduled  ScheduleState = "NOT_SCHEDULED"
	ScheduleStateTransitioning ScheduleState = "TRANSITIONING"
)

func (enum ScheduleState) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ScheduleState) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type Sort string

// Enum values for Sort
const (
	SortAsc  Sort = "ASC"
	SortDesc Sort = "DESC"
)

func (enum Sort) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum Sort) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type SortDirectionType string

// Enum values for SortDirectionType
const (
	SortDirectionTypeDescending SortDirectionType = "DESCENDING"
	SortDirectionTypeAscending  SortDirectionType = "ASCENDING"
)

func (enum SortDirectionType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum SortDirectionType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type TaskRunSortColumnType string

// Enum values for TaskRunSortColumnType
const (
	TaskRunSortColumnTypeTaskRunType TaskRunSortColumnType = "TASK_RUN_TYPE"
	TaskRunSortColumnTypeStatus      TaskRunSortColumnType = "STATUS"
	TaskRunSortColumnTypeStarted     TaskRunSortColumnType = "STARTED"
)

func (enum TaskRunSortColumnType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum TaskRunSortColumnType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type TaskStatusType string

// Enum values for TaskStatusType
const (
	TaskStatusTypeStarting  TaskStatusType = "STARTING"
	TaskStatusTypeRunning   TaskStatusType = "RUNNING"
	TaskStatusTypeStopping  TaskStatusType = "STOPPING"
	TaskStatusTypeStopped   TaskStatusType = "STOPPED"
	TaskStatusTypeSucceeded TaskStatusType = "SUCCEEDED"
	TaskStatusTypeFailed    TaskStatusType = "FAILED"
	TaskStatusTypeTimeout   TaskStatusType = "TIMEOUT"
)

func (enum TaskStatusType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum TaskStatusType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type TaskType string

// Enum values for TaskType
const (
	TaskTypeEvaluation            TaskType = "EVALUATION"
	TaskTypeLabelingSetGeneration TaskType = "LABELING_SET_GENERATION"
	TaskTypeImportLabels          TaskType = "IMPORT_LABELS"
	TaskTypeExportLabels          TaskType = "EXPORT_LABELS"
	TaskTypeFindMatches           TaskType = "FIND_MATCHES"
)

func (enum TaskType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum TaskType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type TransformSortColumnType string

// Enum values for TransformSortColumnType
const (
	TransformSortColumnTypeName          TransformSortColumnType = "NAME"
	TransformSortColumnTypeTransformType TransformSortColumnType = "TRANSFORM_TYPE"
	TransformSortColumnTypeStatus        TransformSortColumnType = "STATUS"
	TransformSortColumnTypeCreated       TransformSortColumnType = "CREATED"
	TransformSortColumnTypeLastModified  TransformSortColumnType = "LAST_MODIFIED"
)

func (enum TransformSortColumnType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum TransformSortColumnType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type TransformStatusType string

// Enum values for TransformStatusType
const (
	TransformStatusTypeNotReady TransformStatusType = "NOT_READY"
	TransformStatusTypeReady    TransformStatusType = "READY"
	TransformStatusTypeDeleting TransformStatusType = "DELETING"
)

func (enum TransformStatusType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum TransformStatusType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type TransformType string

// Enum values for TransformType
const (
	TransformTypeFindMatches TransformType = "FIND_MATCHES"
)

func (enum TransformType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum TransformType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type TriggerState string

// Enum values for TriggerState
const (
	TriggerStateCreating     TriggerState = "CREATING"
	TriggerStateCreated      TriggerState = "CREATED"
	TriggerStateActivating   TriggerState = "ACTIVATING"
	TriggerStateActivated    TriggerState = "ACTIVATED"
	TriggerStateDeactivating TriggerState = "DEACTIVATING"
	TriggerStateDeactivated  TriggerState = "DEACTIVATED"
	TriggerStateDeleting     TriggerState = "DELETING"
	TriggerStateUpdating     TriggerState = "UPDATING"
)

func (enum TriggerState) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum TriggerState) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type TriggerType string

// Enum values for TriggerType
const (
	TriggerTypeScheduled   TriggerType = "SCHEDULED"
	TriggerTypeConditional TriggerType = "CONDITIONAL"
	TriggerTypeOnDemand    TriggerType = "ON_DEMAND"
)

func (enum TriggerType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum TriggerType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type UpdateBehavior string

// Enum values for UpdateBehavior
const (
	UpdateBehaviorLog              UpdateBehavior = "LOG"
	UpdateBehaviorUpdateInDatabase UpdateBehavior = "UPDATE_IN_DATABASE"
)

func (enum UpdateBehavior) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum UpdateBehavior) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type WorkerType string

// Enum values for WorkerType
const (
	WorkerTypeStandard WorkerType = "Standard"
	WorkerTypeG1x      WorkerType = "G.1X"
	WorkerTypeG2x      WorkerType = "G.2X"
)

func (enum WorkerType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum WorkerType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type WorkflowRunStatus string

// Enum values for WorkflowRunStatus
const (
	WorkflowRunStatusRunning   WorkflowRunStatus = "RUNNING"
	WorkflowRunStatusCompleted WorkflowRunStatus = "COMPLETED"
)

func (enum WorkflowRunStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum WorkflowRunStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}
