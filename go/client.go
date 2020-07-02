// client.go
package main

import (
    "fmt"
    "net"
    "io"
    "errors"
    "encoding/binary"
    "strings"
    "os"
    "bufio"
)

const MAX_DATA_LEN = 10*1024

const (
	CMD_EXIT       = uint16(0x0000)
	CMD_CHECK      = uint16(0x0001)
	CMD_VERSION    = uint16(0x0101)
	CMD_UPDATE     = uint16(0x0102)
	CMD_RECONF     = uint16(0x0103)
	CMD_RELOAD_LIC = uint16(0x0104)
)

const (
	RESULT_SUCCESS      = uint16(0x0000)
	RESULT_FAILURE      = uint16(0x0101)
	RESULT_IOC_LOW      = uint16(0x0201)
	RESULT_IOC_MEDIUM   = uint16(0x0202)
	RESULT_IOC_HIGH     = uint16(0x0203)
	RESULT_IOC_CRITICAL = uint16(0x0204)
)

type Request struct {
	id      uint16
	command uint16
	dataLen uint32
	data    string
}

type Response struct {
	id      uint16
	result  uint16
	dataLen uint32
	data    string
}

func BytesToUint16(bs []byte) uint16 {
	return uint16(binary.BigEndian.Uint16(bs))
}

func BytesToUint32(bs []byte) uint32 {
	return uint32(binary.BigEndian.Uint32(bs))
}

func Uint16ToBytes(i uint16, bs []byte) error {
	if len(bs) != 2 {
		return errors.New(fmt.Sprintf("[]byte len error. [ len=%d ]", len(bs)))
	}
	binary.BigEndian.PutUint16(bs, i)
	return nil
}

func Uint32ToBytes(i uint32, bs []byte) error {
	if len(bs) != 4 {
		return errors.New(fmt.Sprintf("[]byte len error. [ len=%d ]", len(bs)))
	}
	binary.BigEndian.PutUint32(bs, i)
	return nil
}

func sendRequest(w io.Writer, id uint16, command uint16, data string) error {
	buf := make([]byte, 8)

	dataLen := uint32(len(data))
	Uint16ToBytes(id, buf[0:2])
	Uint16ToBytes(command, buf[2:4])
	Uint32ToBytes(dataLen, buf[4:8])

	n, err := w.Write([]byte(buf))
	if err != nil {
		return err
	}
	if n != 8 {
		err = errors.New(fmt.Sprintf("write error [%d], want write [%d].", n, 8))
		return err
	}
	if dataLen > 0 {
		if _, err := w.Write([]byte(data)); err != nil {
			return err
		}
	}
	return nil
}

func recvResponse(r io.Reader) (*Response, error) {
	var resp Response
	buf := make([]byte, MAX_DATA_LEN)

	_, err := io.ReadFull(r, buf[0:8])
	if err != nil {
		if err != io.EOF {
		}
		return nil, err
	}

	resp.id      = BytesToUint16(buf[0:2])
	resp.result  = BytesToUint16(buf[2:4])
	resp.dataLen = BytesToUint32(buf[4:8])

	if resp.dataLen < 0 || resp.dataLen > MAX_DATA_LEN {
		err = errors.New(fmt.Sprintf("request body len error (dataLen=%d) (MAX_DATA_LEN=%d).",
		                              resp.dataLen, MAX_DATA_LEN))
		return nil, err
	}

	if resp.dataLen != 0 {
		_, err = io.ReadFull(r, buf[0:resp.dataLen])
		if err != nil {
			return nil, err
		}
		resp.data = string(buf)
	}
	return &resp, nil
}

func main() {
	c, err := net.Dial("unix", "/tmp/iocengine.sock")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer c.Close()

	requestID := uint16(0)
	inputReader := bufio.NewReader(os.Stdin)
	for {
		command := CMD_CHECK
		data    := ""
		fmt.Printf("iocengine client > ")
		inputStr, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		strLen := len(inputStr)
		if strLen <= 0 {
			continue
		}else if inputStr[strLen-1] == '\n' {
			inputStr = inputStr[:strLen-1]
		}
		substr := strings.Split(inputStr, " ")
		switch substr[0] {
		case "check":
			command = CMD_CHECK
		case "reconfig":
			command = CMD_RECONF
		case "reloadLic":
			command = CMD_RELOAD_LIC
		case "update":
			command = CMD_UPDATE
		case "ver":
			command = CMD_VERSION
		case "quit":
			command = CMD_EXIT
		case "exit":
			command = CMD_EXIT
		default:
			fmt.Println("Unknown command.")
			continue
		}
		if command == CMD_EXIT {
			break
		}
		if len(substr) >= 2 {
			data = substr[1]
		}
		requestID++
		if err = sendRequest( c, requestID, command, data ); err != nil {
			fmt.Println(err.Error())
			break;
		}
		var resp *Response
		if resp, err = recvResponse(c); err != nil {
			fmt.Println(err.Error())
			break
		}
		fmt.Printf("id=%d, result=%0x data=%s\n", resp.id, resp.result, resp.data)
	} 
}
