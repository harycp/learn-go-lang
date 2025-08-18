# Belajar Go Lang -- Session: *Golang Dasar*

## 1. Pengenalan Go

-   **Go (Golang)** adalah bahasa pemrograman yang dikembangkan oleh
    **Google** (dirilis 2009).
-   Terinspirasi dari C/C++, tetapi **Go adalah bahasa mandiri** dengan
    tujuan: sederhana, cepat dikompilasi, dan efisien dalam *concurrent
    programming*.
-   **Fokus utama:** banyak digunakan di **backend development** dan
    **microservices**.
-   **Ekstensi file:** `.go`

------------------------------------------------------------------------

## 2. Struktur Project

-   Go modern menggunakan **module** sebagai pengelola dependensi.

-   Untuk memulai project baru:

    ``` bash
    go mod init nama-project
    ```

-   Modul di Go mirip dengan `npm init` di Node.js.

------------------------------------------------------------------------

## 3. Program Pertama (Hello World)

``` go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

-   `package main` → titik masuk eksekusi program.
-   `func main()` → hanya boleh ada satu dalam package `main`.
-   `fmt.Println()` → fungsi untuk mencetak ke layar.

------------------------------------------------------------------------

## 4. Aturan Dasar

-   **Satu `main function` saja:**
    Tidak boleh ada lebih dari satu `main` function dalam satu package
    `main`.
-   **Nama function harus unik:**
    Jika ada function dengan nama sama di satu package, akan error saat
    `go build`.
-   **Titik koma (`;`):** opsional, compiler Go otomatis menanganinya.

------------------------------------------------------------------------

## 5. Menjalankan Program

-   Langsung jalankan tanpa build:

    ``` bash
    go run nama-file.go
    ```

-   Atau build menjadi binary:

    ``` bash
    go build nama-file.go
    ./nama-file   # eksekusi hasil build
    ```

------------------------------------------------------------------------

## 6. Tipe Data di Go

### a. Number

-   **Integer:** `int`, `int8`, `int16`, `int32`, `int64`
-   **Unsigned Integer:** `uint`, `uint8`, `uint16`, `uint32`, `uint64`
-   **Floating Point:** `float32`, `float64`

**Alias:** - `byte` → alias untuk `uint8`
- `rune` → alias untuk `int32`
- `int` / `uint` → otomatis menyesuaikan 32-bit atau 64-bit OS

**Contoh:**

``` go
package main

import "fmt"

func main() {
    var a int = 10
    var b float64 = 3.14
    var c byte = 255
    var d rune = 'A'

    fmt.Println(a, b, c, d) // Output: 10 3.14 255 65
}
```

------------------------------------------------------------------------

### b. Boolean

-   **Tipe:** `bool`
-   **Nilai yang mungkin:**
    -   `true` → benar
    -   `false` → salah

**Contoh:**

``` go
package main

import "fmt"

func main() {
    var isActive bool = true
    var isDone bool = false

    fmt.Println("isActive:", isActive)
    fmt.Println("isDone:", isDone)
}
```

Output:

    isActive: true
    isDone: false

------------------------------------------------------------------------

### c. String

-   **Tipe:** `string`
-   **Definisi:** kumpulan karakter.
-   **Ditulis dengan tanda kutip ganda (`" "`).**
-   Go juga mendukung **raw string literal** dengan backtick (`` ` ``) →
    isi string ditulis apa adanya, termasuk baris baru.

**Operasi pada String:** 1. **Menghitung panjang string:** gunakan
fungsi `len()`.
- `len("Hary")` → `4`
2. **Mengakses karakter string:** gunakan indeks `[ ]`, hasilnya berupa
   **byte** (tipe `uint8`), bukan karakter langsung.
- `"Hary"[1]` → `97` (karena `'a'` dalam ASCII adalah 97)
- Jika ingin karakter aslinya, perlu konversi: `string(name[1])` → `"a"`

**Contoh:**

``` go
package main

import "fmt"

func main() {
    var name string = "Hary"
    var age string = "Dua Puluh Satu"

    fmt.Println("Nama saya ", name)
    fmt.Println("Umur saya ", age)

    var panjang int = len(name)
    fmt.Println("Panjang nama saya ", panjang)

    var kedua int = int(name[1]) // dapat byte, bukan langsung string
    fmt.Println("Huruf kedua (ASCII) nama saya ", kedua)

    var hurufKedua string = string(name[1]) // konversi ke string
    fmt.Println("Huruf kedua nama saya adalah", hurufKedua)
}
```

**Output:**

    Nama saya  Hary
    Umur saya  Dua Puluh Satu
    Panjang nama saya  4
    Huruf kedua (ASCII) nama saya  97
    Huruf kedua nama saya adalah a

**Insight tambahan:** - `len()` menghitung panjang dalam **byte**, bukan
jumlah karakter Unicode.
- Untuk teks non-ASCII (emoji, huruf Jepang, dll.), jumlah byte bisa
  lebih besar dari jumlah huruf.
- Gunakan `rune` untuk representasi karakter Unicode agar hasil sesuai
  jumlah karakter sebenarnya.

------------------------------------------------------------------------

## 7. Insight Sementara

-   Go adalah bahasa yang **strictly typed** → tipe data harus jelas.
-   Struktur project Go sederhana: butuh modul (`go mod init`) dan hanya
    satu `main function`.
-   Tipe data dasar: Number, Boolean, dan String adalah fondasi awal
    sebelum masuk ke array, slice, map, dan struct.
-   Sangat efisien untuk membuat aplikasi skala besar seperti
    **microservices** karena kompile cepat dan binary ringan.

------------------------------------------------------------------------

## 8. Deklarasi & Inisialisasi Variabel

-   Di Go, variabel bisa dibuat dengan dua cara:

### a. Menggunakan `var`

``` go
var name string = "Hary"
var age int = 21
```

-   Kelebihan: bisa eksplisit menentukan tipe data.

------------------------------------------------------------------------

### b. Tanpa `var` → menggunakan `:=`

``` go
name := "Hary Capri"
alas := 21
```

-   Go otomatis mendeteksi tipe data dari nilai yang diberikan.
    -   `"Hary Capri"` → `string`
    -   `21` → `int`
-   Hanya bisa digunakan **di dalam function** (tidak bisa untuk
    deklarasi global).

------------------------------------------------------------------------

### c. Contoh Lengkap

``` go
package main

import "fmt"

func main() {
    // Deklarasi eksplisit dengan var
    var umur int = 22

    // Inisialisasi singkat dengan :=
    name := "Hary Capri"
    alas := 21

    fmt.Println("Nama:", name)
    fmt.Println("Umur:", umur)
    fmt.Println("Alas:", alas)
}
```

**Output:**

    Nama: Hary Capri
    Umur: 22
    Alas: 21

------------------------------------------------------------------------
## d. Multiple Variable Declaration di Go

Go juga mendukung deklarasi **multiple variable** dalam satu blok menggunakan tanda kurung `()`. Hal ini mempermudah pengelompokan variabel yang terkait.

### Contoh:
```go
package main

import "fmt"

func main() {
    var (
        firstName = "Hary"
        lastName  = "Capri"
        Age       = 21
    )

    fmt.Println(firstName, lastName, Age)
}
```

### Output:
```
Hary Capri 21
```
------------------------------------------------------------------------
## 9. Constant di Go

Selain variabel, Go juga memiliki **constant** yaitu nilai yang tidak bisa diubah lagi setelah pertama kali dibuat dan di-assign.

### Aturan Constant
- Dideklarasikan menggunakan keyword `const`.
- Nilai harus sudah diketahui saat compile-time.
- Tidak bisa diubah setelah didefinisikan.

### Contoh:
```go
package main

import "fmt"

func main() {
    const name string = "Hary Capri"
    const age = 25

    const (
        University = "UPNVJ"
        IPK        = 3.92
    )

    fmt.Println(name)
    fmt.Println(age)
    fmt.Println(IPK)
    fmt.Println(University)
}
```

### Output:
```
Hary Capri
25
3.92
UPNVJ
```
------------------------------------------------------------------------
### 10. Type Conversion
Go bersifat **strongly typed**, jadi tipe data berbeda tidak bisa langsung dicampur. Harus dilakukan **konversi tipe** secara eksplisit.

```go
package main

import "fmt"

func main() {
    var a int32 = 90
    var b = int64(a) // konversi int32 ke int64

    fmt.Println(b)

    var name string = "Hary Capri"
    var pick = name[2] // ambil byte ke-2 (index dimulai dari 0)

    var conversion = string(pick) // konversi byte ke string
    fmt.Println(conversion)
}
```

**Output:**
```
90
r
```
------------------------------------------------------------------------
### 11. Type Declaration
Go memungkinkan membuat **tipe data baru** berdasarkan tipe bawaan.

```go
package main

import "fmt"

type noNim string
type noKk int

func main() {
    var nim noNim = "2210511023"
    var kk noKk = 22892929291

    fmt.Println(nim)
    fmt.Println(kk)
}
```

**Output:**
```
2210511023
22892929291
```

**Insight:**
- Berguna untuk memberikan konteks pada data (misalnya `noNim` dan `noKk`), walaupun dasarnya tetap string atau int.
- Membantu readability & keamanan kode.
------------------------------------------------------------------------
### 12. Operasi Matematika
Go mendukung operasi matematika standar: `+`, `-`, `*`, `/`, dan `%` (modulus).

```go
package main

import "fmt"

func main() {
    a := 2
    b := 4
    c := 10
    d := 9
    e := 7
	f := 2

    result := (a + (b * c) - d) / e
	modulus := result % f
    fmt.Println(result)
    fmt.Println(modulus)
}
```

**Output:**
```
4
0
```

---

### Insight
- **Konversi tipe** wajib di Go, tidak ada auto-casting.
- **Type declaration** memudahkan memberi identitas baru pada tipe bawaan.
- **Operasi matematika** sama seperti bahasa lain, namun tipe data harus konsisten (tidak bisa langsung mencampur `int` dan `float64` tanpa konversi).
------------------------------------------------------------------------
### 13. Operator Perbandingan:
- `==` : sama dengan
- `!=` : tidak sama dengan
- `>`  : lebih besar
- `<`  : lebih kecil
- `>=` : lebih besar atau sama dengan
- `<=` : lebih kecil atau sama dengan

### Contoh:
```go
package main

import "fmt"

func main() {
    a := "Hary"
    b := "Hary"
    c := 2
    d := 8

    equal := a == b
    notEqual := a != b
    greater := c > d
    less := c < d
    greaterEqual := c >= d
    lessEqual := c <= d

    fmt.Println(equal)
    fmt.Println(notEqual)
    fmt.Println(greater)
    fmt.Println(less)
    fmt.Println(greaterEqual)
    fmt.Println(lessEqual)
}
```

### Output:
```
true
false
false
true
false
true
```

### Insight:
- Operator perbandingan di Go bekerja pada tipe data yang **sama** (misalnya string dengan string, int dengan int).
- Hasil evaluasi selalu berupa **boolean**.
- Berguna untuk logika percabangan (if-else) dan kontrol program lainnya.

### 14. Operasi Boolean di Go

Go menyediakan operator boolean standar: `&&` (AND), `||` (OR), dan `!` (NOT). Operator ini digunakan untuk menggabungkan atau membalik nilai boolean.

### Contoh:
```go
package main

import "fmt"

func main() {
    var option1, option2, option3, option4, option5, option6, option7, option8, option9 bool
    var benar bool = true
    var salah bool = false

    option1 = benar && benar
    option2 = benar && salah
    option3 = salah && benar
    option4 = salah && salah

    option5 = benar || benar
    option6 = benar || salah
    option7 = salah || benar
    option8 = salah || salah

    option9 = !(benar && benar)
    
    fmt.Println(option1)
    fmt.Println(option2)
    fmt.Println(option3)
    fmt.Println(option4)
    fmt.Println(option5)
    fmt.Println(option6)
    fmt.Println(option7)
    fmt.Println(option8)
    fmt.Println(option9)
}
```

### Output:
```
true
false
false
false
true
true
true
false
false
```

### Insight:
- `&&` menghasilkan `true` hanya jika **kedua operand true**.
- `||` menghasilkan `true` jika **salah satu operand true**.
- `!` membalik nilai boolean.
- Operasi boolean sangat penting dalam pengendalian alur program (misalnya `if`, `for`, atau logika kompleks lainnya).
