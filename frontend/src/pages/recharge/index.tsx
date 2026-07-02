import { useState, useEffect } from 'react'
import { Button, Cell, CellGroup, Loading, Empty, Badge } from '@nutui/nutui-react'
import { showToast } from '@/components/Toast'
import { membershipApi, RechargePackage } from '@/api'

export default function RechargePage() {
  const [packages, setPackages] = useState<RechargePackage[]>([])
  const [loading, setLoading] = useState(true)
  const [selectedId, setSelectedId] = useState<number | null>(null)

  useEffect(() => {
    fetchPackages()
  }, [])

  const fetchPackages = async () => {
    try {
      setLoading(true)
      const data = await membershipApi.getRechargePackages()
      setPackages(data)
    } catch (error) {
      console.error('Failed to fetch recharge packages:', error)
      setPackages([
        { id: 1, name: '新手礼包', amount: 100, gift_amount: 10, gift_points: 100, description: '充值100元送10元', sort_order: 1 },
        { id: 2, name: '银卡套餐', amount: 300, gift_amount: 50, gift_points: 300, description: '充值300元送50元+300积分', sort_order: 2 },
        { id: 3, name: '金卡套餐', amount: 500, gift_amount: 100, gift_points: 500, description: '充值500元送100元+500积分', sort_order: 3 },
        { id: 4, name: '钻石套餐', amount: 1000, gift_amount: 250, gift_points: 1000, description: '充值1000元送250元+1000积分', sort_order: 4 }
      ])
    } finally {
      setLoading(false)
    }
  }

  const handleSelect = (id: number) => {
    setSelectedId(id)
  }

  const handleRecharge = () => {
    if (!selectedId) {
      showToast({ message: '请选择充值套餐', type: 'warning' })
      return
    }
    const selected = packages.find(p => p.id === selectedId)
    if (selected) {
      showToast({ message: `正在充值 ¥${selected.amount}`, type: 'info' })
      setTimeout(() => {
        showToast({ message: '充值成功', type: 'success' })
      }, 1500)
    }
  }

  return (
    <div className="page">
      <div className="header">
        <div className="header-content">
          <span className="title">💳 会员充值</span>
          <span className="subtitle">选择套餐享受更多优惠</span>
        </div>
      </div>

      <div className="container" style={{ overflowY: 'auto', height: 'calc(100vh - 200px)' }}>
        {loading ? (
          <Loading type="circular" color="#667eea" style={{ marginTop: '100rpx' }} />
        ) : (
          <CellGroup>
            {packages.map(pkg => (
              <Cell
                key={pkg.id}
                onClick={() => handleSelect(pkg.id)}
                className="package-cell"
              >
                <div className={`package-card ${selectedId === pkg.id ? 'selected' : ''}`}>
                  <div className="package-header">
                    <span className="package-name">{pkg.name}</span>
                    {selectedId === pkg.id && (
                      <Badge value="已选择" color="#667eea" />
                    )}
                  </div>
                  <div className="package-price">
                    <span className="price-symbol">¥</span>
                    <span className="price-value">{pkg.amount}</span>
                  </div>
                  <div className="package-gifts">
                    {pkg.gift_amount > 0 && (
                      <span className="gift-item">赠送 ¥{pkg.gift_amount}</span>
                    )}
                    {pkg.gift_points > 0 && (
                      <span className="gift-item">赠送 {pkg.gift_points} 积分</span>
                    )}
                  </div>
                  <span className="package-desc">{pkg.description}</span>
                  <div className="package-select">
                    <div className={`select-circle ${selectedId === pkg.id ? 'active' : ''}`}>
                      {selectedId === pkg.id && '✓'}
                    </div>
                  </div>
                </div>
              </Cell>
            ))}
          </CellGroup>
        )}

        {!loading && packages.length === 0 && (
          <Empty description="暂无充值套餐" image="empty" />
        )}
      </div>

      <div className="footer">
        <Button type="primary" block onClick={handleRecharge}>
          确认充值
        </Button>
      </div>
    </div>
  )
}
