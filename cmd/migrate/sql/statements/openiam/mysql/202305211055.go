package mysql

/*
  @Author : lanyulei
  @Desc :
*/

const SQL202305211055 = `-- 开始同步

INSERT INTO system_menu (id, path, name, component, redirect, title, hyperlink, is_hide, is_keep_alive, is_affix, is_iframe, auth, icon, parent, type, sort, app_id, create_time, update_time) VALUES (106, '', '', '', '', '更新工单', '', false, false, false, false, '["workOrder:list:update"]', '', 39, 3, 0, 1, now(), now());

INSERT INTO system_api (id, title, url, method, ` + "`" + "group" + "`" + `, remark, create_time, update_time) VALUES (137, '更新工单信息', '/api/v1/workflow/order/update', 'PATCH', 2, '', now(), now());
INSERT INTO system_api (id, title, url, method, ` + "`" + "group" + "`" + `, remark, create_time, update_time) VALUES (139, '获取当前用户工单统计数据', '/api/v1/workflow/statistics/workbench', 'GET', 1, '', now(), now());

INSERT INTO system_menu_api (id, menu, api, create_time, update_time) VALUES (157, 1, 139, now(), now());
INSERT INTO system_menu_api (id, menu, api, create_time, update_time) VALUES (158, 39, 137, now(), now());

-- 结束同步`
