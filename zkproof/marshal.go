package zkproof

import (
	"fmt"
	"math/big"
)

// Marshal converts proof to bytes
func (p *Proof) Marshal() []byte {
	// Ensure fixed-size fields are padded to their full length
	cBytes := make([]byte, 32)
	sBytes := make([]byte, 32)

	// Right-pad the big.Int bytes
	copy(cBytes[32-len(p.C.Bytes()):], p.C.Bytes())
	copy(sBytes[32-len(p.S.Bytes()):], p.S.Bytes())

	// Calculate total size: C(32) + S(32) + R1(65) + R2(65) + BuyerPubKey(65) + SellerPubKey(65) + BuyerCipherHash(32)
	totalSize := 32*3 + 65*4

	// Build the complete buffer
	buf := make([]byte, 0, totalSize)
	buf = append(buf, cBytes...)            // 32 bytes
	buf = append(buf, sBytes...)            // 32 bytes
	buf = append(buf, p.R1...)              // 65 bytes
	buf = append(buf, p.R2...)              // 65 bytes
	buf = append(buf, p.BuyerPubKey...)     // 65 bytes
	buf = append(buf, p.SellerPubKey...)    // 65 bytes
	buf = append(buf, p.BuyerCipherHash...) // 32 bytes
	return buf
}

// Unmarshal recreates proof from bytes
func (p *Proof) Unmarshal(data []byte) error {
	// Expected size: C(32) + S(32) + R1(65) + R2(65) + BuyerPubKey(65) + SellerPubKey(65) + BuyerCipherHash(32)
	expectedSize := 32*3 + 65*4
	if len(data) != expectedSize {
		return fmt.Errorf("invalid proof data size: got %d bytes, want %d", len(data), expectedSize)
	}

	offset := 0

	// Read C (32 bytes)
	p.C = new(big.Int).SetBytes(data[offset : offset+32])
	offset += 32

	// Read S (32 bytes)
	p.S = new(big.Int).SetBytes(data[offset : offset+32])
	offset += 32

	// Read R1 (65 bytes)
	p.R1 = make([]byte, 65)
	copy(p.R1, data[offset:offset+65])
	offset += 65

	// Read R2 (65 bytes)
	p.R2 = make([]byte, 65)
	copy(p.R2, data[offset:offset+65])
	offset += 65

	// Read BuyerPubKey (65 bytes)
	p.BuyerPubKey = make([]byte, 65)
	copy(p.BuyerPubKey, data[offset:offset+65])
	offset += 65

	// Read SellerPubKey (65 bytes)
	p.SellerPubKey = make([]byte, 65)
	copy(p.SellerPubKey, data[offset:offset+65])
	offset += 65

	// Read BuyerCipherHash (32 bytes)
	p.BuyerCipherHash = make([]byte, 32)
	copy(p.BuyerCipherHash, data[offset:offset+32])

	return nil
}
