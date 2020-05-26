//Package shoset :
package shoset

import (
	"encoding/json"
	"errors"
	"fmt"
	"gandalf-core/models"
	"log"
	"shoset/msg"
	"shoset/net"
)

var configSendIndex = 0

// HandleConnectorConfig : Connector handle connector config.
func HandleConnectorConfig(c *net.ShosetConn, message msg.Message) (err error) {
	conf := message.(msg.Config)
	ch := c.GetCh()
	err = nil

	log.Println("Handle connector config")
	log.Println(conf)

	if conf.GetCommand() == "CONF_REPLY" {
		var connectorConfig = ch.Context["connectorConfig"].(models.ConnectorConfig)
		err = json.Unmarshal([]byte(conf.GetPayload()), &connectorConfig)
	}

	return err
}

//TODO REVOIR CONDITION
//SendConnectorConfig : Connector send connector config function.
func SendConnectorConfig(shoset *net.Shoset, timeoutMax int64) (err error) {
	conf := msg.NewConfig("", "CONFIG", "")
	conf.Tenant = shoset.Context["tenant"].(string)

	shosets := net.GetByType(shoset.ConnsByAddr, "a")

	fmt.Println("SH")
	fmt.Println(shoset)
	fmt.Println(shoset.ConnsByAddr)
	fmt.Println(shosets)

	if len(shosets) != 0 {
		if conf.GetTimeout() > timeoutMax {
			conf.Timeout = timeoutMax
		}
		shosets[0].SendMessage(conf)
		log.Printf("%s : send command %s to %s\n", shoset.GetBindAddr(), conf.GetCommand(), shosets[0])

	} else {
		log.Println("can't find aggregators to send")
		err = errors.New("can't find aggregators to send")
	}

	return err
}

/* //SendConnectorConfig : Connector send connector config function.
func SendConnectorConfig(shoset *net.Shoset, timeoutMax int64) (err error) {
	conf := msg.NewConfig("", "CONFIG", "")
	conf.Tenant = shoset.Context["tenant"].(string)

	shosets := net.GetByType(shoset.ConnsByAddr, "a")

	fmt.Println("SH")
	fmt.Println(shoset)
	fmt.Println(shoset.ConnsByAddr)
	fmt.Println(shosets)

	if len(shosets) != 0 {
		if conf.GetTimeout() > timeoutMax {
			conf.Timeout = timeoutMax
		}

		notSend := true
		for notSend {
			index := getConfigSendIndex(shosets)
			shosets[index].SendMessage(conf)
			log.Printf("%s : send command %s to %s\n", shoset.GetBindAddr(), conf.GetCommand(), shosets[index])

			timeoutSend := time.Duration((int(conf.GetTimeout()) / len(shosets)))

			if shoset.Context["connectorConfig"].(models.ConnectorConfig) != nil {
				notSend = false
				break
			}
			time.Sleep(timeoutSend)
		}

		if notSend {
			return nil
		}

	} else {
		log.Println("can't find aggregators to send")
		err = errors.New("can't find aggregators to send")
	}

	return err
}

// getSendIndex : Cluster getSendIndex function.
func getConfigSendIndex(conns []*net.ShosetConn) int {
	aux := configSendIndex
	configSendIndex++

	if configSendIndex >= len(conns) {
		configSendIndex = 0
	}

	return aux
} */
