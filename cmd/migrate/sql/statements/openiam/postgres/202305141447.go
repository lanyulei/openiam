package postgres

/*
  @Author : lanyulei
  @Desc :
*/

const SQL202305141447 = `-- 开始同步

UPDATE public.system_api SET "group" = 2::integer WHERE "group" in (3::bigint, 4::bigint);
DELETE FROM public.system_api_group WHERE id in (3::bigint, 4::bigint);

INSERT INTO public.system_api_group (id, "name", sort, remark, create_time, update_time) VALUES (5, '资源管理', 10, '', now(), now());

SELECT setval('system_api_group_id_seq', (SELECT MAX(id) FROM public.system_api_group));

INSERT INTO public.system_api (id, title, url, method, "group", remark, create_time, update_time) VALUES (109, '获取被指派人列表', '/api/v1/workflow/flow/assign-user', 'GET', 2, '', now(), now());
INSERT INTO public.system_api (id, title, url, method, "group", remark, create_time, update_time) VALUES (110, '获取需要导出的数据', '/api/v1/workflow/order/export', 'POST', 2, '', now(), now());

SELECT setval('system_api_id_seq', (SELECT MAX(id) FROM public.system_api));

INSERT INTO public.system_menu_api (id, menu, api, create_time, update_time) VALUES (127, 51, 109, now(), now());
INSERT INTO public.system_menu_api (id, menu, api, create_time, update_time) VALUES (128, 39, 110, now(), now());

SELECT setval('system_menu_api_id_seq', (SELECT MAX(id) FROM public.system_menu_api));

-- 结束同步`
