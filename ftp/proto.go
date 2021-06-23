package ftp

const (
	// FTP Commands
	User             = "USER"
	Password         = "PASS"
	Account          = "ACCT"
	ChangeWorkingDir = "CWD"
	ChangeDirUp      = "CDUP"
	StructureMount   = "SMNT"
	Quit             = "QUIT"
	Reinitialize     = "REIN"
	Port             = "PORT"
	Passive          = "PASV"
	Type             = "TYPE"
	Structure        = "STRU"
	Mode             = "MODE"
	Retrieve         = "RETR"
	Store            = "STOR"
	StoreUnique      = "STOU"
	Append           = "APPE"
	Allocate         = "ALLO"
	Restart          = "REST"
	RenameFrom       = "RNFR"
	RenameTo         = "RNTO"
	Abort            = "ABOR"
	Delete           = "DELE"
	RemoveDir        = "RMD"
	MakeDir          = "MKD"
	PrintWorkingDir  = "PWD"
	List             = "LIST"
	NameList         = "NLST"
	SiteParams       = "SITE"
	System           = "SYST"
	Status           = "STAT"
	Help             = "HELP"
	Noop             = "NOOP"
)

var Commands = map[string]bool{
	User:             true,
	Password:         true,
	Account:          true,
	ChangeWorkingDir: true,
	ChangeDirUp:      true,
	StructureMount:   true,
	Quit:             true,
	Reinitialize:     true,
	Port:             true,
	Passive:          true,
	Type:             true,
	Structure:        true,
	Mode:             true,
	Retrieve:         true,
	Store:            true,
	StoreUnique:      true,
	Append:           true,
	Allocate:         true,
	Restart:          true,
	RenameFrom:       true,
	RenameTo:         true,
	Abort:            true,
	Delete:           true,
	RemoveDir:        true,
	MakeDir:          true,
	PrintWorkingDir:  true,
	List:             true,
	NameList:         true,
	SiteParams:       true,
	System:           true,
	Status:           true,
	Help:             true,
	Noop:             true,
}

var ReplyCodes = map[int]string{
	200: "Command OK",
	500: "Syntax error, command unrecognized.",
	501: "Syntax error in parameters or arguments",
	202: "Command not implemented, superfluous at this site",
	502: "Command not implemented",
	503: "Bad sequence of commands",
	504: "Command not implemented for that parameter",
	110: "Restart marker reply", // text matches MARK YYYY = mmmm
	211: "System status, or system help reply",
	212: "Directory status",
	213: "File status",
	214: "Help message",
	215: "<NAME> system type", // where NAME is the official system type name
	120: "Service ready in nnn minutes",
	220: "Service ready for new user",
	221: "Service closing control connection",
	421: "Service not available",
	125: "Data connection already open; transfer starting",
	225: "Data connection open; no transfer in progress.",
	425: "Can't open data connection",
	226: "Closing data connection",
	426: "Connection closed; transfer aborted",
	227: "Entering passive mode",
	230: "User logged in, proceed",
	530: "Not logged in",
	331: "User name okay, need password",
	332: "Need account for login",
	532: "Need account for storing files",
	150: "File status okay; about to open connection",
	250: "Requested file action okay, completed",
	257: "PATHNAME created",
	350: "Requested file action pending further information",
	450: "Requested file action not taken",
	550: "Requested action not taken",
	451: "Requested action aborted. Local error in processing",
	551: "Requested action aborted. Page type unknown",
	452: "Requested action not taken",
	552: "Requested file action aborted. Exceeded storage allocation",
	553: "Requested action not taken. File name not allowed",
}
