import { defineStore } from 'pinia'
import { ref } from 'vue'
import type { Order } from '@/types'
import { orderApi } from '@/api'

export const useOrderStore = defineStore('order', () => {
  const orders = ref<Order[]>([])
  const total = ref(0)

  const loadOrders = async (params?: { status?: number; page?: number; page_size?: number }) => {
    const res = await orderApi.getList(params)
    orders.value = res.data
    total.value = res.total
  }

  const createOrder = async (data: { room_id: number; start_time: string; end_time: string; remark?: string }) => {
    return await orderApi.create(data)
  }

  const cancelOrder = async (id: number) => {
    await orderApi.cancel(id)
    await loadOrders()
  }

  const confirmOrder = async (id: number) => {
    await orderApi.confirm(id)
    await loadOrders()
  }

  const completeOrder = async (id: number) => {
    await orderApi.complete(id)
    await loadOrders()
  }

  const getOrderById = (id: number) => {
    return orders.value.find(o => o.id === id)
  }

  return {
    orders,
    total,
    loadOrders,
    createOrder,
    cancelOrder,
    confirmOrder,
    completeOrder,
    getOrderById
  }
})
