import { useState } from 'react'
import { Input, Select, Button, Row, Col, Space } from 'antd'
import { FilterOutlined } from '@ant-design/icons'

interface SearchField {
  key: string
  label: string
  type: 'input' | 'select' | 'number'
  options?: { label: string; value: string }[]
  placeholder?: string
}

interface SearchBarProps {
  fields: SearchField[]
  onSearch: (values: Record<string, string>) => void
  initialValues?: Record<string, string>
}

export default function SearchBar({ fields, onSearch, initialValues }: SearchBarProps) {
  const [values, setValues] = useState<Record<string, string>>(initialValues || {})

  const handleSearch = () => {
    const searchValues: Record<string, string> = {}
    Object.keys(values).forEach(key => {
      if (values[key] !== '' && values[key] !== undefined) {
        searchValues[key] = values[key]
      }
    })
    onSearch(searchValues)
  }

  const handleReset = () => {
    setValues({})
    onSearch({})
  }

  const handleChange = (key: string, value: string) => {
    setValues(prev => ({ ...prev, [key]: value }))
  }

  return (
    <div style={{ width: '100%', padding: 16, background: '#f5f5f5', borderRadius: 4 }}>
      <Row gutter={[16, 16]}>
        {fields.map(field => (
          <Col key={field.key} span={6}>
            <Space direction="vertical" style={{ width: '100%' }}>
              <span style={{ fontSize: 12, color: '#666' }}>{field.label}</span>
              {field.type === 'select' && field.options ? (
                <Select
                  style={{ width: '100%' }}
                  placeholder={field.placeholder || `请选择${field.label}`}
                  value={values[field.key] || undefined}
                  onChange={(value) => handleChange(field.key, value || '')}
                  options={field.options}
                />
              ) : (
                <Input
                  type={field.type === 'number' ? 'number' : 'text'}
                  placeholder={field.placeholder || `请输入${field.label}`}
                  value={values[field.key] || ''}
                  onChange={(e) => handleChange(field.key, e.target.value)}
                />
              )}
            </Space>
          </Col>
        ))}
      </Row>
      <Row style={{ marginTop: 16 }}>
        <Col offset={16} span={8}>
          <Space>
            <Button icon={<FilterOutlined />} onClick={handleSearch}>搜索</Button>
            <Button onClick={handleReset}>重置</Button>
          </Space>
        </Col>
      </Row>
    </div>
  )
}