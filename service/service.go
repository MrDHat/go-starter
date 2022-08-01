package service

type Services interface {
}

type services struct {
}

func Init() Services {
	return &services{}
}
