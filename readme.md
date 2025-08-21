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
------------------------------------------------------------------------
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
------------------------------------------------------------------------
### 15. Array di Go

Array adalah kumpulan elemen dengan jumlah tetap dan bertipe data sama.

### Karakteristik Array di Go
- Panjang array harus ditentukan dan **tidak bisa berubah** setelah deklarasi.
- Semua elemen array harus memiliki tipe data yang sama.
- Index dimulai dari `0`.

### Contoh:
```go
package main

import "fmt"

func main() {
    // Deklarasi array sekaligus inisialisasi
    var names = [4]string{
        "John",
        "Paul",
        "George",
        "Ringo",
    }

    fmt.Println(names)

    // Deklarasi array kosong dengan panjang 5
    var car [5]string

    car[0] = "Hyundai"
    car[1] = "Toyota"
    car[2] = "BYD"
    car[3] = "Honda"
    car[4] = "Tesla"

    fmt.Println(car)
}
```

### Output:
```
[John Paul George Ringo]
[Hyundai Toyota BYD Honda Tesla]
```

### Insight:
- Array berguna untuk menyimpan data yang jumlahnya sudah pasti.
- Jika butuh ukuran dinamis, biasanya digunakan **slice** (nanti akan dibahas setelah array).
- Mengakses elemen array: `array[index]`.
- Panjang array bisa diperiksa dengan fungsi `len(array)`. 


### 16. Slice di Go — Penjelasan Lengkap

Slice adalah struktur data dinamis yang **mereferensikan** array di belakang layar. Secara internal slice terdiri dari:
- **Pointer** → alamat elemen pertama yang dapat diakses slice pada array asal
- **Length (len)** → jumlah elemen aktual dalam slice
- **Capacity (cap)** → jumlah maksimum elemen dari posisi pointer hingga akhir array asal

> Zero value slice adalah `nil` (len = 0, cap = 0). Berbeda dengan empty slice `[]T{}` yang len = 0 tapi **bukan** nil.

---

### Cara Membuat Slice
1) **Slice literal** (langsung):
```go
names := []string{"Hary", "Capri", "Sukro"}
```
2) **Slicing** dari array/slice:
```go
arr := [5]int{10, 20, 30, 40, 50}
s := arr[1:4] // -> {20,30,40}
```
3) **make** (dengan panjang & kapasitas):
```go
s := make([]int, 3)        // len=3, cap=3, berisi zero value {0,0,0}
t := make([]int, 3, 5)     // len=3, cap=5
```

---

### Notasi Slicing & Dampaknya
Misal `base := []string{"A","B","C","D","E","F"}` (len=6, cap=6)

- **`base[low:high]`** → elemen index `low` s.d. `high-1`
    - `u := base[2:5]` → `{C,D,E}`  (len=3, cap=4 → sisa dari pos 2 hingga akhir)
- **`base[low:]`** → dari `low` hingga akhir
    - `v := base[3:]` → `{D,E,F}`   (len=3, cap=3)
- **`base[:high]`** → dari awal hingga `high-1`
    - `w := base[:4]` → `{A,B,C,D}` (len=4, cap=6)
- **`base[:]`** → seluruh isi (re-slice penuh)
    - `x := base[:]`  → `{A,B,C,D,E,F}` (len=6, cap=6)
- **Full slice 3-indeks `base[low:high:max]`** → mengatur **capacity** slice (`cap = max - low`). Berguna agar `append` **tidak** menimpa array asal.
    - `y := base[1:3:3]` → `{B,C}` (len=2, cap=2) → `append(y, "Z")` pasti **alokasi array baru**.

> **Catatan penting:** `cap` dihitung relatif dari posisi `low` menuju batas `max` (atau akhir array jika 2-indeks). Kapasitas menentukan apakah `append` akan menimpa array lama atau mengalokasikan array baru.

---

### `len()` dan `cap()`
```go
s := []int{1,2,3}
fmt.Println(len(s)) // 3
fmt.Println(cap(s)) // 3 (tergantung asalnya)
```

---

### `append` — Menambah Elemen
- `append(dst, elems...)` **mengembalikan slice baru** (bisa menunjuk array lama atau array baru). **Selalu simpan hasilnya!**
- Jika masih ada kapasitas, elemen baru diletakkan pada array yang sama → dapat **menimpa** data jika ada slice lain yang mereferensikan array tersebut.
- Jika kapasitas habis, Go mengalokasikan **array baru**, menyalin isi lama, lalu menambahkan elemen.

Contoh dampak `append` pada array asal vs array baru:
```go
base := [...]string{"A","B","C","D"}
s := base[1:3]           // {B,C}, len=2 cap=3 (B..D)
s[0] = "B*"              // base -> {A,"B*","C","D"}

s2 := append(s, "X")     // masih muat (cap=3), menulis ke base index 3
fmt.Println(base)         // {A B* C X}  <- base ikut berubah!

// Batasi kapasitas dengan full slice 3-indeks
s3 := base[0:2:2]         // {A, B*}, cap=2
s4 := append(s3, "Y")    // cap penuh -> alokasi array baru
fmt.Println(base)         // {A B* C X}  (tidak berubah)
fmt.Println(s4)           // {A B* Y}
```

---

### `make` — Membuat Slice Kosong Berkapasitas
```go
nums := make([]int, 0, 4) // len=0 cap=4
nums = append(nums, 1,2,3)
fmt.Println(len(nums), cap(nums)) // 3 4
```
> `make` menginisialisasi backing array dengan zero value dan mengatur len/cap sesuai parameter.

---

### `copy` — Menyalin Isi Slice (Detach)
Gunakan `copy` untuk **menyalin** isi slice ke backing array baru sehingga perubahan tidak saling memengaruhi.
```go
src := []int{1,2,3}
dst := make([]int, len(src))
copy(dst, src)            // mengembalikan jumlah elemen yang disalin
src[0] = 99
fmt.Println(src)          // [99 2 3]
fmt.Println(dst)          // [1 2 3]  (terpisah)
```

---

### Contoh
```go
package main

import "fmt"

func main() {
    var name = []string{
        "Hary","Capri","Sukro","Makro","Surtono","Tartono","Hendrawan","Hensubroto","Nabuoloto",
    }

    slice1 := name[3:9]  // [Makro .. Nabuoloto], len=6 cap=6
    slice2 := name[5:]   // [Tartono .. Nabuoloto], len=4 cap=4
    slice3 := name[:4]   // [Hary Capri Sukro Makro]
    slice4 := name[:]    // seluruh isi

    fmt.Println(slice1)
    fmt.Println(slice2)
    fmt.Println(slice3)
    fmt.Println(slice4)

    // len & cap
    fmt.Println(len(slice1))
    fmt.Println(cap(slice1))

    // Contoh array -> slice dan efek edit
    var days = [...]string{"Senin","Selasa","Rabu","Kamis","Jumat","Sabtu","Minggu"}
    daysSlice1 := days[0:7]
    daysSlice1[1] = "Selasa Baru"
    daysSlice1[2] = "Rabu Baru"
    fmt.Println(daysSlice1)

    // append mungkin memodifikasi array asal jika masih ada cap
    daysSlice2 := append(daysSlice1, "Hari Baru")
    fmt.Println(daysSlice2)
    fmt.Println(days)

    // Gunakan full slice 3-indeks untuk membatasi cap -> paksa alokasi baru saat append
    safe := days[0:2:2]                 // len=2 cap=2
    safeAppended := append(safe, "X")  // alokasi baru
    fmt.Println(safe, safeAppended)

    // make & append
    newSlice := make([]string, 9, 11)
    for i := 0; i < 9; i++ { newSlice[i] = "Hari" }
    appendSlice := append(newSlice, "Hari")
    fmt.Println(newSlice)
    fmt.Println(appendSlice)

    // copy untuk detach
    cloned := make([]string, len(slice2))
    copy(cloned, slice2)
    slice2[0] = "(Edited)"
    fmt.Println("slice2:", slice2)
    fmt.Println("cloned:", cloned) // tetap data lama
}
```

---

### Insight
- Slice = pointer + len + cap.
- `len` = jumlah elemen; `cap` = ruang yang tersedia hingga batas backing array.
- `append` harus ditampung hasilnya; bisa menimpa array asal jika masih ada kapasitas.
- Batasi efek samping dengan full slice `a[low:high:max]` atau gunakan `copy` untuk memisahkan data.
- Gunakan `make` untuk menyiapkan slice dengan kapasitas yang direncanakan agar mengurangi alokasi ulang.
------------------------------------------------------------------------
### 16. Map di Go

**Map** adalah struktur data yang menyimpan pasangan **key-value** (kunci-nilai). Berbeda dengan array/slice yang diakses dengan index angka, map diakses menggunakan kunci (key) yang bisa berupa tipe data tertentu.

### Karakteristik Map
- Kunci (key) bersifat **unik**, tidak boleh duplikat.
- Nilai (value) bisa bertipe data apa saja, tergantung deklarasi.
- Jika mengakses key yang tidak ada, Go mengembalikan nilai **zero value** dari tipe data value.
- Jumlah data di map bisa bertambah/berkurang dinamis.

---

### Cara Membuat Map
1. Menggunakan `make`:
```go
student := make(map[string]string)
student["name"] = "Hary"
student["age"] = "20"
```

2. Menggunakan literal map:
```go
students := map[string]string{
    "name":     "Hary Capri",
    "age":      "18",
    "semester": "7",
    "major":    "Informatics",
}
```

---

### Contoh Lengkap:
```go
package main

import "fmt"

func main() {
    var students = map[string]string{
        "name":     "Hary Capri",
        "age":      "18",
        "semester": "7",
        "major":    "Informatics",
    }

    fmt.Println(students["name"])
    fmt.Println(students["age"])
    fmt.Println(students["semester"])
    fmt.Println(students["major"])
    fmt.Println(students["majorx"]) // key tidak ada -> hasil zero value ("")
    fmt.Println(students)

    // Function Map
    fmt.Println(len(students))       // jumlah elemen map
    fmt.Println(students["major"])  // akses key

    students["majorx"] = "InformaticsX" // menambah key baru
    fmt.Println(students["majorx"])

    var studentCopy = make(map[string]string)
    studentCopy = students            // referensi, bukan copy data baru
    fmt.Println(studentCopy)

    delete(studentCopy, "name")      // menghapus key "name"
    fmt.Println("Hapus name :", studentCopy["name"]) // akan tampilkan ""
}
```

### Output (ringkas):
```
Hary Capri
18
7
Informatics

map[age:18 major:Informatics majorx:InformaticsX semester:7]
4
Informatics
InformaticsX
map[age:18 major:Informatics majorx:InformaticsX semester:7]
Hapus name : 
```

---

### Fungsi Bawaan Map
- `len(map)` → menghitung jumlah elemen.
- `delete(map, key)` → menghapus elemen berdasarkan key.
- Akses `map[key]` dengan key yang tidak ada → mengembalikan zero value.

---

### Insight
- Map sangat cocok untuk data dengan pasangan **key-value** seperti dictionary, config, dan data lookup.
- Hati-hati: assignment `mapB = mapA` membuat referensi, bukan salinan baru. Perubahan di salah satu akan memengaruhi yang lain.
- Untuk benar-benar meng-copy isi map, perlu iterasi manual atau utility tertentu.
------------------------------------------------------------------------
### 17. If Statement di Go

Go menyediakan percabangan menggunakan `if`, `else if`, dan `else` untuk mengeksekusi blok kode berdasarkan kondisi tertentu.

### Karakteristik If Statement
- Kondisi `if` harus berupa ekspresi boolean (`true` atau `false`).
- Bisa menggunakan satu atau beberapa blok `else if`.
- `else` akan dieksekusi jika semua kondisi sebelumnya salah.
- Tidak menggunakan tanda kurung `()` untuk kondisi, hanya `{}` untuk blok kode.

---

### Contoh:
```go
package main

import "fmt"

func main() {
    input := "tambah"
    var a, b, result int

    a = 2
    b = 9

    if input == "tambah" {
        result = a + b
    } else if input == "kurang" {
        result = a - b
    } else {
        result = a * b
    }

    fmt.Println("Input :", input, "Hasil :", result)
}
```

### Output:
```
Input : tambah Hasil : 11
```

---

### Insight
- `if` di Go sederhana: tidak perlu `()` seperti di Java/C.
- Bisa dipakai dengan deklarasi singkat sebelum kondisi:
  ```go
  if x := 10; x > 5 {
      fmt.Println("x lebih besar dari 5")
  }
  ```
- Sangat umum digunakan untuk logika percabangan dalam program.
------------------------------------------------------------------------
### 18. Switch Statement di Go

Selain `if-else`, Go juga menyediakan **switch statement** untuk percabangan yang lebih sederhana saat membandingkan satu variabel dengan banyak kemungkinan nilai.

### Karakteristik Switch di Go
- Mengevaluasi ekspresi sekali, lalu dibandingkan dengan setiap `case`.
- Jika cocok, blok kode di `case` tersebut dieksekusi.
- Secara default, Go **tidak melakukan fallthrough** ke case berikutnya (berbeda dengan C/Java). Jika ingin melanjutkan ke case berikutnya, harus menuliskan `fallthrough`.
- `default` dieksekusi jika tidak ada case yang cocok.

---

### Contoh:
```go
package main

import "fmt"

func main() {
    var option string
    var a, b, result float64

    option = "bagi"
    a = 9
    b = 4

    switch option {
    case "tambah":
        result = a + b
    case "kurang":
        result = a - b
    case "bagi":
        result = a / b
    case "kali":
        result = a * b
    default:
        result = 0
    }

    fmt.Println(result)
}
```

### Output:
```
2.25
```

---

### Insight
- Gunakan `switch` saat ada banyak percabangan dari satu variabel.
- Lebih rapi daripada `if-else if` yang panjang.
- Mendukung penggunaan tanpa ekspresi langsung:
  ```go
  switch {
  case a > b:
      fmt.Println("a lebih besar")
  case a < b:
      fmt.Println("a lebih kecil")
  default:
      fmt.Println("sama")
  }
  ```
------------------------------------------------------------------------
### 19. For Loop, Continue, dan Break di Go

Go hanya memiliki satu jenis loop, yaitu **for**, namun dapat digunakan dalam berbagai bentuk.

---

### Bentuk For Loop
1. **While-like loop** (hanya kondisi):
```go
count := 1
for count <= 10 {
    fmt.Println(count)
    count++
}
```

2. **Classic for** (init; condition; post):
```go
for total := 0; total < count; total++ {
    if total%2 == 0 {
        fmt.Println("Genap:", total)
    }
    fmt.Println("Ganjil:", total)
}
```

3. **Loop dengan slice menggunakan index:**
```go
names := []string{"Alice", "Bob", "Charlie", "Ucup"}
for i := 0; i < len(names); i++ {
    fmt.Println("Index ke-", i, "=", names[i])
}
```

4. **Range loop:**
```go
for index, name := range names {
    fmt.Println("Index ke-", index, "=", name)
}

for _, name := range names {
    fmt.Println("names =", name)
}
```

---

### Break
`break` digunakan untuk menghentikan loop lebih awal.
```go
for i := 0; i < 10; i++ {
    if i == 5 {
        fmt.Println(i)
        break
    }
}
```

**Output:**
```
5
```

---

### Continue
`continue` digunakan untuk melewati iterasi saat ini dan lanjut ke iterasi berikutnya.
```go
for i := 0; i < 10; i++ {
    if i%2 == 1 {
        continue
    }
    fmt.Println(i)
}
```

**Output:**
```
0
2
4
6
8
```

---

### Insight
- **for** di Go fleksibel, bisa menggantikan `while` maupun `for` klasik di bahasa lain.
- `range` mempermudah iterasi pada array, slice, map, atau string.
- `break` dan `continue` bekerja sama seperti bahasa lain (C/Java/JS).
- Untuk loop tak terbatas, gunakan `for { ... }` lalu hentikan dengan `break`.
------------------------------------------------------------------------
### 20. Function di Go

Function adalah blok kode yang bisa dipanggil berulang kali untuk menjalankan tugas tertentu. Go mendukung function dengan parameter, return value, multiple return, dan named return value.

---

### 1. Function Sederhana
```go
func sayHello() {
    fmt.Println("Hello World")
}
```
Pemanggilan:
```go
sayHello()
```

---

### 2. Function dengan Parameter
```go
func multiply(variable1 int, variable2 int) {
    result := variable1 * variable2
    fmt.Println(result)
}
```
Pemanggilan:
```go
multiply(4, 2) // Output: 8
```

---

### 3. Function dengan Return Value
```go
func divide(variable1 int, variable2 int) int {
    result := variable1 / variable2
    return result
}
```
Pemanggilan:
```go
result := divide(4, 2)
fmt.Println(result) // Output: 2
```

---

### 4. Function dengan Multiple Return
Go memungkinkan function mengembalikan lebih dari satu nilai.
```go
func oddEvenNumber(variable1 int, variable2 int) ([]int, []int) {
    var oddNumber, evenNumber []int

    if variable1%2 == 0 {
        evenNumber = append(evenNumber, variable1)
    } else {
        oddNumber = append(oddNumber, variable2)
    }

    if variable2%2 == 0 {
        evenNumber = append(evenNumber, variable2)
    } else {
        oddNumber = append(oddNumber, variable1)
    }

    return oddNumber, evenNumber
}
```
Pemanggilan:
```go
odd, even := oddEvenNumber(4, 2)
fmt.Println("Ganjil", odd)
fmt.Println("Genap", even)

// Ignore return value dengan `_`
oddOnly, _ := oddEvenNumber(9, 2)
fmt.Println(oddOnly)
```

---

### 5. Named Return Value
Function bisa menggunakan **named return**. Nilai akan otomatis dikembalikan sesuai nama variabel return.
```go
func calculate(x, y int) (sum, diff int) {
    sum = x + y
    diff = x - y
    return // tidak perlu tulis variabel lagi
}
```
Pemanggilan:
```go
s, d := calculate(10, 5)
fmt.Println("Sum:", s, "Diff:", d)
```

---

### 6. Variadic Function
Variadic function menerima jumlah parameter yang **bervariasi**.
```go
func multipleAll(numbers ...int) int {
total := 1
for _, number := range numbers {
total *= number
}
return total
}
```
Pemanggilan:
```go
// Variadic function langsung
fmt.Println(multipleAll(2, 4, 1, 5, 23, 28, 2, 9, 50))


// Slice sebagai parameter (gunakan ...)
numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
fmt.Println(multipleAll(numbers...))
```
**Output:**
```
8280000
362880
```


---


### 7. Function Value
Function di Go adalah **first-class citizen**, artinya function bisa disimpan di variabel, dijadikan parameter, atau bahkan return value.


```go
func sum(numbers ...int) int {
total := 0
for _, number := range numbers {
total += number
}
return total
}


func main() {
// Function disimpan dalam variabel
sumAll := sum
fmt.Println(sumAll(2, 4, 6, 23, 5))
}
```


**Output:**
```
40
```


---


### Insight
- Function di Go mendukung return lebih dari satu nilai, hal ini memudahkan error handling (`value, err`).
- Bisa mengabaikan return value dengan `_`.
- Named return membuat kode lebih jelas, tapi jangan terlalu banyak digunakan agar tidak membingungkan.
- Variadic function sangat berguna untuk operasi dengan parameter dinamis, dan bisa menerima slice menggunakan `...`.
- Function bisa disimpan ke variabel (**function value**) → memungkinkan membuat callback, handler, atau pola functional programming.

