# Belajar Go Lang -- Session: *Golang Dasar*

## 1. Pengenalan Go

-   **Go (Golang)** adalah bahasa pemrograman yang dikembangkan oleh
    **Google** (dirilis 2009).\
-   Terinspirasi dari C/C++, tetapi **Go adalah bahasa mandiri** dengan
    tujuan: sederhana, cepat dikompilasi, dan efisien dalam *concurrent
    programming*.\
-   **Fokus utama:** banyak digunakan di **backend development** dan
    **microservices**.\
-   **Ekstensi file:** `.go`

------------------------------------------------------------------------

## 2. Struktur Project

-   Go modern menggunakan **module** sebagai pengelola dependensi.\

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

-   `package main` → titik masuk eksekusi program.\
-   `func main()` → hanya boleh ada satu dalam package `main`.\
-   `fmt.Println()` → fungsi untuk mencetak ke layar.

------------------------------------------------------------------------

## 4. Aturan Dasar

-   **Satu `main function` saja:**\
    Tidak boleh ada lebih dari satu `main` function dalam satu package
    `main`.\
-   **Nama function harus unik:**\
    Jika ada function dengan nama sama di satu package, akan error saat
    `go build`.\
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

-   **Number:**
    -   **Integer:** `int`, `int8`, `int16`, `int32`, `int64`
    -   **Unsigned Integer:** `uint`, `uint8`, `uint16`, `uint32`,
        `uint64`
    -   **Floating Point:** `float32`, `float64`
-   **Alias:**
    -   `byte` → alias untuk `uint8`
    -   `rune` → alias untuk `int32`
    -   `int` / `uint` → otomatis menyesuaikan 32-bit atau 64-bit OS

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

## 7. Insight Sementara

-   Go adalah bahasa yang **strictly typed** → tipe data harus jelas.\
-   Struktur project Go sederhana: butuh modul (`go mod init`) dan hanya
    satu `main function`.\
-   Sangat efisien untuk membuat aplikasi skala besar seperti
    **microservices** karena kompile cepat dan binary ringan.
