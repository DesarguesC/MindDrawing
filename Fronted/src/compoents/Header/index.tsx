import { Avatar, Dropdown, Menu } from 'antd';
import React, { FC } from 'react';
import './index.less';
import { UserOutlined } from '@ant-design/icons';
import { useHistory } from 'react-router-dom';
import { dispatchLogin } from '@/store/Actions';
import useLocalStorage from '@/hooks/useLocalStorage';
import { useDispatch, useSelector } from 'react-redux';
import { userInfoType, StoreState } from '@/store/StoreState';
import logo from '@/assets/images/logo.jpg';
import { HandPaintedPlate, History, Me } from '@icon-park/react';
import { NavLink } from 'react-router-dom';

const Header: FC = () => {
  const [, writeState] = useLocalStorage('token', '');
  const history = useHistory();
  const dispatch = useDispatch();
  const userInfo: userInfoType = useSelector(
    (state: StoreState) => state.userInfo
  );
  const onBack = () => {
    dispatch(dispatchLogin({ isLogin: false, userInfo: null }));
    writeState(null);
    history.push('/login');
  };
  const menu = (
    <Menu>
      <Menu.Item>
        <a onClick={onBack}>退出系统</a>
      </Menu.Item>
    </Menu>
  );

  return (
    <div className="header-warp flex">
      <div className="flex-1 in-flex-c">
        <img className="logo" src={logo} alt="logo" />
        <h1 className="h1">MindDrawing智绘</h1>
      </div>

      <div className="nav">
        <div>
          <NavLink exact to="/">
            <HandPaintedPlate theme="outline" size="24" fill="#6e41ff" />
            <span>创作空间</span>
          </NavLink>
        </div>
        <div className="history">
          <NavLink to="/history">
            <History theme="outline" size="24" fill="#6e41ff" />
            <span>历史画作</span>
          </NavLink>
        </div>
        {/* <div>
          <NavLink to="/gallery">
            <Me theme="outline" size="24" fill="#6e41ff" />
            <span>我的绘本</span>
          </NavLink>
        </div> */}
        <div>
          <NavLink to="/personal">
            <Me theme="outline" size="24" fill="#6e41ff" />
            <span>个人中心</span>
          </NavLink>
        </div>
      </div>
      <Dropdown overlay={menu}>
        <div className="pointer" style={{ marginLeft: '30px' }}>
          {/* <span className="m-r-10">{userInfo.userName}</span> */}
          <Avatar src={userInfo.pic} size={60} icon={<UserOutlined />} />
        </div>
      </Dropdown>
    </div>
  );
};
export default Header;
