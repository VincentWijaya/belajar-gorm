package main

import (
	"belajar-gorm/database"
	"belajar-gorm/models"
	"fmt"
)

func main() {
	database.StartDB()
	// updateUserByUserID(3, "admin@mail.com")
	// getUserByUserID(1)
	// getUserByUserID(3)
	createProduct(1, "Kecap Manis", "Bangau")
	// getUsersWithProducts()
	// deleteProductByProductID(1)
	// getUsersWithProducts()

}

func createUser(email string) {
	db := database.GetDB()
	user := models.User{
		Email: email,
	}

	if err := db.Create(&user).Error; err != nil {
		panic(err)
	}

	fmt.Printf("Successfully save user data: %+v\n", user)
}

func getUserByUserID(userID uint) {
	db := database.GetDB()
	user := models.User{}

	if err := db.First(&user, "id=?", userID).Error; err != nil {
		panic(err)
	}

	fmt.Printf("User data: %+v\n", user)
}

func updateUserByUserID(id uint, email string) {
	db := database.GetDB()
	user := models.User{}

	err := db.Model(&user).Where("id = ?", id).Update("email", email).Error
	if err != nil {
		panic(err)
	}

	fmt.Printf("Success to update user data: %+v\n", user)
}

func createProduct(userID uint, name, brand string) {
	db := database.GetDB()
	product := models.Product{
		UserID: userID,
		Name:   name,
		Brand:  brand,
	}

	if err := db.Create(&product).Error; err != nil {
		panic(err)
	}

	fmt.Printf("Success to save product: %+v\n", product)
}

func getUsersWithProducts() {
	db := database.GetDB()
	user := models.User{}

	err := db.Preload("Products").Find(&user).Error
	if err != nil {
		panic(err)
	}

	fmt.Printf("Users with products: %+v\n", user)
}

func deleteProductByProductID(productID uint) {
	db := database.GetDB()
	product := models.Product{}

	err := db.Where("id=?", productID).Delete(&product).Error
	if err != nil {
		panic(err)
	}

	fmt.Println("Success to detele product with ID: ", productID)
}
