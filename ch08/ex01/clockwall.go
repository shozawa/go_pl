package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/exec"
	"time"
)

type timeScanenr struct {
	scanner  *bufio.Scanner
	reader   io.Reader
	port     int
	TimeZone string
}

func New(timeZone string, port int) *timeScanenr {
	scanenr := &timeScanenr{TimeZone: timeZone, port: port}
	conn, err := net.Dial("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatal(err)
	}
	scanenr.reader = conn
	scanenr.scanner = bufio.NewScanner(conn)
	return scanenr
}

func (s *timeScanenr) Scan() bool {
	return s.scanner.Scan()
}

func (s *timeScanenr) Text() string {
	return s.scanner.Text()
}

func main() {
	scanners := []*timeScanenr{
		New("US/Eastern", 8010),
		New("Asia/Tokyo", 8020),
		New("Europe/Londl", 8030),
	}
	for {
		clear()
		for _, scanner := range scanners {
			scanner.Scan()
			fmt.Printf("%s\t%s\n", scanner.TimeZone, scanner.Text())
		}
		time.Sleep(1 * time.Second)
	}
}

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
