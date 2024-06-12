import { useState } from 'react'
import './App.css'
import NavBar from './components/navbar'

function App() {
  const [count, setCount] = useState(0)

  return (
    <div className="w-full h-screen">
      <NavBar />
    </div>
  )
}

export default App
