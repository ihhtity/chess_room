USE chess_room;

INSERT IGNORE INTO admins (username, password, realname, role, status, created_at, updated_at) VALUES
('admin', '$2a$10$N9qo8uLOickgx2ZMRZoMye.IjzqAKL9xL5jvMFVdNJHvGCgTq/VEq', '管理员', 1, 1, NOW(), NOW()),
('admin2', '$2a$10$N9qo8uLOickgx2ZMRZoMye.IjzqAKL9xL5jvMFVdNJHvGCgTq/VEq', '李经理', 1, 1, NOW(), NOW()),
('admin3', '$2a$10$N9qo8uLOickgx2ZMRZoMye.IjzqAKL9xL5jvMFVdNJHvGCgTq/VEq', '王主管', 2, 1, NOW(), NOW()),
('admin4', '$2a$10$N9qo8uLOickgx2ZMRZoMye.IjzqAKL9xL5jvMFVdNJHvGCgTq/VEq', '张店长', 1, 1, NOW(), NOW()),
('admin5', '$2a$10$N9qo8uLOickgx2ZMRZoMye.IjzqAKL9xL5jvMFVdNJHvGCgTq/VEq', '陈助理', 2, 1, NOW(), NOW()),
('admin6', '$2a$10$N9qo8uLOickgx2ZMRZoMye.IjzqAKL9xL5jvMFVdNJHvGCgTq/VEq', '刘会计', 2, 1, NOW(), NOW()),
('admin7', '$2a$10$N9qo8uLOickgx2ZMRZoMye.IjzqAKL9xL5jvMFVdNJHvGCgTq/VEq', '赵销售', 3, 1, NOW(), NOW()),
('admin8', '$2a$10$N9qo8uLOickgx2ZMRZoMye.IjzqAKL9xL5jvMFVdNJHvGCgTq/VEq', '孙客服', 3, 1, NOW(), NOW()),
('admin9', '$2a$10$N9qo8uLOickgx2ZMRZoMye.IjzqAKL9xL5jvMFVdNJHvGCgTq/VEq', '周后勤', 3, 0, NOW(), NOW()),
('admin10', '$2a$10$N9qo8uLOickgx2ZMRZoMye.IjzqAKL9xL5jvMFVdNJHvGCgTq/VEq', '吴保洁', 3, 1, NOW(), NOW());

INSERT IGNORE INTO room_types (name, description, base_price, max_people, sort_order, status, created_at, updated_at) VALUES
('小包间', '适合2-4人', 88.00, 4, 1, 1, NOW(), NOW()),
('中包间', '适合4-6人', 128.00, 6, 2, 1, NOW(), NOW()),
('大包间', '适合6-10人', 188.00, 10, 3, 1, NOW(), NOW()),
('VIP包间', '豪华装修，适合商务聚会', 288.00, 12, 4, 1, NOW(), NOW()),
('棋牌室', '适合多人娱乐，配备麻将桌', 258.00, 8, 5, 1, NOW(), NOW()),
('茶室', '安静雅致，适合品茶聊天', 168.00, 6, 6, 1, NOW(), NOW()),
('会议室', '商务会议专用，配备投影设备', 388.00, 15, 7, 1, NOW(), NOW()),
('桌游室', '桌游爱好者聚集地', 128.00, 6, 8, 1, NOW(), NOW()),
('电竞室', '配备高性能电脑和电竞设备', 198.00, 4, 9, 1, NOW(), NOW()),
('KTV包间', '卡拉OK娱乐', 328.00, 10, 10, 0, NOW(), NOW());

INSERT IGNORE INTO rooms (name, type_id, floor, capacity, equipment, images, description, status, sort_order, created_at, updated_at) VALUES
('清雅阁', 1, '1F', 4, '空调,WiFi,茶具', 'https://neeko-copilot.bytedance.net/api/text_to_image?prompt=elegant%20small%20chess%20room%20interior%20with%20wooden%20furniture&image_size=square', '环境优雅，适合朋友聚会', 1, 1, NOW(), NOW()),
('怡心轩', 1, '1F', 4, '空调,WiFi,茶具', 'https://neeko-copilot.bytedance.net/api/text_to_image?prompt=cozy%20chess%20room%20with%20comfortable%20seating&image_size=square', '温馨舒适，适合休闲娱乐', 1, 2, NOW(), NOW()),
('聚贤堂', 2, '2F', 6, '空调,WiFi,茶具,投影仪', 'https://neeko-copilot.bytedance.net/api/text_to_image?prompt=medium%20chess%20room%20with%20projection%20screen&image_size=square', '空间宽敞，适合家庭聚会', 1, 3, NOW(), NOW()),
('鸿福厅', 2, '2F', 6, '空调,WiFi,茶具,投影仪', 'https://neeko-copilot.bytedance.net/api/text_to_image?prompt=spacious%20chess%20room%20with%20luxury%20decor&image_size=square', '豪华装修，适合商务洽谈', 2, 4, NOW(), NOW()),
('帝王阁', 3, '3F', 10, '空调,WiFi,茶具,投影仪,音响', 'https://neeko-copilot.bytedance.net/api/text_to_image?prompt=large%20luxury%20chess%20room%20with%20audio%20system&image_size=square', '高端大气，适合大型聚会', 1, 5, NOW(), NOW()),
('至尊殿', 4, '3F', 12, '空调,WiFi,茶具,投影仪,音响,麻将机', 'https://neeko-copilot.bytedance.net/api/text_to_image?prompt=VIP%20luxury%20chess%20room%20with%20mahjong%20table&image_size=square', '顶级配置，尊享体验', 1, 6, NOW(), NOW()),
('清风苑', 5, '1F', 8, '空调,WiFi,麻将机,茶具', 'https://neeko-copilot.bytedance.net/api/text_to_image?prompt=traditional%20mahjong%20room%20with%20chinese%20decor&image_size=square', '传统棋牌室风格', 1, 7, NOW(), NOW()),
('茗香阁', 6, '2F', 6, '空调,WiFi,茶具,茶叶', 'https://neeko-copilot.bytedance.net/api/text_to_image?prompt=chinese%20tea%20room%20with%20traditional%20style&image_size=square', '茶香四溢，品味人生', 1, 8, NOW(), NOW()),
('商务中心', 7, '3F', 15, '空调,WiFi,投影仪,白板,音响', 'https://neeko-copilot.bytedance.net/api/text_to_image?prompt=modern%20business%20meeting%20room&image_size=square', '专业商务会议环境', 1, 9, NOW(), NOW()),
('欢乐屋', 8, '1F', 6, '空调,WiFi,桌游', 'https://neeko-copilot.bytedance.net/api/text_to_image?prompt=colorful%20board%20game%20room&image_size=square', '欢乐桌游，益智娱乐', 0, 10, NOW(), NOW());

INSERT IGNORE INTO time_slots (type_id, name, start_time, end_time, price, weekday_price, weekend_price, holiday_price, sort_order, status, created_at, updated_at) VALUES
(1, '上午场', '09:00', '12:00', 88.00, 88.00, 108.00, 128.00, 1, 1, NOW(), NOW()),
(1, '下午场', '13:00', '17:00', 108.00, 108.00, 128.00, 148.00, 2, 1, NOW(), NOW()),
(1, '晚场', '18:00', '22:00', 128.00, 128.00, 148.00, 168.00, 3, 1, NOW(), NOW()),
(2, '上午场', '09:00', '12:00', 128.00, 128.00, 148.00, 168.00, 1, 1, NOW(), NOW()),
(2, '下午场', '13:00', '17:00', 148.00, 148.00, 168.00, 188.00, 2, 1, NOW(), NOW()),
(2, '晚场', '18:00', '22:00', 168.00, 168.00, 188.00, 208.00, 3, 1, NOW(), NOW()),
(3, '上午场', '09:00', '12:00', 188.00, 188.00, 208.00, 228.00, 1, 1, NOW(), NOW()),
(3, '下午场', '13:00', '17:00', 208.00, 208.00, 228.00, 248.00, 2, 1, NOW(), NOW()),
(3, '晚场', '18:00', '22:00', 228.00, 228.00, 248.00, 268.00, 3, 1, NOW(), NOW()),
(4, '全天场', '09:00', '22:00', 288.00, 288.00, 328.00, 368.00, 1, 1, NOW(), NOW());

INSERT IGNORE INTO users (openid, phone, password, nickname, realname, avatar, gender, status, created_at, updated_at) VALUES
('openid_001', '13800138001', '$2a$10$N9qo8uLOickgx2ZMRZoMye.IjzqAKL9xL5jvMFVdNJHvGCgTq/VEq', '张三', '张三', 'https://neeko-copilot.bytedance.net/api/text_to_image?prompt=asian%20male%20avatar%20portrait&image_size=square', 1, 1, NOW(), NOW()),
('openid_002', '13800138002', '$2a$10$N9qo8uLOickgx2ZMRZoMye.IjzqAKL9xL5jvMFVdNJHvGCgTq/VEq', '李四', '李四', 'https://neeko-copilot.bytedance.net/api/text_to_image?prompt=asian%20female%20avatar%20portrait&image_size=square', 2, 1, NOW(), NOW()),
('openid_003', '13800138003', '$2a$10$N9qo8uLOickgx2ZMRZoMye.IjzqAKL9xL5jvMFVdNJHvGCgTq/VEq', '王五', '王五', 'https://neeko-copilot.bytedance.net/api/text_to_image?prompt=asian%20male%20business%20avatar&image_size=square', 1, 1, NOW(), NOW()),
('openid_004', '13800138004', '$2a$10$N9qo8uLOickgx2ZMRZoMye.IjzqAKL9xL5jvMFVdNJHvGCgTq/VEq', '赵六', '赵六', 'https://neeko-copilot.bytedance.net/api/text_to_image?prompt=asian%20female%20elegant%20avatar&image_size=square', 2, 1, NOW(), NOW()),
('openid_005', '13800138005', '$2a$10$N9qo8uLOickgx2ZMRZoMye.IjzqAKL9xL5jvMFVdNJHvGCgTq/VEq', '钱七', '钱七', 'https://neeko-copilot.bytedance.net/api/text_to_image?prompt=asian%20male%20young%20avatar&image_size=square', 1, 1, NOW(), NOW()),
('openid_006', '13800138006', '$2a$10$N9qo8uLOickgx2ZMRZoMye.IjzqAKL9xL5jvMFVdNJHvGCgTq/VEq', '孙八', '孙八', 'https://neeko-copilot.bytedance.net/api/text_to_image?prompt=asian%20female%20professional%20avatar&image_size=square', 2, 1, NOW(), NOW()),
('openid_007', '13800138007', '$2a$10$N9qo8uLOickgx2ZMRZoMye.IjzqAKL9xL5jvMFVdNJHvGCgTq/VEq', '周九', '周九', 'https://neeko-copilot.bytedance.net/api/text_to_image?prompt=asian%20male%20casual%20avatar&image_size=square', 1, 1, NOW(), NOW()),
('openid_008', '13800138008', '$2a$10$N9qo8uLOickgx2ZMRZoMye.IjzqAKL9xL5jvMFVdNJHvGCgTq/VEq', '吴十', '吴十', 'https://neeko-copilot.bytedance.net/api/text_to_image?prompt=asian%20female%20young%20avatar&image_size=square', 2, 0, NOW(), NOW()),
('openid_009', '13800138009', '$2a$10$N9qo8uLOickgx2ZMRZoMye.IjzqAKL9xL5jvMFVdNJHvGCgTq/VEq', '郑十一', '郑十一', 'https://neeko-copilot.bytedance.net/api/text_to_image?prompt=asian%20male%20senior%20avatar&image_size=square', 1, 1, NOW(), NOW()),
('openid_010', '13800138010', '$2a$10$N9qo8uLOickgx2ZMRZoMye.IjzqAKL9xL5jvMFVdNJHvGCgTq/VEq', '王十二', '王十二', 'https://neeko-copilot.bytedance.net/api/text_to_image?prompt=asian%20female%20elegant%20avatar&image_size=square', 2, 1, NOW(), NOW());

INSERT IGNORE INTO memberships (user_id, level, points, balance, total_consumed, total_recharged, membership_status, joined_at, expired_at, created_at, updated_at) VALUES
(1, 2, 1500, 500.00, 1200.00, 1000.00, 1, NOW(), DATE_ADD(NOW(), INTERVAL 1 YEAR), NOW(), NOW()),
(2, 1, 800, 200.00, 600.00, 500.00, 1, NOW(), DATE_ADD(NOW(), INTERVAL 6 MONTH), NOW(), NOW()),
(3, 3, 3000, 1500.00, 3500.00, 2000.00, 1, NOW(), DATE_ADD(NOW(), INTERVAL 2 YEAR), NOW(), NOW()),
(4, 0, 200, 50.00, 300.00, 100.00, 1, NOW(), DATE_ADD(NOW(), INTERVAL 3 MONTH), NOW(), NOW()),
(5, 1, 600, 100.00, 400.00, 200.00, 1, NOW(), DATE_ADD(NOW(), INTERVAL 6 MONTH), NOW(), NOW()),
(6, 2, 1200, 400.00, 900.00, 800.00, 1, NOW(), DATE_ADD(NOW(), INTERVAL 1 YEAR), NOW(), NOW()),
(7, 0, 100, 20.00, 150.00, 50.00, 1, NOW(), DATE_ADD(NOW(), INTERVAL 1 MONTH), NOW(), NOW()),
(8, 3, 5000, 3000.00, 6000.00, 5000.00, 2, NOW(), DATE_SUB(NOW(), INTERVAL 1 DAY), NOW(), NOW()),
(9, 1, 700, 150.00, 500.00, 300.00, 1, NOW(), DATE_ADD(NOW(), INTERVAL 4 MONTH), NOW(), NOW()),
(10, 2, 1800, 600.00, 1500.00, 1200.00, 1, NOW(), DATE_ADD(NOW(), INTERVAL 1 YEAR), NOW(), NOW());

INSERT IGNORE INTO activities (name, description, image, discount, valid_from, valid_to, status, sort_order, created_at, updated_at) VALUES
('新会员首单优惠', '新注册会员享受8折优惠', 'https://neeko-copilot.bytedance.net/api/text_to_image?prompt=promotion%20banner%20for%20new%20member%20discount&image_size=landscape_16_9', 0.80, NOW(), DATE_ADD(NOW(), INTERVAL 30 DAY), 1, 1, NOW(), NOW()),
('周末特惠', '周末包间全场9折', 'https://neeko-copilot.bytedance.net/api/text_to_image?prompt=weekend%20special%20promotion%20banner&image_size=landscape_16_9', 0.90, DATE_ADD(NOW(), INTERVAL 1 DAY), DATE_ADD(NOW(), INTERVAL 90 DAY), 1, 2, NOW(), NOW()),
('会员日活动', '每月8号会员享受双倍积分', 'https://neeko-copilot.bytedance.net/api/text_to_image?prompt=member%20day%20promotion%20banner&image_size=landscape_16_9', 1.00, NOW(), DATE_ADD(NOW(), INTERVAL 1 YEAR), 1, 3, NOW(), NOW()),
('夏季清凉活动', '暑期包间享7折优惠', 'https://neeko-copilot.bytedance.net/api/text_to_image?prompt=summer%20cool%20promotion%20banner&image_size=landscape_16_9', 0.70, DATE_ADD(NOW(), INTERVAL 10 DAY), DATE_ADD(NOW(), INTERVAL 60 DAY), 1, 4, NOW(), NOW()),
('生日特权', '生日当天包间免费3小时', 'https://neeko-copilot.bytedance.net/api/text_to_image?prompt=birthday%20privilege%20promotion%20banner&image_size=landscape_16_9', 0.50, NOW(), DATE_ADD(NOW(), INTERVAL 1 YEAR), 1, 5, NOW(), NOW()),
('团建优惠', '团队预订10人以上享85折', 'https://neeko-copilot.bytedance.net/api/text_to_image?prompt=team%20building%20discount%20banner&image_size=landscape_16_9', 0.85, NOW(), DATE_ADD(NOW(), INTERVAL 30 DAY), 1, 6, NOW(), NOW()),
('老带新活动', '推荐好友注册各得50积分', 'https://neeko-copilot.bytedance.net/api/text_to_image?prompt=referral%20program%20promotion%20banner&image_size=landscape_16_9', 1.00, NOW(), DATE_ADD(NOW(), INTERVAL 90 DAY), 1, 7, NOW(), NOW()),
('节假日特惠', '法定节假日包间8折', 'https://neeko-copilot.bytedance.net/api/text_to_image?prompt=holiday%20special%20discount%20banner&image_size=landscape_16_9', 0.80, NOW(), DATE_ADD(NOW(), INTERVAL 1 YEAR), 0, 8, NOW(), NOW()),
('深夜优惠', '22:00后包间半价', 'https://neeko-copilot.bytedance.net/api/text_to_image?prompt=late%20night%20discount%20promotion%20banner&image_size=landscape_16_9', 0.50, NOW(), DATE_ADD(NOW(), INTERVAL 60 DAY), 1, 9, NOW(), NOW()),
('周年庆活动', '开业周年庆全场6折', 'https://neeko-copilot.bytedance.net/api/text_to_image?prompt=anniversary%20celebration%20promotion%20banner&image_size=landscape_16_9', 0.60, DATE_ADD(NOW(), INTERVAL 20 DAY), DATE_ADD(NOW(), INTERVAL 30 DAY), 0, 10, NOW(), NOW());

INSERT IGNORE INTO announcements (title, content, type, status, sort_order, created_at, updated_at) VALUES
('春节营业时间调整通知', '春节期间营业时间调整为10:00-21:00，请各位顾客提前预约。', 1, 1, 1, NOW(), NOW()),
('新店开业优惠', '新店开业期间，所有包间享受7折优惠，欢迎光临！', 0, 1, 2, NOW(), NOW()),
('温馨提示', '请在离开时将物品带走，感谢您的理解与配合。', 0, 1, 3, NOW(), NOW()),
('暑期活动上线', '暑期清凉活动已上线，包间享7折优惠，详情请查看活动页面。', 0, 1, 4, NOW(), NOW()),
('会员系统升级', '会员系统将于本周日进行升级维护，期间可能影响会员权益查询，请提前知悉。', 1, 1, 5, NOW(), NOW()),
('设施维护通知', '清雅阁包间将于明天进行空调维护，暂停使用一天，给您带来不便敬请谅解。', 1, 1, 6, NOW(), NOW()),
('安全提醒', '为了您的安全，请不要在包间内吸烟，感谢配合！', 0, 1, 7, NOW(), NOW()),
('积分兑换活动', '积分兑换活动开始啦！1000积分可兑换50元代金券，限时一周！', 0, 1, 8, NOW(), NOW()),
('中秋特惠', '中秋佳节期间，包间预订享受8折优惠，提前预订更有好礼相送！', 0, 1, 9, NOW(), NOW()),
('停车场调整', '停车场入口临时调整，请按指示牌行驶，不便之处敬请谅解。', 0, 1, 10, NOW(), NOW());

INSERT IGNORE INTO recharge_packages (name, amount, gift_amount, gift_points, description, sort_order, status, created_at, updated_at) VALUES
('新手礼包', 100.00, 10.00, 100, '适合初次充值的用户', 1, 1, NOW(), NOW()),
('超值套餐', 200.00, 30.00, 200, '赠送30元，性价比超高', 2, 1, NOW(), NOW()),
('豪华套餐', 500.00, 100.00, 500, '赠送100元，尊享特权', 3, 1, NOW(), NOW()),
('至尊套餐', 1000.00, 250.00, 1000, '赠送250元，成为VIP会员', 4, 1, NOW(), NOW()),
('小额充值', 50.00, 5.00, 50, '少量充值，灵活使用', 5, 1, NOW(), NOW()),
('季度套餐', 300.00, 50.00, 300, '季度优惠，划算之选', 6, 1, NOW(), NOW()),
('半年套餐', 600.00, 150.00, 600, '半年畅玩，超值优惠', 7, 1, NOW(), NOW()),
('年度套餐', 1200.00, 350.00, 1200, '年度会员，尊享全年', 8, 1, NOW(), NOW()),
('商务套餐', 2000.00, 600.00, 2000, '商务人士首选', 9, 1, NOW(), NOW()),
('家庭套餐', 800.00, 200.00, 800, '家庭娱乐，欢乐共享', 10, 0, NOW(), NOW());

INSERT IGNORE INTO orders (order_no, user_id, room_id, start_time, end_time, duration, status, total_amount, paid_amount, remark, paid_at, created_at, updated_at) VALUES
('ORD202607020001', 1, 1, '2026-07-02 14:00:00', '2026-07-02 18:00:00', 240, 2, 108.00, 108.00, '', NOW(), NOW(), NOW()),
('ORD202607020002', 2, 3, '2026-07-02 18:00:00', '2026-07-02 22:00:00', 240, 1, 168.00, 168.00, '朋友聚会', NOW(), NOW(), NOW()),
('ORD202607030001', 3, 5, '2026-07-03 10:00:00', '2026-07-03 14:00:00', 240, 0, 208.00, 0.00, '', NULL, NOW(), NOW()),
('ORD202607030002', 1, 2, '2026-07-03 14:00:00', '2026-07-03 18:00:00', 240, 0, 108.00, 0.00, '', NULL, NOW(), NOW()),
('ORD202607010001', 4, 4, '2026-07-01 18:00:00', '2026-07-01 22:00:00', 240, 3, 168.00, 168.00, '临时有事取消', NOW(), NOW(), NOW()),
('ORD202607040001', 5, 6, '2026-07-04 09:00:00', '2026-07-04 22:00:00', 780, 2, 288.00, 288.00, '全天使用', NOW(), NOW(), NOW()),
('ORD202607040002', 6, 7, '2026-07-04 14:00:00', '2026-07-04 18:00:00', 240, 1, 258.00, 258.00, '麻将娱乐', NOW(), NOW(), NOW()),
('ORD202607050001', 7, 8, '2026-07-05 10:00:00', '2026-07-05 14:00:00', 240, 0, 168.00, 0.00, '', NULL, NOW(), NOW()),
('ORD202607050002', 9, 9, '2026-07-05 14:00:00', '2026-07-05 17:00:00', 180, 2, 388.00, 388.00, '商务会议', NOW(), NOW(), NOW()),
('ORD202607060001', 10, 1, '2026-07-06 18:00:00', '2026-07-06 22:00:00', 240, 4, 128.00, 128.00, '申请退款', NOW(), NOW(), NOW());

INSERT IGNORE INTO payments (order_id, user_id, amount, payment_type, status, transaction_no, paid_at, created_at, updated_at) VALUES
(1, 1, 108.00, 1, 1, 'TXN202607020001', NOW(), NOW(), NOW()),
(2, 2, 168.00, 1, 1, 'TXN202607020002', NOW(), NOW(), NOW()),
(5, 4, 168.00, 1, 1, 'TXN202607010001', NOW(), NOW(), NOW()),
(6, 5, 288.00, 2, 1, 'TXN202607040001', NOW(), NOW(), NOW()),
(7, 6, 258.00, 1, 1, 'TXN202607040002', NOW(), NOW(), NOW()),
(9, 9, 388.00, 1, 1, 'TXN202607050002', NOW(), NOW(), NOW()),
(10, 10, 128.00, 3, 2, 'TXN202607060001', NOW(), NOW(), NOW()),
(3, 3, 208.00, 1, 0, 'TXN202607030001', NULL, NOW(), NOW()),
(4, 1, 108.00, 2, 0, 'TXN202607030002', NULL, NOW(), NOW()),
(8, 7, 168.00, 1, 0, 'TXN202607050001', NULL, NOW(), NOW());

INSERT IGNORE INTO reviews (order_id, user_id, room_id, rating, content, images, status, created_at, updated_at) VALUES
(1, 1, 1, 5, '环境非常好，服务周到，下次还会再来！', '', 1, NOW(), NOW()),
(2, 2, 3, 5, '朋友聚会选这里很合适，大家都很满意！', '', 1, NOW(), NOW()),
(5, 4, 4, 4, '包间很大，设施齐全，就是价格有点贵。', '', 1, NOW(), NOW()),
(6, 5, 6, 5, '至尊殿确实名不虚传，豪华大气！', '', 1, NOW(), NOW()),
(7, 6, 7, 4, '麻将桌很舒服，玩得很开心！', '', 1, NOW(), NOW()),
(9, 9, 9, 5, '商务会议环境很好，投影仪清晰，音响效果不错。', '', 1, NOW(), NOW()),
(10, 10, 1, 3, '环境不错，但是退款处理有点慢。', '', 1, NOW(), NOW()),
(1, 1, 1, 5, '第二次来了，依然满意！', '', 1, NOW(), NOW()),
(2, 2, 3, 4, '整体不错，就是空调有点冷。', '', 1, NOW(), NOW()),
(7, 6, 7, 5, '强烈推荐，服务态度很好！', '', 1, NOW(), NOW());

INSERT IGNORE INTO recharge_records (user_id, amount, gift_amount, payment_id, status, created_at) VALUES
(1, 500.00, 100.00, 1, 1, NOW()),
(2, 200.00, 30.00, 2, 1, NOW()),
(3, 1000.00, 250.00, 3, 1, NOW()),
(5, 100.00, 10.00, 4, 1, NOW()),
(6, 500.00, 100.00, 5, 1, NOW()),
(9, 200.00, 30.00, 6, 1, NOW()),
(10, 1000.00, 250.00, 7, 2, NOW()),
(4, 50.00, 5.00, 8, 0, NOW()),
(7, 300.00, 50.00, 9, 0, NOW()),
(8, 1200.00, 350.00, 10, 1, NOW());

INSERT IGNORE INTO holidays (name, date, is_holiday, description, created_at, updated_at) VALUES
('元旦', '2026-01-01', 1, '新年第一天', NOW(), NOW()),
('春节', '2026-02-17', 1, '农历新年', NOW(), NOW()),
('清明节', '2026-04-04', 1, '祭祖扫墓', NOW(), NOW()),
('劳动节', '2026-05-01', 1, '五一劳动节', NOW(), NOW()),
('端午节', '2026-06-19', 1, '传统节日', NOW(), NOW()),
('中秋节', '2026-09-25', 1, '团圆佳节', NOW(), NOW()),
('国庆节', '2026-10-01', 1, '国庆假期', NOW(), NOW()),
('除夕', '2026-02-16', 1, '除夕之夜', NOW(), NOW()),
('七夕节', '2026-08-20', 1, '中国情人节', NOW(), NOW()),
('平安夜', '2026-12-24', 0, '圣诞节前夕', NOW(), NOW());

INSERT IGNORE INTO notifications (user_id, type, title, content, read_status, link, created_at) VALUES
(1, 0, '系统维护通知', '系统将于今晚23:00进行维护，预计持续2小时。', 0, '', NOW()),
(1, 1, '订单确认通知', '您的订单 ORD202607020001 已确认，请准时到达。', 1, '/order/1', NOW()),
(2, 1, '订单完成通知', '您的订单 ORD202607020002 已完成，感谢您的光临。', 1, '/order/2', NOW()),
(3, 2, '新活动上线', '周末特惠活动已上线，全场9折优惠！', 0, '/activity/2', NOW()),
(3, 3, '会员积分更新', '您的积分已增加100分，当前积分3100分。', 0, '/member/profile', NOW()),
(4, 0, '会员到期提醒', '您的会员将于3个月后到期，请及时续费。', 0, '/member/renew', NOW()),
(5, 1, '订单取消通知', '您的订单 ORD202607030002 已取消，款项将原路退回。', 0, '/order/4', NOW()),
(0, 0, '重要公告', '春节期间营业时间调整，请留意公告详情。', 0, '/announcement/1', NOW()),
(0, 2, '新店开业', '新店盛大开业，开业期间享受7折优惠！', 0, '/activity/2', NOW()),
(1, 3, '生日特权', '生日快乐！本月享受生日专属折扣。', 0, '/member/birthday', NOW());

INSERT IGNORE INTO operation_logs (admin_id, action, module, target_id, content, ip, created_at) VALUES
(1, 'login', 'admin', 1, '管理员登录系统', '192.168.1.100', NOW()),
(1, 'create', 'room', 7, '创建新包间：清风苑', '192.168.1.100', NOW()),
(1, 'update', 'room', 1, '更新包间信息：清雅阁', '192.168.1.100', NOW()),
(1, 'delete', 'room', 10, '删除包间：欢乐屋', '192.168.1.100', NOW()),
(2, 'create', 'activity', 4, '创建新活动：夏季清凉活动', '192.168.1.101', NOW()),
(2, 'update', 'activity', 1, '更新活动：新会员首单优惠', '192.168.1.101', NOW()),
(2, 'create', 'announcement', 4, '发布公告：暑期活动上线', '192.168.1.101', NOW()),
(3, 'update', 'membership', 1, '提升会员等级：张三 VIP', '192.168.1.102', NOW()),
(3, 'view', 'order', 0, '查看订单列表', '192.168.1.102', NOW()),
(4, 'logout', 'admin', 1, '管理员退出系统', '192.168.1.103', NOW());
