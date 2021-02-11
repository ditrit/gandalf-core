/*
Copyright © 2020 DitRit community <contact@ditrit.io>
This file is part of Gandalf
*/

// Package cmd manages commands and configuration
package configuration

import (
	"fmt"

	"github.com/ditrit/gandalf/core/cli/client"
	cmodels "github.com/ditrit/gandalf/core/configuration/models"
	"github.com/ditrit/gandalf/core/models"

	"github.com/ditrit/gandalf/core/configuration/config"
	"github.com/spf13/viper"
)

// cliCmd represents the cli command
var cliCfg = config.NewConfigCmd("cli", "Launch gandalf in 'cli' mode.", `Gandalf is launched as CLI (Command Line Interface) to interact with a Gandalf system.`, nil)

var cliCreate = config.NewConfigCmd("create", "create user|tenant|role|domain", "create command allows the creation of Gandalf objects (users, tenants, roles and domains).", nil)
var cliList = config.NewConfigCmd("list", "list users|tenants|roles|domains", "list command allows to list Gandalf objects (users, tenants, roles and domains).", nil)
var cliUpdate = config.NewConfigCmd("update", "update user|tenant|role|domain", "update command allows update of Gandalf objects (users, tenants, roles and domains).", nil)
var cliDelete = config.NewConfigCmd("delete", "delete user|tenant|role|domain", "update command allows deleting of Gandalf objects (users, tenants, roles and domains).", nil)
var cliDeclare = config.NewConfigCmd("declare", "declare cluster|agregator|connector", "declare command allows to declare a new Gandalf component name or member (cluster, aggregator, connector).", nil)
var cliLogin = config.NewConfigCmd("login", "log in as a user into Gandalf", "login command allows user to authenticate using its credentials.", runLogin)

var cliCreateUser = config.NewConfigCmd("user", "create user <username> <email> <password>", "create user command allows the creation of a new user", runCreateUser)
var cliListUsers = config.NewConfigCmd("user", "list users", "list users command allows to list Gandalf users.", runListUsers)
var cliUpdateUser = config.NewConfigCmd("user", "update user <username> [options]", "update user command allows to update a Gandalf user.", runUpdateUser)
var cliDeleteUser = config.NewConfigCmd("user", "delete user <username>", "delete user command allows to delete a Gandalf user.", runDeleteUser)

var cliCreateTenant = config.NewConfigCmd("tenant", "create tenant <tenantname>", "create tenant command allows the creation of a new tenant", runCreateTenant)
var cliListTenants = config.NewConfigCmd("tenant", "list tenants <tenantname>", "list tenants command allows to list Gandalf tenants.", runListTenants)
var cliUpdateTenant = config.NewConfigCmd("tenant", "update tenant <tenantname> [options]", "update tenant command allows to update a Gandalf tenant.", runUpdateTenant)
var cliDeleteTenant = config.NewConfigCmd("tenant", "delete tenant <tenantname>", "delete tenant command allows to delete a Gandalf tenant.", runDeleteTenant)

var cliCreateRole = config.NewConfigCmd("role", "create role <rolename> ", "create role command allows the creation of a new role", runCreateRole)
var cliListRoles = config.NewConfigCmd("role", "list roles <rolename> ", "list roles command allows to list Gandalf roles.", runListRoles)
var cliUpdateRole = config.NewConfigCmd("role", "update role <rolename> [options]", "update role command allows to update a Gandalf role.", runUpdateRole)
var cliDeleteRole = config.NewConfigCmd("role", "delete role <rolename>", "delete role command allows to delete a Gandalf role.", runDeleteRole)

var cliCreateDomain = config.NewConfigCmd("domain", "create domain <domainname>", "create domain command allows the creation of a new domain (in the form <[name.]*name>)", runCreateDomain)
var cliListDomains = config.NewConfigCmd("domain", "list domains ", "list domains command allows to list Gandalf domains.", runListDomains)
var cliUpdateDomain = config.NewConfigCmd("domain", "update domain <domainname> [options]", "update domain command allows to update a Gandalf domain.", runUpdateDomain)
var cliDeleteDomain = config.NewConfigCmd("domain", "delete domain <domainname>", "delete domain command allows to delete a Gandalf domain.", runDeleteDomain)

var cliDeclareCluster = config.NewConfigCmd("custer", "declare cluster", "declare cluster command allows to declare a new cluster memeber", nil)
var cliDeclareAggregator = config.NewConfigCmd("aggregator", "declare aggregator name|member", "declare aggregator command allows to declare the name or a new member for an aggragator.", nil)
var cliDeclareConnector = config.NewConfigCmd("connector", "declare connector name|member", "declare connector command allows to declare the name or a new member for a connector.", nil)

var cliDeclareClusterMember = config.NewConfigCmd("member", "declare cluster member", "declare cluster command allows to declare a new cluster member", runDeclareClusterMember)

var cliDeclareAggregatorName = config.NewConfigCmd("name", "declare aggregator name <name>", "declare aggregator name command allows to declare the name of a new aggregator.", runDeclareAggregatorName)
var cliDeclareAggregatorMember = config.NewConfigCmd("member", "declare aggregator member <name>", "declare aggregator member command allows to declare a new member for an existing aggregator.", runDeclareAggregatorMember)

var cliDeclareConnectorName = config.NewConfigCmd("name", "declare  name <name>", "declare  name command allows to declare the name of a new connector.", runDeclareConnectorName)
var cliDeclareConnectorMember = config.NewConfigCmd("member", "declare  member <name>", "declare  member command allows to declare a new member for an existing connector.", runDeclareAggregatorName)

func init() {
	rootCfg.AddConfig(cliCfg)
	cliCfg.AddConfig(cliCreate)
	cliCfg.AddConfig(cliList)
	cliCfg.AddConfig(cliUpdate)
	cliCfg.AddConfig(cliDelete)
	cliCfg.AddConfig(cliDeclare)
	cliCfg.AddConfig(cliLogin)

	cliCreate.AddConfig(cliCreateUser)
	cliList.AddConfig(cliListUsers)
	cliUpdate.AddConfig(cliUpdateUser)
	cliDelete.AddConfig(cliDeleteUser)

	cliCreate.AddConfig(cliCreateTenant)
	cliList.AddConfig(cliListTenants)
	cliUpdate.AddConfig(cliUpdateTenant)
	cliDelete.AddConfig(cliDeleteTenant)

	cliCreate.AddConfig(cliCreateRole)
	cliList.AddConfig(cliListRoles)
	cliUpdate.AddConfig(cliUpdateRole)
	cliDelete.AddConfig(cliDeleteRole)

	cliCreate.AddConfig(cliCreateDomain)
	cliList.AddConfig(cliListDomains)
	cliUpdate.AddConfig(cliUpdateDomain)
	cliDelete.AddConfig(cliDeleteDomain)

	cliDeclare.AddConfig(cliDeclareCluster)
	cliDeclare.AddConfig(cliDeclareAggregator)
	cliDeclare.AddConfig(cliDeclareConnector)

	cliDeclareCluster.AddConfig(cliDeclareClusterMember)

	cliDeclareAggregator.AddConfig(cliDeclareAggregatorName)
	cliDeclareAggregator.AddConfig(cliDeclareAggregatorMember)

	cliDeclareConnector.AddConfig(cliDeclareConnectorName)
	cliDeclareConnector.AddConfig(cliDeclareConnectorMember)

	cliCfg.Key("endpoint", config.IsStr, "e", "Gandalf endpoint")
	cliCfg.SetRequired("endpoint")
	cliCfg.Key("token", config.IsStr, "t", "Gandalf auth token")
	cliCfg.SetRequired("token")

	cliLogin.SetNbArgs(2)

	cliCreateUser.SetNbArgs(3)
	cliListUsers.SetNbArgs(0)
	cliUpdateUser.SetNbArgs(1)
	cliDeleteUser.SetNbArgs(1)
	cliUpdateUser.Key("username", config.IsStr, "u", "name of the user")
	cliUpdateUser.Key("email", config.IsStr, "m", "mail of the user")
	cliUpdateUser.Key("password", config.IsStr, "p", "password of the user")

	cliCreateTenant.SetNbArgs(1)
	cliListTenants.SetNbArgs(0)
	cliUpdateTenant.SetNbArgs(1)
	cliDeleteTenant.SetNbArgs(1)
	cliUpdateTenant.Key("tenantname", config.IsStr, "t", "name of the Tenant")

	cliCreateRole.SetNbArgs(1)
	cliListRoles.SetNbArgs(0)
	cliUpdateRole.SetNbArgs(1)
	cliDeleteRole.SetNbArgs(1)
	cliUpdateRole.Key("rolename", config.IsStr, "r", "name of the Role")

	cliCreateDomain.SetNbArgs(1)
	cliListDomains.SetNbArgs(0)
	cliUpdateDomain.SetNbArgs(1)
	cliDeleteDomain.SetNbArgs(1)
	cliUpdateDomain.Key("domainname", config.IsStr, "d", "name of the Domain")

	cliDeclareClusterMember.SetNbArgs(0)
	cliDeclareAggregatorName.SetNbArgs(1)
	cliDeclareAggregatorMember.SetNbArgs(1)
	cliDeclareConnectorName.SetNbArgs(1)
	cliDeclareConnectorMember.SetNbArgs(1)
}

func runLogin(cfg *config.ConfigCmd, args []string) {
	name := args[0]
	password := args[1]

	fmt.Printf("gandalf cli login called with username=%s and password=%s\n", name, password)
	configurationCli := cmodels.NewConfigurationCli()
	cliClient := client.NewClient(configurationCli.GetEndpoint())

	user := models.NewUser(name, "", password)
	token, err := cliClient.GandalfAuthenticationService.Login(user)
	if err == nil {
		fmt.Println(token)
	} else {
		fmt.Println(err)
	}
}

func runCreateUser(cfg *config.ConfigCmd, args []string) {
	name := args[0]
	email := args[1]
	password := args[2]

	fmt.Printf("gandalf cli create user called with username=%s, email=%s, password=%s\n", name, email, password)
	configurationCli := cmodels.NewConfigurationCli()
	cliClient := client.NewClient(configurationCli.GetEndpoint())

	user := models.NewUser(name, email, password)
	err := cliClient.GandalfUserService.Create(configurationCli.GetToken(), user)
	if err != nil {
		fmt.Println(err)
	}
}

func runListUsers(cfg *config.ConfigCmd, args []string) {
	fmt.Printf("gandalf cli list users\n")
	configurationCli := cmodels.NewConfigurationCli()
	cliClient := client.NewClient(configurationCli.GetEndpoint())

	users, err := cliClient.GandalfUserService.List(configurationCli.GetToken())
	if err == nil {
		for user := range users {
			fmt.Println(user)
		}
	} else {
		fmt.Println(err)
	}
}

func runUpdateUser(cfg *config.ConfigCmd, args []string) {
	name := args[0]
	newName := viper.GetString("username")
	email := viper.GetViper().GetString("email")
	password := viper.GetViper().GetString("password")
	fmt.Printf("gandalf cli update user called with username=%s, newname=%s, email=%s, password=%s\n", name, newName, email, password)
	configurationCli := cmodels.NewConfigurationCli()
	cliClient := client.NewClient(configurationCli.GetEndpoint())

	oldUser, err := cliClient.GandalfUserService.ReadByName(configurationCli.GetToken(), name)
	if err == nil {
		user := models.NewUser(newName, email, password)
		cliClient.GandalfUserService.Update(configurationCli.GetToken(), int(oldUser.ID), user)
	}
}

func runDeleteUser(cfg *config.ConfigCmd, args []string) {
	name := args[0]
	fmt.Printf("gandalf cli delete user called with username=%s\n", name)
	configurationCli := cmodels.NewConfigurationCli()
	cliClient := client.NewClient(configurationCli.GetEndpoint())

	oldUser, err := cliClient.GandalfUserService.ReadByName(configurationCli.GetToken(), name)
	if err == nil {
		cliClient.GandalfUserService.Delete(configurationCli.GetToken(), int(oldUser.ID))
	}
}

func runCreateTenant(cfg *config.ConfigCmd, args []string) {
	name := args[0]
	fmt.Printf("gandalf cli create tenant called with tenant=%s\n", name)
	configurationCli := cmodels.NewConfigurationCli()
	cliClient := client.NewClient(configurationCli.GetEndpoint())

	tenant := models.Tenant{Name: name}
	login, password, err := cliClient.GandalfTenantService.Create(configurationCli.GetToken(), tenant)
	if err == nil {
		fmt.Println("login : " + login)
		fmt.Println("password : " + password)
	} else {
		fmt.Println(err)
	}
}

func runListTenants(cfg *config.ConfigCmd, args []string) {
	fmt.Printf("gandalf cli list tenants\n")
	configurationCli := cmodels.NewConfigurationCli()
	cliClient := client.NewClient(configurationCli.GetEndpoint())

	tenants, err := cliClient.GandalfTenantService.List(configurationCli.GetToken())
	if err == nil {
		for tenant := range tenants {
			fmt.Println(tenant)
		}
	} else {
		fmt.Println(err)
	}
}

func runUpdateTenant(cfg *config.ConfigCmd, args []string) {
	name := args[0]
	newName := viper.GetString("name")
	fmt.Printf("gandalf cli update tenant called with tenant=%s, newName=%s\n", name, newName)
}

func runDeleteTenant(cfg *config.ConfigCmd, args []string) {
	name := args[0]
	fmt.Printf("gandalf cli delete tenant called with tenant=%s\n", name)
}

func runCreateRole(cfg *config.ConfigCmd, args []string) {
	name := args[0]
	fmt.Printf("gandalf cli create role called with role=%s\n", name)
}

func runListRoles(cfg *config.ConfigCmd, args []string) {
	fmt.Printf("gandalf cli list roles\n")
}

func runUpdateRole(cfg *config.ConfigCmd, args []string) {
	name := args[0]
	newName := viper.GetString("name")
	fmt.Printf("gandalf cli update role called with role=%s, newName=%s\n", name, newName)
}

func runDeleteRole(cfg *config.ConfigCmd, args []string) {
	name := args[0]
	fmt.Printf("gandalf cli delete role called with role=%s\n", name)
}

func runCreateDomain(cfg *config.ConfigCmd, args []string) {
	name := args[0]
	parent := viper.GetString("parent")
	fmt.Printf("gandalf cli create domain called with domain=%s parent=%s\n", name, parent)
}

func runListDomains(cfg *config.ConfigCmd, args []string) {
	fmt.Printf("gandalf cli list domains\n")
}

func runUpdateDomain(cfg *config.ConfigCmd, args []string) {
	name := args[0]
	newName := viper.GetString("name")
	parent := viper.GetString("parent")
	fmt.Printf("gandalf cli update domain called with domain=%s, newName=%s, parent=%s\n", name, newName, parent)
}

func runDeleteDomain(cfg *config.ConfigCmd, args []string) {
	name := args[0]
	fmt.Printf("gandalf cli delete domain called with domain=%s\n", name)
}

func runDeclareClusterMember(cfg *config.ConfigCmd, args []string) {
	fmt.Printf("gandalf declare cluster member\n")
}

func runDeclareAggregatorName(cfg *config.ConfigCmd, args []string) {
	name := args[0]
	fmt.Printf("gandalf declare aggregator name with name=%s\n", name)
}

func runDeclareAggregatorMember(cfg *config.ConfigCmd, args []string) {
	name := args[0]
	fmt.Printf("gandalf declare aggregator member with name=%s\n", name)
}

func runDeclareConnectorName(cfg *config.ConfigCmd, args []string) {
	name := args[0]
	fmt.Printf("gandalf declare connector name with name=%s\n", name)
}

func runDeclareConnectorMember(cfg *config.ConfigCmd, args []string) {
	name := args[0]
	fmt.Printf("gandalf declare connector member with name=%s\n", name)
}

/*
func init() {
	rootCfg.AddConfig(cliCfg)

	cliCfg.Key("api_port", config.IsInt, "", "Address to bind (default is *:9199)")
	cliCfg.SetDefault("api_port", 9199+config.GetOffset())

	cliCfg.Key("database_mode", config.IsStr, "", "database mode (gandalf|tenant)")
	cliCfg.SetDefault("database_mode", "gandalf")
	cliCfg.SetCheck("database_mode", func(val interface{}) bool {
		strVal := strings.ToLower(strings.TrimSpace(val.(string)))
		return map[string]bool{"gandalf": true, "tenant": true}[strVal]
	})

	cliCfg.Key("tenant", config.IsStr, "", "database mode (gandalf|tenant)")
	cliCfg.SetConstraint("a tenant should be provided if database_mode == tenant",
		func() bool {
			return viper.IsSet("database_mode") && viper.GetString("database_mode") == "tenant" && viper.IsSet("tenant")
		})

	cliCfg.Key("model", config.IsStr, "", "models  gandalf(authentication|cluster|tenant|role|user) || tenant(authentication|aggregator|connector|role|user)")
	cliCfg.SetCheck("model", func(val interface{}) bool {
		strVal := strings.ToLower(strings.TrimSpace(val.(string)))
		return map[string]bool{"authentication": true, "cluster": true, "tenant": true, "role": true, "user": true, "aggregator": true, "connector": true}[strVal]
	})
*/

//TODO REVOIR
/* 	cliCfg.SetConstraint("a id should be provided if command == (delete|update|read)",
func() bool {
	return viper.IsSet("database_mode") || viper.GetString("database_mode") == "gandalf" || viper.GetString("command") == "update" || viper.GetString("command") == "read" || viper.IsSet("id")
}) */

/*
	cliCfg.Key("command", config.IsStr, "", "command  (list|read|create|update|delete|upload)")
	cliCfg.SetCheck("command", func(val interface{}) bool {
		strVal := strings.ToLower(strings.TrimSpace(val.(string)))
		return map[string]bool{"list": true, "read": true, "create": true, "update": true, "delete": true, "upload": true}[strVal]
	})

	cliCfg.Key("token", config.IsStr, "", "")
	cliCfg.SetConstraint("a token should be provided if model != authenticaion",
		func() bool {
			return viper.IsSet("model") && viper.GetString("model") != "authentication" && viper.IsSet("token")
		})

	cliCfg.Key("id", config.IsStr, "", "id")
	cliCfg.SetConstraint("a id should be provided if command == (delete|update|read)",
		func() bool {
			return viper.IsSet("command") && (viper.GetString("command") == "delete" || viper.GetString("command") == "update" || viper.GetString("command") == "read") && viper.IsSet("id")
		})

	cliCfg.Key("value", config.IsStr, "", "json")
	cliCfg.SetConstraint("a value should be provided if command == (create|update)",
		func() bool {
			return viper.IsSet("command") && (viper.GetString("command") == "create" || viper.GetString("command") == "update") && viper.IsSet("value")
		})
}
*/
