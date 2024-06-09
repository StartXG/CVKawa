interface CVBaseInfoSchema {
  name: string | null
  birthday: number | null
  mobile: number | null
  email: string | null
  photo_path?: string | null
}

interface CVEduInfoSchema {
  school: string | null
  major: string | null
  academic_degree: string | null
  start_time: number | null
  end_time: number | null
}

interface CVResumeSchema {
  company: string
  resume_steps: CVResumeStepSchema[]
  achievement: string
}

interface CVResumeStepSchema {
  position: string | null
  department: string | null
  start_time: number | null
  end_time: number | null
  responsibilities: string | null
}

interface CVDataSchema {
  cv_base_info: CVBaseInfoSchema
  cv_edu_info: CVEduInfoSchema[]
  cv_resume: CVResumeSchema[]
  cv_skills: string[]
}

export type { CVBaseInfoSchema, CVEduInfoSchema, CVResumeSchema, CVDataSchema, CVResumeStepSchema }
