package main

import (
	"crypto/ed25519"

	"bufio"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/threefoldtech/zos/pkg/gridtypes"
	"github.com/threefoldtech/zos/pkg/gridtypes/zos"
	"os"
)

func main() {
	// size: 10485760
	// zmount := zos.ZMount{
	// 	Size: 1024 * 1024 * 10,
	// }
	fmt.Println(zos.ZMountType)
	workload := gridtypes.Workload{
		Version:     3,
		Name:        "zmountiaia",
		Type:        zos.ZMountType,
		Data:        json.RawMessage(`{"size": 10485760}`),
		Metadata:    "zm",
		Description: "zm test",
	}

	workloads := []gridtypes.Workload{workload}

	deployment := gridtypes.Deployment{
		Version:     3,
		TwinID:      17,
		Expiration:  1626394539, //time.now().unix_time() + 11111,
		Metadata:    "zm dep",
		Description: "zm test",
		Workloads:   workloads,
	}

	w := bufio.NewWriter(os.Stdout)
	fmt.Fprintf(w, "%s", "\"")
	deployment.Challenge(w)
	fmt.Fprintf(w, "%s", "\"\n")
	w.Flush()

	fmt.Println(deployment.ChallengeHash())
	challengeHash, _ := deployment.ChallengeHash()
	hexHash := hex.EncodeToString(challengeHash)
	fmt.Println(hexHash)

	deployment.ContractID = 14
	seed := []byte{171, 101, 228, 213, 178, 56, 187, 250, 175, 19, 223, 79, 12, 92, 149, 56, 221, 186, 188, 41, 119, 82, 88, 84, 191, 11, 119, 28, 6, 131, 8, 40}
	privKey := ed25519.NewKeyFromSeed(seed)
	fmt.Println(privKey.Public())
	deployment.Sign(17, privKey)
	fmt.Println(deployment)

	// test sign
	msg := []byte{1, 2, 3}
	bytes := ed25519.Sign(privKey, msg)
	fmt.Printf("signature bytes: %b", bytes)
	s := hex.EncodeToString(bytes)
	fmt.Printf("signature: %s", s)

}
