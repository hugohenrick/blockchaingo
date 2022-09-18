package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
	"fmt"

	"github.com/btcsuite/btcutil/base58"
	"github.com/hugohenrick/blockchaingo/utils"
	"golang.org/x/crypto/ripemd160"
)

type Wallet struct {
	privatekey        *ecdsa.PrivateKey
	publicKey         *ecdsa.PublicKey
	blockChainAddress string
}

type Transaction struct {
	senderPrivateKey           *ecdsa.PrivateKey
	senderPublicKey            *ecdsa.PublicKey
	senderBlockChainAddress    string
	recipientBlockChainAddress string
	value                      float64
}

func NewWallet() *Wallet {
	// Creating ECDSA private key(32 bytes) public key(64 bytes)
	w := new(Wallet)
	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	w.privatekey = privateKey
	w.publicKey = &w.privatekey.PublicKey

	// Perform SHA256 hashing on the buplic key(32 bytes)
	h2 := sha256.New()
	h2.Write(w.publicKey.X.Bytes())
	h2.Write(w.publicKey.Y.Bytes())
	digest2 := h2.Sum(nil)

	// Perform Ripemd-160 hashing on the result of SHA-256(20 bytes)
	h3 := ripemd160.New()
	h3.Write(digest2)
	digest3 := h3.Sum(nil)

	// Add version byte in front of Ripemd-160 hash(0x00 for Main Network)
	vd4 := make([]byte, 21)
	vd4[0] = 0x00
	copy(vd4[1:], digest3[:])

	// Perform SHA-256 hash on the extended Ripemd-160 result
	h5 := sha256.New()
	h5.Write(vd4)
	digest5 := h5.Sum(nil)

	// Perform SHA-256 hash on the result of the previous SHA-256 hash
	h6 := sha256.New()
	h6.Write(digest5)
	digest6 := h6.Sum(nil)

	// Take the first 4 bytes of the second SHA-256 hash for checksum
	chsum := digest6[:4]

	// Add the 4 checksum bytes from 7 at the end of extended Ripemd-160 hash from 4(25 bytes)
	dc8 := make([]byte, 25)
	copy(dc8[:21], vd4[:])
	copy(dc8[21:], chsum[:])

	// Convert the result from a byte string into base58
	addres := base58.Encode(dc8)
	w.blockChainAddress = addres
	return w
}

func (w *Wallet) PrivateKey() *ecdsa.PrivateKey {
	return w.privatekey
}

func (w *Wallet) PrivateKeyStr() string {
	return fmt.Sprintf("%x", w.privatekey.D.Bytes())
}

func (w *Wallet) PublicKey() *ecdsa.PublicKey {
	return w.publicKey
}

func (w *Wallet) PublicKeyStr() string {
	return fmt.Sprintf("%x%x", w.publicKey.X.Bytes(), w.privatekey.Y.Bytes())
}

func (w *Wallet) BlockChainAddress() string {
	return w.blockChainAddress
}

func NewTranascation(privateKey *ecdsa.PrivateKey, publicKey *ecdsa.PublicKey,
	sender string, recipient string, value float64) *Transaction {

	return &Transaction{privateKey, publicKey, sender, recipient, value}
}

func (t *Transaction) GenerateSiginature() *utils.Signature {
	m, _ := json.Marshal(t)
	h := sha256.Sum256([]byte(m))
	r, s, _ := ecdsa.Sign(rand.Reader, t.senderPrivateKey, h[:])
	return &utils.Signature{R: r, S: s}
}

func (t *Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Sender    string  `json:"sender_blockchain_address"`
		Recipient string  `json:"recipient_blockchain_address"`
		Value     float64 `json:"value"`
	}{
		Sender:    t.senderBlockChainAddress,
		Recipient: t.recipientBlockChainAddress,
		Value:     t.value,
	})
}
