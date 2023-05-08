import React, { FC } from 'react';
import routes from '@/routes/index';
import NoMatch from '../pages/NoMatch';
import { Redirect, Route, Switch } from 'react-router-dom';
import { RouterType } from '@/routes/interface';
import Mens from '@/compoents/Mens';
import Breadcrumbs from '@/compoents/Breadcrumbs';
import Header from '@/compoents/Header';
import { useSelector } from 'react-redux';
import { userInfoType, StoreState } from '@/store/StoreState';
import { routesEnum } from '@/routes/config';
import Canvas from '../compoents/Canvas/index.js';
import Content from '../compoents/Content';
import History from '../compoents/History';
import Personal from '../compoents/Personal';
/**
/**
 * 获取所有路由页面
 * @returns
 */
const getRouters = (): Array<RouterType> => {
  let list = [];
  routes.map(item => {
    list = list.concat(item.children || []);
    return item;
  });
  return list;
};

const filterRoutes = (
  routes: Array<RouterType>,
  jurisdictions: string
): Array<RouterType> => {
  //因为key存储关键词重复的可能，所以转成数组
  const jurList: Array<string> = (jurisdictions || '').split(',');
  return routes.filter(route => {
    //如果包含当前路由的key则算具有对应权限
    return jurList.includes(route.key);
  });
};

const MainLayout: FC = () => {
  //获取所有路由页面
  const routes: Array<RouterType> = getRouters();
  //获取登录后设置的用户信息
  const userInfo: userInfoType = useSelector(
    (state: StoreState) => state.userInfo
  );
  //根据权限字符串过滤路由
  const routeList: Array<RouterType> = filterRoutes(
    routes,
    userInfo.jurisdictions
  );
  //默认路由
  const redirectRoute = routesEnum.defaultPath;
  return (
    <section>
      {/*
      <div className="main-wrapper">
        <div className="main-mens">
          <Mens jurisdictions={userInfo.jurisdictions}></Mens>
        </div>
        <div className="main-com">
          <Breadcrumbs></Breadcrumbs>
          <Route path="/no-match" key="no-match">
            <NoMatch></NoMatch>
          </Route>
          <Redirect path="/" to={redirectRoute} />
          {routeList.map(route => {
            return (
              <Route
                path={route.path}
                key={route.key}
                component={route.component}></Route>
            );
          })}
        </div>
      </div>
        */}

      <Header></Header>
      <Switch>
        <Route path="/" exact component={Content} />
        <Route path="/history" component={History} />
        <Route path="/personal" component={Personal} />
        <Route path="/no" component={NoMatch} />
      </Switch>
    </section>
  );
};

export default MainLayout;
