package main

import (
	"school/infrastructure"
	"school/interfaces/controllers"
)

func main() {
	teacherRepository := infrastructure.NewRealTeacherRepository()
	classRoomRepository := infrastructure.NewRealClassRoomRepository()
	lessonRepository := infrastructure.NewRealLessonRepository()
	pupilRepository := infrastructure.NewRealPupilRepository()
	cliController := controllers.NewCLIController(
		teacherRepository,
		lessonRepository,
		pupilRepository,
		classRoomRepository,
	)
	cliController.Run()
}
