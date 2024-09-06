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
		if len(data) < 4 {
			fmt.Println("[FAILED] your input command is incorrect")
			break

		}
		var SKU = data[0]

		var itemName = data[1]

		var price, error_msg = strconv.Atoi(data[2])
		if error_msg != nil {
			PrintMessage("", error_msg)
			break
		}
		var stockQty, error_msg2 = strconv.Atoi(data[3])
		if error_msg2 != nil {
			PrintMessage("", error_msg)

		}

		result, error := AddItem(SKU, itemName, int32(price), int32(stockQty))
		if error != nil {
			PrintMessage("", error_msg)

		} else {
			PrintMessage(result, nil)

		}
	case "DELETE_ITEM":
		var SKU = data[0]
		result, error := DeleteItem(SKU)
		if error != nil {
			fmt.Println("[FAILED] your input command is incorrect")
			break
		} else {
			PrintMessage(result, nil)

		}

	case "ADD_MEMBER":
		var idMember = data[0]
		var memberName = data[1]
		result, err := AddMember(idMember, memberName)
		if err != nil {
			PrintMessage("", err)

			break
		} else {
			PrintMessage(result, nil)

		}

	case "DELETE_MEMBER":
		var idMember = data[0]
		result, error := DeleteMember(idMember)
		if error != nil {
			PrintMessage("", error)

			break
		} else {
			PrintMessage(result, nil)

		}
	case "ADD_TRANSACTION":
		var qty, err = strconv.Atoi(data[0])
		if err != nil {
			fmt.Println("[FAILED] yosur input command is incorrect")
			break
		} else {
			qtyValid := int32(qty)
			var SKU = data[1]
			var idMember = "-"

			if len(data) > 2 {
				idMember = data[2]
			}

			result, err := AddTransaction(qtyValid, SKU, idMember)

			if err != nil {
				PrintMessage("", err)

				break
			} else {
				fmt.Println("[SUCCESS] ", result)

			}
		}
	case "RESTOCK_ITEM":
		var SKU = data[0]
		var qty, err = strconv.Atoi(data[1])
		if err != nil {
			fmt.Println("[FAILED] your input command is incorrect")
			break
		} else {
			qtyValid := int32(qty)
			result, err := RestockItem(SKU, qtyValid)
			if err != nil {
				fmt.Println("[FAILED] your input command is incorrect")
				break
			} else {
				PrintMessage(result, nil)

			}
		}

	case "TRANSACTION_ITEM_RECAP":
		var SKU = data[0]
		transactions, err := GetTransactionItem(SKU)
		PrintTransactionRecap(transactions, err)

	case "TRANSACTION_MEMBER_RECAP":
		var idMember = data[0]
		transactions, err := GetTransactionMember(idMember)
		PrintTransactionRecap(transactions, err)

	case "EXIT":
		os.Exit(1)
	default:
		os.Exit(1)
	}
}

func PrintMessage(successMsg string, errMsg error) {
	if errMsg != nil {
		fmt.Println("[FAILED] ", errMsg)
	} else { /*  */
		fmt.Println("[SUCCESS] ", successMsg)

	}
}

func PrintTransactionRecap(transactions []Transaction, errMsg error) {
	if errMsg != nil {
		fmt.Println("[FAILED] your input command is incorrect")
	} else { /*  */
		for i := 0; i < len(transactions); i++ {
			fmt.Printf("SKU: %s, ID Member: %s, QTY: %.d, Total Price : %d\n ", transactions[i].SKU, *transactions[i].IdMember, transactions[i].Qty, transactions[i].Price)
		}
	}

}
