import { Modal, Form, Input, Select, InputNumber, DatePicker, message, Table } from 'antd'
import dayjs from 'dayjs'

interface FieldConfig {
  key: string
  label: string
  type: 'input' | 'select' | 'number' | 'date' | 'textarea'
  options?: { label: string; value: string | number }[]
  placeholder?: string
  required?: boolean
  min?: number
  max?: number
}

interface BatchEditModalProps {
  visible: boolean
  onCancel: () => void
  onOk: (data: Record<string, any>[]) => void
  records: any[]
  fields: FieldConfig[]
  title?: string
}

export default function BatchEditModal({ visible, onCancel, onOk, records, fields, title = '批量编辑' }: BatchEditModalProps) {
  const [form] = Form.useForm()

  const handleSubmit = async () => {
    try {
      const values = await form.validateFields()
      const updatedRecords = records.map(record => {
        const updated: Record<string, any> = { id: record.id }
        fields.forEach(field => {
          const key = `${field.key}_${record.id}`
          if (values[key] !== undefined) {
            updated[field.key] = values[key]
          }
        })
        return updated
      })
      onOk(updatedRecords)
    } catch (error) {
      message.error('批量编辑失败')
    }
  }

  const columns = fields.map(field => ({
    title: field.label,
    dataIndex: field.key,
    key: field.key,
    render: (_: any, record: any) => {
      const fieldKey = `${field.key}_${record.id}`
      switch (field.type) {
        case 'select':
          return (
            <Form.Item name={fieldKey}>
              <Select defaultValue={record[field.key]} options={field.options} style={{ width: '100%' }} />
            </Form.Item>
          )
        case 'number':
          return (
            <Form.Item name={fieldKey}>
              <InputNumber defaultValue={record[field.key]} min={field.min} max={field.max} style={{ width: '100%' }} />
            </Form.Item>
          )
        case 'date':
          return (
            <Form.Item name={fieldKey}>
              <DatePicker showTime format="YYYY-MM-DD HH:mm:ss" defaultValue={record[field.key] ? dayjs(record[field.key]) : null} style={{ width: '100%' }} />
            </Form.Item>
          )
        case 'textarea':
          return (
            <Form.Item name={fieldKey}>
              <Input.TextArea defaultValue={record[field.key]} style={{ width: '100%' }} />
            </Form.Item>
          )
        default:
          return (
            <Form.Item name={fieldKey}>
              <Input defaultValue={record[field.key]} style={{ width: '100%' }} />
            </Form.Item>
          )
      }
    }
  }))

  return (
    <Modal
      title={title}
      open={visible}
      onCancel={onCancel}
      onOk={handleSubmit}
      width={800}
      okText="保存"
      cancelText="取消"
    >
      <Form form={form} layout="vertical">
        <Table
          dataSource={records}
          columns={columns}
          rowKey="id"
          pagination={false}
          size="small"
        />
      </Form>
    </Modal>
  )
}