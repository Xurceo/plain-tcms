import { Outlet } from 'react-router-dom'

export default function Layout() {
  return (
    <div className="min-h-screen bg-background text-foreground">
      <nav className=" shadow p-4">
        <h1 className="text-xl font-bold">TCMS</h1>
      </nav>
      <main className="p-4">
        <Outlet />
      </main>
    </div>
  )
}
