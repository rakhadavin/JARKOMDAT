package main

import "fmt"

type Tool interface {
	AddTransaction(data any)
	GetData() any
}

type Transaction struct {
	IdMember *string
	SKU      string
	Qty      int32
	Price    int32
}

type Member struct {
	IdMember     string
	MemberName   string
	Transactions []Transaction
}

// Tes implements Tool.

func (m *Member) AddTransaction(data any) {

	// Cek jika data adalah tipe Transaction
	if transaction, ok := data.(Transaction); ok {
		// Tambahkan transaksi ke slice Transactions
		m.Transactions = append(m.Transactions, transaction)
	} else if transactions, ok := data.([]Transaction); ok {
		// Jika data adalah slice dari Transaction, tambahkan semua transaksi
		m.Transactions = append(m.Transactions, transactions...)
	} else {
		fmt.Println("Error: Data is not of type Transaction or []Transaction")
	}
}

func (m *Member) GetData() any {
	return map[string]any{
		"IdMember":     m.IdMember,
		"MemberName":   m.MemberName,
		"Transactions": m.Transactions,
	}
}

type Item struct {
	SKU          string
	ItemName     string
	StockQty     int32
	Transactions []Transaction
	Price        int32
}

func (it *Item) AddTransaction(data any) {

	// Cek jika data adalah tipe Transaction
	if transaction, ok := data.(Transaction); ok {
		// Tambahkan transaksi ke slice Transactions
		it.Transactions = append(it.Transactions, transaction)
		it.StockQty -= transaction.Qty
	} else if transactions, ok := data.([]Transaction); ok {
		// Jika data adalah slice dari Transaction, tambahkan semua transaksi
		it.Transactions = append(it.Transactions, transactions...)
	} else {
		fmt.Println("Error: Data is not of type Transaction or []Transaction")
	}
}

func (it *Item) GetData() any {
	return map[string]any{
		"SKU":          it.SKU,
		"ItemName":     it.ItemName,
		"StockQty":     it.StockQty,
		"Transactions": it.Transactions,
		"Price":        it.Price,
	}
}
