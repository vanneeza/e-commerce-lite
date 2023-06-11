package entity

type Stock struct {
	Id      string
	Stock   int
	Product Product
	Store   Store
}
