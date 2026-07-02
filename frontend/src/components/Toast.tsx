import { useEffect } from 'react'

interface ToastProps {
  message: string
  type?: 'success' | 'error' | 'warning' | 'info'
  duration?: number
  onClose: () => void
}

export default function Toast({ message, type = 'info', duration = 2000, onClose }: ToastProps) {
  useEffect(() => {
    const timer = setTimeout(() => {
      onClose()
    }, duration)
    return () => clearTimeout(timer)
  }, [duration, onClose])

  const typeStyles: Record<string, { bgColor: string; icon: string }> = {
    success: { bgColor: '#52c41a', icon: '✓' },
    error: { bgColor: '#ff4d4f', icon: '✕' },
    warning: { bgColor: '#faad14', icon: '⚠' },
    info: { bgColor: '#1890ff', icon: 'ℹ' }
  }

  const style = typeStyles[type]

  return (
    <div className="toast-overlay">
      <div className="toast-content" style={{ backgroundColor: style.bgColor }}>
        <span className="toast-icon">{style.icon}</span>
        <span className="toast-message">{message}</span>
      </div>
    </div>
  )
}

let toastContainer: HTMLElement | null = null
let currentToast: { message: string; type: string; duration: number } | null = null

function createToastContainer() {
  if (!toastContainer) {
    toastContainer = document.createElement('div')
    toastContainer.className = 'toast-container'
    document.body.appendChild(toastContainer)
  }
  return toastContainer
}

type ToastType = 'success' | 'error' | 'warning' | 'info'

interface ToastOptions {
  message: string
  type?: ToastType
  duration?: number
}

export function showToast(options: ToastOptions) {
  const container = createToastContainer()
  
  currentToast = {
    message: options.message,
    type: options.type || 'info',
    duration: options.duration || 2000
  }

  const toastEl = document.createElement('div')
  toastEl.className = `toast-item toast-${currentToast.type}`
  toastEl.innerHTML = `
    <span class="toast-icon">${currentToast.type === 'success' ? '✓' : currentToast.type === 'error' ? '✕' : currentToast.type === 'warning' ? '⚠' : 'ℹ'}</span>
    <span class="toast-message">${currentToast.message}</span>
  `
  
  container.appendChild(toastEl)
  
  setTimeout(() => {
    toastEl.classList.add('toast-hide')
    setTimeout(() => {
      if (toastEl.parentNode) {
        toastEl.parentNode.removeChild(toastEl)
      }
    }, 300)
  }, currentToast.duration)
}
