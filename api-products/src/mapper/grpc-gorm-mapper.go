package mapper

import (
	"products/src/model"
	"products/src/pb/products"
)

func ConvertGORMToProto(gormProduct *model.Product) *products.Product {
	return &products.Product{
		Id:          int32(gormProduct.ID),
		Name:        gormProduct.Name,
		Description: gormProduct.Description,
		Price:       gormProduct.Price,
		Quantity:    gormProduct.Quantity,
	}
}

func ConvertProtoToGORM(protoProduct *products.Product) *model.Product {
	product := &model.Product{
		Name:        protoProduct.Name,
		Description: protoProduct.Description,
		Price:       protoProduct.Price,
		Quantity:    protoProduct.Quantity,
	}

	var produtoNaoExiste bool = protoProduct.Id > 0
	if produtoNaoExiste {
		product.ID = uint(protoProduct.Id)
	}

	return product
}

func ConvertGORMListToProto(gormProducts []model.Product) *products.ProductList {
	protoProducts := make([]*products.Product, len(gormProducts))
	for i, product := range gormProducts {
		protoProducts[i] = ConvertGORMToProto(&product)
	}

	return &products.ProductList{
		Products: protoProducts,
	}
}
