package docker

import (
	"archive/tar"
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

 

type DockerRuntime struct{}



func (d *DockerRuntime) BuildImage(options BuildImageOptions) error {
	ctx := context.Background()
    cli, err := client.NewEnvClient()
    if err != nil {
        log.Fatal(err, " :unable to init client")
    }

    buf := new(bytes.Buffer)
    tw := tar.NewWriter(buf)
    defer tw.Close() 

    lang := options.Language

    dockerF :=[]byte(fmt.Sprintf(`
	FROM %s:%s
	RUN echo "%s" > code%s
	CMD ["%s", "code%s"]
	`,lang.Name,lang.Runtimes[0],options.Code,lang.Extension,lang.Activator,lang.Extension))

    tarHeader := &tar.Header{
        Name: "Dockerfile",
        Size: int64(len(dockerF)), 
    }
    err = tw.WriteHeader(tarHeader)
    if err != nil {
        log.Fatal(err, " :unable to write tar header")
    }
    _, err = tw.Write(dockerF)
    if err != nil {
        log.Fatal(err, " :unable to write tar body")
    }
    dockerFileTarReader := bytes.NewReader(buf.Bytes())

    imageBuildResponse, err := cli.ImageBuild(
        ctx,
        dockerFileTarReader,
        types.ImageBuildOptions{
            Context:    dockerFileTarReader,
            Dockerfile: "Dockerfile",
			Tags: []string{strings.ToLower(options.ImageName)},
            Remove:     false})
    if err != nil {
        log.Fatal(err, " :unable to build docker image ",options.ImageName)
    }
    defer imageBuildResponse.Body.Close()
    _, err = io.Copy(os.Stdout, imageBuildResponse.Body)
    if err != nil {
        log.Fatal(err, " :unable to read image build response")
    }
    
	return nil 
}
func (d *DockerRuntime) PullImage(imageName string) error {
	time.Sleep(time.Second * 2)
	fmt.Println("Image Pull")
	return nil
} 
func (d *DockerRuntime) RunContainer(options RunContainerOptions) error {
	time.Sleep(time.Second * 2)
	fmt.Println("Run Container")
	return nil
}
func (d *DockerRuntime) RemoveImage(imageName string) error {
	time.Sleep(time.Second * 2)
	fmt.Println("Remove Image")
	return nil
}

