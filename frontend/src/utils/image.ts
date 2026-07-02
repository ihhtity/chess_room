export function getRandomImage(width: number, height: number, seed?: string): string {
  if (seed) {
    const hash = seed.split('').reduce((acc, char) => acc + char.charCodeAt(0), 0)
    return `https://picsum.photos/seed/${hash}/${width}/${height}`
  }
  return `https://picsum.photos/${width}/${height}`
}

export function getRoomImage(images: string | null | undefined, width: number = 400, height: number = 300): string {
  if (!images) {
    return getRandomImage(width, height, 'room-default')
  }
  try {
    const parsed = JSON.parse(images)
    if (Array.isArray(parsed) && parsed.length > 0 && parsed[0]) {
      return parsed[0]
    }
  } catch {
    if (images.startsWith('http')) {
      return images
    }
  }
  return getRandomImage(width, height, images)
}

export function getUserAvatar(avatar: string | null | undefined): string {
  if (avatar && avatar.startsWith('http')) {
    return avatar
  }
  return getRandomImage(100, 100, avatar || 'user-default')
}

export const ImageSize = {
  RoomCover: { width: 400, height: 300 },
  RoomDetail: { width: 600, height: 450 },
  RoomThumbnail: { width: 200, height: 150 },
  Avatar: { width: 100, height: 100 },
  OrderThumbnail: { width: 120, height: 90 }
}
