package main

import (
	"context"
	"fmt"
	"github.com/skip2/go-qrcode"
	"log"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/pakaiwa/pakaiwa"
	"github.com/pakaiwa/pakaiwa/store/sqlstore"
	"github.com/pakaiwa/pakaiwa/types/events"
	waLog "github.com/pakaiwa/pakaiwa/util/log"
)

func eventHandler(evt interface{}) {
	switch v := evt.(type) {
	case *events.Message:
		fmt.Println("Pesan asli:", v.Message)
		sender := v.Info.Sender.String()
		switch msg := v.Message.GetConversation(); {
		case msg != "":
			fmt.Printf("Pesan dari %s: %s\n", sender, msg)

		case v.Message.GetExtendedTextMessage() != nil:
			ext := v.Message.GetExtendedTextMessage()
			fmt.Printf("Pesan dari %s: %s\n", sender, ext.GetText())

		default:
			fmt.Println("Tidak dikenal / format lain")
		}

	}
}

func main() {
	dbLog := waLog.Stdout("Database", "DEBUG", true)
	container, err := sqlstore.New("pgx", "postgres://kanggara:susah@kvm2.pakaiwa.my.id:5432/dev", dbLog)
	if err != nil {
		log.Fatal("failed to connect:", err)
	}

	// If you want multiple sessions, remember their JIDs and use .GetDevice(jid) or .GetAllDevices() instead.
	deviceStore, err := container.GetFirstDevice()
	if err != nil {
		panic(err)
	}
	clientLog := waLog.Stdout("Client", "INFO", true)
	client := pakaiwa.NewClient(deviceStore, clientLog)
	client.AddEventHandler(eventHandler)

	if client.Store.ID == nil {
		// No ID stored, new login
		qrChan, _ := client.GetQRChannel(context.Background())
		err = client.Connect()
		if err != nil {
			panic(err)
		}
		for evt := range qrChan {
			if evt.Event == "code" {
				fmt.Println("QR code:", evt.Code)

				qr, err := qrcode.New(evt.Code, qrcode.Medium)
				if err != nil {
					panic(err)
				}

				ascii := qr.ToString(false)
				fmt.Println(ascii)
			} else {
				fmt.Println("Login event:", evt.Event)
			}
		}
	} else {
		// Already logged in, just connect
		err = client.Connect()
		if err != nil {
			panic(err)
		}
	}

	// Listen to Ctrl+C (you can also do something else that prevents the program from exiting)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	client.Disconnect()
}
