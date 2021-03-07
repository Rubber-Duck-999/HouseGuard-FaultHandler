// rabbitmq_test.go

package main

import (
	"strings"
	"testing"
)

// Check that State is set
// then run this test will prove it is set
/*
func TestPublishFailRabbit(t *testing.T) {
	file := "../FH.yml"
	var data houseguard.FHConfigTypes
	if houseguard.Exists(file) {
		houseguard.GetData(&data, file)
	} else {
		t.Error("File doesn't exist")
	}
	if data.EmailSettings.Email != "" {
		SetEmailSettings(data.EmailSettings.Email,
			data.EmailSettings.Password,
			data.EmailSettings.Name,
			data.EmailSettings.ToEmail)
		SetPassword(data.MessageSettings.Password)
	}
}*/

func TestLogicNetwork(t *testing.T) {
	value := "{ 'time': 12:00:34, 'type': 'Camera', 'severity': 3 }"
	messages(FAILURENETWORK, value)
	checkState()
	if SubscribedMessagesMap[0].valid == true {
		t.Error("Failure")
	} else if SubscribedMessagesMap[0].routing_key != FAILURENETWORK {
		t.Log(SubscribedMessagesMap[0].routing_key)
		t.Error("Failure")
	}
}

/*func TestLogicNotExpected(t *testing.T) {
	value := "{ 'time': 12:00:34, 'type': 'Camera', 'severity': 3 }"
	messages("Event.DBM", value)
	checkState()
	if SubscribedMessagesMap[1].valid == true {
		t.Error("Failure")
	} else if SubscribedMessagesMap[1].routing_key == ALARMEVENT {
		t.Log(SubscribedMessagesMap[1].routing_key)
		t.Error("Failure")
	}
}*/

/*func TestLogicValid(t *testing.T) {
	value := "{ 'time': 12:00:34, 'type': 'Camera', 'severity': 3 }"
	messages("Motion.Response", value)
	if SubscribedMessagesMap[2].valid == false {
		t.Error("Failure")
	} else if SubscribedMessagesMap[2].routing_key == CAMERASTART {
		t.Log(SubscribedMessagesMap[2].routing_key)
		t.Error("Failure")
	}
}*/

/*func TestLogicRequestPower(t *testing.T) {
	value := "{ 'time': 12:00:34, 'type': 'Camera', 'severity': 3 }"
	messages("Motion.Respons", value)
	checkState()
	if SubscribedMessagesMap[3].valid == true {
		t.Error("Failure")
	} else if SubscribedMessagesMap[2].routing_key == CAMERASTART {
		t.Log(SubscribedMessagesMap[2].routing_key)
		t.Error("Failure")
	}
}*/

func TestGetTime(t *testing.T) {
	time := getTime()
	if !strings.Contains(time, "2021") {
		t.Error("Failure")
	}
}

func TestGetTimeFail(t *testing.T) {
	time := getTime()
	if strings.Contains(time, "-") {
		t.Error("Failure")
	}
}

/*func TestStatusFH(t *testing.T) {
	valid := publishStatusFH()
	if valid != "" {
		t.Error("Failure")
	}
}*/

func TestEmailSettingsFail(t *testing.T) {
	shutdown_valid := SetEmailSettings("email_to", "password", "from_name", "to_email")
	if !shutdown_valid {
		t.Error("Failure")
	}
}

/*func TestIssueNotice(t *testing.T) {
	value := "{ 'severity': 1, 'component': 'CM', 'action': null }"
	messages("Issue.Notice", value)
	checkState()
	if SubscribedMessagesMap[4].valid != false {
		t.Error("Failure")
	} else if SubscribedMessagesMap[4].routing_key == CAMERASTART {
		t.Log(SubscribedMessagesMap[4].routing_key)
		t.Error("Failure")
	}
}*/

/*func TestRequestPower(t *testing.T) {
	value := "{ 'time': '14:00:20', 'failure_type': 'Power loss' }"
	messages(FAILURECAMERA, value)
	checkState()
	if SubscribedMessagesMap[5].valid != false {
		t.Log(SubscribedMessagesMap[5].routing_key)
		t.Error("Failure")
	} else if SubscribedMessagesMap[5].routing_key != FAILURECAMERA {
		t.Log(SubscribedMessagesMap[5].routing_key)
		t.Error("Failure")
	}
}*/

/*func TestAllFailures(t *testing.T) {
	value := "{ 'time': '14:00:20', 'failure_type': 'Power loss' }"
	messages(FAILURENETWORK, value)
	messages(FAILURECOMPONENT, value)
	messages(FAILUREACCESS, value)
	value = "{ 'state': true }"
	messages(MONITORSTATE, value)
	checkState()
	if SubscribedMessagesMap[9].valid != false {
		t.Log(SubscribedMessagesMap[5].routing_key)
		t.Error("Failure")
	} else if SubscribedMessagesMap[9].routing_key == FAILURECAMERA {
		t.Log(SubscribedMessagesMap[9].routing_key)
		t.Error("Failure")
	}
}*/

func TestEmailSettings(t *testing.T) {
	failure := SetEmailSettings("", "", "", "")
	if failure == false {
		t.Error("Failure")
	}
}
