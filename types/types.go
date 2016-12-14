package types

import (
	"time"
)

type PodList struct {
	TypeMeta `json:",inline"`
	ListMeta `json:"metadata,omitempty"`

	Items []Pod `json:"items"`
}

type Pod struct {
	TypeMeta `json:",inline"`
	Metadata ObjectMeta `json:"metadata,omitempty"`

	Spec   PodSpec   `json:"spec,omitempty"`
	Status PodStatus `json:"status,omitempty"`
	Ages   string    `json:"ages,omitempty"`
}

type PodStatus struct {
	Phase      PodPhase       `json:"phase,omitempty"`
	Conditions []PodCondition `json:"conditions,omitempty"`
	Message    string         `json:"message,omitempty"`
	Reason     string         `json:"reason,omitempty"`

	HostIP string `json:"hostIP,omitempty"`
	PodIP  string `json:"podIP,omitempty"`

	StartTime             time.Time         `json:"startTime,omitempty"`
	InitContainerStatuses []ContainerStatus `json:"-"`
	ContainerStatuses     []ContainerStatus `json:"containerStatuses,omitempty"`
}
type ContainerStatus struct {
	Name                 string         `json:"name"`
	State                ContainerState `json:"state,omitempty"`
	LastTerminationState ContainerState `json:"lastState,omitempty"`
	Ready                bool           `json:"ready"`
	RestartCount         int32          `json:"restartCount"`
	Image                string         `json:"image"`
	ImageID              string         `json:"imageID"`
	ContainerID          string         `json:"containerID,omitempty"`
}
type ContainerState struct {
	Waiting    *ContainerStateWaiting    `json:"waiting,omitempty"`
	Running    *ContainerStateRunning    `json:"running,omitempty"`
	Terminated *ContainerStateTerminated `json:"terminated,omitempty"`
}
type ContainerStateWaiting struct {
	Reason  string `json:"reason,omitempty"`
	Message string `json:"message,omitempty"`
}
type ContainerStateRunning struct {
	StartedAt time.Time `json:"startedAt,omitempty"`
}

type ContainerStateTerminated struct {
	ExitCode    int32     `json:"exitCode"`
	Signal      int32     `json:"signal,omitempty"`
	Reason      string    `json:"reason,omitempty"`
	Message     string    `json:"message,omitempty"`
	StartedAt   time.Time `json:"startedAt,omitempty"`
	FinishedAt  time.Time `json:"finishedAt,omitempty"`
	ContainerID string    `json:"containerID,omitempty"`
}

type ConditionStatus string

const (
	ConditionTrue    ConditionStatus = "True"
	ConditionFalse   ConditionStatus = "False"
	ConditionUnknown ConditionStatus = "Unknown"
)

type PodCondition struct {
	Type               PodConditionType `json:"type"`
	Status             ConditionStatus  `json:"status"`
	LastProbeTime      time.Time        `json:"lastProbeTime,omitempty"`
	LastTransitionTime time.Time        `json:"lastTransitionTime,omitempty"`
	Reason             string           `json:"reason,omitempty"`
	Message            string           `json:"message,omitempty"`
}
type PodConditionType string

const (
	PodScheduled   PodConditionType = "PodScheduled"
	PodReady       PodConditionType = "Ready"
	PodInitialized PodConditionType = "Initialized"
)

type PodPhase string

const (
	PodPending   PodPhase = "Pending"
	PodRunning   PodPhase = "Running"
	PodSucceeded PodPhase = "Succeeded"
	PodFailed    PodPhase = "Failed"
	PodUnknown   PodPhase = "Unknown"
)

type NodeList struct {
	TypeMeta `json:",inline"`
	ListMeta `json:"metadata,omitempty"`

	Items []Node `json:"items"`
}
type Node struct {
	TypeMeta `json:",inline"`
	Metadata ObjectMeta `json:"metadata,omitempty"`
	Spec     NodeSpec   `json:"spec,omitempty"`
	Status   NodeStatus `json:"status,omitempty"`
	Ages     string     `json:ages`
}
type NodeSpec struct {
	PodCIDR       string `json:"podCIDR,omitempty"`
	ExternalID    string `json:"externalID,omitempty"`
	ProviderID    string `json:"providerID,omitempty"`
	Unschedulable bool   `json:"unschedulable,omitempty"`
}
type NodeStatus struct {
	Capacity        map[string]string   `json:"capacity,omitempty"`
	Allocatable     map[string]string   `json:"allocatable,omitempty"`
	Phase           string              `json:"phase,omitempty"`
	Conditions      []NodeCondition     `json:"conditions,omitempty"`
	Addresses       []NodeAddress       `json:"addresses,omitempty"`
	DaemonEndpoints NodeDaemonEndpoints `json:"daemonEndpoints,omitempty"`
	NodeInfo        NodeSystemInfo      `json:"nodeInfo,omitempty"`
	Images          []ContainerImage    `json:"images,omitempty"`
	VolumesInUse    []UniqueVolumeName  `json:"volumesInUse,omitempty"`
	VolumesAttached []AttachedVolume    `json:"volumesAttached,omitempty"`
}
type NodeDaemonEndpoints struct {
	KubeletEndpoint DaemonEndpoint `json:"kubeletEndpoint,omitempty"`
}
type NodeSystemInfo struct {
	MachineID               string `json:"machineID"`
	SystemUUID              string `json:"systemUUID"`
	BootID                  string `json:"bootID"`
	KernelVersion           string `json:"kernelVersion"`
	OSImage                 string `json:"osImage"`
	ContainerRuntimeVersion string `json:"containerRuntimeVersion"`
	KubeletVersion          string `json:"kubeletVersion"`
	KubeProxyVersion        string `json:"kubeProxyVersion"`
	OperatingSystem         string `json:"operatingSystem"`
	Architecture            string `json:"architecture"`
}

type DaemonEndpoint struct {
	Port int32 `json:"Port"`
}
type UniqueVolumeName string

type AttachedVolume struct {
	Name       UniqueVolumeName `json:"name"`
	DevicePath string           `json:"devicePath"`
}

type NodeAddress struct {
	Type    string `json:"type"`
	Address string `json:"address"`
}
type ContainerImage struct {
	Names     []string `json:"names"`
	SizeBytes int64    `json:"sizeBytes,omitempty"`
	SizeAuto  string   `json:sizeAuto`
}

type NodeCondition struct {
	Type               string    `json:"type"`
	Status             string    `json:"status"`
	LastHeartbeatTime  time.Time `json:"lastHeartbeatTime,omitempty"`
	LastTransitionTime time.Time `json:"lastTransitionTime,omitempty"`
	Reason             string    `json:"reason,omitempty"`
	Message            string    `json:"message,omitempty"`
}

type ComponentStatusList struct {
	TypeMeta `json:",inline"`
	ListMeta `json:"metadata,omitempty"`

	Items []ComponentStatus `json:"items"`
}
type ComponentStatus struct {
	TypeMeta `json:",inline"`
	Metadata ObjectMeta `json:"metadata,omitempty"`

	Conditions []ComponentCondition `json:"conditions,omitempty"`
}

type ComponentCondition struct {
	Type    string `json:"type"`
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}

type EndpointsList struct {
	TypeMeta `json:",inline"`
	ListMeta `json:"metadata,omitempty"`

	Items []Endpoints `json:"items"`
}

type Endpoints struct {
	TypeMeta `json:",inline"`
	Metadata ObjectMeta
	Subsets  []EndpointSubset
	Ages     string `json:ages`
}
type EndpointSubset struct {
	Addresses         []EndpointAddress
	NotReadyAddresses []EndpointAddress
	Ports             []EndpointPort
}

type EndpointAddress struct {
	Ip        string          `json:ip`
	Hostname  string          `json:hostname,omitempty`
	TargetRef ObjectReference `json:targetRef,omitempty`
}
type EndpointPort struct {
	Name     string `json:name,omitempty`
	Port     int32  `json:port`
	Protocol string `json:Protocol,omitempty`
}

type ServiceList struct {
	TypeMeta
	Metadata ListMeta
	Items    []Service
}
type Service struct {
	TypeMeta
	Metadata ObjectMeta
	Spec     ServiceSpec
	Status   ServiceStatus
	Ages     string `json:ages`
}

type ServiceSpec struct {
	Ports                    []ServicePort
	Selector                 map[string]string `json:selector`
	ClusterIP                string            `json:clusterIP`
	Type                     string            `json:type`
	ExternalIPs              []string          `json:ExternalIPs`
	DeprecatedPublicIPs      []string          `json:deprecatedPublicIPs`
	SessionAffinity          string            `json:sessionAffinity`
	LoadBalancerIP           string            `json:loadBalancerIP`
	LoadBalancerSourceRanges []string          `json:loadBalancerSourceRanges`
}
type ServiceStatus struct {
	LoadBalancer LoadBalancerStatus
}
type ServicePort struct {
	Name       string `json:name`
	Protocol   string `json:protocol`
	Port       int32  `json:port`
	TargetPort int32  `json:targetPort`
	NodePort   int32  `json:nodePort`
}

type ReplicationControllerList struct {
	TypeMeta `json:",inline"`
	ListMeta `json:"metadata,omitempty"`
	Items    []ReplicationController `json:"items"`
}
type ReplicationController struct {
	TypeMeta `json:",inline"`
	Metadata ObjectMeta                  `json:"metadata,omitempty"`
	Ages     string                      `json:ages`
	Spec     ReplicationControllerSpec   `json:"spec,omitempty"`
	Status   ReplicationControllerStatus `json:"status,omitempty"`
}

type ReplicationControllerSpec struct {
	Replicas int32             `json:"replicas"`
	Selector map[string]string `json:"selector"`
	Template *PodTemplateSpec  `json:"template,omitempty"`
}
type ReplicationControllerStatus struct {
	Replicas             int32 `json:"replicas"`
	FullyLabeledReplicas int32 `json:"fullyLabeledReplicas,omitempty"`
	ObservedGeneration   int64 `json:"observedGeneration,omitempty"`
}
type PodTemplateSpec struct {
	Metadata ObjectMeta `json:"metadata,omitempty"`
	Spec     PodSpec    `json:"spec,omitempty"`
}
type PodSpec struct {
	Volumes                       []Volume               `json:"volumes"`
	InitContainers                []Container            `json:"-"`
	Containers                    []Container            `json:"containers"`
	RestartPolicy                 string                 `json:"restartPolicy,omitempty"`
	TerminationGracePeriodSeconds *int64                 `json:"terminationGracePeriodSeconds,omitempty"`
	ActiveDeadlineSeconds         *int64                 `json:"activeDeadlineSeconds,omitempty"`
	DNSPolicy                     string                 `json:"dnsPolicy,omitempty"`
	NodeSelector                  map[string]string      `json:"nodeSelector,omitempty"`
	ServiceAccountName            string                 `json:"serviceAccountName"`
	NodeName                      string                 `json:"nodeName,omitempty"`
	SecurityContext               *PodSecurityContext    `json:"securityContext,omitempty"`
	ImagePullSecrets              []LocalObjectReference `json:"imagePullSecrets,omitempty"`
	Hostname                      string                 `json:"hostname,omitempty"`
	Subdomain                     string                 `json:"subdomain,omitempty"`
}
type Container struct {
	Name                   string               `json:"name"`
	Image                  string               `json:"image"`
	Command                []string             `json:"command,omitempty"`
	Args                   []string             `json:"args,omitempty"`
	WorkingDir             string               `json:"workingDir,omitempty"`
	Ports                  []ContainerPort      `json:"ports,omitempty"`
	Env                    []EnvVar             `json:"env,omitempty"`
	Resources              ResourceRequirements `json:"resources,omitempty"`
	VolumeMounts           []VolumeMount        `json:"volumeMounts,omitempty"`
	LivenessProbe          *Probe               `json:"livenessProbe,omitempty"`
	ReadinessProbe         *Probe               `json:"readinessProbe,omitempty"`
	Lifecycle              *Lifecycle           `json:"lifecycle,omitempty"`
	TerminationMessagePath string               `json:"terminationMessagePath,omitempty"`
	ImagePullPolicy        string               `json:"imagePullPolicy"`
	SecurityContext        *SecurityContext     `json:"securityContext,omitempty"`
	Stdin                  bool                 `json:"stdin,omitempty"`
	StdinOnce              bool                 `json:"stdinOnce,omitempty"`
	TTY                    bool                 `json:"tty,omitempty"`
}
type SecurityContext struct {
	Capabilities           *Capabilities   `json:"capabilities,omitempty"`
	Privileged             *bool           `json:"privileged,omitempty"`
	SELinuxOptions         *SELinuxOptions `json:"seLinuxOptions,omitempty"`
	RunAsUser              *int64          `json:"runAsUser,omitempty"`
	RunAsNonRoot           *bool           `json:"runAsNonRoot,omitempty"`
	ReadOnlyRootFilesystem *bool           `json:"readOnlyRootFilesystem,omitempty"`
}
type SELinuxOptions struct {
	User  string `json:"user,omitempty"`
	Role  string `json:"role,omitempty"`
	Type  string `json:"type,omitempty"`
	Level string `json:"level,omitempty"`
}

type Capability string
type Capabilities struct {
	Add  []Capability `json:"add,omitempty"`
	Drop []Capability `json:"drop,omitempty"`
}

type Lifecycle struct {
	PostStart *Handler `json:"postStart,omitempty"`
	PreStop   *Handler `json:"preStop,omitempty"`
}
type ExecAction struct {
	Command []string `json:"command,omitempty"`
}
type EnvVar struct {
	Name      string        `json:"name"`
	Value     string        `json:"value,omitempty"`
	ValueFrom *EnvVarSource `json:"valueFrom,omitempty"`
}
type EnvVarSource struct {
	FieldRef         *ObjectFieldSelector   `json:"fieldRef,omitempty"`
	ResourceFieldRef *ResourceFieldSelector `json:"resourceFieldRef,omitempty"`
	ConfigMapKeyRef  *ConfigMapKeySelector  `json:"configMapKeyRef,omitempty"`
	SecretKeyRef     *SecretKeySelector     `json:"secretKeyRef,omitempty"`
}

type Handler struct {
	Exec      *ExecAction      `json:"exec,omitempty"`
	HTTPGet   *HTTPGetAction   `json:"httpGet,omitempty"`
	TCPSocket *TCPSocketAction `json:"tcpSocket,omitempty"`
}
type HTTPGetAction struct {
	Path        string       `json:"path,omitempty"`
	Port        IntOrString  `json:"port,omitempty"`
	Host        string       `json:"host,omitempty"`
	Scheme      string       `json:"scheme,omitempty"`
	HTTPHeaders []HTTPHeader `json:"httpHeaders,omitempty"`
}
type HTTPHeader struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type TCPSocketAction struct {
	Port IntOrString `json:"port,omitempty"`
}
type IntOrString struct {
	Type   Type   `protobuf:"varint,1,opt,name=type,casttype=Type"`
	IntVal int32  `protobuf:"varint,2,opt,name=intVal"`
	StrVal string `protobuf:"bytes,3,opt,name=strVal"`
}

type PodSecurityContext struct {
	HostNetwork        bool            `json:"hostNetwork,omitempty"`
	HostPID            bool            `json:"hostPID,omitempty"`
	HostIPC            bool            `json:"hostIPC,omitempty"`
	SELinuxOptions     *SELinuxOptions `json:"seLinuxOptions,omitempty"`
	RunAsUser          *int64          `json:"runAsUser,omitempty"`
	RunAsNonRoot       *bool           `json:"runAsNonRoot,omitempty"`
	SupplementalGroups []int64         `json:"supplementalGroups,omitempty"`
	FSGroup            *int64          `json:"fsGroup,omitempty"`
}

type Type int

const (
	Int    Type = iota // The IntOrString holds an int.
	String             // The IntOrString holds a string.
)

type RestartPolicy string

const (
	RestartPolicyAlways    RestartPolicy = "Always"
	RestartPolicyOnFailure RestartPolicy = "OnFailure"
	RestartPolicyNever     RestartPolicy = "Never"
)

type DNSPolicy string

const (
	DNSClusterFirst DNSPolicy = "ClusterFirst"
	DNSDefault      DNSPolicy = "Default"
)

type PullPolicy string

const (
	PullAlways       PullPolicy = "Always"
	PullNever        PullPolicy = "Never"
	PullIfNotPresent PullPolicy = "IfNotPresent"
)

type Probe struct {
	Handler             `json:",inline"`
	InitialDelaySeconds int32 `json:"initialDelaySeconds,omitempty"`
	TimeoutSeconds      int32 `json:"timeoutSeconds,omitempty"`
	PeriodSeconds       int32 `json:"periodSeconds,omitempty"`
	SuccessThreshold    int32 `json:"successThreshold,omitempty"`
	FailureThreshold    int32 `json:"failureThreshold,omitempty"`
}

type ResourceRequirements struct {
	Limits   map[string]string `json:"limits,omitempty"`
	Requests map[string]string `json:"requests,omitempty"`
}
type VolumeMount struct {
	Name      string `json:"name"`
	ReadOnly  bool   `json:"readOnly,omitempty"`
	MountPath string `json:"mountPath"`
	SubPath   string `json:"subPath,omitempty"`
}

type ContainerPort struct {
	Name          string `json:"name,omitempty"`
	HostPort      int32  `json:"hostPort,omitempty"`
	ContainerPort int32  `json:"containerPort"`
	Protocol      string `json:"protocol,omitempty"`
	HostIP        string `json:"hostIP,omitempty"`
}
type ConfigMapKeySelector struct {
	LocalObjectReference `json:",inline"`
	Key                  string `json:"key"`
}
type SecretKeySelector struct {
	LocalObjectReference `json:",inline"`
	Key                  string `json:"key"`
}
type Volume struct {
	Name         string `json:"name"`
	VolumeSource `json:",inline,omitempty"`
}
type VolumeSource struct {
	HostPath              *HostPathVolumeSource              `json:"hostPath,omitempty"`
	EmptyDir              *EmptyDirVolumeSource              `json:"emptyDir,omitempty"`
	GCEPersistentDisk     *GCEPersistentDiskVolumeSource     `json:"gcePersistentDisk,omitempty"`
	AWSElasticBlockStore  *AWSElasticBlockStoreVolumeSource  `json:"awsElasticBlockStore,omitempty"`
	GitRepo               *GitRepoVolumeSource               `json:"gitRepo,omitempty"`
	Secret                *SecretVolumeSource                `json:"secret,omitempty"`
	NFS                   *NFSVolumeSource                   `json:"nfs,omitempty"`
	ISCSI                 *ISCSIVolumeSource                 `json:"iscsi,omitempty"`
	Glusterfs             *GlusterfsVolumeSource             `json:"glusterfs,omitempty"`
	PersistentVolumeClaim *PersistentVolumeClaimVolumeSource `json:"persistentVolumeClaim,omitempty"`
	RBD                   *RBDVolumeSource                   `json:"rbd,omitempty"`
	FlexVolume            *FlexVolumeSource                  `json:"flexVolume,omitempty"`
	Cinder                *CinderVolumeSource                `json:"cinder,omitempty"`
	CephFS                *CephFSVolumeSource                `json:"cephfs,omitempty"`
	Flocker               *FlockerVolumeSource               `json:"flocker,omitempty"`
	DownwardAPI           *DownwardAPIVolumeSource           `json:"downwardAPI,omitempty"`
	FC                    *FCVolumeSource                    `json:"fc,omitempty"`
	AzureFile             *AzureFileVolumeSource             `json:"azureFile,omitempty"`
	ConfigMap             *ConfigMapVolumeSource             `json:"configMap,omitempty"`
	VsphereVolume         *VsphereVirtualDiskVolumeSource    `json:"vsphereVolume,omitempty"`
}
type VsphereVirtualDiskVolumeSource struct {
	VolumePath string `json:"volumePath"`
	FSType     string `json:"fsType,omitempty"`
}

type ConfigMapVolumeSource struct {
	LocalObjectReference `json:",inline"`
	Items                []KeyToPath `json:"items,omitempty"`
}

type AzureFileVolumeSource struct {
	SecretName string `json:"secretName"`
	ShareName  string `json:"shareName"`
	ReadOnly   bool   `json:"readOnly,omitempty"`
}

type FCVolumeSource struct {
	TargetWWNs []string `json:"targetWWNs"`
	Lun        *int32   `json:"lun"`
	FSType     string   `json:"fsType,omitempty"`
	ReadOnly   bool     `json:"readOnly,omitempty"`
}

type DownwardAPIVolumeSource struct {
	Items []DownwardAPIVolumeFile `json:"items,omitempty"`
}

type DownwardAPIVolumeFile struct {
	Path             string                 `json:"path"`
	FieldRef         *ObjectFieldSelector   `json:"fieldRef,omitempty"`
	ResourceFieldRef *ResourceFieldSelector `json:"resourceFieldRef,omitempty"`
}
type ResourceFieldSelector struct {
	ContainerName string `json:"containerName,omitempty"`
	Resource      string `json:"resource"`
	Divisor       string `json:"divisor,omitempty"`
}

type ObjectFieldSelector struct {
	APIVersion string `json:"apiVersion"`
	FieldPath  string `json:"fieldPath"`
}

type FlockerVolumeSource struct {
	DatasetName string `json:"datasetName"`
}

type CephFSVolumeSource struct {
	Monitors   []string              `json:"monitors"`
	Path       string                `json:"path,omitempty"`
	User       string                `json:"user,omitempty"`
	SecretFile string                `json:"secretFile,omitempty"`
	SecretRef  *LocalObjectReference `json:"secretRef,omitempty"`
	ReadOnly   bool                  `json:"readOnly,omitempty"`
}
type CinderVolumeSource struct {
	VolumeID string `json:"volumeID"`
	FSType   string `json:"fsType,omitempty"`
	ReadOnly bool   `json:"readOnly,omitempty"`
}

type FlexVolumeSource struct {
	Driver    string                `json:"driver"`
	FSType    string                `json:"fsType,omitempty"`
	SecretRef *LocalObjectReference `json:"secretRef,omitempty"`
	ReadOnly  bool                  `json:"readOnly,omitempty"`
	Options   map[string]string     `json:"options,omitempty"`
}

type RBDVolumeSource struct {
	CephMonitors []string              `json:"monitors"`
	RBDImage     string                `json:"image"`
	FSType       string                `json:"fsType,omitempty"`
	RBDPool      string                `json:"pool,omitempty"`
	RadosUser    string                `json:"user,omitempty"`
	Keyring      string                `json:"keyring,omitempty"`
	SecretRef    *LocalObjectReference `json:"secretRef,omitempty"`
	ReadOnly     bool                  `json:"readOnly,omitempty"`
}
type LocalObjectReference struct {
	Name string
}

type PersistentVolumeClaimVolumeSource struct {
	ClaimName string `json:"claimName"`
	ReadOnly  bool   `json:"readOnly,omitempty"`
}

type GlusterfsVolumeSource struct {
	EndpointsName string `json:"endpoints"`
	Path          string `json:"path"`
	ReadOnly      bool   `json:"readOnly,omitempty"`
}

type ISCSIVolumeSource struct {
	TargetPortal   string `json:"targetPortal,omitempty"`
	IQN            string `json:"iqn,omitempty"`
	Lun            int32  `json:"lun,omitempty"`
	ISCSIInterface string `json:"iscsiInterface,omitempty"`
	FSType         string `json:"fsType,omitempty"`
	ReadOnly       bool   `json:"readOnly,omitempty"`
}
type NFSVolumeSource struct {
	Server   string `json:"server"`
	Path     string `json:"path"`
	ReadOnly bool   `json:"readOnly,omitempty"`
}
type SecretVolumeSource struct {
	SecretName string      `json:"secretName,omitempty"`
	Items      []KeyToPath `json:"items,omitempty"`
}
type KeyToPath struct {
	Key  string `json:"key"`
	Path string `json:"path"`
}

type GitRepoVolumeSource struct {
	Repository string `json:"repository"`
	Revision   string `json:"revision,omitempty"`
	Directory  string `json:"directory,omitempty"`
}
type AWSElasticBlockStoreVolumeSource struct {
	VolumeID  string `json:"volumeID"`
	FSType    string `json:"fsType,omitempty"`
	Partition int32  `json:"partition,omitempty"`
	ReadOnly  bool   `json:"readOnly,omitempty"`
}
type GCEPersistentDiskVolumeSource struct {
	PDName    string `json:"pdName"`
	FSType    string `json:"fsType,omitempty"`
	Partition int32  `json:"partition,omitempty"`
	ReadOnly  bool   `json:"readOnly,omitempty"`
}
type EmptyDirVolumeSource struct {
	Medium StorageMedium `json:"medium,omitempty"`
}
type StorageMedium string

const (
	StorageMediumDefault StorageMedium = ""       // use whatever the default is for the node
	StorageMediumMemory  StorageMedium = "Memory" // use memory (tmpfs)
)

type HostPathVolumeSource struct {
	Path string `json:"path"`
}

type LoadBalancerStatus struct {
	Ingress LoadBalancerIngress
}
type LoadBalancerIngress struct {
	Ip       string `json:ip`
	Hostname string `json:hostname`
}

type TypeMeta struct {
	Kind       string `json:"kind,omitempty"`
	APIVersion string `json:"apiVersion,omitempty"`
}

type ListMeta struct {
	SelfLink        string `json:"selfLink,omitempty"`
	ResourceVersion string `json:"resourceVersion,omitempty"`
}

type ObjectMeta struct {
	Name                       string            `json:"name,omitempty"`
	GenerateName               string            `json:"generateName,omitempty"`
	Namespace                  string            `json:"namespace,omitempty"`
	SelfLink                   string            `json:"selfLink,omitempty"`
	UID                        string            `json:"uid,omitempty"`
	ResourceVersion            string            `json:"resourceVersion,omitempty"`
	Generation                 int64             `json:"generation,omitempty"`
	CreationTimestamp          string            `json:"creationTimestamp,omitempty"`
	DeletionTimestamp          string            `json:"deletionTimestamp,omitempty"`
	DeletionGracePeriodSeconds int64             `json:"deletionGracePeriodSeconds,omitempty"`
	Labels                     map[string]string `json:"labels,omitempty"`
	Annotations                map[string]string `json:"annotations,omitempty"`
	OwnerReferences            []OwnerReference  `json:"ownerReferences,omitempty"`
	Finalizers                 []string          `json:"finalizers,omitempty"`
}
type OwnerReference struct {
	APIVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
	Name       string `json:"name"`
	Uid        string `json:"uid"`
	Controller bool   `json:"controller,omitempty"`
}

type ObjectReference struct {
	Kind            string `json:"kind,omitempty"`
	Namespace       string `json:"namespace,omitempty"`
	Name            string `json:"name,omitempty"`
	Uid             string `json:"uid,omitempty"`
	APIVersion      string `json:"apiVersion,omitempty"`
	ResourceVersion string `json:"resourceVersion,omitempty"`
	FieldPath       string `json:"fieldPath,omitempty"`
}
