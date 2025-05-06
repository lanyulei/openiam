package mysql

/*
  @Author : lanyulei
  @Desc :
*/

const SQL202311210010 = `-- 开始同步

UPDATE system_api SET title = '获取部门树', ` + "`" + "group" + "`" + ` = 1 WHERE url = '/api/v1/system/department/tree' AND method = 'GET';

-- 结束同步`
