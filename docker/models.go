package docker



type ContainerConf struct {
	
} 

type BuildImageOptions struct {
	Tags           []string
	CPUSetCPUs     string
	CPUSetMems     string
	CPUShares      int64
	CPUQuota       int64
	CPUPeriod      int64
	Memory         int64
	CgroupParent   string
	NetworkMode    string
	ShmSize        int64
	Dockerfile     string
}


type RunContainerOptions struct {              // Domainname
		User            string               // Attach the standard error
		ExposedPorts    []int            // If true, close stdin after the 1 attached client disconnects.
		Env             []string            // List of environment variable to set in the container
		Cmd             []string   // Command to run when starting the container
	//	Healthcheck     *HealthConfig       `json:",omitempty"` // Healthcheck describes how to check the container is healthy
		ArgsEscaped     bool                `json:",omitempty"` // True if command is already escaped (meaning treat as a command line) (Windows specific).
		Image           string              // Name of the image as it was passed by the operator (e.g. could be symbolic)
		Volumes         map[string]struct{} // List of volumes (mounts) used for the container
		Labels          map[string]string   // List of labels set to this container
}


type ContainerRuntime interface {
	BuildImage(code string,imageName string,runtime string,options BuildImageOptions) error
	PullImage(imageName string) error 
	RunContainer(options RunContainerOptions) error
	RemoveImage(imageName string) error
}

type DockerRuntime struct{}
