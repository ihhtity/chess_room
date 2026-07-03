import { Button, Space, Popconfirm } from 'antd'
import { DeleteOutlined, EditOutlined } from '@ant-design/icons'

interface BatchActionsProps {
  selectedRowKeys: string[]
  onBatchDelete: (ids: string[]) => void
  onBatchEdit?: (ids: string[]) => void
  showEdit?: boolean
}

export default function BatchActions({ selectedRowKeys, onBatchDelete, onBatchEdit, showEdit = false }: BatchActionsProps) {
  const hasSelected = selectedRowKeys.length > 0

  return (
    <Space>
      {hasSelected && (
        <>
          {showEdit && onBatchEdit && (
            <Button icon={<EditOutlined />} onClick={() => onBatchEdit(selectedRowKeys)}>
              批量编辑 ({selectedRowKeys.length})
            </Button>
          )}
          <Popconfirm
            title={`确定删除选中的 ${selectedRowKeys.length} 条记录吗？`}
            onConfirm={() => onBatchDelete(selectedRowKeys)}
            okText="确定"
            cancelText="取消"
          >
            <Button danger icon={<DeleteOutlined />}>
              批量删除 ({selectedRowKeys.length})
            </Button>
          </Popconfirm>
        </>
      )}
    </Space>
  )
}