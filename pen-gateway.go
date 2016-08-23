package gateway

import (
	"github.com/tarm/serial"
	"log"
)

const (
	PN532_START_CODE_1        string = "\x00"
	PN532_START_CODE_2        string = "\xFF"
	PN532_START_CODE          string = PN532_START_CODE_1 + PN532_START_CODE_2
	ACK_FRAME                 string = PN532_START_CODE_1 + PN532_START_CODE_2 + PN532_START_CODE_1 + PN532_START_CODE_2
	ERROR_FRAME               string = PN532_START_CODE_1 + PN532_START_CODE_2 + "\x01" + PN532_START_CODE_2
	END_OF_TEXT               string = "\x03"
	HOST_TO_GATEWAY           string = "\xD4"
	GATEWAY_TO_HOST           string = "\xD5"
	P2P_RESPONSE              string = "\xC3"
	P2P_ETX_END_OF_PEN_RECORD string = "\x03"
	REQUEST_GATEWAY_ID        string = "\xA6"
	REQUEST_MEMORY_FRAMES     string = "\xC2"
	SET_VISIT_ID_TAG          string = "\xC0"
	TASK_ASSOCIATE            string = "\x00"
	ACK_INTERRUPT             string = "\x55\x55\x00\x00\x00\x00\x00\x00\x00\x00\xFF\x00\xFF\x00\x00\x00\xFF\x00\xFF\x00"
	P2P_STATUS_SUCCESS        string = "\x00"
)

func GenerateFrame(direction string, command string) []byte {

	frame := []byte(PN532_START_CODE_1 + PN532_START_CODE_2)

	direction_length := 1
	command_length := 1
	length := direction_length + command_length
	length_checksum := 0xFF & (0x00 - length)
	field_checksum := 0xFF & (0x00 - []byte(direction)[0] - []byte(command)[0])

	frame = append(frame, byte(length), byte(length_checksum), []byte(direction)[0], []byte(command)[0], byte(field_checksum))

	return frame

}

func SendFrame(frame []byte, port string) []byte {

	c := &serial.Config{Name: port, Baud: 9600}
	s, err := serial.OpenPort(c)

	if err != nil {
		log.Fatal(err)
	}

	n, err := s.Write(frame)
	if err != nil {
		log.Fatal(err)
	}

	data := ""
	sum := 1

	for sum < 4 {

		buf := make([]byte, 128)
		n, err = s.Read(buf)

		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Reading Response: %x\n", buf[:n])

		if buf[2] == 0x01 {
			log.Fatalf("Received error frame: %x\n", buf[1])
		}

		data = string(buf[:n])
		sum += sum

	}

	log.Printf("Response: %x\n", data)
	return []byte(data)

}

func DecodeD5Response(frame []byte) string {
	return string(frame[6:])
}
