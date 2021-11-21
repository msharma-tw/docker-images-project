package ruby

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/docker"
	"github.com/stretchr/testify/assert"
)

func TestSample(t *testing.T) {
	//Below method is known as buildDockerImage Method

	// runs in parallel if we have multiple tests
	// t.Parallel()

	// Configure the tag to use on the Docker image.
	tag := "testruby"
	buildOptions := &docker.BuildOptions{
		Tags: []string{tag},
	}

	// Build the Docker image.
	path := "../../images/ruby/2.6-builder-ubi8"
	docker.Build(t, path, buildOptions)

	// Run the Docker image, execute the command, and make sure it contains the expected output.
	opts := &docker.RunOptions{Command: []string{"ruby", "--version"}}
	output := docker.Run(t, tag, opts)
	assert.Equal(t, "ruby 2.6.7p197 (2021-04-05 revision 67941) [x86_64-linux]", output)
}