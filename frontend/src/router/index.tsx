import { createBrowserRouter } from 'react-router-dom'
import App from '@/app'
import Index from '@/pages/index'
import RoomDetail from '@/pages/room/detail'
import BookingCreate from '@/pages/booking/create'
import OrderList from '@/pages/order/list'
import OrderDetail from '@/pages/order/detail'
import MemberIndex from '@/pages/member/index'
import UserProfile from '@/pages/user/profile'
import LoginPage from '@/pages/user/login'
import ActivityPage from '@/pages/activity/index'
import AnnouncementPage from '@/pages/announcement/index'
import RechargePage from '@/pages/recharge/index'
import PrivateRoute from '@/components/PrivateRoute'

const router = createBrowserRouter([
  {
    path: '/login',
    element: <LoginPage />
  },
  {
    path: '/',
    element: <App />,
    children: [
      { path: '/', element: <Index /> },
      { path: '/rooms/:id', element: <RoomDetail /> },
      { 
        path: '/booking/:room_id', 
        element: (
          <PrivateRoute>
            <BookingCreate />
          </PrivateRoute>
        ) 
      },
      { 
        path: '/orders', 
        element: (
          <PrivateRoute>
            <OrderList />
          </PrivateRoute>
        ) 
      },
      { 
        path: '/orders/:id', 
        element: (
          <PrivateRoute>
            <OrderDetail />
          </PrivateRoute>
        ) 
      },
      { 
        path: '/member', 
        element: (
          <PrivateRoute>
            <MemberIndex />
          </PrivateRoute>
        ) 
      },
      { 
        path: '/profile', 
        element: (
          <PrivateRoute>
            <UserProfile />
          </PrivateRoute>
        ) 
      },
      { 
        path: '/activities', 
        element: <ActivityPage /> 
      },
      { 
        path: '/announcements', 
        element: <AnnouncementPage /> 
      },
      { 
        path: '/recharge', 
        element: (
          <PrivateRoute>
            <RechargePage />
          </PrivateRoute>
        ) 
      }
    ]
  }
])

export default router
