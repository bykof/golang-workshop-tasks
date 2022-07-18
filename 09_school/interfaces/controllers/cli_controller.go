package controllers

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"school/application"
	"school/domain"
	"strconv"
	"strings"
)

type CLIController struct {
	reader *bufio.Reader

	teacherRepository   application.TeacherRepository
	lessonRepository    application.LessonRepository
	pupilRepository     application.PupilRepository
	classRoomRepository application.ClassRoomRepository
}

func NewCLIController(
	teacherRepository application.TeacherRepository,
	lessonRepository application.LessonRepository,
	pupilRepository application.PupilRepository,
	classRoomRepository application.ClassRoomRepository,
) *CLIController {
	return &CLIController{
		reader:              bufio.NewReader(os.Stdin),
		teacherRepository:   teacherRepository,
		lessonRepository:    lessonRepository,
		pupilRepository:     pupilRepository,
		classRoomRepository: classRoomRepository,
	}
}

func (clic *CLIController) readUntil(message string) string {
	for {
		fmt.Print(message)
		value, _ := clic.reader.ReadString('\n')
		value = strings.Trim(value, " \n")
		if value != "" {
			return value
		}
	}
}

func (clic *CLIController) OnCommand(command string) error {
	switch command {
	case AddPupilCommand:
		fmt.Println("Add a pupil:")

		firstName := clic.readUntil("> Enter first name: ")
		lastName := clic.readUntil("> Enter last name: ")
		gradeString := clic.readUntil("> Enter grade: ")
		grade, err := strconv.Atoi(gradeString)
		if err != nil {
			return err
		}
		pupil := clic.pupilRepository.Add(domain.Pupil{
			FirstName: firstName,
			LastName:  lastName,
			Grade:     int64(grade),
		})

		if !pupil.HasId() {
			return errors.New("pupil has no ID, something went wrong")
		}
		fmt.Printf("%s successfully added\n", pupil.String())

	case AddTeacherCommand:
		fmt.Println("Add a teacher:")

		firstName := clic.readUntil("> Enter first name: ")
		lastName := clic.readUntil("> Enter last name: ")
		teachingArea := clic.readUntil("> Enter teaching area: ")

		teacher := clic.teacherRepository.Add(domain.Teacher{
			FirstName:    firstName,
			LastName:     lastName,
			TeachingArea: teachingArea,
		})
		fmt.Printf("%s successfully added\n", teacher.String())
	case AddClassRoomCommand:
		fmt.Println("Add a classroom:")

		roomNumber := clic.readUntil("> Enter room number: ")
		location := clic.readUntil("> Enter location: ")

		classRoom := clic.classRoomRepository.Add(domain.ClassRoom{
			Number:   roomNumber,
			Location: location,
		})
		fmt.Printf("%s successfully added\n", classRoom.String())
	case AddLessonCommand:
		var classRoom domain.ClassRoom
		var teacher domain.Teacher

		fmt.Println("Add a lesson:")

		name := clic.readUntil("> Enter lesson name: ")
		semester := clic.readUntil("> Enter semester: ")

		for {
			classRoomIdString := clic.readUntil("> Enter class room id: ")
			classRoomId, err := strconv.Atoi(classRoomIdString)
			if err != nil {
				return err
			}
			foundClassRoom := clic.classRoomRepository.Find(int64(classRoomId))
			if foundClassRoom != nil {
				classRoom = *foundClassRoom
				break
			} else {
				fmt.Println("Could not find classroom with given id...")
			}
		}

		for {
			teacherIdString := clic.readUntil("> Enter teacher id: ")
			teacherId, err := strconv.Atoi(teacherIdString)
			if err != nil {
				return err
			}
			foundTeacher := clic.teacherRepository.Find(int64(teacherId))
			if foundTeacher != nil {
				teacher = *foundTeacher
				break
			} else {
				fmt.Println("Could not find teacher with given id...")
			}
		}

		lesson := clic.lessonRepository.Add(domain.Lesson{
			Name:      name,
			Semester:  semester,
			ClassRoom: classRoom,
			Teacher:   teacher,
		})
		fmt.Printf("%s successfully added\n", lesson.String())
	case AddPupilToLessonCommand:
		var pupils domain.Pupils
		var lesson domain.Lesson

		pupilIdsString := clic.readUntil("> Enter pupil ids seperated by comma: ")
		for _, pupilIdString := range strings.Split(pupilIdsString, ",") {
			pupilId, err := strconv.Atoi(pupilIdString)
			if err != nil {
				return err
			}
			foundPupil := clic.pupilRepository.Find(int64(pupilId))
			if foundPupil != nil {
				pupils = append(pupils, *foundPupil)

			} else {
				fmt.Println("Could not find pupil with given id...")
			}
		}

		for {
			lessonString := clic.readUntil("> Enter lesson id: ")
			lessonId, err := strconv.Atoi(lessonString)
			if err != nil {
				return err
			}
			foundLesson := clic.lessonRepository.Find(int64(lessonId))
			if foundLesson != nil {
				lesson = *foundLesson
				break
			} else {
				fmt.Println("Could not find lesson with given id...")
			}
		}

		lesson.Pupils = append(lesson.Pupils, pupils...)
		clic.lessonRepository.Update(lesson)
		fmt.Println("Pupils:")
		for _, pupil := range pupils {
			fmt.Println("-", pupil.String())
		}
		fmt.Printf("were added to %s\n", lesson.String())

	case ListAllPupilsCommand:
		pupils := clic.pupilRepository.All()
		fmt.Println("List of pupils:")
		for _, pupil := range pupils {
			fmt.Println(pupil.String())
		}
	case ListAllTeachersCommand:
		teachers := clic.teacherRepository.All()
		fmt.Println("List of teachers:")
		for _, teacher := range teachers {
			fmt.Println(teacher.String())
		}
	case ListAllClassRoomsCommand:
		classRooms := clic.classRoomRepository.All()
		fmt.Println("List of class rooms:")
		for _, classRoom := range classRooms {
			fmt.Println(classRoom.String())
		}
	case ListAllLessonsCommand:
		lessons := clic.lessonRepository.All()
		fmt.Println("List of lessons:")
		for _, lesson := range lessons {
			fmt.Println(lesson.String())
		}
	case ListLessonsOfClassRoomCommand:
		var classRoom domain.ClassRoom
		for {
			classRoomString := clic.readUntil("> Enter class room id: ")
			classRoomId, err := strconv.Atoi(classRoomString)
			if err != nil {
				return err
			}
			foundClassRoom := clic.classRoomRepository.Find(int64(classRoomId))
			if foundClassRoom != nil {
				classRoom = *foundClassRoom
				break
			} else {
				fmt.Println("Could not find class room with given id...")
			}
		}
		lessons := clic.lessonRepository.FilterByClassRoom(*classRoom.ID)
		fmt.Println(classRoom.String())
		for _, lesson := range lessons {
			fmt.Println("-", lesson.String())
			fmt.Printf("\t- %s\n", lesson.Teacher.String())
			fmt.Println("\t- Pupils:")
			for _, pupil := range lesson.Pupils {
				fmt.Printf("\t\t- %s\n", pupil.String())
			}
		}
	default:
		fmt.Printf("Don't know command %#v\n", command)
	}
	return nil
}

func (clic *CLIController) Run() {
	for {
		fmt.Print(Menu)
		command, _ := clic.reader.ReadString('\n')
		command = strings.Trim(command, " \n")
		err := clic.OnCommand(command)
		if err != nil {
			log.Fatal(err)
		}
	}
}
