// console_test.go

package exporter

import (
	"testing"
	"time"

	"github.com/asiffer/netspot/config"
)

const consolePrefix = "exporter.console"

func init() {
	m := map[string]interface{}{
		consolePrefix + ".data":  true,
		consolePrefix + ".alarm": true,
	}
	if err := config.LoadForTest(m); err != nil {
		panic(err)
	}
}

func TestConsoleBadInit(t *testing.T) {
	title(t.Name())
	m := map[string]interface{}{
		consolePrefix + ".data":  false,
		consolePrefix + ".alarm": false,
	}
	if err := config.LoadForTest(m); err != nil {
		panic(err)
	}

	c := available["console"]
	checkTitle("Empty config")
	if c.Init() != nil {
		testERROR()
	} else {
		testOK()
	}
}

func TestInitStartCloseConsole(t *testing.T) {
	title(t.Name())

	if err := setFullConfig(); err != nil {
		t.Fatal(err)
	}

	c := Console{}

	checkTitle("Initialization")
	if err := c.Init(); err != nil {
		testERROR()
		t.Fatal(err)
	} else {
		testOK()
		defer Unload(c.Name())
	}

	checkTitle("Check exporting data")
	if !c.LogsData() {
		testERROR()
		t.Fatalf("The exporting module does not send data")
	} else {
		testOK()
	}

	checkTitle("Check exporting alarm")
	if !c.LogsAlarm() {
		testERROR()
		t.Fatalf("The exporting module does not send alarms")
	} else {
		testOK()
	}

	checkTitle("Start")
	if err := c.Start("wtf"); err != nil {
		testERROR()
		t.Error(err)
	} else {
		testOK()
	}

	checkTitle("Close")
	if err := c.Close(); err != nil {
		testERROR()
		t.Error(err)
	} else {
		testOK()
	}
}

func TestConsoleWriteAndWarn(t *testing.T) {
	title(t.Name())
	if err := setFullConfig(); err != nil {
		t.Error(err)
	}

	c := &Console{}

	if err := c.Init(); err != nil {
		t.Error(err)
	}
	defer Unload(c.Name())
	t.Log(c)

	// prepare data
	now := time.Now()
	data := map[string]float64{"stat0": 15.2, "stat1": -3.33333333}

	checkTitle("Writing data")

	if err := c.Write(now, data); err != nil {
		testERROR()
		t.Error(err)
	} else {
		testOK()
	}

	// prepare data
	alert := SpotAlert{
		Status:      "DOWN_ALERT",
		Stat:        "R_ACK",
		Value:       0.00371,
		Code:        -1,
		Probability: 1.2e-12,
	}
	checkTitle("Sending alarm")

	if err := c.Warn(now, &alert); err != nil {
		testERROR()
		t.Error(err)
	} else {
		testOK()
	}

	checkTitle("Closing")

	if err := c.Close(); err != nil {
		testERROR()
		t.Error(err)
	} else {
		testOK()
	}

}
