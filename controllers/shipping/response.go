package shipping

type PostShipType struct {
	Name string `json:"name" form:"name"`
}

type PostShipping struct {
	Name        string `json:"name" form:"name"`
	Cost        int    `json:"cost" form:"cost"`
	ShipType_ID int    `json:"shiptype_id" form:"shiptype_id"`
}
type EditShipping struct {
	Name        string `json:"name" form:"name"`
	Cost        int    `json:"cost" form:"cost"`
	ShipType_ID int    `json:"shiptype_id" form:"shiptype_id"`
}
