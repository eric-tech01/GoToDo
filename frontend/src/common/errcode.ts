export enum ErrorCode {
  TOKEN_INVALID = 'token.is.invalid',
  USER_LOGIN_FAILED = 'user.login.failed',

  USER_PWD_VERIFY_FAILED = 'user.pwd.verify.failed',

  SYS_USER_NOT_INIT = 'sys.user.not.init',

  USER_LOGIN_EXPIRE = 'sys.user.session.expired',
  USER_LOGIN_SESSION_INVALID = 'sys.user.session.invalid',
  USER_NOT_EXIST = 'sys.user.not.exist',
  SYS_EXCEPTION_PARAM_ERR = 'system.exception.error.param',
  SYS_EXCEPTION = 'system.exception',
  INVENTORY_NOT_IN_WAREHOUSE = 'inventory.not.in.warehouse',
}

export const errorMap: Map<ErrorCode, string> = new Map([
  [ErrorCode.TOKEN_INVALID, '登录超时，请重新登录'],
  [ErrorCode.USER_LOGIN_FAILED, '用户名或密码错误'],
  [ErrorCode.USER_PWD_VERIFY_FAILED, '旧密码验证失败'],

  [ErrorCode.USER_NOT_EXIST, '用户账号或密码错误'],
  [ErrorCode.SYS_EXCEPTION_PARAM_ERR, '输入错误，请重新输入'],
  [ErrorCode.SYS_EXCEPTION, '系统异常，请联系管理员或稍后操作'],
  [ErrorCode.USER_LOGIN_EXPIRE, '登录过期，请重新登录'],
  [ErrorCode.INVENTORY_NOT_IN_WAREHOUSE, '物料不在仓库，请走入库登记'],
]);
