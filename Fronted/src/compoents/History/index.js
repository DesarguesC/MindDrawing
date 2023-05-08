import React, { useState } from 'react';
import { Pagination, Col, Row, Card, Space, Image } from 'antd';
import './index.less';
import ironMan from '../../assets/images/ironman.png';
import ironMan2 from '../../assets/images/ironman2.png';
import spring from '../../assets/images/spring.png';
import summer from '../../assets/images/summer.jpg';
import autumn from '../../assets/images/autumn.png';
import winter from '../../assets/images/winter.jpg';

const src =
  'https://img1.baidu.com/it/u=413643897,2296924942&fm=253&fmt=auto&app=138&f=JPEG?w=800&h=500';

const imageData = [
  { id: 1, url: ironMan, context: '钢铁侠长出了翅膀，拿起了宝剑和盾牌' },
  { id: 2, url: ironMan2, context: '钢铁侠长出了翅膀，拿起了宝剑和盾牌' },
  {
    id: 3,
    url: spring,
    context:
      '处处春光明媚，处处春意盎然。小溪的冰融化了，溪水清澈见底，小鱼在溪水里欢快的玩耍。小草偷偷的从土钻出来，伸展嫩绿的身躯。各色的花争奇斗艳，散在草丛里，像无数的星星眨呀眨'
  },
  {
    id: 4,
    url: summer,
    context:
      '夏天是属于花的季节。各种各样的花儿这儿一朵，那儿一簇，大片大片的开出无限的美丽。红的、黄的、蓝的、粉的、白的..把夏天装扮的五彩缤纷'
  },
  {
    id: 5,
    url: autumn,
    context:
      '秋天撒下丰收的果实，天更蓝了，更纯洁，更明净了。太阳是那样明亮，亮得更加柔和。云是那样无暇，给人一种玲珑剔透的感觉，，心情也变得舒畅起来'
  },
  {
    id: 6,
    url: winter,
    context:
      '下雪的时候，一片片雪花从天上落下来，一会儿，山头白了，房子白了，窗外的一切都白了雪花落在我们的手心上化成了一滴水，洁白无瑕，晶莹剔透'
  }
];

function History() {
  const [currentPage, setCurrentPage] = useState(1);
  const pageSize = 2; // number of images to display per page

  const startIndex = (currentPage - 1) * pageSize;
  const endIndex = startIndex + pageSize;
  const currentImages = imageData.slice(startIndex, endIndex);

  return (
    <div className="container">
      <Row justify="space-around" align="middle" className="wrapper">
        {currentImages.map(image => (
          <Col span={11} key={image.id}>
            <Image className="pic" key={image.id} src={image.url} />
            <div className="text">{image.context}</div>
          </Col>
        ))}
      </Row>

      <div
        style={{
          display: 'flex',
          justifyContent: 'center',
          position: 'absolute',
          bottom: 70,
          left: 0,
          right: 0
        }}>
        <Pagination
          // className="page"
          current={currentPage}
          pageSize={pageSize}
          total={imageData.length}
          onChange={page => setCurrentPage(page)}
        />
      </div>
    </div>
  );
}

export default History;
