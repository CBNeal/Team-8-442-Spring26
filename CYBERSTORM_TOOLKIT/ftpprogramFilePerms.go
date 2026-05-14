package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
)

const (
	addr = "localhost"
	port = "21"
	dir  = "files/10" //files/7 or files/10

	METHOD = "10bit" // For 7 bit, comment out this line
	//METHOD = "7bit"   // and uncomment this line

	user  = "anonymous"
	passw = "anon"
)

func main() {
	// Connect to FTP control channel
	conn, err := net.Dial("tcp", addr+":"+port)
	if err != nil {
		return
	}
	defer conn.Close()

	reader := bufio.NewReader(conn)
	readResponse(reader)

	// Login
	fmt.Fprintf(conn, "USER %s\r\n", user)
	readResponse(reader)
	fmt.Fprintf(conn, "PASS %s\r\n", passw)
	readResponse(reader)

	// Change directory
	fmt.Fprintf(conn, "CWD %s\r\n", dir)
	readResponse(reader)

	// Enter Passive Mode to open data connection
	fmt.Fprintf(conn, "PASV\r\n")
	pasvResp := readResponse(reader)

	// Parse PASV response to get data connection IP and Port
	start := strings.Index(pasvResp, "(")
	end := strings.Index(pasvResp, ")")
	if start == -1 || end == -1 {
		return
	}
	parts := strings.Split(pasvResp[start+1:end], ",")
	if len(parts) != 6 {
		return
	}

	p1, _ := strconv.Atoi(parts[4])
	p2, _ := strconv.Atoi(parts[5])
	dataPort := (p1 * 256) + p2
	dataAddr := fmt.Sprintf("%s.%s.%s.%s:%d", parts[0], parts[1], parts[2], parts[3], dataPort)

	// Connect to data channel
	dataConn, err := net.Dial("tcp", dataAddr)
	if err != nil {
		return
	}

	// Request directory listing
	fmt.Fprintf(conn, "LIST\r\n")
	readResponse(reader) // 150 Here comes the directory listing

	// Read files from data channel
	var lines []string
	dataScanner := bufio.NewScanner(dataConn)
	for dataScanner.Scan() {
		lines = append(lines, dataScanner.Text())
	}
	dataConn.Close()

	// Wait for transfer complete on control channel
	readResponse(reader) // 226 Directory send OK

	// Quit
	fmt.Fprintf(conn, "QUIT\r\n")
	readResponse(reader)
	conn.Close()

	decode(lines)
}

// Helper function to read multi-line FTP responses
func readResponse(reader *bufio.Reader) string {
	var response string
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		response = line
		// Standard FTP response ends with 3 digits and a space
		if len(line) >= 4 && line[3] == ' ' {
			break
		}
	}
	return response
}

func decode(lines []string) {
	bits := ""
	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) == 0 {
			continue
		}

		perm := fields[0]
		// A standard permission string has at least 10 chars
		if len(perm) < 10 {
			continue
		}

		if METHOD == "7bit" {
			// Filter out any file with any of the first three bits set
			if perm[0] != '-' || perm[1] != '-' || perm[2] != '-' {
				continue
			}
			// Use the 7 right-most bits
			bits += perm[3:10]
		} else if METHOD == "10bit" {
			// Utilize all ten bits
			bits += perm[:10]
		}
	}

	fmt.Println(binaryToText(bits))
}

func binaryToText(input string) string {
	clean := ""
	// Convert permission characters to binary string
	for _, c := range input {
		if c == '-' {
			clean += "0"
		} else {
			clean += "1" // Any set permission (d, r, w, x) becomes 1
		}
	}

	var out string
	// Parse string 7 bits at a time to form ASCII characters
	for i := 0; i+7 <= len(clean); i += 7 {
		byteData := clean[i : i+7]
		val := 0
		for _, b := range byteData {
			val <<= 1
			if b == '1' {
				val |= 1
			}
		}
		out += string(rune(val))
	}
	return out
}
