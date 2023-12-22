# pritunl-api-cli
## _Simple, Easy, CLI_

[![Build Status](https://travis-ci.org/jhonroun/pritunl-api-cli.svg?branch=master)](https://travis-ci.org/jhonroun/pritunl-api-cli)

This app provide basic API for Pritunl Server to manage your OpenVPN backend.
Support Pritunl ver. 1.32.3504.68
- Writen in Go
- Contain all API docs
- ✨Magic ✨

## Features

- Simple to use: init credential and use it
- All API handlers fully described
- You know what you get. It means that all API handlers and action description containt return data type info

## Overview
- [Instaltions](##-1.-Instalations)
  - [Build from source](####-Building-for-source)
  - [Download](####-Download)
- [Usage](##-Usage)
  - [Init](###-Init-first)
  - [Check](###-Now-check-how-it-works)
- [API CLI Information](##-2.-API-CLI-Information)
  - [help](##-Avalible-handlers)
  - [errors](##-Typical-errors:)
  - [host](###-Host)
  - [server](###-Server)
  - [organization](###-Organization)
  - [user](###-User)
  - [device](###-Device)
  - [key](###-Key) 
- [Feedback](##-PM-me)
- [in the end](##-License)

## 1. Instalations

#### Building for source

Just:

```sh
mkdir pritunl-api
cd pritunl-api
git clone https://github.com/jhonroun/pritunl-api-cli.git
cd pritunl-api-cli
go build
```
#### Download

- See in [Realeses](https://github.com/jhonroun/pritunl-api-cli/releases)

## Usage

### Init first
> [!NOTE]
> Before use this API client you must configure your credential (baseUrl, apiToken, apiSecret).

To do that just call it like:
```sh
pritunl-api-client https://vpn.server.com token secret
```
### Now check how it works
```sh
pritunl-api-client status
```

Output must look like that:
```javascript
{
    "org_count": 4,
    "users_online": 5,
    "user_count": 20,
    "servers_online": 4,
    "server_count": 4,
    "hosts_online": 1,
    "host_count": 1,
    "server_version": "1.32.3504.68",
    "current_host": "host_id",
    "public_ip": "xxx.xxx.xxx.xxx",
    "local_networks": [
        "xxx.xxx.xxx.xxx/24"
    ],
    "notification": ""
}
```
## 2. API-CLI Information

## Avalible handlers

- [host](###-Host)
- [server](###-Server)
- [organization](###-Organization)
- [user](###-User)
- [device](###-Device)
- [key](###-Key) 

Otherwise you can call specified help using 
```sh
pritunl-api-client ${handler} help
```
or for more information type
```sh
pritunl-api-client all
```

## Typical errors:

| Error | Desribe |
| ------ | ------ |
|  Error on marshalling data             | look into input JSON or string|
|  Error on HTTP request                 | looks like internet coonection was interrupt or - check baseUrl |
| Non-200 response                       | maybe an errors on server side? Look server logs or check baseUrl|
| Error on unmarshalling API response    | server return wrong JSON data. Try to restart server|
| Empty response                         | server return empty response. Try to restart server or check baseUrl|
| Unauthorized: Invalid token or secret  | looks like your credential (baseUrl, apiToken, apiSecret)|

### Host

Handler host provide managing Pritunl Host.
Additionaly handler accept action with parametrs:
| Parametr | Type | Describe |
| ------ | ------ | ------ |
| HostId       |`string`| Host idetificator|
| ServerId     |`string`| Server idetificator|
| HostData     |`Json` | Array of Host parametrs|

Json describe as:
```javascript
JSON HostData {
	"ID"               string  "id,omitempty"
	"Hostname"         string  "hostname"
}
```
This handler contains next action:
| Action | Describe | Parameters | Return type |
| ------ | ------ | ------ | ------ |
|  get | Get all Hosts | Not require | `string` [error](##Typicalerrors) or `[]Json` [HostData](###HostData)|
|  byserver | Get determine Host by `ServersId` | `string` ServerId | `string` [error](##Typicalerrors) or `Json` [HostData](###HostData)|
| attach | Attach an Host to an Server | Require `string` HostId and `string` ServerId separated by string | `null` (if success) or `string` [error](##Typicalerrors).
|  dettach | Dettach an Host from an Server | Require string HostId and string ServerId separated by string| `null` (if success) or `string` error|

### Server
Handler server provide managing Pritunl Servers.
Additionaly handler accept action with parametrs:
| Parametr | Type | Describe |
| ------ | ------ | ------ |
|ServerId|       `string`| Server idetificator|
|OrganizationId| `string` |Organization identificator|
|ServerData|     `Json`| Array of Server parametrs (Name, ans others)| 

Json describe as:
```javascript
JSON ServerData {
	"ID"               string   "id,omitempty"
	"Name"             string   "name"
	"Protocol"         string   "protocol,omitempty"
	"Cipher"           string   "cipher,omitempty"
	"Hash"             string   "hash,omitempty"
	"Port"             int      "port,omitempty"
	"Network"          string   "network,omitempty"
	"WG"               bool     "wg,omitempty"
	"PortWG"           int      "port_wg,omitempty"
	"NetworkWG"        string   "network_wg,omitempty"
	"NetworkMode"      string   "network_mode,omitempty"
	"NetworkStart"     string   "network_start,omitempty"
	"NetworkEnd"       string   "network_end,omitempty"
	"RestrictRoutes"   bool     "restrict_routes,omitempty"
	"IPv6"             bool     "ipv6,omitempty"
	"IPv6Firewall"     bool     "ipv6_firewall,omitempty"
	"BindAddress"      string   "bind_address,omitempty"
	"DhParamBits"      int      "dh_param_bits,omitempty"
	"Groups"           []string "groups,omitempty"
	"MultiDevice"      bool     "multi_device,omitempty"
	"DnsServers"       []string "dns_servers,omitempty"
	"SearchDomain"     string   "search_domain,omitempty"
	"InterClient"      bool     "inter_client,omitempty"
	"PingInterval"     int      "ping_interval,omitempty"
	"PingTimeout"      int      "ping_timeout,omitempty"
	"LinkPingInterval" int      "link_ping_interval,omitempty"
	"LinkPingTimeout"  int      "link_ping_timeout,omitempty"
	"InactiveTimeout"  int      "inactive_timeout,omitempty"
	"SessionTimeout"   int      "session_timeout,omitempty"
	"AllowedDevices"   string   "allowed_devices,omitempty"
	"MaxClients"       int      "max_clients,omitempty"
	"MaxDevices"       int      "max_devices,omitempty"
	"ReplicaCount"     int      "replica_count,omitempty"
	"VxLan"            bool     "vxlan,omitempty"
	"DnsMapping"       bool     "dns_mapping,omitempty"
	"PreConnectMsg"    string   "pre_connect_msg,omitempty"
	"OtpAuth"          bool     "otp_auth,omitempty"
	"MssFix"           int      "mss_fix,omitempty"
	"LzoCompression"   bool     "lzo_compression,omitempty"
	"BlockOutsideDns"  bool     "block_outside_dns,omitempty"
	"JumboFrames"      bool     "jumbo_frames,omitempty"
	"Debug"            bool     "debug,omitempty"
	"Status"           string   "status,omitempty"
}
```

This handler contains next action:
| Action | Describe | Parameters | Return type |
| ------ | ------ | ------ | ------ |
| get    | Get info about determine server by Id| Require `string` ServerId| `string` [error](##Typicalerrors) or `Json` [ServerData](###ServerData)|
| getall | Get all Servers| Not require parametrs.| `string` [error](##Typicalerrors) or `[]Json` [ServerData](###ServerData)|
| create | Create new Server by information provided by ServerData| Require `[]string` of `JSON` [ServerData](###ServerData) serialized as `string`| `string` [error](##Typicalerrors) or `[]Json` [ServerData](###ServerData)|
| update | Update determine (by ServerId) Server by information provided by [ServerData](###ServerData) | Require `string` ServerId and [ServerData](###ServerData) serialized as string, all parametrs separate by space |  `null` (if success) or `string` [error](##Typicalerrors)|
| delete | Delete determine Server by Id| Require `string` ServerId | `null` (if success) or `string` [error](##Typicalerrors)|
| start  | Start determine Server by Id| Require `string` ServerId |`null` (if success) or `string` [error](##Typicalerrors)|
| stop   | Stop determine Server by Id| Require `string` ServerId |`null` (if success) or `string` [error](##Typicalerrors)|

### Route

Handler route provide managing Pritunl Routes.
Additionaly handler accept action with parametrs:
| Parametr | Type | Describe |
| ------ | ------ | ------ |
|ServerId     |`string`| Server idetificator|
|RoteData     |`Json`| Array of Route parametrs| 

Json describe as:
```javascript
JSON RouteData {
	"Network"        string "network"
	"Nat"            bool   "nat"
	"Comment"        string "comment,omitempty"
	"VirtualNetwork" bool   "virtual_network,omitempty"
	"WgNetwork"      string "wg_network,omitempty"
	"NetworkLink"    bool   "network_link,omitempty"
	"ServerLink"     bool   "server_link,omitempty"
	"NetGateway"     bool   "net_gateway,omitempty"
	"VpcID"          string "vpc_id,omitempty"
	"VpcRegion"      string "vpc_region,omitempty"
	"Metric"         string "metric,omitempty"
	"Advertise"      bool   "advertise,omitempty"
	"NatInterface"   string "nat_interface,omitempty"
	"NatNetmap"      string "nat_netmap,omitempty"
}
```

This handler contains next action:
| Action | Describe | Parameters | Return type |
| ------ | ------ | ------ | ------ |
| get    | Get info about determine server by Id| Require `string` ServerId| `string` [error](##Typicalerrors) or `Json` [ServerData](###ServerData)|
|  byserver    | Get Routers from Server| Require `string` ServerId| `null` (if success) or `string` [error](##Typicalerrors)|
|  add         | Add an route to an Server| Require `string` ServerId and serealized as `string`  [RouteData](###RouteData)| `null` (if success) or `string` [error](##Typicalerrors)|
|  adds        | Add few routers to an Server| Require `string` ServerId and serealized as `string` [][RouteData](###RouteData)| `null` (if success) or `string` [error](##Typicalerrors)|
|  apdate      | Delete an Route from an Server| Require `string` ServerId and serealized as `string` [RouteData](###RouteData)| `null` (if success) or `string` [error](##Typicalerrors)|
|  delete      | Update an Route on an Server| Require `string` ServerId and serealized as string [RouteData](###RouteData)| `null` (if success) or `string` [error](##Typicalerrors)|

### Organization

Handler organization provide managing Pritunl Organization.
Additionaly handler accept action with parametrs:
| Parametr | Type | Describe |
| ------ | ------ | ------ |
|OrganizationId       | `string` |Organization idetificator|
|ServerId             | `string` |Server idetificator|
|OrganizationName     | `string` |Name of orgranization|
|OrganizationData     | `Json` | Array of Organization parametrs|

Json describe as:
```javascript
JSON OrganizationData {
	"ID"           string   "id,omitempty"
	"Name"         string   "Name"
}
```


This handler contains next action:
| Action | Describe | Parameters | Return type |
| ------ | ------ | ------ | ------ |
|  get      | Get determine Organization by OrganizationId| Require `string` OrganizationId| `string` [error](##Typicalerrors) or `Json` [OrganizationData](###OrganizationData)|
|  getall   | Get all Organization| Not require parametrs| `string` [error](##Typicalerrors) or `[]Json` [OrganizationData](###OrganizationData)|
|  byserver | Get all Organization attached to and Server by ServerId| Require `string` ServerId| `string` [error](##Typicalerrors) or `[]Json` [OrganizationData](###OrganizationData)|
|  create   | Create Organization| Require `string` OrganizationName| `string` [error](##Typicalerrors) or `Json` [OrganizationData](###OrganizationData)|
|  update   | Update determine Organization by OrganizationId| Require `string` OrganizationId and serialized as string [OrganizationData](###OrganizationData)| `null` (if success) or `string` [error](##Typicalerrors)|
|  delete   | Delete determine Organization by OrganizationName| Require `string` OrganizationName| `null` (if success) or `string` [error](##Typicalerrors)|
|  attach   | Attach an Host from an Server| Require `string` OrganizationId and `string` ServerId separated by space| `null` (if success) or `string` [error](##Typicalerrors)|
|  dettach  | Dettach an Host from an Server| Require `string` OrganizationId and `string` ServerId separated by string| `null` (if success) or `string` [error](##Typicalerrors)|

### User
Handler user provide managing Pritunl Users.
Additionaly handler accept action with parametrs:
| Parametr | Type | Describe |
| ------ | ------ | ------ |
|OrganizationId     | `string`| Organization idetificator|
|UserId             | `string`| User idetificator|
|UserData             | `JSON`|Array of User parameters|

Json describe as:
```javascript
JSON UserData {
  "ID"              string                   "id,omitempty"
  "Name"            string                   "name"
  "Type"            string                   "type,omitempty"
  "AuthType"        string                   "auth_type,omitempty"
  "DnsServers"      []string                 "dns_servers,omitempty"
  "Pin"             bool                     "pin,omitempty"
  "DnsSuffix"       string                   "dns_suffix,omitempty"
  "DnsMapping"      string                   "dns_mapping,omitempty"
  "Disabled"        bool                     "disabled,omitempty"
  "NetworkLinks"    []string                 "network_links,omitempty"
  "PortForwarding"  []map(string)             "port_forwarding,omitempty"
  "Email"           string                   "email,omitempty"
  "Status"          bool                     "status,omitempty"
  "OtpSecret"       string                   "otp_secret,omitempty"
  "ClientToClient"  bool                     "client_to_client,omitempty"
  "MacAddresses"    []string                 "mac_addresses,omitempty"
  "YubicoID"        string                   "yubico_id,omitempty"
  "SSO"             string                   "sso,omitempty"
  "BypassSecondary" bool                     "bypass_secondary,omitempty"
  "Groups"          []string                 "groups,omitempty"
  "Audit"           bool                     "audit,omitempty"
  "Gravatar"        bool                     "gravatar,omitempty"
  "OtpAuth"         bool                     "otp_auth,omitempty"
  "Organization"    string                   "organization,omitempty"
}
```

This handler contains next action:
| Action | Describe | Parameters | Return type |
| ------ | ------ | ------ | ------ |
|  get      | Get determine User attached to an Organization| Require `string` UserId and `string` OrganizationId separeted by space|  `string` [error](##Typicalerrors) or `Json` [UserData](###UserData)|
|  getall   | Get all User attached to an Organization|  Require `string` OrganizationId|  `string` [error](##Typicalerrors) or `[]Json` [UserData](###UserData)|
|  create   |Create an User| Require serialized as `string` [UserData](###UserData)|  `string` [error](##Typicalerrors) or `Json` [UserData](###UserData)|
|  update   | Update an User| Require `string` UserId serialized as `string` [UserData](###UserData) separated by space|  `null` (if success) or `string` [error](##Typicalerrors)|
|  delete   | Delete an User| Require `string` UserId serialized as `string` [UserData](###UserData) separated by space| `null` (if success) or `string` [error](##Typicalerrors)|

### Device 

> [!WARNING]
> Under develop. Not implemented yet.

```javascript
JSON DeviceData {
	"Name"   string "name""
	"RegKey" string "reg_key"
}
```
1. Unregister
> Method: "GET",
>	Path:   "/device/unregistered"
2. Update
> Method: "PUT",
> Path:   "/device/register/" + orgId + "/" + userId + "/" + deviceId
3. Delete
> Method: "DELETE",
> Path:   "/device/register/" + orgId + "/" + userId + "/" + deviceId

### Key

Handler key provide access Pritunl keys (ovpn config).
Additionaly handler accept action with parametrs:
| Parametr | Type | Describe |
| ------ | ------ | ------ |
|OrganizationId     | `string`| Organization idetificator|
|UserId             | `string`| User idetificator|
|KeyData            |`Json` | Array of Host parametrs.|

Json describe as:
```javascript
JSON KeyData {
	ID               string   "id,omitempty"
	KeyUrl           string   "key_url"
	KeyZipUrl        string   "key_zip_url"
	KeyOncURL        string   "key_onc_url"
	ViewUrl          string   "view_url"
	UriUrl           string   "uri_url"
}
```
This handler contains next action:
| Action | Describe | Parameters | Return type |
| ------ | ------ | ------ | ------ |
| get      | Get `raw` User keys (`ovpn` config file)| Require `string` UserId and `string` OrganizationId separated by space| `string` [error] or `map` of `strings`|
| urls     | Get url to access User keys (ovpn config file)| Require `string` UserId and `string` OrganizationId separated by space| `string` [error] or `[]Json` [KeyData](###KeyData).

## PM me
[Telegram](t.me/lemtech)

## License

MIT

Inspired by [go-pritunl](https://github.com/drdaeman/go-pritunl)

