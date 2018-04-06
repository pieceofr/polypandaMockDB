package main

import (
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"

	"github.com/ethereum/go-ethereum/common/number"
	haikunator "github.com/yelinaung/go-haikunator"
)

const genZeroInEveryX = 10
const userNamePrefix = "user"

/*EthAddr ethereum address which is 20bytes*/
type EthAddr [20]byte

/*Panda struct*/
type Panda struct {
	PandaIndex uint32
	Genes      *number.Number
	Birthtime  uint64
	Cooldown   uint64
	Rank       uint32
	MotherID   uint32
	FatherID   uint32
	Generation uint16
	Owner      EthAddr
	Ownername  string
	Photourl   string
	Snapurl    string
}

/*CreatePanda create a Panda by assigned start number*/
func CreatePanda(startIdx uint32) Panda {
	rand.Seed(time.Now().UTC().UnixNano())
	now := uint64(time.Now().Unix())
	randCoolIndex := uint16(rand.Intn(14))

	//Still needs to complete motherID, fatherID, owner, photourl, snapurl
	panda := Panda{PandaIndex: startIdx, Genes: genGene(), Birthtime: now,
		Cooldown: cooldownArray[randCoolIndex], Generation: genGeneration(),
		Owner: createOwnerAddress(), Ownername: genRandomName()}
	return panda
}

/*const value for bit operation */
const (
	MaxValueInByte         = 1<<7 - 1
	MaxValueInUint64 int64 = 1<<63 - 1
)

func genGene() *number.Number {
	rand.Seed(time.Now().UTC().UnixNano())
	ret := number.Uint256(rand.Int63n(MaxValueInUint64))
	nMult := number.Uint256(rand.Int63n(MaxValueInUint64))
	return ret.Mul(ret, nMult)
}

func genGeneration() uint16 {
	randGenZeron := uint16(rand.Intn(genZeroInEveryX))
	if randGenZeron%genZeroInEveryX == 0 {
		randGenZeron = 0
	}
	return randGenZeron
}

func createOwnerAddress() EthAddr {
	var addr [20]byte
	rand.Seed(time.Now().UTC().UnixNano())
	for idx := range addr {
		addr[idx] = byte(rand.Intn(MaxValueInByte))
	}
	return addr
}
func genRandomName() string {
	haikunator := haikunator.New(time.Now().UTC().UnixNano())
	return fmt.Sprintf("%s", haikunator.Haikunate())
}

/*EncodeNumberToHexString encode number to hex string */
func encodeNumberToHexString(n *number.Number) string {
	geneBytes := n.Bytes()
	return hex.EncodeToString(geneBytes)
}

func ethAddrToHexString(addr EthAddr) string {
	arrayToSlice := addr[:]
	return hex.EncodeToString(arrayToSlice)
}

/*GetString printout panda in json format*/
func (p *Panda) GetString() string {
	print := fmt.Sprintf("{\n PandaIndex:%v\n Genes:%s\n Birthtime:%v\n Cooldown:%v\n Rank:%v\n "+
		"MotherID:%v\n FatherID:%v\n Generation:%v\n Owner:%v\n Ownername:%v\n Photourl:%v\n Snapurl:%v\n}",
		p.PandaIndex, geneToStringInBinary(p.Genes), p.Birthtime, p.Cooldown, p.Rank,
		p.MotherID, p.FatherID, p.Generation, ethAddrToHexString(p.Owner), p.Ownername, p.Photourl, p.Snapurl)
	return print
}

func geneToStringInBinary(gene *number.Number) string {
	binStr := ""
	fmt.Println(len(gene.Bytes()))
	for _, val := range gene.Bytes() {
		binStr = fmt.Sprintf("%s%08b", binStr, val)
	}
	return fmt.Sprintf("%0256s", binStr)
}

/*secs in one minute , one hour and one day*/
const (
	Sec1Min  = 60
	Sec1Hour = Sec1Min * 60
	Sec1Day  = Sec1Hour * 24
)

var cooldownArray = [15]uint64{
	Sec1Min,
	2 * Sec1Min,
	5 * Sec1Min,
	10 * Sec1Min,
	30 * Sec1Min,
	Sec1Hour,
	2 * Sec1Hour,
	4 * Sec1Hour,
	8 * Sec1Hour,
	16 * Sec1Hour,
	Sec1Day,
	2 * Sec1Day,
	4 * Sec1Day,
	7 * Sec1Day,
}
