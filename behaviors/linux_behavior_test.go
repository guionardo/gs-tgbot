package behaviors

import (
	"testing"
)

func TestLinuxBehavior_Command(t *testing.T) {
	linuxBehavior:=LinuxBehavior{}
	result:=linuxBehavior.Command()
	if len(result)==0{
		t.Error("Erro em LinuxBehavior")
	} else{
		t.Logf("LinuxBehavior: %s",result)
	}
}
