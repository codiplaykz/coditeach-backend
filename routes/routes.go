package routes

import (
	"coditeach/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	api := app.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			auth := v1.Group("/auth")
			{
				//Authorization routes
				auth.Post("/sign_up", controllers.SignUp)
				auth.Post("/sign_in", controllers.SignIn)
				auth.Post("/refresh", controllers.Refresh)
			}

			curriculum := v1.Group("/curriculum", controllers.UserIdentifyMiddleware)
			{
				//Curriculum routes
				curriculum.Get("/get", controllers.GetCurriculum)
				curriculum.Post("/create", controllers.CreateCurriculum)
				curriculum.Put("/update", controllers.UpdateCurriculum)
				curriculum.Delete("/delete", controllers.DeleteCurriculum)

				//Curriculum lessons routes
				app.Get("/lessons/get", controllers.GetCurriculumLesson)
				app.Post("/lessons/create", controllers.CreateCurriculumLesson)
				app.Put("/lessons/update", controllers.UpdateCurriculumLesson)
				app.Delete("/lessons/delete", controllers.DeleteCurriculumLesson)
			}

			module := v1.Group("/module", controllers.UserIdentifyMiddleware)
			{
				//Module routes
				module.Get("/get", controllers.GetModule)
				module.Post("/create", controllers.CreateModule)
				module.Put("/update", controllers.UpdateModule)
				module.Delete("/delete", controllers.DeleteModule)
			}

			block := v1.Group("/block", controllers.UserIdentifyMiddleware)
			{
				//Block routes
				block.Get("/get", controllers.GetBlock)
				block.Post("/create", controllers.CreateBlock)
				block.Put("/update", controllers.UpdateBlock)
				block.Delete("/delete", controllers.DeleteBlock)
			}

			schools := v1.Group("/school", controllers.UserIdentifyMiddleware)
			{
				//School routes
				schools.Post("/create", controllers.CreateSchool)
				schools.Get("/get", controllers.GetSchool)
				schools.Put("/update", controllers.UpdateSchool)
				schools.Delete("/delete", controllers.DeleteSchool)
			}

			teachers := v1.Group("/teachers", controllers.UserIdentifyMiddleware)
			{
				//Teacher routes
				teachers.Post("/create", controllers.CreateTeacher)
				teachers.Post("/create_account", controllers.CreateTeacherAccount)
				teachers.Get("/get", controllers.GetTeacher)
				teachers.Put("/update", controllers.UpdateTeacher)
				teachers.Delete("/delete", controllers.DeleteTeacher)
			}

			classes := v1.Group("/classes", controllers.UserIdentifyMiddleware)
			{
				//Class routes
				classes.Post("/create", controllers.CreateClass)
				classes.Get("/get", controllers.GetClass)
				classes.Get("/get_by_code", controllers.GetClassByCode)
				classes.Put("/update", controllers.UpdateClass)
				classes.Delete("/delete", controllers.DeleteClass)
			}

			students := v1.Group("/students", controllers.UserIdentifyMiddleware)
			{
				//Student routes
				students.Post("/create", controllers.CreateStudent)
				students.Post("/create_account", controllers.CreateStudentAccount)
				students.Post("/register_account", controllers.RegisterStudentAccount)
				students.Get("/get", controllers.GetStudent)
				students.Put("/update", controllers.UpdateStudent)
				students.Delete("/delete", controllers.DeleteStudent)
			}

			parents := v1.Group("/parents", controllers.UserIdentifyMiddleware)
			{
				//Parent routes
				parents.Post("api/parent/create", controllers.CreateParent)
				parents.Get("api/parent/get", controllers.GetParent)
				parents.Put("api/parent/update", controllers.UpdateParent)
				parents.Delete("api/parent/delete", controllers.DeleteParent)
			}

			subjects := v1.Group("", controllers.UserIdentifyMiddleware)
			{
				//Subject routes
				subjects.Post("/create", controllers.CreateSubject)
				subjects.Get("/get", controllers.GetSubject)
				subjects.Put("/update", controllers.UpdateSubject)
				subjects.Delete("/delete", controllers.DeleteSubject)
			}

			homeworks := v1.Group("/homeworks", controllers.UserIdentifyMiddleware)
			{
				//Homework routes
				homeworks.Post("/create", controllers.CreateHomework)
				homeworks.Get("/get", controllers.GetHomework)
				homeworks.Put("/update", controllers.UpdateHomework)
				homeworks.Delete("/delete", controllers.DeleteHomework)
			}

			schedule := v1.Group("/schedule", controllers.UserIdentifyMiddleware)
			{
				//Schedule routes
				schedule.Post("/create", controllers.CreateSchedule)
				schedule.Get("/get", controllers.GetSchedule)
				schedule.Put("/update", controllers.UpdateSchedule)
				schedule.Delete("/delete", controllers.DeleteSchedule)
			}

			schedule_lesson := v1.Group("/schedule_lesson", controllers.UserIdentifyMiddleware)
			{
				//Schedule lesson routes
				schedule_lesson.Post("/create", controllers.CreateScheduleLesson)
				schedule_lesson.Get("/get", controllers.GetScheduleLesson)
				schedule_lesson.Put("/update", controllers.UpdateScheduleLesson)
				schedule_lesson.Delete("/delete", controllers.DeleteScheduleLesson)
			}

			tests := v1.Group("/tests", controllers.UserIdentifyMiddleware)
			{
				//Test routes
				tests.Post("/create", controllers.CreateTest)
				tests.Get("/get", controllers.GetTest)
				tests.Put("/update", controllers.UpdateTest)
				tests.Delete("/delete", controllers.DeleteTest)
			}

			questions := v1.Group("/", controllers.UserIdentifyMiddleware)
			{
				//Question routes
				questions.Post("/create", controllers.CreateQuestion)
				questions.Get("/get", controllers.GetQuestion)
				questions.Put("/update", controllers.UpdateQuestion)
				questions.Delete("/delete", controllers.DeleteQuestion)
			}

			options := v1.Group("/options", controllers.UserIdentifyMiddleware)
			{
				//Option routes
				options.Post("/create", controllers.CreateOption)
				options.Get("/get", controllers.GetOption)
				options.Put("/update", controllers.UpdateOption)
				options.Delete("/delete", controllers.DeleteOption)
			}

			events := v1.Group("/events", controllers.UserIdentifyMiddleware)
			{
				//Event routes
				events.Post("/create", controllers.CreateEvent)
				events.Get("/get", controllers.GetEvent)
				events.Put("/update", controllers.UpdateEvent)
				events.Delete("/delete", controllers.DeleteEvent)
			}

			test_results := v1.Group("/test_results", controllers.UserIdentifyMiddleware)
			{
				//Test result routes
				test_results.Post("/create", controllers.CreateTestResult)
				test_results.Get("/get", controllers.GetTestResult)
				test_results.Put("/update", controllers.UpdateTestResult)
				test_results.Delete("/delete", controllers.DeleteTestResult)
			}

			statistics := v1.Group("/statistics", controllers.UserIdentifyMiddleware)
			{
				//Statistics
				statistics.Get("/get_total", controllers.GenerateTotalStatistics)
				statistics.Get("/get_by_school", controllers.GenerateSchoolStatistics)
				statistics.Get("/get", controllers.GenerateReport)
			}
		}
	}
}
