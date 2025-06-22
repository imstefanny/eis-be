package models

import (
	"time"

	"gorm.io/gorm"
)

type StudentBehaviourActivities struct {
	ID                                    uint           `json:"id" gorm:"primaryKey"`
	AcademicID                            uint           `json:"academic_id" gorm:"default:null"`
	Academic                              Academics      `json:"academic" gorm:"foreignKey:AcademicID"`
	TermID                                uint           `json:"term_id" gorm:"default:null"`
	Term                                  Terms          `json:"term" gorm:"foreignKey:TermID"`
	StudentID                             uint           `json:"student_id" gorm:"default:null"`
	Student                               Students       `json:"student" gorm:"foreignKey:StudentID"`
	FirstBehaviour                        string         `json:"first_behaviour" gorm:"default:''"`
	FirstNeatness                         string         `json:"first_neatness" gorm:"default:''"`
	FirstCrafts                           string         `json:"first_crafts" gorm:"default:''"`
	FirstMonthExtracurricularFirst        string         `json:"first_month_extracurricular_first" gorm:"default:''"`
	FirstMonthExtracurricularScoreFirst   string         `json:"first_month_extracurricular_score_first" gorm:"default:''"`
	FirstMonthExtracurricularSecond       string         `json:"first_month_extracurricular_second" gorm:"default:''"`
	FirstMonthExtracurricularScoreSecond  string         `json:"first_month_extracurricular_score_second" gorm:"default:''"`
	SecondBehaviour                       string         `json:"second_behaviour" gorm:"default:''"`
	SecondNeatness                        string         `json:"second_neatness" gorm:"default:''"`
	SecondCrafts                          string         `json:"second_crafts" gorm:"default:''"`
	SecondMonthExtracurricularFirst       string         `json:"second_month_extracurricular_first" gorm:"default:''"`
	SecondMonthExtracurricularScoreFirst  string         `json:"second_month_extracurricular_score_first" gorm:"default:''"`
	SecondMonthExtracurricularSecond      string         `json:"second_month_extracurricular_second" gorm:"default:''"`
	SecondMonthExtracurricularScoreSecond string         `json:"second_month_extracurricular_score_second" gorm:"default:''"`
	FirstNotes                            string         `json:"first_notes" gorm:"default:''"`
	SecondNotes                           string         `json:"second_notes" gorm:"default:''"`
	CreatedAt                             time.Time      `json:"created_at"`
	UpdatedAt                             time.Time      `json:"updated_at"`
	DeletedAt                             gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
