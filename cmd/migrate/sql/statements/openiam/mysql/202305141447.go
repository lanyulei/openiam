package mysql

/*
  @Author : lanyulei
  @Desc :
*/

const SQL202305141447 = `-- 开始同步

UPDATE system_api SET ` + "`" + "group" + "`" + ` = 2 WHERE ` + "`" + "group" + "`" + ` in (3, 4);
DELETE FROM system_api_group WHERE id in (3, 4);

INSERT INTO system_api_group (id, name, ` + "`" + "sort" + "`" + `, remark, create_time, update_time) VALUES (5, '资源管理', 10, '', now(), now());

INSERT INTO system_api (id, title, url, method, ` + "`" + "group" + "`" + `, remark, create_time, update_time) VALUES (109, '获取被指派人列表', '/api/v1/workflow/flow/assign-user', 'GET', 2, '', now(), now());
INSERT INTO system_api (id, title, url, method, ` + "`" + "group" + "`" + `, remark, create_time, update_time) VALUES (110, '获取需要导出的数据', '/api/v1/workflow/order/export', 'POST', 2, '', now(), now());

INSERT INTO system_menu_api (id, menu, api, create_time, update_time) VALUES (127, 51, 109, now(), now());
INSERT INTO system_menu_api (id, menu, api, create_time, update_time) VALUES (128, 39, 110, now(), now());

-- 结束同步`
