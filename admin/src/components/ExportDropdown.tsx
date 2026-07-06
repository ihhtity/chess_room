import { Button, Dropdown, MenuProps } from 'antd'
import { DownloadOutlined, FileOutlined, FileTextOutlined, FileExcelOutlined, FileWordOutlined } from '@ant-design/icons'
import { saveAs } from 'file-saver'
import * as XLSX from 'xlsx'

interface ExportDropdownProps {
  data: any[]
  filename: string
}

const exportToJson = (data: any[], filename: string) => {
  const content = JSON.stringify(data, null, 2)
  const blob = new Blob([content], { type: 'application/json' })
  saveAs(blob, `${filename}.json`)
}

const exportToXml = (data: any[], filename: string) => {
  const objToXml = (obj: any, indent: string = '') => {
    let xml = ''
    for (const key in obj) {
      if (obj.hasOwnProperty(key)) {
        const value = obj[key]
        xml += `${indent}<${key}>`
        if (Array.isArray(value)) {
          xml += '\n'
          value.forEach(item => {
            xml += objToXml(item, indent + '  ')
          })
          xml += `${indent}</${key}>\n`
        } else if (typeof value === 'object' && value !== null) {
          xml += '\n' + objToXml(value, indent + '  ') + `${indent}</${key}>\n`
        } else {
          xml += value + `</${key}>\n`
        }
      }
    }
    return xml
  }
  const content = `<?xml version="1.0" encoding="UTF-8"?>\n<data>\n${objToXml({ items: data }, '  ')}</data>`
  const blob = new Blob([content], { type: 'application/xml' })
  saveAs(blob, `${filename}.xml`)
}

const exportToCsv = (data: any[], filename: string) => {
  if (data.length === 0) return
  
  const headers = Object.keys(data[0])
  const rows = data.map(row => headers.map(header => {
    const value = row[header]
    const escaped = String(value).replace(/"/g, '""')
    return `"${escaped}"`
  }).join(','))
  
  const content = [headers.join(','), ...rows].join('\n')
  const blob = new Blob([`\uFEFF${content}`], { type: 'text/csv;charset=utf-8' })
  saveAs(blob, `${filename}.csv`)
}

const exportToTxt = (data: any[], filename: string) => {
  const content = data.map((row, index) => {
    return `--- 第 ${index + 1} 条记录 ---\n` + 
      Object.entries(row).map(([key, value]) => `${key}: ${value}`).join('\n')
  }).join('\n\n')
  const blob = new Blob([content], { type: 'text/plain;charset=utf-8' })
  saveAs(blob, `${filename}.txt`)
}

const exportToWord = (data: any[], filename: string) => {
  if (data.length === 0) return
  
  const headers = Object.keys(data[0])
  const html = `
    <!DOCTYPE html>
    <html>
    <head>
      <meta charset="UTF-8">
      <style>
        table { border-collapse: collapse; width: 100%; }
        th, td { border: 1px solid #ddd; padding: 8px; text-align: left; }
        th { background-color: #f2f2f2; }
      </style>
    </head>
    <body>
      <h1>${filename}</h1>
      <table>
        <thead><tr>${headers.map(h => `<th>${h}</th>`).join('')}</tr></thead>
        <tbody>${data.map(row => `<tr>${headers.map(h => `<td>${row[h]}</td>`).join('')}</tr>`).join('')}</tbody>
      </table>
    </body>
    </html>
  `
  const blob = new Blob(['\uFEFF' + html], { type: 'application/msword' })
  saveAs(blob, `${filename}.doc`)
}

const exportToExcel = (data: any[], filename: string) => {
  if (data.length === 0) return
  
  const worksheet = XLSX.utils.json_to_sheet(data)
  const workbook = XLSX.utils.book_new()
  XLSX.utils.book_append_sheet(workbook, worksheet, 'Data')
  XLSX.writeFile(workbook, `${filename}.xlsx`)
}

const exportMenuItems: MenuProps['items'] = [
  {
    key: 'json',
    icon: <FileOutlined />,
    label: 'JSON'
  },
  {
    key: 'xml',
    icon: <FileTextOutlined />,
    label: 'XML'
  },
  {
    key: 'csv',
    icon: <FileTextOutlined />,
    label: 'CSV'
  },
  {
    key: 'txt',
    icon: <FileTextOutlined />,
    label: 'TXT'
  },
  {
    key: 'word',
    icon: <FileWordOutlined />,
    label: 'MS Word'
  },
  {
    key: 'excel',
    icon: <FileExcelOutlined />,
    label: 'MS Excel'
  }
]

export default function ExportDropdown({ data, filename }: ExportDropdownProps) {
  const handleExport: MenuProps['onClick'] = ({ key }) => {
    switch (key) {
      case 'json':
        exportToJson(data, filename)
        break
      case 'xml':
        exportToXml(data, filename)
        break
      case 'csv':
        exportToCsv(data, filename)
        break
      case 'txt':
        exportToTxt(data, filename)
        break
      case 'word':
        exportToWord(data, filename)
        break
      case 'excel':
        exportToExcel(data, filename)
        break
    }
  }

  return (
    <Dropdown
      menu={{ items: exportMenuItems, onClick: handleExport }}
    >
      <Button icon={<DownloadOutlined />} style={{ backgroundColor: '#1890ff', color: '#fff', borderColor: '#1890ff' }}>
        导出
      </Button>
    </Dropdown>
  )
}