package main

import (
	"fmt"
	"os"
)

func helpHandler() {
	switch os.Args[1] {
	case "host":
		helpHost()
		helpErr()
	case "server":
		helpServer()
		helpErr()
	case "organization":
		helpOrganization()
		helpErr()
	case "user":
		helpUser()
		helpErr()
	case "key":
		helpKey()
		helpErr()
	case "route":
		helpRoute()
		helpErr()
	case "init":
		notConfigure()
	default:
		help()
	}
}

func notConfigure() {
	fmt.Println("")
	fmt.Println("Error.")
	fmt.Println("Before use this API client you must configure your credential (baseUrl, apiToken, apiSecret).")
	fmt.Println("To do that just call:")
	fmt.Println("")
	fmt.Println("app baseUrl apiToken apiSecret")
	fmt.Println("")
	fmt.Println("Like:")
	fmt.Println("pritunl-api-client https://vpn.server.com token secret")
}

func help() {
	fmt.Println("Error! Invalid arguments!")
	newline()
	fmt.Println("This app provide basic API for Pritunl Server to manage your OpenVPN backend.")
	newline()
	fmt.Println("You can use all API like this structure:")
	fmt.Println("pritunl-api-client {handler} {action} {parametrs[]}")
	newline()
	fmt.Println("Avalible handlers: host, server, organization, user, key.")
	fmt.Println("Otherwise you can call specified help using \"pritunl-api-client {handler} help\".")
	newline()
	fmt.Println("For more information type \"pritunl-api-client all\"")
}

func helpErr() {
	fmt.Println("")
	fmt.Println("Typical errors:")
	fmt.Println("")
	fmt.Println("    Error on marshalling data             - look into input JSON or string.")
	fmt.Println("    Error on HTTP request                 - looks like internet coonection was interrupt or check baseUrl.")
	fmt.Println("    Non-200 response                      - maybe an errors on server side? Look server logs or check baseUrl.")
	fmt.Println("    Error on unmarshalling API response   - server return wrong JSON data. Try to restart server.")
	fmt.Println("    Empty response                        - server return empty response. Try to restart server or check baseUrl.")
	fmt.Println("    Unauthorized: Invalid token or secret - looks like your credential (baseUrl, apiToken, apiSecret).")
	fmt.Println("")
}

func hasNeeded(mustBe int) bool {
	switch mustBe {
	case len(os.Args):
		return true
	default:
		return false
	}
}

func checkArgs(mustBe int) {
	if !hasNeeded(mustBe) {
		if len(os.Args) == 2 {
			helpHandler()
		} else {
			help()
		}
		os.Exit(2)
	}
}

func helpHost() {
	fmt.Println("Handler host provide managing Pritunl Host.")
	fmt.Println("Additionaly handler accept action with parametrs:")
	fmt.Println("HostId       - (string) Host idetificator.")
	fmt.Println("ServerId     - (string) Server idetificator.")
	fmt.Println("HostData     - (json serialize to string) array of Host parametrs. Json describe as:")
	fmt.Println("")
	fmt.Println("JSON HostData {")
	fmt.Println("	ID               string   \"id,omitempty\"")
	fmt.Println("	Hostname         string   \"hostname\"")
	fmt.Println("}")
	fmt.Println("")
	fmt.Println("This handler contains next action:")
	fmt.Println("  get      - Get all Hosts. Not require parametrs. Return string error or []Json HostData.")
	fmt.Println("  byserver - Get determine Host by ServersId. Require string ServerId. Return string error or []Json HostData.")
	fmt.Println("  attach   - Attach an Host to an Server. Require string HostId and string ServerId separated by string. Return null (if success) or string error.")
	fmt.Println("  dettach  - Dettach an Host from an Server. Require string HostId and string ServerId separated by string. Return null (if success) or string error.")
}

func helpServer() {
	fmt.Println("Handler server provide managing Pritunl Servers.")
	fmt.Println("Additionaly handler accept action with parametrs:")
	fmt.Println("ServerId       - (string) Server idetificator.")
	fmt.Println("OrganizationId - (string) Organization identificator.")
	fmt.Println("ServerData     - (json serialize to string) array of Server parametrs. Json describe as:")
	fmt.Println("")
	fmt.Println("JSON ServerData {")
	fmt.Println("	ID               string   \"id,omitempty\"")
	fmt.Println("	Name             string   \"name\"")
	fmt.Println("	Protocol         string   \"protocol,omitempty\"")
	fmt.Println("	Cipher           string   \"cipher,omitempty\"")
	fmt.Println("	Hash             string   \"hash,omitempty\"")
	fmt.Println("	Port             int      \"port,omitempty\"")
	fmt.Println("	Network          string   \"network,omitempty\"")
	fmt.Println("	WG               bool     \"wg,omitempty\"")
	fmt.Println("	PortWG           int      \"port_wg,omitempty\"")
	fmt.Println("	NetworkWG        string   \"network_wg,omitempty\"")
	fmt.Println("	NetworkMode      string   \"network_mode,omitempty\"")
	fmt.Println("	NetworkStart     string   \"network_start,omitempty\"")
	fmt.Println("	NetworkEnd       string   \"network_end,omitempty\"")
	fmt.Println("	RestrictRoutes   bool     \"restrict_routes,omitempty\"")
	fmt.Println("	IPv6             bool     \"ipv6,omitempty\"")
	fmt.Println("	IPv6Firewall     bool     \"ipv6_firewall,omitempty\"")
	fmt.Println("	BindAddress      string   \"bind_address,omitempty\"")
	fmt.Println("	DhParamBits      int      \"dh_param_bits,omitempty\"")
	fmt.Println("	Groups           []string \"groups,omitempty\"")
	fmt.Println("	MultiDevice      bool     \"multi_device,omitempty\"")
	fmt.Println("	DnsServers       []string \"dns_servers,omitempty\"")
	fmt.Println("	SearchDomain     string   \"search_domain,omitempty\"")
	fmt.Println("	InterClient      bool     \"inter_client,omitempty\"")
	fmt.Println("	PingInterval     int      \"ping_interval,omitempty\"")
	fmt.Println("	PingTimeout      int      \"ping_timeout,omitempty\"")
	fmt.Println("	LinkPingInterval int      \"link_ping_interval,omitempty\"")
	fmt.Println("	LinkPingTimeout  int      \"link_ping_timeout,omitempty\"")
	fmt.Println("	InactiveTimeout  int      \"inactive_timeout,omitempty\"")
	fmt.Println("	SessionTimeout   int      \"session_timeout,omitempty\"")
	fmt.Println("	AllowedDevices   string   \"allowed_devices,omitempty\"")
	fmt.Println("	MaxClients       int      \"max_clients,omitempty\"")
	fmt.Println("	MaxDevices       int      \"max_devices,omitempty\"")
	fmt.Println("	ReplicaCount     int      \"replica_count,omitempty\"")
	fmt.Println("	VxLan            bool     \"vxlan,omitempty\"")
	fmt.Println("	DnsMapping       bool     \"dns_mapping,omitempty\"")
	fmt.Println("	PreConnectMsg    string   \"pre_connect_msg,omitempty\"")
	fmt.Println("	OtpAuth          bool     \"otp_auth,omitempty\"")
	fmt.Println("	MssFix           int      \"mss_fix,omitempty\"")
	fmt.Println("	LzoCompression   bool     \"lzo_compression,omitempty\"")
	fmt.Println("	BlockOutsideDns  bool     \"block_outside_dns,omitempty\"")
	fmt.Println("	JumboFrames      bool     \"jumbo_frames,omitempty\"")
	fmt.Println("	Debug            bool     \"debug,omitempty\"")
	fmt.Println("	Status           string   \"status,omitempty\"")
	fmt.Println("}")
	fmt.Println("")
	fmt.Println("This handler contains next action:")
	fmt.Println("  get    - Get info about determine server by Id. Require string ServerId. Return string error or Json ServerData.")
	fmt.Println("  getall - Get all Servers. Not require parametrs. Return string error or []Json ServerData.")
	fmt.Println("  create - Create new Server by information provided by ServerData. Require []string of JSON ServerData serialized as string. Return string error or Json ServerData.")
	fmt.Println("  update - Update determine (by ServerId) Server by information provided by ServerData. Require string ServerId and JSON ServerData serialized as string, all parametrs separate by space. Return null (if success) or string error.")
	fmt.Println("  delete - Delete determine Server by Id. Require ServerId as string. Return null (if success) or string error.")
	fmt.Println("  start  - Start determine Server by Id. Require ServerId as string. Return null (if success) or string error.")
	fmt.Println("  stop   - Stop determine Server by Id. Require ServerId as string. Return null (if success) or string error.")
}

func helpOrganization() {
	fmt.Println("Handler organization provide managing Pritunl Organization.")
	fmt.Println("Additionaly handler accept action with parametrs:")
	fmt.Println("OrganizationId       - (string) Organization idetificator.")
	fmt.Println("ServerId             - (string) Server idetificator.")
	fmt.Println("OrganizationName     - (string) Name of orgranization.")
	fmt.Println("OrganizationData     - (json serialize to string) array of Organization parametrs. Json describe as:")
	fmt.Println("")
	fmt.Println("JSON OrganizationData {")
	fmt.Println("	ID           string   \"id,omitempty\"")
	fmt.Println("	Name         string   \"Name\"")
	fmt.Println("}")
	fmt.Println("")
	fmt.Println("This handler contains next action:")
	fmt.Println("  get      - Get determine Organization by OrganizationId. Require string OrganizationId. Return string error or Json OrganizationData.")
	fmt.Println("  getall   - Get all Organization. Not require parametrs. Return string error or []Json OrganizationData.")
	fmt.Println("  byserver - Get all Organization attached to and Server by ServerId. Require string ServerId. Return string error or []Json OrganizationData.")
	fmt.Println("  create   - Create Organization. Require string OrganizationName. Return error or OrganiationData.")
	fmt.Println("  update   - Update determine Organization by OrganizationId. Require string OrganizationId and serialized as string OrganizationData. Return null (if success) or string error.")
	fmt.Println("  delete   - Delete determine Organization by OrganizationName. Require string OrganizationName. Return null (if success) or string error.")
	fmt.Println("  attach   - Attach an Host from an Server. Require string OrganizationId and string ServerId separated by string. Return null (if success) or string error.")
	fmt.Println("  dettach  - Dettach an Host from an Server. Require string OrganizationId and string ServerId separated by string. Return null (if success) or string error.")
}

func helpUser() {
	fmt.Println("Handler user provide managing Pritunl Users.")
	fmt.Println("Additionaly handler accept action with parametrs:")
	fmt.Println("OrganizationId       - (string) Organization idetificator.")
	fmt.Println("UserId             - (string) User idetificator.")
	fmt.Println("")
	fmt.Println("JSON UserData {")
	fmt.Println("   ID              string                   \"id,omitempty\"")
	fmt.Println("   Name            string                   \"name\"")
	fmt.Println("   Type            string                   \"type,omitempty\"")
	fmt.Println("   AuthType        string                   \"auth_type,omitempty\"")
	fmt.Println("   DnsServers      []string                 \"dns_servers,omitempty\"")
	fmt.Println("   Pin             bool                     \"pin,omitempty\"")
	fmt.Println("   DnsSuffix       string                   \"dns_suffix,omitempty\"")
	fmt.Println("   DnsMapping      string                   \"dns_mapping,omitempty\"")
	fmt.Println("   Disabled        bool                     \"disabled,omitempty\"")
	fmt.Println("   NetworkLinks    []string                 \"network_links,omitempty\"")
	fmt.Println("   PortForwarding  []map[string]interface{} \"port_forwarding,omitempty\"")
	fmt.Println("   Email           string                   \"email,omitempty\"")
	fmt.Println("   Status          bool                     \"status,omitempty\"")
	fmt.Println("   OtpSecret       string                   \"otp_secret,omitempty\"")
	fmt.Println("   ClientToClient  bool                     \"client_to_client,omitempty\"")
	fmt.Println("   MacAddresses    []string                 \"mac_addresses,omitempty\"")
	fmt.Println("   YubicoID        string                   \"yubico_id,omitempty\"")
	fmt.Println("   SSO             string                   \"sso,omitempty\"")
	fmt.Println("   BypassSecondary bool                     \"bypass_secondary,omitempty\"")
	fmt.Println("   Groups          []string                 \"groups,omitempty\"")
	fmt.Println("   Audit           bool                     \"audit,omitempty\"")
	fmt.Println("   Gravatar        bool                     \"gravatar,omitempty\"")
	fmt.Println("   OtpAuth         bool                     \"otp_auth,omitempty\"")
	fmt.Println("   Organization    string                   \"organization,omitempty\"")
	fmt.Println("}")
	fmt.Println("")
	fmt.Println("This handler contains next action:")
	fmt.Println("  get      - Get determine User attached to an Organization. Require string UserId and OrganizationId separeted by space. Return string error or Json UserData.")
	fmt.Println("  getall   - Get all User attached to an Organization.  Require string OrganizationId. Return string error or []Json UserData.")
	fmt.Println("  create   - Create an User. Require serialized as string UserData. Return string error or Json UserData.")
	fmt.Println("  update   - Update an User. Require string UserId serialized as string UserData separated by space. Return null (if success) or string error.")
	fmt.Println("  delete   - Delete an User. Require string UserId serialized as string UserData separated by space. Return null (if success) or string error.")
}

func helpKey() {
	fmt.Println("Handler key provide access Pritunl keys (ovpn config).")
	fmt.Println("Additionaly handler accept action with parametrs:")
	fmt.Println("UserId           - (string) Host idetificator.")
	fmt.Println("OrganizationId   - (string) Server idetificator.")
	fmt.Println("KeyData          - (json serialize to string) array of Host parametrs. Json describe as:")
	fmt.Println("")
	fmt.Println("JSON KeyData {")
	fmt.Println("	ID               string   \"id,omitempty\"")
	fmt.Println("	KeyUrl           string   \"key_url\"")
	fmt.Println("	KeyZipUrl        string   \"key_zip_url\"")
	fmt.Println("	KeyOncURL        string   \"key_onc_url\"")
	fmt.Println("	ViewUrl          string   \"view_url\"")
	fmt.Println("	UriUrl           string   \"uri_url\"")
	fmt.Println("}")
	fmt.Println("")
	fmt.Println("This handler contains next action:")
	fmt.Println("  get      - Get raw User keys (ovpn config file). Require string UserId and string OrganizationId separated by space. Return string error or map of strings.")
	fmt.Println("  urls     - Get url to access User keys (ovpn config file). Require string UserId and string OrganizationId separated by space. Return string error or []Json KeyData.")
}

func helpRoute() {
	fmt.Println("Handler route provide managing Pritunl Routes.")
	fmt.Println("Additionaly handler accept action with parametrs:")
	fmt.Println("HostId       - (string) Host idetificator.")
	fmt.Println("ServerId     - (string) Server idetificator.")
	fmt.Println("HostData     - (json serialize to string) array of Host parametrs. Json describe as:")
	fmt.Println("")
	fmt.Println("JSON HostData {")
	fmt.Println("	ID               string   \"id,omitempty\"")
	fmt.Println("	Hostname         string   \"hostname\"")
	fmt.Println("}")
	fmt.Println("")
	fmt.Println("This handler contains next action:")
	fmt.Println("  byserver    - Get Routers from Server. Require string ServerId. Return null (if success) or string error.")
	fmt.Println("  add         - Add an route to an Server. Require string ServerId and serealized as string RouteData. Return null (if success) or string error.")
	fmt.Println("  adds        - Add few routers to an Server. Require string ServerId and serealized as string []RouteData. Return null (if success) or string error.")
	fmt.Println("  apdate      - Delete an Route from an Server. Require string ServerId and serealized as string RouteData. Return null (if success) or string error.")
	fmt.Println("  delete      - Update an Route on an Server. Require string ServerId and serealized as string RouteData. Return null (if success) or string error.")
}

func newline() {
	fmt.Println("")
}

func helpAllHandlers() {
	help()
	newline()
	notConfigure()
	newline()
	helpErr()
	newline()
	helpHost()
	newline()
	helpServer()
	newline()
	helpRoute()
	newline()
	helpOrganization()
	newline()
	helpUser()
	newline()
	helpKey()
}
