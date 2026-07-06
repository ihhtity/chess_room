import { Button, Space, Popconfirm } from 'antd'
import { DeleteOutlined, EditOutlined } from '@ant-design/icons'

interface BatchActionsProps {
  selectedRowKeys: string[]
  onBatchDelete: (ids: string[]) => void
  onBatchEdit?: (ids: string[]) => void
  showEdit?: boolean
}

export default function BatchActions({ selectedRowKeys, onBatchDelete, onBatchEdit, showEdit = true }: BatchActionsProps) {
  const hasSelection = selectedRowKeys.length > 0

  return (
    <Space>
      {showEdit && (
        <Popconfirm
          title={`确定编辑选中的 ${selectedRowKeys.length} 条记录吗？`}
          onConfirm={() => onBatchEdit?.(selectedRowKeys)}
          okText="确定"
          cancelText="取消"
          disabled={!hasSelection}
        >
          <Button 
            icon={<EditOutlined />} 
            disabled={!hasSelection}
            style={{ backgroundColor: '#52c41a', color: '#fff', borderColor: '#52c41a' }}
          >
            编辑 ({selectedRowKeys.length})
          </Button>
        </Popconfirm>
      )}
      <Popconfirm
        title={`确定删除选中的 ${selectedRowKeys.length} 条记录吗？`}
        onConfirm={() => onBatchDelete?.(selectedRowKeys)}
        okText="确定"
        cancelText="取消"
        disabled={!hasSelection}
      >
        <Button danger icon={<DeleteOutlined />} disabled={!hasSelection}>
          删除 ({selectedRowKeys.length})
        </Button>
      </Popconfirm>
    </Space>
  )
}