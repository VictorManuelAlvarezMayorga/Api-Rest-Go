package models

type Car struct {
	ID       int    `json:"id"`
	Brand    string `json:"brand"`
	Model    string `json:"model"`
	Version  string `json:"version"`
	Year     int    `json:"year"`
	ImageURL string `json:"image_url"`
}

var Cars = []Car{
	{ID: 1, Brand: "Chevrolet", Model: "Camaro", Version: "RS", Year: 1967, ImageURL: "https://pacificclassics.com/inventory/1967-chevrolet-camaro-rs-orange/"},
	{ID: 2, Brand: "Ford", Model: "Mustang", Version: "Boss 429", Year: 1969, ImageURL: "https://espirituracer.com/cochedeldia/coche-del-dia-ford-mustang-boss-429-by-classics-recreations/"},
	{ID: 3, Brand: "Mazda", Model: "RX7", Version: "Coupé", Year: 1977, ImageURL: "https://soymotor.com/coches/articulos/visita-museo-mazda-prueba-rx-7-950980"},
	{ID: 4, Brand: "Dodge", Model: "Charger", Version: "Daytona", Year: 1970, ImageURL: "https://rmsothebys.com/auctions/st12/lots/r159-1970-dodge-charger-rt-daytona-hardtop-coupe-recreation/"},
}
