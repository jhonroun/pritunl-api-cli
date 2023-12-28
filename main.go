package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strconv"
)

type credential struct {
	Url    string `json:"url"`
	Token  string `json:"token"`
	Secret string `json:"secret"`
}

var ApiClient client

func main() {
	for i := 0; i < len(os.Args); i++ {
		if os.Args[i] == "debug" {
			debug_args()
		}
	}
	if len(os.Args) < 2 {
		help()
	}
	switch os.Args[1] {
	case "host":
		switch os.Args[2] {
		case "get":
			host_get()
		case "byserver":
			host_byserver()
		case "attach":
			host_attach()
		case "dettach":
			host_dettach()
		default:
			helpHandler()
		}
	case "server":
		switch os.Args[2] {
		case "get":
			server_get()
		case "getall":
			server_getall()
		case "create":
			server_create()
		case "update":
			server_update()
		case "delete":
			server_delete()
		case "start":
			server_start()
		case "stop":
			server_stop()
		default:
			helpHandler()
		}
	case "organization":
		switch os.Args[2] {
		case "get":
			organization_get()
		case "getall":
			organization_getall()
		case "byserver":
			organization_byserver()
		case "create":
			organiation_create()
		case "update":
			organization_update()
		case "delete":
			organization_delete()
		case "attach":
			organization_attach()
		case "detach":
			organization_dettach()
		default:
			helpHandler()
		}
	case "user":
		switch os.Args[2] {
		case "get":
			user_get()
		case "getall":
			user_getall()
		case "create":
			user_create()
		case "update":
			user_update()
		case "delete":
			user_delete()
		default:
			helpHandler()
		}
	case "key":
		switch os.Args[2] {
		case "get":
			key_get()
		case "urls":
			key_url()
		default:
			helpHandler()
		}
	case "route":
		switch os.Args[2] {
		case "byserver":
			route_byserver()
		case "add":
			route_add()
		case "adds":
			route_adds()
		case "update":
			route_update()
		case "delete":
			route_delete()
		default:
			helpHandler()
		}
	case "init":
		init_credentials()
	case "status":
		getStatus()
	case "all":
		helpAllHandlers()
	case "help":
		help()
	default:
		help()
	}
}

func init_credentials() {
	if len(os.Args) == 5 {
		credentials := credential{}
		credentials.Url = os.Args[2]
		credentials.Token = os.Args[3]
		credentials.Secret = os.Args[4]
		file, _ := json.Marshal(credentials)
		_ = ioutil.WriteFile("credential.json", file, 0644)
		fmt.Println("")
		fmt.Println("Credential saved!")
	} else {
		notConfigure()
	}
}

func load_credential() *credential {
	jsonFile, err := os.Open("credential.json")
	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var credentials credential
	json.Unmarshal(byteValue, &credentials)
	defer jsonFile.Close()
	return &credentials
}

func getStatus() {
	cred := load_credential()
	ApiClient := NewClient(cred.Url, cred.Token, cred.Secret, false)
	ApiClient.Status()
}

func debug_args() {
	newline()
	fmt.Println("Count of arguments: " + strconv.Itoa((len(os.Args))))
	for i := 0; i < len(os.Args); i++ {
		fmt.Println("Args # " + strconv.Itoa(i) + " is: " + os.Args[i])
	}
	newline()
}

func jsonToMap(jsonStr string) map[string]interface{} {
	result := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonStr), &result)
	if err == nil {
		fmt.Println(err)
	}
	return result
}

// GetHosts() ([]Host, error)
func host_get() {
	cred := load_credential()
	ApiClient := NewClient(cred.Url, cred.Token, cred.Secret, false)
	Hosts, err := ApiClient.GetHosts()
	if err != nil {
		fmt.Println(err)
	} else {
		for i := 0; i < len(Hosts); i++ {
			pret, _ := json.Marshal(Hosts[i])
			fmt.Println(string(pret))
		}
	}
}

// GetHostsByServer(serverId string) ([]Host, error)
func host_byserver() {
	if len(os.Args) < 4 {
		helpHost()
		os.Exit(2)
	}
	cred := load_credential()
	ApiClient := NewClient(cred.Url, cred.Token, cred.Secret, false)
	Hosts, err := ApiClient.GetHostsByServer(os.Args[3])
	if err != nil {
		fmt.Println(err)
	} else {
		for i := 0; i < len(Hosts); i++ {
			pret, _ := json.Marshal(Hosts[i])
			fmt.Println(string(pret))
		}
	}
}

// AttachHostToServer(hostId, serverId string) error
func host_attach() {
	if len(os.Args) < 5 {
		helpHost()
		os.Exit(2)
	}
	cred := load_credential()
	ApiClient := NewClient(cred.Url, cred.Token, cred.Secret, false)
	err := ApiClient.AttachHostToServer(os.Args[3], os.Args[3])
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Done")
	}
}

// DetachHostFromServer(hostId, serverId string) error
func host_dettach() {
	if len(os.Args) < 5 {
		helpHost()
		os.Exit(2)
	}
	cred := load_credential()
	ApiClient := NewClient(cred.Url, cred.Token, cred.Secret, false)
	err := ApiClient.DetachHostFromServer(os.Args[3], os.Args[3])
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Done")
	}
}

// GetServer(id string) (*Server, error)
func server_get() {
	if len(os.Args) < 4 {
		helpServer()
		os.Exit(2)
	}
	cred := load_credential()
	ApiClient := NewClient(cred.Url, cred.Token, cred.Secret, false)
	Server, err := ApiClient.GetServer(os.Args[3])
	if err != nil {
		fmt.Println(err)
	} else {
		pret, _ := json.Marshal(Server)
		fmt.Println(string(pret))
	}
}

// GetServers() ([]Server, error)
func server_getall() {
	cred := load_credential()
	ApiClient := NewClient(cred.Url, cred.Token, cred.Secret, false)
	Server, err := ApiClient.GetServers()
	if err != nil {
		fmt.Println(err)
	} else {
		for i := 0; i < len(Server); i++ {
			pret, _ := json.Marshal(Server[i])
			fmt.Println(string(pret))
		}
	}
}

// CreateServer(serverData map[string]interface{}) (*Server, error)
func server_create() {
	if len(os.Args) < 4 {
		helpServer()
		os.Exit(2)
	}
	cred := load_credential()
	ApiClient := NewClient(cred.Url, cred.Token, cred.Secret, false)
	ServerData := jsonToMap(os.Args[3])
	Server, err := ApiClient.CreateServer(ServerData)
	if err != nil {
		fmt.Println(err)
	} else {
		pret, _ := json.Marshal(Server)
		fmt.Println(string(pret))
	}
}

// UpdateServer(id string, server *Server) error
func server_update() {
	if len(os.Args) < 5 {
		helpServer()
		os.Exit(2)
	}
	cred := load_credential()
	ApiClient := NewClient(cred.Url, cred.Token, cred.Secret, false)
	ServerId := os.Args[3]
	var ServerData Server
	err := json.Unmarshal([]byte(os.Args[4]), &ServerData)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	err = ApiClient.UpdateServer(ServerId, &ServerData)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Done")
	}
}

// DeleteServer(id string) error
func server_delete() {
	if len(os.Args) < 4 {
		helpServer()
		os.Exit(2)
	}
	cred := load_credential()
	ApiClient := NewClient(cred.Url, cred.Token, cred.Secret, false)
	err := ApiClient.DeleteServer(os.Args[3])
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Done")
	}
}

// StartServer(serverId string) error
func server_start() {
	if len(os.Args) < 4 {
		helpServer()
		os.Exit(2)
	}
	cred := load_credential()
	ApiClient := NewClient(cred.Url, cred.Token, cred.Secret, false)
	err := ApiClient.StartServer(os.Args[3])
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Done")
	}
}

// StopServer(serverId string) error
func server_stop() {
	if len(os.Args) < 4 {
		helpServer()
		os.Exit(2)
	}
	cred := load_credential()
	ApiClient := NewClient(cred.Url, cred.Token, cred.Secret, false)
	err := ApiClient.StopServer(os.Args[3])
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Done")
	}
}

// GetOrganization(id string) (*Organization, error)
func organization_get() {
	if len(os.Args) < 4 {
		helpOrganization()
		os.Exit(2)
	}
	cred := load_credential()
	ApiClient := NewClient(cred.Url, cred.Token, cred.Secret, false)
	Organization, err := ApiClient.GetOrganization(os.Args[3])
	if err != nil {
		fmt.Println(err)
	} else {
		pret, _ := json.Marshal(Organization)
		fmt.Println(string(pret))
	}
}

// GetOrganizations() ([]Organization, error)
func organization_getall() {
	cred := load_credential()
	ApiClient := NewClient(cred.Url, cred.Token, cred.Secret, false)
	Organization, err := ApiClient.GetOrganizations()
	if err != nil {
		fmt.Println(err)
	} else {
		for i := 0; i < len(Organization); i++ {
			pret, _ := json.Marshal(Organization[i])
			fmt.Println(string(pret))
		}
	}
}

// GetOrganizationsByServer(serverId string) ([]Organization, error)
func organization_byserver() {
	if len(os.Args) < 4 {
		helpOrganization()
		os.Exit(2)
	}
	cred := load_credential()
	ApiClient := NewClient(cred.Url, cred.Token, cred.Secret, false)
	Organization, err := ApiClient.GetOrganizationsByServer(os.Args[3])
	if err != nil {
		fmt.Println(err)
	} else {
		for i := 0; i < len(Organization); i++ {
			pret, _ := json.Marshal(Organization[i])
			fmt.Println(string(pret))
		}
	}
}

// CreateOrganization(name string) (*Organization, error)
func organiation_create() {
	if len(os.Args) < 4 {
		helpOrganization()
		os.Exit(2)
	}
	cred := load_credential()
	ApiClient := NewClient(cred.Url, cred.Token, cred.Secret, false)
	Organization, err := ApiClient.CreateOrganization(os.Args[3])
	if err != nil {
		fmt.Println(err)
	} else {
		pret, _ := json.Marshal(Organization)
		fmt.Println(string(pret))
	}
}

// UpdateOrganization(id string, organization *Organization) error
func organization_update() {
	if len(os.Args) < 5 {
		helpOrganization()
		os.Exit(2)
	}
	cred := load_credential()
	ApiClient := NewClient(cred.Url, cred.Token, cred.Secret, false)
	OrganizationId := os.Args[3]
	var OrganizationData Organization
	err := json.Unmarshal([]byte(os.Args[4]), &OrganizationData)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	err = ApiClient.UpdateOrganization(OrganizationId, &OrganizationData)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Done")
	}
}

// DeleteOrganization(name string) error
func organization_delete() {
	if len(os.Args) < 4 {
		helpOrganization()
		os.Exit(2)
	}
	cred := load_credential()
	ApiClient := NewClient(cred.Url, cred.Token, cred.Secret, false)
	err := ApiClient.DeleteOrganization(os.Args[3])
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Done")
	}
}

// AttachOrganizationToServer(organizationId, serverId string) error
func organization_attach() {
	if len(os.Args) < 5 {
		helpOrganization()
		os.Exit(2)
	}
	cred := load_credential()
	ApiClient := NewClient(cred.Url, cred.Token, cred.Secret, false)
	err := ApiClient.AttachOrganizationToServer(os.Args[3], os.Args[4])
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Done")
	}
}

// DetachOrganizationFromServer(organizationId, serverId string) error
func organization_dettach() {
	if len(os.Args) < 5 {
		helpOrganization()
		os.Exit(2)
	}
	cred := load_credential()
	ApiClient := NewClient(cred.Url, cred.Token, cred.Secret, false)
	err := ApiClient.DetachOrganizationFromServer(os.Args[3], os.Args[4])
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Done")
	}
}

// GetUser(id string, orgId string) (*User, error)
func user_get() {
	if len(os.Args) < 5 {
		helpUser()
		os.Exit(2)
	}
	cred := load_credential()
	ApiClient := NewClient(cred.Url, cred.Token, cred.Secret, false)
	User, err := ApiClient.GetUser(os.Args[3], os.Args[4])
	if err != nil {
		fmt.Println(err)
	} else {
		pret, _ := json.Marshal(User)
		fmt.Println(string(pret))
	}
}

// GetUsers(orgId string) ([]User, error)
func user_getall() {
	if len(os.Args) < 4 {
		helpUser()
		os.Exit(2)
	}
	cred := load_credential()
	ApiClient := NewClient(cred.Url, cred.Token, cred.Secret, false)
	User, err := ApiClient.GetUsers(os.Args[3])
	if err != nil {
		fmt.Println(err)
	} else {
		for i := 0; i < len(User); i++ {
			pret, _ := json.Marshal(User[i])
			fmt.Println(string(pret))
		}
	}
}

// CreateUser(newUser User) (*User, error)
func user_create() {
	if len(os.Args) < 4 {
		helpUser()
		os.Exit(2)
	}
	cred := load_credential()
	ApiClient := NewClient(cred.Url, cred.Token, cred.Secret, false)
	var UserData User
	err := json.Unmarshal([]byte(os.Args[3]), &UserData)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	User, errx := ApiClient.CreateUser(UserData)
	if errx != nil {
		fmt.Println(err)
	} else {
		pret, _ := json.Marshal(User)
		fmt.Println(string(pret))
	}
}

// UpdateUser(id string, user *User) error
func user_update() {
	if len(os.Args) < 5 {
		helpUser()
		os.Exit(2)
	}
	cred := load_credential()
	ApiClient := NewClient(cred.Url, cred.Token, cred.Secret, false)
	UserId := os.Args[3]
	var UserData User
	err := json.Unmarshal([]byte(os.Args[4]), &UserData)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	err = ApiClient.UpdateUser(UserId, &UserData)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Done")
	}
}

// DeleteUser(id string, orgId string) error
func user_delete() {
	if len(os.Args) < 5 {
		helpUser()
		os.Exit(2)
	}
	cred := load_credential()
	ApiClient := NewClient(cred.Url, cred.Token, cred.Secret, false)
	err := ApiClient.DeleteUser(os.Args[3], os.Args[4])
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Done")
	}
}

// GetUserKeys(id string, orgId string) (map[string]string, error)
func key_get() {
	if len(os.Args) < 5 {
		helpKey()
		os.Exit(2)
	}
	cred := load_credential()
	ApiClient := NewClient(cred.Url, cred.Token, cred.Secret, false)
	Key, err := ApiClient.GetUserKeys(os.Args[3], os.Args[4])
	if err != nil {
		fmt.Println(err)
	} else {
		keys := reflect.ValueOf(Key).MapKeys()
		valueOf := keys[0]
		fmt.Println(Key[valueOf.String()])
	}
}

// GetUserKeyUrls(id string, orgId string) (*Key, error)
func key_url() {
	if len(os.Args) < 5 {
		helpKey()
		os.Exit(2)
	}
	cred := load_credential()
	ApiClient := NewClient(cred.Url, cred.Token, cred.Secret, false)
	Key, err := ApiClient.GetUserKeyUrls(os.Args[3], os.Args[4])
	if err != nil {
		fmt.Println(err)
	} else {
		pret, _ := json.Marshal(Key)
		fmt.Println(string(pret))
	}
}

// GetRoutesByServer(serverId string) ([]Route, error)
func route_byserver() {
	if len(os.Args) < 4 {
		helpRoute()
		os.Exit(2)
	}
	cred := load_credential()
	ApiClient := NewClient(cred.Url, cred.Token, cred.Secret, false)
	Routes, err := ApiClient.GetRoutesByServer(os.Args[3])
	if err != nil {
		fmt.Println(err)
	} else {
		for i := 0; i < len(Routes); i++ {
			pret, _ := json.Marshal(Routes[i])
			fmt.Println(string(pret))
		}
	}
}

// AddRouteToServer(serverId string, route Route) error
func route_add() {

}

// AddRoutesToServer(serverId string, route []Route) error
func route_adds() {
	if len(os.Args) < 5 {
		helpRoute()
		os.Exit(2)
	}
	cred := load_credential()
	ApiClient := NewClient(cred.Url, cred.Token, cred.Secret, false)
	ServerId := os.Args[3]
	var RouteData Route
	err := json.Unmarshal([]byte(os.Args[4]), &RouteData)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	var Routes []Route
	Routes = append(Routes, RouteData) // temp
	err = ApiClient.AddRoutesToServer(ServerId, Routes)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Done")
	}
}

// UpdateRouteOnServer(serverId string, route Route) error
func route_update() {
	if len(os.Args) < 5 {
		helpRoute()
		os.Exit(2)
	}
	cred := load_credential()
	ApiClient := NewClient(cred.Url, cred.Token, cred.Secret, false)
	ServerId := os.Args[3]
	var RouteData Route
	err := json.Unmarshal([]byte(os.Args[4]), &RouteData)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	err = ApiClient.UpdateRouteOnServer(ServerId, RouteData)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Done")
	}
}

// DeleteRouteFromServer(serverId string, route Route) error
func route_delete() {
	if len(os.Args) < 5 {
		helpRoute()
		os.Exit(2)
	}
	cred := load_credential()
	ApiClient := NewClient(cred.Url, cred.Token, cred.Secret, false)
	ServerId := os.Args[3]
	var RouteData Route
	err := json.Unmarshal([]byte(os.Args[4]), &RouteData)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	err = ApiClient.DeleteRouteFromServer(ServerId, RouteData)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Done")
	}
}
