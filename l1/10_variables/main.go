package main

/* комментарий
в несколько строк
 */

func main() {
	// целые числа
	var i = 10 							// платформозависимый тип
	var autoInt = -10
	var bigInt int64 = 1<<32 - 1			// int8, int16, int32, 64
	var unsignedInt uint = 100500			// платформонезависимый типб 32 или 64 бита
	var unsignedBigInt uint64 = 1<<64 - 1	// uint8, uint16, uint32, uint64
	println("integers: ", i, autoInt, bigInt, unsignedInt, unsignedBigInt)

	// числа с плавающей точкой
	var p float32 = 3.14					// float = float32, float64
	println("float: ", p)

	// булевы переменные
	var b = true
	println("bool variable: ", b)

	// строки (в двойных кавычках)
	var hi string = "Hi\n\t"
	var me = "me"
	println(hi, me)

	// бинарные данные
	var rawBinary byte = '\x22'
	println("rawBinary: ", rawBinary)

	// приведение типов
	println("float to int", int(p))
	var u1 uint = 17
	var i1 = 21
	println(i1 + int(u1))
	println("int to string", string(48))

	// комплексные числа
	z := 2 + 2i
	println("complex number: ", z)
}
