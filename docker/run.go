package docker

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func m() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	imageName := "mcuve/hello"

	image := `
		FROM node
		RUN echo "(()=>{console.log('HelloWorld')})()" > node.js
		CMD ["node", "node.js"]
	`

	reader :=  bytes.NewReader([]byte(image))

	imageBuildResponse,err:=cli.ImageBuild(ctx,reader,types.ImageBuildOptions{
		Dockerfile: imageName,
	})

	if err != nil {
        log.Fatal(err, " :unable to build docker image")
    }
    defer imageBuildResponse.Body.Close()
    _, err = io.Copy(os.Stdout, imageBuildResponse.Body)
    if err != nil {
        log.Fatal(err, " :unable to read image build response")
    }


	out, err := cli.ImagePull(ctx, imageName, types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, out)

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: imageName,
		
	}, &container.HostConfig{
		Resources: container.Resources{
			Memory: 256 * 1024 * 1024,
		},
	}, nil, nil, "testing")
	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{
			
	}); err != nil {
		panic(err)
	}

	fmt.Println(resp.ID)
}