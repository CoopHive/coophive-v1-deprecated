package check
import (
	"fmt"
	"os/exec"
	"github.com/pbnjay/memory"
	"runtime"
)

type CheckOptions struct {
}

func RunDockerCommand() (string, error) {
	cmd:= exec.Command("docker", "run", "--rm", "--runtime=nvidia", "--gpus=all", "ubuntu", "nvidia-smi")
	stdout, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to run docker command: %w", err)
	}
	fmt.Printf("Total system memory: %d\n", memory.TotalMemory())
	fmt.Printf("number of cpus : %d\n", runtime.NumCPU())
	return string(stdout), nil
}