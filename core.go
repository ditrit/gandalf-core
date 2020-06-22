package main

import (
	"core/aggregator"
	"core/cluster"
	"core/configuration"
	"core/connector"
	"core/database"
	"fmt"
	"log"
	"shoset/net"
)

func main() {

	configuration.ConfigMain()

	gandalfLogicalName , err := configuration.GetStringConfig("logical_name")
	if err != nil {
		log.Fatalf("No logical name: %v",err)
	}
	gandalfLogPath, err := configuration.GetStringConfig("gandalf_log")
	if err != nil{
		log.Fatalf("No valid log path : %v",err)
	}
	gandalfBindAddress ,err := configuration.GetStringConfig("bind_address")
	if err != nil {
		log.Fatalf("No valid bind address : %v",err)
	}
	gandalfType, err := configuration.GetStringConfig("gandalf_type")
	if err == nil {

		switch gandalfType {
		case "cluster":
			gandalfDBPath, err := configuration.GetStringConfig("gandalf_db")
			if err != nil {
				log.Fatalf("No valid database path : %v",err)
			}
			gandalfJoin, err := configuration.GetStringConfig("cluster_join")
			if err == nil {
				if gandalfJoin == "" {
					done := make(chan bool)
					cluster.ClusterMemberInit(gandalfLogicalName, gandalfBindAddress, gandalfLogPath)
					add, _ := net.DeltaAddress(gandalfBindAddress, 1000)
					go database.DatabaseMemberInit(add, gandalfDBPath, 1)
					<- done
				} else {
					done := make(chan bool)
					member := cluster.ClusterMemberJoin(gandalfLogicalName, gandalfBindAddress, gandalfJoin, gandalfLogPath)
					add, _ := net.DeltaAddress(gandalfBindAddress, 1000)
					id := len(*member.Store)

					go database.DatabaseMemberInit(add, gandalfDBPath, id)

					_ = database.AddNodesToLeader(id, add, *member.Store)
					<- done
				}
			}
			break
		case "aggregator":
			gandalfTenant,err := configuration.GetStringConfig("tenant")
			if err !=nil {
				log.Fatalf("no valid tenant : %v",err)
			}
			gandalfClusterLink,err := configuration.GetStringConfig("clusters")
			if err != nil {
				log.Fatalf("no valid cluster address: %v",err)
			}
			done := make(chan bool)
			aggregator.AggregatorMemberInit(gandalfLogicalName, gandalfTenant, gandalfBindAddress, gandalfClusterLink, gandalfLogPath)
			<-done
			break
		case "connector":
			gandalfTenant,err := configuration.GetStringConfig("tenant")
			if err !=nil {
				log.Fatalf("no valid tenant : %v",err)
			}
			gandalfGRPCBindAddress, err := configuration.GetStringConfig("grpc_bind_address")
			if err !=nil {
				log.Fatalf("no valid tenant : %v",err)
			}
			gandalfAggregatorLink,err := configuration.GetStringConfig("aggregators")
			if err !=nil {
				log.Fatalf("no valid tenant : %v",err)
			}
			gandalfMaxTimeout,err := configuration.GetIntegerConfig("max_timeout")
			if err !=nil {
				log.Fatalf("no valid tenant : %v",err)
			}
			done := make(chan bool)
			connector.ConnectorMemberInit(gandalfLogicalName, gandalfTenant, gandalfBindAddress, gandalfGRPCBindAddress, gandalfAggregatorLink, gandalfLogPath, int64(gandalfMaxTimeout))
			<- done
			break

		default:
			break
		}
	}
	/*
	//CREATE CLUSTER
	fmt.Println("Running Gandalf with:")
	fmt.Println("  Mode : " + mode)
	fmt.Println("  Logical Name : " + LogicalName)
	fmt.Println("  Bind Address : " + BindAdd)
	fmt.Println("  Log Path : " + LogPath)
	fmt.Println("  Db Path : " + dbPath)
	fmt.Println("  Config : " + config)

	<-done

	//CREATE CLUSTER
	fmt.Println("Running Gandalf with:")
	fmt.Println("  Mode : " + mode)
	fmt.Println("  Logical Name : " + LogicalName)
	fmt.Println("  Bind Address : " + BindAdd)
	fmt.Println("  Join Address : " + JoinAdd)
	fmt.Println("  Log Path : " + LogPath)
	fmt.Println("  Db Path : " + dbPath)
	fmt.Println("  Config : " + config)

	<-done

	//CREATE AGGREGATOR
	fmt.Println("Running Gandalf with:")
	fmt.Println("  Logical Name : " + LogicalName)
	fmt.Println("  Tenant : " + Tenant)
	fmt.Println("  Bind Address : " + BindAdd)
	fmt.Println("  Link Address : " + LinkAdd)
	fmt.Println("  Log Path : " + LogPath)
	fmt.Println("  Config : " + config)

	aggregator.AggregatorMemberInit(LogicalName, Tenant, BindAdd, LinkAdd, LogPath)

	<-done
	//CREATE CONNECTOR
	fmt.Println("Running Gandalf with:")
	fmt.Println("  Logical Name : " + LogicalName)
	fmt.Println("  Tenant : " + Tenant)
	fmt.Println("  Bind Address : " + BindAdd)
	fmt.Println("  Grpc Bind Address : " + GrpcBindAdd)
	fmt.Println("  Link Address : " + LinkAdd)
	fmt.Println("  Log Path : " + LogPath)
	fmt.Printf("   Timeout Max : %d \n", TimeoutMax)
	fmt.Println("  Config : " + config)

	connector.ConnectorMemberInit(LogicalName, Tenant, BindAdd, GrpcBindAdd, LinkAdd, LogPath, TimeoutMax)

	<-done
 	*/
}
