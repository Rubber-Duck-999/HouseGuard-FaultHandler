package main

type FHConfigTypes struct {
	EmailSettings struct {
		Email    string `yaml:"email"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
		ToEmail  string `yaml:"to_email"`
	} `yaml:"email_settings"`
	MessageSettings struct {
		Password string `yaml:"passcode"`
	} `yaml:"message_settings"`
}

type Basic struct {
	Message string `yaml:"message"`
}

type FailureMessage struct {
	Time         string `json:"time"`
	Failure_type string `json:"type_of_failure"`
}

type MotionResponse struct {
	File     string `json:"file"`
	Time     string `json:"time"`
	Severity int    `json:"severity"`
}
type MonitorState struct {
	State bool
}

type AlarmEvent struct {
	User  string `string:"user"`
	State string `string:"state"`
}

type MapMessage struct {
	message     string
	routing_key string
	time        string
	valid       bool
}

type DeviceFound struct {
	Device_name string `json:"name"`
	Ip_address  string `json:"address"`
	Status      int    `json:"status"`
}

type StatusFH struct {
	LastFault string `json:"last_fault"`
}

type Fault struct {
	Count int
	Name  string
}

const FAILURE string = "Failure.*"
const FAILURENETWORK string = "Failure.Network"     //Level 5
const FAILURECOMPONENT string = "Failure.Component" //Level 3
const FAILUREACCESS string = "Failure.Access"       //Level 6
const FAILURECAMERA string = "Failure.Camera"       // Level 2
const MOTIONRESPONSE string = "Motion.Response"
const MOTIONEVENT string = "Motion.Event"

const DEVICEFOUND string = "Device.Found"
const MONITORSTATE string = "Monitor.State"
const CAMERASTART string = "Camera.Start"
const CAMERASTOP string = "Camera.Stop"
const ALARMEVENT string = "Alarm.Event"
const STATUSFH string = "Status.FH"
const EXCHANGENAME string = "topics"
const EXCHANGETYPE string = "topic"
const TIMEFORMAT string = "2006/01/02 15:04:05"

//
const DEVICE_TITLE string = "New Device - Network"
const DEVICEBLOCKED_MESSAGE string = "A blocked device has joined the\n" +
	"network. Device name: "
const DEVICEUNKNOWN_MESSAGE string = "A unknown device has joined the\n" +
	"network. Device name: "
const IPADDRESS string = ". IP Address: "

//
const ACTIVATE_TITLE string = "Alarm has been activated"
const ACT_MESSAGE string = "The alarm state has been changed.\n" +
	"Please ensure that whoever enacted this " +
	"was authorised to do so"
const DEACTIVATE_TITLE string = "Alarm has been deactivated"

//
const UPDATESTATE string = "Motion state changed"
const SERVERERROR string = "Server is failing to send"
const MOTIONMESSAGE string = "There was movement in the property. \n Head to the drive space and check the image taken by HouseGuard"

//
const GUIDUPDATE_TITLE string = "Daily GUID Key Inside"
const GUIDUPDATE_MESSAGE string = "Key: "

//
const BOTH_ROLE string = "BOTH"
const ADMIN_ROLE string = "ADMIN"

//
const STATEUPDATESEVERITY int = 2
const SERVERSEVERITY int = 4
const BLOCKED int = 2
const UNKNOWN int = 4
const FAILURECONVERT string = "Failed to convert"
const FAILUREPUBLISH string = "Failed to publish"

var SubscribedMessagesMap map[uint32]*MapMessage
var key_id uint32 = 0
