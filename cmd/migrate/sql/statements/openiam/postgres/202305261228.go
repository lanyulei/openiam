package postgres

/*
  @Author : lanyulei
  @Desc :
*/

const SQL202305261228 = `-- 开始同步

INSERT INTO public.system_menu (id, "path", "name", component, redirect, title, hyperlink, is_hide, is_keep_alive, is_affix, is_iframe, auth, icon, parent, "type", sort, app_id, create_time, update_time) VALUES (113, '', '', '', '', '删除用户分组关联', '', false, false, false, false, '["system:userGroupRelated:delete"]', '', 108, 3, 0, 1, now(), now());
INSERT INTO public.system_menu (id, "path", "name", component, redirect, title, hyperlink, is_hide, is_keep_alive, is_affix, is_iframe, auth, icon, parent, "type", sort, app_id, create_time, update_time) VALUES (112, '', '', '', '', '新建用户分组关联', '', false, false, false, false, '["system:userGroupRelated:add"]', '', 108, 3, 0, 1, now(), now());
INSERT INTO public.system_menu (id, "path", "name", component, redirect, title, hyperlink, is_hide, is_keep_alive, is_affix, is_iframe, auth, icon, parent, "type", sort, app_id, create_time, update_time) VALUES (111, '', '', '', '', '删除用户分组', '', false, false, false, false, '["system:userGroup:delete"]', '', 108, 3, 0, 1, now(), now());
INSERT INTO public.system_menu (id, "path", "name", component, redirect, title, hyperlink, is_hide, is_keep_alive, is_affix, is_iframe, auth, icon, parent, "type", sort, app_id, create_time, update_time) VALUES (110, '', '', '', '', '编辑用户分组', '', false, false, false, false, '["system:userGroup:edit"]', '', 108, 3, 0, 1, now(), now());
INSERT INTO public.system_menu (id, "path", "name", component, redirect, title, hyperlink, is_hide, is_keep_alive, is_affix, is_iframe, auth, icon, parent, "type", sort, app_id, create_time, update_time) VALUES (109, '', '', '', '', '新建用户分组', '', false, false, false, false, '["system:userGroup:create"]', '', 108, 3, 0, 1, now(), now());
INSERT INTO public.system_menu (id, "path", "name", component, redirect, title, hyperlink, is_hide, is_keep_alive, is_affix, is_iframe, auth, icon, parent, "type", sort, app_id, create_time, update_time) VALUES (108, '/workflow/user-group', 'WorkflowUserGroup', 'workflow/userGroup/index', '', '用户分组', '', false, true, false, false, '["workflow:user:group"]', 'ele-User', 40, 2, 1, 1, now(), now());
SELECT setval('system_menu_id_seq', (SELECT MAX(id) FROM public.system_menu));

INSERT INTO public.system_api (id, title, url, method, "group", remark, create_time, update_time) VALUES (148, '通过分组ID获取对应的用户列表', '/api/v1/system/user/group', 'GET', 1, '', now(), now());
INSERT INTO public.system_api (id, title, url, method, "group", remark, create_time, update_time) VALUES (147, '删除用户与分组的关联', '/api/v1/system/user-group-related', 'DELETE', 1, '', now(), now());
INSERT INTO public.system_api (id, title, url, method, "group", remark, create_time, update_time) VALUES (146, '删除用户分组', '/api/v1/system/user-group/:id', 'DELETE', 1, '', now(), now());
INSERT INTO public.system_api (id, title, url, method, "group", remark, create_time, update_time) VALUES (145, '更新用户分组信息', '/api/v1/system/user-group/:id', 'PUT', 1, '', now(), now());
INSERT INTO public.system_api (id, title, url, method, "group", remark, create_time, update_time) VALUES (144, '创建用户与分组的关联', '/api/v1/system/user-group-related', 'POST', 1, '', now(), now());
INSERT INTO public.system_api (id, title, url, method, "group", remark, create_time, update_time) VALUES (143, '创建用户分组', '/api/v1/system/user-group', 'POST', 1, '', now(), now());
INSERT INTO public.system_api (id, title, url, method, "group", remark, create_time, update_time) VALUES (142, '通过多个ID获取用户分组列表', '/api/v1/system/user-group/list', 'GET', 1, '', now(), now());
INSERT INTO public.system_api (id, title, url, method, "group", remark, create_time, update_time) VALUES (141, '获取用户分组列表', '/api/v1/system/user-group', 'GET', 1, '', now(), now());
INSERT INTO public.system_api (id, title, url, method, "group", remark, create_time, update_time) VALUES (140, '通过多个ID获取用户列表', '/api/v1/system/user/list', 'GET', 1, '', now(), now());
SELECT setval('system_api_id_seq', (SELECT MAX(id) FROM public.system_api));

INSERT INTO public.system_menu_api (id, menu, api, create_time, update_time) VALUES (170, 108, 142, now(), now());
INSERT INTO public.system_menu_api (id, menu, api, create_time, update_time) VALUES (169, 108, 148, now(), now());
INSERT INTO public.system_menu_api (id, menu, api, create_time, update_time) VALUES (168, 51, 140, now(), now());
INSERT INTO public.system_menu_api (id, menu, api, create_time, update_time) VALUES (167, 51, 142, now(), now());
INSERT INTO public.system_menu_api (id, menu, api, create_time, update_time) VALUES (166, 43, 142, now(), now());
INSERT INTO public.system_menu_api (id, menu, api, create_time, update_time) VALUES (164, 108, 143, now(), now());
INSERT INTO public.system_menu_api (id, menu, api, create_time, update_time) VALUES (163, 108, 144, now(), now());
INSERT INTO public.system_menu_api (id, menu, api, create_time, update_time) VALUES (162, 108, 145, now(), now());
INSERT INTO public.system_menu_api (id, menu, api, create_time, update_time) VALUES (161, 108, 146, now(), now());
INSERT INTO public.system_menu_api (id, menu, api, create_time, update_time) VALUES (160, 108, 147, now(), now());
SELECT setval('system_menu_api_id_seq', (SELECT MAX(id) FROM public.system_menu_api));

-- 结束同步`
