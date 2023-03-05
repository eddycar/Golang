/* Impuestos de salario
Una empresa de chocolates necesita calcular el impuesto de sus empleados al momento de depositar el sueldo.
Para cumplir el objetivo es necesario crear una función que devuelva el impuesto de un salario, teniendo en cuenta que si la persona gana más de $50.000 se le descontará un 17 % del sueldo y si gana más de $150.000 se le descontará, además, un 10 % (27 % en total).
*/

package impuesto_salario

func ImpuestoSalario(salario float32) float32{
	var salarioConImpuesto float32
	if salario >= 50000 {
		salarioConImpuesto = salario / 1.17
	} else if salario >= 150000{
		salarioConImpuesto = salario / 1.27
	}
	return salarioConImpuesto
}