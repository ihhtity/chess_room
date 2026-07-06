export function formatDateTime(dateStr: string | null | undefined): string {
  if (!dateStr) return '-'
	return dateStr.replace('T', ' ').substring(0, 19)
}