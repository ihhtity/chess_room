import { createContext, useContext, useState, useEffect, ReactNode } from 'react'
import { permissionApi } from '@/api'

interface PermissionContextType {
  permissions: string[]
  hasPermission: (code: string) => boolean
  hasAnyPermission: (codes: string[]) => boolean
  hasAllPermissions: (codes: string[]) => boolean
  loadPermissions: () => Promise<void>
}

const PermissionContext = createContext<PermissionContextType | undefined>(undefined)

export const PermissionProvider = ({ children }: { children: ReactNode }) => {
  const [permissions, setPermissions] = useState<string[]>([])

  const loadPermissions = async () => {
    try {
      const codes = await permissionApi.getMyPermissions()
      if (Array.isArray(codes)) {
        setPermissions(codes)
      } else {
        setPermissions([])
      }
    } catch (error) {
      console.error('加载权限失败:', error)
      setPermissions([])
    }
  }

  useEffect(() => {
    loadPermissions()
  }, [])

  const hasPermission = (code: string) => {
    return permissions.includes(code)
  }

  const hasAnyPermission = (codes: string[]) => {
    return codes.some(code => permissions.includes(code))
  }

  const hasAllPermissions = (codes: string[]) => {
    return codes.every(code => permissions.includes(code))
  }

  return (
    <PermissionContext.Provider value={{ permissions, hasPermission, hasAnyPermission, hasAllPermissions, loadPermissions }}>
      {children}
    </PermissionContext.Provider>
  )
}

export const usePermission = () => {
  const context = useContext(PermissionContext)
  if (!context) {
    throw new Error('usePermission must be used within a PermissionProvider')
  }
  return context
}