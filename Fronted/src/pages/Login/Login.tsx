import {
  Layout,
  Form,
  Input,
  Button,
  Checkbox,
  notification,
  message
} from 'antd';
import React, { FC, useState } from 'react';
import './Login.less';
import { useDispatch } from 'react-redux';
import { useHistory } from 'react-router-dom';
import routes from '@/routes/index';
import { dispatchLogin } from '@/store/Actions';
import { userInfoType } from '@/store/StoreState';
import { filterRoute2Path } from '@/uilts/index';
import useLocalStorage from '@/hooks/useLocalStorage';
import UserApi from '@/apis/UserApi';
import { LoginType } from '@/Interface/common';

const Login: FC = () => {
  const [, writeState] = useLocalStorage('token', '');
  const [loading, setLoading] = useState(false);
  const dispatch = useDispatch();
  const history = useHistory();
  const onShowMsg = (msg: string) => {
    notification.error({
      message: '温馨提示',
      description: msg
    });
  };

  const onFinish = (values: LoginType) => {
    UserApi.login(values)
      .then((res: any) => {
        const userInfo: userInfoType = {
          userName: values.N_E,
          ...res.data
        };
        if (userInfo.jurisdictions) {
          //存储用户信息
          dispatch(dispatchLogin({ isLogin: true, userInfo }));
          //独立存储token
          writeState(userInfo.token);
          //根据账号权限获取默认第一个跳转页面
          const homePath = filterRoute2Path(routes, userInfo.jurisdictions);
          history.push('/');
          message.success('登录成功!');
        } else {
          onShowMsg('暂无权限!');
        }
      })
      .catch(() => {
        writeState('');
        onShowMsg('账号或密码错误!');
      })
      .finally(() => {
        setLoading(false);
      });
  };

  return (
    <>
      <Layout className="login-warp">
        <div className="login-container">
          <div className="login-left"></div>
          <div className="login-right">
            <Form
              className="login-form"
              name="basic"
              labelCol={{ span: 8 }}
              wrapperCol={{ span: 16 }}
              initialValues={{ userSetNumber: true }}
              onFinish={onFinish}
              autoComplete="off">
              <Form.Item
                label="用户名"
                name="N_E"
                rules={[{ required: true, message: '请输入用户名!' }]}>
                <Input placeholder="请输入用户名" size="large" />
              </Form.Item>

              <Form.Item
                label="密码"
                name="Pwd"
                rules={[{ required: true, message: '请输入密码!' }]}>
                <Input.Password placeholder="请输入密码" size="large" />
              </Form.Item>

              <Form.Item wrapperCol={{ offset: 8, span: 16 }}>
                <a
                  onClick={() => {
                    history.push('/register');
                  }}>
                  点击跳转到注册页
                </a>
              </Form.Item>

              <Form.Item wrapperCol={{ offset: 8, span: 16 }}>
                <Button
                  loading={loading}
                  size="large"
                  type="primary"
                  htmlType="submit"
                  style={{ width: '100%' }}>
                  登录
                </Button>
              </Form.Item>
            </Form>
          </div>
        </div>
      </Layout>
    </>
  );
};

export default Login;
