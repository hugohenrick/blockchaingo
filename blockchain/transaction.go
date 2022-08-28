package blockchain

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Transaction struct {
	senderBlockchainAddres     string
	recipientBlockchainAddress string
	value                      float64
}

func NewTranascation(sender, recipient string, value float64) *Transaction {
	return &Transaction{sender, recipient, value}
}

func (t *Transaction) Print() {
	fmt.Printf("%s\n", strings.Repeat("-", 40))
	fmt.Printf(" sender_blockchain_address %s\n", t.senderBlockchainAddres)
	fmt.Printf(" recipient_blockchain_address %s\n", t.recipientBlockchainAddress)
	fmt.Printf(" value %1f\n", t.value)
}

func (t *Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		SenderBlockchainAddres     string  `json:"sender_blockchain_address"`
		RecipientBlockchainAddress string  `json:"recipient_blockchain_address"`
		Value                      float64 `json:"value"`
	}{
		SenderBlockchainAddres:     t.senderBlockchainAddres,
		RecipientBlockchainAddress: t.recipientBlockchainAddress,
		Value:                      t.value,
	})
}
