package api

import (
	"math/big"
	"encoding/json"
	"strconv"
	"github.com/ranchimall/blockbook/bchain"
)

// ScriptSigV1 is used for legacy api v1
type ScriptSigV1 struct {
	Hex string `json:"hex,omitempty"`
	Asm string `json:"asm,omitempty"`
}

// VinV1 is used for legacy api v1
type VinV1 struct {
	Txid      string                   `json:"txid"`
	Vout      uint32                   `json:"vout"`
	Sequence  int64                    `json:"sequence,omitempty"`
	N         int                      `json:"n"`
	ScriptSig ScriptSigV1              `json:"scriptSig"`
	AddrDesc  bchain.AddressDescriptor `json:"-"`
	Addresses []string                 `json:"addresses"`
	IsAddress bool                     `json:"isAddress"`
	Value     float64                  `json:"value"`
	ValueSat  big.Int                  `json:"valueSat"`
	Coinbase  string                   `json:"coinbase,omitempty"`
}

// ScriptPubKeyV1 is used for legacy api v1
type ScriptPubKeyV1 struct {
	Hex       string                   `json:"hex,omitempty"`
	Asm       string                   `json:"asm,omitempty"`
	AddrDesc  bchain.AddressDescriptor `json:"-"`
	Addresses []string                 `json:"addresses"`
	IsAddress bool                     `json:"isAddress"`
	Type      string                   `json:"type,omitempty"`
}

// VoutV1 is used for legacy api v1
type VoutV1 struct {
	Value        float64        `json:"value"`
	ValueSat     big.Int        `json:"valueSat"`
	N            int            `json:"n"`
	ScriptPubKey ScriptPubKeyV1 `json:"scriptPubKey"`
	Spent        bool           `json:"spent"`
	SpentTxID    string         `json:"spentTxId,omitempty"`
	SpentIndex   int            `json:"spentIndex,omitempty"`
	SpentHeight  int            `json:"spentHeight,omitempty"`
}

// TxV1 is used for legacy api v1
type TxV1 struct {
	Txid          string   `json:"txid"`
	Version       int32    `json:"version,omitempty"`
	Locktime      uint32   `json:"locktime"`
	Vin           []VinV1  `json:"vin"`
	Vout          []VoutV1 `json:"vout"`
	Blockhash     string   `json:"blockhash,omitempty"`
	Blockheight   int      `json:"blockheight"`
	Confirmations uint32   `json:"confirmations"`
	Time          int64    `json:"time,omitempty"`
	Blocktime     int64    `json:"blocktime"`
	ValueOut      float64  `json:"valueOut"`
	ValueOutSat   big.Int  `json:"valueOutSat"`
	Size          int      `json:"size,omitempty"`
	ValueIn       float64  `json:"valueIn"`
	ValueInSat    big.Int  `json:"valueInSat"`
	Fees          float64  `json:"fees"`
	FeesSat       big.Int  `json:"feesSat"`
	Hex           string   `json:"hex"`
	FloData       string   `json:"floData,omitempty"`
}

// AddressV1 is used for legacy api v1
type AddressV1 struct {
	Paging
	AddrStr                 string   `json:"addrStr"`
	Balance                 float64  `json:"balance"`
	BalanceSat              big.Int  `json:"balanceSat"`
	TotalReceived           float64  `json:"totalReceived"`
	TotalReceivedSat        big.Int  `json:"totalReceivedSat"`
	TotalSent               float64  `json:"totalSent"`
	TotalSentSat            big.Int  `json:"totalSentSat"`
	UnconfirmedBalance      float64  `json:"unconfirmedBalance"`
	UnconfirmedBalanceSat   big.Int  `json:"unconfirmedBalanceSat"`
	UnconfirmedTxApperances int      `json:"unconfirmedTxApperances"`
	TxApperances            int      `json:"txApperances"`
	Transactions            []*TxV1  `json:"txs,omitempty"`
	Txids                   []string `json:"transactions,omitempty"`
}

// AddressUtxoV1 is used for legacy api v1
type AddressUtxoV1 struct {
	Txid          string  `json:"txid"`
	Vout          uint32  `json:"vout"`
	Amount        float64 `json:"amount"`
	AmountSat     big.Int `json:"satoshis"`
	Height        int     `json:"height,omitempty"`
	Confirmations int     `json:"confirmations"`
}

// BlockV1 contains information about block
type BlockV1 struct {
	Paging
	BlockInfo
	TxCount      int     `json:"txCount"`
	Transactions []*TxV1 `json:"txs,omitempty"`
}

type CoinSpecificDataV0 struct {
	FloData       string
}

func stringToFloat(f string) float64 {
	s, _ := strconv.ParseFloat(f, 64)
	return s
}

// TxToV1 converts Tx to TxV1
func (w *Worker) TxToV1(tx *Tx) *TxV1 {
	d := w.chainParser.AmountDecimals()
	vinV1 := make([]VinV1, len(tx.Vin))
	for i := range tx.Vin {
		v := &tx.Vin[i]
		vinV1[i] = VinV1{
			AddrDesc:  v.AddrDesc,
			Addresses: v.Addresses,
			N:         v.N,
			ScriptSig: ScriptSigV1{
				Asm: v.Asm,
				Hex: v.Hex,
			},
			IsAddress: v.IsAddress,
			Sequence:  v.Sequence,
			Txid:      v.Txid,
			Value:     stringToFloat(v.ValueSat.DecimalString(d)),
			ValueSat:  v.ValueSat.AsBigInt(),
			Vout:      v.Vout,
			Coinbase:  v.Coinbase,
		}
	}
	voutV1 := make([]VoutV1, len(tx.Vout))
	for i := range tx.Vout {
		v := &tx.Vout[i]
		voutV1[i] = VoutV1{
			N: v.N,
			ScriptPubKey: ScriptPubKeyV1{
				AddrDesc:  v.AddrDesc,
				Addresses: v.Addresses,
				Asm:       v.Asm,
				Hex:       v.Hex,
				IsAddress: v.IsAddress,
				Type:      v.Type,
			},
			Spent:       v.Spent,
			SpentHeight: v.SpentHeight,
			SpentIndex:  v.SpentIndex,
			SpentTxID:   v.SpentTxID,
			Value:       stringToFloat(v.ValueSat.DecimalString(d)),
			ValueSat:    v.ValueSat.AsBigInt(),
		}
	}
	
	//floData
	var coinData CoinSpecificDataV0
    // Notice the dereferencing asterisk *
    err := json.Unmarshal(tx.CoinSpecificData, &coinData)
    if err != nil {
        return nil
    }

	return &TxV1{
		Blockhash:     tx.Blockhash,
		Blockheight:   tx.Blockheight,
		Blocktime:     tx.Blocktime,
		Confirmations: tx.Confirmations,
		Fees:          stringToFloat(tx.FeesSat.DecimalString(d)),
		FeesSat:       tx.FeesSat.AsBigInt(),
		Hex:           tx.Hex,
		Locktime:      tx.Locktime,
		Size:          tx.Size,
		Time:          tx.Blocktime,
		Txid:          tx.Txid,
		ValueIn:       stringToFloat(tx.ValueInSat.DecimalString(d)),
		ValueInSat:    tx.ValueInSat.AsBigInt(),
		ValueOut:      stringToFloat(tx.ValueOutSat.DecimalString(d)),
		ValueOutSat:   tx.ValueOutSat.AsBigInt(),
		Version:       tx.Version,
		Vin:           vinV1,
		Vout:          voutV1,
		FloData:       coinData.FloData,	//(tx.CoinSpecificData).FloData
	}
}

func (w *Worker) transactionsToV1(txs []*Tx) []*TxV1 {
	v1 := make([]*TxV1, len(txs))
	for i := range txs {
		v1[i] = w.TxToV1(txs[i])
	}
	return v1
}

// AddressToV1 converts Address to AddressV1
func (w *Worker) AddressToV1(a *Address) *AddressV1 {
	d := w.chainParser.AmountDecimals()
	return &AddressV1{
		AddrStr:                 a.AddrStr,
		Balance:                 stringToFloat(a.BalanceSat.DecimalString(d)),
		BalanceSat:              a.BalanceSat.AsBigInt(),
		Paging:                  a.Paging,
		TotalReceived:           stringToFloat(a.TotalReceivedSat.DecimalString(d)),
		TotalReceivedSat:         a.TotalReceivedSat.AsBigInt(),
		TotalSent:               stringToFloat(a.TotalSentSat.DecimalString(d)),
		TotalSentSat:            a.TotalSentSat.AsBigInt(),
		Transactions:            w.transactionsToV1(a.Transactions),
		TxApperances:            a.Txs,
		Txids:                   a.Txids,
		UnconfirmedBalance:      stringToFloat(a.UnconfirmedBalanceSat.DecimalString(d)),
		UnconfirmedBalanceSat:   a.UnconfirmedBalanceSat.AsBigInt(),
		UnconfirmedTxApperances: a.UnconfirmedTxs,
	}
}

// AddressUtxoToV1 converts []AddressUtxo to []AddressUtxoV1
func (w *Worker) AddressUtxoToV1(au Utxos) []AddressUtxoV1 {
	d := w.chainParser.AmountDecimals()
	v1 := make([]AddressUtxoV1, len(au))
	for i := range au {
		utxo := &au[i]
		v1[i] = AddressUtxoV1{
			AmountSat:     utxo.AmountSat.AsBigInt(),
			Amount:        stringToFloat(utxo.AmountSat.DecimalString(d)),
			Confirmations: utxo.Confirmations,
			Height:        utxo.Height,
			Txid:          utxo.Txid,
			Vout:          uint32(utxo.Vout),
		}
	}
	return v1
}

// BlockToV1 converts Address to Address1
func (w *Worker) BlockToV1(b *Block) *BlockV1 {
	return &BlockV1{
		BlockInfo:    b.BlockInfo,
		Paging:       b.Paging,
		Transactions: w.transactionsToV1(b.Transactions),
		TxCount:      b.TxCount,
	}
}
