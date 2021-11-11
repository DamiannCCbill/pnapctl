package rancher

import (
	"context"
	"net/http"

	ranchersdk "gitlab.com/phoenixnap/bare-metal-cloud/go-sdk.git/ranchersolutionapi"
	"golang.org/x/oauth2/clientcredentials"
	configuration "phoenixnap.com/pnap-cli/configs"
)

var Client RancherSdkClient

type RancherSdkClient interface {
	ClusterPost(clusterCreate ranchersdk.Cluster) (ranchersdk.Cluster, *http.Response, error)
	ClustersGet() ([]ranchersdk.Cluster, *http.Response, error)
	ClusterGetById(clusterId string) (ranchersdk.Cluster, *http.Response, error)
	ClusterDelete(clusterId string) (ranchersdk.DeleteResult, *http.Response, error)
}

type MainClient struct {
	RancherSdkClient ranchersdk.DefaultApi
}

func NewMainClient(clientId string, clientSecret string) RancherSdkClient {
	rancherConfiguration := ranchersdk.NewConfiguration()

	config := clientcredentials.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		TokenURL:     configuration.TokenURL,
		Scopes:       []string{"bmc", "bmc.read"},
	}

	rancherConfiguration.HTTPClient = config.Client(context.Background())

	api_client := ranchersdk.NewAPIClient(rancherConfiguration)

	return MainClient{
		RancherSdkClient: api_client.DefaultApi,
	}
}

func (m MainClient) ClusterPost(cluster ranchersdk.Cluster) (ranchersdk.Cluster, *http.Response, error) {
	return m.RancherSdkClient.ClustersPost(context.Background()).Cluster(cluster).Execute()
}

func (m MainClient) ClustersGet() ([]ranchersdk.Cluster, *http.Response, error) {
	return m.RancherSdkClient.ClustersGet(context.Background()).Execute()
}

func (m MainClient) ClusterGetById(clusterId string) (ranchersdk.Cluster, *http.Response, error) {
	return m.RancherSdkClient.ClustersIdGet(context.Background(), clusterId).Execute()
}

func (m MainClient) ClusterDelete(clusterId string) (ranchersdk.DeleteResult, *http.Response, error) {
	return m.RancherSdkClient.ClustersIdDelete(context.Background(), clusterId).Execute()
}
