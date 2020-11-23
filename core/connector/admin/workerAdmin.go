package admin

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/ditrit/shoset/msg"

	"github.com/ditrit/gandalf/core/configuration"
	"github.com/ditrit/gandalf/core/connector/shoset"
	"github.com/ditrit/gandalf/core/connector/utils"
	"github.com/ditrit/gandalf/libraries/goclient"

	"github.com/ditrit/gandalf/core/models"
	goclientmodels "github.com/ditrit/gandalf/libraries/goclient/models"
	net "github.com/ditrit/shoset"
	"gopkg.in/yaml.v2"
)

type WorkerAdmin struct {
	logicalName     string
	connectorType   string
	product         string
	baseurl         string
	workerPath      string
	grpcBindAddress string
	chaussette      *net.Shoset
	timeoutMax      int64
	versions        []models.Version
	clientGandalf   *goclient.ClientGandalf

	major int64

	CommandsFuncs map[string]func(clientGandalf *goclient.ClientGandalf, major int64, command msg.Command) int
	//mapVersionConfigurationKeys map[models.Version][]models.ConfigurationKeys
}

//NewWorker : NewWorker
func NewWorkerAdmin(logicalName, connectorType, product, baseurl, workerPath, grpcBindAddress string, timeoutMax int64, chaussette *net.Shoset, versions []models.Version) *WorkerAdmin {
	workerAdmin := new(WorkerAdmin)
	workerAdmin.logicalName = logicalName
	workerAdmin.connectorType = connectorType
	workerAdmin.product = product
	workerAdmin.baseurl = baseurl
	workerAdmin.workerPath = workerPath
	workerAdmin.grpcBindAddress = grpcBindAddress
	workerAdmin.timeoutMax = timeoutMax
	workerAdmin.chaussette = chaussette
	workerAdmin.versions = versions

	workerAdmin.major = 0

	workerAdmin.clientGandalf = goclient.NewClientGandalf(workerAdmin.logicalName, strconv.FormatInt(workerAdmin.timeoutMax, 10), strings.Split(workerAdmin.grpcBindAddress, ","))
	workerAdmin.CommandsFuncs = make(map[string]func(clientGandalf *goclient.ClientGandalf, major int64, command msg.Command) int)

	return workerAdmin
}

//GetClientGandalf : GetClientGandalf
func (w WorkerAdmin) GetClientGandalf() *goclient.ClientGandalf {
	return w.clientGandalf
}

//RegisterCommandsFuncs : RegisterCommandsFuncs
func (w WorkerAdmin) RegisterCommandsFuncs(command string, function func(clientGandalf *goclient.ClientGandalf, major int64, command msg.Command) int) {
	fmt.Println("REGISTER")
	w.CommandsFuncs[command] = function
}

//Run : Run
func (w WorkerAdmin) Run() {

	//GET CONFIGURATION
	w.getConfiguration()

	for _, version := range w.versions {
		err := w.getWorkerConfiguration(version)
		if err == nil {
			err = w.getWorker(version)
			if err == nil {
				go w.startWorker(version)
			}
		}
	}

	//
	w.RegisterCommandsFuncs("ADMIN_GET_WORKER", w.GetWorker)
	w.RegisterCommandsFuncs("ADMIN_START_WORKER", w.StartWorker)
	w.RegisterCommandsFuncs("ADMIN_STOP_WORKER", w.StopWorker)
	w.RegisterCommandsFuncs("ADMIN_GET_LAST_VERSION_WORKER", w.GetLastVersionsWorker)

	for key, function := range w.CommandsFuncs {
		id := w.clientGandalf.CreateIteratorCommand()

		go w.waitCommands(id, key, function)
	}
	//TODO REVOIR CONDITION SORTIE
	for true {
		time.Sleep(1 * time.Millisecond)
	}
	fmt.Println("END WORKER ADMIN")
}

func (w WorkerAdmin) waitCommands(id, commandName string, function func(clientGandalf *goclient.ClientGandalf, major int64, command msg.Command) int) {

	for true {

		fmt.Println("wait " + commandName)
		command := w.clientGandalf.WaitCommand(commandName, id, w.major)
		fmt.Println("command")
		fmt.Println(command)

		go w.executeCommands(command, function)

	}
	fmt.Println("END WAIT")
}

func (w WorkerAdmin) executeCommands(command msg.Command, function func(clientGandalf *goclient.ClientGandalf, major int64, command msg.Command) int) {
	fmt.Println("execute")
	result := function(w.clientGandalf, w.major, command)
	if result == 0 {
		w.clientGandalf.SendReply(command.GetCommand(), "SUCCES", command.GetUUID(), goclientmodels.NewOptions("", ""))
	} else {
		w.clientGandalf.SendReply(command.GetCommand(), "FAIL", command.GetUUID(), goclientmodels.NewOptions("", ""))
	}
}

//
func (w WorkerAdmin) StopWorker(clientGandalf *goclient.ClientGandalf, major int64, command msg.Command) int {
	var versionPayload models.Version
	fmt.Println("COMMAND")
	fmt.Println(command)
	fmt.Println("PAYLOAD")
	fmt.Println(command.GetPayload())
	err := json.Unmarshal([]byte(command.GetPayload()), &versionPayload)
	fmt.Println("ERROR STOP WORKER")
	fmt.Println(err)
	if err == nil {
		err = w.stopWorker(versionPayload)
		if err == nil {
			return 0
		}
	}

	return 1
}

//
func (w WorkerAdmin) GetWorker(clientGandalf *goclient.ClientGandalf, major int64, command msg.Command) int {
	var versionPayload models.Version
	fmt.Println("COMMAND")
	fmt.Println(command)
	fmt.Println("PAYLOAD")
	fmt.Println(command.GetPayload())
	err := json.Unmarshal([]byte(command.GetPayload()), &versionPayload)
	fmt.Println("ERROR STOP WORKER")
	fmt.Println(err)
	if err == nil {
		err = w.getWorkerConfiguration(versionPayload)
		if err == nil {
			err = w.getWorker(versionPayload)
			if err == nil {
				return 0
			}
		}
	}

	return 1
}

func (w WorkerAdmin) GetLastVersionsWorker(clientGandalf *goclient.ClientGandalf, major int64, command msg.Command) int {

	lastVersion, err := w.getLastVersion()
	fmt.Println("lastVersion")
	fmt.Println(lastVersion)
	if err == nil {
		err = w.getWorkerConfiguration(lastVersion)
		if err == nil {
			err = w.getWorker(lastVersion)
			if err == nil {
				return 0
			}
		}
	}

	return 1
}

func (w WorkerAdmin) StartWorker(clientGandalf *goclient.ClientGandalf, major int64, command msg.Command) int {
	var versionPayload models.Version
	err := json.Unmarshal([]byte(command.GetPayload()), &versionPayload)

	if err == nil {
		err = w.startWorker(versionPayload)
		if err == nil {
			return 0
		}
	}
	return 1
}

//stopWorker()
func (w WorkerAdmin) stopWorker(version models.Version) (err error) {
	activeWorkers := w.chaussette.Context["mapActiveWorkers"].(map[models.Version]bool)
	activeWorkers[version] = false
	w.chaussette.Context["mapActiveWorkers"] = activeWorkers

	return
}

//getConfiguration()
// GetKeys : Get keys from baseurl/connectorType/ and baseurl/connectorType/product/
func (w WorkerAdmin) getConfiguration() (err error) {

	shoset.SendConnectorConfig(w.chaussette, w.timeoutMax)
	time.Sleep(time.Second * time.Duration(5))

	config := w.chaussette.Context["mapConnectorsConfig"].(map[string][]*models.ConnectorConfig)
	fmt.Println("config")
	fmt.Println(config)
	if config != nil {
		connectorConfig := utils.GetConnectorTypeConfigByVersion(int8(w.major), config["Admin"])
		if connectorConfig == nil {
			fmt.Println("DOWNLOAD")

			dir, err := os.Getwd()
			fmt.Println("dir")
			fmt.Println(dir)
			dat, err := ioutil.ReadFile(dir + "/connector/admin/configuration.yaml")

			fmt.Print("string(dat)")
			fmt.Print(string(dat))

			err = yaml.Unmarshal(dat, &connectorConfig)
			if err != nil {
				fmt.Println(err)
				log.Fatal(err)
			}

			fmt.Println("connectorConfig")
			fmt.Println(connectorConfig)
			connectorConfig.ConnectorType.Name = "Admin"
			connectorConfig.Major = int8(w.major)

			//ADD COMMMANDS ADMIN
			//addCommandsAdmin(connectorConfig)

			shoset.SendSaveConnectorConfig(w.chaussette, w.timeoutMax, connectorConfig)
		}

		config[w.connectorType] = append(config[w.connectorType], connectorConfig)
		w.chaussette.Context["mapConnectorsConfig"] = config
	}

	return
}

//getConfiguration()
// GetKeys : Get keys from baseurl/connectorType/ and baseurl/connectorType/product/
func (w WorkerAdmin) getWorkerConfiguration(version models.Version) (err error) {

	shoset.SendConnectorConfig(w.chaussette, w.timeoutMax)
	time.Sleep(time.Second * time.Duration(5))

	config := w.chaussette.Context["mapConnectorsConfig"].(map[string][]*models.ConnectorConfig)
	fmt.Println("config")
	fmt.Println(config)
	if config != nil {

		configConnectorTypeKeys, _ := utils.DownloadConfigurationsKeys(w.baseurl, "/"+strings.ToLower(w.connectorType)+"/keys.yaml")
		configProductKeys, _ := utils.DownloadConfigurationsKeys(w.baseurl, "/"+strings.ToLower(w.connectorType)+"/"+strings.ToLower(w.product)+"/keys.yaml")

		connectorConfig := utils.GetConnectorTypeConfigByVersion(version.Major, config[w.connectorType])
		if connectorConfig == nil {
			fmt.Println("DOWNLOAD")

			fmt.Println("url")
			fmt.Println(w.baseurl, "/"+strings.ToLower(w.connectorType)+"/"+strings.ToLower(w.product)+"/"+strconv.Itoa(int(version.Major))+"_configuration.yaml")

			connectorConfig, _ = utils.DownloadConfiguration(w.baseurl, "/"+strings.ToLower(w.connectorType)+"/"+strings.ToLower(w.product)+"/"+strconv.Itoa(int(version.Major))+"_configuration.yaml")
			fmt.Println("connectorConfig")
			fmt.Println(connectorConfig)
			connectorConfig.ConnectorType.Name = w.connectorType
			connectorConfig.Major = version.Major

			connectorConfig.ConnectorTypeKeys = configConnectorTypeKeys
			connectorConfig.ProductKeys = configProductKeys

			connectorConfig.VersionMajorKeys, _ = utils.DownloadConfigurationsKeys(w.baseurl, "/"+strings.ToLower(w.connectorType)+"/"+strings.ToLower(w.product)+"/"+strconv.Itoa(int(version.Major))+"_keys.yaml")
			connectorConfig.VersionMinorKeys, _ = utils.DownloadConfigurationsKeys(w.baseurl, "/"+strings.ToLower(w.connectorType)+"/"+strings.ToLower(w.product)+"/"+strconv.Itoa(int(version.Major))+"_"+strconv.Itoa(int(version.Minor))+"_keys.yaml")

			//ADD COMMMANDS ADMIN
			//addCommandsAdmin(connectorConfig)

			shoset.SendSaveConnectorConfig(w.chaussette, w.timeoutMax, connectorConfig)
		}

		config[w.connectorType] = append(config[w.connectorType], connectorConfig)
		w.chaussette.Context["mapConnectorsConfig"] = config
	}

	return
}

/* func addCommandsAdmin(config *models.ConnectorConfig) {

	schemaVersion := `{"$schema": "http://json-schema.org/draft-04/schema#","type": "object","properties": {"Major": { "type": "integer" },"Minor": { "type": "integer" }},"required": ["Major","Minor"]}`
	schemaString := `{"type":"string"}`
	actionExecute := models.Action{Name: "Execute"}

	commandAdminGetWorker := models.Object{Name: "ADMIN_GET_WORKER", Schema: schemaVersion, Actions: []models.Action{actionExecute}}
	commandAdminStartWorker := models.Object{Name: "ADMIN_START_WORKER", Schema: schemaVersion, Actions: []models.Action{actionExecute}}
	commandAdminStopWorker := models.Object{Name: "ADMIN_STOP_WORKER", Schema: schemaVersion, Actions: []models.Action{actionExecute}}
	commandAdminGetLastVersionWorker := models.Object{Name: "ADMIN_GET_LAST_VERSION_WORKER", Schema: schemaString, Actions: []models.Action{actionExecute}}

	config.ConnectorCommands = append(config.ConnectorCommands, commandAdminGetWorker)
	config.ConnectorCommands = append(config.ConnectorCommands, commandAdminStartWorker)
	config.ConnectorCommands = append(config.ConnectorCommands, commandAdminStopWorker)
	config.ConnectorCommands = append(config.ConnectorCommands, commandAdminGetLastVersionWorker)
} */

//getWorker()
func (w WorkerAdmin) getWorker(version models.Version) (err error) {

	ressourceDir := "/" + strings.ToLower(w.connectorType) + "/" + strings.ToLower(w.product) + "/" + strconv.Itoa(int(version.Major)) + "/" + strconv.Itoa(int(version.Minor)) + "/"
	fileWorkersPathVersion := w.workerPath + ressourceDir + "worker"

	if !utils.CheckFileExistAndIsExecAll(fileWorkersPathVersion) {
		fmt.Println("DOWNLOAD")
		ressourceURL := "/" + strings.ToLower(w.connectorType) + "/" + strings.ToLower(w.product) + "/" + strconv.Itoa(int(version.Major)) + "_" + strconv.Itoa(int(version.Minor)) + "_"

		url := w.baseurl + ressourceURL + "worker.zip"
		fmt.Println("url")
		fmt.Println(url)
		src := w.workerPath + ressourceDir + "worker.zip"
		dest := w.workerPath + ressourceDir

		if _, err := os.Stat(dest); os.IsNotExist(err) {
			os.MkdirAll(dest, os.ModePerm)
		}

		err = utils.DownloadWorkers(url, src)

		if err == nil {
			_, err = utils.Unzip(src, dest)
			if err != nil {
				log.Println("Can't unzip workers")
			}
		} else {
			log.Println("Can't download workers")
		}
	}

	return
}

//startWorker()
func (w WorkerAdmin) startWorker(version models.Version) (err error) {

	config := w.chaussette.Context["mapConnectorsConfig"].(map[string][]*models.ConnectorConfig)
	fmt.Println("config")
	fmt.Println(config)
	if config != nil {
		connectorConfig := utils.GetConnectorTypeConfigByVersion(version.Major, config[w.connectorType])

		if connectorConfig != nil {

			var listConfigurationKeys []models.ConfigurationKeys

			var listConfigurationConnectorTypeKeys []models.ConfigurationKeys
			err = yaml.Unmarshal([]byte(connectorConfig.ConnectorTypeKeys), &listConfigurationConnectorTypeKeys)
			if err != nil {
				fmt.Println(err)
			}

			var listConfigurationProductKeys []models.ConfigurationKeys
			err = yaml.Unmarshal([]byte(connectorConfig.ProductKeys), &listConfigurationProductKeys)
			if err != nil {
				fmt.Println(err)
			}
			var listConfigurationVersionMajorKeys []models.ConfigurationKeys
			err = yaml.Unmarshal([]byte(connectorConfig.VersionMajorKeys), &listConfigurationVersionMajorKeys)
			if err != nil {
				fmt.Println(err)
			}

			var listConfigurationVersionMinorKeys []models.ConfigurationKeys
			err = yaml.Unmarshal([]byte(connectorConfig.VersionMinorKeys), &listConfigurationVersionMinorKeys)
			if err != nil {
				fmt.Println(err)
			}

			listConfigurationKeys = append(listConfigurationKeys, listConfigurationConnectorTypeKeys...)
			listConfigurationKeys = append(listConfigurationKeys, listConfigurationProductKeys...)
			listConfigurationKeys = append(listConfigurationKeys, listConfigurationVersionMajorKeys...)
			listConfigurationKeys = append(listConfigurationKeys, listConfigurationVersionMinorKeys...)

			configuration.WorkerKeyParse(listConfigurationKeys)
			err = configuration.IsConfigValid()
			if err == nil {
				fmt.Println("listConfigurationKeys")
				fmt.Println(listConfigurationKeys)

				var stdinargs string
				stdinargs = utils.GetConfigurationKeys(listConfigurationKeys)
				fmt.Println("stdinargs")
				fmt.Println(stdinargs)

				workersPathVersion := w.workerPath + "/" + strings.ToLower(w.connectorType) + "/" + strings.ToLower(w.product) + "/" + strconv.Itoa(int(version.Major)) + "/" + strconv.Itoa(int(version.Minor))
				fileWorkersPathVersion := workersPathVersion + "/worker"

				if utils.CheckFileExistAndIsExecAll(fileWorkersPathVersion) {
					args := []string{w.logicalName, strconv.FormatInt(w.timeoutMax, 10), w.grpcBindAddress}

					cmd := exec.Command("./worker", args...)
					cmd.Dir = workersPathVersion
					cmd.Stdout = os.Stdout

					stdin, err := cmd.StdinPipe()
					if err != nil {
						fmt.Println(err)
					}

					err = cmd.Start()
					if err != nil {
						log.Printf("Can't start worker %s", fileWorkersPathVersion)
					}
					time.Sleep(time.Second * time.Duration(5))

					go func() {
						defer stdin.Close()
						fmt.Println("Write")
						io.WriteString(stdin, stdinargs)
					}()
				}
			}

		}
	}

	return
}

func (w WorkerAdmin) getLastVersion() (lastVersion models.Version, err error) {
	versions, _ := utils.DownloadVersions(w.baseurl, "/"+strings.ToLower(w.connectorType)+"/"+strings.ToLower(w.product)+"/versions.yaml")
	fmt.Println("versions")
	fmt.Println(versions)

	versionSplit := strings.Split(versions[len(versions)-1], ".")
	major8, err := strconv.ParseInt(versionSplit[0], 10, 8)
	minor8, err := strconv.ParseInt(versionSplit[1], 10, 8)
	lastVersion.Major = int8(major8)
	lastVersion.Minor = int8(minor8)

	return
}
