use serde::{Deserialize, Serialize};

#[derive(Default, Debug, Serialize, Deserialize, Clone)]
pub struct CV {
    pub cv_base_info: CVBaseInfo,
    pub cv_edu_info: Vec<CVEduInfo>,
    pub cv_resume: Vec<CVResume>,
    pub cv_skills: Vec<String>,
}

#[derive(Default,Debug, Serialize, Deserialize,Clone)]
pub struct CVResume {
    pub company: String,
    pub resume_steps: Vec<CVResumeStep>,
    pub achievement: String,
}

#[derive(Default,Debug, Serialize, Deserialize,Clone)]
pub struct CVResumeStep{
    pub position: String,
    pub department: String,
    pub start_time: i64,
    pub end_time: i64,
    pub responsibilities: String,
}

#[derive(Default,Debug, Serialize, Deserialize,Clone)]
pub struct CVEduInfo {
    pub school: String,
    pub major: String,
    pub academic_degree: String,
    pub start_time: i64,
    pub end_time: i64,
}

#[derive(Default,Debug, Serialize, Deserialize,Clone)]
pub struct CVBaseInfo {
    pub name: String,
    pub birthday: i64,
    pub mobile: i64,
    pub email: String,
    pub photo_path: Option<String>,
}
