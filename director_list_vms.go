package gogobosh

import (
	"fmt"
)

func (repo BoshDirectorRepository) ListVMs(deploymentName string) (deploymentVMs []DeploymentVM, apiResponse ApiResponse) {
	resources := []DeploymentVMResponse{}

	path := fmt.Sprintf("/deployments/%s/vms", deploymentName)
	apiResponse = repo.gateway.GetResource(repo.config.TargetURL+path, repo.config.Username, repo.config.Password, &resources)
	if apiResponse.IsNotSuccessful() {
		return
	}

	for _, resource := range resources {
		deploymentVMs = append(deploymentVMs, resource.ToModel())
	}

	return
}

type DeploymentVMResponse struct {
	JobName string  `json:"job"`
	Index int       `json:"index"`
	VMCid string    `json:"cid"`
	AgentID string  `json:"agent_id"`
}

func (resource DeploymentVMResponse) ToModel() (vm DeploymentVM) {
	vm = DeploymentVM{}
	vm.JobName  = resource.JobName
	vm.Index    = resource.Index
	vm.VMCid    = resource.VMCid
	vm.AgentID  = resource.AgentID

	return
}