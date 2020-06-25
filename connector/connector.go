//Package connector : Main function for connector
package connector

import (
	"fmt"
	//"github.com/ditrit/gandalf-core/configuration"
	"github.com/ditrit/gandalf-core/connector/grpc"
	"github.com/ditrit/gandalf-core/connector/shoset"
	"github.com/ditrit/gandalf-core/connector/utils"
	coreLog "github.com/ditrit/gandalf-core/log"
	"github.com/ditrit/gandalf-core/models"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"reflect"
	net "github.com/ditrit/shoset"
	"strconv"
	"time"
)

// ConnectorMember : Connector struct.
type ConnectorMember struct {
	chaussette        *net.Shoset
	connectorGrpc     grpc.ConnectorGrpc
	connectorType     string
	timeoutMax        int64
	connectorConfig   *models.ConnectorConfig
	connectorCommands []string
}
/*
func InitConnectorKeys(){
	_ = configuration.SetStringKeyConfig("connector","tenant","t","tenant1","tenant of the connector")
	_ = configuration.SetStringKeyConfig("connector","category","c","svn","category of the connector")
	_ = configuration.SetStringKeyConfig("connector", "product","p","product1","product of the connector")
	_ = configuration.SetStringKeyConfig("connector","aggregators", "a","address1:9800,address2:6400,address3","aggregators addresses linked to the connector")
	_ = configuration.SetStringKeyConfig("connector","gandalf_secret","s","/etc/gandalf/gandalfSecret","path of the gandalf secret")
	_ = configuration.SetStringKeyConfig("connector","product_url","u","url1,url2,url3","product url list of the connector")
	_ = configuration.SetStringKeyConfig("connector","connector_log","","/etc/gandalf/log","path of the log file")
	_ = configuration.SetIntegerKeyConfig("connector","max_timeout","",100,"maximum timeout of the connector")
}*/

// NewConnectorMember : Connector struct constructor.
func NewConnectorMember(logicalName, tenant, connectorType, logPath string) *ConnectorMember {

	member := new(ConnectorMember)
	member.connectorType = connectorType
	member.chaussette = net.NewShoset(logicalName, "c")
	member.chaussette.Context["tenant"] = tenant
	member.chaussette.Context["connectorType"] = connectorType
	member.chaussette.Context["connectorConfig"] = member.connectorConfig
	member.chaussette.Context["connectorCommands"] = member.connectorCommands
	member.chaussette.Handle["cfgjoin"] = shoset.HandleConfigJoin
	member.chaussette.Handle["cmd"] = shoset.HandleCommand
	member.chaussette.Handle["evt"] = shoset.HandleEvent
	member.chaussette.Handle["config"] = shoset.HandleConnectorConfig

	coreLog.OpenLogFile(logPath)

	return member
}

// GetChaussette : Connector chaussette getter.
func (m *ConnectorMember) GetChaussette() *net.Shoset {
	return m.chaussette
}

// GetConnectorGrpc : Connector grpc getter.
func (m *ConnectorMember) GetConnectorGrpc() grpc.ConnectorGrpc {
	return m.connectorGrpc
}

// GetTimeoutMax : Connector timeoutmax getter.
func (m *ConnectorMember) GetTimeoutMax() int64 {
	return m.timeoutMax
}

// Bind : Connector bind function.
func (m *ConnectorMember) Bind(addr string) error {
	ipAddr, err := net.GetIP(addr)
	if err == nil {
		err = m.chaussette.Bind(ipAddr)
	}

	return err
}

// GrpcBind : Connector grpcbind function.
func (m *ConnectorMember) GrpcBind(addr string) (err error) {
	m.connectorGrpc, err = grpc.NewConnectorGrpc(addr, m.timeoutMax, m.chaussette)
	go m.connectorGrpc.StartGrpcServer()

	return err
}

// Join : Connector join function.
func (m *ConnectorMember) Join(addr string) (*net.ShosetConn, error) {
	return m.chaussette.Join(addr)
}

// Link : Connector link function.
func (m *ConnectorMember) Link(addr string) (*net.ShosetConn, error) {
	return m.chaussette.Link(addr)
}

// GetConfiguration : GetConfiguration
func (m *ConnectorMember) GetConfiguration(nshoset *net.Shoset, timeoutMax int64) (err error) {
	return shoset.SendConnectorConfig(nshoset, timeoutMax)
}

// StartWorkers : start workers
func (m *ConnectorMember) StartWorkers(logicalName, grpcBindAddress, targetAdd, workersPath string) (err error) {
	files, err := ioutil.ReadDir(workersPath)

	if err != nil {
		panic(err)
	}
	args := []string{logicalName, strconv.FormatInt(m.GetTimeoutMax(), 10), grpcBindAddress}

	for _, fileInfo := range files {
		if !fileInfo.IsDir() {
			if utils.IsExecAll(fileInfo.Mode().Perm()) {
				cmd := exec.Command("./"+fileInfo.Name(), args...)
				cmd.Dir = workersPath
				cmd.Stdout = os.Stdout
				err := cmd.Start()
				fmt.Println(err)
			}
		}
	}

	return nil
}

// ConfigurationValidation : validation configuration
func (m *ConnectorMember) ConfigurationValidation(tenant, connectorType string) (result bool) {
	commands := m.chaussette.Context["connectorCommands"].([]string)
	config := m.chaussette.Context["connectorConfig"].(models.ConnectorConfig)

	var configCommands []string
	for _, command := range config.ConnectorTypeCommands {
		configCommands = append(configCommands, command.Name)
	}

	return reflect.DeepEqual(commands, configCommands)
}

// getBrothers : Connector list brothers function.
func getBrothers(address string, member *ConnectorMember) []string {
	bros := []string{address}

	member.chaussette.ConnsJoin.Iterate(
		func(key string, val *net.ShosetConn) {
			bros = append(bros, key)
		})

	return bros
}

// ConnectorMemberInit : Connector init function.
func ConnectorMemberInit(logicalName, tenant, bindAddress, grpcBindAddress, linkAddress, connectorType, targetAdd, workerPath, logPath string, timeoutMax int64) *ConnectorMember {
	member := NewConnectorMember(logicalName, tenant, connectorType, logPath)
	member.timeoutMax = timeoutMax

	err := member.Bind(bindAddress)
	if err == nil {
		err = member.GrpcBind(grpcBindAddress)
		if err == nil {
			_, err = member.Link(linkAddress)
			time.Sleep(time.Second * time.Duration(5))
			if err == nil {
				err = member.GetConfiguration(member.GetChaussette(), timeoutMax)
				if err == nil {
					err = member.StartWorkers(logicalName, grpcBindAddress, targetAdd, workerPath)
					if err == nil {
						time.Sleep(time.Second * time.Duration(5))
						result := member.ConfigurationValidation(tenant, connectorType)
						if result {
							log.Printf("New Connector member %s for tenant %s bind on %s GrpcBind on %s link on %s \n", logicalName, tenant, bindAddress, grpcBindAddress, linkAddress)

							//time.Sleep(time.Second * time.Duration(5))
							fmt.Printf("%s.JoinBrothers Init(%#v)\n", bindAddress, getBrothers(bindAddress, member))
						} else {
							log.Printf("Configuration validation failed")
						}
					} else {
						log.Printf("Can't start workers in %s", workerPath)
					}
				} else {
					log.Printf("Can't get configuration in %s", workerPath)
				}
			} else {
				log.Printf("Can't link shoset on %s", linkAddress)
			}
		} else {
			log.Printf("Can't Grpc bind shoset on %s", grpcBindAddress)
		}
	} else {
		log.Printf("Can't bind shoset on %s", bindAddress)
	}

	return member
}

/* func ConnectorMemberJoin(logicalName, tenant, bindAddress, grpcBindAddress, linkAddress, joinAddress string, timeoutMax int64) (connectorMember *ConnectorMember) {

	member := NewConnectorMember(logicalName, tenant)
	member.timeoutMax = timeoutMax

	member.Bind(bindAddress)
	member.GrpcBind(grpcBindAddress)
	member.Link(linkAddress)
	member.Join(joinAddress)

	time.Sleep(time.Second * time.Duration(5))
	fmt.Printf("%s.JoinBrothers Join(%#v)\n", bindAddress, getBrothers(bindAddress, member))

	return member
}
*/
