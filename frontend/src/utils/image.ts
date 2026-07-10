export function getRandomImage(id: string | number, width = 400, height = 300): string {
  return `https://picsum.photos/seed/${id}/${width}/${height}`
}

export function getAvatar(id: string | number): string {
  return `https://picsum.photos/seed/avatar${id}/100/100`
}

export function getRoomImage(id: string | number): string {
  return `https://picsum.photos/seed/room${id}/800/600`
}

export function getActivityImage(id: string | number): string {
  return `https://picsum.photos/seed/activity${id}/600/400`
}

export function getCouponImage(id: string | number): string {
  return `https://picsum.photos/seed/coupon${id}/300/200`
}