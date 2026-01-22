package model

type Produk struct {
	ID    int    `sql:"id" json:"id"`
	Nama  string `sql:"nama" json:"nama"`
	Harga int    `sql:"harga" json:"harga"`
	Stok  int    `sql:"stok" json:"stok"`
}
