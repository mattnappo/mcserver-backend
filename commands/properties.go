package commands

import (
	"errors"
	"strconv"
)

// Properties contains the data to edit the server.properties file.
type Properties struct {
	GeneratorSettings           string
	OpPermissionLevel           string
	AllowNether                 string
	LevelName                   string
	EnableQuery                 string
	AllowFlight                 string
	PreventProxyConnections     string
	ServerPort                  string
	MaxWorldSize                string
	LevelType                   string
	EnableRcon                  string
	ForceGamemode               string
	LevelSeed                   string
	ServerIP                    string
	NetworkCompressionThreshold string
	MaxBuildHeight              string
	SpawnNPCs                   string
	WhiteList                   string
	SpawnAnimals                string
	Hardcore                    string
	SnooperEnabled              string
	ResourcePackSha1            string
	OnlineMode                  string
	ResourcePack                string
	PVP                         string
	Difficulty                  string
	EnableCommandBlock          string
	Gamemode                    string
	PlayerIdleTimeout           string
	MaxPlayers                  string
	MaxTickTime                 string
	SpawnMonsters               string
	ViewDistance                string
	GenerateStructures          string
	Motd                        string
}

// DefaultProperties creates a new Properties struct with the default Minecraft properties
// execpt for server port and motd.
func DefaultProperties(port int, motd string) *Properties {
	defaultProperties := &Properties{
		GeneratorSettings:           "",
		OpPermissionLevel:           "4",
		AllowNether:                 "true",
		LevelName:                   "world",
		EnableQuery:                 "false",
		AllowFlight:                 "false",
		PreventProxyConnections:     "false",
		ServerPort:                  strconv.Itoa(port),
		MaxWorldSize:                "29999984",
		LevelType:                   "DEFAULT",
		EnableRcon:                  "false",
		ForceGamemode:               "false",
		LevelSeed:                   "",
		ServerIP:                    "",
		NetworkCompressionThreshold: "256",
		MaxBuildHeight:              "256",
		SpawnNPCs:                   "true",
		WhiteList:                   "false",
		SpawnAnimals:                "true",
		Hardcore:                    "false",
		SnooperEnabled:              "true",
		ResourcePackSha1:            "",
		OnlineMode:                  "true",
		ResourcePack:                "",
		PVP:                         "true",
		Difficulty:                  "1",
		EnableCommandBlock:          "false",
		Gamemode:                    "0",
		PlayerIdleTimeout:           "0",
		MaxPlayers:                  "20",
		MaxTickTime:                 "60000",
		SpawnMonsters:               "true",
		ViewDistance:                "10",
		GenerateStructures:          "true",
		Motd:                        motd,
	}

	return defaultProperties
}

// GetFile puts the properties from a Properties struct into
// the format to be written to a file.
func (p *Properties) GetFile() string {

	data := `#Minecraft server properties
generator-settings=` + p.GeneratorSettings + `
op-permission-level=` + p.OpPermissionLevel + `
allow-nether=` + p.OpPermissionLevel + `
level-name=` + p.LevelName + `
enable-query=` + p.EnableQuery + `
allow-flight=` + p.AllowFlight + `
prevent-proxy-connections=` + p.PreventProxyConnections + `
server-port=` + p.ServerPort + `
max-world-size=` + p.MaxWorldSize + `
level-type=` + p.LevelType + `
enable-rcon=` + p.EnableRcon + `
force-gamemode=` + p.ForceGamemode + `
level-seed=` + p.LevelSeed + `
server-ip=` + p.ServerIP + `
network-compression-threshold=` + p.NetworkCompressionThreshold + `
max-build-height=` + p.MaxBuildHeight + `
spawn-npcs=` + p.SpawnNPCs + `
white-list=` + p.WhiteList + `
spawn-animals=` + p.SpawnAnimals + `
hardcore=` + p.Hardcore + `
snooper-enabled=` + p.SnooperEnabled + `
resource-pack-sha1=` + p.ResourcePackSha1 + `
online-mode=` + p.OnlineMode + `
resource-pack=` + p.ResourcePack + `
pvp=` + p.PVP + `
difficulty=` + p.Difficulty + `
enable-command-block=` + p.EnableCommandBlock + `
gamemode=` + p.Gamemode + `
player-idle-timeout=` + p.PlayerIdleTimeout + `
max-players=` + p.MaxPlayers + `
max-tick-time=` + p.MaxTickTime + `
spawn-monsters=` + p.SpawnMonsters + `
view-distance=` + p.ViewDistance + `
generate-structures=` + p.GenerateStructures + `
motd=` + p.Motd

	return data
}

// ChangeProperty changes a single property in a Properties struct pointer.
func (p *Properties) ChangeProperty(property, newValue string) error {
	switch property {
	case "GeneratorSettings":
		p.GeneratorSettings = newValue
	case "OpPermissionLevel":
		p.OpPermissionLevel = newValue
	case "AllowNether":
		p.AllowNether = newValue
	case "LevelName":
		p.LevelName = newValue
	case "EnableQuery":
		p.EnableQuery = newValue
	case "AllowFlight":
		p.AllowFlight = newValue
	case "PreventProxyConnections":
		p.PreventProxyConnections = newValue
	case "ServerPort":
		p.ServerPort = newValue
	case "MaxWorldSize":
		p.MaxWorldSize = newValue
	case "LevelType":
		p.LevelType = newValue
	case "EnableRcon":
		p.EnableRcon = newValue
	case "ForceGamemode":
		p.ForceGamemode = newValue
	case "LevelSeed":
		p.LevelSeed = newValue
	case "ServerIP":
		p.ServerIP = newValue
	case "NetworkCompressionThreshold":
		p.NetworkCompressionThreshold = newValue
	case "MaxBuildHeight":
		p.MaxBuildHeight = newValue
	case "SpawnNPCs":
		p.SpawnNPCs = newValue
	case "WhiteList":
		p.WhiteList = newValue
	case "SpawnAnimals":
		p.SpawnAnimals = newValue
	case "Hardcore":
		p.Hardcore = newValue
	case "SnooperEnabled":
		p.SnooperEnabled = newValue
	case "ResourcePackSha1":
		p.ResourcePackSha1 = newValue
	case "OnlineMode":
		p.OnlineMode = newValue
	case "ResourcePack":
		p.ResourcePack = newValue
	case "PVP":
		p.PVP = newValue
	case "Difficulty":
		p.Difficulty = newValue
	case "EnableCommandBlock":
		p.EnableCommandBlock = newValue
	case "Gamemode":
		p.Gamemode = newValue
	case "PlayerIdleTimeout":
		p.PlayerIdleTimeout = newValue
	case "MaxPlayers":
		p.MaxPlayers = newValue
	case "MaxTickTime":
		p.MaxTickTime = newValue
	case "SpawnMonsters":
		p.SpawnMonsters = newValue
	case "ViewDistance":
		p.ViewDistance = newValue
	case "GenerateStructures":
		p.GenerateStructures = newValue
	case "Motd":
		p.Motd = newValue
	default:
		return errors.New("that is not a valid property")
	}

	return nil
}
