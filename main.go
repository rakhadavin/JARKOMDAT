package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	// defer func() {
	// 	if r := recover(); r != nil {
	// 		fmt.Println("[CRASH] ", r)
	// 	}
	// }()

	fmt.Printf("Name: %s, ID Student: %s\n", Name, IdStudent)
	fmt.Println("========================================")
	fmt.Println("Welcome to Sigmart Point of Sales")
	fmt.Println("Please input your command below")
	fmt.Println("========================================")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		line := scanner.Text()
		err := scanner.Err()
		if err != nil {
			fmt.Println("[CRASH] ", err.Error())
			os.Exit(1)
		}

		spl := strings.Split(line, " ")
		executeCommand(spl[0], spl[1:])
	}
}

func executeCommand(command string, data []string) {
	command = strings.ToUpper(command)
	switch command {

	case "ADD_ITEM":
		var SKU = data[0]
		fmt.Printf("SKU : %s\n", SKU)
		for _, item := range Items {
			fmt.Println(item.GetData()) // Prints the memory address and type
		}

		var itemName = data[1]

		var price, error_msg = strconv.Atoi(data[2])
		if error_msg != nil {
			fmt.Println("Error:", error_msg)
			panic("Cannot parse from that type")
		}
		var stockQty, error_msg2 = strconv.Atoi(data[3])
		if error_msg2 != nil {
			fmt.Println("Error:", error_msg)
			panic("Cannot parse from that type")
		}

		result, error := AddItem(SKU, itemName, int32(price), int32(stockQty))
		if error != nil {
			fmt.Println("Error:", error)
		} else {
			fmt.Println(result)
		}
	case "DELETE_ITEM":
		var SKU = data[0]
		result, error := DeleteItem(SKU)
		if error != nil {
			fmt.Println("Error:", error)
		} else {
			fmt.Println(result)
		}

	case "ADD_MEMBER":
		var idMember = data[0]
		var memberName = data[0]
		result, error := AddMember(idMember, memberName)
		if error != nil {
			fmt.Println("Error:", error)
		} else {
			fmt.Println(result)
		}

	case "DELETE_MEMBER":
		var idMember = data[0]
		result, error := DeleteMember(idMember)
		if error != nil {
			fmt.Println("Error:", error)
		} else {
			fmt.Println(result)
		}
	case "ADD_TRANSACTION":
		var qty, err = strconv.Atoi(data[0])
		if err != nil {
			fmt.Errorf("[FAILED] your command is incorrect")
		} else {
			qtyValid := int32(qty)
			var SKU = data[1]
			var idMember = ""
			fmt.Println("idmember : ", idMember)

			if len(data) > 2 {
				fmt.Println(len(data))
				idMember = data[2]
			}

			AddTransaction(qtyValid, SKU, idMember)
		}
	case "RESTOCK_ITEM":
		var SKU = data[0]
		var qty, err = strconv.Atoi(data[1])
		if err != nil {
			fmt.Errorf("[FAILED] your command is incorrect")
		} else {
			qtyValid := int32(qty)
			result, err := RestockItem(SKU, qtyValid)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println(result)
			}
		}

	case "TRANSACTION_ITEM_RECAP":
		var SKU = data[0]
		transactions, err := GetTransactionItem(SKU)
		if err != nil {
			fmt.Println("Error:", err)
		} else { /*  Nanti benerin pin disini, outputnya*/
			for i := 0; i < len(transactions); i++ {
				fmt.Printf("%d. IdMember :  %d - SKU: %s, QTY: %.d - Price : %d \n ", i+1, transactions[i].IdMember, transactions[i].SKU, transactions[i].Qty, transactions[i].Price)
			}
		}

	case "TRANSACTION_MEMBER_RECAP":
		var idMember = data[0]
		transactions, err := GetTransactionMember(idMember)
		if err != nil {
			fmt.Println("Error:", err)
		} else { /*  */
			for i := 0; i < len(transactions); i++ {
				fmt.Printf("%d. IdMember :  %d - SKU: %s, QTY: %.d - Price : %d \n ", i+1, transactions[i].IdMember, transactions[i].SKU, transactions[i].Qty, transactions[i].Price)
			}
		}

	case "EXIT":
		os.Exit(1)
	default:
		os.Exit(1)
	}
}

func PrintMessage(successMsg string, errMsg error) {
	panic("fix me")
}

func PrintTransactionRecap(transactions []Transaction, errMsg error) {
	panic("fix me")

}
