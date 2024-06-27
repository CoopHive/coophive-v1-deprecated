package check
import (
	"fmt"
	"os/exec"
	"github.com/pbnjay/memory"
	_ "github.com/davecgh/go-spew/spew"
	"runtime"
	"encoding/json"
)

type CheckOptions struct {
}

func RunDockerCommand() error  {
	gpu := ""
	gpucmd := exec.Command("nvidia-smi", "--query-gpu=name,uuid,temperature.gpu,utilization.gpu,memory.used,utilization.memory,timestamp", "--format=json")
	_, err := gpucmd.Output()
	if err != nil {
		gpu = "none"
	} else {
		cmd:= exec.Command("docker", "run", "--rm", "--runtime=nvidia", "--gpus=all", "ubuntu", "nvidia-smi")
		stdout, err := cmd.Output()
		if err != nil {
			gpu = "please install nvidia container runtime and cuda"
		} else {
			gpu = string(stdout)
		}
	}
	jsonBytes, err := json.Marshal( struct {
		Gpu string `json:"gpu"`
		Cpu int   `json:"milli-cpu"`
		Ram uint64 `json:"ram"`
	}{
		Gpu: gpu,
		Cpu: runtime.NumCPU() * 1000,
		Ram: memory.TotalMemory(),
	})
	if err != nil {
		return err
	}
	fmt.Println(string(jsonBytes))
	return nil
}