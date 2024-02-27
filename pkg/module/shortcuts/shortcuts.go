package shortcuts

import (
	"fmt"
	"strings"

	"github.com/CoopHive/coophive/pkg/data"
)

// allow shortcode github.com/lukemarsden/coophive-sdxl:v0.0.1 (tag),
// TODO: enforce sha1 for tags on the server side (like a pin file)

// parse something with no slashes in it as
// github.com/coophive/coophive-module-<shortcode>

const COOPHIVE_MODULE_CONFIG_PATH = "/module.coophive"

func GetModule(name string) (data.ModuleConfig, error) {
	// parse name per following valid formats
	// github.com/user/repo:tag --> Repo: https://github.com/user/repo; Hash = tag
	// bar:tag --> Repo = https://github.com/coophive/coophive-module-<bar>, Hash = tag
	if name == "" {
		return data.ModuleConfig{}, fmt.Errorf("module name is empty")
	}
	parts := strings.Split(name, ":")
	if len(parts) != 2 {
		return data.ModuleConfig{}, fmt.Errorf("invalid module name format: %s", name)
	}
	repo, hash := parts[0], parts[1]
	if strings.Contains(name, "/") {
		// 3rd party module
		repo = fmt.Sprintf("https://%s", repo)
	} else {
		// coophive std module
		repo = fmt.Sprintf("https://github.com/coophive/coophive-module-%s", repo)
	}

	fmt.Println("+====================")

	// TODO: docs for authoring a module
	module := data.ModuleConfig{
		Name: "", // TODO:
		Repo: repo,
		Hash: hash,
		Path: COOPHIVE_MODULE_CONFIG_PATH,
	}

	return module, nil
}
