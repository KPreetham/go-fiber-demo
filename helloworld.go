package main

import (
	"fmt"
	s "helloworld/sample"
	"log"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/google/uuid"
)

func main() {
	app := fiber.New(fiber.Config{
		BodyLimit: 20 * 1024 * 1024,
	})

	app.Use(logger.New())

	app.Get("/err", func(c *fiber.Ctx) error {
		return c.SendString(s.HelloWorld())
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("GET request")
	})

	app.Post("/upload", func(c *fiber.Ctx) error {
		// Get first file from form field "document":
		file, err := c.FormFile("file")
		if err != nil {
			return err
		}
		// Save file to root directory:
		c.SaveFile(file, fmt.Sprintf("./%s-%s", uuid.New(), file.Filename))

		return c.JSON(&fiber.Map{
			"success": true,
			"message": "file uploaded",
		})
	})

	// Start server

	app.Get("/ws", websocket.New(func(c *websocket.Conn) {
		log.Printf("connection received")
		// Websocket logic
		for {
			// mtype, msg, err := c.ReadMessage()
			// if err != nil {
			// 	break
			// }
			// log.Printf("Read: %s", msg)

			err := c.WriteMessage(1, []byte("hello"))
			if err != nil {
				break
			}
		}
	}))

	app.Get("/1", func(c *fiber.Ctx) error {
		return c.SendString(func_1())
	})

	app.Get("/2", func(c *fiber.Ctx) error {
		return c.SendString(func_2())
	})

	app.Get("/3", func(c *fiber.Ctx) error {
		return c.SendString(func_3())
	})

	app.Get("/4", func(c *fiber.Ctx) error {
		return c.SendString(func_4())
	})

	app.Get("/5", func(c *fiber.Ctx) error {
		return c.SendString(func_5())
	})

	app.Get("/6", func(c *fiber.Ctx) error {
		return c.SendString(func_6())
	})

	app.Get("/7", func(c *fiber.Ctx) error {
		return c.SendString(func_7())
	})

	app.Get("/8", func(c *fiber.Ctx) error {
		return c.SendString(func_8())
	})

	app.Get("/9", func(c *fiber.Ctx) error {
		return c.SendString(func_9())
	})

	app.Get("/10", func(c *fiber.Ctx) error {
		return c.SendString(func_10())
	})

	app.Get("/11", func(c *fiber.Ctx) error {
		return c.SendString(func_11())
	})

	app.Get("/12", func(c *fiber.Ctx) error {
		return c.SendString(func_12())
	})

	app.Get("/13", func(c *fiber.Ctx) error {
		return c.SendString(func_13())
	})

	app.Get("/14", func(c *fiber.Ctx) error {
		return c.SendString(func_14())
	})

	app.Get("/15", func(c *fiber.Ctx) error {
		return c.SendString(func_15())
	})

	app.Get("/16", func(c *fiber.Ctx) error {
		return c.SendString(func_16())
	})

	app.Get("/17", func(c *fiber.Ctx) error {
		return c.SendString(func_17())
	})

	app.Get("/18", func(c *fiber.Ctx) error {
		return c.SendString(func_18())
	})

	app.Get("/19", func(c *fiber.Ctx) error {
		return c.SendString(func_19())
	})

	app.Get("/20", func(c *fiber.Ctx) error {
		return c.SendString(func_20())
	})

	app.Get("/21", func(c *fiber.Ctx) error {
		return c.SendString(func_21())
	})

	app.Get("/22", func(c *fiber.Ctx) error {
		return c.SendString(func_22())
	})

	app.Get("/23", func(c *fiber.Ctx) error {
		return c.SendString(func_23())
	})

	app.Get("/24", func(c *fiber.Ctx) error {
		return c.SendString(func_24())
	})

	app.Get("/25", func(c *fiber.Ctx) error {
		return c.SendString(func_25())
	})

	app.Get("/26", func(c *fiber.Ctx) error {
		return c.SendString(func_26())
	})

	app.Get("/27", func(c *fiber.Ctx) error {
		return c.SendString(func_27())
	})

	app.Get("/28", func(c *fiber.Ctx) error {
		return c.SendString(func_28())
	})

	app.Get("/29", func(c *fiber.Ctx) error {
		return c.SendString(func_29())
	})

	app.Get("/30", func(c *fiber.Ctx) error {
		return c.SendString(func_30())
	})

	app.Get("/31", func(c *fiber.Ctx) error {
		return c.SendString(func_31())
	})

	app.Get("/32", func(c *fiber.Ctx) error {
		return c.SendString(func_32())
	})

	app.Get("/33", func(c *fiber.Ctx) error {
		return c.SendString(func_33())
	})

	app.Get("/34", func(c *fiber.Ctx) error {
		return c.SendString(func_34())
	})

	app.Get("/35", func(c *fiber.Ctx) error {
		return c.SendString(func_35())
	})

	app.Get("/36", func(c *fiber.Ctx) error {
		return c.SendString(func_36())
	})

	app.Get("/37", func(c *fiber.Ctx) error {
		return c.SendString(func_37())
	})

	app.Get("/38", func(c *fiber.Ctx) error {
		return c.SendString(func_38())
	})

	app.Get("/39", func(c *fiber.Ctx) error {
		return c.SendString(func_39())
	})

	app.Get("/40", func(c *fiber.Ctx) error {
		return c.SendString(func_40())
	})

	app.Get("/41", func(c *fiber.Ctx) error {
		return c.SendString(func_41())
	})

	app.Get("/42", func(c *fiber.Ctx) error {
		return c.SendString(func_42())
	})

	app.Get("/43", func(c *fiber.Ctx) error {
		return c.SendString(func_43())
	})

	app.Get("/44", func(c *fiber.Ctx) error {
		return c.SendString(func_44())
	})

	app.Get("/45", func(c *fiber.Ctx) error {
		return c.SendString(func_45())
	})

	app.Get("/46", func(c *fiber.Ctx) error {
		return c.SendString(func_46())
	})

	app.Get("/47", func(c *fiber.Ctx) error {
		return c.SendString(func_47())
	})

	app.Get("/48", func(c *fiber.Ctx) error {
		return c.SendString(func_48())
	})

	app.Get("/49", func(c *fiber.Ctx) error {
		return c.SendString(func_49())
	})

	app.Get("/50", func(c *fiber.Ctx) error {
		return c.SendString(func_50())
	})

	app.Get("/51", func(c *fiber.Ctx) error {
		return c.SendString(func_51())
	})

	app.Get("/52", func(c *fiber.Ctx) error {
		return c.SendString(func_52())
	})

	app.Get("/53", func(c *fiber.Ctx) error {
		return c.SendString(func_53())
	})

	app.Get("/54", func(c *fiber.Ctx) error {
		return c.SendString(func_54())
	})

	app.Get("/55", func(c *fiber.Ctx) error {
		return c.SendString(func_55())
	})

	app.Get("/56", func(c *fiber.Ctx) error {
		return c.SendString(func_56())
	})

	app.Get("/57", func(c *fiber.Ctx) error {
		return c.SendString(func_57())
	})

	app.Get("/58", func(c *fiber.Ctx) error {
		return c.SendString(func_58())
	})

	app.Get("/59", func(c *fiber.Ctx) error {
		return c.SendString(func_59())
	})

	app.Get("/60", func(c *fiber.Ctx) error {
		return c.SendString(func_60())
	})

	app.Get("/61", func(c *fiber.Ctx) error {
		return c.SendString(func_61())
	})

	app.Get("/62", func(c *fiber.Ctx) error {
		return c.SendString(func_62())
	})

	app.Get("/63", func(c *fiber.Ctx) error {
		return c.SendString(func_63())
	})

	app.Get("/64", func(c *fiber.Ctx) error {
		return c.SendString(func_64())
	})

	app.Get("/65", func(c *fiber.Ctx) error {
		return c.SendString(func_65())
	})

	app.Get("/66", func(c *fiber.Ctx) error {
		return c.SendString(func_66())
	})

	app.Get("/67", func(c *fiber.Ctx) error {
		return c.SendString(func_67())
	})

	app.Get("/68", func(c *fiber.Ctx) error {
		return c.SendString(func_68())
	})

	app.Get("/69", func(c *fiber.Ctx) error {
		return c.SendString(func_69())
	})

	app.Get("/70", func(c *fiber.Ctx) error {
		return c.SendString(func_70())
	})

	app.Get("/71", func(c *fiber.Ctx) error {
		return c.SendString(func_71())
	})

	app.Get("/72", func(c *fiber.Ctx) error {
		return c.SendString(func_72())
	})

	app.Get("/73", func(c *fiber.Ctx) error {
		return c.SendString(func_73())
	})

	app.Get("/74", func(c *fiber.Ctx) error {
		return c.SendString(func_74())
	})

	app.Get("/75", func(c *fiber.Ctx) error {
		return c.SendString(func_75())
	})

	app.Get("/76", func(c *fiber.Ctx) error {
		return c.SendString(func_76())
	})

	app.Get("/77", func(c *fiber.Ctx) error {
		return c.SendString(func_77())
	})

	app.Get("/78", func(c *fiber.Ctx) error {
		return c.SendString(func_78())
	})

	app.Get("/79", func(c *fiber.Ctx) error {
		return c.SendString(func_79())
	})

	app.Get("/80", func(c *fiber.Ctx) error {
		return c.SendString(func_80())
	})

	app.Get("/81", func(c *fiber.Ctx) error {
		return c.SendString(func_81())
	})

	app.Get("/82", func(c *fiber.Ctx) error {
		return c.SendString(func_82())
	})

	app.Get("/83", func(c *fiber.Ctx) error {
		return c.SendString(func_83())
	})

	app.Get("/84", func(c *fiber.Ctx) error {
		return c.SendString(func_84())
	})

	app.Get("/85", func(c *fiber.Ctx) error {
		return c.SendString(func_85())
	})

	app.Get("/86", func(c *fiber.Ctx) error {
		return c.SendString(func_86())
	})

	app.Get("/87", func(c *fiber.Ctx) error {
		return c.SendString(func_87())
	})

	app.Get("/88", func(c *fiber.Ctx) error {
		return c.SendString(func_88())
	})

	app.Get("/89", func(c *fiber.Ctx) error {
		return c.SendString(func_89())
	})

	app.Get("/90", func(c *fiber.Ctx) error {
		return c.SendString(func_90())
	})

	app.Get("/91", func(c *fiber.Ctx) error {
		return c.SendString(func_91())
	})

	app.Get("/92", func(c *fiber.Ctx) error {
		return c.SendString(func_92())
	})

	app.Get("/93", func(c *fiber.Ctx) error {
		return c.SendString(func_93())
	})

	app.Get("/94", func(c *fiber.Ctx) error {
		return c.SendString(func_94())
	})

	app.Get("/95", func(c *fiber.Ctx) error {
		return c.SendString(func_95())
	})

	app.Get("/96", func(c *fiber.Ctx) error {
		return c.SendString(func_96())
	})

	app.Get("/97", func(c *fiber.Ctx) error {
		return c.SendString(func_97())
	})

	app.Get("/98", func(c *fiber.Ctx) error {
		return c.SendString(func_98())
	})

	app.Get("/99", func(c *fiber.Ctx) error {
		return c.SendString(func_99())
	})

	app.Get("/100", func(c *fiber.Ctx) error {
		return c.SendString(func_100())
	})

	log.Fatal(app.Listen(":3000"))
}
