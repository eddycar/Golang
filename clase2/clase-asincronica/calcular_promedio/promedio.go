/* Calcular promedio
Un colegio necesita calcular el promedio (por estudiante) de sus calificaciones. Se solicita generar una función en la cual se le pueda pasar N cantidad de enteros y devuelva el promedio. No se pueden introducir notas negativas.
*/

package calcular_promedio

func PromedioNotas(notas ...float32)(float32,string){
	var promedio float32
	for _, nota := range notas{
		if nota < 0 {
			return 0, "Se ha introducido una nota con valor negativo"
		}else {
			promedio += nota
		}
	}
	return promedio,""
}