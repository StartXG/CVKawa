package common

type JsonSchema struct {
	CvBaseInfo CvBaseInfo  `json:"cv_base_info"`
	CvEduInfo  []CvEduInfo `json:"cv_edu_info"`
	CvResume   []CvResume  `json:"cv_resume"`
	CvSkills   CvSkills    `json:"cv_skills"`
}

type CvBaseInfo struct {
	Name      string `json:"name"`
	Birthday  int64  `json:"birthday"`
	Mobile    int64  `json:"mobile"`
	Email     string `json:"email"`
	PhotoPath string `json:"photo_path"`
} // CvBaseInfo

type CvEduInfo struct {
	School         string `json:"school"`
	Major          string `json:"major"`
	AcademicDegree string `json:"academic_degree"`
	StartTime      int64  `json:"start_time"`
	EndTime        int64  `json:"end_time"`
} // CvEduInfo

type CvResume struct {
	Company     string        `json:"company"`
	ResumeSteps []ResumeSteps `json:"resume_steps"`
	Achievement string        `json:"achievement"`
} // CvResume

type ResumeSteps struct {
	Position         string `json:"position"`
	Department       string `json:"department"`
	StartTime        int64  `json:"start_time"`
	EndTime          int64  `json:"end_time"`
	Responsibilities string `json:"responsibilities"`
} // ResumeSteps

type CvSkills []string // Skills
