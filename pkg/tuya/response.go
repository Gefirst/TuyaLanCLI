package tuya

import "github.com/Gefest/TuyaLanCLI/internal/commands"

type response struct {
	payload
	returnCode    int
	newSequenceNr int
	commandByte   commands.Type
}
