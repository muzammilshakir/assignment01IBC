package assignment01IBC

import (
	"crypto/sha256"
	"fmt"
)

type BlockData struct {
	Transactions []string
}
type Block struct {
	Data        BlockData
	PrevPointer *Block
	PrevHash    string
	CurrentHash string
}

func CalculateHash(inputBlock *Block) string {
	data := fmt.Sprintf("%v", inputBlock.Data.Transactions)
	calHash := fmt.Sprintf("%x\n", sha256.Sum256([]byte(data)))
	return calHash
}
func InsertBlock(dataToInsert BlockData, chainHead *Block) *Block {
	if chainHead == nil {
		temp := &Block{
			Data:        dataToInsert,
			PrevPointer: nil,
			CurrentHash: "",
			PrevHash:    "",
		}
		currentHash := CalculateHash(temp)
		chainHead = &Block{
			Data:        dataToInsert,
			PrevPointer: nil,
			CurrentHash: currentHash,
			PrevHash:    "",
		}

	} else {
		temp := &Block{
			Data:        dataToInsert,
			PrevPointer: chainHead,
			CurrentHash: "",
			PrevHash:    chainHead.CurrentHash,
		}
		currentHash := CalculateHash(temp)
		chainHead = &Block{
			Data:        dataToInsert,
			PrevPointer: chainHead,
			CurrentHash: currentHash,
			PrevHash:    chainHead.CurrentHash,
		}
	}
	return chainHead
}
func ChangeBlock(oldTrans string, newTrans string, chainHead *Block) {
	tempHead := chainHead
	for tempHead.PrevPointer != nil {
		for index, element := range tempHead.Data.Transactions {
			if oldTrans == element {
				tempHead.Data.Transactions[index] = newTrans
			}
		}
		tempHead = tempHead.PrevPointer
	}
	for index, element := range tempHead.Data.Transactions {
		if oldTrans == element {
			tempHead.Data.Transactions[index] = newTrans
		}
	}
}
func ListBlocks(chainHead *Block) {
	tempHead := chainHead
	for tempHead.PrevPointer != nil {
		fmt.Println(tempHead.Data)
		tempHead = tempHead.PrevPointer
	}
	fmt.Println(tempHead.Data)
}
func VerifyChain(chainHead *Block) {
	tempHead := chainHead
	for tempHead.PrevPointer != nil {
		VerHash := CalculateHash(tempHead)
		if tempHead.CurrentHash != VerHash {
			fmt.Println("Block Chain Crompromised")
			return
		}
		tempHead = tempHead.PrevPointer
	}
	VerHash := CalculateHash(tempHead)
	if tempHead.CurrentHash != VerHash {
		fmt.Println("Block Chain Crompromised")
		return
	} else {
		fmt.Println("Block Chian OK")
	}

}
