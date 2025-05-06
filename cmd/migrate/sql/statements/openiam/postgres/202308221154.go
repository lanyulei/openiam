package postgres

/*
  @Author : lanyulei
  @Desc :
*/

const SQL202308221154 = `-- 开始同步

INSERT INTO public.system_api (id, title, url, method, "group", no_forensics, remark, create_time, update_time) VALUES (158, '批量创建 API', '/api/v1/system/api/batch', 'POST', 1, false, '', now(), now());
INSERT INTO public.system_api (id, title, url, method, "group", no_forensics, remark, create_time, update_time) VALUES (160, '通过用户名更新用户信息', '/api/v1/system/user/details/:username', 'PUT', 1, false, '', now(), now());
INSERT INTO public.system_api (id, title, url, method, "group", no_forensics, remark, create_time, update_time) VALUES (161, '通过用户名获取用户详情', '/api/v1/system/user/details/:username', 'GET', 1, false, '', now(), now());
INSERT INTO public.system_api (id, title, url, method, "group", no_forensics, remark, create_time, update_time) VALUES (163, '更新工单评分', '/api/v1/workflow/order/score', 'POST', 2, false, '', now(), now());
SELECT setval('system_api_id_seq', (SELECT MAX(id) FROM public.system_api));

INSERT INTO public.system_menu_api (id, menu, api, create_time, update_time) VALUES (203, 51, 163, now(), now());
SELECT setval('system_menu_api_id_seq', (SELECT MAX(id) FROM public.system_menu_api));

-- 结束同步`
