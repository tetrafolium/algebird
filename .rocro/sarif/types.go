package sarif

// Static Analysis Results Interchange Format (SARIF) Version 2.1.0
// - https://docs.oasis-open.org/sarif/sarif/v2.1.0/os/sarif-v2.1.0-os.html

const (
	VERSION = "2.1.0"
	SCHEMA  = "https://raw.githubusercontent.com/oasis-tcs/sarif-spec/master/Schemata/sarif-external-property-file-schema-2.1.0.json"
)

// ArtifactContent ... 3.3 artifactContent object
type ArtifactContent struct {
	Text       string                    `json:"text,omitempty"`
	Binary     string                    `json:"binary,omitempty"` // MIME Base64 encoding
	Rendered   *MultiformatMessageString `json:"rendered,omitempty"`
	Properties PropertyBag               `json:"properties,omitempty"`
}

// ArtifactLocation ... 3.4 artifactLocation object
type ArtifactLocation struct {
	URI         URIString   `json:"uri,omitempty"`
	URIBaseID   string      `json:"uriBaseId,omitempty"`
	Index       int         `json:"index,omitempty"`
	Description *Message    `json:"description,omitempty"`
	Properties  PropertyBag `json:"properties,omitempty"`
}

// GUIDString ... 3.5.3 GUID-valued strings
type GUIDString = string

// PropertyBag ... 3.8 Property bags
type PropertyBag map[string]interface{}

// DateTimeString ... 3.9 Date/time properties
type DateTimeString = string

// URIString ... 3.10 URI-valued properties
type URIString = string

// Message ... 3.11 message object
type Message struct {
	Text       string      `json:"text,omitempty"`
	Markdown   string      `json:"markdown,omitempty"`
	ID         string      `json:"id,omitempty"`
	Arguments  []string    `json:"arguments,omitempty"`
	Properties PropertyBag `json:"properties,omitempty"`
}

// MultiformatMessageString ... 3.12 multiformatMessageString object
type MultiformatMessageString struct {
	Text       string      `json:"text,omitempty"`
	Markdown   string      `json:"markdown,omitempty"`
	Properties PropertyBag `json:"properties,omitempty"`
}

// Log ... 3.13 sarifLog object
type Log struct {
	Version                  string              `json:"version"`
	Schema                   URIString           `json:"$schema,omitempty"`
	Runs                     []Run               `json:"runs"`
	InlineExternalProperties *ExternalProperties `json:"inlineExternalProperties,omitempty"`
	Properties               PropertyBag         `json:"properties,omitempty"`
}

// Run ... 3.14 run object
type Run struct {
	ExternalPropertyFileReferences *ExternalPropertyFileReferences `json:"externalPropertyFileReferences,omitempty"`
	AutomationDetails              *RunAutomationDetails           `json:"automationDetails,omitempty"`
	RunAggregates                  []RunAutomationDetails          `json:"runAggregates,omitempty"`
	BaselineGUID                   GUIDString                      `json:"baselineGuid,omitempty"`
	Tool                           *Tool                           `json:"tool"`
	Language                       string                          `json:"language,omitempty"`
	Taxonomies                     []ToolComponent                 `json:"taxonomies,omitempty"`
	Translations                   []ToolComponent                 `json:"translations,omitempty"`
	Policies                       []ToolComponent                 `json:"policies,omitempty"`
	Invocations                    []Invocation                    `json:"invocations,omitempty"`
	Conversion                     *Conversion                     `json:"conversion,omitempty"`
	VersionControlProvenance       []VersionControlDetails         `json:"versionControlProvenance,omitempty"`
	OriginalURIBaseIDs             map[string]ArtifactLocation     `json:"originalUriBaseIds,omitempty"`
	Artifacts                      []Artifact                      `json:"artifacts,omitempty"`
	SpecialLocations               *SpecialLocations               `json:"specialLocations,omitempty"`
	LogicalLocations               []LogicalLocation               `json:"logicalLocations,omitempty"`
	Addresses                      []Address                       `json:"addresses,omitempty"`
	ThreadFlowLocations            []ThreadFlowLocation            `json:"threadFlowLocations,omitempty"`
	Graphs                         []Graph                         `json:"graphs,omitempty"`
	WebRequests                    []WebRequest                    `json:"webRequests,omitempty"`
	WebResponses                   []WebResponse                   `json:"webResponses,omitempty"`
	Results                        *[]Result                       `json:"results,omitempty"` // (nullable)
	DefaultEncoding                string                          `json:"defaultEncoding,omitempty"`
	DefaultSourceLanguage          string                          `json:"defaultSourceLanguage,omitempty"`
	NewlineSequences               []string                        `json:"newlineSequences,omitempty"`
	ColumnKind                     string                          `json:"columnKind,omitempty"`
	RedactionTokens                []string                        `json:"redactionTokens,omitempty"`
	Properties                     PropertyBag                     `json:"properties,omitempty"`
}

// ExternalPropertyFileReferences ... 3.15 externalPropertyFileReferences object
type ExternalPropertyFileReferences struct {
	Taxonomies             []ToolComponent      `json:"taxonomies,omitempty"`
	Translations           []ToolComponent      `json:"translations,omitempty"`
	Policies               []ToolComponent      `json:"policies,omitempty"`
	Invocations            []Invocation         `json:"invocations,omitempty"`
	Conversion             *Conversion          `json:"conversion,omitempty"`
	Artifacts              []Artifact           `json:"artifacts,omitempty"`
	LogicalLocations       []LogicalLocation    `json:"logicalLocations,omitempty"`
	Addresses              []Address            `json:"addresses,omitempty"`
	ThreadFlowLocations    []ThreadFlowLocation `json:"threadFlowLocations,omitempty"`
	Graphs                 []Graph              `json:"graphs,omitempty"`
	WebRequests            []WebRequest         `json:"webRequests,omitempty"`
	WebResponses           []WebResponse        `json:"webResponses,omitempty"`
	Results                *[]Result            `json:"results,omitempty"` // (nullable)
	Driver                 *ToolComponent       `json:"driver,omitempty"`
	Extensions             []ToolComponent      `json:"extensions,omitempty"`
	ExternalizedProperties PropertyBag          `json:"externalizedProperties,omitempty"`
}

// ExternalPropertyFileReference ... 3.16 externalPropertyFileReference object
type ExternalPropertyFileReference struct {
	Location   *ArtifactLocation `json:"location,omitempty"`
	GUID       *GUIDString       `json:"guid,omitempty"`
	ItemCount  int               `json:"itemCount,omitempty"`
	Properties PropertyBag       `json:"properties,omitempty"`
}

// RunAutomationDetails ... 3.17 runAutomationDetails object
type RunAutomationDetails struct {
	Description     *Message    `json:"description"`
	ID              string      `json:"id,omitempty"`
	GUID            GUIDString  `json:"guid,omitempty"`
	CorrelationGUID GUIDString  `json:"correlationGuid,omitempty"`
	Properties      PropertyBag `json:"properties,omitempty"`
}

// Tool ... 3.18 tool object
type Tool struct {
	Driver     *ToolComponent  `json:"driver"`
	Extensions []ToolComponent `json:"extensions,omitempty"`
	Properties PropertyBag     `json:"properties,omitempty"`
}

// ToolComponent ... 3.19 toolComponent object
type ToolComponent struct {
	GUID                                        GUIDString                          `json:"guid,omitempty"`
	Name                                        string                              `json:"name"`
	FullName                                    string                              `json:"fullName,omitempty"`
	Product                                     string                              `json:"product,omitempty"`
	ProductSuite                                string                              `json:"productSuite,omitempty"`
	Version                                     string                              `json:"version,omitempty"`
	SemanticVersion                             string                              `json:"semanticVersion,omitempty"`
	DottedQuadFileVersion                       string                              `json:"dottedQuadFileVersion,omitempty"`
	ReleaseDateUTC                              DateTimeString                      `json:"releaseDateUtc,omitempty"`
	DownloadURI                                 string                              `json:"downloadUri,omitempty"`
	InformationURI                              string                              `json:"informationUri,omitempty"`
	Organization                                string                              `json:"organization,omitempty"`
	ShortDescription                            *MultiformatMessageString           `json:"shortDescription,omitempty"`
	FullDescription                             *MultiformatMessageString           `json:"fullDescription,omitempty"`
	Language                                    string                              `json:"language,omitempty"`
	GlobalMessageStrings                        map[string]MultiformatMessageString `json:"globalMessageStrings,omitempty"`
	Rules                                       []ReportingDescriptor               `json:"rules,omitempty"`
	Notifications                               []ReportingDescriptor               `json:"notifications,omitempty"`
	Taxa                                        []ReportingDescriptor               `json:"taxa,omitempty"`
	SupportedTaxonomies                         []ToolComponentReference            `json:"supportedTaxonomies,omitempty"`
	TranslationMetadata                         TranslationMetadata                 `json:"translationMetadata,omitempty"`
	Locations                                   []ArtifactLocation                  `json:"locations,omitempty"`
	Contents                                    []string                            `json:"contents"`
	IsComprehensive                             bool                                `json:"isComprehensive"`
	LocalizedDataSemanticVersion                string                              `json:"localizedDataSemanticVersion,omitempty"`
	MinimumRequiredLocalizedDataSemanticVersion string                              `json:"localiminimumRequiredLocalizedDataSemanticVersionzedDataSemanticVersion,omitempty"`
	AssociatedComponent                         *ToolComponentReference             `json:"associatedComponent,omitempty"`
	Properties                                  PropertyBag                         `json:"properties,omitempty"`
}

// Invocation ... 3.20 invocation object
type Invocation struct {
	CommandLine                        string                  `json:"commandLine,omitempty"`
	Arguments                          []string                `json:"arguments,omitempty"`
	ResponseFiles                      []ArtifactLocation      `json:"responseFiles,omitempty"`
	RuleConfigurationOverrides         []ConfigurationOverride `json:"ruleConfigurationOverrides,omitempty"`
	NotificationConfigurationOverrides []ConfigurationOverride `json:"notificationConfigurationOverrides,omitempty"`
	StartTimeUTC                       DateTimeString          `json:"startTimeUtc,omitempty"`
	EndTimeUTC                         DateTimeString          `json:"endTimeUtc,omitempty"`
	ExitCode                           int                     `json:"exitCode,omitempty"`
	ExitCodeDescription                string                  `json:"exitCodeDescription,omitempty"`
	ExitSignalName                     string                  `json:"exitSignalName,omitempty"`
	ExitSignalNumber                   int                     `json:"exitSignalNumber,omitempty"`
	ProcessStartFailureMessage         string                  `json:"processStartFailureMessage,omitempty"`
	ExecutionSuccessful                bool                    `json:"executionSuccessful"`
	Machine                            string                  `json:"machine,omitempty"`
	Account                            string                  `json:"account,omitempty"`
	ProcessID                          int                     `json:"processId,omitempty"`
	ExecutableLocation                 *ArtifactLocation       `json:"executableLocation,omitempty"`
	WorkingDirectory                   *ArtifactLocation       `json:"workingDirectory,omitempty"`
	EnvironmentVariables               map[string]string       `json:"environmentVariables,omitempty"`
	ToolExecutionNotifications         []Notification          `json:"toolExecutionNotifications,omitempty"`
	ToolConfigurationNotifications     []Notification          `json:"toolConfigurationNotifications,omitempty"`
	Stdin                              *ArtifactLocation       `json:"stdin,omitempty"`
	Stdout                             *ArtifactLocation       `json:"stdout,omitempty"`
	Stderr                             *ArtifactLocation       `json:"stderr,omitempty"`
	StdoutStderr                       *ArtifactLocation       `json:"stdoutStderr,omitempty"`
	Properties                         PropertyBag             `json:"properties,omitempty"`
}

// Attachment ... 3.21 attachment object
type Attachment struct {
	Description *Message          `json:"description"`
	Location    *ArtifactLocation `json:"location,omitempty"`
	Regions     []Region          `json:"regions,omitempty"`
	Rectangles  []Rectangle       `json:"rectangle,omitempty"`
	Properties  PropertyBag       `json:"properties,omitempty"`
}

// Conversion ... 3.22 conversion object
type Conversion struct {
	Tool                 *Tool              `json:"tool"`
	Invocation           *Invocation        `json:"invocation,omitempty"`
	AnalysisToolLogFiles []ArtifactLocation `json:"analysisToolLogFiles,omitempty"`
	Properties           PropertyBag        `json:"properties,omitempty"`
}

// VersionControlDetails ... 3.23 versionControlDetails object
type VersionControlDetails struct {
	RepositoryURI URIString         `json:"repositoryUri"`
	RevisionID    string            `json:"revisionId"`
	Branch        string            `json:"branch,omitempty"`
	RevisionTag   string            `json:"revisionTag,omitempty"`
	AsOfTimeUTC   DateTimeString    `json:"asOfTimeUtc,omitempty"`
	MappedTo      *ArtifactLocation `json:"mappedTo,omitempty"`
	Properties    PropertyBag       `json:"properties,omitempty"`
}

// Artifact ... 3.24 artifact object
type Artifact struct {
	Location            *Location         `json:"location,omitempty"`
	ParentIndex         int               `json:"parentIndex,omitempty"`
	Offset              uint              `json:"offset,omitempty"`
	Length              uint              `json:"length,omitempty"`
	Roles               []string          `json:"roles,omitempty"`
	MIMEType            string            `json:"mimeType,omitempty"`
	Contents            *ArtifactContent  `json:"contents,omitempty"`
	Encoding            string            `json:"encoding,omitempty"`
	SourceLanguage      string            `json:"sourceLanguage,omitempty"`
	Hashes              map[string]string `json:"hashes,omitempty"`
	LastModifiedTimeUtc DateTimeString    `json:"lastModifiedTimeUtc,omitempty"`
	Description         *Message          `json:"description,omitempty"`
	Properties          PropertyBag       `json:"properties,omitempty"`
}

// SpecialLocations .... 3.25 specialLocations object
type SpecialLocations struct {
	DisplayBase *ArtifactLocation `json:"displayBase,omitempty"`
	Properties  PropertyBag       `json:"properties,omitempty"`
}

// TranslationMetadata ... 3.26 translationMetadata object
type TranslationMetadata struct {
	Name             string                    `json:"name"`
	FullName         string                    `json:"fullName,omitempty"`
	ShortDescription *MultiformatMessageString `json:"shortDescription,omitempty"`
	FullDescription  *MultiformatMessageString `json:"fullDescription,omitempty"`
	DownloadURI      URIString                 `json:"downloadUri,omitempty"`
	InformationURI   URIString                 `json:"informationUri,omitempty"`
	Properties       PropertyBag               `json:"properties,omitempty"`
}

// Result ... 3.27 result object
type Result struct {
	GUID                GUIDString                     `json:"guid,omitempty"`
	CorrelationGUID     GUIDString                     `json:"correlationGuid,omitempty"`
	RuleID              string                         `json:"ruleId,omitempty"`
	RuleIndex           string                         `json:"ruleIndex,omitempty"`
	Rule                *ReportingDescriptorReference  `json:"rule,omitempty"`
	Taxa                []ReportingDescriptorReference `json:"taxa,omitempty"`
	Kind                string                         `json:"kind,omitempty"`
	Level               string                         `json:"level,omitempty"`
	Message             *Message                       `json:"message"`
	Locations           []Location                     `json:"locations"`
	AnalysisTarget      *ArtifactLocation              `json:"analysisTarget,omitempty"`
	WebRequest          *WebRequest                    `json:"webRequest,omitempty"`
	WebResponse         *WebResponse                   `json:"webResponse,omitempty"`
	Fingerprints        map[string]string              `json:"fingerprints,omitempty"`
	PartialFingerprints map[string]string              `json:"partialFingerprints,omitempty"`
	CodeFlows           *CodeFlow                      `json:"codeFlows,omitempty"`
	Graphs              []Graph                        `json:"graphs,omitempty"`
	GraphTraversals     []GraphTraversal               `json:"graphTraversals,omitempty"`
	Stack               *Stack                         `json:"stack,omitempty"`
	RelatedLocations    []Location                     `json:"relatedLocations,omitempty"`
	Suppressions        []Suppression                  `json:"suppressions,omitempty"`
	BaselineState       string                         `json:"baselineState,omitempty"`
	Rank                float64                        `json:"rank,omitempty"`
	Attachments         []Attachment                   `json:"attachments,omitempty"`
	WorkItemURIs        *[]URIString                   `json:"workItemUris,omitempty"` // (nullable)
	HostedViewerURI     URIString                      `json:"hostedViewerUri,omitempty"`
	Provenance          *ResultProvenance              `json:"provenance,omitempty"`
	Fixes               []Fix                          `json:"fixes,omitempty"`
	OccurrenceCount     int                            `json:"occurrenceCount,omitempty"`
	Properties          PropertyBag                    `json:"properties,omitempty"`
}

// Location ... 3.28 location object
type Location struct {
	ID               int                    `json:"id,omitempty"`
	PhysicalLocation *PhysicalLocation      `json:"physicalLocation,omitempty"`
	LogicalLocations []LogicalLocation      `json:"logicalLocations,omitempty"`
	Message          *Message               `json:"message,omitempty"`
	Annotations      []Region               `json:"annotations,omitempty"`
	Relationships    []LocationRelationship `json:"relationships,omitempty"`
	Properties       PropertyBag            `json:"properties,omitempty"`
}

// PhysicalLocation ... 3.29 physicalLocation object
type PhysicalLocation struct {
	ArtifactLocation *ArtifactLocation `json:"atifactLocation,omitempty"`
	Region           *Region           `json:"region,omitempty"`
	ContextRegion    *Region           `json:"contextRegion,omitempty"`
	Address          *Address          `json:"address,omitempty"`
	Properties       PropertyBag       `json:"properties,omitempty"`
}

// Region ... 3.30 region object
type Region struct {
	StartLine      int              `json:"startLine,omitempty"`
	StartColumn    int              `json:"startColumn,omitempty"`
	EndLine        int              `json:"endLine,omitempty"`
	EndColumn      int              `json:"endColumn,omitempty"`
	CharOffset     int              `json:"charOffset,omitempty"`
	CharLength     int              `json:"charLength,omitempty"`
	ByteOffset     int              `json:"byteOffset,omitempty"`
	ByteLength     int              `json:"byteLength,omitempty"`
	Snippet        *ArtifactContent `json:"snippet,omitempty"`
	Message        *Message         `json:"message,omitempty"`
	SourceLanguage string           `json:"sourceLanguage,omitempty"`
	Properties     PropertyBag      `json:"properties,omitempty"`
}

// Rectangle ... 3.31 rectangle object
type Rectangle struct {
	Top        int         `json:"top,omitempty"`
	Left       int         `json:"left,omitempty"`
	Bottom     int         `json:"bottom,omitempty"`
	Right      int         `json:"right,omitempty"`
	Message    *Message    `json:"message,omitempty"`
	Properties PropertyBag `json:"properties,omitempty"`
}

// Address ... 3.32 address object
type Address struct {
	Index              int         `json:"index,omitempty"`
	AbsoluteAddress    int         `json:"absoluteAddress,omitempty"`
	RelativeAddress    int         `json:"relativeAddress,omitempty"`
	OffsetFromParent   int         `json:"offsetFromParent,omitempty"`
	Length             int         `json:"length,omitempty"`
	Name               string      `json:"name,omitempty"`
	FullyQualifiedName string      `json:"fullyQualifiedName,omitempty"`
	Kind               string      `json:"kind,omitempty"`
	ParentIndex        int         `json:"parentIndex,omitempty"`
	Properties         PropertyBag `json:"properties,omitempty"`
}

// LogicalLocation ... 3.33 logicalLocation object
type LogicalLocation struct {
	Index              int         `json:"index,omitempty"`
	Name               string      `json:"name,omitempty"`
	FullyQualifiedName string      `json:"fullyQualifiedName,omitempty"`
	DecoratedName      string      `json:"decoratedName,omitempty"`
	Kind               string      `json:"kind,omitempty"`
	ParentIndex        int         `json:"parentIndex,omitempty"`
	Properties         PropertyBag `json:"properties,omitempty"`
}

// LocationRelationship ... 3.34 locationRelationship object
type LocationRelationship struct {
	Target      int         `json:"target,omitempty"`
	Kinds       []string    `json:"kinds,omitempty"`
	Description *Message    `json:"description,omitempty"`
	Properties  PropertyBag `json:"properties,omitempty"`
}

// Suppression ... 3.35 suppression object
type Suppression struct {
	Kind          string      `json:"kind"`
	Status        string      `json:"status,omitempty"`
	Location      *Location   `json:"location,omitempty"`
	GUID          GUIDString  `json:"guid,omitempty"`
	Justification string      `json:"justification,omitempty"`
	Properties    PropertyBag `json:"properties,omitempty"`
}

// CodeFlow ... 3.36 codeFlow object
type CodeFlow struct {
	Message     *Message     `json:"message,omitempty"`
	ThreadFlows []ThreadFlow `json:"threadFlows"`
	Properties  PropertyBag  `json:"properties,omitempty"`
}

// ThreadFlow ... 3.37 threadFlow object
type ThreadFlow struct {
	ID             string                              `json:"id,omitempty"`
	Message        *Message                            `json:"message,omitempty"`
	InitialState   map[string]MultiformatMessageString `json:"initialState,omitempty"`
	ImmutableState map[string]MultiformatMessageString `json:"immutableState,omitempty"`
	Locations      []ThreadFlowLocation                `json:"locations,omitempty"`
	Properties     PropertyBag                         `json:"properties,omitempty"`
}

// ThreadFlowLocation ... 3.38 threadFlowLocation object
type ThreadFlowLocation struct {
	Index            int                                 `json:"index,omitempty"`
	Location         *Location                           `json:"location,omitempty"`
	Module           string                              `json:"module,omitempty"`
	Stack            *Stack                              `json:"stack,omitempty"`
	WebRequest       *WebRequest                         `json:"webRequest,omitempty"`
	WebResponse      *WebResponse                        `json:"webResponse,omitempty"`
	Kinds            []string                            `json:"kinds,omitempty"`
	State            map[string]MultiformatMessageString `json:"state,omitempty"`
	NestingLevel     int                                 `json:"nestingLevel,omitempty"`
	ExecutionOrder   int                                 `json:"executionOrder,omitempty"`
	ExecutionTimeUTC DateTimeString                      `json:"executionTimeUtc,omitempty"`
	Importance       string                              `json:"importance,omitempty"`
	Taxa             []ReportingDescriptorReference      `json:"taxa,omitempty"`
	Properties       PropertyBag                         `json:"properties,omitempty"`
}

// Graph ... 3.39 graph object
type Graph struct {
	Description *Message    `json:"description,omitempty"`
	Nodes       []Node      `json:"nodes,omitempty"`
	Edges       []Edge      `json:"edges,omitempty"`
	Properties  PropertyBag `json:"properties,omitempty"`
}

// Node ... 3.40 node object
type Node struct {
	ID         string      `json:"id"`
	Label      *Message    `json:"label,omitempty"`
	Location   *Location   `json:"location"`
	Children   []Node      `json:"children,omitempty"`
	Properties PropertyBag `json:"properties,omitempty"`
}

// Edge ... 3.41 edge object
type Edge struct {
	ID           string      `json:"id"`
	Label        *Message    `json:"label,omitempty"`
	SourceNodeID string      `json:"sourceNodeId"`
	TargetNodeID string      `json:"targetNodeId"`
	Properties   PropertyBag `json:"properties,omitempty"`
}

// GraphTraversal ... 3.42 graphTraversal object
type GraphTraversal struct {
	ResultGraphIndex int                                 `json:"resultGraphIndex"`
	RunGraphIndex    int                                 `json:"runGraphIndex"`
	Description      *Message                            `json:"description,omitempty"`
	InitialState     map[string]MultiformatMessageString `json:"initialState,omitempty"`
	ImmutableState   map[string]MultiformatMessageString `json:"immutableState,omitempty"`
	EdgeTraversals   []EdgeTraversal                     `json:"edgeTraversals,omitempty"`
	Properties       PropertyBag                         `json:"properties,omitempty"`
}

// EdgeTraversal ... 3.43 edgeTraversal object
type EdgeTraversal struct {
	EdgeID            string                              `json:"edgeId"`
	Message           *Message                            `json:"message,omitempty"`
	FinalState        map[string]MultiformatMessageString `json:"finalState,omitempty"`
	StepOverEdgeCount int                                 `json:"stepOverEdgeCount,omitempty"`
	Properties        PropertyBag                         `json:"properties,omitempty"`
}

// Stack ... 3.44 stack object
type Stack struct {
	Message    *Message     `json:"message,omitempty"`
	Frames     []StackFrame `json:"frames"`
	Properties PropertyBag  `json:"properties,omitempty"`
}

// StackFrame ... 3.45 stackFrame object
type StackFrame struct {
	Location   *Location   `json:"location,omitempty"`
	Module     string      `json:"module,omitempty"`
	ThreadID   int         `json:"threadId,omitempty"`
	Parameters []string    `json:"parameters,omitempty"`
	Properties PropertyBag `json:"properties,omitempty"`
}

// WebRequest ... 3.46 webRequest object
type WebRequest struct {
	Index      int                 `json:"index,omitempty"`
	Protocol   string              `json:"protocol"`
	Version    string              `json:"version"`
	Target     string              `json:"target"`
	Method     string              `json:"method"`
	Headers    map[string][]string `json:"headers"`
	Parameters map[string]string   `json:"parameters,omitempty"`
	Body       *ArtifactContent    `json:"body,omitempty"`
	Properties PropertyBag         `json:"properties,omitempty"`
}

// WebResponse ... 3.47 webResponse object
type WebResponse struct {
	Index              int                 `json:"index,omitempty"`
	Protocol           string              `json:"protocol"`
	Version            string              `json:"version"`
	StatusCode         int                 `json:"statusCode"`
	ReasonPhrase       string              `json:"reasonPhrase"`
	Headers            map[string][]string `json:"headers"`
	Body               *ArtifactContent    `json:"body,omitempty"`
	NoResponseReceived bool                `json:"noResponseReceived"`
	Properties         PropertyBag         `json:"properties,omitempty"`
}

// ResultProvenance ... 3.48 resultProvenance object
type ResultProvenance struct {
	FirstDetectionTimeUTC DateTimeString     `json:"firstDetectionTimeUtc,omitempty"`
	LastDetectionTimeUTC  DateTimeString     `json:"lastDetectionTimeUtc,omitempty"`
	FirstDetectionRunGUID GUIDString         `json:"firstDetectionRunGuid,omitempty"`
	LastDetectionRunGUID  GUIDString         `json:"lastDetectionRunGuid,omitempty"`
	InvocationIndex       int                `json:"invocationIndex,omitempty"`
	ConversionSources     []PhysicalLocation `json:"conversionSources,omitempty"`
	Properties            PropertyBag        `json:"properties,omitempty"`
}

// ReportingDescriptor ... 3.49 reportingDescriptor object
type ReportingDescriptor struct {
	ID                   string                              `json:"id"`
	DeprecatedIDs        []string                            `json:"deprecatedIds,omitempty"`
	GUID                 GUIDString                          `json:"guid,omitempty"`
	DeprecatedGUIDs      []GUIDString                        `json:"deprecatedGuids,omitempty"`
	Name                 string                              `json:"name,omitempty"`
	DeprecatedNames      []string                            `json:"deprecatedNames,omitempty"`
	ShortDescription     *MultiformatMessageString           `json:"shortDescription,omitempty"`
	FullDescription      *MultiformatMessageString           `json:"fullDescription,omitempty"`
	MessageStrings       map[string]MultiformatMessageString `json:"messageStrings,omitempty"`
	HelpURI              URIString                           `json:"helpUri,omitempty"`
	Help                 *MultiformatMessageString           `json:"help,omitempty"`
	DefaultConfiguration *ReportingConfiguration             `json:"defaultConfiguration,omitempty"`
	Relationships        []ReportingDescriptorRelationship   `json:"relationships,omitempty"`
	Properties           PropertyBag                         `json:"properties,omitempty"`
}

// ReportingConfiguration ... 3.50 reportingConfiguration object
type ReportingConfiguration struct {
	Enabled    bool        `json:"enabled,omitempty"`
	Level      string      `json:"level,omitempty"`
	Rank       float64     `json:"rank,omitempty"`
	Parameters PropertyBag `json:"parameters,omitempty"`
	Properties PropertyBag `json:"properties,omitempty"`
}

// ConfigurationOverride ... 3.51 configurationOverride object
type ConfigurationOverride struct {
	Descriptor    *ReportingDescriptorReference `json:"descriptor"`
	Configuration *ReportingConfiguration       `json:"configuration"`
	Properties    PropertyBag                   `json:"properties,omitempty"`
}

// ReportingDescriptorReference ... 3.52 reportingDescriptorReference object
type ReportingDescriptorReference struct {
	ID            string                  `json:"id,omitempty"`
	Index         int                     `json:"index,omitempty"`
	GUID          GUIDString              `json:"guid,omitempty"`
	ToolComponent *ToolComponentReference `json:"toolComponent,omitempty"`
	Properties    PropertyBag             `json:"properties,omitempty"`
}

// ReportingDescriptorRelationship ... 3.53 reportingDescriptorRelationship object
type ReportingDescriptorRelationship struct {
	Target      *ReportingDescriptorReference `json:"target"`
	Kinds       []string                      `json:"kinds,omitempty"`
	Description *Message                      `json:"description"`
	Properties  PropertyBag                   `json:"properties,omitempty"`
}

// ToolComponentReference ... 3.54 toolComponentReference object
type ToolComponentReference struct {
	Name       string      `json:"name,omitempty"`
	Index      int         `json:"index,omitempty"`
	GUID       GUIDString  `json:"guid,omitempty"`
	Properties PropertyBag `json:"properties,omitempty"`
}

// Fix ... 3.55 fix object
type Fix struct {
	Description     *Message         `json:"description"`
	ArtifactChanges []ArtifactChange `json:"artifactChanges"`
	Properties      PropertyBag      `json:"properties,omitempty"`
}

// ArtifactChange ... 3.56 artifactChange object
type ArtifactChange struct {
	ArtifactLocation *ArtifactLocation `json:"artifactLocation"`
	Replacements     []Replacement     `json:"replacements"`
	Properties       PropertyBag       `json:"properties,omitempty"`
}

// Replacement ... 3.57 replacement object
type Replacement struct {
	DeletedRegion   *Region          `json:"deletedRegion"`
	InsertedContent *ArtifactContent `json:"insertedContent,omitempty"`
	Properties      PropertyBag      `json:"properties,omitempty"`
}

// Notification ... 3.58 notification object
type Notification struct {
	Descriptor     *ReportingDescriptorReference `json:"descriptor"`
	AssociatedRule *ReportingDescriptorReference `json:"associatedRule,omitempty"`
	Locations      []Location                    `json:"locations"`
	Message        *Message                      `json:"message"`
	Level          string                        `json:"level,omitempty"`
	ThreadID       int                           `json:"threadId,omitempty"`
	TimeUTC        DateTimeString                `json:"timeUtc,omitempty"`
	Exception      *Exception                    `json:"exception,omitempty"`
	Properties     PropertyBag                   `json:"properties,omitempty"`
}

// Exception ... 3.59 exception object
type Exception struct {
	Kind            string      `json:"kind"`
	Message         *Message    `json:"message"`
	Stack           *Stack      `json:"stack,omitempty"`
	InnerExceptions []Exception `json:"innerExceptions,omitempty"`
	Properties      PropertyBag `json:"properties,omitempty"`
}

// ExternalProperties ... 4.3 externalProperties object
type ExternalProperties struct {
	Schema  URIString  `json:"$schema,omitempty"`
	Version string     `json:"version,omitempty"`
	GUID    GUIDString `json:"guid,omitempty"`
	RunGUID GUIDString `json:"runGuid,omitempty"`
}
