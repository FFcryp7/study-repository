package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generatePrice() float64 {
	var newPrice = float64(20000 + rand.Intn(50001))
	return newPrice
}

type Account struct {
	BalanceUSDT float64
	BalanceBTC  float64
}

func readCmd() string {
	var cmd string
	fmt.Println("Buy or sell: ")
	fmt.Scan(&cmd)
	return cmd
}

func (u *Account) funcBuy(price *float64, oldPrice *float64) {
	var usdt float64
	fmt.Println("-------------------")
	fmt.Println("Введите сумму в USDT: ")
	fmt.Scan(&usdt)
	if u.BalanceUSDT < usdt {
		fmt.Println("❌ Не хватает USDT!")
	} else {
		btcBought := usdt / *price
		u.BalanceBTC += btcBought
		u.BalanceUSDT -= usdt
		fmt.Printf("✅ Куплено %.8f BTC\n", btcBought)
		*oldPrice = *price
		*price = generatePrice()
		percentPrice := ((*price - *oldPrice) / *oldPrice) * 100
		time.Sleep(1000 * time.Millisecond)
		fmt.Print("💰 Цена обновилась!\n", *oldPrice, " --- ")
		fmt.Printf("%.2f%% ---> ", percentPrice)
		fmt.Printf("%.f \n", *price)
	}
}

func (u *Account) funcSell(price *float64, oldPrice *float64) {
	var usdt float64
	fmt.Println("--------------------")
	fmt.Println("Введите сумму в USDT: ")
	fmt.Scan(&usdt)
	needBTC := usdt / *price
	if needBTC > u.BalanceBTC {
		fmt.Println("❌ Не хватает BTC!")
	} else {
		u.BalanceBTC -= needBTC
		u.BalanceUSDT += usdt
		fmt.Printf("✅ Продано %.8f BTC\n", needBTC)
		ptrOldPrice := *price
		*oldPrice = ptrOldPrice
		*price = generatePrice()
		percentPrice := ((*price - *oldPrice) / *oldPrice) * 100
		time.Sleep(1000 * time.Millisecond)
		fmt.Print("💰 Цена обновилась!\n", *oldPrice, " --- ")
		fmt.Printf("%.2f%% ---> ", percentPrice)
		fmt.Printf("%.f \n", *price)
	}
}

func (u Account) funcBalance(price float64) {
	stats := ((((u.BalanceBTC * price) + u.BalanceUSDT) - 1000) / 1000) * 100
	allBalance := (u.BalanceBTC * price) + u.BalanceUSDT
	fmt.Printf("💎 Общий баланс: ~ %.2f USDT\n", allBalance)
	fmt.Println("💵 USDT: ", u.BalanceUSDT, "USDT")
	fmt.Printf("🪙  BTC: %.8f", u.BalanceUSDT)
	fmt.Printf(" (~ %.2f USDT)\n", (u.BalanceBTC * price))
	if stats > 0 {
		fmt.Printf("📈 Прибыль: %.2f%%\n", stats)
		fmt.Println("--------------------")
	} else if stats < 0 {
		fmt.Printf("📉 Убыток: %.2f%%\n", stats)
		fmt.Println("--------------------")
	} else {
		fmt.Println("⚖️ В нуле")
		fmt.Println("--------------------")
	}
	fmt.Println("Нажмите Enter для продолжения...")
	fmt.Scanln()
}

func (u Account) funcExit(price float64) {
	fmt.Println("--------------------")
	fmt.Println("🎯 Твой результат: ")
	fmt.Println("--------------------")
	u.funcBalance(price)
}

func main() {
	user := Account{BalanceUSDT: 1000}
	price := generatePrice()
	oldPrice := price
	var cmd string
	fmt.Println("--------------------")
	fmt.Println("Добрый день! Ты попал на мини биржу!")
	fmt.Println("Актуальная цена BTC = ", price, "USDT")
	for {
		fmt.Println("--------------------")
		fmt.Printf("💵 Доступно в USDT: %.2f\n", user.BalanceUSDT)
		fmt.Printf("🪙 Доступно в BTC: %.8f", user.BalanceBTC)
		fmt.Printf(" (~ %.2f USDT)\n", (user.BalanceBTC * price))
		cmd = readCmd()
		fmt.Println("--------------------")
		switch cmd {
		case "Buy", "buy":
			user.funcBuy(&price, &oldPrice)
		case "Sell", "sell":
			user.funcSell(&price, &oldPrice)
		case "Balance", "balance":
			user.funcBalance(price)
		case "Exit", "exit":
			user.funcExit(price)
			return
		default:
			fmt.Println("Что то я не распознал...")
		}
	}
}
