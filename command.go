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

func AddItem(SKU string, itemName string, price int32, stockQty int32) (string, error) {
	var isExist = isExistSKU(SKU) // Cheking an existing
	if isExist {
		return "", (fmt.Errorf(" item %s is already in the list of items", SKU))
	}

	var itemMasuk = Item{SKU: SKU, ItemName: itemName, Price: price, StockQty: stockQty}
	Items = append(Items, &itemMasuk)
	return fmt.Sprintf("SuccSuccessfully added item %s to list item", SKU), nil
}

func DeleteItem(SKU string) (string, error) {
	var detailItem = getDetailItem(SKU)
	if detailItem != nil {
		var transaction, _ = GetTransactionItem(SKU)
		if len(transaction) != 0 {
			return "", fmt.Errorf("[FAILED] there is at least one transaction taking item %s", SKU)
		}
		for i, item := range Items {
			detailItem := item.GetData().(map[string]any)
			oldSKU := detailItem["SKU"]
			if SKU == oldSKU {
				Items = append(Items[:i], Items[i+1:]...)

				return fmt.Sprintf("[SUCCESS] Successfully deleted item %s to list item", SKU), nil

			}
		}
	}
	return "", fmt.Errorf("[FAILED] item %s not in list of items", SKU)
}

func AddMember(idMember string, memberName string) (string, error) {
	var isExist = isExistIdMember(idMember) // Cheking an existing member
	if isExist {
		return "", (fmt.Errorf("[FAILED] Member %s is already in the list of members", idMember))
	}
	var memberMasuk = Member{IdMember: idMember, MemberName: memberName, Transactions: []Transaction{}} //ready for adding item
	Members = append(Members, &memberMasuk)
	return fmt.Sprintf("[SUCCESS] Successfully added item %s to list member", idMember), nil
}

func DeleteMember(idMember string) (string, error) {
	var detailMember = getDetailMember(idMember)
	if detailMember != nil {
		var transaction, _ = GetTransactionMember(idMember)
		if len(transaction) != 0 {
			return "", fmt.Errorf("[FAILED] there is at least one transaction taking item %s", idMember)
		}
		for i, member := range Members {
			detailMember := member.GetData().(map[string]any)
			idMemberOld := detailMember["IdMember"]
			if idMember == idMemberOld {
				Members = append(Members[:i], Members[i+1:]...)
				fmt.Sprintf("[SUCCESS] Successfully deleted member %s from list item", idMember)
				return fmt.Sprintf("[SUCCESS] Successfully deleted member %s from list item", idMember), nil
			}
		}
	}
	return "", fmt.Errorf("[FAILED] member %s not in list of items", idMember)
}

func AddTransaction(qty int32, data ...string) (string, error) {
	var SKU = data[0]
	var idMember = data[1]
	var itemSelected any
	var SKUExist = isExistSKU(SKU)
	if SKUExist {
		for _, item := range Items {
			detailItem := item.GetData().(map[string]any)
			oldSKU := detailItem["SKU"]
			if SKU == oldSKU && (SKU != "") {
				//harus dua kali kerja, karena dalam GO tidak bisa langsung a+=b, bisa tapi hanya menyimpan nilai sementara
				fmt.Println("INI ITEM YANG TERBARU -->> ", detailItem)
			}
		}

		var idMemberExist = isExistIdMember(idMember)
		itemSelected = getDetailItem(SKU)
		for _, item := range Items {
			detailItem := item.GetData().(map[string]any)
			oldSKU := detailItem["SKU"]
			if SKU == oldSKU && (SKU != "") {
			}
		}
		if idMemberExist || (idMember == "") {
			if itemMap, ok := itemSelected.(map[string]any); ok {
				fmt.Println("ITEM MAP --> ", itemMap)
				Price := itemMap["Price"].(int32)
				if qty > Price {
					return "", fmt.Errorf("[FAILED] insufficient stock for SKU %s", SKU)
				}
				for _, member := range Members {
					detailMember := member.GetData().(map[string]any)
					oldId := detailMember["IdMember"]
					if idMember == oldId && (SKU != "") {
						newTransaction := Transaction{IdMember: &idMember, Qty: qty, Price: Price * qty, SKU: SKU}
						member.AddTransaction(newTransaction)
					}
				}

				for _, item := range Items {
					detailItem := item.GetData().(map[string]any)
					oldSKU := detailItem["SKU"]
					if SKU == oldSKU && (SKU != "") {
						newTransaction := Transaction{IdMember: &idMember, Qty: qty, Price: Price * qty, SKU: SKU}
						item.AddTransaction(newTransaction)
					}
				}
				return fmt.Sprintf("[SUCCESS] Successfully added transaction item  %s", SKU), nil
			}

		} else if !idMemberExist && idMember != "" {

			return "", fmt.Errorf("[FAILED] member %s is not in list of members", idMember)
		}
	} else { // jika SKU tidak ditemukan
		return "", fmt.Errorf("[FAILED] item %s is not in list of items", SKU)

	}
	return "", fmt.Errorf("[FAILED] item %s is not in list of items", SKU)

}

func RestockItem(SKU string, qty int32) (string, error) {
	for _, item := range Items {
		detailItem := item.GetData().(map[string]any)
		oldSKU := detailItem["SKU"]
		if SKU == oldSKU && (SKU != "") {
			//harus dua kali kerja, karena dalam GO tidak bisa langsung a+=b, bisa tapi hanya menyimpan nilai sementara
			newStock := detailItem["StockQty"].(int32)
			newStock += qty
			detailItem["StockQty"] = newStock

			fmt.Println("INI ITEM : ", detailItem)
			return fmt.Sprintf("[SUCCESS] Successfully added [%d] Stocks to  Item", detailItem["StockQty"].(int32)), nil
		}
	}
	return "", fmt.Errorf("[FAILED] item %s is not in list of items", SKU)

}

func GetTransactionItem(SKU string) ([]Transaction, error) {
	var detailItem = getDetailItem(SKU)
	if detailItem != nil {
		detailMap := detailItem.(map[string]any)
		// aman
		var transaction = detailMap["Transactions"].([]Transaction)
		fmt.Println("TRANSACTION ITEM : ", transaction)

		return transaction, nil
	}
	return nil, fmt.Errorf("[FAILED] no SKU match")

}

func GetTransactionMember(idMember string) ([]Transaction, error) {
	var detailMember = getDetailMember(idMember)
	if detailMember != nil {
		detailMap := detailMember.(map[string]any)
		// aman
		var transaction = detailMap["Transactions"].([]Transaction)
		fmt.Println("TRANSACTION MEMBER : ", transaction)

		return transaction, nil
	}
	return nil, fmt.Errorf("no IdMatch match")
}
