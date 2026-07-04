import { useEffect, useState } from 'react'
import api from '../../api.ts'
import type { Project } from '../../types/Project.ts'

export default function ProjectsPage() {
  const [projects, setProjects] = useState<Project[]>([])

  useEffect(() => {
    api.get('/projects').then((res) => setProjects(res.data))
  }, [])

  return (
    <div>
      <h1>Projects</h1>
      {projects.map((p) => (
        <div key={p.id}>{p.name}</div>
      ))}
    </div>
  )
}
