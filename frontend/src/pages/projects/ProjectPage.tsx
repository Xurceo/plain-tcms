/*
 * // TCMS - Test Case Management System
 * // Copyright (C) 2026 Pavlo Shnal
 * //
 * // This program is free software: you can redistribute it and/or modify
 * // it under the terms of the GNU Affero General Public License as published
 * // by the Free Software Foundation, either version 3 of the License, or
 * // (at your option) any later version.
 * //
 * // This program is distributed in the hope that it will be useful,
 * // but WITHOUT ANY WARRANTY; without even the implied warranty of
 * // MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * // GNU Affero General Public License for more details.
 * //
 * // You should have received a copy of the GNU Affero General Public License
 * // along with this program. If not, see <https://www.gnu.org/licenses/>.
 */

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
