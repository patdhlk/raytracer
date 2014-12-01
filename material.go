package main

type Material interface {
	getKa() float64
	getKd() float64
	getKs() float64
}
