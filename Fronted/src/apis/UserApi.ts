/**
 * 用户模块接口
 */
import { QueryTableType, LoginType, RegisterType } from '@/Interface/common';
import { httpPost, ResponetFrom } from './request';

/**
 * 登录
 * @param params
 * @returns
 */
const login = (params: LoginType): Promise<ResponetFrom> => {
  return httpPost('/user/login', params);
};

/**
 * 注册
 * @param params
 * @returns
 */
const register = (params: RegisterType): Promise<ResponetFrom> => {
  return httpPost('/api/user/register', params);
};

/**
 * 查询用户列表
 * @param params
 * @returns
 */
const getUserList = (params: QueryTableType): Promise<ResponetFrom> => {
  return httpPost('/user/getUserList', params);
};

/**
 * 查询角色列表
 * @param params
 * @returns
 */
const getRoleList = (params: QueryTableType): Promise<ResponetFrom> => {
  return httpPost('/user/getRoleList', params);
};

export default {
  login,
  getUserList,
  getRoleList,
  register
};
