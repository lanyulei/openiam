package postgres

/*
  @Author : lanyulei
  @Desc :
*/

const SQL202302250147 = `-- 开始同步

INSERT INTO public.system_api (id, title, url, "method", "group", remark, create_time, update_time) VALUES (106, '初始化用户密码', '/api/v1/system/user/init-password/:id', 'PUT', 1, '', now(), now());
INSERT INTO public.system_api (id, title, url, "method", "group", remark, create_time, update_time) VALUES (107, '更新当前用户密码', '/api/v1/system/user/update-password', 'PUT', 1, '', now(), now());
INSERT INTO public.system_api (id, title, url, "method", "group", remark, create_time, update_time) VALUES (108, '更新全局默认接口权限', '/api/v1/system/api/no-forensics', 'PUT', 1, '', now(), now());

SELECT setval('system_api_id_seq', (SELECT MAX(id) FROM public.system_api));

INSERT INTO public.system_menu_api (id, menu, api, create_time, update_time) VALUES (113, 91, 97, now(), now());
INSERT INTO public.system_menu_api (id, menu, api, create_time, update_time) VALUES (114, 91, 96, now(), now());
INSERT INTO public.system_menu_api (id, menu, api, create_time, update_time) VALUES (115, 91, 98, now(), now());
INSERT INTO public.system_menu_api (id, menu, api, create_time, update_time) VALUES (116, 91, 99, now(), now());
INSERT INTO public.system_menu_api (id, menu, api, create_time, update_time) VALUES (117, 92, 105, now(), now());
INSERT INTO public.system_menu_api (id, menu, api, create_time, update_time) VALUES (118, 92, 104, now(), now());
INSERT INTO public.system_menu_api (id, menu, api, create_time, update_time) VALUES (119, 92, 103, now(), now());
INSERT INTO public.system_menu_api (id, menu, api, create_time, update_time) VALUES (120, 92, 102, now(), now());
INSERT INTO public.system_menu_api (id, menu, api, create_time, update_time) VALUES (121, 92, 101, now(), now());
INSERT INTO public.system_menu_api (id, menu, api, create_time, update_time) VALUES (122, 92, 100, now(), now());
INSERT INTO public.system_menu_api (id, menu, api, create_time, update_time) VALUES (123, 4, 106, now(), now());
INSERT INTO public.system_menu_api (id, menu, api, create_time, update_time) VALUES (124, 54, 107, now(), now());
INSERT INTO public.system_menu_api (id, menu, api, create_time, update_time) VALUES (125, 4, 2, now(), now());
INSERT INTO public.system_menu_api (id, menu, api, create_time, update_time) VALUES (126, 53, 108, now(), now());

SELECT setval('system_menu_api_id_seq', (SELECT MAX(id) FROM public.system_menu_api));

INSERT INTO public.system_menu (id, "path", "name", component, redirect, title, hyperlink, is_hide, is_keep_alive, is_affix, is_iframe, auth, icon, parent, "type", sort, app_id, create_time, update_time) VALUES (93, '', '', '', '', '初始化密码', '', false, false, false, false, '["system:user:initPassword"]', '', 4, 3, 0, 2, now(), now());
INSERT INTO public.system_menu (id, "path", "name", component, redirect, title, hyperlink, is_hide, is_keep_alive, is_affix, is_iframe, auth, icon, parent, "type", sort, app_id, create_time, update_time) VALUES (94, '', '', '', '', '配置全局默认接口权限', '', false, false, false, false, '["system:settings:api-create"]', '', 53, 3, 0, 2, now(), now());
INSERT INTO public.system_menu (id, "path", "name", component, redirect, title, hyperlink, is_hide, is_keep_alive, is_affix, is_iframe, auth, icon, parent, "type", sort, app_id, create_time, update_time) VALUES (95, '', '', '', '', '删除全局默认接口权限', '', false, false, false, false, '["system:settings:api-delete"]', '', 53, 3, 0, 2, now(), now());

SELECT setval('system_menu_id_seq', (SELECT MAX(id) FROM public.system_menu));

UPDATE public.system_menu SET sort = 88::smallint WHERE id = 53::bigint;

UPDATE public.system_api SET no_forensics = true::boolean WHERE id = 2::bigint;
UPDATE public.system_api SET no_forensics = true::boolean WHERE id = 95::bigint;
UPDATE public.system_api SET no_forensics = true::boolean WHERE id = 102::bigint;

-- 结束同步`
