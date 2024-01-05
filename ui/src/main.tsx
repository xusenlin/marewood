import React from 'react'
import ReactDOM from 'react-dom/client'
import App from '@/components/layout/app.tsx'
import { Toaster } from "sonner"
import "nprogress/nprogress.css"
import '../app/globals.css'
import '../app/style.css'

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.Fragment>
    <App />
    <Toaster position="bottom-left" closeButton richColors/>
  </React.Fragment>,
)


if (localStorage.theme === 'dark' || (!('theme' in localStorage) && window.matchMedia('(prefers-color-scheme: dark)').matches)) {
  document.documentElement.classList.add('dark')
} else {
  document.documentElement.classList.remove('dark')
}
