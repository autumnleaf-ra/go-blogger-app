package controller

import (
	"fmt"
	"math"
	"strconv"

	"github.com/autumnleaf-ra/go-blogger-app/database"
	"github.com/autumnleaf-ra/go-blogger-app/models"
	"github.com/autumnleaf-ra/go-blogger-app/util"
	"github.com/gofiber/fiber/v2"
)

func CreatePost(c *fiber.Ctx) error {
	var blogpost models.Blog
	if err := c.BodyParser(&blogpost); err != nil {
		fmt.Println("Unable to parse body")
	}
	if err := database.DB.Create(&blogpost).Error; err != nil {
		c.Status(401)
		return c.JSON(fiber.Map{
			"message": "Invalid Payload",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Congratulation, post successfully",
	})
}

func AllPost(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit := 5
	offset := (page - 1) * limit
	var total int64
	var getblog []models.Blog
	database.DB.Preload("User").Offset(offset).Limit(limit).Find(&getblog)
	database.DB.Model(&models.Blog{}).Count(&total)
	return c.JSON(fiber.Map{
		"data": getblog,
		"meta": fiber.Map{
			"total":     total,
			"page":      page,
			"last_page": math.Ceil(float64(int(total) / limit)),
		},
	})
}

func DetailPost(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var blogpost models.Blog
	database.DB.Where("id = ?", id).Preload("User").First(&blogpost)
	return c.JSON(fiber.Map{
		"data": blogpost,
	})
}

func UpdatePost(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	blog := models.Blog{
		Id: uint(id),
	}

	if err := c.BodyParser(&blog); err != nil {
		fmt.Println("Unable to parse body")
	}

	database.DB.Model(&blog).Updates(blog)
	return c.JSON(fiber.Map{
		"message": "Post updated successfully",
	})
}

func UniquePost(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	id, _ := util.ParseJWT(cookie)
	var blog []models.Blog
	database.DB.Model(&blog).Where("user_id = ?", id).Preload("User").Find(&blog)
	return c.JSON(blog)
}

func DeletePost(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	blog := models.Blog{
		Id: uint(id),
	}

	deleteQuery := database.DB.Delete(&blog)
	if deleteQuery.RowsAffected == 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Post not found!",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Post deleted successfully",
	})
}
