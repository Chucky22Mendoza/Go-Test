package repositories

import (
	"github.com/Chucky22Mendoza/Rest-api/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Task struct {
	DB *gorm.DB
}

func (r *Task) GetAll(c *fiber.Ctx) error {
	tasks := []models.Task{}
	err := r.DB.Find(&tasks).Error

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"error": "Get all tasks. " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": "Data tasks loading successfully",
		"data":    tasks,
	})
}

func (r *Task) Get(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"error": "missing parameter :id",
		})
	}
	task := &models.Task{
		Id: id,
	}
	err := r.DB.First(&task).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}
	if task.Id == "" {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"error": "Task not found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": "Task found",
		"data":    task,
	})
}

func (r *Task) Create(c *fiber.Ctx) error {
	task := models.Task{}
	if err := c.BodyParser(&task); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}
	err := r.DB.Create(&task).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"error": "Task not created. " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": "Task found",
		"task_id": task.Id,
	})
}

func (r *Task) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"error": "missing parameter :id",
		})
	}

	task := &models.Task{
		Id: id,
	}
	body := &models.Task{}
	err := c.BodyParser(body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}

	err = r.DB.First(task).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}

	if task.Id == "" {
		return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"message": "Task not found",
		})
	}

	if body.Name != "" {
		task.Name = body.Name
	}
	if body.Estatus != "" {
		task.Estatus = body.Estatus
	}
	err = r.DB.Save(task).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": "The media has been updated succesfully",
		"data":    task,
	})
}

func (r *Task) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"error": "missing parameter :id",
		})
	}
	media := &models.Task{
		Id: id,
	}
	err := r.DB.Delete(media).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": "The task has been deleted succesfully",
	})
}

func (r *Task) SetUpRoutes(router fiber.Router) {
	tasks := router.Group("/tasks")
	tasks.Get("/", r.GetAll)
	tasks.Get("/get/:id", r.Get)
	tasks.Post("/create", r.Create)
	tasks.Put("/update/:id", r.Update)
	tasks.Delete("/delete/:id", r.Delete)
}
