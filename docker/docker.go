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
	"github.com/docker/docker/api/types/container"
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

    dockerF :=[]byte(options.DockerFile)

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

    // should be placed on a cloudwatch service
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
    ctx := context.Background()
    cli, err := client.NewEnvClient()
    if err != nil {
        log.Fatal(err, " :unable to init client")
    }

    env := d.GetEnv(options.Environment)

    resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: options.Image,
        Env: env,
        
	}, &container.HostConfig{
		Resources: container.Resources{
			Memory: options.Ram,
		},
        NetworkMode: "host",
	},nil, nil, options.Name)
	if err != nil {
        fmt.Println(err)
		return err
	}

    fmt.Println("Container id=" + resp.ID)


	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{		
	}); err != nil {
        fmt.Println(err)
		return err
	}

	return nil
}
func (d *DockerRuntime) RemoveImage(imageName string) error {
	time.Sleep(time.Second * 2)
	fmt.Println("Remove Image")
	return nil
}



func (d *DockerRuntime) GetEnv(env map[string]interface{}) []string {
    var environ []string

    for k,v := range env {
        environ = append(environ, fmt.Sprintf("%s=%s", k, v))
    }

    return environ
}