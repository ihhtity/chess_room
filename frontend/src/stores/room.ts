import { defineStore } from 'pinia'
import { ref } from 'vue'
import type { Room, RoomType, TimeSlot } from '@/types'
import { roomApi, timeSlotApi } from '@/api'

export const useRoomStore = defineStore('room', () => {
  const rooms = ref<Room[]>([])
  const roomTypes = ref<RoomType[]>([])
  const timeSlots = ref<TimeSlot[]>([])
  const total = ref(0)

  const loadRooms = async (params?: { type_id?: number; floor?: string; status?: number; name?: string; page?: number; page_size?: number }) => {
    const res = await roomApi.getList(params)
    rooms.value = res.data
    total.value = res.total
  }

  const loadRoomTypes = async () => {
    const res = await roomApi.getTypes()
    roomTypes.value = res.data
  }

  const loadTimeSlots = async () => {
    const res = await timeSlotApi.getList()
    timeSlots.value = res.data
  }

  const getRoomById = (id: number) => {
    return rooms.value.find(r => r.id === id)
  }

  const getRoomTypeById = (id: number) => {
    return roomTypes.value.find(t => t.id === id)
  }

  return {
    rooms,
    roomTypes,
    timeSlots,
    total,
    loadRooms,
    loadRoomTypes,
    loadTimeSlots,
    getRoomById,
    getRoomTypeById
  }
})
