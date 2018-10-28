package main

import (
	"flag"
	"fmt"
	"github.com/quickfixgo/fix40/heartbeat"
	"os"
	//"github.com/quickfixgo/examples/cmd/internal"
	"./internal"
	"github.com/quickfixgo/quickfix"
)

//TradeClient implements the quickfix.Application interface
type TradeClient struct {
}

//OnCreate implemented as part of Application interface
func (e TradeClient) OnCreate(sessionID quickfix.SessionID) {
	return
}

//OnLogon implemented as part of Application interface
func (e TradeClient) OnLogon(sessionID quickfix.SessionID) {
	return
}

//OnLogout implemented as part of Application interface
func (e TradeClient) OnLogout(sessionID quickfix.SessionID) {
	return
}

//FromAdmin implemented as part of Application interface
func (e TradeClient) FromAdmin(msg *quickfix.Message, sessionID quickfix.SessionID) (reject quickfix.MessageRejectError) {

	//t, _ := msg.MsgType()
	//fmt.Printf("From Admin type %s\n", t)
	//fmt.Printf("From Admin %s\n", msg.String())
	return
}

//ToAdmin implemented as part of Application interface
func (e TradeClient) ToAdmin(msg *quickfix.Message, sessionID quickfix.SessionID) {
	//fmt.Printf("To Admin %s\n", msg)
	return
}

//ToApp implemented as part of Application interface
func (e TradeClient) ToApp(msg *quickfix.Message, sessionID quickfix.SessionID) (err error) {
	//fmt.Printf("Sending %s\n", msg)
	return
}

//FromApp implemented as part of Application interface. This is the callback for all Application level messages from the counter party.
func (e TradeClient) FromApp(msg *quickfix.Message, sessionID quickfix.SessionID) (reject quickfix.MessageRejectError) {

	fmt.Printf("\n\nFromApp: %s\n", msg.String())
	//showMenu()
	return
}

func main() {
	flag.Parse()

	//cfgFileName := path.Join("config", "tradeclient.cfg")
	//cfgFileName := path.Join("tradeclient.cfg")
	//cfgFileName, err := filepath.Abs("tradeclient.cfg")
	//cfgFileName := path.Join("tradeclient.cfg")
	cfgFileName := "config/tradeclient.cfg"
	if flag.NArg() > 0 {
		fmt.Printf("Error opening " + flag.Arg(0))
		cfgFileName = flag.Arg(0)
	}

	cfg, err := os.Open(cfgFileName)
	if err != nil {
		fmt.Printf("Error opening %v, %v\n", cfgFileName, err)
		return
	}

	appSettings, err := quickfix.ParseSettings(cfg)
	if err != nil {
		fmt.Println("Error reading cfg,", err)
		return
	}

	app := TradeClient{}
	fileLogFactory, err := quickfix.NewFileLogFactory(appSettings)

	if err != nil {
		fmt.Println("Error creating file log factory,", err)
		return
	}

	initiator, err := quickfix.NewInitiator(app, quickfix.NewMemoryStoreFactory(), appSettings, fileLogFactory)
	if err != nil {
		fmt.Printf("Unable to create Initiator: %s\n", err)
		return
	}

	h := heartbeat.New()

	initiator.Start()
	quickfix.Send(h)

	//test:= orderstatusrequest.New()
	SenderCompID, _ := appSettings.GlobalSettings().Setting("SenderCompID")
	TargetCompID, _ := appSettings.GlobalSettings().Setting("TargetCompID")
	internal.SenderId = SenderCompID
	internal.TargetId = TargetCompID
	showMenu()

	initiator.Stop()
}


func showMenu(){
Loop:
	for {
		action, err := internal.QueryAction()
		if err != nil {
			break
		}

		switch action {
		case "1":
			err = internal.QueryEnterOrder()

		case "2":
			err = internal.QueryCancelOrder()

		case "3":
			err = internal.QueryStatusOrder()

		case "4":
			//quit
			break Loop

		default:
			err = fmt.Errorf("unknown action: '%v'", action)
		}

		if err != nil {
			fmt.Printf("%v\n", err)
		}
	}
}