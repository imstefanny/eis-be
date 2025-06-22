package dto

type StudentBehaviourActivityRequest struct {
	ID          uint   `json:"id"`
	AcademicID  uint   `json:"academic_id" validate:"required"`
	TermID      uint   `json:"term_id" validate:"required"`
	StudentID   uint   `json:"student_id"`
	StudentName string `json:"student_name"`

	FirstNeatness   string `json:"first_neatness"`
	FirstCrafts     string `json:"first_crafts"`
	FirstBehaviour  string `json:"first_behaviour"`
	SecondNeatness  string `json:"second_neatness"`
	SecondCrafts    string `json:"second_crafts"`
	SecondBehaviour string `json:"second_behaviour"`

	FirstMonthExtracurricularFirst       string `json:"first_month_extracurricular_first"`
	FirstMonthExtracurricularScoreFirst  string `json:"first_month_extracurricular_score_first"`
	FirstMonthExtracurricularSecond      string `json:"first_month_extracurricular_second"`
	FirstMonthExtracurricularScoreSecond string `json:"first_month_extracurricular_score_second"`

	SecondMonthExtracurricularFirst       string `json:"second_month_extracurricular_first"`
	SecondMonthExtracurricularScoreFirst  string `json:"second_month_extracurricular_score_first"`
	SecondMonthExtracurricularSecond      string `json:"second_month_extracurricular_second"`
	SecondMonthExtracurricularScoreSecond string `json:"second_month_extracurricular_score_second"`
}
