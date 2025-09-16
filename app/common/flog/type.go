package flog

type DataType struct {
	Msg   string `json:"msg" bson:"msg"`
	Track any    `json:"track" bson:"track"`
}
