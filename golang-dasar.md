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

### Insight Tambahan

-   Gunakan `var` kalau butuh deklarasi tanpa langsung memberi nilai
    (misalnya `var count int`).
-   Gunakan `:=` untuk kode yang lebih singkat & clean, terutama saat
    deklarasi langsung dengan nilai.
-   Best practice: di Go, lebih sering digunakan `:=` di dalam function,
    sementara `var` dipakai untuk variabel global atau ketika butuh tipe
    data tertentu.