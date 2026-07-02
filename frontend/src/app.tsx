import { useEffect } from 'react'
import { Outlet } from 'react-router-dom'
import './app.scss'

function App() {
  useEffect(() => {
    console.log('App initialized')
  }, [])

  return <Outlet />
}

export default App