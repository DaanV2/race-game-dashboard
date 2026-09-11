/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/base64"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	udp "github.com/daanv2/race-game-dashboard/infrastructure/transport/upd"
	"github.com/daanv2/race-game-dashboard/pkg/extensions/xsync"
	"github.com/daanv2/race-game-dashboard/pkg/f1game/f12026s"
	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "A brief description of your command",
	Long:  `A longer description`,
	RunE:  runServer,
}

func init() {
	rootCmd.AddCommand(serverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	serverCmd.Flags().StringP("output-file", "o", "", "A file that will store all the packets, as base64, message per line")
}

func runServer(cmd *cobra.Command, args []string) error {
	ctx, stop := signal.NotifyContext(cmd.Context(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	hand := &f12026s.PacketHandler{}

	hand.CarDamage.Register(logMsg)
	hand.CarSetups.Register(logMsg)
	hand.CarStatus.Register(logMsg)
	hand.CarTelemetry.Register(logMsg)
	hand.CarTelemetry2.Register(logMsg)
	hand.Event.Register(logMsg)
	hand.FinalClassification.Register(logMsg)
	hand.LapData.Register(logMsg)
	hand.LapPositions.Register(logMsg)
	hand.LobbyInfo.Register(logMsg)
	hand.Motion.Register(logMsg)
	hand.MotionEx.Register(logMsg)
	hand.Participants.Register(logMsg)
	hand.Session.Register(logMsg)
	hand.SessionHistory.Register(logMsg)
	hand.TimeTrial.Register(logMsg)
	hand.TyreSets.Register(logMsg)

	handlefn := hand.Ingest

	outputfile, _ := cmd.Flags().GetString("output-file")
	if outputfile != "" {
		outputfile, err := filepath.Abs(outputfile)
		if err != nil {
			return fmt.Errorf("couldn't make path absolute: %w", err)
		}

		outputfile = filepath.Clean(outputfile)
		w, err := os.Create(outputfile)
		if err != nil {
			return fmt.Errorf("couldn't open or create file: \nfile: %s\nerror: %w", outputfile, err)
		}

		pcks := xsync.NewMap[f12026s.PacketID, string]()

		defer writeOutputfileAndClose(w, pcks)

		oldHandlefn := handlefn
		handlefn = func(data []byte) {
			p := hand.ParsePacket(data)
			if p == nil {
				return
			}
			id := p.GetPacketID()
			pcks.Set(id, base64.RawStdEncoding.EncodeToString(data))

			oldHandlefn(data)
		}
	}

	udphandle := udp.HandleFunc(handlefn)

	const PORT = 20777
	server, err := udp.NewServer(PORT, udphandle)
	if err != nil {
		return fmt.Errorf("error with setting up server: %w", err)
	}

	go func() {
		fmt.Println("listing on: ", PORT)
		err := server.Listen()
		if err != nil {
			fmt.Println("Error during shutdown", err)
		}
	}()

	<-ctx.Done()

	err = server.Close()

	return err
}

func writeOutputfileAndClose(w *os.File, pcks *xsync.Map[f12026s.PacketID, string]) {
	for _, p := range pcks.Values() {
		_, err := w.WriteString(p + "\n")
		if err != nil {
			fmt.Println("error writing to file", err)
		}
	}

	cerr := w.Close()
	if cerr != nil {
		fmt.Println("error closing file", cerr)
	}
}

func logMsg[T any](data T) {
	fmt.Printf("%v\n", data)
}
