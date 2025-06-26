package dto

type CreateStudentGradesRequest struct {
	AcademicID uint                               `json:"academic_id" validate:"required"`
	TermID     uint                               `json:"term_id" validate:"required"`
	Details    []CreateStudentGradesDetailRequest `json:"details" validate:"required,dive"`
}

type CreateStudentGradesDetailRequest struct {
	SubjectID uint                              `json:"subject_id" validate:"required"`
	Students  []CreateStudentGradesEntryRequest `json:"students" validate:"required,dive"`
}

type CreateStudentGradesEntryRequest struct {
	StudentID   uint    `json:"student_id" validate:"required"`
	FirstQuiz   float64 `json:"first_quiz"`
	SecondQuiz  float64 `json:"second_quiz"`
	FirstMonth  float64 `json:"first_month"`
	SecondMonth float64 `json:"second_month"`
	Finals      float64 `json:"finals"`
	Remarks     string  `json:"remarks"`
}

type GetStudentGradesResponse struct {
	AcademicID   uint                             `json:"academic_id"`
	Academic     string                           `json:"academic"`
	TermID       uint                             `json:"term_id"`
	Term         string                           `json:"term"`
	Details      []GetStudentGradesDetailResponse `json:"details"`
	TeacherNotes []GetTeacherNotesResponse        `json:"teacher_notes"`
}
type GetStudentGradesDetailResponse struct {
	SubjectID uint                            `json:"subject_id"`
	Subject   string                          `json:"subject"`
	Students  []GetStudentGradesEntryResponse `json:"students"`
}
type GetStudentGradesEntryResponse struct {
	ID          uint    `json:"id"`
	StudentID   uint    `json:"student_id"`
	StudentName string  `json:"student_name"`
	DisplayName string  `json:"display_name"`
	FirstQuiz   float64 `json:"first_quiz"`
	SecondQuiz  float64 `json:"second_quiz"`
	FirstMonth  float64 `json:"first_month"`
	SecondMonth float64 `json:"second_month"`
	Finals      float64 `json:"finals"`
	FinalGrade  float64 `json:"final_grade"`
	Remarks     string  `json:"remarks"`
}
type GetTeacherNotesResponse struct {
	ID        uint   `json:"id"`
	StudentID uint   `json:"student_id"`
	Student   string `json:"student"`
	Notes     string `json:"notes"`
}

type UpdateStudentGradesRequest struct {
	AcademicID uint                               `json:"academic_id" validate:"required"`
	TermID     uint                               `json:"term_id" validate:"required"`
	Details    []UpdateStudentGradesDetailRequest `json:"details" validate:"required,dive"`
}
type UpdateStudentGradesDetailRequest struct {
	SubjectID uint                              `json:"subject_id" validate:"required"`
	Subject   string                            `json:"subject"`
	Students  []UpdateStudentGradesEntryRequest `json:"students" validate:"required,dive"`
}
type UpdateStudentGradesEntryRequest struct {
	ID          uint    `json:"id"`
	StudentID   uint    `json:"student_id" validate:"required"`
	FirstQuiz   float64 `json:"first_quiz"`
	SecondQuiz  float64 `json:"second_quiz"`
	FirstMonth  float64 `json:"first_month"`
	SecondMonth float64 `json:"second_month"`
	Finals      float64 `json:"finals"`
	Remarks     string  `json:"remarks"`
}

type StudentScoreResponse struct {
	SubjectName string  `json:"subject_name"`
	FirstMonth  float64 `json:"first_month"`
	SecondMonth float64 `json:"second_month"`
	FirstQuiz   float64 `json:"first_quiz"`
	SecondQuiz  float64 `json:"second_quiz"`
	Finals      float64 `json:"finals"`
	FinalGrade  float64 `json:"final_grade"`
	Remarks     string  `json:"remarks"`
}

type StudentGradesReport struct {
	Class    string                          `json:"class"`
	Average  float64                         `json:"average"`
	Students []StudentGradesReportTopStudent `json:"students"`
}
type StudentGradesReportTopStudent struct {
	Rank    int     `json:"rank"`
	Student string  `json:"student"`
	NIS     string  `json:"nis"`
	Class   string  `json:"class"`
	Finals  float64 `json:"finals"`
}
type StudentGradesReportQuery struct {
	StudentID uint    `json:"student_id"`
	Student   string  `json:"student"`
	NIS       string  `json:"nis"`
	ClassID   uint    `json:"class_id"`
	Class     string  `json:"class"`
	Finals    float64 `json:"finals"`
}

type GetPrintReportByStudent struct {
	Name             string                          `json:"name"`
	NIS              string                          `json:"nis"`
	NISN             string                          `json:"nisn"`
	Level            string                          `json:"level"`
	Class            string                          `json:"class"`
	Fase             string                          `json:"fase"`
	Term             int                             `json:"term"`
	AcademicYear     string                          `json:"academic_year"`
	Grades           []GetPrintReportGrade           `json:"grades"`
	Extracurriculars []GetPrintReportExtracurricular `json:"extracurriculars"`
	Sick             int                             `json:"sick"`
	Permission       int                             `json:"permission"`
	Absent           int                             `json:"absent"`
	HomeRoomTeacher  string                          `json:"home_room_teacher"`
	Principal        string                          `json:"principal"`
	TeacherNotes     string                          `json:"teacher_notes"`
}
type GetPrintReportGrade struct {
	Subject string  `json:"subject"`
	Finals  float64 `json:"finals"`
	Remarks string  `json:"remarks"`
}
type GetPrintReportExtracurricular struct {
	Name  string `json:"name"`
	Score string `json:"score"`
}

type GetPrintMonthlyReportByStudent struct {
	Name            string                       `json:"name"`
	NIS             string                       `json:"nis"`
	Class           string                       `json:"class"`
	AcademicYear    string                       `json:"academic_year"`
	Grades          []GetPrintMonthlyReportGrade `json:"grades"`
	HomeRoomTeacher string                       `json:"home_room_teacher"`

	StFirstBehavior  string `json:"st_first_behavior"`
	StSecondBehavior string `json:"st_second_behavior"`
	StFirstCraft     string `json:"st_first_craft"`
	StSecondCraft    string `json:"st_second_craft"`
	StFirstTidiness  string `json:"st_first_tidiness"`
	StSecondTidiness string `json:"st_second_tidiness"`

	StFirstExtracurricularFirst        string `json:"st_first_extracurricular_first"`
	StFirstExtracurricularScoreFirst   string `json:"st_first_extracurricular_score_first"`
	StFirstExtracurricularSecond       string `json:"st_first_extracurricular_second"`
	StFirstExtracurricularScoreSecond  string `json:"st_first_extracurricular_score_second"`
	StSecondExtracurricularFirst       string `json:"st_second_extracurricular_first"`
	StSecondExtracurricularScoreFirst  string `json:"st_second_extracurricular_score_first"`
	StSecondExtracurricularSecond      string `json:"st_second_extracurricular_second"`
	StSecondExtracurricularScoreSecond string `json:"st_second_extracurricular_score_second"`
	StFirstNotes                       string `json:"st_first_notes"`
	StSecondNotes                      string `json:"st_second_notes"`

	StFirstSick        int `json:"st_first_sick"`
	StSecondSick       int `json:"st_second_sick"`
	StFirstPermission  int `json:"st_first_permission"`
	StSecondPermission int `json:"st_second_permission"`
	StFirstAbsent      int `json:"st_first_absent"`
	StSecondAbsent     int `json:"st_second_absent"`

	NdFirstBehavior  string `json:"nd_first_behavior"`
	NdSecondBehavior string `json:"nd_second_behavior"`
	NdFirstCraft     string `json:"nd_first_craft"`
	NdSecondCraft    string `json:"nd_second_craft"`
	NdFirstTidiness  string `json:"nd_first_tidiness"`
	NdSecondTidiness string `json:"nd_second_tidiness"`

	NdFirstExtracurricularFirst        string `json:"nd_first_extracurricular_first"`
	NdFirstExtracurricularScoreFirst   string `json:"nd_first_extracurricular_score_first"`
	NdFirstExtracurricularSecond       string `json:"nd_first_extracurricular_second"`
	NdFirstExtracurricularScoreSecond  string `json:"nd_first_extracurricular_score_second"`
	NdSecondExtracurricularFirst       string `json:"nd_second_extracurricular_first"`
	NdSecondExtracurricularScoreFirst  string `json:"nd_second_extracurricular_score_first"`
	NdSecondExtracurricularSecond      string `json:"nd_second_extracurricular_second"`
	NdSecondExtracurricularScoreSecond string `json:"nd_second_extracurricular_score_second"`
	NdFirstNotes                       string `json:"nd_first_notes"`
	NdSecondNotes                      string `json:"nd_second_notes"`

	NdFirstSick        int `json:"nd_first_sick"`
	NdSecondSick       int `json:"nd_second_sick"`
	NdFirstPermission  int `json:"nd_first_permission"`
	NdSecondPermission int `json:"nd_second_permission"`
	NdFirstAbsent      int `json:"nd_first_absent"`
	NdSecondAbsent     int `json:"nd_second_absent"`

	StFirstDate  string `json:"st_first_date"`
	StSecondDate string `json:"st_second_date"`
	NdFirstDate  string `json:"nd_first_date"`
	NdSecondDate string `json:"nd_second_date"`
}
type GetPrintMonthlyReportGrade struct {
	Subject       string  `json:"subject"`
	StFirstQuiz   float64 `json:"st_first_quiz"`
	StSecondQuiz  float64 `json:"st_second_quiz"`
	StFirstMonth  float64 `json:"st_first_month"`
	StSecondMonth float64 `json:"st_second_month"`
	NdFirstQuiz   float64 `json:"nd_first_quiz"`
	NdSecondQuiz  float64 `json:"nd_second_quiz"`
	NdFirstMonth  float64 `json:"nd_first_month"`
	NdSecondMonth float64 `json:"nd_second_month"`
}
