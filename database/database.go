package database

import (
	"eis-be/config"
	"eis-be/models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	InitDB()
	InitialMigration()
	PopulateRolesPermissions()
	PopulateLevels()
	PopulateDocTypes()
}

func InitDB() *gorm.DB {
	dbconfig := config.ReadEnv()
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbconfig.DB_USERNAME,
		dbconfig.DB_PASSWORD,
		dbconfig.DB_HOSTNAME,
		dbconfig.DB_PORT,
		dbconfig.DB_NAME,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return DB
}

func InitialMigration() {
	DB.AutoMigrate(&models.Blogs{})
	DB.AutoMigrate(&models.Users{})
	DB.AutoMigrate(&models.Applicants{})
	DB.AutoMigrate(&models.Students{})
	DB.AutoMigrate(&models.StudentAttendances{})
	DB.AutoMigrate(&models.Guardians{})
	DB.AutoMigrate(&models.DocTypes{})
	DB.AutoMigrate(&models.Documents{})
	DB.AutoMigrate(&models.WorkScheds{})
	DB.AutoMigrate(&models.WorkSchedDetails{})
	DB.AutoMigrate(&models.Subjects{})
	DB.AutoMigrate(&models.Levels{})
	DB.AutoMigrate(&models.LevelHistories{})
	DB.AutoMigrate(&models.Classrooms{})
	DB.AutoMigrate(&models.Teachers{})
	DB.AutoMigrate(&models.TeacherAttendances{})
	DB.AutoMigrate(&models.Academics{})
	DB.AutoMigrate(&models.SubjectSchedules{})
	DB.AutoMigrate(&models.ClassNotes{})
	DB.AutoMigrate(&models.ClassNotesDetails{})
	DB.AutoMigrate(&models.StudentGrades{})
	DB.AutoMigrate(&models.Roles{}, &models.Permissions{})
}

func PopulateRolesPermissions() {
	var count int64
	if err := DB.Model(&models.Roles{}).Count(&count).Error; err != nil || count > 0 {
		return
	}

	permNames := []string{
		"registration:read","registration:write",
		"student:read","student:write",
		"academic:read","academic:write",
		"studentatt:read","studentatt:write",
		"subject:read","subject:write",
		"classnote:read","classnote:write",
		"studentattrep:read","studentattrep:write",
		"examrecap:read","examrecap:write",
		"grade:read","grade:write",
		"class:read","class:write",
		"subjsched:read","subjsched:write",
		"document:read","document:write",
		"doctype:read","doctype:write",
		"teacher:read","teacher:write",
		"teacheratt:read","teacheratt:write",
		"teacherattrep:read","teacherattrep:write",
		"worksched:read","worksched:write",
		"news:read","news:write",
		"users:read","users:write",
		"accessrights:read","accessrights:write",
	}

	permissions := []models.Permissions{}
    for _, name := range permNames {
        permissions = append(permissions, models.Permissions{Name: name})
    }
    if err := DB.Create(&permissions).Error; err != nil {
		fmt.Printf("Error creating permissions: %v\n", err)
		return
	}

	permMap := make(map[string]models.Permissions)
	if err := DB.Find(&permissions).Error; err != nil {
		fmt.Printf("Error retrieving permissions: %v\n", err)
		return
	}
	for _, perm := range permissions {
		permMap[perm.Name] = perm
	}

	roles := []models.Roles{
		{Name: "Admin", Permissions: []models.Permissions{
			permMap["registration:read"],
			permMap["registration:write"],
			permMap["student:read"],
			permMap["student:write"],
			permMap["academic:read"],
			permMap["academic:write"],
			permMap["studentatt:read"],
			permMap["studentatt:write"],
			permMap["subject:read"],
			permMap["subject:write"],
			permMap["classnote:read"],
			permMap["classnote:write"],
			permMap["studentattrep:read"],
			permMap["studentattrep:write"],
			permMap["examrecap:read"],
			permMap["examrecap:write"],
			permMap["grade:read"],
			permMap["grade:write"],
			permMap["class:read"],
			permMap["class:write"],
			permMap["subjsched:read"],
			permMap["subjsched:write"],
			permMap["document:read"],
			permMap["document:write"],
			permMap["doctype:read"],
			permMap["doctype:write"],
			permMap["teacher:read"],
			permMap["teacher:write"],
			permMap["teacheratt:read"],
			permMap["teacheratt:write"],
			permMap["teacherattrep:read"],
			permMap["teacherattrep:write"],
			permMap["worksched:read"],
			permMap["worksched:write"],
			permMap["news:read"],
			permMap["news:write"],
			permMap["users:read"],
			permMap["users:write"],
			permMap["accessrights:read"],
			permMap["accessrights:write"],
		}},
		{Name: "Teacher", Permissions: []models.Permissions{
			permMap["student:read"],
			permMap["academic:read"],
			permMap["academic:write"],
			permMap["studentatt:read"],
			permMap["studentatt:write"],
			permMap["classnote:read"],
			permMap["classnote:write"],
			permMap["studentattrep:read"],
			permMap["examrecap:read"],
			permMap["teacher:read"],
			permMap["teacher:write"],
			permMap["teacheratt:read"],
			permMap["teacherattrep:read"],
			permMap["worksched:read"],
			permMap["news:read"],
		}},
		{Name: "Student", Permissions: []models.Permissions{
			permMap["student:read"],
			permMap["student:write"],
			permMap["studentatt:read"],
			permMap["studentattrep:read"],
			permMap["examrecap:read"],
			permMap["subjsched:read"],
			permMap["news:read"],
		}},
		{Name: "Principal", Permissions: []models.Permissions{
			permMap["registration:read"],
			permMap["student:read"],
			permMap["academic:read"],
			permMap["studentatt:read"],
			permMap["subject:read"],
			permMap["classnote:read"],
			permMap["studentattrep:read"],
			permMap["examrecap:read"],
			permMap["grade:read"],
			permMap["class:read"],
			permMap["subjsched:read"],
			permMap["document:read"],
			permMap["doctype:read"],
			permMap["teacher:read"],
			permMap["teacheratt:read"],
			permMap["teacherattrep:read"],
			permMap["worksched:read"],
			permMap["news:read"],
		}},
		{Name: "Applicant", Permissions: []models.Permissions{
			permMap["news:read"],
			permMap["registration:read"],
			permMap["registration:write"],
		}},
	}
	for _, role := range roles {
		if err := DB.Create(&role).Error; err != nil {
			fmt.Printf("Error creating role %s: %v\n", role.Name, err)
			continue
		}
	}
}

func PopulateLevels() {
	var count int64
	if err := DB.Model(&models.Roles{}).Count(&count).Error; err != nil || count > 0 {
		return
	}
	levels := []models.Levels{
		{Name: "TK"},
		{Name: "SD"},
		{Name: "SMP"},
		{Name: "SMA"},
	}
	for _, level := range levels {
		if err := DB.Create(&level).Error; err != nil {
			fmt.Printf("Error creating level %s: %v\n", level.Name, err)
			continue
		}
	}
}

func PopulateDocTypes() {
	var count int64
	if err := DB.Model(&models.DocTypes{}).Count(&count).Error; err != nil || count > 0 {
		return
	}
	docTypes := []models.DocTypes{
		{Name: "Kartu Keluarga"},
		{Name: "Akta Kelahiran"},
		{Name: "KTP Orang Tua"},
		{Name: "Bukti Pembayaran"},
	}
	for _, docType := range docTypes {
		if err := DB.Create(&docType).Error; err != nil {
			fmt.Printf("Error creating document type %s: %v\n", docType.Name, err)
			continue
		}
	}
}
