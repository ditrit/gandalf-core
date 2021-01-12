package api

// Urls :
type Urls struct {
	STATIC_PATH  string
	ROOT_PATH    string
	GANDALF_PATH string
	TENANTS_PATH string

	GANDALF_LOGIN_PATH                        string
	GANDALF_CLUSTER_PATH                      string
	GANDALF_CLUSTER_PATH_LIST                 string
	GANDALF_CLUSTER_PATH_CREATE               string
	GANDALF_CLUSTER_PATH_READ                 string
	GANDALF_CLUSTER_PATH_UPDATE               string
	GANDALF_CLUSTER_PATH_DELETE               string
	GANDALF_ROLE_PATH                         string
	GANDALF_ROLE_PATH_LIST                    string
	GANDALF_ROLE_PATH_CREATE                  string
	GANDALF_ROLE_PATH_READ                    string
	GANDALF_ROLE_PATH_UPDATE                  string
	GANDALF_ROLE_PATH_DELETE                  string
	GANDALF_USER_PATH                         string
	GANDALF_USER_PATH_LIST                    string
	GANDALF_USER_PATH_CREATE                  string
	GANDALF_USER_PATH_READ                    string
	GANDALF_USER_PATH_UPDATE                  string
	GANDALF_USER_PATH_DELETE                  string
	GANDALF_TENANT_PATH                       string
	GANDALF_TENANT_PATH_LIST                  string
	GANDALF_TENANT_PATH_CREATE                string
	GANDALF_TENANT_PATH_READ                  string
	GANDALF_TENANT_PATH_UPDATE                string
	GANDALF_TENANT_PATH_DELETE                string
	GANDALF_CONFIGURATION_PATH                string
	GANDALF_CONFIGURATION_CLUSTER_PATH_UPLOAD string
	GANDALF_CONFIGURATION_CLUSTER_PATH_READ   string

	TENANTS_LOGIN_PATH                           string
	TENANTS_CONNECTOR_PATH                       string
	TENANTS_CONNECTOR_PATH_LIST                  string
	TENANTS_CONNECTOR_PATH_CREATE                string
	TENANTS_CONNECTOR_PATH_READ                  string
	TENANTS_CONNECTOR_PATH_UPDATE                string
	TENANTS_CONNECTOR_PATH_DELETE                string
	TENANTS_AGGREGATOR_PATH                      string
	TENANTS_AGGREGATOR_PATH_LIST                 string
	TENANTS_AGGREGATOR_PATH_CREATE               string
	TENANTS_AGGREGATOR_PATH_READ                 string
	TENANTS_AGGREGATOR_PATH_UPDATE               string
	TENANTS_AGGREGATOR_PATH_DELETE               string
	TENANTS_ROLE_PATH                            string
	TENANTS_ROLE_PATH_LIST                       string
	TENANTS_ROLE_PATH_CREATE                     string
	TENANTS_ROLE_PATH_READ                       string
	TENANTS_ROLE_PATH_UPDATE                     string
	TENANTS_ROLE_PATH_DELETE                     string
	TENANTS_USER_PATH                            string
	TENANTS_USER_PATH_LIST                       string
	TENANTS_USER_PATH_CREATE                     string
	TENANTS_USER_PATH_READ                       string
	TENANTS_USER_PATH_UPDATE                     string
	TENANTS_USER_PATH_DELETE                     string
	TENANTS_CONFIGURATION_PATH                   string
	TENANTS_CONFIGURATION_AGGREGATOR_PATH_UPLOAD string
	TENANTS_CONFIGURATION_AGGREGATOR_PATH_READ   string
	TENANTS_CONFIGURATION_CONNECTOR_PATH_UPLOAD  string
	TENANTS_CONFIGURATION_CONNECTOR_PATH_READ    string
}

// ReturnURLS :
func ReturnURLS() *Urls {

	//BASE
	apiurls := new(Urls)
	apiurls.ROOT_PATH = "/"
	apiurls.GANDALF_PATH = "/gandalf"
	apiurls.TENANTS_PATH = "/tenants/{tenant}"

	//GANDALF
	apiurls.GANDALF_LOGIN_PATH = apiurls.GANDALF_PATH + "/login/"
	apiurls.GANDALF_CLUSTER_PATH = apiurls.GANDALF_PATH + "/clusters"
	apiurls.GANDALF_CLUSTER_PATH_LIST = apiurls.GANDALF_CLUSTER_PATH + "/"
	apiurls.GANDALF_CLUSTER_PATH_CREATE = apiurls.GANDALF_CLUSTER_PATH + "/"
	apiurls.GANDALF_CLUSTER_PATH_READ = apiurls.GANDALF_CLUSTER_PATH + "/{id}"
	apiurls.GANDALF_CLUSTER_PATH_UPDATE = apiurls.GANDALF_CLUSTER_PATH + "/{id:[0-9]+}"
	apiurls.GANDALF_CLUSTER_PATH_DELETE = apiurls.GANDALF_CLUSTER_PATH + "/{id:[0-9]+}"
	apiurls.GANDALF_ROLE_PATH = apiurls.GANDALF_PATH + "/roles"
	apiurls.GANDALF_ROLE_PATH_LIST = apiurls.GANDALF_ROLE_PATH + "/"
	apiurls.GANDALF_ROLE_PATH_CREATE = apiurls.GANDALF_ROLE_PATH + "/"
	apiurls.GANDALF_ROLE_PATH_READ = apiurls.GANDALF_ROLE_PATH + "/{id:[0-9]+}"
	apiurls.GANDALF_ROLE_PATH_UPDATE = apiurls.GANDALF_ROLE_PATH + "/{id:[0-9]+}"
	apiurls.GANDALF_ROLE_PATH_DELETE = apiurls.GANDALF_ROLE_PATH + "/{id:[0-9]+}"
	apiurls.GANDALF_USER_PATH = apiurls.GANDALF_PATH + "/users"
	apiurls.GANDALF_USER_PATH_LIST = apiurls.GANDALF_USER_PATH + "/"
	apiurls.GANDALF_USER_PATH_CREATE = apiurls.GANDALF_USER_PATH + "/"
	apiurls.GANDALF_USER_PATH_READ = apiurls.GANDALF_USER_PATH + "/{id:[0-9]+}"
	apiurls.GANDALF_USER_PATH_UPDATE = apiurls.GANDALF_USER_PATH + "/{id:[0-9]+}"
	apiurls.GANDALF_USER_PATH_DELETE = apiurls.GANDALF_USER_PATH + "/{id:[0-9]+}"
	apiurls.GANDALF_TENANT_PATH = apiurls.GANDALF_PATH + "/tenants"
	apiurls.GANDALF_TENANT_PATH_LIST = apiurls.GANDALF_TENANT_PATH + "/"
	apiurls.GANDALF_TENANT_PATH_CREATE = apiurls.GANDALF_TENANT_PATH + "/"
	apiurls.GANDALF_TENANT_PATH_READ = apiurls.GANDALF_TENANT_PATH + "/{id:[0-9]+}"
	apiurls.GANDALF_TENANT_PATH_UPDATE = apiurls.GANDALF_TENANT_PATH + "/{id:[0-9]+}"
	apiurls.GANDALF_TENANT_PATH_DELETE = apiurls.GANDALF_TENANT_PATH + "/{id:[0-9]+}"
	apiurls.GANDALF_CONFIGURATION_PATH = apiurls.GANDALF_PATH + "/configurations"
	apiurls.GANDALF_CONFIGURATION_CLUSTER_PATH_UPLOAD = apiurls.GANDALF_CONFIGURATION_PATH + "/"
	apiurls.GANDALF_CONFIGURATION_CLUSTER_PATH_READ = apiurls.GANDALF_CONFIGURATION_PATH + "/{id:[0-9]+}"

	//TENANTS
	apiurls.TENANTS_LOGIN_PATH = apiurls.TENANTS_PATH + "/login/"
	apiurls.TENANTS_CONNECTOR_PATH = apiurls.TENANTS_PATH + "/connectors"
	apiurls.TENANTS_CONNECTOR_PATH_LIST = apiurls.TENANTS_CONNECTOR_PATH + "/"
	apiurls.TENANTS_CONNECTOR_PATH_CREATE = apiurls.TENANTS_CONNECTOR_PATH + "/"
	apiurls.TENANTS_CONNECTOR_PATH_READ = apiurls.TENANTS_CONNECTOR_PATH + "/{id:[0-9]+}"
	apiurls.TENANTS_CONNECTOR_PATH_UPDATE = apiurls.TENANTS_CONNECTOR_PATH + "/{id:[0-9]+}"
	apiurls.TENANTS_CONNECTOR_PATH_DELETE = apiurls.TENANTS_CONNECTOR_PATH + "/{id:[0-9]+}"
	apiurls.TENANTS_AGGREGATOR_PATH = apiurls.TENANTS_PATH + "/aggregators"
	apiurls.TENANTS_AGGREGATOR_PATH_LIST = apiurls.TENANTS_AGGREGATOR_PATH + "/"
	apiurls.TENANTS_AGGREGATOR_PATH_CREATE = apiurls.TENANTS_AGGREGATOR_PATH + "/"
	apiurls.TENANTS_AGGREGATOR_PATH_READ = apiurls.TENANTS_AGGREGATOR_PATH + "/{id:[0-9]+}"
	apiurls.TENANTS_AGGREGATOR_PATH_UPDATE = apiurls.TENANTS_AGGREGATOR_PATH + "/{id:[0-9]+}"
	apiurls.TENANTS_AGGREGATOR_PATH_DELETE = apiurls.TENANTS_AGGREGATOR_PATH + "/{id:[0-9]+}"
	apiurls.TENANTS_ROLE_PATH = apiurls.TENANTS_PATH + "/roles"
	apiurls.TENANTS_ROLE_PATH_LIST = apiurls.TENANTS_ROLE_PATH + "/"
	apiurls.TENANTS_ROLE_PATH_CREATE = apiurls.TENANTS_ROLE_PATH + "/"
	apiurls.TENANTS_ROLE_PATH_READ = apiurls.TENANTS_ROLE_PATH + "/{id:[0-9]+}"
	apiurls.TENANTS_ROLE_PATH_UPDATE = apiurls.TENANTS_ROLE_PATH + "/{id:[0-9]+}"
	apiurls.TENANTS_ROLE_PATH_DELETE = apiurls.TENANTS_ROLE_PATH + "/{id:[0-9]+}"
	apiurls.TENANTS_USER_PATH = apiurls.TENANTS_PATH + "/users"
	apiurls.TENANTS_USER_PATH_LIST = apiurls.TENANTS_USER_PATH + "/"
	apiurls.TENANTS_USER_PATH_CREATE = apiurls.TENANTS_USER_PATH + "/"
	apiurls.TENANTS_USER_PATH_READ = apiurls.TENANTS_USER_PATH + "/{id:[0-9]+}"
	apiurls.TENANTS_USER_PATH_UPDATE = apiurls.TENANTS_USER_PATH + "/{id:[0-9]+}"
	apiurls.TENANTS_USER_PATH_DELETE = apiurls.TENANTS_USER_PATH + "/{id:[0-9]+}"
	apiurls.TENANTS_CONFIGURATION_PATH = apiurls.TENANTS_PATH + "/configurations"
	apiurls.TENANTS_CONFIGURATION_AGGREGATOR_PATH_UPLOAD = apiurls.TENANTS_CONFIGURATION_PATH + "/"
	apiurls.TENANTS_CONFIGURATION_AGGREGATOR_PATH_READ = apiurls.TENANTS_CONFIGURATION_PATH + "/{id:[0-9]+}"
	apiurls.TENANTS_CONFIGURATION_CONNECTOR_PATH_UPLOAD = apiurls.TENANTS_CONFIGURATION_PATH + "/"
	apiurls.TENANTS_CONFIGURATION_CONNECTOR_PATH_READ = apiurls.TENANTS_CONFIGURATION_PATH + "/{id:[0-9]+}"

	return apiurls
}
