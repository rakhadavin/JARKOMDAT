package main

import (
	"fmt"
)

var (
	Name      string = "Rakha Davin Bani Alamsyah" // please insert your name here
	IdStudent string = "2206082650"                // please insert your id student here
	Items     []Tool                               // contain array of item pointer
	Members   []Tool                               // contain array of member pointer
)

//	func isExist(SKU string) bool {
//		for _, transaction := range Items {
//			if item, ok := transaction.(*Transaction); ok {
//				if item.SKU == SKU {
//					return "", fmt.Errorf("item %s is already in the list of items", SKU)
//				}
//			}
//		}
//		return false
//	}

func isExistSKU(SKU string) bool {
	for _, item := range Items {
		detailItem := item.GetData().(map[string]any)
		oldSKU := detailItem["SKU"]
		if SKU == oldSKU {
			return true
		}
	}
	return false
}
func isExistIdMember(idMember string) bool {
	for _, member := range Members {
		detailMember := member.GetData().(map[string]any)
		fmt.Print("DETAOL MEMBER : ", detailMember)
		idMemberOld := detailMember["IdMember"]
		if idMember == idMemberOld && (idMember != "") {
			return true
		}
	}
	return false
}

func getDetailItem(SKU string) any {
	for _, item := range Items {
		detailItem := item.GetData().(map[string]any)
		oldSKU := detailItem["SKU"]
		if SKU == oldSKU && (SKU != "") {
			return detailItem
		}
	}
	return nil
}
func getDetailMember(IdMember string) any {
	for _, member := range Members {
		detailMember := member.GetData().(map[string]any)
		idMemberOld := detailMember["IdMember"]
		if idMemberOld == IdMember {
			return detailMember
		}
	}
	return nil
}

func getIndex(SKU string) int {
	var index = 0
	for _, item := range Items {
		detailItem := item.GetData().(map[string]any)
		oldSKU := detailItem["SKU"]
		if SKU == oldSKU {
			return index
		}
		index++
	}
	return -1
}
func AddItem(SKU string, itemName string, price int32, stockQty int32) (string, error) {
	var isExist = isExistSKU(SKU) // Cheking an existing
	if isExist {
		fmt.Println("SUDAH ADA")
		return "", (fmt.Errorf("item %s is already in the list of items", SKU))
	}

	// fmt.Println(Items)
	// // var item = Items.GetData()
	// for _, item := range Items {
	// 	detailItem := item.GetData().(map[string]any)
	// 	oldSKU := detailItem["SKU"]
	// 	fmt.Println(detailItem) //contoh : map[ItemName:jmal Price:12 SKU:agus StockQty:12 Transactions:[]]
	// 	if SKU == oldSKU {
	// 		return "", (fmt.Errorf("item %s is already in the list of items", SKU))
	// 	}
	// }
	var itemMasuk = Item{SKU: SKU, ItemName: itemName, Price: price, StockQty: stockQty}
	Items = append(Items, &itemMasuk)
	return fmt.Sprintf("Successfully added item %s to list item", SKU), nil
}

func DeleteItem(SKU string) (string, error) {
	var detailItem = getDetailItem(SKU)
	if detailItem != nil {
		var transaction, _ = GetTransactionItem(SKU)
		if len(transaction) != 0 {
			return "", fmt.Errorf("there is at least one transaction taking item %s", SKU)
		}
		for i, item := range Items {
			detailItem := item.GetData().(map[string]any)
			oldSKU := detailItem["SKU"]
			if SKU == oldSKU {
				Items = append(Items[:i], Items[i+1:]...)
				fmt.Println(detailItem)

				return fmt.Sprintf("Successfully deleted item %s to list item", SKU), nil

			}
		}
	}
	return "", fmt.Errorf("item %s not in list of items", SKU)
}

func AddMember(idMember string, memberName string) (string, error) {
	var isExist = isExistIdMember(idMember) // Cheking an existing member
	if isExist {
		fmt.Println("SUDAH ADA")
		return "", (fmt.Errorf("Member %s is already in the list of members", idMember))
	}
	var memberMasuk = Member{IdMember: idMember, MemberName: memberName, Transactions: []Transaction{}} //ready for adding item
	Members = append(Members, &memberMasuk)
	fmt.Println(Members)
	return fmt.Sprintf("Successfully added item %s to list item", idMember), nil
}

func DeleteMember(idMember string) (string, error) {
	var detailMember = getDetailMember(idMember)
	if detailMember != nil {
		var transaction, _ = GetTransactionMember(idMember)
		if len(transaction) != 0 {
			return "", fmt.Errorf("there is at least one transaction taking item %s", idMember)
		}
		for i, member := range Members {
			detailMember := member.GetData().(map[string]any)
			idMemberOld := detailMember["IdMember"]
			if idMember == idMemberOld {
				Members = append(Members[:i], Members[i+1:]...)
				fmt.Println(detailMember)
				return fmt.Sprintf("Successfully deleted member %s from list item", idMember), nil
			}
		}
	}
	return "", fmt.Errorf("member %s not in list of items", idMember)
}

func AddTransaction(qty int32, data ...string) (string, error) {
	var SKU = data[0]
	print(SKU)
	var idMember = data[1]
	var itemSelected any
	var SKUExist = isExistSKU(SKU)
	if SKUExist {
		var idMemberExist = isExistIdMember(idMember)
		itemSelected = getDetailItem(SKU)
		if idMemberExist || (idMember == "") {
			// fmt.Printf("idMemberExist: %v\n", idMemberExist)

			if itemMap, ok := itemSelected.(map[string]any); ok {
				Price := itemMap["Price"].(int32)
				if qty > itemMap["StockQty"].(int32) {
					fmt.Print("STOK ABIS ")
					return "", fmt.Errorf("insufficient stock for SKU %s", SKU)
				}
				itemTransaction, _ := GetTransactionItem(SKU)
				memberTransaction, _ := GetTransactionMember(idMember)
				fmt.Print("\nItem Transaction : ", itemTransaction)
				fmt.Print("\nMember Transaction : ", memberTransaction)
				// add member trtansacation
				for _, member := range Members {
					detailMember := member.GetData().(map[string]any)
					oldId := detailMember["IdMember"]
					if idMember == oldId && (SKU != "") {
						newTransaction := Transaction{IdMember: &idMember, Qty: qty, Price: Price, SKU: SKU}
						member.AddTransaction(newTransaction)
					}
				}
				// add item trtansacation

				// memberTransaction = append(memberTransaction, Transaction{IdMember: &idMember, Qty: qty, Price: Price, SKU: SKU})
				for _, item := range Items {
					detailItem := item.GetData().(map[string]any)
					oldSKU := detailItem["SKU"]
					if SKU == oldSKU && (SKU != "") {
						newTransaction := Transaction{IdMember: &idMember, Qty: qty, Price: Price, SKU: SKU}
						item.AddTransaction(newTransaction)
					}
				}

				fmt.Println("OK NIH")

				return fmt.Sprintf("[SUCCESS] Successfully added transaction item  %s", SKU), nil
			}

		} else if !idMemberExist && idMember != "" {
			fmt.Print("GANEMU MEMBER")

			return "", fmt.Errorf("member %s is not in list of members", idMember)
		}
	} else { // jika SKU tidak ditemukan
		fmt.Print("GANEMU SKU")
		return "", fmt.Errorf("item %s is not in list of items", SKU)

	}
	return "", fmt.Errorf("item %s is not in list of items", SKU)

}
func updateItemStock(SKU string, newStockQty int32) {
	for _, item := range Items {
		detailItem := item.GetData().(map[string]any)
		if detailItem["SKU"] == SKU {
			detailItem["StockQty"] = newStockQty
			break
		}
	}
}

func RestockItem(SKU string, qty int32) (string, error) {
	panic("fix me")
}

func GetTransactionItem(SKU string) ([]Transaction, error) {
	var detailItem = getDetailItem(SKU)
	if detailItem != nil {
		detailMap := detailItem.(map[string]any)
		// aman
		var transaction = detailMap["Transactions"].([]Transaction)
		fmt.Println("masunngkngkgn")
		return transaction, nil
	}
	fmt.Print("GAMASUK")
	return nil, fmt.Errorf("no SKU match")

}

func GetTransactionMember(idMember string) ([]Transaction, error) {
	var detailMember = getDetailMember(idMember)
	if detailMember != nil {
		detailMap := detailMember.(map[string]any)
		// aman
		var transaction = detailMap["Transactions"].([]Transaction)
		fmt.Println(transaction)
		return transaction, nil
	}
	return nil, fmt.Errorf("no IdMatch match")
}
