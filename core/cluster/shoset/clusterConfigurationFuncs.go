//Package shoset :
package shoset

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	cmodels "github.com/ditrit/gandalf/core/configuration/models"
	"github.com/ditrit/gandalf/core/models"
	cmsg "github.com/ditrit/gandalf/core/msg"
	"github.com/jinzhu/gorm"

	"github.com/ditrit/gandalf/core/cluster/database"
	cutils "github.com/ditrit/gandalf/core/cluster/utils"

	net "github.com/ditrit/shoset"
	"github.com/ditrit/shoset/msg"
)

var configurationSendIndex = 0

func GetConfiguration(c *net.ShosetConn) (msg.Message, error) {
	var configuration cmsg.Configuration
	err := c.ReadMessage(&configuration)
	return configuration, err
}

// WaitConfig :
func WaitConfiguration(c *net.Shoset, replies *msg.Iterator, args map[string]string, timeout int) *msg.Message {
	commandName, ok := args["name"]
	if !ok {
		return nil
	}
	term := make(chan *msg.Message, 1)
	cont := true
	go func() {
		for cont {
			message := replies.Get().GetMessage()
			if message != nil {
				config := message.(cmsg.Configuration)
				if config.GetCommand() == commandName {
					term <- &message
				}
			} else {
				time.Sleep(time.Duration(10) * time.Millisecond)
			}
		}
	}()
	select {
	case res := <-term:
		cont = false
		return res
	case <-time.After(time.Duration(timeout) * time.Second):
		return nil
	}
}

// HandleConnectorConfig : Cluster handle connector config function.
func HandleConfiguration(c *net.ShosetConn, message msg.Message) (err error) {
	configuration := message.(cmsg.Configuration)
	ch := c.GetCh()

	err = nil

	log.Println("Handle configuration")
	log.Println(configuration)
	fmt.Println("HANDLE CONFIGURATION")
	fmt.Println(configuration)

	if configuration.GetCommand() == "PIVOT_CONFIGURATION" {

		componentType := configuration.Context["componentType"].(string)
		databaseConnection := ch.Context["databaseConnection"].(*database.DatabaseConnection)
		//mapDatabaseClient := ch.Context["tenantDatabases"].(map[string]*gorm.DB)
		//databaseBindAddr := ch.Context["databaseBindAddr"].(string)
		//configurationCluster := ch.Context["configuration"].(*cmodels.ConfigurationCluster)

		if databaseConnection != nil {
			//databaseClient := databaseConnection.GetDatabaseClientByTenant(configuration.GetTenant())
			var databaseClient *gorm.DB
			if componentType == "cluster" {
				databaseClient = databaseConnection.GetGandalfDatabaseClient()
			} else {
				databaseClient = databaseConnection.GetDatabaseClientByTenant(configuration.GetTenant())
			}
			if databaseClient != nil {
				//componentType := configuration.Context["componentType"].(string)
				var version models.Version
				jsonVersion := configuration.Context["version"].([]byte)
				json.Unmarshal(jsonVersion, &version)

				pivot := cutils.GetPivots(databaseClient, componentType, version)
				jsonData, err := json.Marshal(pivot)

				if err == nil {
					switch componentType {
					case "cluster":
						cmdReply := cmsg.NewConfiguration("", "PIVOT_CONFIGURATION_REPLY", string(jsonData))
						cmdReply.Tenant = configuration.GetTenant()
						shoset := ch.ConnsJoin.Get(configuration.Context["bindAddress"].(string))
						fmt.Println("shoset")
						fmt.Println(shoset)
						shoset.SendMessage(cmdReply)
						break
					case "aggregator":
						cmdReply := cmsg.NewConfiguration("", "PIVOT_CONFIGURATION_REPLY", string(jsonData))
						cmdReply.Tenant = configuration.GetTenant()
						shoset := ch.ConnsByAddr.Get(c.GetBindAddr())

						shoset.SendMessage(cmdReply)
						break
					case "connector":
						cmdReply := cmsg.NewConfiguration(configuration.GetTarget(), "PIVOT_CONFIGURATION_REPLY", string(jsonData))
						cmdReply.Tenant = configuration.GetTenant()
						cmdReply.GetContext()["componentType"] = "connector"

						shoset := ch.ConnsByAddr.Get(c.GetBindAddr())

						shoset.SendMessage(cmdReply)
						break
					case "admin":
						cmdReply := cmsg.NewConfiguration(configuration.GetTarget(), "PIVOT_CONFIGURATION_REPLY", string(jsonData))
						cmdReply.Tenant = configuration.GetTenant()
						cmdReply.GetContext()["componentType"] = "admin"

						shoset := ch.ConnsByAddr.Get(c.GetBindAddr())

						shoset.SendMessage(cmdReply)
					default:
						cmdReply := cmsg.NewConfiguration(configuration.GetTarget(), "PIVOT_CONFIGURATION_REPLY", string(jsonData))
						cmdReply.Tenant = configuration.GetTenant()
						cmdReply.GetContext()["componentType"] = "worker"

						shoset := ch.ConnsByAddr.Get(c.GetBindAddr())

						shoset.SendMessage(cmdReply)
						break
					}

				} else {
					log.Println("Can't unmarshall configuration")
					err = errors.New("Can't unmarshall configuration")
				}
			} else {
				log.Println("Can't get database client by tenant")
				err = errors.New("Can't get database client by tenant")
			}
		} else {
			log.Println("Can't get database clients")
			err = errors.New("Can't get database clients")
		}

	} else if configuration.GetCommand() == "CONNECTOR_PRODUCT_CONFIGURATION" {

		componentType := configuration.Context["product"].(string)
		databaseConnection := ch.Context["databaseConnection"].(*database.DatabaseConnection)
		//mapDatabaseClient := ch.Context["tenantDatabases"].(map[string]*gorm.DB)
		//databaseBindAddr := ch.Context["databaseBindAddr"].(string)
		//configurationCluster := ch.Context["configuration"].(*cmodels.ConfigurationCluster)

		if databaseConnection != nil {
			//databaseClient := databaseConnection.GetDatabaseClientByTenant(configuration.GetTenant())
			var databaseClient *gorm.DB
			if componentType == "cluster" {
				databaseClient = databaseConnection.GetGandalfDatabaseClient()
			} else {
				databaseClient = databaseConnection.GetDatabaseClientByTenant(configuration.GetTenant())
			}
			if databaseClient != nil {
				product := configuration.Context["product"].(string)
				var version models.Version
				jsonVersion := configuration.Context["version"].([]byte)
				json.Unmarshal(jsonVersion, &version)

				productConnectors := cutils.GetProductConnectors(databaseClient, product, version)
				jsonData, err := json.Marshal(productConnectors)

				if err == nil {
					cmdReply := cmsg.NewConfiguration(configuration.GetTarget(), "CONNECTOR_PRODUCT_CONFIGURATION_REPLY", string(jsonData))
					cmdReply.Tenant = configuration.GetTenant()
					shoset := ch.ConnsByAddr.Get(c.GetBindAddr())
					fmt.Println("shoset")
					fmt.Println(shoset)
					shoset.SendMessage(cmdReply)
				} else {
					log.Println("Can't unmarshall configuration")
					err = errors.New("Can't unmarshall configuration")
				}
			} else {
				log.Println("Can't get database client by tenant")
				err = errors.New("Can't get database client by tenant")
			}
		} else {
			log.Println("Can't get database clients")
			err = errors.New("Can't get database clients")
		}

	} else if configuration.GetCommand() == "PIVOT_CONFIGURATION_REPLY" {
		fmt.Println("REPLY")
		var pivot *models.Pivot
		err = json.Unmarshal([]byte(configuration.GetPayload()), &pivot)
		if err == nil {
			ch.Context["pivot"] = pivot
		}
	} else if configuration.GetCommand() == "SAVE_PIVOT_CONFIGURATION" {
		fmt.Println("SAVE")
		databaseConnection := ch.Context["databaseConnection"].(*database.DatabaseConnection)
		//mapDatabaseClient := ch.Context["tenantDatabases"].(map[string]*gorm.DB)
		//databaseBindAddr := ch.Context["databaseBindAddr"].(string)
		//configurationCluster := ch.Context["configuration"].(*cmodels.ConfigurationCluster)

		if databaseConnection != nil {
			//databaseClient := databaseConnection.GetDatabaseClientByTenant(configuration.GetTenant())
			var databaseClient *gorm.DB
			databaseClient = databaseConnection.GetDatabaseClientByTenant(configuration.GetTenant())

			if databaseClient != nil {
				//connectorConfig := conf.GetContext()["connectorConfig"].(models.ConnectorConfig)
				var pivot models.Pivot
				err = json.Unmarshal([]byte(configuration.GetPayload()), &pivot)
				fmt.Println("PIVOT SAVE")
				fmt.Println(pivot)
				if err == nil {
					cutils.SavePivot(pivot, databaseClient)
				}
			} else {
				log.Println("Can't get database client by tenant")
				err = errors.New("Can't get database client by tenant")
			}
		} else {
			log.Println("Can't get database clients")
			err = errors.New("Can't get database clients")
		}

	} else if configuration.GetCommand() == "SAVE_PRODUCT_CONNECTOR_CONFIGURATION" {

		componentType := configuration.Context["componentType"].(string)
		databaseConnection := ch.Context["databaseConnection"].(*database.DatabaseConnection)
		//mapDatabaseClient := ch.Context["tenantDatabases"].(map[string]*gorm.DB)
		//databaseBindAddr := ch.Context["databaseBindAddr"].(string)
		//configurationCluster := ch.Context["configuration"].(*cmodels.ConfigurationCluster)

		if databaseConnection != nil {
			//databaseClient := databaseConnection.GetDatabaseClientByTenant(configuration.GetTenant())
			var databaseClient *gorm.DB
			if componentType == "cluster" {
				databaseClient = databaseConnection.GetGandalfDatabaseClient()
			} else {
				databaseClient = databaseConnection.GetDatabaseClientByTenant(configuration.GetTenant())
			}
			if databaseClient != nil {
				var productConnector *models.ProductConnector
				err = json.Unmarshal([]byte(configuration.GetPayload()), &productConnector)
				if err == nil {
					cutils.SaveProductConnector(productConnector, databaseClient)
				}
			} else {
				log.Println("Can't get database client by tenant")
				err = errors.New("Can't get database client by tenant")
			}
		} else {
			log.Println("Can't get database clients")
			err = errors.New("Can't get database clients")
		}
		//connectorConfig := conf.GetContext()["connectorConfig"].(models.ConnectorConfig)

	}

	/* 	gandalfdatabaseClient := cutils.GetGandalfDatabaseClient(databasePath)
	   	if gandalfdatabaseClient != nil {

	   	} else {
	   		log.Println("Can't get gandalf database client")
	   		err = errors.New("Can't get gandalf database client")
	   	} */

	return err
}

//SendPivotConfiguration :
func SendClusterPivotConfiguration(shoset *net.Shoset) (err error) {
	version := shoset.Context["version"].(models.Version)
	jsonVersion, err := json.Marshal(version)
	if err == nil {
		conf := cmsg.NewConfiguration("", "PIVOT_CONFIGURATION", "")
		configurationCluster := shoset.Context["configuration"].(*cmodels.ConfigurationCluster)

		conf.GetContext()["componentType"] = "cluster"
		conf.GetContext()["version"] = jsonVersion
		conf.GetContext()["bindAddress"] = configurationCluster.GetBindAddress()

		//conf.GetContext()["product"] = shoset.Context["product"]

		shosets := net.GetByType(shoset.ConnsJoin, "")

		if len(shosets) != 0 {
			if conf.GetTimeout() > configurationCluster.GetMaxTimeout() {
				conf.Timeout = configurationCluster.GetMaxTimeout()
			}

			notSend := true
			for start := time.Now(); time.Since(start) < time.Duration(conf.GetTimeout())*time.Millisecond; {
				index := getConfigurationSendIndex(shosets)
				shosets[index].SendMessage(conf)
				log.Printf("%s : send command %s to %s\n", shoset.GetBindAddr(), conf.GetCommand(), shosets[index])

				timeoutSend := time.Duration((int(conf.GetTimeout()) / len(shosets)))

				time.Sleep(timeoutSend * time.Millisecond)

				if shoset.Context["pivot"] != nil {
					notSend = false
					break
				}
			}

			if notSend {
				return nil
			}

		} else {
			log.Println("can't find aggregators to send")
			err = errors.New("can't find aggregators to send")
		}
	}

	return err
}

// getCommandSendIndex : Aggregator getSendIndex function.
func getConfigurationSendIndex(conns []*net.ShosetConn) int {
	aux := configurationSendIndex
	configurationSendIndex++

	if configurationSendIndex >= len(conns) {
		configurationSendIndex = 0
	}

	return aux
}
