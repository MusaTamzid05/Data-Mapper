package lib

// bride pattern

type Table interface {
	Create(map[string]interface{})
	Retrive(id int)
	Update(map[string]interface{})
	Delete(id int)
}
