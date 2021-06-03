/*
Copyright © 2020 DitRit community <contact@ditrit.io>
This file is part of Gandalf
*/

// Package configuration manages commands and configuration
package configuration

import (
	"fmt"

	"github.com/ditrit/gandalf/verdeter"

	"github.com/ditrit/gandalf/core/cli"
	cmodels "github.com/ditrit/gandalf/core/configuration/models"
	"github.com/ditrit/gandalf/core/models"

	"github.com/spf13/viper"
)

// cliCmd represents the cli command
var cliCfg = verdeter.NewConfigCmd("cli", "Launch gandalf in 'cli' mode.", `Gandalf is launched as CLI (Command Line Interface) to interact with a Gandalf system.`, nil)

var cliCreate = verdeter.NewConfigCmd("create", "create user|tenant|role|domain|resource|eventTypeToPoll|resourceType|eventType", "create command allows the creation of Gandalf objects (users, tenants, roles and domains).", nil)
var cliList = verdeter.NewConfigCmd("list", "list users|tenants|roles|domain|resource|eventTypeToPoll|resourceType|eventType", "list command allows to list Gandalf objects (users, tenants, roles and domains).", nil)
var cliUpdate = verdeter.NewConfigCmd("update", "update user|tenant|role|domain|resource|eventTypeToPoll|resourceType|eventType", "update command allows update of Gandalf objects (users, tenants, roles and domains).", nil)
var cliDelete = verdeter.NewConfigCmd("delete", "delete user|tenant|role|domain|resource|eventTypeToPoll|resourceType|eventType", "update command allows deleting of Gandalf objects (users, tenants, roles and domains).", nil)
var cliLogin = verdeter.NewConfigCmd("login", "log in as a user into Gandalf", "login command allows user to authenticate using its credentials.", runLogin)

var cliCreateUser = verdeter.NewConfigCmd("user", "create user <username> <email> <password>", "create user command allows the creation of a new user", runCreateUser)
var cliListUsers = verdeter.NewConfigCmd("user", "list users", "list users command allows to list Gandalf users.", runListUsers)
var cliUpdateUser = verdeter.NewConfigCmd("user", "update user <username> [options]", "update user command allows to update a Gandalf user.", runUpdateUser)
var cliDeleteUser = verdeter.NewConfigCmd("user", "delete user <username>", "delete user command allows to delete a Gandalf user.", runDeleteUser)

var cliCreateTenant = verdeter.NewConfigCmd("tenant", "create tenant <tenantname>", "create tenant command allows the creation of a new tenant", runCreateTenant)
var cliListTenants = verdeter.NewConfigCmd("tenant", "list tenants <tenantname>", "list tenants command allows to list Gandalf tenants.", runListTenants)
var cliUpdateTenant = verdeter.NewConfigCmd("tenant", "update tenant <tenantname> [options]", "update tenant command allows to update a Gandalf tenant.", runUpdateTenant)
var cliDeleteTenant = verdeter.NewConfigCmd("tenant", "delete tenant <tenantname>", "delete tenant command allows to delete a Gandalf tenant.", runDeleteTenant)

var cliCreateRole = verdeter.NewConfigCmd("role", "create role <rolename> ", "create role command allows the creation of a new role", runCreateRole)
var cliListRoles = verdeter.NewConfigCmd("role", "list roles <rolename> ", "list roles command allows to list Gandalf roles.", runListRoles)
var cliUpdateRole = verdeter.NewConfigCmd("role", "update role <rolename> [options]", "update role command allows to update a Gandalf role.", runUpdateRole)
var cliDeleteRole = verdeter.NewConfigCmd("role", "delete role <rolename>", "delete role command allows to delete a Gandalf role.", runDeleteRole)

var cliCreateDomain = verdeter.NewConfigCmd("domain", "create domain <domainname>", "create domain command allows the creation of a new domain (in the form <[name.]*name>)", runCreateDomain)
var cliListDomains = verdeter.NewConfigCmd("domain", "list domains ", "list domains command allows to list Gandalf domains.", runListDomains)
var cliUpdateDomain = verdeter.NewConfigCmd("domain", "update domain <domainname> [options]", "update domain command allows to update a Gandalf domain.", runUpdateDomain)
var cliDeleteDomain = verdeter.NewConfigCmd("domain", "delete domain <domainname>", "delete domain command allows to delete a Gandalf domain.", runDeleteDomain)

var cliCreateResource = verdeter.NewConfigCmd("resource", "create resource <resourcename>", "create resource command allows the creation of a new resource (in the form <[name.]*name>)", runCreateResource)
var cliListResources = verdeter.NewConfigCmd("resource", "list resources ", "list resource command allows to list Gandalf resources.", runListResources)
var cliUpdateResource = verdeter.NewConfigCmd("resource", "update resource <resourcename> [options]", "update resource command allows to update a Gandalf resource.", runUpdateResource)
var cliDeleteResource = verdeter.NewConfigCmd("resource", "delete resource <resourcename>", "delete resource command allows to delete a Gandalf resource.", runDeleteResource)

var cliCreateEventTypeToPoll = verdeter.NewConfigCmd("eventtypetopoll", "create eventtypetopoll <eventtypetopollname>", "create eventtypetopoll command allows the creation of a new resource (in the form <[name.]*name>)", runCreateEventTypeToPoll)
var cliListEventTypeToPolls = verdeter.NewConfigCmd("eventtypetopoll", "list eventtypetopolls ", "list eventtypetopolls command allows to list Gandalf eventtypetopolls.", runListEventTypeToPolls)
var cliUpdateEventTypeToPoll = verdeter.NewConfigCmd("eventtypetopoll", "update eventtypetopoll <eventtypetopollname> [options]", "update resource command allows to update a Gandalf eventtypetopoll.", runUpdateEventTypeToPoll)
var cliDeleteEventTypeToPoll = verdeter.NewConfigCmd("eventtypetopoll", "delete eventtypetopoll <eventtypetopollname>", "delete eventtypetopoll command allows to delete a Gandalf eventtypetopoll.", runDeleteEventTypeToPoll)

//
var cliCreateResourceType = verdeter.NewConfigCmd("resourcetype", "create resourcetype <resourcetypename>", "create resourcetype command allows the creation of a new resource (in the form <[name.]*name>)", runCreateResourceType)
var cliListResourceTypes = verdeter.NewConfigCmd("resourcetype", "list resourcetype ", "list resourcetype command allows to list Gandalf resourcetypes.", runListResourceTypes)
var cliUpdateResourceType = verdeter.NewConfigCmd("resourcetype", "update resourcetype <resourcetypename> [options]", "update resource command allows to update a Gandalf resourcetype.", runUpdateResourceType)
var cliDeleteResourceType = verdeter.NewConfigCmd("resourcetype", "delete resourcetype <resourcetypename>", "delete resourcetype command allows to delete a Gandalf resourcetype.", runDeleteResourceType)

var cliCreateEventType = verdeter.NewConfigCmd("eventtype", "create eventtype <eventtypename>", "create eventtype command allows the creation of a new resource (in the form <[name.]*name>)", runCreateEventType)
var cliListEventTypes = verdeter.NewConfigCmd("eventtype", "list eventtypes ", "list eventtypes command allows to list Gandalf eventtype.", runListEventTypes)
var cliUpdateEventType = verdeter.NewConfigCmd("eventtype", "update eventtype <eventtypename> [options]", "update resource command allows to update a Gandalf eventtype.", runUpdateEventType)
var cliDeleteEventType = verdeter.NewConfigCmd("eventtype", "delete eventtype <eventtypename>", "delete eventtype command allows to delete a Gandalf eventtype.", runDeleteEventType)

//
var cliCreateSecret = verdeter.NewConfigCmd("secret", "create secret", "declare  name command allows to declare the name of a new connector.", runCreateSecret)
var cliListSecret = verdeter.NewConfigCmd("secret", "list secret", "declare  member command allows to declare a new member for an existing connector.", runListSecret)

func init() {

	rootCfg.AddConfig(cliCfg)

	cliCfg.GKey("endpoint", verdeter.IsStr, "e", "Gandalf endpoint")
	cliCfg.SetRequired("endpoint")
	cliCfg.GKey("token", verdeter.IsStr, "t", "Gandalf auth token")
	//cliCfg.SetRequired("token")

	cliCfg.AddConfig(cliCreate)
	cliCfg.AddConfig(cliList)
	cliCfg.AddConfig(cliUpdate)
	cliCfg.AddConfig(cliDelete)
	cliCfg.AddConfig(cliLogin)

	cliCreate.AddConfig(cliCreateSecret)
	cliList.AddConfig(cliListSecret)

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

	cliCreate.AddConfig(cliCreateResource)
	cliList.AddConfig(cliListResources)
	cliUpdate.AddConfig(cliUpdateResource)
	cliDelete.AddConfig(cliDeleteResource)

	cliCreate.AddConfig(cliCreateEventTypeToPoll)
	cliList.AddConfig(cliListEventTypeToPolls)
	cliUpdate.AddConfig(cliUpdateEventTypeToPoll)
	cliDelete.AddConfig(cliDeleteEventTypeToPoll)

	cliCreate.AddConfig(cliCreateResourceType)
	cliList.AddConfig(cliListResourceTypes)
	cliUpdate.AddConfig(cliUpdateResourceType)
	cliDelete.AddConfig(cliDeleteResourceType)

	cliCreate.AddConfig(cliCreateEventType)
	cliList.AddConfig(cliListEventTypes)
	cliUpdate.AddConfig(cliUpdateEventType)
	cliDelete.AddConfig(cliDeleteEventType)

	cliLogin.SetNbArgs(2)

	cliCreateSecret.SetNbArgs(0)
	cliListSecret.SetNbArgs(0)

	cliCreateUser.SetNbArgs(3)
	cliListUsers.SetNbArgs(0)
	cliUpdateUser.SetNbArgs(1)
	cliDeleteUser.SetNbArgs(1)
	cliUpdateUser.LKey("username", verdeter.IsStr, "u", "name of the user")
	cliUpdateUser.LKey("email", verdeter.IsStr, "m", "mail of the user")
	cliUpdateUser.LKey("password", verdeter.IsStr, "p", "password of the user")

	cliCreateTenant.SetNbArgs(1)
	cliListTenants.SetNbArgs(0)
	cliUpdateTenant.SetNbArgs(1)
	cliDeleteTenant.SetNbArgs(1)
	cliUpdateTenant.LKey("tenantname", verdeter.IsStr, "t", "name of the Tenant")

	cliCreateRole.SetNbArgs(1)
	cliListRoles.SetNbArgs(0)
	cliUpdateRole.SetNbArgs(1)
	cliDeleteRole.SetNbArgs(1)
	cliUpdateRole.LKey("rolename", verdeter.IsStr, "r", "name of the Role")

	cliCreateDomain.SetNbArgs(2)
	cliListDomains.SetNbArgs(0)
	cliUpdateDomain.SetNbArgs(1)
	cliDeleteDomain.SetNbArgs(1)
	cliUpdateDomain.LKey("domainname", verdeter.IsStr, "d", "name of the Domain")

	cliCreateResource.SetNbArgs(4)
	cliListResources.SetNbArgs(0)
	cliUpdateResource.SetNbArgs(1)
	cliDeleteResource.SetNbArgs(1)
	cliUpdateResource.LKey("resourceName", verdeter.IsStr, "r", "name of the Resource")
	cliUpdateResource.LKey("logicalComponentName", verdeter.IsStr, "l", "name of the LogicalComponent")
	cliUpdateResource.LKey("domainName", verdeter.IsStr, "d", "name of the Domain")
	cliUpdateResource.LKey("resourceTypeName2", verdeter.IsStr, "y", "name of the ResourceType")

	cliCreateEventTypeToPoll.SetNbArgs(2)
	cliListEventTypeToPolls.SetNbArgs(0)
	cliUpdateEventTypeToPoll.SetNbArgs(1)
	cliDeleteEventTypeToPoll.SetNbArgs(1)
	cliUpdateEventTypeToPoll.LKey("eventtypetopollname", verdeter.IsStr, "r", "name of the EventTypeToPoll")

	cliCreateResourceType.SetNbArgs(3)
	cliListResourceTypes.SetNbArgs(0)
	cliUpdateResourceType.SetNbArgs(4)
	cliDeleteResourceType.SetNbArgs(1)
	cliUpdateResourceType.LKey("resourcetypename", verdeter.IsStr, "r", "name of the ResourceType")

	cliCreateEventType.SetNbArgs(4)
	cliListEventTypes.SetNbArgs(0)
	cliUpdateEventType.SetNbArgs(2)
	cliDeleteEventType.SetNbArgs(1)
	cliUpdateEventType.LKey("eventypename", verdeter.IsStr, "r", "name of the EventType")
}

func runLogin(cfg *verdeter.ConfigCmd, args []string) {
	name := args[0]
	password := args[1]

	fmt.Printf("gandalf cli login called with username=%s and password=%s\n", name, password)
	configurationCli := cmodels.NewConfigurationCli()
	fmt.Println(configurationCli.GetEndpoint())
	cliClient := cli.NewClient(configurationCli.GetEndpoint())

	var user models.User
	user.Name = name
	//user.Email = name
	user.Password = password
	//user := models.NewUser(name, name, password)
	token, err := cliClient.AuthenticationService.Login(user)
	if err == nil {
		fmt.Println("Token: " + token)
	} else {
		fmt.Println(err)
	}
}

func runCreateSecret(cfg *verdeter.ConfigCmd, args []string) {

	fmt.Printf("gandalf cli create secret called \n")
	configurationCli := cmodels.NewConfigurationCli()
	fmt.Println(configurationCli.GetEndpoint())
	cliClient := cli.NewClient(configurationCli.GetEndpoint())

	secret, err := cliClient.SecretAssignementService.Create(configurationCli.GetToken())
	if err == nil {
		fmt.Println(secret)
	} else {
		fmt.Println(err)
	}
}

func runListSecret(cfg *verdeter.ConfigCmd, args []string) {
	fmt.Printf("gandalf cli list secret called \n")
	configurationCli := cmodels.NewConfigurationCli()
	fmt.Println(configurationCli.GetEndpoint())
	cliClient := cli.NewClient(configurationCli.GetEndpoint())

	secrets, err := cliClient.SecretAssignementService.List(configurationCli.GetToken())
	if err == nil {
		for _, secret := range secrets {
			fmt.Println(secret)
		}
	} else {
		fmt.Println(err)
	}
}

func runCreateUser(cfg *verdeter.ConfigCmd, args []string) {
	name := args[0]
	email := args[1]
	password := args[2]

	fmt.Printf("gandalf cli create user called with username=%s, email=%s, password=%s\n", name, email, password)
	configurationCli := cmodels.NewConfigurationCli()
	cliClient := cli.NewClient(configurationCli.GetEndpoint())

	user := models.NewUser(name, email, password)
	err := cliClient.UserService.Create(configurationCli.GetToken(), user)
	if err != nil {
		fmt.Println(err)
	}

}

func runListUsers(cfg *verdeter.ConfigCmd, args []string) {
	fmt.Printf("gandalf cli list users\n")
	configurationCli := cmodels.NewConfigurationCli()
	cliClient := cli.NewClient(configurationCli.GetEndpoint())

	users, err := cliClient.UserService.List(configurationCli.GetToken())
	if err == nil {
		for _, user := range users {
			fmt.Println(user)
		}
	} else {
		fmt.Println(err)
	}

}

func runUpdateUser(cfg *verdeter.ConfigCmd, args []string) {
	name := args[0]
	newName := viper.GetString("username")
	email := viper.GetViper().GetString("email")
	password := viper.GetViper().GetString("password")
	fmt.Printf("gandalf cli update user called with username=%s, newname=%s, email=%s, password=%s\n", name, newName, email, password)
	configurationCli := cmodels.NewConfigurationCli()
	cliClient := cli.NewClient(configurationCli.GetEndpoint())

	oldUser, err := cliClient.UserService.ReadByName(configurationCli.GetToken(), name)
	if err == nil {
		user := models.NewUser(newName, email, password)
		err = cliClient.UserService.Update(configurationCli.GetToken(), int(oldUser.ID), user)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err)
	}

}

func runDeleteUser(cfg *verdeter.ConfigCmd, args []string) {
	name := args[0]
	fmt.Printf("gandalf cli delete user called with username=%s\n", name)
	configurationCli := cmodels.NewConfigurationCli()
	cliClient := cli.NewClient(configurationCli.GetEndpoint())

	oldUser, err := cliClient.UserService.ReadByName(configurationCli.GetToken(), name)
	if err == nil {
		err = cliClient.UserService.Delete(configurationCli.GetToken(), int(oldUser.ID))
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err)
	}

}

func runCreateTenant(cfg *verdeter.ConfigCmd, args []string) {
	name := args[0]
	fmt.Printf("gandalf cli create tenant called with tenant=%s\n", name)
	configurationCli := cmodels.NewConfigurationCli()
	cliClient := cli.NewClient(configurationCli.GetEndpoint())

	result, err := cliClient.CliService.Cli()
	if err == nil {
		if result == "cluster" {
			tenant := models.Tenant{Name: name}
			login, password, err := cliClient.TenantService.Create(configurationCli.GetToken(), tenant)
			if err == nil {
				fmt.Println("login : " + login)
				fmt.Println("password : " + password)
			} else {
				fmt.Println(err)
			}
		} else if result == "aggregator" {
			fmt.Println("Error: Not allowed")
		}
	}
}

func runListTenants(cfg *verdeter.ConfigCmd, args []string) {
	fmt.Printf("gandalf cli list tenants\n")
	configurationCli := cmodels.NewConfigurationCli()
	cliClient := cli.NewClient(configurationCli.GetEndpoint())

	result, err := cliClient.CliService.Cli()
	if err == nil {
		if result == "cluster" {
			tenants, err := cliClient.TenantService.List(configurationCli.GetToken())
			if err == nil {
				for _, tenant := range tenants {
					fmt.Println(tenant)
				}
			} else {
				fmt.Println(err)
			}
		} else if result == "aggregator" {
			fmt.Println("Error: Not allowed")
		}
	}

}

func runUpdateTenant(cfg *verdeter.ConfigCmd, args []string) {
	/* name := args[0]
	password := args[1]
	newName := viper.GetString("name")
	fmt.Printf("gandalf cli update tenant called with tenant=%s, newName=%s\n", name, newName)
	configurationCli := cmodels.NewConfigurationCli()
	cliClient := cli.NewClient(configurationCli.GetEndpoint())

	oldTenant, err := cliClient.TenantService.ReadByName(configurationCli.GetToken(), name)
	if err == nil {
		tenant := models.Tenant{Name: newName}
		err = cliClient.TenantService.Update(configurationCli.GetToken(), int(oldTenant.ID), tenant)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err)
	} */
}

func runDeleteTenant(cfg *verdeter.ConfigCmd, args []string) {
	name := args[0]
	fmt.Printf("gandalf cli delete tenant called with tenant=%s\n", name)

	configurationCli := cmodels.NewConfigurationCli()
	cliClient := cli.NewClient(configurationCli.GetEndpoint())

	oldTenant, err := cliClient.TenantService.ReadByName(configurationCli.GetToken(), name)
	if err == nil {
		err = cliClient.TenantService.Delete(configurationCli.GetToken(), int(oldTenant.ID))
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err)
	}

}

func runCreateRole(cfg *verdeter.ConfigCmd, args []string) {

	name := args[0]

	fmt.Printf("gandalf cli create role called with role=%s\n", name)
	configurationCli := cmodels.NewConfigurationCli()
	cliClient := cli.NewClient(configurationCli.GetEndpoint())

	role := models.Role{Name: name}
	err := cliClient.RoleService.Create(configurationCli.GetToken(), role)
	if err != nil {
		fmt.Println(err)
	}
}

func runListRoles(cfg *verdeter.ConfigCmd, args []string) {
	fmt.Printf("gandalf cli list roles\n")
	configurationCli := cmodels.NewConfigurationCli()
	cliClient := cli.NewClient(configurationCli.GetEndpoint())

	roles, err := cliClient.RoleService.List(configurationCli.GetToken())
	if err == nil {
		for _, role := range roles {
			fmt.Println(role)
		}
	} else {
		fmt.Println(err)
	}
}

func runUpdateRole(cfg *verdeter.ConfigCmd, args []string) {
	name := args[0]
	newName := viper.GetString("rolename")
	fmt.Printf("gandalf cli update role called with role=%s, newName=%s\n", name, newName)
	configurationCli := cmodels.NewConfigurationCli()
	cliClient := cli.NewClient(configurationCli.GetEndpoint())

	oldRole, err := cliClient.RoleService.ReadByName(configurationCli.GetToken(), name)
	if err == nil {
		role := models.Role{Name: newName}
		err = cliClient.RoleService.Update(configurationCli.GetToken(), int(oldRole.ID), role)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err)
	}
}

func runDeleteRole(cfg *verdeter.ConfigCmd, args []string) {
	name := args[0]
	fmt.Printf("gandalf cli delete role called with role=%s\n", name)
	configurationCli := cmodels.NewConfigurationCli()
	cliClient := cli.NewClient(configurationCli.GetEndpoint())

	oldRole, err := cliClient.RoleService.ReadByName(configurationCli.GetToken(), name)
	if err == nil {
		err = cliClient.RoleService.Delete(configurationCli.GetToken(), int(oldRole.ID))
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err)
	}

}

func runCreateDomain(cfg *verdeter.ConfigCmd, args []string) {
	name := args[0]
	parentName := args[1]
	fmt.Printf("gandalf cli create domain called with domain=%s parent=%s\n", name, parentName)

	configurationCli := cmodels.NewConfigurationCli()
	cliClient := cli.NewClient(configurationCli.GetEndpoint())
	domain := models.Domain{Name: name}
	err := cliClient.DomainService.Create(configurationCli.GetToken(), domain, parentName)
	if err != nil {
		fmt.Println(err)
	}
}

func runListDomains(cfg *verdeter.ConfigCmd, args []string) {
	fmt.Printf("gandalf cli list domains\n")
	configurationCli := cmodels.NewConfigurationCli()
	cliClient := cli.NewClient(configurationCli.GetEndpoint())

	domains, err := cliClient.DomainService.List(configurationCli.GetToken())
	if err == nil {
		for _, domain := range domains {
			fmt.Println(domain)
		}
	} else {
		fmt.Println(err)
	}

}

func runUpdateDomain(cfg *verdeter.ConfigCmd, args []string) {
	name := args[0]
	newName := viper.GetString("name")
	parent := viper.GetString("parent")
	fmt.Printf("gandalf cli update domain called with domain=%s, newName=%s, parent=%s\n", name, newName, parent)
	configurationCli := cmodels.NewConfigurationCli()
	cliClient := cli.NewClient(configurationCli.GetEndpoint())

	oldDomain, err := cliClient.DomainService.ReadByName(configurationCli.GetToken(), name)
	if err == nil {
		domain := models.Domain{Name: newName}
		err = cliClient.DomainService.Update(configurationCli.GetToken(), int(oldDomain.ID), domain)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err)
	}
}

func runDeleteDomain(cfg *verdeter.ConfigCmd, args []string) {
	name := args[0]
	fmt.Printf("gandalf cli delete domain called with domain=%s\n", name)
	configurationCli := cmodels.NewConfigurationCli()
	cliClient := cli.NewClient(configurationCli.GetEndpoint())

	oldDomain, err := cliClient.DomainService.ReadByName(configurationCli.GetToken(), name)
	if err == nil {
		err = cliClient.DomainService.Delete(configurationCli.GetToken(), int(oldDomain.ID))
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err)
	}

}

// NEW
func runCreateResource(cfg *verdeter.ConfigCmd, args []string) {
	name := args[0]
	logicalComponentName := args[1]
	domainName := args[2]
	resourceTypeName := args[3]
	fmt.Printf("gandalf cli create resource called with resource=%s", name)

	configurationCli := cmodels.NewConfigurationCli()
	cliClient := cli.NewClient(configurationCli.GetEndpoint())

	logicalComponent, err := cliClient.LogicalComponentService.ReadByName(configurationCli.GetToken(), logicalComponentName)
	fmt.Println("err")
	fmt.Println(err)
	if err == nil {
		domain, err := cliClient.DomainService.ReadByName(configurationCli.GetToken(), domainName)
		if err == nil {
			resourceType, err := cliClient.ResourceTypeService.ReadByName(configurationCli.GetToken(), resourceTypeName)
			if err == nil {
				resource := models.Resource{Name: name, LogicalComponent: *logicalComponent, Domain: *domain, ResourceType: *resourceType}
				err = cliClient.ResourceService.Create(configurationCli.GetToken(), resource)
				if err != nil {
					fmt.Println(err)
				}
			} else {
				fmt.Printf("gandalf cli can't find resourceType name=%s", resourceTypeName)
			}
		} else {
			fmt.Printf("gandalf cli can't find domain name=%s", domainName)
		}
	} else {
		fmt.Printf("gandalf cli can't find logicalcomponent name=%s", logicalComponentName)
	}
}

func runListResources(cfg *verdeter.ConfigCmd, args []string) {
	fmt.Printf("gandalf cli list resources\n")
	configurationCli := cmodels.NewConfigurationCli()
	cliClient := cli.NewClient(configurationCli.GetEndpoint())

	resources, err := cliClient.ResourceService.List(configurationCli.GetToken())
	if err == nil {
		for _, resource := range resources {
			fmt.Println(resource)
		}
	} else {
		fmt.Println(err)
	}

}

func runUpdateResource(cfg *verdeter.ConfigCmd, args []string) {
	name := args[0]
	newName := viper.GetString("resourceName")
	newLogicalComponentName := viper.GetString("logicalComponentName")
	newDomainName := viper.GetString("domainName")
	newResourceTypeName := viper.GetString("resourceTypeName2")

	fmt.Printf("gandalf cli update resource called with resource=%s, newName=%s, newLogicalComponent=%s, newDomain=%s, newResourceType=%s\n", name, newName, newLogicalComponentName, newDomainName, newResourceTypeName)
	configurationCli := cmodels.NewConfigurationCli()
	cliClient := cli.NewClient(configurationCli.GetEndpoint())

	oldResource, err := cliClient.ResourceService.ReadByName(configurationCli.GetToken(), name)
	if err == nil {
		newLogicalComponent, err := cliClient.LogicalComponentService.ReadByName(configurationCli.GetToken(), newLogicalComponentName)
		if err == nil {
			newDomain, err := cliClient.DomainService.ReadByName(configurationCli.GetToken(), newDomainName)
			if err == nil {
				newResourceType, err := cliClient.ResourceTypeService.ReadByName(configurationCli.GetToken(), newResourceTypeName)
				if err == nil {
					resource := models.Resource{Name: newName, LogicalComponent: *newLogicalComponent, Domain: *newDomain, ResourceType: *newResourceType}
					err = cliClient.ResourceService.Update(configurationCli.GetToken(), int(oldResource.ID), resource)
					if err != nil {
						fmt.Println(err)
					}
				} else {
					fmt.Printf("gandalf cli can't find resourceType name=%s", newResourceTypeName)
				}
			} else {
				fmt.Printf("gandalf cli can't find domain name=%s", newDomainName)
			}
		} else {
			fmt.Printf("gandalf cli can't find logicalcomponent name=%s", newLogicalComponentName)
		}
	} else {
		fmt.Println(err)
	}
}

func runDeleteResource(cfg *verdeter.ConfigCmd, args []string) {
	name := args[0]
	fmt.Printf("gandalf cli delete resource called with resource=%s\n", name)
	configurationCli := cmodels.NewConfigurationCli()
	cliClient := cli.NewClient(configurationCli.GetEndpoint())

	oldResource, err := cliClient.ResourceService.ReadByName(configurationCli.GetToken(), name)
	if err == nil {
		err = cliClient.ResourceService.Delete(configurationCli.GetToken(), int(oldResource.ID))
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err)
	}
}

func runCreateEventTypeToPoll(cfg *verdeter.ConfigCmd, args []string) {
	resourceName := args[0]
	eventTypeName := args[1]
	fmt.Printf("gandalf cli create eventtypetopoll called with resource=%s and eventtype=%s", resourceName, eventTypeName)

	configurationCli := cmodels.NewConfigurationCli()
	cliClient := cli.NewClient(configurationCli.GetEndpoint())
	resource, err := cliClient.ResourceService.ReadByName(configurationCli.GetToken(), resourceName)
	if err == nil {
		eventType, err := cliClient.EventTypeService.ReadByName(configurationCli.GetToken(), eventTypeName)
		if err == nil {
			eventTypeToPoll := models.EventTypeToPoll{Resource: *resource, EventType: *eventType}
			err := cliClient.EventTypeToPollService.Create(configurationCli.GetToken(), eventTypeToPoll)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err)
	}
}

func runListEventTypeToPolls(cfg *verdeter.ConfigCmd, args []string) {

	fmt.Printf("gandalf cli list eventtypetopoll\n")
	configurationCli := cmodels.NewConfigurationCli()
	cliClient := cli.NewClient(configurationCli.GetEndpoint())

	eventtypetopolls, err := cliClient.EventTypeToPollService.List(configurationCli.GetToken())
	if err == nil {
		for _, eventtypetopoll := range eventtypetopolls {
			fmt.Println(eventtypetopoll)
		}
	} else {
		fmt.Println(err)
	}

	/* fmt.Printf("gandalf cli list resources\n")
	configurationCli := cmodels.NewConfigurationCli()
	cliClient := cli.NewClient(configurationCli.GetEndpoint())

	eventTypeToPolls, err := cliClient.EventTypeToPollService.List(configurationCli.GetToken())
	if err == nil {
		for _, eventTypeToPoll := range eventTypeToPolls {
			fmt.Println(eventTypeToPoll)
		}
	} else {
		fmt.Println(err)
	} */

}

// WIP_3
func runUpdateEventTypeToPoll(cfg *verdeter.ConfigCmd, args []string) {
	/* name := args[0]
	fmt.Printf("gandalf cli update eventtypetopoll called with =%seventtypename", name)
	configurationCli := cmodels.NewConfigurationCli()
	cliClient := cli.NewClient(configurationCli.GetEndpoint())

	oldEventTypeToPoll, err := cliClient.EventTypeToPollService.ReadByName(configurationCli.GetToken(), name)
	if err == nil {
		eventTypeToPoll := models.EventTypeToPoll{Name: name}
		err = cliClient.EventTypeToPollService.Update(configurationCli.GetToken(), int(oldEventTypeToPoll.ID), eventTypeToPoll)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err)
	} */

}

func runDeleteEventTypeToPoll(cfg *verdeter.ConfigCmd, args []string) {
	resourceName := args[0]

	fmt.Printf("gandalf cli delete eventtypetopoll called with eventtypetopoll=%s\n", resourceName)

	configurationCli := cmodels.NewConfigurationCli()
	cliClient := cli.NewClient(configurationCli.GetEndpoint())

	oldEventTypeToPoll, err := cliClient.EventTypeToPollService.ReadByName(configurationCli.GetToken(), resourceName)
	if err == nil {
		err = cliClient.EventTypeToPollService.Delete(configurationCli.GetToken(), int(oldEventTypeToPoll.ID))
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err)
	}
}

func runCreateResourceType(cfg *verdeter.ConfigCmd, args []string) {

	name := args[0]
	pivotProductConnectorName := args[1]
	typeName := args[2]

	fmt.Printf("gandalf cli create eventtype called with name=%s, pivotProductConnectorName=%s", name, pivotProductConnectorName)

	configurationCli := cmodels.NewConfigurationCli()
	cliClient := cli.NewClient(configurationCli.GetEndpoint())

	if typeName == "pivot" {
		pivot, err := cliClient.PivotService.ReadByName(configurationCli.GetToken(), pivotProductConnectorName)

		if err == nil {
			resourceType := models.ResourceType{Name: name, Pivot: *pivot}
			err := cliClient.ResourceTypeService.Create(configurationCli.GetToken(), resourceType)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println(err)
		}
	} else if typeName == "productConnector" {

		productConnector, err := cliClient.ProductConnectorService.ReadByName(configurationCli.GetToken(), pivotProductConnectorName)
		if err == nil {
			resourceType := models.ResourceType{Name: name, ProductConnector: *productConnector}
			err := cliClient.ResourceTypeService.Create(configurationCli.GetToken(), resourceType)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println(err)
		}
	} else {
		fmt.Println("Error: must be connectorProduct or pivot.")
	}
}

func runListResourceTypes(cfg *verdeter.ConfigCmd, args []string) {
	fmt.Printf("gandalf cli list resources\n")
	configurationCli := cmodels.NewConfigurationCli()
	cliClient := cli.NewClient(configurationCli.GetEndpoint())

	resourceTypes, err := cliClient.ResourceTypeService.List(configurationCli.GetToken())
	if err == nil {
		for _, resourceTypes := range resourceTypes {
			fmt.Println(resourceTypes)
		}
	} else {
		fmt.Println(err)
	}

}

// WAIT - PRIO 3 - PROBLEM
func runUpdateResourceType(cfg *verdeter.ConfigCmd, args []string) {
	name := args[0]
	newName := args[1]
	pivotProductConnectorName := args[2]
	typeName := args[3]
	fmt.Printf("gandalf cli update resourcetype called with name=%s, newName=%s", name, newName)
	configurationCli := cmodels.NewConfigurationCli()
	cliClient := cli.NewClient(configurationCli.GetEndpoint())

	oldResourceType, err := cliClient.ResourceTypeService.ReadByName(configurationCli.GetToken(), name)
	if err == nil {

		if typeName == "pivot" {
			pivot, err := cliClient.PivotService.ReadByName(configurationCli.GetToken(), pivotProductConnectorName)
			if err == nil {
				resourceType := models.ResourceType{Name: newName, Pivot: *pivot}
				err = cliClient.ResourceTypeService.Update(configurationCli.GetToken(), int(oldResourceType.ID), resourceType)
				fmt.Println(err)
			} else {
				fmt.Println(err)
				fmt.Println("ERROR 404: CANNOT FIND SPECIFIED PIVOT")
			}

		} else if typeName == "productConnector" {
			productConnector, err := cliClient.ProductConnectorService.ReadByName(configurationCli.GetToken(), pivotProductConnectorName)
			if err == nil {
				resourceType := models.ResourceType{Name: newName, ProductConnector: *productConnector}
				err = cliClient.ResourceTypeService.Update(configurationCli.GetToken(), int(oldResourceType.ID), resourceType)
				fmt.Println(err)
			}
		} else {
			fmt.Println(err)
			fmt.Println("ERROR 404: CANNOT FIND SPECIFIED PRODUCTCONNECTOR")
		}
	} else {
		fmt.Println(err)
	}

}

func runDeleteResourceType(cfg *verdeter.ConfigCmd, args []string) {
	name := args[0]
	fmt.Printf("gandalf cli delete eventtypetopoll called with eventtypetopoll=%s\n", name)

	configurationCli := cmodels.NewConfigurationCli()
	cliClient := cli.NewClient(configurationCli.GetEndpoint())

	ResourceType, err := cliClient.ResourceTypeService.ReadByName(configurationCli.GetToken(), name)
	if err == nil {
		err = cliClient.ResourceTypeService.Delete(configurationCli.GetToken(), int(ResourceType.ID))
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err)
	}
}

func runCreateEventType(cfg *verdeter.ConfigCmd, args []string) {
	name := args[0]
	schema := args[1]
	pivotProductConnectorName := args[2]
	typeName := args[3]

	fmt.Printf("gandalf cli create eventtype called with name=%s, pivotProductConnectorName=%s, schema=%s\n", name, pivotProductConnectorName, schema)

	configurationCli := cmodels.NewConfigurationCli()
	cliClient := cli.NewClient(configurationCli.GetEndpoint())

	if typeName == "pivot" {
		pivot, err := cliClient.PivotService.ReadByName(configurationCli.GetToken(), pivotProductConnectorName)

		if err == nil {
			eventType := models.EventType{Name: name, Schema: schema, Pivot: *pivot}
			err := cliClient.EventTypeService.Create(configurationCli.GetToken(), eventType)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println(err)
		}
	} else if typeName == "productConnector" {

		productConnector, err := cliClient.ProductConnectorService.ReadByName(configurationCli.GetToken(), pivotProductConnectorName)
		if err == nil {
			eventType := models.EventType{Name: name, Schema: schema, ProductConnector: *productConnector}
			err := cliClient.EventTypeService.Create(configurationCli.GetToken(), eventType)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println(err)
		}
	} else {
		fmt.Println("Error: must be connectorProduct or pivot.")
	}
}

func runListEventTypes(cfg *verdeter.ConfigCmd, args []string) {
	fmt.Printf("gandalf cli list eventtypes\n")
	configurationCli := cmodels.NewConfigurationCli()
	cliClient := cli.NewClient(configurationCli.GetEndpoint())

	eventTypeToPolls, err := cliClient.EventTypeService.List(configurationCli.GetToken())
	if err == nil {
		for _, eventTypeToPoll := range eventTypeToPolls {
			fmt.Println(eventTypeToPoll)
		}
	} else {
		fmt.Println(err)
	}

}

func runUpdateEventType(cfg *verdeter.ConfigCmd, args []string) {
	name := args[0]
	newName := args[1]
	fmt.Printf("gandalf cli update eventtypetopoll called with eventtypetopoll=%s, newName=%s\n", name, newName)
	configurationCli := cmodels.NewConfigurationCli()
	cliClient := cli.NewClient(configurationCli.GetEndpoint())

	oldEventType, err := cliClient.EventTypeService.ReadByName(configurationCli.GetToken(), name)
	if err == nil {
		resourceType := models.EventType{Name: newName}
		err = cliClient.EventTypeService.Update(configurationCli.GetToken(), int(oldEventType.ID), resourceType)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err)
	}

}

func runDeleteEventType(cfg *verdeter.ConfigCmd, args []string) {
	name := args[0]
	fmt.Printf("gandalf cli delete eventtypetopoll called with eventtypetopoll=%s\n", name)

	configurationCli := cmodels.NewConfigurationCli()
	cliClient := cli.NewClient(configurationCli.GetEndpoint())

	oldEventType, err := cliClient.EventTypeService.ReadByName(configurationCli.GetToken(), name)
	if err == nil {
		err = cliClient.EventTypeService.Delete(configurationCli.GetToken(), int(oldEventType.ID))
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err)
	}
}

/*
func init() {
	rootCfg.AddConfig(cliCfg)

	cliCfg.Key("api_port", verdeter.IsInt, "", "Address to bind (default is *:9199)")
	cliCfg.SetDefault("api_port", 9199+verdeter.GetOffset())

	cliCfg.Key("database_mode", verdeter.IsStr, "", "database mode (gandalf|tenant)")
	cliCfg.SetDefault("database_mode", "gandalf")
	cliCfg.SetCheck("database_mode", func(val interface{}) bool {
		strVal := strings.ToLower(strings.TrimSpace(val.(string)))
		return map[string]bool{"gandalf": true, "tenant": true}[strVal]
	})

	cliCfg.Key("tenant", verdeter.IsStr, "", "database mode (gandalf|tenant)")
	cliCfg.SetConstraint("a tenant should be provided if database_mode == tenant",
		func() bool {
			return viper.IsSet("database_mode") && viper.GetString("database_mode") == "tenant" && viper.IsSet("tenant")
		})

	cliCfg.Key("model", verdeter.IsStr, "", "models  gandalf(authentication|cluster|tenant|role|user) || tenant(authentication|aggregator|connector|role|user)")
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
	cliCfg.Key("command", verdeter.IsStr, "", "command  (list|read|create|update|delete|upload)")
	cliCfg.SetCheck("command", func(val interface{}) bool {
		strVal := strings.ToLower(strings.TrimSpace(val.(string)))
		return map[string]bool{"list": true, "read": true, "create": true, "update": true, "delete": true, "upload": true}[strVal]
	})

	cliCfg.Key("token", verdeter.IsStr, "", "")
	cliCfg.SetConstraint("a token should be provided if model != authenticaion",
		func() bool {
			return viper.IsSet("model") && viper.GetString("model") != "authentication" && viper.IsSet("token")
		})

	cliCfg.Key("id", verdeter.IsStr, "", "id")
	cliCfg.SetConstraint("a id should be provided if command == (delete|update|read)",
		func() bool {
			return viper.IsSet("command") && (viper.GetString("command") == "delete" || viper.GetString("command") == "update" || viper.GetString("command") == "read") && viper.IsSet("id")
		})

	cliCfg.Key("value", verdeter.IsStr, "", "json")
	cliCfg.SetConstraint("a value should be provided if command == (create|update)",
		func() bool {
			return viper.IsSet("command") && (viper.GetString("command") == "create" || viper.GetString("command") == "update") && viper.IsSet("value")
		})
}
*/
