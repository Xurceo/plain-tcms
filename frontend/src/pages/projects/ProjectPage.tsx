import { useEffect, useState } from 'react'
import api from '../../api.ts'
import type { Project } from '../../types/Project.ts'
import { useParams } from 'react-router-dom'

export default function ProjectPage() {
  const [project, setProject] = useState<Project>()
  const { id } = useParams()

  useEffect(() => {
    api.get(`/projects/${id}`).then((res) => setProject(res.data))
  }, [])

  return (
    <div>
      <h1>Project: {project?.name}</h1>
      <p>
        Created: {new Date(project?.created_at!).toLocaleDateString('uk-UA')}
      </p>
    </div>
  )
}
