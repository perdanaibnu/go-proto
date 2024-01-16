package main

import (
	"fmt"
	pb "go-proto/pb"
	"log"

	"google.golang.org/protobuf/proto"
)

func main() {
	products := &pb.Products{
		Data: []*pb.Product{
			{
				Id:    1,
				Name:  "Nike Air Max",
				Price: 1100000.00,
				Stock: 100.00,
				Category: &pb.Category{
					Id:   1,
					Name: "Shoes",
				},
			},
			{
				Id:    2,
				Name:  "Adidas UB",
				Price: 1100000.00,
				Stock: 100.00,
				Category: &pb.Category{
					Id:   2,
					Name: "Shoes",
				},
			},
		},
	}

	data, err := proto.Marshal(products)
	if err != nil {
		log.Fatal("Marshal error", err)
	}

	// compact binary wire format
	fmt.Println(data)

	testProduct := &pb.Products{}
	if err := proto.Unmarshal(data, testProduct); err != nil {
		log.Fatal("Unmarshal", err)
	}

	//fmt.Print(testProduct)

	for _, product := range testProduct.GetData() {
		fmt.Println(product.GetName())
		fmt.Println(product.GetCategory().GetName())
	}
}
