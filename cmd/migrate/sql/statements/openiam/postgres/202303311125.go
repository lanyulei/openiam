package postgres

/*
  @Author : lanyulei
  @Desc :
*/

const SQL202303311125 = `-- 开始同步

UPDATE public.system_menu
SET path      = '/workflow/task/worker'::varchar(200),
    component = 'workflow/task/worker'::varchar(200)
WHERE id = 48::bigint;

UPDATE public.system_menu
SET path      = '/workflow/home'::varchar(200),
    component = 'workflow/home/index'::varchar(200)
WHERE id = 1::bigint;

UPDATE public.system_menu
SET path      = '/workflow/work-order/list'::varchar(200),
    component = 'workflow/workOrder/list'::varchar(200)
WHERE id = 39::bigint;

UPDATE public.system_menu
SET path      = '/workflow/task/history'::varchar(200),
    component = 'workflow/task/history'::varchar(200)
WHERE id = 47::bigint;

UPDATE public.system_menu
SET path      = '/workflow/flow/template'::varchar(200),
    component = 'workflow/flow/templates'::varchar(200)
WHERE id = 42::bigint;

UPDATE public.system_menu
SET path      = '/workflow/flow/group'::varchar(200),
    component = 'workflow/flow/group'::varchar(200)
WHERE id = 41::bigint;

UPDATE public.system_menu
SET path      = '/workflow/work-order/create/:flowId'::varchar(200),
    component = 'workflow/workOrder/create'::varchar(200)
WHERE id = 49::bigint;

UPDATE public.system_menu
SET path      = '/workflow/flow/list'::varchar(200),
    component = 'workflow/flow/list'::varchar(200)
WHERE id = 43::bigint;

UPDATE public.system_menu
SET path      = '/workflow/work-order/apply'::varchar(200),
    component = 'workflow/workOrder/apply'::varchar(200)
WHERE id = 38::bigint;

UPDATE public.system_menu
SET path      = '/workflow/work-order/details/:id'::varchar(200),
    component = 'workflow/workOrder/details'::varchar(200)
WHERE id = 51::bigint;

UPDATE public.system_menu
SET path      = '/workflow/task/list'::varchar(200),
    component = 'workflow/task/list'::varchar(200)
WHERE id = 46::bigint;

UPDATE public.system_app
SET link = '/workflow/home'::varchar(256)
WHERE id = 1::bigint;

-- 结束同步`
