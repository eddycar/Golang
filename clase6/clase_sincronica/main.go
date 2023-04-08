/* Algunas tiendas ecommerce necesitan realizar una funcionalidad en Go para administrar los diferentes métodos
de pago que se ofrecen al momento de realizar una nueva compra.
Actualmente, las empresas cuentan con 3 métodos de pago:
	*Tarjeta de Crédito, Transferencia Bancaria y Efectivo (Se espera que sean muchos más).
	*Cada método de pago, cuenta con su propia comisión por la utilización del servicio.
	Dichas comisiones son:
	*Tarjeta de Crédito: cuenta con un 10% de recargo sobre el precio publicado.
	*Transferencia Bancaria: no presenta recargo sobre el precio publicado.
	*Efectivo: cuenta con un 5% de descuento sobre el precio publicado.

Requerimientos:
	*Crear tres estructuras que representan los diferentes métodos de pagos permitidos.
	*Crear una estructura Purchase que tenga el precio y un arreglo de métodos de pagos disponibles para la compra.
	Deberá crear una función que permita setear el precio de la compra al momento de crear la misma.
	*Crear una interfaz “PaymentMethod” que defina el método Pay(purchase *purchase) a ser implementado por cada método de pago disponible.
	*Crear una función Process que nos permita procesar el pago de la compra recibiendo como parámetro el método de pago elegido.
	*Solicitar al usuario ingresar el método de pago seleccionado para la compra. */

package main

import (
	"errors"
	"fmt"
	"log"
)

const(
	CreditCardCommission   = 1.10
	BankTransferCommission = 1.0
	CashCommission         = 0.95
)


func main() {
	purchase := NewPurchase(1000)
	err := purchase.RegisterPaymentMethod("cash", &cash{CashCommission})
	if err != nil {
		log.Fatalln(err)
	}
	err = purchase.RegisterPaymentMethod("creditCard", &creditCard{CreditCardCommission})
	if err != nil {
		log.Fatalln(err)
	}
	err = purchase.RegisterPaymentMethod("bank", &bankTransfer{BankTransferCommission})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Enter the payment method you want to use (%v): \n", purchase.getAvailablePaymentMethods())
	var paymentMethod string
	fmt.Scanln(&paymentMethod)

	if err := purchase.Process(paymentMethod); 
	err != nil {
		log.Fatalln(err)
	}
}

type PaymentMethod interface{
	Pay(purchase *purchase)
}

type purchase struct{
	Price float64
	PaymentMethod map[string]PaymentMethod
}
func (p *purchase)SetPrice(price float64){
	p.Price = price
}
func (p *purchase)GetPrice() float64{
	return p.Price
}
func (p *purchase)Process(paymentMethod string) error{
	method, ok := p.PaymentMethod[paymentMethod]
	if !ok {
		return errors.New("Method not found")
	}

	method.Pay(p)

	fmt.Printf("Amount to pay: %.2f\n", p.GetPrice())
	fmt.Printf("Purchase with %s completed\n", paymentMethod)
	return nil
}

func NewPurchase(price float64) purchase {
	return purchase{
		PaymentMethod: map[string]PaymentMethod{},
		Price:         price,
	}
}

func (p *purchase) RegisterPaymentMethod(name string, paymentMethod PaymentMethod) error {
	_, ok := p.PaymentMethod[name]
	if ok {
		return errors.New("method alredy exist")
	}

	p.PaymentMethod[name] = paymentMethod
	return nil
}

func (p *purchase) getAvailablePaymentMethods() []string {
	var availableMethod []string
	for name := range p.PaymentMethod {
		availableMethod = append(availableMethod, name)
	}
	return availableMethod
}

type cash struct {
	Commission float64
}
func (c *cash)Pay(purchase *purchase){
	purchase.SetPrice(purchase.GetPrice() * CashCommission)
}

type bankTransfer struct {
	Commission float64
}
func (b *bankTransfer)Pay(purchase *purchase){
	purchase.SetPrice(purchase.GetPrice() * BankTransferCommission)
}

type creditCard struct {
	Commission float64
}
func (cred *creditCard)Pay(purchase *purchase){
	purchase.SetPrice(purchase.GetPrice() * CreditCardCommission)
}





