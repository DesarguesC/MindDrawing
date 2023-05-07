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
import './Register.less';
import { useDispatch } from 'react-redux';
import { useHistory } from 'react-router-dom';
import routes from '@/routes/index';
import { dispatchLogin } from '@/store/Actions';
import { userInfoType } from '@/store/StoreState';
import { filterRoute2Path } from '@/uilts/index';
import useLocalStorage from '@/hooks/useLocalStorage';
import UserApi from '@/apis/UserApi';
import { RegisterType } from '@/Interface/common';

const Register: FC = () => {
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

  const onFinish = (values: RegisterType) => {
    message.success('注册成功！请登录');
    history.push('/login');
    // UserApi.register(values).then((res: any) => {
    //   console.log(res.date);
    // });
  };

  return (
    <Layout className="register-warp">
      <div className="register-container">
        <div className="register-left"></div>
        <div className="register-right">
          <Form
            name="basic"
            labelCol={{ span: 8 }}
            wrapperCol={{ span: 16 }}
            initialValues={{ userSetNumber: true }}
            onFinish={onFinish}
            autoComplete="off">
            <Form.Item
              label="用户名"
              name="Name"
              rules={[{ required: true, message: '请输入用户名!' }]}>
              <Input placeholder="请输入用户名" size="large" />
            </Form.Item>

            <Form.Item
              label="邮箱"
              name="Email"
              rules={[{ required: true, message: '请输入邮箱!' }]}>
              <Input placeholder="请输入邮箱" size="large" />
            </Form.Item>

            <Form.Item
              label="密码"
              name="Pwd"
              rules={[{ required: true, message: '请输入密码!' }]}>
              <Input.Password placeholder="请输入密码" size="large" />
            </Form.Item>

            <Form.Item
              label="重复密码"
              name="Pwd_confirm"
              rules={[{ required: true, message: '请重复密码!' }]}>
              <Input.Password placeholder="请重复密码" size="large" />
            </Form.Item>

            <Form.Item
              label="密保问题"
              name="SecQ"
              rules={[{ required: true, message: '请输入邮箱密保问题!' }]}>
              <Input placeholder="请输入密保问题" size="large" />
            </Form.Item>

            <Form.Item
              label="密保答案"
              name="SecA"
              rules={[{ required: true, message: '请输入邮箱密保答案!' }]}>
              <Input placeholder="请输入密保答案" size="large" />
            </Form.Item>

            <Form.Item wrapperCol={{ offset: 8, span: 16 }}>
              <a
                onClick={() => {
                  history.push('/login');
                }}>
                点击跳转到登录页
              </a>
            </Form.Item>

            <Form.Item wrapperCol={{ offset: 8, span: 16 }}>
              <Button
                loading={loading}
                size="large"
                type="primary"
                htmlType="submit"
                style={{ width: '100%' }}>
                注册
              </Button>
            </Form.Item>
          </Form>
        </div>
      </div>
    </Layout>
  );
};

export default Register;
