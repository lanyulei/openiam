package mysql

/*
  @Author : lanyulei
  @Desc :
*/

const SQL202303311125 = `-- 开始同步

UPDATE system_menu SET path = '/workflow/task/worker', component = 'workflow/task/worker' WHERE id = 48;
UPDATE system_menu SET path = '/workflow/home', component = 'workflow/home/index' WHERE id = 1;
UPDATE system_menu SET path = '/workflow/work-order/list', component = 'workflow/workOrder/list' WHERE id = 39;
UPDATE system_menu SET path = '/workflow/task/history', component = 'workflow/task/history' WHERE id = 47;
UPDATE system_menu SET path = '/workflow/flow/template', component = 'workflow/flow/templates' WHERE id = 42;
UPDATE system_menu SET path = '/workflow/flow/group', component = 'workflow/flow/group' WHERE id = 41;
UPDATE system_menu SET path = '/workflow/work-order/create/:flowId', component = 'workflow/workOrder/create' WHERE id = 49;
UPDATE system_menu SET path = '/workflow/flow/list', component = 'workflow/flow/list' WHERE id = 43;
UPDATE system_menu SET path = '/workflow/work-order/apply', component = 'workflow/workOrder/apply' WHERE id = 38;
UPDATE system_menu SET path = '/workflow/work-order/details/:id', component = 'workflow/workOrder/details' WHERE id = 51;
UPDATE system_menu SET path = '/workflow/task/list', component = 'workflow/task/list' WHERE id = 46;
UPDATE system_app SET link = '/workflow/home' WHERE id = 1;

-- 结束同步`
