import { Pagination, InputNumber, Button, Select } from 'antd'
import { useEffect, useState, useRef } from 'react'

interface CustomPaginationProps {
  current: number
  pageSize: number
  total: number
  onChange: (page: number, pageSize: number) => void
  showSizeChanger?: boolean
}

export default function CustomPagination({
  current,
  pageSize,
  total,
  onChange,
  showSizeChanger = true
}: CustomPaginationProps) {
  const [jumpPage, setJumpPage] = useState(current)
  const [localPageSize, setLocalPageSize] = useState(pageSize)
  const scrollRef = useRef<HTMLDivElement>(null)

  useEffect(() => {
    setJumpPage(current)
  }, [current])

  useEffect(() => {
    setLocalPageSize(pageSize)
  }, [pageSize])

  const handleJump = () => {
    if (jumpPage > 0 && jumpPage <= Math.ceil(total / pageSize)) {
      onChange(jumpPage, pageSize)
    }
  }

  const handleSizeChange = (value: number) => {
    setLocalPageSize(value)
    onChange(1, value)
  }

  const pageSizeOptions = ['10', '20', '50', '100']

  const startNum = total > 0 ? (current - 1) * pageSize + 1 : 0
  const endNum = Math.min(current * pageSize, total)

  return (
    <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', marginTop: 16, flexWrap: 'wrap', gap: 12 }}>
      <div style={{ display: 'flex', alignItems: 'center', gap: 8, flexWrap: 'wrap' }}>
        <span style={{ color: '#666', fontSize: 14 }}>
          显示第 <strong>{startNum}</strong> 到第 <strong>{endNum}</strong> 条记录，总共 <strong>{total}</strong> 条记录
        </span>
        {showSizeChanger && (
          <>
            <span style={{ color: '#666', fontSize: 14 }}>每页显示</span>
            <Select
              value={String(localPageSize)}
              onChange={(value) => handleSizeChange(Number(value))}
              style={{ width: 80 }}
              options={pageSizeOptions.map(size => ({
                value: size,
                label: size
              }))}
            />
            <span style={{ color: '#666', fontSize: 14 }}>条记录</span>
          </>
        )}
      </div>

      <div style={{ display: 'flex', alignItems: 'center', gap: 8 }}>
        <div ref={scrollRef} style={{ overflowX: 'auto', whiteSpace: 'nowrap', flex: 1, minWidth: 0 }}>
          <Pagination
            current={current}
            pageSize={pageSize}
            total={total}
            onChange={(page, size) => onChange(page, size)}
            showSizeChanger={false}
            showQuickJumper={false}
            itemRender={(_, _page, originalElement) => originalElement}
          />
        </div>

        <InputNumber
          min={1}
          max={Math.ceil(total / pageSize) || 1}
          value={jumpPage}
          onChange={(value) => setJumpPage(value || 1)}
          onKeyDown={(e) => {
            if (e.key === 'Enter') {
              handleJump()
            }
          }}
          style={{ width: 80 }}
        />
        <Button onClick={handleJump}>跳转</Button>
      </div>
    </div>
  )
}