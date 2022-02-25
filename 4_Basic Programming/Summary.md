Basic programing

package fmt
terdiri dari: 1. Output: 1. fmt.Printf()
						 2. fmt.Println()
						 3. fmt.Sprintf()
						 etc
			  2. Scanning: fmt.Scanln()
			  3. Format Verb: %T,%v,%s,%q,etc

Tipe data: 1. Numerik desimal
		   2. Boolean
		   3. String
		   4. Non desimal
		   5. Nilai kosong

Variabel Declaration

var <nama-variabel> <tipe-data>
var <nama-variabel> <tipe-data> = <nilai>

Cons : 1. Single Constans: cons pi float64 = 3.14
	   2. Multiple Constans: cons(
		   pi float64 = 3.14
		   pi2
		   age=10
	   )

Arithmetic Operators
+ penjumlahan
- pengurangan
/ pembagian
* perkalian
% sisa pembagian
++ penjumlahan +1
-- pengurangan -1

Comparison Operators
>
<
>=
<=

If else statement
if number%2 == 0 {
	fmt.Println("Genap")
} else {
	fmt.Printf("Ganjil")
}
Nested if
var v1 int = 400
var v2 int = 700

if v1 == 400 {
	if v2 == 700 {
		fmt.Printf("Value of v1 is 400 and v2 is 700\n");
	}
}

Swith
switch expression {
case condition:
	statement(s)
}
Contoh:
var today int = 2
switch today {
case 1:
	fmt.Printf("Today is Monday")
case 2:
	fmt.Printf("Today is Tuesday")
case 3:
	fmt.Printf("Today is Wednesday")
default:
	fmt.Printf("Invalid Date")
}

Swith without expression
var point = 6

switch {
case point == 8:
    fmt.Println("perfect")
case (point < 8) && (point > 3):
    fmt.Println("awesome")
default:
    {
        fmt.Println("not bad")
        fmt.Println("you need to learn more")
    }
}

Falltrough
var point = 6

switch {
case point == 8:
    fmt.Println("perfect")
case (point < 8) && (point > 3):
    fmt.Println("awesome")
    fallthrough
case point < 5:
    fmt.Println("you need to learn more")
default:
    {
        fmt.Println("not bad")
        fmt.Println("you need to learn more")
    }
}

Looping
for i := 0; i < 5; i++ {
    fmt.Println("Angka", i)
}

While
for main(){
	sum := 1
	for sum < 10 {
		sum += sum
	}
	fmt.Printf(sum)
}

Continue and break

for i := 1; i <= 10; i++ {
    if i % 2 == 1 {
        continue
    }

    if i > 8 {
        break
    }

    fmt.Println("Angka", i)
}