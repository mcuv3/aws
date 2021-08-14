package lambda



type ImagePullOptions struct {
	ImageName     string
	Platform      string
}



type ContainerRuntine interface {
	BuildImage() error
	PullImage(ImagePullOptions) error
	RunContainer() error
	RemoveImage() error

}


type DockerRuntime struct {}