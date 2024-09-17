package controller

import (
	"fmt"
	"net/http"
	"git.intelbras.com.br/isec/linha-future/jovens/gestoredu.git/api/entities"
	_ "git.intelbras.com.br/isec/linha-future/jovens/gestoredu.git/docs"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UsersController struct {
	db *gorm.DB
}

func NewUsersController(db *gorm.DB) *UsersController {
	return &UsersController{db: db}
}

// ShowAccount godoc
// @Summary      Criar usuários
// @Description  Faz a criação de usuários
// @Tags         Usuários
// @Accept       json
// @Produce      json
// @Param        Body   body      entities.UserRequest  true  "dados do usuario"
// @Success      200  {object}  entities.Users
// @Failure      400 "error"
// @Router       /user [post]
func (uc *UsersController) CreateUser(c *fiber.Ctx) error {
	var userRequest entities.UserRequest

	if err := c.BodyParser(&userRequest); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse body",
		})
	}

	user := &entities.Users{
		ID:             uuid.New().String(),
		Nome:           userRequest.Nome,
		Sobrenome:      userRequest.Sobrenome,
		Telefone:       userRequest.Telefone,
		Email:          userRequest.Email,
		Sexo:           userRequest.Sexo,
		Idade:          userRequest.Idade,
		CPF:            userRequest.CPF,
		Endereco:       userRequest.Endereco,
		DataNascimento: userRequest.DataNascimento,
		Matricula:      userRequest.Matricula,
	}
	if err:= uc.db.Save(&user).Error; err != nil{
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "erro ao salvar no db",
		})
	}
	return c.Status(http.StatusCreated).JSON(&user)
}

// ShowAccount godoc
// @Summary      Buscar usuários
// @Description  Faz a busca de usuários por ID
// @Tags         Usuários
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "ID do Usuário"
// @Success      200  {object}  entities.Users "Usuário encontrado com sucesso"
// @Failure      404  {object}  fiber.Map  "Usuário não encontrado"
// @Router       /user/{id} [get]
func (uc *UsersController) GetUser(c *fiber.Ctx) error {
	var user entities.Users
	id := c.Params("id")
	err := uc.db.Model(&entities.Users{}).Where("id = ?", id).First(&user).Error
	if err != nil {
		fmt.Println(err.Error())
	}

	if user.ID != id {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"erro": "usuario não encontrado",
		})
	}

	return c.JSON(&user)
}

// ShowAccount godoc
// @Summary      Atualizar usuários
// @Description  Atualiza dados do usuário por ID
// @Tags         Usuários
// @Accept       json
// @Produce      json
// @Param        id   body      string  true  "ID"
// @Success      200  {object}  entities.Users
// @Router       /user/{id} [patch]
func (uc *UsersController) UpdateUser(c *fiber.Ctx) error {
	var user entities.Users
	id := c.Params("id")

	err := uc.db.Model(&entities.Users{}).Where("id = ?", id).First(&user).Error
	if err != nil {
		fmt.Println(err.Error())
	}

	if user.ID != id {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"erro": "usuario não encontrado",
		})
	}
	if err := c.BodyParser(&user); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	uc.db.Save(&user)
	return c.JSON(&user)
}

// ShowAccount godoc
// @Summary      Deletar usuários
// @Description  Deleta usuário por ID
// @Tags         Usuários
// @Accept       json
// @Produce      json
// @Param        id   body      string  true  "ID"
// @Success      200  {object}  entities.Users
// @Router       /user/{id} [delete]
func (uc *UsersController) DeleteUser(c *fiber.Ctx) error {
	var user entities.Users
	id := c.Params("id")
	err := uc.db.Model(&entities.Users{}).Where("id = ?", id).First(&user).Error
	if err != nil {
		fmt.Println(err.Error())
	}
	if user.ID != id {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"erro": "usuario não encontrado",
		})
	}

	// if err := uc.db.Delete(&user).Error ; err != nil {
	// 	return c.Status(http.StatusNotFound).JSON(fiber.Map{
	// 		"erro": err.Error(),
	// 	})
	// }
	uc.db.Delete(&user)
	return c.SendString("usuario deletado")
}
