package controller

import (
	"context"
	"fmt"
	"sort"

	"git.amabanana.com/plancks-cloud/pc-go-daemon/model"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/swarm"
	"github.com/docker/docker/client"
)

//DockerListServices lists all Docker services
func DockerListServices() []model.Service {
	var services []model.Service
	services = append(services, model.Service{})
	AllDockerServices()
	return services
}

//AllDockerServices gets all running docker services
func AllDockerServices() {

	cli, err := client.NewEnvClient()

	ctx := context.Background()

	if err != nil {
		panic(fmt.Sprintf("Error getting docker client environment: %s", err))
	}

	services, err := cli.ServiceList(context.Background(), types.ServiceListOptions{})
	if err != nil {
		fmt.Printf("Error getting docker client environment: %s", err)
	}

	sort.Sort(model.ByName(services))
	if len(services) > 0 {
		// only non-empty services and not quiet, should we call TaskList and NodeList api
		taskFilter := filters.NewArgs()
		for _, service := range services {
			taskFilter.Add("service", service.ID)
		}

		tasks, err := cli.TaskList(ctx, types.TaskListOptions{Filters: taskFilter})
		if err != nil {
			fmt.Println("Error getting tasks")
		}

		nodes, err := cli.NodeList(ctx, types.NodeListOptions{})
		if err != nil {
			fmt.Println("Error getting nodes")
		}

		info := TotalReplicas(services, nodes, tasks)

		for _, item := range info {
			fmt.Printf("ID: %s\n", item.ID)
			fmt.Printf("Service: %s, Replicas: %d/%d", item.Name, item.ReplicasRunning, item.ReplicasRequired)
		}
	}
}

//TotalReplicas returns the total number of replicas running for a service
func TotalReplicas(services []swarm.Service, nodes []swarm.Node, tasks []swarm.Task) map[string]model.ServiceState {
	running := map[string]int{}
	tasksNoShutdown := map[string]int{}
	activeNodes := make(map[string]struct{})
	replicaState := make(map[string]model.ServiceState)

	for _, n := range nodes {
		if n.Status.State != swarm.NodeStateDown {
			activeNodes[n.ID] = struct{}{}
		}
	}

	for _, task := range tasks {
		if task.DesiredState != swarm.TaskStateShutdown {
			tasksNoShutdown[task.ServiceID]++
		}
		if _, nodeActive := activeNodes[task.NodeID]; nodeActive && task.Status.State == swarm.TaskStateRunning {
			running[task.ServiceID]++
		}
	}

	for _, service := range services {
		if service.Spec.Mode.Replicated != nil && service.Spec.Mode.Replicated.Replicas != nil {
			replicaState[service.ID] = model.ServiceState{
				ID:               service.ID,
				Name:             service.Spec.Name,
				ReplicasRunning:  running[service.ID],
				ReplicasRequired: *service.Spec.Mode.Replicated.Replicas}
		}
	}
	return replicaState
}
