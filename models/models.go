package models


var UserTable = struct {
	TableName string
	firstname string
	surname string
	password1 string
	email string
	 
}{
	TableName:"UserSQLModel",
	firstname: "firstname",
	surname: "surname",
	password1: "paswoord1",
	email: "email",


}




var BuildingTable = struct {
	TableName string
	gid string
	shape__are string
	shape__len string
	globalid string
	creationda string
	creator string
	editdate string
	editor string
	num_storey string
	building_t string
	ghanapost_ string
	plot_numbe string
	developmen string
	name string
	parcel_id string
	exact_use string
	building_u string
	remarks string
	other_info string
	other_in_1 string
	geom string

}{
	TableName: "building",
	gid :"gid",
	shape__are: "shape__are",
	shape__len: "shape__len",
	globalid :"globalid",
	creationda: "creationda",
	creator: "creator",
	editdate: "editdate",
	editor : "editor",
	num_storey : "num_storey",
	building_t : "building_t",
	ghanapost_ :"ghanapost_",
	plot_numbe :"plot_numbe",
	developmen :"developmen",
	name : "name",
	parcel_id : "parcel_id",
	exact_use : "exact_use",
	building_u: "building_u",
	remarks: "remarks",
	other_info: "other_info",
	other_in_1 : "other_in_1",
	geom : "geom",

}



var OtherPolygonTable = struct {
	TableName string
	gid string
	shape__are string
	shape__len string
	globalid string
	creationda string
	creator string
	editdate string
	editor string
    usage1 string
	ghanapostg string
	developmen string
	name string
	structure_ string
	parcel_id string
	exact_use string
	remarks string
	other_info string
	other_in_1 string
	geom string
	street_nam string
	mixed_usag string

}{
	TableName: "other_polygon_structure",
	gid :"gid",
	shape__are: "shape__are",
	shape__len: "shape__len",
	globalid :"globalid",
	creationda: "creationda",
	creator: "creator",
	editdate: "editdate",
	editor : "editor",
	usage1: "usage1",
	structure_:"structure_",
	ghanapostg :"ghanapostg",
	developmen :"developmen",
	name : "name",
	parcel_id : "parcel_id",
	exact_use : "exact_use",
	remarks: "remarks",
	other_info: "other_info",
	other_in_1 : "other_in_1",
	street_nam : "street_nam",
	geom : "geom",
	mixed_usag  : "mixed_usag",

}



