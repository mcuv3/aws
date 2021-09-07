package main

import (
	"archive/tar"
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/MauricioAntonioMartinez/aws/docker"
	aws "github.com/MauricioAntonioMartinez/aws/proto"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/golang/protobuf/proto"
)

func d() {

	dis := docker.NewContainerDispatcher(3, &docker.DockerRuntime{})

	dis.Start()

	dis.BuildImage(docker.BuildImageOptions{})
	dis.BuildImage(docker.BuildImageOptions{})
	dis.BuildImage(docker.BuildImageOptions{})
	dis.BuildImage(docker.BuildImageOptions{})
	dis.BuildImage(docker.BuildImageOptions{})
	dis.BuildImage(docker.BuildImageOptions{})
	dis.BuildImage(docker.BuildImageOptions{})
	dis.BuildImage(docker.BuildImageOptions{})
	dis.BuildImage(docker.BuildImageOptions{})
	dis.BuildImage(docker.BuildImageOptions{})
	dis.BuildImage(docker.BuildImageOptions{})
	dis.BuildImage(docker.BuildImageOptions{})
	dis.BuildImage(docker.BuildImageOptions{})
	dis.BuildImage(docker.BuildImageOptions{})
	dis.BuildImage(docker.BuildImageOptions{})
	dis.BuildImage(docker.BuildImageOptions{})
	dis.BuildImage(docker.BuildImageOptions{})
	dis.BuildImage(docker.BuildImageOptions{})
	dis.BuildImage(docker.BuildImageOptions{})
	dis.BuildImage(docker.BuildImageOptions{})

	time.Sleep(time.Second * 15)

}

func main() {
	res := aws.CreateUserRequest{
		Name:        "Test",
		Description: "New",
		Password:    "123",
		Polices:     []*aws.Policy{},
	}
	value, err := proto.Marshal(&res)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(value)
	newRes := &aws.CreateUserRequest{}
	err = proto.Unmarshal(value, newRes)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(newRes)

}

func dock() {
	ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		log.Fatal(err, " :unable to init client")
	}

	buf := new(bytes.Buffer)
	tw := tar.NewWriter(buf)
	defer tw.Close()

	dockerF := []byte(`
	FROM node
	RUN echo "(()=>{console.log('HelloWorld')})()" > node.js
	CMD ["node", "node.js"]
	`)

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
			Tags:       []string{"test"},
			Remove:     false})
	if err != nil {
		log.Fatal(err, " :unable to build docker image")
	}
	defer imageBuildResponse.Body.Close()
	_, err = io.Copy(os.Stdout, imageBuildResponse.Body)
	if err != nil {
		log.Fatal(err, " :unable to read image build response")
	}
}
